// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        (unknown)
// source: commurzpb/v1/commurzpb.proto

package commurzpbv1

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	_ "google.golang.org/protobuf/types/known/timestamppb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Resource definition
type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id    string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Email string `protobuf:"bytes,2,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_commurzpb_v1_commurzpb_proto_rawDescGZIP(), []int{0}
}

func (x *User) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *User) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type Product struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price        int64  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	CurrentStock int64  `protobuf:"varint,4,opt,name=current_stock,json=currentStock,proto3" json:"current_stock,omitempty"`
}

func (x *Product) Reset() {
	*x = Product{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Product) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Product) ProtoMessage() {}

func (x *Product) ProtoReflect() protoreflect.Message {
	mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Product.ProtoReflect.Descriptor instead.
func (*Product) Descriptor() ([]byte, []int) {
	return file_commurzpb_v1_commurzpb_proto_rawDescGZIP(), []int{1}
}

func (x *Product) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Product) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Product) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *Product) GetCurrentStock() int64 {
	if x != nil {
		return x.CurrentStock
	}
	return 0
}

type OrderProduct struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name         string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	Price        int64  `protobuf:"varint,3,opt,name=price,proto3" json:"price,omitempty"`
	CurrentStock int64  `protobuf:"varint,4,opt,name=current_stock,json=currentStock,proto3" json:"current_stock,omitempty"`
}

func (x *OrderProduct) Reset() {
	*x = OrderProduct{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderProduct) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderProduct) ProtoMessage() {}

func (x *OrderProduct) ProtoReflect() protoreflect.Message {
	mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderProduct.ProtoReflect.Descriptor instead.
func (*OrderProduct) Descriptor() ([]byte, []int) {
	return file_commurzpb_v1_commurzpb_proto_rawDescGZIP(), []int{2}
}

func (x *OrderProduct) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderProduct) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *OrderProduct) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

func (x *OrderProduct) GetCurrentStock() int64 {
	if x != nil {
		return x.CurrentStock
	}
	return 0
}

type Cart struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id     string      `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId string      `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	User   *User       `protobuf:"bytes,3,opt,name=user,proto3" json:"user,omitempty"`
	Items  []*CartItem `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
}

func (x *Cart) Reset() {
	*x = Cart{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Cart) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Cart) ProtoMessage() {}

func (x *Cart) ProtoReflect() protoreflect.Message {
	mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Cart.ProtoReflect.Descriptor instead.
func (*Cart) Descriptor() ([]byte, []int) {
	return file_commurzpb_v1_commurzpb_proto_rawDescGZIP(), []int{3}
}

func (x *Cart) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Cart) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Cart) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Cart) GetItems() []*CartItem {
	if x != nil {
		return x.Items
	}
	return nil
}

type CartItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	CartId       string        `protobuf:"bytes,2,opt,name=cart_id,json=cartId,proto3" json:"cart_id,omitempty"`
	ProductId    string        `protobuf:"bytes,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity     int64         `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	ProductPrice int64         `protobuf:"varint,5,opt,name=product_price,json=productPrice,proto3" json:"product_price,omitempty"`
	Product      *OrderProduct `protobuf:"bytes,6,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *CartItem) Reset() {
	*x = CartItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CartItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CartItem) ProtoMessage() {}

func (x *CartItem) ProtoReflect() protoreflect.Message {
	mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CartItem.ProtoReflect.Descriptor instead.
func (*CartItem) Descriptor() ([]byte, []int) {
	return file_commurzpb_v1_commurzpb_proto_rawDescGZIP(), []int{4}
}

func (x *CartItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *CartItem) GetCartId() string {
	if x != nil {
		return x.CartId
	}
	return ""
}

func (x *CartItem) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *CartItem) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *CartItem) GetProductPrice() int64 {
	if x != nil {
		return x.ProductPrice
	}
	return 0
}

func (x *CartItem) GetProduct() *OrderProduct {
	if x != nil {
		return x.Product
	}
	return nil
}

type Order struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	UserId     string `protobuf:"bytes,2,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"`
	TotalPrice int64  `protobuf:"varint,3,opt,name=total_price,json=totalPrice,proto3" json:"total_price,omitempty"`
}

func (x *Order) Reset() {
	*x = Order{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Order) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Order) ProtoMessage() {}

func (x *Order) ProtoReflect() protoreflect.Message {
	mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Order.ProtoReflect.Descriptor instead.
func (*Order) Descriptor() ([]byte, []int) {
	return file_commurzpb_v1_commurzpb_proto_rawDescGZIP(), []int{5}
}

func (x *Order) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Order) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

func (x *Order) GetTotalPrice() int64 {
	if x != nil {
		return x.TotalPrice
	}
	return 0
}

type OrderItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id           string        `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	OrderId      string        `protobuf:"bytes,2,opt,name=order_id,json=orderId,proto3" json:"order_id,omitempty"`
	ProductId    string        `protobuf:"bytes,3,opt,name=product_id,json=productId,proto3" json:"product_id,omitempty"`
	Quantity     int64         `protobuf:"varint,4,opt,name=quantity,proto3" json:"quantity,omitempty"`
	ProductPrice int64         `protobuf:"varint,5,opt,name=product_price,json=productPrice,proto3" json:"product_price,omitempty"`
	Product      *OrderProduct `protobuf:"bytes,6,opt,name=product,proto3" json:"product,omitempty"`
}

