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
	"fmt"
)

// checks if the UpdateAssetRequestPayload type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UpdateAssetRequestPayload{}

// UpdateAssetRequestPayload struct for UpdateAssetRequestPayload
type UpdateAssetRequestPayload struct {
	Code                 *string         `json:"code,omitempty"`
	Origin               DamAssetOrigin  `json:"origin"`
	Metadata             []AssetMetadata `json:"metadata,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _UpdateAssetRequestPayload UpdateAssetRequestPayload

// NewUpdateAssetRequestPayload instantiates a new UpdateAssetRequestPayload object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateAssetRequestPayload(origin DamAssetOrigin) *UpdateAssetRequestPayload {
	this := UpdateAssetRequestPayload{}
	this.Origin = origin
	return &this
}

// NewUpdateAssetRequestPayloadWithDefaults instantiates a new UpdateAssetRequestPayload object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateAssetRequestPayloadWithDefaults() *UpdateAssetRequestPayload {
	this := UpdateAssetRequestPayload{}
	return &this
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *UpdateAssetRequestPayload) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAssetRequestPayload) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *UpdateAssetRequestPayload) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *UpdateAssetRequestPayload) SetCode(v string) {
	o.Code = &v
}

// GetOrigin returns the Origin field value
func (o *UpdateAssetRequestPayload) GetOrigin() DamAssetOrigin {
	if o == nil {
		var ret DamAssetOrigin
		return ret
	}

	return o.Origin
}

// GetOriginOk returns a tuple with the Origin field value
// and a boolean to check if the value has been set.
func (o *UpdateAssetRequestPayload) GetOriginOk() (*DamAssetOrigin, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Origin, true
}

// SetOrigin sets field value
func (o *UpdateAssetRequestPayload) SetOrigin(v DamAssetOrigin) {
	o.Origin = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UpdateAssetRequestPayload) GetMetadata() []AssetMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret []AssetMetadata
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateAssetRequestPayload) GetMetadataOk() ([]AssetMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UpdateAssetRequestPayload) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []AssetMetadata and assigns it to the Metadata field.
func (o *UpdateAssetRequestPayload) SetMetadata(v []AssetMetadata) {
	o.Metadata = v
}

func (o UpdateAssetRequestPayload) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UpdateAssetRequestPayload) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	toSerialize["origin"] = o.Origin
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *UpdateAssetRequestPayload) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"origin",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUpdateAssetRequestPayload := _UpdateAssetRequestPayload{}

	err = json.Unmarshal(data, &varUpdateAssetRequestPayload)

	if err != nil {
		return err
	}

	*o = UpdateAssetRequestPayload(varUpdateAssetRequestPayload)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "code")
		delete(additionalProperties, "origin")
		delete(additionalProperties, "metadata")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *UpdateAssetRequestPayload) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *UpdateAssetRequestPayload) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableUpdateAssetRequestPayload struct {
	value *UpdateAssetRequestPayload
	isSet bool
}

func (v NullableUpdateAssetRequestPayload) Get() *UpdateAssetRequestPayload {
	return v.value
}

func (v *NullableUpdateAssetRequestPayload) Set(val *UpdateAssetRequestPayload) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateAssetRequestPayload) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateAssetRequestPayload) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateAssetRequestPayload(val *UpdateAssetRequestPayload) *NullableUpdateAssetRequestPayload {
	return &NullableUpdateAssetRequestPayload{value: val, isSet: true}
}

func (v NullableUpdateAssetRequestPayload) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateAssetRequestPayload) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
