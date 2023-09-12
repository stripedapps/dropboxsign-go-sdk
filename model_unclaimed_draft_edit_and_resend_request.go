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

// checks if the UnclaimedDraftEditAndResendRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UnclaimedDraftEditAndResendRequest{}

// UnclaimedDraftEditAndResendRequest struct for UnclaimedDraftEditAndResendRequest
type UnclaimedDraftEditAndResendRequest struct {
	// Client id of the app used to create the draft. Used to apply the branding and callback url defined for the app.
	ClientId string `json:"client_id"`
	EditorOptions *SubEditorOptions `json:"editor_options,omitempty"`
	// The request created from this draft will also be signable in embedded mode if set to `true`.
	IsForEmbeddedSigning *bool `json:"is_for_embedded_signing,omitempty"`
	// The email address of the user that should be designated as the requester of this draft. If not set, original requester's email address will be used.
	RequesterEmailAddress *string `json:"requester_email_address,omitempty"`
	// The URL you want signers redirected to after they successfully request a signature.
	RequestingRedirectUrl *string `json:"requesting_redirect_url,omitempty"`
	// When only one step remains in the signature request process and this parameter is set to `false` then the progress stepper will be hidden.
	ShowProgressStepper *bool `json:"show_progress_stepper,omitempty"`
	// The URL you want signers redirected to after they successfully sign.
	SigningRedirectUrl *string `json:"signing_redirect_url,omitempty"`
	// Whether this is a test, the signature request created from this draft will not be legally binding if set to `true`. Defaults to `false`.
	TestMode *bool `json:"test_mode,omitempty"`
}

// NewUnclaimedDraftEditAndResendRequest instantiates a new UnclaimedDraftEditAndResendRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUnclaimedDraftEditAndResendRequest(clientId string) *UnclaimedDraftEditAndResendRequest {
	this := UnclaimedDraftEditAndResendRequest{}
	this.ClientId = clientId
	var showProgressStepper bool = true
	this.ShowProgressStepper = &showProgressStepper
	var testMode bool = false
	this.TestMode = &testMode
	return &this
}

// NewUnclaimedDraftEditAndResendRequestWithDefaults instantiates a new UnclaimedDraftEditAndResendRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUnclaimedDraftEditAndResendRequestWithDefaults() *UnclaimedDraftEditAndResendRequest {
	this := UnclaimedDraftEditAndResendRequest{}
	var showProgressStepper bool = true
	this.ShowProgressStepper = &showProgressStepper
	var testMode bool = false
	this.TestMode = &testMode
	return &this
}

// GetClientId returns the ClientId field value
func (o *UnclaimedDraftEditAndResendRequest) GetClientId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ClientId
}

// GetClientIdOk returns a tuple with the ClientId field value
// and a boolean to check if the value has been set.
func (o *UnclaimedDraftEditAndResendRequest) GetClientIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ClientId, true
}

// SetClientId sets field value
func (o *UnclaimedDraftEditAndResendRequest) SetClientId(v string) {
	o.ClientId = v
}

// GetEditorOptions returns the EditorOptions field value if set, zero value otherwise.
func (o *UnclaimedDraftEditAndResendRequest) GetEditorOptions() SubEditorOptions {
	if o == nil || IsNil(o.EditorOptions) {
		var ret SubEditorOptions
		return ret
	}
	return *o.EditorOptions
}

// GetEditorOptionsOk returns a tuple with the EditorOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnclaimedDraftEditAndResendRequest) GetEditorOptionsOk() (*SubEditorOptions, bool) {
	if o == nil || IsNil(o.EditorOptions) {
		return nil, false
	}
	return o.EditorOptions, true
}

// HasEditorOptions returns a boolean if a field has been set.
func (o *UnclaimedDraftEditAndResendRequest) HasEditorOptions() bool {
	if o != nil && !IsNil(o.EditorOptions) {
		return true
	}

	return false
}

// SetEditorOptions gets a reference to the given SubEditorOptions and assigns it to the EditorOptions field.
func (o *UnclaimedDraftEditAndResendRequest) SetEditorOptions(v SubEditorOptions) {
	o.EditorOptions = &v
}

// GetIsForEmbeddedSigning returns the IsForEmbeddedSigning field value if set, zero value otherwise.
func (o *UnclaimedDraftEditAndResendRequest) GetIsForEmbeddedSigning() bool {
	if o == nil || IsNil(o.IsForEmbeddedSigning) {
		var ret bool
		return ret
	}
	return *o.IsForEmbeddedSigning
}

