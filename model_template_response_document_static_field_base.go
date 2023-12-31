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

// checks if the TemplateResponseDocumentStaticFieldBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateResponseDocumentStaticFieldBase{}

// TemplateResponseDocumentStaticFieldBase An array describing static overlay fields. **Note** only available for certain subscriptions.
type TemplateResponseDocumentStaticFieldBase struct {
	// A unique id for the static field.
	ApiId *string `json:"api_id,omitempty"`
	// The name of the static field.
	Name *string `json:"name,omitempty"`
	Type string `json:"type"`
	// The signer of the Static Field.
	Signer *string `json:"signer,omitempty"`
	// The horizontal offset in pixels for this static field.
	X *int32 `json:"x,omitempty"`
	// The vertical offset in pixels for this static field.
	Y *int32 `json:"y,omitempty"`
	// The width in pixels of this static field.
	Width *int32 `json:"width,omitempty"`
	// The height in pixels of this static field.
	Height *int32 `json:"height,omitempty"`
	// Boolean showing whether or not this field is required.
	Required *bool `json:"required,omitempty"`
	// The name of the group this field is in. If this field is not a group, this defaults to `null`.
	Group NullableString `json:"group,omitempty"`
}

// NewTemplateResponseDocumentStaticFieldBase instantiates a new TemplateResponseDocumentStaticFieldBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateResponseDocumentStaticFieldBase(type_ string) *TemplateResponseDocumentStaticFieldBase {
	this := TemplateResponseDocumentStaticFieldBase{}
	this.Type = type_
	var signer string = "me_now"
	this.Signer = &signer
	return &this
}

// NewTemplateResponseDocumentStaticFieldBaseWithDefaults instantiates a new TemplateResponseDocumentStaticFieldBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateResponseDocumentStaticFieldBaseWithDefaults() *TemplateResponseDocumentStaticFieldBase {
	this := TemplateResponseDocumentStaticFieldBase{}
	var signer string = "me_now"
	this.Signer = &signer
	return &this
}

// GetApiId returns the ApiId field value if set, zero value otherwise.
func (o *TemplateResponseDocumentStaticFieldBase) GetApiId() string {
	if o == nil || IsNil(o.ApiId) {
		var ret string
		return ret
	}
	return *o.ApiId
}

// GetApiIdOk returns a tuple with the ApiId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentStaticFieldBase) GetApiIdOk() (*string, bool) {
	if o == nil || IsNil(o.ApiId) {
		return nil, false
	}
	return o.ApiId, true
}

// HasApiId returns a boolean if a field has been set.
func (o *TemplateResponseDocumentStaticFieldBase) HasApiId() bool {
	if o != nil && !IsNil(o.ApiId) {
		return true
	}

	return false
}

// SetApiId gets a reference to the given string and assigns it to the ApiId field.
func (o *TemplateResponseDocumentStaticFieldBase) SetApiId(v string) {
	o.ApiId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplateResponseDocumentStaticFieldBase) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentStaticFieldBase) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplateResponseDocumentStaticFieldBase) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplateResponseDocumentStaticFieldBase) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value
func (o *TemplateResponseDocumentStaticFieldBase) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentStaticFieldBase) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *TemplateResponseDocumentStaticFieldBase) SetType(v string) {
	o.Type = v
}

// GetSigner returns the Signer field value if set, zero value otherwise.
func (o *TemplateResponseDocumentStaticFieldBase) GetSigner() string {
	if o == nil || IsNil(o.Signer) {
		var ret string
		return ret
	}
	return *o.Signer
}

// GetSignerOk returns a tuple with the Signer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentStaticFieldBase) GetSignerOk() (*string, bool) {
	if o == nil || IsNil(o.Signer) {
		return nil, false
	}
	return o.Signer, true
}

// HasSigner returns a boolean if a field has been set.
func (o *TemplateResponseDocumentStaticFieldBase) HasSigner() bool {
	if o != nil && !IsNil(o.Signer) {
		return true
	}

	return false
}

// SetSigner gets a reference to the given string and assigns it to the Signer field.
func (o *TemplateResponseDocumentStaticFieldBase) SetSigner(v string) {
	o.Signer = &v
}

// GetX returns the X field value if set, zero value otherwise.
func (o *TemplateResponseDocumentStaticFieldBase) GetX() int32 {
	if o == nil || IsNil(o.X) {
		var ret int32
		return ret
	}
	return *o.X
}

// GetXOk returns a tuple with the X field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentStaticFieldBase) GetXOk() (*int32, bool) {
	if o == nil || IsNil(o.X) {
		return nil, false
	}
	return o.X, true
}

