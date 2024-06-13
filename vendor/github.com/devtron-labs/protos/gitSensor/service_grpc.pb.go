// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.4.0
// - protoc             v3.9.1
// source: gitSensor/service.proto

package gitSensor

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.62.0 or later.
const _ = grpc.SupportPackageIsVersion8

const (
	GitSensorService_SaveGitProvider_FullMethodName                                  = "/gitService.GitSensorService/SaveGitProvider"
	GitSensorService_AddRepo_FullMethodName                                          = "/gitService.GitSensorService/AddRepo"
	GitSensorService_UpdateRepo_FullMethodName                                       = "/gitService.GitSensorService/UpdateRepo"
	GitSensorService_SavePipelineMaterial_FullMethodName                             = "/gitService.GitSensorService/SavePipelineMaterial"
	GitSensorService_FetchChanges_FullMethodName                                     = "/gitService.GitSensorService/FetchChanges"
	GitSensorService_GetHeadForPipelineMaterials_FullMethodName                      = "/gitService.GitSensorService/GetHeadForPipelineMaterials"
	GitSensorService_GetCommitMetadata_FullMethodName                                = "/gitService.GitSensorService/GetCommitMetadata"
	GitSensorService_GetCommitMetadataForPipelineMaterial_FullMethodName             = "/gitService.GitSensorService/GetCommitMetadataForPipelineMaterial"
	GitSensorService_GetCommitInfoForTag_FullMethodName                              = "/gitService.GitSensorService/GetCommitInfoForTag"
	GitSensorService_RefreshGitMaterial_FullMethodName                               = "/gitService.GitSensorService/RefreshGitMaterial"
	GitSensorService_ReloadAllMaterial_FullMethodName                                = "/gitService.GitSensorService/ReloadAllMaterial"
	GitSensorService_ReloadMaterial_FullMethodName                                   = "/gitService.GitSensorService/ReloadMaterial"
	GitSensorService_GetChangesInRelease_FullMethodName                              = "/gitService.GitSensorService/GetChangesInRelease"
	GitSensorService_GetWebhookData_FullMethodName                                   = "/gitService.GitSensorService/GetWebhookData"
	GitSensorService_GetAllWebhookEventConfigForHost_FullMethodName                  = "/gitService.GitSensorService/GetAllWebhookEventConfigForHost"
	GitSensorService_GetWebhookEventConfig_FullMethodName                            = "/gitService.GitSensorService/GetWebhookEventConfig"
	GitSensorService_GetWebhookPayloadDataForPipelineMaterialId_FullMethodName       = "/gitService.GitSensorService/GetWebhookPayloadDataForPipelineMaterialId"
	GitSensorService_GetWebhookPayloadFilterDataForPipelineMaterialId_FullMethodName = "/gitService.GitSensorService/GetWebhookPayloadFilterDataForPipelineMaterialId"
)

