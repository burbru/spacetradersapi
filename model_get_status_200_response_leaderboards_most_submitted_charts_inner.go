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

// checks if the GetStatus200ResponseLeaderboardsMostSubmittedChartsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatus200ResponseLeaderboardsMostSubmittedChartsInner{}

// GetStatus200ResponseLeaderboardsMostSubmittedChartsInner struct for GetStatus200ResponseLeaderboardsMostSubmittedChartsInner
type GetStatus200ResponseLeaderboardsMostSubmittedChartsInner struct {
	// Symbol of the agent.
	AgentSymbol string `json:"agentSymbol"`
	// Amount of charts done by the agent.
	ChartCount int32 `json:"chartCount"`
}

type _GetStatus200ResponseLeaderboardsMostSubmittedChartsInner GetStatus200ResponseLeaderboardsMostSubmittedChartsInner

// NewGetStatus200ResponseLeaderboardsMostSubmittedChartsInner instantiates a new GetStatus200ResponseLeaderboardsMostSubmittedChartsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatus200ResponseLeaderboardsMostSubmittedChartsInner(agentSymbol string, chartCount int32) *GetStatus200ResponseLeaderboardsMostSubmittedChartsInner {
	this := GetStatus200ResponseLeaderboardsMostSubmittedChartsInner{}
	this.AgentSymbol = agentSymbol
	this.ChartCount = chartCount
	return &this
}

// NewGetStatus200ResponseLeaderboardsMostSubmittedChartsInnerWithDefaults instantiates a new GetStatus200ResponseLeaderboardsMostSubmittedChartsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatus200ResponseLeaderboardsMostSubmittedChartsInnerWithDefaults() *GetStatus200ResponseLeaderboardsMostSubmittedChartsInner {
	this := GetStatus200ResponseLeaderboardsMostSubmittedChartsInner{}
	return &this
}

// GetAgentSymbol returns the AgentSymbol field value
func (o *GetStatus200ResponseLeaderboardsMostSubmittedChartsInner) GetAgentSymbol() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AgentSymbol
}

// GetAgentSymbolOk returns a tuple with the AgentSymbol field value
// and a boolean to check if the value has been set.
func (o *GetStatus200ResponseLeaderboardsMostSubmittedChartsInner) GetAgentSymbolOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AgentSymbol, true
}

// SetAgentSymbol sets field value
func (o *GetStatus200ResponseLeaderboardsMostSubmittedChartsInner) SetAgentSymbol(v string) {
	o.AgentSymbol = v
}

// GetChartCount returns the ChartCount field value
func (o *GetStatus200ResponseLeaderboardsMostSubmittedChartsInner) GetChartCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.ChartCount
}

// GetChartCountOk returns a tuple with the ChartCount field value
// and a boolean to check if the value has been set.
func (o *GetStatus200ResponseLeaderboardsMostSubmittedChartsInner) GetChartCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ChartCount, true
}

// SetChartCount sets field value
func (o *GetStatus200ResponseLeaderboardsMostSubmittedChartsInner) SetChartCount(v int32) {
	o.ChartCount = v
}

func (o GetStatus200ResponseLeaderboardsMostSubmittedChartsInner) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatus200ResponseLeaderboardsMostSubmittedChartsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agentSymbol"] = o.AgentSymbol
	toSerialize["chartCount"] = o.ChartCount
	return toSerialize, nil
}

func (o *GetStatus200ResponseLeaderboardsMostSubmittedChartsInner) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"agentSymbol",
		"chartCount",
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

	varGetStatus200ResponseLeaderboardsMostSubmittedChartsInner := _GetStatus200ResponseLeaderboardsMostSubmittedChartsInner{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetStatus200ResponseLeaderboardsMostSubmittedChartsInner)

	if err != nil {
		return err
	}

	*o = GetStatus200ResponseLeaderboardsMostSubmittedChartsInner(varGetStatus200ResponseLeaderboardsMostSubmittedChartsInner)

	return err
}

type NullableGetStatus200ResponseLeaderboardsMostSubmittedChartsInner struct {
	value *GetStatus200ResponseLeaderboardsMostSubmittedChartsInner
	isSet bool
}

func (v NullableGetStatus200ResponseLeaderboardsMostSubmittedChartsInner) Get() *GetStatus200ResponseLeaderboardsMostSubmittedChartsInner {
	return v.value
}

func (v *NullableGetStatus200ResponseLeaderboardsMostSubmittedChartsInner) Set(val *GetStatus200ResponseLeaderboardsMostSubmittedChartsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatus200ResponseLeaderboardsMostSubmittedChartsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatus200ResponseLeaderboardsMostSubmittedChartsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatus200ResponseLeaderboardsMostSubmittedChartsInner(val *GetStatus200ResponseLeaderboardsMostSubmittedChartsInner) *NullableGetStatus200ResponseLeaderboardsMostSubmittedChartsInner {
	return &NullableGetStatus200ResponseLeaderboardsMostSubmittedChartsInner{value: val, isSet: true}
}

func (v NullableGetStatus200ResponseLeaderboardsMostSubmittedChartsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatus200ResponseLeaderboardsMostSubmittedChartsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


