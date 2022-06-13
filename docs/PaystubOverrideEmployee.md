# PaystubOverrideEmployee

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the employee. | [optional] 
**Address** | Pointer to [**PaystubOverrideEmployeeAddress**](PaystubOverrideEmployeeAddress.md) |  | [optional] 

## Methods

### NewPaystubOverrideEmployee

`func NewPaystubOverrideEmployee() *PaystubOverrideEmployee`

NewPaystubOverrideEmployee instantiates a new PaystubOverrideEmployee object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPaystubOverrideEmployeeWithDefaults

`func NewPaystubOverrideEmployeeWithDefaults() *PaystubOverrideEmployee`

NewPaystubOverrideEmployeeWithDefaults instantiates a new PaystubOverrideEmployee object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PaystubOverrideEmployee) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PaystubOverrideEmployee) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PaystubOverrideEmployee) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *PaystubOverrideEmployee) HasName() bool`

HasName returns a boolean if a field has been set.

### GetAddress

`func (o *PaystubOverrideEmployee) GetAddress() PaystubOverrideEmployeeAddress`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *PaystubOverrideEmployee) GetAddressOk() (*PaystubOverrideEmployeeAddress, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *PaystubOverrideEmployee) SetAddress(v PaystubOverrideEmployeeAddress)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *PaystubOverrideEmployee) HasAddress() bool`

HasAddress returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


