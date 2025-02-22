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

// checks if the GetStatus200ResponseStats type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetStatus200ResponseStats{}

// GetStatus200ResponseStats struct for GetStatus200ResponseStats
type GetStatus200ResponseStats struct {
	// Number of registered agents in the game.
	Agents int32 `json:"agents"`
	// Total number of ships in the game.
	Ships int32 `json:"ships"`
	// Total number of systems in the game.
	Systems int32 `json:"systems"`
	// Total number of waypoints in the game.
	Waypoints int32 `json:"waypoints"`
}

type _GetStatus200ResponseStats GetStatus200ResponseStats

// NewGetStatus200ResponseStats instantiates a new GetStatus200ResponseStats object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetStatus200ResponseStats(agents int32, ships int32, systems int32, waypoints int32) *GetStatus200ResponseStats {
	this := GetStatus200ResponseStats{}
	this.Agents = agents
	this.Ships = ships
	this.Systems = systems
	this.Waypoints = waypoints
	return &this
}

// NewGetStatus200ResponseStatsWithDefaults instantiates a new GetStatus200ResponseStats object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetStatus200ResponseStatsWithDefaults() *GetStatus200ResponseStats {
	this := GetStatus200ResponseStats{}
	return &this
}

// GetAgents returns the Agents field value
func (o *GetStatus200ResponseStats) GetAgents() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Agents
}

// GetAgentsOk returns a tuple with the Agents field value
// and a boolean to check if the value has been set.
func (o *GetStatus200ResponseStats) GetAgentsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Agents, true
}

// SetAgents sets field value
func (o *GetStatus200ResponseStats) SetAgents(v int32) {
	o.Agents = v
}

// GetShips returns the Ships field value
func (o *GetStatus200ResponseStats) GetShips() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Ships
}

// GetShipsOk returns a tuple with the Ships field value
// and a boolean to check if the value has been set.
func (o *GetStatus200ResponseStats) GetShipsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Ships, true
}

// SetShips sets field value
func (o *GetStatus200ResponseStats) SetShips(v int32) {
	o.Ships = v
}

// GetSystems returns the Systems field value
func (o *GetStatus200ResponseStats) GetSystems() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Systems
}

// GetSystemsOk returns a tuple with the Systems field value
// and a boolean to check if the value has been set.
func (o *GetStatus200ResponseStats) GetSystemsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Systems, true
}

// SetSystems sets field value
func (o *GetStatus200ResponseStats) SetSystems(v int32) {
	o.Systems = v
}

// GetWaypoints returns the Waypoints field value
func (o *GetStatus200ResponseStats) GetWaypoints() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Waypoints
}

// GetWaypointsOk returns a tuple with the Waypoints field value
// and a boolean to check if the value has been set.
func (o *GetStatus200ResponseStats) GetWaypointsOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Waypoints, true
}

// SetWaypoints sets field value
func (o *GetStatus200ResponseStats) SetWaypoints(v int32) {
	o.Waypoints = v
}

func (o GetStatus200ResponseStats) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetStatus200ResponseStats) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["agents"] = o.Agents
	toSerialize["ships"] = o.Ships
	toSerialize["systems"] = o.Systems
	toSerialize["waypoints"] = o.Waypoints
	return toSerialize, nil
}

func (o *GetStatus200ResponseStats) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"agents",
		"ships",
		"systems",
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

	varGetStatus200ResponseStats := _GetStatus200ResponseStats{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetStatus200ResponseStats)

	if err != nil {
		return err
	}

	*o = GetStatus200ResponseStats(varGetStatus200ResponseStats)

	return err
}

type NullableGetStatus200ResponseStats struct {
	value *GetStatus200ResponseStats
	isSet bool
}

func (v NullableGetStatus200ResponseStats) Get() *GetStatus200ResponseStats {
	return v.value
}

func (v *NullableGetStatus200ResponseStats) Set(val *GetStatus200ResponseStats) {
	v.value = val
	v.isSet = true
}

func (v NullableGetStatus200ResponseStats) IsSet() bool {
	return v.isSet
}

func (v *NullableGetStatus200ResponseStats) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetStatus200ResponseStats(val *GetStatus200ResponseStats) *NullableGetStatus200ResponseStats {
	return &NullableGetStatus200ResponseStats{value: val, isSet: true}
}

func (v NullableGetStatus200ResponseStats) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetStatus200ResponseStats) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


