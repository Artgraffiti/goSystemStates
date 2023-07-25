// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.12.4
// source: proto/gss.proto

package proto

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

type Empty struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Empty) Reset() {
	*x = Empty{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gss_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Empty) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Empty) ProtoMessage() {}

func (x *Empty) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gss_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Empty.ProtoReflect.Descriptor instead.
func (*Empty) Descriptor() ([]byte, []int) {
	return file_proto_gss_proto_rawDescGZIP(), []int{0}
}

type UsersMetrics struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data map[string]*MetricStorageArr `protobuf:"bytes,1,rep,name=data,proto3" json:"data,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *UsersMetrics) Reset() {
	*x = UsersMetrics{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gss_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UsersMetrics) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UsersMetrics) ProtoMessage() {}

func (x *UsersMetrics) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gss_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UsersMetrics.ProtoReflect.Descriptor instead.
func (*UsersMetrics) Descriptor() ([]byte, []int) {
	return file_proto_gss_proto_rawDescGZIP(), []int{1}
}

func (x *UsersMetrics) GetData() map[string]*MetricStorageArr {
	if x != nil {
		return x.Data
	}
	return nil
}

type MetricStorageArr struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	DataArr []*MetricStorage `protobuf:"bytes,1,rep,name=dataArr,proto3" json:"dataArr,omitempty"`
}

func (x *MetricStorageArr) Reset() {
	*x = MetricStorageArr{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gss_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricStorageArr) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricStorageArr) ProtoMessage() {}

func (x *MetricStorageArr) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gss_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricStorageArr.ProtoReflect.Descriptor instead.
func (*MetricStorageArr) Descriptor() ([]byte, []int) {
	return file_proto_gss_proto_rawDescGZIP(), []int{2}
}

func (x *MetricStorageArr) GetDataArr() []*MetricStorage {
	if x != nil {
		return x.DataArr
	}
	return nil
}

type UserMetricStorage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID          string         `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
	MetricStorage *MetricStorage `protobuf:"bytes,2,opt,name=MetricStorage,proto3" json:"MetricStorage,omitempty"`
}

func (x *UserMetricStorage) Reset() {
	*x = UserMetricStorage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gss_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserMetricStorage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserMetricStorage) ProtoMessage() {}

