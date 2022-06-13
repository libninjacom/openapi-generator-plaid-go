# ProcessorBankTransferCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | [optional] 
**IdempotencyKey** | **string** | A random key provided by the client, per unique bank transfer. Maximum of 50 characters.  The API supports idempotency for safely retrying requests without accidentally performing the same operation twice. For example, if a request to create a bank transfer fails due to a network connection error, you can retry the request with the same idempotency key to guarantee that only a single bank transfer is created. | 
**ProcessorToken** | **string** | The processor token obtained from the Plaid integration partner. Processor tokens are in the format: &#x60;processor-&lt;environment&gt;-&lt;identifier&gt;&#x60; | 
**Type** | [**BankTransferType**](BankTransferType.md) |  | 
**Network** | [**BankTransferNetwork**](BankTransferNetwork.md) |  | 
**Amount** | **string** | The amount of the bank transfer (decimal string with two digits of precision e.g. \&quot;10.00\&quot;). | 
**IsoCurrencyCode** | **string** | The currency of the transfer amount – should be set to \&quot;USD\&quot;. | 
**Description** | **string** | The transfer description. Maximum of 10 characters. | 
**AchClass** | Pointer to [**ACHClass**](ACHClass.md) |  | [optional] 
**User** | [**BankTransferUser**](BankTransferUser.md) |  | 
**CustomTag** | Pointer to **NullableString** | An arbitrary string provided by the client for storage with the bank transfer. May be up to 100 characters. | [optional] 
**Metadata** | Pointer to **map[string]string** | The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply: - The JSON values must be Strings (no nested JSON objects allowed) - Only ASCII characters may be used - Maximum of 50 key/value pairs - Maximum key length of 40 characters - Maximum value length of 500 characters  | [optional] 
**OriginationAccountId** | Pointer to **NullableString** | Plaid’s unique identifier for the origination account for this transfer. If you have more than one origination account, this value must be specified. | [optional] 

## Methods

### NewProcessorBankTransferCreateRequest

`func NewProcessorBankTransferCreateRequest(idempotencyKey string, processorToken string, type_ BankTransferType, network BankTransferNetwork, amount string, isoCurrencyCode string, description string, user BankTransferUser, ) *ProcessorBankTransferCreateRequest`

NewProcessorBankTransferCreateRequest instantiates a new ProcessorBankTransferCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorBankTransferCreateRequestWithDefaults

`func NewProcessorBankTransferCreateRequestWithDefaults() *ProcessorBankTransferCreateRequest`

NewProcessorBankTransferCreateRequestWithDefaults instantiates a new ProcessorBankTransferCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *ProcessorBankTransferCreateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *ProcessorBankTransferCreateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *ProcessorBankTransferCreateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *ProcessorBankTransferCreateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *ProcessorBankTransferCreateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *ProcessorBankTransferCreateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *ProcessorBankTransferCreateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *ProcessorBankTransferCreateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetIdempotencyKey

`func (o *ProcessorBankTransferCreateRequest) GetIdempotencyKey() string`

GetIdempotencyKey returns the IdempotencyKey field if non-nil, zero value otherwise.

### GetIdempotencyKeyOk

`func (o *ProcessorBankTransferCreateRequest) GetIdempotencyKeyOk() (*string, bool)`

GetIdempotencyKeyOk returns a tuple with the IdempotencyKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdempotencyKey

`func (o *ProcessorBankTransferCreateRequest) SetIdempotencyKey(v string)`

SetIdempotencyKey sets IdempotencyKey field to given value.


### GetProcessorToken

`func (o *ProcessorBankTransferCreateRequest) GetProcessorToken() string`

GetProcessorToken returns the ProcessorToken field if non-nil, zero value otherwise.

### GetProcessorTokenOk

`func (o *ProcessorBankTransferCreateRequest) GetProcessorTokenOk() (*string, bool)`

GetProcessorTokenOk returns a tuple with the ProcessorToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProcessorToken

`func (o *ProcessorBankTransferCreateRequest) SetProcessorToken(v string)`

SetProcessorToken sets ProcessorToken field to given value.


### GetType

`func (o *ProcessorBankTransferCreateRequest) GetType() BankTransferType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *ProcessorBankTransferCreateRequest) GetTypeOk() (*BankTransferType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *ProcessorBankTransferCreateRequest) SetType(v BankTransferType)`

SetType sets Type field to given value.


### GetNetwork

`func (o *ProcessorBankTransferCreateRequest) GetNetwork() BankTransferNetwork`