// HasX returns a boolean if a field has been set.
func (o *TemplateResponseDocumentStaticFieldBase) HasX() bool {
	if o != nil && !IsNil(o.X) {
		return true
	}

	return false
}

// SetX gets a reference to the given int32 and assigns it to the X field.
func (o *TemplateResponseDocumentStaticFieldBase) SetX(v int32) {
	o.X = &v
}

// GetY returns the Y field value if set, zero value otherwise.
func (o *TemplateResponseDocumentStaticFieldBase) GetY() int32 {
	if o == nil || IsNil(o.Y) {
		var ret int32
		return ret
	}
	return *o.Y
}

// GetYOk returns a tuple with the Y field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentStaticFieldBase) GetYOk() (*int32, bool) {
	if o == nil || IsNil(o.Y) {
		return nil, false
	}
	return o.Y, true
}

// HasY returns a boolean if a field has been set.
func (o *TemplateResponseDocumentStaticFieldBase) HasY() bool {
	if o != nil && !IsNil(o.Y) {
		return true
	}

	return false
}

// SetY gets a reference to the given int32 and assigns it to the Y field.
func (o *TemplateResponseDocumentStaticFieldBase) SetY(v int32) {
	o.Y = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *TemplateResponseDocumentStaticFieldBase) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentStaticFieldBase) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *TemplateResponseDocumentStaticFieldBase) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *TemplateResponseDocumentStaticFieldBase) SetWidth(v int32) {
	o.Width = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *TemplateResponseDocumentStaticFieldBase) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentStaticFieldBase) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *TemplateResponseDocumentStaticFieldBase) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *TemplateResponseDocumentStaticFieldBase) SetHeight(v int32) {
	o.Height = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *TemplateResponseDocumentStaticFieldBase) GetRequired() bool {
	if o == nil || IsNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocumentStaticFieldBase) GetRequiredOk() (*bool, bool) {
	if o == nil || IsNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *TemplateResponseDocumentStaticFieldBase) HasRequired() bool {
	if o != nil && !IsNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *TemplateResponseDocumentStaticFieldBase) SetRequired(v bool) {
	o.Required = &v
}

// GetGroup returns the Group field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateResponseDocumentStaticFieldBase) GetGroup() string {
	if o == nil || IsNil(o.Group.Get()) {
		var ret string
		return ret
	}
	return *o.Group.Get()
}

// GetGroupOk returns a tuple with the Group field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateResponseDocumentStaticFieldBase) GetGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Group.Get(), o.Group.IsSet()
}

// HasGroup returns a boolean if a field has been set.
func (o *TemplateResponseDocumentStaticFieldBase) HasGroup() bool {
	if o != nil && o.Group.IsSet() {
		return true
	}

	return false
}

// SetGroup gets a reference to the given NullableString and assigns it to the Group field.
func (o *TemplateResponseDocumentStaticFieldBase) SetGroup(v string) {
	o.Group.Set(&v)
}
// SetGroupNil sets the value for Group to be an explicit nil
func (o *TemplateResponseDocumentStaticFieldBase) SetGroupNil() {
	o.Group.Set(nil)
}

// UnsetGroup ensures that no value is present for Group, not even an explicit nil
func (o *TemplateResponseDocumentStaticFieldBase) UnsetGroup() {
	o.Group.Unset()
}

func (o TemplateResponseDocumentStaticFieldBase) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateResponseDocumentStaticFieldBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ApiId) {
		toSerialize["api_id"] = o.ApiId
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["type"] = o.Type
	if !IsNil(o.Signer) {
		toSerialize["signer"] = o.Signer
	}
	if !IsNil(o.X) {
		toSerialize["x"] = o.X
	}
	if !IsNil(o.Y) {
		toSerialize["y"] = o.Y
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if o.Group.IsSet() {
		toSerialize["group"] = o.Group.Get()
	}
	return toSerialize, nil
}

type NullableTemplateResponseDocumentStaticFieldBase struct {
	value *TemplateResponseDocumentStaticFieldBase
	isSet bool
}

func (v NullableTemplateResponseDocumentStaticFieldBase) Get() *TemplateResponseDocumentStaticFieldBase {
	return v.value
}

func (v *NullableTemplateResponseDocumentStaticFieldBase) Set(val *TemplateResponseDocumentStaticFieldBase) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateResponseDocumentStaticFieldBase) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateResponseDocumentStaticFieldBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateResponseDocumentStaticFieldBase(val *TemplateResponseDocumentStaticFieldBase) *NullableTemplateResponseDocumentStaticFieldBase {
	return &NullableTemplateResponseDocumentStaticFieldBase{value: val, isSet: true}
}

func (v NullableTemplateResponseDocumentStaticFieldBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateResponseDocumentStaticFieldBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


