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

// checks if the FileResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FileResponse{}

// FileResponse struct for FileResponse
type FileResponse struct {
	// URL to the file.
	FileUrl *string `json:"file_url,omitempty"`
	// When the link expires.
	ExpiresAt *int32 `json:"expires_at,omitempty"`
}

// NewFileResponse instantiates a new FileResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileResponse() *FileResponse {
	this := FileResponse{}
	return &this
}

// NewFileResponseWithDefaults instantiates a new FileResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileResponseWithDefaults() *FileResponse {
	this := FileResponse{}
	return &this
}

// GetFileUrl returns the FileUrl field value if set, zero value otherwise.
func (o *FileResponse) GetFileUrl() string {
	if o == nil || IsNil(o.FileUrl) {
		var ret string
		return ret
	}
	return *o.FileUrl
}

// GetFileUrlOk returns a tuple with the FileUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileResponse) GetFileUrlOk() (*string, bool) {
	if o == nil || IsNil(o.FileUrl) {
		return nil, false
	}
	return o.FileUrl, true
}

// HasFileUrl returns a boolean if a field has been set.
func (o *FileResponse) HasFileUrl() bool {
	if o != nil && !IsNil(o.FileUrl) {
		return true
	}

	return false
}

// SetFileUrl gets a reference to the given string and assigns it to the FileUrl field.
func (o *FileResponse) SetFileUrl(v string) {
	o.FileUrl = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *FileResponse) GetExpiresAt() int32 {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret int32
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FileResponse) GetExpiresAtOk() (*int32, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *FileResponse) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given int32 and assigns it to the ExpiresAt field.
func (o *FileResponse) SetExpiresAt(v int32) {
	o.ExpiresAt = &v
}

func (o FileResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FileResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.FileUrl) {
		toSerialize["file_url"] = o.FileUrl
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expires_at"] = o.ExpiresAt
	}
	return toSerialize, nil
}

type NullableFileResponse struct {
	value *FileResponse
	isSet bool
}

func (v NullableFileResponse) Get() *FileResponse {
	return v.value
}

func (v *NullableFileResponse) Set(val *FileResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableFileResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableFileResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileResponse(val *FileResponse) *NullableFileResponse {
	return &NullableFileResponse{value: val, isSet: true}
}

func (v NullableFileResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


