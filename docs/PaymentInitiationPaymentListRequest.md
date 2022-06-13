# PaymentInitiationPaymentListRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | [optional] 
**Count** | Pointer to **NullableInt32** | The maximum number of payments to return. If &#x60;count&#x60; is not specified, a maximum of 10 payments will be returned, beginning with the most recent payment before the cursor (if specified). | [optional] [default to 10]
**Cursor** | Pointer to **NullableTime** | A string in RFC 3339 format (i.e. \&quot;2019-12-06T22:35:49Z\&quot;). Only payments created before the cursor will be returned. | [optional] 

## Methods

### NewPaymentInitiationPaymentListRequest

`func NewPaymentInitiationPaymentListRequest() *PaymentInitiationPaymentListRequest`

NewPaymentInitiationPaymentListRequest instantiates a new PaymentInitiationPaymentListRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiationPaymentListRequestWithDefaults

`func NewPaymentInitiationPaymentListRequestWithDefaults() *PaymentInitiationPaymentListRequest`

NewPaymentInitiationPaymentListRequestWithDefaults instantiates a new PaymentInitiationPaymentListRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *PaymentInitiationPaymentListRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *PaymentInitiationPaymentListRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *PaymentInitiationPaymentListRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *PaymentInitiationPaymentListRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *PaymentInitiationPaymentListRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *PaymentInitiationPaymentListRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *PaymentInitiationPaymentListRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *PaymentInitiationPaymentListRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetCount

`func (o *PaymentInitiationPaymentListRequest) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *PaymentInitiationPaymentListRequest) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *PaymentInitiationPaymentListRequest) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *PaymentInitiationPaymentListRequest) HasCount() bool`

HasCount returns a boolean if a field has been set.

### SetCountNil

`func (o *PaymentInitiationPaymentListRequest) SetCountNil(b bool)`

 SetCountNil sets the value for Count to be an explicit nil

### UnsetCount
`func (o *PaymentInitiationPaymentListRequest) UnsetCount()`

UnsetCount ensures that no value is present for Count, not even an explicit nil
### GetCursor

`func (o *PaymentInitiationPaymentListRequest) GetCursor() time.Time`

GetCursor returns the Cursor field if non-nil, zero value otherwise.

### GetCursorOk

`func (o *PaymentInitiationPaymentListRequest) GetCursorOk() (*time.Time, bool)`

GetCursorOk returns a tuple with the Cursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCursor

`func (o *PaymentInitiationPaymentListRequest) SetCursor(v time.Time)`

SetCursor sets Cursor field to given value.

### HasCursor

`func (o *PaymentInitiationPaymentListRequest) HasCursor() bool`

HasCursor returns a boolean if a field has been set.

### SetCursorNil

`func (o *PaymentInitiationPaymentListRequest) SetCursorNil(b bool)`

 SetCursorNil sets the value for Cursor to be an explicit nil

### UnsetCursor
`func (o *PaymentInitiationPaymentListRequest) UnsetCursor()`

UnsetCursor ensures that no value is present for Cursor, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


