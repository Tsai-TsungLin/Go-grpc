// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        v3.19.4
// source: work.proto

package __

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

type GetProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id int32 `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
}

func (x *GetProductRequest) Reset() {
	*x = GetProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_work_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductRequest) ProtoMessage() {}

func (x *GetProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_work_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductRequest.ProtoReflect.Descriptor instead.
func (*GetProductRequest) Descriptor() ([]byte, []int) {
	return file_work_proto_rawDescGZIP(), []int{0}
}

func (x *GetProductRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

type GetProductReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price     int32  `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	Inventory int32  `protobuf:"varint,3,opt,name=inventory,proto3" json:"inventory,omitempty"`
}

func (x *GetProductReply) Reset() {
	*x = GetProductReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_work_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetProductReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetProductReply) ProtoMessage() {}

func (x *GetProductReply) ProtoReflect() protoreflect.Message {
	mi := &file_work_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetProductReply.ProtoReflect.Descriptor instead.
func (*GetProductReply) Descriptor() ([]byte, []int) {
	return file_work_proto_rawDescGZIP(), []int{1}
}

func (x *GetProductReply) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *GetProductReply) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *GetProductReply) GetInventory() int32 {
	if x != nil {
		return x.Inventory
	}
	return 0
}

type GetAllProductReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	AllProduct map[int32]string `protobuf:"bytes,1,rep,name=AllProduct,proto3" json:"AllProduct,omitempty" protobuf_key:"varint,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *GetAllProductReply) Reset() {
	*x = GetAllProductReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_work_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *GetAllProductReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAllProductReply) ProtoMessage() {}

func (x *GetAllProductReply) ProtoReflect() protoreflect.Message {
	mi := &file_work_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAllProductReply.ProtoReflect.Descriptor instead.
func (*GetAllProductReply) Descriptor() ([]byte, []int) {
	return file_work_proto_rawDescGZIP(), []int{2}
}

func (x *GetAllProductReply) GetAllProduct() map[int32]string {
	if x != nil {
		return x.AllProduct
	}
	return nil
}

type InsertRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name      string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price     int32  `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
	Inventory int32  `protobuf:"varint,3,opt,name=inventory,proto3" json:"inventory,omitempty"`
}

func (x *InsertRequest) Reset() {
	*x = InsertRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_work_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *InsertRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*InsertRequest) ProtoMessage() {}

func (x *InsertRequest) ProtoReflect() protoreflect.Message {
	mi := &file_work_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use InsertRequest.ProtoReflect.Descriptor instead.
func (*InsertRequest) Descriptor() ([]byte, []int) {
	return file_work_proto_rawDescGZIP(), []int{3}
}

func (x *InsertRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *InsertRequest) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *InsertRequest) GetInventory() int32 {
	if x != nil {
		return x.Inventory
	}
	return 0
}

type ModifyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id        int32  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`
	Name      string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price     int32  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	Inventory int32  `protobuf:"varint,4,opt,name=inventory,proto3" json:"inventory,omitempty"`
}

func (x *ModifyRequest) Reset() {
	*x = ModifyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_work_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ModifyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ModifyRequest) ProtoMessage() {}

func (x *ModifyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_work_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ModifyRequest.ProtoReflect.Descriptor instead.
func (*ModifyRequest) Descriptor() ([]byte, []int) {
	return file_work_proto_rawDescGZIP(), []int{4}
}

func (x *ModifyRequest) GetId() int32 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *ModifyRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ModifyRequest) GetPrice() int32 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *ModifyRequest) GetInventory() int32 {
	if x != nil {
		return x.Inventory
	}
	return 0
}

type StatusReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code   int32  `protobuf:"varint,1,opt,name=code,proto3" json:"code,omitempty"`
	Status string `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
}

func (x *StatusReply) Reset() {
	*x = StatusReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_work_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *StatusReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*StatusReply) ProtoMessage() {}

func (x *StatusReply) ProtoReflect() protoreflect.Message {
	mi := &file_work_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use StatusReply.ProtoReflect.Descriptor instead.
func (*StatusReply) Descriptor() ([]byte, []int) {
	return file_work_proto_rawDescGZIP(), []int{5}
}

func (x *StatusReply) GetCode() int32 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *StatusReply) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

type BuyRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Customer   string `protobuf:"bytes,1,opt,name=customer,proto3" json:"customer,omitempty"`
	Production string `protobuf:"bytes,2,opt,name=production,proto3" json:"production,omitempty"`
}

func (x *BuyRequest) Reset() {
	*x = BuyRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_work_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BuyRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BuyRequest) ProtoMessage() {}

func (x *BuyRequest) ProtoReflect() protoreflect.Message {
	mi := &file_work_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BuyRequest.ProtoReflect.Descriptor instead.
func (*BuyRequest) Descriptor() ([]byte, []int) {
	return file_work_proto_rawDescGZIP(), []int{6}
}

func (x *BuyRequest) GetCustomer() string {
	if x != nil {
		return x.Customer
	}
	return ""
}

func (x *BuyRequest) GetProduction() string {
	if x != nil {
		return x.Production
	}
	return ""
}

type TurnoverRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Date string `protobuf:"bytes,1,opt,name=date,proto3" json:"date,omitempty"`
}

