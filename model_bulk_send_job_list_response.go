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

// checks if the BulkSendJobListResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkSendJobListResponse{}

// BulkSendJobListResponse struct for BulkSendJobListResponse
type BulkSendJobListResponse struct {
	// Contains a list of BulkSendJobs that the API caller has access to.
	BulkSendJobs []BulkSendJobResponse `json:"bulk_send_jobs,omitempty"`
	ListInfo *ListInfoResponse `json:"list_info,omitempty"`
	// A list of warnings.
	Warnings []WarningResponse `json:"warnings,omitempty"`
}

// NewBulkSendJobListResponse instantiates a new BulkSendJobListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkSendJobListResponse() *BulkSendJobListResponse {
	this := BulkSendJobListResponse{}
	return &this
}

// NewBulkSendJobListResponseWithDefaults instantiates a new BulkSendJobListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkSendJobListResponseWithDefaults() *BulkSendJobListResponse {
	this := BulkSendJobListResponse{}
	return &this
}

// GetBulkSendJobs returns the BulkSendJobs field value if set, zero value otherwise.
func (o *BulkSendJobListResponse) GetBulkSendJobs() []BulkSendJobResponse {
	if o == nil || IsNil(o.BulkSendJobs) {
		var ret []BulkSendJobResponse
		return ret
	}
	return o.BulkSendJobs
}

// GetBulkSendJobsOk returns a tuple with the BulkSendJobs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendJobListResponse) GetBulkSendJobsOk() ([]BulkSendJobResponse, bool) {
	if o == nil || IsNil(o.BulkSendJobs) {
		return nil, false
	}
	return o.BulkSendJobs, true
}

// HasBulkSendJobs returns a boolean if a field has been set.
func (o *BulkSendJobListResponse) HasBulkSendJobs() bool {
	if o != nil && !IsNil(o.BulkSendJobs) {
		return true
	}

	return false
}

// SetBulkSendJobs gets a reference to the given []BulkSendJobResponse and assigns it to the BulkSendJobs field.
func (o *BulkSendJobListResponse) SetBulkSendJobs(v []BulkSendJobResponse) {
	o.BulkSendJobs = v
}

// GetListInfo returns the ListInfo field value if set, zero value otherwise.
func (o *BulkSendJobListResponse) GetListInfo() ListInfoResponse {
	if o == nil || IsNil(o.ListInfo) {
		var ret ListInfoResponse
		return ret
	}
	return *o.ListInfo
}

// GetListInfoOk returns a tuple with the ListInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendJobListResponse) GetListInfoOk() (*ListInfoResponse, bool) {
	if o == nil || IsNil(o.ListInfo) {
		return nil, false
	}
	return o.ListInfo, true
}

// HasListInfo returns a boolean if a field has been set.
func (o *BulkSendJobListResponse) HasListInfo() bool {
	if o != nil && !IsNil(o.ListInfo) {
		return true
	}

	return false
}

// SetListInfo gets a reference to the given ListInfoResponse and assigns it to the ListInfo field.
func (o *BulkSendJobListResponse) SetListInfo(v ListInfoResponse) {
	o.ListInfo = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *BulkSendJobListResponse) GetWarnings() []WarningResponse {
	if o == nil || IsNil(o.Warnings) {
		var ret []WarningResponse
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BulkSendJobListResponse) GetWarningsOk() ([]WarningResponse, bool) {
	if o == nil || IsNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *BulkSendJobListResponse) HasWarnings() bool {
	if o != nil && !IsNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []WarningResponse and assigns it to the Warnings field.
func (o *BulkSendJobListResponse) SetWarnings(v []WarningResponse) {
	o.Warnings = v
}

func (o BulkSendJobListResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkSendJobListResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BulkSendJobs) {
		toSerialize["bulk_send_jobs"] = o.BulkSendJobs
	}
	if !IsNil(o.ListInfo) {
		toSerialize["list_info"] = o.ListInfo
	}
	if !IsNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	return toSerialize, nil
}

type NullableBulkSendJobListResponse struct {
	value *BulkSendJobListResponse
	isSet bool
}

func (v NullableBulkSendJobListResponse) Get() *BulkSendJobListResponse {
	return v.value
}

func (v *NullableBulkSendJobListResponse) Set(val *BulkSendJobListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkSendJobListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkSendJobListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkSendJobListResponse(val *BulkSendJobListResponse) *NullableBulkSendJobListResponse {
	return &NullableBulkSendJobListResponse{value: val, isSet: true}
}

func (v NullableBulkSendJobListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkSendJobListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

