# PhoneNumber

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **string** | The phone number. | 
**Primary** | **bool** | When &#x60;true&#x60;, identifies the phone number as the primary number on an account. | 
**Type** | **string** | The type of phone number. | 

## Methods

### NewPhoneNumber

`func NewPhoneNumber(data string, primary bool, type_ string, ) *PhoneNumber`

NewPhoneNumber instantiates a new PhoneNumber object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPhoneNumberWithDefaults

`func NewPhoneNumberWithDefaults() *PhoneNumber`

NewPhoneNumberWithDefaults instantiates a new PhoneNumber object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *PhoneNumber) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PhoneNumber) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PhoneNumber) SetData(v string)`

SetData sets Data field to given value.


### GetPrimary

`func (o *PhoneNumber) GetPrimary() bool`

GetPrimary returns the Primary field if non-nil, zero value otherwise.

### GetPrimaryOk

`func (o *PhoneNumber) GetPrimaryOk() (*bool, bool)`

GetPrimaryOk returns a tuple with the Primary field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPrimary

`func (o *PhoneNumber) SetPrimary(v bool)`

SetPrimary sets Primary field to given value.


### GetType

`func (o *PhoneNumber) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PhoneNumber) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PhoneNumber) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


