/*
Okta Admin Management

Allows customers to easily access the Okta Management APIs

Copyright 2018 - Present Okta, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

API version: 2024.06.1
Contact: devex-public@okta.com
*/

package okta

import (
	"encoding/json"
	"reflect"
	"strings"
)

// MfaEnrollPolicyRule struct for MfaEnrollPolicyRule
type MfaEnrollPolicyRule struct {
	PolicyRule
	// The enrollment actions for the MFA enrollment policy rule.
	Actions map[string]interface{} `json:"actions,omitempty"`
	Conditions *PolicyRuleConditions `json:"conditions,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MfaEnrollPolicyRule MfaEnrollPolicyRule

// NewMfaEnrollPolicyRule instantiates a new MfaEnrollPolicyRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMfaEnrollPolicyRule() *MfaEnrollPolicyRule {
	this := MfaEnrollPolicyRule{}
	var system bool = false
	this.System = &system
	return &this
}

// NewMfaEnrollPolicyRuleWithDefaults instantiates a new MfaEnrollPolicyRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMfaEnrollPolicyRuleWithDefaults() *MfaEnrollPolicyRule {
	this := MfaEnrollPolicyRule{}
	return &this
}

// GetActions returns the Actions field value if set, zero value otherwise.
func (o *MfaEnrollPolicyRule) GetActions() map[string]interface{} {
	if o == nil || o.Actions == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.Actions
}

// GetActionsOk returns a tuple with the Actions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaEnrollPolicyRule) GetActionsOk() (map[string]interface{}, bool) {
	if o == nil || o.Actions == nil {
		return nil, false
	}
	return o.Actions, true
}

// HasActions returns a boolean if a field has been set.
func (o *MfaEnrollPolicyRule) HasActions() bool {
	if o != nil && o.Actions != nil {
		return true
	}

	return false
}

// SetActions gets a reference to the given map[string]interface{} and assigns it to the Actions field.
func (o *MfaEnrollPolicyRule) SetActions(v map[string]interface{}) {
	o.Actions = v
}

// GetConditions returns the Conditions field value if set, zero value otherwise.
func (o *MfaEnrollPolicyRule) GetConditions() PolicyRuleConditions {
	if o == nil || o.Conditions == nil {
		var ret PolicyRuleConditions
		return ret
	}
	return *o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaEnrollPolicyRule) GetConditionsOk() (*PolicyRuleConditions, bool) {
	if o == nil || o.Conditions == nil {
		return nil, false
	}
	return o.Conditions, true
}

// HasConditions returns a boolean if a field has been set.
func (o *MfaEnrollPolicyRule) HasConditions() bool {
	if o != nil && o.Conditions != nil {
		return true
	}

	return false
}

// SetConditions gets a reference to the given PolicyRuleConditions and assigns it to the Conditions field.
func (o *MfaEnrollPolicyRule) SetConditions(v PolicyRuleConditions) {
	o.Conditions = &v
}

func (o MfaEnrollPolicyRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	serializedPolicyRule, errPolicyRule := json.Marshal(o.PolicyRule)
	if errPolicyRule != nil {
		return []byte{}, errPolicyRule
	}
	errPolicyRule = json.Unmarshal([]byte(serializedPolicyRule), &toSerialize)
	if errPolicyRule != nil {
		return []byte{}, errPolicyRule
	}
	if o.Actions != nil {
		toSerialize["actions"] = o.Actions
	}
	if o.Conditions != nil {
		toSerialize["conditions"] = o.Conditions
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MfaEnrollPolicyRule) UnmarshalJSON(bytes []byte) (err error) {
	type MfaEnrollPolicyRuleWithoutEmbeddedStruct struct {
		// The enrollment actions for the MFA enrollment policy rule.
		Actions map[string]interface{} `json:"actions,omitempty"`
		Conditions *PolicyRuleConditions `json:"conditions,omitempty"`
	}

	varMfaEnrollPolicyRuleWithoutEmbeddedStruct := MfaEnrollPolicyRuleWithoutEmbeddedStruct{}

	err = json.Unmarshal(bytes, &varMfaEnrollPolicyRuleWithoutEmbeddedStruct)
	if err == nil {
		varMfaEnrollPolicyRule := _MfaEnrollPolicyRule{}
		varMfaEnrollPolicyRule.Actions = varMfaEnrollPolicyRuleWithoutEmbeddedStruct.Actions
		varMfaEnrollPolicyRule.Conditions = varMfaEnrollPolicyRuleWithoutEmbeddedStruct.Conditions
		*o = MfaEnrollPolicyRule(varMfaEnrollPolicyRule)
	} else {
		return err
	}

	varMfaEnrollPolicyRule := _MfaEnrollPolicyRule{}

	err = json.Unmarshal(bytes, &varMfaEnrollPolicyRule)
	if err == nil {
		o.PolicyRule = varMfaEnrollPolicyRule.PolicyRule
	} else {
		return err
	}

	additionalProperties := make(map[string]interface{})

	err = json.Unmarshal(bytes, &additionalProperties)
	if err == nil {
		delete(additionalProperties, "actions")
		delete(additionalProperties, "conditions")

		// remove fields from embedded structs
		reflectPolicyRule := reflect.ValueOf(o.PolicyRule)
		for i := 0; i < reflectPolicyRule.Type().NumField(); i++ {
			t := reflectPolicyRule.Type().Field(i)

			if jsonTag := t.Tag.Get("json"); jsonTag != "" {
				fieldName := ""
				if commaIdx := strings.Index(jsonTag, ","); commaIdx > 0 {
					fieldName = jsonTag[:commaIdx]
				} else {
					fieldName = jsonTag
				}
				if fieldName != "AdditionalProperties" {
					delete(additionalProperties, fieldName)
				}
			}
		}

		o.AdditionalProperties = additionalProperties
	} else {
		return err
	}

	return err
}

type NullableMfaEnrollPolicyRule struct {
	value *MfaEnrollPolicyRule
	isSet bool
}

func (v NullableMfaEnrollPolicyRule) Get() *MfaEnrollPolicyRule {
	return v.value
}

func (v *NullableMfaEnrollPolicyRule) Set(val *MfaEnrollPolicyRule) {
	v.value = val
	v.isSet = true
}

func (v NullableMfaEnrollPolicyRule) IsSet() bool {
	return v.isSet
}

func (v *NullableMfaEnrollPolicyRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMfaEnrollPolicyRule(val *MfaEnrollPolicyRule) *NullableMfaEnrollPolicyRule {
	return &NullableMfaEnrollPolicyRule{value: val, isSet: true}
}

func (v NullableMfaEnrollPolicyRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMfaEnrollPolicyRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

