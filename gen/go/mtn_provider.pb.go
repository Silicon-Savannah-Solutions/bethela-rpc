// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.36.6
// 	protoc        v5.29.3
// source: mtn_provider.proto

package bethela_rpc

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
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

// Request to Pay - Customer to Business payment
type RequestToPayRequest struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	ExternalId        string                 `protobuf:"bytes,1,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`                      // Reference ID provided by the client
	TargetEnvironment string                 `protobuf:"bytes,2,opt,name=target_environment,json=targetEnvironment,proto3" json:"target_environment,omitempty"` // sandbox or production
	Amount            float64                `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency          string                 `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`                            // e.g., "EUR"
	PayerId           string                 `protobuf:"bytes,5,opt,name=payer_id,json=payerId,proto3" json:"payer_id,omitempty"`               // Phone number with country code
	PayerIdType       string                 `protobuf:"bytes,6,opt,name=payer_id_type,json=payerIdType,proto3" json:"payer_id_type,omitempty"` // MSISDN, EMAIL, or PARTY_CODE
	PayerMessage      string                 `protobuf:"bytes,7,opt,name=payer_message,json=payerMessage,proto3" json:"payer_message,omitempty"`
	PayeeNote         string                 `protobuf:"bytes,8,opt,name=payee_note,json=payeeNote,proto3" json:"payee_note,omitempty"`
	CallbackUrl       string                 `protobuf:"bytes,9,opt,name=callback_url,json=callbackUrl,proto3" json:"callback_url,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *RequestToPayRequest) Reset() {
	*x = RequestToPayRequest{}
	mi := &file_mtn_provider_proto_msgTypes[0]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestToPayRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestToPayRequest) ProtoMessage() {}

func (x *RequestToPayRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mtn_provider_proto_msgTypes[0]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestToPayRequest.ProtoReflect.Descriptor instead.
func (*RequestToPayRequest) Descriptor() ([]byte, []int) {
	return file_mtn_provider_proto_rawDescGZIP(), []int{0}
}

func (x *RequestToPayRequest) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *RequestToPayRequest) GetTargetEnvironment() string {
	if x != nil {
		return x.TargetEnvironment
	}
	return ""
}

func (x *RequestToPayRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *RequestToPayRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *RequestToPayRequest) GetPayerId() string {
	if x != nil {
		return x.PayerId
	}
	return ""
}

func (x *RequestToPayRequest) GetPayerIdType() string {
	if x != nil {
		return x.PayerIdType
	}
	return ""
}

func (x *RequestToPayRequest) GetPayerMessage() string {
	if x != nil {
		return x.PayerMessage
	}
	return ""
}

func (x *RequestToPayRequest) GetPayeeNote() string {
	if x != nil {
		return x.PayeeNote
	}
	return ""
}

func (x *RequestToPayRequest) GetCallbackUrl() string {
	if x != nil {
		return x.CallbackUrl
	}
	return ""
}

type RequestToPayResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReferenceId   string                 `protobuf:"bytes,1,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"` // Transaction reference ID from MTN
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`                              // PENDING, SUCCESSFUL, FAILED
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *RequestToPayResponse) Reset() {
	*x = RequestToPayResponse{}
	mi := &file_mtn_provider_proto_msgTypes[1]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *RequestToPayResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*RequestToPayResponse) ProtoMessage() {}

func (x *RequestToPayResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mtn_provider_proto_msgTypes[1]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use RequestToPayResponse.ProtoReflect.Descriptor instead.
func (*RequestToPayResponse) Descriptor() ([]byte, []int) {
	return file_mtn_provider_proto_rawDescGZIP(), []int{1}
}

func (x *RequestToPayResponse) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (x *RequestToPayResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *RequestToPayResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

// Transfer - Business to Customer payment
type TransferRequest struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	ExternalId        string                 `protobuf:"bytes,1,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`                      // Reference ID provided by the client
	TargetEnvironment string                 `protobuf:"bytes,2,opt,name=target_environment,json=targetEnvironment,proto3" json:"target_environment,omitempty"` // sandbox or production
	Amount            float64                `protobuf:"fixed64,3,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency          string                 `protobuf:"bytes,4,opt,name=currency,proto3" json:"currency,omitempty"`
	PayeeId           string                 `protobuf:"bytes,5,opt,name=payee_id,json=payeeId,proto3" json:"payee_id,omitempty"`               // Phone number with country code
	PayeeIdType       string                 `protobuf:"bytes,6,opt,name=payee_id_type,json=payeeIdType,proto3" json:"payee_id_type,omitempty"` // MSISDN, EMAIL, or PARTY_CODE
	PayerMessage      string                 `protobuf:"bytes,7,opt,name=payer_message,json=payerMessage,proto3" json:"payer_message,omitempty"`
	PayeeNote         string                 `protobuf:"bytes,8,opt,name=payee_note,json=payeeNote,proto3" json:"payee_note,omitempty"`
	CallbackUrl       string                 `protobuf:"bytes,9,opt,name=callback_url,json=callbackUrl,proto3" json:"callback_url,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *TransferRequest) Reset() {
	*x = TransferRequest{}
	mi := &file_mtn_provider_proto_msgTypes[2]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransferRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferRequest) ProtoMessage() {}

