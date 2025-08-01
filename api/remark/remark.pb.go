// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v6.31.1
// source: api/remark/remark.proto

package remark

import (
	_ "google.golang.org/genproto/googleapis/api/annotations"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
	unsafe "unsafe"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type AddRemarkRequest struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	StudentId        int64                  `protobuf:"varint,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	Text             string                 `protobuf:"bytes,2,opt,name=text,proto3" json:"text,omitempty"`
	Level            string                 `protobuf:"bytes,3,opt,name=level,proto3" json:"level,omitempty"`
	AddedByTeacherId int64                  `protobuf:"varint,4,opt,name=added_by_teacher_id,json=addedByTeacherId,proto3" json:"added_by_teacher_id,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *AddRemarkRequest) Reset() {
	*x = AddRemarkRequest{}
	mi := &file_api_remark_remark_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddRemarkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRemarkRequest) ProtoMessage() {}

func (x *AddRemarkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_remark_remark_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRemarkRequest.ProtoReflect.Descriptor instead.
func (*AddRemarkRequest) Descriptor() ([]byte, []int) {
	return file_api_remark_remark_proto_rawDescGZIP(), []int{0}
}

func (x *AddRemarkRequest) GetStudentId() int64 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

func (x *AddRemarkRequest) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *AddRemarkRequest) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *AddRemarkRequest) GetAddedByTeacherId() int64 {
	if x != nil {
		return x.AddedByTeacherId
	}
	return 0
}

type AddRemarkResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	Id            int64                  `protobuf:"varint,3,opt,name=id,proto3" json:"id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *AddRemarkResponse) Reset() {
	*x = AddRemarkResponse{}
	mi := &file_api_remark_remark_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *AddRemarkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddRemarkResponse) ProtoMessage() {}

func (x *AddRemarkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_remark_remark_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddRemarkResponse.ProtoReflect.Descriptor instead.
func (*AddRemarkResponse) Descriptor() ([]byte, []int) {
	return file_api_remark_remark_proto_rawDescGZIP(), []int{1}
}

func (x *AddRemarkResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *AddRemarkResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

func (x *AddRemarkResponse) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetRemarksForStudentRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	StudentId     int64                  `protobuf:"varint,1,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRemarksForStudentRequest) Reset() {
	*x = GetRemarksForStudentRequest{}
	mi := &file_api_remark_remark_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRemarksForStudentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRemarksForStudentRequest) ProtoMessage() {}

func (x *GetRemarksForStudentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_remark_remark_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRemarksForStudentRequest.ProtoReflect.Descriptor instead.
func (*GetRemarksForStudentRequest) Descriptor() ([]byte, []int) {
	return file_api_remark_remark_proto_rawDescGZIP(), []int{2}
}

func (x *GetRemarksForStudentRequest) GetStudentId() int64 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

type GetRemarksForStudentResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Remarks       []*Remark              `protobuf:"bytes,1,rep,name=remarks,proto3" json:"remarks,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRemarksForStudentResponse) Reset() {
	*x = GetRemarksForStudentResponse{}
	mi := &file_api_remark_remark_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRemarksForStudentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRemarksForStudentResponse) ProtoMessage() {}

func (x *GetRemarksForStudentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_remark_remark_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRemarksForStudentResponse.ProtoReflect.Descriptor instead.
func (*GetRemarksForStudentResponse) Descriptor() ([]byte, []int) {
	return file_api_remark_remark_proto_rawDescGZIP(), []int{3}
}

func (x *GetRemarksForStudentResponse) GetRemarks() []*Remark {
	if x != nil {
		return x.Remarks
	}
	return nil
}

type ApproveRemarkRequest struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	RemarkId      int64                  `protobuf:"varint,1,opt,name=remark_id,json=remarkId,proto3" json:"remark_id,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApproveRemarkRequest) Reset() {
	*x = ApproveRemarkRequest{}
	mi := &file_api_remark_remark_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApproveRemarkRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveRemarkRequest) ProtoMessage() {}

