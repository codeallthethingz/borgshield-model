// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.25.0
// 	protoc        v3.13.0
// source: model.proto

package model

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

type VirusFoundRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Viruses []*Virus `protobuf:"bytes,1,rep,name=viruses,proto3" json:"viruses,omitempty"`
	Machine *Machine `protobuf:"bytes,2,opt,name=machine,proto3" json:"machine,omitempty"`
	MacHash string   `protobuf:"bytes,3,opt,name=macHash,proto3" json:"macHash,omitempty"`
}

func (x *VirusFoundRequest) Reset() {
	*x = VirusFoundRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirusFoundRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirusFoundRequest) ProtoMessage() {}

func (x *VirusFoundRequest) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirusFoundRequest.ProtoReflect.Descriptor instead.
func (*VirusFoundRequest) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{0}
}

func (x *VirusFoundRequest) GetViruses() []*Virus {
	if x != nil {
		return x.Viruses
	}
	return nil
}

func (x *VirusFoundRequest) GetMachine() *Machine {
	if x != nil {
		return x.Machine
	}
	return nil
}

func (x *VirusFoundRequest) GetMacHash() string {
	if x != nil {
		return x.MacHash
	}
	return ""
}

type Machine struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MacHash  string `protobuf:"bytes,1,opt,name=macHash,proto3" json:"macHash,omitempty"`
	Hostname string `protobuf:"bytes,2,opt,name=hostname,proto3" json:"hostname,omitempty"`
}

func (x *Machine) Reset() {
	*x = Machine{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Machine) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Machine) ProtoMessage() {}

func (x *Machine) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Machine.ProtoReflect.Descriptor instead.
func (*Machine) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{1}
}

func (x *Machine) GetMacHash() string {
	if x != nil {
		return x.MacHash
	}
	return ""
}

func (x *Machine) GetHostname() string {
	if x != nil {
		return x.Hostname
	}
	return ""
}

type Virus struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
}

func (x *Virus) Reset() {
	*x = Virus{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Virus) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Virus) ProtoMessage() {}

func (x *Virus) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Virus.ProtoReflect.Descriptor instead.
func (*Virus) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{2}
}

func (x *Virus) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Virus) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

type VirusFoundReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *VirusFoundReply) Reset() {
	*x = VirusFoundReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VirusFoundReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VirusFoundReply) ProtoMessage() {}

func (x *VirusFoundReply) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VirusFoundReply.ProtoReflect.Descriptor instead.
func (*VirusFoundReply) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{3}
}

type MachineListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Machines []*MachineViruses `protobuf:"bytes,1,rep,name=machines,proto3" json:"machines,omitempty"`
}

func (x *MachineListResponse) Reset() {
	*x = MachineListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineListResponse) ProtoMessage() {}

func (x *MachineListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineListResponse.ProtoReflect.Descriptor instead.
func (*MachineListResponse) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{4}
}

func (x *MachineListResponse) GetMachines() []*MachineViruses {
	if x != nil {
		return x.Machines
	}
	return nil
}