// GetIsForEmbeddedSigningOk returns a tuple with the IsForEmbeddedSigning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnclaimedDraftEditAndResendRequest) GetIsForEmbeddedSigningOk() (*bool, bool) {
	if o == nil || IsNil(o.IsForEmbeddedSigning) {
		return nil, false
	}
	return o.IsForEmbeddedSigning, true
}

// HasIsForEmbeddedSigning returns a boolean if a field has been set.
func (o *UnclaimedDraftEditAndResendRequest) HasIsForEmbeddedSigning() bool {
	if o != nil && !IsNil(o.IsForEmbeddedSigning) {
		return true
	}

	return false
}

// SetIsForEmbeddedSigning gets a reference to the given bool and assigns it to the IsForEmbeddedSigning field.
func (o *UnclaimedDraftEditAndResendRequest) SetIsForEmbeddedSigning(v bool) {
	o.IsForEmbeddedSigning = &v
}

// GetRequesterEmailAddress returns the RequesterEmailAddress field value if set, zero value otherwise.
func (o *UnclaimedDraftEditAndResendRequest) GetRequesterEmailAddress() string {
	if o == nil || IsNil(o.RequesterEmailAddress) {
		var ret string
		return ret
	}
	return *o.RequesterEmailAddress
}

// GetRequesterEmailAddressOk returns a tuple with the RequesterEmailAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnclaimedDraftEditAndResendRequest) GetRequesterEmailAddressOk() (*string, bool) {
	if o == nil || IsNil(o.RequesterEmailAddress) {
		return nil, false
	}
	return o.RequesterEmailAddress, true
}

// HasRequesterEmailAddress returns a boolean if a field has been set.
func (o *UnclaimedDraftEditAndResendRequest) HasRequesterEmailAddress() bool {
	if o != nil && !IsNil(o.RequesterEmailAddress) {
		return true
	}

	return false
}

// SetRequesterEmailAddress gets a reference to the given string and assigns it to the RequesterEmailAddress field.
func (o *UnclaimedDraftEditAndResendRequest) SetRequesterEmailAddress(v string) {
	o.RequesterEmailAddress = &v
}

// GetRequestingRedirectUrl returns the RequestingRedirectUrl field value if set, zero value otherwise.
func (o *UnclaimedDraftEditAndResendRequest) GetRequestingRedirectUrl() string {
	if o == nil || IsNil(o.RequestingRedirectUrl) {
		var ret string
		return ret
	}
	return *o.RequestingRedirectUrl
}

// GetRequestingRedirectUrlOk returns a tuple with the RequestingRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnclaimedDraftEditAndResendRequest) GetRequestingRedirectUrlOk() (*string, bool) {
	if o == nil || IsNil(o.RequestingRedirectUrl) {
		return nil, false
	}
	return o.RequestingRedirectUrl, true
}

// HasRequestingRedirectUrl returns a boolean if a field has been set.
func (o *UnclaimedDraftEditAndResendRequest) HasRequestingRedirectUrl() bool {
	if o != nil && !IsNil(o.RequestingRedirectUrl) {
		return true
	}

	return false
}

// SetRequestingRedirectUrl gets a reference to the given string and assigns it to the RequestingRedirectUrl field.
func (o *UnclaimedDraftEditAndResendRequest) SetRequestingRedirectUrl(v string) {
	o.RequestingRedirectUrl = &v
}

// GetShowProgressStepper returns the ShowProgressStepper field value if set, zero value otherwise.
func (o *UnclaimedDraftEditAndResendRequest) GetShowProgressStepper() bool {
	if o == nil || IsNil(o.ShowProgressStepper) {
		var ret bool
		return ret
	}
	return *o.ShowProgressStepper
}

// GetShowProgressStepperOk returns a tuple with the ShowProgressStepper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnclaimedDraftEditAndResendRequest) GetShowProgressStepperOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowProgressStepper) {
		return nil, false
	}
	return o.ShowProgressStepper, true
}

// HasShowProgressStepper returns a boolean if a field has been set.
func (o *UnclaimedDraftEditAndResendRequest) HasShowProgressStepper() bool {
	if o != nil && !IsNil(o.ShowProgressStepper) {
		return true
	}

	return false
}

