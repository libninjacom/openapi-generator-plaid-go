/*
The Plaid API

The Plaid REST API. Please see https://plaid.com/docs/api for more details.

API version: 2020-09-14_1.64.13
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// AccountSubtype See the [Account type schema](https://plaid.com/docs/api/accounts/#account-type-schema) for a full listing of account types and corresponding subtypes.
type AccountSubtype string

// List of AccountSubtype
const (
	_401A AccountSubtype = "401a"
	_401K AccountSubtype = "401k"
	_403_B AccountSubtype = "403B"
	_457B AccountSubtype = "457b"
	_529 AccountSubtype = "529"
	BROKERAGE AccountSubtype = "brokerage"
	CASH_ISA AccountSubtype = "cash isa"
	EDUCATION_SAVINGS_ACCOUNT AccountSubtype = "education savings account"
	EBT AccountSubtype = "ebt"
	FIXED_ANNUITY AccountSubtype = "fixed annuity"
	GIC AccountSubtype = "gic"
	HEALTH_REIMBURSEMENT_ARRANGEMENT AccountSubtype = "health reimbursement arrangement"
	HSA AccountSubtype = "hsa"
	ISA AccountSubtype = "isa"
	IRA AccountSubtype = "ira"
	LIF AccountSubtype = "lif"
	LIFE_INSURANCE AccountSubtype = "life insurance"
	LIRA AccountSubtype = "lira"
	LRIF AccountSubtype = "lrif"
	LRSP AccountSubtype = "lrsp"
	NON_TAXABLE_BROKERAGE_ACCOUNT AccountSubtype = "non-taxable brokerage account"
	OTHER AccountSubtype = "other"
	OTHER_INSURANCE AccountSubtype = "other insurance"
	OTHER_ANNUITY AccountSubtype = "other annuity"
	PRIF AccountSubtype = "prif"
	RDSP AccountSubtype = "rdsp"
	RESP AccountSubtype = "resp"
	RLIF AccountSubtype = "rlif"
	RRIF AccountSubtype = "rrif"
	PENSION AccountSubtype = "pension"
	PROFIT_SHARING_PLAN AccountSubtype = "profit sharing plan"
	RETIREMENT AccountSubtype = "retirement"
	ROTH AccountSubtype = "roth"
	ROTH_401K AccountSubtype = "roth 401k"
	RRSP AccountSubtype = "rrsp"
	SEP_IRA AccountSubtype = "sep ira"
	SIMPLE_IRA AccountSubtype = "simple ira"
	SIPP AccountSubtype = "sipp"
	STOCK_PLAN AccountSubtype = "stock plan"
	THRIFT_SAVINGS_PLAN AccountSubtype = "thrift savings plan"
	TFSA AccountSubtype = "tfsa"
	TRUST AccountSubtype = "trust"
	UGMA AccountSubtype = "ugma"
	UTMA AccountSubtype = "utma"
	VARIABLE_ANNUITY AccountSubtype = "variable annuity"
	CREDIT_CARD AccountSubtype = "credit card"
	PAYPAL AccountSubtype = "paypal"
	CD AccountSubtype = "cd"
	CHECKING AccountSubtype = "checking"
	SAVINGS AccountSubtype = "savings"
	MONEY_MARKET AccountSubtype = "money market"
	PREPAID AccountSubtype = "prepaid"
	AUTO AccountSubtype = "auto"
	BUSINESS AccountSubtype = "business"
	COMMERCIAL AccountSubtype = "commercial"
	CONSTRUCTION AccountSubtype = "construction"
	CONSUMER AccountSubtype = "consumer"
	HOME_EQUITY AccountSubtype = "home equity"
	LOAN AccountSubtype = "loan"
	MORTGAGE AccountSubtype = "mortgage"
	OVERDRAFT AccountSubtype = "overdraft"
	LINE_OF_CREDIT AccountSubtype = "line of credit"
	STUDENT AccountSubtype = "student"
	CASH_MANAGEMENT AccountSubtype = "cash management"
	KEOGH AccountSubtype = "keogh"
	MUTUAL_FUND AccountSubtype = "mutual fund"
	RECURRING AccountSubtype = "recurring"
	REWARDS AccountSubtype = "rewards"
	SAFE_DEPOSIT AccountSubtype = "safe deposit"
	SARSEP AccountSubtype = "sarsep"
	PAYROLL AccountSubtype = "payroll"
	NULL AccountSubtype = "null"
)

// All allowed values of AccountSubtype enum
var AllowedAccountSubtypeEnumValues = []AccountSubtype{
	"401a",
	"401k",
	"403B",
	"457b",
	"529",
	"brokerage",
	"cash isa",
	"education savings account",
	"ebt",
	"fixed annuity",
	"gic",
	"health reimbursement arrangement",
	"hsa",
	"isa",
	"ira",
	"lif",
	"life insurance",
	"lira",
	"lrif",
	"lrsp",
	"non-taxable brokerage account",
	"other",
	"other insurance",
	"other annuity",
	"prif",
	"rdsp",
	"resp",
	"rlif",
	"rrif",
	"pension",
	"profit sharing plan",
	"retirement",
	"roth",
	"roth 401k",
	"rrsp",
	"sep ira",
	"simple ira",
	"sipp",
	"stock plan",
	"thrift savings plan",
	"tfsa",
	"trust",
	"ugma",
	"utma",
	"variable annuity",
	"credit card",
	"paypal",
	"cd",
	"checking",
	"savings",
	"money market",
	"prepaid",
	"auto",
	"business",
	"commercial",
	"construction",
	"consumer",
	"home equity",
	"loan",
	"mortgage",
	"overdraft",
	"line of credit",
	"student",
	"cash management",
	"keogh",
	"mutual fund",
	"recurring",
	"rewards",
	"safe deposit",
	"sarsep",
	"payroll",
	"null",
}

func (v *AccountSubtype) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AccountSubtype(value)
	for _, existing := range AllowedAccountSubtypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AccountSubtype", value)
}

// NewAccountSubtypeFromValue returns a pointer to a valid AccountSubtype
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAccountSubtypeFromValue(v string) (*AccountSubtype, error) {
	ev := AccountSubtype(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AccountSubtype: valid values are %v", v, AllowedAccountSubtypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AccountSubtype) IsValid() bool {
	for _, existing := range AllowedAccountSubtypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AccountSubtype value
func (v AccountSubtype) Ptr() *AccountSubtype {
	return &v
}

type NullableAccountSubtype struct {
	value *AccountSubtype
	isSet bool
}

func (v NullableAccountSubtype) Get() *AccountSubtype {
	return v.value
}

func (v *NullableAccountSubtype) Set(val *AccountSubtype) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountSubtype) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountSubtype) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountSubtype(val *AccountSubtype) *NullableAccountSubtype {
	return &NullableAccountSubtype{value: val, isSet: true}
}

func (v NullableAccountSubtype) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountSubtype) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
