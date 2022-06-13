# DocumentMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | The name of the document. | [optional] 
**Status** | Pointer to **string** | The processing status of the document. | [optional] 
**DocId** | Pointer to **string** | An identifier of the document that is also present in the paystub response. | [optional] 
**DocType** | Pointer to [**DocType**](DocType.md) |  | [optional] 

## Methods

### NewDocumentMetadata

`func NewDocumentMetadata() *DocumentMetadata`

NewDocumentMetadata instantiates a new DocumentMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDocumentMetadataWithDefaults

`func NewDocumentMetadataWithDefaults() *DocumentMetadata`

NewDocumentMetadataWithDefaults instantiates a new DocumentMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *DocumentMetadata) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DocumentMetadata) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DocumentMetadata) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DocumentMetadata) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *DocumentMetadata) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DocumentMetadata) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DocumentMetadata) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DocumentMetadata) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDocId

`func (o *DocumentMetadata) GetDocId() string`

GetDocId returns the DocId field if non-nil, zero value otherwise.

### GetDocIdOk

`func (o *DocumentMetadata) GetDocIdOk() (*string, bool)`

GetDocIdOk returns a tuple with the DocId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocId

`func (o *DocumentMetadata) SetDocId(v string)`

SetDocId sets DocId field to given value.

### HasDocId

`func (o *DocumentMetadata) HasDocId() bool`

HasDocId returns a boolean if a field has been set.

### GetDocType

`func (o *DocumentMetadata) GetDocType() DocType`

GetDocType returns the DocType field if non-nil, zero value otherwise.

### GetDocTypeOk

`func (o *DocumentMetadata) GetDocTypeOk() (*DocType, bool)`

GetDocTypeOk returns a tuple with the DocType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDocType

`func (o *DocumentMetadata) SetDocType(v DocType)`

SetDocType sets DocType field to given value.

### HasDocType

`func (o *DocumentMetadata) HasDocType() bool`

HasDocType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