// SetShowProgressStepper gets a reference to the given bool and assigns it to the ShowProgressStepper field.
func (o *UnclaimedDraftEditAndResendRequest) SetShowProgressStepper(v bool) {
	o.ShowProgressStepper = &v
}

// GetSigningRedirectUrl returns the SigningRedirectUrl field value if set, zero value otherwise.
func (o *UnclaimedDraftEditAndResendRequest) GetSigningRedirectUrl() string {
	if o == nil || IsNil(o.SigningRedirectUrl) {
		var ret string
		return ret
	}
	return *o.SigningRedirectUrl
}

// GetSigningRedirectUrlOk returns a tuple with the SigningRedirectUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnclaimedDraftEditAndResendRequest) GetSigningRedirectUrlOk() (*string, bool) {
	if o == nil || IsNil(o.SigningRedirectUrl) {
		return nil, false
	}
	return o.SigningRedirectUrl, true
}

// HasSigningRedirectUrl returns a boolean if a field has been set.
func (o *UnclaimedDraftEditAndResendRequest) HasSigningRedirectUrl() bool {
	if o != nil && !IsNil(o.SigningRedirectUrl) {
		return true
	}

	return false
}

// SetSigningRedirectUrl gets a reference to the given string and assigns it to the SigningRedirectUrl field.
func (o *UnclaimedDraftEditAndResendRequest) SetSigningRedirectUrl(v string) {
	o.SigningRedirectUrl = &v
}

// GetTestMode returns the TestMode field value if set, zero value otherwise.
func (o *UnclaimedDraftEditAndResendRequest) GetTestMode() bool {
	if o == nil || IsNil(o.TestMode) {
		var ret bool
		return ret
	}
	return *o.TestMode
}

// GetTestModeOk returns a tuple with the TestMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UnclaimedDraftEditAndResendRequest) GetTestModeOk() (*bool, bool) {
	if o == nil || IsNil(o.TestMode) {
		return nil, false
	}
	return o.TestMode, true
}

// HasTestMode returns a boolean if a field has been set.
func (o *UnclaimedDraftEditAndResendRequest) HasTestMode() bool {
	if o != nil && !IsNil(o.TestMode) {
		return true
	}

	return false
}

// SetTestMode gets a reference to the given bool and assigns it to the TestMode field.
func (o *UnclaimedDraftEditAndResendRequest) SetTestMode(v bool) {
	o.TestMode = &v
}

func (o UnclaimedDraftEditAndResendRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UnclaimedDraftEditAndResendRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["client_id"] = o.ClientId
	if !IsNil(o.EditorOptions) {
		toSerialize["editor_options"] = o.EditorOptions
	}
	if !IsNil(o.IsForEmbeddedSigning) {
		toSerialize["is_for_embedded_signing"] = o.IsForEmbeddedSigning
	}
	if !IsNil(o.RequesterEmailAddress) {
		toSerialize["requester_email_address"] = o.RequesterEmailAddress
	}
	if !IsNil(o.RequestingRedirectUrl) {
		toSerialize["requesting_redirect_url"] = o.RequestingRedirectUrl
	}
	if !IsNil(o.ShowProgressStepper) {
		toSerialize["show_progress_stepper"] = o.ShowProgressStepper
	}
	if !IsNil(o.SigningRedirectUrl) {
		toSerialize["signing_redirect_url"] = o.SigningRedirectUrl
	}
	if !IsNil(o.TestMode) {
		toSerialize["test_mode"] = o.TestMode
	}
	return toSerialize, nil
}

type NullableUnclaimedDraftEditAndResendRequest struct {
	value *UnclaimedDraftEditAndResendRequest
	isSet bool
}

func (v NullableUnclaimedDraftEditAndResendRequest) Get() *UnclaimedDraftEditAndResendRequest {
	return v.value
}

func (v *NullableUnclaimedDraftEditAndResendRequest) Set(val *UnclaimedDraftEditAndResendRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUnclaimedDraftEditAndResendRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUnclaimedDraftEditAndResendRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUnclaimedDraftEditAndResendRequest(val *UnclaimedDraftEditAndResendRequest) *NullableUnclaimedDraftEditAndResendRequest {
	return &NullableUnclaimedDraftEditAndResendRequest{value: val, isSet: true}
}

func (v NullableUnclaimedDraftEditAndResendRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUnclaimedDraftEditAndResendRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

