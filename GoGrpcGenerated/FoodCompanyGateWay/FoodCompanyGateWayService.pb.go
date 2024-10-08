// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v5.28.2
// source: FoodCompanyGateWayService.proto

package FoodCompanyGateWay

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

type GetOrderStatusResponse_Status int32

const (
	GetOrderStatusResponse_Done      GetOrderStatusResponse_Status = 0
	GetOrderStatusResponse_Pending   GetOrderStatusResponse_Status = 1
	GetOrderStatusResponse_Cancelled GetOrderStatusResponse_Status = 2
)

// Enum value maps for GetOrderStatusResponse_Status.
var (
	GetOrderStatusResponse_Status_name = map[int32]string{
		0: "Done",
		1: "Pending",
		2: "Cancelled",
	}
	GetOrderStatusResponse_Status_value = map[string]int32{
		"Done":      0,
		"Pending":   1,
		"Cancelled": 2,
	}
)

func (x GetOrderStatusResponse_Status) Enum() *GetOrderStatusResponse_Status {
	p := new(GetOrderStatusResponse_Status)
	*p = x
	return p
}

func (x GetOrderStatusResponse_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (GetOrderStatusResponse_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_FoodCompanyGateWayService_proto_enumTypes[0].Descriptor()
}

func (GetOrderStatusResponse_Status) Type() protoreflect.EnumType {
	return &file_FoodCompanyGateWayService_proto_enumTypes[0]
}

func (x GetOrderStatusResponse_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use GetOrderStatusResponse_Status.Descriptor instead.
func (GetOrderStatusResponse_Status) EnumDescriptor() ([]byte, []int) {
	return file_FoodCompanyGateWayService_proto_rawDescGZIP(), []int{1, 0}
}

type GetOrderStatusRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetOrderStatusRequest) Reset() {
	*x = GetOrderStatusRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FoodCompanyGateWayService_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderStatusRequest) ProtoMessage() {}

func (x *GetOrderStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_FoodCompanyGateWayService_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderStatusRequest.ProtoReflect.Descriptor instead.
func (*GetOrderStatusRequest) Descriptor() ([]byte, []int) {
	return file_FoodCompanyGateWayService_proto_rawDescGZIP(), []int{0}
}

func (x *GetOrderStatusRequest) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

type GetOrderStatusResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Status GetOrderStatusResponse_Status `protobuf:"varint,1,opt,name=status,proto3,enum=GetOrderStatusResponse_Status" json:"status,omitempty"`
	Owner  *string                       `protobuf:"bytes,2,opt,name=owner,proto3,oneof" json:"owner,omitempty"`
}

func (x *GetOrderStatusResponse) Reset() {
	*x = GetOrderStatusResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_FoodCompanyGateWayService_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetOrderStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetOrderStatusResponse) ProtoMessage() {}

func (x *GetOrderStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_FoodCompanyGateWayService_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetOrderStatusResponse.ProtoReflect.Descriptor instead.
func (*GetOrderStatusResponse) Descriptor() ([]byte, []int) {
	return file_FoodCompanyGateWayService_proto_rawDescGZIP(), []int{1}
}

func (x *GetOrderStatusResponse) GetStatus() GetOrderStatusResponse_Status {
	if x != nil {
		return x.Status
	}
	return GetOrderStatusResponse_Done
}

func (x *GetOrderStatusResponse) GetOwner() string {
	if x != nil && x.Owner != nil {
		return *x.Owner
	}
	return ""
}

var File_FoodCompanyGateWayService_proto protoreflect.FileDescriptor

var file_FoodCompanyGateWayService_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x46, 0x6f, 0x6f, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e, 0x79, 0x47, 0x61, 0x74,
	0x65, 0x57, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x27, 0x0a, 0x15, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x22, 0xa5, 0x01, 0x0a, 0x16, 0x47,
	0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73,
	0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x36, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x1e, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x19, 0x0a,
	0x05, 0x6f, 0x77, 0x6e, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x48, 0x00, 0x52, 0x05,
	0x6f, 0x77, 0x6e, 0x65, 0x72, 0x88, 0x01, 0x01, 0x22, 0x2e, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x12, 0x08, 0x0a, 0x04, 0x44, 0x6f, 0x6e, 0x65, 0x10, 0x00, 0x12, 0x0b, 0x0a, 0x07,
	0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x10, 0x01, 0x12, 0x0d, 0x0a, 0x09, 0x43, 0x61, 0x6e,
	0x63, 0x65, 0x6c, 0x6c, 0x65, 0x64, 0x10, 0x02, 0x42, 0x08, 0x0a, 0x06, 0x5f, 0x6f, 0x77, 0x6e,
	0x65, 0x72, 0x32, 0x5e, 0x0a, 0x19, 0x46, 0x6f, 0x6f, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x47, 0x61, 0x74, 0x65, 0x57, 0x61, 0x79, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12,
	0x41, 0x0a, 0x0e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75,
	0x73, 0x12, 0x16, 0x2e, 0x47, 0x65, 0x74, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x17, 0x2e, 0x47, 0x65, 0x74, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x42, 0x24, 0x5a, 0x22, 0x47, 0x6f, 0x47, 0x72, 0x70, 0x63, 0x47, 0x65, 0x6e, 0x65,
	0x72, 0x61, 0x74, 0x65, 0x64, 0x2f, 0x46, 0x6f, 0x6f, 0x64, 0x43, 0x6f, 0x6d, 0x70, 0x61, 0x6e,
	0x79, 0x47, 0x61, 0x74, 0x65, 0x57, 0x61, 0x79, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_FoodCompanyGateWayService_proto_rawDescOnce sync.Once
	file_FoodCompanyGateWayService_proto_rawDescData = file_FoodCompanyGateWayService_proto_rawDesc
)

func file_FoodCompanyGateWayService_proto_rawDescGZIP() []byte {
	file_FoodCompanyGateWayService_proto_rawDescOnce.Do(func() {
		file_FoodCompanyGateWayService_proto_rawDescData = protoimpl.X.CompressGZIP(file_FoodCompanyGateWayService_proto_rawDescData)
	})
	return file_FoodCompanyGateWayService_proto_rawDescData
}

var file_FoodCompanyGateWayService_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_FoodCompanyGateWayService_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_FoodCompanyGateWayService_proto_goTypes = []any{
	(GetOrderStatusResponse_Status)(0), // 0: GetOrderStatusResponse.Status
	(*GetOrderStatusRequest)(nil),      // 1: GetOrderStatusRequest
	(*GetOrderStatusResponse)(nil),     // 2: GetOrderStatusResponse
}
var file_FoodCompanyGateWayService_proto_depIdxs = []int32{
	0, // 0: GetOrderStatusResponse.status:type_name -> GetOrderStatusResponse.Status
	1, // 1: FoodCompanyGateWayService.GetOrderStatus:input_type -> GetOrderStatusRequest
	2, // 2: FoodCompanyGateWayService.GetOrderStatus:output_type -> GetOrderStatusResponse
	2, // [2:3] is the sub-list for method output_type
	1, // [1:2] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_FoodCompanyGateWayService_proto_init() }
func file_FoodCompanyGateWayService_proto_init() {
	if File_FoodCompanyGateWayService_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_FoodCompanyGateWayService_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*GetOrderStatusRequest); i {
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
		file_FoodCompanyGateWayService_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*GetOrderStatusResponse); i {
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
	file_FoodCompanyGateWayService_proto_msgTypes[1].OneofWrappers = []any{}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_FoodCompanyGateWayService_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_FoodCompanyGateWayService_proto_goTypes,
		DependencyIndexes: file_FoodCompanyGateWayService_proto_depIdxs,
		EnumInfos:         file_FoodCompanyGateWayService_proto_enumTypes,
		MessageInfos:      file_FoodCompanyGateWayService_proto_msgTypes,
	}.Build()
	File_FoodCompanyGateWayService_proto = out.File
	file_FoodCompanyGateWayService_proto_rawDesc = nil
	file_FoodCompanyGateWayService_proto_goTypes = nil
	file_FoodCompanyGateWayService_proto_depIdxs = nil
}
