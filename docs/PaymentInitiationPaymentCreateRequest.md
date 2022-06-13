# PaymentInitiationPaymentCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | [optional] 
**RecipientId** | **string** | The ID of the recipient the payment is for. | 
**Reference** | **string** | A reference for the payment. This must be an alphanumeric string with at most 18 characters and must not contain any special characters (since not all institutions support them). | 
**Amount** | [**PaymentAmount**](PaymentAmount.md) |  | 
**Schedule** | Pointer to [**ExternalPaymentScheduleRequest**](ExternalPaymentScheduleRequest.md) |  | [optional] 
**Options** | Pointer to [**NullableExternalPaymentOptions**](ExternalPaymentOptions.md) |  | [optional] 

## Methods

### NewPaymentInitiationPaymentCreateRequest

`func NewPaymentInitiationPaymentCreateRequest(recipientId string, reference string, amount PaymentAmount, ) *PaymentInitiationPaymentCreateRequest`

NewPaymentInitiationPaymentCreateRequest instantiates a new PaymentInitiationPaymentCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiationPaymentCreateRequestWithDefaults

`func NewPaymentInitiationPaymentCreateRequestWithDefaults() *PaymentInitiationPaymentCreateRequest`

NewPaymentInitiationPaymentCreateRequestWithDefaults instantiates a new PaymentInitiationPaymentCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *PaymentInitiationPaymentCreateRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PaymentInitiationPaymentCreateRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PaymentInitiationPaymentCreateRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PaymentInitiationPaymentCreateRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *PaymentInitiationPaymentCreateRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PaymentInitiationPaymentCreateRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PaymentInitiationPaymentCreateRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PaymentInitiationPaymentCreateRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetRecipientId

`func (o *PaymentInitiationPaymentCreateRequest) GetRecipientId() string`

GetRecipientId returns the RecipientId field if non-nil, zero value otherwise.

### GetRecipientIdOk

`func (o *PaymentInitiationPaymentCreateRequest) GetRecipientIdOk() (*string, bool)`

GetRecipientIdOk returns a tuple with the RecipientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientId

`func (o *PaymentInitiationPaymentCreateRequest) SetRecipientId(v string)`

SetRecipientId sets RecipientId field to given value.


### GetReference

`func (o *PaymentInitiationPaymentCreateRequest) GetReference() string`

GetReference returns the Reference field if non-nil, zero value otherwise.

### GetReferenceOk

`func (o *PaymentInitiationPaymentCreateRequest) GetReferenceOk() (*string, bool)`

GetReferenceOk returns a tuple with the Reference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReference

`func (o *PaymentInitiationPaymentCreateRequest) SetReference(v string)`

SetReference sets Reference field to given value.


### GetAmount

`func (o *PaymentInitiationPaymentCreateRequest) GetAmount() PaymentAmount`

GetAmount returns the Amount field if non-nil, zero value otherwise.

### GetAmountOk

`func (o *PaymentInitiationPaymentCreateRequest) GetAmountOk() (*PaymentAmount, bool)`

GetAmountOk returns a tuple with the Amount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmount

`func (o *PaymentInitiationPaymentCreateRequest) SetAmount(v PaymentAmount)`

SetAmount sets Amount field to given value.


### GetSchedule

`func (o *PaymentInitiationPaymentCreateRequest) GetSchedule() ExternalPaymentScheduleRequest`

GetSchedule returns the Schedule field if non-nil, zero value otherwise.

### GetScheduleOk

`func (o *PaymentInitiationPaymentCreateRequest) GetScheduleOk() (*ExternalPaymentScheduleRequest, bool)`

GetScheduleOk returns a tuple with the Schedule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedule

`func (o *PaymentInitiationPaymentCreateRequest) SetSchedule(v ExternalPaymentScheduleRequest)`

SetSchedule sets Schedule field to given value.

### HasSchedule

`func (o *PaymentInitiationPaymentCreateRequest) HasSchedule() bool`

HasSchedule returns a boolean if a field has been set.

### GetOptions

`func (o *PaymentInitiationPaymentCreateRequest) GetOptions() ExternalPaymentOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *PaymentInitiationPaymentCreateRequest) GetOptionsOk() (*ExternalPaymentOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *PaymentInitiationPaymentCreateRequest) SetOptions(v ExternalPaymentOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *PaymentInitiationPaymentCreateRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### SetOptionsNil

`func (o *PaymentInitiationPaymentCreateRequest) SetOptionsNil(b bool)`

 SetOptionsNil sets the value for Options to be an explicit nil

### UnsetOptions
`func (o *PaymentInitiationPaymentCreateRequest) UnsetOptions()`

UnsetOptions ensures that no value is present for Options, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


