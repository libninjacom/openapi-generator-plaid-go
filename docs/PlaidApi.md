# \PlaidApi

All URIs are relative to *https://production.plaid.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AccountsBalanceGet**](PlaidApi.md#AccountsBalanceGet) | **Post** /accounts/balance/get | Retrieve real-time balance data
[**AccountsGet**](PlaidApi.md#AccountsGet) | **Post** /accounts/get | Retrieve accounts
[**ApplicationGet**](PlaidApi.md#ApplicationGet) | **Post** /application/get | Retrieve information about a Plaid application
[**AssetReportAuditCopyCreate**](PlaidApi.md#AssetReportAuditCopyCreate) | **Post** /asset_report/audit_copy/create | Create Asset Report Audit Copy
[**AssetReportAuditCopyGet**](PlaidApi.md#AssetReportAuditCopyGet) | **Post** /asset_report/audit_copy/get | Retrieve an Asset Report Audit Copy
[**AssetReportAuditCopyRemove**](PlaidApi.md#AssetReportAuditCopyRemove) | **Post** /asset_report/audit_copy/remove | Remove Asset Report Audit Copy
[**AssetReportCreate**](PlaidApi.md#AssetReportCreate) | **Post** /asset_report/create | Create an Asset Report
[**AssetReportFilter**](PlaidApi.md#AssetReportFilter) | **Post** /asset_report/filter | Filter Asset Report
[**AssetReportGet**](PlaidApi.md#AssetReportGet) | **Post** /asset_report/get | Retrieve an Asset Report
[**AssetReportPdfGet**](PlaidApi.md#AssetReportPdfGet) | **Post** /asset_report/pdf/get | Retrieve a PDF Asset Report
[**AssetReportRefresh**](PlaidApi.md#AssetReportRefresh) | **Post** /asset_report/refresh | Refresh an Asset Report
[**AssetReportRemove**](PlaidApi.md#AssetReportRemove) | **Post** /asset_report/remove | Delete an Asset Report
[**AuthGet**](PlaidApi.md#AuthGet) | **Post** /auth/get | Retrieve auth data
[**BankTransferBalanceGet**](PlaidApi.md#BankTransferBalanceGet) | **Post** /bank_transfer/balance/get | Get balance of your Bank Transfer account
[**BankTransferCancel**](PlaidApi.md#BankTransferCancel) | **Post** /bank_transfer/cancel | Cancel a bank transfer
[**BankTransferCreate**](PlaidApi.md#BankTransferCreate) | **Post** /bank_transfer/create | Create a bank transfer
[**BankTransferEventList**](PlaidApi.md#BankTransferEventList) | **Post** /bank_transfer/event/list | List bank transfer events
[**BankTransferEventSync**](PlaidApi.md#BankTransferEventSync) | **Post** /bank_transfer/event/sync | Sync bank transfer events
[**BankTransferGet**](PlaidApi.md#BankTransferGet) | **Post** /bank_transfer/get | Retrieve a bank transfer
[**BankTransferList**](PlaidApi.md#BankTransferList) | **Post** /bank_transfer/list | List bank transfers
[**BankTransferMigrateAccount**](PlaidApi.md#BankTransferMigrateAccount) | **Post** /bank_transfer/migrate_account | Migrate account into Bank Transfers
[**BankTransferSweepGet**](PlaidApi.md#BankTransferSweepGet) | **Post** /bank_transfer/sweep/get | Retrieve a sweep
[**BankTransferSweepList**](PlaidApi.md#BankTransferSweepList) | **Post** /bank_transfer/sweep/list | List sweeps
[**CategoriesGet**](PlaidApi.md#CategoriesGet) | **Post** /categories/get | Get Categories
[**CreatePaymentToken**](PlaidApi.md#CreatePaymentToken) | **Post** /payment_initiation/payment/token/create | Create payment token
[**DepositSwitchAltCreate**](PlaidApi.md#DepositSwitchAltCreate) | **Post** /deposit_switch/alt/create | Create a deposit switch without using Plaid Exchange
[**DepositSwitchCreate**](PlaidApi.md#DepositSwitchCreate) | **Post** /deposit_switch/create | Create a deposit switch
[**DepositSwitchGet**](PlaidApi.md#DepositSwitchGet) | **Post** /deposit_switch/get | Retrieve a deposit switch
[**DepositSwitchTokenCreate**](PlaidApi.md#DepositSwitchTokenCreate) | **Post** /deposit_switch/token/create | Create a deposit switch token
[**EmployersSearch**](PlaidApi.md#EmployersSearch) | **Post** /employers/search | Search employer database
[**EmploymentVerificationGet**](PlaidApi.md#EmploymentVerificationGet) | **Post** /employment/verification/get | Retrieve a summary of an individual&#39;s employment information
[**IdentityGet**](PlaidApi.md#IdentityGet) | **Post** /identity/get | Retrieve identity data
[**IncomeVerificationCreate**](PlaidApi.md#IncomeVerificationCreate) | **Post** /income/verification/create | (Deprecated) Create an income verification instance
[**IncomeVerificationDocumentsDownload**](PlaidApi.md#IncomeVerificationDocumentsDownload) | **Post** /income/verification/documents/download | Download the original documents used for income verification
[**IncomeVerificationPaystubGet**](PlaidApi.md#IncomeVerificationPaystubGet) | **Post** /income/verification/paystub/get | (Deprecated) Retrieve information from a single paystub used for income verification
[**IncomeVerificationPaystubsGet**](PlaidApi.md#IncomeVerificationPaystubsGet) | **Post** /income/verification/paystubs/get | Retrieve information from the paystubs used for income verification
[**IncomeVerificationPrecheck**](PlaidApi.md#IncomeVerificationPrecheck) | **Post** /income/verification/precheck | Check digital income verification eligibility and optimize conversion
[**IncomeVerificationRefresh**](PlaidApi.md#IncomeVerificationRefresh) | **Post** /income/verification/refresh | Refresh an income verification
[**IncomeVerificationSummaryGet**](PlaidApi.md#IncomeVerificationSummaryGet) | **Post** /income/verification/summary/get | (Deprecated) Retrieve a summary of information derived from income verification
[**IncomeVerificationTaxformsGet**](PlaidApi.md#IncomeVerificationTaxformsGet) | **Post** /income/verification/taxforms/get | Retrieve information from the tax documents used for income verification
[**InstitutionsGet**](PlaidApi.md#InstitutionsGet) | **Post** /institutions/get | Get details of all supported institutions
[**InstitutionsGetById**](PlaidApi.md#InstitutionsGetById) | **Post** /institutions/get_by_id | Get details of an institution
[**InstitutionsSearch**](PlaidApi.md#InstitutionsSearch) | **Post** /institutions/search | Search institutions
[**InvestmentsHoldingsGet**](PlaidApi.md#InvestmentsHoldingsGet) | **Post** /investments/holdings/get | Get Investment holdings
[**InvestmentsTransactionsGet**](PlaidApi.md#InvestmentsTransactionsGet) | **Post** /investments/transactions/get | Get investment transactions
[**ItemAccessTokenInvalidate**](PlaidApi.md#ItemAccessTokenInvalidate) | **Post** /item/access_token/invalidate | Invalidate access_token
[**ItemApplicationList**](PlaidApi.md#ItemApplicationList) | **Post** /item/application/list | List a user’s connected applications
[**ItemApplicationScopesUpdate**](PlaidApi.md#ItemApplicationScopesUpdate) | **Post** /item/application/scopes/update | Update the scopes of access for a particular application
[**ItemCreatePublicToken**](PlaidApi.md#ItemCreatePublicToken) | **Post** /item/public_token/create | Create public token
[**ItemGet**](PlaidApi.md#ItemGet) | **Post** /item/get | Retrieve an Item
[**ItemImport**](PlaidApi.md#ItemImport) | **Post** /item/import | Import Item
[**ItemPublicTokenExchange**](PlaidApi.md#ItemPublicTokenExchange) | **Post** /item/public_token/exchange | Exchange public token for an access token
[**ItemRemove**](PlaidApi.md#ItemRemove) | **Post** /item/remove | Remove an Item
[**ItemWebhookUpdate**](PlaidApi.md#ItemWebhookUpdate) | **Post** /item/webhook/update | Update Webhook URL
[**LiabilitiesGet**](PlaidApi.md#LiabilitiesGet) | **Post** /liabilities/get | Retrieve Liabilities data
[**LinkTokenCreate**](PlaidApi.md#LinkTokenCreate) | **Post** /link/token/create | Create Link Token
[**LinkTokenGet**](PlaidApi.md#LinkTokenGet) | **Post** /link/token/get | Get Link Token
[**PaymentInitiationPaymentCreate**](PlaidApi.md#PaymentInitiationPaymentCreate) | **Post** /payment_initiation/payment/create | Create a payment
[**PaymentInitiationPaymentGet**](PlaidApi.md#PaymentInitiationPaymentGet) | **Post** /payment_initiation/payment/get | Get payment details
[**PaymentInitiationPaymentList**](PlaidApi.md#PaymentInitiationPaymentList) | **Post** /payment_initiation/payment/list | List payments
[**PaymentInitiationPaymentReverse**](PlaidApi.md#PaymentInitiationPaymentReverse) | **Post** /payment_initiation/payment/reverse | Reverse an existing payment
[**PaymentInitiationRecipientCreate**](PlaidApi.md#PaymentInitiationRecipientCreate) | **Post** /payment_initiation/recipient/create | Create payment recipient
[**PaymentInitiationRecipientGet**](PlaidApi.md#PaymentInitiationRecipientGet) | **Post** /payment_initiation/recipient/get | Get payment recipient
[**PaymentInitiationRecipientList**](PlaidApi.md#PaymentInitiationRecipientList) | **Post** /payment_initiation/recipient/list | List payment recipients
[**ProcessorApexProcessorTokenCreate**](PlaidApi.md#ProcessorApexProcessorTokenCreate) | **Post** /processor/apex/processor_token/create | Create Apex bank account token
[**ProcessorAuthGet**](PlaidApi.md#ProcessorAuthGet) | **Post** /processor/auth/get | Retrieve Auth data
[**ProcessorBalanceGet**](PlaidApi.md#ProcessorBalanceGet) | **Post** /processor/balance/get | Retrieve Balance data
[**ProcessorBankTransferCreate**](PlaidApi.md#ProcessorBankTransferCreate) | **Post** /processor/bank_transfer/create | Create a bank transfer as a processor
[**ProcessorIdentityGet**](PlaidApi.md#ProcessorIdentityGet) | **Post** /processor/identity/get | Retrieve Identity data
[**ProcessorStripeBankAccountTokenCreate**](PlaidApi.md#ProcessorStripeBankAccountTokenCreate) | **Post** /processor/stripe/bank_account_token/create | Create Stripe bank account token
[**ProcessorTokenCreate**](PlaidApi.md#ProcessorTokenCreate) | **Post** /processor/token/create | Create processor token
[**SandboxBankTransferFireWebhook**](PlaidApi.md#SandboxBankTransferFireWebhook) | **Post** /sandbox/bank_transfer/fire_webhook | Manually fire a Bank Transfer webhook
[**SandboxBankTransferSimulate**](PlaidApi.md#SandboxBankTransferSimulate) | **Post** /sandbox/bank_transfer/simulate | Simulate a bank transfer event in Sandbox
[**SandboxIncomeFireWebhook**](PlaidApi.md#SandboxIncomeFireWebhook) | **Post** /sandbox/income/fire_webhook | Manually fire an Income webhook
[**SandboxItemFireWebhook**](PlaidApi.md#SandboxItemFireWebhook) | **Post** /sandbox/item/fire_webhook | Fire a test webhook
[**SandboxItemResetLogin**](PlaidApi.md#SandboxItemResetLogin) | **Post** /sandbox/item/reset_login | Force a Sandbox Item into an error state
[**SandboxItemSetVerificationStatus**](PlaidApi.md#SandboxItemSetVerificationStatus) | **Post** /sandbox/item/set_verification_status | Set verification status for Sandbox account
[**SandboxOauthSelectAccounts**](PlaidApi.md#SandboxOauthSelectAccounts) | **Post** /sandbox/oauth/select_accounts | Save the selected accounts when connecting to the Platypus Oauth institution
[**SandboxProcessorTokenCreate**](PlaidApi.md#SandboxProcessorTokenCreate) | **Post** /sandbox/processor_token/create | Create a test Item and processor token
[**SandboxPublicTokenCreate**](PlaidApi.md#SandboxPublicTokenCreate) | **Post** /sandbox/public_token/create | Create a test Item
[**SandboxTransferRepaymentSimulate**](PlaidApi.md#SandboxTransferRepaymentSimulate) | **Post** /sandbox/transfer/repayment/simulate | Trigger the creation of a repayment
[**SandboxTransferSimulate**](PlaidApi.md#SandboxTransferSimulate) | **Post** /sandbox/transfer/simulate | Simulate a transfer event in Sandbox
[**SandboxTransferSweepSimulate**](PlaidApi.md#SandboxTransferSweepSimulate) | **Post** /sandbox/transfer/sweep/simulate | Simulate creating a sweep
[**SignalDecisionReport**](PlaidApi.md#SignalDecisionReport) | **Post** /signal/decision/report | Report whether you initiated an ACH transaction
[**SignalEvaluate**](PlaidApi.md#SignalEvaluate) | **Post** /signal/evaluate | Evaluate a planned ACH transaction
[**SignalReturnReport**](PlaidApi.md#SignalReturnReport) | **Post** /signal/return/report | Report a return for an ACH transaction
[**TransactionsGet**](PlaidApi.md#TransactionsGet) | **Post** /transactions/get | Get transaction data
[**TransactionsRecurringGet**](PlaidApi.md#TransactionsRecurringGet) | **Post** /transactions/recurring/get | Get streams of recurring transactions
[**TransactionsRefresh**](PlaidApi.md#TransactionsRefresh) | **Post** /transactions/refresh | Refresh transaction data
[**TransactionsSync**](PlaidApi.md#TransactionsSync) | **Post** /transactions/sync | Get incremental transaction updates on an Item
[**TransferAuthorizationCreate**](PlaidApi.md#TransferAuthorizationCreate) | **Post** /transfer/authorization/create | Create a transfer authorization
[**TransferCancel**](PlaidApi.md#TransferCancel) | **Post** /transfer/cancel | Cancel a transfer
[**TransferCreate**](PlaidApi.md#TransferCreate) | **Post** /transfer/create | Create a transfer
[**TransferEventList**](PlaidApi.md#TransferEventList) | **Post** /transfer/event/list | List transfer events
[**TransferEventSync**](PlaidApi.md#TransferEventSync) | **Post** /transfer/event/sync | Sync transfer events
[**TransferGet**](PlaidApi.md#TransferGet) | **Post** /transfer/get | Retrieve a transfer
[**TransferIntentCreate**](PlaidApi.md#TransferIntentCreate) | **Post** /transfer/intent/create | Create a transfer intent object to invoke the Transfer UI
[**TransferIntentGet**](PlaidApi.md#TransferIntentGet) | **Post** /transfer/intent/get | Retrieve more information about a transfer intent
[**TransferList**](PlaidApi.md#TransferList) | **Post** /transfer/list | List transfers
[**TransferRepaymentList**](PlaidApi.md#TransferRepaymentList) | **Post** /transfer/repayment/list | Lists historical repayments
[**TransferRepaymentReturnList**](PlaidApi.md#TransferRepaymentReturnList) | **Post** /transfer/repayment/return/list | List the returns included in a repayment
[**TransferSweepGet**](PlaidApi.md#TransferSweepGet) | **Post** /transfer/sweep/get | Retrieve a sweep
[**TransferSweepList**](PlaidApi.md#TransferSweepList) | **Post** /transfer/sweep/list | List sweeps
[**WalletGet**](PlaidApi.md#WalletGet) | **Post** /wallet/get | Fetch an e-wallet
[**WalletTransactionExecute**](PlaidApi.md#WalletTransactionExecute) | **Post** /wallet/transaction/execute | Execute a transaction using an e-wallet
[**WalletTransactionsList**](PlaidApi.md#WalletTransactionsList) | **Post** /wallet/transactions/list | List e-wallet transactions
[**WebhookVerificationKeyGet**](PlaidApi.md#WebhookVerificationKeyGet) | **Post** /webhook_verification_key/get | Get webhook verification key



## AccountsBalanceGet

> AccountsGetResponse AccountsBalanceGet(ctx).AccountsBalanceGetRequest(accountsBalanceGetRequest).Execute()

Retrieve real-time balance data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountsBalanceGetRequest := *openapiclient.NewAccountsBalanceGetRequest("AccessToken_example") // AccountsBalanceGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.AccountsBalanceGet(context.Background()).AccountsBalanceGetRequest(accountsBalanceGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AccountsBalanceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsBalanceGet`: AccountsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AccountsBalanceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsBalanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountsBalanceGetRequest** | [**AccountsBalanceGetRequest**](AccountsBalanceGetRequest.md) |  | 

### Return type

[**AccountsGetResponse**](AccountsGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AccountsGet

> AccountsGetResponse AccountsGet(ctx).AccountsGetRequest(accountsGetRequest).Execute()

Retrieve accounts



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    accountsGetRequest := *openapiclient.NewAccountsGetRequest("AccessToken_example") // AccountsGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.AccountsGet(context.Background()).AccountsGetRequest(accountsGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AccountsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AccountsGet`: AccountsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AccountsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAccountsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accountsGetRequest** | [**AccountsGetRequest**](AccountsGetRequest.md) |  | 

### Return type

[**AccountsGetResponse**](AccountsGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ApplicationGet

> ApplicationGetResponse ApplicationGet(ctx).ApplicationGetRequest(applicationGetRequest).Execute()

Retrieve information about a Plaid application



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    applicationGetRequest := *openapiclient.NewApplicationGetRequest("ClientId_example", "Secret_example", "ApplicationId_example") // ApplicationGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.ApplicationGet(context.Background()).ApplicationGetRequest(applicationGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ApplicationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ApplicationGet`: ApplicationGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ApplicationGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiApplicationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **applicationGetRequest** | [**ApplicationGetRequest**](ApplicationGetRequest.md) |  | 

### Return type

[**ApplicationGetResponse**](ApplicationGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetReportAuditCopyCreate

> AssetReportAuditCopyCreateResponse AssetReportAuditCopyCreate(ctx).AssetReportAuditCopyCreateRequest(assetReportAuditCopyCreateRequest).Execute()

Create Asset Report Audit Copy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetReportAuditCopyCreateRequest := *openapiclient.NewAssetReportAuditCopyCreateRequest("AssetReportToken_example", "AuditorId_example") // AssetReportAuditCopyCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.AssetReportAuditCopyCreate(context.Background()).AssetReportAuditCopyCreateRequest(assetReportAuditCopyCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AssetReportAuditCopyCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetReportAuditCopyCreate`: AssetReportAuditCopyCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AssetReportAuditCopyCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssetReportAuditCopyCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetReportAuditCopyCreateRequest** | [**AssetReportAuditCopyCreateRequest**](AssetReportAuditCopyCreateRequest.md) |  | 

### Return type

[**AssetReportAuditCopyCreateResponse**](AssetReportAuditCopyCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetReportAuditCopyGet

> AssetReportGetResponse AssetReportAuditCopyGet(ctx).AssetReportAuditCopyGetRequest(assetReportAuditCopyGetRequest).Execute()

Retrieve an Asset Report Audit Copy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetReportAuditCopyGetRequest := *openapiclient.NewAssetReportAuditCopyGetRequest("AuditCopyToken_example") // AssetReportAuditCopyGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.AssetReportAuditCopyGet(context.Background()).AssetReportAuditCopyGetRequest(assetReportAuditCopyGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AssetReportAuditCopyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetReportAuditCopyGet`: AssetReportGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AssetReportAuditCopyGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssetReportAuditCopyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetReportAuditCopyGetRequest** | [**AssetReportAuditCopyGetRequest**](AssetReportAuditCopyGetRequest.md) |  | 

### Return type

[**AssetReportGetResponse**](AssetReportGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetReportAuditCopyRemove

> AssetReportAuditCopyRemoveResponse AssetReportAuditCopyRemove(ctx).AssetReportAuditCopyRemoveRequest(assetReportAuditCopyRemoveRequest).Execute()

Remove Asset Report Audit Copy



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetReportAuditCopyRemoveRequest := *openapiclient.NewAssetReportAuditCopyRemoveRequest("AuditCopyToken_example") // AssetReportAuditCopyRemoveRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.AssetReportAuditCopyRemove(context.Background()).AssetReportAuditCopyRemoveRequest(assetReportAuditCopyRemoveRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AssetReportAuditCopyRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetReportAuditCopyRemove`: AssetReportAuditCopyRemoveResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AssetReportAuditCopyRemove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssetReportAuditCopyRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetReportAuditCopyRemoveRequest** | [**AssetReportAuditCopyRemoveRequest**](AssetReportAuditCopyRemoveRequest.md) |  | 

### Return type

[**AssetReportAuditCopyRemoveResponse**](AssetReportAuditCopyRemoveResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetReportCreate

> AssetReportCreateResponse AssetReportCreate(ctx).AssetReportCreateRequest(assetReportCreateRequest).Execute()

Create an Asset Report



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetReportCreateRequest := *openapiclient.NewAssetReportCreateRequest([]string{"AccessTokens_example"}, int32(123)) // AssetReportCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.AssetReportCreate(context.Background()).AssetReportCreateRequest(assetReportCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AssetReportCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetReportCreate`: AssetReportCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AssetReportCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssetReportCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetReportCreateRequest** | [**AssetReportCreateRequest**](AssetReportCreateRequest.md) |  | 

### Return type

[**AssetReportCreateResponse**](AssetReportCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetReportFilter

> AssetReportFilterResponse AssetReportFilter(ctx).AssetReportFilterRequest(assetReportFilterRequest).Execute()

Filter Asset Report



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetReportFilterRequest := *openapiclient.NewAssetReportFilterRequest("AssetReportToken_example", []string{"AccountIdsToExclude_example"}) // AssetReportFilterRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.AssetReportFilter(context.Background()).AssetReportFilterRequest(assetReportFilterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AssetReportFilter``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetReportFilter`: AssetReportFilterResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AssetReportFilter`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssetReportFilterRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetReportFilterRequest** | [**AssetReportFilterRequest**](AssetReportFilterRequest.md) |  | 

### Return type

[**AssetReportFilterResponse**](AssetReportFilterResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetReportGet

> AssetReportGetResponse AssetReportGet(ctx).AssetReportGetRequest(assetReportGetRequest).Execute()

Retrieve an Asset Report



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetReportGetRequest := *openapiclient.NewAssetReportGetRequest("AssetReportToken_example") // AssetReportGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.AssetReportGet(context.Background()).AssetReportGetRequest(assetReportGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AssetReportGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetReportGet`: AssetReportGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AssetReportGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssetReportGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetReportGetRequest** | [**AssetReportGetRequest**](AssetReportGetRequest.md) |  | 

### Return type

[**AssetReportGetResponse**](AssetReportGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetReportPdfGet

> *os.File AssetReportPdfGet(ctx).AssetReportPDFGetRequest(assetReportPDFGetRequest).Execute()

Retrieve a PDF Asset Report



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetReportPDFGetRequest := *openapiclient.NewAssetReportPDFGetRequest("AssetReportToken_example") // AssetReportPDFGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.AssetReportPdfGet(context.Background()).AssetReportPDFGetRequest(assetReportPDFGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AssetReportPdfGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetReportPdfGet`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AssetReportPdfGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssetReportPdfGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetReportPDFGetRequest** | [**AssetReportPDFGetRequest**](AssetReportPDFGetRequest.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/pdf

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetReportRefresh

> AssetReportRefreshResponse AssetReportRefresh(ctx).AssetReportRefreshRequest(assetReportRefreshRequest).Execute()

Refresh an Asset Report



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetReportRefreshRequest := *openapiclient.NewAssetReportRefreshRequest("AssetReportToken_example") // AssetReportRefreshRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.AssetReportRefresh(context.Background()).AssetReportRefreshRequest(assetReportRefreshRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AssetReportRefresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetReportRefresh`: AssetReportRefreshResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AssetReportRefresh`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssetReportRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetReportRefreshRequest** | [**AssetReportRefreshRequest**](AssetReportRefreshRequest.md) |  | 

### Return type

[**AssetReportRefreshResponse**](AssetReportRefreshResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AssetReportRemove

> AssetReportRemoveResponse AssetReportRemove(ctx).AssetReportRemoveRequest(assetReportRemoveRequest).Execute()

Delete an Asset Report



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    assetReportRemoveRequest := *openapiclient.NewAssetReportRemoveRequest("AssetReportToken_example") // AssetReportRemoveRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.AssetReportRemove(context.Background()).AssetReportRemoveRequest(assetReportRemoveRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AssetReportRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AssetReportRemove`: AssetReportRemoveResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AssetReportRemove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAssetReportRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **assetReportRemoveRequest** | [**AssetReportRemoveRequest**](AssetReportRemoveRequest.md) |  | 

### Return type

[**AssetReportRemoveResponse**](AssetReportRemoveResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AuthGet

> AuthGetResponse AuthGet(ctx).AuthGetRequest(authGetRequest).Execute()

Retrieve auth data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authGetRequest := *openapiclient.NewAuthGetRequest("AccessToken_example") // AuthGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.AuthGet(context.Background()).AuthGetRequest(authGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.AuthGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthGet`: AuthGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.AuthGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authGetRequest** | [**AuthGetRequest**](AuthGetRequest.md) |  | 

### Return type

[**AuthGetResponse**](AuthGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankTransferBalanceGet

> BankTransferBalanceGetResponse BankTransferBalanceGet(ctx).BankTransferBalanceGetRequest(bankTransferBalanceGetRequest).Execute()

Get balance of your Bank Transfer account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bankTransferBalanceGetRequest := *openapiclient.NewBankTransferBalanceGetRequest() // BankTransferBalanceGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.BankTransferBalanceGet(context.Background()).BankTransferBalanceGetRequest(bankTransferBalanceGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.BankTransferBalanceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankTransferBalanceGet`: BankTransferBalanceGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.BankTransferBalanceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBankTransferBalanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankTransferBalanceGetRequest** | [**BankTransferBalanceGetRequest**](BankTransferBalanceGetRequest.md) |  | 

### Return type

[**BankTransferBalanceGetResponse**](BankTransferBalanceGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankTransferCancel

> BankTransferCancelResponse BankTransferCancel(ctx).BankTransferCancelRequest(bankTransferCancelRequest).Execute()

Cancel a bank transfer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bankTransferCancelRequest := *openapiclient.NewBankTransferCancelRequest("BankTransferId_example") // BankTransferCancelRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.BankTransferCancel(context.Background()).BankTransferCancelRequest(bankTransferCancelRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.BankTransferCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankTransferCancel`: BankTransferCancelResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.BankTransferCancel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBankTransferCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankTransferCancelRequest** | [**BankTransferCancelRequest**](BankTransferCancelRequest.md) |  | 

### Return type

[**BankTransferCancelResponse**](BankTransferCancelResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankTransferCreate

> BankTransferCreateResponse BankTransferCreate(ctx).BankTransferCreateRequest(bankTransferCreateRequest).Execute()

Create a bank transfer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bankTransferCreateRequest := *openapiclient.NewBankTransferCreateRequest("IdempotencyKey_example", "AccessToken_example", "AccountId_example", openapiclient.BankTransferType("debit"), openapiclient.BankTransferNetwork("ach"), "Amount_example", "IsoCurrencyCode_example", "Description_example", *openapiclient.NewBankTransferUser("LegalName_example")) // BankTransferCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.BankTransferCreate(context.Background()).BankTransferCreateRequest(bankTransferCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.BankTransferCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankTransferCreate`: BankTransferCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.BankTransferCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBankTransferCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankTransferCreateRequest** | [**BankTransferCreateRequest**](BankTransferCreateRequest.md) |  | 

### Return type

[**BankTransferCreateResponse**](BankTransferCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankTransferEventList

> BankTransferEventListResponse BankTransferEventList(ctx).BankTransferEventListRequest(bankTransferEventListRequest).Execute()

List bank transfer events



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bankTransferEventListRequest := *openapiclient.NewBankTransferEventListRequest() // BankTransferEventListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.BankTransferEventList(context.Background()).BankTransferEventListRequest(bankTransferEventListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.BankTransferEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankTransferEventList`: BankTransferEventListResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.BankTransferEventList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBankTransferEventListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankTransferEventListRequest** | [**BankTransferEventListRequest**](BankTransferEventListRequest.md) |  | 

### Return type

[**BankTransferEventListResponse**](BankTransferEventListResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankTransferEventSync

> BankTransferEventSyncResponse BankTransferEventSync(ctx).BankTransferEventSyncRequest(bankTransferEventSyncRequest).Execute()

Sync bank transfer events



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bankTransferEventSyncRequest := *openapiclient.NewBankTransferEventSyncRequest(int32(123)) // BankTransferEventSyncRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.BankTransferEventSync(context.Background()).BankTransferEventSyncRequest(bankTransferEventSyncRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.BankTransferEventSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankTransferEventSync`: BankTransferEventSyncResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.BankTransferEventSync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBankTransferEventSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankTransferEventSyncRequest** | [**BankTransferEventSyncRequest**](BankTransferEventSyncRequest.md) |  | 

### Return type

[**BankTransferEventSyncResponse**](BankTransferEventSyncResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankTransferGet

> BankTransferGetResponse BankTransferGet(ctx).BankTransferGetRequest(bankTransferGetRequest).Execute()

Retrieve a bank transfer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bankTransferGetRequest := *openapiclient.NewBankTransferGetRequest("BankTransferId_example") // BankTransferGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.BankTransferGet(context.Background()).BankTransferGetRequest(bankTransferGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.BankTransferGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankTransferGet`: BankTransferGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.BankTransferGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBankTransferGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankTransferGetRequest** | [**BankTransferGetRequest**](BankTransferGetRequest.md) |  | 

### Return type

[**BankTransferGetResponse**](BankTransferGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankTransferList

> BankTransferListResponse BankTransferList(ctx).BankTransferListRequest(bankTransferListRequest).Execute()

List bank transfers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bankTransferListRequest := *openapiclient.NewBankTransferListRequest() // BankTransferListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.BankTransferList(context.Background()).BankTransferListRequest(bankTransferListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.BankTransferList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankTransferList`: BankTransferListResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.BankTransferList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBankTransferListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankTransferListRequest** | [**BankTransferListRequest**](BankTransferListRequest.md) |  | 

### Return type

[**BankTransferListResponse**](BankTransferListResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankTransferMigrateAccount

> BankTransferMigrateAccountResponse BankTransferMigrateAccount(ctx).BankTransferMigrateAccountRequest(bankTransferMigrateAccountRequest).Execute()

Migrate account into Bank Transfers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bankTransferMigrateAccountRequest := *openapiclient.NewBankTransferMigrateAccountRequest("AccountNumber_example", "RoutingNumber_example", "AccountType_example") // BankTransferMigrateAccountRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.BankTransferMigrateAccount(context.Background()).BankTransferMigrateAccountRequest(bankTransferMigrateAccountRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.BankTransferMigrateAccount``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankTransferMigrateAccount`: BankTransferMigrateAccountResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.BankTransferMigrateAccount`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBankTransferMigrateAccountRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankTransferMigrateAccountRequest** | [**BankTransferMigrateAccountRequest**](BankTransferMigrateAccountRequest.md) |  | 

### Return type

[**BankTransferMigrateAccountResponse**](BankTransferMigrateAccountResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankTransferSweepGet

> BankTransferSweepGetResponse BankTransferSweepGet(ctx).BankTransferSweepGetRequest(bankTransferSweepGetRequest).Execute()

Retrieve a sweep



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bankTransferSweepGetRequest := *openapiclient.NewBankTransferSweepGetRequest("SweepId_example") // BankTransferSweepGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.BankTransferSweepGet(context.Background()).BankTransferSweepGetRequest(bankTransferSweepGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.BankTransferSweepGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankTransferSweepGet`: BankTransferSweepGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.BankTransferSweepGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBankTransferSweepGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankTransferSweepGetRequest** | [**BankTransferSweepGetRequest**](BankTransferSweepGetRequest.md) |  | 

### Return type

[**BankTransferSweepGetResponse**](BankTransferSweepGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## BankTransferSweepList

> BankTransferSweepListResponse BankTransferSweepList(ctx).BankTransferSweepListRequest(bankTransferSweepListRequest).Execute()

List sweeps



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    bankTransferSweepListRequest := *openapiclient.NewBankTransferSweepListRequest() // BankTransferSweepListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.BankTransferSweepList(context.Background()).BankTransferSweepListRequest(bankTransferSweepListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.BankTransferSweepList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `BankTransferSweepList`: BankTransferSweepListResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.BankTransferSweepList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiBankTransferSweepListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **bankTransferSweepListRequest** | [**BankTransferSweepListRequest**](BankTransferSweepListRequest.md) |  | 

### Return type

[**BankTransferSweepListResponse**](BankTransferSweepListResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CategoriesGet

> CategoriesGetResponse CategoriesGet(ctx).Body(body).Execute()

Get Categories



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    body := map[string]interface{}{ ... } // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.CategoriesGet(context.Background()).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.CategoriesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CategoriesGet`: CategoriesGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.CategoriesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCategoriesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** |  | 

### Return type

[**CategoriesGetResponse**](CategoriesGetResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePaymentToken

> PaymentInitiationPaymentTokenCreateResponse CreatePaymentToken(ctx).PaymentInitiationPaymentTokenCreateRequest(paymentInitiationPaymentTokenCreateRequest).Execute()

Create payment token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    paymentInitiationPaymentTokenCreateRequest := *openapiclient.NewPaymentInitiationPaymentTokenCreateRequest("PaymentId_example") // PaymentInitiationPaymentTokenCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.CreatePaymentToken(context.Background()).PaymentInitiationPaymentTokenCreateRequest(paymentInitiationPaymentTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.CreatePaymentToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePaymentToken`: PaymentInitiationPaymentTokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.CreatePaymentToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePaymentTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentInitiationPaymentTokenCreateRequest** | [**PaymentInitiationPaymentTokenCreateRequest**](PaymentInitiationPaymentTokenCreateRequest.md) |  | 

### Return type

[**PaymentInitiationPaymentTokenCreateResponse**](PaymentInitiationPaymentTokenCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DepositSwitchAltCreate

> DepositSwitchAltCreateResponse DepositSwitchAltCreate(ctx).DepositSwitchAltCreateRequest(depositSwitchAltCreateRequest).Execute()

Create a deposit switch without using Plaid Exchange



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    depositSwitchAltCreateRequest := *openapiclient.NewDepositSwitchAltCreateRequest(*openapiclient.NewDepositSwitchTargetAccount("AccountNumber_example", "RoutingNumber_example", "AccountName_example", "AccountSubtype_example"), *openapiclient.NewDepositSwitchTargetUser("GivenName_example", "FamilyName_example", "Phone_example", "Email_example")) // DepositSwitchAltCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.DepositSwitchAltCreate(context.Background()).DepositSwitchAltCreateRequest(depositSwitchAltCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.DepositSwitchAltCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DepositSwitchAltCreate`: DepositSwitchAltCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.DepositSwitchAltCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDepositSwitchAltCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **depositSwitchAltCreateRequest** | [**DepositSwitchAltCreateRequest**](DepositSwitchAltCreateRequest.md) |  | 

### Return type

[**DepositSwitchAltCreateResponse**](DepositSwitchAltCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DepositSwitchCreate

> DepositSwitchCreateResponse DepositSwitchCreate(ctx).DepositSwitchCreateRequest(depositSwitchCreateRequest).Execute()

Create a deposit switch



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    depositSwitchCreateRequest := *openapiclient.NewDepositSwitchCreateRequest("TargetAccessToken_example", "TargetAccountId_example") // DepositSwitchCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.DepositSwitchCreate(context.Background()).DepositSwitchCreateRequest(depositSwitchCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.DepositSwitchCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DepositSwitchCreate`: DepositSwitchCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.DepositSwitchCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDepositSwitchCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **depositSwitchCreateRequest** | [**DepositSwitchCreateRequest**](DepositSwitchCreateRequest.md) |  | 

### Return type

[**DepositSwitchCreateResponse**](DepositSwitchCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DepositSwitchGet

> DepositSwitchGetResponse DepositSwitchGet(ctx).DepositSwitchGetRequest(depositSwitchGetRequest).Execute()

Retrieve a deposit switch



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    depositSwitchGetRequest := *openapiclient.NewDepositSwitchGetRequest("DepositSwitchId_example") // DepositSwitchGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.DepositSwitchGet(context.Background()).DepositSwitchGetRequest(depositSwitchGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.DepositSwitchGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DepositSwitchGet`: DepositSwitchGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.DepositSwitchGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDepositSwitchGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **depositSwitchGetRequest** | [**DepositSwitchGetRequest**](DepositSwitchGetRequest.md) |  | 

### Return type

[**DepositSwitchGetResponse**](DepositSwitchGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DepositSwitchTokenCreate

> DepositSwitchTokenCreateResponse DepositSwitchTokenCreate(ctx).DepositSwitchTokenCreateRequest(depositSwitchTokenCreateRequest).Execute()

Create a deposit switch token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    depositSwitchTokenCreateRequest := *openapiclient.NewDepositSwitchTokenCreateRequest("DepositSwitchId_example") // DepositSwitchTokenCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.DepositSwitchTokenCreate(context.Background()).DepositSwitchTokenCreateRequest(depositSwitchTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.DepositSwitchTokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DepositSwitchTokenCreate`: DepositSwitchTokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.DepositSwitchTokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDepositSwitchTokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **depositSwitchTokenCreateRequest** | [**DepositSwitchTokenCreateRequest**](DepositSwitchTokenCreateRequest.md) |  | 

### Return type

[**DepositSwitchTokenCreateResponse**](DepositSwitchTokenCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmployersSearch

> EmployersSearchResponse EmployersSearch(ctx).EmployersSearchRequest(employersSearchRequest).Execute()

Search employer database



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    employersSearchRequest := *openapiclient.NewEmployersSearchRequest("Query_example", []string{"Products_example"}) // EmployersSearchRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.EmployersSearch(context.Background()).EmployersSearchRequest(employersSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.EmployersSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmployersSearch`: EmployersSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.EmployersSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmployersSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **employersSearchRequest** | [**EmployersSearchRequest**](EmployersSearchRequest.md) |  | 

### Return type

[**EmployersSearchResponse**](EmployersSearchResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EmploymentVerificationGet

> EmploymentVerificationGetResponse EmploymentVerificationGet(ctx).EmploymentVerificationGetRequest(employmentVerificationGetRequest).Execute()

Retrieve a summary of an individual's employment information



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    employmentVerificationGetRequest := *openapiclient.NewEmploymentVerificationGetRequest("AccessToken_example") // EmploymentVerificationGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.EmploymentVerificationGet(context.Background()).EmploymentVerificationGetRequest(employmentVerificationGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.EmploymentVerificationGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `EmploymentVerificationGet`: EmploymentVerificationGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.EmploymentVerificationGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiEmploymentVerificationGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **employmentVerificationGetRequest** | [**EmploymentVerificationGetRequest**](EmploymentVerificationGetRequest.md) |  | 

### Return type

[**EmploymentVerificationGetResponse**](EmploymentVerificationGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IdentityGet

> IdentityGetResponse IdentityGet(ctx).IdentityGetRequest(identityGetRequest).Execute()

Retrieve identity data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    identityGetRequest := *openapiclient.NewIdentityGetRequest("AccessToken_example") // IdentityGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.IdentityGet(context.Background()).IdentityGetRequest(identityGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.IdentityGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IdentityGet`: IdentityGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.IdentityGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIdentityGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **identityGetRequest** | [**IdentityGetRequest**](IdentityGetRequest.md) |  | 

### Return type

[**IdentityGetResponse**](IdentityGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IncomeVerificationCreate

> IncomeVerificationCreateResponse IncomeVerificationCreate(ctx).IncomeVerificationCreateRequest(incomeVerificationCreateRequest).Execute()

(Deprecated) Create an income verification instance



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    incomeVerificationCreateRequest := *openapiclient.NewIncomeVerificationCreateRequest("Webhook_example") // IncomeVerificationCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.IncomeVerificationCreate(context.Background()).IncomeVerificationCreateRequest(incomeVerificationCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.IncomeVerificationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IncomeVerificationCreate`: IncomeVerificationCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.IncomeVerificationCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIncomeVerificationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **incomeVerificationCreateRequest** | [**IncomeVerificationCreateRequest**](IncomeVerificationCreateRequest.md) |  | 

### Return type

[**IncomeVerificationCreateResponse**](IncomeVerificationCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IncomeVerificationDocumentsDownload

> *os.File IncomeVerificationDocumentsDownload(ctx).IncomeVerificationDocumentsDownloadRequest(incomeVerificationDocumentsDownloadRequest).Execute()

Download the original documents used for income verification



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    incomeVerificationDocumentsDownloadRequest := *openapiclient.NewIncomeVerificationDocumentsDownloadRequest() // IncomeVerificationDocumentsDownloadRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.IncomeVerificationDocumentsDownload(context.Background()).IncomeVerificationDocumentsDownloadRequest(incomeVerificationDocumentsDownloadRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.IncomeVerificationDocumentsDownload``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IncomeVerificationDocumentsDownload`: *os.File
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.IncomeVerificationDocumentsDownload`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIncomeVerificationDocumentsDownloadRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **incomeVerificationDocumentsDownloadRequest** | [**IncomeVerificationDocumentsDownloadRequest**](IncomeVerificationDocumentsDownloadRequest.md) |  | 

### Return type

[***os.File**](*os.File.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IncomeVerificationPaystubGet

> IncomeVerificationPaystubGetResponse IncomeVerificationPaystubGet(ctx).IncomeVerificationPaystubGetRequest(incomeVerificationPaystubGetRequest).Execute()

(Deprecated) Retrieve information from a single paystub used for income verification



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    incomeVerificationPaystubGetRequest := *openapiclient.NewIncomeVerificationPaystubGetRequest() // IncomeVerificationPaystubGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.IncomeVerificationPaystubGet(context.Background()).IncomeVerificationPaystubGetRequest(incomeVerificationPaystubGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.IncomeVerificationPaystubGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IncomeVerificationPaystubGet`: IncomeVerificationPaystubGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.IncomeVerificationPaystubGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIncomeVerificationPaystubGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **incomeVerificationPaystubGetRequest** | [**IncomeVerificationPaystubGetRequest**](IncomeVerificationPaystubGetRequest.md) |  | 

### Return type

[**IncomeVerificationPaystubGetResponse**](IncomeVerificationPaystubGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IncomeVerificationPaystubsGet

> IncomeVerificationPaystubsGetResponse IncomeVerificationPaystubsGet(ctx).IncomeVerificationPaystubsGetRequest(incomeVerificationPaystubsGetRequest).Execute()

Retrieve information from the paystubs used for income verification



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    incomeVerificationPaystubsGetRequest := *openapiclient.NewIncomeVerificationPaystubsGetRequest() // IncomeVerificationPaystubsGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.IncomeVerificationPaystubsGet(context.Background()).IncomeVerificationPaystubsGetRequest(incomeVerificationPaystubsGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.IncomeVerificationPaystubsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IncomeVerificationPaystubsGet`: IncomeVerificationPaystubsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.IncomeVerificationPaystubsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIncomeVerificationPaystubsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **incomeVerificationPaystubsGetRequest** | [**IncomeVerificationPaystubsGetRequest**](IncomeVerificationPaystubsGetRequest.md) |  | 

### Return type

[**IncomeVerificationPaystubsGetResponse**](IncomeVerificationPaystubsGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IncomeVerificationPrecheck

> IncomeVerificationPrecheckResponse IncomeVerificationPrecheck(ctx).IncomeVerificationPrecheckRequest(incomeVerificationPrecheckRequest).Execute()

Check digital income verification eligibility and optimize conversion



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    incomeVerificationPrecheckRequest := *openapiclient.NewIncomeVerificationPrecheckRequest() // IncomeVerificationPrecheckRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.IncomeVerificationPrecheck(context.Background()).IncomeVerificationPrecheckRequest(incomeVerificationPrecheckRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.IncomeVerificationPrecheck``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IncomeVerificationPrecheck`: IncomeVerificationPrecheckResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.IncomeVerificationPrecheck`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIncomeVerificationPrecheckRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **incomeVerificationPrecheckRequest** | [**IncomeVerificationPrecheckRequest**](IncomeVerificationPrecheckRequest.md) |  | 

### Return type

[**IncomeVerificationPrecheckResponse**](IncomeVerificationPrecheckResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IncomeVerificationRefresh

> IncomeVerificationRefreshResponse IncomeVerificationRefresh(ctx).IncomeVerificationRefreshRequest(incomeVerificationRefreshRequest).Execute()

Refresh an income verification



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    incomeVerificationRefreshRequest := *openapiclient.NewIncomeVerificationRefreshRequest() // IncomeVerificationRefreshRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.IncomeVerificationRefresh(context.Background()).IncomeVerificationRefreshRequest(incomeVerificationRefreshRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.IncomeVerificationRefresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IncomeVerificationRefresh`: IncomeVerificationRefreshResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.IncomeVerificationRefresh`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIncomeVerificationRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **incomeVerificationRefreshRequest** | [**IncomeVerificationRefreshRequest**](IncomeVerificationRefreshRequest.md) |  | 

### Return type

[**IncomeVerificationRefreshResponse**](IncomeVerificationRefreshResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IncomeVerificationSummaryGet

> IncomeVerificationSummaryGetResponse IncomeVerificationSummaryGet(ctx).IncomeVerificationSummaryGetRequest(incomeVerificationSummaryGetRequest).Execute()

(Deprecated) Retrieve a summary of information derived from income verification



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    incomeVerificationSummaryGetRequest := *openapiclient.NewIncomeVerificationSummaryGetRequest() // IncomeVerificationSummaryGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.IncomeVerificationSummaryGet(context.Background()).IncomeVerificationSummaryGetRequest(incomeVerificationSummaryGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.IncomeVerificationSummaryGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IncomeVerificationSummaryGet`: IncomeVerificationSummaryGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.IncomeVerificationSummaryGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIncomeVerificationSummaryGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **incomeVerificationSummaryGetRequest** | [**IncomeVerificationSummaryGetRequest**](IncomeVerificationSummaryGetRequest.md) |  | 

### Return type

[**IncomeVerificationSummaryGetResponse**](IncomeVerificationSummaryGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IncomeVerificationTaxformsGet

> IncomeVerificationTaxformsGetResponse IncomeVerificationTaxformsGet(ctx).RequestBody(requestBody).Execute()

Retrieve information from the tax documents used for income verification



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.IncomeVerificationTaxformsGet(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.IncomeVerificationTaxformsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `IncomeVerificationTaxformsGet`: IncomeVerificationTaxformsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.IncomeVerificationTaxformsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiIncomeVerificationTaxformsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**IncomeVerificationTaxformsGetResponse**](IncomeVerificationTaxformsGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstitutionsGet

> InstitutionsGetResponse InstitutionsGet(ctx).InstitutionsGetRequest(institutionsGetRequest).Execute()

Get details of all supported institutions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    institutionsGetRequest := *openapiclient.NewInstitutionsGetRequest(int32(123), int32(123), []openapiclient.CountryCode{openapiclient.CountryCode("US")}) // InstitutionsGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.InstitutionsGet(context.Background()).InstitutionsGetRequest(institutionsGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.InstitutionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstitutionsGet`: InstitutionsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.InstitutionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstitutionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **institutionsGetRequest** | [**InstitutionsGetRequest**](InstitutionsGetRequest.md) |  | 

### Return type

[**InstitutionsGetResponse**](InstitutionsGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstitutionsGetById

> InstitutionsGetByIdResponse InstitutionsGetById(ctx).InstitutionsGetByIdRequest(institutionsGetByIdRequest).Execute()

Get details of an institution



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    institutionsGetByIdRequest := *openapiclient.NewInstitutionsGetByIdRequest("InstitutionId_example", []openapiclient.CountryCode{openapiclient.CountryCode("US")}) // InstitutionsGetByIdRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.InstitutionsGetById(context.Background()).InstitutionsGetByIdRequest(institutionsGetByIdRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.InstitutionsGetById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstitutionsGetById`: InstitutionsGetByIdResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.InstitutionsGetById`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstitutionsGetByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **institutionsGetByIdRequest** | [**InstitutionsGetByIdRequest**](InstitutionsGetByIdRequest.md) |  | 

### Return type

[**InstitutionsGetByIdResponse**](InstitutionsGetByIdResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InstitutionsSearch

> InstitutionsSearchResponse InstitutionsSearch(ctx).InstitutionsSearchRequest(institutionsSearchRequest).Execute()

Search institutions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    institutionsSearchRequest := *openapiclient.NewInstitutionsSearchRequest("Query_example", []openapiclient.Products{openapiclient.Products("assets")}, []openapiclient.CountryCode{openapiclient.CountryCode("US")}) // InstitutionsSearchRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.InstitutionsSearch(context.Background()).InstitutionsSearchRequest(institutionsSearchRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.InstitutionsSearch``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InstitutionsSearch`: InstitutionsSearchResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.InstitutionsSearch`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInstitutionsSearchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **institutionsSearchRequest** | [**InstitutionsSearchRequest**](InstitutionsSearchRequest.md) |  | 

### Return type

[**InstitutionsSearchResponse**](InstitutionsSearchResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvestmentsHoldingsGet

> InvestmentsHoldingsGetResponse InvestmentsHoldingsGet(ctx).InvestmentsHoldingsGetRequest(investmentsHoldingsGetRequest).Execute()

Get Investment holdings



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    investmentsHoldingsGetRequest := *openapiclient.NewInvestmentsHoldingsGetRequest("AccessToken_example") // InvestmentsHoldingsGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.InvestmentsHoldingsGet(context.Background()).InvestmentsHoldingsGetRequest(investmentsHoldingsGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.InvestmentsHoldingsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvestmentsHoldingsGet`: InvestmentsHoldingsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.InvestmentsHoldingsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvestmentsHoldingsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **investmentsHoldingsGetRequest** | [**InvestmentsHoldingsGetRequest**](InvestmentsHoldingsGetRequest.md) |  | 

### Return type

[**InvestmentsHoldingsGetResponse**](InvestmentsHoldingsGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## InvestmentsTransactionsGet

> InvestmentsTransactionsGetResponse InvestmentsTransactionsGet(ctx).InvestmentsTransactionsGetRequest(investmentsTransactionsGetRequest).Execute()

Get investment transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    investmentsTransactionsGetRequest := *openapiclient.NewInvestmentsTransactionsGetRequest("AccessToken_example", time.Now(), time.Now()) // InvestmentsTransactionsGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.InvestmentsTransactionsGet(context.Background()).InvestmentsTransactionsGetRequest(investmentsTransactionsGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.InvestmentsTransactionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `InvestmentsTransactionsGet`: InvestmentsTransactionsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.InvestmentsTransactionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiInvestmentsTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **investmentsTransactionsGetRequest** | [**InvestmentsTransactionsGetRequest**](InvestmentsTransactionsGetRequest.md) |  | 

### Return type

[**InvestmentsTransactionsGetResponse**](InvestmentsTransactionsGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ItemAccessTokenInvalidate

> ItemAccessTokenInvalidateResponse ItemAccessTokenInvalidate(ctx).ItemAccessTokenInvalidateRequest(itemAccessTokenInvalidateRequest).Execute()

Invalidate access_token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    itemAccessTokenInvalidateRequest := *openapiclient.NewItemAccessTokenInvalidateRequest("AccessToken_example") // ItemAccessTokenInvalidateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.ItemAccessTokenInvalidate(context.Background()).ItemAccessTokenInvalidateRequest(itemAccessTokenInvalidateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ItemAccessTokenInvalidate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ItemAccessTokenInvalidate`: ItemAccessTokenInvalidateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ItemAccessTokenInvalidate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiItemAccessTokenInvalidateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemAccessTokenInvalidateRequest** | [**ItemAccessTokenInvalidateRequest**](ItemAccessTokenInvalidateRequest.md) |  | 

### Return type

[**ItemAccessTokenInvalidateResponse**](ItemAccessTokenInvalidateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ItemApplicationList

> ItemApplicationListResponse ItemApplicationList(ctx).ItemApplicationListRequest(itemApplicationListRequest).Execute()

List a user’s connected applications



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    itemApplicationListRequest := *openapiclient.NewItemApplicationListRequest() // ItemApplicationListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.ItemApplicationList(context.Background()).ItemApplicationListRequest(itemApplicationListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ItemApplicationList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ItemApplicationList`: ItemApplicationListResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ItemApplicationList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiItemApplicationListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemApplicationListRequest** | [**ItemApplicationListRequest**](ItemApplicationListRequest.md) |  | 

### Return type

[**ItemApplicationListResponse**](ItemApplicationListResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ItemApplicationScopesUpdate

> ItemApplicationScopesUpdateResponse ItemApplicationScopesUpdate(ctx).ItemApplicationScopesUpdateRequest(itemApplicationScopesUpdateRequest).Execute()

Update the scopes of access for a particular application



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    itemApplicationScopesUpdateRequest := *openapiclient.NewItemApplicationScopesUpdateRequest("AccessToken_example", "ApplicationId_example", *openapiclient.NewScopes(), openapiclient.ScopesContext("ENROLLMENT")) // ItemApplicationScopesUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.ItemApplicationScopesUpdate(context.Background()).ItemApplicationScopesUpdateRequest(itemApplicationScopesUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ItemApplicationScopesUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ItemApplicationScopesUpdate`: ItemApplicationScopesUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ItemApplicationScopesUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiItemApplicationScopesUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemApplicationScopesUpdateRequest** | [**ItemApplicationScopesUpdateRequest**](ItemApplicationScopesUpdateRequest.md) |  | 

### Return type

[**ItemApplicationScopesUpdateResponse**](ItemApplicationScopesUpdateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ItemCreatePublicToken

> ItemPublicTokenCreateResponse ItemCreatePublicToken(ctx).ItemPublicTokenCreateRequest(itemPublicTokenCreateRequest).Execute()

Create public token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    itemPublicTokenCreateRequest := *openapiclient.NewItemPublicTokenCreateRequest("AccessToken_example") // ItemPublicTokenCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.ItemCreatePublicToken(context.Background()).ItemPublicTokenCreateRequest(itemPublicTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ItemCreatePublicToken``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ItemCreatePublicToken`: ItemPublicTokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ItemCreatePublicToken`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiItemCreatePublicTokenRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemPublicTokenCreateRequest** | [**ItemPublicTokenCreateRequest**](ItemPublicTokenCreateRequest.md) |  | 

### Return type

[**ItemPublicTokenCreateResponse**](ItemPublicTokenCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ItemGet

> ItemGetResponse ItemGet(ctx).ItemGetRequest(itemGetRequest).Execute()

Retrieve an Item



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    itemGetRequest := *openapiclient.NewItemGetRequest("AccessToken_example") // ItemGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.ItemGet(context.Background()).ItemGetRequest(itemGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ItemGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ItemGet`: ItemGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ItemGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiItemGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemGetRequest** | [**ItemGetRequest**](ItemGetRequest.md) |  | 

### Return type

[**ItemGetResponse**](ItemGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ItemImport

> ItemImportResponse ItemImport(ctx).ItemImportRequest(itemImportRequest).Execute()

Import Item



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    itemImportRequest := *openapiclient.NewItemImportRequest([]openapiclient.Products{openapiclient.Products("assets")}, *openapiclient.NewItemImportRequestUserAuth("UserId_example", "AuthToken_example")) // ItemImportRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.ItemImport(context.Background()).ItemImportRequest(itemImportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ItemImport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ItemImport`: ItemImportResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ItemImport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiItemImportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemImportRequest** | [**ItemImportRequest**](ItemImportRequest.md) |  | 

### Return type

[**ItemImportResponse**](ItemImportResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ItemPublicTokenExchange

> ItemPublicTokenExchangeResponse ItemPublicTokenExchange(ctx).ItemPublicTokenExchangeRequest(itemPublicTokenExchangeRequest).Execute()

Exchange public token for an access token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    itemPublicTokenExchangeRequest := *openapiclient.NewItemPublicTokenExchangeRequest("PublicToken_example") // ItemPublicTokenExchangeRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.ItemPublicTokenExchange(context.Background()).ItemPublicTokenExchangeRequest(itemPublicTokenExchangeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ItemPublicTokenExchange``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ItemPublicTokenExchange`: ItemPublicTokenExchangeResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ItemPublicTokenExchange`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiItemPublicTokenExchangeRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemPublicTokenExchangeRequest** | [**ItemPublicTokenExchangeRequest**](ItemPublicTokenExchangeRequest.md) |  | 

### Return type

[**ItemPublicTokenExchangeResponse**](ItemPublicTokenExchangeResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ItemRemove

> ItemRemoveResponse ItemRemove(ctx).ItemRemoveRequest(itemRemoveRequest).Execute()

Remove an Item



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    itemRemoveRequest := *openapiclient.NewItemRemoveRequest("AccessToken_example") // ItemRemoveRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.ItemRemove(context.Background()).ItemRemoveRequest(itemRemoveRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ItemRemove``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ItemRemove`: ItemRemoveResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ItemRemove`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiItemRemoveRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemRemoveRequest** | [**ItemRemoveRequest**](ItemRemoveRequest.md) |  | 

### Return type

[**ItemRemoveResponse**](ItemRemoveResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ItemWebhookUpdate

> ItemWebhookUpdateResponse ItemWebhookUpdate(ctx).ItemWebhookUpdateRequest(itemWebhookUpdateRequest).Execute()

Update Webhook URL



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    itemWebhookUpdateRequest := *openapiclient.NewItemWebhookUpdateRequest("AccessToken_example") // ItemWebhookUpdateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.ItemWebhookUpdate(context.Background()).ItemWebhookUpdateRequest(itemWebhookUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ItemWebhookUpdate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ItemWebhookUpdate`: ItemWebhookUpdateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ItemWebhookUpdate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiItemWebhookUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **itemWebhookUpdateRequest** | [**ItemWebhookUpdateRequest**](ItemWebhookUpdateRequest.md) |  | 

### Return type

[**ItemWebhookUpdateResponse**](ItemWebhookUpdateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LiabilitiesGet

> LiabilitiesGetResponse LiabilitiesGet(ctx).LiabilitiesGetRequest(liabilitiesGetRequest).Execute()

Retrieve Liabilities data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    liabilitiesGetRequest := *openapiclient.NewLiabilitiesGetRequest("AccessToken_example") // LiabilitiesGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.LiabilitiesGet(context.Background()).LiabilitiesGetRequest(liabilitiesGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.LiabilitiesGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LiabilitiesGet`: LiabilitiesGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.LiabilitiesGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLiabilitiesGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **liabilitiesGetRequest** | [**LiabilitiesGetRequest**](LiabilitiesGetRequest.md) |  | 

### Return type

[**LiabilitiesGetResponse**](LiabilitiesGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkTokenCreate

> LinkTokenCreateResponse LinkTokenCreate(ctx).LinkTokenCreateRequest(linkTokenCreateRequest).Execute()

Create Link Token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    linkTokenCreateRequest := *openapiclient.NewLinkTokenCreateRequest("ClientName_example", "Language_example", []openapiclient.CountryCode{openapiclient.CountryCode("US")}, *openapiclient.NewLinkTokenCreateRequestUser("ClientUserId_example")) // LinkTokenCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.LinkTokenCreate(context.Background()).LinkTokenCreateRequest(linkTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.LinkTokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkTokenCreate`: LinkTokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.LinkTokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkTokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkTokenCreateRequest** | [**LinkTokenCreateRequest**](LinkTokenCreateRequest.md) |  | 

### Return type

[**LinkTokenCreateResponse**](LinkTokenCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LinkTokenGet

> LinkTokenGetResponse LinkTokenGet(ctx).LinkTokenGetRequest(linkTokenGetRequest).Execute()

Get Link Token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    linkTokenGetRequest := *openapiclient.NewLinkTokenGetRequest("LinkToken_example") // LinkTokenGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.LinkTokenGet(context.Background()).LinkTokenGetRequest(linkTokenGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.LinkTokenGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `LinkTokenGet`: LinkTokenGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.LinkTokenGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLinkTokenGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **linkTokenGetRequest** | [**LinkTokenGetRequest**](LinkTokenGetRequest.md) |  | 

### Return type

[**LinkTokenGetResponse**](LinkTokenGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentInitiationPaymentCreate

> PaymentInitiationPaymentCreateResponse PaymentInitiationPaymentCreate(ctx).PaymentInitiationPaymentCreateRequest(paymentInitiationPaymentCreateRequest).Execute()

Create a payment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    paymentInitiationPaymentCreateRequest := *openapiclient.NewPaymentInitiationPaymentCreateRequest("RecipientId_example", "Reference_example", *openapiclient.NewPaymentAmount("Currency_example", float32(123))) // PaymentInitiationPaymentCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.PaymentInitiationPaymentCreate(context.Background()).PaymentInitiationPaymentCreateRequest(paymentInitiationPaymentCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.PaymentInitiationPaymentCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentInitiationPaymentCreate`: PaymentInitiationPaymentCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.PaymentInitiationPaymentCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentInitiationPaymentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentInitiationPaymentCreateRequest** | [**PaymentInitiationPaymentCreateRequest**](PaymentInitiationPaymentCreateRequest.md) |  | 

### Return type

[**PaymentInitiationPaymentCreateResponse**](PaymentInitiationPaymentCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentInitiationPaymentGet

> PaymentInitiationPaymentGetResponse PaymentInitiationPaymentGet(ctx).PaymentInitiationPaymentGetRequest(paymentInitiationPaymentGetRequest).Execute()

Get payment details



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    paymentInitiationPaymentGetRequest := *openapiclient.NewPaymentInitiationPaymentGetRequest("PaymentId_example") // PaymentInitiationPaymentGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.PaymentInitiationPaymentGet(context.Background()).PaymentInitiationPaymentGetRequest(paymentInitiationPaymentGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.PaymentInitiationPaymentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentInitiationPaymentGet`: PaymentInitiationPaymentGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.PaymentInitiationPaymentGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentInitiationPaymentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentInitiationPaymentGetRequest** | [**PaymentInitiationPaymentGetRequest**](PaymentInitiationPaymentGetRequest.md) |  | 

### Return type

[**PaymentInitiationPaymentGetResponse**](PaymentInitiationPaymentGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentInitiationPaymentList

> PaymentInitiationPaymentListResponse PaymentInitiationPaymentList(ctx).PaymentInitiationPaymentListRequest(paymentInitiationPaymentListRequest).Execute()

List payments



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    paymentInitiationPaymentListRequest := *openapiclient.NewPaymentInitiationPaymentListRequest() // PaymentInitiationPaymentListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.PaymentInitiationPaymentList(context.Background()).PaymentInitiationPaymentListRequest(paymentInitiationPaymentListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.PaymentInitiationPaymentList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentInitiationPaymentList`: PaymentInitiationPaymentListResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.PaymentInitiationPaymentList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentInitiationPaymentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentInitiationPaymentListRequest** | [**PaymentInitiationPaymentListRequest**](PaymentInitiationPaymentListRequest.md) |  | 

### Return type

[**PaymentInitiationPaymentListResponse**](PaymentInitiationPaymentListResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentInitiationPaymentReverse

> PaymentInitiationPaymentReverseResponse PaymentInitiationPaymentReverse(ctx).PaymentInitiationPaymentReverseRequest(paymentInitiationPaymentReverseRequest).Execute()

Reverse an existing payment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    paymentInitiationPaymentReverseRequest := *openapiclient.NewPaymentInitiationPaymentReverseRequest("PaymentId_example") // PaymentInitiationPaymentReverseRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.PaymentInitiationPaymentReverse(context.Background()).PaymentInitiationPaymentReverseRequest(paymentInitiationPaymentReverseRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.PaymentInitiationPaymentReverse``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentInitiationPaymentReverse`: PaymentInitiationPaymentReverseResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.PaymentInitiationPaymentReverse`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentInitiationPaymentReverseRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentInitiationPaymentReverseRequest** | [**PaymentInitiationPaymentReverseRequest**](PaymentInitiationPaymentReverseRequest.md) |  | 

### Return type

[**PaymentInitiationPaymentReverseResponse**](PaymentInitiationPaymentReverseResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentInitiationRecipientCreate

> PaymentInitiationRecipientCreateResponse PaymentInitiationRecipientCreate(ctx).PaymentInitiationRecipientCreateRequest(paymentInitiationRecipientCreateRequest).Execute()

Create payment recipient



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    paymentInitiationRecipientCreateRequest := *openapiclient.NewPaymentInitiationRecipientCreateRequest("Name_example") // PaymentInitiationRecipientCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.PaymentInitiationRecipientCreate(context.Background()).PaymentInitiationRecipientCreateRequest(paymentInitiationRecipientCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.PaymentInitiationRecipientCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentInitiationRecipientCreate`: PaymentInitiationRecipientCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.PaymentInitiationRecipientCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentInitiationRecipientCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentInitiationRecipientCreateRequest** | [**PaymentInitiationRecipientCreateRequest**](PaymentInitiationRecipientCreateRequest.md) |  | 

### Return type

[**PaymentInitiationRecipientCreateResponse**](PaymentInitiationRecipientCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentInitiationRecipientGet

> PaymentInitiationRecipientGetResponse PaymentInitiationRecipientGet(ctx).PaymentInitiationRecipientGetRequest(paymentInitiationRecipientGetRequest).Execute()

Get payment recipient



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    paymentInitiationRecipientGetRequest := *openapiclient.NewPaymentInitiationRecipientGetRequest("RecipientId_example") // PaymentInitiationRecipientGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.PaymentInitiationRecipientGet(context.Background()).PaymentInitiationRecipientGetRequest(paymentInitiationRecipientGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.PaymentInitiationRecipientGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentInitiationRecipientGet`: PaymentInitiationRecipientGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.PaymentInitiationRecipientGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentInitiationRecipientGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentInitiationRecipientGetRequest** | [**PaymentInitiationRecipientGetRequest**](PaymentInitiationRecipientGetRequest.md) |  | 

### Return type

[**PaymentInitiationRecipientGetResponse**](PaymentInitiationRecipientGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PaymentInitiationRecipientList

> PaymentInitiationRecipientListResponse PaymentInitiationRecipientList(ctx).PaymentInitiationRecipientListRequest(paymentInitiationRecipientListRequest).Execute()

List payment recipients



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    paymentInitiationRecipientListRequest := *openapiclient.NewPaymentInitiationRecipientListRequest() // PaymentInitiationRecipientListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.PaymentInitiationRecipientList(context.Background()).PaymentInitiationRecipientListRequest(paymentInitiationRecipientListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.PaymentInitiationRecipientList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PaymentInitiationRecipientList`: PaymentInitiationRecipientListResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.PaymentInitiationRecipientList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPaymentInitiationRecipientListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **paymentInitiationRecipientListRequest** | [**PaymentInitiationRecipientListRequest**](PaymentInitiationRecipientListRequest.md) |  | 

### Return type

[**PaymentInitiationRecipientListResponse**](PaymentInitiationRecipientListResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessorApexProcessorTokenCreate

> ProcessorTokenCreateResponse ProcessorApexProcessorTokenCreate(ctx).ProcessorApexProcessorTokenCreateRequest(processorApexProcessorTokenCreateRequest).Execute()

Create Apex bank account token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    processorApexProcessorTokenCreateRequest := *openapiclient.NewProcessorApexProcessorTokenCreateRequest("AccessToken_example", "AccountId_example") // ProcessorApexProcessorTokenCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.ProcessorApexProcessorTokenCreate(context.Background()).ProcessorApexProcessorTokenCreateRequest(processorApexProcessorTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ProcessorApexProcessorTokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessorApexProcessorTokenCreate`: ProcessorTokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ProcessorApexProcessorTokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessorApexProcessorTokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processorApexProcessorTokenCreateRequest** | [**ProcessorApexProcessorTokenCreateRequest**](ProcessorApexProcessorTokenCreateRequest.md) |  | 

### Return type

[**ProcessorTokenCreateResponse**](ProcessorTokenCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessorAuthGet

> ProcessorAuthGetResponse ProcessorAuthGet(ctx).ProcessorAuthGetRequest(processorAuthGetRequest).Execute()

Retrieve Auth data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    processorAuthGetRequest := *openapiclient.NewProcessorAuthGetRequest("ProcessorToken_example") // ProcessorAuthGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.ProcessorAuthGet(context.Background()).ProcessorAuthGetRequest(processorAuthGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ProcessorAuthGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessorAuthGet`: ProcessorAuthGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ProcessorAuthGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessorAuthGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processorAuthGetRequest** | [**ProcessorAuthGetRequest**](ProcessorAuthGetRequest.md) |  | 

### Return type

[**ProcessorAuthGetResponse**](ProcessorAuthGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessorBalanceGet

> ProcessorBalanceGetResponse ProcessorBalanceGet(ctx).ProcessorBalanceGetRequest(processorBalanceGetRequest).Execute()

Retrieve Balance data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    processorBalanceGetRequest := *openapiclient.NewProcessorBalanceGetRequest("ProcessorToken_example") // ProcessorBalanceGetRequest | The `/processor/balance/get` endpoint returns the real-time balance for the account associated with a given `processor_token`.  The current balance is the total amount of funds in the account. The available balance is the current balance less any outstanding holds or debits that have not yet posted to the account.  Note that not all institutions calculate the available balance. In the event that available balance is unavailable from the institution, Plaid will return an available balance value of `null`.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.ProcessorBalanceGet(context.Background()).ProcessorBalanceGetRequest(processorBalanceGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ProcessorBalanceGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessorBalanceGet`: ProcessorBalanceGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ProcessorBalanceGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessorBalanceGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processorBalanceGetRequest** | [**ProcessorBalanceGetRequest**](ProcessorBalanceGetRequest.md) | The &#x60;/processor/balance/get&#x60; endpoint returns the real-time balance for the account associated with a given &#x60;processor_token&#x60;.  The current balance is the total amount of funds in the account. The available balance is the current balance less any outstanding holds or debits that have not yet posted to the account.  Note that not all institutions calculate the available balance. In the event that available balance is unavailable from the institution, Plaid will return an available balance value of &#x60;null&#x60;. | 

### Return type

[**ProcessorBalanceGetResponse**](ProcessorBalanceGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessorBankTransferCreate

> ProcessorBankTransferCreateResponse ProcessorBankTransferCreate(ctx).ProcessorBankTransferCreateRequest(processorBankTransferCreateRequest).Execute()

Create a bank transfer as a processor



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    processorBankTransferCreateRequest := *openapiclient.NewProcessorBankTransferCreateRequest("IdempotencyKey_example", "ProcessorToken_example", openapiclient.BankTransferType("debit"), openapiclient.BankTransferNetwork("ach"), "Amount_example", "IsoCurrencyCode_example", "Description_example", *openapiclient.NewBankTransferUser("LegalName_example")) // ProcessorBankTransferCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.ProcessorBankTransferCreate(context.Background()).ProcessorBankTransferCreateRequest(processorBankTransferCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ProcessorBankTransferCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessorBankTransferCreate`: ProcessorBankTransferCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ProcessorBankTransferCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessorBankTransferCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processorBankTransferCreateRequest** | [**ProcessorBankTransferCreateRequest**](ProcessorBankTransferCreateRequest.md) |  | 

### Return type

[**ProcessorBankTransferCreateResponse**](ProcessorBankTransferCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessorIdentityGet

> ProcessorIdentityGetResponse ProcessorIdentityGet(ctx).ProcessorIdentityGetRequest(processorIdentityGetRequest).Execute()

Retrieve Identity data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    processorIdentityGetRequest := *openapiclient.NewProcessorIdentityGetRequest("ProcessorToken_example") // ProcessorIdentityGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.ProcessorIdentityGet(context.Background()).ProcessorIdentityGetRequest(processorIdentityGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ProcessorIdentityGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessorIdentityGet`: ProcessorIdentityGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ProcessorIdentityGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessorIdentityGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processorIdentityGetRequest** | [**ProcessorIdentityGetRequest**](ProcessorIdentityGetRequest.md) |  | 

### Return type

[**ProcessorIdentityGetResponse**](ProcessorIdentityGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessorStripeBankAccountTokenCreate

> ProcessorStripeBankAccountTokenCreateResponse ProcessorStripeBankAccountTokenCreate(ctx).ProcessorStripeBankAccountTokenCreateRequest(processorStripeBankAccountTokenCreateRequest).Execute()

Create Stripe bank account token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    processorStripeBankAccountTokenCreateRequest := *openapiclient.NewProcessorStripeBankAccountTokenCreateRequest("AccessToken_example", "AccountId_example") // ProcessorStripeBankAccountTokenCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.ProcessorStripeBankAccountTokenCreate(context.Background()).ProcessorStripeBankAccountTokenCreateRequest(processorStripeBankAccountTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ProcessorStripeBankAccountTokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessorStripeBankAccountTokenCreate`: ProcessorStripeBankAccountTokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ProcessorStripeBankAccountTokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessorStripeBankAccountTokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processorStripeBankAccountTokenCreateRequest** | [**ProcessorStripeBankAccountTokenCreateRequest**](ProcessorStripeBankAccountTokenCreateRequest.md) |  | 

### Return type

[**ProcessorStripeBankAccountTokenCreateResponse**](ProcessorStripeBankAccountTokenCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ProcessorTokenCreate

> ProcessorTokenCreateResponse ProcessorTokenCreate(ctx).ProcessorTokenCreateRequest(processorTokenCreateRequest).Execute()

Create processor token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    processorTokenCreateRequest := *openapiclient.NewProcessorTokenCreateRequest("AccessToken_example", "AccountId_example", "Processor_example") // ProcessorTokenCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.ProcessorTokenCreate(context.Background()).ProcessorTokenCreateRequest(processorTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.ProcessorTokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ProcessorTokenCreate`: ProcessorTokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.ProcessorTokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiProcessorTokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **processorTokenCreateRequest** | [**ProcessorTokenCreateRequest**](ProcessorTokenCreateRequest.md) |  | 

### Return type

[**ProcessorTokenCreateResponse**](ProcessorTokenCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxBankTransferFireWebhook

> SandboxBankTransferFireWebhookResponse SandboxBankTransferFireWebhook(ctx).SandboxBankTransferFireWebhookRequest(sandboxBankTransferFireWebhookRequest).Execute()

Manually fire a Bank Transfer webhook



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sandboxBankTransferFireWebhookRequest := *openapiclient.NewSandboxBankTransferFireWebhookRequest("Webhook_example") // SandboxBankTransferFireWebhookRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.SandboxBankTransferFireWebhook(context.Background()).SandboxBankTransferFireWebhookRequest(sandboxBankTransferFireWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.SandboxBankTransferFireWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SandboxBankTransferFireWebhook`: SandboxBankTransferFireWebhookResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.SandboxBankTransferFireWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxBankTransferFireWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxBankTransferFireWebhookRequest** | [**SandboxBankTransferFireWebhookRequest**](SandboxBankTransferFireWebhookRequest.md) |  | 

### Return type

[**SandboxBankTransferFireWebhookResponse**](SandboxBankTransferFireWebhookResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxBankTransferSimulate

> SandboxBankTransferSimulateResponse SandboxBankTransferSimulate(ctx).SandboxBankTransferSimulateRequest(sandboxBankTransferSimulateRequest).Execute()

Simulate a bank transfer event in Sandbox



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sandboxBankTransferSimulateRequest := *openapiclient.NewSandboxBankTransferSimulateRequest("BankTransferId_example", "EventType_example") // SandboxBankTransferSimulateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.SandboxBankTransferSimulate(context.Background()).SandboxBankTransferSimulateRequest(sandboxBankTransferSimulateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.SandboxBankTransferSimulate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SandboxBankTransferSimulate`: SandboxBankTransferSimulateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.SandboxBankTransferSimulate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxBankTransferSimulateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxBankTransferSimulateRequest** | [**SandboxBankTransferSimulateRequest**](SandboxBankTransferSimulateRequest.md) |  | 

### Return type

[**SandboxBankTransferSimulateResponse**](SandboxBankTransferSimulateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxIncomeFireWebhook

> SandboxIncomeFireWebhookResponse SandboxIncomeFireWebhook(ctx).SandboxIncomeFireWebhookRequest(sandboxIncomeFireWebhookRequest).Execute()

Manually fire an Income webhook



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sandboxIncomeFireWebhookRequest := *openapiclient.NewSandboxIncomeFireWebhookRequest("IncomeVerificationId_example", "ItemId_example", "Webhook_example", "VerificationStatus_example") // SandboxIncomeFireWebhookRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.SandboxIncomeFireWebhook(context.Background()).SandboxIncomeFireWebhookRequest(sandboxIncomeFireWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.SandboxIncomeFireWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SandboxIncomeFireWebhook`: SandboxIncomeFireWebhookResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.SandboxIncomeFireWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxIncomeFireWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxIncomeFireWebhookRequest** | [**SandboxIncomeFireWebhookRequest**](SandboxIncomeFireWebhookRequest.md) |  | 

### Return type

[**SandboxIncomeFireWebhookResponse**](SandboxIncomeFireWebhookResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxItemFireWebhook

> SandboxItemFireWebhookResponse SandboxItemFireWebhook(ctx).SandboxItemFireWebhookRequest(sandboxItemFireWebhookRequest).Execute()

Fire a test webhook



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sandboxItemFireWebhookRequest := *openapiclient.NewSandboxItemFireWebhookRequest("AccessToken_example", "WebhookCode_example") // SandboxItemFireWebhookRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.SandboxItemFireWebhook(context.Background()).SandboxItemFireWebhookRequest(sandboxItemFireWebhookRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.SandboxItemFireWebhook``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SandboxItemFireWebhook`: SandboxItemFireWebhookResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.SandboxItemFireWebhook`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxItemFireWebhookRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxItemFireWebhookRequest** | [**SandboxItemFireWebhookRequest**](SandboxItemFireWebhookRequest.md) |  | 

### Return type

[**SandboxItemFireWebhookResponse**](SandboxItemFireWebhookResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxItemResetLogin

> SandboxItemResetLoginResponse SandboxItemResetLogin(ctx).SandboxItemResetLoginRequest(sandboxItemResetLoginRequest).Execute()

Force a Sandbox Item into an error state



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sandboxItemResetLoginRequest := *openapiclient.NewSandboxItemResetLoginRequest("AccessToken_example") // SandboxItemResetLoginRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.SandboxItemResetLogin(context.Background()).SandboxItemResetLoginRequest(sandboxItemResetLoginRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.SandboxItemResetLogin``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SandboxItemResetLogin`: SandboxItemResetLoginResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.SandboxItemResetLogin`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxItemResetLoginRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxItemResetLoginRequest** | [**SandboxItemResetLoginRequest**](SandboxItemResetLoginRequest.md) |  | 

### Return type

[**SandboxItemResetLoginResponse**](SandboxItemResetLoginResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxItemSetVerificationStatus

> SandboxItemSetVerificationStatusResponse SandboxItemSetVerificationStatus(ctx).SandboxItemSetVerificationStatusRequest(sandboxItemSetVerificationStatusRequest).Execute()

Set verification status for Sandbox account



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sandboxItemSetVerificationStatusRequest := *openapiclient.NewSandboxItemSetVerificationStatusRequest("AccessToken_example", "AccountId_example", "VerificationStatus_example") // SandboxItemSetVerificationStatusRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.SandboxItemSetVerificationStatus(context.Background()).SandboxItemSetVerificationStatusRequest(sandboxItemSetVerificationStatusRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.SandboxItemSetVerificationStatus``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SandboxItemSetVerificationStatus`: SandboxItemSetVerificationStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.SandboxItemSetVerificationStatus`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxItemSetVerificationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxItemSetVerificationStatusRequest** | [**SandboxItemSetVerificationStatusRequest**](SandboxItemSetVerificationStatusRequest.md) |  | 

### Return type

[**SandboxItemSetVerificationStatusResponse**](SandboxItemSetVerificationStatusResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxOauthSelectAccounts

> map[string]interface{} SandboxOauthSelectAccounts(ctx).SandboxOauthSelectAccountsRequest(sandboxOauthSelectAccountsRequest).Execute()

Save the selected accounts when connecting to the Platypus Oauth institution



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sandboxOauthSelectAccountsRequest := *openapiclient.NewSandboxOauthSelectAccountsRequest("OauthStateId_example", []string{"Accounts_example"}) // SandboxOauthSelectAccountsRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.SandboxOauthSelectAccounts(context.Background()).SandboxOauthSelectAccountsRequest(sandboxOauthSelectAccountsRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.SandboxOauthSelectAccounts``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SandboxOauthSelectAccounts`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.SandboxOauthSelectAccounts`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxOauthSelectAccountsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxOauthSelectAccountsRequest** | [**SandboxOauthSelectAccountsRequest**](SandboxOauthSelectAccountsRequest.md) |  | 

### Return type

**map[string]interface{}**

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxProcessorTokenCreate

> SandboxProcessorTokenCreateResponse SandboxProcessorTokenCreate(ctx).SandboxProcessorTokenCreateRequest(sandboxProcessorTokenCreateRequest).Execute()

Create a test Item and processor token



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sandboxProcessorTokenCreateRequest := *openapiclient.NewSandboxProcessorTokenCreateRequest("InstitutionId_example") // SandboxProcessorTokenCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.SandboxProcessorTokenCreate(context.Background()).SandboxProcessorTokenCreateRequest(sandboxProcessorTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.SandboxProcessorTokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SandboxProcessorTokenCreate`: SandboxProcessorTokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.SandboxProcessorTokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxProcessorTokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxProcessorTokenCreateRequest** | [**SandboxProcessorTokenCreateRequest**](SandboxProcessorTokenCreateRequest.md) |  | 

### Return type

[**SandboxProcessorTokenCreateResponse**](SandboxProcessorTokenCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxPublicTokenCreate

> SandboxPublicTokenCreateResponse SandboxPublicTokenCreate(ctx).SandboxPublicTokenCreateRequest(sandboxPublicTokenCreateRequest).Execute()

Create a test Item



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sandboxPublicTokenCreateRequest := *openapiclient.NewSandboxPublicTokenCreateRequest("InstitutionId_example", []openapiclient.Products{openapiclient.Products("assets")}) // SandboxPublicTokenCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.SandboxPublicTokenCreate(context.Background()).SandboxPublicTokenCreateRequest(sandboxPublicTokenCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.SandboxPublicTokenCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SandboxPublicTokenCreate`: SandboxPublicTokenCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.SandboxPublicTokenCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxPublicTokenCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxPublicTokenCreateRequest** | [**SandboxPublicTokenCreateRequest**](SandboxPublicTokenCreateRequest.md) |  | 

### Return type

[**SandboxPublicTokenCreateResponse**](SandboxPublicTokenCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxTransferRepaymentSimulate

> SandboxTransferRepaymentSimulateResponse SandboxTransferRepaymentSimulate(ctx).SandboxTransferRepaymentSimulateRequest(sandboxTransferRepaymentSimulateRequest).Execute()

Trigger the creation of a repayment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sandboxTransferRepaymentSimulateRequest := *openapiclient.NewSandboxTransferRepaymentSimulateRequest() // SandboxTransferRepaymentSimulateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.SandboxTransferRepaymentSimulate(context.Background()).SandboxTransferRepaymentSimulateRequest(sandboxTransferRepaymentSimulateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.SandboxTransferRepaymentSimulate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SandboxTransferRepaymentSimulate`: SandboxTransferRepaymentSimulateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.SandboxTransferRepaymentSimulate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxTransferRepaymentSimulateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxTransferRepaymentSimulateRequest** | [**SandboxTransferRepaymentSimulateRequest**](SandboxTransferRepaymentSimulateRequest.md) |  | 

### Return type

[**SandboxTransferRepaymentSimulateResponse**](SandboxTransferRepaymentSimulateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxTransferSimulate

> SandboxTransferSimulateResponse SandboxTransferSimulate(ctx).SandboxTransferSimulateRequest(sandboxTransferSimulateRequest).Execute()

Simulate a transfer event in Sandbox



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sandboxTransferSimulateRequest := *openapiclient.NewSandboxTransferSimulateRequest("TransferId_example", "EventType_example") // SandboxTransferSimulateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.SandboxTransferSimulate(context.Background()).SandboxTransferSimulateRequest(sandboxTransferSimulateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.SandboxTransferSimulate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SandboxTransferSimulate`: SandboxTransferSimulateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.SandboxTransferSimulate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxTransferSimulateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxTransferSimulateRequest** | [**SandboxTransferSimulateRequest**](SandboxTransferSimulateRequest.md) |  | 

### Return type

[**SandboxTransferSimulateResponse**](SandboxTransferSimulateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SandboxTransferSweepSimulate

> SandboxTransferSweepSimulateResponse SandboxTransferSweepSimulate(ctx).SandboxTransferSweepSimulateRequest(sandboxTransferSweepSimulateRequest).Execute()

Simulate creating a sweep



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    sandboxTransferSweepSimulateRequest := *openapiclient.NewSandboxTransferSweepSimulateRequest() // SandboxTransferSweepSimulateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.SandboxTransferSweepSimulate(context.Background()).SandboxTransferSweepSimulateRequest(sandboxTransferSweepSimulateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.SandboxTransferSweepSimulate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SandboxTransferSweepSimulate`: SandboxTransferSweepSimulateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.SandboxTransferSweepSimulate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSandboxTransferSweepSimulateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **sandboxTransferSweepSimulateRequest** | [**SandboxTransferSweepSimulateRequest**](SandboxTransferSweepSimulateRequest.md) |  | 

### Return type

[**SandboxTransferSweepSimulateResponse**](SandboxTransferSweepSimulateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignalDecisionReport

> SignalDecisionReportResponse SignalDecisionReport(ctx).SignalDecisionReportRequest(signalDecisionReportRequest).Execute()

Report whether you initiated an ACH transaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    signalDecisionReportRequest := *openapiclient.NewSignalDecisionReportRequest("ClientTransactionId_example", false) // SignalDecisionReportRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.SignalDecisionReport(context.Background()).SignalDecisionReportRequest(signalDecisionReportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.SignalDecisionReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SignalDecisionReport`: SignalDecisionReportResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.SignalDecisionReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignalDecisionReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signalDecisionReportRequest** | [**SignalDecisionReportRequest**](SignalDecisionReportRequest.md) |  | 

### Return type

[**SignalDecisionReportResponse**](SignalDecisionReportResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignalEvaluate

> SignalEvaluateResponse SignalEvaluate(ctx).SignalEvaluateRequest(signalEvaluateRequest).Execute()

Evaluate a planned ACH transaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    signalEvaluateRequest := *openapiclient.NewSignalEvaluateRequest("AccessToken_example", "AccountId_example", "ClientTransactionId_example", float32(123)) // SignalEvaluateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.SignalEvaluate(context.Background()).SignalEvaluateRequest(signalEvaluateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.SignalEvaluate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SignalEvaluate`: SignalEvaluateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.SignalEvaluate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignalEvaluateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signalEvaluateRequest** | [**SignalEvaluateRequest**](SignalEvaluateRequest.md) |  | 

### Return type

[**SignalEvaluateResponse**](SignalEvaluateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SignalReturnReport

> SignalReturnReportResponse SignalReturnReport(ctx).SignalReturnReportRequest(signalReturnReportRequest).Execute()

Report a return for an ACH transaction



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    signalReturnReportRequest := *openapiclient.NewSignalReturnReportRequest("ClientTransactionId_example", "ReturnCode_example") // SignalReturnReportRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.SignalReturnReport(context.Background()).SignalReturnReportRequest(signalReturnReportRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.SignalReturnReport``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SignalReturnReport`: SignalReturnReportResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.SignalReturnReport`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSignalReturnReportRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **signalReturnReportRequest** | [**SignalReturnReportRequest**](SignalReturnReportRequest.md) |  | 

### Return type

[**SignalReturnReportResponse**](SignalReturnReportResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsGet

> TransactionsGetResponse TransactionsGet(ctx).TransactionsGetRequest(transactionsGetRequest).Execute()

Get transaction data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    "time"
    openapiclient "./openapi"
)

func main() {
    transactionsGetRequest := *openapiclient.NewTransactionsGetRequest("AccessToken_example", time.Now(), time.Now()) // TransactionsGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.TransactionsGet(context.Background()).TransactionsGetRequest(transactionsGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.TransactionsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransactionsGet`: TransactionsGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.TransactionsGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionsGetRequest** | [**TransactionsGetRequest**](TransactionsGetRequest.md) |  | 

### Return type

[**TransactionsGetResponse**](TransactionsGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsRecurringGet

> TransactionsRecurringGetResponse TransactionsRecurringGet(ctx).TransactionsRecurringGetRequest(transactionsRecurringGetRequest).Execute()

Get streams of recurring transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transactionsRecurringGetRequest := *openapiclient.NewTransactionsRecurringGetRequest("AccessToken_example", []string{"AccountIds_example"}) // TransactionsRecurringGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.TransactionsRecurringGet(context.Background()).TransactionsRecurringGetRequest(transactionsRecurringGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.TransactionsRecurringGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransactionsRecurringGet`: TransactionsRecurringGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.TransactionsRecurringGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsRecurringGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionsRecurringGetRequest** | [**TransactionsRecurringGetRequest**](TransactionsRecurringGetRequest.md) |  | 

### Return type

[**TransactionsRecurringGetResponse**](TransactionsRecurringGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsRefresh

> TransactionsRefreshResponse TransactionsRefresh(ctx).TransactionsRefreshRequest(transactionsRefreshRequest).Execute()

Refresh transaction data



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transactionsRefreshRequest := *openapiclient.NewTransactionsRefreshRequest("AccessToken_example") // TransactionsRefreshRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.TransactionsRefresh(context.Background()).TransactionsRefreshRequest(transactionsRefreshRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.TransactionsRefresh``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransactionsRefresh`: TransactionsRefreshResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.TransactionsRefresh`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsRefreshRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionsRefreshRequest** | [**TransactionsRefreshRequest**](TransactionsRefreshRequest.md) |  | 

### Return type

[**TransactionsRefreshResponse**](TransactionsRefreshResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransactionsSync

> TransactionsSyncResponse TransactionsSync(ctx).TransactionsSyncRequest(transactionsSyncRequest).Execute()

Get incremental transaction updates on an Item



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transactionsSyncRequest := *openapiclient.NewTransactionsSyncRequest("AccessToken_example") // TransactionsSyncRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.TransactionsSync(context.Background()).TransactionsSyncRequest(transactionsSyncRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.TransactionsSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransactionsSync`: TransactionsSyncResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.TransactionsSync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransactionsSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transactionsSyncRequest** | [**TransactionsSyncRequest**](TransactionsSyncRequest.md) |  | 

### Return type

[**TransactionsSyncResponse**](TransactionsSyncResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferAuthorizationCreate

> TransferAuthorizationCreateResponse TransferAuthorizationCreate(ctx).TransferAuthorizationCreateRequest(transferAuthorizationCreateRequest).Execute()

Create a transfer authorization



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transferAuthorizationCreateRequest := *openapiclient.NewTransferAuthorizationCreateRequest("AccessToken_example", "AccountId_example", openapiclient.TransferType("debit"), openapiclient.TransferNetwork("ach"), "Amount_example", openapiclient.ACHClass("arc"), *openapiclient.NewTransferUserInRequest("LegalName_example")) // TransferAuthorizationCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.TransferAuthorizationCreate(context.Background()).TransferAuthorizationCreateRequest(transferAuthorizationCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.TransferAuthorizationCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferAuthorizationCreate`: TransferAuthorizationCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.TransferAuthorizationCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferAuthorizationCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferAuthorizationCreateRequest** | [**TransferAuthorizationCreateRequest**](TransferAuthorizationCreateRequest.md) |  | 

### Return type

[**TransferAuthorizationCreateResponse**](TransferAuthorizationCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferCancel

> TransferCancelResponse TransferCancel(ctx).TransferCancelRequest(transferCancelRequest).Execute()

Cancel a transfer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transferCancelRequest := *openapiclient.NewTransferCancelRequest("TransferId_example") // TransferCancelRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.TransferCancel(context.Background()).TransferCancelRequest(transferCancelRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.TransferCancel``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferCancel`: TransferCancelResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.TransferCancel`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferCancelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferCancelRequest** | [**TransferCancelRequest**](TransferCancelRequest.md) |  | 

### Return type

[**TransferCancelResponse**](TransferCancelResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferCreate

> TransferCreateResponse TransferCreate(ctx).TransferCreateRequest(transferCreateRequest).Execute()

Create a transfer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transferCreateRequest := *openapiclient.NewTransferCreateRequest("AccessToken_example", "AccountId_example", "AuthorizationId_example", openapiclient.TransferType("debit"), openapiclient.TransferNetwork("ach"), "Amount_example", "Description_example", openapiclient.ACHClass("arc"), *openapiclient.NewTransferUserInRequest("LegalName_example")) // TransferCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.TransferCreate(context.Background()).TransferCreateRequest(transferCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.TransferCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferCreate`: TransferCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.TransferCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferCreateRequest** | [**TransferCreateRequest**](TransferCreateRequest.md) |  | 

### Return type

[**TransferCreateResponse**](TransferCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferEventList

> TransferEventListResponse TransferEventList(ctx).TransferEventListRequest(transferEventListRequest).Execute()

List transfer events



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transferEventListRequest := *openapiclient.NewTransferEventListRequest() // TransferEventListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.TransferEventList(context.Background()).TransferEventListRequest(transferEventListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.TransferEventList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferEventList`: TransferEventListResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.TransferEventList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferEventListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferEventListRequest** | [**TransferEventListRequest**](TransferEventListRequest.md) |  | 

### Return type

[**TransferEventListResponse**](TransferEventListResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferEventSync

> TransferEventSyncResponse TransferEventSync(ctx).TransferEventSyncRequest(transferEventSyncRequest).Execute()

Sync transfer events



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transferEventSyncRequest := *openapiclient.NewTransferEventSyncRequest(int32(123)) // TransferEventSyncRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.TransferEventSync(context.Background()).TransferEventSyncRequest(transferEventSyncRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.TransferEventSync``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferEventSync`: TransferEventSyncResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.TransferEventSync`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferEventSyncRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferEventSyncRequest** | [**TransferEventSyncRequest**](TransferEventSyncRequest.md) |  | 

### Return type

[**TransferEventSyncResponse**](TransferEventSyncResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferGet

> TransferGetResponse TransferGet(ctx).TransferGetRequest(transferGetRequest).Execute()

Retrieve a transfer



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transferGetRequest := *openapiclient.NewTransferGetRequest("TransferId_example") // TransferGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.TransferGet(context.Background()).TransferGetRequest(transferGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.TransferGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferGet`: TransferGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.TransferGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferGetRequest** | [**TransferGetRequest**](TransferGetRequest.md) |  | 

### Return type

[**TransferGetResponse**](TransferGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferIntentCreate

> TransferIntentCreateResponse TransferIntentCreate(ctx).TransferIntentCreateRequest(transferIntentCreateRequest).Execute()

Create a transfer intent object to invoke the Transfer UI



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transferIntentCreateRequest := *openapiclient.NewTransferIntentCreateRequest("ClientId_example", "Secret_example", openapiclient.TransferIntentCreateMode("PAYMENT"), "Amount_example", "Description_example", openapiclient.ACHClass("arc"), *openapiclient.NewTransferUserInRequest("LegalName_example")) // TransferIntentCreateRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.TransferIntentCreate(context.Background()).TransferIntentCreateRequest(transferIntentCreateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.TransferIntentCreate``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferIntentCreate`: TransferIntentCreateResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.TransferIntentCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferIntentCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferIntentCreateRequest** | [**TransferIntentCreateRequest**](TransferIntentCreateRequest.md) |  | 

### Return type

[**TransferIntentCreateResponse**](TransferIntentCreateResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferIntentGet

> TransferIntentGetResponse TransferIntentGet(ctx).RequestBody(requestBody).Execute()

Retrieve more information about a transfer intent



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    requestBody := map[string]interface{}{"key": interface{}(123)} // map[string]interface{} | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.TransferIntentGet(context.Background()).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.TransferIntentGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferIntentGet`: TransferIntentGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.TransferIntentGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferIntentGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **requestBody** | **map[string]interface{}** |  | 

### Return type

[**TransferIntentGetResponse**](TransferIntentGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferList

> TransferListResponse TransferList(ctx).TransferListRequest(transferListRequest).Execute()

List transfers



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transferListRequest := *openapiclient.NewTransferListRequest() // TransferListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.TransferList(context.Background()).TransferListRequest(transferListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.TransferList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferList`: TransferListResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.TransferList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferListRequest** | [**TransferListRequest**](TransferListRequest.md) |  | 

### Return type

[**TransferListResponse**](TransferListResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferRepaymentList

> TransferRepaymentListResponse TransferRepaymentList(ctx).TransferRepaymentListRequest(transferRepaymentListRequest).Execute()

Lists historical repayments



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transferRepaymentListRequest := *openapiclient.NewTransferRepaymentListRequest() // TransferRepaymentListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.TransferRepaymentList(context.Background()).TransferRepaymentListRequest(transferRepaymentListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.TransferRepaymentList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferRepaymentList`: TransferRepaymentListResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.TransferRepaymentList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferRepaymentListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferRepaymentListRequest** | [**TransferRepaymentListRequest**](TransferRepaymentListRequest.md) |  | 

### Return type

[**TransferRepaymentListResponse**](TransferRepaymentListResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferRepaymentReturnList

> TransferRepaymentReturnListResponse TransferRepaymentReturnList(ctx).TransferRepaymentReturnListRequest(transferRepaymentReturnListRequest).Execute()

List the returns included in a repayment



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transferRepaymentReturnListRequest := *openapiclient.NewTransferRepaymentReturnListRequest("RepaymentId_example") // TransferRepaymentReturnListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.TransferRepaymentReturnList(context.Background()).TransferRepaymentReturnListRequest(transferRepaymentReturnListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.TransferRepaymentReturnList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferRepaymentReturnList`: TransferRepaymentReturnListResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.TransferRepaymentReturnList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferRepaymentReturnListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferRepaymentReturnListRequest** | [**TransferRepaymentReturnListRequest**](TransferRepaymentReturnListRequest.md) |  | 

### Return type

[**TransferRepaymentReturnListResponse**](TransferRepaymentReturnListResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferSweepGet

> TransferSweepGetResponse TransferSweepGet(ctx).TransferSweepGetRequest(transferSweepGetRequest).Execute()

Retrieve a sweep



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transferSweepGetRequest := *openapiclient.NewTransferSweepGetRequest("SweepId_example") // TransferSweepGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.TransferSweepGet(context.Background()).TransferSweepGetRequest(transferSweepGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.TransferSweepGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferSweepGet`: TransferSweepGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.TransferSweepGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferSweepGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferSweepGetRequest** | [**TransferSweepGetRequest**](TransferSweepGetRequest.md) |  | 

### Return type

[**TransferSweepGetResponse**](TransferSweepGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TransferSweepList

> TransferSweepListResponse TransferSweepList(ctx).TransferSweepListRequest(transferSweepListRequest).Execute()

List sweeps



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    transferSweepListRequest := *openapiclient.NewTransferSweepListRequest() // TransferSweepListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.TransferSweepList(context.Background()).TransferSweepListRequest(transferSweepListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.TransferSweepList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TransferSweepList`: TransferSweepListResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.TransferSweepList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTransferSweepListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **transferSweepListRequest** | [**TransferSweepListRequest**](TransferSweepListRequest.md) |  | 

### Return type

[**TransferSweepListResponse**](TransferSweepListResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletGet

> WalletGetResponse WalletGet(ctx).WalletGetRequest(walletGetRequest).Execute()

Fetch an e-wallet



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    walletGetRequest := *openapiclient.NewWalletGetRequest("WalletId_example") // WalletGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.WalletGet(context.Background()).WalletGetRequest(walletGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.WalletGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WalletGet`: WalletGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.WalletGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletGetRequest** | [**WalletGetRequest**](WalletGetRequest.md) |  | 

### Return type

[**WalletGetResponse**](WalletGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletTransactionExecute

> WalletTransactionExecuteResponse WalletTransactionExecute(ctx).WalletTransactionExecuteRequest(walletTransactionExecuteRequest).Execute()

Execute a transaction using an e-wallet



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    walletTransactionExecuteRequest := *openapiclient.NewWalletTransactionExecuteRequest("IdempotencyKey_example", "WalletId_example", *openapiclient.NewWalletTransactionCounterparty("Name_example", *openapiclient.NewWalletTransactionCounterpartyNumbers(*openapiclient.NewWalletTransactionCounterpartyBACS())), *openapiclient.NewWalletTransactionAmount("IsoCurrencyCode_example", float32(123)), "Reference_example") // WalletTransactionExecuteRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.WalletTransactionExecute(context.Background()).WalletTransactionExecuteRequest(walletTransactionExecuteRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.WalletTransactionExecute``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WalletTransactionExecute`: WalletTransactionExecuteResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.WalletTransactionExecute`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletTransactionExecuteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletTransactionExecuteRequest** | [**WalletTransactionExecuteRequest**](WalletTransactionExecuteRequest.md) |  | 

### Return type

[**WalletTransactionExecuteResponse**](WalletTransactionExecuteResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WalletTransactionsList

> WalletTransactionsListResponse WalletTransactionsList(ctx).WalletTransactionsListRequest(walletTransactionsListRequest).Execute()

List e-wallet transactions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    walletTransactionsListRequest := *openapiclient.NewWalletTransactionsListRequest("WalletId_example") // WalletTransactionsListRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.WalletTransactionsList(context.Background()).WalletTransactionsListRequest(walletTransactionsListRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.WalletTransactionsList``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WalletTransactionsList`: WalletTransactionsListResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.WalletTransactionsList`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWalletTransactionsListRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **walletTransactionsListRequest** | [**WalletTransactionsListRequest**](WalletTransactionsListRequest.md) |  | 

### Return type

[**WalletTransactionsListResponse**](WalletTransactionsListResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## WebhookVerificationKeyGet

> WebhookVerificationKeyGetResponse WebhookVerificationKeyGet(ctx).WebhookVerificationKeyGetRequest(webhookVerificationKeyGetRequest).Execute()

Get webhook verification key



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    webhookVerificationKeyGetRequest := *openapiclient.NewWebhookVerificationKeyGetRequest("KeyId_example") // WebhookVerificationKeyGetRequest | 

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlaidApi.WebhookVerificationKeyGet(context.Background()).WebhookVerificationKeyGetRequest(webhookVerificationKeyGetRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlaidApi.WebhookVerificationKeyGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `WebhookVerificationKeyGet`: WebhookVerificationKeyGetResponse
    fmt.Fprintf(os.Stdout, "Response from `PlaidApi.WebhookVerificationKeyGet`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiWebhookVerificationKeyGetRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **webhookVerificationKeyGetRequest** | [**WebhookVerificationKeyGetRequest**](WebhookVerificationKeyGetRequest.md) |  | 

### Return type

[**WebhookVerificationKeyGetResponse**](WebhookVerificationKeyGetResponse.md)

### Authorization

[clientId](../README.md#clientId), [plaidVersion](../README.md#plaidVersion), [secret](../README.md#secret)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

