# ProcessorBalanceGetRequestOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MinLastUpdatedDatetime** | Pointer to **time.Time** | Timestamp in [ISO 8601](https://wikipedia.org/wiki/ISO_8601) format (&#x60;YYYY-MM-DDTHH:mm:ssZ&#x60;) indicating the oldest acceptable balance when making a request to &#x60;/accounts/balance/get&#x60;.  If the balance that is pulled for &#x60;ins_128026&#x60; (Capital One) is older than the given timestamp, an &#x60;INVALID_REQUEST&#x60; error with the code of &#x60;LAST_UPDATED_DATETIME_OUT_OF_RANGE&#x60; will be returned with the most recent timestamp for the requested account contained in the response.  This field is only used when the institution is &#x60;ins_128026&#x60; (Capital One), in which case a value must be provided or an &#x60;INVALID_REQUEST&#x60; error with the code of &#x60;INVALID_FIELD&#x60; will be returned. For all other institutions, this field is ignored. | [optional] 

## Methods

### NewProcessorBalanceGetRequestOptions

`func NewProcessorBalanceGetRequestOptions() *ProcessorBalanceGetRequestOptions`

NewProcessorBalanceGetRequestOptions instantiates a new ProcessorBalanceGetRequestOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewProcessorBalanceGetRequestOptionsWithDefaults

`func NewProcessorBalanceGetRequestOptionsWithDefaults() *ProcessorBalanceGetRequestOptions`

NewProcessorBalanceGetRequestOptionsWithDefaults instantiates a new ProcessorBalanceGetRequestOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMinLastUpdatedDatetime

`func (o *ProcessorBalanceGetRequestOptions) GetMinLastUpdatedDatetime() time.Time`

GetMinLastUpdatedDatetime returns the MinLastUpdatedDatetime field if non-nil, zero value otherwise.

### GetMinLastUpdatedDatetimeOk

`func (o *ProcessorBalanceGetRequestOptions) GetMinLastUpdatedDatetimeOk() (*time.Time, bool)`

GetMinLastUpdatedDatetimeOk returns a tuple with the MinLastUpdatedDatetime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMinLastUpdatedDatetime

`func (o *ProcessorBalanceGetRequestOptions) SetMinLastUpdatedDatetime(v time.Time)`

SetMinLastUpdatedDatetime sets MinLastUpdatedDatetime field to given value.

### HasMinLastUpdatedDatetime

`func (o *ProcessorBalanceGetRequestOptions) HasMinLastUpdatedDatetime() bool`

HasMinLastUpdatedDatetime returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


