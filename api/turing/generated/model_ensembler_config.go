/*
 * Turing Minimal Openapi Spec for SDK
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 0.0.1
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// EnsemblerConfig struct for EnsemblerConfig
type EnsemblerConfig struct {
	Version  string              `json:"version"`
	Kind     EnsemblerConfigKind `json:"kind"`
	Metadata *EnsemblingJobMeta  `json:"metadata,omitempty"`
	Spec     EnsemblingJobSpec   `json:"spec"`
}

// NewEnsemblerConfig instantiates a new EnsemblerConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnsemblerConfig(version string, kind EnsemblerConfigKind, spec EnsemblingJobSpec) *EnsemblerConfig {
	this := EnsemblerConfig{}
	this.Version = version
	this.Kind = kind
	this.Spec = spec
	return &this
}

// NewEnsemblerConfigWithDefaults instantiates a new EnsemblerConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnsemblerConfigWithDefaults() *EnsemblerConfig {
	this := EnsemblerConfig{}
	return &this
}

// GetVersion returns the Version field value
func (o *EnsemblerConfig) GetVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *EnsemblerConfig) GetVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *EnsemblerConfig) SetVersion(v string) {
	o.Version = v
}

// GetKind returns the Kind field value
func (o *EnsemblerConfig) GetKind() EnsemblerConfigKind {
	if o == nil {
		var ret EnsemblerConfigKind
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *EnsemblerConfig) GetKindOk() (*EnsemblerConfigKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *EnsemblerConfig) SetKind(v EnsemblerConfigKind) {
	o.Kind = v
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *EnsemblerConfig) GetMetadata() EnsemblingJobMeta {
	if o == nil || o.Metadata == nil {
		var ret EnsemblingJobMeta
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EnsemblerConfig) GetMetadataOk() (*EnsemblingJobMeta, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *EnsemblerConfig) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given EnsemblingJobMeta and assigns it to the Metadata field.
func (o *EnsemblerConfig) SetMetadata(v EnsemblingJobMeta) {
	o.Metadata = &v
}

// GetSpec returns the Spec field value
func (o *EnsemblerConfig) GetSpec() EnsemblingJobSpec {
	if o == nil {
		var ret EnsemblingJobSpec
		return ret
	}

	return o.Spec
}

// GetSpecOk returns a tuple with the Spec field value
// and a boolean to check if the value has been set.
func (o *EnsemblerConfig) GetSpecOk() (*EnsemblingJobSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Spec, true
}

// SetSpec sets field value
func (o *EnsemblerConfig) SetSpec(v EnsemblingJobSpec) {
	o.Spec = v
}

func (o EnsemblerConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["version"] = o.Version
	}
	if true {
		toSerialize["kind"] = o.Kind
	}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if true {
		toSerialize["spec"] = o.Spec
	}
	return json.Marshal(toSerialize)
}

type NullableEnsemblerConfig struct {
	value *EnsemblerConfig
	isSet bool
}

func (v NullableEnsemblerConfig) Get() *EnsemblerConfig {
	return v.value
}

func (v *NullableEnsemblerConfig) Set(val *EnsemblerConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableEnsemblerConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableEnsemblerConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnsemblerConfig(val *EnsemblerConfig) *NullableEnsemblerConfig {
	return &NullableEnsemblerConfig{value: val, isSet: true}
}

func (v NullableEnsemblerConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnsemblerConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}