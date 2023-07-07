// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.31.0
// 	protoc        v3.21.12
// source: proto/pool/price.proto

package pool

import (
	_ "github.com/cosmos/gogoproto/gogoproto"
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

type EstimateSwapExactAmountInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PoolId  uint64               `protobuf:"varint,2,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	TokenIn string               `protobuf:"bytes,3,opt,name=token_in,json=tokenIn,proto3" json:"token_in,omitempty"`
	Routes  []*SwapAmountInRoute `protobuf:"bytes,4,rep,name=routes,proto3" json:"routes,omitempty"`
}

func (x *EstimateSwapExactAmountInRequest) Reset() {
	*x = EstimateSwapExactAmountInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pool_price_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EstimateSwapExactAmountInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstimateSwapExactAmountInRequest) ProtoMessage() {}

func (x *EstimateSwapExactAmountInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pool_price_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EstimateSwapExactAmountInRequest.ProtoReflect.Descriptor instead.
func (*EstimateSwapExactAmountInRequest) Descriptor() ([]byte, []int) {
	return file_proto_pool_price_proto_rawDescGZIP(), []int{0}
}

func (x *EstimateSwapExactAmountInRequest) GetPoolId() uint64 {
	if x != nil {
		return x.PoolId
	}
	return 0
}

func (x *EstimateSwapExactAmountInRequest) GetTokenIn() string {
	if x != nil {
		return x.TokenIn
	}
	return ""
}

func (x *EstimateSwapExactAmountInRequest) GetRoutes() []*SwapAmountInRoute {
	if x != nil {
		return x.Routes
	}
	return nil
}

type EstimateSinglePoolSwapExactAmountInRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PoolId        uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	TokenIn       string `protobuf:"bytes,2,opt,name=token_in,json=tokenIn,proto3" json:"token_in,omitempty"`
	TokenOutDenom string `protobuf:"bytes,3,opt,name=token_out_denom,json=tokenOutDenom,proto3" json:"token_out_denom,omitempty"`
}

func (x *EstimateSinglePoolSwapExactAmountInRequest) Reset() {
	*x = EstimateSinglePoolSwapExactAmountInRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pool_price_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EstimateSinglePoolSwapExactAmountInRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstimateSinglePoolSwapExactAmountInRequest) ProtoMessage() {}

func (x *EstimateSinglePoolSwapExactAmountInRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pool_price_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EstimateSinglePoolSwapExactAmountInRequest.ProtoReflect.Descriptor instead.
func (*EstimateSinglePoolSwapExactAmountInRequest) Descriptor() ([]byte, []int) {
	return file_proto_pool_price_proto_rawDescGZIP(), []int{1}
}

func (x *EstimateSinglePoolSwapExactAmountInRequest) GetPoolId() uint64 {
	if x != nil {
		return x.PoolId
	}
	return 0
}

func (x *EstimateSinglePoolSwapExactAmountInRequest) GetTokenIn() string {
	if x != nil {
		return x.TokenIn
	}
	return ""
}

func (x *EstimateSinglePoolSwapExactAmountInRequest) GetTokenOutDenom() string {
	if x != nil {
		return x.TokenOutDenom
	}
	return ""
}

type EstimateSwapExactAmountInResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TokenOutAmount string `protobuf:"bytes,1,opt,name=token_out_amount,json=tokenOutAmount,proto3" json:"token_out_amount,omitempty"`
}

func (x *EstimateSwapExactAmountInResponse) Reset() {
	*x = EstimateSwapExactAmountInResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pool_price_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *EstimateSwapExactAmountInResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*EstimateSwapExactAmountInResponse) ProtoMessage() {}

func (x *EstimateSwapExactAmountInResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pool_price_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use EstimateSwapExactAmountInResponse.ProtoReflect.Descriptor instead.
func (*EstimateSwapExactAmountInResponse) Descriptor() ([]byte, []int) {
	return file_proto_pool_price_proto_rawDescGZIP(), []int{2}
}

func (x *EstimateSwapExactAmountInResponse) GetTokenOutAmount() string {
	if x != nil {
		return x.TokenOutAmount
	}
	return ""
}

// SpotPriceRequest defines the gRPC request structure for a SpotPrice
// query.
type SpotPriceRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PoolId          uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	BaseAssetDenom  string `protobuf:"bytes,2,opt,name=base_asset_denom,json=baseAssetDenom,proto3" json:"base_asset_denom,omitempty"`
	QuoteAssetDenom string `protobuf:"bytes,3,opt,name=quote_asset_denom,json=quoteAssetDenom,proto3" json:"quote_asset_denom,omitempty"`
}

func (x *SpotPriceRequest) Reset() {
	*x = SpotPriceRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pool_price_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpotPriceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpotPriceRequest) ProtoMessage() {}

func (x *SpotPriceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pool_price_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpotPriceRequest.ProtoReflect.Descriptor instead.
func (*SpotPriceRequest) Descriptor() ([]byte, []int) {
	return file_proto_pool_price_proto_rawDescGZIP(), []int{3}
}

func (x *SpotPriceRequest) GetPoolId() uint64 {
	if x != nil {
		return x.PoolId
	}
	return 0
}

func (x *SpotPriceRequest) GetBaseAssetDenom() string {
	if x != nil {
		return x.BaseAssetDenom
	}
	return ""
}

func (x *SpotPriceRequest) GetQuoteAssetDenom() string {
	if x != nil {
		return x.QuoteAssetDenom
	}
	return ""
}

// SpotPriceResponse defines the gRPC response structure for a SpotPrice
// query.
type SpotPriceResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// String of the Dec. Ex) 10.203uatom
	SpotPrice string `protobuf:"bytes,1,opt,name=spot_price,json=spotPrice,proto3" json:"spot_price,omitempty"`
}

func (x *SpotPriceResponse) Reset() {
	*x = SpotPriceResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pool_price_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SpotPriceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SpotPriceResponse) ProtoMessage() {}

func (x *SpotPriceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pool_price_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SpotPriceResponse.ProtoReflect.Descriptor instead.
func (*SpotPriceResponse) Descriptor() ([]byte, []int) {
	return file_proto_pool_price_proto_rawDescGZIP(), []int{4}
}

func (x *SpotPriceResponse) GetSpotPrice() string {
	if x != nil {
		return x.SpotPrice
	}
	return ""
}

type SwapAmountInRoute struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PoolId        uint64 `protobuf:"varint,1,opt,name=pool_id,json=poolId,proto3" json:"pool_id,omitempty"`
	TokenOutDenom string `protobuf:"bytes,2,opt,name=token_out_denom,json=tokenOutDenom,proto3" json:"token_out_denom,omitempty"`
}

func (x *SwapAmountInRoute) Reset() {
	*x = SwapAmountInRoute{}
	if protoimpl.UnsafeEnabled {
		mi := &file_proto_pool_price_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SwapAmountInRoute) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SwapAmountInRoute) ProtoMessage() {}

func (x *SwapAmountInRoute) ProtoReflect() protoreflect.Message {
	mi := &file_proto_pool_price_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SwapAmountInRoute.ProtoReflect.Descriptor instead.
func (*SwapAmountInRoute) Descriptor() ([]byte, []int) {
	return file_proto_pool_price_proto_rawDescGZIP(), []int{5}
}

func (x *SwapAmountInRoute) GetPoolId() uint64 {
	if x != nil {
		return x.PoolId
	}
	return 0
}

func (x *SwapAmountInRoute) GetTokenOutDenom() string {
	if x != nil {
		return x.TokenOutDenom
	}
	return ""
}

var File_proto_pool_price_proto protoreflect.FileDescriptor

var file_proto_pool_price_proto_rawDesc = []byte{
	0x0a, 0x16, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x2f, 0x70, 0x72, 0x69,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x1b, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x69,
	0x73, 0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31,
	0x62, 0x65, 0x74, 0x61, 0x31, 0x1a, 0x14, 0x67, 0x6f, 0x67, 0x6f, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x2f, 0x67, 0x6f, 0x67, 0x6f, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xec, 0x01, 0x0a, 0x20,
	0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x53, 0x77, 0x61, 0x70, 0x45, 0x78, 0x61, 0x63,
	0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x12, 0x2b, 0x0a, 0x07, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x04, 0x42, 0x12, 0xf2, 0xde, 0x1f, 0x0e, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x70, 0x6f, 0x6f,
	0x6c, 0x5f, 0x69, 0x64, 0x22, 0x52, 0x06, 0x70, 0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x12, 0x2e, 0x0a,
	0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x13, 0xf2, 0xde, 0x1f, 0x0f, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x69, 0x6e, 0x22, 0x52, 0x07, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x12, 0x5d, 0x0a,
	0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x2e, 0x2e,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x6d, 0x61, 0x6e, 0x61,
	0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x53, 0x77, 0x61, 0x70,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x42, 0x15, 0xc8,
	0xde, 0x1f, 0x00, 0xf2, 0xde, 0x1f, 0x0d, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x72, 0x6f, 0x75,
	0x74, 0x65, 0x73, 0x22, 0x52, 0x06, 0x72, 0x6f, 0x75, 0x74, 0x65, 0x73, 0x4a, 0x04, 0x08, 0x01,
	0x10, 0x02, 0x52, 0x06, 0x73, 0x65, 0x6e, 0x64, 0x65, 0x72, 0x22, 0xcd, 0x01, 0x0a, 0x2a, 0x45,
	0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50, 0x6f, 0x6f,
	0x6c, 0x53, 0x77, 0x61, 0x70, 0x45, 0x78, 0x61, 0x63, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x07, 0x70, 0x6f, 0x6f,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x12, 0xf2, 0xde, 0x1f, 0x0e,
	0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x22, 0x52, 0x06,
	0x70, 0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x12, 0x2e, 0x0a, 0x08, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x69, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42, 0x13, 0xf2, 0xde, 0x1f, 0x0f, 0x79, 0x61,
	0x6d, 0x6c, 0x3a, 0x22, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x69, 0x6e, 0x22, 0x52, 0x07, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x49, 0x6e, 0x12, 0x42, 0x0a, 0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x6f, 0x75, 0x74, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x1a, 0xf2, 0xde, 0x1f, 0x16, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x22, 0x52, 0x0d, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x4f, 0x75, 0x74, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x22, 0x98, 0x01, 0x0a, 0x21, 0x45,
	0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x53, 0x77, 0x61, 0x70, 0x45, 0x78, 0x61, 0x63, 0x74,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65,
	0x12, 0x73, 0x0a, 0x10, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x42, 0x49, 0xc8, 0xde, 0x1f, 0x00,
	0xda, 0xde, 0x1f, 0x26, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x63,
	0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2f, 0x63, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x2d, 0x73, 0x64, 0x6b,
	0x2f, 0x74, 0x79, 0x70, 0x65, 0x73, 0x2e, 0x49, 0x6e, 0x74, 0xf2, 0xde, 0x1f, 0x17, 0x79, 0x61,
	0x6d, 0x6c, 0x3a, 0x22, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x61, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x22, 0x52, 0x0e, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x4f, 0x75, 0x74, 0x41,
	0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0xd0, 0x01, 0x0a, 0x10, 0x53, 0x70, 0x6f, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x2b, 0x0a, 0x07, 0x70, 0x6f,
	0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x12, 0xf2, 0xde, 0x1f,
	0x0e, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x22, 0x52,
	0x06, 0x70, 0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x12, 0x45, 0x0a, 0x10, 0x62, 0x61, 0x73, 0x65, 0x5f,
	0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x1b, 0xf2, 0xde, 0x1f, 0x17, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x22, 0x52, 0x0e,
	0x62, 0x61, 0x73, 0x65, 0x41, 0x73, 0x73, 0x65, 0x74, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x12, 0x48,
	0x0a, 0x11, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74, 0x5f, 0x64, 0x65,
	0x6e, 0x6f, 0x6d, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x42, 0x1c, 0xf2, 0xde, 0x1f, 0x18, 0x79,
	0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x5f, 0x61, 0x73, 0x73, 0x65, 0x74,
	0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x22, 0x52, 0x0f, 0x71, 0x75, 0x6f, 0x74, 0x65, 0x41, 0x73,
	0x73, 0x65, 0x74, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x22, 0x49, 0x0a, 0x11, 0x53, 0x70, 0x6f, 0x74,
	0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x34, 0x0a,
	0x0a, 0x73, 0x70, 0x6f, 0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x09, 0x42, 0x15, 0xf2, 0xde, 0x1f, 0x11, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x73, 0x70, 0x6f,
	0x74, 0x5f, 0x70, 0x72, 0x69, 0x63, 0x65, 0x22, 0x52, 0x09, 0x73, 0x70, 0x6f, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x22, 0x84, 0x01, 0x0a, 0x11, 0x53, 0x77, 0x61, 0x70, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x49, 0x6e, 0x52, 0x6f, 0x75, 0x74, 0x65, 0x12, 0x2b, 0x0a, 0x07, 0x70, 0x6f, 0x6f,
	0x6c, 0x5f, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x04, 0x42, 0x12, 0xf2, 0xde, 0x1f, 0x0e,
	0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x70, 0x6f, 0x6f, 0x6c, 0x5f, 0x69, 0x64, 0x22, 0x52, 0x06,
	0x70, 0x6f, 0x6f, 0x6c, 0x49, 0x64, 0x12, 0x42, 0x0a, 0x0f, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x5f,
	0x6f, 0x75, 0x74, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x42,
	0x1a, 0xf2, 0xde, 0x1f, 0x16, 0x79, 0x61, 0x6d, 0x6c, 0x3a, 0x22, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x5f, 0x6f, 0x75, 0x74, 0x5f, 0x64, 0x65, 0x6e, 0x6f, 0x6d, 0x22, 0x52, 0x0d, 0x74, 0x6f, 0x6b,
	0x65, 0x6e, 0x4f, 0x75, 0x74, 0x44, 0x65, 0x6e, 0x6f, 0x6d, 0x32, 0xc7, 0x03, 0x0a, 0x05, 0x51,
	0x75, 0x65, 0x72, 0x79, 0x12, 0x9c, 0x01, 0x0a, 0x19, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74,
	0x65, 0x53, 0x77, 0x61, 0x70, 0x45, 0x78, 0x61, 0x63, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74,
	0x49, 0x6e, 0x12, 0x3d, 0x2e, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x6f, 0x6f,
	0x6c, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x53, 0x77, 0x61, 0x70, 0x45, 0x78, 0x61,
	0x63, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x1a, 0x3e, 0x2e, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x6f, 0x6f, 0x6c,
	0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e,
	0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x53, 0x77, 0x61, 0x70, 0x45, 0x78, 0x61, 0x63,
	0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73,
	0x65, 0x22, 0x00, 0x12, 0xb0, 0x01, 0x0a, 0x23, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65,
	0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x77, 0x61, 0x70, 0x45, 0x78,
	0x61, 0x63, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x12, 0x47, 0x2e, 0x6f, 0x73,
	0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x6f, 0x6f, 0x6c, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65,
	0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31, 0x2e, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61,
	0x74, 0x65, 0x53, 0x69, 0x6e, 0x67, 0x6c, 0x65, 0x50, 0x6f, 0x6f, 0x6c, 0x53, 0x77, 0x61, 0x70,
	0x45, 0x78, 0x61, 0x63, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x52, 0x65, 0x71,
	0x75, 0x65, 0x73, 0x74, 0x1a, 0x3e, 0x2e, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2e, 0x70,
	0x6f, 0x6f, 0x6c, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74,
	0x61, 0x31, 0x2e, 0x45, 0x73, 0x74, 0x69, 0x6d, 0x61, 0x74, 0x65, 0x53, 0x77, 0x61, 0x70, 0x45,
	0x78, 0x61, 0x63, 0x74, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x49, 0x6e, 0x52, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x6c, 0x0a, 0x09, 0x53, 0x70, 0x6f, 0x74, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x2d, 0x2e, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x6f,
	0x6f, 0x6c, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61,
	0x31, 0x2e, 0x53, 0x70, 0x6f, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x2e, 0x2e, 0x6f, 0x73, 0x6d, 0x6f, 0x73, 0x69, 0x73, 0x2e, 0x70, 0x6f, 0x6f,
	0x6c, 0x6d, 0x61, 0x6e, 0x61, 0x67, 0x65, 0x72, 0x2e, 0x76, 0x31, 0x62, 0x65, 0x74, 0x61, 0x31,
	0x2e, 0x53, 0x70, 0x6f, 0x74, 0x50, 0x72, 0x69, 0x63, 0x65, 0x52, 0x65, 0x73, 0x70, 0x6f, 0x6e,
	0x73, 0x65, 0x22, 0x00, 0x42, 0x36, 0x5a, 0x34, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x64, 0x6f, 0x67, 0x67, 0x79, 0x73, 0x74, 0x79, 0x6c, 0x65, 0x7a, 0x2f, 0x69,
	0x6e, 0x74, 0x65, 0x72, 0x73, 0x74, 0x65, 0x6c, 0x6c, 0x61, 0x72, 0x2d, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x2f, 0x71, 0x75, 0x65, 0x72, 0x79, 0x2f, 0x70, 0x6f, 0x6f, 0x6c, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_proto_pool_price_proto_rawDescOnce sync.Once
	file_proto_pool_price_proto_rawDescData = file_proto_pool_price_proto_rawDesc
)

func file_proto_pool_price_proto_rawDescGZIP() []byte {
	file_proto_pool_price_proto_rawDescOnce.Do(func() {
		file_proto_pool_price_proto_rawDescData = protoimpl.X.CompressGZIP(file_proto_pool_price_proto_rawDescData)
	})
	return file_proto_pool_price_proto_rawDescData
}

var file_proto_pool_price_proto_msgTypes = make([]protoimpl.MessageInfo, 6)
var file_proto_pool_price_proto_goTypes = []interface{}{
	(*EstimateSwapExactAmountInRequest)(nil),           // 0: osmosis.poolmanager.v1beta1.EstimateSwapExactAmountInRequest
	(*EstimateSinglePoolSwapExactAmountInRequest)(nil), // 1: osmosis.poolmanager.v1beta1.EstimateSinglePoolSwapExactAmountInRequest
	(*EstimateSwapExactAmountInResponse)(nil),          // 2: osmosis.poolmanager.v1beta1.EstimateSwapExactAmountInResponse
	(*SpotPriceRequest)(nil),                           // 3: osmosis.poolmanager.v1beta1.SpotPriceRequest
	(*SpotPriceResponse)(nil),                          // 4: osmosis.poolmanager.v1beta1.SpotPriceResponse
	(*SwapAmountInRoute)(nil),                          // 5: osmosis.poolmanager.v1beta1.SwapAmountInRoute
}
var file_proto_pool_price_proto_depIdxs = []int32{
	5, // 0: osmosis.poolmanager.v1beta1.EstimateSwapExactAmountInRequest.routes:type_name -> osmosis.poolmanager.v1beta1.SwapAmountInRoute
	0, // 1: osmosis.poolmanager.v1beta1.Query.EstimateSwapExactAmountIn:input_type -> osmosis.poolmanager.v1beta1.EstimateSwapExactAmountInRequest
	1, // 2: osmosis.poolmanager.v1beta1.Query.EstimateSinglePoolSwapExactAmountIn:input_type -> osmosis.poolmanager.v1beta1.EstimateSinglePoolSwapExactAmountInRequest
	3, // 3: osmosis.poolmanager.v1beta1.Query.SpotPrice:input_type -> osmosis.poolmanager.v1beta1.SpotPriceRequest
	2, // 4: osmosis.poolmanager.v1beta1.Query.EstimateSwapExactAmountIn:output_type -> osmosis.poolmanager.v1beta1.EstimateSwapExactAmountInResponse
	2, // 5: osmosis.poolmanager.v1beta1.Query.EstimateSinglePoolSwapExactAmountIn:output_type -> osmosis.poolmanager.v1beta1.EstimateSwapExactAmountInResponse
	4, // 6: osmosis.poolmanager.v1beta1.Query.SpotPrice:output_type -> osmosis.poolmanager.v1beta1.SpotPriceResponse
	4, // [4:7] is the sub-list for method output_type
	1, // [1:4] is the sub-list for method input_type
	1, // [1:1] is the sub-list for extension type_name
	1, // [1:1] is the sub-list for extension extendee
	0, // [0:1] is the sub-list for field type_name
}

func init() { file_proto_pool_price_proto_init() }
func file_proto_pool_price_proto_init() {
	if File_proto_pool_price_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_proto_pool_price_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EstimateSwapExactAmountInRequest); i {
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
		file_proto_pool_price_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EstimateSinglePoolSwapExactAmountInRequest); i {
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
		file_proto_pool_price_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*EstimateSwapExactAmountInResponse); i {
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
		file_proto_pool_price_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpotPriceRequest); i {
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
		file_proto_pool_price_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SpotPriceResponse); i {
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
		file_proto_pool_price_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SwapAmountInRoute); i {
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
			RawDescriptor: file_proto_pool_price_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   6,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_proto_pool_price_proto_goTypes,
		DependencyIndexes: file_proto_pool_price_proto_depIdxs,
		MessageInfos:      file_proto_pool_price_proto_msgTypes,
	}.Build()
	File_proto_pool_price_proto = out.File
	file_proto_pool_price_proto_rawDesc = nil
	file_proto_pool_price_proto_goTypes = nil
	file_proto_pool_price_proto_depIdxs = nil
}