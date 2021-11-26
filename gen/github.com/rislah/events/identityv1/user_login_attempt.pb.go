// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        v3.17.3
// source: tracking/identity/v1/user_login_attempt.proto

package identityv1

import (
	_ "github.com/rislah/events/gen/protobuf/extension/v1"
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

type UserLoginAttempt_Result int32

const (
	UserLoginAttempt_RESULT_INVALID UserLoginAttempt_Result = 0
	UserLoginAttempt_RESULT_SUCCESS UserLoginAttempt_Result = 1
	UserLoginAttempt_RESULT_FAILED  UserLoginAttempt_Result = 2
)

// Enum value maps for UserLoginAttempt_Result.
var (
	UserLoginAttempt_Result_name = map[int32]string{
		0: "RESULT_INVALID",
		1: "RESULT_SUCCESS",
		2: "RESULT_FAILED",
	}
	UserLoginAttempt_Result_value = map[string]int32{
		"RESULT_INVALID": 0,
		"RESULT_SUCCESS": 1,
		"RESULT_FAILED":  2,
	}
)

func (x UserLoginAttempt_Result) Enum() *UserLoginAttempt_Result {
	p := new(UserLoginAttempt_Result)
	*p = x
	return p
}

func (x UserLoginAttempt_Result) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (UserLoginAttempt_Result) Descriptor() protoreflect.EnumDescriptor {
	return file_tracking_identity_v1_user_login_attempt_proto_enumTypes[0].Descriptor()
}

func (UserLoginAttempt_Result) Type() protoreflect.EnumType {
	return &file_tracking_identity_v1_user_login_attempt_proto_enumTypes[0]
}

func (x UserLoginAttempt_Result) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use UserLoginAttempt_Result.Descriptor instead.
func (UserLoginAttempt_Result) EnumDescriptor() ([]byte, []int) {
	return file_tracking_identity_v1_user_login_attempt_proto_rawDescGZIP(), []int{0, 0}
}

type UserLoginAttempt struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string                  `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Ip     string                  `protobuf:"bytes,2,opt,name=ip,proto3" json:"ip,omitempty"`
	Result UserLoginAttempt_Result `protobuf:"varint,3,opt,name=result,proto3,enum=tracking.identity.v1.UserLoginAttempt_Result" json:"result,omitempty"`
}

func (x *UserLoginAttempt) Reset() {
	*x = UserLoginAttempt{}
	if protoimpl.UnsafeEnabled {
		mi := &file_tracking_identity_v1_user_login_attempt_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserLoginAttempt) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserLoginAttempt) ProtoMessage() {}

func (x *UserLoginAttempt) ProtoReflect() protoreflect.Message {
	mi := &file_tracking_identity_v1_user_login_attempt_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserLoginAttempt.ProtoReflect.Descriptor instead.
func (*UserLoginAttempt) Descriptor() ([]byte, []int) {
	return file_tracking_identity_v1_user_login_attempt_proto_rawDescGZIP(), []int{0}
}

func (x *UserLoginAttempt) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *UserLoginAttempt) GetIp() string {
	if x != nil {
		return x.Ip
	}
	return ""
}

func (x *UserLoginAttempt) GetResult() UserLoginAttempt_Result {
	if x != nil {
		return x.Result
	}
	return UserLoginAttempt_RESULT_INVALID
}

var File_tracking_identity_v1_user_login_attempt_proto protoreflect.FileDescriptor

var file_tracking_identity_v1_user_login_attempt_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x69, 0x74, 0x79, 0x2f, 0x76, 0x31, 0x2f, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x6c, 0x6f, 0x67, 0x69,
	0x6e, 0x5f, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x14, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69,
	0x74, 0x79, 0x2e, 0x76, 0x31, 0x1a, 0x25, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f,
	0x65, 0x78, 0x74, 0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2f, 0x76, 0x31, 0x2f, 0x65, 0x78, 0x74,
	0x65, 0x6e, 0x73, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x01, 0x0a,
	0x10, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x74, 0x74, 0x65, 0x6d, 0x70,
	0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x70,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x70, 0x12, 0x45, 0x0a, 0x06, 0x72, 0x65,
	0x73, 0x75, 0x6c, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x2d, 0x2e, 0x74, 0x72, 0x61,
	0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x76,
	0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x4c, 0x6f, 0x67, 0x69, 0x6e, 0x41, 0x74, 0x74, 0x65, 0x6d,
	0x70, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x06, 0x72, 0x65, 0x73, 0x75, 0x6c,
	0x74, 0x22, 0x43, 0x0a, 0x06, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x12, 0x12, 0x0a, 0x0e, 0x52,
	0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x49, 0x4e, 0x56, 0x41, 0x4c, 0x49, 0x44, 0x10, 0x00, 0x12,
	0x12, 0x0a, 0x0e, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x53, 0x55, 0x43, 0x43, 0x45, 0x53,
	0x53, 0x10, 0x01, 0x12, 0x11, 0x0a, 0x0d, 0x52, 0x45, 0x53, 0x55, 0x4c, 0x54, 0x5f, 0x46, 0x41,
	0x49, 0x4c, 0x45, 0x44, 0x10, 0x02, 0x3a, 0x23, 0x82, 0xb5, 0x18, 0x1f, 0x74, 0x72, 0x61, 0x63,
	0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x2e, 0x6c, 0x6f,
	0x67, 0x69, 0x6e, 0x2d, 0x61, 0x74, 0x74, 0x65, 0x6d, 0x70, 0x74, 0x42, 0x41, 0x0a, 0x18, 0x63,
	0x6f, 0x6d, 0x2e, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x69, 0x6e, 0x67, 0x2e, 0x69, 0x64, 0x65, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x2e, 0x76, 0x31, 0x50, 0x01, 0x5a, 0x23, 0x67, 0x69, 0x74, 0x68, 0x75,
	0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x72, 0x69, 0x73, 0x6c, 0x61, 0x68, 0x2f, 0x65, 0x76, 0x65,
	0x6e, 0x74, 0x73, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x76, 0x31, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_tracking_identity_v1_user_login_attempt_proto_rawDescOnce sync.Once
	file_tracking_identity_v1_user_login_attempt_proto_rawDescData = file_tracking_identity_v1_user_login_attempt_proto_rawDesc
)

func file_tracking_identity_v1_user_login_attempt_proto_rawDescGZIP() []byte {
	file_tracking_identity_v1_user_login_attempt_proto_rawDescOnce.Do(func() {
		file_tracking_identity_v1_user_login_attempt_proto_rawDescData = protoimpl.X.CompressGZIP(file_tracking_identity_v1_user_login_attempt_proto_rawDescData)
	})
	return file_tracking_identity_v1_user_login_attempt_proto_rawDescData
}

var file_tracking_identity_v1_user_login_attempt_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_tracking_identity_v1_user_login_attempt_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_tracking_identity_v1_user_login_attempt_proto_goTypes = []interface{}{
	(UserLoginAttempt_Result)(0), // 0: tracking.identity.v1.UserLoginAttempt.Result
	(*UserLoginAttempt)(nil),     // 1: tracking.identity.v1.UserLoginAttempt
}
var file_tracking_identity_v1_user_login_attempt_proto_depIdxs = []int32{
	0, // 0: tracking.identity.v1.UserLoginAttempt.result:type_name -> tracking.identity.v1.UserLoginAttempt.Result
	1, // [1:1] is the sub-list for method output_type
	1, // [1:1] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_tracking_identity_v1_user_login_attempt_proto_init() }
func file_tracking_identity_v1_user_login_attempt_proto_init() {
	if File_tracking_identity_v1_user_login_attempt_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_tracking_identity_v1_user_login_attempt_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserLoginAttempt); i {
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
			RawDescriptor: file_tracking_identity_v1_user_login_attempt_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_tracking_identity_v1_user_login_attempt_proto_goTypes,
		DependencyIndexes: file_tracking_identity_v1_user_login_attempt_proto_depIdxs,
		EnumInfos:         file_tracking_identity_v1_user_login_attempt_proto_enumTypes,
		MessageInfos:      file_tracking_identity_v1_user_login_attempt_proto_msgTypes,
	}.Build()
	File_tracking_identity_v1_user_login_attempt_proto = out.File
	file_tracking_identity_v1_user_login_attempt_proto_rawDesc = nil
	file_tracking_identity_v1_user_login_attempt_proto_goTypes = nil
	file_tracking_identity_v1_user_login_attempt_proto_depIdxs = nil
}
