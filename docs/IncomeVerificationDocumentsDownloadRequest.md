# IncomeVerificationDocumentsDownloadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | Your Plaid API &#x60;client_id&#x60;. The &#x60;client_id&#x60; is required and may be provided either in the &#x60;PLAID-CLIENT-ID&#x60; header or as part of a request body. | [optional] 
**Secret** | Pointer to **string** | Your Plaid API &#x60;secret&#x60;. The &#x60;secret&#x60; is required and may be provided either in the &#x60;PLAID-SECRET&#x60; header or as part of a request body. | [optional] 
**IncomeVerificationId** | Pointer to **NullableString** | The ID of the verification. | [optional] 
**AccessToken** | Pointer to **NullableString** | The access token associated with the Item data is being requested for. | [optional] 
**DocumentId** | Pointer to **NullableString** | The document ID to download. If passed, a single document will be returned in the resulting zip file, rather than all document | [optional] 

## Methods

### NewIncomeVerificationDocumentsDownloadRequest

`func NewIncomeVerificationDocumentsDownloadRequest() *IncomeVerificationDocumentsDownloadRequest`

NewIncomeVerificationDocumentsDownloadRequest instantiates a new IncomeVerificationDocumentsDownloadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeVerificationDocumentsDownloadRequestWithDefaults

`func NewIncomeVerificationDocumentsDownloadRequestWithDefaults() *IncomeVerificationDocumentsDownloadRequest`

NewIncomeVerificationDocumentsDownloadRequestWithDefaults instantiates a new IncomeVerificationDocumentsDownloadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *IncomeVerificationDocumentsDownloadRequest) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *IncomeVerificationDocumentsDownloadRequest) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *IncomeVerificationDocumentsDownloadRequest) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *IncomeVerificationDocumentsDownloadRequest) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetSecret

`func (o *IncomeVerificationDocumentsDownloadRequest) GetSecret() string`

GetSecret returns the Secret field if non-nil, zero value otherwise.

### GetSecretOk

`func (o *IncomeVerificationDocumentsDownloadRequest) GetSecretOk() (*string, bool)`

GetSecretOk returns a tuple with the Secret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSecret

`func (o *IncomeVerificationDocumentsDownloadRequest) SetSecret(v string)`

SetSecret sets Secret field to given value.

### HasSecret

`func (o *IncomeVerificationDocumentsDownloadRequest) HasSecret() bool`

HasSecret returns a boolean if a field has been set.

### GetIncomeVerificationId

`func (o *IncomeVerificationDocumentsDownloadRequest) GetIncomeVerificationId() string`

GetIncomeVerificationId returns the IncomeVerificationId field if non-nil, zero value otherwise.

### GetIncomeVerificationIdOk

`func (o *IncomeVerificationDocumentsDownloadRequest) GetIncomeVerificationIdOk() (*string, bool)`

GetIncomeVerificationIdOk returns a tuple with the IncomeVerificationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncomeVerificationId

`func (o *IncomeVerificationDocumentsDownloadRequest) SetIncomeVerificationId(v string)`

SetIncomeVerificationId sets IncomeVerificationId field to given value.

### HasIncomeVerificationId

`func (o *IncomeVerificationDocumentsDownloadRequest) HasIncomeVerificationId() bool`

HasIncomeVerificationId returns a boolean if a field has been set.

### SetIncomeVerificationIdNil

`func (o *IncomeVerificationDocumentsDownloadRequest) SetIncomeVerificationIdNil(b bool)`

 SetIncomeVerificationIdNil sets the value for IncomeVerificationId to be an explicit nil

### UnsetIncomeVerificationId
`func (o *IncomeVerificationDocumentsDownloadRequest) UnsetIncomeVerificationId()`

UnsetIncomeVerificationId ensures that no value is present for IncomeVerificationId, not even an explicit nil
### GetAccessToken

`func (o *IncomeVerificationDocumentsDownloadRequest) GetAccessToken() string`

GetAccessToken returns the AccessToken field if non-nil, zero value otherwise.

### GetAccessTokenOk

`func (o *IncomeVerificationDocumentsDownloadRequest) GetAccessTokenOk() (*string, bool)`

GetAccessTokenOk returns a tuple with the AccessToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccessToken

`func (o *IncomeVerificationDocumentsDownloadRequest) SetAccessToken(v string)`

SetAccessToken sets AccessToken field to given value.

### HasAccessToken

`func (o *IncomeVerificationDocumentsDownloadRequest) HasAccessToken() bool`

HasAccessToken returns a boolean if a field has been set.

### SetAccessTokenNil

`func (o *IncomeVerificationDocumentsDownloadRequest) SetAccessTokenNil(b bool)`

 SetAccessTokenNil sets the value for AccessToken to be an explicit nil

### UnsetAccessToken
`func (o *IncomeVerificationDocumentsDownloadRequest) UnsetAccessToken()`

UnsetAccessToken ensures that no value is present for AccessToken, not even an explicit nil
### GetDocumentId

`func (o *IncomeVerificationDocumentsDownloadRequest) GetDocumentId() string`

GetDocumentId returns the DocumentId field if non-nil, zero value otherwise.

### GetDocumentIdOk

`func (o *IncomeVerificationDocumentsDownloadRequest) GetDocumentIdOk() (*string, bool)`

GetDocumentIdOk returns a tuple with the DocumentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentId

`func (o *IncomeVerificationDocumentsDownloadRequest) SetDocumentId(v string)`

SetDocumentId sets DocumentId field to given value.

### HasDocumentId

`func (o *IncomeVerificationDocumentsDownloadRequest) HasDocumentId() bool`

HasDocumentId returns a boolean if a field has been set.

### SetDocumentIdNil

`func (o *IncomeVerificationDocumentsDownloadRequest) SetDocumentIdNil(b bool)`

 SetDocumentIdNil sets the value for DocumentId to be an explicit nil

### UnsetDocumentId
`func (o *IncomeVerificationDocumentsDownloadRequest) UnsetDocumentId()`

UnsetDocumentId ensures that no value is present for DocumentId, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


