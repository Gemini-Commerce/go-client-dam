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

// checks if the DamListAssetsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DamListAssetsResponse{}

// DamListAssetsResponse struct for DamListAssetsResponse
type DamListAssetsResponse struct {
	Assets []DamAsset `json:"assets,omitempty"`
	// A token that can be sent as `page_token` to retrieve the next page. If this field is omitted, there are no subsequent pages.
	NextPageToken *string `json:"nextPageToken,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DamListAssetsResponse DamListAssetsResponse

// NewDamListAssetsResponse instantiates a new DamListAssetsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDamListAssetsResponse() *DamListAssetsResponse {
	this := DamListAssetsResponse{}
	return &this
}

// NewDamListAssetsResponseWithDefaults instantiates a new DamListAssetsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDamListAssetsResponseWithDefaults() *DamListAssetsResponse {
	this := DamListAssetsResponse{}
	return &this
}

// GetAssets returns the Assets field value if set, zero value otherwise.
func (o *DamListAssetsResponse) GetAssets() []DamAsset {
	if o == nil || IsNil(o.Assets) {
		var ret []DamAsset
		return ret
	}
	return o.Assets
}

// GetAssetsOk returns a tuple with the Assets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamListAssetsResponse) GetAssetsOk() ([]DamAsset, bool) {
	if o == nil || IsNil(o.Assets) {
		return nil, false
	}
	return o.Assets, true
}

// HasAssets returns a boolean if a field has been set.
func (o *DamListAssetsResponse) HasAssets() bool {
	if o != nil && !IsNil(o.Assets) {
		return true
	}

	return false
}

// SetAssets gets a reference to the given []DamAsset and assigns it to the Assets field.
func (o *DamListAssetsResponse) SetAssets(v []DamAsset) {
	o.Assets = v
}

// GetNextPageToken returns the NextPageToken field value if set, zero value otherwise.
func (o *DamListAssetsResponse) GetNextPageToken() string {
	if o == nil || IsNil(o.NextPageToken) {
		var ret string
		return ret
	}
	return *o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamListAssetsResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil || IsNil(o.NextPageToken) {
		return nil, false
	}
	return o.NextPageToken, true
}

// HasNextPageToken returns a boolean if a field has been set.
func (o *DamListAssetsResponse) HasNextPageToken() bool {
	if o != nil && !IsNil(o.NextPageToken) {
		return true
	}

	return false
}

// SetNextPageToken gets a reference to the given string and assigns it to the NextPageToken field.
func (o *DamListAssetsResponse) SetNextPageToken(v string) {
	o.NextPageToken = &v
}

func (o DamListAssetsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DamListAssetsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Assets) {
		toSerialize["assets"] = o.Assets
	}
	if !IsNil(o.NextPageToken) {
		toSerialize["nextPageToken"] = o.NextPageToken
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DamListAssetsResponse) UnmarshalJSON(data []byte) (err error) {
	varDamListAssetsResponse := _DamListAssetsResponse{}

	err = json.Unmarshal(data, &varDamListAssetsResponse)

	if err != nil {
		return err
	}

	*o = DamListAssetsResponse(varDamListAssetsResponse)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "assets")
		delete(additionalProperties, "nextPageToken")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *DamListAssetsResponse) GetValue() interface{} {
	if o == nil || IsNil(o.Type) || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}
// SetValue populate the value of well-known types
func (o *DamListAssetsResponse) SetValue(value interface{}) {
	if o == nil || IsNil(o.Type) || IsNil(value) {
		return
	}
    if IsNil(o.AdditionalProperties) {
        o.AdditionalProperties = map[string]interface{}{}
    }
	o.AdditionalProperties["value"] = value
	return
}
type NullableDamListAssetsResponse struct {
	value *DamListAssetsResponse
	isSet bool
}

func (v NullableDamListAssetsResponse) Get() *DamListAssetsResponse {
	return v.value
}

func (v *NullableDamListAssetsResponse) Set(val *DamListAssetsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDamListAssetsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDamListAssetsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDamListAssetsResponse(val *DamListAssetsResponse) *NullableDamListAssetsResponse {
	return &NullableDamListAssetsResponse{value: val, isSet: true}
}

func (v NullableDamListAssetsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDamListAssetsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


