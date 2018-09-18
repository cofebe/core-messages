// Code generated by protoc-gen-go. DO NOT EDIT.
// source: events/recording.proto

package events

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type RecordingImageExtracted struct {
	JobId                string   `protobuf:"bytes,10,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	TaskId               string   `protobuf:"bytes,11,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	SampleIndex          int32    `protobuf:"varint,12,opt,name=sample_index,json=sampleIndex,proto3" json:"sample_index,omitempty"`
	SampleRatePerSec     float64  `protobuf:"fixed64,13,opt,name=sample_rate_per_sec,json=sampleRatePerSec,proto3" json:"sample_rate_per_sec,omitempty"`
	Uri                  string   `protobuf:"bytes,14,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordingImageExtracted) Reset()         { *m = RecordingImageExtracted{} }
func (m *RecordingImageExtracted) String() string { return proto.CompactTextString(m) }
func (*RecordingImageExtracted) ProtoMessage()    {}
func (*RecordingImageExtracted) Descriptor() ([]byte, []int) {
	return fileDescriptor_recording_cad610edf16006da, []int{0}
}
func (m *RecordingImageExtracted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordingImageExtracted.Unmarshal(m, b)
}
func (m *RecordingImageExtracted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordingImageExtracted.Marshal(b, m, deterministic)
}
func (dst *RecordingImageExtracted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordingImageExtracted.Merge(dst, src)
}
func (m *RecordingImageExtracted) XXX_Size() int {
	return xxx_messageInfo_RecordingImageExtracted.Size(m)
}
func (m *RecordingImageExtracted) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordingImageExtracted.DiscardUnknown(m)
}

var xxx_messageInfo_RecordingImageExtracted proto.InternalMessageInfo

func (m *RecordingImageExtracted) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *RecordingImageExtracted) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *RecordingImageExtracted) GetSampleIndex() int32 {
	if m != nil {
		return m.SampleIndex
	}
	return 0
}

func (m *RecordingImageExtracted) GetSampleRatePerSec() float64 {
	if m != nil {
		return m.SampleRatePerSec
	}
	return 0
}

func (m *RecordingImageExtracted) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

type RecordingVideoFragmentCaptured struct {
	JobId                string   `protobuf:"bytes,10,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	TaskId               string   `protobuf:"bytes,11,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	SampleIndex          int32    `protobuf:"varint,12,opt,name=sample_index,json=sampleIndex,proto3" json:"sample_index,omitempty"`
	RelativeStartTimeMs  int32    `protobuf:"varint,13,opt,name=relative_start_time_ms,json=relativeStartTimeMs,proto3" json:"relative_start_time_ms,omitempty"`
	RelativeEndTimeMs    int32    `protobuf:"varint,14,opt,name=relative_end_time_ms,json=relativeEndTimeMs,proto3" json:"relative_end_time_ms,omitempty"`
	Uri                  string   `protobuf:"bytes,15,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordingVideoFragmentCaptured) Reset()         { *m = RecordingVideoFragmentCaptured{} }
func (m *RecordingVideoFragmentCaptured) String() string { return proto.CompactTextString(m) }
func (*RecordingVideoFragmentCaptured) ProtoMessage()    {}
func (*RecordingVideoFragmentCaptured) Descriptor() ([]byte, []int) {
	return fileDescriptor_recording_cad610edf16006da, []int{1}
}
func (m *RecordingVideoFragmentCaptured) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordingVideoFragmentCaptured.Unmarshal(m, b)
}
func (m *RecordingVideoFragmentCaptured) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordingVideoFragmentCaptured.Marshal(b, m, deterministic)
}
func (dst *RecordingVideoFragmentCaptured) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordingVideoFragmentCaptured.Merge(dst, src)
}
func (m *RecordingVideoFragmentCaptured) XXX_Size() int {
	return xxx_messageInfo_RecordingVideoFragmentCaptured.Size(m)
}
func (m *RecordingVideoFragmentCaptured) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordingVideoFragmentCaptured.DiscardUnknown(m)
}

var xxx_messageInfo_RecordingVideoFragmentCaptured proto.InternalMessageInfo

func (m *RecordingVideoFragmentCaptured) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *RecordingVideoFragmentCaptured) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *RecordingVideoFragmentCaptured) GetSampleIndex() int32 {
	if m != nil {
		return m.SampleIndex
	}
	return 0
}

func (m *RecordingVideoFragmentCaptured) GetRelativeStartTimeMs() int32 {
	if m != nil {
		return m.RelativeStartTimeMs
	}
	return 0
}

func (m *RecordingVideoFragmentCaptured) GetRelativeEndTimeMs() int32 {
	if m != nil {
		return m.RelativeEndTimeMs
	}
	return 0
}

func (m *RecordingVideoFragmentCaptured) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

type RecordingAudioFragmentCaptured struct {
	JobId                string   `protobuf:"bytes,10,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	TaskId               string   `protobuf:"bytes,11,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	SampleIndex          int32    `protobuf:"varint,12,opt,name=sample_index,json=sampleIndex,proto3" json:"sample_index,omitempty"`
	RelativeStartTimeMs  int32    `protobuf:"varint,13,opt,name=relative_start_time_ms,json=relativeStartTimeMs,proto3" json:"relative_start_time_ms,omitempty"`
	RelativeEndTimeMs    int32    `protobuf:"varint,14,opt,name=relative_end_time_ms,json=relativeEndTimeMs,proto3" json:"relative_end_time_ms,omitempty"`
	Uri                  string   `protobuf:"bytes,15,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordingAudioFragmentCaptured) Reset()         { *m = RecordingAudioFragmentCaptured{} }
func (m *RecordingAudioFragmentCaptured) String() string { return proto.CompactTextString(m) }
func (*RecordingAudioFragmentCaptured) ProtoMessage()    {}
func (*RecordingAudioFragmentCaptured) Descriptor() ([]byte, []int) {
	return fileDescriptor_recording_cad610edf16006da, []int{2}
}
func (m *RecordingAudioFragmentCaptured) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordingAudioFragmentCaptured.Unmarshal(m, b)
}
func (m *RecordingAudioFragmentCaptured) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordingAudioFragmentCaptured.Marshal(b, m, deterministic)
}
func (dst *RecordingAudioFragmentCaptured) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordingAudioFragmentCaptured.Merge(dst, src)
}
func (m *RecordingAudioFragmentCaptured) XXX_Size() int {
	return xxx_messageInfo_RecordingAudioFragmentCaptured.Size(m)
}
func (m *RecordingAudioFragmentCaptured) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordingAudioFragmentCaptured.DiscardUnknown(m)
}

var xxx_messageInfo_RecordingAudioFragmentCaptured proto.InternalMessageInfo

func (m *RecordingAudioFragmentCaptured) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *RecordingAudioFragmentCaptured) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *RecordingAudioFragmentCaptured) GetSampleIndex() int32 {
	if m != nil {
		return m.SampleIndex
	}
	return 0
}

func (m *RecordingAudioFragmentCaptured) GetRelativeStartTimeMs() int32 {
	if m != nil {
		return m.RelativeStartTimeMs
	}
	return 0
}

func (m *RecordingAudioFragmentCaptured) GetRelativeEndTimeMs() int32 {
	if m != nil {
		return m.RelativeEndTimeMs
	}
	return 0
}

func (m *RecordingAudioFragmentCaptured) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

type RecordingCompleted struct {
	JobId                string   `protobuf:"bytes,10,opt,name=job_id,json=jobId,proto3" json:"job_id,omitempty"`
	TaskId               string   `protobuf:"bytes,11,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	RecordingId          string   `protobuf:"bytes,12,opt,name=recording_id,json=recordingId,proto3" json:"recording_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordingCompleted) Reset()         { *m = RecordingCompleted{} }