func (x *TransferRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mtn_provider_proto_msgTypes[2]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferRequest.ProtoReflect.Descriptor instead.
func (*TransferRequest) Descriptor() ([]byte, []int) {
	return file_mtn_provider_proto_rawDescGZIP(), []int{2}
}

func (x *TransferRequest) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *TransferRequest) GetTargetEnvironment() string {
	if x != nil {
		return x.TargetEnvironment
	}
	return ""
}

func (x *TransferRequest) GetAmount() float64 {
	if x != nil {
		return x.Amount
	}
	return 0
}

func (x *TransferRequest) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *TransferRequest) GetPayeeId() string {
	if x != nil {
		return x.PayeeId
	}
	return ""
}

func (x *TransferRequest) GetPayeeIdType() string {
	if x != nil {
		return x.PayeeIdType
	}
	return ""
}

func (x *TransferRequest) GetPayerMessage() string {
	if x != nil {
		return x.PayerMessage
	}
	return ""
}

func (x *TransferRequest) GetPayeeNote() string {
	if x != nil {
		return x.PayeeNote
	}
	return ""
}

func (x *TransferRequest) GetCallbackUrl() string {
	if x != nil {
		return x.CallbackUrl
	}
	return ""
}

type TransferResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ReferenceId   string                 `protobuf:"bytes,1,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"` // Transaction reference ID from MTN
	Status        string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`                              // PENDING, SUCCESSFUL, FAILED
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,3,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *TransferResponse) Reset() {
	*x = TransferResponse{}
	mi := &file_mtn_provider_proto_msgTypes[3]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *TransferResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransferResponse) ProtoMessage() {}

func (x *TransferResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mtn_provider_proto_msgTypes[3]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransferResponse.ProtoReflect.Descriptor instead.
func (*TransferResponse) Descriptor() ([]byte, []int) {
	return file_mtn_provider_proto_rawDescGZIP(), []int{3}
}

func (x *TransferResponse) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (x *TransferResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *TransferResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

// Get Request to Pay Status
type GetRequestToPayStatusRequest struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	ReferenceId       string                 `protobuf:"bytes,1,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
	TargetEnvironment string                 `protobuf:"bytes,2,opt,name=target_environment,json=targetEnvironment,proto3" json:"target_environment,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *GetRequestToPayStatusRequest) Reset() {
	*x = GetRequestToPayStatusRequest{}
	mi := &file_mtn_provider_proto_msgTypes[4]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequestToPayStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequestToPayStatusRequest) ProtoMessage() {}

func (x *GetRequestToPayStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mtn_provider_proto_msgTypes[4]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequestToPayStatusRequest.ProtoReflect.Descriptor instead.
func (*GetRequestToPayStatusRequest) Descriptor() ([]byte, []int) {
	return file_mtn_provider_proto_rawDescGZIP(), []int{4}
}

func (x *GetRequestToPayStatusRequest) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (x *GetRequestToPayStatusRequest) GetTargetEnvironment() string {
	if x != nil {
		return x.TargetEnvironment
	}
	return ""
}

type GetRequestToPayStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ExternalId    string                 `protobuf:"bytes,1,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	Amount        string                 `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency      string                 `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	PayerId       string                 `protobuf:"bytes,4,opt,name=payer_id,json=payerId,proto3" json:"payer_id,omitempty"`
	PayerIdType   string                 `protobuf:"bytes,5,opt,name=payer_id_type,json=payerIdType,proto3" json:"payer_id_type,omitempty"`
	Status        string                 `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Reason        string                 `protobuf:"bytes,7,opt,name=reason,proto3" json:"reason,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ModifiedAt    *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=modified_at,json=modifiedAt,proto3" json:"modified_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetRequestToPayStatusResponse) Reset() {
	*x = GetRequestToPayStatusResponse{}
	mi := &file_mtn_provider_proto_msgTypes[5]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetRequestToPayStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetRequestToPayStatusResponse) ProtoMessage() {}

