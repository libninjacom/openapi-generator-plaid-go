# TransferAuthorizationCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | [optional] 
**AccessToken** | **string** | The Plaid &#x60;access_token&#x60; for the account that will be debited or credited. | 
**AccountId** | **string** | The Plaid &#x60;account_id&#x60; for the account that will be debited or credited. | 
**Type** | [**TransferType**](TransferType.md) |  | 
**Network** | [**TransferNetwork**](TransferNetwork.md) |  | 
**Amount** | **string** | The amount of the transfer (decimal string with two digits of precision e.g. \&quot;10.00\&quot;). | 
**AchClass** | [**ACHClass**](ACHClass.md) |  | 
**User** | [**TransferUserInRequest**](TransferUserInRequest.md) |  | 
**Device** | Pointer to [**TransferAuthorizationDevice**](TransferAuthorizationDevice.md) |  | [optional] 
**OriginationAccountId** | Pointer to **string** | Plaid&#39;s unique identifier for the origination account for this authorization. If not specified, the default account will be used. | [optional] 
**IsoCurrencyCode** | Pointer to **string** | The currency of the transfer amount. The default value is \&quot;USD\&quot;. | [optional] 

## Methods

### NewTransferAuthorizationCreateRequest

`func NewTransferAuthorizationCreateRequest(accessToken string, accountId string, type_ TransferType, network TransferNetwork, amount string, achClass ACHClass, user TransferUserInRequest, ) *TransferAuthorizationCreateRequest`

NewTransferAuthorizationCreateRequest instantiates a new TransferAuthorizationCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferAuthorizationCreateRequestWithDefaults

`func NewTransferAuthorizationCreateRequestWithDefaults() *TransferAuthorizationCreateRequest`

NewTransferAuthorizationCreateRequestWithDefaults instantiates a new TransferAuthorizationCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *TransferAuthorizationCreateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TransferAuthorizationCreateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TransferAuthorizationCreateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *TransferAuthorizationCreateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *TransferAuthorizationCreateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *TransferAuthorizationCreateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *TransferAuthorizationCreateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *TransferAuthorizationCreateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetAccessToken

`func (o *TransferAuthorizationCreateRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *TransferAuthorizationCreateRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *TransferAuthorizationCreateRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.


### GetAccountId

`func (o *TransferAuthorizationCreateRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransferAuthorizationCreateRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransferAuthorizationCreateRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.


### GetType

`func (o *TransferAuthorizationCreateRequest) GetType() TransferType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *TransferAuthorizationCreateRequest) GetTypeOk() (*TransferType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *TransferAuthorizationCreateRequest) SetType(v TransferType)`

SetType sets Type field to given value.


### GetNetwork

`func (o *TransferAuthorizationCreateRequest) GetNetwork() TransferNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *TransferAuthorizationCreateRequest) GetNetworkOk() (*TransferNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *TransferAuthorizationCreateRequest) SetNetwork(v TransferNetwork)`

SetNetwork sets Network field to given value.


### GetAmount

`func (o *TransferAuthorizationCreateRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferAuthorizationCreateRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferAuthorizationCreateRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetAchClass

`func (o *TransferAuthorizationCreateRequest) GetAchClass() ACHClass`

GetAchClass returns the AchClass field if non-nil, zero value otherwise.

### GetAchClassOk

`func (o *TransferAuthorizationCreateRequest) GetAchClassOk() (*ACHClass, bool)`

GetAchClassOk returns a tuple with the AchClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchClass

`func (o *TransferAuthorizationCreateRequest) SetAchClass(v ACHClass)`

SetAchClass sets AchClass field to given value.


### GetUser

`func (o *TransferAuthorizationCreateRequest) GetUser() TransferUserInRequest`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TransferAuthorizationCreateRequest) GetUserOk() (*TransferUserInRequest, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TransferAuthorizationCreateRequest) SetUser(v TransferUserInRequest)`

SetUser sets User field to given value.


### GetDevice

`func (o *TransferAuthorizationCreateRequest) GetDevice() TransferAuthorizationDevice`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *TransferAuthorizationCreateRequest) GetDeviceOk() (*TransferAuthorizationDevice, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *TransferAuthorizationCreateRequest) SetDevice(v TransferAuthorizationDevice)`

SetDevice sets Device field to given value.

### HasDevice

`func (o *TransferAuthorizationCreateRequest) HasDevice() bool`

HasDevice returns a boolean if a field has been set.

### GetOriginationAccountId

`func (o *TransferAuthorizationCreateRequest) GetOriginationAccountId() string`

GetOriginationAccountId returns the OriginationAccountId field if non-nil, zero value otherwise.

### GetOriginationAccountIdOk

`func (o *TransferAuthorizationCreateRequest) GetOriginationAccountIdOk() (*string, bool)`

GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationAccountId

`func (o *TransferAuthorizationCreateRequest) SetOriginationAccountId(v string)`

SetOriginationAccountId sets OriginationAccountId field to given value.

### HasOriginationAccountId

`func (o *TransferAuthorizationCreateRequest) HasOriginationAccountId() bool`

HasOriginationAccountId returns a boolean if a field has been set.

### GetIsoCurrencyCode

`func (o *TransferAuthorizationCreateRequest) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *TransferAuthorizationCreateRequest) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *TransferAuthorizationCreateRequest) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.

### HasIsoCurrencyCode

`func (o *TransferAuthorizationCreateRequest) HasIsoCurrencyCode() bool`

HasIsoCurrencyCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


