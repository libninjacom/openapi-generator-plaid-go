# TransactionsSyncResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Added** | [**[]Transaction**](Transaction.md) | Transactions that have been added to the item since &#x60;cursor&#x60; ordered by ascending last modified time. | 
**Modified** | [**[]Transaction**](Transaction.md) | Transactions that have been modified on the item since &#x60;cursor&#x60; ordered by ascending last modified time. | 
**Removed** | [**[]RemovedTransaction**](RemovedTransaction.md) | Transactions that have been removed from the item since &#x60;cursor&#x60; ordered by ascending last modified time. | 
**NextCursor** | **string** | Cursor used for fetching any future updates after the latest update provided in this response. | 
**HasMore** | **bool** | Represents if more than requested count of transaction updates exist. If true, the additional updates can be fetched by making an additional request with &#x60;cursor&#x60; set to &#x60;next_cursor&#x60;. | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewTransactionsSyncResponse

`func NewTransactionsSyncResponse(added []Transaction, modified []Transaction, removed []RemovedTransaction, nextCursor string, hasMore bool, requestId string, ) *TransactionsSyncResponse`

NewTransactionsSyncResponse instantiates a new TransactionsSyncResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransactionsSyncResponseWithDefaults

`func NewTransactionsSyncResponseWithDefaults() *TransactionsSyncResponse`

NewTransactionsSyncResponseWithDefaults instantiates a new TransactionsSyncResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdded

`func (o *TransactionsSyncResponse) GetAdded() []Transaction`

GetAdded returns the Added field if non-nil, zero value otherwise.

### GetAddedOk

`func (o *TransactionsSyncResponse) GetAddedOk() (*[]Transaction, bool)`

GetAddedOk returns a tuple with the Added field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdded

`func (o *TransactionsSyncResponse) SetAdded(v []Transaction)`

SetAdded sets Added field to given value.


### GetModified

`func (o *TransactionsSyncResponse) GetModified() []Transaction`

GetModified returns the Modified field if non-nil, zero value otherwise.

### GetModifiedOk

`func (o *TransactionsSyncResponse) GetModifiedOk() (*[]Transaction, bool)`

GetModifiedOk returns a tuple with the Modified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModified

`func (o *TransactionsSyncResponse) SetModified(v []Transaction)`

SetModified sets Modified field to given value.


### GetRemoved

`func (o *TransactionsSyncResponse) GetRemoved() []RemovedTransaction`

GetRemoved returns the Removed field if non-nil, zero value otherwise.

### GetRemovedOk

`func (o *TransactionsSyncResponse) GetRemovedOk() (*[]RemovedTransaction, bool)`

GetRemovedOk returns a tuple with the Removed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoved

`func (o *TransactionsSyncResponse) SetRemoved(v []RemovedTransaction)`

SetRemoved sets Removed field to given value.


### GetNextCursor

`func (o *TransactionsSyncResponse) GetNextCursor() string`

GetNextCursor returns the NextCursor field if non-nil, zero value otherwise.

### GetNextCursorOk

`func (o *TransactionsSyncResponse) GetNextCursorOk() (*string, bool)`

GetNextCursorOk returns a tuple with the NextCursor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextCursor

`func (o *TransactionsSyncResponse) SetNextCursor(v string)`

SetNextCursor sets NextCursor field to given value.


### GetHasMore

`func (o *TransactionsSyncResponse) GetHasMore() bool`

GetHasMore returns the HasMore field if non-nil, zero value otherwise.

### GetHasMoreOk

`func (o *TransactionsSyncResponse) GetHasMoreOk() (*bool, bool)`

GetHasMoreOk returns a tuple with the HasMore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHasMore

`func (o *TransactionsSyncResponse) SetHasMore(v bool)`

SetHasMore sets HasMore field to given value.


### GetRequestId

`func (o *TransactionsSyncResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TransactionsSyncResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TransactionsSyncResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


