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

// checks if the SubWhiteLabelingOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubWhiteLabelingOptions{}

// SubWhiteLabelingOptions An array of elements and values serialized to a string, to be used to customize the app's signer page. (Only applies to some API plans)  Take a look at our [white labeling guide](https://developers.hellosign.com/api/reference/premium-branding/) to learn more.
type SubWhiteLabelingOptions struct {
	HeaderBackgroundColor *string `json:"header_background_color,omitempty"`
	LegalVersion *string `json:"legal_version,omitempty"`
	LinkColor *string `json:"link_color,omitempty"`
	PageBackgroundColor *string `json:"page_background_color,omitempty"`
	PrimaryButtonColor *string `json:"primary_button_color,omitempty"`
	PrimaryButtonColorHover *string `json:"primary_button_color_hover,omitempty"`
	PrimaryButtonTextColor *string `json:"primary_button_text_color,omitempty"`
	PrimaryButtonTextColorHover *string `json:"primary_button_text_color_hover,omitempty"`
	SecondaryButtonColor *string `json:"secondary_button_color,omitempty"`
	SecondaryButtonColorHover *string `json:"secondary_button_color_hover,omitempty"`
	SecondaryButtonTextColor *string `json:"secondary_button_text_color,omitempty"`
	SecondaryButtonTextColorHover *string `json:"secondary_button_text_color_hover,omitempty"`
	TextColor1 *string `json:"text_color1,omitempty"`
	TextColor2 *string `json:"text_color2,omitempty"`
	// Resets white labeling options to defaults. Only useful when updating an API App.
	ResetToDefault *bool `json:"reset_to_default,omitempty"`
}

// NewSubWhiteLabelingOptions instantiates a new SubWhiteLabelingOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubWhiteLabelingOptions() *SubWhiteLabelingOptions {
	this := SubWhiteLabelingOptions{}
	var headerBackgroundColor string = "#1A1A1A"
	this.HeaderBackgroundColor = &headerBackgroundColor
	var legalVersion string = "terms1"
	this.LegalVersion = &legalVersion
	var linkColor string = "#00B3E6"
	this.LinkColor = &linkColor
	var pageBackgroundColor string = "#F7F8F9"
	this.PageBackgroundColor = &pageBackgroundColor
	var primaryButtonColor string = "#00B3E6"
	this.PrimaryButtonColor = &primaryButtonColor
	var primaryButtonColorHover string = "#00B3E6"
	this.PrimaryButtonColorHover = &primaryButtonColorHover
	var primaryButtonTextColor string = "#FFFFFF"
	this.PrimaryButtonTextColor = &primaryButtonTextColor
	var primaryButtonTextColorHover string = "#FFFFFF"
	this.PrimaryButtonTextColorHover = &primaryButtonTextColorHover
	var secondaryButtonColor string = "#FFFFFF"
	this.SecondaryButtonColor = &secondaryButtonColor
	var secondaryButtonColorHover string = "#FFFFFF"
	this.SecondaryButtonColorHover = &secondaryButtonColorHover
	var secondaryButtonTextColor string = "#00B3E6"
	this.SecondaryButtonTextColor = &secondaryButtonTextColor
	var secondaryButtonTextColorHover string = "#00B3E6"
	this.SecondaryButtonTextColorHover = &secondaryButtonTextColorHover
	var textColor1 string = "#808080"
	this.TextColor1 = &textColor1
	var textColor2 string = "#FFFFFF"
	this.TextColor2 = &textColor2
	return &this
}