type MachineViruses struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Viruses     []*Virus `protobuf:"bytes,1,rep,name=viruses,proto3" json:"viruses,omitempty"`
	Machine     *Machine `protobuf:"bytes,2,opt,name=machine,proto3" json:"machine,omitempty"`
	LastUpdated int64    `protobuf:"varint,3,opt,name=lastUpdated,proto3" json:"lastUpdated,omitempty"`
	MacHash     string   `protobuf:"bytes,4,opt,name=macHash,proto3" json:"macHash,omitempty"`
	Domain      string   `protobuf:"bytes,5,opt,name=domain,proto3" json:"domain,omitempty"`
	Email       string   `protobuf:"bytes,6,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *MachineViruses) Reset() {
	*x = MachineViruses{}
	if protoimpl.UnsafeEnabled {
		mi := &file_model_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MachineViruses) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MachineViruses) ProtoMessage() {}

func (x *MachineViruses) ProtoReflect() protoreflect.Message {
	mi := &file_model_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MachineViruses.ProtoReflect.Descriptor instead.
func (*MachineViruses) Descriptor() ([]byte, []int) {
	return file_model_proto_rawDescGZIP(), []int{5}
}

func (x *MachineViruses) GetViruses() []*Virus {
	if x != nil {
		return x.Viruses
	}
	return nil
}

func (x *MachineViruses) GetMachine() *Machine {
	if x != nil {
		return x.Machine
	}
	return nil
}

func (x *MachineViruses) GetLastUpdated() int64 {
	if x != nil {
		return x.LastUpdated
	}
	return 0
}

func (x *MachineViruses) GetMacHash() string {
	if x != nil {
		return x.MacHash
	}
	return ""
}

func (x *MachineViruses) GetDomain() string {
	if x != nil {
		return x.Domain
	}
	return ""
}

func (x *MachineViruses) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

var File_model_proto protoreflect.FileDescriptor

var file_model_proto_rawDesc = []byte{
	0x0a, 0x0b, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x05, 0x6d,
	0x6f, 0x64, 0x65, 0x6c, 0x22, 0x7f, 0x0a, 0x11, 0x56, 0x69, 0x72, 0x75, 0x73, 0x46, 0x6f, 0x75,
	0x6e, 0x64, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x26, 0x0a, 0x07, 0x76, 0x69, 0x72,
	0x75, 0x73, 0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64,
	0x65, 0x6c, 0x2e, 0x56, 0x69, 0x72, 0x75, 0x73, 0x52, 0x07, 0x76, 0x69, 0x72, 0x75, 0x73, 0x65,
	0x73, 0x12, 0x28, 0x0a, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69,
	0x6e, 0x65, 0x52, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x61, 0x63, 0x48, 0x61, 0x73, 0x68, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61,
	0x63, 0x48, 0x61, 0x73, 0x68, 0x22, 0x3f, 0x0a, 0x07, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x12, 0x18, 0x0a, 0x07, 0x6d, 0x61, 0x63, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x07, 0x6d, 0x61, 0x63, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x68, 0x6f,
	0x73, 0x74, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x2f, 0x0a, 0x05, 0x56, 0x69, 0x72, 0x75, 0x73, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x22, 0x11, 0x0a, 0x0f, 0x56, 0x69, 0x72, 0x75, 0x73,
	0x46, 0x6f, 0x75, 0x6e, 0x64, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x48, 0x0a, 0x13, 0x4d, 0x61,
	0x63, 0x68, 0x69, 0x6e, 0x65, 0x4c, 0x69, 0x73, 0x74, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x12, 0x31, 0x0a, 0x08, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x56, 0x69, 0x72, 0x75, 0x73, 0x65, 0x73, 0x52, 0x08, 0x6d, 0x61, 0x63, 0x68,
	0x69, 0x6e, 0x65, 0x73, 0x22, 0xcc, 0x01, 0x0a, 0x0e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x56, 0x69, 0x72, 0x75, 0x73, 0x65, 0x73, 0x12, 0x26, 0x0a, 0x07, 0x76, 0x69, 0x72, 0x75, 0x73,
	0x65, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x0c, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c,
	0x2e, 0x56, 0x69, 0x72, 0x75, 0x73, 0x52, 0x07, 0x76, 0x69, 0x72, 0x75, 0x73, 0x65, 0x73, 0x12,
	0x28, 0x0a, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0e, 0x2e, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x2e, 0x4d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65,
	0x52, 0x07, 0x6d, 0x61, 0x63, 0x68, 0x69, 0x6e, 0x65, 0x12, 0x20, 0x0a, 0x0b, 0x6c, 0x61, 0x73,
	0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b,
	0x6c, 0x61, 0x73, 0x74, 0x55, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x6d,
	0x61, 0x63, 0x48, 0x61, 0x73, 0x68, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6d, 0x61,
	0x63, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x18,
	0x05, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x64, 0x6f, 0x6d, 0x61, 0x69, 0x6e, 0x12, 0x14, 0x0a,
	0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d,
	0x61, 0x69, 0x6c, 0x42, 0x07, 0x5a, 0x05, 0x6d, 0x6f, 0x64, 0x65, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_model_proto_rawDescOnce sync.Once
	file_model_proto_rawDescData = file_model_proto_rawDesc
)

func file_model_proto_rawDescGZIP() []byte {
	file_model_proto_rawDescOnce.Do(func() {
		file_model_proto_rawDescData = protoimpl.X.CompressGZIP(file_model_proto_rawDescData)
	})
	return file_model_proto_rawDescData
}

var file_model_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_model_proto_goTypes = []interface{}{
	(*VirusFoundRequest)(nil),   // 0: model.VirusFoundRequest
	(*Machine)(nil),             // 1: model.Machine
	(*Virus)(nil),               // 2: model.Virus
	(*VirusFoundReply)(nil),     // 3: model.VirusFoundReply
	(*MachineListResponse)(nil), // 4: model.MachineListResponse
	(*MachineViruses)(nil),      // 5: model.MachineViruses
}
var file_model_proto_depIdxs = []int32{
	2, // 0: model.VirusFoundRequest.viruses:type_name -> model.Virus
	1, // 1: model.VirusFoundRequest.machine:type_name -> model.Machine
	5, // 2: model.MachineListResponse.machines:type_name -> model.MachineViruses
	2, // 3: model.MachineViruses.viruses:type_name -> model.Virus
	1, // 4: model.MachineViruses.machine:type_name -> model.Machine
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_model_proto_init() }
func file_model_proto_init() {
	if File_model_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_model_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirusFoundRequest); i {
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
		file_model_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Machine); i {
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
		file_model_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Virus); i {
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
		file_model_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VirusFoundReply); i {
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
		file_model_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineListResponse); i {
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
		file_model_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MachineViruses); i {
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
			RawDescriptor: file_model_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_model_proto_goTypes,
		DependencyIndexes: file_model_proto_depIdxs,
		MessageInfos:      file_model_proto_msgTypes,
	}.Build()
	File_model_proto = out.File
	file_model_proto_rawDesc = nil
	file_model_proto_goTypes = nil
	file_model_proto_depIdxs = nil
}
