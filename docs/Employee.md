# Employee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | [**PaystubAddress**](PaystubAddress.md) |  | 
**Name** | **NullableString** | The name of the employee. | 
**MaritalStatus** | Pointer to **NullableString** | Marital status of the employee - either &#x60;single&#x60; or &#x60;married&#x60;. | [optional] 
**TaxpayerId** | Pointer to [**TaxpayerID**](TaxpayerID.md) |  | [optional] 

## Methods

### NewEmployee

`func NewEmployee(address PaystubAddress, name NullableString, ) *Employee`

NewEmployee instantiates a new Employee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmployeeWithDefaults

`func NewEmployeeWithDefaults() *Employee`

NewEmployeeWithDefaults instantiates a new Employee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *Employee) GetAddress() PaystubAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *Employee) GetAddressOk() (*PaystubAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *Employee) SetAddress(v PaystubAddress)`

SetAddress sets Address field to given value.


### GetName

`func (o *Employee) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Employee) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Employee) SetName(v string)`

SetName sets Name field to given value.


### SetNameNil

`func (o *Employee) SetNameNil(b bool)`

 SetNameNil sets the value for Name to be an explicit nil

### UnsetName
`func (o *Employee) UnsetName()`

UnsetName ensures that no value is present for Name, not even an explicit nil
### GetMaritalStatus

`func (o *Employee) GetMaritalStatus() string`

GetMaritalStatus returns the MaritalStatus field if non-nil, zero value otherwise.

### GetMaritalStatusOk

`func (o *Employee) GetMaritalStatusOk() (*string, bool)`

GetMaritalStatusOk returns a tuple with the MaritalStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaritalStatus

`func (o *Employee) SetMaritalStatus(v string)`

SetMaritalStatus sets MaritalStatus field to given value.

### HasMaritalStatus

`func (o *Employee) HasMaritalStatus() bool`

HasMaritalStatus returns a boolean if a field has been set.

### SetMaritalStatusNil

`func (o *Employee) SetMaritalStatusNil(b bool)`

 SetMaritalStatusNil sets the value for MaritalStatus to be an explicit nil

### UnsetMaritalStatus
`func (o *Employee) UnsetMaritalStatus()`

UnsetMaritalStatus ensures that no value is present for MaritalStatus, not even an explicit nil
### GetTaxpayerId

`func (o *Employee) GetTaxpayerId() TaxpayerID`

GetTaxpayerId returns the TaxpayerId field if non-nil, zero value otherwise.

### GetTaxpayerIdOk

`func (o *Employee) GetTaxpayerIdOk() (*TaxpayerID, bool)`

GetTaxpayerIdOk returns a tuple with the TaxpayerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxpayerId

`func (o *Employee) SetTaxpayerId(v TaxpayerID)`

SetTaxpayerId sets TaxpayerId field to given value.

### HasTaxpayerId

`func (o *Employee) HasTaxpayerId() bool`

HasTaxpayerId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


