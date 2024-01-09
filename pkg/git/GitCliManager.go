package git

import (
	"encoding/json"
	"fmt"
	"go.uber.org/zap"
	"gopkg.in/src-d/go-billy.v4/osfs"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
)

type GitCliManager interface {
	GitManager
}

type GitCliManagerImpl struct {
	GitManagerBaseImpl
}

func NewGitCliManagerImpl(logger *zap.SugaredLogger) *GitCliManagerImpl {
	return &GitCliManagerImpl{
		GitManagerBaseImpl: GitManagerBaseImpl{logger: logger},
	}
}

const (
	GIT_ASK_PASS                = "/git-ask-pass.sh"
	AUTHENTICATION_FAILED_ERROR = "Authentication failed"
)

func (impl *GitCliManagerImpl) Init(rootDir string, remoteUrl string, isBare bool) error {
	//-----------------

	err := os.MkdirAll(rootDir, 0755)
	if err != nil {
		return err
	}

	err = impl.GitInit(rootDir)
	if err != nil {
		return err
	}
	return impl.GitCreateRemote(rootDir, remoteUrl)

}

func (impl *GitCliManagerImpl) OpenRepoPlain(checkoutPath string) (*GitRepository, error) {

	err := openGitRepo(checkoutPath)
	if err != nil {
		return nil, err
	}
	return &GitRepository{
		rootDir: checkoutPath,
	}, nil
}

func (impl *GitCliManagerImpl) GetCommitsForTag(checkoutPath, tag string) (GitCommit, error) {
	return impl.GitShow(checkoutPath, tag)
}

func (impl *GitCliManagerImpl) GetCommitForHash(checkoutPath, commitHash string) (GitCommit, error) {

	return impl.GitShow(checkoutPath, commitHash)
}
func (impl *GitCliManagerImpl) GetCommitIterator(repository *GitRepository, iteratorRequest IteratorRequest) (CommitIterator, error) {

	commits, err := impl.GetCommits(iteratorRequest.BranchRef, iteratorRequest.Branch, repository.rootDir, iteratorRequest.CommitCount, iteratorRequest.FromCommitHash, iteratorRequest.ToCommitHash)
	if err != nil {
		impl.logger.Errorw("error in fetching commits for", "err", err, "path", repository.rootDir)
		return nil, err
	}
	return &CommitCliIterator{
		commits: commits,
	}, nil
}

func openGitRepo(path string) error {
	if _, err := filepath.Abs(path); err != nil {
		return err
	}
	fst := osfs.New(path)
	_, err := fst.Stat(".git")
	if !os.IsNotExist(err) {
		return err
	}
	return nil
}
func (impl *GitCliManagerImpl) GitInit(rootDir string) error {
	impl.logger.Debugw("git", "-C", rootDir, "init")
	cmd := exec.Command("git", "-C", rootDir, "init")
	output, errMsg, err := impl.runCommand(cmd)
	impl.logger.Debugw("root", rootDir, "opt", output, "errMsg", errMsg, "error", err)
	return err
}

func (impl *GitCliManagerImpl) GitCreateRemote(rootDir string, url string) error {
	impl.logger.Debugw("git", "-C", rootDir, "remote", "add", "origin", url)
	cmd := exec.Command("git", "-C", rootDir, "remote", "add", "origin", url)
	output, errMsg, err := impl.runCommand(cmd)
	impl.logger.Debugw("url", url, "opt", output, "errMsg", errMsg, "error", err)
	return err
}

func (impl *GitCliManagerImpl) GetCommits(branchRef string, branch string, rootDir string, numCommits int, from string, to string) ([]GitCommit, error) {
	baseCmdArgs := []string{"-C", rootDir, "log"}
	rangeCmdArgs := []string{branchRef}
	extraCmdArgs := []string{"-n", strconv.Itoa(numCommits), "--date=iso-strict", GITFORMAT}
	cmdArgs := impl.getCommandForLogRange(branchRef, from, to, rangeCmdArgs, baseCmdArgs, extraCmdArgs)

	impl.logger.Debugw("git", cmdArgs)
	cmd := exec.Command("git", cmdArgs...)
	output, errMsg, err := impl.runCommand(cmd)
	impl.logger.Debugw("root", rootDir, "opt", output, "errMsg", errMsg, "error", err)
	if err != nil {
		return nil, err
	}
	commits, err := impl.processGitLogOutput(output, rootDir)
	if err != nil {
		return nil, err
	}
	return commits, nil
}

