// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: internal/grpc/pb/grpc_pb.proto

package grpc_pb_proto

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

type RegistrateRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RegData string `protobuf:"bytes,1,opt,name=regData,proto3" json:"regData,omitempty"`
}

func (x *RegistrateRequest) Reset() {
	*x = RegistrateRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_pb_grpc_pb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrateRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrateRequest) ProtoMessage() {}

func (x *RegistrateRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_pb_grpc_pb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrateRequest.ProtoReflect.Descriptor instead.
func (*RegistrateRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpc_pb_grpc_pb_proto_rawDescGZIP(), []int{0}
}

func (x *RegistrateRequest) GetRegData() string {
	if x != nil {
		return x.RegData
	}
	return ""
}

type RegistrateResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Status    string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	ErrorText string `protobuf:"bytes,3,opt,name=error_text,json=errorText,proto3" json:"error_text,omitempty"`
}

func (x *RegistrateResponse) Reset() {
	*x = RegistrateResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_pb_grpc_pb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *RegistrateResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RegistrateResponse) ProtoMessage() {}

func (x *RegistrateResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_pb_grpc_pb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RegistrateResponse.ProtoReflect.Descriptor instead.
func (*RegistrateResponse) Descriptor() ([]byte, []int) {
	return file_internal_grpc_pb_grpc_pb_proto_rawDescGZIP(), []int{1}
}

func (x *RegistrateResponse) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *RegistrateResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *RegistrateResponse) GetErrorText() string {
	if x != nil {
		return x.ErrorText
	}
	return ""
}

type GetDataRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"` // сюда передаём _id
}

func (x *GetDataRequest) Reset() {
	*x = GetDataRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_pb_grpc_pb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataRequest) ProtoMessage() {}

func (x *GetDataRequest) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_pb_grpc_pb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataRequest.ProtoReflect.Descriptor instead.
func (*GetDataRequest) Descriptor() ([]byte, []int) {
	return file_internal_grpc_pb_grpc_pb_proto_rawDescGZIP(), []int{2}
}

