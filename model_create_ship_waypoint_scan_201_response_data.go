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

// checks if the CreateShipWaypointScan201ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CreateShipWaypointScan201ResponseData{}

// CreateShipWaypointScan201ResponseData struct for CreateShipWaypointScan201ResponseData
type CreateShipWaypointScan201ResponseData struct {
	Cooldown Cooldown `json:"cooldown"`
	// List of scanned waypoints.
	Waypoints []ScannedWaypoint `json:"waypoints"`
}

type _CreateShipWaypointScan201ResponseData CreateShipWaypointScan201ResponseData

// NewCreateShipWaypointScan201ResponseData instantiates a new CreateShipWaypointScan201ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateShipWaypointScan201ResponseData(cooldown Cooldown, waypoints []ScannedWaypoint) *CreateShipWaypointScan201ResponseData {
	this := CreateShipWaypointScan201ResponseData{}
	this.Cooldown = cooldown
	this.Waypoints = waypoints
	return &this
}

// NewCreateShipWaypointScan201ResponseDataWithDefaults instantiates a new CreateShipWaypointScan201ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateShipWaypointScan201ResponseDataWithDefaults() *CreateShipWaypointScan201ResponseData {
	this := CreateShipWaypointScan201ResponseData{}
	return &this
}

// GetCooldown returns the Cooldown field value
func (o *CreateShipWaypointScan201ResponseData) GetCooldown() Cooldown {
	if o == nil {
		var ret Cooldown
		return ret
	}

	return o.Cooldown
}

// GetCooldownOk returns a tuple with the Cooldown field value
// and a boolean to check if the value has been set.
func (o *CreateShipWaypointScan201ResponseData) GetCooldownOk() (*Cooldown, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cooldown, true
}

// SetCooldown sets field value
func (o *CreateShipWaypointScan201ResponseData) SetCooldown(v Cooldown) {
	o.Cooldown = v
}

// GetWaypoints returns the Waypoints field value
func (o *CreateShipWaypointScan201ResponseData) GetWaypoints() []ScannedWaypoint {
	if o == nil {
		var ret []ScannedWaypoint
		return ret
	}

	return o.Waypoints
}

// GetWaypointsOk returns a tuple with the Waypoints field value
// and a boolean to check if the value has been set.
func (o *CreateShipWaypointScan201ResponseData) GetWaypointsOk() ([]ScannedWaypoint, bool) {
	if o == nil {
		return nil, false
	}
	return o.Waypoints, true
}

// SetWaypoints sets field value
func (o *CreateShipWaypointScan201ResponseData) SetWaypoints(v []ScannedWaypoint) {
	o.Waypoints = v
}

func (o CreateShipWaypointScan201ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CreateShipWaypointScan201ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cooldown"] = o.Cooldown
	toSerialize["waypoints"] = o.Waypoints
	return toSerialize, nil
}

func (o *CreateShipWaypointScan201ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"cooldown",
		"waypoints",
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

	varCreateShipWaypointScan201ResponseData := _CreateShipWaypointScan201ResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCreateShipWaypointScan201ResponseData)

	if err != nil {
		return err
	}

	*o = CreateShipWaypointScan201ResponseData(varCreateShipWaypointScan201ResponseData)

	return err
}

type NullableCreateShipWaypointScan201ResponseData struct {
	value *CreateShipWaypointScan201ResponseData
	isSet bool
}

func (v NullableCreateShipWaypointScan201ResponseData) Get() *CreateShipWaypointScan201ResponseData {
	return v.value
}

func (v *NullableCreateShipWaypointScan201ResponseData) Set(val *CreateShipWaypointScan201ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateShipWaypointScan201ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateShipWaypointScan201ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateShipWaypointScan201ResponseData(val *CreateShipWaypointScan201ResponseData) *NullableCreateShipWaypointScan201ResponseData {
	return &NullableCreateShipWaypointScan201ResponseData{value: val, isSet: true}
}

func (v NullableCreateShipWaypointScan201ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateShipWaypointScan201ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


