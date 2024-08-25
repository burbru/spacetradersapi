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

// checks if the GetMarket200Response type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &GetMarket200Response{}

// GetMarket200Response struct for GetMarket200Response
type GetMarket200Response struct {
	Data Market `json:"data"`
}

type _GetMarket200Response GetMarket200Response

// NewGetMarket200Response instantiates a new GetMarket200Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGetMarket200Response(data Market) *GetMarket200Response {
	this := GetMarket200Response{}
	this.Data = data
	return &this
}

// NewGetMarket200ResponseWithDefaults instantiates a new GetMarket200Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGetMarket200ResponseWithDefaults() *GetMarket200Response {
	this := GetMarket200Response{}
	return &this
}

// GetData returns the Data field value
func (o *GetMarket200Response) GetData() Market {
	if o == nil {
		var ret Market
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *GetMarket200Response) GetDataOk() (*Market, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *GetMarket200Response) SetData(v Market) {
	o.Data = v
}

func (o GetMarket200Response) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o GetMarket200Response) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	return toSerialize, nil
}

func (o *GetMarket200Response) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"data",
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

	varGetMarket200Response := _GetMarket200Response{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varGetMarket200Response)

	if err != nil {
		return err
	}

	*o = GetMarket200Response(varGetMarket200Response)

	return err
}

type NullableGetMarket200Response struct {
	value *GetMarket200Response
	isSet bool
}

func (v NullableGetMarket200Response) Get() *GetMarket200Response {
	return v.value
}

func (v *NullableGetMarket200Response) Set(val *GetMarket200Response) {
	v.value = val
	v.isSet = true
}

func (v NullableGetMarket200Response) IsSet() bool {
	return v.isSet
}

func (v *NullableGetMarket200Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGetMarket200Response(val *GetMarket200Response) *NullableGetMarket200Response {
	return &NullableGetMarket200Response{value: val, isSet: true}
}

func (v NullableGetMarket200Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGetMarket200Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


