// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: games.proto

package utils

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

type Response struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code int32  `protobuf:"varint,1,opt,name=Code,proto3" json:"Code,omitempty"` // 返回代碼
	Msg  string `protobuf:"bytes,2,opt,name=Msg,proto3" json:"Msg,omitempty"`    // 返回信息
}

func (x *Response) Reset() {
	*x = Response{}
	if protoimpl.UnsafeEnabled {
		mi := &file_games_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Response) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Response) ProtoMessage() {}

func (x *Response) ProtoReflect() protoreflect.Message {
	mi := &file_games_proto_msgTypes[0]
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
	return file_games_proto_rawDescGZIP(), []int{0}
}

func (x *Response) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *Response) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

type UserInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	NickName string `protobuf:"bytes,1,opt,name=NickName,proto3" json:"NickName,omitempty"`
	Icon     int32  `protobuf:"varint,2,opt,name=Icon,proto3" json:"Icon,omitempty"`
	Money    int64  `protobuf:"varint,3,opt,name=Money,proto3" json:"Money,omitempty"`
}

func (x *UserInfo) Reset() {
	*x = UserInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_games_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *UserInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*UserInfo) ProtoMessage() {}

func (x *UserInfo) ProtoReflect() protoreflect.Message {
	mi := &file_games_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use UserInfo.ProtoReflect.Descriptor instead.
func (*UserInfo) Descriptor() ([]byte, []int) {
	return file_games_proto_rawDescGZIP(), []int{1}
}

func (x *UserInfo) GetNickName() string {
	if x != nil {
		return x.NickName
	}
	return ""
}

func (x *UserInfo) GetIcon() int32 {
	if x != nil {
		return x.Icon
	}
	return 0
}

func (x *UserInfo) GetMoney() int64 {
	if x != nil {
		return x.Money
	}
	return 0
}

type HistoryInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RoomId      int64   `protobuf:"varint,1,opt,name=RoomId,proto3" json:"RoomId,omitempty"`
	RoomType    int64   `protobuf:"varint,2,opt,name=RoomType,proto3" json:"RoomType,omitempty"`
	TableStatus int32   `protobuf:"varint,3,opt,name=TableStatus,proto3" json:"TableStatus,omitempty"`
	TableTime   int32   `protobuf:"varint,4,opt,name=TableTime,proto3" json:"TableTime,omitempty"`
	HistoryList []int32 `protobuf:"varint,5,rep,packed,name=HistoryList,proto3" json:"HistoryList,omitempty"`
	TotalTime   int32   `protobuf:"varint,6,opt,name=TotalTime,proto3" json:"TotalTime,omitempty"`
}

func (x *HistoryInfo) Reset() {
	*x = HistoryInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_games_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *HistoryInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*HistoryInfo) ProtoMessage() {}

func (x *HistoryInfo) ProtoReflect() protoreflect.Message {
	mi := &file_games_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use HistoryInfo.ProtoReflect.Descriptor instead.
func (*HistoryInfo) Descriptor() ([]byte, []int) {
	return file_games_proto_rawDescGZIP(), []int{2}
}

func (x *HistoryInfo) GetRoomId() int64 {
	if x != nil {
		return x.RoomId
	}
	return 0
}

func (x *HistoryInfo) GetRoomType() int64 {
	if x != nil {
		return x.RoomType
	}
	return 0
}

func (x *HistoryInfo) GetTableStatus() int32 {
	if x != nil {
		return x.TableStatus
	}
	return 0
}

func (x *HistoryInfo) GetTableTime() int32 {
	if x != nil {
		return x.TableTime
	}
	return 0
}

func (x *HistoryInfo) GetHistoryList() []int32 {
	if x != nil {
		return x.HistoryList
	}
	return nil
}

func (x *HistoryInfo) GetTotalTime() int32 {
	if x != nil {
		return x.TotalTime
	}
	return 0
}

var File_games_proto protoreflect.FileDescriptor

var file_games_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x67, 0x61, 0x6d, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x75,
	0x74, 0x69, 0x6c, 0x73, 0x22, 0x30, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04,
	0x43, 0x6f, 0x64, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x4d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x03, 0x4d, 0x73, 0x67, 0x22, 0x50, 0x0a, 0x08, 0x55, 0x73, 0x65, 0x72, 0x49, 0x6e,
	0x66, 0x6f, 0x12, 0x1a, 0x0a, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x4e, 0x69, 0x63, 0x6b, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x12,
	0x0a, 0x04, 0x49, 0x63, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x49, 0x63,
	0x6f, 0x6e, 0x12, 0x14, 0x0a, 0x05, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x05, 0x4d, 0x6f, 0x6e, 0x65, 0x79, 0x22, 0xc1, 0x01, 0x0a, 0x0b, 0x48, 0x69, 0x73,
	0x74, 0x6f, 0x72, 0x79, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x16, 0x0a, 0x06, 0x52, 0x6f, 0x6f, 0x6d,
	0x49, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x52, 0x6f, 0x6f, 0x6d, 0x49, 0x64,
	0x12, 0x1a, 0x0a, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x08, 0x52, 0x6f, 0x6f, 0x6d, 0x54, 0x79, 0x70, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x0b, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x1c,
	0x0a, 0x09, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x54, 0x61, 0x62, 0x6c, 0x65, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x18, 0x05, 0x20, 0x03, 0x28,
	0x05, 0x52, 0x0b, 0x48, 0x69, 0x73, 0x74, 0x6f, 0x72, 0x79, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1c,
	0x0a, 0x09, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x09, 0x54, 0x6f, 0x74, 0x61, 0x6c, 0x54, 0x69, 0x6d, 0x65, 0x42, 0x0d, 0x5a, 0x0b,
	0x2e, 0x2e, 0x2f, 0x2e, 0x2e, 0x2f, 0x75, 0x74, 0x69, 0x6c, 0x73, 0x62, 0x06, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x33,
}

var (
	file_games_proto_rawDescOnce sync.Once
	file_games_proto_rawDescData = file_games_proto_rawDesc
)

func file_games_proto_rawDescGZIP() []byte {
	file_games_proto_rawDescOnce.Do(func() {
		file_games_proto_rawDescData = protoimpl.X.CompressGZIP(file_games_proto_rawDescData)
	})
	return file_games_proto_rawDescData
}

var file_games_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_games_proto_goTypes = []interface{}{
	(*Response)(nil),    // 0: utils.Response
	(*UserInfo)(nil),    // 1: utils.UserInfo
	(*HistoryInfo)(nil), // 2: utils.HistoryInfo
}
var file_games_proto_depIdxs = []int32{
	0, // [0:0] is the sub-list for method output_type
	0, // [0:0] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_games_proto_init() }
func file_games_proto_init() {
	if File_games_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_games_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_games_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*UserInfo); i {
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
		file_games_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*HistoryInfo); i {
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
			RawDescriptor: file_games_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_games_proto_goTypes,
		DependencyIndexes: file_games_proto_depIdxs,
		MessageInfos:      file_games_proto_msgTypes,
	}.Build()
	File_games_proto = out.File
	file_games_proto_rawDesc = nil
	file_games_proto_goTypes = nil
	file_games_proto_depIdxs = nil
}
