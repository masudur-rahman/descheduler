//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2024 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha2

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeschedulerPolicy) DeepCopyInto(out *DeschedulerPolicy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	if in.Profiles != nil {
		in, out := &in.Profiles, &out.Profiles
		*out = make([]DeschedulerProfile, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NodeSelector != nil {
		in, out := &in.NodeSelector, &out.NodeSelector
		*out = new(string)
		**out = **in
	}
	if in.MaxNoOfPodsToEvictPerNode != nil {
		in, out := &in.MaxNoOfPodsToEvictPerNode, &out.MaxNoOfPodsToEvictPerNode
		*out = new(uint)
		**out = **in
	}
	if in.MaxNoOfPodsToEvictPerNamespace != nil {
		in, out := &in.MaxNoOfPodsToEvictPerNamespace, &out.MaxNoOfPodsToEvictPerNamespace
		*out = new(uint)
		**out = **in
	}
	if in.ProcessNotReadyNodes != nil {
		in, out := &in.ProcessNotReadyNodes, &out.ProcessNotReadyNodes
		*out = new(bool)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeschedulerPolicy.
func (in *DeschedulerPolicy) DeepCopy() *DeschedulerPolicy {
	if in == nil {
		return nil
	}
	out := new(DeschedulerPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *DeschedulerPolicy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DeschedulerProfile) DeepCopyInto(out *DeschedulerProfile) {
	*out = *in
	if in.PluginConfigs != nil {
		in, out := &in.PluginConfigs, &out.PluginConfigs
		*out = make([]PluginConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Plugins.DeepCopyInto(&out.Plugins)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DeschedulerProfile.
func (in *DeschedulerProfile) DeepCopy() *DeschedulerProfile {
	if in == nil {
		return nil
	}
	out := new(DeschedulerProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginConfig) DeepCopyInto(out *PluginConfig) {
	*out = *in
	in.Args.DeepCopyInto(&out.Args)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginConfig.
func (in *PluginConfig) DeepCopy() *PluginConfig {
	if in == nil {
		return nil
	}
	out := new(PluginConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PluginSet) DeepCopyInto(out *PluginSet) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Disabled != nil {
		in, out := &in.Disabled, &out.Disabled
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PluginSet.
func (in *PluginSet) DeepCopy() *PluginSet {
	if in == nil {
		return nil
	}
	out := new(PluginSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Plugins) DeepCopyInto(out *Plugins) {
	*out = *in
	in.PreSort.DeepCopyInto(&out.PreSort)
	in.Sort.DeepCopyInto(&out.Sort)
	in.Deschedule.DeepCopyInto(&out.Deschedule)
	in.Balance.DeepCopyInto(&out.Balance)
	in.Filter.DeepCopyInto(&out.Filter)
	in.PreEvictionFilter.DeepCopyInto(&out.PreEvictionFilter)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Plugins.
func (in *Plugins) DeepCopy() *Plugins {
	if in == nil {
		return nil
	}
	out := new(Plugins)
	in.DeepCopyInto(out)
	return out
}
