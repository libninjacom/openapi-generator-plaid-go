# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | **NullableString** | The street address where the transaction occurred. | 
**City** | **NullableString** | The city where the transaction occurred. | 
**Region** | **NullableString** | The region or state where the transaction occurred. In API versions 2018-05-22 and earlier, this field is called &#x60;state&#x60;. | 
**PostalCode** | **NullableString** | The postal code where the transaction occurred. In API versions 2018-05-22 and earlier, this field is called &#x60;zip&#x60;. | 
**Country** | **NullableString** | The ISO 3166-1 alpha-2 country code where the transaction occurred. | 
**Lat** | **NullableFloat32** | The latitude where the transaction occurred. | 
**Lon** | **NullableFloat32** | The longitude where the transaction occurred. | 
**StoreNumber** | **NullableString** | The merchant defined store number where the transaction occurred. | 

## Methods

### NewLocation

`func NewLocation(address NullableString, city NullableString, region NullableString, postalCode NullableString, country NullableString, lat NullableFloat32, lon NullableFloat32, storeNumber NullableString, ) *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Location) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Location) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Location) SetAddress(v string)`

SetAddress sets Address field to given value.


### SetAddressNil

`func (o *Location) SetAddressNil(b bool)`

 SetAddressNil sets the value for Address to be an explicit nil

### UnsetAddress
`func (o *Location) UnsetAddress()`

UnsetAddress ensures that no value is present for Address, not even an explicit nil
### GetCity

`func (o *Location) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *Location) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *Location) SetCity(v string)`

SetCity sets City field to given value.


### SetCityNil

`func (o *Location) SetCityNil(b bool)`

 SetCityNil sets the value for City to be an explicit nil

### UnsetCity
`func (o *Location) UnsetCity()`

UnsetCity ensures that no value is present for City, not even an explicit nil
### GetRegion

`func (o *Location) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *Location) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *Location) SetRegion(v string)`

SetRegion sets Region field to given value.


### SetRegionNil

`func (o *Location) SetRegionNil(b bool)`

 SetRegionNil sets the value for Region to be an explicit nil

### UnsetRegion
`func (o *Location) UnsetRegion()`

UnsetRegion ensures that no value is present for Region, not even an explicit nil
### GetPostalCode

`func (o *Location) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *Location) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *Location) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### SetPostalCodeNil

`func (o *Location) SetPostalCodeNil(b bool)`

 SetPostalCodeNil sets the value for PostalCode to be an explicit nil

### UnsetPostalCode
`func (o *Location) UnsetPostalCode()`

UnsetPostalCode ensures that no value is present for PostalCode, not even an explicit nil
### GetCountry

`func (o *Location) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Location) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Location) SetCountry(v string)`

SetCountry sets Country field to given value.


### SetCountryNil

`func (o *Location) SetCountryNil(b bool)`

 SetCountryNil sets the value for Country to be an explicit nil

### UnsetCountry
`func (o *Location) UnsetCountry()`

UnsetCountry ensures that no value is present for Country, not even an explicit nil
### GetLat

`func (o *Location) GetLat() float32`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *Location) GetLatOk() (*float32, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *Location) SetLat(v float32)`

SetLat sets Lat field to given value.


### SetLatNil

`func (o *Location) SetLatNil(b bool)`

 SetLatNil sets the value for Lat to be an explicit nil

### UnsetLat
`func (o *Location) UnsetLat()`

UnsetLat ensures that no value is present for Lat, not even an explicit nil
### GetLon

`func (o *Location) GetLon() float32`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *Location) GetLonOk() (*float32, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLon

`func (o *Location) SetLon(v float32)`

SetLon sets Lon field to given value.


### SetLonNil

`func (o *Location) SetLonNil(b bool)`

 SetLonNil sets the value for Lon to be an explicit nil

### UnsetLon
`func (o *Location) UnsetLon()`

UnsetLon ensures that no value is present for Lon, not even an explicit nil
### GetStoreNumber

`func (o *Location) GetStoreNumber() string`

GetStoreNumber returns the StoreNumber field if non-nil, zero value otherwise.

### GetStoreNumberOk

`func (o *Location) GetStoreNumberOk() (*string, bool)`

GetStoreNumberOk returns a tuple with the StoreNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStoreNumber

`func (o *Location) SetStoreNumber(v string)`

SetStoreNumber sets StoreNumber field to given value.


### SetStoreNumberNil

`func (o *Location) SetStoreNumberNil(b bool)`

 SetStoreNumberNil sets the value for StoreNumber to be an explicit nil

### UnsetStoreNumber
`func (o *Location) UnsetStoreNumber()`

UnsetStoreNumber ensures that no value is present for StoreNumber, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


