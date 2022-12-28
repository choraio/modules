// Code generated by protoc-gen-go-pulsar. DO NOT EDIT.
package contentv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
)

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.0
// 	protoc        (unknown)
// source: chora/content/v1/doc.proto

// chora.content.v1 describes the API for the content module.
//
// The following outlines revisions and tagged versions of the compiled code in
// relation to tagged releases in https://github.com/choraio/mods:
//
// - Revision 0: (in progress)
//

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

var File_chora_content_v1_doc_proto protoreflect.FileDescriptor

var file_chora_content_v1_doc_proto_rawDesc = []byte{
	0x0a, 0x1a, 0x63, 0x68, 0x6f, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f,
	0x76, 0x31, 0x2f, 0x64, 0x6f, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x10, 0x63, 0x68,
	0x6f, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0xbd,
	0x01, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x2e, 0x63, 0x68, 0x6f, 0x72, 0x61, 0x2e, 0x63, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2e, 0x76, 0x31, 0x42, 0x08, 0x44, 0x6f, 0x63, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x50, 0x01, 0x5a, 0x39, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f,
	0x63, 0x68, 0x6f, 0x72, 0x61, 0x69, 0x6f, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x2f,
	0x61, 0x70, 0x69, 0x2f, 0x63, 0x68, 0x6f, 0x72, 0x61, 0x2f, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e,
	0x74, 0x2f, 0x76, 0x31, 0x3b, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x76, 0x31, 0xa2, 0x02,
	0x03, 0x43, 0x43, 0x58, 0xaa, 0x02, 0x10, 0x43, 0x68, 0x6f, 0x72, 0x61, 0x2e, 0x43, 0x6f, 0x6e,
	0x74, 0x65, 0x6e, 0x74, 0x2e, 0x56, 0x31, 0xca, 0x02, 0x10, 0x43, 0x68, 0x6f, 0x72, 0x61, 0x5c,
	0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0xe2, 0x02, 0x1c, 0x43, 0x68, 0x6f,
	0x72, 0x61, 0x5c, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x5c, 0x56, 0x31, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x12, 0x43, 0x68, 0x6f, 0x72,
	0x61, 0x3a, 0x3a, 0x43, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x3a, 0x3a, 0x56, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var file_chora_content_v1_doc_proto_goTypes = []interface{}{}
var file_chora_content_v1_doc_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_chora_content_v1_doc_proto_init() }
func file_chora_content_v1_doc_proto_init() {
	if File_chora_content_v1_doc_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_chora_content_v1_doc_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   0,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_chora_content_v1_doc_proto_goTypes,
		DependencyIndexes: file_chora_content_v1_doc_proto_depIdxs,
	}.Build()
	File_chora_content_v1_doc_proto = out.File
	file_chora_content_v1_doc_proto_rawDesc = nil
	file_chora_content_v1_doc_proto_goTypes = nil
	file_chora_content_v1_doc_proto_depIdxs = nil
}
