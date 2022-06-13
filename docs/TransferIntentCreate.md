# TransferIntentCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Plaid&#39;s unique identifier for the transfer intent object. | 
**Created** | **time.Time** | The datetime the transfer was created. This will be of the form &#x60;2006-01-02T15:04:05Z&#x60;. | 
**Status** | **string** | The status of the transfer intent.  - &#x60;PENDING&#x60; – The transfer intent is pending. - &#x60;SUCCEEDED&#x60; – The transfer intent was successfully created. - &#x60;FAILED&#x60; – The transfer intent was unable to be created. | 
**AccountId** | Pointer to **NullableString** | The Plaid &#x60;account_id&#x60; for the account that will be debited or credited. Returned only if &#x60;account_id&#x60; was set on intent creation. | [optional] 
**OriginationAccountId** | **string** | Plaid’s unique identifier for the origination account for the intent. If not provided, the default account will be used. | 
**Amount** | **string** | The amount of the transfer (decimal string with two digits of precision e.g. \&quot;10.00\&quot;). | 
**Mode** | [**TransferIntentCreateMode**](TransferIntentCreateMode.md) |  | 
**AchClass** | [**ACHClass**](ACHClass.md) |  | 
**User** | [**TransferUserInResponse**](TransferUserInResponse.md) |  | 
**Description** | **string** | A description for the underlying transfer. Maximum of 8 characters. | 
**Metadata** | Pointer to **map[string]string** | The Metadata object is a mapping of client-provided string fields to any string value. The following limitations apply: - The JSON values must be Strings (no nested JSON objects allowed) - Only ASCII characters may be used - Maximum of 50 key/value pairs - Maximum key length of 40 characters - Maximum value length of 500 characters  | [optional] 
**IsoCurrencyCode** | **string** | The currency of the transfer amount, e.g. \&quot;USD\&quot; | 

## Methods

### NewTransferIntentCreate

`func NewTransferIntentCreate(id string, created time.Time, status string, originationAccountId string, amount string, mode TransferIntentCreateMode, achClass ACHClass, user TransferUserInResponse, description string, isoCurrencyCode string, ) *TransferIntentCreate`

NewTransferIntentCreate instantiates a new TransferIntentCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferIntentCreateWithDefaults

`func NewTransferIntentCreateWithDefaults() *TransferIntentCreate`

NewTransferIntentCreateWithDefaults instantiates a new TransferIntentCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TransferIntentCreate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TransferIntentCreate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TransferIntentCreate) SetId(v string)`

SetId sets Id field to given value.


### GetCreated

`func (o *TransferIntentCreate) GetCreated() time.Time`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *TransferIntentCreate) GetCreatedOk() (*time.Time, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *TransferIntentCreate) SetCreated(v time.Time)`

SetCreated sets Created field to given value.


### GetStatus

`func (o *TransferIntentCreate) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *TransferIntentCreate) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *TransferIntentCreate) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetAccountId

`func (o *TransferIntentCreate) GetAccountId() string`

GetAccountId returns the AccountId field if non-nil, zero value otherwise.

### GetAccountIdOk

`func (o *TransferIntentCreate) GetAccountIdOk() (*string, bool)`

GetAccountIdOk returns a tuple with the AccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountId

`func (o *TransferIntentCreate) SetAccountId(v string)`

SetAccountId sets AccountId field to given value.

### HasAccountId

`func (o *TransferIntentCreate) HasAccountId() bool`

HasAccountId returns a boolean if a field has been set.

### SetAccountIdNil

`func (o *TransferIntentCreate) SetAccountIdNil(b bool)`

 SetAccountIdNil sets the value for AccountId to be an explicit nil

### UnsetAccountId
`func (o *TransferIntentCreate) UnsetAccountId()`

UnsetAccountId ensures that no value is present for AccountId, not even an explicit nil
### GetOriginationAccountId

`func (o *TransferIntentCreate) GetOriginationAccountId() string`

GetOriginationAccountId returns the OriginationAccountId field if non-nil, zero value otherwise.

### GetOriginationAccountIdOk

`func (o *TransferIntentCreate) GetOriginationAccountIdOk() (*string, bool)`

GetOriginationAccountIdOk returns a tuple with the OriginationAccountId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOriginationAccountId

`func (o *TransferIntentCreate) SetOriginationAccountId(v string)`

SetOriginationAccountId sets OriginationAccountId field to given value.


### GetAmount

`func (o *TransferIntentCreate) GetAmount() string`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *TransferIntentCreate) GetAmountOk() (*string, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *TransferIntentCreate) SetAmount(v string)`

SetAmount sets Amount field to given value.


### GetMode

`func (o *TransferIntentCreate) GetMode() TransferIntentCreateMode`

GetMode returns the Mode field if non-nil, zero value otherwise.

### GetModeOk

`func (o *TransferIntentCreate) GetModeOk() (*TransferIntentCreateMode, bool)`

GetModeOk returns a tuple with the Mode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMode

`func (o *TransferIntentCreate) SetMode(v TransferIntentCreateMode)`

SetMode sets Mode field to given value.


### GetAchClass

`func (o *TransferIntentCreate) GetAchClass() ACHClass`

GetAchClass returns the AchClass field if non-nil, zero value otherwise.

### GetAchClassOk

`func (o *TransferIntentCreate) GetAchClassOk() (*ACHClass, bool)`

GetAchClassOk returns a tuple with the AchClass field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchClass

`func (o *TransferIntentCreate) SetAchClass(v ACHClass)`

SetAchClass sets AchClass field to given value.


### GetUser

`func (o *TransferIntentCreate) GetUser() TransferUserInResponse`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *TransferIntentCreate) GetUserOk() (*TransferUserInResponse, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *TransferIntentCreate) SetUser(v TransferUserInResponse)`

SetUser sets User field to given value.


### GetDescription

`func (o *TransferIntentCreate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransferIntentCreate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransferIntentCreate) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetMetadata

`func (o *TransferIntentCreate) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *TransferIntentCreate) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *TransferIntentCreate) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *TransferIntentCreate) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### SetMetadataNil

`func (o *TransferIntentCreate) SetMetadataNil(b bool)`

 SetMetadataNil sets the value for Metadata to be an explicit nil

### UnsetMetadata
`func (o *TransferIntentCreate) UnsetMetadata()`

UnsetMetadata ensures that no value is present for Metadata, not even an explicit nil
### GetIsoCurrencyCode

`func (o *TransferIntentCreate) GetIsoCurrencyCode() string`

GetIsoCurrencyCode returns the IsoCurrencyCode field if non-nil, zero value otherwise.

### GetIsoCurrencyCodeOk

`func (o *TransferIntentCreate) GetIsoCurrencyCodeOk() (*string, bool)`

GetIsoCurrencyCodeOk returns a tuple with the IsoCurrencyCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsoCurrencyCode

`func (o *TransferIntentCreate) SetIsoCurrencyCode(v string)`

SetIsoCurrencyCode sets IsoCurrencyCode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


