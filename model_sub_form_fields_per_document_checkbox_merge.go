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

// checks if the SubFormFieldsPerDocumentCheckboxMerge type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubFormFieldsPerDocumentCheckboxMerge{}

// SubFormFieldsPerDocumentCheckboxMerge This class extends `SubFormFieldsPerDocumentBase`.
type SubFormFieldsPerDocumentCheckboxMerge struct {
	SubFormFieldsPerDocumentBase
	// A checkbox field that has default value set using pre-filled data. Use the `SubFormFieldsPerDocumentCheckboxMerge` class.
	Type string `json:"type"`
}

// NewSubFormFieldsPerDocumentCheckboxMerge instantiates a new SubFormFieldsPerDocumentCheckboxMerge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubFormFieldsPerDocumentCheckboxMerge(type_ string, documentIndex int32, apiId string, height int32, required bool, signer string, width int32, x int32, y int32) *SubFormFieldsPerDocumentCheckboxMerge {
	this := SubFormFieldsPerDocumentCheckboxMerge{}
	this.DocumentIndex = documentIndex
	this.ApiId = apiId
	this.Height = height
	this.Required = required
	this.Signer = signer
	this.Type = type_
	this.Width = width
	this.X = x
	this.Y = y
	return &this
}

// NewSubFormFieldsPerDocumentCheckboxMergeWithDefaults instantiates a new SubFormFieldsPerDocumentCheckboxMerge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubFormFieldsPerDocumentCheckboxMergeWithDefaults() *SubFormFieldsPerDocumentCheckboxMerge {
	this := SubFormFieldsPerDocumentCheckboxMerge{}
	var type_ string = "checkbox-merge"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *SubFormFieldsPerDocumentCheckboxMerge) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentCheckboxMerge) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubFormFieldsPerDocumentCheckboxMerge) SetType(v string) {
	o.Type = v
}

func (o SubFormFieldsPerDocumentCheckboxMerge) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubFormFieldsPerDocumentCheckboxMerge) ToMap() (map[string]interface{}, error) {
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
	return toSerialize, nil
}

type NullableSubFormFieldsPerDocumentCheckboxMerge struct {
	value *SubFormFieldsPerDocumentCheckboxMerge
	isSet bool
}

func (v NullableSubFormFieldsPerDocumentCheckboxMerge) Get() *SubFormFieldsPerDocumentCheckboxMerge {
	return v.value
}

func (v *NullableSubFormFieldsPerDocumentCheckboxMerge) Set(val *SubFormFieldsPerDocumentCheckboxMerge) {
	v.value = val
	v.isSet = true
}

func (v NullableSubFormFieldsPerDocumentCheckboxMerge) IsSet() bool {
	return v.isSet
}

func (v *NullableSubFormFieldsPerDocumentCheckboxMerge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubFormFieldsPerDocumentCheckboxMerge(val *SubFormFieldsPerDocumentCheckboxMerge) *NullableSubFormFieldsPerDocumentCheckboxMerge {
	return &NullableSubFormFieldsPerDocumentCheckboxMerge{value: val, isSet: true}
}

func (v NullableSubFormFieldsPerDocumentCheckboxMerge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubFormFieldsPerDocumentCheckboxMerge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


