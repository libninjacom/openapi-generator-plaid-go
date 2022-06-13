# TransferAuthorizationProposedTransfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchClass** | [**ACHClass**](ACHClass.md) |  | 
**AccountId** | **string** | The Plaid &#x60;account_id&#x60; for the account that will be debited or credited. | 
**Type** | [**TransferType**](TransferType.md) |  | 
**User** | [**TransferUserInResponse**](TransferUserInResponse.md) |  | 
**Amount** | **string** | The amount of the transfer (decimal string with two digits of precision e.g. \&quot;10.00\&quot;). | 
**Network** | **string** | The network or rails used for the transfer. | 
**OriginationAccountId** | **string** | Plaid&#39;s unique identifier for the origination account that was used for this transfer. | 
**IsoCurrencyCode** | **string** | The currency of the transfer amount. The default value is \&quot;USD\&quot;. | 

## Methods

### NewTransferAuthorizationProposedTransfer

`func NewTransferAuthorizationProposedTransfer(achClass ACHClass, accountId string, type_ TransferType, user TransferUserInResponse, amount string, network string, originationAccountId string, isoCurrencyCode string, ) *TransferAuthorizationProposedTransfer`

NewTransferAuthorizationProposedTransfer instantiates a new TransferAuthorizationProposedTransfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferAuthorizationProposedTransferWithDefaults

`func NewTransferAuthorizationProposedTransferWithDefaults() *TransferAuthorizationProposedTransfer`

NewTransferAuthorizationProposedTransferWithDefaults instantiates a new TransferAuthorizationProposedTransfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchClass

`func (o *TransferAuthorizationProposedTransfer) GetAchClass() ACHClass`

GetAchClass returns the AchClass field if non-nil, zero value otherwise.

### GetAchClassOk

`func (o *TransferAuthorizationProposedTransfer) GetAchClassOk() (*ACHClass, bool)`

GetAchClassOk returns a tuple with the AchClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchClass

`func (o *TransferAuthorizationProposedTransfer) SetAchClass(v ACHClass)`

SetAchClass sets AchClass field to given value.


### GetAccountId

`func (o *TransferAuthorizationProposedTransfer) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransferAuthorizationProposedTransfer) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransferAuthorizationProposedTransfer) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetType

`func (o *TransferAuthorizationProposedTransfer) GetType() TransferType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransferAuthorizationProposedTransfer) GetTypeOk() (*TransferType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransferAuthorizationProposedTransfer) SetType(v TransferType)`

SetType sets Type field to given value.


### GetUser

`func (o *TransferAuthorizationProposedTransfer) GetUser() TransferUserInResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TransferAuthorizationProposedTransfer) GetUserOk() (*TransferUserInResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TransferAuthorizationProposedTransfer) SetUser(v TransferUserInResponse)`

SetUser sets User field to given value.


### GetAmount

`func (o *TransferAuthorizationProposedTransfer) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferAuthorizationProposedTransfer) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferAuthorizationProposedTransfer) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetNetwork

`func (o *TransferAuthorizationProposedTransfer) GetNetwork() string`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *TransferAuthorizationProposedTransfer) GetNetworkOk() (*string, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *TransferAuthorizationProposedTransfer) SetNetwork(v string)`

SetNetwork sets Network field to given value.


### GetOriginationAccountId

`func (o *TransferAuthorizationProposedTransfer) GetOriginationAccountId() string`

GetOriginationAccountId returns the OriginationAccountId field if non-nil, zero value otherwise.

### GetOriginationAccountIdOk

`func (o *TransferAuthorizationProposedTransfer) GetOriginationAccountIdOk() (*string, bool)`

GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationAccountId

`func (o *TransferAuthorizationProposedTransfer) SetOriginationAccountId(v string)`

SetOriginationAccountId sets OriginationAccountId field to given value.


### GetIsoCurrencyCode

`func (o *TransferAuthorizationProposedTransfer) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *TransferAuthorizationProposedTransfer) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *TransferAuthorizationProposedTransfer) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


