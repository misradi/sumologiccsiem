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

// checks if the GetRulesResponseDataObjectsInnerOneOfStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetRulesResponseDataObjectsInnerOneOfStatus{}

// GetRulesResponseDataObjectsInnerOneOfStatus struct for GetRulesResponseDataObjectsInnerOneOfStatus
type GetRulesResponseDataObjectsInnerOneOfStatus struct {
	Status string `json:"status"`
	Message *string `json:"message,omitempty"`
}

type _GetRulesResponseDataObjectsInnerOneOfStatus GetRulesResponseDataObjectsInnerOneOfStatus

// NewGetRulesResponseDataObjectsInnerOneOfStatus instantiates a new GetRulesResponseDataObjectsInnerOneOfStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetRulesResponseDataObjectsInnerOneOfStatus(status string) *GetRulesResponseDataObjectsInnerOneOfStatus {
	this := GetRulesResponseDataObjectsInnerOneOfStatus{}
	this.Status = status
	return &this
}

// NewGetRulesResponseDataObjectsInnerOneOfStatusWithDefaults instantiates a new GetRulesResponseDataObjectsInnerOneOfStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetRulesResponseDataObjectsInnerOneOfStatusWithDefaults() *GetRulesResponseDataObjectsInnerOneOfStatus {
	this := GetRulesResponseDataObjectsInnerOneOfStatus{}
	return &this
}

// GetStatus returns the Status field value
func (o *GetRulesResponseDataObjectsInnerOneOfStatus) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *GetRulesResponseDataObjectsInnerOneOfStatus) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *GetRulesResponseDataObjectsInnerOneOfStatus) SetStatus(v string) {
	o.Status = v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *GetRulesResponseDataObjectsInnerOneOfStatus) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GetRulesResponseDataObjectsInnerOneOfStatus) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *GetRulesResponseDataObjectsInnerOneOfStatus) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *GetRulesResponseDataObjectsInnerOneOfStatus) SetMessage(v string) {
	o.Message = &v
}

func (o GetRulesResponseDataObjectsInnerOneOfStatus) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetRulesResponseDataObjectsInnerOneOfStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["status"] = o.Status
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

func (o *GetRulesResponseDataObjectsInnerOneOfStatus) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"status",
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

	varGetRulesResponseDataObjectsInnerOneOfStatus := _GetRulesResponseDataObjectsInnerOneOfStatus{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetRulesResponseDataObjectsInnerOneOfStatus)

	if err != nil {
		return err
	}

	*o = GetRulesResponseDataObjectsInnerOneOfStatus(varGetRulesResponseDataObjectsInnerOneOfStatus)

	return err
}

type NullableGetRulesResponseDataObjectsInnerOneOfStatus struct {
	value *GetRulesResponseDataObjectsInnerOneOfStatus
	isSet bool
}

func (v NullableGetRulesResponseDataObjectsInnerOneOfStatus) Get() *GetRulesResponseDataObjectsInnerOneOfStatus {
	return v.value
}

func (v *NullableGetRulesResponseDataObjectsInnerOneOfStatus) Set(val *GetRulesResponseDataObjectsInnerOneOfStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableGetRulesResponseDataObjectsInnerOneOfStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableGetRulesResponseDataObjectsInnerOneOfStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetRulesResponseDataObjectsInnerOneOfStatus(val *GetRulesResponseDataObjectsInnerOneOfStatus) *NullableGetRulesResponseDataObjectsInnerOneOfStatus {
	return &NullableGetRulesResponseDataObjectsInnerOneOfStatus{value: val, isSet: true}
}

func (v NullableGetRulesResponseDataObjectsInnerOneOfStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetRulesResponseDataObjectsInnerOneOfStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


