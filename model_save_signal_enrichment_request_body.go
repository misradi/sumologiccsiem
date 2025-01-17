/*
Sumo Logic CSE API

 https://help.sumologic.com/APIs 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologiccsiem

import (
	"encoding/json"
	"time"
	"bytes"
	"fmt"
)

// checks if the SaveSignalEnrichmentRequestBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaveSignalEnrichmentRequestBody{}

// SaveSignalEnrichmentRequestBody struct for SaveSignalEnrichmentRequestBody
type SaveSignalEnrichmentRequestBody struct {
	Detail map[string]interface{} `json:"detail"`
	Raw *string `json:"raw,omitempty"`
	ExpiresAt *time.Time `json:"expiresAt,omitempty"`
	Reputation *string `json:"reputation,omitempty"`
	ExternalUrl *string `json:"externalUrl,omitempty"`
}

type _SaveSignalEnrichmentRequestBody SaveSignalEnrichmentRequestBody

// NewSaveSignalEnrichmentRequestBody instantiates a new SaveSignalEnrichmentRequestBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaveSignalEnrichmentRequestBody(detail map[string]interface{}) *SaveSignalEnrichmentRequestBody {
	this := SaveSignalEnrichmentRequestBody{}
	this.Detail = detail
	return &this
}

// NewSaveSignalEnrichmentRequestBodyWithDefaults instantiates a new SaveSignalEnrichmentRequestBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaveSignalEnrichmentRequestBodyWithDefaults() *SaveSignalEnrichmentRequestBody {
	this := SaveSignalEnrichmentRequestBody{}
	return &this
}

// GetDetail returns the Detail field value
func (o *SaveSignalEnrichmentRequestBody) GetDetail() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Detail
}

// GetDetailOk returns a tuple with the Detail field value
// and a boolean to check if the value has been set.
func (o *SaveSignalEnrichmentRequestBody) GetDetailOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Detail, true
}

// SetDetail sets field value
func (o *SaveSignalEnrichmentRequestBody) SetDetail(v map[string]interface{}) {
	o.Detail = v
}

// GetRaw returns the Raw field value if set, zero value otherwise.
func (o *SaveSignalEnrichmentRequestBody) GetRaw() string {
	if o == nil || IsNil(o.Raw) {
		var ret string
		return ret
	}
	return *o.Raw
}

// GetRawOk returns a tuple with the Raw field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveSignalEnrichmentRequestBody) GetRawOk() (*string, bool) {
	if o == nil || IsNil(o.Raw) {
		return nil, false
	}
	return o.Raw, true
}

// HasRaw returns a boolean if a field has been set.
func (o *SaveSignalEnrichmentRequestBody) HasRaw() bool {
	if o != nil && !IsNil(o.Raw) {
		return true
	}

	return false
}

// SetRaw gets a reference to the given string and assigns it to the Raw field.
func (o *SaveSignalEnrichmentRequestBody) SetRaw(v string) {
	o.Raw = &v
}

// GetExpiresAt returns the ExpiresAt field value if set, zero value otherwise.
func (o *SaveSignalEnrichmentRequestBody) GetExpiresAt() time.Time {
	if o == nil || IsNil(o.ExpiresAt) {
		var ret time.Time
		return ret
	}
	return *o.ExpiresAt
}

// GetExpiresAtOk returns a tuple with the ExpiresAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveSignalEnrichmentRequestBody) GetExpiresAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ExpiresAt) {
		return nil, false
	}
	return o.ExpiresAt, true
}

// HasExpiresAt returns a boolean if a field has been set.
func (o *SaveSignalEnrichmentRequestBody) HasExpiresAt() bool {
	if o != nil && !IsNil(o.ExpiresAt) {
		return true
	}

	return false
}

// SetExpiresAt gets a reference to the given time.Time and assigns it to the ExpiresAt field.
func (o *SaveSignalEnrichmentRequestBody) SetExpiresAt(v time.Time) {
	o.ExpiresAt = &v
}

// GetReputation returns the Reputation field value if set, zero value otherwise.
func (o *SaveSignalEnrichmentRequestBody) GetReputation() string {
	if o == nil || IsNil(o.Reputation) {
		var ret string
		return ret
	}
	return *o.Reputation
}

// GetReputationOk returns a tuple with the Reputation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveSignalEnrichmentRequestBody) GetReputationOk() (*string, bool) {
	if o == nil || IsNil(o.Reputation) {
		return nil, false
	}
	return o.Reputation, true
}

// HasReputation returns a boolean if a field has been set.
func (o *SaveSignalEnrichmentRequestBody) HasReputation() bool {
	if o != nil && !IsNil(o.Reputation) {
		return true
	}

	return false
}

// SetReputation gets a reference to the given string and assigns it to the Reputation field.
func (o *SaveSignalEnrichmentRequestBody) SetReputation(v string) {
	o.Reputation = &v
}

// GetExternalUrl returns the ExternalUrl field value if set, zero value otherwise.
func (o *SaveSignalEnrichmentRequestBody) GetExternalUrl() string {
	if o == nil || IsNil(o.ExternalUrl) {
		var ret string
		return ret
	}
	return *o.ExternalUrl
}

// GetExternalUrlOk returns a tuple with the ExternalUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SaveSignalEnrichmentRequestBody) GetExternalUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalUrl) {
		return nil, false
	}
	return o.ExternalUrl, true
}

// HasExternalUrl returns a boolean if a field has been set.
func (o *SaveSignalEnrichmentRequestBody) HasExternalUrl() bool {
	if o != nil && !IsNil(o.ExternalUrl) {
		return true
	}

	return false
}

// SetExternalUrl gets a reference to the given string and assigns it to the ExternalUrl field.
func (o *SaveSignalEnrichmentRequestBody) SetExternalUrl(v string) {
	o.ExternalUrl = &v
}

func (o SaveSignalEnrichmentRequestBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaveSignalEnrichmentRequestBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["detail"] = o.Detail
	if !IsNil(o.Raw) {
		toSerialize["raw"] = o.Raw
	}
	if !IsNil(o.ExpiresAt) {
		toSerialize["expiresAt"] = o.ExpiresAt
	}
	if !IsNil(o.Reputation) {
		toSerialize["reputation"] = o.Reputation
	}
	if !IsNil(o.ExternalUrl) {
		toSerialize["externalUrl"] = o.ExternalUrl
	}
	return toSerialize, nil
}

func (o *SaveSignalEnrichmentRequestBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"detail",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSaveSignalEnrichmentRequestBody := _SaveSignalEnrichmentRequestBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSaveSignalEnrichmentRequestBody)

	if err != nil {
		return err
	}

	*o = SaveSignalEnrichmentRequestBody(varSaveSignalEnrichmentRequestBody)

	return err
}

type NullableSaveSignalEnrichmentRequestBody struct {
	value *SaveSignalEnrichmentRequestBody
	isSet bool
}

func (v NullableSaveSignalEnrichmentRequestBody) Get() *SaveSignalEnrichmentRequestBody {
	return v.value
}

func (v *NullableSaveSignalEnrichmentRequestBody) Set(val *SaveSignalEnrichmentRequestBody) {
	v.value = val
	v.isSet = true
}

func (v NullableSaveSignalEnrichmentRequestBody) IsSet() bool {
	return v.isSet
}

func (v *NullableSaveSignalEnrichmentRequestBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaveSignalEnrichmentRequestBody(val *SaveSignalEnrichmentRequestBody) *NullableSaveSignalEnrichmentRequestBody {
	return &NullableSaveSignalEnrichmentRequestBody{value: val, isSet: true}
}

func (v NullableSaveSignalEnrichmentRequestBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaveSignalEnrichmentRequestBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