func (x *OrderItem) Reset() {
	*x = OrderItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *OrderItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*OrderItem) ProtoMessage() {}

func (x *OrderItem) ProtoReflect() protoreflect.Message {
	mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use OrderItem.ProtoReflect.Descriptor instead.
func (*OrderItem) Descriptor() ([]byte, []int) {
	return file_commurzpb_v1_commurzpb_proto_rawDescGZIP(), []int{6}
}

func (x *OrderItem) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *OrderItem) GetOrderId() string {
	if x != nil {
		return x.OrderId
	}
	return ""
}

func (x *OrderItem) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *OrderItem) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *OrderItem) GetProductPrice() int64 {
	if x != nil {
		return x.ProductPrice
	}
	return 0
}

func (x *OrderItem) GetProduct() *OrderProduct {
	if x != nil {
		return x.Product
	}
	return nil
}

// Requests definition
type CreateUserRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Email string `protobuf:"bytes,1,opt,name=email,proto3" json:"email,omitempty"`
}

func (x *CreateUserRequest) Reset() {
	*x = CreateUserRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateUserRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateUserRequest) ProtoMessage() {}

func (x *CreateUserRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateUserRequest.ProtoReflect.Descriptor instead.
func (*CreateUserRequest) Descriptor() ([]byte, []int) {
	return file_commurzpb_v1_commurzpb_proto_rawDescGZIP(), []int{7}
}

func (x *CreateUserRequest) GetEmail() string {
	if x != nil {
		return x.Email
	}
	return ""
}

type CreateProductRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name  string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Price int64  `protobuf:"varint,2,opt,name=price,proto3" json:"price,omitempty"`
}

func (x *CreateProductRequest) Reset() {
	*x = CreateProductRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateProductRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateProductRequest) ProtoMessage() {}

func (x *CreateProductRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateProductRequest.ProtoReflect.Descriptor instead.
func (*CreateProductRequest) Descriptor() ([]byte, []int) {
	return file_commurzpb_v1_commurzpb_proto_rawDescGZIP(), []int{8}
}

func (x *CreateProductRequest) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *CreateProductRequest) GetPrice() int64 {
	if x != nil {
		return x.Price
	}
	return 0
}

type ChangeProductStockRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId     string `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
	StockQuantity int64  `protobuf:"varint,2,opt,name=stock_quantity,json=stockQuantity,proto3" json:"stock_quantity,omitempty"`
}

func (x *ChangeProductStockRequest) Reset() {
	*x = ChangeProductStockRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[9]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChangeProductStockRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChangeProductStockRequest) ProtoMessage() {}

func (x *ChangeProductStockRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[9]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChangeProductStockRequest.ProtoReflect.Descriptor instead.
func (*ChangeProductStockRequest) Descriptor() ([]byte, []int) {
	return file_commurzpb_v1_commurzpb_proto_rawDescGZIP(), []int{9}
}

func (x *ChangeProductStockRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *ChangeProductStockRequest) GetStockQuantity() int64 {
	if x != nil {
		return x.StockQuantity
	}
	return 0
}

type AddItemToCartRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ProductId string `protobuf:"bytes,1,opt,name=productId,proto3" json:"productId,omitempty"`
	Quantity  int64  `protobuf:"varint,2,opt,name=quantity,proto3" json:"quantity,omitempty"`
	UserId    string `protobuf:"bytes,3,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // TODO: remove this, change to auth token/session
}

func (x *AddItemToCartRequest) Reset() {
	*x = AddItemToCartRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[10]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *AddItemToCartRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*AddItemToCartRequest) ProtoMessage() {}

