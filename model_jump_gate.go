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

// checks if the JumpGate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JumpGate{}

// JumpGate 
type JumpGate struct {
	// The symbol of the waypoint.
	Symbol string `json:"symbol"`
	// All the gates that are connected to this waypoint.
	Connections []string `json:"connections"`
}

type _JumpGate JumpGate

// NewJumpGate instantiates a new JumpGate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJumpGate(symbol string, connections []string) *JumpGate {
	this := JumpGate{}
	this.Symbol = symbol
	this.Connections = connections
	return &this
}

// NewJumpGateWithDefaults instantiates a new JumpGate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJumpGateWithDefaults() *JumpGate {
	this := JumpGate{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *JumpGate) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *JumpGate) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *JumpGate) SetSymbol(v string) {
	o.Symbol = v
}

// GetConnections returns the Connections field value
func (o *JumpGate) GetConnections() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Connections
}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
func (o *JumpGate) GetConnectionsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Connections, true
}

// SetConnections sets field value
func (o *JumpGate) SetConnections(v []string) {
	o.Connections = v
}

func (o JumpGate) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JumpGate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	toSerialize["connections"] = o.Connections
	return toSerialize, nil
}

func (o *JumpGate) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
		"connections",
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

	varJumpGate := _JumpGate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJumpGate)

	if err != nil {
		return err
	}

	*o = JumpGate(varJumpGate)

	return err
}

type NullableJumpGate struct {
	value *JumpGate
	isSet bool
}

func (v NullableJumpGate) Get() *JumpGate {
	return v.value
}

func (v *NullableJumpGate) Set(val *JumpGate) {
	v.value = val
	v.isSet = true
}

func (v NullableJumpGate) IsSet() bool {
	return v.isSet
}

func (v *NullableJumpGate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJumpGate(val *JumpGate) *NullableJumpGate {
	return &NullableJumpGate{value: val, isSet: true}
}

func (v NullableJumpGate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJumpGate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


