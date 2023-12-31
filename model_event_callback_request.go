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

// checks if the EventCallbackRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventCallbackRequest{}

// EventCallbackRequest struct for EventCallbackRequest
type EventCallbackRequest struct {
	Event EventCallbackRequestEvent `json:"event"`
	Account *AccountResponse `json:"account,omitempty"`
	SignatureRequest *SignatureRequestResponse `json:"signature_request,omitempty"`
	Template *TemplateResponse `json:"template,omitempty"`
}

// NewEventCallbackRequest instantiates a new EventCallbackRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventCallbackRequest(event EventCallbackRequestEvent) *EventCallbackRequest {
	this := EventCallbackRequest{}
	this.Event = event
	return &this
}

// NewEventCallbackRequestWithDefaults instantiates a new EventCallbackRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventCallbackRequestWithDefaults() *EventCallbackRequest {
	this := EventCallbackRequest{}
	return &this
}

// GetEvent returns the Event field value
func (o *EventCallbackRequest) GetEvent() EventCallbackRequestEvent {
	if o == nil {
		var ret EventCallbackRequestEvent
		return ret
	}

	return o.Event
}

// GetEventOk returns a tuple with the Event field value
// and a boolean to check if the value has been set.
func (o *EventCallbackRequest) GetEventOk() (*EventCallbackRequestEvent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Event, true
}

// SetEvent sets field value
func (o *EventCallbackRequest) SetEvent(v EventCallbackRequestEvent) {
	o.Event = v
}

// GetAccount returns the Account field value if set, zero value otherwise.
func (o *EventCallbackRequest) GetAccount() AccountResponse {
	if o == nil || IsNil(o.Account) {
		var ret AccountResponse
		return ret
	}
	return *o.Account
}

// GetAccountOk returns a tuple with the Account field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCallbackRequest) GetAccountOk() (*AccountResponse, bool) {
	if o == nil || IsNil(o.Account) {
		return nil, false
	}
	return o.Account, true
}

// HasAccount returns a boolean if a field has been set.
func (o *EventCallbackRequest) HasAccount() bool {
	if o != nil && !IsNil(o.Account) {
		return true
	}

	return false
}

// SetAccount gets a reference to the given AccountResponse and assigns it to the Account field.
func (o *EventCallbackRequest) SetAccount(v AccountResponse) {
	o.Account = &v
}

// GetSignatureRequest returns the SignatureRequest field value if set, zero value otherwise.
func (o *EventCallbackRequest) GetSignatureRequest() SignatureRequestResponse {
	if o == nil || IsNil(o.SignatureRequest) {
		var ret SignatureRequestResponse
		return ret
	}
	return *o.SignatureRequest
}

// GetSignatureRequestOk returns a tuple with the SignatureRequest field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCallbackRequest) GetSignatureRequestOk() (*SignatureRequestResponse, bool) {
	if o == nil || IsNil(o.SignatureRequest) {
		return nil, false
	}
	return o.SignatureRequest, true
}

// HasSignatureRequest returns a boolean if a field has been set.
func (o *EventCallbackRequest) HasSignatureRequest() bool {
	if o != nil && !IsNil(o.SignatureRequest) {
		return true
	}

	return false
}

// SetSignatureRequest gets a reference to the given SignatureRequestResponse and assigns it to the SignatureRequest field.
func (o *EventCallbackRequest) SetSignatureRequest(v SignatureRequestResponse) {
	o.SignatureRequest = &v
}

// GetTemplate returns the Template field value if set, zero value otherwise.
func (o *EventCallbackRequest) GetTemplate() TemplateResponse {
	if o == nil || IsNil(o.Template) {
		var ret TemplateResponse
		return ret
	}
	return *o.Template
}

// GetTemplateOk returns a tuple with the Template field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EventCallbackRequest) GetTemplateOk() (*TemplateResponse, bool) {
	if o == nil || IsNil(o.Template) {
		return nil, false
	}
	return o.Template, true
}

// HasTemplate returns a boolean if a field has been set.
func (o *EventCallbackRequest) HasTemplate() bool {
	if o != nil && !IsNil(o.Template) {
		return true
	}

	return false
}

// SetTemplate gets a reference to the given TemplateResponse and assigns it to the Template field.
func (o *EventCallbackRequest) SetTemplate(v TemplateResponse) {
	o.Template = &v
}

func (o EventCallbackRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventCallbackRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["event"] = o.Event
	if !IsNil(o.Account) {
		toSerialize["account"] = o.Account
	}
	if !IsNil(o.SignatureRequest) {
		toSerialize["signature_request"] = o.SignatureRequest
	}
	if !IsNil(o.Template) {
		toSerialize["template"] = o.Template
	}
	return toSerialize, nil
}

type NullableEventCallbackRequest struct {
	value *EventCallbackRequest
	isSet bool
}

func (v NullableEventCallbackRequest) Get() *EventCallbackRequest {
	return v.value
}

func (v *NullableEventCallbackRequest) Set(val *EventCallbackRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEventCallbackRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEventCallbackRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventCallbackRequest(val *EventCallbackRequest) *NullableEventCallbackRequest {
	return &NullableEventCallbackRequest{value: val, isSet: true}
}

func (v NullableEventCallbackRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventCallbackRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


