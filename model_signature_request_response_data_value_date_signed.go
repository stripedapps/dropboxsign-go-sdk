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

// checks if the SignatureRequestResponseDataValueDateSigned type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignatureRequestResponseDataValueDateSigned{}

// SignatureRequestResponseDataValueDateSigned struct for SignatureRequestResponseDataValueDateSigned
type SignatureRequestResponseDataValueDateSigned struct {
	// A date
	Type *string `json:"type,omitempty"`
	// The value of the form field.
	Value *string `json:"value,omitempty"`
}

// NewSignatureRequestResponseDataValueDateSigned instantiates a new SignatureRequestResponseDataValueDateSigned object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignatureRequestResponseDataValueDateSigned() *SignatureRequestResponseDataValueDateSigned {
	this := SignatureRequestResponseDataValueDateSigned{}
	var type_ string = "date_signed"
	this.Type = &type_
	return &this
}

// NewSignatureRequestResponseDataValueDateSignedWithDefaults instantiates a new SignatureRequestResponseDataValueDateSigned object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureRequestResponseDataValueDateSignedWithDefaults() *SignatureRequestResponseDataValueDateSigned {
	this := SignatureRequestResponseDataValueDateSigned{}
	var type_ string = "date_signed"
	this.Type = &type_
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *SignatureRequestResponseDataValueDateSigned) GetType() string {
	if o == nil || IsNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestResponseDataValueDateSigned) GetTypeOk() (*string, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *SignatureRequestResponseDataValueDateSigned) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *SignatureRequestResponseDataValueDateSigned) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SignatureRequestResponseDataValueDateSigned) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestResponseDataValueDateSigned) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SignatureRequestResponseDataValueDateSigned) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SignatureRequestResponseDataValueDateSigned) SetValue(v string) {
	o.Value = &v
}

func (o SignatureRequestResponseDataValueDateSigned) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignatureRequestResponseDataValueDateSigned) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSignatureRequestResponseDataValueDateSigned struct {
	value *SignatureRequestResponseDataValueDateSigned
	isSet bool
}

func (v NullableSignatureRequestResponseDataValueDateSigned) Get() *SignatureRequestResponseDataValueDateSigned {
	return v.value
}

func (v *NullableSignatureRequestResponseDataValueDateSigned) Set(val *SignatureRequestResponseDataValueDateSigned) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureRequestResponseDataValueDateSigned) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureRequestResponseDataValueDateSigned) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureRequestResponseDataValueDateSigned(val *SignatureRequestResponseDataValueDateSigned) *NullableSignatureRequestResponseDataValueDateSigned {
	return &NullableSignatureRequestResponseDataValueDateSigned{value: val, isSet: true}
}

func (v NullableSignatureRequestResponseDataValueDateSigned) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureRequestResponseDataValueDateSigned) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


