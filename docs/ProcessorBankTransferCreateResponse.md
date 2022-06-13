# ProcessorBankTransferCreateResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BankTransfer** | [**BankTransfer**](BankTransfer.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewProcessorBankTransferCreateResponse

`func NewProcessorBankTransferCreateResponse(bankTransfer BankTransfer, requestId string, ) *ProcessorBankTransferCreateResponse`

NewProcessorBankTransferCreateResponse instantiates a new ProcessorBankTransferCreateResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorBankTransferCreateResponseWithDefaults

`func NewProcessorBankTransferCreateResponseWithDefaults() *ProcessorBankTransferCreateResponse`

NewProcessorBankTransferCreateResponseWithDefaults instantiates a new ProcessorBankTransferCreateResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBankTransfer

`func (o *ProcessorBankTransferCreateResponse) GetBankTransfer() BankTransfer`

GetBankTransfer returns the BankTransfer field if non-nil, zero value otherwise.

### GetBankTransferOk

`func (o *ProcessorBankTransferCreateResponse) GetBankTransferOk() (*BankTransfer, bool)`

GetBankTransferOk returns a tuple with the BankTransfer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBankTransfer

`func (o *ProcessorBankTransferCreateResponse) SetBankTransfer(v BankTransfer)`

SetBankTransfer sets BankTransfer field to given value.


### GetRequestId

`func (o *ProcessorBankTransferCreateResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ProcessorBankTransferCreateResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ProcessorBankTransferCreateResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