func (m *RecordingCompleted) String() string { return proto.CompactTextString(m) }
func (*RecordingCompleted) ProtoMessage()    {}
func (*RecordingCompleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_recording_cad610edf16006da, []int{3}
}
func (m *RecordingCompleted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordingCompleted.Unmarshal(m, b)
}
func (m *RecordingCompleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordingCompleted.Marshal(b, m, deterministic)
}
func (dst *RecordingCompleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordingCompleted.Merge(dst, src)
}
func (m *RecordingCompleted) XXX_Size() int {
	return xxx_messageInfo_RecordingCompleted.Size(m)
}
func (m *RecordingCompleted) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordingCompleted.DiscardUnknown(m)
}

var xxx_messageInfo_RecordingCompleted proto.InternalMessageInfo

func (m *RecordingCompleted) GetJobId() string {
	if m != nil {
		return m.JobId
	}
	return ""
}

func (m *RecordingCompleted) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *RecordingCompleted) GetRecordingId() string {
	if m != nil {
		return m.RecordingId
	}
	return ""
}

type RecordingAssetCreated struct {
	AssetId              string   `protobuf:"bytes,10,opt,name=asset_id,json=assetId,proto3" json:"asset_id,omitempty"`
	TaskId               string   `protobuf:"bytes,11,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	RecordingId          string   `protobuf:"bytes,12,opt,name=recording_id,json=recordingId,proto3" json:"recording_id,omitempty"`
	AssetType            string   `protobuf:"bytes,13,opt,name=asset_type,json=assetType,proto3" json:"asset_type,omitempty"`
	Uri                  string   `protobuf:"bytes,14,opt,name=uri,proto3" json:"uri,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordingAssetCreated) Reset()         { *m = RecordingAssetCreated{} }
func (m *RecordingAssetCreated) String() string { return proto.CompactTextString(m) }
func (*RecordingAssetCreated) ProtoMessage()    {}
func (*RecordingAssetCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_recording_cad610edf16006da, []int{4}
}
func (m *RecordingAssetCreated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordingAssetCreated.Unmarshal(m, b)
}
func (m *RecordingAssetCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordingAssetCreated.Marshal(b, m, deterministic)
}
func (dst *RecordingAssetCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordingAssetCreated.Merge(dst, src)
}
func (m *RecordingAssetCreated) XXX_Size() int {
	return xxx_messageInfo_RecordingAssetCreated.Size(m)
}
func (m *RecordingAssetCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordingAssetCreated.DiscardUnknown(m)
}

var xxx_messageInfo_RecordingAssetCreated proto.InternalMessageInfo

func (m *RecordingAssetCreated) GetAssetId() string {
	if m != nil {
		return m.AssetId
	}
	return ""
}

func (m *RecordingAssetCreated) GetTaskId() string {
	if m != nil {
		return m.TaskId
	}
	return ""
}

func (m *RecordingAssetCreated) GetRecordingId() string {
	if m != nil {
		return m.RecordingId
	}
	return ""
}

func (m *RecordingAssetCreated) GetAssetType() string {
	if m != nil {
		return m.AssetType
	}
	return ""
}

func (m *RecordingAssetCreated) GetUri() string {
	if m != nil {
		return m.Uri
	}
	return ""
}

type RecordingCreated struct {
	RecordingId          string   `protobuf:"bytes,10,opt,name=recording_id,json=recordingId,proto3" json:"recording_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordingCreated) Reset()         { *m = RecordingCreated{} }
func (m *RecordingCreated) String() string { return proto.CompactTextString(m) }
func (*RecordingCreated) ProtoMessage()    {}
func (*RecordingCreated) Descriptor() ([]byte, []int) {
	return fileDescriptor_recording_cad610edf16006da, []int{5}
}
func (m *RecordingCreated) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordingCreated.Unmarshal(m, b)
}
func (m *RecordingCreated) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordingCreated.Marshal(b, m, deterministic)
}
func (dst *RecordingCreated) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordingCreated.Merge(dst, src)
}
func (m *RecordingCreated) XXX_Size() int {
	return xxx_messageInfo_RecordingCreated.Size(m)
}
func (m *RecordingCreated) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordingCreated.DiscardUnknown(m)
}

var xxx_messageInfo_RecordingCreated proto.InternalMessageInfo

func (m *RecordingCreated) GetRecordingId() string {
	if m != nil {
		return m.RecordingId
	}
	return ""
}

type RecordingInserted struct {
	RecordingId          string   `protobuf:"bytes,10,opt,name=recording_id,json=recordingId,proto3" json:"recording_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordingInserted) Reset()         { *m = RecordingInserted{} }
func (m *RecordingInserted) String() string { return proto.CompactTextString(m) }
func (*RecordingInserted) ProtoMessage()    {}
func (*RecordingInserted) Descriptor() ([]byte, []int) {
	return fileDescriptor_recording_cad610edf16006da, []int{6}
}
func (m *RecordingInserted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordingInserted.Unmarshal(m, b)
}
func (m *RecordingInserted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordingInserted.Marshal(b, m, deterministic)
}
func (dst *RecordingInserted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordingInserted.Merge(dst, src)
}
func (m *RecordingInserted) XXX_Size() int {
	return xxx_messageInfo_RecordingInserted.Size(m)
}
func (m *RecordingInserted) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordingInserted.DiscardUnknown(m)
}

var xxx_messageInfo_RecordingInserted proto.InternalMessageInfo

func (m *RecordingInserted) GetRecordingId() string {
	if m != nil {
		return m.RecordingId
	}
	return ""
}

type RecordingDeleted struct {
	RecordingId          string   `protobuf:"bytes,10,opt,name=recording_id,json=recordingId,proto3" json:"recording_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordingDeleted) Reset()         { *m = RecordingDeleted{} }
func (m *RecordingDeleted) String() string { return proto.CompactTextString(m) }
func (*RecordingDeleted) ProtoMessage()    {}
func (*RecordingDeleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_recording_cad610edf16006da, []int{7}
}
func (m *RecordingDeleted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordingDeleted.Unmarshal(m, b)
}
func (m *RecordingDeleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordingDeleted.Marshal(b, m, deterministic)
}
func (dst *RecordingDeleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordingDeleted.Merge(dst, src)
}
func (m *RecordingDeleted) XXX_Size() int {
	return xxx_messageInfo_RecordingDeleted.Size(m)
}
func (m *RecordingDeleted) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordingDeleted.DiscardUnknown(m)
}

var xxx_messageInfo_RecordingDeleted proto.InternalMessageInfo

func (m *RecordingDeleted) GetRecordingId() string {
	if m != nil {
		return m.RecordingId
	}
	return ""
}

type RecordingCognitionCompleted struct {
	RecordingId          string   `protobuf:"bytes,10,opt,name=recording_id,json=recordingId,proto3" json:"recording_id,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *RecordingCognitionCompleted) Reset()         { *m = RecordingCognitionCompleted{} }
func (m *RecordingCognitionCompleted) String() string { return proto.CompactTextString(m) }
func (*RecordingCognitionCompleted) ProtoMessage()    {}
func (*RecordingCognitionCompleted) Descriptor() ([]byte, []int) {
	return fileDescriptor_recording_cad610edf16006da, []int{8}
}
func (m *RecordingCognitionCompleted) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_RecordingCognitionCompleted.Unmarshal(m, b)
}
func (m *RecordingCognitionCompleted) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_RecordingCognitionCompleted.Marshal(b, m, deterministic)
}
func (dst *RecordingCognitionCompleted) XXX_Merge(src proto.Message) {
	xxx_messageInfo_RecordingCognitionCompleted.Merge(dst, src)
}
func (m *RecordingCognitionCompleted) XXX_Size() int {
	return xxx_messageInfo_RecordingCognitionCompleted.Size(m)
}
func (m *RecordingCognitionCompleted) XXX_DiscardUnknown() {
	xxx_messageInfo_RecordingCognitionCompleted.DiscardUnknown(m)
}

