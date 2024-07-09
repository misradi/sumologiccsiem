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

// checks if the BulkAddEntityTagsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BulkAddEntityTagsResponse{}

// BulkAddEntityTagsResponse struct for BulkAddEntityTagsResponse
type BulkAddEntityTagsResponse struct {
	Data []BulkAddEntityTagsResponseDataInner `json:"data"`
}

type _BulkAddEntityTagsResponse BulkAddEntityTagsResponse

// NewBulkAddEntityTagsResponse instantiates a new BulkAddEntityTagsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBulkAddEntityTagsResponse(data []BulkAddEntityTagsResponseDataInner) *BulkAddEntityTagsResponse {
	this := BulkAddEntityTagsResponse{}
	this.Data = data
	return &this
}

// NewBulkAddEntityTagsResponseWithDefaults instantiates a new BulkAddEntityTagsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBulkAddEntityTagsResponseWithDefaults() *BulkAddEntityTagsResponse {
	this := BulkAddEntityTagsResponse{}
	return &this
}

// GetData returns the Data field value
func (o *BulkAddEntityTagsResponse) GetData() []BulkAddEntityTagsResponseDataInner {
	if o == nil {
		var ret []BulkAddEntityTagsResponseDataInner
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *BulkAddEntityTagsResponse) GetDataOk() ([]BulkAddEntityTagsResponseDataInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Data, true
}

// SetData sets field value
func (o *BulkAddEntityTagsResponse) SetData(v []BulkAddEntityTagsResponseDataInner) {
	o.Data = v
}

func (o BulkAddEntityTagsResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BulkAddEntityTagsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *BulkAddEntityTagsResponse) UnmarshalJSON(data []byte) (err error) {
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

	varBulkAddEntityTagsResponse := _BulkAddEntityTagsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBulkAddEntityTagsResponse)

	if err != nil {
		return err
	}

	*o = BulkAddEntityTagsResponse(varBulkAddEntityTagsResponse)

	return err
}

type NullableBulkAddEntityTagsResponse struct {
	value *BulkAddEntityTagsResponse
	isSet bool
}

func (v NullableBulkAddEntityTagsResponse) Get() *BulkAddEntityTagsResponse {
	return v.value
}

func (v *NullableBulkAddEntityTagsResponse) Set(val *BulkAddEntityTagsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBulkAddEntityTagsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBulkAddEntityTagsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBulkAddEntityTagsResponse(val *BulkAddEntityTagsResponse) *NullableBulkAddEntityTagsResponse {
	return &NullableBulkAddEntityTagsResponse{value: val, isSet: true}
}

func (v NullableBulkAddEntityTagsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBulkAddEntityTagsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


