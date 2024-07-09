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

// checks if the BulkDeleteMatchListItemsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkDeleteMatchListItemsResponse{}

// BulkDeleteMatchListItemsResponse struct for BulkDeleteMatchListItemsResponse
type BulkDeleteMatchListItemsResponse struct {
	Data BulkDeleteMatchListItemsResponseData `json:"data"`
}

type _BulkDeleteMatchListItemsResponse BulkDeleteMatchListItemsResponse

// NewBulkDeleteMatchListItemsResponse instantiates a new BulkDeleteMatchListItemsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkDeleteMatchListItemsResponse(data BulkDeleteMatchListItemsResponseData) *BulkDeleteMatchListItemsResponse {
	this := BulkDeleteMatchListItemsResponse{}
	this.Data = data
	return &this
}

// NewBulkDeleteMatchListItemsResponseWithDefaults instantiates a new BulkDeleteMatchListItemsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkDeleteMatchListItemsResponseWithDefaults() *BulkDeleteMatchListItemsResponse {
	this := BulkDeleteMatchListItemsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BulkDeleteMatchListItemsResponse) GetData() BulkDeleteMatchListItemsResponseData {
	if o == nil {
		var ret BulkDeleteMatchListItemsResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BulkDeleteMatchListItemsResponse) GetDataOk() (*BulkDeleteMatchListItemsResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *BulkDeleteMatchListItemsResponse) SetData(v BulkDeleteMatchListItemsResponseData) {
	o.Data = v
}

func (o BulkDeleteMatchListItemsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkDeleteMatchListItemsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *BulkDeleteMatchListItemsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varBulkDeleteMatchListItemsResponse := _BulkDeleteMatchListItemsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBulkDeleteMatchListItemsResponse)

	if err != nil {
		return err
	}

	*o = BulkDeleteMatchListItemsResponse(varBulkDeleteMatchListItemsResponse)

	return err
}

type NullableBulkDeleteMatchListItemsResponse struct {
	value *BulkDeleteMatchListItemsResponse
	isSet bool
}

func (v NullableBulkDeleteMatchListItemsResponse) Get() *BulkDeleteMatchListItemsResponse {
	return v.value
}

func (v *NullableBulkDeleteMatchListItemsResponse) Set(val *BulkDeleteMatchListItemsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkDeleteMatchListItemsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkDeleteMatchListItemsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkDeleteMatchListItemsResponse(val *BulkDeleteMatchListItemsResponse) *NullableBulkDeleteMatchListItemsResponse {
	return &NullableBulkDeleteMatchListItemsResponse{value: val, isSet: true}
}

func (v NullableBulkDeleteMatchListItemsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkDeleteMatchListItemsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


