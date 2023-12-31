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

// checks if the TemplateResponseDocumentFormFieldHyperlink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateResponseDocumentFormFieldHyperlink{}

// TemplateResponseDocumentFormFieldHyperlink This class extends `TemplateResponseDocumentFormFieldBase`
type TemplateResponseDocumentFormFieldHyperlink struct {
	// The type of this form field. See [field types](/api/reference/constants/#field-types).  * Text Field uses `TemplateResponseDocumentFormFieldText` * Dropdown Field uses `TemplateResponseDocumentFormFieldDropdown` * Hyperlink Field uses `TemplateResponseDocumentFormFieldHyperlink` * Checkbox Field uses `TemplateResponseDocumentFormFieldCheckbox` * Radio Field uses `TemplateResponseDocumentFormFieldRadio` * Signature Field uses `TemplateResponseDocumentFormFieldSignature` * Date Signed Field uses `TemplateResponseDocumentFormFieldDateSigned` * Initials Field uses `TemplateResponseDocumentFormFieldInitials`
	Type string `json:"type"`
	AvgTextLength *TemplateResponseFieldAvgTextLength `json:"avg_text_length,omitempty"`
	// Whether this form field is multiline text.
	IsMultiline *bool `json:"isMultiline,omitempty"`
	// Original font size used in this form field's text.
	OriginalFontSize *int32 `json:"originalFontSize,omitempty"`
	// Font family used in this form field's text.
	FontFamily *string `json:"fontFamily,omitempty"`
}

// NewTemplateResponseDocumentFormFieldHyperlink instantiates a new TemplateResponseDocumentFormFieldHyperlink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateResponseDocumentFormFieldHyperlink(type_ string) *TemplateResponseDocumentFormFieldHyperlink {
	this := TemplateResponseDocumentFormFieldHyperlink{}
	this.Type = type_
	return &this
}

// NewTemplateResponseDocumentFormFieldHyperlinkWithDefaults instantiates a new TemplateResponseDocumentFormFieldHyperlink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateResponseDocumentFormFieldHyperlinkWithDefaults() *TemplateResponseDocumentFormFieldHyperlink {
	this := TemplateResponseDocumentFormFieldHyperlink{}
	var type_ string = "hyperlink"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *TemplateResponseDocumentFormFieldHyperlink) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentFormFieldHyperlink) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TemplateResponseDocumentFormFieldHyperlink) SetType(v string) {
	o.Type = v
}

// GetAvgTextLength returns the AvgTextLength field value if set, zero value otherwise.
func (o *TemplateResponseDocumentFormFieldHyperlink) GetAvgTextLength() TemplateResponseFieldAvgTextLength {
	if o == nil || IsNil(o.AvgTextLength) {
		var ret TemplateResponseFieldAvgTextLength
		return ret
	}
	return *o.AvgTextLength
}

// GetAvgTextLengthOk returns a tuple with the AvgTextLength field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentFormFieldHyperlink) GetAvgTextLengthOk() (*TemplateResponseFieldAvgTextLength, bool) {
	if o == nil || IsNil(o.AvgTextLength) {
		return nil, false
	}
	return o.AvgTextLength, true
}

// HasAvgTextLength returns a boolean if a field has been set.
func (o *TemplateResponseDocumentFormFieldHyperlink) HasAvgTextLength() bool {
	if o != nil && !IsNil(o.AvgTextLength) {
		return true
	}

	return false
}

// SetAvgTextLength gets a reference to the given TemplateResponseFieldAvgTextLength and assigns it to the AvgTextLength field.
func (o *TemplateResponseDocumentFormFieldHyperlink) SetAvgTextLength(v TemplateResponseFieldAvgTextLength) {
	o.AvgTextLength = &v
}

// GetIsMultiline returns the IsMultiline field value if set, zero value otherwise.
func (o *TemplateResponseDocumentFormFieldHyperlink) GetIsMultiline() bool {
	if o == nil || IsNil(o.IsMultiline) {
		var ret bool
		return ret
	}
	return *o.IsMultiline
}

