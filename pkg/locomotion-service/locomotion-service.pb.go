// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v4.25.2
// source: locomotion-server/locomotion-server.proto

package locomotion_service

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Locomotion struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId       int64                  `protobuf:"varint,2,opt,name=user_id,proto3" json:"user_id,omitempty"`
	Country      string                 `protobuf:"bytes,3,opt,name=country,proto3" json:"country,omitempty"`
	LengthKm     int64                  `protobuf:"varint,4,opt,name=length_km,proto3" json:"length_km,omitempty"`
	DurationDays int64                  `protobuf:"varint,5,opt,name=duration_days,proto3" json:"duration_days,omitempty"`
	CreatedAt    *timestamppb.Timestamp `protobuf:"bytes,6,opt,name=created_at,proto3" json:"created_at,omitempty"`
}

func (x *Locomotion) Reset() {
	*x = Locomotion{}
	if protoimpl.UnsafeEnabled {
		mi := &file_locomotion_service_locomotion_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Locomotion) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Locomotion) ProtoMessage() {}

func (x *Locomotion) ProtoReflect() protoreflect.Message {
	mi := &file_locomotion_service_locomotion_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Locomotion.ProtoReflect.Descriptor instead.
func (*Locomotion) Descriptor() ([]byte, []int) {
	return file_locomotion_service_locomotion_service_proto_rawDescGZIP(), []int{0}
}

func (x *Locomotion) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Locomotion) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *Locomotion) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *Locomotion) GetLengthKm() int64 {
	if x != nil {
		return x.LengthKm
	}
	return 0
}

func (x *Locomotion) GetDurationDays() int64 {
	if x != nil {
		return x.DurationDays
	}
	return 0
}

func (x *Locomotion) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

type LocomotionCreate struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId       int64  `protobuf:"varint,1,opt,name=user_id,proto3" json:"user_id,omitempty"`
	Country      string `protobuf:"bytes,2,opt,name=country,proto3" json:"country,omitempty"`
	LengthKm     int64  `protobuf:"varint,3,opt,name=length_km,proto3" json:"length_km,omitempty"`
	DurationDays int64  `protobuf:"varint,4,opt,name=duration_days,proto3" json:"duration_days,omitempty"`
}

func (x *LocomotionCreate) Reset() {
	*x = LocomotionCreate{}
	if protoimpl.UnsafeEnabled {
		mi := &file_locomotion_service_locomotion_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LocomotionCreate) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LocomotionCreate) ProtoMessage() {}

func (x *LocomotionCreate) ProtoReflect() protoreflect.Message {
	mi := &file_locomotion_service_locomotion_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LocomotionCreate.ProtoReflect.Descriptor instead.
func (*LocomotionCreate) Descriptor() ([]byte, []int) {
	return file_locomotion_service_locomotion_service_proto_rawDescGZIP(), []int{1}
}

func (x *LocomotionCreate) GetUserId() int64 {
	if x != nil {
		return x.UserId
	}
	return 0
}

func (x *LocomotionCreate) GetCountry() string {
	if x != nil {
		return x.Country
	}
	return ""
}

func (x *LocomotionCreate) GetLengthKm() int64 {
	if x != nil {
		return x.LengthKm
	}
	return 0
}

func (x *LocomotionCreate) GetDurationDays() int64 {
	if x != nil {
		return x.DurationDays
	}
	return 0
}

type GetLocomotionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Locomotion *Locomotion `protobuf:"bytes,1,opt,name=locomotion,proto3" json:"locomotion,omitempty"`
}

func (x *GetLocomotionRequest) Reset() {
	*x = GetLocomotionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_locomotion_service_locomotion_service_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLocomotionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLocomotionRequest) ProtoMessage() {}

func (x *GetLocomotionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_locomotion_service_locomotion_service_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLocomotionRequest.ProtoReflect.Descriptor instead.
func (*GetLocomotionRequest) Descriptor() ([]byte, []int) {
	return file_locomotion_service_locomotion_service_proto_rawDescGZIP(), []int{2}
}

func (x *GetLocomotionRequest) GetLocomotion() *Locomotion {
	if x != nil {
		return x.Locomotion
	}
	return nil
}

type CreateLocomotionResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Response bool `protobuf:"varint,1,opt,name=response,proto3" json:"response,omitempty"`
}

func (x *CreateLocomotionResponse) Reset() {
	*x = CreateLocomotionResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_locomotion_service_locomotion_service_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateLocomotionResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateLocomotionResponse) ProtoMessage() {}

