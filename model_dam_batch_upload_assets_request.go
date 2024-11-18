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

// checks if the DamBatchUploadAssetsRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DamBatchUploadAssetsRequest{}

// DamBatchUploadAssetsRequest struct for DamBatchUploadAssetsRequest
type DamBatchUploadAssetsRequest struct {
	TenantId             string                          `json:"tenantId"`
	Files                []BatchUploadAssetsRequestFiles `json:"files"`
	AdditionalProperties map[string]interface{}
}

type _DamBatchUploadAssetsRequest DamBatchUploadAssetsRequest

// NewDamBatchUploadAssetsRequest instantiates a new DamBatchUploadAssetsRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDamBatchUploadAssetsRequest(tenantId string, files []BatchUploadAssetsRequestFiles) *DamBatchUploadAssetsRequest {
	this := DamBatchUploadAssetsRequest{}
	this.TenantId = tenantId
	this.Files = files
	return &this
}

// NewDamBatchUploadAssetsRequestWithDefaults instantiates a new DamBatchUploadAssetsRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDamBatchUploadAssetsRequestWithDefaults() *DamBatchUploadAssetsRequest {
	this := DamBatchUploadAssetsRequest{}
	return &this
}

// GetTenantId returns the TenantId field value
func (o *DamBatchUploadAssetsRequest) GetTenantId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value
// and a boolean to check if the value has been set.
func (o *DamBatchUploadAssetsRequest) GetTenantIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TenantId, true
}

// SetTenantId sets field value
func (o *DamBatchUploadAssetsRequest) SetTenantId(v string) {
	o.TenantId = v
}

// GetFiles returns the Files field value
func (o *DamBatchUploadAssetsRequest) GetFiles() []BatchUploadAssetsRequestFiles {
	if o == nil {
		var ret []BatchUploadAssetsRequestFiles
		return ret
	}

	return o.Files
}

// GetFilesOk returns a tuple with the Files field value
// and a boolean to check if the value has been set.
func (o *DamBatchUploadAssetsRequest) GetFilesOk() ([]BatchUploadAssetsRequestFiles, bool) {
	if o == nil {
		return nil, false
	}
	return o.Files, true
}

// SetFiles sets field value
func (o *DamBatchUploadAssetsRequest) SetFiles(v []BatchUploadAssetsRequestFiles) {
	o.Files = v
}

func (o DamBatchUploadAssetsRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DamBatchUploadAssetsRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["tenantId"] = o.TenantId
	toSerialize["files"] = o.Files

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DamBatchUploadAssetsRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"tenantId",
		"files",
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

	varDamBatchUploadAssetsRequest := _DamBatchUploadAssetsRequest{}

	err = json.Unmarshal(data, &varDamBatchUploadAssetsRequest)

	if err != nil {
		return err
	}

	*o = DamBatchUploadAssetsRequest(varDamBatchUploadAssetsRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "tenantId")
		delete(additionalProperties, "files")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *DamBatchUploadAssetsRequest) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *DamBatchUploadAssetsRequest) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableDamBatchUploadAssetsRequest struct {
	value *DamBatchUploadAssetsRequest
	isSet bool
}

func (v NullableDamBatchUploadAssetsRequest) Get() *DamBatchUploadAssetsRequest {
	return v.value
}

func (v *NullableDamBatchUploadAssetsRequest) Set(val *DamBatchUploadAssetsRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableDamBatchUploadAssetsRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableDamBatchUploadAssetsRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDamBatchUploadAssetsRequest(val *DamBatchUploadAssetsRequest) *NullableDamBatchUploadAssetsRequest {
	return &NullableDamBatchUploadAssetsRequest{value: val, isSet: true}
}

func (v NullableDamBatchUploadAssetsRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDamBatchUploadAssetsRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
