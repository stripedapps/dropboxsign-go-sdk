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

// checks if the SubFormFieldRuleAction type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubFormFieldRuleAction{}

// SubFormFieldRuleAction struct for SubFormFieldRuleAction
type SubFormFieldRuleAction struct {
	// **field_id** or **group_id** is required, but not both.  Must reference the `api_id` of an existing field defined within `form_fields_per_document`.  Cannot use with `group_id`. Trigger and action fields must belong to the same signer.
	FieldId *string `json:"field_id,omitempty"`
	// **group_id** or **field_id** is required, but not both.  Must reference the ID of an existing group defined within `form_field_groups`.  Cannot use with `field_id`. Trigger and action fields and groups must belong to the same signer.
	GroupId *string `json:"group_id,omitempty"`
	// `true` to hide the target field when rule is satisfied, otherwise `false`.
	Hidden bool `json:"hidden"`
	Type string `json:"type"`
}

// NewSubFormFieldRuleAction instantiates a new SubFormFieldRuleAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubFormFieldRuleAction(hidden bool, type_ string) *SubFormFieldRuleAction {
	this := SubFormFieldRuleAction{}
	this.Hidden = hidden
	this.Type = type_
	return &this
}

// NewSubFormFieldRuleActionWithDefaults instantiates a new SubFormFieldRuleAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubFormFieldRuleActionWithDefaults() *SubFormFieldRuleAction {
	this := SubFormFieldRuleAction{}
	return &this
}

// GetFieldId returns the FieldId field value if set, zero value otherwise.
func (o *SubFormFieldRuleAction) GetFieldId() string {
	if o == nil || IsNil(o.FieldId) {
		var ret string
		return ret
	}
	return *o.FieldId
}

// GetFieldIdOk returns a tuple with the FieldId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubFormFieldRuleAction) GetFieldIdOk() (*string, bool) {
	if o == nil || IsNil(o.FieldId) {
		return nil, false
	}
	return o.FieldId, true
}

// HasFieldId returns a boolean if a field has been set.
func (o *SubFormFieldRuleAction) HasFieldId() bool {
	if o != nil && !IsNil(o.FieldId) {
		return true
	}

	return false
}

// SetFieldId gets a reference to the given string and assigns it to the FieldId field.
func (o *SubFormFieldRuleAction) SetFieldId(v string) {
	o.FieldId = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *SubFormFieldRuleAction) GetGroupId() string {
	if o == nil || IsNil(o.GroupId) {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubFormFieldRuleAction) GetGroupIdOk() (*string, bool) {
	if o == nil || IsNil(o.GroupId) {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *SubFormFieldRuleAction) HasGroupId() bool {
	if o != nil && !IsNil(o.GroupId) {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *SubFormFieldRuleAction) SetGroupId(v string) {
	o.GroupId = &v
}

// GetHidden returns the Hidden field value
func (o *SubFormFieldRuleAction) GetHidden() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldRuleAction) GetHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hidden, true
}

// SetHidden sets field value
func (o *SubFormFieldRuleAction) SetHidden(v bool) {
	o.Hidden = v
}

// GetType returns the Type field value
func (o *SubFormFieldRuleAction) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldRuleAction) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubFormFieldRuleAction) SetType(v string) {
	o.Type = v
}

func (o SubFormFieldRuleAction) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubFormFieldRuleAction) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FieldId) {
		toSerialize["field_id"] = o.FieldId
	}
	if !IsNil(o.GroupId) {
		toSerialize["group_id"] = o.GroupId
	}
	toSerialize["hidden"] = o.Hidden
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

type NullableSubFormFieldRuleAction struct {
	value *SubFormFieldRuleAction
	isSet bool
}

func (v NullableSubFormFieldRuleAction) Get() *SubFormFieldRuleAction {
	return v.value
}

func (v *NullableSubFormFieldRuleAction) Set(val *SubFormFieldRuleAction) {
	v.value = val
	v.isSet = true
}

func (v NullableSubFormFieldRuleAction) IsSet() bool {
	return v.isSet
}

func (v *NullableSubFormFieldRuleAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubFormFieldRuleAction(val *SubFormFieldRuleAction) *NullableSubFormFieldRuleAction {
	return &NullableSubFormFieldRuleAction{value: val, isSet: true}
}

func (v NullableSubFormFieldRuleAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubFormFieldRuleAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


