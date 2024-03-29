// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.12
// source: playlist_service.proto

package plservice

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ControlRequest_Action int32

const (
	ControlRequest_PLAY  ControlRequest_Action = 0
	ControlRequest_PAUSE ControlRequest_Action = 1
	ControlRequest_NEXT  ControlRequest_Action = 2
	ControlRequest_PREV  ControlRequest_Action = 3
)

// Enum value maps for ControlRequest_Action.
var (
	ControlRequest_Action_name = map[int32]string{
		0: "PLAY",
		1: "PAUSE",
		2: "NEXT",
		3: "PREV",
	}
	ControlRequest_Action_value = map[string]int32{
		"PLAY":  0,
		"PAUSE": 1,
		"NEXT":  2,
		"PREV":  3,
	}
)

func (x ControlRequest_Action) Enum() *ControlRequest_Action {
	p := new(ControlRequest_Action)
	*p = x
	return p
}

func (x ControlRequest_Action) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ControlRequest_Action) Descriptor() protoreflect.EnumDescriptor {
	return file_playlist_service_proto_enumTypes[0].Descriptor()
}

func (ControlRequest_Action) Type() protoreflect.EnumType {
	return &file_playlist_service_proto_enumTypes[0]
}

func (x ControlRequest_Action) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ControlRequest_Action.Descriptor instead.
func (ControlRequest_Action) EnumDescriptor() ([]byte, []int) {
	return file_playlist_service_proto_rawDescGZIP(), []int{8, 0}
}

type CreateSongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Song *Song `protobuf:"bytes,1,opt,name=song,proto3" json:"song,omitempty"`
}

func (x *CreateSongRequest) Reset() {
	*x = CreateSongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSongRequest) ProtoMessage() {}

func (x *CreateSongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSongRequest.ProtoReflect.Descriptor instead.
func (*CreateSongRequest) Descriptor() ([]byte, []int) {
	return file_playlist_service_proto_rawDescGZIP(), []int{0}
}

func (x *CreateSongRequest) GetSong() *Song {
	if x != nil {
		return x.Song
	}
	return nil
}

type CreateSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *CreateSongResponse) Reset() {
	*x = CreateSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateSongResponse) ProtoMessage() {}

