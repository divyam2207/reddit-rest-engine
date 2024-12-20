// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.3
// source: proto/engine.proto

package proto

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	RedditEngine_RegisterUser_FullMethodName        = "/reddit.RedditEngine/RegisterUser"
	RedditEngine_CreateSubreddit_FullMethodName     = "/reddit.RedditEngine/CreateSubreddit"
	RedditEngine_JoinSubreddit_FullMethodName       = "/reddit.RedditEngine/JoinSubreddit"
	RedditEngine_Post_FullMethodName                = "/reddit.RedditEngine/Post"
	RedditEngine_Comment_FullMethodName             = "/reddit.RedditEngine/Comment"
	RedditEngine_Vote_FullMethodName                = "/reddit.RedditEngine/Vote"
	RedditEngine_FetchSubredditState_FullMethodName = "/reddit.RedditEngine/FetchSubredditState"
	RedditEngine_SendDirectMessage_FullMethodName   = "/reddit.RedditEngine/SendDirectMessage"
)

// RedditEngineClient is the client API for RedditEngine service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// Service Declaration (for reference, even though Proto.Actor is used)
type RedditEngineClient interface {
	RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error)
	CreateSubreddit(ctx context.Context, in *CreateSubredditRequest, opts ...grpc.CallOption) (*CreateSubredditResponse, error)
	JoinSubreddit(ctx context.Context, in *JoinSubredditRequest, opts ...grpc.CallOption) (*JoinSubredditResponse, error)
	Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error)
	Comment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error)
	Vote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error)
	FetchSubredditState(ctx context.Context, in *FetchSubredditStateRequest, opts ...grpc.CallOption) (*FetchSubredditStateResponse, error)
	SendDirectMessage(ctx context.Context, in *DirectMessageRequest, opts ...grpc.CallOption) (*DirectMessageResponse, error)
}

type redditEngineClient struct {
	cc grpc.ClientConnInterface
}

func NewRedditEngineClient(cc grpc.ClientConnInterface) RedditEngineClient {
	return &redditEngineClient{cc}
}

