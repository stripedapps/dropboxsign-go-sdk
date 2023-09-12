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

// checks if the EmbeddedEditUrlRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EmbeddedEditUrlRequest{}

// EmbeddedEditUrlRequest struct for EmbeddedEditUrlRequest
type EmbeddedEditUrlRequest struct {
	// This allows the requester to enable/disable to add or change CC roles when editing the template.
	AllowEditCcs *bool `json:"allow_edit_ccs,omitempty"`
	// The CC roles that must be assigned when using the template to send a signature request. To remove all CC roles, pass in a single role with no name. For use in a POST request.
	CcRoles []string `json:"cc_roles,omitempty"`
	EditorOptions *SubEditorOptions `json:"editor_options,omitempty"`
	// Provide users the ability to review/edit the template signer roles.
	ForceSignerRoles *bool `json:"force_signer_roles,omitempty"`
	// Provide users the ability to review/edit the template subject and message.
	ForceSubjectMessage *bool `json:"force_subject_message,omitempty"`
	// Add additional merge fields to the template, which can be used used to pre-fill data by passing values into signature requests made with that template.  Remove all merge fields on the template by passing an empty array `[]`.
	MergeFields []SubMergeField `json:"merge_fields,omitempty"`
	// This allows the requester to enable the preview experience (i.e. does not allow the requester's end user to add any additional fields via the editor).  **Note**: This parameter overwrites `show_preview=true` (if set).
	PreviewOnly *bool `json:"preview_only,omitempty"`
	// This allows the requester to enable the editor/preview experience.
	ShowPreview *bool `json:"show_preview,omitempty"`
	// When only one step remains in the signature request process and this parameter is set to `false` then the progress stepper will be hidden.
	ShowProgressStepper *bool `json:"show_progress_stepper,omitempty"`
	// Whether this is a test, locked templates will only be available for editing if this is set to `true`. Defaults to `false`.
	TestMode *bool `json:"test_mode,omitempty"`
}

// NewEmbeddedEditUrlRequest instantiates a new EmbeddedEditUrlRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmbeddedEditUrlRequest() *EmbeddedEditUrlRequest {
	this := EmbeddedEditUrlRequest{}
	var allowEditCcs bool = false
	this.AllowEditCcs = &allowEditCcs
	var forceSignerRoles bool = false
	this.ForceSignerRoles = &forceSignerRoles
	var forceSubjectMessage bool = false
	this.ForceSubjectMessage = &forceSubjectMessage
	var previewOnly bool = false
	this.PreviewOnly = &previewOnly
	var showPreview bool = false
	this.ShowPreview = &showPreview
	var showProgressStepper bool = true
	this.ShowProgressStepper = &showProgressStepper
	var testMode bool = false
	this.TestMode = &testMode
	return &this
}

// NewEmbeddedEditUrlRequestWithDefaults instantiates a new EmbeddedEditUrlRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmbeddedEditUrlRequestWithDefaults() *EmbeddedEditUrlRequest {
	this := EmbeddedEditUrlRequest{}
	var allowEditCcs bool = false
	this.AllowEditCcs = &allowEditCcs
	var forceSignerRoles bool = false
	this.ForceSignerRoles = &forceSignerRoles
	var forceSubjectMessage bool = false
	this.ForceSubjectMessage = &forceSubjectMessage
	var previewOnly bool = false
	this.PreviewOnly = &previewOnly
	var showPreview bool = false
	this.ShowPreview = &showPreview
	var showProgressStepper bool = true
	this.ShowProgressStepper = &showProgressStepper
	var testMode bool = false
	this.TestMode = &testMode
	return &this
}

// GetAllowEditCcs returns the AllowEditCcs field value if set, zero value otherwise.
func (o *EmbeddedEditUrlRequest) GetAllowEditCcs() bool {
	if o == nil || IsNil(o.AllowEditCcs) {
		var ret bool
		return ret
	}
	return *o.AllowEditCcs
}