// NewSubWhiteLabelingOptionsWithDefaults instantiates a new SubWhiteLabelingOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubWhiteLabelingOptionsWithDefaults() *SubWhiteLabelingOptions {
	this := SubWhiteLabelingOptions{}
	var headerBackgroundColor string = "#1A1A1A"
	this.HeaderBackgroundColor = &headerBackgroundColor
	var legalVersion string = "terms1"
	this.LegalVersion = &legalVersion
	var linkColor string = "#00B3E6"
	this.LinkColor = &linkColor
	var pageBackgroundColor string = "#F7F8F9"
	this.PageBackgroundColor = &pageBackgroundColor
	var primaryButtonColor string = "#00B3E6"
	this.PrimaryButtonColor = &primaryButtonColor
	var primaryButtonColorHover string = "#00B3E6"
	this.PrimaryButtonColorHover = &primaryButtonColorHover
	var primaryButtonTextColor string = "#FFFFFF"
	this.PrimaryButtonTextColor = &primaryButtonTextColor
	var primaryButtonTextColorHover string = "#FFFFFF"
	this.PrimaryButtonTextColorHover = &primaryButtonTextColorHover
	var secondaryButtonColor string = "#FFFFFF"
	this.SecondaryButtonColor = &secondaryButtonColor
	var secondaryButtonColorHover string = "#FFFFFF"
	this.SecondaryButtonColorHover = &secondaryButtonColorHover
	var secondaryButtonTextColor string = "#00B3E6"
	this.SecondaryButtonTextColor = &secondaryButtonTextColor
	var secondaryButtonTextColorHover string = "#00B3E6"
	this.SecondaryButtonTextColorHover = &secondaryButtonTextColorHover
	var textColor1 string = "#808080"
	this.TextColor1 = &textColor1
	var textColor2 string = "#FFFFFF"
	this.TextColor2 = &textColor2
	return &this
}

// GetHeaderBackgroundColor returns the HeaderBackgroundColor field value if set, zero value otherwise.
func (o *SubWhiteLabelingOptions) GetHeaderBackgroundColor() string {
	if o == nil || IsNil(o.HeaderBackgroundColor) {
		var ret string
		return ret
	}
	return *o.HeaderBackgroundColor
}

// GetHeaderBackgroundColorOk returns a tuple with the HeaderBackgroundColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubWhiteLabelingOptions) GetHeaderBackgroundColorOk() (*string, bool) {
	if o == nil || IsNil(o.HeaderBackgroundColor) {
		return nil, false
	}
	return o.HeaderBackgroundColor, true
}

// HasHeaderBackgroundColor returns a boolean if a field has been set.
func (o *SubWhiteLabelingOptions) HasHeaderBackgroundColor() bool {
	if o != nil && !IsNil(o.HeaderBackgroundColor) {
		return true
	}

	return false
}

// SetHeaderBackgroundColor gets a reference to the given string and assigns it to the HeaderBackgroundColor field.
func (o *SubWhiteLabelingOptions) SetHeaderBackgroundColor(v string) {
	o.HeaderBackgroundColor = &v
}

// GetLegalVersion returns the LegalVersion field value if set, zero value otherwise.
func (o *SubWhiteLabelingOptions) GetLegalVersion() string {
	if o == nil || IsNil(o.LegalVersion) {
		var ret string
		return ret
	}
	return *o.LegalVersion
}

// GetLegalVersionOk returns a tuple with the LegalVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubWhiteLabelingOptions) GetLegalVersionOk() (*string, bool) {
	if o == nil || IsNil(o.LegalVersion) {
		return nil, false
	}
	return o.LegalVersion, true
}

// HasLegalVersion returns a boolean if a field has been set.
func (o *SubWhiteLabelingOptions) HasLegalVersion() bool {
	if o != nil && !IsNil(o.LegalVersion) {
		return true
	}

	return false
}

// SetLegalVersion gets a reference to the given string and assigns it to the LegalVersion field.
func (o *SubWhiteLabelingOptions) SetLegalVersion(v string) {
	o.LegalVersion = &v
}

// GetLinkColor returns the LinkColor field value if set, zero value otherwise.
func (o *SubWhiteLabelingOptions) GetLinkColor() string {
	if o == nil || IsNil(o.LinkColor) {
		var ret string
		return ret
	}
	return *o.LinkColor
}

// GetLinkColorOk returns a tuple with the LinkColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubWhiteLabelingOptions) GetLinkColorOk() (*string, bool) {
	if o == nil || IsNil(o.LinkColor) {
		return nil, false
	}
	return o.LinkColor, true
}

// HasLinkColor returns a boolean if a field has been set.
func (o *SubWhiteLabelingOptions) HasLinkColor() bool {
	if o != nil && !IsNil(o.LinkColor) {
		return true
	}

	return false
}

// SetLinkColor gets a reference to the given string and assigns it to the LinkColor field.
func (o *SubWhiteLabelingOptions) SetLinkColor(v string) {
	o.LinkColor = &v
}