var xxx_messageInfo_RecordingCognitionCompleted proto.InternalMessageInfo

func (m *RecordingCognitionCompleted) GetRecordingId() string {
	if m != nil {
		return m.RecordingId
	}
	return ""
}

func init() {
	proto.RegisterType((*RecordingImageExtracted)(nil), "events.RecordingImageExtracted")
	proto.RegisterType((*RecordingVideoFragmentCaptured)(nil), "events.RecordingVideoFragmentCaptured")
	proto.RegisterType((*RecordingAudioFragmentCaptured)(nil), "events.RecordingAudioFragmentCaptured")
	proto.RegisterType((*RecordingCompleted)(nil), "events.RecordingCompleted")
	proto.RegisterType((*RecordingAssetCreated)(nil), "events.RecordingAssetCreated")
	proto.RegisterType((*RecordingCreated)(nil), "events.RecordingCreated")
	proto.RegisterType((*RecordingInserted)(nil), "events.RecordingInserted")
	proto.RegisterType((*RecordingDeleted)(nil), "events.RecordingDeleted")
	proto.RegisterType((*RecordingCognitionCompleted)(nil), "events.RecordingCognitionCompleted")
}

func init() { proto.RegisterFile("events/recording.proto", fileDescriptor_recording_cad610edf16006da) }

