# AddressDataNullable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | **string** | The full city name | 
**Region** | **NullableString** | The region or state. In API versions 2018-05-22 and earlier, this field is called &#x60;state&#x60;. Example: &#x60;\&quot;NC\&quot;&#x60; | 
**Street** | **string** | The full street address Example: &#x60;\&quot;564 Main Street, APT 15\&quot;&#x60; | 
**PostalCode** | **NullableString** | The postal code. In API versions 2018-05-22 and earlier, this field is called &#x60;zip&#x60;. | 
**Country** | **NullableString** | The ISO 3166-1 alpha-2 country code | 

## Methods

### NewAddressDataNullable

`func NewAddressDataNullable(city string, region NullableString, street string, postalCode NullableString, country NullableString, ) *AddressDataNullable`

NewAddressDataNullable instantiates a new AddressDataNullable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressDataNullableWithDefaults

`func NewAddressDataNullableWithDefaults() *AddressDataNullable`

NewAddressDataNullableWithDefaults instantiates a new AddressDataNullable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *AddressDataNullable) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *AddressDataNullable) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *AddressDataNullable) SetCity(v string)`

SetCity sets City field to given value.


### GetRegion

`func (o *AddressDataNullable) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *AddressDataNullable) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *AddressDataNullable) SetRegion(v string)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *AddressDataNullable) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *AddressDataNullable) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetStreet

`func (o *AddressDataNullable) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *AddressDataNullable) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *AddressDataNullable) SetStreet(v string)`

SetStreet sets Street field to given value.


### GetPostalCode

`func (o *AddressDataNullable) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *AddressDataNullable) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *AddressDataNullable) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### SetPostalCodeNil

`func (o *AddressDataNullable) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *AddressDataNullable) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountry

`func (o *AddressDataNullable) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AddressDataNullable) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AddressDataNullable) SetCountry(v string)`

SetCountry sets Country field to given value.


### SetCountryNil

`func (o *AddressDataNullable) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *AddressDataNullable) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


