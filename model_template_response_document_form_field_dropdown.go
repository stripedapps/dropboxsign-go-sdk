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

// checks if the TemplateResponseDocumentFormFieldDropdown type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateResponseDocumentFormFieldDropdown{}

// TemplateResponseDocumentFormFieldDropdown This class extends `TemplateResponseDocumentFormFieldBase`
type TemplateResponseDocumentFormFieldDropdown struct {
	// The type of this form field. See [field types](/api/reference/constants/#field-types).  * Text Field uses `TemplateResponseDocumentFormFieldText` * Dropdown Field uses `TemplateResponseDocumentFormFieldDropdown` * Hyperlink Field uses `TemplateResponseDocumentFormFieldHyperlink` * Checkbox Field uses `TemplateResponseDocumentFormFieldCheckbox` * Radio Field uses `TemplateResponseDocumentFormFieldRadio` * Signature Field uses `TemplateResponseDocumentFormFieldSignature` * Date Signed Field uses `TemplateResponseDocumentFormFieldDateSigned` * Initials Field uses `TemplateResponseDocumentFormFieldInitials`
	Type string `json:"type"`
}

// NewTemplateResponseDocumentFormFieldDropdown instantiates a new TemplateResponseDocumentFormFieldDropdown object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateResponseDocumentFormFieldDropdown(type_ string) *TemplateResponseDocumentFormFieldDropdown {
	this := TemplateResponseDocumentFormFieldDropdown{}
	this.Type = type_
	return &this
}

// NewTemplateResponseDocumentFormFieldDropdownWithDefaults instantiates a new TemplateResponseDocumentFormFieldDropdown object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateResponseDocumentFormFieldDropdownWithDefaults() *TemplateResponseDocumentFormFieldDropdown {
	this := TemplateResponseDocumentFormFieldDropdown{}
	var type_ string = "dropdown"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *TemplateResponseDocumentFormFieldDropdown) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentFormFieldDropdown) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TemplateResponseDocumentFormFieldDropdown) SetType(v string) {
	o.Type = v
}

func (o TemplateResponseDocumentFormFieldDropdown) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateResponseDocumentFormFieldDropdown) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableTemplateResponseDocumentFormFieldDropdown struct {
	value *TemplateResponseDocumentFormFieldDropdown
	isSet bool
}

func (v NullableTemplateResponseDocumentFormFieldDropdown) Get() *TemplateResponseDocumentFormFieldDropdown {
	return v.value
}

func (v *NullableTemplateResponseDocumentFormFieldDropdown) Set(val *TemplateResponseDocumentFormFieldDropdown) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateResponseDocumentFormFieldDropdown) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateResponseDocumentFormFieldDropdown) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateResponseDocumentFormFieldDropdown(val *TemplateResponseDocumentFormFieldDropdown) *NullableTemplateResponseDocumentFormFieldDropdown {
	return &NullableTemplateResponseDocumentFormFieldDropdown{value: val, isSet: true}
}

func (v NullableTemplateResponseDocumentFormFieldDropdown) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateResponseDocumentFormFieldDropdown) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


