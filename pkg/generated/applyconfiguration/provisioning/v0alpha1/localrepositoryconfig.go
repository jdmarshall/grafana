// SPDX-License-Identifier: AGPL-3.0-only

// Code generated by applyconfiguration-gen. DO NOT EDIT.

package v0alpha1

// LocalRepositoryConfigApplyConfiguration represents a declarative configuration of the LocalRepositoryConfig type for use
// with apply.
type LocalRepositoryConfigApplyConfiguration struct {
	Path *string `json:"path,omitempty"`
}

// LocalRepositoryConfigApplyConfiguration constructs a declarative configuration of the LocalRepositoryConfig type for use with
// apply.
func LocalRepositoryConfig() *LocalRepositoryConfigApplyConfiguration {
	return &LocalRepositoryConfigApplyConfiguration{}
}

// WithPath sets the Path field in the declarative configuration to the given value
// and returns the receiver, so that objects can be built by chaining "With" function invocations.
// If called multiple times, the Path field is set to the value of the last call.
func (b *LocalRepositoryConfigApplyConfiguration) WithPath(value string) *LocalRepositoryConfigApplyConfiguration {
	b.Path = &value
	return b
}
