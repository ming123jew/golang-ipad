// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.23.0
// 	protoc        v3.6.1
// source: get_login_qr_code_request.proto

package mm_pb

import (
	proto "github.com/golang/protobuf/proto"
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

type GetLoginQRCodeRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseRequest     *BaseRequest `protobuf:"bytes,1,req,name=baseRequest" json:"baseRequest,omitempty"`
	Aes             *AesKey      `protobuf:"bytes,2,req,name=aes" json:"aes,omitempty"`
	Opcode          *uint32      `protobuf:"varint,3,req,name=opcode" json:"opcode,omitempty"`
	DeviceName      *string      `protobuf:"bytes,4,opt,name=deviceName" json:"deviceName,omitempty"`
	UserName        *string      `protobuf:"bytes,5,opt,name=userName" json:"userName,omitempty"`
	ExtDevLoginType *uint32      `protobuf:"varint,6,req,name=extDevLoginType" json:"extDevLoginType,omitempty"`
	HardwareExtra   *string      `protobuf:"bytes,7,opt,name=hardwareExtra" json:"hardwareExtra,omitempty"`
	SoftType        *string      `protobuf:"bytes,8,opt,name=softType" json:"softType,omitempty"`
	Rsa             *RSAPem      `protobuf:"bytes,9,opt,name=rsa" json:"rsa,omitempty"`
}

func (x *GetLoginQRCodeRequest) Reset() {
	*x = GetLoginQRCodeRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_get_login_qr_code_request_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetLoginQRCodeRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetLoginQRCodeRequest) ProtoMessage() {}