func (x *CreateSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateSongResponse.ProtoReflect.Descriptor instead.
func (*CreateSongResponse) Descriptor() ([]byte, []int) {
	return file_playlist_service_proto_rawDescGZIP(), []int{1}
}

func (x *CreateSongResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ReadSongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id []uint64 `protobuf:"varint,1,rep,packed,name=id,proto3" json:"id,omitempty"`
}

func (x *ReadSongRequest) Reset() {
	*x = ReadSongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadSongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadSongRequest) ProtoMessage() {}

func (x *ReadSongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadSongRequest.ProtoReflect.Descriptor instead.
func (*ReadSongRequest) Descriptor() ([]byte, []int) {
	return file_playlist_service_proto_rawDescGZIP(), []int{2}
}

func (x *ReadSongRequest) GetId() []uint64 {
	if x != nil {
		return x.Id
	}
	return nil
}

type ReadSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Song *Song `protobuf:"bytes,1,opt,name=song,proto3" json:"song,omitempty"`
}

func (x *ReadSongResponse) Reset() {
	*x = ReadSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ReadSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ReadSongResponse) ProtoMessage() {}

func (x *ReadSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ReadSongResponse.ProtoReflect.Descriptor instead.
func (*ReadSongResponse) Descriptor() ([]byte, []int) {
	return file_playlist_service_proto_rawDescGZIP(), []int{3}
}

func (x *ReadSongResponse) GetSong() *Song {
	if x != nil {
		return x.Song
	}
	return nil
}

type UpdateSongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Song *Song `protobuf:"bytes,1,opt,name=song,proto3" json:"song,omitempty"`
}

func (x *UpdateSongRequest) Reset() {
	*x = UpdateSongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_service_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSongRequest) ProtoMessage() {}

func (x *UpdateSongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_service_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSongRequest.ProtoReflect.Descriptor instead.
func (*UpdateSongRequest) Descriptor() ([]byte, []int) {
	return file_playlist_service_proto_rawDescGZIP(), []int{4}
}

func (x *UpdateSongRequest) GetSong() *Song {
	if x != nil {
		return x.Song
	}
	return nil
}

type UpdateSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *UpdateSongResponse) Reset() {
	*x = UpdateSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_service_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UpdateSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UpdateSongResponse) ProtoMessage() {}

func (x *UpdateSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_service_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UpdateSongResponse.ProtoReflect.Descriptor instead.
func (*UpdateSongResponse) Descriptor() ([]byte, []int) {
	return file_playlist_service_proto_rawDescGZIP(), []int{5}
}

func (x *UpdateSongResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteSongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSongRequest) Reset() {
	*x = DeleteSongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_service_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSongRequest) ProtoMessage() {}

func (x *DeleteSongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_service_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSongRequest.ProtoReflect.Descriptor instead.
func (*DeleteSongRequest) Descriptor() ([]byte, []int) {
	return file_playlist_service_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteSongRequest) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type DeleteSongResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id uint64 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *DeleteSongResponse) Reset() {
	*x = DeleteSongResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_service_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteSongResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteSongResponse) ProtoMessage() {}

func (x *DeleteSongResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_service_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteSongResponse.ProtoReflect.Descriptor instead.
func (*DeleteSongResponse) Descriptor() ([]byte, []int) {
	return file_playlist_service_proto_rawDescGZIP(), []int{7}
}

func (x *DeleteSongResponse) GetId() uint64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type ControlRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Action ControlRequest_Action `protobuf:"varint,1,opt,name=action,proto3,enum=plservice.ControlRequest_Action" json:"action,omitempty"`
}

func (x *ControlRequest) Reset() {
	*x = ControlRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_service_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControlRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlRequest) ProtoMessage() {}

func (x *ControlRequest) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_service_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlRequest.ProtoReflect.Descriptor instead.
func (*ControlRequest) Descriptor() ([]byte, []int) {
	return file_playlist_service_proto_rawDescGZIP(), []int{8}
}

func (x *ControlRequest) GetAction() ControlRequest_Action {
	if x != nil {
		return x.Action
	}
	return ControlRequest_PLAY
}

type ControlResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status string `protobuf:"bytes,1,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *ControlResponse) Reset() {
	*x = ControlResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_playlist_service_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ControlResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ControlResponse) ProtoMessage() {}

func (x *ControlResponse) ProtoReflect() protoreflect.Message {
	mi := &file_playlist_service_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ControlResponse.ProtoReflect.Descriptor instead.
func (*ControlResponse) Descriptor() ([]byte, []int) {
	return file_playlist_service_proto_rawDescGZIP(), []int{9}
}

func (x *ControlResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

var File_playlist_service_proto protoreflect.FileDescriptor

var file_playlist_service_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x6c, 0x61, 0x79, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x09, 0x70, 0x6c, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x1a, 0x12, 0x73, 0x6f, 0x6e, 0x67, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x38, 0x0a, 0x11, 0x43, 0x72, 0x65, 0x61, 0x74,
	0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04,
	0x73, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6c, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x04, 0x73, 0x6f, 0x6e,
	0x67, 0x22, 0x24, 0x0a, 0x12, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x21, 0x0a, 0x0f, 0x52, 0x65, 0x61, 0x64, 0x53,
	0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x37, 0x0a, 0x10, 0x52, 0x65,
	0x61, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x23,
	0x0a, 0x04, 0x73, 0x6f, 0x6e, 0x67, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70,
	0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x04, 0x73,
	0x6f, 0x6e, 0x67, 0x22, 0x38, 0x0a, 0x11, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x04, 0x73, 0x6f, 0x6e, 0x67,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0f, 0x2e, 0x70, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x04, 0x73, 0x6f, 0x6e, 0x67, 0x22, 0x24, 0x0a,
	0x12, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f,
	0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x02, 0x69, 0x64, 0x22, 0x23, 0x0a, 0x11, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x24, 0x0a, 0x12, 0x44, 0x65, 0x6c, 0x65,
	0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e,
	0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x52, 0x02, 0x69, 0x64, 0x22, 0x7d,
	0x0a, 0x0e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x38, 0x0a, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x20, 0x2e, 0x70, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x06, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x31, 0x0a, 0x06, 0x41, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x4c, 0x41, 0x59, 0x10, 0x00, 0x12, 0x09,
	0x0a, 0x05, 0x50, 0x41, 0x55, 0x53, 0x45, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x45, 0x58,
	0x54, 0x10, 0x02, 0x12, 0x08, 0x0a, 0x04, 0x50, 0x52, 0x45, 0x56, 0x10, 0x03, 0x22, 0x29, 0x0a,
	0x0f, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x32, 0x85, 0x03, 0x0a, 0x0f, 0x50, 0x6c, 0x61,
	0x79, 0x6c, 0x69, 0x73, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x4b, 0x0a, 0x0a,
	0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x1c, 0x2e, 0x70, 0x6c, 0x73,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x6c, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x47, 0x0a, 0x08, 0x52, 0x65, 0x61,
	0x64, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x1a, 0x2e, 0x70, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x2e, 0x52, 0x65, 0x61, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x1b, 0x2e, 0x70, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x52, 0x65,
	0x61, 0x64, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x30, 0x01, 0x12, 0x4b, 0x0a, 0x0a, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67,
	0x12, 0x1c, 0x2e, 0x70, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64,
	0x61, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d,
	0x2e, 0x70, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x55, 0x70, 0x64, 0x61, 0x74,
	0x65, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12,
	0x4b, 0x0a, 0x0a, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x1c, 0x2e,
	0x70, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1d, 0x2e, 0x70, 0x6c,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x53, 0x6f,
	0x6e, 0x67, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x42, 0x0a, 0x07,
	0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x12, 0x19, 0x2e, 0x70, 0x6c, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2e, 0x43, 0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x70, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x43,
	0x6f, 0x6e, 0x74, 0x72, 0x6f, 0x6c, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x42, 0x0d, 0x5a, 0x0b, 0x2e, 0x2f, 0x70, 0x6c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_playlist_service_proto_rawDescOnce sync.Once
	file_playlist_service_proto_rawDescData = file_playlist_service_proto_rawDesc
)

func file_playlist_service_proto_rawDescGZIP() []byte {
	file_playlist_service_proto_rawDescOnce.Do(func() {
		file_playlist_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_playlist_service_proto_rawDescData)
	})
	return file_playlist_service_proto_rawDescData
}

var file_playlist_service_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_playlist_service_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_playlist_service_proto_goTypes = []interface{}{
	(ControlRequest_Action)(0), // 0: plservice.ControlRequest.Action
	(*CreateSongRequest)(nil),  // 1: plservice.CreateSongRequest
	(*CreateSongResponse)(nil), // 2: plservice.CreateSongResponse
	(*ReadSongRequest)(nil),    // 3: plservice.ReadSongRequest
	(*ReadSongResponse)(nil),   // 4: plservice.ReadSongResponse
	(*UpdateSongRequest)(nil),  // 5: plservice.UpdateSongRequest
	(*UpdateSongResponse)(nil), // 6: plservice.UpdateSongResponse
	(*DeleteSongRequest)(nil),  // 7: plservice.DeleteSongRequest
	(*DeleteSongResponse)(nil), // 8: plservice.DeleteSongResponse
	(*ControlRequest)(nil),     // 9: plservice.ControlRequest
	(*ControlResponse)(nil),    // 10: plservice.ControlResponse
	(*Song)(nil),               // 11: plservice.Song
}
var file_playlist_service_proto_depIdxs = []int32{
	11, // 0: plservice.CreateSongRequest.song:type_name -> plservice.Song
	11, // 1: plservice.ReadSongResponse.song:type_name -> plservice.Song
	11, // 2: plservice.UpdateSongRequest.song:type_name -> plservice.Song
	0,  // 3: plservice.ControlRequest.action:type_name -> plservice.ControlRequest.Action
	1,  // 4: plservice.PlaylistService.CreateSong:input_type -> plservice.CreateSongRequest
	3,  // 5: plservice.PlaylistService.ReadSong:input_type -> plservice.ReadSongRequest
	5,  // 6: plservice.PlaylistService.UpdateSong:input_type -> plservice.UpdateSongRequest
	7,  // 7: plservice.PlaylistService.DeleteSong:input_type -> plservice.DeleteSongRequest
	9,  // 8: plservice.PlaylistService.Control:input_type -> plservice.ControlRequest
	2,  // 9: plservice.PlaylistService.CreateSong:output_type -> plservice.CreateSongResponse
	4,  // 10: plservice.PlaylistService.ReadSong:output_type -> plservice.ReadSongResponse
	6,  // 11: plservice.PlaylistService.UpdateSong:output_type -> plservice.UpdateSongResponse
	8,  // 12: plservice.PlaylistService.DeleteSong:output_type -> plservice.DeleteSongResponse
	10, // 13: plservice.PlaylistService.Control:output_type -> plservice.ControlResponse
	9,  // [9:14] is the sub-list for method output_type
	4,  // [4:9] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_playlist_service_proto_init() }
func file_playlist_service_proto_init() {
	if File_playlist_service_proto != nil {
		return
	}
	file_song_message_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_playlist_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSongRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_playlist_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateSongResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_playlist_service_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadSongRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_playlist_service_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ReadSongResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_playlist_service_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSongRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_playlist_service_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UpdateSongResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_playlist_service_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSongRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_playlist_service_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteSongResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_playlist_service_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControlRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_playlist_service_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ControlResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_playlist_service_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_playlist_service_proto_goTypes,
		DependencyIndexes: file_playlist_service_proto_depIdxs,
		EnumInfos:         file_playlist_service_proto_enumTypes,
		MessageInfos:      file_playlist_service_proto_msgTypes,
	}.Build()
	File_playlist_service_proto = out.File
	file_playlist_service_proto_rawDesc = nil
	file_playlist_service_proto_goTypes = nil
	file_playlist_service_proto_depIdxs = nil
}