// GitSensorServiceClient is the client API for GitSensorService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type GitSensorServiceClient interface {
	// Saves Git credentials
	SaveGitProvider(ctx context.Context, in *GitProvider, opts ...grpc.CallOption) (*Empty, error)
	// Add Repo
	AddRepo(ctx context.Context, in *AddRepoRequest, opts ...grpc.CallOption) (*Empty, error)
	// Update Repo
	UpdateRepo(ctx context.Context, in *GitMaterial, opts ...grpc.CallOption) (*Empty, error)
	// Save CI pipeline material
	SavePipelineMaterial(ctx context.Context, in *SavePipelineMaterialRequest, opts ...grpc.CallOption) (*Empty, error)
	// Fetch SCM changes
	FetchChanges(ctx context.Context, in *FetchScmChangesRequest, opts ...grpc.CallOption) (*MaterialChangeResponse, error)
	// Get Head for pipeline materials
	GetHeadForPipelineMaterials(ctx context.Context, in *HeadRequest, opts ...grpc.CallOption) (*GetHeadForPipelineMaterialsResponse, error)
	// Get commit metadata
	GetCommitMetadata(ctx context.Context, in *CommitMetadataRequest, opts ...grpc.CallOption) (*GitCommit, error)
	// Get commit metadata for pipeline material
	GetCommitMetadataForPipelineMaterial(ctx context.Context, in *CommitMetadataRequest, opts ...grpc.CallOption) (*GitCommit, error)
	// Get commit info for Tag
	GetCommitInfoForTag(ctx context.Context, in *CommitMetadataRequest, opts ...grpc.CallOption) (*GitCommit, error)
	// Refresh git material
	RefreshGitMaterial(ctx context.Context, in *RefreshGitMaterialRequest, opts ...grpc.CallOption) (*RefreshGitMaterialResponse, error)
	// Reload all material
	ReloadAllMaterial(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error)
	// Reload a specific material
	ReloadMaterial(ctx context.Context, in *ReloadMaterialRequest, opts ...grpc.CallOption) (*GenericResponse, error)
	// Get changes in release
	GetChangesInRelease(ctx context.Context, in *ReleaseChangeRequest, opts ...grpc.CallOption) (*GitChanges, error)
	// Get webhook data
	GetWebhookData(ctx context.Context, in *WebhookDataRequest, opts ...grpc.CallOption) (*WebhookAndCiData, error)
	// Get all webhook event config for host
	GetAllWebhookEventConfigForHost(ctx context.Context, in *WebhookEventConfigRequest, opts ...grpc.CallOption) (*WebhookEventConfigResponse, error)
	// Get webhook event config
	GetWebhookEventConfig(ctx context.Context, in *WebhookEventConfigRequest, opts ...grpc.CallOption) (*WebhookEventConfig, error)
	// Get webhook payload data by pipeline material id
	GetWebhookPayloadDataForPipelineMaterialId(ctx context.Context, in *WebhookPayloadDataRequest, opts ...grpc.CallOption) (*WebhookPayloadDataResponse, error)
	// Get webhook payload data by pipeline material id with filter
	GetWebhookPayloadFilterDataForPipelineMaterialId(ctx context.Context, in *WebhookPayloadFilterDataRequest, opts ...grpc.CallOption) (*WebhookPayloadFilterDataResponse, error)
}

type gitSensorServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewGitSensorServiceClient(cc grpc.ClientConnInterface) GitSensorServiceClient {
	return &gitSensorServiceClient{cc}
}

