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

// checks if the ScannedShipEngine type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScannedShipEngine{}

// ScannedShipEngine The engine of the ship.
type ScannedShipEngine struct {
	// The symbol of the engine.
	Symbol string `json:"symbol"`
}

type _ScannedShipEngine ScannedShipEngine

// NewScannedShipEngine instantiates a new ScannedShipEngine object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScannedShipEngine(symbol string) *ScannedShipEngine {
	this := ScannedShipEngine{}
	this.Symbol = symbol
	return &this
}

// NewScannedShipEngineWithDefaults instantiates a new ScannedShipEngine object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScannedShipEngineWithDefaults() *ScannedShipEngine {
	this := ScannedShipEngine{}
	return &this
}

// GetSymbol returns the Symbol field value
func (o *ScannedShipEngine) GetSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Symbol
}

// GetSymbolOk returns a tuple with the Symbol field value
// and a boolean to check if the value has been set.
func (o *ScannedShipEngine) GetSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Symbol, true
}

// SetSymbol sets field value
func (o *ScannedShipEngine) SetSymbol(v string) {
	o.Symbol = v
}

func (o ScannedShipEngine) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScannedShipEngine) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["symbol"] = o.Symbol
	return toSerialize, nil
}

func (o *ScannedShipEngine) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"symbol",
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

	varScannedShipEngine := _ScannedShipEngine{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varScannedShipEngine)

	if err != nil {
		return err
	}

	*o = ScannedShipEngine(varScannedShipEngine)

	return err
}

type NullableScannedShipEngine struct {
	value *ScannedShipEngine
	isSet bool
}

func (v NullableScannedShipEngine) Get() *ScannedShipEngine {
	return v.value
}

func (v *NullableScannedShipEngine) Set(val *ScannedShipEngine) {
	v.value = val
	v.isSet = true
}

func (v NullableScannedShipEngine) IsSet() bool {
	return v.isSet
}

func (v *NullableScannedShipEngine) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScannedShipEngine(val *ScannedShipEngine) *NullableScannedShipEngine {
	return &NullableScannedShipEngine{value: val, isSet: true}
}

func (v NullableScannedShipEngine) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScannedShipEngine) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


