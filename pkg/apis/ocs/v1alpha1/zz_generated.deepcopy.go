// +build !ignore_autogenerated

// Code generated by operator-sdk. DO NOT EDIT.

package v1alpha1

import (
	v1 "github.com/openshift/custom-resource-status/conditions/v1"
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCSInitialization) DeepCopyInto(out *OCSInitialization) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCSInitialization.
func (in *OCSInitialization) DeepCopy() *OCSInitialization {
	if in == nil {
		return nil
	}
	out := new(OCSInitialization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OCSInitialization) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCSInitializationList) DeepCopyInto(out *OCSInitializationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OCSInitialization, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCSInitializationList.
func (in *OCSInitializationList) DeepCopy() *OCSInitializationList {
	if in == nil {
		return nil
	}
	out := new(OCSInitializationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OCSInitializationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCSInitializationSpec) DeepCopyInto(out *OCSInitializationSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCSInitializationSpec.
func (in *OCSInitializationSpec) DeepCopy() *OCSInitializationSpec {
	if in == nil {
		return nil
	}
	out := new(OCSInitializationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCSInitializationStatus) DeepCopyInto(out *OCSInitializationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RelatedObjects != nil {
		in, out := &in.RelatedObjects, &out.RelatedObjects
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCSInitializationStatus.
func (in *OCSInitializationStatus) DeepCopy() *OCSInitializationStatus {
	if in == nil {
		return nil
	}
	out := new(OCSInitializationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageCluster) DeepCopyInto(out *StorageCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageCluster.
func (in *StorageCluster) DeepCopy() *StorageCluster {
	if in == nil {
		return nil
	}
	out := new(StorageCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClusterInitialization) DeepCopyInto(out *StorageClusterInitialization) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClusterInitialization.
func (in *StorageClusterInitialization) DeepCopy() *StorageClusterInitialization {
	if in == nil {
		return nil
	}
	out := new(StorageClusterInitialization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageClusterInitialization) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClusterInitializationList) DeepCopyInto(out *StorageClusterInitializationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StorageClusterInitialization, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClusterInitializationList.
func (in *StorageClusterInitializationList) DeepCopy() *StorageClusterInitializationList {
	if in == nil {
		return nil
	}
	out := new(StorageClusterInitializationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageClusterInitializationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClusterInitializationSpec) DeepCopyInto(out *StorageClusterInitializationSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClusterInitializationSpec.
func (in *StorageClusterInitializationSpec) DeepCopy() *StorageClusterInitializationSpec {
	if in == nil {
		return nil
	}
	out := new(StorageClusterInitializationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClusterInitializationStatus) DeepCopyInto(out *StorageClusterInitializationStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClusterInitializationStatus.
func (in *StorageClusterInitializationStatus) DeepCopy() *StorageClusterInitializationStatus {
	if in == nil {
		return nil
	}
	out := new(StorageClusterInitializationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClusterList) DeepCopyInto(out *StorageClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StorageCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClusterList.
func (in *StorageClusterList) DeepCopy() *StorageClusterList {
	if in == nil {
		return nil
	}
	out := new(StorageClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClusterSpec) DeepCopyInto(out *StorageClusterSpec) {
	*out = *in
	if in.StorageDeviceSets != nil {
		in, out := &in.StorageDeviceSets, &out.StorageDeviceSets
		*out = make([]StorageDeviceSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MonPVCTemplate != nil {
		in, out := &in.MonPVCTemplate, &out.MonPVCTemplate
		*out = new(corev1.PersistentVolumeClaim)
		(*in).DeepCopyInto(*out)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClusterSpec.
func (in *StorageClusterSpec) DeepCopy() *StorageClusterSpec {
	if in == nil {
		return nil
	}
	out := new(StorageClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClusterStatus) DeepCopyInto(out *StorageClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]v1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RelatedObjects != nil {
		in, out := &in.RelatedObjects, &out.RelatedObjects
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClusterStatus.
func (in *StorageClusterStatus) DeepCopy() *StorageClusterStatus {
	if in == nil {
		return nil
	}
	out := new(StorageClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageDeviceSet) DeepCopyInto(out *StorageDeviceSet) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	in.Placement.DeepCopyInto(&out.Placement)
	out.Config = in.Config
	in.DataPVCTemplate.DeepCopyInto(&out.DataPVCTemplate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageDeviceSet.
func (in *StorageDeviceSet) DeepCopy() *StorageDeviceSet {
	if in == nil {
		return nil
	}
	out := new(StorageDeviceSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageDeviceSetConfig) DeepCopyInto(out *StorageDeviceSetConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageDeviceSetConfig.
func (in *StorageDeviceSetConfig) DeepCopy() *StorageDeviceSetConfig {
	if in == nil {
		return nil
	}
	out := new(StorageDeviceSetConfig)
	in.DeepCopyInto(out)
	return out
}
