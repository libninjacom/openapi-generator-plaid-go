# InstitutionsGetRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Products** | Pointer to [**[]Products**](Products.md) | Filter the Institutions based on which products they support.  | [optional] 
**RoutingNumbers** | Pointer to **[]string** | Specify an array of routing numbers to filter institutions. The response will only return institutions that match all of the routing numbers in the array. Routing number records used for this matching are not comprehensive; failure to match a given routing number to an institution does not mean that the institution is unsupported by Plaid. | [optional] 
**Oauth** | Pointer to **NullableBool** | Limit results to institutions with or without OAuth login flows. | [optional] 
**IncludeOptionalMetadata** | Pointer to **bool** | When &#x60;true&#x60;, return the institution&#39;s homepage URL, logo and primary brand color.  Note that Plaid does not own any of the logos shared by the API, and that by accessing or using these logos, you agree that you are doing so at your own risk and will, if necessary, obtain all required permissions from the appropriate rights holders and adhere to any applicable usage guidelines. Plaid disclaims all express or implied warranties with respect to the logos. | [optional] 
**IncludeAuthMetadata** | Pointer to **bool** | When &#x60;true&#x60;, returns metadata related to the Auth product indicating which auth methods are supported. | [optional] [default to false]
**IncludePaymentInitiationMetadata** | Pointer to **bool** | When &#x60;true&#x60;, returns metadata related to the Payment Initiation product indicating which payment configurations are supported. | [optional] [default to false]

## Methods

### NewInstitutionsGetRequestOptions

`func NewInstitutionsGetRequestOptions() *InstitutionsGetRequestOptions`

NewInstitutionsGetRequestOptions instantiates a new InstitutionsGetRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInstitutionsGetRequestOptionsWithDefaults

`func NewInstitutionsGetRequestOptionsWithDefaults() *InstitutionsGetRequestOptions`

NewInstitutionsGetRequestOptionsWithDefaults instantiates a new InstitutionsGetRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProducts

`func (o *InstitutionsGetRequestOptions) GetProducts() []Products`

GetProducts returns the Products field if non-nil, zero value otherwise.

### GetProductsOk

`func (o *InstitutionsGetRequestOptions) GetProductsOk() (*[]Products, bool)`

GetProductsOk returns a tuple with the Products field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProducts

`func (o *InstitutionsGetRequestOptions) SetProducts(v []Products)`

SetProducts sets Products field to given value.

### HasProducts

`func (o *InstitutionsGetRequestOptions) HasProducts() bool`

HasProducts returns a boolean if a field has been set.

### SetProductsNil

`func (o *InstitutionsGetRequestOptions) SetProductsNil(b bool)`

 SetProductsNil sets the value for Products to be an explicit nil

### UnsetProducts
`func (o *InstitutionsGetRequestOptions) UnsetProducts()`

UnsetProducts ensures that no value is present for Products, not even an explicit nil
### GetRoutingNumbers

`func (o *InstitutionsGetRequestOptions) GetRoutingNumbers() []string`

GetRoutingNumbers returns the RoutingNumbers field if non-nil, zero value otherwise.

### GetRoutingNumbersOk

`func (o *InstitutionsGetRequestOptions) GetRoutingNumbersOk() (*[]string, bool)`

GetRoutingNumbersOk returns a tuple with the RoutingNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingNumbers

`func (o *InstitutionsGetRequestOptions) SetRoutingNumbers(v []string)`

SetRoutingNumbers sets RoutingNumbers field to given value.

### HasRoutingNumbers

`func (o *InstitutionsGetRequestOptions) HasRoutingNumbers() bool`

HasRoutingNumbers returns a boolean if a field has been set.

### SetRoutingNumbersNil

`func (o *InstitutionsGetRequestOptions) SetRoutingNumbersNil(b bool)`

 SetRoutingNumbersNil sets the value for RoutingNumbers to be an explicit nil

