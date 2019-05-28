// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/shared_collection.proto

package events

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type UpdateCollectionMentions_UpdateType int32

const (
	UpdateCollectionMentions_AddMention    UpdateCollectionMentions_UpdateType = 0
	UpdateCollectionMentions_RemoveMention UpdateCollectionMentions_UpdateType = 1
	UpdateCollectionMentions_UpdateMention UpdateCollectionMentions_UpdateType = 2
)

var UpdateCollectionMentions_UpdateType_name = map[int32]string{
	0: "AddMention",
	1: "RemoveMention",
	2: "UpdateMention",
}

var UpdateCollectionMentions_UpdateType_value = map[string]int32{
	"AddMention":    0,
	"RemoveMention": 1,
	"UpdateMention": 2,
}

func (x UpdateCollectionMentions_UpdateType) String() string {
	return proto.EnumName(UpdateCollectionMentions_UpdateType_name, int32(x))
}

func (UpdateCollectionMentions_UpdateType) EnumDescriptor() ([]byte, []int) {
	return fileDescriptor_b84d98cc82013cd4, []int{1, 0}
}

type NewCollectionShare struct {
	FolderId             string   `protobuf:"bytes,10,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	ShareId              string   `protobuf:"bytes,11,opt,name=share_id,json=shareId,proto3" json:"share_id,omitempty"`
	HistoryId            string   `protobuf:"bytes,12,opt,name=history_id,json=historyId,proto3" json:"history_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *NewCollectionShare) Reset()         { *m = NewCollectionShare{} }
func (m *NewCollectionShare) String() string { return proto.CompactTextString(m) }
func (*NewCollectionShare) ProtoMessage()    {}
func (*NewCollectionShare) Descriptor() ([]byte, []int) {
	return fileDescriptor_b84d98cc82013cd4, []int{0}
}

func (m *NewCollectionShare) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_NewCollectionShare.Unmarshal(m, b)
}
func (m *NewCollectionShare) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_NewCollectionShare.Marshal(b, m, deterministic)
}
func (m *NewCollectionShare) XXX_Merge(src proto.Message) {
	xxx_messageInfo_NewCollectionShare.Merge(m, src)
}
func (m *NewCollectionShare) XXX_Size() int {
	return xxx_messageInfo_NewCollectionShare.Size(m)
}
func (m *NewCollectionShare) XXX_DiscardUnknown() {
	xxx_messageInfo_NewCollectionShare.DiscardUnknown(m)
}

var xxx_messageInfo_NewCollectionShare proto.InternalMessageInfo

func (m *NewCollectionShare) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *NewCollectionShare) GetShareId() string {
	if m != nil {
		return m.ShareId
	}
	return ""
}

func (m *NewCollectionShare) GetHistoryId() string {
	if m != nil {
		return m.HistoryId
	}
	return ""
}

