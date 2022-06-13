# TransferUserAddressInRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Street** | Pointer to **string** | The street number and name (i.e., \&quot;100 Market St.\&quot;). | [optional] 
**City** | Pointer to **string** | Ex. \&quot;San Francisco\&quot; | [optional] 
**Region** | Pointer to **string** | The state or province (e.g., \&quot;California\&quot;). | [optional] 
**PostalCode** | Pointer to **string** | The postal code (e.g., \&quot;94103\&quot;). | [optional] 
**Country** | Pointer to **string** | A two-letter country code (e.g., \&quot;US\&quot;). | [optional] 

## Methods

### NewTransferUserAddressInRequest

`func NewTransferUserAddressInRequest() *TransferUserAddressInRequest`

NewTransferUserAddressInRequest instantiates a new TransferUserAddressInRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferUserAddressInRequestWithDefaults

`func NewTransferUserAddressInRequestWithDefaults() *TransferUserAddressInRequest`

NewTransferUserAddressInRequestWithDefaults instantiates a new TransferUserAddressInRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreet

`func (o *TransferUserAddressInRequest) GetStreet() string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *TransferUserAddressInRequest) GetStreetOk() (*string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *TransferUserAddressInRequest) SetStreet(v string)`

SetStreet sets Street field to given value.

### HasStreet

`func (o *TransferUserAddressInRequest) HasStreet() bool`

HasStreet returns a boolean if a field has been set.

### GetCity

`func (o *TransferUserAddressInRequest) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *TransferUserAddressInRequest) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *TransferUserAddressInRequest) SetCity(v string)`

SetCity sets City field to given value.

### HasCity

`func (o *TransferUserAddressInRequest) HasCity() bool`

HasCity returns a boolean if a field has been set.

### GetRegion

`func (o *TransferUserAddressInRequest) GetRegion() string`

GetRegion returns the Region field if non-nil, zero value otherwise.

### GetRegionOk

`func (o *TransferUserAddressInRequest) GetRegionOk() (*string, bool)`

GetRegionOk returns a tuple with the Region field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegion

`func (o *TransferUserAddressInRequest) SetRegion(v string)`

SetRegion sets Region field to given value.

### HasRegion

`func (o *TransferUserAddressInRequest) HasRegion() bool`

HasRegion returns a boolean if a field has been set.

### GetPostalCode

`func (o *TransferUserAddressInRequest) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *TransferUserAddressInRequest) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *TransferUserAddressInRequest) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.

### HasPostalCode

`func (o *TransferUserAddressInRequest) HasPostalCode() bool`

HasPostalCode returns a boolean if a field has been set.

### GetCountry

`func (o *TransferUserAddressInRequest) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *TransferUserAddressInRequest) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *TransferUserAddressInRequest) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *TransferUserAddressInRequest) HasCountry() bool`

HasCountry returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


