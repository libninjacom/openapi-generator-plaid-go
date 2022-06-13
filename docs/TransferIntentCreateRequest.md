# TransferIntentCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | 
**Secret** | **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | 
**AccountId** | Pointer to **NullableString** | The Plaid &#x60;account_id&#x60; for the account that will be debited or credited. | [optional] 
**Mode** | [**TransferIntentCreateMode**](TransferIntentCreateMode.md) |  | 
**Amount** | **string** | The amount of the transfer (decimal string with two digits of precision e.g. \&quot;10.00\&quot;). | 
**Description** | **string** | A description for the underlying transfer. Maximum of 8 characters. | 
**AchClass** | [**ACHClass**](ACHClass.md) |  | 
**OriginationAccountId** | Pointer to **NullableString** | Plaid’s unique identifier for the origination account for the intent. If not provided, the default account will be used. | [optional] 
**User** | [**TransferUserInRequest**](TransferUserInRequest.md) |  | 
**Metadata** | Pointer to **map[string]string** | The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply: - The JSON values must be Strings (no nested JSON objects allowed) - Only ASCII characters may be used - Maximum of 50 key/value pairs - Maximum key length of 40 characters - Maximum value length of 500 characters  | [optional] 
**IsoCurrencyCode** | Pointer to **string** | The currency of the transfer amount, e.g. \&quot;USD\&quot; | [optional] 

## Methods

### NewTransferIntentCreateRequest

`func NewTransferIntentCreateRequest(clientId string, secret string, mode TransferIntentCreateMode, amount string, description string, achClass ACHClass, user TransferUserInRequest, ) *TransferIntentCreateRequest`

NewTransferIntentCreateRequest instantiates a new TransferIntentCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferIntentCreateRequestWithDefaults

`func NewTransferIntentCreateRequestWithDefaults() *TransferIntentCreateRequest`

NewTransferIntentCreateRequestWithDefaults instantiates a new TransferIntentCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *TransferIntentCreateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *TransferIntentCreateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *TransferIntentCreateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.


### GetSecret

`func (o *TransferIntentCreateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *TransferIntentCreateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *TransferIntentCreateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.


### GetAccountId

`func (o *TransferIntentCreateRequest) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransferIntentCreateRequest) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransferIntentCreateRequest) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *TransferIntentCreateRequest) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *TransferIntentCreateRequest) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *TransferIntentCreateRequest) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetMode

`func (o *TransferIntentCreateRequest) GetMode() TransferIntentCreateMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TransferIntentCreateRequest) GetModeOk() (*TransferIntentCreateMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TransferIntentCreateRequest) SetMode(v TransferIntentCreateMode)`

SetMode sets Mode field to given value.


### GetAmount

`func (o *TransferIntentCreateRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferIntentCreateRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferIntentCreateRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetDescription

`func (o *TransferIntentCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransferIntentCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransferIntentCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAchClass

`func (o *TransferIntentCreateRequest) GetAchClass() ACHClass`

GetAchClass returns the AchClass field if non-nil, zero value otherwise.

### GetAchClassOk

`func (o *TransferIntentCreateRequest) GetAchClassOk() (*ACHClass, bool)`

GetAchClassOk returns a tuple with the AchClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchClass

`func (o *TransferIntentCreateRequest) SetAchClass(v ACHClass)`

SetAchClass sets AchClass field to given value.


### GetOriginationAccountId

`func (o *TransferIntentCreateRequest) GetOriginationAccountId() string`

GetOriginationAccountId returns the OriginationAccountId field if non-nil, zero value otherwise.

### GetOriginationAccountIdOk

`func (o *TransferIntentCreateRequest) GetOriginationAccountIdOk() (*string, bool)`

GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationAccountId

`func (o *TransferIntentCreateRequest) SetOriginationAccountId(v string)`

SetOriginationAccountId sets OriginationAccountId field to given value.

### HasOriginationAccountId

`func (o *TransferIntentCreateRequest) HasOriginationAccountId() bool`

HasOriginationAccountId returns a boolean if a field has been set.

### SetOriginationAccountIdNil

`func (o *TransferIntentCreateRequest) SetOriginationAccountIdNil(b bool)`

 SetOriginationAccountIdNil sets the value for OriginationAccountId to be an explicit nil

### UnsetOriginationAccountId
`func (o *TransferIntentCreateRequest) UnsetOriginationAccountId()`

UnsetOriginationAccountId ensures that no value is present for OriginationAccountId, not even an explicit nil
### GetUser

`func (o *TransferIntentCreateRequest) GetUser() TransferUserInRequest`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TransferIntentCreateRequest) GetUserOk() (*TransferUserInRequest, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TransferIntentCreateRequest) SetUser(v TransferUserInRequest)`

SetUser sets User field to given value.


### GetMetadata

`func (o *TransferIntentCreateRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TransferIntentCreateRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TransferIntentCreateRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TransferIntentCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *TransferIntentCreateRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *TransferIntentCreateRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetIsoCurrencyCode

`func (o *TransferIntentCreateRequest) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *TransferIntentCreateRequest) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *TransferIntentCreateRequest) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.

### HasIsoCurrencyCode

`func (o *TransferIntentCreateRequest) HasIsoCurrencyCode() bool`

HasIsoCurrencyCode returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