func (x *GetRequestToPayStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mtn_provider_proto_msgTypes[5]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetRequestToPayStatusResponse.ProtoReflect.Descriptor instead.
func (*GetRequestToPayStatusResponse) Descriptor() ([]byte, []int) {
	return file_mtn_provider_proto_rawDescGZIP(), []int{5}
}

func (x *GetRequestToPayStatusResponse) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *GetRequestToPayStatusResponse) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *GetRequestToPayStatusResponse) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *GetRequestToPayStatusResponse) GetPayerId() string {
	if x != nil {
		return x.PayerId
	}
	return ""
}

func (x *GetRequestToPayStatusResponse) GetPayerIdType() string {
	if x != nil {
		return x.PayerIdType
	}
	return ""
}

func (x *GetRequestToPayStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetRequestToPayStatusResponse) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *GetRequestToPayStatusResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *GetRequestToPayStatusResponse) GetModifiedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ModifiedAt
	}
	return nil
}

// Get Transfer Status
type GetTransferStatusRequest struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	ReferenceId       string                 `protobuf:"bytes,1,opt,name=reference_id,json=referenceId,proto3" json:"reference_id,omitempty"`
	TargetEnvironment string                 `protobuf:"bytes,2,opt,name=target_environment,json=targetEnvironment,proto3" json:"target_environment,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *GetTransferStatusRequest) Reset() {
	*x = GetTransferStatusRequest{}
	mi := &file_mtn_provider_proto_msgTypes[6]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTransferStatusRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransferStatusRequest) ProtoMessage() {}

func (x *GetTransferStatusRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mtn_provider_proto_msgTypes[6]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransferStatusRequest.ProtoReflect.Descriptor instead.
func (*GetTransferStatusRequest) Descriptor() ([]byte, []int) {
	return file_mtn_provider_proto_rawDescGZIP(), []int{6}
}

func (x *GetTransferStatusRequest) GetReferenceId() string {
	if x != nil {
		return x.ReferenceId
	}
	return ""
}

func (x *GetTransferStatusRequest) GetTargetEnvironment() string {
	if x != nil {
		return x.TargetEnvironment
	}
	return ""
}

type GetTransferStatusResponse struct {
	state         protoimpl.MessageState `protogen:"open.v1"`
	ExternalId    string                 `protobuf:"bytes,1,opt,name=external_id,json=externalId,proto3" json:"external_id,omitempty"`
	Amount        string                 `protobuf:"bytes,2,opt,name=amount,proto3" json:"amount,omitempty"`
	Currency      string                 `protobuf:"bytes,3,opt,name=currency,proto3" json:"currency,omitempty"`
	PayeeId       string                 `protobuf:"bytes,4,opt,name=payee_id,json=payeeId,proto3" json:"payee_id,omitempty"`
	PayeeIdType   string                 `protobuf:"bytes,5,opt,name=payee_id_type,json=payeeIdType,proto3" json:"payee_id_type,omitempty"`
	Status        string                 `protobuf:"bytes,6,opt,name=status,proto3" json:"status,omitempty"`
	Reason        string                 `protobuf:"bytes,7,opt,name=reason,proto3" json:"reason,omitempty"`
	CreatedAt     *timestamppb.Timestamp `protobuf:"bytes,8,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	ModifiedAt    *timestamppb.Timestamp `protobuf:"bytes,9,opt,name=modified_at,json=modifiedAt,proto3" json:"modified_at,omitempty"`
	unknownFields protoimpl.UnknownFields
	sizeCache     protoimpl.SizeCache
}

