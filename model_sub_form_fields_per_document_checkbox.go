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

// checks if the SubFormFieldsPerDocumentCheckbox type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubFormFieldsPerDocumentCheckbox{}

// SubFormFieldsPerDocumentCheckbox This class extends `SubFormFieldsPerDocumentBase`.
type SubFormFieldsPerDocumentCheckbox struct {
	SubFormFieldsPerDocumentBase
	// A yes/no checkbox. Use the `SubFormFieldsPerDocumentCheckbox` class.
	Type string `json:"type"`
	// String referencing group defined in `form_field_groups` parameter.
	Group *string `json:"group,omitempty"`
	// `true` for checking the checkbox field by default, otherwise `false`.
	IsChecked bool `json:"is_checked"`
}

// NewSubFormFieldsPerDocumentCheckbox instantiates a new SubFormFieldsPerDocumentCheckbox object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubFormFieldsPerDocumentCheckbox(type_ string, isChecked bool, documentIndex int32, apiId string, height int32, required bool, signer string, width int32, x int32, y int32) *SubFormFieldsPerDocumentCheckbox {
	this := SubFormFieldsPerDocumentCheckbox{}
	this.DocumentIndex = documentIndex
	this.ApiId = apiId
	this.Height = height
	this.Required = required
	this.Signer = signer
	this.Type = type_
	this.Width = width
	this.X = x
	this.Y = y
	this.IsChecked = isChecked
	return &this
}

// NewSubFormFieldsPerDocumentCheckboxWithDefaults instantiates a new SubFormFieldsPerDocumentCheckbox object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubFormFieldsPerDocumentCheckboxWithDefaults() *SubFormFieldsPerDocumentCheckbox {
	this := SubFormFieldsPerDocumentCheckbox{}
	var type_ string = "checkbox"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *SubFormFieldsPerDocumentCheckbox) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentCheckbox) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubFormFieldsPerDocumentCheckbox) SetType(v string) {
	o.Type = v
}

// GetGroup returns the Group field value if set, zero value otherwise.
func (o *SubFormFieldsPerDocumentCheckbox) GetGroup() string {
	if o == nil || IsNil(o.Group) {
		var ret string
		return ret
	}
	return *o.Group
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentCheckbox) GetGroupOk() (*string, bool) {
	if o == nil || IsNil(o.Group) {
		return nil, false
	}
	return o.Group, true
}

// HasGroup returns a boolean if a field has been set.
func (o *SubFormFieldsPerDocumentCheckbox) HasGroup() bool {
	if o != nil && !IsNil(o.Group) {
		return true
	}

	return false
}

// SetGroup gets a reference to the given string and assigns it to the Group field.
func (o *SubFormFieldsPerDocumentCheckbox) SetGroup(v string) {
	o.Group = &v
}

// GetIsChecked returns the IsChecked field value
func (o *SubFormFieldsPerDocumentCheckbox) GetIsChecked() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsChecked
}

// GetIsCheckedOk returns a tuple with the IsChecked field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentCheckbox) GetIsCheckedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsChecked, true
}

// SetIsChecked sets field value
func (o *SubFormFieldsPerDocumentCheckbox) SetIsChecked(v bool) {
	o.IsChecked = v
}

func (o SubFormFieldsPerDocumentCheckbox) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubFormFieldsPerDocumentCheckbox) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSubFormFieldsPerDocumentBase, errSubFormFieldsPerDocumentBase := json.Marshal(o.SubFormFieldsPerDocumentBase)
	if errSubFormFieldsPerDocumentBase != nil {
		return map[string]interface{}{}, errSubFormFieldsPerDocumentBase
	}
	errSubFormFieldsPerDocumentBase = json.Unmarshal([]byte(serializedSubFormFieldsPerDocumentBase), &toSerialize)
	if errSubFormFieldsPerDocumentBase != nil {
		return map[string]interface{}{}, errSubFormFieldsPerDocumentBase
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Group) {
		toSerialize["group"] = o.Group
	}
	toSerialize["is_checked"] = o.IsChecked
	return toSerialize, nil
}

type NullableSubFormFieldsPerDocumentCheckbox struct {
	value *SubFormFieldsPerDocumentCheckbox
	isSet bool
}

func (v NullableSubFormFieldsPerDocumentCheckbox) Get() *SubFormFieldsPerDocumentCheckbox {
	return v.value
}

func (v *NullableSubFormFieldsPerDocumentCheckbox) Set(val *SubFormFieldsPerDocumentCheckbox) {
	v.value = val
	v.isSet = true
}

func (v NullableSubFormFieldsPerDocumentCheckbox) IsSet() bool {
	return v.isSet
}

func (v *NullableSubFormFieldsPerDocumentCheckbox) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubFormFieldsPerDocumentCheckbox(val *SubFormFieldsPerDocumentCheckbox) *NullableSubFormFieldsPerDocumentCheckbox {
	return &NullableSubFormFieldsPerDocumentCheckbox{value: val, isSet: true}
}

func (v NullableSubFormFieldsPerDocumentCheckbox) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubFormFieldsPerDocumentCheckbox) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


