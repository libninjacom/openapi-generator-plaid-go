# PaymentInitiationAddress

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Street** | **[]string** | An array of length 1-2 representing the street address where the recipient is located. Maximum of 70 characters. | 
**City** | **string** | The city where the recipient is located. Maximum of 35 characters. | 
**PostalCode** | **string** | The postal code where the recipient is located. Maximum of 16 characters. | 
**Country** | **string** | The ISO 3166-1 alpha-2 country code where the recipient is located. | 

## Methods

### NewPaymentInitiationAddress

`func NewPaymentInitiationAddress(street []string, city string, postalCode string, country string, ) *PaymentInitiationAddress`

NewPaymentInitiationAddress instantiates a new PaymentInitiationAddress object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaymentInitiationAddressWithDefaults

`func NewPaymentInitiationAddressWithDefaults() *PaymentInitiationAddress`

NewPaymentInitiationAddressWithDefaults instantiates a new PaymentInitiationAddress object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreet

`func (o *PaymentInitiationAddress) GetStreet() []string`

GetStreet returns the Street field if non-nil, zero value otherwise.

### GetStreetOk

`func (o *PaymentInitiationAddress) GetStreetOk() (*[]string, bool)`

GetStreetOk returns a tuple with the Street field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreet

`func (o *PaymentInitiationAddress) SetStreet(v []string)`

SetStreet sets Street field to given value.


### GetCity

`func (o *PaymentInitiationAddress) GetCity() string`

GetCity returns the City field if non-nil, zero value otherwise.

### GetCityOk

`func (o *PaymentInitiationAddress) GetCityOk() (*string, bool)`

GetCityOk returns a tuple with the City field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCity

`func (o *PaymentInitiationAddress) SetCity(v string)`

SetCity sets City field to given value.


### GetPostalCode

`func (o *PaymentInitiationAddress) GetPostalCode() string`

GetPostalCode returns the PostalCode field if non-nil, zero value otherwise.

### GetPostalCodeOk

`func (o *PaymentInitiationAddress) GetPostalCodeOk() (*string, bool)`

GetPostalCodeOk returns a tuple with the PostalCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPostalCode

`func (o *PaymentInitiationAddress) SetPostalCode(v string)`

SetPostalCode sets PostalCode field to given value.


### GetCountry

`func (o *PaymentInitiationAddress) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *PaymentInitiationAddress) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *PaymentInitiationAddress) SetCountry(v string)`

SetCountry sets Country field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


