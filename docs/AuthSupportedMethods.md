# AuthSupportedMethods

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**InstantAuth** | **bool** | Indicates if instant auth is supported. | 
**InstantMatch** | **bool** | Indicates if instant match is supported. | 
**AutomatedMicroDeposits** | **bool** | Indicates if automated microdeposits are supported. | 

## Methods

### NewAuthSupportedMethods

`func NewAuthSupportedMethods(instantAuth bool, instantMatch bool, automatedMicroDeposits bool, ) *AuthSupportedMethods`

NewAuthSupportedMethods instantiates a new AuthSupportedMethods object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthSupportedMethodsWithDefaults

`func NewAuthSupportedMethodsWithDefaults() *AuthSupportedMethods`

NewAuthSupportedMethodsWithDefaults instantiates a new AuthSupportedMethods object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInstantAuth

`func (o *AuthSupportedMethods) GetInstantAuth() bool`

GetInstantAuth returns the InstantAuth field if non-nil, zero value otherwise.

### GetInstantAuthOk

`func (o *AuthSupportedMethods) GetInstantAuthOk() (*bool, bool)`

GetInstantAuthOk returns a tuple with the InstantAuth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantAuth

`func (o *AuthSupportedMethods) SetInstantAuth(v bool)`

SetInstantAuth sets InstantAuth field to given value.


### GetInstantMatch

`func (o *AuthSupportedMethods) GetInstantMatch() bool`

GetInstantMatch returns the InstantMatch field if non-nil, zero value otherwise.

### GetInstantMatchOk

`func (o *AuthSupportedMethods) GetInstantMatchOk() (*bool, bool)`

GetInstantMatchOk returns a tuple with the InstantMatch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInstantMatch

`func (o *AuthSupportedMethods) SetInstantMatch(v bool)`

SetInstantMatch sets InstantMatch field to given value.


### GetAutomatedMicroDeposits

`func (o *AuthSupportedMethods) GetAutomatedMicroDeposits() bool`

GetAutomatedMicroDeposits returns the AutomatedMicroDeposits field if non-nil, zero value otherwise.

### GetAutomatedMicroDepositsOk

`func (o *AuthSupportedMethods) GetAutomatedMicroDepositsOk() (*bool, bool)`

GetAutomatedMicroDepositsOk returns a tuple with the AutomatedMicroDeposits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutomatedMicroDeposits

`func (o *AuthSupportedMethods) SetAutomatedMicroDeposits(v bool)`

SetAutomatedMicroDeposits sets AutomatedMicroDeposits field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