// GetAllowEditCcsOk returns a tuple with the AllowEditCcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedEditUrlRequest) GetAllowEditCcsOk() (*bool, bool) {
	if o == nil || IsNil(o.AllowEditCcs) {
		return nil, false
	}
	return o.AllowEditCcs, true
}

// HasAllowEditCcs returns a boolean if a field has been set.
func (o *EmbeddedEditUrlRequest) HasAllowEditCcs() bool {
	if o != nil && !IsNil(o.AllowEditCcs) {
		return true
	}

	return false
}

// SetAllowEditCcs gets a reference to the given bool and assigns it to the AllowEditCcs field.
func (o *EmbeddedEditUrlRequest) SetAllowEditCcs(v bool) {
	o.AllowEditCcs = &v
}

// GetCcRoles returns the CcRoles field value if set, zero value otherwise.
func (o *EmbeddedEditUrlRequest) GetCcRoles() []string {
	if o == nil || IsNil(o.CcRoles) {
		var ret []string
		return ret
	}
	return o.CcRoles
}

// GetCcRolesOk returns a tuple with the CcRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedEditUrlRequest) GetCcRolesOk() ([]string, bool) {
	if o == nil || IsNil(o.CcRoles) {
		return nil, false
	}
	return o.CcRoles, true
}

// HasCcRoles returns a boolean if a field has been set.
func (o *EmbeddedEditUrlRequest) HasCcRoles() bool {
	if o != nil && !IsNil(o.CcRoles) {
		return true
	}

	return false
}

// SetCcRoles gets a reference to the given []string and assigns it to the CcRoles field.
func (o *EmbeddedEditUrlRequest) SetCcRoles(v []string) {
	o.CcRoles = v
}

// GetEditorOptions returns the EditorOptions field value if set, zero value otherwise.
func (o *EmbeddedEditUrlRequest) GetEditorOptions() SubEditorOptions {
	if o == nil || IsNil(o.EditorOptions) {
		var ret SubEditorOptions
		return ret
	}
	return *o.EditorOptions
}

// GetEditorOptionsOk returns a tuple with the EditorOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedEditUrlRequest) GetEditorOptionsOk() (*SubEditorOptions, bool) {
	if o == nil || IsNil(o.EditorOptions) {
		return nil, false
	}
	return o.EditorOptions, true
}

// HasEditorOptions returns a boolean if a field has been set.
func (o *EmbeddedEditUrlRequest) HasEditorOptions() bool {
	if o != nil && !IsNil(o.EditorOptions) {
		return true
	}

	return false
}

// SetEditorOptions gets a reference to the given SubEditorOptions and assigns it to the EditorOptions field.
func (o *EmbeddedEditUrlRequest) SetEditorOptions(v SubEditorOptions) {
	o.EditorOptions = &v
}

// GetForceSignerRoles returns the ForceSignerRoles field value if set, zero value otherwise.
func (o *EmbeddedEditUrlRequest) GetForceSignerRoles() bool {
	if o == nil || IsNil(o.ForceSignerRoles) {
		var ret bool
		return ret
	}
	return *o.ForceSignerRoles
}

// GetForceSignerRolesOk returns a tuple with the ForceSignerRoles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedEditUrlRequest) GetForceSignerRolesOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceSignerRoles) {
		return nil, false
	}
	return o.ForceSignerRoles, true
}

// HasForceSignerRoles returns a boolean if a field has been set.
func (o *EmbeddedEditUrlRequest) HasForceSignerRoles() bool {
	if o != nil && !IsNil(o.ForceSignerRoles) {
		return true
	}

	return false
}

// SetForceSignerRoles gets a reference to the given bool and assigns it to the ForceSignerRoles field.
func (o *EmbeddedEditUrlRequest) SetForceSignerRoles(v bool) {
	o.ForceSignerRoles = &v
}

// GetForceSubjectMessage returns the ForceSubjectMessage field value if set, zero value otherwise.
func (o *EmbeddedEditUrlRequest) GetForceSubjectMessage() bool {
	if o == nil || IsNil(o.ForceSubjectMessage) {
		var ret bool
		return ret
	}
	return *o.ForceSubjectMessage
}

