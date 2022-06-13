# IncomeVerificationPaystubsGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocumentMetadata** | Pointer to [**[]DocumentMetadata**](DocumentMetadata.md) | Metadata for an income document. | [optional] 
**Paystubs** | [**[]Paystub**](Paystub.md) |  | 
**Error** | Pointer to [**PlaidError**](PlaidError.md) |  | [optional] 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewIncomeVerificationPaystubsGetResponse

`func NewIncomeVerificationPaystubsGetResponse(paystubs []Paystub, requestId string, ) *IncomeVerificationPaystubsGetResponse`

NewIncomeVerificationPaystubsGetResponse instantiates a new IncomeVerificationPaystubsGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewIncomeVerificationPaystubsGetResponseWithDefaults

`func NewIncomeVerificationPaystubsGetResponseWithDefaults() *IncomeVerificationPaystubsGetResponse`

NewIncomeVerificationPaystubsGetResponseWithDefaults instantiates a new IncomeVerificationPaystubsGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocumentMetadata

`func (o *IncomeVerificationPaystubsGetResponse) GetDocumentMetadata() []DocumentMetadata`

GetDocumentMetadata returns the DocumentMetadata field if non-nil, zero value otherwise.

### GetDocumentMetadataOk

`func (o *IncomeVerificationPaystubsGetResponse) GetDocumentMetadataOk() (*[]DocumentMetadata, bool)`

GetDocumentMetadataOk returns a tuple with the DocumentMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentMetadata

`func (o *IncomeVerificationPaystubsGetResponse) SetDocumentMetadata(v []DocumentMetadata)`

SetDocumentMetadata sets DocumentMetadata field to given value.

### HasDocumentMetadata

`func (o *IncomeVerificationPaystubsGetResponse) HasDocumentMetadata() bool`

HasDocumentMetadata returns a boolean if a field has been set.

### GetPaystubs

`func (o *IncomeVerificationPaystubsGetResponse) GetPaystubs() []Paystub`

GetPaystubs returns the Paystubs field if non-nil, zero value otherwise.

### GetPaystubsOk

`func (o *IncomeVerificationPaystubsGetResponse) GetPaystubsOk() (*[]Paystub, bool)`

GetPaystubsOk returns a tuple with the Paystubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPaystubs

`func (o *IncomeVerificationPaystubsGetResponse) SetPaystubs(v []Paystub)`

SetPaystubs sets Paystubs field to given value.


### GetError

`func (o *IncomeVerificationPaystubsGetResponse) GetError() PlaidError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *IncomeVerificationPaystubsGetResponse) GetErrorOk() (*PlaidError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *IncomeVerificationPaystubsGetResponse) SetError(v PlaidError)`

SetError sets Error field to given value.

### HasError

`func (o *IncomeVerificationPaystubsGetResponse) HasError() bool`

HasError returns a boolean if a field has been set.

### GetRequestId

`func (o *IncomeVerificationPaystubsGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *IncomeVerificationPaystubsGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *IncomeVerificationPaystubsGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


