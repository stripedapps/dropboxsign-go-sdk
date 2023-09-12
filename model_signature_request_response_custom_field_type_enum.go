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
	"fmt"
)

// SignatureRequestResponseCustomFieldTypeEnum the model 'SignatureRequestResponseCustomFieldTypeEnum'
type SignatureRequestResponseCustomFieldTypeEnum string

// List of SignatureRequestResponseCustomFieldTypeEnum
const (
	TEXT SignatureRequestResponseCustomFieldTypeEnum = "text"
	CHECKBOX SignatureRequestResponseCustomFieldTypeEnum = "checkbox"
)

// All allowed values of SignatureRequestResponseCustomFieldTypeEnum enum
var AllowedSignatureRequestResponseCustomFieldTypeEnumEnumValues = []SignatureRequestResponseCustomFieldTypeEnum{
	"text",
	"checkbox",
}

func (v *SignatureRequestResponseCustomFieldTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SignatureRequestResponseCustomFieldTypeEnum(value)
	for _, existing := range AllowedSignatureRequestResponseCustomFieldTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SignatureRequestResponseCustomFieldTypeEnum", value)
}

// NewSignatureRequestResponseCustomFieldTypeEnumFromValue returns a pointer to a valid SignatureRequestResponseCustomFieldTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSignatureRequestResponseCustomFieldTypeEnumFromValue(v string) (*SignatureRequestResponseCustomFieldTypeEnum, error) {
	ev := SignatureRequestResponseCustomFieldTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SignatureRequestResponseCustomFieldTypeEnum: valid values are %v", v, AllowedSignatureRequestResponseCustomFieldTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SignatureRequestResponseCustomFieldTypeEnum) IsValid() bool {
	for _, existing := range AllowedSignatureRequestResponseCustomFieldTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SignatureRequestResponseCustomFieldTypeEnum value
func (v SignatureRequestResponseCustomFieldTypeEnum) Ptr() *SignatureRequestResponseCustomFieldTypeEnum {
	return &v
}

type NullableSignatureRequestResponseCustomFieldTypeEnum struct {
	value *SignatureRequestResponseCustomFieldTypeEnum
	isSet bool
}

func (v NullableSignatureRequestResponseCustomFieldTypeEnum) Get() *SignatureRequestResponseCustomFieldTypeEnum {
	return v.value
}

func (v *NullableSignatureRequestResponseCustomFieldTypeEnum) Set(val *SignatureRequestResponseCustomFieldTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureRequestResponseCustomFieldTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureRequestResponseCustomFieldTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureRequestResponseCustomFieldTypeEnum(val *SignatureRequestResponseCustomFieldTypeEnum) *NullableSignatureRequestResponseCustomFieldTypeEnum {
	return &NullableSignatureRequestResponseCustomFieldTypeEnum{value: val, isSet: true}
}

func (v NullableSignatureRequestResponseCustomFieldTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureRequestResponseCustomFieldTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

