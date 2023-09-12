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

// checks if the SubFormFieldRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubFormFieldRule{}

// SubFormFieldRule struct for SubFormFieldRule
type SubFormFieldRule struct {
	// Must be unique across all defined rules.
	Id string `json:"id"`
	// Currently only `AND` is supported. Support for `OR` is being worked on.
	TriggerOperator string `json:"trigger_operator"`
	// An array of trigger definitions, the \"if this\" part of \"**if this**, then that\". Currently only a single trigger per rule is allowed.
	Triggers []SubFormFieldRuleTrigger `json:"triggers"`
	// An array of action definitions, the \"then that\" part of \"if this, **then that**\". Any number of actions may be attached to a single rule.
	Actions []SubFormFieldRuleAction `json:"actions"`
}

// NewSubFormFieldRule instantiates a new SubFormFieldRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubFormFieldRule(id string, triggerOperator string, triggers []SubFormFieldRuleTrigger, actions []SubFormFieldRuleAction) *SubFormFieldRule {
	this := SubFormFieldRule{}
	this.Id = id
	this.TriggerOperator = triggerOperator
	this.Triggers = triggers
	this.Actions = actions
	return &this
}

// NewSubFormFieldRuleWithDefaults instantiates a new SubFormFieldRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubFormFieldRuleWithDefaults() *SubFormFieldRule {
	this := SubFormFieldRule{}
	var triggerOperator string = "AND"
	this.TriggerOperator = triggerOperator
	return &this
}

// GetId returns the Id field value
func (o *SubFormFieldRule) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldRule) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SubFormFieldRule) SetId(v string) {
	o.Id = v
}

// GetTriggerOperator returns the TriggerOperator field value
func (o *SubFormFieldRule) GetTriggerOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TriggerOperator
}

// GetTriggerOperatorOk returns a tuple with the TriggerOperator field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldRule) GetTriggerOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TriggerOperator, true
}

// SetTriggerOperator sets field value
func (o *SubFormFieldRule) SetTriggerOperator(v string) {
	o.TriggerOperator = v
}

// GetTriggers returns the Triggers field value
func (o *SubFormFieldRule) GetTriggers() []SubFormFieldRuleTrigger {
	if o == nil {
		var ret []SubFormFieldRuleTrigger
		return ret
	}

	return o.Triggers
}

// GetTriggersOk returns a tuple with the Triggers field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldRule) GetTriggersOk() ([]SubFormFieldRuleTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return o.Triggers, true
}

// SetTriggers sets field value
func (o *SubFormFieldRule) SetTriggers(v []SubFormFieldRuleTrigger) {
	o.Triggers = v
}

// GetActions returns the Actions field value
func (o *SubFormFieldRule) GetActions() []SubFormFieldRuleAction {
	if o == nil {
		var ret []SubFormFieldRuleAction
		return ret
	}

	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldRule) GetActionsOk() ([]SubFormFieldRuleAction, bool) {
	if o == nil {
		return nil, false
	}
	return o.Actions, true
}

// SetActions sets field value
func (o *SubFormFieldRule) SetActions(v []SubFormFieldRuleAction) {
	o.Actions = v
}

func (o SubFormFieldRule) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubFormFieldRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["trigger_operator"] = o.TriggerOperator
	toSerialize["triggers"] = o.Triggers
	toSerialize["actions"] = o.Actions
	return toSerialize, nil
}

type NullableSubFormFieldRule struct {
	value *SubFormFieldRule
	isSet bool
}

func (v NullableSubFormFieldRule) Get() *SubFormFieldRule {
	return v.value
}

func (v *NullableSubFormFieldRule) Set(val *SubFormFieldRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSubFormFieldRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSubFormFieldRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubFormFieldRule(val *SubFormFieldRule) *NullableSubFormFieldRule {
	return &NullableSubFormFieldRule{value: val, isSet: true}
}

func (v NullableSubFormFieldRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubFormFieldRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


