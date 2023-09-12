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

// checks if the SubFormFieldsPerDocumentTextMerge type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubFormFieldsPerDocumentTextMerge{}

// SubFormFieldsPerDocumentTextMerge This class extends `SubFormFieldsPerDocumentBase`.
type SubFormFieldsPerDocumentTextMerge struct {
	SubFormFieldsPerDocumentBase
	// A text field that has default text set using pre-filled data. Use the `SubFormFieldsPerDocumentTextMerge` class.
	Type string `json:"type"`
	// Font family for the field.
	FontFamily *string `json:"font_family,omitempty"`
	// The initial px font size for the field contents. Can be any integer value between `7` and `49`.  **NOTE**: Font size may be reduced during processing in order to fit the contents within the dimensions of the field.
	FontSize *int32 `json:"font_size,omitempty"`
}

// NewSubFormFieldsPerDocumentTextMerge instantiates a new SubFormFieldsPerDocumentTextMerge object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubFormFieldsPerDocumentTextMerge(type_ string, documentIndex int32, apiId string, height int32, required bool, signer string, width int32, x int32, y int32) *SubFormFieldsPerDocumentTextMerge {
	this := SubFormFieldsPerDocumentTextMerge{}
	this.DocumentIndex = documentIndex
	this.ApiId = apiId
	this.Height = height
	this.Required = required
	this.Signer = signer
	this.Type = type_
	this.Width = width
	this.X = x
	this.Y = y
	return &this
}

// NewSubFormFieldsPerDocumentTextMergeWithDefaults instantiates a new SubFormFieldsPerDocumentTextMerge object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubFormFieldsPerDocumentTextMergeWithDefaults() *SubFormFieldsPerDocumentTextMerge {
	this := SubFormFieldsPerDocumentTextMerge{}
	var type_ string = "text-merge"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *SubFormFieldsPerDocumentTextMerge) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentTextMerge) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *SubFormFieldsPerDocumentTextMerge) SetType(v string) {
	o.Type = v
}

// GetFontFamily returns the FontFamily field value if set, zero value otherwise.
func (o *SubFormFieldsPerDocumentTextMerge) GetFontFamily() string {
	if o == nil || IsNil(o.FontFamily) {
		var ret string
		return ret
	}
	return *o.FontFamily
}

// GetFontFamilyOk returns a tuple with the FontFamily field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentTextMerge) GetFontFamilyOk() (*string, bool) {
	if o == nil || IsNil(o.FontFamily) {
		return nil, false
	}
	return o.FontFamily, true
}

// HasFontFamily returns a boolean if a field has been set.
func (o *SubFormFieldsPerDocumentTextMerge) HasFontFamily() bool {
	if o != nil && !IsNil(o.FontFamily) {
		return true
	}

	return false
}

// SetFontFamily gets a reference to the given string and assigns it to the FontFamily field.
func (o *SubFormFieldsPerDocumentTextMerge) SetFontFamily(v string) {
	o.FontFamily = &v
}

// GetFontSize returns the FontSize field value if set, zero value otherwise.
func (o *SubFormFieldsPerDocumentTextMerge) GetFontSize() int32 {
	if o == nil || IsNil(o.FontSize) {
		var ret int32
		return ret
	}
	return *o.FontSize
}

// GetFontSizeOk returns a tuple with the FontSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubFormFieldsPerDocumentTextMerge) GetFontSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.FontSize) {
		return nil, false
	}
	return o.FontSize, true
}

// HasFontSize returns a boolean if a field has been set.
func (o *SubFormFieldsPerDocumentTextMerge) HasFontSize() bool {
	if o != nil && !IsNil(o.FontSize) {
		return true
	}

	return false
}

// SetFontSize gets a reference to the given int32 and assigns it to the FontSize field.
func (o *SubFormFieldsPerDocumentTextMerge) SetFontSize(v int32) {
	o.FontSize = &v
}

func (o SubFormFieldsPerDocumentTextMerge) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubFormFieldsPerDocumentTextMerge) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	serializedSubFormFieldsPerDocumentBase, errSubFormFieldsPerDocumentBase := json.Marshal(o.SubFormFieldsPerDocumentBase)
	if errSubFormFieldsPerDocumentBase != nil {
		return map[string]interface{}{}, errSubFormFieldsPerDocumentBase
	}
	errSubFormFieldsPerDocumentBase = json.Unmarshal([]byte(serializedSubFormFieldsPerDocumentBase), &toSerialize)
	if errSubFormFieldsPerDocumentBase != nil {
		return map[string]interface{}{}, errSubFormFieldsPerDocumentBase
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.FontFamily) {
		toSerialize["font_family"] = o.FontFamily
	}
	if !IsNil(o.FontSize) {
		toSerialize["font_size"] = o.FontSize
	}
	return toSerialize, nil
}

type NullableSubFormFieldsPerDocumentTextMerge struct {
	value *SubFormFieldsPerDocumentTextMerge
	isSet bool
}

func (v NullableSubFormFieldsPerDocumentTextMerge) Get() *SubFormFieldsPerDocumentTextMerge {
	return v.value
}

func (v *NullableSubFormFieldsPerDocumentTextMerge) Set(val *SubFormFieldsPerDocumentTextMerge) {
	v.value = val
	v.isSet = true
}

func (v NullableSubFormFieldsPerDocumentTextMerge) IsSet() bool {
	return v.isSet
}

func (v *NullableSubFormFieldsPerDocumentTextMerge) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubFormFieldsPerDocumentTextMerge(val *SubFormFieldsPerDocumentTextMerge) *NullableSubFormFieldsPerDocumentTextMerge {
	return &NullableSubFormFieldsPerDocumentTextMerge{value: val, isSet: true}
}

func (v NullableSubFormFieldsPerDocumentTextMerge) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubFormFieldsPerDocumentTextMerge) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


