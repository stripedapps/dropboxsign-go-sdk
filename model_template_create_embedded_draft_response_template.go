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

// checks if the TemplateCreateEmbeddedDraftResponseTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateCreateEmbeddedDraftResponseTemplate{}

// TemplateCreateEmbeddedDraftResponseTemplate Template object with parameters: `template_id`, `edit_url`, `expires_at`.
type TemplateCreateEmbeddedDraftResponseTemplate struct {
	// The id of the Template.
	TemplateId *string `json:"template_id,omitempty"`
	// Link to edit the template.
	EditUrl *string `json:"edit_url,omitempty"`
	// When the link expires.
	ExpiresAt *int32 `json:"expires_at,omitempty"`
	// A list of warnings.
	// Deprecated
	Warnings []WarningResponse `json:"warnings,omitempty"`
}

// NewTemplateCreateEmbeddedDraftResponseTemplate instantiates a new TemplateCreateEmbeddedDraftResponseTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateCreateEmbeddedDraftResponseTemplate() *TemplateCreateEmbeddedDraftResponseTemplate {
	this := TemplateCreateEmbeddedDraftResponseTemplate{}
	return &this
}

// NewTemplateCreateEmbeddedDraftResponseTemplateWithDefaults instantiates a new TemplateCreateEmbeddedDraftResponseTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateCreateEmbeddedDraftResponseTemplateWithDefaults() *TemplateCreateEmbeddedDraftResponseTemplate {
	this := TemplateCreateEmbeddedDraftResponseTemplate{}
	return &this
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *TemplateCreateEmbeddedDraftResponseTemplate) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateCreateEmbeddedDraftResponseTemplate) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *TemplateCreateEmbeddedDraftResponseTemplate) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *TemplateCreateEmbeddedDraftResponseTemplate) SetTemplateId(v string) {
	o.TemplateId = &v
}

// GetEditUrl returns the EditUrl field value if set, zero value otherwise.
func (o *TemplateCreateEmbeddedDraftResponseTemplate) GetEditUrl() string {
	if o == nil || IsNil(o.EditUrl) {
		var ret string
		return ret
	}
	return *o.EditUrl
}

// GetEditUrlOk returns a tuple with the EditUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateCreateEmbeddedDraftResponseTemplate) GetEditUrlOk() (*string, bool) {
	if o == nil || IsNil(o.EditUrl) {
		return nil, false
	}
	return o.EditUrl, true
}

// HasEditUrl returns a boolean if a field has been set.
func (o *TemplateCreateEmbeddedDraftResponseTemplate) HasEditUrl() bool {
	if o != nil && !IsNil(o.EditUrl) {
		return true
	}

	return false
}

// SetEditUrl gets a reference to the given string and assigns it to the EditUrl field.
func (o *TemplateCreateEmbeddedDraftResponseTemplate) SetEditUrl(v string) {
	o.EditUrl = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *TemplateCreateEmbeddedDraftResponseTemplate) GetExpiresAt() int32 {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret int32
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateCreateEmbeddedDraftResponseTemplate) GetExpiresAtOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *TemplateCreateEmbeddedDraftResponseTemplate) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int32 and assigns it to the ExpiresAt field.
func (o *TemplateCreateEmbeddedDraftResponseTemplate) SetExpiresAt(v int32) {
	o.ExpiresAt = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
// Deprecated
func (o *TemplateCreateEmbeddedDraftResponseTemplate) GetWarnings() []WarningResponse {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningResponse
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
// Deprecated
func (o *TemplateCreateEmbeddedDraftResponseTemplate) GetWarningsOk() ([]WarningResponse, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *TemplateCreateEmbeddedDraftResponseTemplate) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningResponse and assigns it to the Warnings field.
// Deprecated
func (o *TemplateCreateEmbeddedDraftResponseTemplate) SetWarnings(v []WarningResponse) {
	o.Warnings = v
}

func (o TemplateCreateEmbeddedDraftResponseTemplate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateCreateEmbeddedDraftResponseTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	if !IsNil(o.EditUrl) {
		toSerialize["edit_url"] = o.EditUrl
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableTemplateCreateEmbeddedDraftResponseTemplate struct {
	value *TemplateCreateEmbeddedDraftResponseTemplate
	isSet bool
}

func (v NullableTemplateCreateEmbeddedDraftResponseTemplate) Get() *TemplateCreateEmbeddedDraftResponseTemplate {
	return v.value
}

func (v *NullableTemplateCreateEmbeddedDraftResponseTemplate) Set(val *TemplateCreateEmbeddedDraftResponseTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateCreateEmbeddedDraftResponseTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateCreateEmbeddedDraftResponseTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateCreateEmbeddedDraftResponseTemplate(val *TemplateCreateEmbeddedDraftResponseTemplate) *NullableTemplateCreateEmbeddedDraftResponseTemplate {
	return &NullableTemplateCreateEmbeddedDraftResponseTemplate{value: val, isSet: true}
}

func (v NullableTemplateCreateEmbeddedDraftResponseTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateCreateEmbeddedDraftResponseTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


