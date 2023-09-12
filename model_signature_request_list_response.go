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

// checks if the SignatureRequestListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SignatureRequestListResponse{}

// SignatureRequestListResponse struct for SignatureRequestListResponse
type SignatureRequestListResponse struct {
	// Contains information about signature requests.
	SignatureRequests []SignatureRequestResponse `json:"signature_requests,omitempty"`
	ListInfo *ListInfoResponse `json:"list_info,omitempty"`
	// A list of warnings.
	Warnings []WarningResponse `json:"warnings,omitempty"`
}

// NewSignatureRequestListResponse instantiates a new SignatureRequestListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSignatureRequestListResponse() *SignatureRequestListResponse {
	this := SignatureRequestListResponse{}
	return &this
}

// NewSignatureRequestListResponseWithDefaults instantiates a new SignatureRequestListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSignatureRequestListResponseWithDefaults() *SignatureRequestListResponse {
	this := SignatureRequestListResponse{}
	return &this
}

// GetSignatureRequests returns the SignatureRequests field value if set, zero value otherwise.
func (o *SignatureRequestListResponse) GetSignatureRequests() []SignatureRequestResponse {
	if o == nil || IsNil(o.SignatureRequests) {
		var ret []SignatureRequestResponse
		return ret
	}
	return o.SignatureRequests
}

// GetSignatureRequestsOk returns a tuple with the SignatureRequests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestListResponse) GetSignatureRequestsOk() ([]SignatureRequestResponse, bool) {
	if o == nil || IsNil(o.SignatureRequests) {
		return nil, false
	}
	return o.SignatureRequests, true
}

// HasSignatureRequests returns a boolean if a field has been set.
func (o *SignatureRequestListResponse) HasSignatureRequests() bool {
	if o != nil && !IsNil(o.SignatureRequests) {
		return true
	}

	return false
}

// SetSignatureRequests gets a reference to the given []SignatureRequestResponse and assigns it to the SignatureRequests field.
func (o *SignatureRequestListResponse) SetSignatureRequests(v []SignatureRequestResponse) {
	o.SignatureRequests = v
}

// GetListInfo returns the ListInfo field value if set, zero value otherwise.
func (o *SignatureRequestListResponse) GetListInfo() ListInfoResponse {
	if o == nil || IsNil(o.ListInfo) {
		var ret ListInfoResponse
		return ret
	}
	return *o.ListInfo
}

// GetListInfoOk returns a tuple with the ListInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestListResponse) GetListInfoOk() (*ListInfoResponse, bool) {
	if o == nil || IsNil(o.ListInfo) {
		return nil, false
	}
	return o.ListInfo, true
}

// HasListInfo returns a boolean if a field has been set.
func (o *SignatureRequestListResponse) HasListInfo() bool {
	if o != nil && !IsNil(o.ListInfo) {
		return true
	}

	return false
}

// SetListInfo gets a reference to the given ListInfoResponse and assigns it to the ListInfo field.
func (o *SignatureRequestListResponse) SetListInfo(v ListInfoResponse) {
	o.ListInfo = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *SignatureRequestListResponse) GetWarnings() []WarningResponse {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningResponse
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SignatureRequestListResponse) GetWarningsOk() ([]WarningResponse, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *SignatureRequestListResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningResponse and assigns it to the Warnings field.
func (o *SignatureRequestListResponse) SetWarnings(v []WarningResponse) {
	o.Warnings = v
}

func (o SignatureRequestListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SignatureRequestListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SignatureRequests) {
		toSerialize["signature_requests"] = o.SignatureRequests
	}
	if !IsNil(o.ListInfo) {
		toSerialize["list_info"] = o.ListInfo
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableSignatureRequestListResponse struct {
	value *SignatureRequestListResponse
	isSet bool
}

func (v NullableSignatureRequestListResponse) Get() *SignatureRequestListResponse {
	return v.value
}

func (v *NullableSignatureRequestListResponse) Set(val *SignatureRequestListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureRequestListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureRequestListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureRequestListResponse(val *SignatureRequestListResponse) *NullableSignatureRequestListResponse {
	return &NullableSignatureRequestListResponse{value: val, isSet: true}
}

func (v NullableSignatureRequestListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureRequestListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


