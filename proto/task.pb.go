// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.34.2
// 	protoc        v3.18.0
// source: proto/task.proto

package proto

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

type TaskPayload struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid          string   `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	Msg           string   `protobuf:"bytes,2,opt,name=msg,proto3" json:"msg,omitempty"`
	Format        string   `protobuf:"bytes,3,opt,name=format,proto3" json:"format,omitempty"`
	ScanType      string   `protobuf:"bytes,4,opt,name=scan_type,json=scanType,proto3" json:"scan_type,omitempty"`
	Plugins       []string `protobuf:"bytes,5,rep,name=plugins,proto3" json:"plugins,omitempty"`
	ExecutionTime int32    `protobuf:"varint,6,opt,name=execution_time,json=executionTime,proto3" json:"execution_time,omitempty"`
	Delay         int32    `protobuf:"varint,7,opt,name=delay,proto3" json:"delay,omitempty"`
	Implement     bool     `protobuf:"varint,8,opt,name=implement,proto3" json:"implement,omitempty"`
}

func (x *TaskPayload) Reset() {
	*x = TaskPayload{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskPayload) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskPayload) ProtoMessage() {}

func (x *TaskPayload) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskPayload.ProtoReflect.Descriptor instead.
func (*TaskPayload) Descriptor() ([]byte, []int) {
	return file_proto_task_proto_rawDescGZIP(), []int{0}
}

func (x *TaskPayload) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *TaskPayload) GetMsg() string {
	if x != nil {
		return x.Msg
	}
	return ""
}

func (x *TaskPayload) GetFormat() string {
	if x != nil {
		return x.Format
	}
	return ""
}

func (x *TaskPayload) GetScanType() string {
	if x != nil {
		return x.ScanType
	}
	return ""
}

func (x *TaskPayload) GetPlugins() []string {
	if x != nil {
		return x.Plugins
	}
	return nil
}

func (x *TaskPayload) GetExecutionTime() int32 {
	if x != nil {
		return x.ExecutionTime
	}
	return 0
}

func (x *TaskPayload) GetDelay() int32 {
	if x != nil {
		return x.Delay
	}
	return 0
}

func (x *TaskPayload) GetImplement() bool {
	if x != nil {
		return x.Implement
	}
	return false
}

type TaskResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TaskId string `protobuf:"bytes,1,opt,name=task_id,json=taskId,proto3" json:"task_id,omitempty"`
	Queue  string `protobuf:"bytes,2,opt,name=queue,proto3" json:"queue,omitempty"`
	Status string `protobuf:"bytes,3,opt,name=status,proto3" json:"status,omitempty"`
	Uuid   string `protobuf:"bytes,4,opt,name=uuid,proto3" json:"uuid,omitempty"`
}

func (x *TaskResponse) Reset() {
	*x = TaskResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_task_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TaskResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TaskResponse) ProtoMessage() {}

func (x *TaskResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_task_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TaskResponse.ProtoReflect.Descriptor instead.
func (*TaskResponse) Descriptor() ([]byte, []int) {
	return file_proto_task_proto_rawDescGZIP(), []int{1}
}

func (x *TaskResponse) GetTaskId() string {
	if x != nil {
		return x.TaskId
	}
	return ""
}

func (x *TaskResponse) GetQueue() string {
	if x != nil {
		return x.Queue
	}
	return ""
}

func (x *TaskResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TaskResponse) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

var File_proto_task_proto protoreflect.FileDescriptor

var file_proto_task_proto_rawDesc = []byte{
	0x0a, 0x10, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x74, 0x61, 0x73, 0x6b, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x12, 0x0b, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x22,
	0xdd, 0x01, 0x0a, 0x0b, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x79, 0x6c, 0x6f, 0x61, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x12, 0x10, 0x0a, 0x03, 0x6d, 0x73, 0x67, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x03, 0x6d, 0x73, 0x67, 0x12, 0x16, 0x0a, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x66, 0x6f, 0x72, 0x6d, 0x61, 0x74, 0x12, 0x1b, 0x0a,
	0x09, 0x73, 0x63, 0x61, 0x6e, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x08, 0x73, 0x63, 0x61, 0x6e, 0x54, 0x79, 0x70, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x70, 0x6c,
	0x75, 0x67, 0x69, 0x6e, 0x73, 0x18, 0x05, 0x20, 0x03, 0x28, 0x09, 0x52, 0x07, 0x70, 0x6c, 0x75,
	0x67, 0x69, 0x6e, 0x73, 0x12, 0x25, 0x0a, 0x0e, 0x65, 0x78, 0x65, 0x63, 0x75, 0x74, 0x69, 0x6f,
	0x6e, 0x5f, 0x74, 0x69, 0x6d, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x65, 0x78,
	0x65, 0x63, 0x75, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x64,
	0x65, 0x6c, 0x61, 0x79, 0x18, 0x07, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x64, 0x65, 0x6c, 0x61,
	0x79, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x08,
	0x20, 0x01, 0x28, 0x08, 0x52, 0x09, 0x69, 0x6d, 0x70, 0x6c, 0x65, 0x6d, 0x65, 0x6e, 0x74, 0x22,
	0x69, 0x0a, 0x0c, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12,
	0x17, 0x0a, 0x07, 0x74, 0x61, 0x73, 0x6b, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x06, 0x74, 0x61, 0x73, 0x6b, 0x49, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x71, 0x75, 0x65, 0x75,
	0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x71, 0x75, 0x65, 0x75, 0x65, 0x12, 0x16,
	0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12, 0x12, 0x0a, 0x04, 0x75, 0x75, 0x69, 0x64, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x32, 0x51, 0x0a, 0x0b, 0x54, 0x61,
	0x73, 0x6b, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x42, 0x0a, 0x0b, 0x45, 0x6e, 0x71,
	0x75, 0x65, 0x75, 0x65, 0x54, 0x61, 0x73, 0x6b, 0x12, 0x18, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d,
	0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x54, 0x61, 0x73, 0x6b, 0x50, 0x61, 0x79, 0x6c, 0x6f,
	0x61, 0x64, 0x1a, 0x19, 0x2e, 0x74, 0x61, 0x73, 0x6b, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72,
	0x2e, 0x54, 0x61, 0x73, 0x6b, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x42, 0x09, 0x5a,
	0x07, 0x2e, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_task_proto_rawDescOnce sync.Once
	file_proto_task_proto_rawDescData = file_proto_task_proto_rawDesc
)

func file_proto_task_proto_rawDescGZIP() []byte {
	file_proto_task_proto_rawDescOnce.Do(func() {
		file_proto_task_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_task_proto_rawDescData)
	})
	return file_proto_task_proto_rawDescData
}

var file_proto_task_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_proto_task_proto_goTypes = []any{
	(*TaskPayload)(nil),  // 0: taskmanager.TaskPayload
	(*TaskResponse)(nil), // 1: taskmanager.TaskResponse
}
var file_proto_task_proto_depIdxs = []int32{
	0, // 0: taskmanager.TaskService.EnqueueTask:input_type -> taskmanager.TaskPayload
	1, // 1: taskmanager.TaskService.EnqueueTask:output_type -> taskmanager.TaskResponse
	1, // [1:2] is the sub-list for method output_type
	0, // [0:1] is the sub-list for method input_type
	0, // [0:0] is the sub-list for extension type_name
	0, // [0:0] is the sub-list for extension extendee
	0, // [0:0] is the sub-list for field type_name
}

func init() { file_proto_task_proto_init() }
func file_proto_task_proto_init() {
	if File_proto_task_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_task_proto_msgTypes[0].Exporter = func(v any, i int) any {
			switch v := v.(*TaskPayload); i {
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
		file_proto_task_proto_msgTypes[1].Exporter = func(v any, i int) any {
			switch v := v.(*TaskResponse); i {
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
			RawDescriptor: file_proto_task_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_task_proto_goTypes,
		DependencyIndexes: file_proto_task_proto_depIdxs,
		MessageInfos:      file_proto_task_proto_msgTypes,
	}.Build()
	File_proto_task_proto = out.File
	file_proto_task_proto_rawDesc = nil
	file_proto_task_proto_goTypes = nil
	file_proto_task_proto_depIdxs = nil
}
