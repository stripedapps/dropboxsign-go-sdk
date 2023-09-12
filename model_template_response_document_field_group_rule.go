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

// checks if the TemplateResponseDocumentFieldGroupRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateResponseDocumentFieldGroupRule{}

// TemplateResponseDocumentFieldGroupRule The rule used to validate checkboxes in the form field group. See [checkbox field grouping](/api/reference/constants/#checkbox-field-grouping).
type TemplateResponseDocumentFieldGroupRule struct {
	// Examples: `require_0-1` `require_1` `require_1-ormore`  - Check out the list of [acceptable `requirement` checkbox type values](/api/reference/constants/#checkbox-field-grouping). - Check out the list of [acceptable `requirement` radio type fields](/api/reference/constants/#radio-field-grouping). - Radio groups require **at least** two fields per group.
	Requirement *string `json:"requirement,omitempty"`
	// Name of the group
	GroupLabel *string `json:"groupLabel,omitempty"`
}

// NewTemplateResponseDocumentFieldGroupRule instantiates a new TemplateResponseDocumentFieldGroupRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateResponseDocumentFieldGroupRule() *TemplateResponseDocumentFieldGroupRule {
	this := TemplateResponseDocumentFieldGroupRule{}
	return &this
}

// NewTemplateResponseDocumentFieldGroupRuleWithDefaults instantiates a new TemplateResponseDocumentFieldGroupRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateResponseDocumentFieldGroupRuleWithDefaults() *TemplateResponseDocumentFieldGroupRule {
	this := TemplateResponseDocumentFieldGroupRule{}
	return &this
}

// GetRequirement returns the Requirement field value if set, zero value otherwise.
func (o *TemplateResponseDocumentFieldGroupRule) GetRequirement() string {
	if o == nil || IsNil(o.Requirement) {
		var ret string
		return ret
	}
	return *o.Requirement
}

// GetRequirementOk returns a tuple with the Requirement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentFieldGroupRule) GetRequirementOk() (*string, bool) {
	if o == nil || IsNil(o.Requirement) {
		return nil, false
	}
	return o.Requirement, true
}

// HasRequirement returns a boolean if a field has been set.
func (o *TemplateResponseDocumentFieldGroupRule) HasRequirement() bool {
	if o != nil && !IsNil(o.Requirement) {
		return true
	}

	return false
}

// SetRequirement gets a reference to the given string and assigns it to the Requirement field.
func (o *TemplateResponseDocumentFieldGroupRule) SetRequirement(v string) {
	o.Requirement = &v
}

// GetGroupLabel returns the GroupLabel field value if set, zero value otherwise.
func (o *TemplateResponseDocumentFieldGroupRule) GetGroupLabel() string {
	if o == nil || IsNil(o.GroupLabel) {
		var ret string
		return ret
	}
	return *o.GroupLabel
}

// GetGroupLabelOk returns a tuple with the GroupLabel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentFieldGroupRule) GetGroupLabelOk() (*string, bool) {
	if o == nil || IsNil(o.GroupLabel) {
		return nil, false
	}
	return o.GroupLabel, true
}

// HasGroupLabel returns a boolean if a field has been set.
func (o *TemplateResponseDocumentFieldGroupRule) HasGroupLabel() bool {
	if o != nil && !IsNil(o.GroupLabel) {
		return true
	}

	return false
}

// SetGroupLabel gets a reference to the given string and assigns it to the GroupLabel field.
func (o *TemplateResponseDocumentFieldGroupRule) SetGroupLabel(v string) {
	o.GroupLabel = &v
}

func (o TemplateResponseDocumentFieldGroupRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateResponseDocumentFieldGroupRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Requirement) {
		toSerialize["requirement"] = o.Requirement
	}
	if !IsNil(o.GroupLabel) {
		toSerialize["groupLabel"] = o.GroupLabel
	}
	return toSerialize, nil
}

type NullableTemplateResponseDocumentFieldGroupRule struct {
	value *TemplateResponseDocumentFieldGroupRule
	isSet bool
}

func (v NullableTemplateResponseDocumentFieldGroupRule) Get() *TemplateResponseDocumentFieldGroupRule {
	return v.value
}

func (v *NullableTemplateResponseDocumentFieldGroupRule) Set(val *TemplateResponseDocumentFieldGroupRule) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateResponseDocumentFieldGroupRule) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateResponseDocumentFieldGroupRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateResponseDocumentFieldGroupRule(val *TemplateResponseDocumentFieldGroupRule) *NullableTemplateResponseDocumentFieldGroupRule {
	return &NullableTemplateResponseDocumentFieldGroupRule{value: val, isSet: true}
}

func (v NullableTemplateResponseDocumentFieldGroupRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateResponseDocumentFieldGroupRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

