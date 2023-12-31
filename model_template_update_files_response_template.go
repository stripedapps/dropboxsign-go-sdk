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

// checks if the TemplateUpdateFilesResponseTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateUpdateFilesResponseTemplate{}

// TemplateUpdateFilesResponseTemplate Contains template id
type TemplateUpdateFilesResponseTemplate struct {
	// The id of the Template.
	TemplateId *string `json:"template_id,omitempty"`
	// A list of warnings.
	// Deprecated
	Warnings []WarningResponse `json:"warnings,omitempty"`
}

// NewTemplateUpdateFilesResponseTemplate instantiates a new TemplateUpdateFilesResponseTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateUpdateFilesResponseTemplate() *TemplateUpdateFilesResponseTemplate {
	this := TemplateUpdateFilesResponseTemplate{}
	return &this
}

// NewTemplateUpdateFilesResponseTemplateWithDefaults instantiates a new TemplateUpdateFilesResponseTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateUpdateFilesResponseTemplateWithDefaults() *TemplateUpdateFilesResponseTemplate {
	this := TemplateUpdateFilesResponseTemplate{}
	return &this
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *TemplateUpdateFilesResponseTemplate) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateUpdateFilesResponseTemplate) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *TemplateUpdateFilesResponseTemplate) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *TemplateUpdateFilesResponseTemplate) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
// Deprecated
func (o *TemplateUpdateFilesResponseTemplate) GetWarnings() []WarningResponse {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningResponse
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *TemplateUpdateFilesResponseTemplate) GetWarningsOk() ([]WarningResponse, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *TemplateUpdateFilesResponseTemplate) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningResponse and assigns it to the Warnings field.
// Deprecated
func (o *TemplateUpdateFilesResponseTemplate) SetWarnings(v []WarningResponse) {
	o.Warnings = v
}

func (o TemplateUpdateFilesResponseTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateUpdateFilesResponseTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableTemplateUpdateFilesResponseTemplate struct {
	value *TemplateUpdateFilesResponseTemplate
	isSet bool
}

func (v NullableTemplateUpdateFilesResponseTemplate) Get() *TemplateUpdateFilesResponseTemplate {
	return v.value
}

func (v *NullableTemplateUpdateFilesResponseTemplate) Set(val *TemplateUpdateFilesResponseTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateUpdateFilesResponseTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateUpdateFilesResponseTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateUpdateFilesResponseTemplate(val *TemplateUpdateFilesResponseTemplate) *NullableTemplateUpdateFilesResponseTemplate {
	return &NullableTemplateUpdateFilesResponseTemplate{value: val, isSet: true}
}

func (v NullableTemplateUpdateFilesResponseTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateUpdateFilesResponseTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


