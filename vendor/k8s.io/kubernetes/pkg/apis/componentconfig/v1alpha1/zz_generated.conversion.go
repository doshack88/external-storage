// +build !ignore_autogenerated

/*
Copyright 2018 The Kubernetes Authors.

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

// This file was autogenerated by conversion-gen. Do not edit it manually!

package v1alpha1

import (
	unsafe "unsafe"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	componentconfig "k8s.io/kubernetes/pkg/apis/componentconfig"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(scheme *runtime.Scheme) error {
	return scheme.AddGeneratedConversionFuncs(
		Convert_v1alpha1_ClientConnectionConfiguration_To_componentconfig_ClientConnectionConfiguration,
		Convert_componentconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration,
		Convert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration,
		Convert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration,
		Convert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_componentconfig_KubeSchedulerLeaderElectionConfiguration,
		Convert_componentconfig_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration,
		Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration,
		Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration,
		Convert_v1alpha1_SchedulerAlgorithmSource_To_componentconfig_SchedulerAlgorithmSource,
		Convert_componentconfig_SchedulerAlgorithmSource_To_v1alpha1_SchedulerAlgorithmSource,
		Convert_v1alpha1_SchedulerPolicyConfigMapSource_To_componentconfig_SchedulerPolicyConfigMapSource,
		Convert_componentconfig_SchedulerPolicyConfigMapSource_To_v1alpha1_SchedulerPolicyConfigMapSource,
		Convert_v1alpha1_SchedulerPolicyFileSource_To_componentconfig_SchedulerPolicyFileSource,
		Convert_componentconfig_SchedulerPolicyFileSource_To_v1alpha1_SchedulerPolicyFileSource,
		Convert_v1alpha1_SchedulerPolicySource_To_componentconfig_SchedulerPolicySource,
		Convert_componentconfig_SchedulerPolicySource_To_v1alpha1_SchedulerPolicySource,
	)
}

func autoConvert_v1alpha1_ClientConnectionConfiguration_To_componentconfig_ClientConnectionConfiguration(in *ClientConnectionConfiguration, out *componentconfig.ClientConnectionConfiguration, s conversion.Scope) error {
	out.KubeConfigFile = in.KubeConfigFile
	out.AcceptContentTypes = in.AcceptContentTypes
	out.ContentType = in.ContentType
	out.QPS = in.QPS
	out.Burst = in.Burst
	return nil
}

// Convert_v1alpha1_ClientConnectionConfiguration_To_componentconfig_ClientConnectionConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_ClientConnectionConfiguration_To_componentconfig_ClientConnectionConfiguration(in *ClientConnectionConfiguration, out *componentconfig.ClientConnectionConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_ClientConnectionConfiguration_To_componentconfig_ClientConnectionConfiguration(in, out, s)
}

func autoConvert_componentconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(in *componentconfig.ClientConnectionConfiguration, out *ClientConnectionConfiguration, s conversion.Scope) error {
	out.KubeConfigFile = in.KubeConfigFile
	out.AcceptContentTypes = in.AcceptContentTypes
	out.ContentType = in.ContentType
	out.QPS = in.QPS
	out.Burst = in.Burst
	return nil
}

// Convert_componentconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration is an autogenerated conversion function.
func Convert_componentconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(in *componentconfig.ClientConnectionConfiguration, out *ClientConnectionConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(in, out, s)
}

func autoConvert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration(in *KubeSchedulerConfiguration, out *componentconfig.KubeSchedulerConfiguration, s conversion.Scope) error {
	out.SchedulerName = in.SchedulerName
	if err := Convert_v1alpha1_SchedulerAlgorithmSource_To_componentconfig_SchedulerAlgorithmSource(&in.AlgorithmSource, &out.AlgorithmSource, s); err != nil {
		return err
	}
	out.HardPodAffinitySymmetricWeight = in.HardPodAffinitySymmetricWeight
	if err := Convert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_componentconfig_KubeSchedulerLeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ClientConnectionConfiguration_To_componentconfig_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	out.HealthzBindAddress = in.HealthzBindAddress
	out.MetricsBindAddress = in.MetricsBindAddress
	out.EnableProfiling = in.EnableProfiling
	out.EnableContentionProfiling = in.EnableContentionProfiling
	out.FailureDomains = in.FailureDomains
	return nil
}

// Convert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration(in *KubeSchedulerConfiguration, out *componentconfig.KubeSchedulerConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeSchedulerConfiguration_To_componentconfig_KubeSchedulerConfiguration(in, out, s)
}

func autoConvert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in *componentconfig.KubeSchedulerConfiguration, out *KubeSchedulerConfiguration, s conversion.Scope) error {
	out.SchedulerName = in.SchedulerName
	if err := Convert_componentconfig_SchedulerAlgorithmSource_To_v1alpha1_SchedulerAlgorithmSource(&in.AlgorithmSource, &out.AlgorithmSource, s); err != nil {
		return err
	}
	out.HardPodAffinitySymmetricWeight = in.HardPodAffinitySymmetricWeight
	if err := Convert_componentconfig_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration(&in.LeaderElection, &out.LeaderElection, s); err != nil {
		return err
	}
	if err := Convert_componentconfig_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	out.HealthzBindAddress = in.HealthzBindAddress
	out.MetricsBindAddress = in.MetricsBindAddress
	out.EnableProfiling = in.EnableProfiling
	out.EnableContentionProfiling = in.EnableContentionProfiling
	out.FailureDomains = in.FailureDomains
	return nil
}

// Convert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration is an autogenerated conversion function.
func Convert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in *componentconfig.KubeSchedulerConfiguration, out *KubeSchedulerConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_KubeSchedulerConfiguration_To_v1alpha1_KubeSchedulerConfiguration(in, out, s)
}

func autoConvert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_componentconfig_KubeSchedulerLeaderElectionConfiguration(in *KubeSchedulerLeaderElectionConfiguration, out *componentconfig.KubeSchedulerLeaderElectionConfiguration, s conversion.Scope) error {
	if err := Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(&in.LeaderElectionConfiguration, &out.LeaderElectionConfiguration, s); err != nil {
		return err
	}
	out.LockObjectNamespace = in.LockObjectNamespace
	out.LockObjectName = in.LockObjectName
	return nil
}

// Convert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_componentconfig_KubeSchedulerLeaderElectionConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_componentconfig_KubeSchedulerLeaderElectionConfiguration(in *KubeSchedulerLeaderElectionConfiguration, out *componentconfig.KubeSchedulerLeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_KubeSchedulerLeaderElectionConfiguration_To_componentconfig_KubeSchedulerLeaderElectionConfiguration(in, out, s)
}

func autoConvert_componentconfig_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration(in *componentconfig.KubeSchedulerLeaderElectionConfiguration, out *KubeSchedulerLeaderElectionConfiguration, s conversion.Scope) error {
	if err := Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(&in.LeaderElectionConfiguration, &out.LeaderElectionConfiguration, s); err != nil {
		return err
	}
	out.LockObjectNamespace = in.LockObjectNamespace
	out.LockObjectName = in.LockObjectName
	return nil
}

// Convert_componentconfig_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration is an autogenerated conversion function.
func Convert_componentconfig_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration(in *componentconfig.KubeSchedulerLeaderElectionConfiguration, out *KubeSchedulerLeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_KubeSchedulerLeaderElectionConfiguration_To_v1alpha1_KubeSchedulerLeaderElectionConfiguration(in, out, s)
}

func autoConvert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in *LeaderElectionConfiguration, out *componentconfig.LeaderElectionConfiguration, s conversion.Scope) error {
	if err := v1.Convert_Pointer_bool_To_bool(&in.LeaderElect, &out.LeaderElect, s); err != nil {
		return err
	}
	out.LeaseDuration = in.LeaseDuration
	out.RenewDeadline = in.RenewDeadline
	out.RetryPeriod = in.RetryPeriod
	out.ResourceLock = in.ResourceLock
	return nil
}

// Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration is an autogenerated conversion function.
func Convert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in *LeaderElectionConfiguration, out *componentconfig.LeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha1_LeaderElectionConfiguration_To_componentconfig_LeaderElectionConfiguration(in, out, s)
}

func autoConvert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in *componentconfig.LeaderElectionConfiguration, out *LeaderElectionConfiguration, s conversion.Scope) error {
	if err := v1.Convert_bool_To_Pointer_bool(&in.LeaderElect, &out.LeaderElect, s); err != nil {
		return err
	}
	out.LeaseDuration = in.LeaseDuration
	out.RenewDeadline = in.RenewDeadline
	out.RetryPeriod = in.RetryPeriod
	out.ResourceLock = in.ResourceLock
	return nil
}

// Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration is an autogenerated conversion function.
func Convert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in *componentconfig.LeaderElectionConfiguration, out *LeaderElectionConfiguration, s conversion.Scope) error {
	return autoConvert_componentconfig_LeaderElectionConfiguration_To_v1alpha1_LeaderElectionConfiguration(in, out, s)
}

func autoConvert_v1alpha1_SchedulerAlgorithmSource_To_componentconfig_SchedulerAlgorithmSource(in *SchedulerAlgorithmSource, out *componentconfig.SchedulerAlgorithmSource, s conversion.Scope) error {
	out.Policy = (*componentconfig.SchedulerPolicySource)(unsafe.Pointer(in.Policy))
	out.Provider = (*string)(unsafe.Pointer(in.Provider))
	return nil
}

// Convert_v1alpha1_SchedulerAlgorithmSource_To_componentconfig_SchedulerAlgorithmSource is an autogenerated conversion function.
func Convert_v1alpha1_SchedulerAlgorithmSource_To_componentconfig_SchedulerAlgorithmSource(in *SchedulerAlgorithmSource, out *componentconfig.SchedulerAlgorithmSource, s conversion.Scope) error {
	return autoConvert_v1alpha1_SchedulerAlgorithmSource_To_componentconfig_SchedulerAlgorithmSource(in, out, s)
}

func autoConvert_componentconfig_SchedulerAlgorithmSource_To_v1alpha1_SchedulerAlgorithmSource(in *componentconfig.SchedulerAlgorithmSource, out *SchedulerAlgorithmSource, s conversion.Scope) error {
	out.Policy = (*SchedulerPolicySource)(unsafe.Pointer(in.Policy))
	out.Provider = (*string)(unsafe.Pointer(in.Provider))
	return nil
}

// Convert_componentconfig_SchedulerAlgorithmSource_To_v1alpha1_SchedulerAlgorithmSource is an autogenerated conversion function.
func Convert_componentconfig_SchedulerAlgorithmSource_To_v1alpha1_SchedulerAlgorithmSource(in *componentconfig.SchedulerAlgorithmSource, out *SchedulerAlgorithmSource, s conversion.Scope) error {
	return autoConvert_componentconfig_SchedulerAlgorithmSource_To_v1alpha1_SchedulerAlgorithmSource(in, out, s)
}

func autoConvert_v1alpha1_SchedulerPolicyConfigMapSource_To_componentconfig_SchedulerPolicyConfigMapSource(in *SchedulerPolicyConfigMapSource, out *componentconfig.SchedulerPolicyConfigMapSource, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	return nil
}

// Convert_v1alpha1_SchedulerPolicyConfigMapSource_To_componentconfig_SchedulerPolicyConfigMapSource is an autogenerated conversion function.
func Convert_v1alpha1_SchedulerPolicyConfigMapSource_To_componentconfig_SchedulerPolicyConfigMapSource(in *SchedulerPolicyConfigMapSource, out *componentconfig.SchedulerPolicyConfigMapSource, s conversion.Scope) error {
	return autoConvert_v1alpha1_SchedulerPolicyConfigMapSource_To_componentconfig_SchedulerPolicyConfigMapSource(in, out, s)
}

func autoConvert_componentconfig_SchedulerPolicyConfigMapSource_To_v1alpha1_SchedulerPolicyConfigMapSource(in *componentconfig.SchedulerPolicyConfigMapSource, out *SchedulerPolicyConfigMapSource, s conversion.Scope) error {
	out.Namespace = in.Namespace
	out.Name = in.Name
	return nil
}

// Convert_componentconfig_SchedulerPolicyConfigMapSource_To_v1alpha1_SchedulerPolicyConfigMapSource is an autogenerated conversion function.
func Convert_componentconfig_SchedulerPolicyConfigMapSource_To_v1alpha1_SchedulerPolicyConfigMapSource(in *componentconfig.SchedulerPolicyConfigMapSource, out *SchedulerPolicyConfigMapSource, s conversion.Scope) error {
	return autoConvert_componentconfig_SchedulerPolicyConfigMapSource_To_v1alpha1_SchedulerPolicyConfigMapSource(in, out, s)
}

func autoConvert_v1alpha1_SchedulerPolicyFileSource_To_componentconfig_SchedulerPolicyFileSource(in *SchedulerPolicyFileSource, out *componentconfig.SchedulerPolicyFileSource, s conversion.Scope) error {
	out.Path = in.Path
	return nil
}

// Convert_v1alpha1_SchedulerPolicyFileSource_To_componentconfig_SchedulerPolicyFileSource is an autogenerated conversion function.
func Convert_v1alpha1_SchedulerPolicyFileSource_To_componentconfig_SchedulerPolicyFileSource(in *SchedulerPolicyFileSource, out *componentconfig.SchedulerPolicyFileSource, s conversion.Scope) error {
	return autoConvert_v1alpha1_SchedulerPolicyFileSource_To_componentconfig_SchedulerPolicyFileSource(in, out, s)
}

func autoConvert_componentconfig_SchedulerPolicyFileSource_To_v1alpha1_SchedulerPolicyFileSource(in *componentconfig.SchedulerPolicyFileSource, out *SchedulerPolicyFileSource, s conversion.Scope) error {
	out.Path = in.Path
	return nil
}

// Convert_componentconfig_SchedulerPolicyFileSource_To_v1alpha1_SchedulerPolicyFileSource is an autogenerated conversion function.
func Convert_componentconfig_SchedulerPolicyFileSource_To_v1alpha1_SchedulerPolicyFileSource(in *componentconfig.SchedulerPolicyFileSource, out *SchedulerPolicyFileSource, s conversion.Scope) error {
	return autoConvert_componentconfig_SchedulerPolicyFileSource_To_v1alpha1_SchedulerPolicyFileSource(in, out, s)
}

func autoConvert_v1alpha1_SchedulerPolicySource_To_componentconfig_SchedulerPolicySource(in *SchedulerPolicySource, out *componentconfig.SchedulerPolicySource, s conversion.Scope) error {
	out.File = (*componentconfig.SchedulerPolicyFileSource)(unsafe.Pointer(in.File))
	out.ConfigMap = (*componentconfig.SchedulerPolicyConfigMapSource)(unsafe.Pointer(in.ConfigMap))
	return nil
}

// Convert_v1alpha1_SchedulerPolicySource_To_componentconfig_SchedulerPolicySource is an autogenerated conversion function.
func Convert_v1alpha1_SchedulerPolicySource_To_componentconfig_SchedulerPolicySource(in *SchedulerPolicySource, out *componentconfig.SchedulerPolicySource, s conversion.Scope) error {
	return autoConvert_v1alpha1_SchedulerPolicySource_To_componentconfig_SchedulerPolicySource(in, out, s)
}

func autoConvert_componentconfig_SchedulerPolicySource_To_v1alpha1_SchedulerPolicySource(in *componentconfig.SchedulerPolicySource, out *SchedulerPolicySource, s conversion.Scope) error {
	out.File = (*SchedulerPolicyFileSource)(unsafe.Pointer(in.File))
	out.ConfigMap = (*SchedulerPolicyConfigMapSource)(unsafe.Pointer(in.ConfigMap))
	return nil
}

// Convert_componentconfig_SchedulerPolicySource_To_v1alpha1_SchedulerPolicySource is an autogenerated conversion function.
func Convert_componentconfig_SchedulerPolicySource_To_v1alpha1_SchedulerPolicySource(in *componentconfig.SchedulerPolicySource, out *SchedulerPolicySource, s conversion.Scope) error {
	return autoConvert_componentconfig_SchedulerPolicySource_To_v1alpha1_SchedulerPolicySource(in, out, s)
}