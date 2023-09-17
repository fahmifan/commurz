// @generated by protoc-gen-es v1.3.0 with parameter "target=ts"
// @generated from file commurz/v1/commurz.proto (package commurz.v1, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";

/**
 * Resource definition
 *
 * @generated from message commurz.v1.User
 */
export class User extends Message<User> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string email = 2;
   */
  email = "";

  /**
   * @generated from field: string role = 3;
   */
  role = "";

  constructor(data?: PartialMessage<User>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "commurz.v1.User";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "email", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "role", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): User {
    return new User().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): User {
    return new User().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): User {
    return new User().fromJsonString(jsonString, options);
  }

  static equals(a: User | PlainMessage<User> | undefined, b: User | PlainMessage<User> | undefined): boolean {
    return proto3.util.equals(User, a, b);
  }
}

/**
 * @generated from message commurz.v1.Product
 */
export class Product extends Message<Product> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: int64 price = 3;
   */
  price = protoInt64.zero;

  /**
   * @generated from field: int64 current_stock = 4;
   */
  currentStock = protoInt64.zero;

  /**
   * price in text format, e.g. "Rp 10.000"
   *
   * @generated from field: string text_price_idr = 5;
   */
  textPriceIdr = "";

  /**
   * @generated from field: int64 version = 6;
   */
  version = protoInt64.zero;

  constructor(data?: PartialMessage<Product>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "commurz.v1.Product";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "price", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 4, name: "current_stock", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 5, name: "text_price_idr", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "version", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Product {
    return new Product().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Product {
    return new Product().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Product {
    return new Product().fromJsonString(jsonString, options);
  }

  static equals(a: Product | PlainMessage<Product> | undefined, b: Product | PlainMessage<Product> | undefined): boolean {
    return proto3.util.equals(Product, a, b);
  }
}

/**
 * @generated from message commurz.v1.OrderProduct
 */
export class OrderProduct extends Message<OrderProduct> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: int64 price = 3;
   */
  price = protoInt64.zero;

  /**
   * @generated from field: int64 current_stock = 4;
   */
  currentStock = protoInt64.zero;

  constructor(data?: PartialMessage<OrderProduct>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "commurz.v1.OrderProduct";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "price", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 4, name: "current_stock", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OrderProduct {
    return new OrderProduct().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OrderProduct {
    return new OrderProduct().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OrderProduct {
    return new OrderProduct().fromJsonString(jsonString, options);
  }

  static equals(a: OrderProduct | PlainMessage<OrderProduct> | undefined, b: OrderProduct | PlainMessage<OrderProduct> | undefined): boolean {
    return proto3.util.equals(OrderProduct, a, b);
  }
}

/**
 * @generated from message commurz.v1.Cart
 */
export class Cart extends Message<Cart> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string user_id = 2;
   */
  userId = "";

  /**
   * @generated from field: commurz.v1.User user = 3;
   */
  user?: User;

  /**
   * @generated from field: repeated commurz.v1.CartItem items = 4;
   */
  items: CartItem[] = [];

  constructor(data?: PartialMessage<Cart>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "commurz.v1.Cart";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "user", kind: "message", T: User },
    { no: 4, name: "items", kind: "message", T: CartItem, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Cart {
    return new Cart().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Cart {
    return new Cart().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Cart {
    return new Cart().fromJsonString(jsonString, options);
  }

  static equals(a: Cart | PlainMessage<Cart> | undefined, b: Cart | PlainMessage<Cart> | undefined): boolean {
    return proto3.util.equals(Cart, a, b);
  }
}

/**
 * @generated from message commurz.v1.CartItem
 */
export class CartItem extends Message<CartItem> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string cart_id = 2;
   */
  cartId = "";

  /**
   * @generated from field: string product_id = 3;
   */
  productId = "";

  /**
   * @generated from field: int64 quantity = 4;
   */
  quantity = protoInt64.zero;

  /**
   * @generated from field: int64 product_price = 5;
   */
  productPrice = protoInt64.zero;

  /**
   * @generated from field: commurz.v1.OrderProduct product = 6;
   */
  product?: OrderProduct;

  constructor(data?: PartialMessage<CartItem>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "commurz.v1.CartItem";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "cart_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "product_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "quantity", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 5, name: "product_price", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 6, name: "product", kind: "message", T: OrderProduct },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CartItem {
    return new CartItem().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CartItem {
    return new CartItem().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CartItem {
    return new CartItem().fromJsonString(jsonString, options);
  }

  static equals(a: CartItem | PlainMessage<CartItem> | undefined, b: CartItem | PlainMessage<CartItem> | undefined): boolean {
    return proto3.util.equals(CartItem, a, b);
  }
}

/**
 * @generated from message commurz.v1.Order
 */
export class Order extends Message<Order> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string user_id = 2;
   */
  userId = "";

  /**
   * @generated from field: int64 total_price = 3;
   */
  totalPrice = protoInt64.zero;

  constructor(data?: PartialMessage<Order>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "commurz.v1.Order";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "total_price", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Order {
    return new Order().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Order {
    return new Order().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Order {
    return new Order().fromJsonString(jsonString, options);
  }

  static equals(a: Order | PlainMessage<Order> | undefined, b: Order | PlainMessage<Order> | undefined): boolean {
    return proto3.util.equals(Order, a, b);
  }
}

/**
 * @generated from message commurz.v1.OrderItem
 */
export class OrderItem extends Message<OrderItem> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string order_id = 2;
   */
  orderId = "";

  /**
   * @generated from field: string product_id = 3;
   */
  productId = "";

  /**
   * @generated from field: int64 quantity = 4;
   */
  quantity = protoInt64.zero;

  /**
   * @generated from field: int64 product_price = 5;
   */
  productPrice = protoInt64.zero;

  /**
   * @generated from field: commurz.v1.OrderProduct product = 6;
   */
  product?: OrderProduct;

  constructor(data?: PartialMessage<OrderItem>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "commurz.v1.OrderItem";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "order_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "product_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "quantity", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 5, name: "product_price", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 6, name: "product", kind: "message", T: OrderProduct },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): OrderItem {
    return new OrderItem().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): OrderItem {
    return new OrderItem().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): OrderItem {
    return new OrderItem().fromJsonString(jsonString, options);
  }

  static equals(a: OrderItem | PlainMessage<OrderItem> | undefined, b: OrderItem | PlainMessage<OrderItem> | undefined): boolean {
    return proto3.util.equals(OrderItem, a, b);
  }
}

/**
 * Requests definition
 *
 * @generated from message commurz.v1.CreateUserRequest
 */
export class CreateUserRequest extends Message<CreateUserRequest> {
  /**
   * @generated from field: string email = 1;
   */
  email = "";

  /**
   * @generated from field: string password = 2;
   */
  password = "";

  constructor(data?: PartialMessage<CreateUserRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "commurz.v1.CreateUserRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "email", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "password", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateUserRequest {
    return new CreateUserRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateUserRequest {
    return new CreateUserRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateUserRequest {
    return new CreateUserRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateUserRequest | PlainMessage<CreateUserRequest> | undefined, b: CreateUserRequest | PlainMessage<CreateUserRequest> | undefined): boolean {
    return proto3.util.equals(CreateUserRequest, a, b);
  }
}

/**
 * @generated from message commurz.v1.CreateProductRequest
 */
export class CreateProductRequest extends Message<CreateProductRequest> {
  /**
   * @generated from field: string name = 1;
   */
  name = "";

  /**
   * @generated from field: int64 price = 2;
   */
  price = protoInt64.zero;

  constructor(data?: PartialMessage<CreateProductRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "commurz.v1.CreateProductRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "price", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CreateProductRequest {
    return new CreateProductRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CreateProductRequest {
    return new CreateProductRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CreateProductRequest {
    return new CreateProductRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CreateProductRequest | PlainMessage<CreateProductRequest> | undefined, b: CreateProductRequest | PlainMessage<CreateProductRequest> | undefined): boolean {
    return proto3.util.equals(CreateProductRequest, a, b);
  }
}

/**
 * @generated from message commurz.v1.ChangeProductStockRequest
 */
export class ChangeProductStockRequest extends Message<ChangeProductStockRequest> {
  /**
   * @generated from field: string product_id = 1;
   */
  productId = "";

  /**
   * @generated from field: int64 stock_quantity = 2;
   */
  stockQuantity = protoInt64.zero;

  constructor(data?: PartialMessage<ChangeProductStockRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "commurz.v1.ChangeProductStockRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "product_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "stock_quantity", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ChangeProductStockRequest {
    return new ChangeProductStockRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ChangeProductStockRequest {
    return new ChangeProductStockRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ChangeProductStockRequest {
    return new ChangeProductStockRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ChangeProductStockRequest | PlainMessage<ChangeProductStockRequest> | undefined, b: ChangeProductStockRequest | PlainMessage<ChangeProductStockRequest> | undefined): boolean {
    return proto3.util.equals(ChangeProductStockRequest, a, b);
  }
}

/**
 * @generated from message commurz.v1.UpdateProductStockRequest
 */
export class UpdateProductStockRequest extends Message<UpdateProductStockRequest> {
  /**
   * @generated from field: string product_id = 1;
   */
  productId = "";

  /**
   * @generated from field: int64 version = 2;
   */
  version = protoInt64.zero;

  /**
   * @generated from field: int64 stock_in = 3;
   */
  stockIn = protoInt64.zero;

  /**
   * @generated from field: int64 stock_out = 4;
   */
  stockOut = protoInt64.zero;

  constructor(data?: PartialMessage<UpdateProductStockRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "commurz.v1.UpdateProductStockRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "product_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "version", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "stock_in", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 4, name: "stock_out", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): UpdateProductStockRequest {
    return new UpdateProductStockRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): UpdateProductStockRequest {
    return new UpdateProductStockRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): UpdateProductStockRequest {
    return new UpdateProductStockRequest().fromJsonString(jsonString, options);
  }

  static equals(a: UpdateProductStockRequest | PlainMessage<UpdateProductStockRequest> | undefined, b: UpdateProductStockRequest | PlainMessage<UpdateProductStockRequest> | undefined): boolean {
    return proto3.util.equals(UpdateProductStockRequest, a, b);
  }
}

/**
 * @generated from message commurz.v1.Empty
 */
export class Empty extends Message<Empty> {
  constructor(data?: PartialMessage<Empty>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "commurz.v1.Empty";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Empty {
    return new Empty().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Empty {
    return new Empty().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Empty {
    return new Empty().fromJsonString(jsonString, options);
  }

  static equals(a: Empty | PlainMessage<Empty> | undefined, b: Empty | PlainMessage<Empty> | undefined): boolean {
    return proto3.util.equals(Empty, a, b);
  }
}

/**
 * @generated from message commurz.v1.AddProductToCartRequest
 */
export class AddProductToCartRequest extends Message<AddProductToCartRequest> {
  /**
   * @generated from field: string product_id = 1;
   */
  productId = "";

  /**
   * @generated from field: int64 quantity = 2;
   */
  quantity = protoInt64.zero;

  /**
   * TODO: remove this, change to auth token/session
   *
   * @generated from field: string user_id = 3;
   */
  userId = "";

  constructor(data?: PartialMessage<AddProductToCartRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "commurz.v1.AddProductToCartRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "product_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "quantity", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 3, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddProductToCartRequest {
    return new AddProductToCartRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddProductToCartRequest {
    return new AddProductToCartRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddProductToCartRequest {
    return new AddProductToCartRequest().fromJsonString(jsonString, options);
  }

  static equals(a: AddProductToCartRequest | PlainMessage<AddProductToCartRequest> | undefined, b: AddProductToCartRequest | PlainMessage<AddProductToCartRequest> | undefined): boolean {
    return proto3.util.equals(AddProductToCartRequest, a, b);
  }
}

/**
 * @generated from message commurz.v1.CheckoutAllRequest
 */
export class CheckoutAllRequest extends Message<CheckoutAllRequest> {
  /**
   * TODO: remove this, change to auth token/session
   *
   * @generated from field: string user_id = 1;
   */
  userId = "";

  constructor(data?: PartialMessage<CheckoutAllRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "commurz.v1.CheckoutAllRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "user_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): CheckoutAllRequest {
    return new CheckoutAllRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): CheckoutAllRequest {
    return new CheckoutAllRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): CheckoutAllRequest {
    return new CheckoutAllRequest().fromJsonString(jsonString, options);
  }

  static equals(a: CheckoutAllRequest | PlainMessage<CheckoutAllRequest> | undefined, b: CheckoutAllRequest | PlainMessage<CheckoutAllRequest> | undefined): boolean {
    return proto3.util.equals(CheckoutAllRequest, a, b);
  }
}

/**
 * @generated from message commurz.v1.FindByIDRequest
 */
export class FindByIDRequest extends Message<FindByIDRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  constructor(data?: PartialMessage<FindByIDRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "commurz.v1.FindByIDRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FindByIDRequest {
    return new FindByIDRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FindByIDRequest {
    return new FindByIDRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FindByIDRequest {
    return new FindByIDRequest().fromJsonString(jsonString, options);
  }

  static equals(a: FindByIDRequest | PlainMessage<FindByIDRequest> | undefined, b: FindByIDRequest | PlainMessage<FindByIDRequest> | undefined): boolean {
    return proto3.util.equals(FindByIDRequest, a, b);
  }
}

/**
 * @generated from message commurz.v1.ListUsersRequest
 */
export class ListUsersRequest extends Message<ListUsersRequest> {
  constructor(data?: PartialMessage<ListUsersRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "commurz.v1.ListUsersRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListUsersRequest {
    return new ListUsersRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListUsersRequest {
    return new ListUsersRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListUsersRequest {
    return new ListUsersRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListUsersRequest | PlainMessage<ListUsersRequest> | undefined, b: ListUsersRequest | PlainMessage<ListUsersRequest> | undefined): boolean {
    return proto3.util.equals(ListUsersRequest, a, b);
  }
}

/**
 * @generated from message commurz.v1.ListUsersResponse
 */
export class ListUsersResponse extends Message<ListUsersResponse> {
  /**
   * @generated from field: repeated commurz.v1.User users = 1;
   */
  users: User[] = [];

  constructor(data?: PartialMessage<ListUsersResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "commurz.v1.ListUsersResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "users", kind: "message", T: User, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListUsersResponse {
    return new ListUsersResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListUsersResponse {
    return new ListUsersResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListUsersResponse {
    return new ListUsersResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListUsersResponse | PlainMessage<ListUsersResponse> | undefined, b: ListUsersResponse | PlainMessage<ListUsersResponse> | undefined): boolean {
    return proto3.util.equals(ListUsersResponse, a, b);
  }
}

/**
 * @generated from message commurz.v1.Pagination
 */
export class Pagination extends Message<Pagination> {
  /**
   * @generated from field: int32 page = 1;
   */
  page = 0;

  /**
   * @generated from field: int32 size = 2;
   */
  size = 0;

  constructor(data?: PartialMessage<Pagination>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "commurz.v1.Pagination";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "page", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
    { no: 2, name: "size", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Pagination {
    return new Pagination().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Pagination {
    return new Pagination().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Pagination {
    return new Pagination().fromJsonString(jsonString, options);
  }

  static equals(a: Pagination | PlainMessage<Pagination> | undefined, b: Pagination | PlainMessage<Pagination> | undefined): boolean {
    return proto3.util.equals(Pagination, a, b);
  }
}

/**
 * @generated from message commurz.v1.ListBackofficeProductsRequest
 */
export class ListBackofficeProductsRequest extends Message<ListBackofficeProductsRequest> {
  /**
   * @generated from field: commurz.v1.Pagination pagination = 1;
   */
  pagination?: Pagination;

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  constructor(data?: PartialMessage<ListBackofficeProductsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "commurz.v1.ListBackofficeProductsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "pagination", kind: "message", T: Pagination },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListBackofficeProductsRequest {
    return new ListBackofficeProductsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListBackofficeProductsRequest {
    return new ListBackofficeProductsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListBackofficeProductsRequest {
    return new ListBackofficeProductsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListBackofficeProductsRequest | PlainMessage<ListBackofficeProductsRequest> | undefined, b: ListBackofficeProductsRequest | PlainMessage<ListBackofficeProductsRequest> | undefined): boolean {
    return proto3.util.equals(ListBackofficeProductsRequest, a, b);
  }
}

/**
 * @generated from message commurz.v1.ListBackofficeProductsResponse
 */
export class ListBackofficeProductsResponse extends Message<ListBackofficeProductsResponse> {
  /**
   * @generated from field: repeated commurz.v1.Product products = 1;
   */
  products: Product[] = [];

  /**
   * @generated from field: int32 count = 2;
   */
  count = 0;

  constructor(data?: PartialMessage<ListBackofficeProductsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "commurz.v1.ListBackofficeProductsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "products", kind: "message", T: Product, repeated: true },
    { no: 2, name: "count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListBackofficeProductsResponse {
    return new ListBackofficeProductsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListBackofficeProductsResponse {
    return new ListBackofficeProductsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListBackofficeProductsResponse {
    return new ListBackofficeProductsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListBackofficeProductsResponse | PlainMessage<ListBackofficeProductsResponse> | undefined, b: ListBackofficeProductsResponse | PlainMessage<ListBackofficeProductsResponse> | undefined): boolean {
    return proto3.util.equals(ListBackofficeProductsResponse, a, b);
  }
}

/**
 * @generated from message commurz.v1.ListAppProductsRequest
 */
export class ListAppProductsRequest extends Message<ListAppProductsRequest> {
  /**
   * @generated from field: commurz.v1.Pagination pagination = 1;
   */
  pagination?: Pagination;

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  constructor(data?: PartialMessage<ListAppProductsRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "commurz.v1.ListAppProductsRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "pagination", kind: "message", T: Pagination },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListAppProductsRequest {
    return new ListAppProductsRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListAppProductsRequest {
    return new ListAppProductsRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListAppProductsRequest {
    return new ListAppProductsRequest().fromJsonString(jsonString, options);
  }

  static equals(a: ListAppProductsRequest | PlainMessage<ListAppProductsRequest> | undefined, b: ListAppProductsRequest | PlainMessage<ListAppProductsRequest> | undefined): boolean {
    return proto3.util.equals(ListAppProductsRequest, a, b);
  }
}

/**
 * @generated from message commurz.v1.ProductListing
 */
export class ProductListing extends Message<ProductListing> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: int64 price = 3;
   */
  price = protoInt64.zero;

  /**
   * @generated from field: int64 latest_stock = 4;
   */
  latestStock = protoInt64.zero;

  /**
   * price in text format, e.g. "Rp 10.000"
   *
   * @generated from field: string text_price_idr = 5;
   */
  textPriceIdr = "";

  constructor(data?: PartialMessage<ProductListing>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "commurz.v1.ProductListing";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "price", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 4, name: "latest_stock", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
    { no: 5, name: "text_price_idr", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ProductListing {
    return new ProductListing().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ProductListing {
    return new ProductListing().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ProductListing {
    return new ProductListing().fromJsonString(jsonString, options);
  }

  static equals(a: ProductListing | PlainMessage<ProductListing> | undefined, b: ProductListing | PlainMessage<ProductListing> | undefined): boolean {
    return proto3.util.equals(ProductListing, a, b);
  }
}

/**
 * @generated from message commurz.v1.ListAppProductsResponse
 */
export class ListAppProductsResponse extends Message<ListAppProductsResponse> {
  /**
   * @generated from field: repeated commurz.v1.ProductListing products = 1;
   */
  products: ProductListing[] = [];

  /**
   * @generated from field: int32 count = 2;
   */
  count = 0;

  constructor(data?: PartialMessage<ListAppProductsResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "commurz.v1.ListAppProductsResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "products", kind: "message", T: ProductListing, repeated: true },
    { no: 2, name: "count", kind: "scalar", T: 5 /* ScalarType.INT32 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): ListAppProductsResponse {
    return new ListAppProductsResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): ListAppProductsResponse {
    return new ListAppProductsResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): ListAppProductsResponse {
    return new ListAppProductsResponse().fromJsonString(jsonString, options);
  }

  static equals(a: ListAppProductsResponse | PlainMessage<ListAppProductsResponse> | undefined, b: ListAppProductsResponse | PlainMessage<ListAppProductsResponse> | undefined): boolean {
    return proto3.util.equals(ListAppProductsResponse, a, b);
  }
}
