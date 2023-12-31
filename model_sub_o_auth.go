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

// checks if the SubOAuth type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubOAuth{}

// SubOAuth OAuth related parameters.
type SubOAuth struct {
	// The callback URL to be used for OAuth flows. (Required if `oauth[scopes]` is provided)
	CallbackUrl *string `json:"callback_url,omitempty"`
	// A list of [OAuth scopes](/api/reference/tag/OAuth) to be granted to the app. (Required if `oauth[callback_url]` is provided).
	Scopes []string `json:"scopes,omitempty"`
}

// NewSubOAuth instantiates a new SubOAuth object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubOAuth() *SubOAuth {
	this := SubOAuth{}
	return &this
}

// NewSubOAuthWithDefaults instantiates a new SubOAuth object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubOAuthWithDefaults() *SubOAuth {
	this := SubOAuth{}
	return &this
}

// GetCallbackUrl returns the CallbackUrl field value if set, zero value otherwise.
func (o *SubOAuth) GetCallbackUrl() string {
	if o == nil || IsNil(o.CallbackUrl) {
		var ret string
		return ret
	}
	return *o.CallbackUrl
}

// GetCallbackUrlOk returns a tuple with the CallbackUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubOAuth) GetCallbackUrlOk() (*string, bool) {
	if o == nil || IsNil(o.CallbackUrl) {
		return nil, false
	}
	return o.CallbackUrl, true
}

// HasCallbackUrl returns a boolean if a field has been set.
func (o *SubOAuth) HasCallbackUrl() bool {
	if o != nil && !IsNil(o.CallbackUrl) {
		return true
	}

	return false
}

// SetCallbackUrl gets a reference to the given string and assigns it to the CallbackUrl field.
func (o *SubOAuth) SetCallbackUrl(v string) {
	o.CallbackUrl = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *SubOAuth) GetScopes() []string {
	if o == nil || IsNil(o.Scopes) {
		var ret []string
		return ret
	}
	return o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubOAuth) GetScopesOk() ([]string, bool) {
	if o == nil || IsNil(o.Scopes) {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *SubOAuth) HasScopes() bool {
	if o != nil && !IsNil(o.Scopes) {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []string and assigns it to the Scopes field.
func (o *SubOAuth) SetScopes(v []string) {
	o.Scopes = v
}

func (o SubOAuth) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubOAuth) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallbackUrl) {
		toSerialize["callback_url"] = o.CallbackUrl
	}
	if !IsNil(o.Scopes) {
		toSerialize["scopes"] = o.Scopes
	}
	return toSerialize, nil
}

type NullableSubOAuth struct {
	value *SubOAuth
	isSet bool
}

func (v NullableSubOAuth) Get() *SubOAuth {
	return v.value
}

func (v *NullableSubOAuth) Set(val *SubOAuth) {
	v.value = val
	v.isSet = true
}

func (v NullableSubOAuth) IsSet() bool {
	return v.isSet
}

func (v *NullableSubOAuth) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubOAuth(val *SubOAuth) *NullableSubOAuth {
	return &NullableSubOAuth{value: val, isSet: true}
}

func (v NullableSubOAuth) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubOAuth) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


