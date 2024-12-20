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
	"time"
)

// checks if the DamAsset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DamAsset{}

// DamAsset struct for DamAsset
type DamAsset struct {
	CreatedAt            *time.Time      `json:"createdAt,omitempty"`
	UpdatedAt            *time.Time      `json:"updatedAt,omitempty"`
	Id                   *string         `json:"id,omitempty"`
	Type                 *DamAssetType   `json:"type,omitempty"`
	Code                 *string         `json:"code,omitempty"`
	Metadata             []AssetMetadata `json:"metadata,omitempty"`
	Grn                  *string         `json:"grn,omitempty"`
	PublicUrl            *string         `json:"publicUrl,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _DamAsset DamAsset

// NewDamAsset instantiates a new DamAsset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDamAsset() *DamAsset {
	this := DamAsset{}
	var type_ DamAssetType = DAMASSETTYPE_UNKNOWN
	this.Type = &type_
	return &this
}

// NewDamAssetWithDefaults instantiates a new DamAsset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDamAssetWithDefaults() *DamAsset {
	this := DamAsset{}
	var type_ DamAssetType = DAMASSETTYPE_UNKNOWN
	this.Type = &type_
	return &this
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *DamAsset) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamAsset) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *DamAsset) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *DamAsset) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *DamAsset) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamAsset) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *DamAsset) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *DamAsset) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *DamAsset) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamAsset) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *DamAsset) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *DamAsset) SetId(v string) {
	o.Id = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DamAsset) GetType() DamAssetType {
	if o == nil || IsNil(o.Type) {
		var ret DamAssetType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamAsset) GetTypeOk() (*DamAssetType, bool) {
	if o == nil || IsNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DamAsset) HasType() bool {
	if o != nil && !IsNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given DamAssetType and assigns it to the Type field.
func (o *DamAsset) SetType(v DamAssetType) {
	o.Type = &v
}

// GetCode returns the Code field value if set, zero value otherwise.
func (o *DamAsset) GetCode() string {
	if o == nil || IsNil(o.Code) {
		var ret string
		return ret
	}
	return *o.Code
}

// GetCodeOk returns a tuple with the Code field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamAsset) GetCodeOk() (*string, bool) {
	if o == nil || IsNil(o.Code) {
		return nil, false
	}
	return o.Code, true
}

// HasCode returns a boolean if a field has been set.
func (o *DamAsset) HasCode() bool {
	if o != nil && !IsNil(o.Code) {
		return true
	}

	return false
}

// SetCode gets a reference to the given string and assigns it to the Code field.
func (o *DamAsset) SetCode(v string) {
	o.Code = &v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *DamAsset) GetMetadata() []AssetMetadata {
	if o == nil || IsNil(o.Metadata) {
		var ret []AssetMetadata
		return ret
	}
	return o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamAsset) GetMetadataOk() ([]AssetMetadata, bool) {
	if o == nil || IsNil(o.Metadata) {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *DamAsset) HasMetadata() bool {
	if o != nil && !IsNil(o.Metadata) {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given []AssetMetadata and assigns it to the Metadata field.
func (o *DamAsset) SetMetadata(v []AssetMetadata) {
	o.Metadata = v
}

// GetGrn returns the Grn field value if set, zero value otherwise.
func (o *DamAsset) GetGrn() string {
	if o == nil || IsNil(o.Grn) {
		var ret string
		return ret
	}
	return *o.Grn
}

// GetGrnOk returns a tuple with the Grn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamAsset) GetGrnOk() (*string, bool) {
	if o == nil || IsNil(o.Grn) {
		return nil, false
	}
	return o.Grn, true
}

// HasGrn returns a boolean if a field has been set.
func (o *DamAsset) HasGrn() bool {
	if o != nil && !IsNil(o.Grn) {
		return true
	}

	return false
}

// SetGrn gets a reference to the given string and assigns it to the Grn field.
func (o *DamAsset) SetGrn(v string) {
	o.Grn = &v
}

// GetPublicUrl returns the PublicUrl field value if set, zero value otherwise.
func (o *DamAsset) GetPublicUrl() string {
	if o == nil || IsNil(o.PublicUrl) {
		var ret string
		return ret
	}
	return *o.PublicUrl
}

// GetPublicUrlOk returns a tuple with the PublicUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DamAsset) GetPublicUrlOk() (*string, bool) {
	if o == nil || IsNil(o.PublicUrl) {
		return nil, false
	}
	return o.PublicUrl, true
}

// HasPublicUrl returns a boolean if a field has been set.
func (o *DamAsset) HasPublicUrl() bool {
	if o != nil && !IsNil(o.PublicUrl) {
		return true
	}

	return false
}

// SetPublicUrl gets a reference to the given string and assigns it to the PublicUrl field.
func (o *DamAsset) SetPublicUrl(v string) {
	o.PublicUrl = &v
}

func (o DamAsset) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DamAsset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !IsNil(o.Code) {
		toSerialize["code"] = o.Code
	}
	if !IsNil(o.Metadata) {
		toSerialize["metadata"] = o.Metadata
	}
	if !IsNil(o.Grn) {
		toSerialize["grn"] = o.Grn
	}
	if !IsNil(o.PublicUrl) {
		toSerialize["publicUrl"] = o.PublicUrl
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *DamAsset) UnmarshalJSON(data []byte) (err error) {
	varDamAsset := _DamAsset{}

	err = json.Unmarshal(data, &varDamAsset)

	if err != nil {
		return err
	}

	*o = DamAsset(varDamAsset)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "createdAt")
		delete(additionalProperties, "updatedAt")
		delete(additionalProperties, "id")
		delete(additionalProperties, "type")
		delete(additionalProperties, "code")
		delete(additionalProperties, "metadata")
		delete(additionalProperties, "grn")
		delete(additionalProperties, "publicUrl")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

// GetValue returns the value of well-known types
func (o *DamAsset) GetValue() interface{} {
	if o == nil || IsNil(o.AdditionalProperties) {
		return nil
	}
	return o.AdditionalProperties["value"]
}

// SetValue populates the value of well-known types
func (o *DamAsset) SetValue(value interface{}) {
	if o == nil || IsNil(value) {
		return
	}
	if IsNil(o.AdditionalProperties) {
		o.AdditionalProperties = map[string]interface{}{}
	}
	o.AdditionalProperties["value"] = value
	return
}

type NullableDamAsset struct {
	value *DamAsset
	isSet bool
}

func (v NullableDamAsset) Get() *DamAsset {
	return v.value
}

func (v *NullableDamAsset) Set(val *DamAsset) {
	v.value = val
	v.isSet = true
}

func (v NullableDamAsset) IsSet() bool {
	return v.isSet
}

func (v *NullableDamAsset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDamAsset(val *DamAsset) *NullableDamAsset {
	return &NullableDamAsset{value: val, isSet: true}
}

func (v NullableDamAsset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDamAsset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
