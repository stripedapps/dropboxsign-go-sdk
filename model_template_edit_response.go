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

// checks if the TemplateEditResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateEditResponse{}

// TemplateEditResponse struct for TemplateEditResponse
type TemplateEditResponse struct {
	// The id of the Template.
	TemplateId *string `json:"template_id,omitempty"`
}

// NewTemplateEditResponse instantiates a new TemplateEditResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateEditResponse() *TemplateEditResponse {
	this := TemplateEditResponse{}
	return &this
}

// NewTemplateEditResponseWithDefaults instantiates a new TemplateEditResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateEditResponseWithDefaults() *TemplateEditResponse {
	this := TemplateEditResponse{}
	return &this
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *TemplateEditResponse) GetTemplateId() string {
	if o == nil || IsNil(o.TemplateId) {
		var ret string
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateEditResponse) GetTemplateIdOk() (*string, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *TemplateEditResponse) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given string and assigns it to the TemplateId field.
func (o *TemplateEditResponse) SetTemplateId(v string) {
	o.TemplateId = &v
}

func (o TemplateEditResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateEditResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	return toSerialize, nil
}

type NullableTemplateEditResponse struct {
	value *TemplateEditResponse
	isSet bool
}

func (v NullableTemplateEditResponse) Get() *TemplateEditResponse {
	return v.value
}

func (v *NullableTemplateEditResponse) Set(val *TemplateEditResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateEditResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateEditResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateEditResponse(val *TemplateEditResponse) *NullableTemplateEditResponse {
	return &NullableTemplateEditResponse{value: val, isSet: true}
}

func (v NullableTemplateEditResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateEditResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


