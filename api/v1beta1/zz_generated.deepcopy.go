//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright 2022.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1beta1

import (
	"github.com/openstack-k8s-operators/lib-common/modules/common/condition"
	"github.com/openstack-k8s-operators/lib-common/modules/common/service"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PasswordSelector) DeepCopyInto(out *PasswordSelector) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PasswordSelector.
func (in *PasswordSelector) DeepCopy() *PasswordSelector {
	if in == nil {
		return nil
	}
	out := new(PasswordSelector)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ProxyOverrideSpec) DeepCopyInto(out *ProxyOverrideSpec) {
	*out = *in
	if in.Service != nil {
		in, out := &in.Service, &out.Service
		*out = make(map[service.Endpoint]service.RoutedOverrideSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ProxyOverrideSpec.
func (in *ProxyOverrideSpec) DeepCopy() *ProxyOverrideSpec {
	if in == nil {
		return nil
	}
	out := new(ProxyOverrideSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Swift) DeepCopyInto(out *Swift) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Swift.
func (in *Swift) DeepCopy() *Swift {
	if in == nil {
		return nil
	}
	out := new(Swift)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Swift) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftDefaults) DeepCopyInto(out *SwiftDefaults) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftDefaults.
func (in *SwiftDefaults) DeepCopy() *SwiftDefaults {
	if in == nil {
		return nil
	}
	out := new(SwiftDefaults)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftDisk) DeepCopyInto(out *SwiftDisk) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftDisk.
func (in *SwiftDisk) DeepCopy() *SwiftDisk {
	if in == nil {
		return nil
	}
	out := new(SwiftDisk)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftList) DeepCopyInto(out *SwiftList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Swift, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftList.
func (in *SwiftList) DeepCopy() *SwiftList {
	if in == nil {
		return nil
	}
	out := new(SwiftList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SwiftList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftProxy) DeepCopyInto(out *SwiftProxy) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftProxy.
func (in *SwiftProxy) DeepCopy() *SwiftProxy {
	if in == nil {
		return nil
	}
	out := new(SwiftProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SwiftProxy) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftProxyList) DeepCopyInto(out *SwiftProxyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SwiftProxy, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftProxyList.
func (in *SwiftProxyList) DeepCopy() *SwiftProxyList {
	if in == nil {
		return nil
	}
	out := new(SwiftProxyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SwiftProxyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftProxySpec) DeepCopyInto(out *SwiftProxySpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	out.PasswordSelectors = in.PasswordSelectors
	in.Override.DeepCopyInto(&out.Override)
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.TLS.DeepCopyInto(&out.TLS)
	if in.DefaultConfigOverwrite != nil {
		in, out := &in.DefaultConfigOverwrite, &out.DefaultConfigOverwrite
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftProxySpec.
func (in *SwiftProxySpec) DeepCopy() *SwiftProxySpec {
	if in == nil {
		return nil
	}
	out := new(SwiftProxySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftProxyStatus) DeepCopyInto(out *SwiftProxyStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftProxyStatus.
func (in *SwiftProxyStatus) DeepCopy() *SwiftProxyStatus {
	if in == nil {
		return nil
	}
	out := new(SwiftProxyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftRing) DeepCopyInto(out *SwiftRing) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftRing.
func (in *SwiftRing) DeepCopy() *SwiftRing {
	if in == nil {
		return nil
	}
	out := new(SwiftRing)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SwiftRing) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftRingList) DeepCopyInto(out *SwiftRingList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SwiftRing, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftRingList.
func (in *SwiftRingList) DeepCopy() *SwiftRingList {
	if in == nil {
		return nil
	}
	out := new(SwiftRingList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SwiftRingList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftRingSpec) DeepCopyInto(out *SwiftRingSpec) {
	*out = *in
	if in.RingReplicas != nil {
		in, out := &in.RingReplicas, &out.RingReplicas
		*out = new(int64)
		**out = **in
	}
	if in.PartPower != nil {
		in, out := &in.PartPower, &out.PartPower
		*out = new(int64)
		**out = **in
	}
	if in.MinPartHours != nil {
		in, out := &in.MinPartHours, &out.MinPartHours
		*out = new(int64)
		**out = **in
	}
	out.TLS = in.TLS
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftRingSpec.
func (in *SwiftRingSpec) DeepCopy() *SwiftRingSpec {
	if in == nil {
		return nil
	}
	out := new(SwiftRingSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftRingStatus) DeepCopyInto(out *SwiftRingStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Hash != nil {
		in, out := &in.Hash, &out.Hash
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftRingStatus.
func (in *SwiftRingStatus) DeepCopy() *SwiftRingStatus {
	if in == nil {
		return nil
	}
	out := new(SwiftRingStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftSpec) DeepCopyInto(out *SwiftSpec) {
	*out = *in
	in.SwiftRing.DeepCopyInto(&out.SwiftRing)
	in.SwiftStorage.DeepCopyInto(&out.SwiftStorage)
	in.SwiftProxy.DeepCopyInto(&out.SwiftProxy)
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftSpec.
func (in *SwiftSpec) DeepCopy() *SwiftSpec {
	if in == nil {
		return nil
	}
	out := new(SwiftSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftStatus) DeepCopyInto(out *SwiftStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftStatus.
func (in *SwiftStatus) DeepCopy() *SwiftStatus {
	if in == nil {
		return nil
	}
	out := new(SwiftStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftStorage) DeepCopyInto(out *SwiftStorage) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftStorage.
func (in *SwiftStorage) DeepCopy() *SwiftStorage {
	if in == nil {
		return nil
	}
	out := new(SwiftStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SwiftStorage) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftStorageList) DeepCopyInto(out *SwiftStorageList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SwiftStorage, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftStorageList.
func (in *SwiftStorageList) DeepCopy() *SwiftStorageList {
	if in == nil {
		return nil
	}
	out := new(SwiftStorageList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SwiftStorageList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftStorageSpec) DeepCopyInto(out *SwiftStorageSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.DefaultConfigOverwrite != nil {
		in, out := &in.DefaultConfigOverwrite, &out.DefaultConfigOverwrite
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftStorageSpec.
func (in *SwiftStorageSpec) DeepCopy() *SwiftStorageSpec {
	if in == nil {
		return nil
	}
	out := new(SwiftStorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SwiftStorageStatus) DeepCopyInto(out *SwiftStorageStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make(condition.Conditions, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.NetworkAttachments != nil {
		in, out := &in.NetworkAttachments, &out.NetworkAttachments
		*out = make(map[string][]string, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make([]string, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SwiftStorageStatus.
func (in *SwiftStorageStatus) DeepCopy() *SwiftStorageStatus {
	if in == nil {
		return nil
	}
	out := new(SwiftStorageStatus)
	in.DeepCopyInto(out)
	return out
}
