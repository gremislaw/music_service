// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: main.proto

package api

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
	MusicService_Play_FullMethodName                   = "/music_service.api.MusicService/Play"
	MusicService_Pause_FullMethodName                  = "/music_service.api.MusicService/Pause"
	MusicService_AddSong_FullMethodName                = "/music_service.api.MusicService/AddSong"
	MusicService_DeleteSong_FullMethodName             = "/music_service.api.MusicService/DeleteSong"
	MusicService_AddSongToPlaylist_FullMethodName      = "/music_service.api.MusicService/AddSongToPlaylist"
	MusicService_DeleteSongFromPlaylist_FullMethodName = "/music_service.api.MusicService/DeleteSongFromPlaylist"
	MusicService_AddPlaylist_FullMethodName            = "/music_service.api.MusicService/AddPlaylist"
	MusicService_GetPlaylist_FullMethodName            = "/music_service.api.MusicService/GetPlaylist"
	MusicService_DeletePlaylist_FullMethodName         = "/music_service.api.MusicService/DeletePlaylist"
	MusicService_PrintPlaylist_FullMethodName          = "/music_service.api.MusicService/PrintPlaylist"
	MusicService_UpdatePlaylist_FullMethodName         = "/music_service.api.MusicService/UpdatePlaylist"
	MusicService_GetSong_FullMethodName                = "/music_service.api.MusicService/GetSong"
	MusicService_UpdateSong_FullMethodName             = "/music_service.api.MusicService/UpdateSong"
	MusicService_Next_FullMethodName                   = "/music_service.api.MusicService/Next"
	MusicService_Prev_FullMethodName                   = "/music_service.api.MusicService/Prev"
)

