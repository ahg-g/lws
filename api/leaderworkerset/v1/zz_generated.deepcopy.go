//go:build !ignore_autogenerated

/*
Copyright 2023.

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

package v1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeaderWorkerSet) DeepCopyInto(out *LeaderWorkerSet) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeaderWorkerSet.
func (in *LeaderWorkerSet) DeepCopy() *LeaderWorkerSet {
	if in == nil {
		return nil
	}
	out := new(LeaderWorkerSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LeaderWorkerSet) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeaderWorkerSetList) DeepCopyInto(out *LeaderWorkerSetList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]LeaderWorkerSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeaderWorkerSetList.
func (in *LeaderWorkerSetList) DeepCopy() *LeaderWorkerSetList {
	if in == nil {
		return nil
	}
	out := new(LeaderWorkerSetList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LeaderWorkerSetList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeaderWorkerSetSpec) DeepCopyInto(out *LeaderWorkerSetSpec) {
	*out = *in
	if in.Replicas != nil {
		in, out := &in.Replicas, &out.Replicas
		*out = new(int32)
		**out = **in
	}
	in.LeaderWorkerTemplate.DeepCopyInto(&out.LeaderWorkerTemplate)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeaderWorkerSetSpec.
func (in *LeaderWorkerSetSpec) DeepCopy() *LeaderWorkerSetSpec {
	if in == nil {
		return nil
	}
	out := new(LeaderWorkerSetSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeaderWorkerSetStatus) DeepCopyInto(out *LeaderWorkerSetStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeaderWorkerSetStatus.
func (in *LeaderWorkerSetStatus) DeepCopy() *LeaderWorkerSetStatus {
	if in == nil {
		return nil
	}
	out := new(LeaderWorkerSetStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LeaderWorkerTemplate) DeepCopyInto(out *LeaderWorkerTemplate) {
	*out = *in
	if in.LeaderTemplate != nil {
		in, out := &in.LeaderTemplate, &out.LeaderTemplate
		*out = new(corev1.PodTemplateSpec)
		(*in).DeepCopyInto(*out)
	}
	in.WorkerTemplate.DeepCopyInto(&out.WorkerTemplate)
	if in.Size != nil {
		in, out := &in.Size, &out.Size
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LeaderWorkerTemplate.
func (in *LeaderWorkerTemplate) DeepCopy() *LeaderWorkerTemplate {
	if in == nil {
		return nil
	}
	out := new(LeaderWorkerTemplate)
	in.DeepCopyInto(out)
	return out
}
