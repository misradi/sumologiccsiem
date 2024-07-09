/*
Sumo Logic CSE API

 https://help.sumologic.com/APIs 

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the DeleteCustomerSourcedEntityLookupTableResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeleteCustomerSourcedEntityLookupTableResponse{}

// DeleteCustomerSourcedEntityLookupTableResponse struct for DeleteCustomerSourcedEntityLookupTableResponse
type DeleteCustomerSourcedEntityLookupTableResponse struct {
	Data DeleteContextActionResponseData `json:"data"`
}

type _DeleteCustomerSourcedEntityLookupTableResponse DeleteCustomerSourcedEntityLookupTableResponse

// NewDeleteCustomerSourcedEntityLookupTableResponse instantiates a new DeleteCustomerSourcedEntityLookupTableResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeleteCustomerSourcedEntityLookupTableResponse(data DeleteContextActionResponseData) *DeleteCustomerSourcedEntityLookupTableResponse {
	this := DeleteCustomerSourcedEntityLookupTableResponse{}
	this.Data = data
	return &this
}

// NewDeleteCustomerSourcedEntityLookupTableResponseWithDefaults instantiates a new DeleteCustomerSourcedEntityLookupTableResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeleteCustomerSourcedEntityLookupTableResponseWithDefaults() *DeleteCustomerSourcedEntityLookupTableResponse {
	this := DeleteCustomerSourcedEntityLookupTableResponse{}
	return &this
}

// GetData returns the Data field value
func (o *DeleteCustomerSourcedEntityLookupTableResponse) GetData() DeleteContextActionResponseData {
	if o == nil {
		var ret DeleteContextActionResponseData
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *DeleteCustomerSourcedEntityLookupTableResponse) GetDataOk() (*DeleteContextActionResponseData, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *DeleteCustomerSourcedEntityLookupTableResponse) SetData(v DeleteContextActionResponseData) {
	o.Data = v
}

func (o DeleteCustomerSourcedEntityLookupTableResponse) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeleteCustomerSourcedEntityLookupTableResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *DeleteCustomerSourcedEntityLookupTableResponse) UnmarshalJSON(data []byte) (err error) {
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

	varDeleteCustomerSourcedEntityLookupTableResponse := _DeleteCustomerSourcedEntityLookupTableResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeleteCustomerSourcedEntityLookupTableResponse)

	if err != nil {
		return err
	}

	*o = DeleteCustomerSourcedEntityLookupTableResponse(varDeleteCustomerSourcedEntityLookupTableResponse)

	return err
}

type NullableDeleteCustomerSourcedEntityLookupTableResponse struct {
	value *DeleteCustomerSourcedEntityLookupTableResponse
	isSet bool
}

func (v NullableDeleteCustomerSourcedEntityLookupTableResponse) Get() *DeleteCustomerSourcedEntityLookupTableResponse {
	return v.value
}

func (v *NullableDeleteCustomerSourcedEntityLookupTableResponse) Set(val *DeleteCustomerSourcedEntityLookupTableResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableDeleteCustomerSourcedEntityLookupTableResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableDeleteCustomerSourcedEntityLookupTableResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeleteCustomerSourcedEntityLookupTableResponse(val *DeleteCustomerSourcedEntityLookupTableResponse) *NullableDeleteCustomerSourcedEntityLookupTableResponse {
	return &NullableDeleteCustomerSourcedEntityLookupTableResponse{value: val, isSet: true}
}

func (v NullableDeleteCustomerSourcedEntityLookupTableResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeleteCustomerSourcedEntityLookupTableResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