func (c *redditEngineClient) RegisterUser(ctx context.Context, in *RegisterUserRequest, opts ...grpc.CallOption) (*RegisterUserResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(RegisterUserResponse)
	err := c.cc.Invoke(ctx, RedditEngine_RegisterUser_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redditEngineClient) CreateSubreddit(ctx context.Context, in *CreateSubredditRequest, opts ...grpc.CallOption) (*CreateSubredditResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CreateSubredditResponse)
	err := c.cc.Invoke(ctx, RedditEngine_CreateSubreddit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redditEngineClient) JoinSubreddit(ctx context.Context, in *JoinSubredditRequest, opts ...grpc.CallOption) (*JoinSubredditResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(JoinSubredditResponse)
	err := c.cc.Invoke(ctx, RedditEngine_JoinSubreddit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redditEngineClient) Post(ctx context.Context, in *PostRequest, opts ...grpc.CallOption) (*PostResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PostResponse)
	err := c.cc.Invoke(ctx, RedditEngine_Post_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redditEngineClient) Comment(ctx context.Context, in *CommentRequest, opts ...grpc.CallOption) (*CommentResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(CommentResponse)
	err := c.cc.Invoke(ctx, RedditEngine_Comment_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redditEngineClient) Vote(ctx context.Context, in *VoteRequest, opts ...grpc.CallOption) (*VoteResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(VoteResponse)
	err := c.cc.Invoke(ctx, RedditEngine_Vote_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redditEngineClient) FetchSubredditState(ctx context.Context, in *FetchSubredditStateRequest, opts ...grpc.CallOption) (*FetchSubredditStateResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(FetchSubredditStateResponse)
	err := c.cc.Invoke(ctx, RedditEngine_FetchSubredditState_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *redditEngineClient) SendDirectMessage(ctx context.Context, in *DirectMessageRequest, opts ...grpc.CallOption) (*DirectMessageResponse, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(DirectMessageResponse)
	err := c.cc.Invoke(ctx, RedditEngine_SendDirectMessage_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// RedditEngineServer is the server API for RedditEngine service.
// All implementations must embed UnimplementedRedditEngineServer
// for forward compatibility.
//
// Service Declaration (for reference, even though Proto.Actor is used)
type RedditEngineServer interface {
	RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error)
	CreateSubreddit(context.Context, *CreateSubredditRequest) (*CreateSubredditResponse, error)
	JoinSubreddit(context.Context, *JoinSubredditRequest) (*JoinSubredditResponse, error)
	Post(context.Context, *PostRequest) (*PostResponse, error)
	Comment(context.Context, *CommentRequest) (*CommentResponse, error)
	Vote(context.Context, *VoteRequest) (*VoteResponse, error)
	FetchSubredditState(context.Context, *FetchSubredditStateRequest) (*FetchSubredditStateResponse, error)
	SendDirectMessage(context.Context, *DirectMessageRequest) (*DirectMessageResponse, error)
	mustEmbedUnimplementedRedditEngineServer()
}

// UnimplementedRedditEngineServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedRedditEngineServer struct{}

func (UnimplementedRedditEngineServer) RegisterUser(context.Context, *RegisterUserRequest) (*RegisterUserResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RegisterUser not implemented")
}
func (UnimplementedRedditEngineServer) CreateSubreddit(context.Context, *CreateSubredditRequest) (*CreateSubredditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method CreateSubreddit not implemented")
}
func (UnimplementedRedditEngineServer) JoinSubreddit(context.Context, *JoinSubredditRequest) (*JoinSubredditResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method JoinSubreddit not implemented")
}
func (UnimplementedRedditEngineServer) Post(context.Context, *PostRequest) (*PostResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (UnimplementedRedditEngineServer) Comment(context.Context, *CommentRequest) (*CommentResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Comment not implemented")
}
func (UnimplementedRedditEngineServer) Vote(context.Context, *VoteRequest) (*VoteResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Vote not implemented")
}
func (UnimplementedRedditEngineServer) FetchSubredditState(context.Context, *FetchSubredditStateRequest) (*FetchSubredditStateResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method FetchSubredditState not implemented")
}
func (UnimplementedRedditEngineServer) SendDirectMessage(context.Context, *DirectMessageRequest) (*DirectMessageResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method SendDirectMessage not implemented")
}
func (UnimplementedRedditEngineServer) mustEmbedUnimplementedRedditEngineServer() {}
func (UnimplementedRedditEngineServer) testEmbeddedByValue()                      {}

// UnsafeRedditEngineServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to RedditEngineServer will
// result in compilation errors.
type UnsafeRedditEngineServer interface {
	mustEmbedUnimplementedRedditEngineServer()
}

func RegisterRedditEngineServer(s grpc.ServiceRegistrar, srv RedditEngineServer) {
	// If the following call pancis, it indicates UnimplementedRedditEngineServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&RedditEngine_ServiceDesc, srv)
}

func _RedditEngine_RegisterUser_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RegisterUserRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedditEngineServer).RegisterUser(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedditEngine_RegisterUser_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedditEngineServer).RegisterUser(ctx, req.(*RegisterUserRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedditEngine_CreateSubreddit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CreateSubredditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedditEngineServer).CreateSubreddit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedditEngine_CreateSubreddit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedditEngineServer).CreateSubreddit(ctx, req.(*CreateSubredditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedditEngine_JoinSubreddit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(JoinSubredditRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedditEngineServer).JoinSubreddit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedditEngine_JoinSubreddit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedditEngineServer).JoinSubreddit(ctx, req.(*JoinSubredditRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedditEngine_Post_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PostRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedditEngineServer).Post(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedditEngine_Post_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedditEngineServer).Post(ctx, req.(*PostRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedditEngine_Comment_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommentRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedditEngineServer).Comment(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedditEngine_Comment_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedditEngineServer).Comment(ctx, req.(*CommentRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedditEngine_Vote_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(VoteRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedditEngineServer).Vote(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedditEngine_Vote_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedditEngineServer).Vote(ctx, req.(*VoteRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedditEngine_FetchSubredditState_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(FetchSubredditStateRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedditEngineServer).FetchSubredditState(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedditEngine_FetchSubredditState_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedditEngineServer).FetchSubredditState(ctx, req.(*FetchSubredditStateRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _RedditEngine_SendDirectMessage_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(DirectMessageRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(RedditEngineServer).SendDirectMessage(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: RedditEngine_SendDirectMessage_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(RedditEngineServer).SendDirectMessage(ctx, req.(*DirectMessageRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// RedditEngine_ServiceDesc is the grpc.ServiceDesc for RedditEngine service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var RedditEngine_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "reddit.RedditEngine",
	HandlerType: (*RedditEngineServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RegisterUser",
			Handler:    _RedditEngine_RegisterUser_Handler,
		},
		{
			MethodName: "CreateSubreddit",
			Handler:    _RedditEngine_CreateSubreddit_Handler,
		},
		{
			MethodName: "JoinSubreddit",
			Handler:    _RedditEngine_JoinSubreddit_Handler,
		},
		{
			MethodName: "Post",
			Handler:    _RedditEngine_Post_Handler,
		},
		{
			MethodName: "Comment",
			Handler:    _RedditEngine_Comment_Handler,
		},
		{
			MethodName: "Vote",
			Handler:    _RedditEngine_Vote_Handler,
		},
		{
			MethodName: "FetchSubredditState",
			Handler:    _RedditEngine_FetchSubredditState_Handler,
		},
		{
			MethodName: "SendDirectMessage",
			Handler:    _RedditEngine_SendDirectMessage_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "proto/engine.proto",
}
