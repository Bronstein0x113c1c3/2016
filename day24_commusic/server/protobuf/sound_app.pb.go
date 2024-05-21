// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.32.0
// 	protoc        v3.21.12
// source: sound_app.proto

package protobuf

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

type Signal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Signal) Reset() {
	*x = Signal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sound_app_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Signal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Signal) ProtoMessage() {}

func (x *Signal) ProtoReflect() protoreflect.Message {
	mi := &file_sound_app_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Signal.ProtoReflect.Descriptor instead.
func (*Signal) Descriptor() ([]byte, []int) {
	return file_sound_app_proto_rawDescGZIP(), []int{0}
}

type SongRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Request:
	//
	//	*SongRequest_Id
	//	*SongRequest_S
	Request isSongRequest_Request `protobuf_oneof:"Request"`
}

func (x *SongRequest) Reset() {
	*x = SongRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sound_app_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SongRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SongRequest) ProtoMessage() {}

func (x *SongRequest) ProtoReflect() protoreflect.Message {
	mi := &file_sound_app_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SongRequest.ProtoReflect.Descriptor instead.
func (*SongRequest) Descriptor() ([]byte, []int) {
	return file_sound_app_proto_rawDescGZIP(), []int{1}
}

func (m *SongRequest) GetRequest() isSongRequest_Request {
	if m != nil {
		return m.Request
	}
	return nil
}

func (x *SongRequest) GetId() int32 {
	if x, ok := x.GetRequest().(*SongRequest_Id); ok {
		return x.Id
	}
	return 0
}

func (x *SongRequest) GetS() *Signal {
	if x, ok := x.GetRequest().(*SongRequest_S); ok {
		return x.S
	}
	return nil
}

type isSongRequest_Request interface {
	isSongRequest_Request()
}

type SongRequest_Id struct {
	Id int32 `protobuf:"varint,1,opt,name=id,proto3,oneof"`
}

type SongRequest_S struct {
	S *Signal `protobuf:"bytes,2,opt,name=s,proto3,oneof"`
}

func (*SongRequest_Id) isSongRequest_Request() {}

func (*SongRequest_S) isSongRequest_Request() {}

type SongInfo struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ArtistName    string `protobuf:"bytes,3,opt,name=artist_name,json=artistName,proto3" json:"artist_name,omitempty"`
	AlbumName     string `protobuf:"bytes,4,opt,name=album_name,json=albumName,proto3" json:"album_name,omitempty"`
	TrackNumber   int32  `protobuf:"varint,5,opt,name=track_number,json=trackNumber,proto3" json:"track_number,omitempty"`
	PublishedYear int32  `protobuf:"varint,6,opt,name=published_year,json=publishedYear,proto3" json:"published_year,omitempty"`
}

func (x *SongInfo) Reset() {
	*x = SongInfo{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sound_app_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SongInfo) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SongInfo) ProtoMessage() {}

func (x *SongInfo) ProtoReflect() protoreflect.Message {
	mi := &file_sound_app_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SongInfo.ProtoReflect.Descriptor instead.
func (*SongInfo) Descriptor() ([]byte, []int) {
	return file_sound_app_proto_rawDescGZIP(), []int{2}
}

func (x *SongInfo) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *SongInfo) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *SongInfo) GetArtistName() string {
	if x != nil {
		return x.ArtistName
	}
	return ""
}

func (x *SongInfo) GetAlbumName() string {
	if x != nil {
		return x.AlbumName
	}
	return ""
}

func (x *SongInfo) GetTrackNumber() int32 {
	if x != nil {
		return x.TrackNumber
	}
	return 0
}

func (x *SongInfo) GetPublishedYear() int32 {
	if x != nil {
		return x.PublishedYear
	}
	return 0
}

type AllSong struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Songs []*SongInfo `protobuf:"bytes,1,rep,name=songs,proto3" json:"songs,omitempty"`
}

func (x *AllSong) Reset() {
	*x = AllSong{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sound_app_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AllSong) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AllSong) ProtoMessage() {}

func (x *AllSong) ProtoReflect() protoreflect.Message {
	mi := &file_sound_app_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AllSong.ProtoReflect.Descriptor instead.
func (*AllSong) Descriptor() ([]byte, []int) {
	return file_sound_app_proto_rawDescGZIP(), []int{3}
}

func (x *AllSong) GetSongs() []*SongInfo {
	if x != nil {
		return x.Songs
	}
	return nil
}

type Chunk struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Left  float64 `protobuf:"fixed64,1,opt,name=left,proto3" json:"left,omitempty"`
	Right float64 `protobuf:"fixed64,2,opt,name=right,proto3" json:"right,omitempty"`
}

func (x *Chunk) Reset() {
	*x = Chunk{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sound_app_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunk) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunk) ProtoMessage() {}