func (x *ApproveRemarkRequest) ProtoReflect() protoreflect.Message {
	mi := &file_api_remark_remark_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveRemarkRequest.ProtoReflect.Descriptor instead.
func (*ApproveRemarkRequest) Descriptor() ([]byte, []int) {
	return file_api_remark_remark_proto_rawDescGZIP(), []int{4}
}

func (x *ApproveRemarkRequest) GetRemarkId() int64 {
	if x != nil {
		return x.RemarkId
	}
	return 0
}

type ApproveRemarkResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Success       bool                   `protobuf:"varint,1,opt,name=success,proto3" json:"success,omitempty"`
	Message       string                 `protobuf:"bytes,2,opt,name=message,proto3" json:"message,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *ApproveRemarkResponse) Reset() {
	*x = ApproveRemarkResponse{}
	mi := &file_api_remark_remark_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ApproveRemarkResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ApproveRemarkResponse) ProtoMessage() {}

func (x *ApproveRemarkResponse) ProtoReflect() protoreflect.Message {
	mi := &file_api_remark_remark_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ApproveRemarkResponse.ProtoReflect.Descriptor instead.
func (*ApproveRemarkResponse) Descriptor() ([]byte, []int) {
	return file_api_remark_remark_proto_rawDescGZIP(), []int{5}
}

func (x *ApproveRemarkResponse) GetSuccess() bool {
	if x != nil {
		return x.Success
	}
	return false
}

func (x *ApproveRemarkResponse) GetMessage() string {
	if x != nil {
		return x.Message
	}
	return ""
}

type Remark struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	Id            int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	StudentId     int64                  `protobuf:"varint,2,opt,name=student_id,json=studentId,proto3" json:"student_id,omitempty"`
	TeacherId     int64                  `protobuf:"varint,3,opt,name=teacher_id,json=teacherId,proto3" json:"teacher_id,omitempty"`
	Text          string                 `protobuf:"bytes,4,opt,name=text,proto3" json:"text,omitempty"`
	Level         string                 `protobuf:"bytes,5,opt,name=level,proto3" json:"level,omitempty"`
	Date          string                 `protobuf:"bytes,6,opt,name=date,proto3" json:"date,omitempty"`
	Approved      bool                   `protobuf:"varint,7,opt,name=approved,proto3" json:"approved,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *Remark) Reset() {
	*x = Remark{}
	mi := &file_api_remark_remark_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Remark) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Remark) ProtoMessage() {}

func (x *Remark) ProtoReflect() protoreflect.Message {
	mi := &file_api_remark_remark_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Remark.ProtoReflect.Descriptor instead.
func (*Remark) Descriptor() ([]byte, []int) {
	return file_api_remark_remark_proto_rawDescGZIP(), []int{6}
}

func (x *Remark) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Remark) GetStudentId() int64 {
	if x != nil {
		return x.StudentId
	}
	return 0
}

func (x *Remark) GetTeacherId() int64 {
	if x != nil {
		return x.TeacherId
	}
	return 0
}

func (x *Remark) GetText() string {
	if x != nil {
		return x.Text
	}
	return ""
}

func (x *Remark) GetLevel() string {
	if x != nil {
		return x.Level
	}
	return ""
}

func (x *Remark) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

func (x *Remark) GetApproved() bool {
	if x != nil {
		return x.Approved
	}
	return false
}

var File_api_remark_remark_proto protoreflect.FileDescriptor

