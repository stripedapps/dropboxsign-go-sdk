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

// checks if the EmbeddedEditUrlResponseEmbedded type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmbeddedEditUrlResponseEmbedded{}

// EmbeddedEditUrlResponseEmbedded An embedded template object.
type EmbeddedEditUrlResponseEmbedded struct {
	// A template url that can be opened in an iFrame.
	EditUrl *string `json:"edit_url,omitempty"`
	// The specific time that the the `edit_url` link expires, in epoch.
	ExpiresAt *int32 `json:"expires_at,omitempty"`
}

// NewEmbeddedEditUrlResponseEmbedded instantiates a new EmbeddedEditUrlResponseEmbedded object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbeddedEditUrlResponseEmbedded() *EmbeddedEditUrlResponseEmbedded {
	this := EmbeddedEditUrlResponseEmbedded{}
	return &this
}

// NewEmbeddedEditUrlResponseEmbeddedWithDefaults instantiates a new EmbeddedEditUrlResponseEmbedded object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbeddedEditUrlResponseEmbeddedWithDefaults() *EmbeddedEditUrlResponseEmbedded {
	this := EmbeddedEditUrlResponseEmbedded{}
	return &this
}

// GetEditUrl returns the EditUrl field value if set, zero value otherwise.
func (o *EmbeddedEditUrlResponseEmbedded) GetEditUrl() string {
	if o == nil || IsNil(o.EditUrl) {
		var ret string
		return ret
	}
	return *o.EditUrl
}

// GetEditUrlOk returns a tuple with the EditUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedEditUrlResponseEmbedded) GetEditUrlOk() (*string, bool) {
	if o == nil || IsNil(o.EditUrl) {
		return nil, false
	}
	return o.EditUrl, true
}

// HasEditUrl returns a boolean if a field has been set.
func (o *EmbeddedEditUrlResponseEmbedded) HasEditUrl() bool {
	if o != nil && !IsNil(o.EditUrl) {
		return true
	}

	return false
}

// SetEditUrl gets a reference to the given string and assigns it to the EditUrl field.
func (o *EmbeddedEditUrlResponseEmbedded) SetEditUrl(v string) {
	o.EditUrl = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *EmbeddedEditUrlResponseEmbedded) GetExpiresAt() int32 {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret int32
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedEditUrlResponseEmbedded) GetExpiresAtOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *EmbeddedEditUrlResponseEmbedded) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int32 and assigns it to the ExpiresAt field.
func (o *EmbeddedEditUrlResponseEmbedded) SetExpiresAt(v int32) {
	o.ExpiresAt = &v
}

func (o EmbeddedEditUrlResponseEmbedded) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmbeddedEditUrlResponseEmbedded) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.EditUrl) {
		toSerialize["edit_url"] = o.EditUrl
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	return toSerialize, nil
}

type NullableEmbeddedEditUrlResponseEmbedded struct {
	value *EmbeddedEditUrlResponseEmbedded
	isSet bool
}

func (v NullableEmbeddedEditUrlResponseEmbedded) Get() *EmbeddedEditUrlResponseEmbedded {
	return v.value
}

func (v *NullableEmbeddedEditUrlResponseEmbedded) Set(val *EmbeddedEditUrlResponseEmbedded) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbeddedEditUrlResponseEmbedded) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbeddedEditUrlResponseEmbedded) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbeddedEditUrlResponseEmbedded(val *EmbeddedEditUrlResponseEmbedded) *NullableEmbeddedEditUrlResponseEmbedded {
	return &NullableEmbeddedEditUrlResponseEmbedded{value: val, isSet: true}
}

func (v NullableEmbeddedEditUrlResponseEmbedded) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbeddedEditUrlResponseEmbedded) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