// GetPageBackgroundColor returns the PageBackgroundColor field value if set, zero value otherwise.
func (o *SubWhiteLabelingOptions) GetPageBackgroundColor() string {
	if o == nil || IsNil(o.PageBackgroundColor) {
		var ret string
		return ret
	}
	return *o.PageBackgroundColor
}

// GetPageBackgroundColorOk returns a tuple with the PageBackgroundColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubWhiteLabelingOptions) GetPageBackgroundColorOk() (*string, bool) {
	if o == nil || IsNil(o.PageBackgroundColor) {
		return nil, false
	}
	return o.PageBackgroundColor, true
}

// HasPageBackgroundColor returns a boolean if a field has been set.
func (o *SubWhiteLabelingOptions) HasPageBackgroundColor() bool {
	if o != nil && !IsNil(o.PageBackgroundColor) {
		return true
	}

	return false
}

// SetPageBackgroundColor gets a reference to the given string and assigns it to the PageBackgroundColor field.
func (o *SubWhiteLabelingOptions) SetPageBackgroundColor(v string) {
	o.PageBackgroundColor = &v
}

// GetPrimaryButtonColor returns the PrimaryButtonColor field value if set, zero value otherwise.
func (o *SubWhiteLabelingOptions) GetPrimaryButtonColor() string {
	if o == nil || IsNil(o.PrimaryButtonColor) {
		var ret string
		return ret
	}
	return *o.PrimaryButtonColor
}

// GetPrimaryButtonColorOk returns a tuple with the PrimaryButtonColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubWhiteLabelingOptions) GetPrimaryButtonColorOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryButtonColor) {
		return nil, false
	}
	return o.PrimaryButtonColor, true
}

// HasPrimaryButtonColor returns a boolean if a field has been set.
func (o *SubWhiteLabelingOptions) HasPrimaryButtonColor() bool {
	if o != nil && !IsNil(o.PrimaryButtonColor) {
		return true
	}

	return false
}

// SetPrimaryButtonColor gets a reference to the given string and assigns it to the PrimaryButtonColor field.
func (o *SubWhiteLabelingOptions) SetPrimaryButtonColor(v string) {
	o.PrimaryButtonColor = &v
}

// GetPrimaryButtonColorHover returns the PrimaryButtonColorHover field value if set, zero value otherwise.
func (o *SubWhiteLabelingOptions) GetPrimaryButtonColorHover() string {
	if o == nil || IsNil(o.PrimaryButtonColorHover) {
		var ret string
		return ret
	}
	return *o.PrimaryButtonColorHover
}

// GetPrimaryButtonColorHoverOk returns a tuple with the PrimaryButtonColorHover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubWhiteLabelingOptions) GetPrimaryButtonColorHoverOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryButtonColorHover) {
		return nil, false
	}
	return o.PrimaryButtonColorHover, true
}

// HasPrimaryButtonColorHover returns a boolean if a field has been set.
func (o *SubWhiteLabelingOptions) HasPrimaryButtonColorHover() bool {
	if o != nil && !IsNil(o.PrimaryButtonColorHover) {
		return true
	}

	return false
}

// SetPrimaryButtonColorHover gets a reference to the given string and assigns it to the PrimaryButtonColorHover field.
func (o *SubWhiteLabelingOptions) SetPrimaryButtonColorHover(v string) {
	o.PrimaryButtonColorHover = &v
}

// GetPrimaryButtonTextColor returns the PrimaryButtonTextColor field value if set, zero value otherwise.
func (o *SubWhiteLabelingOptions) GetPrimaryButtonTextColor() string {
	if o == nil || IsNil(o.PrimaryButtonTextColor) {
		var ret string
		return ret
	}
	return *o.PrimaryButtonTextColor
}

// GetPrimaryButtonTextColorOk returns a tuple with the PrimaryButtonTextColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubWhiteLabelingOptions) GetPrimaryButtonTextColorOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryButtonTextColor) {
		return nil, false
	}
	return o.PrimaryButtonTextColor, true
}

