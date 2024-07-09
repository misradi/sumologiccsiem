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

// checks if the CreateSuppressListRequestBodyFields type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateSuppressListRequestBodyFields{}

// CreateSuppressListRequestBodyFields struct for CreateSuppressListRequestBodyFields
type CreateSuppressListRequestBodyFields struct {
	Name string `json:"name"`
	Description *string `json:"description,omitempty"`
	TargetColumn string `json:"targetColumn"`
	// The default time-to-live (in seconds) for new Items added to this List. This default is only used to default the field in the UI, and is not used at all when new Items are added via the API.
	DefaultTtl *int32 `json:"defaultTtl,omitempty"`
	Active *bool `json:"active,omitempty"`
}

type _CreateSuppressListRequestBodyFields CreateSuppressListRequestBodyFields

// NewCreateSuppressListRequestBodyFields instantiates a new CreateSuppressListRequestBodyFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateSuppressListRequestBodyFields(name string, targetColumn string) *CreateSuppressListRequestBodyFields {
	this := CreateSuppressListRequestBodyFields{}
	this.Name = name
	this.TargetColumn = targetColumn
	return &this
}

// NewCreateSuppressListRequestBodyFieldsWithDefaults instantiates a new CreateSuppressListRequestBodyFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateSuppressListRequestBodyFieldsWithDefaults() *CreateSuppressListRequestBodyFields {
	this := CreateSuppressListRequestBodyFields{}
	return &this
}

// GetName returns the Name field value
func (o *CreateSuppressListRequestBodyFields) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *CreateSuppressListRequestBodyFields) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *CreateSuppressListRequestBodyFields) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateSuppressListRequestBodyFields) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSuppressListRequestBodyFields) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateSuppressListRequestBodyFields) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateSuppressListRequestBodyFields) SetDescription(v string) {
	o.Description = &v
}

// GetTargetColumn returns the TargetColumn field value
func (o *CreateSuppressListRequestBodyFields) GetTargetColumn() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetColumn
}

// GetTargetColumnOk returns a tuple with the TargetColumn field value
// and a boolean to check if the value has been set.
func (o *CreateSuppressListRequestBodyFields) GetTargetColumnOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetColumn, true
}

// SetTargetColumn sets field value
func (o *CreateSuppressListRequestBodyFields) SetTargetColumn(v string) {
	o.TargetColumn = v
}

// GetDefaultTtl returns the DefaultTtl field value if set, zero value otherwise.
func (o *CreateSuppressListRequestBodyFields) GetDefaultTtl() int32 {
	if o == nil || IsNil(o.DefaultTtl) {
		var ret int32
		return ret
	}
	return *o.DefaultTtl
}

// GetDefaultTtlOk returns a tuple with the DefaultTtl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSuppressListRequestBodyFields) GetDefaultTtlOk() (*int32, bool) {
	if o == nil || IsNil(o.DefaultTtl) {
		return nil, false
	}
	return o.DefaultTtl, true
}

// HasDefaultTtl returns a boolean if a field has been set.
func (o *CreateSuppressListRequestBodyFields) HasDefaultTtl() bool {
	if o != nil && !IsNil(o.DefaultTtl) {
		return true
	}

	return false
}

// SetDefaultTtl gets a reference to the given int32 and assigns it to the DefaultTtl field.
func (o *CreateSuppressListRequestBodyFields) SetDefaultTtl(v int32) {
	o.DefaultTtl = &v
}

// GetActive returns the Active field value if set, zero value otherwise.
func (o *CreateSuppressListRequestBodyFields) GetActive() bool {
	if o == nil || IsNil(o.Active) {
		var ret bool
		return ret
	}
	return *o.Active
}

// GetActiveOk returns a tuple with the Active field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateSuppressListRequestBodyFields) GetActiveOk() (*bool, bool) {
	if o == nil || IsNil(o.Active) {
		return nil, false
	}
	return o.Active, true
}

// HasActive returns a boolean if a field has been set.
func (o *CreateSuppressListRequestBodyFields) HasActive() bool {
	if o != nil && !IsNil(o.Active) {
		return true
	}

	return false
}

// SetActive gets a reference to the given bool and assigns it to the Active field.
func (o *CreateSuppressListRequestBodyFields) SetActive(v bool) {
	o.Active = &v
}

func (o CreateSuppressListRequestBodyFields) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateSuppressListRequestBodyFields) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["targetColumn"] = o.TargetColumn
	if !IsNil(o.DefaultTtl) {
		toSerialize["defaultTtl"] = o.DefaultTtl
	}
	if !IsNil(o.Active) {
		toSerialize["active"] = o.Active
	}
	return toSerialize, nil
}

func (o *CreateSuppressListRequestBodyFields) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"targetColumn",
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

	varCreateSuppressListRequestBodyFields := _CreateSuppressListRequestBodyFields{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateSuppressListRequestBodyFields)

	if err != nil {
		return err
	}

	*o = CreateSuppressListRequestBodyFields(varCreateSuppressListRequestBodyFields)

	return err
}

type NullableCreateSuppressListRequestBodyFields struct {
	value *CreateSuppressListRequestBodyFields
	isSet bool
}

func (v NullableCreateSuppressListRequestBodyFields) Get() *CreateSuppressListRequestBodyFields {
	return v.value
}

func (v *NullableCreateSuppressListRequestBodyFields) Set(val *CreateSuppressListRequestBodyFields) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateSuppressListRequestBodyFields) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateSuppressListRequestBodyFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateSuppressListRequestBodyFields(val *CreateSuppressListRequestBodyFields) *NullableCreateSuppressListRequestBodyFields {
	return &NullableCreateSuppressListRequestBodyFields{value: val, isSet: true}
}

func (v NullableCreateSuppressListRequestBodyFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateSuppressListRequestBodyFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