func (impl *GitCliManagerImpl) getCommandForLogRange(branchRef string, from string, to string, rangeCmdArgs []string, baseCmdArgs []string, extraCmdArgs []string) []string {
	if from != "" && to != "" {
		rangeCmdArgs = []string{from + "^.." + to}
	} else if from != "" {
		rangeCmdArgs = []string{from + "^.." + branchRef}
	} else if to != "" {
		rangeCmdArgs = []string{to}
	}
	return append(baseCmdArgs, append(rangeCmdArgs, extraCmdArgs...)...)
}

func (impl *GitCliManagerImpl) GitShow(rootDir string, hash string) (GitCommit, error) {
	impl.logger.Debugw("git", "-C", rootDir, "show", hash, "--date=iso-strict", GITFORMAT, "-s")
	cmd := exec.Command("git", "-C", rootDir, "show", hash, "--date=iso-strict", GITFORMAT, "-s")
	output, errMsg, err := impl.runCommand(cmd)
	impl.logger.Debugw("root", rootDir, "opt", output, "errMsg", errMsg, "error", err)
	commits, err := impl.processGitLogOutput(output, rootDir)
	if err != nil || len(commits) == 0 {
		return nil, err
	}

	return commits[0], nil
}

func (impl *GitCliManagerImpl) processGitLogOutput(out string, rootDir string) ([]GitCommit, error) {
	if len(out) == 0 {
		return make([]GitCommit, 0), nil
	}
	logOut := out
	logOut = logOut[:len(logOut)-1]      // Remove the last ","
	logOut = fmt.Sprintf("[%s]", logOut) // Add []

	var gitCommitFormattedList []GitCommitFormat
	err := json.Unmarshal([]byte(logOut), &gitCommitFormattedList)
	if err != nil {
		return nil, err
	}

	gitCommits := make([]GitCommit, 0)
	for _, formattedCommit := range gitCommitFormattedList {

		cm := GitCommitBase{
			Commit:       formattedCommit.Commit,
			Author:       formattedCommit.Commiter.Name + " <" + formattedCommit.Commiter.Email + ">",
			Date:         formattedCommit.Commiter.Date,
			Message:      formattedCommit.Subject + "\n" + formattedCommit.Body,
			CheckoutPath: rootDir,
		}
		gitCommits = append(gitCommits, &GitCommitCli{
			GitCommitBase: cm,
		})
	}
	return gitCommits, nil
}

func (impl *GitCliManagerImpl) FetchDiffStatBetweenCommits(gitContext *GitContext, oldHash string, newHash string, rootDir string) (response, errMsg string, err error) {
	impl.logger.Debugw("git", "-C", rootDir, "diff", "--numstat", oldHash, newHash)

	if newHash == "" {
		newHash = oldHash
		oldHash = oldHash + "^"
	}
	cmd := exec.Command("git", "-C", rootDir, "diff", "--numstat", oldHash, newHash)

	output, errMsg, err := impl.runCommandWithCred(cmd, gitContext.Username, gitContext.Password)
	impl.logger.Debugw("root", rootDir, "opt", output, "errMsg", errMsg, "error", err)
	return output, errMsg, err
}

func (impl *GitCliManagerImpl) GetCommitStats(commit GitCommit) (FileStats, error) {
	gitCommit := commit.GetCommit()
	fileStat, errorMsg, err := impl.FetchDiffStatBetweenCommits(&GitContext{}, gitCommit.Commit, "", gitCommit.CheckoutPath)
	if err != nil {
		impl.logger.Errorw("error in fetching fileStat of commit: ", gitCommit.Commit, "checkoutPath", gitCommit.CheckoutPath, "errorMsg", errorMsg, "err", err)
		return nil, err
	}
	return getFileStat(fileStat)
}