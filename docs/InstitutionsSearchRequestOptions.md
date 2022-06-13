# InstitutionsSearchRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Oauth** | Pointer to **NullableBool** | Limit results to institutions with or without OAuth login flows. | [optional] 
**IncludeOptionalMetadata** | Pointer to **bool** | When true, return the institution&#39;s homepage URL, logo and primary brand color. | [optional] 
**IncludeAuthMetadata** | Pointer to **NullableBool** | When &#x60;true&#x60;, returns metadata related to the Auth product indicating which auth methods are supported. | [optional] [default to false]
**IncludePaymentInitiationMetadata** | Pointer to **NullableBool** | When &#x60;true&#x60;, returns metadata related to the Payment Initiation product indicating which payment configurations are supported. | [optional] [default to false]
**PaymentInitiation** | Pointer to [**NullableInstitutionsSearchPaymentInitiationOptions**](InstitutionsSearchPaymentInitiationOptions.md) |  | [optional] 

## Methods

### NewInstitutionsSearchRequestOptions

`func NewInstitutionsSearchRequestOptions() *InstitutionsSearchRequestOptions`

NewInstitutionsSearchRequestOptions instantiates a new InstitutionsSearchRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionsSearchRequestOptionsWithDefaults

`func NewInstitutionsSearchRequestOptionsWithDefaults() *InstitutionsSearchRequestOptions`

NewInstitutionsSearchRequestOptionsWithDefaults instantiates a new InstitutionsSearchRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOauth

`func (o *InstitutionsSearchRequestOptions) GetOauth() bool`

GetOauth returns the Oauth field if non-nil, zero value otherwise.

### GetOauthOk

`func (o *InstitutionsSearchRequestOptions) GetOauthOk() (*bool, bool)`

GetOauthOk returns a tuple with the Oauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth

`func (o *InstitutionsSearchRequestOptions) SetOauth(v bool)`

SetOauth sets Oauth field to given value.

### HasOauth

`func (o *InstitutionsSearchRequestOptions) HasOauth() bool`

HasOauth returns a boolean if a field has been set.

### SetOauthNil

`func (o *InstitutionsSearchRequestOptions) SetOauthNil(b bool)`

 SetOauthNil sets the value for Oauth to be an explicit nil

### UnsetOauth
`func (o *InstitutionsSearchRequestOptions) UnsetOauth()`

UnsetOauth ensures that no value is present for Oauth, not even an explicit nil
### GetIncludeOptionalMetadata

`func (o *InstitutionsSearchRequestOptions) GetIncludeOptionalMetadata() bool`

GetIncludeOptionalMetadata returns the IncludeOptionalMetadata field if non-nil, zero value otherwise.

### GetIncludeOptionalMetadataOk

`func (o *InstitutionsSearchRequestOptions) GetIncludeOptionalMetadataOk() (*bool, bool)`

GetIncludeOptionalMetadataOk returns a tuple with the IncludeOptionalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOptionalMetadata

`func (o *InstitutionsSearchRequestOptions) SetIncludeOptionalMetadata(v bool)`

SetIncludeOptionalMetadata sets IncludeOptionalMetadata field to given value.

### HasIncludeOptionalMetadata

`func (o *InstitutionsSearchRequestOptions) HasIncludeOptionalMetadata() bool`

HasIncludeOptionalMetadata returns a boolean if a field has been set.

### GetIncludeAuthMetadata

`func (o *InstitutionsSearchRequestOptions) GetIncludeAuthMetadata() bool`

GetIncludeAuthMetadata returns the IncludeAuthMetadata field if non-nil, zero value otherwise.

### GetIncludeAuthMetadataOk

`func (o *InstitutionsSearchRequestOptions) GetIncludeAuthMetadataOk() (*bool, bool)`

GetIncludeAuthMetadataOk returns a tuple with the IncludeAuthMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAuthMetadata

