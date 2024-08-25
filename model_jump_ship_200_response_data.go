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

// checks if the JumpShip200ResponseData type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &JumpShip200ResponseData{}

// JumpShip200ResponseData struct for JumpShip200ResponseData
type JumpShip200ResponseData struct {
	Nav ShipNav `json:"nav"`
	Cooldown Cooldown `json:"cooldown"`
	Transaction MarketTransaction `json:"transaction"`
	Agent Agent `json:"agent"`
}

type _JumpShip200ResponseData JumpShip200ResponseData

// NewJumpShip200ResponseData instantiates a new JumpShip200ResponseData object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJumpShip200ResponseData(nav ShipNav, cooldown Cooldown, transaction MarketTransaction, agent Agent) *JumpShip200ResponseData {
	this := JumpShip200ResponseData{}
	this.Nav = nav
	this.Cooldown = cooldown
	this.Transaction = transaction
	this.Agent = agent
	return &this
}

// NewJumpShip200ResponseDataWithDefaults instantiates a new JumpShip200ResponseData object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJumpShip200ResponseDataWithDefaults() *JumpShip200ResponseData {
	this := JumpShip200ResponseData{}
	return &this
}

// GetNav returns the Nav field value
func (o *JumpShip200ResponseData) GetNav() ShipNav {
	if o == nil {
		var ret ShipNav
		return ret
	}

	return o.Nav
}

// GetNavOk returns a tuple with the Nav field value
// and a boolean to check if the value has been set.
func (o *JumpShip200ResponseData) GetNavOk() (*ShipNav, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Nav, true
}

// SetNav sets field value
func (o *JumpShip200ResponseData) SetNav(v ShipNav) {
	o.Nav = v
}

// GetCooldown returns the Cooldown field value
func (o *JumpShip200ResponseData) GetCooldown() Cooldown {
	if o == nil {
		var ret Cooldown
		return ret
	}

	return o.Cooldown
}

// GetCooldownOk returns a tuple with the Cooldown field value
// and a boolean to check if the value has been set.
func (o *JumpShip200ResponseData) GetCooldownOk() (*Cooldown, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Cooldown, true
}

// SetCooldown sets field value
func (o *JumpShip200ResponseData) SetCooldown(v Cooldown) {
	o.Cooldown = v
}

// GetTransaction returns the Transaction field value
func (o *JumpShip200ResponseData) GetTransaction() MarketTransaction {
	if o == nil {
		var ret MarketTransaction
		return ret
	}

	return o.Transaction
}

// GetTransactionOk returns a tuple with the Transaction field value
// and a boolean to check if the value has been set.
func (o *JumpShip200ResponseData) GetTransactionOk() (*MarketTransaction, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Transaction, true
}

// SetTransaction sets field value
func (o *JumpShip200ResponseData) SetTransaction(v MarketTransaction) {
	o.Transaction = v
}

// GetAgent returns the Agent field value
func (o *JumpShip200ResponseData) GetAgent() Agent {
	if o == nil {
		var ret Agent
		return ret
	}

	return o.Agent
}

// GetAgentOk returns a tuple with the Agent field value
// and a boolean to check if the value has been set.
func (o *JumpShip200ResponseData) GetAgentOk() (*Agent, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Agent, true
}

// SetAgent sets field value
func (o *JumpShip200ResponseData) SetAgent(v Agent) {
	o.Agent = v
}

func (o JumpShip200ResponseData) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o JumpShip200ResponseData) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nav"] = o.Nav
	toSerialize["cooldown"] = o.Cooldown
	toSerialize["transaction"] = o.Transaction
	toSerialize["agent"] = o.Agent
	return toSerialize, nil
}

func (o *JumpShip200ResponseData) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"nav",
		"cooldown",
		"transaction",
		"agent",
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

	varJumpShip200ResponseData := _JumpShip200ResponseData{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varJumpShip200ResponseData)

	if err != nil {
		return err
	}

	*o = JumpShip200ResponseData(varJumpShip200ResponseData)

	return err
}

type NullableJumpShip200ResponseData struct {
	value *JumpShip200ResponseData
	isSet bool
}

func (v NullableJumpShip200ResponseData) Get() *JumpShip200ResponseData {
	return v.value
}

func (v *NullableJumpShip200ResponseData) Set(val *JumpShip200ResponseData) {
	v.value = val
	v.isSet = true
}

func (v NullableJumpShip200ResponseData) IsSet() bool {
	return v.isSet
}

func (v *NullableJumpShip200ResponseData) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJumpShip200ResponseData(val *JumpShip200ResponseData) *NullableJumpShip200ResponseData {
	return &NullableJumpShip200ResponseData{value: val, isSet: true}
}

func (v NullableJumpShip200ResponseData) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJumpShip200ResponseData) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