// HasPrimaryButtonTextColor returns a boolean if a field has been set.
func (o *SubWhiteLabelingOptions) HasPrimaryButtonTextColor() bool {
	if o != nil && !IsNil(o.PrimaryButtonTextColor) {
		return true
	}

	return false
}

// SetPrimaryButtonTextColor gets a reference to the given string and assigns it to the PrimaryButtonTextColor field.
func (o *SubWhiteLabelingOptions) SetPrimaryButtonTextColor(v string) {
	o.PrimaryButtonTextColor = &v
}

// GetPrimaryButtonTextColorHover returns the PrimaryButtonTextColorHover field value if set, zero value otherwise.
func (o *SubWhiteLabelingOptions) GetPrimaryButtonTextColorHover() string {
	if o == nil || IsNil(o.PrimaryButtonTextColorHover) {
		var ret string
		return ret
	}
	return *o.PrimaryButtonTextColorHover
}

// GetPrimaryButtonTextColorHoverOk returns a tuple with the PrimaryButtonTextColorHover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubWhiteLabelingOptions) GetPrimaryButtonTextColorHoverOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryButtonTextColorHover) {
		return nil, false
	}
	return o.PrimaryButtonTextColorHover, true
}

// HasPrimaryButtonTextColorHover returns a boolean if a field has been set.
func (o *SubWhiteLabelingOptions) HasPrimaryButtonTextColorHover() bool {
	if o != nil && !IsNil(o.PrimaryButtonTextColorHover) {
		return true
	}

	return false
}

// SetPrimaryButtonTextColorHover gets a reference to the given string and assigns it to the PrimaryButtonTextColorHover field.
func (o *SubWhiteLabelingOptions) SetPrimaryButtonTextColorHover(v string) {
	o.PrimaryButtonTextColorHover = &v
}

// GetSecondaryButtonColor returns the SecondaryButtonColor field value if set, zero value otherwise.
func (o *SubWhiteLabelingOptions) GetSecondaryButtonColor() string {
	if o == nil || IsNil(o.SecondaryButtonColor) {
		var ret string
		return ret
	}
	return *o.SecondaryButtonColor
}

// GetSecondaryButtonColorOk returns a tuple with the SecondaryButtonColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubWhiteLabelingOptions) GetSecondaryButtonColorOk() (*string, bool) {
	if o == nil || IsNil(o.SecondaryButtonColor) {
		return nil, false
	}
	return o.SecondaryButtonColor, true
}

// HasSecondaryButtonColor returns a boolean if a field has been set.
func (o *SubWhiteLabelingOptions) HasSecondaryButtonColor() bool {
	if o != nil && !IsNil(o.SecondaryButtonColor) {
		return true
	}

	return false
}

// SetSecondaryButtonColor gets a reference to the given string and assigns it to the SecondaryButtonColor field.
func (o *SubWhiteLabelingOptions) SetSecondaryButtonColor(v string) {
	o.SecondaryButtonColor = &v
}

// GetSecondaryButtonColorHover returns the SecondaryButtonColorHover field value if set, zero value otherwise.
func (o *SubWhiteLabelingOptions) GetSecondaryButtonColorHover() string {
	if o == nil || IsNil(o.SecondaryButtonColorHover) {
		var ret string
		return ret
	}
	return *o.SecondaryButtonColorHover
}

// GetSecondaryButtonColorHoverOk returns a tuple with the SecondaryButtonColorHover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubWhiteLabelingOptions) GetSecondaryButtonColorHoverOk() (*string, bool) {
	if o == nil || IsNil(o.SecondaryButtonColorHover) {
		return nil, false
	}
	return o.SecondaryButtonColorHover, true
}

// HasSecondaryButtonColorHover returns a boolean if a field has been set.
func (o *SubWhiteLabelingOptions) HasSecondaryButtonColorHover() bool {
	if o != nil && !IsNil(o.SecondaryButtonColorHover) {
		return true
	}

	return false
}

// SetSecondaryButtonColorHover gets a reference to the given string and assigns it to the SecondaryButtonColorHover field.
func (o *SubWhiteLabelingOptions) SetSecondaryButtonColorHover(v string) {
	o.SecondaryButtonColorHover = &v
}