GetNetwork returns the Network field if non-nil, zero value otherwise.

### GetNetworkOk

`func (o *ProcessorBankTransferCreateRequest) GetNetworkOk() (*BankTransferNetwork, bool)`

GetNetworkOk returns a tuple with the Network field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetwork

`func (o *ProcessorBankTransferCreateRequest) SetNetwork(v BankTransferNetwork)`

SetNetwork sets Network field to given value.


### GetAmount

`func (o *ProcessorBankTransferCreateRequest) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *ProcessorBankTransferCreateRequest) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *ProcessorBankTransferCreateRequest) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetIsoCurrencyCode

`func (o *ProcessorBankTransferCreateRequest) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *ProcessorBankTransferCreateRequest) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *ProcessorBankTransferCreateRequest) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.


### GetDescription

`func (o *ProcessorBankTransferCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ProcessorBankTransferCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ProcessorBankTransferCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetAchClass

`func (o *ProcessorBankTransferCreateRequest) GetAchClass() ACHClass`

GetAchClass returns the AchClass field if non-nil, zero value otherwise.

### GetAchClassOk

`func (o *ProcessorBankTransferCreateRequest) GetAchClassOk() (*ACHClass, bool)`

GetAchClassOk returns a tuple with the AchClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchClass

`func (o *ProcessorBankTransferCreateRequest) SetAchClass(v ACHClass)`

SetAchClass sets AchClass field to given value.

### HasAchClass

`func (o *ProcessorBankTransferCreateRequest) HasAchClass() bool`

HasAchClass returns a boolean if a field has been set.

### GetUser

`func (o *ProcessorBankTransferCreateRequest) GetUser() BankTransferUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *ProcessorBankTransferCreateRequest) GetUserOk() (*BankTransferUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *ProcessorBankTransferCreateRequest) SetUser(v BankTransferUser)`

SetUser sets User field to given value.


### GetCustomTag

`func (o *ProcessorBankTransferCreateRequest) GetCustomTag() string`

GetCustomTag returns the CustomTag field if non-nil, zero value otherwise.

### GetCustomTagOk

`func (o *ProcessorBankTransferCreateRequest) GetCustomTagOk() (*string, bool)`

GetCustomTagOk returns a tuple with the CustomTag field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustomTag

`func (o *ProcessorBankTransferCreateRequest) SetCustomTag(v string)`

SetCustomTag sets CustomTag field to given value.

### HasCustomTag

`func (o *ProcessorBankTransferCreateRequest) HasCustomTag() bool`

HasCustomTag returns a boolean if a field has been set.

### SetCustomTagNil

`func (o *ProcessorBankTransferCreateRequest) SetCustomTagNil(b bool)`

 SetCustomTagNil sets the value for CustomTag to be an explicit nil

### UnsetCustomTag
`func (o *ProcessorBankTransferCreateRequest) UnsetCustomTag()`

UnsetCustomTag ensures that no value is present for CustomTag, not even an explicit nil
### GetMetadata

`func (o *ProcessorBankTransferCreateRequest) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ProcessorBankTransferCreateRequest) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ProcessorBankTransferCreateRequest) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ProcessorBankTransferCreateRequest) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *ProcessorBankTransferCreateRequest) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *ProcessorBankTransferCreateRequest) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetOriginationAccountId

`func (o *ProcessorBankTransferCreateRequest) GetOriginationAccountId() string`

GetOriginationAccountId returns the OriginationAccountId field if non-nil, zero value otherwise.

### GetOriginationAccountIdOk

`func (o *ProcessorBankTransferCreateRequest) GetOriginationAccountIdOk() (*string, bool)`

GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationAccountId

`func (o *ProcessorBankTransferCreateRequest) SetOriginationAccountId(v string)`

SetOriginationAccountId sets OriginationAccountId field to given value.

### HasOriginationAccountId

`func (o *ProcessorBankTransferCreateRequest) HasOriginationAccountId() bool`

HasOriginationAccountId returns a boolean if a field has been set.

### SetOriginationAccountIdNil

`func (o *ProcessorBankTransferCreateRequest) SetOriginationAccountIdNil(b bool)`

 SetOriginationAccountIdNil sets the value for OriginationAccountId to be an explicit nil

### UnsetOriginationAccountId
`func (o *ProcessorBankTransferCreateRequest) UnsetOriginationAccountId()`

UnsetOriginationAccountId ensures that no value is present for OriginationAccountId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