func (x *GetTransferStatusResponse) Reset() {
	*x = GetTransferStatusResponse{}
	mi := &file_mtn_provider_proto_msgTypes[7]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetTransferStatusResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetTransferStatusResponse) ProtoMessage() {}

func (x *GetTransferStatusResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mtn_provider_proto_msgTypes[7]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetTransferStatusResponse.ProtoReflect.Descriptor instead.
func (*GetTransferStatusResponse) Descriptor() ([]byte, []int) {
	return file_mtn_provider_proto_rawDescGZIP(), []int{7}
}

func (x *GetTransferStatusResponse) GetExternalId() string {
	if x != nil {
		return x.ExternalId
	}
	return ""
}

func (x *GetTransferStatusResponse) GetAmount() string {
	if x != nil {
		return x.Amount
	}
	return ""
}

func (x *GetTransferStatusResponse) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

func (x *GetTransferStatusResponse) GetPayeeId() string {
	if x != nil {
		return x.PayeeId
	}
	return ""
}

func (x *GetTransferStatusResponse) GetPayeeIdType() string {
	if x != nil {
		return x.PayeeIdType
	}
	return ""
}

func (x *GetTransferStatusResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *GetTransferStatusResponse) GetReason() string {
	if x != nil {
		return x.Reason
	}
	return ""
}

func (x *GetTransferStatusResponse) GetCreatedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.CreatedAt
	}
	return nil
}

func (x *GetTransferStatusResponse) GetModifiedAt() *timestamppb.Timestamp {
	if x != nil {
		return x.ModifiedAt
	}
	return nil
}

