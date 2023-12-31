/*
Dropbox Sign API

Dropbox Sign v3 API

API version: 3.0.0
Contact: apisupport@hellosign.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// checks if the AccountCreateResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AccountCreateResponse{}

// AccountCreateResponse struct for AccountCreateResponse
type AccountCreateResponse struct {
	Account *AccountResponse `json:"account,omitempty"`
	OauthData *OAuthTokenResponse `json:"oauth_data,omitempty"`
	// A list of warnings.
	Warnings []WarningResponse `json:"warnings,omitempty"`
}

// NewAccountCreateResponse instantiates a new AccountCreateResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountCreateResponse() *AccountCreateResponse {
	this := AccountCreateResponse{}
	return &this
}

// NewAccountCreateResponseWithDefaults instantiates a new AccountCreateResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountCreateResponseWithDefaults() *AccountCreateResponse {
	this := AccountCreateResponse{}
	return &this
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *AccountCreateResponse) GetAccount() AccountResponse {
	if o == nil || IsNil(o.Account) {
		var ret AccountResponse
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateResponse) GetAccountOk() (*AccountResponse, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *AccountCreateResponse) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given AccountResponse and assigns it to the Account field.
func (o *AccountCreateResponse) SetAccount(v AccountResponse) {
	o.Account = &v
}

// GetOauthData returns the OauthData field value if set, zero value otherwise.
func (o *AccountCreateResponse) GetOauthData() OAuthTokenResponse {
	if o == nil || IsNil(o.OauthData) {
		var ret OAuthTokenResponse
		return ret
	}
	return *o.OauthData
}

// GetOauthDataOk returns a tuple with the OauthData field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateResponse) GetOauthDataOk() (*OAuthTokenResponse, bool) {
	if o == nil || IsNil(o.OauthData) {
		return nil, false
	}
	return o.OauthData, true
}

// HasOauthData returns a boolean if a field has been set.
func (o *AccountCreateResponse) HasOauthData() bool {
	if o != nil && !IsNil(o.OauthData) {
		return true
	}

	return false
}

// SetOauthData gets a reference to the given OAuthTokenResponse and assigns it to the OauthData field.
func (o *AccountCreateResponse) SetOauthData(v OAuthTokenResponse) {
	o.OauthData = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *AccountCreateResponse) GetWarnings() []WarningResponse {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningResponse
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountCreateResponse) GetWarningsOk() ([]WarningResponse, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *AccountCreateResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningResponse and assigns it to the Warnings field.
func (o *AccountCreateResponse) SetWarnings(v []WarningResponse) {
	o.Warnings = v
}

func (o AccountCreateResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AccountCreateResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.OauthData) {
		toSerialize["oauth_data"] = o.OauthData
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableAccountCreateResponse struct {
	value *AccountCreateResponse
	isSet bool
}

func (v NullableAccountCreateResponse) Get() *AccountCreateResponse {
	return v.value
}

func (v *NullableAccountCreateResponse) Set(val *AccountCreateResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountCreateResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountCreateResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountCreateResponse(val *AccountCreateResponse) *NullableAccountCreateResponse {
	return &NullableAccountCreateResponse{value: val, isSet: true}
}

func (v NullableAccountCreateResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountCreateResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


