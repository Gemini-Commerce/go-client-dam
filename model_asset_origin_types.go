/*
Dam Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// AssetOriginTypes the model 'AssetOriginTypes'
type AssetOriginTypes string

// List of AssetOriginTypes
const (
	ASSETORIGINTYPES_EXTERNAL AssetOriginTypes = "EXTERNAL"
	ASSETORIGINTYPES_S3 AssetOriginTypes = "S3"
)

// All allowed values of AssetOriginTypes enum
var AllowedAssetOriginTypesEnumValues = []AssetOriginTypes{
	"EXTERNAL",
	"S3",
}

func (v *AssetOriginTypes) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := AssetOriginTypes(value)
	for _, existing := range AllowedAssetOriginTypesEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid AssetOriginTypes", value)
}

// NewAssetOriginTypesFromValue returns a pointer to a valid AssetOriginTypes
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewAssetOriginTypesFromValue(v string) (*AssetOriginTypes, error) {
	ev := AssetOriginTypes(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for AssetOriginTypes: valid values are %v", v, AllowedAssetOriginTypesEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v AssetOriginTypes) IsValid() bool {
	for _, existing := range AllowedAssetOriginTypesEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to AssetOriginTypes value
func (v AssetOriginTypes) Ptr() *AssetOriginTypes {
	return &v
}

type NullableAssetOriginTypes struct {
	value *AssetOriginTypes
	isSet bool
}

func (v NullableAssetOriginTypes) Get() *AssetOriginTypes {
	return v.value
}

func (v *NullableAssetOriginTypes) Set(val *AssetOriginTypes) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetOriginTypes) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetOriginTypes) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetOriginTypes(val *AssetOriginTypes) *NullableAssetOriginTypes {
	return &NullableAssetOriginTypes{value: val, isSet: true}
}

func (v NullableAssetOriginTypes) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetOriginTypes) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

