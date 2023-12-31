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

// checks if the ApiAppResponseOwnerAccount type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApiAppResponseOwnerAccount{}

// ApiAppResponseOwnerAccount An object describing the app's owner
type ApiAppResponseOwnerAccount struct {
	// The owner account's ID
	AccountId *string `json:"account_id,omitempty"`
	// The owner account's email address
	EmailAddress *string `json:"email_address,omitempty"`
}

// NewApiAppResponseOwnerAccount instantiates a new ApiAppResponseOwnerAccount object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiAppResponseOwnerAccount() *ApiAppResponseOwnerAccount {
	this := ApiAppResponseOwnerAccount{}
	return &this
}

// NewApiAppResponseOwnerAccountWithDefaults instantiates a new ApiAppResponseOwnerAccount object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiAppResponseOwnerAccountWithDefaults() *ApiAppResponseOwnerAccount {
	this := ApiAppResponseOwnerAccount{}
	return &this
}

// GetAccountId returns the AccountId field value if set, zero value otherwise.
func (o *ApiAppResponseOwnerAccount) GetAccountId() string {
	if o == nil || IsNil(o.AccountId) {
		var ret string
		return ret
	}
	return *o.AccountId
}

// GetAccountIdOk returns a tuple with the AccountId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppResponseOwnerAccount) GetAccountIdOk() (*string, bool) {
	if o == nil || IsNil(o.AccountId) {
		return nil, false
	}
	return o.AccountId, true
}

// HasAccountId returns a boolean if a field has been set.
func (o *ApiAppResponseOwnerAccount) HasAccountId() bool {
	if o != nil && !IsNil(o.AccountId) {
		return true
	}

	return false
}

// SetAccountId gets a reference to the given string and assigns it to the AccountId field.
func (o *ApiAppResponseOwnerAccount) SetAccountId(v string) {
	o.AccountId = &v
}

// GetEmailAddress returns the EmailAddress field value if set, zero value otherwise.
func (o *ApiAppResponseOwnerAccount) GetEmailAddress() string {
	if o == nil || IsNil(o.EmailAddress) {
		var ret string
		return ret
	}
	return *o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiAppResponseOwnerAccount) GetEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.EmailAddress) {
		return nil, false
	}
	return o.EmailAddress, true
}

// HasEmailAddress returns a boolean if a field has been set.
func (o *ApiAppResponseOwnerAccount) HasEmailAddress() bool {
	if o != nil && !IsNil(o.EmailAddress) {
		return true
	}

	return false
}

// SetEmailAddress gets a reference to the given string and assigns it to the EmailAddress field.
func (o *ApiAppResponseOwnerAccount) SetEmailAddress(v string) {
	o.EmailAddress = &v
}

func (o ApiAppResponseOwnerAccount) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApiAppResponseOwnerAccount) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AccountId) {
		toSerialize["account_id"] = o.AccountId
	}
	if !IsNil(o.EmailAddress) {
		toSerialize["email_address"] = o.EmailAddress
	}
	return toSerialize, nil
}

type NullableApiAppResponseOwnerAccount struct {
	value *ApiAppResponseOwnerAccount
	isSet bool
}

func (v NullableApiAppResponseOwnerAccount) Get() *ApiAppResponseOwnerAccount {
	return v.value
}

func (v *NullableApiAppResponseOwnerAccount) Set(val *ApiAppResponseOwnerAccount) {
	v.value = val
	v.isSet = true
}

func (v NullableApiAppResponseOwnerAccount) IsSet() bool {
	return v.isSet
}

func (v *NullableApiAppResponseOwnerAccount) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiAppResponseOwnerAccount(val *ApiAppResponseOwnerAccount) *NullableApiAppResponseOwnerAccount {
	return &NullableApiAppResponseOwnerAccount{value: val, isSet: true}
}

func (v NullableApiAppResponseOwnerAccount) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiAppResponseOwnerAccount) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