func (x *UserMetricStorage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gss_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserMetricStorage.ProtoReflect.Descriptor instead.
func (*UserMetricStorage) Descriptor() ([]byte, []int) {
	return file_proto_gss_proto_rawDescGZIP(), []int{3}
}

func (x *UserMetricStorage) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

func (x *UserMetricStorage) GetMetricStorage() *MetricStorage {
	if x != nil {
		return x.MetricStorage
	}
	return nil
}

type MetricStorage struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Timestamp        int64              `protobuf:"varint,1,opt,name=timestamp,proto3" json:"timestamp,omitempty"`
	MetricMapUint32  map[string]uint32  `protobuf:"bytes,2,rep,name=MetricMapUint32,proto3" json:"MetricMapUint32,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MetricMapUint64  map[string]uint64  `protobuf:"bytes,3,rep,name=MetricMapUint64,proto3" json:"MetricMapUint64,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	MetricMapFloat64 map[string]float64 `protobuf:"bytes,4,rep,name=MetricMapFloat64,proto3" json:"MetricMapFloat64,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"fixed64,2,opt,name=value,proto3"`
}

func (x *MetricStorage) Reset() {
	*x = MetricStorage{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gss_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricStorage) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricStorage) ProtoMessage() {}

func (x *MetricStorage) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gss_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricStorage.ProtoReflect.Descriptor instead.
func (*MetricStorage) Descriptor() ([]byte, []int) {
	return file_proto_gss_proto_rawDescGZIP(), []int{4}
}

func (x *MetricStorage) GetTimestamp() int64 {
	if x != nil {
		return x.Timestamp
	}
	return 0
}

func (x *MetricStorage) GetMetricMapUint32() map[string]uint32 {
	if x != nil {
		return x.MetricMapUint32
	}
	return nil
}

func (x *MetricStorage) GetMetricMapUint64() map[string]uint64 {
	if x != nil {
		return x.MetricMapUint64
	}
	return nil
}

func (x *MetricStorage) GetMetricMapFloat64() map[string]float64 {
	if x != nil {
		return x.MetricMapFloat64
	}
	return nil
}

type MetricsByUUIDRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UUID string `protobuf:"bytes,1,opt,name=UUID,proto3" json:"UUID,omitempty"`
}

func (x *MetricsByUUIDRequest) Reset() {
	*x = MetricsByUUIDRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gss_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricsByUUIDRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsByUUIDRequest) ProtoMessage() {}

func (x *MetricsByUUIDRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gss_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsByUUIDRequest.ProtoReflect.Descriptor instead.
func (*MetricsByUUIDRequest) Descriptor() ([]byte, []int) {
	return file_proto_gss_proto_rawDescGZIP(), []int{5}
}

func (x *MetricsByUUIDRequest) GetUUID() string {
	if x != nil {
		return x.UUID
	}
	return ""
}

type MetricsByUUIDResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MetricsArray []*MetricStorage `protobuf:"bytes,1,rep,name=metricsArray,proto3" json:"metricsArray,omitempty"`
}

func (x *MetricsByUUIDResponse) Reset() {
	*x = MetricsByUUIDResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_gss_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MetricsByUUIDResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MetricsByUUIDResponse) ProtoMessage() {}

func (x *MetricsByUUIDResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_gss_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MetricsByUUIDResponse.ProtoReflect.Descriptor instead.
func (*MetricsByUUIDResponse) Descriptor() ([]byte, []int) {
	return file_proto_gss_proto_rawDescGZIP(), []int{6}
}

func (x *MetricsByUUIDResponse) GetMetricsArray() []*MetricStorage {
	if x != nil {
		return x.MetricsArray
	}
	return nil
}

var File_proto_gss_proto protoreflect.FileDescriptor

var file_proto_gss_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x67, 0x73, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x03, 0x67, 0x73, 0x73, 0x22, 0x07, 0x0a, 0x05, 0x45, 0x6d, 0x70, 0x74, 0x79, 0x22,
	0x8f, 0x01, 0x0a, 0x0c, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x12, 0x2f, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x1b,
	0x2e, 0x67, 0x73, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x73, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x73, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x04, 0x64, 0x61, 0x74,
	0x61, 0x1a, 0x4e, 0x0a, 0x09, 0x44, 0x61, 0x74, 0x61, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x2b, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x15, 0x2e, 0x67, 0x73, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x41, 0x72, 0x72, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0x40, 0x0a, 0x10, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72, 0x61,
	0x67, 0x65, 0x41, 0x72, 0x72, 0x12, 0x2c, 0x0a, 0x07, 0x64, 0x61, 0x74, 0x61, 0x41, 0x72, 0x72,
	0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x73, 0x73, 0x2e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x07, 0x64, 0x61, 0x74, 0x61,
	0x41, 0x72, 0x72, 0x22, 0x61, 0x0a, 0x11, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x12, 0x38, 0x0a, 0x0d,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x73, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x0d, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x22, 0xf6, 0x03, 0x0a, 0x0d, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x74, 0x69, 0x6d, 0x65,
	0x73, 0x74, 0x61, 0x6d, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x74, 0x69, 0x6d,
	0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x12, 0x51, 0x0a, 0x0f, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x4d, 0x61, 0x70, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x27, 0x2e, 0x67, 0x73, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72,
	0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4d, 0x61, 0x70, 0x55, 0x69, 0x6e,
	0x74, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x4d, 0x61, 0x70, 0x55, 0x69, 0x6e, 0x74, 0x33, 0x32, 0x12, 0x51, 0x0a, 0x0f, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x4d, 0x61, 0x70, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x18, 0x03, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x67, 0x73, 0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53,
	0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4d, 0x61, 0x70,
	0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0f, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x4d, 0x61, 0x70, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x12, 0x54, 0x0a, 0x10,
	0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4d, 0x61, 0x70, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34,
	0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x28, 0x2e, 0x67, 0x73, 0x73, 0x2e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x4d, 0x61, 0x70, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79,
	0x52, 0x10, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4d, 0x61, 0x70, 0x46, 0x6c, 0x6f, 0x61, 0x74,
	0x36, 0x34, 0x1a, 0x42, 0x0a, 0x14, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x4d, 0x61, 0x70, 0x55,
	0x69, 0x6e, 0x74, 0x33, 0x32, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x42, 0x0a, 0x14, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63,
	0x4d, 0x61, 0x70, 0x55, 0x69, 0x6e, 0x74, 0x36, 0x34, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10,
	0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79,
	0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x04, 0x52,
	0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x1a, 0x43, 0x0a, 0x15, 0x4d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x4d, 0x61, 0x70, 0x46, 0x6c, 0x6f, 0x61, 0x74, 0x36, 0x34, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x01, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22,
	0x2a, 0x0a, 0x14, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44,
	0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x55, 0x55, 0x49, 0x44, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x55, 0x55, 0x49, 0x44, 0x22, 0x4f, 0x0a, 0x15, 0x4d,
	0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x41,
	0x72, 0x72, 0x61, 0x79, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x67, 0x73, 0x73,
	0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x52, 0x0c,
	0x6d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x41, 0x72, 0x72, 0x61, 0x79, 0x32, 0xc5, 0x01, 0x0a,
	0x0e, 0x67, 0x6f, 0x53, 0x79, 0x73, 0x74, 0x65, 0x6d, 0x53, 0x74, 0x61, 0x74, 0x65, 0x73, 0x12,
	0x35, 0x0a, 0x0d, 0x75, 0x70, 0x6c, 0x6f, 0x61, 0x64, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73,
	0x12, 0x16, 0x2e, 0x67, 0x73, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4d, 0x65, 0x74, 0x72, 0x69,
	0x63, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x1a, 0x0a, 0x2e, 0x67, 0x73, 0x73, 0x2e, 0x45,
	0x6d, 0x70, 0x74, 0x79, 0x22, 0x00, 0x12, 0x32, 0x0a, 0x0f, 0x67, 0x65, 0x74, 0x55, 0x73, 0x65,
	0x72, 0x73, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x12, 0x0a, 0x2e, 0x67, 0x73, 0x73, 0x2e,
	0x45, 0x6d, 0x70, 0x74, 0x79, 0x1a, 0x11, 0x2e, 0x67, 0x73, 0x73, 0x2e, 0x55, 0x73, 0x65, 0x72,
	0x73, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x22, 0x00, 0x12, 0x48, 0x0a, 0x0d, 0x6d, 0x65,
	0x74, 0x72, 0x69, 0x63, 0x73, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x12, 0x19, 0x2e, 0x67, 0x73,
	0x73, 0x2e, 0x4d, 0x65, 0x74, 0x72, 0x69, 0x63, 0x73, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1a, 0x2e, 0x67, 0x73, 0x73, 0x2e, 0x4d, 0x65, 0x74,
	0x72, 0x69, 0x63, 0x73, 0x42, 0x79, 0x55, 0x55, 0x49, 0x44, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x0b, 0x5a, 0x09, 0x47, 0x53, 0x53, 0x2f, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_gss_proto_rawDescOnce sync.Once
	file_proto_gss_proto_rawDescData = file_proto_gss_proto_rawDesc
)

func file_proto_gss_proto_rawDescGZIP() []byte {
	file_proto_gss_proto_rawDescOnce.Do(func() {
		file_proto_gss_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_gss_proto_rawDescData)
	})
	return file_proto_gss_proto_rawDescData
}

var file_proto_gss_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_proto_gss_proto_goTypes = []interface{}{
	(*Empty)(nil),                 // 0: gss.Empty
	(*UsersMetrics)(nil),          // 1: gss.UsersMetrics
	(*MetricStorageArr)(nil),      // 2: gss.MetricStorageArr
	(*UserMetricStorage)(nil),     // 3: gss.UserMetricStorage
	(*MetricStorage)(nil),         // 4: gss.MetricStorage
	(*MetricsByUUIDRequest)(nil),  // 5: gss.MetricsByUUIDRequest
	(*MetricsByUUIDResponse)(nil), // 6: gss.MetricsByUUIDResponse
	nil,                           // 7: gss.UsersMetrics.DataEntry
	nil,                           // 8: gss.MetricStorage.MetricMapUint32Entry
	nil,                           // 9: gss.MetricStorage.MetricMapUint64Entry
	nil,                           // 10: gss.MetricStorage.MetricMapFloat64Entry
}
var file_proto_gss_proto_depIdxs = []int32{
	7,  // 0: gss.UsersMetrics.data:type_name -> gss.UsersMetrics.DataEntry
	4,  // 1: gss.MetricStorageArr.dataArr:type_name -> gss.MetricStorage
	4,  // 2: gss.UserMetricStorage.MetricStorage:type_name -> gss.MetricStorage
	8,  // 3: gss.MetricStorage.MetricMapUint32:type_name -> gss.MetricStorage.MetricMapUint32Entry
	9,  // 4: gss.MetricStorage.MetricMapUint64:type_name -> gss.MetricStorage.MetricMapUint64Entry
	10, // 5: gss.MetricStorage.MetricMapFloat64:type_name -> gss.MetricStorage.MetricMapFloat64Entry
	4,  // 6: gss.MetricsByUUIDResponse.metricsArray:type_name -> gss.MetricStorage
	2,  // 7: gss.UsersMetrics.DataEntry.value:type_name -> gss.MetricStorageArr
	3,  // 8: gss.goSystemStates.uploadMetrics:input_type -> gss.UserMetricStorage
	0,  // 9: gss.goSystemStates.getUsersMetrics:input_type -> gss.Empty
	5,  // 10: gss.goSystemStates.metricsByUUID:input_type -> gss.MetricsByUUIDRequest
	0,  // 11: gss.goSystemStates.uploadMetrics:output_type -> gss.Empty
	1,  // 12: gss.goSystemStates.getUsersMetrics:output_type -> gss.UsersMetrics
	6,  // 13: gss.goSystemStates.metricsByUUID:output_type -> gss.MetricsByUUIDResponse
	11, // [11:14] is the sub-list for method output_type
	8,  // [8:11] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_proto_gss_proto_init() }
func file_proto_gss_proto_init() {
	if File_proto_gss_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_gss_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Empty); i {
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
		file_proto_gss_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UsersMetrics); i {
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
		file_proto_gss_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricStorageArr); i {
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
		file_proto_gss_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserMetricStorage); i {
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
		file_proto_gss_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricStorage); i {
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
		file_proto_gss_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricsByUUIDRequest); i {
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
		file_proto_gss_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MetricsByUUIDResponse); i {
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
			RawDescriptor: file_proto_gss_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_gss_proto_goTypes,
		DependencyIndexes: file_proto_gss_proto_depIdxs,
		MessageInfos:      file_proto_gss_proto_msgTypes,
	}.Build()
	File_proto_gss_proto = out.File
	file_proto_gss_proto_rawDesc = nil
	file_proto_gss_proto_goTypes = nil
	file_proto_gss_proto_depIdxs = nil
}
