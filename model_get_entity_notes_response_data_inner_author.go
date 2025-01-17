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

// checks if the GetEntityNotesResponseDataInnerAuthor type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetEntityNotesResponseDataInnerAuthor{}

// GetEntityNotesResponseDataInnerAuthor struct for GetEntityNotesResponseDataInnerAuthor
type GetEntityNotesResponseDataInnerAuthor struct {
	Username string `json:"username"`
}

type _GetEntityNotesResponseDataInnerAuthor GetEntityNotesResponseDataInnerAuthor

// NewGetEntityNotesResponseDataInnerAuthor instantiates a new GetEntityNotesResponseDataInnerAuthor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetEntityNotesResponseDataInnerAuthor(username string) *GetEntityNotesResponseDataInnerAuthor {
	this := GetEntityNotesResponseDataInnerAuthor{}
	this.Username = username
	return &this
}

// NewGetEntityNotesResponseDataInnerAuthorWithDefaults instantiates a new GetEntityNotesResponseDataInnerAuthor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetEntityNotesResponseDataInnerAuthorWithDefaults() *GetEntityNotesResponseDataInnerAuthor {
	this := GetEntityNotesResponseDataInnerAuthor{}
	return &this
}

// GetUsername returns the Username field value
func (o *GetEntityNotesResponseDataInnerAuthor) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *GetEntityNotesResponseDataInnerAuthor) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *GetEntityNotesResponseDataInnerAuthor) SetUsername(v string) {
	o.Username = v
}

func (o GetEntityNotesResponseDataInnerAuthor) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetEntityNotesResponseDataInnerAuthor) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

func (o *GetEntityNotesResponseDataInnerAuthor) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"username",
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

	varGetEntityNotesResponseDataInnerAuthor := _GetEntityNotesResponseDataInnerAuthor{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetEntityNotesResponseDataInnerAuthor)

	if err != nil {
		return err
	}

	*o = GetEntityNotesResponseDataInnerAuthor(varGetEntityNotesResponseDataInnerAuthor)

	return err
}

type NullableGetEntityNotesResponseDataInnerAuthor struct {
	value *GetEntityNotesResponseDataInnerAuthor
	isSet bool
}

func (v NullableGetEntityNotesResponseDataInnerAuthor) Get() *GetEntityNotesResponseDataInnerAuthor {
	return v.value
}

func (v *NullableGetEntityNotesResponseDataInnerAuthor) Set(val *GetEntityNotesResponseDataInnerAuthor) {
	v.value = val
	v.isSet = true
}

func (v NullableGetEntityNotesResponseDataInnerAuthor) IsSet() bool {
	return v.isSet
}

func (v *NullableGetEntityNotesResponseDataInnerAuthor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetEntityNotesResponseDataInnerAuthor(val *GetEntityNotesResponseDataInnerAuthor) *NullableGetEntityNotesResponseDataInnerAuthor {
	return &NullableGetEntityNotesResponseDataInnerAuthor{value: val, isSet: true}
}

func (v NullableGetEntityNotesResponseDataInnerAuthor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetEntityNotesResponseDataInnerAuthor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


