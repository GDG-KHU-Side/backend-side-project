// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.35.2
// 	protoc        v3.20.3
// source: proto/reservation.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Reservation struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id               int64                  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	RestId           int64                  `protobuf:"varint,2,opt,name=rest_id,json=restId,proto3" json:"rest_id,omitempty"`
	Count            int32                  `protobuf:"varint,3,opt,name=count,proto3" json:"count,omitempty"`
	CustomerPhoneNum string                 `protobuf:"bytes,4,opt,name=customer_phone_num,json=customerPhoneNum,proto3" json:"customer_phone_num,omitempty"`
	CreatedAt        *timestamppb.Timestamp `protobuf:"bytes,5,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	IsEntry          int32                  `protobuf:"varint,6,opt,name=is_entry,json=isEntry,proto3" json:"is_entry,omitempty"`
}

func (x *Reservation) Reset() {
	*x = Reservation{}
	mi := &file_proto_reservation_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *Reservation) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Reservation) ProtoMessage() {}

func (x *Reservation) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reservation_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Reservation.ProtoReflect.Descriptor instead.
func (*Reservation) Descriptor() ([]byte, []int) {
	return file_proto_reservation_proto_rawDescGZIP(), []int{0}
}

func (x *Reservation) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Reservation) GetRestId() int64 {
	if x != nil {
		return x.RestId
	}
	return 0
}

func (x *Reservation) GetCount() int32 {
	if x != nil {
		return x.Count
	}
	return 0
}

func (x *Reservation) GetCustomerPhoneNum() string {
	if x != nil {
		return x.CustomerPhoneNum
	}
	return ""
}

func (x *Reservation) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *Reservation) GetIsEntry() int32 {
	if x != nil {
		return x.IsEntry
	}
	return 0
}

type GetReservationsRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	RestaurantId int64 `protobuf:"varint,1,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
}

func (x *GetReservationsRequest) Reset() {
	*x = GetReservationsRequest{}
	mi := &file_proto_reservation_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetReservationsRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReservationsRequest) ProtoMessage() {}

func (x *GetReservationsRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reservation_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReservationsRequest.ProtoReflect.Descriptor instead.
func (*GetReservationsRequest) Descriptor() ([]byte, []int) {
	return file_proto_reservation_proto_rawDescGZIP(), []int{1}
}

func (x *GetReservationsRequest) GetRestaurantId() int64 {
	if x != nil {
		return x.RestaurantId
	}
	return 0
}

type GetReservationsResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Reservations []*Reservation `protobuf:"bytes,1,rep,name=reservations,proto3" json:"reservations,omitempty"`
}

func (x *GetReservationsResponse) Reset() {
	*x = GetReservationsResponse{}
	mi := &file_proto_reservation_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetReservationsResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetReservationsResponse) ProtoMessage() {}

func (x *GetReservationsResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_reservation_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetReservationsResponse.ProtoReflect.Descriptor instead.
func (*GetReservationsResponse) Descriptor() ([]byte, []int) {
	return file_proto_reservation_proto_rawDescGZIP(), []int{2}
}

func (x *GetReservationsResponse) GetReservations() []*Reservation {
	if x != nil {
		return x.Reservations
	}
	return nil
}

var File_proto_reservation_proto protoreflect.FileDescriptor

var file_proto_reservation_proto_rawDesc = []byte{
	0x0a, 0x17, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x72, 0x65, 0x73, 0x65, 0x72,
	0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x1a, 0x1f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69, 0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d,
	0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xd0, 0x01, 0x0a, 0x0b, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x72, 0x65, 0x73, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x06, 0x72, 0x65, 0x73, 0x74, 0x49, 0x64,
	0x12, 0x14, 0x0a, 0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x05, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x2c, 0x0a, 0x12, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d,
	0x65, 0x72, 0x5f, 0x70, 0x68, 0x6f, 0x6e, 0x65, 0x5f, 0x6e, 0x75, 0x6d, 0x18, 0x04, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x10, 0x63, 0x75, 0x73, 0x74, 0x6f, 0x6d, 0x65, 0x72, 0x50, 0x68, 0x6f, 0x6e,
	0x65, 0x4e, 0x75, 0x6d, 0x12, 0x39, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f,
	0x61, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c,
	0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x54, 0x69, 0x6d, 0x65, 0x73,
	0x74, 0x61, 0x6d, 0x70, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74, 0x12,
	0x19, 0x0a, 0x08, 0x69, 0x73, 0x5f, 0x65, 0x6e, 0x74, 0x72, 0x79, 0x18, 0x06, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x07, 0x69, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x22, 0x3d, 0x0a, 0x16, 0x47, 0x65,
	0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x22, 0x57, 0x0a, 0x17, 0x47, 0x65, 0x74,
	0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x12, 0x3c, 0x0a, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x73, 0x32, 0x7c, 0x0a, 0x12, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x66, 0x0a, 0x19, 0x47, 0x65, 0x74, 0x52,
	0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x23, 0x2e, 0x72, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65, 0x72, 0x76, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x24, 0x2e, 0x72, 0x65, 0x73,
	0x65, 0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x47, 0x65, 0x74, 0x52, 0x65, 0x73, 0x65,
	0x72, 0x76, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x42, 0x34, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x47,
	0x44, 0x47, 0x2d, 0x4b, 0x48, 0x55, 0x2d, 0x53, 0x69, 0x64, 0x65, 0x2f, 0x62, 0x61, 0x63, 0x6b,
	0x65, 0x6e, 0x64, 0x2d, 0x73, 0x69, 0x64, 0x65, 0x2d, 0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74,
	0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_reservation_proto_rawDescOnce sync.Once
	file_proto_reservation_proto_rawDescData = file_proto_reservation_proto_rawDesc
)

func file_proto_reservation_proto_rawDescGZIP() []byte {
	file_proto_reservation_proto_rawDescOnce.Do(func() {
		file_proto_reservation_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_reservation_proto_rawDescData)
	})
	return file_proto_reservation_proto_rawDescData
}

var file_proto_reservation_proto_msgTypes = make([]protoimpl.MessageInfo, 3)
var file_proto_reservation_proto_goTypes = []any{
	(*Reservation)(nil),             // 0: reservation.Reservation
	(*GetReservationsRequest)(nil),  // 1: reservation.GetReservationsRequest
	(*GetReservationsResponse)(nil), // 2: reservation.GetReservationsResponse
	(*timestamppb.Timestamp)(nil),   // 3: google.protobuf.Timestamp
}
var file_proto_reservation_proto_depIdxs = []int32{
	3, // 0: reservation.Reservation.created_at:type_name -> google.protobuf.Timestamp
	0, // 1: reservation.GetReservationsResponse.reservations:type_name -> reservation.Reservation
	1, // 2: reservation.ReservationService.GetRestaurantReservations:input_type -> reservation.GetReservationsRequest
	2, // 3: reservation.ReservationService.GetRestaurantReservations:output_type -> reservation.GetReservationsResponse
	3, // [3:4] is the sub-list for method output_type
	2, // [2:3] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_proto_reservation_proto_init() }
func file_proto_reservation_proto_init() {
	if File_proto_reservation_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_proto_reservation_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   3,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_reservation_proto_goTypes,
		DependencyIndexes: file_proto_reservation_proto_depIdxs,
		MessageInfos:      file_proto_reservation_proto_msgTypes,
	}.Build()
	File_proto_reservation_proto = out.File
	file_proto_reservation_proto_rawDesc = nil
	file_proto_reservation_proto_goTypes = nil
	file_proto_reservation_proto_depIdxs = nil
}