func (x *CreateLocomotionResponse) ProtoReflect() protoreflect.Message {
	mi := &file_locomotion_service_locomotion_service_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateLocomotionResponse.ProtoReflect.Descriptor instead.
func (*CreateLocomotionResponse) Descriptor() ([]byte, []int) {
	return file_locomotion_service_locomotion_service_proto_rawDescGZIP(), []int{3}
}

func (x *CreateLocomotionResponse) GetResponse() bool {
	if x != nil {
		return x.Response
	}
	return false
}

var File_locomotion_service_locomotion_service_proto protoreflect.FileDescriptor

var file_locomotion_service_locomotion_service_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x6c, 0x6f, 0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x2f, 0x6c, 0x6f, 0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2d,
	0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0a, 0x6c,
	0x6f, 0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1c, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2f, 0x61, 0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61,
	0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x01, 0x0a, 0x0a, 0x4c, 0x6f, 0x63,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69,
	0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6c,
	0x65, 0x6e, 0x67, 0x74, 0x68, 0x5f, 0x6b, 0x6d, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09,
	0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x5f, 0x6b, 0x6d, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x75, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x0d, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x12,
	0x3a, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x52,
	0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x22, 0x8a, 0x01, 0x0a, 0x10,
	0x4c, 0x6f, 0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f,
	0x75, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x72, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x5f, 0x6b,
	0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x6c, 0x65, 0x6e, 0x67, 0x74, 0x68, 0x5f,
	0x6b, 0x6d, 0x12, 0x24, 0x0a, 0x0d, 0x64, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x64,
	0x61, 0x79, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x64, 0x75, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x5f, 0x64, 0x61, 0x79, 0x73, 0x22, 0x4e, 0x0a, 0x14, 0x47, 0x65, 0x74, 0x4c,
	0x6f, 0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x36, 0x0a, 0x0a, 0x6c, 0x6f, 0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x6c, 0x6f, 0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x6c, 0x6f,
	0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x36, 0x0a, 0x18, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x4c, 0x6f, 0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x32, 0xea, 0x01, 0x0a, 0x11, 0x4c, 0x6f, 0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x61, 0x0a, 0x0d, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63,
	0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x20, 0x2e, 0x6c, 0x6f, 0x63, 0x6f, 0x6d, 0x6f,
	0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x4c, 0x6f, 0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x6c, 0x6f, 0x63, 0x6f,
	0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f,
	0x6e, 0x22, 0x16, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x10, 0x3a, 0x01, 0x2a, 0x22, 0x0b, 0x2f, 0x6c,
	0x6f, 0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x72, 0x0a, 0x10, 0x43, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x4c, 0x6f, 0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x1c, 0x2e,
	0x6c, 0x6f, 0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4c, 0x6f, 0x63, 0x6f, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x1a, 0x24, 0x2e, 0x6c, 0x6f,
	0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4c,
	0x6f, 0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x1a, 0x82, 0xd3, 0xe4, 0x93, 0x02, 0x14, 0x3a, 0x01, 0x2a, 0x22, 0x0f, 0x2f, 0x6c,
	0x6f, 0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2f, 0x6e, 0x65, 0x77, 0x42, 0x3e, 0x5a,
	0x3c, 0x6c, 0x6f, 0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76,
	0x69, 0x63, 0x65, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x6c, 0x6f, 0x63, 0x6f, 0x6d, 0x6f, 0x74, 0x69,
	0x6f, 0x6e, 0x2d, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x3b, 0x6c, 0x6f, 0x63, 0x6f, 0x6d,
	0x6f, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x62, 0x06, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_locomotion_service_locomotion_service_proto_rawDescOnce sync.Once
	file_locomotion_service_locomotion_service_proto_rawDescData = file_locomotion_service_locomotion_service_proto_rawDesc
)

func file_locomotion_service_locomotion_service_proto_rawDescGZIP() []byte {
	file_locomotion_service_locomotion_service_proto_rawDescOnce.Do(func() {
		file_locomotion_service_locomotion_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_locomotion_service_locomotion_service_proto_rawDescData)
	})
	return file_locomotion_service_locomotion_service_proto_rawDescData
}

var file_locomotion_service_locomotion_service_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_locomotion_service_locomotion_service_proto_goTypes = []any{
	(*Locomotion)(nil),               // 0: locomotion.Locomotion
	(*LocomotionCreate)(nil),         // 1: locomotion.LocomotionCreate
	(*GetLocomotionRequest)(nil),     // 2: locomotion.GetLocomotionRequest
	(*CreateLocomotionResponse)(nil), // 3: locomotion.CreateLocomotionResponse
	(*timestamppb.Timestamp)(nil),    // 4: google.protobuf.Timestamp
}
var file_locomotion_service_locomotion_service_proto_depIdxs = []int32{
	4, // 0: locomotion.Locomotion.created_at:type_name -> google.protobuf.Timestamp
	0, // 1: locomotion.GetLocomotionRequest.locomotion:type_name -> locomotion.Locomotion
	2, // 2: locomotion.LocomotionService.GetLocomotion:input_type -> locomotion.GetLocomotionRequest
	1, // 3: locomotion.LocomotionService.CreateLocomotion:input_type -> locomotion.LocomotionCreate
	0, // 4: locomotion.LocomotionService.GetLocomotion:output_type -> locomotion.Locomotion
	3, // 5: locomotion.LocomotionService.CreateLocomotion:output_type -> locomotion.CreateLocomotionResponse
	4, // [4:6] is the sub-list for method output_type
	2, // [2:4] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_locomotion_service_locomotion_service_proto_init() }
func file_locomotion_service_locomotion_service_proto_init() {
	if File_locomotion_service_locomotion_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_locomotion_service_locomotion_service_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*Locomotion); i {
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
		file_locomotion_service_locomotion_service_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*LocomotionCreate); i {
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
		file_locomotion_service_locomotion_service_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*GetLocomotionRequest); i {
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
		file_locomotion_service_locomotion_service_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*CreateLocomotionResponse); i {
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
			RawDescriptor: file_locomotion_service_locomotion_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_locomotion_service_locomotion_service_proto_goTypes,
		DependencyIndexes: file_locomotion_service_locomotion_service_proto_depIdxs,
		MessageInfos:      file_locomotion_service_locomotion_service_proto_msgTypes,
	}.Build()
	File_locomotion_service_locomotion_service_proto = out.File
	file_locomotion_service_locomotion_service_proto_rawDesc = nil
	file_locomotion_service_locomotion_service_proto_goTypes = nil
	file_locomotion_service_locomotion_service_proto_depIdxs = nil
}
