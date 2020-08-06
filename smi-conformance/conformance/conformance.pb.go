// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.21.0
// 	protoc        v3.12.3
// source: conformance/conformance.proto

package conformance

import (
	context "context"
	proto "github.com/golang/protobuf/proto"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
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

// This is a compile-time assertion that a sufficiently up-to-date version
// of the legacy proto package is being used.
const _ = proto.ProtoPackageIsVersion4

type Request struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Annotations map[string]string `protobuf:"bytes,1,rep,name=annotations,proto3" json:"annotations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Labels      map[string]string `protobuf:"bytes,2,rep,name=labels,proto3" json:"labels,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Meshname    string            `protobuf:"bytes,3,opt,name=meshname,proto3" json:"meshname,omitempty"`
}

func (x *Request) Reset() {
	*x = Request{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conformance_conformance_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Request) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Request) ProtoMessage() {}

func (x *Request) ProtoReflect() protoreflect.Message {
	mi := &file_conformance_conformance_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Request.ProtoReflect.Descriptor instead.
func (*Request) Descriptor() ([]byte, []int) {
	return file_conformance_conformance_proto_rawDescGZIP(), []int{0}
}

func (x *Request) GetAnnotations() map[string]string {
	if x != nil {
		return x.Annotations
	}
	return nil
}

func (x *Request) GetLabels() map[string]string {
	if x != nil {
		return x.Labels
	}
	return nil
}

func (x *Request) GetMeshname() string {
	if x != nil {
		return x.Meshname
	}
	return ""
}

type SingleTestResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name       string   `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Time       string   `protobuf:"bytes,2,opt,name=time,proto3" json:"time,omitempty"`
	Assertions string   `protobuf:"bytes,3,opt,name=assertions,proto3" json:"assertions,omitempty"`
	Failure    *Failure `protobuf:"bytes,4,opt,name=failure,proto3" json:"failure,omitempty"`
}

func (x *SingleTestResult) Reset() {
	*x = SingleTestResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conformance_conformance_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SingleTestResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SingleTestResult) ProtoMessage() {}

func (x *SingleTestResult) ProtoReflect() protoreflect.Message {
	mi := &file_conformance_conformance_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SingleTestResult.ProtoReflect.Descriptor instead.
func (*SingleTestResult) Descriptor() ([]byte, []int) {
	return file_conformance_conformance_proto_rawDescGZIP(), []int{1}
}

func (x *SingleTestResult) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SingleTestResult) GetTime() string {
	if x != nil {
		return x.Time
	}
	return ""
}

func (x *SingleTestResult) GetAssertions() string {
	if x != nil {
		return x.Assertions
	}
	return ""
}

func (x *SingleTestResult) GetFailure() *Failure {
	if x != nil {
		return x.Failure
	}
	return nil
}

type Failure struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Test    string `protobuf:"bytes,1,opt,name=test,proto3" json:"test,omitempty"`
	Message string `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
}

func (x *Failure) Reset() {
	*x = Failure{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conformance_conformance_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Failure) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Failure) ProtoMessage() {}

