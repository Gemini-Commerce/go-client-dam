/*
Dam Service

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 1.0.0
Contact: info@gemini-commerce.com
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dam

import (
	"encoding/json"
)

// checks if the AssetMetadata type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssetMetadata{}

// AssetMetadata struct for AssetMetadata
type AssetMetadata struct {
	Key                  *string `json:"key,omitempty"`
	Value                *string `json:"value,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AssetMetadata AssetMetadata

// NewAssetMetadata instantiates a new AssetMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetMetadata() *AssetMetadata {
	this := AssetMetadata{}
	return &this
}

// NewAssetMetadataWithDefaults instantiates a new AssetMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetMetadataWithDefaults() *AssetMetadata {
	this := AssetMetadata{}
	return &this
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *AssetMetadata) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetMetadata) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *AssetMetadata) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *AssetMetadata) SetKey(v string) {
	o.Key = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *AssetMetadata) GetValue() string {
	if o == nil || IsNil(o.Value) {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetMetadata) GetValueOk() (*string, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *AssetMetadata) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *AssetMetadata) SetValue(v string) {
	o.Value = &v
}

func (o AssetMetadata) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssetMetadata) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *AssetMetadata) UnmarshalJSON(data []byte) (err error) {
	varAssetMetadata := _AssetMetadata{}

	err = json.Unmarshal(data, &varAssetMetadata)

	if err != nil {
		return err
	}

	*o = AssetMetadata(varAssetMetadata)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "key")
		delete(additionalProperties, "value")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAssetMetadata struct {
	value *AssetMetadata
	isSet bool
}

func (v NullableAssetMetadata) Get() *AssetMetadata {
	return v.value
}

func (v *NullableAssetMetadata) Set(val *AssetMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetMetadata(val *AssetMetadata) *NullableAssetMetadata {
	return &NullableAssetMetadata{value: val, isSet: true}
}

func (v NullableAssetMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