func (x *TurnoverRequest) Reset() {
	*x = TurnoverRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_work_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TurnoverRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TurnoverRequest) ProtoMessage() {}

func (x *TurnoverRequest) ProtoReflect() protoreflect.Message {
	mi := &file_work_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TurnoverRequest.ProtoReflect.Descriptor instead.
func (*TurnoverRequest) Descriptor() ([]byte, []int) {
	return file_work_proto_rawDescGZIP(), []int{7}
}

func (x *TurnoverRequest) GetDate() string {
	if x != nil {
		return x.Date
	}
	return ""
}

type TurnoverReply struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductInfo map[string]int32 `protobuf:"bytes,1,rep,name=product_info,json=productInfo,proto3" json:"product_info,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"varint,2,opt,name=value,proto3"`
	TotalPrice  int32            `protobuf:"varint,2,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
}

func (x *TurnoverReply) Reset() {
	*x = TurnoverReply{}
	if protoimpl.UnsafeEnabled {
		mi := &file_work_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TurnoverReply) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TurnoverReply) ProtoMessage() {}

func (x *TurnoverReply) ProtoReflect() protoreflect.Message {
	mi := &file_work_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TurnoverReply.ProtoReflect.Descriptor instead.
func (*TurnoverReply) Descriptor() ([]byte, []int) {
	return file_work_proto_rawDescGZIP(), []int{8}
}

func (x *TurnoverReply) GetProductInfo() map[string]int32 {
	if x != nil {
		return x.ProductInfo
	}
	return nil
}

func (x *TurnoverReply) GetTotalPrice() int32 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

var File_work_proto protoreflect.FileDescriptor

var file_work_proto_rawDesc = []byte{
	0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x02, 0x70, 0x62,
	0x22, 0x23, 0x0a, 0x11, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x05, 0x52, 0x02, 0x69, 0x64, 0x22, 0x59, 0x0a, 0x0f, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64,
	0x75, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x22, 0x9b, 0x01, 0x0a, 0x12, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x46, 0x0a, 0x0a, 0x41, 0x6c, 0x6c, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x70, 0x62,
	0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x2e, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x45, 0x6e,
	0x74, 0x72, 0x79, 0x52, 0x0a, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x1a,
	0x3d, 0x0a, 0x0f, 0x41, 0x6c, 0x6c, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x45, 0x6e, 0x74,
	0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52,
	0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x22, 0x57,
	0x0a, 0x0d, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x76,
	0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x6e,
	0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x22, 0x67, 0x0a, 0x0d, 0x4d, 0x6f, 0x64, 0x69, 0x66,
	0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x05, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05,
	0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x6e, 0x76, 0x65, 0x6e, 0x74, 0x6f, 0x72, 0x79,
	0x22, 0x39, 0x0a, 0x0b, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12,
	0x12, 0x0a, 0x04, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x04, 0x63,
	0x6f, 0x64, 0x65, 0x12, 0x16, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x48, 0x0a, 0x0a, 0x42,
	0x75, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x1a, 0x0a, 0x08, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x08, 0x63, 0x75, 0x73,
	0x74, 0x6f, 0x6d, 0x65, 0x72, 0x12, 0x1e, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x22, 0x25, 0x0a, 0x0f, 0x54, 0x75, 0x72, 0x6e, 0x6f, 0x76, 0x65,
	0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x64, 0x61, 0x74, 0x65, 0x22, 0xb7, 0x01, 0x0a,
	0x0d, 0x54, 0x75, 0x72, 0x6e, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x12, 0x45,
	0x0a, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x6e, 0x66, 0x6f, 0x18, 0x01,
	0x20, 0x03, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x75, 0x72, 0x6e, 0x6f, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x5f, 0x70,
	0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x1a, 0x3e, 0x0a, 0x10, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x6e, 0x66, 0x6f, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05,
	0x76, 0x61, 0x6c, 0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x05, 0x52, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x3a, 0x02, 0x38, 0x01, 0x32, 0xe0, 0x02, 0x0a, 0x0d, 0x53, 0x65, 0x72, 0x76, 0x69,
	0x63, 0x65, 0x53, 0x65, 0x72, 0x76, 0x65, 0x72, 0x12, 0x36, 0x0a, 0x06, 0x53, 0x65, 0x61, 0x72,
	0x63, 0x68, 0x12, 0x15, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75,
	0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x13, 0x2e, 0x70, 0x62, 0x2e, 0x47,
	0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00,
	0x12, 0x3c, 0x0a, 0x09, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x41, 0x6c, 0x6c, 0x12, 0x15, 0x2e,
	0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x16, 0x2e, 0x70, 0x62, 0x2e, 0x47, 0x65, 0x74, 0x41, 0x6c, 0x6c,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x35,
	0x0a, 0x0d, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12,
	0x11, 0x2e, 0x70, 0x62, 0x2e, 0x49, 0x6e, 0x73, 0x65, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65,
	0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x35, 0x0a, 0x0d, 0x4d, 0x6f, 0x64, 0x69, 0x66, 0x79, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x4d, 0x6f, 0x64, 0x69,
	0x66, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e, 0x53,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x2f, 0x0a, 0x0a,
	0x42, 0x75, 0x79, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x2e, 0x70, 0x62, 0x2e,
	0x42, 0x75, 0x79, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x0f, 0x2e, 0x70, 0x62, 0x2e,
	0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x12, 0x3a, 0x0a,
	0x0e, 0x54, 0x75, 0x72, 0x6e, 0x6f, 0x76, 0x65, 0x72, 0x53, 0x65, 0x61, 0x72, 0x63, 0x68, 0x12,
	0x13, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x75, 0x72, 0x6e, 0x6f, 0x76, 0x65, 0x72, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x11, 0x2e, 0x70, 0x62, 0x2e, 0x54, 0x75, 0x72, 0x6e, 0x6f, 0x76,
	0x65, 0x72, 0x52, 0x65, 0x70, 0x6c, 0x79, 0x22, 0x00, 0x42, 0x03, 0x5a, 0x01, 0x2e, 0x62, 0x06,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_work_proto_rawDescOnce sync.Once
	file_work_proto_rawDescData = file_work_proto_rawDesc
)

func file_work_proto_rawDescGZIP() []byte {
	file_work_proto_rawDescOnce.Do(func() {
		file_work_proto_rawDescData = protoimpl.X.CompressGZIP(file_work_proto_rawDescData)
	})
	return file_work_proto_rawDescData
}

var file_work_proto_msgTypes = make([]protoimpl.MessageInfo, 11)
var file_work_proto_goTypes = []interface{}{
	(*GetProductRequest)(nil),  // 0: pb.GetProductRequest
	(*GetProductReply)(nil),    // 1: pb.GetProductReply
	(*GetAllProductReply)(nil), // 2: pb.GetAllProductReply
	(*InsertRequest)(nil),      // 3: pb.InsertRequest
	(*ModifyRequest)(nil),      // 4: pb.ModifyRequest
	(*StatusReply)(nil),        // 5: pb.StatusReply
	(*BuyRequest)(nil),         // 6: pb.BuyRequest
	(*TurnoverRequest)(nil),    // 7: pb.TurnoverRequest
	(*TurnoverReply)(nil),      // 8: pb.TurnoverReply
	nil,                        // 9: pb.GetAllProductReply.AllProductEntry
	nil,                        // 10: pb.TurnoverReply.ProductInfoEntry
}
var file_work_proto_depIdxs = []int32{
	9,  // 0: pb.GetAllProductReply.AllProduct:type_name -> pb.GetAllProductReply.AllProductEntry
	10, // 1: pb.TurnoverReply.product_info:type_name -> pb.TurnoverReply.ProductInfoEntry
	0,  // 2: pb.ServiceServer.Search:input_type -> pb.GetProductRequest
	0,  // 3: pb.ServiceServer.SearchAll:input_type -> pb.GetProductRequest
	3,  // 4: pb.ServiceServer.InsertProduct:input_type -> pb.InsertRequest
	4,  // 5: pb.ServiceServer.ModifyProduct:input_type -> pb.ModifyRequest
	6,  // 6: pb.ServiceServer.BuyProduct:input_type -> pb.BuyRequest
	7,  // 7: pb.ServiceServer.TurnoverSearch:input_type -> pb.TurnoverRequest
	1,  // 8: pb.ServiceServer.Search:output_type -> pb.GetProductReply
	2,  // 9: pb.ServiceServer.SearchAll:output_type -> pb.GetAllProductReply
	5,  // 10: pb.ServiceServer.InsertProduct:output_type -> pb.StatusReply
	5,  // 11: pb.ServiceServer.ModifyProduct:output_type -> pb.StatusReply
	5,  // 12: pb.ServiceServer.BuyProduct:output_type -> pb.StatusReply
	8,  // 13: pb.ServiceServer.TurnoverSearch:output_type -> pb.TurnoverReply
	8,  // [8:14] is the sub-list for method output_type
	2,  // [2:8] is the sub-list for method input_type
	2,  // [2:2] is the sub-list for extension type_name
	2,  // [2:2] is the sub-list for extension extendee
	0,  // [0:2] is the sub-list for field type_name
}

func init() { file_work_proto_init() }
func file_work_proto_init() {
	if File_work_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_work_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductRequest); i {
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
		file_work_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetProductReply); i {
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
		file_work_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*GetAllProductReply); i {
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
		file_work_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*InsertRequest); i {
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
		file_work_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ModifyRequest); i {
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
		file_work_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*StatusReply); i {
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
		file_work_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BuyRequest); i {
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
		file_work_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TurnoverRequest); i {
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
		file_work_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TurnoverReply); i {
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
			RawDescriptor: file_work_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   11,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_work_proto_goTypes,
		DependencyIndexes: file_work_proto_depIdxs,
		MessageInfos:      file_work_proto_msgTypes,
	}.Build()
	File_work_proto = out.File
	file_work_proto_rawDesc = nil
	file_work_proto_goTypes = nil
	file_work_proto_depIdxs = nil
}