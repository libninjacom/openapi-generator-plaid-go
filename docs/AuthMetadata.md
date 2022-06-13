# AuthMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SupportedMethods** | [**NullableAuthSupportedMethods**](AuthSupportedMethods.md) |  | 

## Methods

### NewAuthMetadata

`func NewAuthMetadata(supportedMethods NullableAuthSupportedMethods, ) *AuthMetadata`

NewAuthMetadata instantiates a new AuthMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthMetadataWithDefaults

`func NewAuthMetadataWithDefaults() *AuthMetadata`

NewAuthMetadataWithDefaults instantiates a new AuthMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSupportedMethods

`func (o *AuthMetadata) GetSupportedMethods() AuthSupportedMethods`

GetSupportedMethods returns the SupportedMethods field if non-nil, zero value otherwise.

### GetSupportedMethodsOk

`func (o *AuthMetadata) GetSupportedMethodsOk() (*AuthSupportedMethods, bool)`

GetSupportedMethodsOk returns a tuple with the SupportedMethods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedMethods

`func (o *AuthMetadata) SetSupportedMethods(v AuthSupportedMethods)`

SetSupportedMethods sets SupportedMethods field to given value.


### SetSupportedMethodsNil

`func (o *AuthMetadata) SetSupportedMethodsNil(b bool)`

 SetSupportedMethodsNil sets the value for SupportedMethods to be an explicit nil

### UnsetSupportedMethods
`func (o *AuthMetadata) UnsetSupportedMethods()`

UnsetSupportedMethods ensures that no value is present for SupportedMethods, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