func (x *AddItemToCartRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[10]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use AddItemToCartRequest.ProtoReflect.Descriptor instead.
func (*AddItemToCartRequest) Descriptor() ([]byte, []int) {
	return file_commurzpb_v1_commurzpb_proto_rawDescGZIP(), []int{10}
}

func (x *AddItemToCartRequest) GetProductId() string {
	if x != nil {
		return x.ProductId
	}
	return ""
}

func (x *AddItemToCartRequest) GetQuantity() int64 {
	if x != nil {
		return x.Quantity
	}
	return 0
}

func (x *AddItemToCartRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

type CheckoutAllRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	UserId string `protobuf:"bytes,1,opt,name=user_id,json=userId,proto3" json:"user_id,omitempty"` // TODO: remove this, change to auth token/session
}

func (x *CheckoutAllRequest) Reset() {
	*x = CheckoutAllRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[11]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CheckoutAllRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CheckoutAllRequest) ProtoMessage() {}

func (x *CheckoutAllRequest) ProtoReflect() protoreflect.Message {
	mi := &file_commurzpb_v1_commurzpb_proto_msgTypes[11]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CheckoutAllRequest.ProtoReflect.Descriptor instead.
func (*CheckoutAllRequest) Descriptor() ([]byte, []int) {
	return file_commurzpb_v1_commurzpb_proto_rawDescGZIP(), []int{11}
}

func (x *CheckoutAllRequest) GetUserId() string {
	if x != nil {
		return x.UserId
	}
	return ""
}

var File_commurzpb_v1_commurzpb_proto protoreflect.FileDescriptor

var file_commurzpb_v1_commurzpb_proto_rawDesc = []byte{
	0x0a, 0x1c, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x72, 0x7a, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x2f, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x72, 0x7a, 0x70, 0x62, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0c,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x72, 0x7a, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x74, 0x69,
	0x6d, 0x65, 0x73, 0x74, 0x61, 0x6d, 0x70, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x2c, 0x0a,
	0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x68, 0x0a, 0x07, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72,
	0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x12, 0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74,
	0x53, 0x74, 0x6f, 0x63, 0x6b, 0x22, 0x6d, 0x0a, 0x0c, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72,
	0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x12,
	0x23, 0x0a, 0x0d, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x5f, 0x73, 0x74, 0x6f, 0x63, 0x6b,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x63, 0x75, 0x72, 0x72, 0x65, 0x6e, 0x74, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x22, 0x85, 0x01, 0x0a, 0x04, 0x43, 0x61, 0x72, 0x74, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a,
	0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06,
	0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x26, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x72, 0x7a, 0x70, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52, 0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x2c,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x16, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x72, 0x7a, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72,
	0x74, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x22, 0xc9, 0x01, 0x0a,
	0x08, 0x43, 0x61, 0x72, 0x74, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x17, 0x0a, 0x07, 0x63, 0x61, 0x72,
	0x74, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x63, 0x61, 0x72, 0x74,
	0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x69, 0x64,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x04, 0x20,
	0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x23, 0x0a,
	0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50, 0x72, 0x69,
	0x63, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x72, 0x7a, 0x70, 0x62, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52,
	0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x51, 0x0a, 0x05, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69,
	0x64, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65, 0x72, 0x49, 0x64, 0x12, 0x1f, 0x0a, 0x0b, 0x74, 0x6f,
	0x74, 0x61, 0x6c, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52,
	0x0a, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x50, 0x72, 0x69, 0x63, 0x65, 0x22, 0xcc, 0x01, 0x0a, 0x09,
	0x4f, 0x72, 0x64, 0x65, 0x72, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x19, 0x0a, 0x08, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x6f, 0x72, 0x64,
	0x65, 0x72, 0x49, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f,
	0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x49, 0x64, 0x12, 0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12,
	0x23, 0x0a, 0x0d, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x05, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0c, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x50,
	0x72, 0x69, 0x63, 0x65, 0x12, 0x34, 0x0a, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x18,
	0x06, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x72, 0x7a, 0x70,
	0x62, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x52, 0x07, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x29, 0x0a, 0x11, 0x43, 0x72,
	0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x65, 0x6d, 0x61, 0x69, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x65, 0x6d, 0x61, 0x69, 0x6c, 0x22, 0x40, 0x0a, 0x14, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x50,
	0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x14, 0x0a, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03,
	0x52, 0x05, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x60, 0x0a, 0x19, 0x43, 0x68, 0x61, 0x6e, 0x67,
	0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49,
	0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74,
	0x49, 0x64, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x6f, 0x63, 0x6b, 0x5f, 0x71, 0x75, 0x61, 0x6e,
	0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x73, 0x74, 0x6f, 0x63,
	0x6b, 0x51, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x22, 0x69, 0x0a, 0x14, 0x41, 0x64, 0x64,
	0x49, 0x74, 0x65, 0x6d, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x49, 0x64, 0x12,
	0x1a, 0x0a, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x08, 0x71, 0x75, 0x61, 0x6e, 0x74, 0x69, 0x74, 0x79, 0x12, 0x17, 0x0a, 0x07, 0x75,
	0x73, 0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73,
	0x65, 0x72, 0x49, 0x64, 0x22, 0x2d, 0x0a, 0x12, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74,
	0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x75, 0x73,
	0x65, 0x72, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x75, 0x73, 0x65,
	0x72, 0x49, 0x64, 0x32, 0xe3, 0x03, 0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x75, 0x72, 0x7a, 0x53,
	0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x43, 0x0a, 0x0a, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65,
	0x55, 0x73, 0x65, 0x72, 0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x72, 0x7a, 0x70, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x55, 0x73, 0x65, 0x72, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x72, 0x7a, 0x70,
	0x62, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x22, 0x00, 0x12, 0x4c, 0x0a, 0x0d, 0x43,
	0x72, 0x65, 0x61, 0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x12, 0x22, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x72, 0x7a, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x72, 0x7a, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x53, 0x0a, 0x0f, 0x41, 0x64, 0x64,
	0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x12, 0x27, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x72, 0x7a, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e,
	0x67, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65,
	0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x72, 0x7a, 0x70,
	0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x56,
	0x0a, 0x12, 0x52, 0x65, 0x64, 0x75, 0x63, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63, 0x74, 0x53,
	0x74, 0x6f, 0x63, 0x6b, 0x12, 0x27, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x72, 0x7a, 0x70, 0x62,
	0x2e, 0x76, 0x31, 0x2e, 0x43, 0x68, 0x61, 0x6e, 0x67, 0x65, 0x50, 0x72, 0x6f, 0x64, 0x75, 0x63,
	0x74, 0x53, 0x74, 0x6f, 0x63, 0x6b, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x15, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x72, 0x7a, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x50, 0x72, 0x6f,
	0x64, 0x75, 0x63, 0x74, 0x22, 0x00, 0x12, 0x49, 0x0a, 0x0d, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65,
	0x6d, 0x54, 0x6f, 0x43, 0x61, 0x72, 0x74, 0x12, 0x22, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x72,
	0x7a, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x41, 0x64, 0x64, 0x49, 0x74, 0x65, 0x6d, 0x54, 0x6f,
	0x43, 0x61, 0x72, 0x74, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x12, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x75, 0x72, 0x7a, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x72, 0x74, 0x22,
	0x00, 0x12, 0x46, 0x0a, 0x0b, 0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x41, 0x6c, 0x6c,
	0x12, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x72, 0x7a, 0x70, 0x62, 0x2e, 0x76, 0x31, 0x2e,
	0x43, 0x68, 0x65, 0x63, 0x6b, 0x6f, 0x75, 0x74, 0x41, 0x6c, 0x6c, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x13, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x72, 0x7a, 0x70, 0x62, 0x2e, 0x76,
	0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x22, 0x00, 0x42, 0x3f, 0x5a, 0x3d, 0x67, 0x69, 0x74,
	0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x61, 0x68, 0x6d, 0x69, 0x66, 0x61, 0x6e,
	0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x72, 0x7a, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x67, 0x65,
	0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x75, 0x72, 0x7a, 0x70, 0x62, 0x2f, 0x76, 0x31, 0x3b, 0x63,
	0x6f, 0x6d, 0x6d, 0x75, 0x72, 0x7a, 0x70, 0x62, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x33,
}

var (
	file_commurzpb_v1_commurzpb_proto_rawDescOnce sync.Once
	file_commurzpb_v1_commurzpb_proto_rawDescData = file_commurzpb_v1_commurzpb_proto_rawDesc
)

func file_commurzpb_v1_commurzpb_proto_rawDescGZIP() []byte {
	file_commurzpb_v1_commurzpb_proto_rawDescOnce.Do(func() {
		file_commurzpb_v1_commurzpb_proto_rawDescData = protoimpl.X.CompressGZIP(file_commurzpb_v1_commurzpb_proto_rawDescData)
	})
	return file_commurzpb_v1_commurzpb_proto_rawDescData
}

var file_commurzpb_v1_commurzpb_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_commurzpb_v1_commurzpb_proto_goTypes = []interface{}{
	(*User)(nil),                      // 0: commurzpb.v1.User
	(*Product)(nil),                   // 1: commurzpb.v1.Product
	(*OrderProduct)(nil),              // 2: commurzpb.v1.OrderProduct
	(*Cart)(nil),                      // 3: commurzpb.v1.Cart
	(*CartItem)(nil),                  // 4: commurzpb.v1.CartItem
	(*Order)(nil),                     // 5: commurzpb.v1.Order
	(*OrderItem)(nil),                 // 6: commurzpb.v1.OrderItem
	(*CreateUserRequest)(nil),         // 7: commurzpb.v1.CreateUserRequest
	(*CreateProductRequest)(nil),      // 8: commurzpb.v1.CreateProductRequest
	(*ChangeProductStockRequest)(nil), // 9: commurzpb.v1.ChangeProductStockRequest
	(*AddItemToCartRequest)(nil),      // 10: commurzpb.v1.AddItemToCartRequest
	(*CheckoutAllRequest)(nil),        // 11: commurzpb.v1.CheckoutAllRequest
}
var file_commurzpb_v1_commurzpb_proto_depIdxs = []int32{
	0,  // 0: commurzpb.v1.Cart.user:type_name -> commurzpb.v1.User
	4,  // 1: commurzpb.v1.Cart.items:type_name -> commurzpb.v1.CartItem
	2,  // 2: commurzpb.v1.CartItem.product:type_name -> commurzpb.v1.OrderProduct
	2,  // 3: commurzpb.v1.OrderItem.product:type_name -> commurzpb.v1.OrderProduct
	7,  // 4: commurzpb.v1.CommurzService.CreateUser:input_type -> commurzpb.v1.CreateUserRequest
	8,  // 5: commurzpb.v1.CommurzService.CreateProduct:input_type -> commurzpb.v1.CreateProductRequest
	9,  // 6: commurzpb.v1.CommurzService.AddProductStock:input_type -> commurzpb.v1.ChangeProductStockRequest
	9,  // 7: commurzpb.v1.CommurzService.ReduceProductStock:input_type -> commurzpb.v1.ChangeProductStockRequest
	10, // 8: commurzpb.v1.CommurzService.AddItemToCart:input_type -> commurzpb.v1.AddItemToCartRequest
	11, // 9: commurzpb.v1.CommurzService.CheckoutAll:input_type -> commurzpb.v1.CheckoutAllRequest
	0,  // 10: commurzpb.v1.CommurzService.CreateUser:output_type -> commurzpb.v1.User
	1,  // 11: commurzpb.v1.CommurzService.CreateProduct:output_type -> commurzpb.v1.Product
	1,  // 12: commurzpb.v1.CommurzService.AddProductStock:output_type -> commurzpb.v1.Product
	1,  // 13: commurzpb.v1.CommurzService.ReduceProductStock:output_type -> commurzpb.v1.Product
	3,  // 14: commurzpb.v1.CommurzService.AddItemToCart:output_type -> commurzpb.v1.Cart
	5,  // 15: commurzpb.v1.CommurzService.CheckoutAll:output_type -> commurzpb.v1.Order
	10, // [10:16] is the sub-list for method output_type
	4,  // [4:10] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_commurzpb_v1_commurzpb_proto_init() }
func file_commurzpb_v1_commurzpb_proto_init() {
	if File_commurzpb_v1_commurzpb_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_commurzpb_v1_commurzpb_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
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
		file_commurzpb_v1_commurzpb_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Product); i {
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
		file_commurzpb_v1_commurzpb_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderProduct); i {
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
		file_commurzpb_v1_commurzpb_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Cart); i {
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
		file_commurzpb_v1_commurzpb_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CartItem); i {
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
		file_commurzpb_v1_commurzpb_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Order); i {
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
		file_commurzpb_v1_commurzpb_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*OrderItem); i {
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
		file_commurzpb_v1_commurzpb_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateUserRequest); i {
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
		file_commurzpb_v1_commurzpb_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateProductRequest); i {
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
		file_commurzpb_v1_commurzpb_proto_msgTypes[9].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChangeProductStockRequest); i {
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
		file_commurzpb_v1_commurzpb_proto_msgTypes[10].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*AddItemToCartRequest); i {
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
		file_commurzpb_v1_commurzpb_proto_msgTypes[11].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CheckoutAllRequest); i {
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
			RawDescriptor: file_commurzpb_v1_commurzpb_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_commurzpb_v1_commurzpb_proto_goTypes,
		DependencyIndexes: file_commurzpb_v1_commurzpb_proto_depIdxs,
		MessageInfos:      file_commurzpb_v1_commurzpb_proto_msgTypes,
	}.Build()
	File_commurzpb_v1_commurzpb_proto = out.File
	file_commurzpb_v1_commurzpb_proto_rawDesc = nil
	file_commurzpb_v1_commurzpb_proto_goTypes = nil
	file_commurzpb_v1_commurzpb_proto_depIdxs = nil
}
