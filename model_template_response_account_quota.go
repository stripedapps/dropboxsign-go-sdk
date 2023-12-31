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

// checks if the TemplateResponseAccountQuota type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TemplateResponseAccountQuota{}

// TemplateResponseAccountQuota An array of the designated CC roles that must be specified when sending a SignatureRequest using this Template.
type TemplateResponseAccountQuota struct {
	// API templates remaining.
	TemplatesLeft *int32 `json:"templates_left,omitempty"`
	// API signature requests remaining.
	ApiSignatureRequestsLeft *int32 `json:"api_signature_requests_left,omitempty"`
	// Signature requests remaining.
	DocumentsLeft *int32 `json:"documents_left,omitempty"`
	// SMS verifications remaining.
	SmsVerificationsLeft *int32 `json:"sms_verifications_left,omitempty"`
}

// NewTemplateResponseAccountQuota instantiates a new TemplateResponseAccountQuota object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTemplateResponseAccountQuota() *TemplateResponseAccountQuota {
	this := TemplateResponseAccountQuota{}
	return &this
}

// NewTemplateResponseAccountQuotaWithDefaults instantiates a new TemplateResponseAccountQuota object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTemplateResponseAccountQuotaWithDefaults() *TemplateResponseAccountQuota {
	this := TemplateResponseAccountQuota{}
	return &this
}

// GetTemplatesLeft returns the TemplatesLeft field value if set, zero value otherwise.
func (o *TemplateResponseAccountQuota) GetTemplatesLeft() int32 {
	if o == nil || IsNil(o.TemplatesLeft) {
		var ret int32
		return ret
	}
	return *o.TemplatesLeft
}

// GetTemplatesLeftOk returns a tuple with the TemplatesLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseAccountQuota) GetTemplatesLeftOk() (*int32, bool) {
	if o == nil || IsNil(o.TemplatesLeft) {
		return nil, false
	}
	return o.TemplatesLeft, true
}

// HasTemplatesLeft returns a boolean if a field has been set.
func (o *TemplateResponseAccountQuota) HasTemplatesLeft() bool {
	if o != nil && !IsNil(o.TemplatesLeft) {
		return true
	}

	return false
}

// SetTemplatesLeft gets a reference to the given int32 and assigns it to the TemplatesLeft field.
func (o *TemplateResponseAccountQuota) SetTemplatesLeft(v int32) {
	o.TemplatesLeft = &v
}

// GetApiSignatureRequestsLeft returns the ApiSignatureRequestsLeft field value if set, zero value otherwise.
func (o *TemplateResponseAccountQuota) GetApiSignatureRequestsLeft() int32 {
	if o == nil || IsNil(o.ApiSignatureRequestsLeft) {
		var ret int32
		return ret
	}
	return *o.ApiSignatureRequestsLeft
}

// GetApiSignatureRequestsLeftOk returns a tuple with the ApiSignatureRequestsLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseAccountQuota) GetApiSignatureRequestsLeftOk() (*int32, bool) {
	if o == nil || IsNil(o.ApiSignatureRequestsLeft) {
		return nil, false
	}
	return o.ApiSignatureRequestsLeft, true
}

// HasApiSignatureRequestsLeft returns a boolean if a field has been set.
func (o *TemplateResponseAccountQuota) HasApiSignatureRequestsLeft() bool {
	if o != nil && !IsNil(o.ApiSignatureRequestsLeft) {
		return true
	}

	return false
}

// SetApiSignatureRequestsLeft gets a reference to the given int32 and assigns it to the ApiSignatureRequestsLeft field.
func (o *TemplateResponseAccountQuota) SetApiSignatureRequestsLeft(v int32) {
	o.ApiSignatureRequestsLeft = &v
}

// GetDocumentsLeft returns the DocumentsLeft field value if set, zero value otherwise.
func (o *TemplateResponseAccountQuota) GetDocumentsLeft() int32 {
	if o == nil || IsNil(o.DocumentsLeft) {
		var ret int32
		return ret
	}
	return *o.DocumentsLeft
}

// GetDocumentsLeftOk returns a tuple with the DocumentsLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseAccountQuota) GetDocumentsLeftOk() (*int32, bool) {
	if o == nil || IsNil(o.DocumentsLeft) {
		return nil, false
	}
	return o.DocumentsLeft, true
}

// HasDocumentsLeft returns a boolean if a field has been set.
func (o *TemplateResponseAccountQuota) HasDocumentsLeft() bool {
	if o != nil && !IsNil(o.DocumentsLeft) {
		return true
	}

	return false
}

// SetDocumentsLeft gets a reference to the given int32 and assigns it to the DocumentsLeft field.
func (o *TemplateResponseAccountQuota) SetDocumentsLeft(v int32) {
	o.DocumentsLeft = &v
}

// GetSmsVerificationsLeft returns the SmsVerificationsLeft field value if set, zero value otherwise.
func (o *TemplateResponseAccountQuota) GetSmsVerificationsLeft() int32 {
	if o == nil || IsNil(o.SmsVerificationsLeft) {
		var ret int32
		return ret
	}
	return *o.SmsVerificationsLeft
}

// GetSmsVerificationsLeftOk returns a tuple with the SmsVerificationsLeft field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TemplateResponseAccountQuota) GetSmsVerificationsLeftOk() (*int32, bool) {
	if o == nil || IsNil(o.SmsVerificationsLeft) {
		return nil, false
	}
	return o.SmsVerificationsLeft, true
}

// HasSmsVerificationsLeft returns a boolean if a field has been set.
func (o *TemplateResponseAccountQuota) HasSmsVerificationsLeft() bool {
	if o != nil && !IsNil(o.SmsVerificationsLeft) {
		return true
	}

	return false
}

// SetSmsVerificationsLeft gets a reference to the given int32 and assigns it to the SmsVerificationsLeft field.
func (o *TemplateResponseAccountQuota) SetSmsVerificationsLeft(v int32) {
	o.SmsVerificationsLeft = &v
}

func (o TemplateResponseAccountQuota) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TemplateResponseAccountQuota) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TemplatesLeft) {
		toSerialize["templates_left"] = o.TemplatesLeft
	}
	if !IsNil(o.ApiSignatureRequestsLeft) {
		toSerialize["api_signature_requests_left"] = o.ApiSignatureRequestsLeft
	}
	if !IsNil(o.DocumentsLeft) {
		toSerialize["documents_left"] = o.DocumentsLeft
	}
	if !IsNil(o.SmsVerificationsLeft) {
		toSerialize["sms_verifications_left"] = o.SmsVerificationsLeft
	}
	return toSerialize, nil
}

type NullableTemplateResponseAccountQuota struct {
	value *TemplateResponseAccountQuota
	isSet bool
}

func (v NullableTemplateResponseAccountQuota) Get() *TemplateResponseAccountQuota {
	return v.value
}

func (v *NullableTemplateResponseAccountQuota) Set(val *TemplateResponseAccountQuota) {
	v.value = val
	v.isSet = true
}

func (v NullableTemplateResponseAccountQuota) IsSet() bool {
	return v.isSet
}

func (v *NullableTemplateResponseAccountQuota) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTemplateResponseAccountQuota(val *TemplateResponseAccountQuota) *NullableTemplateResponseAccountQuota {
	return &NullableTemplateResponseAccountQuota{value: val, isSet: true}
}

func (v NullableTemplateResponseAccountQuota) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTemplateResponseAccountQuota) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


