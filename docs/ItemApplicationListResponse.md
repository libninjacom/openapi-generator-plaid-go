# ItemApplicationListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | Pointer to **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | [optional] 
**Applications** | [**[]ConnectedApplication**](ConnectedApplication.md) | A list of connected applications. | 

## Methods

### NewItemApplicationListResponse

`func NewItemApplicationListResponse(applications []ConnectedApplication, ) *ItemApplicationListResponse`

NewItemApplicationListResponse instantiates a new ItemApplicationListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewItemApplicationListResponseWithDefaults

`func NewItemApplicationListResponseWithDefaults() *ItemApplicationListResponse`

NewItemApplicationListResponseWithDefaults instantiates a new ItemApplicationListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ItemApplicationListResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ItemApplicationListResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ItemApplicationListResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.

### HasRequestId

`func (o *ItemApplicationListResponse) HasRequestId() bool`

HasRequestId returns a boolean if a field has been set.

### GetApplications

`func (o *ItemApplicationListResponse) GetApplications() []ConnectedApplication`

GetApplications returns the Applications field if non-nil, zero value otherwise.

### GetApplicationsOk

`func (o *ItemApplicationListResponse) GetApplicationsOk() (*[]ConnectedApplication, bool)`

GetApplicationsOk returns a tuple with the Applications field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplications

`func (o *ItemApplicationListResponse) SetApplications(v []ConnectedApplication)`

SetApplications sets Applications field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


