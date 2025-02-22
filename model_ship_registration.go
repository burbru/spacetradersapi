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

// checks if the ShipRegistration type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ShipRegistration{}

// ShipRegistration The public registration information of the ship
type ShipRegistration struct {
	// The agent's registered name of the ship
	Name string `json:"name"`
	// The symbol of the faction the ship is registered with
	FactionSymbol string `json:"factionSymbol"`
	Role ShipRole `json:"role"`
}

type _ShipRegistration ShipRegistration

// NewShipRegistration instantiates a new ShipRegistration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewShipRegistration(name string, factionSymbol string, role ShipRole) *ShipRegistration {
	this := ShipRegistration{}
	this.Name = name
	this.FactionSymbol = factionSymbol
	this.Role = role
	return &this
}

// NewShipRegistrationWithDefaults instantiates a new ShipRegistration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewShipRegistrationWithDefaults() *ShipRegistration {
	this := ShipRegistration{}
	return &this
}

// GetName returns the Name field value
func (o *ShipRegistration) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ShipRegistration) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ShipRegistration) SetName(v string) {
	o.Name = v
}

// GetFactionSymbol returns the FactionSymbol field value
func (o *ShipRegistration) GetFactionSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FactionSymbol
}

// GetFactionSymbolOk returns a tuple with the FactionSymbol field value
// and a boolean to check if the value has been set.
func (o *ShipRegistration) GetFactionSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FactionSymbol, true
}

// SetFactionSymbol sets field value
func (o *ShipRegistration) SetFactionSymbol(v string) {
	o.FactionSymbol = v
}

// GetRole returns the Role field value
func (o *ShipRegistration) GetRole() ShipRole {
	if o == nil {
		var ret ShipRole
		return ret
	}

	return o.Role
}

// GetRoleOk returns a tuple with the Role field value
// and a boolean to check if the value has been set.
func (o *ShipRegistration) GetRoleOk() (*ShipRole, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Role, true
}

// SetRole sets field value
func (o *ShipRegistration) SetRole(v ShipRole) {
	o.Role = v
}

func (o ShipRegistration) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ShipRegistration) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["factionSymbol"] = o.FactionSymbol
	toSerialize["role"] = o.Role
	return toSerialize, nil
}

func (o *ShipRegistration) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"factionSymbol",
		"role",
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

	varShipRegistration := _ShipRegistration{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varShipRegistration)

	if err != nil {
		return err
	}

	*o = ShipRegistration(varShipRegistration)

	return err
}

type NullableShipRegistration struct {
	value *ShipRegistration
	isSet bool
}

func (v NullableShipRegistration) Get() *ShipRegistration {
	return v.value
}

func (v *NullableShipRegistration) Set(val *ShipRegistration) {
	v.value = val
	v.isSet = true
}

func (v NullableShipRegistration) IsSet() bool {
	return v.isSet
}

func (v *NullableShipRegistration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableShipRegistration(val *ShipRegistration) *NullableShipRegistration {
	return &NullableShipRegistration{value: val, isSet: true}
}

func (v NullableShipRegistration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableShipRegistration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


