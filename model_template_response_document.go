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

// checks if the TemplateResponseDocument type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateResponseDocument{}

// TemplateResponseDocument struct for TemplateResponseDocument
type TemplateResponseDocument struct {
	// Name of the associated file.
	Name *string `json:"name,omitempty"`
	// Document ordering, the lowest index is displayed first and the highest last (0-based indexing).
	Index *int32 `json:"index,omitempty"`
	// An array of Form Field Group objects.
	FieldGroups []TemplateResponseDocumentFieldGroup `json:"field_groups,omitempty"`
	// An array of Form Field objects containing the name and type of each named field.
	FormFields []TemplateResponseDocumentFormFieldBase `json:"form_fields,omitempty"`
	// An array of Form Field objects containing the name and type of each named field.
	CustomFields []TemplateResponseDocumentCustomFieldBase `json:"custom_fields,omitempty"`
	// An array describing static overlay fields. **Note** only available for certain subscriptions.
	StaticFields []TemplateResponseDocumentStaticFieldBase `json:"static_fields,omitempty"`
}

// NewTemplateResponseDocument instantiates a new TemplateResponseDocument object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateResponseDocument() *TemplateResponseDocument {
	this := TemplateResponseDocument{}
	return &this
}

// NewTemplateResponseDocumentWithDefaults instantiates a new TemplateResponseDocument object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateResponseDocumentWithDefaults() *TemplateResponseDocument {
	this := TemplateResponseDocument{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TemplateResponseDocument) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocument) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TemplateResponseDocument) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TemplateResponseDocument) SetName(v string) {
	o.Name = &v
}

// GetIndex returns the Index field value if set, zero value otherwise.
func (o *TemplateResponseDocument) GetIndex() int32 {
	if o == nil || IsNil(o.Index) {
		var ret int32
		return ret
	}
	return *o.Index
}

// GetIndexOk returns a tuple with the Index field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocument) GetIndexOk() (*int32, bool) {
	if o == nil || IsNil(o.Index) {
		return nil, false
	}
	return o.Index, true
}

// HasIndex returns a boolean if a field has been set.
func (o *TemplateResponseDocument) HasIndex() bool {
	if o != nil && !IsNil(o.Index) {
		return true
	}

	return false
}

// SetIndex gets a reference to the given int32 and assigns it to the Index field.
func (o *TemplateResponseDocument) SetIndex(v int32) {
	o.Index = &v
}

// GetFieldGroups returns the FieldGroups field value if set, zero value otherwise.
func (o *TemplateResponseDocument) GetFieldGroups() []TemplateResponseDocumentFieldGroup {
	if o == nil || IsNil(o.FieldGroups) {
		var ret []TemplateResponseDocumentFieldGroup
		return ret
	}
	return o.FieldGroups
}

// GetFieldGroupsOk returns a tuple with the FieldGroups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocument) GetFieldGroupsOk() ([]TemplateResponseDocumentFieldGroup, bool) {
	if o == nil || IsNil(o.FieldGroups) {
		return nil, false
	}
	return o.FieldGroups, true
}

// HasFieldGroups returns a boolean if a field has been set.
func (o *TemplateResponseDocument) HasFieldGroups() bool {
	if o != nil && !IsNil(o.FieldGroups) {
		return true
	}

	return false
}

// SetFieldGroups gets a reference to the given []TemplateResponseDocumentFieldGroup and assigns it to the FieldGroups field.
func (o *TemplateResponseDocument) SetFieldGroups(v []TemplateResponseDocumentFieldGroup) {
	o.FieldGroups = v
}

// GetFormFields returns the FormFields field value if set, zero value otherwise.
func (o *TemplateResponseDocument) GetFormFields() []TemplateResponseDocumentFormFieldBase {
	if o == nil || IsNil(o.FormFields) {
		var ret []TemplateResponseDocumentFormFieldBase
		return ret
	}
	return o.FormFields
}

// GetFormFieldsOk returns a tuple with the FormFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocument) GetFormFieldsOk() ([]TemplateResponseDocumentFormFieldBase, bool) {
	if o == nil || IsNil(o.FormFields) {
		return nil, false
	}
	return o.FormFields, true
}

// HasFormFields returns a boolean if a field has been set.
func (o *TemplateResponseDocument) HasFormFields() bool {
	if o != nil && !IsNil(o.FormFields) {
		return true
	}

	return false
}

// SetFormFields gets a reference to the given []TemplateResponseDocumentFormFieldBase and assigns it to the FormFields field.
func (o *TemplateResponseDocument) SetFormFields(v []TemplateResponseDocumentFormFieldBase) {
	o.FormFields = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *TemplateResponseDocument) GetCustomFields() []TemplateResponseDocumentCustomFieldBase {
	if o == nil || IsNil(o.CustomFields) {
		var ret []TemplateResponseDocumentCustomFieldBase
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseDocument) GetCustomFieldsOk() ([]TemplateResponseDocumentCustomFieldBase, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return nil, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *TemplateResponseDocument) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given []TemplateResponseDocumentCustomFieldBase and assigns it to the CustomFields field.
func (o *TemplateResponseDocument) SetCustomFields(v []TemplateResponseDocumentCustomFieldBase) {
	o.CustomFields = v
}

// GetStaticFields returns the StaticFields field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *TemplateResponseDocument) GetStaticFields() []TemplateResponseDocumentStaticFieldBase {
	if o == nil {
		var ret []TemplateResponseDocumentStaticFieldBase
		return ret
	}
	return o.StaticFields
}

// GetStaticFieldsOk returns a tuple with the StaticFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *TemplateResponseDocument) GetStaticFieldsOk() ([]TemplateResponseDocumentStaticFieldBase, bool) {
	if o == nil || IsNil(o.StaticFields) {
		return nil, false
	}
	return o.StaticFields, true
}

// HasStaticFields returns a boolean if a field has been set.
func (o *TemplateResponseDocument) HasStaticFields() bool {
	if o != nil && IsNil(o.StaticFields) {
		return true
	}

	return false
}

// SetStaticFields gets a reference to the given []TemplateResponseDocumentStaticFieldBase and assigns it to the StaticFields field.
func (o *TemplateResponseDocument) SetStaticFields(v []TemplateResponseDocumentStaticFieldBase) {
	o.StaticFields = v
}

func (o TemplateResponseDocument) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateResponseDocument) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Index) {
		toSerialize["index"] = o.Index
	}
	if !IsNil(o.FieldGroups) {
		toSerialize["field_groups"] = o.FieldGroups
	}
	if !IsNil(o.FormFields) {
		toSerialize["form_fields"] = o.FormFields
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if o.StaticFields != nil {
		toSerialize["static_fields"] = o.StaticFields
	}
	return toSerialize, nil
}

type NullableTemplateResponseDocument struct {
	value *TemplateResponseDocument
	isSet bool
}

func (v NullableTemplateResponseDocument) Get() *TemplateResponseDocument {
	return v.value
}

func (v *NullableTemplateResponseDocument) Set(val *TemplateResponseDocument) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateResponseDocument) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateResponseDocument) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateResponseDocument(val *TemplateResponseDocument) *NullableTemplateResponseDocument {
	return &NullableTemplateResponseDocument{value: val, isSet: true}
}

func (v NullableTemplateResponseDocument) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateResponseDocument) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