// MusicServiceClient is the client API for MusicService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type MusicServiceClient interface {
	Play(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
	Pause(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
	AddSong(ctx context.Context, in *Song, opts ...grpc.CallOption) (*Response, error)
	DeleteSong(ctx context.Context, in *Song, opts ...grpc.CallOption) (*Response, error)
	AddSongToPlaylist(ctx context.Context, in *SongPlaylist, opts ...grpc.CallOption) (*Response, error)
	DeleteSongFromPlaylist(ctx context.Context, in *SongPlaylist, opts ...grpc.CallOption) (*Response, error)
	AddPlaylist(ctx context.Context, in *Playlist, opts ...grpc.CallOption) (*Response, error)
	GetPlaylist(ctx context.Context, in *Playlist, opts ...grpc.CallOption) (*Response, error)
	DeletePlaylist(ctx context.Context, in *Playlist, opts ...grpc.CallOption) (*Response, error)
	PrintPlaylist(ctx context.Context, in *Playlist, opts ...grpc.CallOption) (*Playlist, error)
	UpdatePlaylist(ctx context.Context, in *Playlist, opts ...grpc.CallOption) (*Response, error)
	GetSong(ctx context.Context, in *Song, opts ...grpc.CallOption) (*Song, error)
	UpdateSong(ctx context.Context, in *Song, opts ...grpc.CallOption) (*Response, error)
	Next(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
	Prev(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error)
}

type musicServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewMusicServiceClient(cc grpc.ClientConnInterface) MusicServiceClient {
	return &musicServiceClient{cc}
}

func (c *musicServiceClient) Play(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, MusicService_Play_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) Pause(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, MusicService_Pause_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) AddSong(ctx context.Context, in *Song, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, MusicService_AddSong_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) DeleteSong(ctx context.Context, in *Song, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, MusicService_DeleteSong_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) AddSongToPlaylist(ctx context.Context, in *SongPlaylist, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, MusicService_AddSongToPlaylist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) DeleteSongFromPlaylist(ctx context.Context, in *SongPlaylist, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, MusicService_DeleteSongFromPlaylist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) AddPlaylist(ctx context.Context, in *Playlist, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, MusicService_AddPlaylist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) GetPlaylist(ctx context.Context, in *Playlist, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, MusicService_GetPlaylist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) DeletePlaylist(ctx context.Context, in *Playlist, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, MusicService_DeletePlaylist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) PrintPlaylist(ctx context.Context, in *Playlist, opts ...grpc.CallOption) (*Playlist, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Playlist)
	err := c.cc.Invoke(ctx, MusicService_PrintPlaylist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) UpdatePlaylist(ctx context.Context, in *Playlist, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, MusicService_UpdatePlaylist_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) GetSong(ctx context.Context, in *Song, opts ...grpc.CallOption) (*Song, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Song)
	err := c.cc.Invoke(ctx, MusicService_GetSong_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) UpdateSong(ctx context.Context, in *Song, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, MusicService_UpdateSong_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) Next(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, MusicService_Next_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *musicServiceClient) Prev(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*Response, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(Response)
	err := c.cc.Invoke(ctx, MusicService_Prev_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// MusicServiceServer is the server API for MusicService service.
// All implementations must embed UnimplementedMusicServiceServer
// for forward compatibility.
type MusicServiceServer interface {
	Play(context.Context, *Empty) (*Response, error)
	Pause(context.Context, *Empty) (*Response, error)
	AddSong(context.Context, *Song) (*Response, error)
	DeleteSong(context.Context, *Song) (*Response, error)
	AddSongToPlaylist(context.Context, *SongPlaylist) (*Response, error)
	DeleteSongFromPlaylist(context.Context, *SongPlaylist) (*Response, error)
	AddPlaylist(context.Context, *Playlist) (*Response, error)
	GetPlaylist(context.Context, *Playlist) (*Response, error)
	DeletePlaylist(context.Context, *Playlist) (*Response, error)
	PrintPlaylist(context.Context, *Playlist) (*Playlist, error)
	UpdatePlaylist(context.Context, *Playlist) (*Response, error)
	GetSong(context.Context, *Song) (*Song, error)
	UpdateSong(context.Context, *Song) (*Response, error)
	Next(context.Context, *Empty) (*Response, error)
	Prev(context.Context, *Empty) (*Response, error)
	mustEmbedUnimplementedMusicServiceServer()
}

// UnimplementedMusicServiceServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedMusicServiceServer struct{}

func (UnimplementedMusicServiceServer) Play(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Play not implemented")
}
func (UnimplementedMusicServiceServer) Pause(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Pause not implemented")
}
func (UnimplementedMusicServiceServer) AddSong(context.Context, *Song) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSong not implemented")
}
func (UnimplementedMusicServiceServer) DeleteSong(context.Context, *Song) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSong not implemented")
}
func (UnimplementedMusicServiceServer) AddSongToPlaylist(context.Context, *SongPlaylist) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddSongToPlaylist not implemented")
}
func (UnimplementedMusicServiceServer) DeleteSongFromPlaylist(context.Context, *SongPlaylist) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeleteSongFromPlaylist not implemented")
}
func (UnimplementedMusicServiceServer) AddPlaylist(context.Context, *Playlist) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method AddPlaylist not implemented")
}
func (UnimplementedMusicServiceServer) GetPlaylist(context.Context, *Playlist) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetPlaylist not implemented")
}
func (UnimplementedMusicServiceServer) DeletePlaylist(context.Context, *Playlist) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method DeletePlaylist not implemented")
}
func (UnimplementedMusicServiceServer) PrintPlaylist(context.Context, *Playlist) (*Playlist, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrintPlaylist not implemented")
}
func (UnimplementedMusicServiceServer) UpdatePlaylist(context.Context, *Playlist) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdatePlaylist not implemented")
}
func (UnimplementedMusicServiceServer) GetSong(context.Context, *Song) (*Song, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetSong not implemented")
}
func (UnimplementedMusicServiceServer) UpdateSong(context.Context, *Song) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method UpdateSong not implemented")
}
func (UnimplementedMusicServiceServer) Next(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Next not implemented")
}
func (UnimplementedMusicServiceServer) Prev(context.Context, *Empty) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prev not implemented")
}
func (UnimplementedMusicServiceServer) mustEmbedUnimplementedMusicServiceServer() {}
func (UnimplementedMusicServiceServer) testEmbeddedByValue()                      {}

// UnsafeMusicServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to MusicServiceServer will
// result in compilation errors.
type UnsafeMusicServiceServer interface {
	mustEmbedUnimplementedMusicServiceServer()
}

func RegisterMusicServiceServer(s grpc.ServiceRegistrar, srv MusicServiceServer) {
	// If the following call pancis, it indicates UnimplementedMusicServiceServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&MusicService_ServiceDesc, srv)
}

func _MusicService_Play_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).Play(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_Play_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).Play(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_Pause_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).Pause(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_Pause_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).Pause(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_AddSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Song)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).AddSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_AddSong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).AddSong(ctx, req.(*Song))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_DeleteSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Song)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).DeleteSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_DeleteSong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).DeleteSong(ctx, req.(*Song))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_AddSongToPlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SongPlaylist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).AddSongToPlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_AddSongToPlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).AddSongToPlaylist(ctx, req.(*SongPlaylist))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_DeleteSongFromPlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(SongPlaylist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).DeleteSongFromPlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_DeleteSongFromPlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).DeleteSongFromPlaylist(ctx, req.(*SongPlaylist))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_AddPlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Playlist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).AddPlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_AddPlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).AddPlaylist(ctx, req.(*Playlist))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_GetPlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Playlist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).GetPlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_GetPlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).GetPlaylist(ctx, req.(*Playlist))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_DeletePlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Playlist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).DeletePlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_DeletePlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).DeletePlaylist(ctx, req.(*Playlist))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_PrintPlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Playlist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).PrintPlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_PrintPlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).PrintPlaylist(ctx, req.(*Playlist))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_UpdatePlaylist_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Playlist)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).UpdatePlaylist(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_UpdatePlaylist_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).UpdatePlaylist(ctx, req.(*Playlist))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_GetSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Song)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).GetSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_GetSong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).GetSong(ctx, req.(*Song))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_UpdateSong_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Song)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).UpdateSong(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_UpdateSong_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).UpdateSong(ctx, req.(*Song))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_Next_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).Next(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_Next_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).Next(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _MusicService_Prev_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(MusicServiceServer).Prev(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: MusicService_Prev_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(MusicServiceServer).Prev(ctx, req.(*Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// MusicService_ServiceDesc is the grpc.ServiceDesc for MusicService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var MusicService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "music_service.api.MusicService",
	HandlerType: (*MusicServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Play",
			Handler:    _MusicService_Play_Handler,
		},
		{
			MethodName: "Pause",
			Handler:    _MusicService_Pause_Handler,
		},
		{
			MethodName: "AddSong",
			Handler:    _MusicService_AddSong_Handler,
		},
		{
			MethodName: "DeleteSong",
			Handler:    _MusicService_DeleteSong_Handler,
		},
		{
			MethodName: "AddSongToPlaylist",
			Handler:    _MusicService_AddSongToPlaylist_Handler,
		},
		{
			MethodName: "DeleteSongFromPlaylist",
			Handler:    _MusicService_DeleteSongFromPlaylist_Handler,
		},
		{
			MethodName: "AddPlaylist",
			Handler:    _MusicService_AddPlaylist_Handler,
		},
		{
			MethodName: "GetPlaylist",
			Handler:    _MusicService_GetPlaylist_Handler,
		},
		{
			MethodName: "DeletePlaylist",
			Handler:    _MusicService_DeletePlaylist_Handler,
		},
		{
			MethodName: "PrintPlaylist",
			Handler:    _MusicService_PrintPlaylist_Handler,
		},
		{
			MethodName: "UpdatePlaylist",
			Handler:    _MusicService_UpdatePlaylist_Handler,
		},
		{
			MethodName: "GetSong",
			Handler:    _MusicService_GetSong_Handler,
		},
		{
			MethodName: "UpdateSong",
			Handler:    _MusicService_UpdateSong_Handler,
		},
		{
			MethodName: "Next",
			Handler:    _MusicService_Next_Handler,
		},
		{
			MethodName: "Prev",
			Handler:    _MusicService_Prev_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "main.proto",
}
