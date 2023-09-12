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

// checks if the TemplateResponseDocumentCustomFieldText type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateResponseDocumentCustomFieldText{}

// TemplateResponseDocumentCustomFieldText This class extends `TemplateResponseDocumentCustomFieldBase`
type TemplateResponseDocumentCustomFieldText struct {
	// The type of this Custom Field. Only `text` and `checkbox` are currently supported.  * Text uses `TemplateResponseDocumentCustomFieldText` * Checkbox uses `TemplateResponseDocumentCustomFieldCheckbox`
	Type string `json:"type"`
	AvgTextLength *TemplateResponseFieldAvgTextLength `json:"avg_text_length,omitempty"`
	// Whether this form field is multiline text.
	IsMultiline *bool `json:"isMultiline,omitempty"`
	// Original font size used in this form field's text.
	OriginalFontSize *int32 `json:"originalFontSize,omitempty"`
	// Font family used in this form field's text.
	FontFamily *string `json:"fontFamily,omitempty"`
}

// NewTemplateResponseDocumentCustomFieldText instantiates a new TemplateResponseDocumentCustomFieldText object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateResponseDocumentCustomFieldText(type_ string) *TemplateResponseDocumentCustomFieldText {
	this := TemplateResponseDocumentCustomFieldText{}
	this.Type = type_
	return &this
}

// NewTemplateResponseDocumentCustomFieldTextWithDefaults instantiates a new TemplateResponseDocumentCustomFieldText object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateResponseDocumentCustomFieldTextWithDefaults() *TemplateResponseDocumentCustomFieldText {
	this := TemplateResponseDocumentCustomFieldText{}
	var type_ string = "text"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *TemplateResponseDocumentCustomFieldText) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentCustomFieldText) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TemplateResponseDocumentCustomFieldText) SetType(v string) {
	o.Type = v
}

// GetAvgTextLength returns the AvgTextLength field value if set, zero value otherwise.
func (o *TemplateResponseDocumentCustomFieldText) GetAvgTextLength() TemplateResponseFieldAvgTextLength {
	if o == nil || IsNil(o.AvgTextLength) {
		var ret TemplateResponseFieldAvgTextLength
		return ret
	}
	return *o.AvgTextLength
}

// GetAvgTextLengthOk returns a tuple with the AvgTextLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentCustomFieldText) GetAvgTextLengthOk() (*TemplateResponseFieldAvgTextLength, bool) {
	if o == nil || IsNil(o.AvgTextLength) {
		return nil, false
	}
	return o.AvgTextLength, true
}

// HasAvgTextLength returns a boolean if a field has been set.
func (o *TemplateResponseDocumentCustomFieldText) HasAvgTextLength() bool {
	if o != nil && !IsNil(o.AvgTextLength) {
		return true
	}

	return false
}

// SetAvgTextLength gets a reference to the given TemplateResponseFieldAvgTextLength and assigns it to the AvgTextLength field.
func (o *TemplateResponseDocumentCustomFieldText) SetAvgTextLength(v TemplateResponseFieldAvgTextLength) {
	o.AvgTextLength = &v
}

// GetIsMultiline returns the IsMultiline field value if set, zero value otherwise.
func (o *TemplateResponseDocumentCustomFieldText) GetIsMultiline() bool {
	if o == nil || IsNil(o.IsMultiline) {
		var ret bool
		return ret
	}
	return *o.IsMultiline
}

// GetIsMultilineOk returns a tuple with the IsMultiline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentCustomFieldText) GetIsMultilineOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMultiline) {
		return nil, false
	}
	return o.IsMultiline, true
}

// HasIsMultiline returns a boolean if a field has been set.
func (o *TemplateResponseDocumentCustomFieldText) HasIsMultiline() bool {
	if o != nil && !IsNil(o.IsMultiline) {
		return true
	}

	return false
}

// SetIsMultiline gets a reference to the given bool and assigns it to the IsMultiline field.
func (o *TemplateResponseDocumentCustomFieldText) SetIsMultiline(v bool) {
	o.IsMultiline = &v
}

// GetOriginalFontSize returns the OriginalFontSize field value if set, zero value otherwise.
func (o *TemplateResponseDocumentCustomFieldText) GetOriginalFontSize() int32 {
	if o == nil || IsNil(o.OriginalFontSize) {
		var ret int32
		return ret
	}
	return *o.OriginalFontSize
}

// GetOriginalFontSizeOk returns a tuple with the OriginalFontSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentCustomFieldText) GetOriginalFontSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginalFontSize) {
		return nil, false
	}
	return o.OriginalFontSize, true
}

// HasOriginalFontSize returns a boolean if a field has been set.
func (o *TemplateResponseDocumentCustomFieldText) HasOriginalFontSize() bool {
	if o != nil && !IsNil(o.OriginalFontSize) {
		return true
	}

	return false
}

// SetOriginalFontSize gets a reference to the given int32 and assigns it to the OriginalFontSize field.
func (o *TemplateResponseDocumentCustomFieldText) SetOriginalFontSize(v int32) {
	o.OriginalFontSize = &v
}

// GetFontFamily returns the FontFamily field value if set, zero value otherwise.
func (o *TemplateResponseDocumentCustomFieldText) GetFontFamily() string {
	if o == nil || IsNil(o.FontFamily) {
		var ret string
		return ret
	}
	return *o.FontFamily
}

// GetFontFamilyOk returns a tuple with the FontFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentCustomFieldText) GetFontFamilyOk() (*string, bool) {
	if o == nil || IsNil(o.FontFamily) {
		return nil, false
	}
	return o.FontFamily, true
}

// HasFontFamily returns a boolean if a field has been set.
func (o *TemplateResponseDocumentCustomFieldText) HasFontFamily() bool {
	if o != nil && !IsNil(o.FontFamily) {
		return true
	}

	return false
}

// SetFontFamily gets a reference to the given string and assigns it to the FontFamily field.
func (o *TemplateResponseDocumentCustomFieldText) SetFontFamily(v string) {
	o.FontFamily = &v
}

func (o TemplateResponseDocumentCustomFieldText) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateResponseDocumentCustomFieldText) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.AvgTextLength) {
		toSerialize["avg_text_length"] = o.AvgTextLength
	}
	if !IsNil(o.IsMultiline) {
		toSerialize["isMultiline"] = o.IsMultiline
	}
	if !IsNil(o.OriginalFontSize) {
		toSerialize["originalFontSize"] = o.OriginalFontSize
	}
	if !IsNil(o.FontFamily) {
		toSerialize["fontFamily"] = o.FontFamily
	}
	return toSerialize, nil
}

type NullableTemplateResponseDocumentCustomFieldText struct {
	value *TemplateResponseDocumentCustomFieldText
	isSet bool
}

func (v NullableTemplateResponseDocumentCustomFieldText) Get() *TemplateResponseDocumentCustomFieldText {
	return v.value
}

func (v *NullableTemplateResponseDocumentCustomFieldText) Set(val *TemplateResponseDocumentCustomFieldText) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateResponseDocumentCustomFieldText) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateResponseDocumentCustomFieldText) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateResponseDocumentCustomFieldText(val *TemplateResponseDocumentCustomFieldText) *NullableTemplateResponseDocumentCustomFieldText {
	return &NullableTemplateResponseDocumentCustomFieldText{value: val, isSet: true}
}

func (v NullableTemplateResponseDocumentCustomFieldText) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateResponseDocumentCustomFieldText) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

