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

// checks if the TemplateResponseDocumentCustomFieldCheckbox type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateResponseDocumentCustomFieldCheckbox{}

// TemplateResponseDocumentCustomFieldCheckbox This class extends `TemplateResponseDocumentCustomFieldBase`
type TemplateResponseDocumentCustomFieldCheckbox struct {
	// The type of this Custom Field. Only `text` and `checkbox` are currently supported.  * Text uses `TemplateResponseDocumentCustomFieldText` * Checkbox uses `TemplateResponseDocumentCustomFieldCheckbox`
	Type string `json:"type"`
}

// NewTemplateResponseDocumentCustomFieldCheckbox instantiates a new TemplateResponseDocumentCustomFieldCheckbox object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateResponseDocumentCustomFieldCheckbox(type_ string) *TemplateResponseDocumentCustomFieldCheckbox {
	this := TemplateResponseDocumentCustomFieldCheckbox{}
	this.Type = type_
	return &this
}

// NewTemplateResponseDocumentCustomFieldCheckboxWithDefaults instantiates a new TemplateResponseDocumentCustomFieldCheckbox object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateResponseDocumentCustomFieldCheckboxWithDefaults() *TemplateResponseDocumentCustomFieldCheckbox {
	this := TemplateResponseDocumentCustomFieldCheckbox{}
	var type_ string = "checkbox"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *TemplateResponseDocumentCustomFieldCheckbox) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentCustomFieldCheckbox) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TemplateResponseDocumentCustomFieldCheckbox) SetType(v string) {
	o.Type = v
}

func (o TemplateResponseDocumentCustomFieldCheckbox) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateResponseDocumentCustomFieldCheckbox) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTemplateResponseDocumentCustomFieldCheckbox struct {
	value *TemplateResponseDocumentCustomFieldCheckbox
	isSet bool
}

func (v NullableTemplateResponseDocumentCustomFieldCheckbox) Get() *TemplateResponseDocumentCustomFieldCheckbox {
	return v.value
}

func (v *NullableTemplateResponseDocumentCustomFieldCheckbox) Set(val *TemplateResponseDocumentCustomFieldCheckbox) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateResponseDocumentCustomFieldCheckbox) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateResponseDocumentCustomFieldCheckbox) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateResponseDocumentCustomFieldCheckbox(val *TemplateResponseDocumentCustomFieldCheckbox) *NullableTemplateResponseDocumentCustomFieldCheckbox {
	return &NullableTemplateResponseDocumentCustomFieldCheckbox{value: val, isSet: true}
}

func (v NullableTemplateResponseDocumentCustomFieldCheckbox) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateResponseDocumentCustomFieldCheckbox) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