func (x *Chunk) ProtoReflect() protoreflect.Message {
	mi := &file_sound_app_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunk.ProtoReflect.Descriptor instead.
func (*Chunk) Descriptor() ([]byte, []int) {
	return file_sound_app_proto_rawDescGZIP(), []int{4}
}

func (x *Chunk) GetLeft() float64 {
	if x != nil {
		return x.Left
	}
	return 0
}

func (x *Chunk) GetRight() float64 {
	if x != nil {
		return x.Right
	}
	return 0
}

type Chunks struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	C []*Chunk `protobuf:"bytes,1,rep,name=c,proto3" json:"c,omitempty"`
}

func (x *Chunks) Reset() {
	*x = Chunks{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sound_app_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Chunks) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Chunks) ProtoMessage() {}

func (x *Chunks) ProtoReflect() protoreflect.Message {
	mi := &file_sound_app_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Chunks.ProtoReflect.Descriptor instead.
func (*Chunks) Descriptor() ([]byte, []int) {
	return file_sound_app_proto_rawDescGZIP(), []int{5}
}

func (x *Chunks) GetC() []*Chunk {
	if x != nil {
		return x.C
	}
	return nil
}

type SongData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to Data:
	//
	//	*SongData_Info
	//	*SongData_Chunks
	Data isSongData_Data `protobuf_oneof:"Data"`
}

func (x *SongData) Reset() {
	*x = SongData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_sound_app_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SongData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SongData) ProtoMessage() {}

