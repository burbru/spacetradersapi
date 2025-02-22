/*
SpaceTraders API

SpaceTraders is an open-universe game and learning platform that offers a set of HTTP endpoints to control a fleet of ships and explore a multiplayer universe.  The API is documented using [OpenAPI](https://github.com/SpaceTradersAPI/api-docs). You can send your first request right here in your browser to check the status of the game server.  ```json http {   \"method\": \"GET\",   \"url\": \"https://api.spacetraders.io/v2\", } ```  Unlike a traditional game, SpaceTraders does not have a first-party client or app to play the game. Instead, you can use the API to build your own client, write a script to automate your ships, or try an app built by the community.  We have a [Discord channel](https://discord.com/invite/jh6zurdWk5) where you can share your projects, ask questions, and get help from other players.   

API version: 2.0.0
Contact: joel@spacetraders.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package spacetradersapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the Jettison200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Jettison200ResponseData{}

// Jettison200ResponseData struct for Jettison200ResponseData
type Jettison200ResponseData struct {
	Cargo ShipCargo `json:"cargo"`
}

type _Jettison200ResponseData Jettison200ResponseData

// NewJettison200ResponseData instantiates a new Jettison200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJettison200ResponseData(cargo ShipCargo) *Jettison200ResponseData {
	this := Jettison200ResponseData{}
	this.Cargo = cargo
	return &this
}

// NewJettison200ResponseDataWithDefaults instantiates a new Jettison200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJettison200ResponseDataWithDefaults() *Jettison200ResponseData {
	this := Jettison200ResponseData{}
	return &this
}

// GetCargo returns the Cargo field value
func (o *Jettison200ResponseData) GetCargo() ShipCargo {
	if o == nil {
		var ret ShipCargo
		return ret
	}

	return o.Cargo
}

// GetCargoOk returns a tuple with the Cargo field value
// and a boolean to check if the value has been set.
func (o *Jettison200ResponseData) GetCargoOk() (*ShipCargo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cargo, true
}

// SetCargo sets field value
func (o *Jettison200ResponseData) SetCargo(v ShipCargo) {
	o.Cargo = v
}

func (o Jettison200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Jettison200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cargo"] = o.Cargo
	return toSerialize, nil
}

func (o *Jettison200ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cargo",
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

	varJettison200ResponseData := _Jettison200ResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJettison200ResponseData)

	if err != nil {
		return err
	}

	*o = Jettison200ResponseData(varJettison200ResponseData)

	return err
}

type NullableJettison200ResponseData struct {
	value *Jettison200ResponseData
	isSet bool
}

func (v NullableJettison200ResponseData) Get() *Jettison200ResponseData {
	return v.value
}

func (v *NullableJettison200ResponseData) Set(val *Jettison200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableJettison200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableJettison200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJettison200ResponseData(val *Jettison200ResponseData) *NullableJettison200ResponseData {
	return &NullableJettison200ResponseData{value: val, isSet: true}
}

func (v NullableJettison200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJettison200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


