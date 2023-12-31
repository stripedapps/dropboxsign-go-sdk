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

// SubFormFieldsPerDocumentTypeEnum the model 'SubFormFieldsPerDocumentTypeEnum'
type SubFormFieldsPerDocumentTypeEnum string

// List of SubFormFieldsPerDocumentTypeEnum
const (
	SUBFORMFIELDSPERDOCUMENTTYPEENUM_CHECKBOX SubFormFieldsPerDocumentTypeEnum = "checkbox"
	SUBFORMFIELDSPERDOCUMENTTYPEENUM_CHECKBOX_MERGE SubFormFieldsPerDocumentTypeEnum = "checkbox-merge"
	SUBFORMFIELDSPERDOCUMENTTYPEENUM_DATE_SIGNED SubFormFieldsPerDocumentTypeEnum = "date_signed"
	SUBFORMFIELDSPERDOCUMENTTYPEENUM_DROPDOWN SubFormFieldsPerDocumentTypeEnum = "dropdown"
	SUBFORMFIELDSPERDOCUMENTTYPEENUM_HYPERLINK SubFormFieldsPerDocumentTypeEnum = "hyperlink"
	SUBFORMFIELDSPERDOCUMENTTYPEENUM_INITIALS SubFormFieldsPerDocumentTypeEnum = "initials"
	SUBFORMFIELDSPERDOCUMENTTYPEENUM_SIGNATURE SubFormFieldsPerDocumentTypeEnum = "signature"
	SUBFORMFIELDSPERDOCUMENTTYPEENUM_RADIO SubFormFieldsPerDocumentTypeEnum = "radio"
	SUBFORMFIELDSPERDOCUMENTTYPEENUM_TEXT SubFormFieldsPerDocumentTypeEnum = "text"
	SUBFORMFIELDSPERDOCUMENTTYPEENUM_TEXT_MERGE SubFormFieldsPerDocumentTypeEnum = "text-merge"
)

// All allowed values of SubFormFieldsPerDocumentTypeEnum enum
var AllowedSubFormFieldsPerDocumentTypeEnumEnumValues = []SubFormFieldsPerDocumentTypeEnum{
	"checkbox",
	"checkbox-merge",
	"date_signed",
	"dropdown",
	"hyperlink",
	"initials",
	"signature",
	"radio",
	"text",
	"text-merge",
}

func (v *SubFormFieldsPerDocumentTypeEnum) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := SubFormFieldsPerDocumentTypeEnum(value)
	for _, existing := range AllowedSubFormFieldsPerDocumentTypeEnumEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid SubFormFieldsPerDocumentTypeEnum", value)
}

// NewSubFormFieldsPerDocumentTypeEnumFromValue returns a pointer to a valid SubFormFieldsPerDocumentTypeEnum
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewSubFormFieldsPerDocumentTypeEnumFromValue(v string) (*SubFormFieldsPerDocumentTypeEnum, error) {
	ev := SubFormFieldsPerDocumentTypeEnum(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for SubFormFieldsPerDocumentTypeEnum: valid values are %v", v, AllowedSubFormFieldsPerDocumentTypeEnumEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v SubFormFieldsPerDocumentTypeEnum) IsValid() bool {
	for _, existing := range AllowedSubFormFieldsPerDocumentTypeEnumEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to SubFormFieldsPerDocumentTypeEnum value
func (v SubFormFieldsPerDocumentTypeEnum) Ptr() *SubFormFieldsPerDocumentTypeEnum {
	return &v
}

type NullableSubFormFieldsPerDocumentTypeEnum struct {
	value *SubFormFieldsPerDocumentTypeEnum
	isSet bool
}

func (v NullableSubFormFieldsPerDocumentTypeEnum) Get() *SubFormFieldsPerDocumentTypeEnum {
	return v.value
}

func (v *NullableSubFormFieldsPerDocumentTypeEnum) Set(val *SubFormFieldsPerDocumentTypeEnum) {
	v.value = val
	v.isSet = true
}

func (v NullableSubFormFieldsPerDocumentTypeEnum) IsSet() bool {
	return v.isSet
}

func (v *NullableSubFormFieldsPerDocumentTypeEnum) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubFormFieldsPerDocumentTypeEnum(val *SubFormFieldsPerDocumentTypeEnum) *NullableSubFormFieldsPerDocumentTypeEnum {
	return &NullableSubFormFieldsPerDocumentTypeEnum{value: val, isSet: true}
}

func (v NullableSubFormFieldsPerDocumentTypeEnum) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubFormFieldsPerDocumentTypeEnum) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