// GetForceSubjectMessageOk returns a tuple with the ForceSubjectMessage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedEditUrlRequest) GetForceSubjectMessageOk() (*bool, bool) {
	if o == nil || IsNil(o.ForceSubjectMessage) {
		return nil, false
	}
	return o.ForceSubjectMessage, true
}

// HasForceSubjectMessage returns a boolean if a field has been set.
func (o *EmbeddedEditUrlRequest) HasForceSubjectMessage() bool {
	if o != nil && !IsNil(o.ForceSubjectMessage) {
		return true
	}

	return false
}

// SetForceSubjectMessage gets a reference to the given bool and assigns it to the ForceSubjectMessage field.
func (o *EmbeddedEditUrlRequest) SetForceSubjectMessage(v bool) {
	o.ForceSubjectMessage = &v
}

// GetMergeFields returns the MergeFields field value if set, zero value otherwise.
func (o *EmbeddedEditUrlRequest) GetMergeFields() []SubMergeField {
	if o == nil || IsNil(o.MergeFields) {
		var ret []SubMergeField
		return ret
	}
	return o.MergeFields
}

// GetMergeFieldsOk returns a tuple with the MergeFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedEditUrlRequest) GetMergeFieldsOk() ([]SubMergeField, bool) {
	if o == nil || IsNil(o.MergeFields) {
		return nil, false
	}
	return o.MergeFields, true
}

// HasMergeFields returns a boolean if a field has been set.
func (o *EmbeddedEditUrlRequest) HasMergeFields() bool {
	if o != nil && !IsNil(o.MergeFields) {
		return true
	}

	return false
}

// SetMergeFields gets a reference to the given []SubMergeField and assigns it to the MergeFields field.
func (o *EmbeddedEditUrlRequest) SetMergeFields(v []SubMergeField) {
	o.MergeFields = v
}

// GetPreviewOnly returns the PreviewOnly field value if set, zero value otherwise.
func (o *EmbeddedEditUrlRequest) GetPreviewOnly() bool {
	if o == nil || IsNil(o.PreviewOnly) {
		var ret bool
		return ret
	}
	return *o.PreviewOnly
}

// GetPreviewOnlyOk returns a tuple with the PreviewOnly field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedEditUrlRequest) GetPreviewOnlyOk() (*bool, bool) {
	if o == nil || IsNil(o.PreviewOnly) {
		return nil, false
	}
	return o.PreviewOnly, true
}

// HasPreviewOnly returns a boolean if a field has been set.
func (o *EmbeddedEditUrlRequest) HasPreviewOnly() bool {
	if o != nil && !IsNil(o.PreviewOnly) {
		return true
	}

	return false
}

// SetPreviewOnly gets a reference to the given bool and assigns it to the PreviewOnly field.
func (o *EmbeddedEditUrlRequest) SetPreviewOnly(v bool) {
	o.PreviewOnly = &v
}

// GetShowPreview returns the ShowPreview field value if set, zero value otherwise.
func (o *EmbeddedEditUrlRequest) GetShowPreview() bool {
	if o == nil || IsNil(o.ShowPreview) {
		var ret bool
		return ret
	}
	return *o.ShowPreview
}

// GetShowPreviewOk returns a tuple with the ShowPreview field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedEditUrlRequest) GetShowPreviewOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowPreview) {
		return nil, false
	}
	return o.ShowPreview, true
}

// HasShowPreview returns a boolean if a field has been set.
func (o *EmbeddedEditUrlRequest) HasShowPreview() bool {
	if o != nil && !IsNil(o.ShowPreview) {
		return true
	}

	return false
}

// SetShowPreview gets a reference to the given bool and assigns it to the ShowPreview field.
func (o *EmbeddedEditUrlRequest) SetShowPreview(v bool) {
	o.ShowPreview = &v
}

