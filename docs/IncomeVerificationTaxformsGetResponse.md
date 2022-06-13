# IncomeVerificationTaxformsGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | [optional] 
**DocumentMetadata** | [**[]DocumentMetadata**](DocumentMetadata.md) |  | 
**Taxforms** | [**[]Taxform**](Taxform.md) | A list of forms. | 
**Error** | Pointer to [**PlaidError**](PlaidError.md) |  | [optional] 

## Methods

### NewIncomeVerificationTaxformsGetResponse

`func NewIncomeVerificationTaxformsGetResponse(documentMetadata []DocumentMetadata, taxforms []Taxform, ) *IncomeVerificationTaxformsGetResponse`

NewIncomeVerificationTaxformsGetResponse instantiates a new IncomeVerificationTaxformsGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeVerificationTaxformsGetResponseWithDefaults

`func NewIncomeVerificationTaxformsGetResponseWithDefaults() *IncomeVerificationTaxformsGetResponse`

NewIncomeVerificationTaxformsGetResponseWithDefaults instantiates a new IncomeVerificationTaxformsGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *IncomeVerificationTaxformsGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *IncomeVerificationTaxformsGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *IncomeVerificationTaxformsGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *IncomeVerificationTaxformsGetResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetDocumentMetadata

`func (o *IncomeVerificationTaxformsGetResponse) GetDocumentMetadata() []DocumentMetadata`

GetDocumentMetadata returns the DocumentMetadata field if non-nil, zero value otherwise.

### GetDocumentMetadataOk

`func (o *IncomeVerificationTaxformsGetResponse) GetDocumentMetadataOk() (*[]DocumentMetadata, bool)`

GetDocumentMetadataOk returns a tuple with the DocumentMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentMetadata

`func (o *IncomeVerificationTaxformsGetResponse) SetDocumentMetadata(v []DocumentMetadata)`

SetDocumentMetadata sets DocumentMetadata field to given value.


### GetTaxforms

`func (o *IncomeVerificationTaxformsGetResponse) GetTaxforms() []Taxform`

GetTaxforms returns the Taxforms field if non-nil, zero value otherwise.

### GetTaxformsOk

`func (o *IncomeVerificationTaxformsGetResponse) GetTaxformsOk() (*[]Taxform, bool)`

GetTaxformsOk returns a tuple with the Taxforms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTaxforms

`func (o *IncomeVerificationTaxformsGetResponse) SetTaxforms(v []Taxform)`

SetTaxforms sets Taxforms field to given value.


### GetError

`func (o *IncomeVerificationTaxformsGetResponse) GetError() PlaidError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *IncomeVerificationTaxformsGetResponse) GetErrorOk() (*PlaidError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *IncomeVerificationTaxformsGetResponse) SetError(v PlaidError)`

SetError sets Error field to given value.

### HasError

`func (o *IncomeVerificationTaxformsGetResponse) HasError() bool`

HasError returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


