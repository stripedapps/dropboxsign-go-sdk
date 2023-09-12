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

// checks if the SubCustomField type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubCustomField{}

// SubCustomField When used together with merge fields, `custom_fields` allows users to add pre-filled data to their signature requests.  Pre-filled data can be used with \"send-once\" signature requests by adding merge fields with `form_fields_per_document` or [Text Tags](https://app.hellosign.com/api/textTagsWalkthrough#TextTagIntro) while passing values back with `custom_fields` together in one API call.  For using pre-filled on repeatable signature requests, merge fields are added to templates in the Dropbox Sign UI or by calling [/template/create_embedded_draft](/api/reference/operation/templateCreateEmbeddedDraft) and then passing `custom_fields` on subsequent signature requests referencing that template.
type SubCustomField struct {
	// Used to create editable merge fields. When the value matches a role passed in with `signers`, that role can edit the data that was pre-filled to that field. This field is optional, but required when this custom field object is set to `required = true`.  **Note**: Editable merge fields are only supported for single signer requests (or the first signer in ordered signature requests). If used when there are multiple signers in an unordered signature request, the editor value is ignored and the field won't be editable.
	Editor *string `json:"editor,omitempty"`
	// The name of a custom field. When working with pre-filled data, the custom field's name must have a matching merge field name or the field will remain empty on the document during signing.
	Name string `json:"name"`
	// Used to set an editable merge field when working with pre-filled data. When `true`, the custom field must specify a signer role in `editor`.
	Required *bool `json:"required,omitempty"`
	// The string that resolves (aka \"pre-fills\") to the merge field on the final document(s) used for signing.
	Value *string `json:"value,omitempty"`
}

// NewSubCustomField instantiates a new SubCustomField object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubCustomField(name string) *SubCustomField {
	this := SubCustomField{}
	this.Name = name
	var required bool = false
	this.Required = &required
	return &this
}

// NewSubCustomFieldWithDefaults instantiates a new SubCustomField object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubCustomFieldWithDefaults() *SubCustomField {
	this := SubCustomField{}
	var required bool = false
	this.Required = &required
	return &this
}

// GetEditor returns the Editor field value if set, zero value otherwise.
func (o *SubCustomField) GetEditor() string {
	if o == nil || IsNil(o.Editor) {
		var ret string
		return ret
	}
	return *o.Editor
}

// GetEditorOk returns a tuple with the Editor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubCustomField) GetEditorOk() (*string, bool) {
	if o == nil || IsNil(o.Editor) {
		return nil, false
	}
	return o.Editor, true
}

// HasEditor returns a boolean if a field has been set.
func (o *SubCustomField) HasEditor() bool {
	if o != nil && !IsNil(o.Editor) {
		return true
	}

	return false
}

// SetEditor gets a reference to the given string and assigns it to the Editor field.
func (o *SubCustomField) SetEditor(v string) {
	o.Editor = &v
}

// GetName returns the Name field value
func (o *SubCustomField) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SubCustomField) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SubCustomField) SetName(v string) {
	o.Name = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *SubCustomField) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubCustomField) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *SubCustomField) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *SubCustomField) SetRequired(v bool) {
	o.Required = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *SubCustomField) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubCustomField) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *SubCustomField) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *SubCustomField) SetValue(v string) {
	o.Value = &v
}

func (o SubCustomField) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubCustomField) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Editor) {
		toSerialize["editor"] = o.Editor
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	return toSerialize, nil
}

type NullableSubCustomField struct {
	value *SubCustomField
	isSet bool
}

func (v NullableSubCustomField) Get() *SubCustomField {
	return v.value
}

func (v *NullableSubCustomField) Set(val *SubCustomField) {
	v.value = val
	v.isSet = true
}

func (v NullableSubCustomField) IsSet() bool {
	return v.isSet
}

func (v *NullableSubCustomField) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubCustomField(val *SubCustomField) *NullableSubCustomField {
	return &NullableSubCustomField{value: val, isSet: true}
}

func (v NullableSubCustomField) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubCustomField) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


