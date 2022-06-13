# JWKPublicKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alg** | **string** | The alg member identifies the cryptographic algorithm family used with the key. | 
**Crv** | **string** | The crv member identifies the cryptographic curve used with the key. | 
**Kid** | **string** | The kid (Key ID) member can be used to match a specific key. This can be used, for instance, to choose among a set of keys within the JWK during key rollover. | 
**Kty** | **string** | The kty (key type) parameter identifies the cryptographic algorithm family used with the key, such as RSA or EC. | 
**Use** | **string** | The use (public key use) parameter identifies the intended use of the public key. | 
**X** | **string** | The x member contains the x coordinate for the elliptic curve point. | 
**Y** | **string** | The y member contains the y coordinate for the elliptic curve point. | 
**CreatedAt** | **int32** | The timestamp when the key was created, in Unix time. | 
**ExpiredAt** | **NullableInt32** | The timestamp when the key expired, in Unix time. | 

## Methods

### NewJWKPublicKey

`func NewJWKPublicKey(alg string, crv string, kid string, kty string, use string, x string, y string, createdAt int32, expiredAt NullableInt32, ) *JWKPublicKey`

NewJWKPublicKey instantiates a new JWKPublicKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewJWKPublicKeyWithDefaults

`func NewJWKPublicKeyWithDefaults() *JWKPublicKey`

NewJWKPublicKeyWithDefaults instantiates a new JWKPublicKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlg

`func (o *JWKPublicKey) GetAlg() string`

GetAlg returns the Alg field if non-nil, zero value otherwise.

### GetAlgOk

`func (o *JWKPublicKey) GetAlgOk() (*string, bool)`

GetAlgOk returns a tuple with the Alg field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlg

`func (o *JWKPublicKey) SetAlg(v string)`

SetAlg sets Alg field to given value.


### GetCrv

`func (o *JWKPublicKey) GetCrv() string`

GetCrv returns the Crv field if non-nil, zero value otherwise.

### GetCrvOk

`func (o *JWKPublicKey) GetCrvOk() (*string, bool)`

GetCrvOk returns a tuple with the Crv field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCrv

`func (o *JWKPublicKey) SetCrv(v string)`

SetCrv sets Crv field to given value.


### GetKid

`func (o *JWKPublicKey) GetKid() string`

GetKid returns the Kid field if non-nil, zero value otherwise.

### GetKidOk

`func (o *JWKPublicKey) GetKidOk() (*string, bool)`

GetKidOk returns a tuple with the Kid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKid

`func (o *JWKPublicKey) SetKid(v string)`

SetKid sets Kid field to given value.


### GetKty

`func (o *JWKPublicKey) GetKty() string`

GetKty returns the Kty field if non-nil, zero value otherwise.

### GetKtyOk

`func (o *JWKPublicKey) GetKtyOk() (*string, bool)`

GetKtyOk returns a tuple with the Kty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKty

`func (o *JWKPublicKey) SetKty(v string)`

SetKty sets Kty field to given value.


### GetUse

`func (o *JWKPublicKey) GetUse() string`

GetUse returns the Use field if non-nil, zero value otherwise.

### GetUseOk

`func (o *JWKPublicKey) GetUseOk() (*string, bool)`

GetUseOk returns a tuple with the Use field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUse

`func (o *JWKPublicKey) SetUse(v string)`

SetUse sets Use field to given value.


### GetX

`func (o *JWKPublicKey) GetX() string`

GetX returns the X field if non-nil, zero value otherwise.

### GetXOk

`func (o *JWKPublicKey) GetXOk() (*string, bool)`

GetXOk returns a tuple with the X field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetX

`func (o *JWKPublicKey) SetX(v string)`

SetX sets X field to given value.


### GetY

`func (o *JWKPublicKey) GetY() string`

GetY returns the Y field if non-nil, zero value otherwise.

### GetYOk

`func (o *JWKPublicKey) GetYOk() (*string, bool)`

GetYOk returns a tuple with the Y field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetY

`func (o *JWKPublicKey) SetY(v string)`

SetY sets Y field to given value.


### GetCreatedAt

`func (o *JWKPublicKey) GetCreatedAt() int32`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *JWKPublicKey) GetCreatedAtOk() (*int32, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *JWKPublicKey) SetCreatedAt(v int32)`

SetCreatedAt sets CreatedAt field to given value.


### GetExpiredAt

`func (o *JWKPublicKey) GetExpiredAt() int32`

GetExpiredAt returns the ExpiredAt field if non-nil, zero value otherwise.

### GetExpiredAtOk

`func (o *JWKPublicKey) GetExpiredAtOk() (*int32, bool)`

GetExpiredAtOk returns a tuple with the ExpiredAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpiredAt

`func (o *JWKPublicKey) SetExpiredAt(v int32)`

SetExpiredAt sets ExpiredAt field to given value.


### SetExpiredAtNil

`func (o *JWKPublicKey) SetExpiredAtNil(b bool)`

 SetExpiredAtNil sets the value for ExpiredAt to be an explicit nil

### UnsetExpiredAt
`func (o *JWKPublicKey) UnsetExpiredAt()`

UnsetExpiredAt ensures that no value is present for ExpiredAt, not even an explicit nil

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


