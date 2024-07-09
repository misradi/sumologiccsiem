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

// checks if the DeleteThreatIntelIndicatorResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteThreatIntelIndicatorResponse{}

// DeleteThreatIntelIndicatorResponse struct for DeleteThreatIntelIndicatorResponse
type DeleteThreatIntelIndicatorResponse struct {
	Data DeleteContextActionResponseData `json:"data"`
}

type _DeleteThreatIntelIndicatorResponse DeleteThreatIntelIndicatorResponse

// NewDeleteThreatIntelIndicatorResponse instantiates a new DeleteThreatIntelIndicatorResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteThreatIntelIndicatorResponse(data DeleteContextActionResponseData) *DeleteThreatIntelIndicatorResponse {
	this := DeleteThreatIntelIndicatorResponse{}
	this.Data = data
	return &this
}

// NewDeleteThreatIntelIndicatorResponseWithDefaults instantiates a new DeleteThreatIntelIndicatorResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteThreatIntelIndicatorResponseWithDefaults() *DeleteThreatIntelIndicatorResponse {
	this := DeleteThreatIntelIndicatorResponse{}
	return &this
}

// GetData returns the Data field value
func (o *DeleteThreatIntelIndicatorResponse) GetData() DeleteContextActionResponseData {
	if o == nil {
		var ret DeleteContextActionResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeleteThreatIntelIndicatorResponse) GetDataOk() (*DeleteContextActionResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DeleteThreatIntelIndicatorResponse) SetData(v DeleteContextActionResponseData) {
	o.Data = v
}

func (o DeleteThreatIntelIndicatorResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteThreatIntelIndicatorResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *DeleteThreatIntelIndicatorResponse) UnmarshalJSON(data []byte) (err error) {
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

	varDeleteThreatIntelIndicatorResponse := _DeleteThreatIntelIndicatorResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteThreatIntelIndicatorResponse)

	if err != nil {
		return err
	}

	*o = DeleteThreatIntelIndicatorResponse(varDeleteThreatIntelIndicatorResponse)

	return err
}

type NullableDeleteThreatIntelIndicatorResponse struct {
	value *DeleteThreatIntelIndicatorResponse
	isSet bool
}

func (v NullableDeleteThreatIntelIndicatorResponse) Get() *DeleteThreatIntelIndicatorResponse {
	return v.value
}

func (v *NullableDeleteThreatIntelIndicatorResponse) Set(val *DeleteThreatIntelIndicatorResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteThreatIntelIndicatorResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteThreatIntelIndicatorResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteThreatIntelIndicatorResponse(val *DeleteThreatIntelIndicatorResponse) *NullableDeleteThreatIntelIndicatorResponse {
	return &NullableDeleteThreatIntelIndicatorResponse{value: val, isSet: true}
}

func (v NullableDeleteThreatIntelIndicatorResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteThreatIntelIndicatorResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


