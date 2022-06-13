# WalletGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**WalletId** | **string** | A unique ID identifying the e-wallet | 
**Balance** | [**NullableWalletBalance**](WalletBalance.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewWalletGetResponse

`func NewWalletGetResponse(walletId string, balance NullableWalletBalance, requestId string, ) *WalletGetResponse`

NewWalletGetResponse instantiates a new WalletGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWalletGetResponseWithDefaults

`func NewWalletGetResponseWithDefaults() *WalletGetResponse`

NewWalletGetResponseWithDefaults instantiates a new WalletGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetWalletId

`func (o *WalletGetResponse) GetWalletId() string`

GetWalletId returns the WalletId field if non-nil, zero value otherwise.

### GetWalletIdOk

`func (o *WalletGetResponse) GetWalletIdOk() (*string, bool)`

GetWalletIdOk returns a tuple with the WalletId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWalletId

`func (o *WalletGetResponse) SetWalletId(v string)`

SetWalletId sets WalletId field to given value.


### GetBalance

`func (o *WalletGetResponse) GetBalance() WalletBalance`

GetBalance returns the Balance field if non-nil, zero value otherwise.

### GetBalanceOk

`func (o *WalletGetResponse) GetBalanceOk() (*WalletBalance, bool)`

GetBalanceOk returns a tuple with the Balance field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBalance

`func (o *WalletGetResponse) SetBalance(v WalletBalance)`

SetBalance sets Balance field to given value.


### SetBalanceNil

`func (o *WalletGetResponse) SetBalanceNil(b bool)`

 SetBalanceNil sets the value for Balance to be an explicit nil

### UnsetBalance
`func (o *WalletGetResponse) UnsetBalance()`

UnsetBalance ensures that no value is present for Balance, not even an explicit nil
### GetRequestId

`func (o *WalletGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *WalletGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *WalletGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


