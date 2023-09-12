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

// checks if the SubCC type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubCC{}

// SubCC struct for SubCC
type SubCC struct {
	// Must match an existing CC role in chosen Template(s). Multiple CC recipients cannot share the same CC role.
	Role string `json:"role"`
	// The email address of the CC recipient.
	EmailAddress string `json:"email_address"`
}

// NewSubCC instantiates a new SubCC object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubCC(role string, emailAddress string) *SubCC {
	this := SubCC{}
	this.Role = role
	this.EmailAddress = emailAddress
	return &this
}

// NewSubCCWithDefaults instantiates a new SubCC object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubCCWithDefaults() *SubCC {
	this := SubCC{}
	return &this
}

// GetRole returns the Role field value
func (o *SubCC) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *SubCC) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *SubCC) SetRole(v string) {
	o.Role = v
}

// GetEmailAddress returns the EmailAddress field value
func (o *SubCC) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *SubCC) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *SubCC) SetEmailAddress(v string) {
	o.EmailAddress = v
}

func (o SubCC) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubCC) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["role"] = o.Role
	toSerialize["email_address"] = o.EmailAddress
	return toSerialize, nil
}

type NullableSubCC struct {
	value *SubCC
	isSet bool
}

func (v NullableSubCC) Get() *SubCC {
	return v.value
}

func (v *NullableSubCC) Set(val *SubCC) {
	v.value = val
	v.isSet = true
}

func (v NullableSubCC) IsSet() bool {
	return v.isSet
}

func (v *NullableSubCC) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubCC(val *SubCC) *NullableSubCC {
	return &NullableSubCC{value: val, isSet: true}
}

func (v NullableSubCC) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubCC) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


