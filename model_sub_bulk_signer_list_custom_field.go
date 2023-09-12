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

// checks if the SubBulkSignerListCustomField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubBulkSignerListCustomField{}

// SubBulkSignerListCustomField struct for SubBulkSignerListCustomField
type SubBulkSignerListCustomField struct {
	// The name of the custom field. Must be the field's `name` or `api_id`.
	Name string `json:"name"`
	// The value of the custom field.
	Value string `json:"value"`
}

// NewSubBulkSignerListCustomField instantiates a new SubBulkSignerListCustomField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubBulkSignerListCustomField(name string, value string) *SubBulkSignerListCustomField {
	this := SubBulkSignerListCustomField{}
	this.Name = name
	this.Value = value
	return &this
}

// NewSubBulkSignerListCustomFieldWithDefaults instantiates a new SubBulkSignerListCustomField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubBulkSignerListCustomFieldWithDefaults() *SubBulkSignerListCustomField {
	this := SubBulkSignerListCustomField{}
	return &this
}

// GetName returns the Name field value
func (o *SubBulkSignerListCustomField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SubBulkSignerListCustomField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SubBulkSignerListCustomField) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *SubBulkSignerListCustomField) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *SubBulkSignerListCustomField) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *SubBulkSignerListCustomField) SetValue(v string) {
	o.Value = v
}

func (o SubBulkSignerListCustomField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubBulkSignerListCustomField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableSubBulkSignerListCustomField struct {
	value *SubBulkSignerListCustomField
	isSet bool
}

func (v NullableSubBulkSignerListCustomField) Get() *SubBulkSignerListCustomField {
	return v.value
}

func (v *NullableSubBulkSignerListCustomField) Set(val *SubBulkSignerListCustomField) {
	v.value = val
	v.isSet = true
}

func (v NullableSubBulkSignerListCustomField) IsSet() bool {
	return v.isSet
}

func (v *NullableSubBulkSignerListCustomField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubBulkSignerListCustomField(val *SubBulkSignerListCustomField) *NullableSubBulkSignerListCustomField {
	return &NullableSubBulkSignerListCustomField{value: val, isSet: true}
}

func (v NullableSubBulkSignerListCustomField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubBulkSignerListCustomField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