func (x *Failure) ProtoReflect() protoreflect.Message {
	mi := &file_conformance_conformance_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Failure.ProtoReflect.Descriptor instead.
func (*Failure) Descriptor() ([]byte, []int) {
	return file_conformance_conformance_proto_rawDescGZIP(), []int{2}
}

func (x *Failure) GetTest() string {
	if x != nil {
		return x.Test
	}
	return ""
}

func (x *Failure) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Tests            string              `protobuf:"bytes,1,opt,name=tests,proto3" json:"tests,omitempty"`
	Failures         string              `protobuf:"bytes,2,opt,name=failures,proto3" json:"failures,omitempty"`
	SingleTestResult []*SingleTestResult `protobuf:"bytes,3,rep,name=singleTestResult,proto3" json:"singleTestResult,omitempty"`
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_conformance_conformance_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_conformance_conformance_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Response.ProtoReflect.Descriptor instead.
func (*Response) Descriptor() ([]byte, []int) {
	return file_conformance_conformance_proto_rawDescGZIP(), []int{3}
}

func (x *Response) GetTests() string {
	if x != nil {
		return x.Tests
	}
	return ""
}

func (x *Response) GetFailures() string {
	if x != nil {
		return x.Failures
	}
	return ""
}

func (x *Response) GetSingleTestResult() []*SingleTestResult {
	if x != nil {
		return x.SingleTestResult
	}
	return nil
}

var File_conformance_conformance_proto protoreflect.FileDescriptor

var file_conformance_conformance_proto_rawDesc = []byte{
	0x0a, 0x1d, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2f, 0x63, 0x6f,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x73, 0x6d, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65,
	0x22, 0xab, 0x02, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x4b, 0x0a, 0x0b,
	0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x29, 0x2e, 0x73, 0x6d, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x41, 0x6e, 0x6e, 0x6f,
	0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x61, 0x6e,
	0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x3c, 0x0a, 0x06, 0x6c, 0x61, 0x62,
	0x65, 0x6c, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x73, 0x6d, 0x69, 0x5f,
	0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x2e, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52,
	0x06, 0x6c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x6d, 0x65, 0x73, 0x68, 0x6e,
	0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6d, 0x65, 0x73, 0x68, 0x6e,
	0x61, 0x6d, 0x65, 0x1a, 0x3e, 0x0a, 0x10, 0x41, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x1a, 0x39, 0x0a, 0x0b, 0x4c, 0x61, 0x62, 0x65, 0x6c, 0x73, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x8e,
	0x01, 0x0a, 0x10, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x69, 0x6d, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x61,
	0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x0a, 0x61, 0x73, 0x73, 0x65, 0x72, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x32, 0x0a, 0x07, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x73,
	0x6d, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x46,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x52, 0x07, 0x66, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x22,
	0x37, 0x0a, 0x07, 0x46, 0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x74, 0x65,
	0x73, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x74, 0x65, 0x73, 0x74, 0x12, 0x18,
	0x0a, 0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x07, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x22, 0x8b, 0x01, 0x0a, 0x08, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x65, 0x73, 0x74, 0x73, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x65, 0x73, 0x74, 0x73, 0x12, 0x1a, 0x0a, 0x08, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x66,
	0x61, 0x69, 0x6c, 0x75, 0x72, 0x65, 0x73, 0x12, 0x4d, 0x0a, 0x10, 0x73, 0x69, 0x6e, 0x67, 0x6c,
	0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x03, 0x28,
	0x0b, 0x32, 0x21, 0x2e, 0x73, 0x6d, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x52, 0x10, 0x73, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x54, 0x65, 0x73, 0x74,
	0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x32, 0x54, 0x0a, 0x12, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72,
	0x6d, 0x61, 0x6e, 0x63, 0x65, 0x54, 0x65, 0x73, 0x74, 0x69, 0x6e, 0x67, 0x12, 0x3e, 0x0a, 0x07,
	0x52, 0x75, 0x6e, 0x54, 0x65, 0x73, 0x74, 0x12, 0x18, 0x2e, 0x73, 0x6d, 0x69, 0x5f, 0x63, 0x6f,
	0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x19, 0x2e, 0x73, 0x6d, 0x69, 0x5f, 0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61,
	0x6e, 0x63, 0x65, 0x2e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x19, 0x5a, 0x17,
	0x63, 0x6f, 0x6e, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x3b, 0x63, 0x6f, 0x6e, 0x66,
	0x6f, 0x72, 0x6d, 0x61, 0x6e, 0x63, 0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_conformance_conformance_proto_rawDescOnce sync.Once
	file_conformance_conformance_proto_rawDescData = file_conformance_conformance_proto_rawDesc
)

func file_conformance_conformance_proto_rawDescGZIP() []byte {
	file_conformance_conformance_proto_rawDescOnce.Do(func() {
		file_conformance_conformance_proto_rawDescData = protoimpl.X.CompressGZIP(file_conformance_conformance_proto_rawDescData)
	})
	return file_conformance_conformance_proto_rawDescData
}

var file_conformance_conformance_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_conformance_conformance_proto_goTypes = []interface{}{
	(*Request)(nil),          // 0: smi_conformance.Request
	(*SingleTestResult)(nil), // 1: smi_conformance.SingleTestResult
	(*Failure)(nil),          // 2: smi_conformance.Failure
	(*Response)(nil),         // 3: smi_conformance.Response
	nil,                      // 4: smi_conformance.Request.AnnotationsEntry
	nil,                      // 5: smi_conformance.Request.LabelsEntry
}
var file_conformance_conformance_proto_depIdxs = []int32{
	4, // 0: smi_conformance.Request.annotations:type_name -> smi_conformance.Request.AnnotationsEntry
	5, // 1: smi_conformance.Request.labels:type_name -> smi_conformance.Request.LabelsEntry
	2, // 2: smi_conformance.SingleTestResult.failure:type_name -> smi_conformance.Failure
	1, // 3: smi_conformance.Response.singleTestResult:type_name -> smi_conformance.SingleTestResult
	0, // 4: smi_conformance.conformanceTesting.RunTest:input_type -> smi_conformance.Request
	3, // 5: smi_conformance.conformanceTesting.RunTest:output_type -> smi_conformance.Response
	5, // [5:6] is the sub-list for method output_type
	4, // [4:5] is the sub-list for method input_type
	4, // [4:4] is the sub-list for extension type_name
	4, // [4:4] is the sub-list for extension extendee
	0, // [0:4] is the sub-list for field type_name
}

func init() { file_conformance_conformance_proto_init() }
func file_conformance_conformance_proto_init() {
	if File_conformance_conformance_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_conformance_conformance_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Request); i {
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
		file_conformance_conformance_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SingleTestResult); i {
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
		file_conformance_conformance_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Failure); i {
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
		file_conformance_conformance_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Response); i {
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
			RawDescriptor: file_conformance_conformance_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_conformance_conformance_proto_goTypes,
		DependencyIndexes: file_conformance_conformance_proto_depIdxs,
		MessageInfos:      file_conformance_conformance_proto_msgTypes,
	}.Build()
	File_conformance_conformance_proto = out.File
	file_conformance_conformance_proto_rawDesc = nil
	file_conformance_conformance_proto_goTypes = nil
	file_conformance_conformance_proto_depIdxs = nil
}

