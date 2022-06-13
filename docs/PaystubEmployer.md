# PaystubEmployer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to [**PaystubAddress**](PaystubAddress.md) |  | [optional] 
**Name** | **NullableString** | The name of the employer on the paystub. | 

## Methods

### NewPaystubEmployer

`func NewPaystubEmployer(name NullableString, ) *PaystubEmployer`

NewPaystubEmployer instantiates a new PaystubEmployer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaystubEmployerWithDefaults

`func NewPaystubEmployerWithDefaults() *PaystubEmployer`

NewPaystubEmployerWithDefaults instantiates a new PaystubEmployer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *PaystubEmployer) GetAddress() PaystubAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PaystubEmployer) GetAddressOk() (*PaystubAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PaystubEmployer) SetAddress(v PaystubAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PaystubEmployer) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetName

`func (o *PaystubEmployer) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaystubEmployer) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaystubEmployer) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *PaystubEmployer) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *PaystubEmployer) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


