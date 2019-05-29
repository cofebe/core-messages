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
	ShareId              string                              `protobuf:"bytes,11,opt,name=share_id,json=shareId,proto3" json:"share_id,omitempty"`
	MentionIds           []string                            `protobuf:"bytes,12,rep,name=mention_ids,json=mentionIds,proto3" json:"mention_ids,omitempty"`
	UpdateType           UpdateCollectionMentions_UpdateType `protobuf:"varint,13,opt,name=update_type,json=updateType,proto3,enum=events.UpdateCollectionMentions_UpdateType" json:"update_type,omitempty"`
	HistoryId            string                              `protobuf:"bytes,14,opt,name=history_id,json=historyId,proto3" json:"history_id,omitempty"`
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

func (m *UpdateCollectionMentions) GetShareId() string {
	if m != nil {
		return m.ShareId
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
	// 265 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x9c, 0x91, 0xcd, 0x4a, 0xc3, 0x40,
	0x10, 0x80, 0x4d, 0x85, 0x9a, 0x4c, 0xda, 0xa0, 0x73, 0x8a, 0x88, 0x1a, 0x7a, 0x0a, 0x08, 0x11,
	0xf4, 0x09, 0xac, 0xa7, 0x80, 0x7a, 0x88, 0x7a, 0xf1, 0x12, 0x6a, 0x67, 0xa4, 0xc1, 0x34, 0x1b,
	0xb2, 0xdb, 0x4a, 0xde, 0xc7, 0x07, 0x95, 0xfd, 0x69, 0x0b, 0x05, 0x2f, 0x3d, 0xed, 0xee, 0xf7,
	0xed, 0xcc, 0xec, 0xec, 0xc0, 0x15, 0xaf, 0xb9, 0x51, 0xf2, 0x56, 0x2e, 0x66, 0x1d, 0x53, 0x39,
	0x17, 0x75, 0xcd, 0x73, 0x55, 0x89, 0x26, 0x6b, 0x3b, 0xa1, 0x04, 0x0e, 0xad, 0x9f, 0x7c, 0x03,
	0xbe, 0xf0, 0xcf, 0xe3, 0x56, 0xbf, 0xea, 0xfb, 0x78, 0x01, 0xc1, 0x97, 0xa8, 0x89, 0xbb, 0xb2,
	0xa2, 0x18, 0x12, 0x2f, 0x0d, 0x0a, 0xdf, 0x82, 0x9c, 0xf0, 0x1c, 0x7c, 0x93, 0x55, 0xbb, 0xd0,
	0xb8, 0x13, 0x73, 0xce, 0x09, 0x2f, 0x01, 0x16, 0x95, 0x54, 0xa2, 0xeb, 0xb5, 0x1c, 0x19, 0x19,
	0x38, 0x92, 0xd3, 0xe4, 0x77, 0x00, 0xf1, 0x7b, 0x4b, 0x33, 0xc5, 0xbb, 0x82, 0xcf, 0xdc, 0xe8,
	0x45, 0x1e, 0x5c, 0xf3, 0x1a, 0xc2, 0xa5, 0xcd, 0x51, 0x56, 0x24, 0xe3, 0x51, 0x72, 0x9c, 0x06,
	0x05, 0x38, 0x94, 0x93, 0xc4, 0x27, 0x08, 0x57, 0xa6, 0x68, 0xa9, 0xfa, 0x96, 0xe3, 0x71, 0xe2,
	0xa5, 0xd1, 0xdd, 0x4d, 0x66, 0x3f, 0x20, 0xfb, 0xef, 0x3d, 0x4e, 0xbc, 0xf5, 0x2d, 0x17, 0xb0,
	0xda, 0xee, 0xf7, 0x5a, 0x8c, 0xf6, 0x5b, 0x9c, 0x02, 0xec, 0x02, 0x31, 0x02, 0x78, 0x20, 0x72,
	0x29, 0x4f, 0x8f, 0xf0, 0x0c, 0xc6, 0x05, 0x2f, 0xc5, 0x9a, 0x37, 0xc8, 0xd3, 0xc8, 0x06, 0x6c,
	0xd0, 0x60, 0xea, 0x7f, 0xb8, 0xe9, 0x7c, 0x0e, 0xcd, 0xb0, 0xee, 0xff, 0x02, 0x00, 0x00, 0xff,
	0xff, 0x7e, 0x7e, 0x28, 0xc4, 0xce, 0x01, 0x00, 0x00,
}