const file_api_remark_remark_proto_rawDesc = "" +
	"\n" +
	"\x17api/remark/remark.proto\x12\x06remark\x1a\x1cgoogle/api/annotations.proto\"\x8a\x01\n" +
	"\x10AddRemarkRequest\x12\x1d\n" +
	"\n" +
	"student_id\x18\x01 \x01(\x03R\tstudentId\x12\x12\n" +
	"\x04text\x18\x02 \x01(\tR\x04text\x12\x14\n" +
	"\x05level\x18\x03 \x01(\tR\x05level\x12-\n" +
	"\x13added_by_teacher_id\x18\x04 \x01(\x03R\x10addedByTeacherId\"W\n" +
	"\x11AddRemarkResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\x12\x0e\n" +
	"\x02id\x18\x03 \x01(\x03R\x02id\"<\n" +
	"\x1bGetRemarksForStudentRequest\x12\x1d\n" +
	"\n" +
	"student_id\x18\x01 \x01(\x03R\tstudentId\"H\n" +
	"\x1cGetRemarksForStudentResponse\x12(\n" +
	"\aremarks\x18\x01 \x03(\v2\x0e.remark.RemarkR\aremarks\"3\n" +
	"\x14ApproveRemarkRequest\x12\x1b\n" +
	"\tremark_id\x18\x01 \x01(\x03R\bremarkId\"K\n" +
	"\x15ApproveRemarkResponse\x12\x18\n" +
	"\asuccess\x18\x01 \x01(\bR\asuccess\x12\x18\n" +
	"\amessage\x18\x02 \x01(\tR\amessage\"\xb0\x01\n" +
	"\x06Remark\x12\x0e\n" +
	"\x02id\x18\x01 \x01(\x03R\x02id\x12\x1d\n" +
	"\n" +
	"student_id\x18\x02 \x01(\x03R\tstudentId\x12\x1d\n" +
	"\n" +
	"teacher_id\x18\x03 \x01(\x03R\tteacherId\x12\x12\n" +
	"\x04text\x18\x04 \x01(\tR\x04text\x12\x14\n" +
	"\x05level\x18\x05 \x01(\tR\x05level\x12\x12\n" +
	"\x04date\x18\x06 \x01(\tR\x04date\x12\x1a\n" +
	"\bapproved\x18\a \x01(\bR\bapproved2\xef\x02\n" +
	"\rRemarkService\x12X\n" +
	"\tAddRemark\x12\x18.remark.AddRemarkRequest\x1a\x19.remark.AddRemarkResponse\"\x16\x82\xd3\xe4\x93\x02\x10:\x01*\"\v/v1/remarks\x12\x8c\x01\n" +
	"\x14GetRemarksForStudent\x12#.remark.GetRemarksForStudentRequest\x1a$.remark.GetRemarksForStudentResponse\")\x82\xd3\xe4\x93\x02#\x12!/v1/students/{student_id}/remarks\x12u\n" +
	"\rApproveRemark\x12\x1c.remark.ApproveRemarkRequest\x1a\x1d.remark.ApproveRemarkResponse\"'\x82\xd3\xe4\x93\x02!\x1a\x1f/v1/remarks/{remark_id}/approveB7Z5github.com/karanparmar2o/student-dashboard/api/remarkb\x06proto3"

var (
	file_api_remark_remark_proto_rawDescOnce sync.Once
	file_api_remark_remark_proto_rawDescData []byte
)

func file_api_remark_remark_proto_rawDescGZIP() []byte {
	file_api_remark_remark_proto_rawDescOnce.Do(func() {
		file_api_remark_remark_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_api_remark_remark_proto_rawDesc), len(file_api_remark_remark_proto_rawDesc)))
	})
	return file_api_remark_remark_proto_rawDescData
}

var file_api_remark_remark_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_api_remark_remark_proto_goTypes = []any{
	(*AddRemarkRequest)(nil),             // 0: remark.AddRemarkRequest
	(*AddRemarkResponse)(nil),            // 1: remark.AddRemarkResponse
	(*GetRemarksForStudentRequest)(nil),  // 2: remark.GetRemarksForStudentRequest
	(*GetRemarksForStudentResponse)(nil), // 3: remark.GetRemarksForStudentResponse
	(*ApproveRemarkRequest)(nil),         // 4: remark.ApproveRemarkRequest
	(*ApproveRemarkResponse)(nil),        // 5: remark.ApproveRemarkResponse
	(*Remark)(nil),                       // 6: remark.Remark
}
var file_api_remark_remark_proto_depIdxs = []int32{
	6, // 0: remark.GetRemarksForStudentResponse.remarks:type_name -> remark.Remark
	0, // 1: remark.RemarkService.AddRemark:input_type -> remark.AddRemarkRequest
	2, // 2: remark.RemarkService.GetRemarksForStudent:input_type -> remark.GetRemarksForStudentRequest
	4, // 3: remark.RemarkService.ApproveRemark:input_type -> remark.ApproveRemarkRequest
	1, // 4: remark.RemarkService.AddRemark:output_type -> remark.AddRemarkResponse
	3, // 5: remark.RemarkService.GetRemarksForStudent:output_type -> remark.GetRemarksForStudentResponse
	5, // 6: remark.RemarkService.ApproveRemark:output_type -> remark.ApproveRemarkResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_api_remark_remark_proto_init() }
func file_api_remark_remark_proto_init() {
	if File_api_remark_remark_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_api_remark_remark_proto_rawDesc), len(file_api_remark_remark_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_api_remark_remark_proto_goTypes,
		DependencyIndexes: file_api_remark_remark_proto_depIdxs,
		MessageInfos:      file_api_remark_remark_proto_msgTypes,
	}.Build()
	File_api_remark_remark_proto = out.File
	file_api_remark_remark_proto_goTypes = nil
	file_api_remark_remark_proto_depIdxs = nil
}
