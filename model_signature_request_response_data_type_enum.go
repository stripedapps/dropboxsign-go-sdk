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

// SignatureRequestResponseDataTypeEnum the model 'SignatureRequestResponseDataTypeEnum'
type SignatureRequestResponseDataTypeEnum string

// List of SignatureRequestResponseDataTypeEnum
const (
	TEXT SignatureRequestResponseDataTypeEnum = "text"
	CHECKBOX SignatureRequestResponseDataTypeEnum = "checkbox"
	DATE_SIGNED SignatureRequestResponseDataTypeEnum = "date_signed"
	DROPDOWN SignatureRequestResponseDataTypeEnum = "dropdown"
	INITIALS SignatureRequestResponseDataTypeEnum = "initials"
	RADIO SignatureRequestResponseDataTypeEnum = "radio"
	SIGNATURE SignatureRequestResponseDataTypeEnum = "signature"
	TEXT_MERGE SignatureRequestResponseDataTypeEnum = "text-merge"
	CHECKBOX_MERGE SignatureRequestResponseDataTypeEnum = "checkbox-merge"
)

// All allowed values of SignatureRequestResponseDataTypeEnum enum
var AllowedSignatureRequestResponseDataTypeEnumEnumValues = []SignatureRequestResponseDataTypeEnum{
	"text",
	"checkbox",
	"date_signed",
	"dropdown",
	"initials",
	"radio",
	"signature",
	"text-merge",
	"checkbox-merge",
}

func (v *SignatureRequestResponseDataTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SignatureRequestResponseDataTypeEnum(value)
	for _, existing := range AllowedSignatureRequestResponseDataTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SignatureRequestResponseDataTypeEnum", value)
}

// NewSignatureRequestResponseDataTypeEnumFromValue returns a pointer to a valid SignatureRequestResponseDataTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSignatureRequestResponseDataTypeEnumFromValue(v string) (*SignatureRequestResponseDataTypeEnum, error) {
	ev := SignatureRequestResponseDataTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SignatureRequestResponseDataTypeEnum: valid values are %v", v, AllowedSignatureRequestResponseDataTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SignatureRequestResponseDataTypeEnum) IsValid() bool {
	for _, existing := range AllowedSignatureRequestResponseDataTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SignatureRequestResponseDataTypeEnum value
func (v SignatureRequestResponseDataTypeEnum) Ptr() *SignatureRequestResponseDataTypeEnum {
	return &v
}

type NullableSignatureRequestResponseDataTypeEnum struct {
	value *SignatureRequestResponseDataTypeEnum
	isSet bool
}

func (v NullableSignatureRequestResponseDataTypeEnum) Get() *SignatureRequestResponseDataTypeEnum {
	return v.value
}

func (v *NullableSignatureRequestResponseDataTypeEnum) Set(val *SignatureRequestResponseDataTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableSignatureRequestResponseDataTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableSignatureRequestResponseDataTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSignatureRequestResponseDataTypeEnum(val *SignatureRequestResponseDataTypeEnum) *NullableSignatureRequestResponseDataTypeEnum {
	return &NullableSignatureRequestResponseDataTypeEnum{value: val, isSet: true}
}

func (v NullableSignatureRequestResponseDataTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSignatureRequestResponseDataTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
