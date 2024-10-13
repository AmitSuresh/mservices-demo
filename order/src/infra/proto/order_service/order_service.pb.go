// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v5.27.2
// source: order_service.proto

package order_service

import (
	reflect "reflect"
	sync "sync"

	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type OrderReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer string `protobuf:"bytes,1,opt,name=Customer,proto3" json:"Customer,omitempty"`
	Quantity uint32 `protobuf:"varint,2,opt,name=Quantity,proto3" json:"Quantity,omitempty"`
	Sku      string `protobuf:"bytes,3,opt,name=Sku,proto3" json:"Sku,omitempty"`
}

func (x *OrderReq) Reset() {
	*x = OrderReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_service_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderReq) ProtoMessage() {}

func (x *OrderReq) ProtoReflect() protoreflect.Message {
	mi := &file_order_service_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderReq.ProtoReflect.Descriptor instead.
func (*OrderReq) Descriptor() ([]byte, []int) {
	return file_order_service_proto_rawDescGZIP(), []int{0}
}

func (x *OrderReq) GetCustomer() string {
	if x != nil {
		return x.Customer
	}
	return ""
}

func (x *OrderReq) GetQuantity() uint32 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderReq) GetSku() string {
	if x != nil {
		return x.Sku
	}
	return ""
}

type OrderResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer string `protobuf:"bytes,1,opt,name=Customer,proto3" json:"Customer,omitempty"`
	Id       string `protobuf:"bytes,2,opt,name=Id,proto3" json:"Id,omitempty"`
	Shipped  bool   `protobuf:"varint,3,opt,name=Shipped,proto3" json:"Shipped,omitempty"`
}

func (x *OrderResp) Reset() {
	*x = OrderResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_order_service_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderResp) ProtoMessage() {}

func (x *OrderResp) ProtoReflect() protoreflect.Message {
	mi := &file_order_service_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderResp.ProtoReflect.Descriptor instead.
func (*OrderResp) Descriptor() ([]byte, []int) {
	return file_order_service_proto_rawDescGZIP(), []int{1}
}

func (x *OrderResp) GetCustomer() string {
	if x != nil {
		return x.Customer
	}
	return ""
}

func (x *OrderResp) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderResp) GetShipped() bool {
	if x != nil {
		return x.Shipped
	}
	return false
}

var File_order_service_proto protoreflect.FileDescriptor

var file_order_service_proto_rawDesc = []byte{
	0x0a, 0x13, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0d, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72,
	0x76, 0x69, 0x63, 0x65, 0x22, 0x54, 0x0a, 0x08, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x12, 0x1a, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1a, 0x0a, 0x08,
	0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0d, 0x52, 0x08,
	0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x53, 0x6b, 0x75, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x53, 0x6b, 0x75, 0x22, 0x51, 0x0a, 0x09, 0x4f, 0x72,
	0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1a, 0x0a, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x43, 0x75, 0x73, 0x74, 0x6f,
	0x6d, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x49, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x02, 0x49, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x07, 0x53, 0x68, 0x69, 0x70, 0x70, 0x65, 0x64, 0x32, 0x50, 0x0a,
	0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x40, 0x0a,
	0x0b, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x12, 0x17, 0x2e, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64,
	0x65, 0x72, 0x52, 0x65, 0x71, 0x1a, 0x18, 0x2e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65,
	0x72, 0x76, 0x69, 0x63, 0x65, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x52, 0x65, 0x73, 0x70, 0x42,
	0x10, 0x5a, 0x0e, 0x2f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63,
	0x65, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_order_service_proto_rawDescOnce sync.Once
	file_order_service_proto_rawDescData = file_order_service_proto_rawDesc
)

func file_order_service_proto_rawDescGZIP() []byte {
	file_order_service_proto_rawDescOnce.Do(func() {
		file_order_service_proto_rawDescData = protoimpl.X.CompressGZIP(file_order_service_proto_rawDescData)
	})
	return file_order_service_proto_rawDescData
}

var (
	file_order_service_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
	file_order_service_proto_goTypes  = []interface{}{
		(*OrderReq)(nil),  // 0: order_service.OrderReq
		(*OrderResp)(nil), // 1: order_service.OrderResp
	}
)
var file_order_service_proto_depIdxs = []int32{
	0, // 0: order_service.OrderService.CreateOrder:input_type -> order_service.OrderReq
	1, // 1: order_service.OrderService.CreateOrder:output_type -> order_service.OrderResp
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_order_service_proto_init() }
func file_order_service_proto_init() {
	if File_order_service_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_order_service_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderReq); i {
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
		file_order_service_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderResp); i {
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
			RawDescriptor: file_order_service_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_order_service_proto_goTypes,
		DependencyIndexes: file_order_service_proto_depIdxs,
		MessageInfos:      file_order_service_proto_msgTypes,
	}.Build()
	File_order_service_proto = out.File
	file_order_service_proto_rawDesc = nil
	file_order_service_proto_goTypes = nil
	file_order_service_proto_depIdxs = nil
}
