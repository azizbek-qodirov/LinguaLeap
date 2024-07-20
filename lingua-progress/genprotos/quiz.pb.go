// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.21.1
// source: lingua-protos/quiz.proto

package genprotos

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

type TestCheckReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LessonId string      `protobuf:"bytes,1,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty"`
	UserId   string      `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	Requests []*CheckReq `protobuf:"bytes,3,rep,name=requests,proto3" json:"requests,omitempty"`
}

func (x *TestCheckReq) Reset() {
	*x = TestCheckReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lingua_protos_quiz_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestCheckReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestCheckReq) ProtoMessage() {}

func (x *TestCheckReq) ProtoReflect() protoreflect.Message {
	mi := &file_lingua_protos_quiz_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestCheckReq.ProtoReflect.Descriptor instead.
func (*TestCheckReq) Descriptor() ([]byte, []int) {
	return file_lingua_protos_quiz_proto_rawDescGZIP(), []int{0}
}

func (x *TestCheckReq) GetLessonId() string {
	if x != nil {
		return x.LessonId
	}
	return ""
}

func (x *TestCheckReq) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *TestCheckReq) GetRequests() []*CheckReq {
	if x != nil {
		return x.Requests
	}
	return nil
}

type TestCheckReqForSwagger struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LessonId string      `protobuf:"bytes,1,opt,name=lesson_id,json=lessonId,proto3" json:"lesson_id,omitempty"`
	Requests []*CheckReq `protobuf:"bytes,2,rep,name=requests,proto3" json:"requests,omitempty"`
}

func (x *TestCheckReqForSwagger) Reset() {
	*x = TestCheckReqForSwagger{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lingua_protos_quiz_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestCheckReqForSwagger) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestCheckReqForSwagger) ProtoMessage() {}

func (x *TestCheckReqForSwagger) ProtoReflect() protoreflect.Message {
	mi := &file_lingua_protos_quiz_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestCheckReqForSwagger.ProtoReflect.Descriptor instead.
func (*TestCheckReqForSwagger) Descriptor() ([]byte, []int) {
	return file_lingua_protos_quiz_proto_rawDescGZIP(), []int{1}
}

func (x *TestCheckReqForSwagger) GetLessonId() string {
	if x != nil {
		return x.LessonId
	}
	return ""
}

func (x *TestCheckReqForSwagger) GetRequests() []*CheckReq {
	if x != nil {
		return x.Requests
	}
	return nil
}

type CheckReq struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ExerciseId    string `protobuf:"bytes,1,opt,name=exercise_id,json=exerciseId,proto3" json:"exercise_id,omitempty"`
	CorrectAnswer string `protobuf:"bytes,2,opt,name=correct_answer,json=correctAnswer,proto3" json:"correct_answer,omitempty"`
}

func (x *CheckReq) Reset() {
	*x = CheckReq{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lingua_protos_quiz_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckReq) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckReq) ProtoMessage() {}

func (x *CheckReq) ProtoReflect() protoreflect.Message {
	mi := &file_lingua_protos_quiz_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckReq.ProtoReflect.Descriptor instead.
func (*CheckReq) Descriptor() ([]byte, []int) {
	return file_lingua_protos_quiz_proto_rawDescGZIP(), []int{2}
}

func (x *CheckReq) GetExerciseId() string {
	if x != nil {
		return x.ExerciseId
	}
	return ""
}

func (x *CheckReq) GetCorrectAnswer() string {
	if x != nil {
		return x.CorrectAnswer
	}
	return ""
}

type TestResultRes struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TestsCount          int32  `protobuf:"varint,1,opt,name=tests_count,json=testsCount,proto3" json:"tests_count,omitempty"`
	CorrectAnswersCount int32  `protobuf:"varint,2,opt,name=correct_answers_count,json=correctAnswersCount,proto3" json:"correct_answers_count,omitempty"`
	XpGiven             int32  `protobuf:"varint,3,opt,name=xp_given,json=xpGiven,proto3" json:"xp_given,omitempty"`
	Feedback            string `protobuf:"bytes,4,opt,name=feedback,proto3" json:"feedback,omitempty"`
}

func (x *TestResultRes) Reset() {
	*x = TestResultRes{}
	if protoimpl.UnsafeEnabled {
		mi := &file_lingua_protos_quiz_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TestResultRes) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TestResultRes) ProtoMessage() {}

func (x *TestResultRes) ProtoReflect() protoreflect.Message {
	mi := &file_lingua_protos_quiz_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TestResultRes.ProtoReflect.Descriptor instead.
func (*TestResultRes) Descriptor() ([]byte, []int) {
	return file_lingua_protos_quiz_proto_rawDescGZIP(), []int{3}
}

func (x *TestResultRes) GetTestsCount() int32 {
	if x != nil {
		return x.TestsCount
	}
	return 0
}

func (x *TestResultRes) GetCorrectAnswersCount() int32 {
	if x != nil {
		return x.CorrectAnswersCount
	}
	return 0
}

func (x *TestResultRes) GetXpGiven() int32 {
	if x != nil {
		return x.XpGiven
	}
	return 0
}

func (x *TestResultRes) GetFeedback() string {
	if x != nil {
		return x.Feedback
	}
	return ""
}

var File_lingua_protos_quiz_proto protoreflect.FileDescriptor

var file_lingua_protos_quiz_proto_rawDesc = []byte{
	0x0a, 0x18, 0x6c, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f,
	0x71, 0x75, 0x69, 0x7a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x6c, 0x69, 0x6e, 0x67,
	0x75, 0x61, 0x22, 0x72, 0x0a, 0x0c, 0x54, 0x65, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52,
	0x65, 0x71, 0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x08, 0x72, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x6c, 0x69, 0x6e,
	0x67, 0x75, 0x61, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x52, 0x08, 0x72, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x63, 0x0a, 0x16, 0x54, 0x65, 0x73, 0x74, 0x43, 0x68,
	0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x46, 0x6f, 0x72, 0x53, 0x77, 0x61, 0x67, 0x67, 0x65, 0x72,
	0x12, 0x1b, 0x0a, 0x09, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x6c, 0x65, 0x73, 0x73, 0x6f, 0x6e, 0x49, 0x64, 0x12, 0x2c, 0x0a,
	0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x10, 0x2e, 0x6c, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x2e, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x52, 0x08, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x73, 0x22, 0x52, 0x0a, 0x08, 0x43,
	0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x12, 0x1f, 0x0a, 0x0b, 0x65, 0x78, 0x65, 0x72, 0x63,
	0x69, 0x73, 0x65, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x65, 0x78,
	0x65, 0x72, 0x63, 0x69, 0x73, 0x65, 0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x6f, 0x72, 0x72,
	0x65, 0x63, 0x74, 0x5f, 0x61, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0d, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72, 0x22,
	0x9b, 0x01, 0x0a, 0x0d, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74, 0x52, 0x65,
	0x73, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x65, 0x73, 0x74, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x65, 0x73, 0x74, 0x73, 0x43, 0x6f, 0x75,
	0x6e, 0x74, 0x12, 0x32, 0x0a, 0x15, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x5f, 0x61, 0x6e,
	0x73, 0x77, 0x65, 0x72, 0x73, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x13, 0x63, 0x6f, 0x72, 0x72, 0x65, 0x63, 0x74, 0x41, 0x6e, 0x73, 0x77, 0x65, 0x72,
	0x73, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x19, 0x0a, 0x08, 0x78, 0x70, 0x5f, 0x67, 0x69, 0x76,
	0x65, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x78, 0x70, 0x47, 0x69, 0x76, 0x65,
	0x6e, 0x12, 0x1a, 0x0a, 0x08, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x08, 0x66, 0x65, 0x65, 0x64, 0x62, 0x61, 0x63, 0x6b, 0x32, 0x49, 0x0a,
	0x0b, 0x51, 0x75, 0x69, 0x7a, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x3a, 0x0a, 0x09,
	0x53, 0x74, 0x61, 0x72, 0x74, 0x54, 0x65, 0x73, 0x74, 0x12, 0x14, 0x2e, 0x6c, 0x69, 0x6e, 0x67,
	0x75, 0x61, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x1a,
	0x15, 0x2e, 0x6c, 0x69, 0x6e, 0x67, 0x75, 0x61, 0x2e, 0x54, 0x65, 0x73, 0x74, 0x52, 0x65, 0x73,
	0x75, 0x6c, 0x74, 0x52, 0x65, 0x73, 0x22, 0x00, 0x42, 0x0c, 0x5a, 0x0a, 0x67, 0x65, 0x6e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_lingua_protos_quiz_proto_rawDescOnce sync.Once
	file_lingua_protos_quiz_proto_rawDescData = file_lingua_protos_quiz_proto_rawDesc
)

func file_lingua_protos_quiz_proto_rawDescGZIP() []byte {
	file_lingua_protos_quiz_proto_rawDescOnce.Do(func() {
		file_lingua_protos_quiz_proto_rawDescData = protoimpl.X.CompressGZIP(file_lingua_protos_quiz_proto_rawDescData)
	})
	return file_lingua_protos_quiz_proto_rawDescData
}

var file_lingua_protos_quiz_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_lingua_protos_quiz_proto_goTypes = []any{
	(*TestCheckReq)(nil),           // 0: lingua.TestCheckReq
	(*TestCheckReqForSwagger)(nil), // 1: lingua.TestCheckReqForSwagger
	(*CheckReq)(nil),               // 2: lingua.CheckReq
	(*TestResultRes)(nil),          // 3: lingua.TestResultRes
}
var file_lingua_protos_quiz_proto_depIdxs = []int32{
	2, // 0: lingua.TestCheckReq.requests:type_name -> lingua.CheckReq
	2, // 1: lingua.TestCheckReqForSwagger.requests:type_name -> lingua.CheckReq
	0, // 2: lingua.QuizService.StartTest:input_type -> lingua.TestCheckReq
	3, // 3: lingua.QuizService.StartTest:output_type -> lingua.TestResultRes
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_lingua_protos_quiz_proto_init() }
func file_lingua_protos_quiz_proto_init() {
	if File_lingua_protos_quiz_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_lingua_protos_quiz_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TestCheckReq); i {
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
		file_lingua_protos_quiz_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TestCheckReqForSwagger); i {
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
		file_lingua_protos_quiz_proto_msgTypes[2].Exporter = func(v any, i int) any {
			switch v := v.(*CheckReq); i {
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
		file_lingua_protos_quiz_proto_msgTypes[3].Exporter = func(v any, i int) any {
			switch v := v.(*TestResultRes); i {
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
			RawDescriptor: file_lingua_protos_quiz_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_lingua_protos_quiz_proto_goTypes,
		DependencyIndexes: file_lingua_protos_quiz_proto_depIdxs,
		MessageInfos:      file_lingua_protos_quiz_proto_msgTypes,
	}.Build()
	File_lingua_protos_quiz_proto = out.File
	file_lingua_protos_quiz_proto_rawDesc = nil
	file_lingua_protos_quiz_proto_goTypes = nil
	file_lingua_protos_quiz_proto_depIdxs = nil
}
