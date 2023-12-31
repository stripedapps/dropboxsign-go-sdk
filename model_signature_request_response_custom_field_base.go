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

// checks if the SignatureRequestResponseCustomFieldBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignatureRequestResponseCustomFieldBase{}

// SignatureRequestResponseCustomFieldBase An array of Custom Field objects containing the name and type of each custom field.  * Text Field uses `SignatureRequestResponseCustomFieldText` * Checkbox Field uses `SignatureRequestResponseCustomFieldCheckbox`
type SignatureRequestResponseCustomFieldBase struct {
	// The type of this Custom Field. Only 'text' and 'checkbox' are currently supported.
	Type string `json:"type"`
	// The name of the Custom Field.
	Name string `json:"name"`
	// A boolean value denoting if this field is required.
	Required *bool `json:"required,omitempty"`
	// The unique ID for this field.
	ApiId *string `json:"api_id,omitempty"`
	// The name of the Role that is able to edit this field.
	Editor *string `json:"editor,omitempty"`
}

// NewSignatureRequestResponseCustomFieldBase instantiates a new SignatureRequestResponseCustomFieldBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignatureRequestResponseCustomFieldBase(type_ string, name string) *SignatureRequestResponseCustomFieldBase {
	this := SignatureRequestResponseCustomFieldBase{}
	this.Type = type_
	this.Name = name
	return &this
}

// NewSignatureRequestResponseCustomFieldBaseWithDefaults instantiates a new SignatureRequestResponseCustomFieldBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureRequestResponseCustomFieldBaseWithDefaults() *SignatureRequestResponseCustomFieldBase {
	this := SignatureRequestResponseCustomFieldBase{}
	return &this
}

// GetType returns the Type field value
func (o *SignatureRequestResponseCustomFieldBase) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SignatureRequestResponseCustomFieldBase) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SignatureRequestResponseCustomFieldBase) SetType(v string) {
	o.Type = v
}

// GetName returns the Name field value
func (o *SignatureRequestResponseCustomFieldBase) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SignatureRequestResponseCustomFieldBase) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SignatureRequestResponseCustomFieldBase) SetName(v string) {
	o.Name = v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *SignatureRequestResponseCustomFieldBase) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestResponseCustomFieldBase) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *SignatureRequestResponseCustomFieldBase) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *SignatureRequestResponseCustomFieldBase) SetRequired(v bool) {
	o.Required = &v
}

// GetApiId returns the ApiId field value if set, zero value otherwise.
func (o *SignatureRequestResponseCustomFieldBase) GetApiId() string {
	if o == nil || IsNil(o.ApiId) {
		var ret string
		return ret
	}
	return *o.ApiId
}

// GetApiIdOk returns a tuple with the ApiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestResponseCustomFieldBase) GetApiIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApiId) {
		return nil, false
	}
	return o.ApiId, true
}

// HasApiId returns a boolean if a field has been set.
func (o *SignatureRequestResponseCustomFieldBase) HasApiId() bool {
	if o != nil && !IsNil(o.ApiId) {
		return true
	}

	return false
}

// SetApiId gets a reference to the given string and assigns it to the ApiId field.
func (o *SignatureRequestResponseCustomFieldBase) SetApiId(v string) {
	o.ApiId = &v
}

// GetEditor returns the Editor field value if set, zero value otherwise.
func (o *SignatureRequestResponseCustomFieldBase) GetEditor() string {
	if o == nil || IsNil(o.Editor) {
		var ret string
		return ret
	}
	return *o.Editor
}

// GetEditorOk returns a tuple with the Editor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestResponseCustomFieldBase) GetEditorOk() (*string, bool) {
	if o == nil || IsNil(o.Editor) {
		return nil, false
	}
	return o.Editor, true
}

// HasEditor returns a boolean if a field has been set.
func (o *SignatureRequestResponseCustomFieldBase) HasEditor() bool {
	if o != nil && !IsNil(o.Editor) {
		return true
	}

	return false
}

// SetEditor gets a reference to the given string and assigns it to the Editor field.
func (o *SignatureRequestResponseCustomFieldBase) SetEditor(v string) {
	o.Editor = &v
}

func (o SignatureRequestResponseCustomFieldBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignatureRequestResponseCustomFieldBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["name"] = o.Name
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !IsNil(o.ApiId) {
		toSerialize["api_id"] = o.ApiId
	}
	if !IsNil(o.Editor) {
		toSerialize["editor"] = o.Editor
	}
	return toSerialize, nil
}

type NullableSignatureRequestResponseCustomFieldBase struct {
	value *SignatureRequestResponseCustomFieldBase
	isSet bool
}

func (v NullableSignatureRequestResponseCustomFieldBase) Get() *SignatureRequestResponseCustomFieldBase {
	return v.value
}

func (v *NullableSignatureRequestResponseCustomFieldBase) Set(val *SignatureRequestResponseCustomFieldBase) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureRequestResponseCustomFieldBase) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureRequestResponseCustomFieldBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureRequestResponseCustomFieldBase(val *SignatureRequestResponseCustomFieldBase) *NullableSignatureRequestResponseCustomFieldBase {
	return &NullableSignatureRequestResponseCustomFieldBase{value: val, isSet: true}
}

func (v NullableSignatureRequestResponseCustomFieldBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureRequestResponseCustomFieldBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