### UnsetRoutingNumbers
`func (o *InstitutionsGetRequestOptions) UnsetRoutingNumbers()`

UnsetRoutingNumbers ensures that no value is present for RoutingNumbers, not even an explicit nil
### GetOauth

`func (o *InstitutionsGetRequestOptions) GetOauth() bool`

GetOauth returns the Oauth field if non-nil, zero value otherwise.

### GetOauthOk

`func (o *InstitutionsGetRequestOptions) GetOauthOk() (*bool, bool)`

GetOauthOk returns a tuple with the Oauth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauth

`func (o *InstitutionsGetRequestOptions) SetOauth(v bool)`

SetOauth sets Oauth field to given value.

### HasOauth

`func (o *InstitutionsGetRequestOptions) HasOauth() bool`

HasOauth returns a boolean if a field has been set.

### SetOauthNil

`func (o *InstitutionsGetRequestOptions) SetOauthNil(b bool)`

 SetOauthNil sets the value for Oauth to be an explicit nil

### UnsetOauth
`func (o *InstitutionsGetRequestOptions) UnsetOauth()`

UnsetOauth ensures that no value is present for Oauth, not even an explicit nil
### GetIncludeOptionalMetadata

`func (o *InstitutionsGetRequestOptions) GetIncludeOptionalMetadata() bool`

GetIncludeOptionalMetadata returns the IncludeOptionalMetadata field if non-nil, zero value otherwise.

### GetIncludeOptionalMetadataOk

`func (o *InstitutionsGetRequestOptions) GetIncludeOptionalMetadataOk() (*bool, bool)`

GetIncludeOptionalMetadataOk returns a tuple with the IncludeOptionalMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeOptionalMetadata

`func (o *InstitutionsGetRequestOptions) SetIncludeOptionalMetadata(v bool)`

SetIncludeOptionalMetadata sets IncludeOptionalMetadata field to given value.

### HasIncludeOptionalMetadata

`func (o *InstitutionsGetRequestOptions) HasIncludeOptionalMetadata() bool`

HasIncludeOptionalMetadata returns a boolean if a field has been set.

### GetIncludeAuthMetadata

`func (o *InstitutionsGetRequestOptions) GetIncludeAuthMetadata() bool`

GetIncludeAuthMetadata returns the IncludeAuthMetadata field if non-nil, zero value otherwise.

### GetIncludeAuthMetadataOk

`func (o *InstitutionsGetRequestOptions) GetIncludeAuthMetadataOk() (*bool, bool)`

GetIncludeAuthMetadataOk returns a tuple with the IncludeAuthMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeAuthMetadata

`func (o *InstitutionsGetRequestOptions) SetIncludeAuthMetadata(v bool)`

SetIncludeAuthMetadata sets IncludeAuthMetadata field to given value.

### HasIncludeAuthMetadata

`func (o *InstitutionsGetRequestOptions) HasIncludeAuthMetadata() bool`

HasIncludeAuthMetadata returns a boolean if a field has been set.

### GetIncludePaymentInitiationMetadata

`func (o *InstitutionsGetRequestOptions) GetIncludePaymentInitiationMetadata() bool`

GetIncludePaymentInitiationMetadata returns the IncludePaymentInitiationMetadata field if non-nil, zero value otherwise.

### GetIncludePaymentInitiationMetadataOk

`func (o *InstitutionsGetRequestOptions) GetIncludePaymentInitiationMetadataOk() (*bool, bool)`

GetIncludePaymentInitiationMetadataOk returns a tuple with the IncludePaymentInitiationMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludePaymentInitiationMetadata

`func (o *InstitutionsGetRequestOptions) SetIncludePaymentInitiationMetadata(v bool)`

SetIncludePaymentInitiationMetadata sets IncludePaymentInitiationMetadata field to given value.

### HasIncludePaymentInitiationMetadata

`func (o *InstitutionsGetRequestOptions) HasIncludePaymentInitiationMetadata() bool`

HasIncludePaymentInitiationMetadata returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


