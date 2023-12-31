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

// checks if the SubUnclaimedDraftTemplateSigner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubUnclaimedDraftTemplateSigner{}

// SubUnclaimedDraftTemplateSigner struct for SubUnclaimedDraftTemplateSigner
type SubUnclaimedDraftTemplateSigner struct {
	// Must match an existing role in chosen Template(s).
	Role string `json:"role"`
	// The name of the signer filling the role of `role`.
	Name string `json:"name"`
	// The email address of the signer filling the role of `role`.
	EmailAddress string `json:"email_address"`
}

// NewSubUnclaimedDraftTemplateSigner instantiates a new SubUnclaimedDraftTemplateSigner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubUnclaimedDraftTemplateSigner(role string, name string, emailAddress string) *SubUnclaimedDraftTemplateSigner {
	this := SubUnclaimedDraftTemplateSigner{}
	this.Role = role
	this.Name = name
	this.EmailAddress = emailAddress
	return &this
}

// NewSubUnclaimedDraftTemplateSignerWithDefaults instantiates a new SubUnclaimedDraftTemplateSigner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubUnclaimedDraftTemplateSignerWithDefaults() *SubUnclaimedDraftTemplateSigner {
	this := SubUnclaimedDraftTemplateSigner{}
	return &this
}

// GetRole returns the Role field value
func (o *SubUnclaimedDraftTemplateSigner) GetRole() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *SubUnclaimedDraftTemplateSigner) GetRoleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *SubUnclaimedDraftTemplateSigner) SetRole(v string) {
	o.Role = v
}

// GetName returns the Name field value
func (o *SubUnclaimedDraftTemplateSigner) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SubUnclaimedDraftTemplateSigner) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SubUnclaimedDraftTemplateSigner) SetName(v string) {
	o.Name = v
}

// GetEmailAddress returns the EmailAddress field value
func (o *SubUnclaimedDraftTemplateSigner) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *SubUnclaimedDraftTemplateSigner) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *SubUnclaimedDraftTemplateSigner) SetEmailAddress(v string) {
	o.EmailAddress = v
}

func (o SubUnclaimedDraftTemplateSigner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubUnclaimedDraftTemplateSigner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["role"] = o.Role
	toSerialize["name"] = o.Name
	toSerialize["email_address"] = o.EmailAddress
	return toSerialize, nil
}

type NullableSubUnclaimedDraftTemplateSigner struct {
	value *SubUnclaimedDraftTemplateSigner
	isSet bool
}

func (v NullableSubUnclaimedDraftTemplateSigner) Get() *SubUnclaimedDraftTemplateSigner {
	return v.value
}

func (v *NullableSubUnclaimedDraftTemplateSigner) Set(val *SubUnclaimedDraftTemplateSigner) {
	v.value = val
	v.isSet = true
}

func (v NullableSubUnclaimedDraftTemplateSigner) IsSet() bool {
	return v.isSet
}

func (v *NullableSubUnclaimedDraftTemplateSigner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubUnclaimedDraftTemplateSigner(val *SubUnclaimedDraftTemplateSigner) *NullableSubUnclaimedDraftTemplateSigner {
	return &NullableSubUnclaimedDraftTemplateSigner{value: val, isSet: true}
}

func (v NullableSubUnclaimedDraftTemplateSigner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubUnclaimedDraftTemplateSigner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


