// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.6.1
// source: errors/shop.proto

package errors

import (
	_ "github.com/go-kratos/kratos/v2/api/kratos/api"
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

type Admin int32

const (
	Admin_UNKNOWN_ERROR Admin = 0
)

// Enum value maps for Admin.
var (
	Admin_name = map[int32]string{
		0: "UNKNOWN_ERROR",
	}
	Admin_value = map[string]int32{
		"UNKNOWN_ERROR": 0,
	}
)

func (x Admin) Enum() *Admin {
	p := new(Admin)
	*p = x
	return p
}

func (x Admin) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (Admin) Descriptor() protoreflect.EnumDescriptor {
	return file_errors_shop_proto_enumTypes[0].Descriptor()
}

func (Admin) Type() protoreflect.EnumType {
	return &file_errors_shop_proto_enumTypes[0]
}

func (x Admin) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use Admin.Descriptor instead.
func (Admin) EnumDescriptor() ([]byte, []int) {
	return file_errors_shop_proto_rawDescGZIP(), []int{0}
}

var File_errors_shop_proto protoreflect.FileDescriptor

var file_errors_shop_proto_rawDesc = []byte{
	0x0a, 0x11, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x2f, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x11, 0x73, 0x68, 0x6f, 0x70, 0x2e, 0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2e,
	0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x1a, 0x1c, 0x6b, 0x72, 0x61, 0x74, 0x6f, 0x73, 0x2f, 0x61,
	0x70, 0x69, 0x2f, 0x61, 0x6e, 0x6e, 0x6f, 0x74, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x2a, 0x1f, 0x0a, 0x05, 0x41, 0x64, 0x6d, 0x69, 0x6e, 0x12, 0x11, 0x0a,
	0x0d, 0x55, 0x4e, 0x4b, 0x4e, 0x4f, 0x57, 0x4e, 0x5f, 0x45, 0x52, 0x52, 0x4f, 0x52, 0x10, 0x00,
	0x1a, 0x03, 0xa0, 0x45, 0x01, 0x42, 0x1c, 0x50, 0x01, 0x5a, 0x18, 0x73, 0x68, 0x6f, 0x70, 0x2f,
	0x61, 0x64, 0x6d, 0x69, 0x6e, 0x2f, 0x65, 0x72, 0x72, 0x6f, 0x72, 0x73, 0x3b, 0x65, 0x72, 0x72,
	0x6f, 0x72, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_errors_shop_proto_rawDescOnce sync.Once
	file_errors_shop_proto_rawDescData = file_errors_shop_proto_rawDesc
)

func file_errors_shop_proto_rawDescGZIP() []byte {
	file_errors_shop_proto_rawDescOnce.Do(func() {
		file_errors_shop_proto_rawDescData = protoimpl.X.CompressGZIP(file_errors_shop_proto_rawDescData)
	})
	return file_errors_shop_proto_rawDescData
}

var file_errors_shop_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_errors_shop_proto_goTypes = []interface{}{
	(Admin)(0), // 0: shop.admin.errors.Admin
}
var file_errors_shop_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_errors_shop_proto_init() }
func file_errors_shop_proto_init() {
	if File_errors_shop_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_errors_shop_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_errors_shop_proto_goTypes,
		DependencyIndexes: file_errors_shop_proto_depIdxs,
		EnumInfos:         file_errors_shop_proto_enumTypes,
	}.Build()
	File_errors_shop_proto = out.File
	file_errors_shop_proto_rawDesc = nil
	file_errors_shop_proto_goTypes = nil
	file_errors_shop_proto_depIdxs = nil
}
