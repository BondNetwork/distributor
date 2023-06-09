// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.21.9
// source: claim/claim.proto

package claim

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

type ProofRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Account string `protobuf:"bytes,1,opt,name=account,proto3" json:"account,omitempty"`
}

func (x *ProofRequest) Reset() {
	*x = ProofRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_claim_claim_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProofRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProofRequest) ProtoMessage() {}

func (x *ProofRequest) ProtoReflect() protoreflect.Message {
	mi := &file_claim_claim_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProofRequest.ProtoReflect.Descriptor instead.
func (*ProofRequest) Descriptor() ([]byte, []int) {
	return file_claim_claim_proto_rawDescGZIP(), []int{0}
}

func (x *ProofRequest) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

type ProofResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BatchId int32    `protobuf:"varint,1,opt,name=batchId,proto3" json:"batchId,omitempty"`
	Index   int32    `protobuf:"varint,2,opt,name=index,proto3" json:"index,omitempty"`
	Account string   `protobuf:"bytes,3,opt,name=account,proto3" json:"account,omitempty"`
	Proofs  []string `protobuf:"bytes,4,rep,name=Proofs,proto3" json:"Proofs,omitempty"`
	Amount  string   `protobuf:"bytes,5,opt,name=amount,proto3" json:"amount,omitempty"`
}

func (x *ProofResponse) Reset() {
	*x = ProofResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_claim_claim_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ProofResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ProofResponse) ProtoMessage() {}

func (x *ProofResponse) ProtoReflect() protoreflect.Message {
	mi := &file_claim_claim_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ProofResponse.ProtoReflect.Descriptor instead.
func (*ProofResponse) Descriptor() ([]byte, []int) {
	return file_claim_claim_proto_rawDescGZIP(), []int{1}
}

func (x *ProofResponse) GetBatchId() int32 {
	if x != nil {
		return x.BatchId
	}
	return 0
}

func (x *ProofResponse) GetIndex() int32 {
	if x != nil {
		return x.Index
	}
	return 0
}

func (x *ProofResponse) GetAccount() string {
	if x != nil {
		return x.Account
	}
	return ""
}

func (x *ProofResponse) GetProofs() []string {
	if x != nil {
		return x.Proofs
	}
	return nil
}

func (x *ProofResponse) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

var File_claim_claim_proto protoreflect.FileDescriptor

var file_claim_claim_proto_rawDesc = []byte{
	0x0a, 0x11, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x2f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x05, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x22, 0x28, 0x0a, 0x0c, 0x50, 0x72,
	0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x89, 0x01, 0x0a, 0x0d, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65,
	0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x62, 0x61, 0x74, 0x63, 0x68, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x62, 0x61, 0x74, 0x63, 0x68, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x69, 0x6e, 0x64, 0x65, 0x78, 0x12, 0x18, 0x0a, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x61, 0x63, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x16, 0x0a, 0x06, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x09,
	0x52, 0x06, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x73, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x32, 0x3e, 0x0a, 0x05, 0x43, 0x6c, 0x61, 0x69, 0x6d, 0x12, 0x35, 0x0a, 0x08, 0x67, 0x65, 0x74,
	0x50, 0x72, 0x6f, 0x6f, 0x66, 0x12, 0x13, 0x2e, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x2e, 0x50, 0x72,
	0x6f, 0x6f, 0x66, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x14, 0x2e, 0x63, 0x6c, 0x61,
	0x69, 0x6d, 0x2e, 0x50, 0x72, 0x6f, 0x6f, 0x66, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x09, 0x5a, 0x07, 0x2e, 0x2f, 0x63, 0x6c, 0x61, 0x69, 0x6d, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_claim_claim_proto_rawDescOnce sync.Once
	file_claim_claim_proto_rawDescData = file_claim_claim_proto_rawDesc
)

func file_claim_claim_proto_rawDescGZIP() []byte {
	file_claim_claim_proto_rawDescOnce.Do(func() {
		file_claim_claim_proto_rawDescData = protoimpl.X.CompressGZIP(file_claim_claim_proto_rawDescData)
	})
	return file_claim_claim_proto_rawDescData
}

var file_claim_claim_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_claim_claim_proto_goTypes = []interface{}{
	(*ProofRequest)(nil),  // 0: claim.ProofRequest
	(*ProofResponse)(nil), // 1: claim.ProofResponse
}
var file_claim_claim_proto_depIdxs = []int32{
	0, // 0: claim.Claim.getProof:input_type -> claim.ProofRequest
	1, // 1: claim.Claim.getProof:output_type -> claim.ProofResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_claim_claim_proto_init() }
func file_claim_claim_proto_init() {
	if File_claim_claim_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_claim_claim_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProofRequest); i {
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
		file_claim_claim_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ProofResponse); i {
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
			RawDescriptor: file_claim_claim_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_claim_claim_proto_goTypes,
		DependencyIndexes: file_claim_claim_proto_depIdxs,
		MessageInfos:      file_claim_claim_proto_msgTypes,
	}.Build()
	File_claim_claim_proto = out.File
	file_claim_claim_proto_rawDesc = nil
	file_claim_claim_proto_goTypes = nil
	file_claim_claim_proto_depIdxs = nil
}
