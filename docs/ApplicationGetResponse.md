# ApplicationGetResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 
**Application** | [**Application**](Application.md) |  | 

## Methods

### NewApplicationGetResponse

`func NewApplicationGetResponse(requestId string, application Application, ) *ApplicationGetResponse`

NewApplicationGetResponse instantiates a new ApplicationGetResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationGetResponseWithDefaults

`func NewApplicationGetResponseWithDefaults() *ApplicationGetResponse`

NewApplicationGetResponseWithDefaults instantiates a new ApplicationGetResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequestId

`func (o *ApplicationGetResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *ApplicationGetResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *ApplicationGetResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.


### GetApplication

`func (o *ApplicationGetResponse) GetApplication() Application`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *ApplicationGetResponse) GetApplicationOk() (*Application, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *ApplicationGetResponse) SetApplication(v Application)`

SetApplication sets Application field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