func (x *SongData) ProtoReflect() protoreflect.Message {
	mi := &file_sound_app_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SongData.ProtoReflect.Descriptor instead.
func (*SongData) Descriptor() ([]byte, []int) {
	return file_sound_app_proto_rawDescGZIP(), []int{6}
}

func (m *SongData) GetData() isSongData_Data {
	if m != nil {
		return m.Data
	}
	return nil
}

func (x *SongData) GetInfo() *SongInfo {
	if x, ok := x.GetData().(*SongData_Info); ok {
		return x.Info
	}
	return nil
}

func (x *SongData) GetChunks() *Chunks {
	if x, ok := x.GetData().(*SongData_Chunks); ok {
		return x.Chunks
	}
	return nil
}

type isSongData_Data interface {
	isSongData_Data()
}

type SongData_Info struct {
	Info *SongInfo `protobuf:"bytes,1,opt,name=info,proto3,oneof"`
}

type SongData_Chunks struct {
	Chunks *Chunks `protobuf:"bytes,2,opt,name=chunks,proto3,oneof"`
}

func (*SongData_Info) isSongData_Data() {}

func (*SongData_Chunks) isSongData_Data() {}

var File_sound_app_proto protoreflect.FileDescriptor

var file_sound_app_proto_rawDesc = []byte{
	0x0a, 0x0f, 0x73, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x61, 0x70, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x08, 0x0a, 0x06, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c, 0x22, 0x43, 0x0a, 0x0b, 0x53,
	0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x10, 0x0a, 0x02, 0x69, 0x64,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x48, 0x00, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x01,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x6c,
	0x48, 0x00, 0x52, 0x01, 0x73, 0x42, 0x09, 0x0a, 0x07, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x22, 0xb8, 0x01, 0x0a, 0x08, 0x53, 0x6f, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x1f, 0x0a, 0x0b, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x61, 0x72, 0x74, 0x69, 0x73, 0x74, 0x4e, 0x61,
	0x6d, 0x65, 0x12, 0x1d, 0x0a, 0x0a, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x5f, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x61, 0x6c, 0x62, 0x75, 0x6d, 0x4e, 0x61, 0x6d,
	0x65, 0x12, 0x21, 0x0a, 0x0c, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x5f, 0x6e, 0x75, 0x6d, 0x62, 0x65,
	0x72, 0x18, 0x05, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0b, 0x74, 0x72, 0x61, 0x63, 0x6b, 0x4e, 0x75,
	0x6d, 0x62, 0x65, 0x72, 0x12, 0x25, 0x0a, 0x0e, 0x70, 0x75, 0x62, 0x6c, 0x69, 0x73, 0x68, 0x65,
	0x64, 0x5f, 0x79, 0x65, 0x61, 0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0d, 0x70, 0x75,
	0x62, 0x6c, 0x69, 0x73, 0x68, 0x65, 0x64, 0x59, 0x65, 0x61, 0x72, 0x22, 0x2a, 0x0a, 0x07, 0x41,
	0x6c, 0x6c, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x1f, 0x0a, 0x05, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x18,
	0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f,
	0x52, 0x05, 0x73, 0x6f, 0x6e, 0x67, 0x73, 0x22, 0x31, 0x0a, 0x05, 0x43, 0x68, 0x75, 0x6e, 0x6b,
	0x12, 0x12, 0x0a, 0x04, 0x6c, 0x65, 0x66, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x01, 0x52, 0x04,
	0x6c, 0x65, 0x66, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x72, 0x69, 0x67, 0x68, 0x74, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x01, 0x52, 0x05, 0x72, 0x69, 0x67, 0x68, 0x74, 0x22, 0x1e, 0x0a, 0x06, 0x43, 0x68,
	0x75, 0x6e, 0x6b, 0x73, 0x12, 0x14, 0x0a, 0x01, 0x63, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32,
	0x06, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x52, 0x01, 0x63, 0x22, 0x56, 0x0a, 0x08, 0x53, 0x6f,
	0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1f, 0x0a, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x09, 0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x49, 0x6e, 0x66, 0x6f, 0x48,
	0x00, 0x52, 0x04, 0x69, 0x6e, 0x66, 0x6f, 0x12, 0x21, 0x0a, 0x06, 0x63, 0x68, 0x75, 0x6e, 0x6b,
	0x73, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x07, 0x2e, 0x43, 0x68, 0x75, 0x6e, 0x6b, 0x73,
	0x48, 0x00, 0x52, 0x06, 0x63, 0x68, 0x75, 0x6e, 0x6b, 0x73, 0x42, 0x06, 0x0a, 0x04, 0x44, 0x61,
	0x74, 0x61, 0x32, 0x55, 0x0a, 0x06, 0x50, 0x6c, 0x61, 0x79, 0x65, 0x72, 0x12, 0x24, 0x0a, 0x0a,
	0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x0c, 0x2e, 0x53, 0x6f, 0x6e,
	0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x08, 0x2e, 0x41, 0x6c, 0x6c, 0x53, 0x6f,
	0x6e, 0x67, 0x12, 0x25, 0x0a, 0x08, 0x50, 0x6c, 0x61, 0x79, 0x53, 0x6f, 0x6e, 0x67, 0x12, 0x0c,
	0x2e, 0x53, 0x6f, 0x6e, 0x67, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x09, 0x2e, 0x53,
	0x6f, 0x6e, 0x67, 0x44, 0x61, 0x74, 0x61, 0x30, 0x01, 0x42, 0x0b, 0x5a, 0x09, 0x2f, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_sound_app_proto_rawDescOnce sync.Once
	file_sound_app_proto_rawDescData = file_sound_app_proto_rawDesc
)

func file_sound_app_proto_rawDescGZIP() []byte {
	file_sound_app_proto_rawDescOnce.Do(func() {
		file_sound_app_proto_rawDescData = protoimpl.X.CompressGZIP(file_sound_app_proto_rawDescData)
	})
	return file_sound_app_proto_rawDescData
}

var file_sound_app_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_sound_app_proto_goTypes = []interface{}{
	(*Signal)(nil),      // 0: Signal
	(*SongRequest)(nil), // 1: SongRequest
	(*SongInfo)(nil),    // 2: SongInfo
	(*AllSong)(nil),     // 3: AllSong
	(*Chunk)(nil),       // 4: Chunk
	(*Chunks)(nil),      // 5: Chunks
	(*SongData)(nil),    // 6: SongData
}
var file_sound_app_proto_depIdxs = []int32{
	0, // 0: SongRequest.s:type_name -> Signal
	2, // 1: AllSong.songs:type_name -> SongInfo
	4, // 2: Chunks.c:type_name -> Chunk
	2, // 3: SongData.info:type_name -> SongInfo
	5, // 4: SongData.chunks:type_name -> Chunks
	1, // 5: Player.GetAllSong:input_type -> SongRequest
	1, // 6: Player.PlaySong:input_type -> SongRequest
	3, // 7: Player.GetAllSong:output_type -> AllSong
	6, // 8: Player.PlaySong:output_type -> SongData
	7, // [7:9] is the sub-list for method output_type
	5, // [5:7] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_sound_app_proto_init() }
func file_sound_app_proto_init() {
	if File_sound_app_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_sound_app_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Signal); i {
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
		file_sound_app_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SongRequest); i {
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
		file_sound_app_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SongInfo); i {
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
		file_sound_app_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AllSong); i {
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
		file_sound_app_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunk); i {
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
		file_sound_app_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Chunks); i {
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
		file_sound_app_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SongData); i {
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
	file_sound_app_proto_msgTypes[1].OneofWrappers = []interface{}{
		(*SongRequest_Id)(nil),
		(*SongRequest_S)(nil),
	}
	file_sound_app_proto_msgTypes[6].OneofWrappers = []interface{}{
		(*SongData_Info)(nil),
		(*SongData_Chunks)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_sound_app_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_sound_app_proto_goTypes,
		DependencyIndexes: file_sound_app_proto_depIdxs,
		MessageInfos:      file_sound_app_proto_msgTypes,
	}.Build()
	File_sound_app_proto = out.File
	file_sound_app_proto_rawDesc = nil
	file_sound_app_proto_goTypes = nil
	file_sound_app_proto_depIdxs = nil
}