// GetSecondaryButtonTextColor returns the SecondaryButtonTextColor field value if set, zero value otherwise.
func (o *SubWhiteLabelingOptions) GetSecondaryButtonTextColor() string {
	if o == nil || IsNil(o.SecondaryButtonTextColor) {
		var ret string
		return ret
	}
	return *o.SecondaryButtonTextColor
}

// GetSecondaryButtonTextColorOk returns a tuple with the SecondaryButtonTextColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubWhiteLabelingOptions) GetSecondaryButtonTextColorOk() (*string, bool) {
	if o == nil || IsNil(o.SecondaryButtonTextColor) {
		return nil, false
	}
	return o.SecondaryButtonTextColor, true
}

// HasSecondaryButtonTextColor returns a boolean if a field has been set.
func (o *SubWhiteLabelingOptions) HasSecondaryButtonTextColor() bool {
	if o != nil && !IsNil(o.SecondaryButtonTextColor) {
		return true
	}

	return false
}

// SetSecondaryButtonTextColor gets a reference to the given string and assigns it to the SecondaryButtonTextColor field.
func (o *SubWhiteLabelingOptions) SetSecondaryButtonTextColor(v string) {
	o.SecondaryButtonTextColor = &v
}

// GetSecondaryButtonTextColorHover returns the SecondaryButtonTextColorHover field value if set, zero value otherwise.
func (o *SubWhiteLabelingOptions) GetSecondaryButtonTextColorHover() string {
	if o == nil || IsNil(o.SecondaryButtonTextColorHover) {
		var ret string
		return ret
	}
	return *o.SecondaryButtonTextColorHover
}

// GetSecondaryButtonTextColorHoverOk returns a tuple with the SecondaryButtonTextColorHover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubWhiteLabelingOptions) GetSecondaryButtonTextColorHoverOk() (*string, bool) {
	if o == nil || IsNil(o.SecondaryButtonTextColorHover) {
		return nil, false
	}
	return o.SecondaryButtonTextColorHover, true
}

// HasSecondaryButtonTextColorHover returns a boolean if a field has been set.
func (o *SubWhiteLabelingOptions) HasSecondaryButtonTextColorHover() bool {
	if o != nil && !IsNil(o.SecondaryButtonTextColorHover) {
		return true
	}

	return false
}

// SetSecondaryButtonTextColorHover gets a reference to the given string and assigns it to the SecondaryButtonTextColorHover field.
func (o *SubWhiteLabelingOptions) SetSecondaryButtonTextColorHover(v string) {
	o.SecondaryButtonTextColorHover = &v
}

// GetTextColor1 returns the TextColor1 field value if set, zero value otherwise.
func (o *SubWhiteLabelingOptions) GetTextColor1() string {
	if o == nil || IsNil(o.TextColor1) {
		var ret string
		return ret
	}
	return *o.TextColor1
}

// GetTextColor1Ok returns a tuple with the TextColor1 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubWhiteLabelingOptions) GetTextColor1Ok() (*string, bool) {
	if o == nil || IsNil(o.TextColor1) {
		return nil, false
	}
	return o.TextColor1, true
}

// HasTextColor1 returns a boolean if a field has been set.
func (o *SubWhiteLabelingOptions) HasTextColor1() bool {
	if o != nil && !IsNil(o.TextColor1) {
		return true
	}

	return false
}

// SetTextColor1 gets a reference to the given string and assigns it to the TextColor1 field.
func (o *SubWhiteLabelingOptions) SetTextColor1(v string) {
	o.TextColor1 = &v
}

// GetTextColor2 returns the TextColor2 field value if set, zero value otherwise.
func (o *SubWhiteLabelingOptions) GetTextColor2() string {
	if o == nil || IsNil(o.TextColor2) {
		var ret string
		return ret
	}
	return *o.TextColor2
}

// GetTextColor2Ok returns a tuple with the TextColor2 field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubWhiteLabelingOptions) GetTextColor2Ok() (*string, bool) {
	if o == nil || IsNil(o.TextColor2) {
		return nil, false
	}
	return o.TextColor2, true
}

// HasTextColor2 returns a boolean if a field has been set.
func (o *SubWhiteLabelingOptions) HasTextColor2() bool {
	if o != nil && !IsNil(o.TextColor2) {
		return true
	}

	return false
}