// GetShowProgressStepper returns the ShowProgressStepper field value if set, zero value otherwise.
func (o *EmbeddedEditUrlRequest) GetShowProgressStepper() bool {
	if o == nil || IsNil(o.ShowProgressStepper) {
		var ret bool
		return ret
	}
	return *o.ShowProgressStepper
}

// GetShowProgressStepperOk returns a tuple with the ShowProgressStepper field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedEditUrlRequest) GetShowProgressStepperOk() (*bool, bool) {
	if o == nil || IsNil(o.ShowProgressStepper) {
		return nil, false
	}
	return o.ShowProgressStepper, true
}

// HasShowProgressStepper returns a boolean if a field has been set.
func (o *EmbeddedEditUrlRequest) HasShowProgressStepper() bool {
	if o != nil && !IsNil(o.ShowProgressStepper) {
		return true
	}

	return false
}

// SetShowProgressStepper gets a reference to the given bool and assigns it to the ShowProgressStepper field.
func (o *EmbeddedEditUrlRequest) SetShowProgressStepper(v bool) {
	o.ShowProgressStepper = &v
}

// GetTestMode returns the TestMode field value if set, zero value otherwise.
func (o *EmbeddedEditUrlRequest) GetTestMode() bool {
	if o == nil || IsNil(o.TestMode) {
		var ret bool
		return ret
	}
	return *o.TestMode
}

// GetTestModeOk returns a tuple with the TestMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmbeddedEditUrlRequest) GetTestModeOk() (*bool, bool) {
	if o == nil || IsNil(o.TestMode) {
		return nil, false
	}
	return o.TestMode, true
}

// HasTestMode returns a boolean if a field has been set.
func (o *EmbeddedEditUrlRequest) HasTestMode() bool {
	if o != nil && !IsNil(o.TestMode) {
		return true
	}

	return false
}

// SetTestMode gets a reference to the given bool and assigns it to the TestMode field.
func (o *EmbeddedEditUrlRequest) SetTestMode(v bool) {
	o.TestMode = &v
}

func (o EmbeddedEditUrlRequest) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmbeddedEditUrlRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AllowEditCcs) {
		toSerialize["allow_edit_ccs"] = o.AllowEditCcs
	}
	if !IsNil(o.CcRoles) {
		toSerialize["cc_roles"] = o.CcRoles
	}
	if !IsNil(o.EditorOptions) {
		toSerialize["editor_options"] = o.EditorOptions
	}
	if !IsNil(o.ForceSignerRoles) {
		toSerialize["force_signer_roles"] = o.ForceSignerRoles
	}
	if !IsNil(o.ForceSubjectMessage) {
		toSerialize["force_subject_message"] = o.ForceSubjectMessage
	}
	if !IsNil(o.MergeFields) {
		toSerialize["merge_fields"] = o.MergeFields
	}
	if !IsNil(o.PreviewOnly) {
		toSerialize["preview_only"] = o.PreviewOnly
	}
	if !IsNil(o.ShowPreview) {
		toSerialize["show_preview"] = o.ShowPreview
	}
	if !IsNil(o.ShowProgressStepper) {
		toSerialize["show_progress_stepper"] = o.ShowProgressStepper
	}
	if !IsNil(o.TestMode) {
		toSerialize["test_mode"] = o.TestMode
	}
	return toSerialize, nil
}

type NullableEmbeddedEditUrlRequest struct {
	value *EmbeddedEditUrlRequest
	isSet bool
}

func (v NullableEmbeddedEditUrlRequest) Get() *EmbeddedEditUrlRequest {
	return v.value
}

func (v *NullableEmbeddedEditUrlRequest) Set(val *EmbeddedEditUrlRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableEmbeddedEditUrlRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableEmbeddedEditUrlRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmbeddedEditUrlRequest(val *EmbeddedEditUrlRequest) *NullableEmbeddedEditUrlRequest {
	return &NullableEmbeddedEditUrlRequest{value: val, isSet: true}
}

func (v NullableEmbeddedEditUrlRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmbeddedEditUrlRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


