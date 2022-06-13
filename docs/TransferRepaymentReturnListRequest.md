# TransferRepaymentReturnListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | [optional] 
**RepaymentId** | **string** | Identifier of the repayment to query. | 
**Count** | Pointer to **NullableInt32** | The maximum number of repayments to return. | [optional] [default to 25]
**Offset** | Pointer to **int32** | The number of repayments to skip before returning results. | [optional] [default to 0]

## Methods

### NewTransferRepaymentReturnListRequest

`func NewTransferRepaymentReturnListRequest(repaymentId string, ) *TransferRepaymentReturnListRequest`

NewTransferRepaymentReturnListRequest instantiates a new TransferRepaymentReturnListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferRepaymentReturnListRequestWithDefaults

`func NewTransferRepaymentReturnListRequestWithDefaults() *TransferRepaymentReturnListRequest`

NewTransferRepaymentReturnListRequestWithDefaults instantiates a new TransferRepaymentReturnListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *TransferRepaymentReturnListRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TransferRepaymentReturnListRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TransferRepaymentReturnListRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *TransferRepaymentReturnListRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *TransferRepaymentReturnListRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *TransferRepaymentReturnListRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *TransferRepaymentReturnListRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *TransferRepaymentReturnListRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetRepaymentId

`func (o *TransferRepaymentReturnListRequest) GetRepaymentId() string`

GetRepaymentId returns the RepaymentId field if non-nil, zero value otherwise.

### GetRepaymentIdOk

`func (o *TransferRepaymentReturnListRequest) GetRepaymentIdOk() (*string, bool)`

GetRepaymentIdOk returns a tuple with the RepaymentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepaymentId

`func (o *TransferRepaymentReturnListRequest) SetRepaymentId(v string)`

SetRepaymentId sets RepaymentId field to given value.


### GetCount

`func (o *TransferRepaymentReturnListRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *TransferRepaymentReturnListRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *TransferRepaymentReturnListRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *TransferRepaymentReturnListRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *TransferRepaymentReturnListRequest) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *TransferRepaymentReturnListRequest) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetOffset

`func (o *TransferRepaymentReturnListRequest) GetOffset() int32`

GetOffset returns the Offset field if non-nil, zero value otherwise.

### GetOffsetOk

`func (o *TransferRepaymentReturnListRequest) GetOffsetOk() (*int32, bool)`

GetOffsetOk returns a tuple with the Offset field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffset

`func (o *TransferRepaymentReturnListRequest) SetOffset(v int32)`

SetOffset sets Offset field to given value.

### HasOffset

`func (o *TransferRepaymentReturnListRequest) HasOffset() bool`

HasOffset returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


