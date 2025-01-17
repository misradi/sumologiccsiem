/*
Sumo Logic CSE API

 https://help.sumologic.com/APIs 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sumologiccsiem

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the SaveEntityEnrichmentResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SaveEntityEnrichmentResponse{}

// SaveEntityEnrichmentResponse struct for SaveEntityEnrichmentResponse
type SaveEntityEnrichmentResponse struct {
	Data Enrichment `json:"data"`
}

type _SaveEntityEnrichmentResponse SaveEntityEnrichmentResponse

// NewSaveEntityEnrichmentResponse instantiates a new SaveEntityEnrichmentResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSaveEntityEnrichmentResponse(data Enrichment) *SaveEntityEnrichmentResponse {
	this := SaveEntityEnrichmentResponse{}
	this.Data = data
	return &this
}

// NewSaveEntityEnrichmentResponseWithDefaults instantiates a new SaveEntityEnrichmentResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSaveEntityEnrichmentResponseWithDefaults() *SaveEntityEnrichmentResponse {
	this := SaveEntityEnrichmentResponse{}
	return &this
}

// GetData returns the Data field value
func (o *SaveEntityEnrichmentResponse) GetData() Enrichment {
	if o == nil {
		var ret Enrichment
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *SaveEntityEnrichmentResponse) GetDataOk() (*Enrichment, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *SaveEntityEnrichmentResponse) SetData(v Enrichment) {
	o.Data = v
}

func (o SaveEntityEnrichmentResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SaveEntityEnrichmentResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *SaveEntityEnrichmentResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varSaveEntityEnrichmentResponse := _SaveEntityEnrichmentResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSaveEntityEnrichmentResponse)

	if err != nil {
		return err
	}

	*o = SaveEntityEnrichmentResponse(varSaveEntityEnrichmentResponse)

	return err
}

type NullableSaveEntityEnrichmentResponse struct {
	value *SaveEntityEnrichmentResponse
	isSet bool
}

func (v NullableSaveEntityEnrichmentResponse) Get() *SaveEntityEnrichmentResponse {
	return v.value
}

func (v *NullableSaveEntityEnrichmentResponse) Set(val *SaveEntityEnrichmentResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSaveEntityEnrichmentResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSaveEntityEnrichmentResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSaveEntityEnrichmentResponse(val *SaveEntityEnrichmentResponse) *NullableSaveEntityEnrichmentResponse {
	return &NullableSaveEntityEnrichmentResponse{value: val, isSet: true}
}

func (v NullableSaveEntityEnrichmentResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSaveEntityEnrichmentResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


