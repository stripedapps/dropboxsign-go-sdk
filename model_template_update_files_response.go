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

// checks if the TemplateUpdateFilesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateUpdateFilesResponse{}

// TemplateUpdateFilesResponse struct for TemplateUpdateFilesResponse
type TemplateUpdateFilesResponse struct {
	Template *TemplateUpdateFilesResponseTemplate `json:"template,omitempty"`
}

// NewTemplateUpdateFilesResponse instantiates a new TemplateUpdateFilesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateUpdateFilesResponse() *TemplateUpdateFilesResponse {
	this := TemplateUpdateFilesResponse{}
	return &this
}

// NewTemplateUpdateFilesResponseWithDefaults instantiates a new TemplateUpdateFilesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateUpdateFilesResponseWithDefaults() *TemplateUpdateFilesResponse {
	this := TemplateUpdateFilesResponse{}
	return &this
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *TemplateUpdateFilesResponse) GetTemplate() TemplateUpdateFilesResponseTemplate {
	if o == nil || IsNil(o.Template) {
		var ret TemplateUpdateFilesResponseTemplate
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateUpdateFilesResponse) GetTemplateOk() (*TemplateUpdateFilesResponseTemplate, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *TemplateUpdateFilesResponse) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given TemplateUpdateFilesResponseTemplate and assigns it to the Template field.
func (o *TemplateUpdateFilesResponse) SetTemplate(v TemplateUpdateFilesResponseTemplate) {
	o.Template = &v
}

func (o TemplateUpdateFilesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateUpdateFilesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	return toSerialize, nil
}

type NullableTemplateUpdateFilesResponse struct {
	value *TemplateUpdateFilesResponse
	isSet bool
}

func (v NullableTemplateUpdateFilesResponse) Get() *TemplateUpdateFilesResponse {
	return v.value
}

func (v *NullableTemplateUpdateFilesResponse) Set(val *TemplateUpdateFilesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateUpdateFilesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateUpdateFilesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateUpdateFilesResponse(val *TemplateUpdateFilesResponse) *NullableTemplateUpdateFilesResponse {
	return &NullableTemplateUpdateFilesResponse{value: val, isSet: true}
}

func (v NullableTemplateUpdateFilesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateUpdateFilesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


