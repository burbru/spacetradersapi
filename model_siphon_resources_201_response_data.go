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

// checks if the SiphonResources201ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SiphonResources201ResponseData{}

// SiphonResources201ResponseData struct for SiphonResources201ResponseData
type SiphonResources201ResponseData struct {
	Cooldown Cooldown `json:"cooldown"`
	Siphon Siphon `json:"siphon"`
	Cargo ShipCargo `json:"cargo"`
	Events []ExtractResources201ResponseDataEventsInner `json:"events"`
}

type _SiphonResources201ResponseData SiphonResources201ResponseData

// NewSiphonResources201ResponseData instantiates a new SiphonResources201ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSiphonResources201ResponseData(cooldown Cooldown, siphon Siphon, cargo ShipCargo, events []ExtractResources201ResponseDataEventsInner) *SiphonResources201ResponseData {
	this := SiphonResources201ResponseData{}
	this.Cooldown = cooldown
	this.Siphon = siphon
	this.Cargo = cargo
	this.Events = events
	return &this
}

// NewSiphonResources201ResponseDataWithDefaults instantiates a new SiphonResources201ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSiphonResources201ResponseDataWithDefaults() *SiphonResources201ResponseData {
	this := SiphonResources201ResponseData{}
	return &this
}

// GetCooldown returns the Cooldown field value
func (o *SiphonResources201ResponseData) GetCooldown() Cooldown {
	if o == nil {
		var ret Cooldown
		return ret
	}

	return o.Cooldown
}

// GetCooldownOk returns a tuple with the Cooldown field value
// and a boolean to check if the value has been set.
func (o *SiphonResources201ResponseData) GetCooldownOk() (*Cooldown, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cooldown, true
}

// SetCooldown sets field value
func (o *SiphonResources201ResponseData) SetCooldown(v Cooldown) {
	o.Cooldown = v
}

// GetSiphon returns the Siphon field value
func (o *SiphonResources201ResponseData) GetSiphon() Siphon {
	if o == nil {
		var ret Siphon
		return ret
	}

	return o.Siphon
}

// GetSiphonOk returns a tuple with the Siphon field value
// and a boolean to check if the value has been set.
func (o *SiphonResources201ResponseData) GetSiphonOk() (*Siphon, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Siphon, true
}

// SetSiphon sets field value
func (o *SiphonResources201ResponseData) SetSiphon(v Siphon) {
	o.Siphon = v
}

// GetCargo returns the Cargo field value
func (o *SiphonResources201ResponseData) GetCargo() ShipCargo {
	if o == nil {
		var ret ShipCargo
		return ret
	}

	return o.Cargo
}

// GetCargoOk returns a tuple with the Cargo field value
// and a boolean to check if the value has been set.
func (o *SiphonResources201ResponseData) GetCargoOk() (*ShipCargo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cargo, true
}

// SetCargo sets field value
func (o *SiphonResources201ResponseData) SetCargo(v ShipCargo) {
	o.Cargo = v
}

// GetEvents returns the Events field value
func (o *SiphonResources201ResponseData) GetEvents() []ExtractResources201ResponseDataEventsInner {
	if o == nil {
		var ret []ExtractResources201ResponseDataEventsInner
		return ret
	}

	return o.Events
}

// GetEventsOk returns a tuple with the Events field value
// and a boolean to check if the value has been set.
func (o *SiphonResources201ResponseData) GetEventsOk() ([]ExtractResources201ResponseDataEventsInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Events, true
}

// SetEvents sets field value
func (o *SiphonResources201ResponseData) SetEvents(v []ExtractResources201ResponseDataEventsInner) {
	o.Events = v
}

func (o SiphonResources201ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SiphonResources201ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cooldown"] = o.Cooldown
	toSerialize["siphon"] = o.Siphon
	toSerialize["cargo"] = o.Cargo
	toSerialize["events"] = o.Events
	return toSerialize, nil
}

func (o *SiphonResources201ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cooldown",
		"siphon",
		"cargo",
		"events",
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

	varSiphonResources201ResponseData := _SiphonResources201ResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSiphonResources201ResponseData)

	if err != nil {
		return err
	}

	*o = SiphonResources201ResponseData(varSiphonResources201ResponseData)

	return err
}

type NullableSiphonResources201ResponseData struct {
	value *SiphonResources201ResponseData
	isSet bool
}

func (v NullableSiphonResources201ResponseData) Get() *SiphonResources201ResponseData {
	return v.value
}

func (v *NullableSiphonResources201ResponseData) Set(val *SiphonResources201ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableSiphonResources201ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableSiphonResources201ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSiphonResources201ResponseData(val *SiphonResources201ResponseData) *NullableSiphonResources201ResponseData {
	return &NullableSiphonResources201ResponseData{value: val, isSet: true}
}

func (v NullableSiphonResources201ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSiphonResources201ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


