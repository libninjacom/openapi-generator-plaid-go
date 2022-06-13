# TransferRepaymentListResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Repayments** | [**[]TransferRepayment**](TransferRepayment.md) |  | 
**RequestId** | **string** | A unique identifier for the request, which can be used for troubleshooting. This identifier, like all Plaid identifiers, is case sensitive. | 

## Methods

### NewTransferRepaymentListResponse

`func NewTransferRepaymentListResponse(repayments []TransferRepayment, requestId string, ) *TransferRepaymentListResponse`

NewTransferRepaymentListResponse instantiates a new TransferRepaymentListResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferRepaymentListResponseWithDefaults

`func NewTransferRepaymentListResponseWithDefaults() *TransferRepaymentListResponse`

NewTransferRepaymentListResponseWithDefaults instantiates a new TransferRepaymentListResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRepayments

`func (o *TransferRepaymentListResponse) GetRepayments() []TransferRepayment`

GetRepayments returns the Repayments field if non-nil, zero value otherwise.

### GetRepaymentsOk

`func (o *TransferRepaymentListResponse) GetRepaymentsOk() (*[]TransferRepayment, bool)`

GetRepaymentsOk returns a tuple with the Repayments field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepayments

`func (o *TransferRepaymentListResponse) SetRepayments(v []TransferRepayment)`

SetRepayments sets Repayments field to given value.


### GetRequestId

`func (o *TransferRepaymentListResponse) GetRequestId() string`

GetRequestId returns the RequestId field if non-nil, zero value otherwise.

### GetRequestIdOk

`func (o *TransferRepaymentListResponse) GetRequestIdOk() (*string, bool)`

GetRequestIdOk returns a tuple with the RequestId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestId

`func (o *TransferRepaymentListResponse) SetRequestId(v string)`

SetRequestId sets RequestId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


