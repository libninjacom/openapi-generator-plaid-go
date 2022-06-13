# Taxform

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DocId** | Pointer to **string** | An identifier of the document referenced by the document metadata. | [optional] 
**DocumentType** | **string** | The type of tax document. Currently, the only supported value is &#x60;w2&#x60;. | 
**W2** | Pointer to [**W2**](W2.md) |  | [optional] 

## Methods

### NewTaxform

`func NewTaxform(documentType string, ) *Taxform`

NewTaxform instantiates a new Taxform object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTaxformWithDefaults

`func NewTaxformWithDefaults() *Taxform`

NewTaxformWithDefaults instantiates a new Taxform object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDocId

`func (o *Taxform) GetDocId() string`

GetDocId returns the DocId field if non-nil, zero value otherwise.

### GetDocIdOk

`func (o *Taxform) GetDocIdOk() (*string, bool)`

GetDocIdOk returns a tuple with the DocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocId

`func (o *Taxform) SetDocId(v string)`

SetDocId sets DocId field to given value.

### HasDocId

`func (o *Taxform) HasDocId() bool`

HasDocId returns a boolean if a field has been set.

### GetDocumentType

`func (o *Taxform) GetDocumentType() string`

GetDocumentType returns the DocumentType field if non-nil, zero value otherwise.

### GetDocumentTypeOk

`func (o *Taxform) GetDocumentTypeOk() (*string, bool)`

GetDocumentTypeOk returns a tuple with the DocumentType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocumentType

`func (o *Taxform) SetDocumentType(v string)`

SetDocumentType sets DocumentType field to given value.


### GetW2

`func (o *Taxform) GetW2() W2`

GetW2 returns the W2 field if non-nil, zero value otherwise.

### GetW2Ok

`func (o *Taxform) GetW2Ok() (*W2, bool)`

GetW2Ok returns a tuple with the W2 field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetW2

`func (o *Taxform) SetW2(v W2)`

SetW2 sets W2 field to given value.

### HasW2

`func (o *Taxform) HasW2() bool`

HasW2 returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


