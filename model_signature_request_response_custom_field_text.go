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

// checks if the SignatureRequestResponseCustomFieldText type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignatureRequestResponseCustomFieldText{}

// SignatureRequestResponseCustomFieldText This class extends `SignatureRequestResponseCustomFieldBase`.
type SignatureRequestResponseCustomFieldText struct {
	// The type of this Custom Field. Only 'text' and 'checkbox' are currently supported.
	Type string `json:"type"`
	// A text string for text fields
	Value *string `json:"value,omitempty"`
}

// NewSignatureRequestResponseCustomFieldText instantiates a new SignatureRequestResponseCustomFieldText object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignatureRequestResponseCustomFieldText(type_ string, name string) *SignatureRequestResponseCustomFieldText {
	this := SignatureRequestResponseCustomFieldText{}
	this.Type = type_
	this.Name = name
	return &this
}

// NewSignatureRequestResponseCustomFieldTextWithDefaults instantiates a new SignatureRequestResponseCustomFieldText object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureRequestResponseCustomFieldTextWithDefaults() *SignatureRequestResponseCustomFieldText {
	this := SignatureRequestResponseCustomFieldText{}
	var type_ string = "text"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *SignatureRequestResponseCustomFieldText) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SignatureRequestResponseCustomFieldText) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SignatureRequestResponseCustomFieldText) SetType(v string) {
	o.Type = v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SignatureRequestResponseCustomFieldText) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestResponseCustomFieldText) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SignatureRequestResponseCustomFieldText) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SignatureRequestResponseCustomFieldText) SetValue(v string) {
	o.Value = &v
}

func (o SignatureRequestResponseCustomFieldText) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignatureRequestResponseCustomFieldText) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSignatureRequestResponseCustomFieldText struct {
	value *SignatureRequestResponseCustomFieldText
	isSet bool
}

func (v NullableSignatureRequestResponseCustomFieldText) Get() *SignatureRequestResponseCustomFieldText {
	return v.value
}

func (v *NullableSignatureRequestResponseCustomFieldText) Set(val *SignatureRequestResponseCustomFieldText) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureRequestResponseCustomFieldText) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureRequestResponseCustomFieldText) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureRequestResponseCustomFieldText(val *SignatureRequestResponseCustomFieldText) *NullableSignatureRequestResponseCustomFieldText {
	return &NullableSignatureRequestResponseCustomFieldText{value: val, isSet: true}
}

func (v NullableSignatureRequestResponseCustomFieldText) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureRequestResponseCustomFieldText) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


