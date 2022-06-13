# AddressNullable

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**AddressData**](AddressData.md) |  | 
**Primary** | Pointer to **bool** | When &#x60;true&#x60;, identifies the address as the primary address on an account. | [optional] 

## Methods

### NewAddressNullable

`func NewAddressNullable(data AddressData, ) *AddressNullable`

NewAddressNullable instantiates a new AddressNullable object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddressNullableWithDefaults

`func NewAddressNullableWithDefaults() *AddressNullable`

NewAddressNullableWithDefaults instantiates a new AddressNullable object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AddressNullable) GetData() AddressData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AddressNullable) GetDataOk() (*AddressData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AddressNullable) SetData(v AddressData)`

SetData sets Data field to given value.


### GetPrimary

`func (o *AddressNullable) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *AddressNullable) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *AddressNullable) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.

### HasPrimary

`func (o *AddressNullable) HasPrimary() bool`

HasPrimary returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


