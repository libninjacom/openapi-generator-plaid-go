# SandboxOauthSelectAccountsRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**OauthStateId** | **string** |  | 
**Accounts** | **[]string** |  | 

## Methods

### NewSandboxOauthSelectAccountsRequest

`func NewSandboxOauthSelectAccountsRequest(oauthStateId string, accounts []string, ) *SandboxOauthSelectAccountsRequest`

NewSandboxOauthSelectAccountsRequest instantiates a new SandboxOauthSelectAccountsRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSandboxOauthSelectAccountsRequestWithDefaults

`func NewSandboxOauthSelectAccountsRequestWithDefaults() *SandboxOauthSelectAccountsRequest`

NewSandboxOauthSelectAccountsRequestWithDefaults instantiates a new SandboxOauthSelectAccountsRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOauthStateId

`func (o *SandboxOauthSelectAccountsRequest) GetOauthStateId() string`

GetOauthStateId returns the OauthStateId field if non-nil, zero value otherwise.

### GetOauthStateIdOk

`func (o *SandboxOauthSelectAccountsRequest) GetOauthStateIdOk() (*string, bool)`

GetOauthStateIdOk returns a tuple with the OauthStateId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOauthStateId

`func (o *SandboxOauthSelectAccountsRequest) SetOauthStateId(v string)`

SetOauthStateId sets OauthStateId field to given value.


### GetAccounts

`func (o *SandboxOauthSelectAccountsRequest) GetAccounts() []string`

GetAccounts returns the Accounts field if non-nil, zero value otherwise.

### GetAccountsOk

`func (o *SandboxOauthSelectAccountsRequest) GetAccountsOk() (*[]string, bool)`

GetAccountsOk returns a tuple with the Accounts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccounts

`func (o *SandboxOauthSelectAccountsRequest) SetAccounts(v []string)`

SetAccounts sets Accounts field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


