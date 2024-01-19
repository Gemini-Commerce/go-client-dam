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

// checks if the DamBatchUploadAssetsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DamBatchUploadAssetsResponse{}

// DamBatchUploadAssetsResponse struct for DamBatchUploadAssetsResponse
type DamBatchUploadAssetsResponse struct {
	PreSignedUrls []string `json:"preSignedUrls,omitempty"`
}

// NewDamBatchUploadAssetsResponse instantiates a new DamBatchUploadAssetsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDamBatchUploadAssetsResponse() *DamBatchUploadAssetsResponse {
	this := DamBatchUploadAssetsResponse{}
	return &this
}

// NewDamBatchUploadAssetsResponseWithDefaults instantiates a new DamBatchUploadAssetsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDamBatchUploadAssetsResponseWithDefaults() *DamBatchUploadAssetsResponse {
	this := DamBatchUploadAssetsResponse{}
	return &this
}

// GetPreSignedUrls returns the PreSignedUrls field value if set, zero value otherwise.
func (o *DamBatchUploadAssetsResponse) GetPreSignedUrls() []string {
	if o == nil || IsNil(o.PreSignedUrls) {
		var ret []string
		return ret
	}
	return o.PreSignedUrls
}

// GetPreSignedUrlsOk returns a tuple with the PreSignedUrls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamBatchUploadAssetsResponse) GetPreSignedUrlsOk() ([]string, bool) {
	if o == nil || IsNil(o.PreSignedUrls) {
		return nil, false
	}
	return o.PreSignedUrls, true
}

// HasPreSignedUrls returns a boolean if a field has been set.
func (o *DamBatchUploadAssetsResponse) HasPreSignedUrls() bool {
	if o != nil && !IsNil(o.PreSignedUrls) {
		return true
	}

	return false
}

// SetPreSignedUrls gets a reference to the given []string and assigns it to the PreSignedUrls field.
func (o *DamBatchUploadAssetsResponse) SetPreSignedUrls(v []string) {
	o.PreSignedUrls = v
}

func (o DamBatchUploadAssetsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DamBatchUploadAssetsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.PreSignedUrls) {
		toSerialize["preSignedUrls"] = o.PreSignedUrls
	}
	return toSerialize, nil
}

type NullableDamBatchUploadAssetsResponse struct {
	value *DamBatchUploadAssetsResponse
	isSet bool
}

func (v NullableDamBatchUploadAssetsResponse) Get() *DamBatchUploadAssetsResponse {
	return v.value
}

func (v *NullableDamBatchUploadAssetsResponse) Set(val *DamBatchUploadAssetsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDamBatchUploadAssetsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDamBatchUploadAssetsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDamBatchUploadAssetsResponse(val *DamBatchUploadAssetsResponse) *NullableDamBatchUploadAssetsResponse {
	return &NullableDamBatchUploadAssetsResponse{value: val, isSet: true}
}

func (v NullableDamBatchUploadAssetsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDamBatchUploadAssetsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