// SetTextColor2 gets a reference to the given string and assigns it to the TextColor2 field.
func (o *SubWhiteLabelingOptions) SetTextColor2(v string) {
	o.TextColor2 = &v
}

// GetResetToDefault returns the ResetToDefault field value if set, zero value otherwise.
func (o *SubWhiteLabelingOptions) GetResetToDefault() bool {
	if o == nil || IsNil(o.ResetToDefault) {
		var ret bool
		return ret
	}
	return *o.ResetToDefault
}

// GetResetToDefaultOk returns a tuple with the ResetToDefault field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubWhiteLabelingOptions) GetResetToDefaultOk() (*bool, bool) {
	if o == nil || IsNil(o.ResetToDefault) {
		return nil, false
	}
	return o.ResetToDefault, true
}

// HasResetToDefault returns a boolean if a field has been set.
func (o *SubWhiteLabelingOptions) HasResetToDefault() bool {
	if o != nil && !IsNil(o.ResetToDefault) {
		return true
	}

	return false
}

// SetResetToDefault gets a reference to the given bool and assigns it to the ResetToDefault field.
func (o *SubWhiteLabelingOptions) SetResetToDefault(v bool) {
	o.ResetToDefault = &v
}

func (o SubWhiteLabelingOptions) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubWhiteLabelingOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.HeaderBackgroundColor) {
		toSerialize["header_background_color"] = o.HeaderBackgroundColor
	}
	if !IsNil(o.LegalVersion) {
		toSerialize["legal_version"] = o.LegalVersion
	}
	if !IsNil(o.LinkColor) {
		toSerialize["link_color"] = o.LinkColor
	}
	if !IsNil(o.PageBackgroundColor) {
		toSerialize["page_background_color"] = o.PageBackgroundColor
	}
	if !IsNil(o.PrimaryButtonColor) {
		toSerialize["primary_button_color"] = o.PrimaryButtonColor
	}
	if !IsNil(o.PrimaryButtonColorHover) {
		toSerialize["primary_button_color_hover"] = o.PrimaryButtonColorHover
	}
	if !IsNil(o.PrimaryButtonTextColor) {
		toSerialize["primary_button_text_color"] = o.PrimaryButtonTextColor
	}
	if !IsNil(o.PrimaryButtonTextColorHover) {
		toSerialize["primary_button_text_color_hover"] = o.PrimaryButtonTextColorHover
	}
	if !IsNil(o.SecondaryButtonColor) {
		toSerialize["secondary_button_color"] = o.SecondaryButtonColor
	}
	if !IsNil(o.SecondaryButtonColorHover) {
		toSerialize["secondary_button_color_hover"] = o.SecondaryButtonColorHover
	}
	if !IsNil(o.SecondaryButtonTextColor) {
		toSerialize["secondary_button_text_color"] = o.SecondaryButtonTextColor
	}
	if !IsNil(o.SecondaryButtonTextColorHover) {
		toSerialize["secondary_button_text_color_hover"] = o.SecondaryButtonTextColorHover
	}
	if !IsNil(o.TextColor1) {
		toSerialize["text_color1"] = o.TextColor1
	}
	if !IsNil(o.TextColor2) {
		toSerialize["text_color2"] = o.TextColor2
	}
	if !IsNil(o.ResetToDefault) {
		toSerialize["reset_to_default"] = o.ResetToDefault
	}
	return toSerialize, nil
}

type NullableSubWhiteLabelingOptions struct {
	value *SubWhiteLabelingOptions
	isSet bool
}

func (v NullableSubWhiteLabelingOptions) Get() *SubWhiteLabelingOptions {
	return v.value
}

func (v *NullableSubWhiteLabelingOptions) Set(val *SubWhiteLabelingOptions) {
	v.value = val
	v.isSet = true
}

func (v NullableSubWhiteLabelingOptions) IsSet() bool {
	return v.isSet
}

func (v *NullableSubWhiteLabelingOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubWhiteLabelingOptions(val *SubWhiteLabelingOptions) *NullableSubWhiteLabelingOptions {
	return &NullableSubWhiteLabelingOptions{value: val, isSet: true}
}

func (v NullableSubWhiteLabelingOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubWhiteLabelingOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


