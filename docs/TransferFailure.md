# TransferFailure

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AchReturnCode** | Pointer to **NullableString** | The ACH return code, e.g. &#x60;R01&#x60;.  A return code will be provided if and only if the transfer status is &#x60;reversed&#x60;. For a full listing of ACH return codes, see [Transfer errors](https://plaid.com/docs/errors/transfer/#ach-return-codes). | [optional] 
**Description** | Pointer to **string** | A human-readable description of the reason for the failure or reversal. | [optional] 

## Methods

### NewTransferFailure

`func NewTransferFailure() *TransferFailure`

NewTransferFailure instantiates a new TransferFailure object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferFailureWithDefaults

`func NewTransferFailureWithDefaults() *TransferFailure`

NewTransferFailureWithDefaults instantiates a new TransferFailure object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAchReturnCode

`func (o *TransferFailure) GetAchReturnCode() string`

GetAchReturnCode returns the AchReturnCode field if non-nil, zero value otherwise.

### GetAchReturnCodeOk

`func (o *TransferFailure) GetAchReturnCodeOk() (*string, bool)`

GetAchReturnCodeOk returns a tuple with the AchReturnCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAchReturnCode

`func (o *TransferFailure) SetAchReturnCode(v string)`

SetAchReturnCode sets AchReturnCode field to given value.

### HasAchReturnCode

`func (o *TransferFailure) HasAchReturnCode() bool`

HasAchReturnCode returns a boolean if a field has been set.

### SetAchReturnCodeNil

`func (o *TransferFailure) SetAchReturnCodeNil(b bool)`

 SetAchReturnCodeNil sets the value for AchReturnCode to be an explicit nil

### UnsetAchReturnCode
`func (o *TransferFailure) UnsetAchReturnCode()`

UnsetAchReturnCode ensures that no value is present for AchReturnCode, not even an explicit nil
### GetDescription

`func (o *TransferFailure) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *TransferFailure) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *TransferFailure) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *TransferFailure) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