// Get Account Balance
type GetAccountBalanceRequest struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	TargetEnvironment string                 `protobuf:"bytes,1,opt,name=target_environment,json=targetEnvironment,proto3" json:"target_environment,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *GetAccountBalanceRequest) Reset() {
	*x = GetAccountBalanceRequest{}
	mi := &file_mtn_provider_proto_msgTypes[8]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAccountBalanceRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountBalanceRequest) ProtoMessage() {}

func (x *GetAccountBalanceRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mtn_provider_proto_msgTypes[8]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountBalanceRequest.ProtoReflect.Descriptor instead.
func (*GetAccountBalanceRequest) Descriptor() ([]byte, []int) {
	return file_mtn_provider_proto_rawDescGZIP(), []int{8}
}

func (x *GetAccountBalanceRequest) GetTargetEnvironment() string {
	if x != nil {
		return x.TargetEnvironment
	}
	return ""
}

type GetAccountBalanceResponse struct {
	state            protoimpl.MessageState `protogen:"open.v1"`
	AvailableBalance string                 `protobuf:"bytes,1,opt,name=available_balance,json=availableBalance,proto3" json:"available_balance,omitempty"`
	Currency         string                 `protobuf:"bytes,2,opt,name=currency,proto3" json:"currency,omitempty"`
	unknownFields    protoimpl.UnknownFields
	sizeCache        protoimpl.SizeCache
}

func (x *GetAccountBalanceResponse) Reset() {
	*x = GetAccountBalanceResponse{}
	mi := &file_mtn_provider_proto_msgTypes[9]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *GetAccountBalanceResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*GetAccountBalanceResponse) ProtoMessage() {}

func (x *GetAccountBalanceResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mtn_provider_proto_msgTypes[9]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use GetAccountBalanceResponse.ProtoReflect.Descriptor instead.
func (*GetAccountBalanceResponse) Descriptor() ([]byte, []int) {
	return file_mtn_provider_proto_rawDescGZIP(), []int{9}
}

func (x *GetAccountBalanceResponse) GetAvailableBalance() string {
	if x != nil {
		return x.AvailableBalance
	}
	return ""
}

func (x *GetAccountBalanceResponse) GetCurrency() string {
	if x != nil {
		return x.Currency
	}
	return ""
}

// Validate Account Holder
type ValidateAccountHolderRequest struct {
	state               protoimpl.MessageState `protogen:"open.v1"`
	AccountHolderId     string                 `protobuf:"bytes,1,opt,name=account_holder_id,json=accountHolderId,proto3" json:"account_holder_id,omitempty"`
	AccountHolderIdType string                 `protobuf:"bytes,2,opt,name=account_holder_id_type,json=accountHolderIdType,proto3" json:"account_holder_id_type,omitempty"` // MSISDN, EMAIL, or PARTY_CODE
	TargetEnvironment   string                 `protobuf:"bytes,3,opt,name=target_environment,json=targetEnvironment,proto3" json:"target_environment,omitempty"`
	unknownFields       protoimpl.UnknownFields
	sizeCache           protoimpl.SizeCache
}

func (x *ValidateAccountHolderRequest) Reset() {
	*x = ValidateAccountHolderRequest{}
	mi := &file_mtn_provider_proto_msgTypes[10]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateAccountHolderRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateAccountHolderRequest) ProtoMessage() {}

func (x *ValidateAccountHolderRequest) ProtoReflect() protoreflect.Message {
	mi := &file_mtn_provider_proto_msgTypes[10]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateAccountHolderRequest.ProtoReflect.Descriptor instead.
func (*ValidateAccountHolderRequest) Descriptor() ([]byte, []int) {
	return file_mtn_provider_proto_rawDescGZIP(), []int{10}
}

func (x *ValidateAccountHolderRequest) GetAccountHolderId() string {
	if x != nil {
		return x.AccountHolderId
	}
	return ""
}

func (x *ValidateAccountHolderRequest) GetAccountHolderIdType() string {
	if x != nil {
		return x.AccountHolderIdType
	}
	return ""
}

func (x *ValidateAccountHolderRequest) GetTargetEnvironment() string {
	if x != nil {
		return x.TargetEnvironment
	}
	return ""
}

type ValidateAccountHolderResponse struct {
	state             protoimpl.MessageState `protogen:"open.v1"`
	IsRegistered      bool                   `protobuf:"varint,1,opt,name=is_registered,json=isRegistered,proto3" json:"is_registered,omitempty"`
	Status            string                 `protobuf:"bytes,2,opt,name=status,proto3" json:"status,omitempty"`
	AccountHolderName string                 `protobuf:"bytes,3,opt,name=account_holder_name,json=accountHolderName,proto3" json:"account_holder_name,omitempty"`
	unknownFields     protoimpl.UnknownFields
	sizeCache         protoimpl.SizeCache
}

func (x *ValidateAccountHolderResponse) Reset() {
	*x = ValidateAccountHolderResponse{}
	mi := &file_mtn_provider_proto_msgTypes[11]
	ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
	ms.StoreMessageInfo(mi)
}

func (x *ValidateAccountHolderResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ValidateAccountHolderResponse) ProtoMessage() {}

func (x *ValidateAccountHolderResponse) ProtoReflect() protoreflect.Message {
	mi := &file_mtn_provider_proto_msgTypes[11]
	if x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ValidateAccountHolderResponse.ProtoReflect.Descriptor instead.
func (*ValidateAccountHolderResponse) Descriptor() ([]byte, []int) {
	return file_mtn_provider_proto_rawDescGZIP(), []int{11}
}

func (x *ValidateAccountHolderResponse) GetIsRegistered() bool {
	if x != nil {
		return x.IsRegistered
	}
	return false
}

func (x *ValidateAccountHolderResponse) GetStatus() string {
	if x != nil {
		return x.Status
	}
	return ""
}

func (x *ValidateAccountHolderResponse) GetAccountHolderName() string {
	if x != nil {
		return x.AccountHolderName
	}
	return ""
}

var File_mtn_provider_proto protoreflect.FileDescriptor

const file_mtn_provider_proto_rawDesc = "" +
	"\n" +
	"\x12mtn_provider.proto\x12\tproviders\x1a\x1fgoogle/protobuf/timestamp.proto\"\xbf\x02\n" +
	"\x13RequestToPayRequest\x12\x1f\n" +
	"\vexternal_id\x18\x01 \x01(\tR\n" +
	"externalId\x12-\n" +
	"\x12target_environment\x18\x02 \x01(\tR\x11targetEnvironment\x12\x16\n" +
	"\x06amount\x18\x03 \x01(\x01R\x06amount\x12\x1a\n" +
	"\bcurrency\x18\x04 \x01(\tR\bcurrency\x12\x19\n" +
	"\bpayer_id\x18\x05 \x01(\tR\apayerId\x12\"\n" +
	"\rpayer_id_type\x18\x06 \x01(\tR\vpayerIdType\x12#\n" +
	"\rpayer_message\x18\a \x01(\tR\fpayerMessage\x12\x1d\n" +
	"\n" +
	"payee_note\x18\b \x01(\tR\tpayeeNote\x12!\n" +
	"\fcallback_url\x18\t \x01(\tR\vcallbackUrl\"\x8c\x01\n" +
	"\x14RequestToPayResponse\x12!\n" +
	"\freference_id\x18\x01 \x01(\tR\vreferenceId\x12\x16\n" +
	"\x06status\x18\x02 \x01(\tR\x06status\x129\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\"\xbb\x02\n" +
	"\x0fTransferRequest\x12\x1f\n" +
	"\vexternal_id\x18\x01 \x01(\tR\n" +
	"externalId\x12-\n" +
	"\x12target_environment\x18\x02 \x01(\tR\x11targetEnvironment\x12\x16\n" +
	"\x06amount\x18\x03 \x01(\x01R\x06amount\x12\x1a\n" +
	"\bcurrency\x18\x04 \x01(\tR\bcurrency\x12\x19\n" +
	"\bpayee_id\x18\x05 \x01(\tR\apayeeId\x12\"\n" +
	"\rpayee_id_type\x18\x06 \x01(\tR\vpayeeIdType\x12#\n" +
	"\rpayer_message\x18\a \x01(\tR\fpayerMessage\x12\x1d\n" +
	"\n" +
	"payee_note\x18\b \x01(\tR\tpayeeNote\x12!\n" +
	"\fcallback_url\x18\t \x01(\tR\vcallbackUrl\"\x88\x01\n" +
	"\x10TransferResponse\x12!\n" +
	"\freference_id\x18\x01 \x01(\tR\vreferenceId\x12\x16\n" +
	"\x06status\x18\x02 \x01(\tR\x06status\x129\n" +
	"\n" +
	"created_at\x18\x03 \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\"p\n" +
	"\x1cGetRequestToPayStatusRequest\x12!\n" +
	"\freference_id\x18\x01 \x01(\tR\vreferenceId\x12-\n" +
	"\x12target_environment\x18\x02 \x01(\tR\x11targetEnvironment\"\xdb\x02\n" +
	"\x1dGetRequestToPayStatusResponse\x12\x1f\n" +
	"\vexternal_id\x18\x01 \x01(\tR\n" +
	"externalId\x12\x16\n" +
	"\x06amount\x18\x02 \x01(\tR\x06amount\x12\x1a\n" +
	"\bcurrency\x18\x03 \x01(\tR\bcurrency\x12\x19\n" +
	"\bpayer_id\x18\x04 \x01(\tR\apayerId\x12\"\n" +
	"\rpayer_id_type\x18\x05 \x01(\tR\vpayerIdType\x12\x16\n" +
	"\x06status\x18\x06 \x01(\tR\x06status\x12\x16\n" +
	"\x06reason\x18\a \x01(\tR\x06reason\x129\n" +
	"\n" +
	"created_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12;\n" +
	"\vmodified_at\x18\t \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"modifiedAt\"l\n" +
	"\x18GetTransferStatusRequest\x12!\n" +
	"\freference_id\x18\x01 \x01(\tR\vreferenceId\x12-\n" +
	"\x12target_environment\x18\x02 \x01(\tR\x11targetEnvironment\"\xd7\x02\n" +
	"\x19GetTransferStatusResponse\x12\x1f\n" +
	"\vexternal_id\x18\x01 \x01(\tR\n" +
	"externalId\x12\x16\n" +
	"\x06amount\x18\x02 \x01(\tR\x06amount\x12\x1a\n" +
	"\bcurrency\x18\x03 \x01(\tR\bcurrency\x12\x19\n" +
	"\bpayee_id\x18\x04 \x01(\tR\apayeeId\x12\"\n" +
	"\rpayee_id_type\x18\x05 \x01(\tR\vpayeeIdType\x12\x16\n" +
	"\x06status\x18\x06 \x01(\tR\x06status\x12\x16\n" +
	"\x06reason\x18\a \x01(\tR\x06reason\x129\n" +
	"\n" +
	"created_at\x18\b \x01(\v2\x1a.google.protobuf.TimestampR\tcreatedAt\x12;\n" +
	"\vmodified_at\x18\t \x01(\v2\x1a.google.protobuf.TimestampR\n" +
	"modifiedAt\"I\n" +
	"\x18GetAccountBalanceRequest\x12-\n" +
	"\x12target_environment\x18\x01 \x01(\tR\x11targetEnvironment\"d\n" +
	"\x19GetAccountBalanceResponse\x12+\n" +
	"\x11available_balance\x18\x01 \x01(\tR\x10availableBalance\x12\x1a\n" +
	"\bcurrency\x18\x02 \x01(\tR\bcurrency\"\xae\x01\n" +
	"\x1cValidateAccountHolderRequest\x12*\n" +
	"\x11account_holder_id\x18\x01 \x01(\tR\x0faccountHolderId\x123\n" +
	"\x16account_holder_id_type\x18\x02 \x01(\tR\x13accountHolderIdType\x12-\n" +
	"\x12target_environment\x18\x03 \x01(\tR\x11targetEnvironment\"\x8c\x01\n" +
	"\x1dValidateAccountHolderResponse\x12#\n" +
	"\ris_registered\x18\x01 \x01(\bR\fisRegistered\x12\x16\n" +
	"\x06status\x18\x02 \x01(\tR\x06status\x12.\n" +
	"\x13account_holder_name\x18\x03 \x01(\tR\x11accountHolderName2\xba\x04\n" +
	"\n" +
	"MTNService\x12O\n" +
	"\fRequestToPay\x12\x1e.providers.RequestToPayRequest\x1a\x1f.providers.RequestToPayResponse\x12C\n" +
	"\bTransfer\x12\x1a.providers.TransferRequest\x1a\x1b.providers.TransferResponse\x12j\n" +
	"\x15GetRequestToPayStatus\x12'.providers.GetRequestToPayStatusRequest\x1a(.providers.GetRequestToPayStatusResponse\x12^\n" +
	"\x11GetTransferStatus\x12#.providers.GetTransferStatusRequest\x1a$.providers.GetTransferStatusResponse\x12^\n" +
	"\x11GetAccountBalance\x12#.providers.GetAccountBalanceRequest\x1a$.providers.GetAccountBalanceResponse\x12j\n" +
	"\x15ValidateAccountHolder\x12'.providers.ValidateAccountHolderRequest\x1a(.providers.ValidateAccountHolderResponseB$Z\"github.com/travoroguna/bethela-rpcb\x06proto3"

var (
	file_mtn_provider_proto_rawDescOnce sync.Once
	file_mtn_provider_proto_rawDescData []byte
)

func file_mtn_provider_proto_rawDescGZIP() []byte {
	file_mtn_provider_proto_rawDescOnce.Do(func() {
		file_mtn_provider_proto_rawDescData = protoimpl.X.CompressGZIP(unsafe.Slice(unsafe.StringData(file_mtn_provider_proto_rawDesc), len(file_mtn_provider_proto_rawDesc)))
	})
	return file_mtn_provider_proto_rawDescData
}

var file_mtn_provider_proto_msgTypes = make([]protoimpl.MessageInfo, 12)
var file_mtn_provider_proto_goTypes = []any{
	(*RequestToPayRequest)(nil),           // 0: providers.RequestToPayRequest
	(*RequestToPayResponse)(nil),          // 1: providers.RequestToPayResponse
	(*TransferRequest)(nil),               // 2: providers.TransferRequest
	(*TransferResponse)(nil),              // 3: providers.TransferResponse
	(*GetRequestToPayStatusRequest)(nil),  // 4: providers.GetRequestToPayStatusRequest
	(*GetRequestToPayStatusResponse)(nil), // 5: providers.GetRequestToPayStatusResponse
	(*GetTransferStatusRequest)(nil),      // 6: providers.GetTransferStatusRequest
	(*GetTransferStatusResponse)(nil),     // 7: providers.GetTransferStatusResponse
	(*GetAccountBalanceRequest)(nil),      // 8: providers.GetAccountBalanceRequest
	(*GetAccountBalanceResponse)(nil),     // 9: providers.GetAccountBalanceResponse
	(*ValidateAccountHolderRequest)(nil),  // 10: providers.ValidateAccountHolderRequest
	(*ValidateAccountHolderResponse)(nil), // 11: providers.ValidateAccountHolderResponse
	(*timestamppb.Timestamp)(nil),         // 12: google.protobuf.Timestamp
}
var file_mtn_provider_proto_depIdxs = []int32{
	12, // 0: providers.RequestToPayResponse.created_at:type_name -> google.protobuf.Timestamp
	12, // 1: providers.TransferResponse.created_at:type_name -> google.protobuf.Timestamp
	12, // 2: providers.GetRequestToPayStatusResponse.created_at:type_name -> google.protobuf.Timestamp
	12, // 3: providers.GetRequestToPayStatusResponse.modified_at:type_name -> google.protobuf.Timestamp
	12, // 4: providers.GetTransferStatusResponse.created_at:type_name -> google.protobuf.Timestamp
	12, // 5: providers.GetTransferStatusResponse.modified_at:type_name -> google.protobuf.Timestamp
	0,  // 6: providers.MTNService.RequestToPay:input_type -> providers.RequestToPayRequest
	2,  // 7: providers.MTNService.Transfer:input_type -> providers.TransferRequest
	4,  // 8: providers.MTNService.GetRequestToPayStatus:input_type -> providers.GetRequestToPayStatusRequest
	6,  // 9: providers.MTNService.GetTransferStatus:input_type -> providers.GetTransferStatusRequest
	8,  // 10: providers.MTNService.GetAccountBalance:input_type -> providers.GetAccountBalanceRequest
	10, // 11: providers.MTNService.ValidateAccountHolder:input_type -> providers.ValidateAccountHolderRequest
	1,  // 12: providers.MTNService.RequestToPay:output_type -> providers.RequestToPayResponse
	3,  // 13: providers.MTNService.Transfer:output_type -> providers.TransferResponse
	5,  // 14: providers.MTNService.GetRequestToPayStatus:output_type -> providers.GetRequestToPayStatusResponse
	7,  // 15: providers.MTNService.GetTransferStatus:output_type -> providers.GetTransferStatusResponse
	9,  // 16: providers.MTNService.GetAccountBalance:output_type -> providers.GetAccountBalanceResponse
	11, // 17: providers.MTNService.ValidateAccountHolder:output_type -> providers.ValidateAccountHolderResponse
	12, // [12:18] is the sub-list for method output_type
	6,  // [6:12] is the sub-list for method input_type
	6,  // [6:6] is the sub-list for extension type_name
	6,  // [6:6] is the sub-list for extension extendee
	0,  // [0:6] is the sub-list for field type_name
}

func init() { file_mtn_provider_proto_init() }
func file_mtn_provider_proto_init() {
	if File_mtn_provider_proto != nil {
		return
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: unsafe.Slice(unsafe.StringData(file_mtn_provider_proto_rawDesc), len(file_mtn_provider_proto_rawDesc)),
			NumEnums:      0,
			NumMessages:   12,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_mtn_provider_proto_goTypes,
		DependencyIndexes: file_mtn_provider_proto_depIdxs,
		MessageInfos:      file_mtn_provider_proto_msgTypes,
	}.Build()
	File_mtn_provider_proto = out.File
	file_mtn_provider_proto_goTypes = nil
	file_mtn_provider_proto_depIdxs = nil
}
