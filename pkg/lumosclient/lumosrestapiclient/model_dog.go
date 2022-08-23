/*
Sample API

Optional multiline or single-line description in [CommonMark](http://commonmark.org/help/) or HTML.

API version: 0.1.9
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// Dog struct for Dog
type Dog struct {
	// name of the pet
	Name string `json:"name"`
	Type PetTypes `json:"type"`
	// age of the pet
	Age int32 `json:"age"`
	Bark bool `json:"bark"`
}

// NewDog instantiates a new Dog object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDog(name string, type_ PetTypes, age int32, bark bool) *Dog {
	this := Dog{}
	this.Name = name
	this.Type = type_
	this.Age = age
	this.Bark = bark
	return &this
}

// NewDogWithDefaults instantiates a new Dog object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDogWithDefaults() *Dog {
	this := Dog{}
	return &this
}

// GetName returns the Name field value
func (o *Dog) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Dog) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Dog) SetName(v string) {
	o.Name = v
}

// GetType returns the Type field value
func (o *Dog) GetType() PetTypes {
	if o == nil {
		var ret PetTypes
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Dog) GetTypeOk() (*PetTypes, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Dog) SetType(v PetTypes) {
	o.Type = v
}

// GetAge returns the Age field value
func (o *Dog) GetAge() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Age
}

// GetAgeOk returns a tuple with the Age field value
// and a boolean to check if the value has been set.
func (o *Dog) GetAgeOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Age, true
}

// SetAge sets field value
func (o *Dog) SetAge(v int32) {
	o.Age = v
}

// GetBark returns the Bark field value
func (o *Dog) GetBark() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Bark
}

// GetBarkOk returns a tuple with the Bark field value
// and a boolean to check if the value has been set.
func (o *Dog) GetBarkOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Bark, true
}

// SetBark sets field value
func (o *Dog) SetBark(v bool) {
	o.Bark = v
}

func (o Dog) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["age"] = o.Age
	}
	if true {
		toSerialize["bark"] = o.Bark
	}
	return json.Marshal(toSerialize)
}

type NullableDog struct {
	value *Dog
	isSet bool
}

func (v NullableDog) Get() *Dog {
	return v.value
}

func (v *NullableDog) Set(val *Dog) {
	v.value = val
	v.isSet = true
}

func (v NullableDog) IsSet() bool {
	return v.isSet
}

func (v *NullableDog) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDog(val *Dog) *NullableDog {
	return &NullableDog{value: val, isSet: true}
}

func (v NullableDog) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDog) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

