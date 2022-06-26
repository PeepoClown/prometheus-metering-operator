//go:build !ignore_autogenerated
// +build !ignore_autogenerated

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GlobalPrometheusConfig) DeepCopyInto(out *GlobalPrometheusConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GlobalPrometheusConfig.
func (in *GlobalPrometheusConfig) DeepCopy() *GlobalPrometheusConfig {
	if in == nil {
		return nil
	}
	out := new(GlobalPrometheusConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusMetering) DeepCopyInto(out *PrometheusMetering) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusMetering.
func (in *PrometheusMetering) DeepCopy() *PrometheusMetering {
	if in == nil {
		return nil
	}
	out := new(PrometheusMetering)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrometheusMetering) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusMeteringConfig) DeepCopyInto(out *PrometheusMeteringConfig) {
	*out = *in
	if in.GlobalConfig != nil {
		in, out := &in.GlobalConfig, &out.GlobalConfig
		*out = new(GlobalPrometheusConfig)
		**out = **in
	}
	if in.RemoteWrite != nil {
		in, out := &in.RemoteWrite, &out.RemoteWrite
		*out = make([]RemoteStorageUrl, len(*in))
		copy(*out, *in)
	}
	if in.RemoteRead != nil {
		in, out := &in.RemoteRead, &out.RemoteRead
		*out = make([]RemoteStorageUrl, len(*in))
		copy(*out, *in)
	}
	if in.ScrapeJobs != nil {
		in, out := &in.ScrapeJobs, &out.ScrapeJobs
		*out = make([]ScrapeJob, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusMeteringConfig.
func (in *PrometheusMeteringConfig) DeepCopy() *PrometheusMeteringConfig {
	if in == nil {
		return nil
	}
	out := new(PrometheusMeteringConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusMeteringList) DeepCopyInto(out *PrometheusMeteringList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]PrometheusMetering, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusMeteringList.
func (in *PrometheusMeteringList) DeepCopy() *PrometheusMeteringList {
	if in == nil {
		return nil
	}
	out := new(PrometheusMeteringList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *PrometheusMeteringList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusMeteringSpec) DeepCopyInto(out *PrometheusMeteringSpec) {
	*out = *in
	if in.PrometheusConfig != nil {
		in, out := &in.PrometheusConfig, &out.PrometheusConfig
		*out = new(PrometheusMeteringConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = new(v1.ResourceRequirements)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusMeteringSpec.
func (in *PrometheusMeteringSpec) DeepCopy() *PrometheusMeteringSpec {
	if in == nil {
		return nil
	}
	out := new(PrometheusMeteringSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PrometheusMeteringStatus) DeepCopyInto(out *PrometheusMeteringStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PrometheusMeteringStatus.
func (in *PrometheusMeteringStatus) DeepCopy() *PrometheusMeteringStatus {
	if in == nil {
		return nil
	}
	out := new(PrometheusMeteringStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RemoteStorageUrl) DeepCopyInto(out *RemoteStorageUrl) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RemoteStorageUrl.
func (in *RemoteStorageUrl) DeepCopy() *RemoteStorageUrl {
	if in == nil {
		return nil
	}
	out := new(RemoteStorageUrl)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ScrapeJob) DeepCopyInto(out *ScrapeJob) {
	*out = *in
	if in.StaticConfigs != nil {
		in, out := &in.StaticConfigs, &out.StaticConfigs
		*out = make([]StaticConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ScrapeJob.
func (in *ScrapeJob) DeepCopy() *ScrapeJob {
	if in == nil {
		return nil
	}
	out := new(ScrapeJob)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StaticConfig) DeepCopyInto(out *StaticConfig) {
	*out = *in
	if in.Targets != nil {
		in, out := &in.Targets, &out.Targets
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StaticConfig.
func (in *StaticConfig) DeepCopy() *StaticConfig {
	if in == nil {
		return nil
	}
	out := new(StaticConfig)
	in.DeepCopyInto(out)
	return out
}
