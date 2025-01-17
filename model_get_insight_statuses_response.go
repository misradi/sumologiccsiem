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

// checks if the GetInsightStatusesResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetInsightStatusesResponse{}

// GetInsightStatusesResponse struct for GetInsightStatusesResponse
type GetInsightStatusesResponse struct {
	// All available Insight statuses.
	Data []InsightStatus `json:"data"`
}

type _GetInsightStatusesResponse GetInsightStatusesResponse

// NewGetInsightStatusesResponse instantiates a new GetInsightStatusesResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetInsightStatusesResponse(data []InsightStatus) *GetInsightStatusesResponse {
	this := GetInsightStatusesResponse{}
	this.Data = data
	return &this
}

// NewGetInsightStatusesResponseWithDefaults instantiates a new GetInsightStatusesResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetInsightStatusesResponseWithDefaults() *GetInsightStatusesResponse {
	this := GetInsightStatusesResponse{}
	return &this
}

// GetData returns the Data field value
func (o *GetInsightStatusesResponse) GetData() []InsightStatus {
	if o == nil {
		var ret []InsightStatus
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetInsightStatusesResponse) GetDataOk() ([]InsightStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *GetInsightStatusesResponse) SetData(v []InsightStatus) {
	o.Data = v
}

func (o GetInsightStatusesResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetInsightStatusesResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetInsightStatusesResponse) UnmarshalJSON(data []byte) (err error) {
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

	varGetInsightStatusesResponse := _GetInsightStatusesResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetInsightStatusesResponse)

	if err != nil {
		return err
	}

	*o = GetInsightStatusesResponse(varGetInsightStatusesResponse)

	return err
}

type NullableGetInsightStatusesResponse struct {
	value *GetInsightStatusesResponse
	isSet bool
}

func (v NullableGetInsightStatusesResponse) Get() *GetInsightStatusesResponse {
	return v.value
}

func (v *NullableGetInsightStatusesResponse) Set(val *GetInsightStatusesResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableGetInsightStatusesResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableGetInsightStatusesResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetInsightStatusesResponse(val *GetInsightStatusesResponse) *NullableGetInsightStatusesResponse {
	return &NullableGetInsightStatusesResponse{value: val, isSet: true}
}

func (v NullableGetInsightStatusesResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetInsightStatusesResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


