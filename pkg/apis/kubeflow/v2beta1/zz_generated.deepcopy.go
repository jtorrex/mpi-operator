//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Copyright 2023 The Kubeflow Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Code generated by deepcopy-gen. DO NOT EDIT.

package v2beta1

import (
	v1 "github.com/kubeflow/common/pkg/apis/common/v1"
	corev1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MPIJob) DeepCopyInto(out *MPIJob) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MPIJob.
func (in *MPIJob) DeepCopy() *MPIJob {
	if in == nil {
		return nil
	}
	out := new(MPIJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MPIJob) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MPIJobList) DeepCopyInto(out *MPIJobList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]MPIJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MPIJobList.
func (in *MPIJobList) DeepCopy() *MPIJobList {
	if in == nil {
		return nil
	}
	out := new(MPIJobList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *MPIJobList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MPIJobSpec) DeepCopyInto(out *MPIJobSpec) {
	*out = *in
	if in.SlotsPerWorker != nil {
		in, out := &in.SlotsPerWorker, &out.SlotsPerWorker
		*out = new(int32)
		**out = **in
	}
	in.RunPolicy.DeepCopyInto(&out.RunPolicy)
	if in.MPIReplicaSpecs != nil {
		in, out := &in.MPIReplicaSpecs, &out.MPIReplicaSpecs
		*out = make(map[MPIReplicaType]*v1.ReplicaSpec, len(*in))
		for key, val := range *in {
			var outVal *v1.ReplicaSpec
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = new(v1.ReplicaSpec)
				(*in).DeepCopyInto(*out)
			}
			(*out)[key] = outVal
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MPIJobSpec.
func (in *MPIJobSpec) DeepCopy() *MPIJobSpec {
	if in == nil {
		return nil
	}
	out := new(MPIJobSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RunPolicy) DeepCopyInto(out *RunPolicy) {
	*out = *in
	if in.CleanPodPolicy != nil {
		in, out := &in.CleanPodPolicy, &out.CleanPodPolicy
		*out = new(CleanPodPolicy)
		**out = **in
	}
	if in.TTLSecondsAfterFinished != nil {
		in, out := &in.TTLSecondsAfterFinished, &out.TTLSecondsAfterFinished
		*out = new(int32)
		**out = **in
	}
	if in.ActiveDeadlineSeconds != nil {
		in, out := &in.ActiveDeadlineSeconds, &out.ActiveDeadlineSeconds
		*out = new(int64)
		**out = **in
	}
	if in.BackoffLimit != nil {
		in, out := &in.BackoffLimit, &out.BackoffLimit
		*out = new(int32)
		**out = **in
	}
	if in.SchedulingPolicy != nil {
		in, out := &in.SchedulingPolicy, &out.SchedulingPolicy
		*out = new(SchedulingPolicy)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RunPolicy.
func (in *RunPolicy) DeepCopy() *RunPolicy {
	if in == nil {
		return nil
	}
	out := new(RunPolicy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulingPolicy) DeepCopyInto(out *SchedulingPolicy) {
	*out = *in
	if in.MinAvailable != nil {
		in, out := &in.MinAvailable, &out.MinAvailable
		*out = new(int32)
		**out = **in
	}
	if in.MinResources != nil {
		in, out := &in.MinResources, &out.MinResources
		*out = new(corev1.ResourceList)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[corev1.ResourceName]resource.Quantity, len(*in))
			for key, val := range *in {
				(*out)[key] = val.DeepCopy()
			}
		}
	}
	if in.ScheduleTimeoutSeconds != nil {
		in, out := &in.ScheduleTimeoutSeconds, &out.ScheduleTimeoutSeconds
		*out = new(int32)
		**out = **in
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulingPolicy.
func (in *SchedulingPolicy) DeepCopy() *SchedulingPolicy {
	if in == nil {
		return nil
	}
	out := new(SchedulingPolicy)
	in.DeepCopyInto(out)
	return out
}