// Reference imports to suppress errors if they are not otherwise used.
var _ context.Context
var _ grpc.ClientConnInterface

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
const _ = grpc.SupportPackageIsVersion6

// ConformanceTestingClient is the client API for ConformanceTesting service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://godoc.org/google.golang.org/grpc#ClientConn.NewStream.
type ConformanceTestingClient interface {
	RunTest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error)
}

type conformanceTestingClient struct {
	cc grpc.ClientConnInterface
}

func NewConformanceTestingClient(cc grpc.ClientConnInterface) ConformanceTestingClient {
	return &conformanceTestingClient{cc}
}

func (c *conformanceTestingClient) RunTest(ctx context.Context, in *Request, opts ...grpc.CallOption) (*Response, error) {
	out := new(Response)
	err := c.cc.Invoke(ctx, "/smi_conformance.conformanceTesting/RunTest", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// ConformanceTestingServer is the server API for ConformanceTesting service.
type ConformanceTestingServer interface {
	RunTest(context.Context, *Request) (*Response, error)
}

// UnimplementedConformanceTestingServer can be embedded to have forward compatible implementations.
type UnimplementedConformanceTestingServer struct {
}

func (*UnimplementedConformanceTestingServer) RunTest(context.Context, *Request) (*Response, error) {
	return nil, status.Errorf(codes.Unimplemented, "method RunTest not implemented")
}

func RegisterConformanceTestingServer(s *grpc.Server, srv ConformanceTestingServer) {
	s.RegisterService(&_ConformanceTesting_serviceDesc, srv)
}

func _ConformanceTesting_RunTest_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(Request)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(ConformanceTestingServer).RunTest(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/smi_conformance.conformanceTesting/RunTest",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(ConformanceTestingServer).RunTest(ctx, req.(*Request))
	}
	return interceptor(ctx, in, info, handler)
}

var _ConformanceTesting_serviceDesc = grpc.ServiceDesc{
	ServiceName: "smi_conformance.conformanceTesting",
	HandlerType: (*ConformanceTestingServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "RunTest",
			Handler:    _ConformanceTesting_RunTest_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "conformance/conformance.proto",
}
