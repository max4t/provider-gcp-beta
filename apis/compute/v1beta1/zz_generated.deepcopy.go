//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022 Max T.
*/

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionTargetTCPProxy) DeepCopyInto(out *RegionTargetTCPProxy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionTargetTCPProxy.
func (in *RegionTargetTCPProxy) DeepCopy() *RegionTargetTCPProxy {
	if in == nil {
		return nil
	}
	out := new(RegionTargetTCPProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RegionTargetTCPProxy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionTargetTCPProxyList) DeepCopyInto(out *RegionTargetTCPProxyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]RegionTargetTCPProxy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionTargetTCPProxyList.
func (in *RegionTargetTCPProxyList) DeepCopy() *RegionTargetTCPProxyList {
	if in == nil {
		return nil
	}
	out := new(RegionTargetTCPProxyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *RegionTargetTCPProxyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionTargetTCPProxyObservation) DeepCopyInto(out *RegionTargetTCPProxyObservation) {
	*out = *in
	if in.CreationTimestamp != nil {
		in, out := &in.CreationTimestamp, &out.CreationTimestamp
		*out = new(string)
		**out = **in
	}
	if in.ID != nil {
		in, out := &in.ID, &out.ID
		*out = new(string)
		**out = **in
	}
	if in.ProxyID != nil {
		in, out := &in.ProxyID, &out.ProxyID
		*out = new(float64)
		**out = **in
	}
	if in.SelfLink != nil {
		in, out := &in.SelfLink, &out.SelfLink
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionTargetTCPProxyObservation.
func (in *RegionTargetTCPProxyObservation) DeepCopy() *RegionTargetTCPProxyObservation {
	if in == nil {
		return nil
	}
	out := new(RegionTargetTCPProxyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionTargetTCPProxyParameters) DeepCopyInto(out *RegionTargetTCPProxyParameters) {
	*out = *in
	if in.BackendService != nil {
		in, out := &in.BackendService, &out.BackendService
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Project != nil {
		in, out := &in.Project, &out.Project
		*out = new(string)
		**out = **in
	}
	if in.ProxyBind != nil {
		in, out := &in.ProxyBind, &out.ProxyBind
		*out = new(bool)
		**out = **in
	}
	if in.ProxyHeader != nil {
		in, out := &in.ProxyHeader, &out.ProxyHeader
		*out = new(string)
		**out = **in
	}
	if in.Region != nil {
		in, out := &in.Region, &out.Region
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionTargetTCPProxyParameters.
func (in *RegionTargetTCPProxyParameters) DeepCopy() *RegionTargetTCPProxyParameters {
	if in == nil {
		return nil
	}
	out := new(RegionTargetTCPProxyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionTargetTCPProxySpec) DeepCopyInto(out *RegionTargetTCPProxySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionTargetTCPProxySpec.
func (in *RegionTargetTCPProxySpec) DeepCopy() *RegionTargetTCPProxySpec {
	if in == nil {
		return nil
	}
	out := new(RegionTargetTCPProxySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RegionTargetTCPProxyStatus) DeepCopyInto(out *RegionTargetTCPProxyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RegionTargetTCPProxyStatus.
func (in *RegionTargetTCPProxyStatus) DeepCopy() *RegionTargetTCPProxyStatus {
	if in == nil {
		return nil
	}
	out := new(RegionTargetTCPProxyStatus)
	in.DeepCopyInto(out)
	return out
}