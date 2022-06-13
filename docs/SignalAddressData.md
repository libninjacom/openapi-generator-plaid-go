# SignalAddressData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**City** | Pointer to **string** | The full city name | [optional] 
**Region** | Pointer to **NullableString** | The region or state Example: &#x60;\&quot;NC\&quot;&#x60; | [optional] 
**Street** | Pointer to **string** | The full street address Example: &#x60;\&quot;564 Main Street, APT 15\&quot;&#x60; | [optional] 
**PostalCode** | Pointer to **NullableString** | The postal code | [optional] 
**Country** | Pointer to **NullableString** | The ISO 3166-1 alpha-2 country code | [optional] 

## Methods

### NewSignalAddressData

`func NewSignalAddressData() *SignalAddressData`

NewSignalAddressData instantiates a new SignalAddressData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSignalAddressDataWithDefaults

`func NewSignalAddressDataWithDefaults() *SignalAddressData`

NewSignalAddressDataWithDefaults instantiates a new SignalAddressData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCity

`func (o *SignalAddressData) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *SignalAddressData) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *SignalAddressData) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *SignalAddressData) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetRegion

`func (o *SignalAddressData) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *SignalAddressData) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *SignalAddressData) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *SignalAddressData) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### SetRegionNil

`func (o *SignalAddressData) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *SignalAddressData) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetStreet

`func (o *SignalAddressData) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *SignalAddressData) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *SignalAddressData) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *SignalAddressData) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetPostalCode

`func (o *SignalAddressData) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *SignalAddressData) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *SignalAddressData) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *SignalAddressData) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### SetPostalCodeNil

`func (o *SignalAddressData) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *SignalAddressData) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountry

`func (o *SignalAddressData) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SignalAddressData) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SignalAddressData) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *SignalAddressData) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### SetCountryNil

`func (o *SignalAddressData) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *SignalAddressData) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