func (x *GetLoginQRCodeRequest) ProtoReflect() protoreflect.Message {
	mi := &file_get_login_qr_code_request_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetLoginQRCodeRequest.ProtoReflect.Descriptor instead.
func (*GetLoginQRCodeRequest) Descriptor() ([]byte, []int) {
	return file_get_login_qr_code_request_proto_rawDescGZIP(), []int{0}
}

func (x *GetLoginQRCodeRequest) GetBaseRequest() *BaseRequest {
	if x != nil {
		return x.BaseRequest
	}
	return nil
}

func (x *GetLoginQRCodeRequest) GetAes() *AesKey {
	if x != nil {
		return x.Aes
	}
	return nil
}

func (x *GetLoginQRCodeRequest) GetOpcode() uint32 {
	if x != nil && x.Opcode != nil {
		return *x.Opcode
	}
	return 0
}

func (x *GetLoginQRCodeRequest) GetDeviceName() string {
	if x != nil && x.DeviceName != nil {
		return *x.DeviceName
	}
	return ""
}

func (x *GetLoginQRCodeRequest) GetUserName() string {
	if x != nil && x.UserName != nil {
		return *x.UserName
	}
	return ""
}

func (x *GetLoginQRCodeRequest) GetExtDevLoginType() uint32 {
	if x != nil && x.ExtDevLoginType != nil {
		return *x.ExtDevLoginType
	}
	return 0
}

func (x *GetLoginQRCodeRequest) GetHardwareExtra() string {
	if x != nil && x.HardwareExtra != nil {
		return *x.HardwareExtra
	}
	return ""
}

func (x *GetLoginQRCodeRequest) GetSoftType() string {
	if x != nil && x.SoftType != nil {
		return *x.SoftType
	}
	return ""
}

func (x *GetLoginQRCodeRequest) GetRsa() *RSAPem {
	if x != nil {
		return x.Rsa
	}
	return nil
}

var File_get_login_qr_code_request_proto protoreflect.FileDescriptor

var file_get_login_qr_code_request_proto_rawDesc = []byte{
	0x0a, 0x1f, 0x67, 0x65, 0x74, 0x5f, 0x6c, 0x6f, 0x67, 0x69, 0x6e, 0x5f, 0x71, 0x72, 0x5f, 0x63,
	0x6f, 0x64, 0x65, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x12, 0x05, 0x6d, 0x6d, 0x5f, 0x70, 0x62, 0x1a, 0x11, 0x62, 0x61, 0x73, 0x65, 0x5f, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcf, 0x02, 0x0a, 0x15,
	0x47, 0x65, 0x74, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x51, 0x52, 0x43, 0x6f, 0x64, 0x65, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x34, 0x0a, 0x0b, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x18, 0x01, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x6d, 0x6d, 0x5f,
	0x70, 0x62, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x52, 0x0b,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1f, 0x0a, 0x03, 0x61,
	0x65, 0x73, 0x18, 0x02, 0x20, 0x02, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6d, 0x5f, 0x70, 0x62,
	0x2e, 0x41, 0x65, 0x73, 0x4b, 0x65, 0x79, 0x52, 0x03, 0x61, 0x65, 0x73, 0x12, 0x16, 0x0a, 0x06,
	0x6f, 0x70, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x03, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x06, 0x6f, 0x70,
	0x63, 0x6f, 0x64, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4e, 0x61,
	0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x64, 0x65, 0x76, 0x69, 0x63, 0x65,
	0x4e, 0x61, 0x6d, 0x65, 0x12, 0x1a, 0x0a, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x75, 0x73, 0x65, 0x72, 0x4e, 0x61, 0x6d, 0x65,
	0x12, 0x28, 0x0a, 0x0f, 0x65, 0x78, 0x74, 0x44, 0x65, 0x76, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54,
	0x79, 0x70, 0x65, 0x18, 0x06, 0x20, 0x02, 0x28, 0x0d, 0x52, 0x0f, 0x65, 0x78, 0x74, 0x44, 0x65,
	0x76, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x24, 0x0a, 0x0d, 0x68, 0x61,
	0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x45, 0x78, 0x74, 0x72, 0x61, 0x18, 0x07, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x68, 0x61, 0x72, 0x64, 0x77, 0x61, 0x72, 0x65, 0x45, 0x78, 0x74, 0x72, 0x61,
	0x12, 0x1a, 0x0a, 0x08, 0x73, 0x6f, 0x66, 0x74, 0x54, 0x79, 0x70, 0x65, 0x18, 0x08, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x73, 0x6f, 0x66, 0x74, 0x54, 0x79, 0x70, 0x65, 0x12, 0x1f, 0x0a, 0x03,
	0x72, 0x73, 0x61, 0x18, 0x09, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x0d, 0x2e, 0x6d, 0x6d, 0x5f, 0x70,
	0x62, 0x2e, 0x52, 0x53, 0x41, 0x50, 0x65, 0x6d, 0x52, 0x03, 0x72, 0x73, 0x61, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x3b, 0x6d, 0x6d, 0x5f, 0x70, 0x62,
}

var (
	file_get_login_qr_code_request_proto_rawDescOnce sync.Once
	file_get_login_qr_code_request_proto_rawDescData = file_get_login_qr_code_request_proto_rawDesc
)

func file_get_login_qr_code_request_proto_rawDescGZIP() []byte {
	file_get_login_qr_code_request_proto_rawDescOnce.Do(func() {
		file_get_login_qr_code_request_proto_rawDescData = protoimpl.X.CompressGZIP(file_get_login_qr_code_request_proto_rawDescData)
	})
	return file_get_login_qr_code_request_proto_rawDescData
}

var file_get_login_qr_code_request_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_get_login_qr_code_request_proto_goTypes = []interface{}{
	(*GetLoginQRCodeRequest)(nil), // 0: mm_pb.GetLoginQRCodeRequest
	(*BaseRequest)(nil),           // 1: mm_pb.BaseRequest
	(*AesKey)(nil),                // 2: mm_pb.AesKey
	(*RSAPem)(nil),                // 3: mm_pb.RSAPem
}
var file_get_login_qr_code_request_proto_depIdxs = []int32{
	1, // 0: mm_pb.GetLoginQRCodeRequest.baseRequest:type_name -> mm_pb.BaseRequest
	2, // 1: mm_pb.GetLoginQRCodeRequest.aes:type_name -> mm_pb.AesKey
	3, // 2: mm_pb.GetLoginQRCodeRequest.rsa:type_name -> mm_pb.RSAPem
	3, // [3:3] is the sub-list for method output_type
	3, // [3:3] is the sub-list for method input_type
	3, // [3:3] is the sub-list for extension type_name
	3, // [3:3] is the sub-list for extension extendee
	0, // [0:3] is the sub-list for field type_name
}

func init() { file_get_login_qr_code_request_proto_init() }
func file_get_login_qr_code_request_proto_init() {
	if File_get_login_qr_code_request_proto != nil {
		return
	}
	file_base_common_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_get_login_qr_code_request_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetLoginQRCodeRequest); i {
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
			RawDescriptor: file_get_login_qr_code_request_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_get_login_qr_code_request_proto_goTypes,
		DependencyIndexes: file_get_login_qr_code_request_proto_depIdxs,
		MessageInfos:      file_get_login_qr_code_request_proto_msgTypes,
	}.Build()
	File_get_login_qr_code_request_proto = out.File
	file_get_login_qr_code_request_proto_rawDesc = nil
	file_get_login_qr_code_request_proto_goTypes = nil
	file_get_login_qr_code_request_proto_depIdxs = nil
}
