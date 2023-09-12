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

// checks if the TemplateResponseDocumentFieldGroup type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateResponseDocumentFieldGroup{}

// TemplateResponseDocumentFieldGroup struct for TemplateResponseDocumentFieldGroup
type TemplateResponseDocumentFieldGroup struct {
	// The name of the form field group.
	Name *string `json:"name,omitempty"`
	Rule *TemplateResponseDocumentFieldGroupRule `json:"rule,omitempty"`
}

// NewTemplateResponseDocumentFieldGroup instantiates a new TemplateResponseDocumentFieldGroup object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateResponseDocumentFieldGroup() *TemplateResponseDocumentFieldGroup {
	this := TemplateResponseDocumentFieldGroup{}
	return &this
}

// NewTemplateResponseDocumentFieldGroupWithDefaults instantiates a new TemplateResponseDocumentFieldGroup object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateResponseDocumentFieldGroupWithDefaults() *TemplateResponseDocumentFieldGroup {
	this := TemplateResponseDocumentFieldGroup{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplateResponseDocumentFieldGroup) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentFieldGroup) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplateResponseDocumentFieldGroup) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplateResponseDocumentFieldGroup) SetName(v string) {
	o.Name = &v
}

// GetRule returns the Rule field value if set, zero value otherwise.
func (o *TemplateResponseDocumentFieldGroup) GetRule() TemplateResponseDocumentFieldGroupRule {
	if o == nil || IsNil(o.Rule) {
		var ret TemplateResponseDocumentFieldGroupRule
		return ret
	}
	return *o.Rule
}

// GetRuleOk returns a tuple with the Rule field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentFieldGroup) GetRuleOk() (*TemplateResponseDocumentFieldGroupRule, bool) {
	if o == nil || IsNil(o.Rule) {
		return nil, false
	}
	return o.Rule, true
}

// HasRule returns a boolean if a field has been set.
func (o *TemplateResponseDocumentFieldGroup) HasRule() bool {
	if o != nil && !IsNil(o.Rule) {
		return true
	}

	return false
}

// SetRule gets a reference to the given TemplateResponseDocumentFieldGroupRule and assigns it to the Rule field.
func (o *TemplateResponseDocumentFieldGroup) SetRule(v TemplateResponseDocumentFieldGroupRule) {
	o.Rule = &v
}

func (o TemplateResponseDocumentFieldGroup) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateResponseDocumentFieldGroup) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Rule) {
		toSerialize["rule"] = o.Rule
	}
	return toSerialize, nil
}

type NullableTemplateResponseDocumentFieldGroup struct {
	value *TemplateResponseDocumentFieldGroup
	isSet bool
}

func (v NullableTemplateResponseDocumentFieldGroup) Get() *TemplateResponseDocumentFieldGroup {
	return v.value
}

func (v *NullableTemplateResponseDocumentFieldGroup) Set(val *TemplateResponseDocumentFieldGroup) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateResponseDocumentFieldGroup) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateResponseDocumentFieldGroup) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateResponseDocumentFieldGroup(val *TemplateResponseDocumentFieldGroup) *NullableTemplateResponseDocumentFieldGroup {
	return &NullableTemplateResponseDocumentFieldGroup{value: val, isSet: true}
}

func (v NullableTemplateResponseDocumentFieldGroup) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateResponseDocumentFieldGroup) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