`func (o *InstitutionsSearchRequestOptions) SetIncludeAuthMetadata(v bool)`

SetIncludeAuthMetadata sets IncludeAuthMetadata field to given value.

### HasIncludeAuthMetadata

`func (o *InstitutionsSearchRequestOptions) HasIncludeAuthMetadata() bool`

HasIncludeAuthMetadata returns a boolean if a field has been set.

### SetIncludeAuthMetadataNil

`func (o *InstitutionsSearchRequestOptions) SetIncludeAuthMetadataNil(b bool)`

 SetIncludeAuthMetadataNil sets the value for IncludeAuthMetadata to be an explicit nil

### UnsetIncludeAuthMetadata
`func (o *InstitutionsSearchRequestOptions) UnsetIncludeAuthMetadata()`

UnsetIncludeAuthMetadata ensures that no value is present for IncludeAuthMetadata, not even an explicit nil
### GetIncludePaymentInitiationMetadata

`func (o *InstitutionsSearchRequestOptions) GetIncludePaymentInitiationMetadata() bool`

GetIncludePaymentInitiationMetadata returns the IncludePaymentInitiationMetadata field if non-nil, zero value otherwise.

### GetIncludePaymentInitiationMetadataOk

`func (o *InstitutionsSearchRequestOptions) GetIncludePaymentInitiationMetadataOk() (*bool, bool)`

GetIncludePaymentInitiationMetadataOk returns a tuple with the IncludePaymentInitiationMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePaymentInitiationMetadata

`func (o *InstitutionsSearchRequestOptions) SetIncludePaymentInitiationMetadata(v bool)`

SetIncludePaymentInitiationMetadata sets IncludePaymentInitiationMetadata field to given value.

### HasIncludePaymentInitiationMetadata

`func (o *InstitutionsSearchRequestOptions) HasIncludePaymentInitiationMetadata() bool`

HasIncludePaymentInitiationMetadata returns a boolean if a field has been set.

### SetIncludePaymentInitiationMetadataNil

`func (o *InstitutionsSearchRequestOptions) SetIncludePaymentInitiationMetadataNil(b bool)`

 SetIncludePaymentInitiationMetadataNil sets the value for IncludePaymentInitiationMetadata to be an explicit nil

### UnsetIncludePaymentInitiationMetadata
`func (o *InstitutionsSearchRequestOptions) UnsetIncludePaymentInitiationMetadata()`

UnsetIncludePaymentInitiationMetadata ensures that no value is present for IncludePaymentInitiationMetadata, not even an explicit nil
### GetPaymentInitiation

`func (o *InstitutionsSearchRequestOptions) GetPaymentInitiation() InstitutionsSearchPaymentInitiationOptions`

GetPaymentInitiation returns the PaymentInitiation field if non-nil, zero value otherwise.

### GetPaymentInitiationOk

`func (o *InstitutionsSearchRequestOptions) GetPaymentInitiationOk() (*InstitutionsSearchPaymentInitiationOptions, bool)`

GetPaymentInitiationOk returns a tuple with the PaymentInitiation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaymentInitiation

`func (o *InstitutionsSearchRequestOptions) SetPaymentInitiation(v InstitutionsSearchPaymentInitiationOptions)`

SetPaymentInitiation sets PaymentInitiation field to given value.

### HasPaymentInitiation

`func (o *InstitutionsSearchRequestOptions) HasPaymentInitiation() bool`

HasPaymentInitiation returns a boolean if a field has been set.

### SetPaymentInitiationNil

`func (o *InstitutionsSearchRequestOptions) SetPaymentInitiationNil(b bool)`

 SetPaymentInitiationNil sets the value for PaymentInitiation to be an explicit nil

### UnsetPaymentInitiation
`func (o *InstitutionsSearchRequestOptions) UnsetPaymentInitiation()`

UnsetPaymentInitiation ensures that no value is present for PaymentInitiation, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