// GetIsMultilineOk returns a tuple with the IsMultiline field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentFormFieldHyperlink) GetIsMultilineOk() (*bool, bool) {
	if o == nil || IsNil(o.IsMultiline) {
		return nil, false
	}
	return o.IsMultiline, true
}

// HasIsMultiline returns a boolean if a field has been set.
func (o *TemplateResponseDocumentFormFieldHyperlink) HasIsMultiline() bool {
	if o != nil && !IsNil(o.IsMultiline) {
		return true
	}

	return false
}

// SetIsMultiline gets a reference to the given bool and assigns it to the IsMultiline field.
func (o *TemplateResponseDocumentFormFieldHyperlink) SetIsMultiline(v bool) {
	o.IsMultiline = &v
}

// GetOriginalFontSize returns the OriginalFontSize field value if set, zero value otherwise.
func (o *TemplateResponseDocumentFormFieldHyperlink) GetOriginalFontSize() int32 {
	if o == nil || IsNil(o.OriginalFontSize) {
		var ret int32
		return ret
	}
	return *o.OriginalFontSize
}

// GetOriginalFontSizeOk returns a tuple with the OriginalFontSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentFormFieldHyperlink) GetOriginalFontSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.OriginalFontSize) {
		return nil, false
	}
	return o.OriginalFontSize, true
}

// HasOriginalFontSize returns a boolean if a field has been set.
func (o *TemplateResponseDocumentFormFieldHyperlink) HasOriginalFontSize() bool {
	if o != nil && !IsNil(o.OriginalFontSize) {
		return true
	}

	return false
}

// SetOriginalFontSize gets a reference to the given int32 and assigns it to the OriginalFontSize field.
func (o *TemplateResponseDocumentFormFieldHyperlink) SetOriginalFontSize(v int32) {
	o.OriginalFontSize = &v
}

// GetFontFamily returns the FontFamily field value if set, zero value otherwise.
func (o *TemplateResponseDocumentFormFieldHyperlink) GetFontFamily() string {
	if o == nil || IsNil(o.FontFamily) {
		var ret string
		return ret
	}
	return *o.FontFamily
}

// GetFontFamilyOk returns a tuple with the FontFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentFormFieldHyperlink) GetFontFamilyOk() (*string, bool) {
	if o == nil || IsNil(o.FontFamily) {
		return nil, false
	}
	return o.FontFamily, true
}

// HasFontFamily returns a boolean if a field has been set.
func (o *TemplateResponseDocumentFormFieldHyperlink) HasFontFamily() bool {
	if o != nil && !IsNil(o.FontFamily) {
		return true
	}

	return false
}

// SetFontFamily gets a reference to the given string and assigns it to the FontFamily field.
func (o *TemplateResponseDocumentFormFieldHyperlink) SetFontFamily(v string) {
	o.FontFamily = &v
}

func (o TemplateResponseDocumentFormFieldHyperlink) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateResponseDocumentFormFieldHyperlink) ToMap() (map[string]interface{}, error) {
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

type NullableTemplateResponseDocumentFormFieldHyperlink struct {
	value *TemplateResponseDocumentFormFieldHyperlink
	isSet bool
}

func (v NullableTemplateResponseDocumentFormFieldHyperlink) Get() *TemplateResponseDocumentFormFieldHyperlink {
	return v.value
}

func (v *NullableTemplateResponseDocumentFormFieldHyperlink) Set(val *TemplateResponseDocumentFormFieldHyperlink) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateResponseDocumentFormFieldHyperlink) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateResponseDocumentFormFieldHyperlink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateResponseDocumentFormFieldHyperlink(val *TemplateResponseDocumentFormFieldHyperlink) *NullableTemplateResponseDocumentFormFieldHyperlink {
	return &NullableTemplateResponseDocumentFormFieldHyperlink{value: val, isSet: true}
}

func (v NullableTemplateResponseDocumentFormFieldHyperlink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateResponseDocumentFormFieldHyperlink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