var fileDescriptor_recording_cad610edf16006da = []byte{
	// 418 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xdc, 0x54, 0x4d, 0x8f, 0xd3, 0x30,
	0x10, 0x55, 0x84, 0x9a, 0xdd, 0x4e, 0xcb, 0xd2, 0xf5, 0xb2, 0xbb, 0x41, 0x08, 0x54, 0x7a, 0xea,
	0x05, 0xf6, 0xb0, 0x82, 0x33, 0x50, 0x8a, 0x94, 0x03, 0x12, 0x4a, 0x2b, 0x0e, 0x5c, 0x22, 0x37,
	0x1e, 0x45, 0x2e, 0x8d, 0x1d, 0xd9, 0xd3, 0xaa, 0xfd, 0x35, 0xfc, 0x02, 0x7e, 0x19, 0x7f, 0x02,
	0xd9, 0x49, 0x93, 0x8a, 0x0f, 0x51, 0x81, 0xb8, 0x70, 0xeb, 0xbc, 0xd7, 0xe7, 0x37, 0x6f, 0x3c,
	0x31, 0x5c, 0xe1, 0x06, 0x15, 0xd9, 0x1b, 0x83, 0x99, 0x36, 0x42, 0xaa, 0xfc, 0x59, 0x69, 0x34,
	0x69, 0x16, 0x56, 0xf8, 0xe8, 0x4b, 0x00, 0xd7, 0xc9, 0x9e, 0x8b, 0x0b, 0x9e, 0xe3, 0x74, 0x4b,
	0x86, 0x67, 0x84, 0x82, 0x5d, 0x42, 0xb8, 0xd4, 0x8b, 0x54, 0x8a, 0x08, 0x86, 0xc1, 0xb8, 0x9b,
	0x74, 0x96, 0x7a, 0x11, 0x0b, 0x76, 0x0d, 0x27, 0xc4, 0xed, 0x27, 0x87, 0xf7, 0x3c, 0x1e, 0xba,
	0x32, 0x16, 0xec, 0x09, 0xf4, 0x2d, 0x2f, 0xca, 0x15, 0xa6, 0x52, 0x09, 0xdc, 0x46, 0xfd, 0x61,
	0x30, 0xee, 0x24, 0xbd, 0x0a, 0x8b, 0x1d, 0xc4, 0x9e, 0xc2, 0x45, 0xfd, 0x17, 0xc3, 0x09, 0xd3,
	0x12, 0x4d, 0x6a, 0x31, 0x8b, 0xee, 0x0e, 0x83, 0x71, 0x90, 0x0c, 0x2a, 0x2a, 0xe1, 0x84, 0xef,
	0xd1, 0xcc, 0x30, 0x63, 0x03, 0xb8, 0xb3, 0x36, 0x32, 0x3a, 0xf3, 0x36, 0xee, 0xe7, 0xe8, 0x6b,
	0x00, 0x8f, 0x9b, 0x7e, 0x3f, 0x48, 0x81, 0xfa, 0xad, 0xe1, 0x79, 0x81, 0x8a, 0x26, 0xbc, 0xa4,
	0xb5, 0xf9, 0x37, 0x6d, 0xdf, 0xc2, 0x95, 0xc1, 0x15, 0x27, 0xb9, 0xc1, 0xd4, 0x12, 0x37, 0x94,
	0x92, 0x2c, 0x30, 0x2d, 0xac, 0xef, 0xbc, 0x93, 0x5c, 0xec, 0xd9, 0x99, 0x23, 0xe7, 0xb2, 0xc0,
	0x77, 0x96, 0xdd, 0xc0, 0xfd, 0x46, 0x84, 0x4a, 0x34, 0x92, 0x33, 0x2f, 0x39, 0xdf, 0x73, 0x53,
	0x25, 0x6a, 0x41, 0x9d, 0xf6, 0xde, 0x2f, 0xd2, 0xbe, 0x5a, 0x0b, 0xf9, 0x5f, 0xa7, 0xcd, 0x81,
	0x35, 0x61, 0x27, 0xda, 0xf5, 0xf3, 0x87, 0x5b, 0xd8, 0x2c, 0xbb, 0x63, 0xfb, 0x9e, 0xed, 0x35,
	0x58, 0x2c, 0x46, 0x9f, 0x03, 0xb8, 0x6c, 0xc7, 0x6a, 0x2d, 0xd2, 0xc4, 0x20, 0x77, 0x66, 0x0f,
	0xe0, 0x94, 0xbb, 0xba, 0xb5, 0x3b, 0xf1, 0xf5, 0xdf, 0x19, 0xb2, 0x47, 0x00, 0xd5, 0xb1, 0xb4,
	0x2b, 0xd1, 0x4f, 0xb1, 0x9b, 0x74, 0x3d, 0x32, 0xdf, 0x95, 0xf8, 0x93, 0x35, 0x7f, 0x0e, 0x83,
	0x76, 0x14, 0x75, 0x6f, 0xdf, 0xfb, 0xc0, 0x8f, 0xc1, 0x5e, 0xc0, 0x79, 0xfb, 0x31, 0x2b, 0x8b,
	0xe6, 0x48, 0xdd, 0xa1, 0xdd, 0x1b, 0xac, 0xe6, 0x7e, 0x84, 0xec, 0x25, 0x3c, 0x3c, 0xb8, 0xb0,
	0x5c, 0x49, 0x92, 0x5a, 0xb5, 0x37, 0xf7, 0xfb, 0x13, 0x5e, 0x9f, 0x7e, 0xac, 0x1f, 0xa2, 0x45,
	0xe8, 0xdf, 0xa5, 0xdb, 0x6f, 0x01, 0x00, 0x00, 0xff, 0xff, 0xcd, 0x50, 0x39, 0xaf, 0xb1, 0x04,
	0x00, 0x00,
}