type UpdateCollectionMentions struct {
	FolderId             string                              `protobuf:"bytes,10,opt,name=folder_id,json=folderId,proto3" json:"folder_id,omitempty"`
	MentionIds           []string                            `protobuf:"bytes,11,rep,name=mention_ids,json=mentionIds,proto3" json:"mention_ids,omitempty"`
	UpdateType           UpdateCollectionMentions_UpdateType `protobuf:"varint,12,opt,name=update_type,json=updateType,proto3,enum=events.UpdateCollectionMentions_UpdateType" json:"update_type,omitempty"`
	HistoryId            string                              `protobuf:"bytes,13,opt,name=history_id,json=historyId,proto3" json:"history_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                            `json:"-"`
	XXX_unrecognized     []byte                              `json:"-"`
	XXX_sizecache        int32                               `json:"-"`
}

func (m *UpdateCollectionMentions) Reset()         { *m = UpdateCollectionMentions{} }
func (m *UpdateCollectionMentions) String() string { return proto.CompactTextString(m) }
func (*UpdateCollectionMentions) ProtoMessage()    {}
func (*UpdateCollectionMentions) Descriptor() ([]byte, []int) {
	return fileDescriptor_b84d98cc82013cd4, []int{1}
}

func (m *UpdateCollectionMentions) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_UpdateCollectionMentions.Unmarshal(m, b)
}
func (m *UpdateCollectionMentions) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_UpdateCollectionMentions.Marshal(b, m, deterministic)
}
func (m *UpdateCollectionMentions) XXX_Merge(src proto.Message) {
	xxx_messageInfo_UpdateCollectionMentions.Merge(m, src)
}
func (m *UpdateCollectionMentions) XXX_Size() int {
	return xxx_messageInfo_UpdateCollectionMentions.Size(m)
}
func (m *UpdateCollectionMentions) XXX_DiscardUnknown() {
	xxx_messageInfo_UpdateCollectionMentions.DiscardUnknown(m)
}

var xxx_messageInfo_UpdateCollectionMentions proto.InternalMessageInfo

func (m *UpdateCollectionMentions) GetFolderId() string {
	if m != nil {
		return m.FolderId
	}
	return ""
}

func (m *UpdateCollectionMentions) GetMentionIds() []string {
	if m != nil {
		return m.MentionIds
	}
	return nil
}

func (m *UpdateCollectionMentions) GetUpdateType() UpdateCollectionMentions_UpdateType {
	if m != nil {
		return m.UpdateType
	}
	return UpdateCollectionMentions_AddMention
}

func (m *UpdateCollectionMentions) GetHistoryId() string {
	if m != nil {
		return m.HistoryId
	}
	return ""
}

func init() {
	proto.RegisterEnum("events.UpdateCollectionMentions_UpdateType", UpdateCollectionMentions_UpdateType_name, UpdateCollectionMentions_UpdateType_value)
	proto.RegisterType((*NewCollectionShare)(nil), "events.NewCollectionShare")
	proto.RegisterType((*UpdateCollectionMentions)(nil), "events.UpdateCollectionMentions")
}

func init() { proto.RegisterFile("events/shared_collection.proto", fileDescriptor_b84d98cc82013cd4) }

var fileDescriptor_b84d98cc82013cd4 = []byte{
	// 261 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x91, 0xcf, 0x4a, 0xc3, 0x40,
	0x10, 0xc6, 0x6d, 0x85, 0x9a, 0x7c, 0xb1, 0x45, 0xf7, 0x14, 0x11, 0xb5, 0xf4, 0x54, 0x10, 0x22,
	0xe8, 0x13, 0x58, 0x4f, 0x01, 0xf5, 0x10, 0xf5, 0xe2, 0x25, 0xd4, 0xce, 0x48, 0x83, 0x69, 0x36,
	0x64, 0xb7, 0x95, 0x3c, 0x82, 0x6f, 0x2d, 0xfb, 0xa7, 0x2d, 0x04, 0xf4, 0x94, 0xec, 0xef, 0x37,
	0x33, 0x1f, 0xc3, 0xe0, 0x92, 0x37, 0x5c, 0x69, 0x75, 0xa3, 0x96, 0xf3, 0x86, 0x29, 0x5f, 0xc8,
	0xb2, 0xe4, 0x85, 0x2e, 0x64, 0x95, 0xd4, 0x8d, 0xd4, 0x52, 0x0c, 0x9c, 0x9f, 0x7c, 0x41, 0x3c,
	0xf3, 0xf7, 0xc3, 0x4e, 0xbf, 0x98, 0x7a, 0x71, 0x8e, 0xf0, 0x53, 0x96, 0xc4, 0x4d, 0x5e, 0x50,
	0x8c, 0x71, 0x6f, 0x1a, 0x66, 0x81, 0x03, 0x29, 0x89, 0x33, 0x04, 0x76, 0xaa, 0x71, 0x91, 0x75,
	0x47, 0xf6, 0x9d, 0x92, 0xb8, 0x00, 0x96, 0x85, 0xd2, 0xb2, 0x69, 0x8d, 0x3c, 0xb6, 0x32, 0xf4,
	0x24, 0xa5, 0xc9, 0x4f, 0x1f, 0xf1, 0x5b, 0x4d, 0x73, 0xcd, 0xfb, 0xc0, 0x27, 0xae, 0xcc, 0x47,
	0xfd, 0x9f, 0x79, 0x85, 0x68, 0xe5, 0x0a, 0xf3, 0x82, 0x54, 0x1c, 0x8d, 0x0f, 0xa7, 0x61, 0x06,
	0x8f, 0x52, 0x52, 0xe2, 0x11, 0xd1, 0xda, 0x4e, 0xce, 0x75, 0x5b, 0xb3, 0x8d, 0x1e, 0xdd, 0x5e,
	0x27, 0x6e, 0xcb, 0xe4, 0xaf, 0x50, 0x2f, 0x5e, 0xdb, 0x9a, 0x33, 0xac, 0x77, 0xff, 0x9d, 0x3d,
	0x86, 0xdd, 0x3d, 0x66, 0xc0, 0xbe, 0x51, 0x8c, 0x80, 0x7b, 0x22, 0x3f, 0xf2, 0xe4, 0x40, 0x9c,
	0x62, 0x98, 0xf1, 0x4a, 0x6e, 0x78, 0x8b, 0x7a, 0x06, 0xb9, 0x86, 0x2d, 0xea, 0xcf, 0x82, 0x77,
	0x7f, 0x82, 0x8f, 0x81, 0xbd, 0xc8, 0xdd, 0x6f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x0e, 0xd1, 0x37,
	0x7b, 0xb3, 0x01, 0x00, 0x00,
}