func (x *GetDataRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetDataResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data      string `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Status    string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	ErrorText string `protobuf:"bytes,3,opt,name=error_text,json=errorText,proto3" json:"error_text,omitempty"`
}

func (x *GetDataResponse) Reset() {
	*x = GetDataResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_internal_grpc_pb_grpc_pb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetDataResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetDataResponse) ProtoMessage() {}

func (x *GetDataResponse) ProtoReflect() protoreflect.Message {
	mi := &file_internal_grpc_pb_grpc_pb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetDataResponse.ProtoReflect.Descriptor instead.
func (*GetDataResponse) Descriptor() ([]byte, []int) {
	return file_internal_grpc_pb_grpc_pb_proto_rawDescGZIP(), []int{3}
}

func (x *GetDataResponse) GetData() string {
	if x != nil {
		return x.Data
	}
	return ""
}

func (x *GetDataResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetDataResponse) GetErrorText() string {
	if x != nil {
		return x.ErrorText
	}
	return ""
}

var File_internal_grpc_pb_grpc_pb_proto protoreflect.FileDescriptor

var file_internal_grpc_pb_grpc_pb_proto_rawDesc = []byte{
	0x0a, 0x1e, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f,
	0x70, 0x62, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x12, 0x04, 0x67, 0x72, 0x70, 0x63, 0x22, 0x2d, 0x0a, 0x11, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74,
	0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x72,
	0x65, 0x67, 0x44, 0x61, 0x74, 0x61, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x72, 0x65,
	0x67, 0x44, 0x61, 0x74, 0x61, 0x22, 0x5b, 0x0a, 0x12, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72,
	0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x0e, 0x0a, 0x02, 0x69,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x65,
	0x78, 0x74, 0x22, 0x20, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x02, 0x69, 0x64, 0x22, 0x5c, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52,
	0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x16, 0x0a, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x12, 0x1d, 0x0a, 0x0a, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x5f, 0x74, 0x65, 0x78,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x54, 0x65,
	0x78, 0x74, 0x32, 0x99, 0x01, 0x0a, 0x10, 0x53, 0x63, 0x61, 0x6e, 0x6e, 0x65, 0x72, 0x53, 0x6d,
	0x61, 0x72, 0x74, 0x43, 0x61, 0x72, 0x64, 0x12, 0x47, 0x0a, 0x10, 0x52, 0x65, 0x67, 0x69, 0x73,
	0x74, 0x65, 0x72, 0x43, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12, 0x17, 0x2e, 0x67, 0x72,
	0x70, 0x63, 0x2e, 0x52, 0x65, 0x67, 0x69, 0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x18, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x52, 0x65, 0x67, 0x69,
	0x73, 0x74, 0x72, 0x61, 0x74, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00,
	0x12, 0x3c, 0x0a, 0x0b, 0x47, 0x65, 0x74, 0x43, 0x61, 0x72, 0x64, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x14, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74, 0x44, 0x61, 0x74, 0x61, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x67, 0x72, 0x70, 0x63, 0x2e, 0x47, 0x65, 0x74,
	0x44, 0x61, 0x74, 0x61, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x1b,
	0x5a, 0x19, 0x70, 0x6b, 0x67, 0x2f, 0x67, 0x72, 0x70, 0x63, 0x2f, 0x70, 0x62, 0x2f, 0x67, 0x72,
	0x70, 0x63, 0x5f, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_internal_grpc_pb_grpc_pb_proto_rawDescOnce sync.Once
	file_internal_grpc_pb_grpc_pb_proto_rawDescData = file_internal_grpc_pb_grpc_pb_proto_rawDesc
)

func file_internal_grpc_pb_grpc_pb_proto_rawDescGZIP() []byte {
	file_internal_grpc_pb_grpc_pb_proto_rawDescOnce.Do(func() {
		file_internal_grpc_pb_grpc_pb_proto_rawDescData = protoimpl.X.CompressGZIP(file_internal_grpc_pb_grpc_pb_proto_rawDescData)
	})
	return file_internal_grpc_pb_grpc_pb_proto_rawDescData
}

var file_internal_grpc_pb_grpc_pb_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_internal_grpc_pb_grpc_pb_proto_goTypes = []interface{}{
	(*RegistrateRequest)(nil),  // 0: grpc.RegistrateRequest
	(*RegistrateResponse)(nil), // 1: grpc.RegistrateResponse
	(*GetDataRequest)(nil),     // 2: grpc.GetDataRequest
	(*GetDataResponse)(nil),    // 3: grpc.GetDataResponse
}
var file_internal_grpc_pb_grpc_pb_proto_depIdxs = []int32{
	0, // 0: grpc.ScannerSmartCard.RegisterCardData:input_type -> grpc.RegistrateRequest
	2, // 1: grpc.ScannerSmartCard.GetCardData:input_type -> grpc.GetDataRequest
	1, // 2: grpc.ScannerSmartCard.RegisterCardData:output_type -> grpc.RegistrateResponse
	3, // 3: grpc.ScannerSmartCard.GetCardData:output_type -> grpc.GetDataResponse
	2, // [2:4] is the sub-list for method output_type
	0, // [0:2] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_internal_grpc_pb_grpc_pb_proto_init() }
func file_internal_grpc_pb_grpc_pb_proto_init() {
	if File_internal_grpc_pb_grpc_pb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_internal_grpc_pb_grpc_pb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrateRequest); i {
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
		file_internal_grpc_pb_grpc_pb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*RegistrateResponse); i {
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
		file_internal_grpc_pb_grpc_pb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataRequest); i {
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
		file_internal_grpc_pb_grpc_pb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetDataResponse); i {
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
			RawDescriptor: file_internal_grpc_pb_grpc_pb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_internal_grpc_pb_grpc_pb_proto_goTypes,
		DependencyIndexes: file_internal_grpc_pb_grpc_pb_proto_depIdxs,
		MessageInfos:      file_internal_grpc_pb_grpc_pb_proto_msgTypes,
	}.Build()
	File_internal_grpc_pb_grpc_pb_proto = out.File
	file_internal_grpc_pb_grpc_pb_proto_rawDesc = nil
	file_internal_grpc_pb_grpc_pb_proto_goTypes = nil
	file_internal_grpc_pb_grpc_pb_proto_depIdxs = nil
}
