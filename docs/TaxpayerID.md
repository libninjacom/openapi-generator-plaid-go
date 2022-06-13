# TaxpayerID

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IdType** | Pointer to **NullableString** | Type of ID, e.g. &#39;SSN&#39; | [optional] 
**IdMask** | Pointer to **NullableString** | ID mask; i.e. last 4 digits of the taxpayer ID | [optional] 
**Last4Digits** | Pointer to **NullableString** | Last 4 digits of unique number of ID. | [optional] 

## Methods

### NewTaxpayerID

`func NewTaxpayerID() *TaxpayerID`

NewTaxpayerID instantiates a new TaxpayerID object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxpayerIDWithDefaults

`func NewTaxpayerIDWithDefaults() *TaxpayerID`

NewTaxpayerIDWithDefaults instantiates a new TaxpayerID object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIdType

`func (o *TaxpayerID) GetIdType() string`

GetIdType returns the IdType field if non-nil, zero value otherwise.

### GetIdTypeOk

`func (o *TaxpayerID) GetIdTypeOk() (*string, bool)`

GetIdTypeOk returns a tuple with the IdType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdType

`func (o *TaxpayerID) SetIdType(v string)`

SetIdType sets IdType field to given value.

### HasIdType

`func (o *TaxpayerID) HasIdType() bool`

HasIdType returns a boolean if a field has been set.

### SetIdTypeNil

`func (o *TaxpayerID) SetIdTypeNil(b bool)`

 SetIdTypeNil sets the value for IdType to be an explicit nil

### UnsetIdType
`func (o *TaxpayerID) UnsetIdType()`

UnsetIdType ensures that no value is present for IdType, not even an explicit nil
### GetIdMask

`func (o *TaxpayerID) GetIdMask() string`

GetIdMask returns the IdMask field if non-nil, zero value otherwise.

### GetIdMaskOk

`func (o *TaxpayerID) GetIdMaskOk() (*string, bool)`

GetIdMaskOk returns a tuple with the IdMask field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdMask

`func (o *TaxpayerID) SetIdMask(v string)`

SetIdMask sets IdMask field to given value.

### HasIdMask

`func (o *TaxpayerID) HasIdMask() bool`

HasIdMask returns a boolean if a field has been set.

### SetIdMaskNil

`func (o *TaxpayerID) SetIdMaskNil(b bool)`

 SetIdMaskNil sets the value for IdMask to be an explicit nil

### UnsetIdMask
`func (o *TaxpayerID) UnsetIdMask()`

UnsetIdMask ensures that no value is present for IdMask, not even an explicit nil
### GetLast4Digits

`func (o *TaxpayerID) GetLast4Digits() string`

GetLast4Digits returns the Last4Digits field if non-nil, zero value otherwise.

### GetLast4DigitsOk

`func (o *TaxpayerID) GetLast4DigitsOk() (*string, bool)`

GetLast4DigitsOk returns a tuple with the Last4Digits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLast4Digits

`func (o *TaxpayerID) SetLast4Digits(v string)`

SetLast4Digits sets Last4Digits field to given value.

### HasLast4Digits

`func (o *TaxpayerID) HasLast4Digits() bool`

HasLast4Digits returns a boolean if a field has been set.

### SetLast4DigitsNil

`func (o *TaxpayerID) SetLast4DigitsNil(b bool)`

 SetLast4DigitsNil sets the value for Last4Digits to be an explicit nil

### UnsetLast4Digits
`func (o *TaxpayerID) UnsetLast4Digits()`

UnsetLast4Digits ensures that no value is present for Last4Digits, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