func (c *gitSensorServiceClient) SaveGitProvider(ctx context.Context, in *GitProvider, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, GitSensorService_SaveGitProvider_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitSensorServiceClient) AddRepo(ctx context.Context, in *AddRepoRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, GitSensorService_AddRepo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitSensorServiceClient) UpdateRepo(ctx context.Context, in *GitMaterial, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, GitSensorService_UpdateRepo_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitSensorServiceClient) SavePipelineMaterial(ctx context.Context, in *SavePipelineMaterialRequest, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, GitSensorService_SavePipelineMaterial_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitSensorServiceClient) FetchChanges(ctx context.Context, in *FetchScmChangesRequest, opts ...grpc.CallOption) (*MaterialChangeResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(MaterialChangeResponse)
	err := c.cc.Invoke(ctx, GitSensorService_FetchChanges_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitSensorServiceClient) GetHeadForPipelineMaterials(ctx context.Context, in *HeadRequest, opts ...grpc.CallOption) (*GetHeadForPipelineMaterialsResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GetHeadForPipelineMaterialsResponse)
	err := c.cc.Invoke(ctx, GitSensorService_GetHeadForPipelineMaterials_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitSensorServiceClient) GetCommitMetadata(ctx context.Context, in *CommitMetadataRequest, opts ...grpc.CallOption) (*GitCommit, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GitCommit)
	err := c.cc.Invoke(ctx, GitSensorService_GetCommitMetadata_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitSensorServiceClient) GetCommitMetadataForPipelineMaterial(ctx context.Context, in *CommitMetadataRequest, opts ...grpc.CallOption) (*GitCommit, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GitCommit)
	err := c.cc.Invoke(ctx, GitSensorService_GetCommitMetadataForPipelineMaterial_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitSensorServiceClient) GetCommitInfoForTag(ctx context.Context, in *CommitMetadataRequest, opts ...grpc.CallOption) (*GitCommit, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GitCommit)
	err := c.cc.Invoke(ctx, GitSensorService_GetCommitInfoForTag_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitSensorServiceClient) RefreshGitMaterial(ctx context.Context, in *RefreshGitMaterialRequest, opts ...grpc.CallOption) (*RefreshGitMaterialResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RefreshGitMaterialResponse)
	err := c.cc.Invoke(ctx, GitSensorService_RefreshGitMaterial_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitSensorServiceClient) ReloadAllMaterial(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Empty)
	err := c.cc.Invoke(ctx, GitSensorService_ReloadAllMaterial_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitSensorServiceClient) ReloadMaterial(ctx context.Context, in *ReloadMaterialRequest, opts ...grpc.CallOption) (*GenericResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GenericResponse)
	err := c.cc.Invoke(ctx, GitSensorService_ReloadMaterial_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitSensorServiceClient) GetChangesInRelease(ctx context.Context, in *ReleaseChangeRequest, opts ...grpc.CallOption) (*GitChanges, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(GitChanges)
	err := c.cc.Invoke(ctx, GitSensorService_GetChangesInRelease_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitSensorServiceClient) GetWebhookData(ctx context.Context, in *WebhookDataRequest, opts ...grpc.CallOption) (*WebhookAndCiData, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WebhookAndCiData)
	err := c.cc.Invoke(ctx, GitSensorService_GetWebhookData_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitSensorServiceClient) GetAllWebhookEventConfigForHost(ctx context.Context, in *WebhookEventConfigRequest, opts ...grpc.CallOption) (*WebhookEventConfigResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WebhookEventConfigResponse)
	err := c.cc.Invoke(ctx, GitSensorService_GetAllWebhookEventConfigForHost_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitSensorServiceClient) GetWebhookEventConfig(ctx context.Context, in *WebhookEventConfigRequest, opts ...grpc.CallOption) (*WebhookEventConfig, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WebhookEventConfig)
	err := c.cc.Invoke(ctx, GitSensorService_GetWebhookEventConfig_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitSensorServiceClient) GetWebhookPayloadDataForPipelineMaterialId(ctx context.Context, in *WebhookPayloadDataRequest, opts ...grpc.CallOption) (*WebhookPayloadDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WebhookPayloadDataResponse)
	err := c.cc.Invoke(ctx, GitSensorService_GetWebhookPayloadDataForPipelineMaterialId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *gitSensorServiceClient) GetWebhookPayloadFilterDataForPipelineMaterialId(ctx context.Context, in *WebhookPayloadFilterDataRequest, opts ...grpc.CallOption) (*WebhookPayloadFilterDataResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(WebhookPayloadFilterDataResponse)
	err := c.cc.Invoke(ctx, GitSensorService_GetWebhookPayloadFilterDataForPipelineMaterialId_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// GitSensorServiceServer is the server API for GitSensorService service.
// All implementations must embed UnimplementedGitSensorServiceServer
// for forward compatibility
type GitSensorServiceServer interface {
	// Saves Git credentials
	SaveGitProvider(context.Context, *GitProvider) (*Empty, error)
	// Add Repo
	AddRepo(context.Context, *AddRepoRequest) (*Empty, error)
	// Update Repo
	UpdateRepo(context.Context, *GitMaterial) (*Empty, error)
	// Save CI pipeline material
	SavePipelineMaterial(context.Context, *SavePipelineMaterialRequest) (*Empty, error)
	// Fetch SCM changes
	FetchChanges(context.Context, *FetchScmChangesRequest) (*MaterialChangeResponse, error)
	// Get Head for pipeline materials
	GetHeadForPipelineMaterials(context.Context, *HeadRequest) (*GetHeadForPipelineMaterialsResponse, error)
	// Get commit metadata
	GetCommitMetadata(context.Context, *CommitMetadataRequest) (*GitCommit, error)
	// Get commit metadata for pipeline material
	GetCommitMetadataForPipelineMaterial(context.Context, *CommitMetadataRequest) (*GitCommit, error)
	// Get commit info for Tag
	GetCommitInfoForTag(context.Context, *CommitMetadataRequest) (*GitCommit, error)
	// Refresh git material
	RefreshGitMaterial(context.Context, *RefreshGitMaterialRequest) (*RefreshGitMaterialResponse, error)
	// Reload all material
	ReloadAllMaterial(context.Context, *Empty) (*Empty, error)
	// Reload a specific material
	ReloadMaterial(context.Context, *ReloadMaterialRequest) (*GenericResponse, error)
	// Get changes in release
	GetChangesInRelease(context.Context, *ReleaseChangeRequest) (*GitChanges, error)
	// Get webhook data
	GetWebhookData(context.Context, *WebhookDataRequest) (*WebhookAndCiData, error)
	// Get all webhook event config for host
	GetAllWebhookEventConfigForHost(context.Context, *WebhookEventConfigRequest) (*WebhookEventConfigResponse, error)
	// Get webhook event config
	GetWebhookEventConfig(context.Context, *WebhookEventConfigRequest) (*WebhookEventConfig, error)
	// Get webhook payload data by pipeline material id
	GetWebhookPayloadDataForPipelineMaterialId(context.Context, *WebhookPayloadDataRequest) (*WebhookPayloadDataResponse, error)
	// Get webhook payload data by pipeline material id with filter
	GetWebhookPayloadFilterDataForPipelineMaterialId(context.Context, *WebhookPayloadFilterDataRequest) (*WebhookPayloadFilterDataResponse, error)
	mustEmbedUnimplementedGitSensorServiceServer()
}

// UnimplementedGitSensorServiceServer must be embedded to have forward compatible implementations.
type UnimplementedGitSensorServiceServer struct {
}

func (UnimplementedGitSensorServiceServer) SaveGitProvider(context.Context, *GitProvider) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SaveGitProvider not implemented")
}
func (UnimplementedGitSensorServiceServer) AddRepo(context.Context, *AddRepoRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddRepo not implemented")
}
func (UnimplementedGitSensorServiceServer) UpdateRepo(context.Context, *GitMaterial) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateRepo not implemented")
}
func (UnimplementedGitSensorServiceServer) SavePipelineMaterial(context.Context, *SavePipelineMaterialRequest) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SavePipelineMaterial not implemented")
}
func (UnimplementedGitSensorServiceServer) FetchChanges(context.Context, *FetchScmChangesRequest) (*MaterialChangeResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchChanges not implemented")
}
func (UnimplementedGitSensorServiceServer) GetHeadForPipelineMaterials(context.Context, *HeadRequest) (*GetHeadForPipelineMaterialsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetHeadForPipelineMaterials not implemented")
}
func (UnimplementedGitSensorServiceServer) GetCommitMetadata(context.Context, *CommitMetadataRequest) (*GitCommit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommitMetadata not implemented")
}
func (UnimplementedGitSensorServiceServer) GetCommitMetadataForPipelineMaterial(context.Context, *CommitMetadataRequest) (*GitCommit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommitMetadataForPipelineMaterial not implemented")
}
func (UnimplementedGitSensorServiceServer) GetCommitInfoForTag(context.Context, *CommitMetadataRequest) (*GitCommit, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetCommitInfoForTag not implemented")
}
func (UnimplementedGitSensorServiceServer) RefreshGitMaterial(context.Context, *RefreshGitMaterialRequest) (*RefreshGitMaterialResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RefreshGitMaterial not implemented")
}
func (UnimplementedGitSensorServiceServer) ReloadAllMaterial(context.Context, *Empty) (*Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReloadAllMaterial not implemented")
}
func (UnimplementedGitSensorServiceServer) ReloadMaterial(context.Context, *ReloadMaterialRequest) (*GenericResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ReloadMaterial not implemented")
}
func (UnimplementedGitSensorServiceServer) GetChangesInRelease(context.Context, *ReleaseChangeRequest) (*GitChanges, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetChangesInRelease not implemented")
}
func (UnimplementedGitSensorServiceServer) GetWebhookData(context.Context, *WebhookDataRequest) (*WebhookAndCiData, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWebhookData not implemented")
}
func (UnimplementedGitSensorServiceServer) GetAllWebhookEventConfigForHost(context.Context, *WebhookEventConfigRequest) (*WebhookEventConfigResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetAllWebhookEventConfigForHost not implemented")
}
func (UnimplementedGitSensorServiceServer) GetWebhookEventConfig(context.Context, *WebhookEventConfigRequest) (*WebhookEventConfig, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWebhookEventConfig not implemented")
}
func (UnimplementedGitSensorServiceServer) GetWebhookPayloadDataForPipelineMaterialId(context.Context, *WebhookPayloadDataRequest) (*WebhookPayloadDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWebhookPayloadDataForPipelineMaterialId not implemented")
}
func (UnimplementedGitSensorServiceServer) GetWebhookPayloadFilterDataForPipelineMaterialId(context.Context, *WebhookPayloadFilterDataRequest) (*WebhookPayloadFilterDataResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetWebhookPayloadFilterDataForPipelineMaterialId not implemented")
}
func (UnimplementedGitSensorServiceServer) mustEmbedUnimplementedGitSensorServiceServer() {}

// UnsafeGitSensorServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to GitSensorServiceServer will
// result in compilation errors.
type UnsafeGitSensorServiceServer interface {
	mustEmbedUnimplementedGitSensorServiceServer()
}

func RegisterGitSensorServiceServer(s grpc.ServiceRegistrar, srv GitSensorServiceServer) {
	s.RegisterService(&GitSensorService_ServiceDesc, srv)
}

func _GitSensorService_SaveGitProvider_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GitProvider)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitSensorServiceServer).SaveGitProvider(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitSensorService_SaveGitProvider_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitSensorServiceServer).SaveGitProvider(ctx, req.(*GitProvider))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitSensorService_AddRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AddRepoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitSensorServiceServer).AddRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitSensorService_AddRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitSensorServiceServer).AddRepo(ctx, req.(*AddRepoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitSensorService_UpdateRepo_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GitMaterial)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitSensorServiceServer).UpdateRepo(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitSensorService_UpdateRepo_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitSensorServiceServer).UpdateRepo(ctx, req.(*GitMaterial))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitSensorService_SavePipelineMaterial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SavePipelineMaterialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitSensorServiceServer).SavePipelineMaterial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitSensorService_SavePipelineMaterial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitSensorServiceServer).SavePipelineMaterial(ctx, req.(*SavePipelineMaterialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitSensorService_FetchChanges_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchScmChangesRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitSensorServiceServer).FetchChanges(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitSensorService_FetchChanges_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitSensorServiceServer).FetchChanges(ctx, req.(*FetchScmChangesRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitSensorService_GetHeadForPipelineMaterials_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(HeadRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitSensorServiceServer).GetHeadForPipelineMaterials(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitSensorService_GetHeadForPipelineMaterials_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitSensorServiceServer).GetHeadForPipelineMaterials(ctx, req.(*HeadRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitSensorService_GetCommitMetadata_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitSensorServiceServer).GetCommitMetadata(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitSensorService_GetCommitMetadata_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitSensorServiceServer).GetCommitMetadata(ctx, req.(*CommitMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitSensorService_GetCommitMetadataForPipelineMaterial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitSensorServiceServer).GetCommitMetadataForPipelineMaterial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitSensorService_GetCommitMetadataForPipelineMaterial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitSensorServiceServer).GetCommitMetadataForPipelineMaterial(ctx, req.(*CommitMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitSensorService_GetCommitInfoForTag_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitMetadataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitSensorServiceServer).GetCommitInfoForTag(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitSensorService_GetCommitInfoForTag_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitSensorServiceServer).GetCommitInfoForTag(ctx, req.(*CommitMetadataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitSensorService_RefreshGitMaterial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RefreshGitMaterialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitSensorServiceServer).RefreshGitMaterial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitSensorService_RefreshGitMaterial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitSensorServiceServer).RefreshGitMaterial(ctx, req.(*RefreshGitMaterialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitSensorService_ReloadAllMaterial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitSensorServiceServer).ReloadAllMaterial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitSensorService_ReloadAllMaterial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitSensorServiceServer).ReloadAllMaterial(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitSensorService_ReloadMaterial_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReloadMaterialRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitSensorServiceServer).ReloadMaterial(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitSensorService_ReloadMaterial_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitSensorServiceServer).ReloadMaterial(ctx, req.(*ReloadMaterialRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitSensorService_GetChangesInRelease_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReleaseChangeRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitSensorServiceServer).GetChangesInRelease(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitSensorService_GetChangesInRelease_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitSensorServiceServer).GetChangesInRelease(ctx, req.(*ReleaseChangeRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitSensorService_GetWebhookData_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitSensorServiceServer).GetWebhookData(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitSensorService_GetWebhookData_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitSensorServiceServer).GetWebhookData(ctx, req.(*WebhookDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitSensorService_GetAllWebhookEventConfigForHost_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookEventConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitSensorServiceServer).GetAllWebhookEventConfigForHost(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitSensorService_GetAllWebhookEventConfigForHost_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitSensorServiceServer).GetAllWebhookEventConfigForHost(ctx, req.(*WebhookEventConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitSensorService_GetWebhookEventConfig_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookEventConfigRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitSensorServiceServer).GetWebhookEventConfig(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitSensorService_GetWebhookEventConfig_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitSensorServiceServer).GetWebhookEventConfig(ctx, req.(*WebhookEventConfigRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitSensorService_GetWebhookPayloadDataForPipelineMaterialId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookPayloadDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitSensorServiceServer).GetWebhookPayloadDataForPipelineMaterialId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitSensorService_GetWebhookPayloadDataForPipelineMaterialId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitSensorServiceServer).GetWebhookPayloadDataForPipelineMaterialId(ctx, req.(*WebhookPayloadDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _GitSensorService_GetWebhookPayloadFilterDataForPipelineMaterialId_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(WebhookPayloadFilterDataRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(GitSensorServiceServer).GetWebhookPayloadFilterDataForPipelineMaterialId(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: GitSensorService_GetWebhookPayloadFilterDataForPipelineMaterialId_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(GitSensorServiceServer).GetWebhookPayloadFilterDataForPipelineMaterialId(ctx, req.(*WebhookPayloadFilterDataRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// GitSensorService_ServiceDesc is the grpc.ServiceDesc for GitSensorService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var GitSensorService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "gitService.GitSensorService",
	HandlerType: (*GitSensorServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "SaveGitProvider",
			Handler:    _GitSensorService_SaveGitProvider_Handler,
		},
		{
			MethodName: "AddRepo",
			Handler:    _GitSensorService_AddRepo_Handler,
		},
		{
			MethodName: "UpdateRepo",
			Handler:    _GitSensorService_UpdateRepo_Handler,
		},
		{
			MethodName: "SavePipelineMaterial",
			Handler:    _GitSensorService_SavePipelineMaterial_Handler,
		},
		{
			MethodName: "FetchChanges",
			Handler:    _GitSensorService_FetchChanges_Handler,
		},
		{
			MethodName: "GetHeadForPipelineMaterials",
			Handler:    _GitSensorService_GetHeadForPipelineMaterials_Handler,
		},
		{
			MethodName: "GetCommitMetadata",
			Handler:    _GitSensorService_GetCommitMetadata_Handler,
		},
		{
			MethodName: "GetCommitMetadataForPipelineMaterial",
			Handler:    _GitSensorService_GetCommitMetadataForPipelineMaterial_Handler,
		},
		{
			MethodName: "GetCommitInfoForTag",
			Handler:    _GitSensorService_GetCommitInfoForTag_Handler,
		},
		{
			MethodName: "RefreshGitMaterial",
			Handler:    _GitSensorService_RefreshGitMaterial_Handler,
		},
		{
			MethodName: "ReloadAllMaterial",
			Handler:    _GitSensorService_ReloadAllMaterial_Handler,
		},
		{
			MethodName: "ReloadMaterial",
			Handler:    _GitSensorService_ReloadMaterial_Handler,
		},
		{
			MethodName: "GetChangesInRelease",
			Handler:    _GitSensorService_GetChangesInRelease_Handler,
		},
		{
			MethodName: "GetWebhookData",
			Handler:    _GitSensorService_GetWebhookData_Handler,
		},
		{
			MethodName: "GetAllWebhookEventConfigForHost",
			Handler:    _GitSensorService_GetAllWebhookEventConfigForHost_Handler,
		},
		{
			MethodName: "GetWebhookEventConfig",
			Handler:    _GitSensorService_GetWebhookEventConfig_Handler,
		},
		{
			MethodName: "GetWebhookPayloadDataForPipelineMaterialId",
			Handler:    _GitSensorService_GetWebhookPayloadDataForPipelineMaterialId_Handler,
		},
		{
			MethodName: "GetWebhookPayloadFilterDataForPipelineMaterialId",
			Handler:    _GitSensorService_GetWebhookPayloadFilterDataForPipelineMaterialId_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "gitSensor/service.proto",
}
