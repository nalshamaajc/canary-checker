// +build !ignore_autogenerated

/*
Copyright 2020 The Kubernetes authors.

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
	"github.com/flanksource/kommons"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Authentication) DeepCopyInto(out *Authentication) {
	*out = *in
	in.Username.DeepCopyInto(&out.Username)
	in.Password.DeepCopyInto(&out.Password)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Authentication.
func (in *Authentication) DeepCopy() *Authentication {
	if in == nil {
		return nil
	}
	out := new(Authentication)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Bucket) DeepCopyInto(out *Bucket) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Bucket.
func (in *Bucket) DeepCopy() *Bucket {
	if in == nil {
		return nil
	}
	out := new(Bucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Canary) DeepCopyInto(out *Canary) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Canary.
func (in *Canary) DeepCopy() *Canary {
	if in == nil {
		return nil
	}
	out := new(Canary)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Canary) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryList) DeepCopyInto(out *CanaryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Canary, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryList.
func (in *CanaryList) DeepCopy() *CanaryList {
	if in == nil {
		return nil
	}
	out := new(CanaryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CanaryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanarySpec) DeepCopyInto(out *CanarySpec) {
	*out = *in
	if in.Env != nil {
		in, out := &in.Env, &out.Env
		*out = make(map[string]VarSource, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.HTTP != nil {
		in, out := &in.HTTP, &out.HTTP
		*out = make([]HTTPCheck, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DNS != nil {
		in, out := &in.DNS, &out.DNS
		*out = make([]DNSCheck, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.DockerPull != nil {
		in, out := &in.DockerPull, &out.DockerPull
		*out = make([]DockerPullCheck, len(*in))
		copy(*out, *in)
	}
	if in.DockerPush != nil {
		in, out := &in.DockerPush, &out.DockerPush
		*out = make([]DockerPushCheck, len(*in))
		copy(*out, *in)
	}
	if in.ContainerdPull != nil {
		in, out := &in.ContainerdPull, &out.ContainerdPull
		*out = make([]ContainerdPullCheck, len(*in))
		copy(*out, *in)
	}
	if in.ContainerdPush != nil {
		in, out := &in.ContainerdPush, &out.ContainerdPush
		*out = make([]ContainerdPushCheck, len(*in))
		copy(*out, *in)
	}
	if in.S3 != nil {
		in, out := &in.S3, &out.S3
		*out = make([]S3Check, len(*in))
		copy(*out, *in)
	}
	if in.S3Bucket != nil {
		in, out := &in.S3Bucket, &out.S3Bucket
		*out = make([]S3BucketCheck, len(*in))
		copy(*out, *in)
	}
	if in.TCP != nil {
		in, out := &in.TCP, &out.TCP
		*out = make([]TCPCheck, len(*in))
		copy(*out, *in)
	}
	if in.Pod != nil {
		in, out := &in.Pod, &out.Pod
		*out = make([]PodCheck, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.LDAP != nil {
		in, out := &in.LDAP, &out.LDAP
		*out = make([]LDAPCheck, len(*in))
		copy(*out, *in)
	}
	if in.ICMP != nil {
		in, out := &in.ICMP, &out.ICMP
		*out = make([]ICMPCheck, len(*in))
		copy(*out, *in)
	}
	if in.Postgres != nil {
		in, out := &in.Postgres, &out.Postgres
		*out = make([]PostgresCheck, len(*in))
		copy(*out, *in)
	}
	if in.Mssql != nil {
		in, out := &in.Mssql, &out.Mssql
		*out = make([]MssqlCheck, len(*in))
		copy(*out, *in)
	}
	if in.Restic != nil {
		in, out := &in.Restic, &out.Restic
		*out = make([]ResticCheck, len(*in))
		copy(*out, *in)
	}
	if in.Jmeter != nil {
		in, out := &in.Jmeter, &out.Jmeter
		*out = make([]JmeterCheck, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Helm != nil {
		in, out := &in.Helm, &out.Helm
		*out = make([]HelmCheck, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Namespace != nil {
		in, out := &in.Namespace, &out.Namespace
		*out = make([]NamespaceCheck, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Redis != nil {
		in, out := &in.Redis, &out.Redis
		*out = make([]RedisCheck, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanarySpec.
func (in *CanarySpec) DeepCopy() *CanarySpec {
	if in == nil {
		return nil
	}
	out := new(CanarySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryStatus) DeepCopyInto(out *CanaryStatus) {
	*out = *in
	if in.LastTransitionedTime != nil {
		in, out := &in.LastTransitionedTime, &out.LastTransitionedTime
		*out = (*in).DeepCopy()
	}
	if in.LastCheck != nil {
		in, out := &in.LastCheck, &out.LastCheck
		*out = (*in).DeepCopy()
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(CanaryStatusCondition)
		**out = **in
	}
	if in.Message != nil {
		in, out := &in.Message, &out.Message
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryStatus.
func (in *CanaryStatus) DeepCopy() *CanaryStatus {
	if in == nil {
		return nil
	}
	out := new(CanaryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Condition) DeepCopyInto(out *Condition) {
	*out = *in
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Condition.
func (in *Condition) DeepCopy() *Condition {
	if in == nil {
		return nil
	}
	out := new(Condition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerdPullCheck) DeepCopyInto(out *ContainerdPullCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerdPullCheck.
func (in *ContainerdPullCheck) DeepCopy() *ContainerdPullCheck {
	if in == nil {
		return nil
	}
	out := new(ContainerdPullCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ContainerdPushCheck) DeepCopyInto(out *ContainerdPushCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ContainerdPushCheck.
func (in *ContainerdPushCheck) DeepCopy() *ContainerdPushCheck {
	if in == nil {
		return nil
	}
	out := new(ContainerdPushCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNS) DeepCopyInto(out *DNS) {
	*out = *in
	in.DNSCheck.DeepCopyInto(&out.DNSCheck)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNS.
func (in *DNS) DeepCopy() *DNS {
	if in == nil {
		return nil
	}
	out := new(DNS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DNSCheck) DeepCopyInto(out *DNSCheck) {
	*out = *in
	if in.ExactReply != nil {
		in, out := &in.ExactReply, &out.ExactReply
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DNSCheck.
func (in *DNSCheck) DeepCopy() *DNSCheck {
	if in == nil {
		return nil
	}
	out := new(DNSCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerPull) DeepCopyInto(out *DockerPull) {
	*out = *in
	out.DockerPullCheck = in.DockerPullCheck
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerPull.
func (in *DockerPull) DeepCopy() *DockerPull {
	if in == nil {
		return nil
	}
	out := new(DockerPull)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerPullCheck) DeepCopyInto(out *DockerPullCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerPullCheck.
func (in *DockerPullCheck) DeepCopy() *DockerPullCheck {
	if in == nil {
		return nil
	}
	out := new(DockerPullCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerPush) DeepCopyInto(out *DockerPush) {
	*out = *in
	out.DockerPushCheck = in.DockerPushCheck
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerPush.
func (in *DockerPush) DeepCopy() *DockerPush {
	if in == nil {
		return nil
	}
	out := new(DockerPush)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DockerPushCheck) DeepCopyInto(out *DockerPushCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DockerPushCheck.
func (in *DockerPushCheck) DeepCopy() *DockerPushCheck {
	if in == nil {
		return nil
	}
	out := new(DockerPushCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTP) DeepCopyInto(out *HTTP) {
	*out = *in
	in.HTTPCheck.DeepCopyInto(&out.HTTPCheck)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTP.
func (in *HTTP) DeepCopy() *HTTP {
	if in == nil {
		return nil
	}
	out := new(HTTP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HTTPCheck) DeepCopyInto(out *HTTPCheck) {
	*out = *in
	if in.ResponseCodes != nil {
		in, out := &in.ResponseCodes, &out.ResponseCodes
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
	out.ResponseJSONContent = in.ResponseJSONContent
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = make([]kommons.EnvVar, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Authentication != nil {
		in, out := &in.Authentication, &out.Authentication
		*out = new(Authentication)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HTTPCheck.
func (in *HTTPCheck) DeepCopy() *HTTPCheck {
	if in == nil {
		return nil
	}
	out := new(HTTPCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Helm) DeepCopyInto(out *Helm) {
	*out = *in
	in.HelmCheck.DeepCopyInto(&out.HelmCheck)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Helm.
func (in *Helm) DeepCopy() *Helm {
	if in == nil {
		return nil
	}
	out := new(Helm)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmCheck) DeepCopyInto(out *HelmCheck) {
	*out = *in
	if in.CaFile != nil {
		in, out := &in.CaFile, &out.CaFile
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmCheck.
func (in *HelmCheck) DeepCopy() *HelmCheck {
	if in == nil {
		return nil
	}
	out := new(HelmCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ICMP) DeepCopyInto(out *ICMP) {
	*out = *in
	out.ICMPCheck = in.ICMPCheck
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ICMP.
func (in *ICMP) DeepCopy() *ICMP {
	if in == nil {
		return nil
	}
	out := new(ICMP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ICMPCheck) DeepCopyInto(out *ICMPCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ICMPCheck.
func (in *ICMPCheck) DeepCopy() *ICMPCheck {
	if in == nil {
		return nil
	}
	out := new(ICMPCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JSONCheck) DeepCopyInto(out *JSONCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JSONCheck.
func (in *JSONCheck) DeepCopy() *JSONCheck {
	if in == nil {
		return nil
	}
	out := new(JSONCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Jmeter) DeepCopyInto(out *Jmeter) {
	*out = *in
	in.JmeterCheck.DeepCopyInto(&out.JmeterCheck)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Jmeter.
func (in *Jmeter) DeepCopy() *Jmeter {
	if in == nil {
		return nil
	}
	out := new(Jmeter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *JmeterCheck) DeepCopyInto(out *JmeterCheck) {
	*out = *in
	in.Jmx.DeepCopyInto(&out.Jmx)
	if in.Properties != nil {
		in, out := &in.Properties, &out.Properties
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.SystemProperties != nil {
		in, out := &in.SystemProperties, &out.SystemProperties
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new JmeterCheck.
func (in *JmeterCheck) DeepCopy() *JmeterCheck {
	if in == nil {
		return nil
	}
	out := new(JmeterCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LDAP) DeepCopyInto(out *LDAP) {
	*out = *in
	out.LDAPCheck = in.LDAPCheck
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LDAP.
func (in *LDAP) DeepCopy() *LDAP {
	if in == nil {
		return nil
	}
	out := new(LDAP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LDAPCheck) DeepCopyInto(out *LDAPCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LDAPCheck.
func (in *LDAPCheck) DeepCopy() *LDAPCheck {
	if in == nil {
		return nil
	}
	out := new(LDAPCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MsSQL) DeepCopyInto(out *MsSQL) {
	*out = *in
	out.MssqlCheck = in.MssqlCheck
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MsSQL.
func (in *MsSQL) DeepCopy() *MsSQL {
	if in == nil {
		return nil
	}
	out := new(MsSQL)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MssqlCheck) DeepCopyInto(out *MssqlCheck) {
	*out = *in
	out.SQLCheck = in.SQLCheck
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MssqlCheck.
func (in *MssqlCheck) DeepCopy() *MssqlCheck {
	if in == nil {
		return nil
	}
	out := new(MssqlCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Namespace) DeepCopyInto(out *Namespace) {
	*out = *in
	in.NamespaceCheck.DeepCopyInto(&out.NamespaceCheck)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Namespace.
func (in *Namespace) DeepCopy() *Namespace {
	if in == nil {
		return nil
	}
	out := new(Namespace)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NamespaceCheck) DeepCopyInto(out *NamespaceCheck) {
	*out = *in
	if in.NamespaceLabels != nil {
		in, out := &in.NamespaceLabels, &out.NamespaceLabels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.NamespaceAnnotations != nil {
		in, out := &in.NamespaceAnnotations, &out.NamespaceAnnotations
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.ExpectedHTTPStatuses != nil {
		in, out := &in.ExpectedHTTPStatuses, &out.ExpectedHTTPStatuses
		*out = make([]int64, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NamespaceCheck.
func (in *NamespaceCheck) DeepCopy() *NamespaceCheck {
	if in == nil {
		return nil
	}
	out := new(NamespaceCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Pod) DeepCopyInto(out *Pod) {
	*out = *in
	in.PodCheck.DeepCopyInto(&out.PodCheck)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Pod.
func (in *Pod) DeepCopy() *Pod {
	if in == nil {
		return nil
	}
	out := new(Pod)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PodCheck) DeepCopyInto(out *PodCheck) {
	*out = *in
	if in.ExpectedHTTPStatuses != nil {
		in, out := &in.ExpectedHTTPStatuses, &out.ExpectedHTTPStatuses
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PodCheck.
func (in *PodCheck) DeepCopy() *PodCheck {
	if in == nil {
		return nil
	}
	out := new(PodCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Postgres) DeepCopyInto(out *Postgres) {
	*out = *in
	out.PostgresCheck = in.PostgresCheck
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Postgres.
func (in *Postgres) DeepCopy() *Postgres {
	if in == nil {
		return nil
	}
	out := new(Postgres)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *PostgresCheck) DeepCopyInto(out *PostgresCheck) {
	*out = *in
	out.SQLCheck = in.SQLCheck
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new PostgresCheck.
func (in *PostgresCheck) DeepCopy() *PostgresCheck {
	if in == nil {
		return nil
	}
	out := new(PostgresCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Redis) DeepCopyInto(out *Redis) {
	*out = *in
	out.RedisCheck = in.RedisCheck
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Redis.
func (in *Redis) DeepCopy() *Redis {
	if in == nil {
		return nil
	}
	out := new(Redis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *RedisCheck) DeepCopyInto(out *RedisCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RedisCheck.
func (in *RedisCheck) DeepCopy() *RedisCheck {
	if in == nil {
		return nil
	}
	out := new(RedisCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Restic) DeepCopyInto(out *Restic) {
	*out = *in
	out.ResticCheck = in.ResticCheck
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Restic.
func (in *Restic) DeepCopy() *Restic {
	if in == nil {
		return nil
	}
	out := new(Restic)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ResticCheck) DeepCopyInto(out *ResticCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ResticCheck.
func (in *ResticCheck) DeepCopy() *ResticCheck {
	if in == nil {
		return nil
	}
	out := new(ResticCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3) DeepCopyInto(out *S3) {
	*out = *in
	out.S3Check = in.S3Check
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3.
func (in *S3) DeepCopy() *S3 {
	if in == nil {
		return nil
	}
	out := new(S3)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Bucket) DeepCopyInto(out *S3Bucket) {
	*out = *in
	out.S3BucketCheck = in.S3BucketCheck
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Bucket.
func (in *S3Bucket) DeepCopy() *S3Bucket {
	if in == nil {
		return nil
	}
	out := new(S3Bucket)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3BucketCheck) DeepCopyInto(out *S3BucketCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3BucketCheck.
func (in *S3BucketCheck) DeepCopy() *S3BucketCheck {
	if in == nil {
		return nil
	}
	out := new(S3BucketCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *S3Check) DeepCopyInto(out *S3Check) {
	*out = *in
	out.Bucket = in.Bucket
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new S3Check.
func (in *S3Check) DeepCopy() *S3Check {
	if in == nil {
		return nil
	}
	out := new(S3Check)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SQLCheck) DeepCopyInto(out *SQLCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SQLCheck.
func (in *SQLCheck) DeepCopy() *SQLCheck {
	if in == nil {
		return nil
	}
	out := new(SQLCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SrvReply) DeepCopyInto(out *SrvReply) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SrvReply.
func (in *SrvReply) DeepCopy() *SrvReply {
	if in == nil {
		return nil
	}
	out := new(SrvReply)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCP) DeepCopyInto(out *TCP) {
	*out = *in
	out.TCPCheck = in.TCPCheck
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCP.
func (in *TCP) DeepCopy() *TCP {
	if in == nil {
		return nil
	}
	out := new(TCP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *TCPCheck) DeepCopyInto(out *TCPCheck) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TCPCheck.
func (in *TCPCheck) DeepCopy() *TCPCheck {
	if in == nil {
		return nil
	}
	out := new(TCPCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VarSource) DeepCopyInto(out *VarSource) {
	*out = *in
	if in.FieldRef != nil {
		in, out := &in.FieldRef, &out.FieldRef
		*out = new(corev1.ObjectFieldSelector)
		**out = **in
	}
	if in.ConfigMapKeyRef != nil {
		in, out := &in.ConfigMapKeyRef, &out.ConfigMapKeyRef
		*out = new(corev1.ConfigMapKeySelector)
		(*in).DeepCopyInto(*out)
	}
	if in.SecretKeyRef != nil {
		in, out := &in.SecretKeyRef, &out.SecretKeyRef
		*out = new(corev1.SecretKeySelector)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VarSource.
func (in *VarSource) DeepCopy() *VarSource {
	if in == nil {
		return nil
	}
	out := new(VarSource)
	in.DeepCopyInto(out)
	return out
}
