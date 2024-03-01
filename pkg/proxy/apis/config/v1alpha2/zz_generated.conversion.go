//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha2

import (
	unsafe "unsafe"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	v1alpha1 "k8s.io/component-base/config/v1alpha1"
	v1alpha2 "k8s.io/kube-proxy/config/v1alpha2"
	config "k8s.io/kubernetes/pkg/proxy/apis/config"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*v1alpha2.DetectLocalConfiguration)(nil), (*config.DetectLocalConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_DetectLocalConfiguration_To_config_DetectLocalConfiguration(a.(*v1alpha2.DetectLocalConfiguration), b.(*config.DetectLocalConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.DetectLocalConfiguration)(nil), (*v1alpha2.DetectLocalConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_DetectLocalConfiguration_To_v1alpha2_DetectLocalConfiguration(a.(*config.DetectLocalConfiguration), b.(*v1alpha2.DetectLocalConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha2.KubeProxyConfiguration)(nil), (*config.KubeProxyConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_KubeProxyConfiguration_To_config_KubeProxyConfiguration(a.(*v1alpha2.KubeProxyConfiguration), b.(*config.KubeProxyConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeProxyConfiguration)(nil), (*v1alpha2.KubeProxyConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeProxyConfiguration_To_v1alpha2_KubeProxyConfiguration(a.(*config.KubeProxyConfiguration), b.(*v1alpha2.KubeProxyConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha2.KubeProxyConntrackConfiguration)(nil), (*config.KubeProxyConntrackConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_KubeProxyConntrackConfiguration_To_config_KubeProxyConntrackConfiguration(a.(*v1alpha2.KubeProxyConntrackConfiguration), b.(*config.KubeProxyConntrackConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeProxyConntrackConfiguration)(nil), (*v1alpha2.KubeProxyConntrackConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeProxyConntrackConfiguration_To_v1alpha2_KubeProxyConntrackConfiguration(a.(*config.KubeProxyConntrackConfiguration), b.(*v1alpha2.KubeProxyConntrackConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha2.KubeProxyIPTablesConfiguration)(nil), (*config.KubeProxyIPTablesConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_KubeProxyIPTablesConfiguration_To_config_KubeProxyIPTablesConfiguration(a.(*v1alpha2.KubeProxyIPTablesConfiguration), b.(*config.KubeProxyIPTablesConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeProxyIPTablesConfiguration)(nil), (*v1alpha2.KubeProxyIPTablesConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeProxyIPTablesConfiguration_To_v1alpha2_KubeProxyIPTablesConfiguration(a.(*config.KubeProxyIPTablesConfiguration), b.(*v1alpha2.KubeProxyIPTablesConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha2.KubeProxyIPVSConfiguration)(nil), (*config.KubeProxyIPVSConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_KubeProxyIPVSConfiguration_To_config_KubeProxyIPVSConfiguration(a.(*v1alpha2.KubeProxyIPVSConfiguration), b.(*config.KubeProxyIPVSConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeProxyIPVSConfiguration)(nil), (*v1alpha2.KubeProxyIPVSConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeProxyIPVSConfiguration_To_v1alpha2_KubeProxyIPVSConfiguration(a.(*config.KubeProxyIPVSConfiguration), b.(*v1alpha2.KubeProxyIPVSConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha2.KubeProxyLinuxConfiguration)(nil), (*config.KubeProxyLinuxConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_KubeProxyLinuxConfiguration_To_config_KubeProxyLinuxConfiguration(a.(*v1alpha2.KubeProxyLinuxConfiguration), b.(*config.KubeProxyLinuxConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeProxyLinuxConfiguration)(nil), (*v1alpha2.KubeProxyLinuxConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeProxyLinuxConfiguration_To_v1alpha2_KubeProxyLinuxConfiguration(a.(*config.KubeProxyLinuxConfiguration), b.(*v1alpha2.KubeProxyLinuxConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha2.KubeProxyNFTablesConfiguration)(nil), (*config.KubeProxyNFTablesConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_KubeProxyNFTablesConfiguration_To_config_KubeProxyNFTablesConfiguration(a.(*v1alpha2.KubeProxyNFTablesConfiguration), b.(*config.KubeProxyNFTablesConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeProxyNFTablesConfiguration)(nil), (*v1alpha2.KubeProxyNFTablesConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeProxyNFTablesConfiguration_To_v1alpha2_KubeProxyNFTablesConfiguration(a.(*config.KubeProxyNFTablesConfiguration), b.(*v1alpha2.KubeProxyNFTablesConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha2.KubeProxyWindowsConfiguration)(nil), (*config.KubeProxyWindowsConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_KubeProxyWindowsConfiguration_To_config_KubeProxyWindowsConfiguration(a.(*v1alpha2.KubeProxyWindowsConfiguration), b.(*config.KubeProxyWindowsConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeProxyWindowsConfiguration)(nil), (*v1alpha2.KubeProxyWindowsConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeProxyWindowsConfiguration_To_v1alpha2_KubeProxyWindowsConfiguration(a.(*config.KubeProxyWindowsConfiguration), b.(*v1alpha2.KubeProxyWindowsConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*v1alpha2.KubeProxyWinkernelConfiguration)(nil), (*config.KubeProxyWinkernelConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha2_KubeProxyWinkernelConfiguration_To_config_KubeProxyWinkernelConfiguration(a.(*v1alpha2.KubeProxyWinkernelConfiguration), b.(*config.KubeProxyWinkernelConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.KubeProxyWinkernelConfiguration)(nil), (*v1alpha2.KubeProxyWinkernelConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_KubeProxyWinkernelConfiguration_To_v1alpha2_KubeProxyWinkernelConfiguration(a.(*config.KubeProxyWinkernelConfiguration), b.(*v1alpha2.KubeProxyWinkernelConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha2_DetectLocalConfiguration_To_config_DetectLocalConfiguration(in *v1alpha2.DetectLocalConfiguration, out *config.DetectLocalConfiguration, s conversion.Scope) error {
	out.BridgeInterface = in.BridgeInterface
	out.InterfaceNamePrefix = in.InterfaceNamePrefix
	out.ClusterCIDRs = *(*[]string)(unsafe.Pointer(&in.ClusterCIDRs))
	return nil
}

// Convert_v1alpha2_DetectLocalConfiguration_To_config_DetectLocalConfiguration is an autogenerated conversion function.
func Convert_v1alpha2_DetectLocalConfiguration_To_config_DetectLocalConfiguration(in *v1alpha2.DetectLocalConfiguration, out *config.DetectLocalConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha2_DetectLocalConfiguration_To_config_DetectLocalConfiguration(in, out, s)
}

func autoConvert_config_DetectLocalConfiguration_To_v1alpha2_DetectLocalConfiguration(in *config.DetectLocalConfiguration, out *v1alpha2.DetectLocalConfiguration, s conversion.Scope) error {
	out.BridgeInterface = in.BridgeInterface
	out.InterfaceNamePrefix = in.InterfaceNamePrefix
	out.ClusterCIDRs = *(*[]string)(unsafe.Pointer(&in.ClusterCIDRs))
	return nil
}

// Convert_config_DetectLocalConfiguration_To_v1alpha2_DetectLocalConfiguration is an autogenerated conversion function.
func Convert_config_DetectLocalConfiguration_To_v1alpha2_DetectLocalConfiguration(in *config.DetectLocalConfiguration, out *v1alpha2.DetectLocalConfiguration, s conversion.Scope) error {
	return autoConvert_config_DetectLocalConfiguration_To_v1alpha2_DetectLocalConfiguration(in, out, s)
}

func autoConvert_v1alpha2_KubeProxyConfiguration_To_config_KubeProxyConfiguration(in *v1alpha2.KubeProxyConfiguration, out *config.KubeProxyConfiguration, s conversion.Scope) error {
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	if err := v1alpha1.Convert_v1alpha1_ClientConnectionConfiguration_To_config_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	out.Logging = in.Logging
	out.HostnameOverride = in.HostnameOverride
	out.NodeIPOverride = *(*[]string)(unsafe.Pointer(&in.NodeIPOverride))
	out.HealthzBindAddresses = *(*[]string)(unsafe.Pointer(&in.HealthzBindAddresses))
	out.HealthzBindPort = in.HealthzBindPort
	out.MetricsBindAddresses = *(*[]string)(unsafe.Pointer(&in.MetricsBindAddresses))
	out.MetricsBindPort = in.MetricsBindPort
	out.ConfigHardFail = (*bool)(unsafe.Pointer(in.ConfigHardFail))
	out.BindAddressHardFail = in.BindAddressHardFail
	out.EnableProfiling = in.EnableProfiling
	out.ShowHiddenMetricsForVersion = in.ShowHiddenMetricsForVersion
	out.Mode = config.ProxyMode(in.Mode)
	if err := Convert_v1alpha2_KubeProxyLinuxConfiguration_To_config_KubeProxyLinuxConfiguration(&in.Linux, &out.Linux, s); err != nil {
		return err
	}
	if err := Convert_v1alpha2_KubeProxyWindowsConfiguration_To_config_KubeProxyWindowsConfiguration(&in.Windows, &out.Windows, s); err != nil {
		return err
	}
	if err := Convert_v1alpha2_KubeProxyIPTablesConfiguration_To_config_KubeProxyIPTablesConfiguration(&in.IPTables, &out.IPTables, s); err != nil {
		return err
	}
	if err := Convert_v1alpha2_KubeProxyIPVSConfiguration_To_config_KubeProxyIPVSConfiguration(&in.IPVS, &out.IPVS, s); err != nil {
		return err
	}
	if err := Convert_v1alpha2_KubeProxyNFTablesConfiguration_To_config_KubeProxyNFTablesConfiguration(&in.NFTables, &out.NFTables, s); err != nil {
		return err
	}
	if err := Convert_v1alpha2_KubeProxyWinkernelConfiguration_To_config_KubeProxyWinkernelConfiguration(&in.Winkernel, &out.Winkernel, s); err != nil {
		return err
	}
	out.DetectLocalMode = config.LocalMode(in.DetectLocalMode)
	if err := Convert_v1alpha2_DetectLocalConfiguration_To_config_DetectLocalConfiguration(&in.DetectLocal, &out.DetectLocal, s); err != nil {
		return err
	}
	out.NodePortAddresses = *(*[]string)(unsafe.Pointer(&in.NodePortAddresses))
	out.SyncPeriod = in.SyncPeriod
	out.MinSyncPeriod = in.MinSyncPeriod
	out.ConfigSyncPeriod = in.ConfigSyncPeriod
	return nil
}

// Convert_v1alpha2_KubeProxyConfiguration_To_config_KubeProxyConfiguration is an autogenerated conversion function.
func Convert_v1alpha2_KubeProxyConfiguration_To_config_KubeProxyConfiguration(in *v1alpha2.KubeProxyConfiguration, out *config.KubeProxyConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha2_KubeProxyConfiguration_To_config_KubeProxyConfiguration(in, out, s)
}

func autoConvert_config_KubeProxyConfiguration_To_v1alpha2_KubeProxyConfiguration(in *config.KubeProxyConfiguration, out *v1alpha2.KubeProxyConfiguration, s conversion.Scope) error {
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	if err := v1alpha1.Convert_config_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(&in.ClientConnection, &out.ClientConnection, s); err != nil {
		return err
	}
	out.Logging = in.Logging
	out.HostnameOverride = in.HostnameOverride
	out.NodeIPOverride = *(*[]string)(unsafe.Pointer(&in.NodeIPOverride))
	out.HealthzBindAddresses = *(*[]string)(unsafe.Pointer(&in.HealthzBindAddresses))
	out.HealthzBindPort = in.HealthzBindPort
	out.MetricsBindAddresses = *(*[]string)(unsafe.Pointer(&in.MetricsBindAddresses))
	out.MetricsBindPort = in.MetricsBindPort
	out.BindAddressHardFail = in.BindAddressHardFail
	out.ConfigHardFail = (*bool)(unsafe.Pointer(in.ConfigHardFail))
	out.EnableProfiling = in.EnableProfiling
	out.ShowHiddenMetricsForVersion = in.ShowHiddenMetricsForVersion
	out.Mode = v1alpha2.ProxyMode(in.Mode)
	if err := Convert_config_KubeProxyLinuxConfiguration_To_v1alpha2_KubeProxyLinuxConfiguration(&in.Linux, &out.Linux, s); err != nil {
		return err
	}
	if err := Convert_config_KubeProxyWindowsConfiguration_To_v1alpha2_KubeProxyWindowsConfiguration(&in.Windows, &out.Windows, s); err != nil {
		return err
	}
	if err := Convert_config_KubeProxyIPTablesConfiguration_To_v1alpha2_KubeProxyIPTablesConfiguration(&in.IPTables, &out.IPTables, s); err != nil {
		return err
	}
	if err := Convert_config_KubeProxyIPVSConfiguration_To_v1alpha2_KubeProxyIPVSConfiguration(&in.IPVS, &out.IPVS, s); err != nil {
		return err
	}
	if err := Convert_config_KubeProxyWinkernelConfiguration_To_v1alpha2_KubeProxyWinkernelConfiguration(&in.Winkernel, &out.Winkernel, s); err != nil {
		return err
	}
	if err := Convert_config_KubeProxyNFTablesConfiguration_To_v1alpha2_KubeProxyNFTablesConfiguration(&in.NFTables, &out.NFTables, s); err != nil {
		return err
	}
	out.DetectLocalMode = v1alpha2.LocalMode(in.DetectLocalMode)
	if err := Convert_config_DetectLocalConfiguration_To_v1alpha2_DetectLocalConfiguration(&in.DetectLocal, &out.DetectLocal, s); err != nil {
		return err
	}
	out.NodePortAddresses = *(*[]string)(unsafe.Pointer(&in.NodePortAddresses))
	out.SyncPeriod = in.SyncPeriod
	out.MinSyncPeriod = in.MinSyncPeriod
	out.ConfigSyncPeriod = in.ConfigSyncPeriod
	return nil
}

// Convert_config_KubeProxyConfiguration_To_v1alpha2_KubeProxyConfiguration is an autogenerated conversion function.
func Convert_config_KubeProxyConfiguration_To_v1alpha2_KubeProxyConfiguration(in *config.KubeProxyConfiguration, out *v1alpha2.KubeProxyConfiguration, s conversion.Scope) error {
	return autoConvert_config_KubeProxyConfiguration_To_v1alpha2_KubeProxyConfiguration(in, out, s)
}

func autoConvert_v1alpha2_KubeProxyConntrackConfiguration_To_config_KubeProxyConntrackConfiguration(in *v1alpha2.KubeProxyConntrackConfiguration, out *config.KubeProxyConntrackConfiguration, s conversion.Scope) error {
	out.MaxPerCore = (*int32)(unsafe.Pointer(in.MaxPerCore))
	out.Min = (*int32)(unsafe.Pointer(in.Min))
	out.TCPEstablishedTimeout = (*v1.Duration)(unsafe.Pointer(in.TCPEstablishedTimeout))
	out.TCPCloseWaitTimeout = (*v1.Duration)(unsafe.Pointer(in.TCPCloseWaitTimeout))
	out.TCPBeLiberal = in.TCPBeLiberal
	out.UDPTimeout = in.UDPTimeout
	out.UDPStreamTimeout = in.UDPStreamTimeout
	return nil
}

// Convert_v1alpha2_KubeProxyConntrackConfiguration_To_config_KubeProxyConntrackConfiguration is an autogenerated conversion function.
func Convert_v1alpha2_KubeProxyConntrackConfiguration_To_config_KubeProxyConntrackConfiguration(in *v1alpha2.KubeProxyConntrackConfiguration, out *config.KubeProxyConntrackConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha2_KubeProxyConntrackConfiguration_To_config_KubeProxyConntrackConfiguration(in, out, s)
}

func autoConvert_config_KubeProxyConntrackConfiguration_To_v1alpha2_KubeProxyConntrackConfiguration(in *config.KubeProxyConntrackConfiguration, out *v1alpha2.KubeProxyConntrackConfiguration, s conversion.Scope) error {
	out.MaxPerCore = (*int32)(unsafe.Pointer(in.MaxPerCore))
	out.Min = (*int32)(unsafe.Pointer(in.Min))
	out.TCPEstablishedTimeout = (*v1.Duration)(unsafe.Pointer(in.TCPEstablishedTimeout))
	out.TCPCloseWaitTimeout = (*v1.Duration)(unsafe.Pointer(in.TCPCloseWaitTimeout))
	out.TCPBeLiberal = in.TCPBeLiberal
	out.UDPTimeout = in.UDPTimeout
	out.UDPStreamTimeout = in.UDPStreamTimeout
	return nil
}

// Convert_config_KubeProxyConntrackConfiguration_To_v1alpha2_KubeProxyConntrackConfiguration is an autogenerated conversion function.
func Convert_config_KubeProxyConntrackConfiguration_To_v1alpha2_KubeProxyConntrackConfiguration(in *config.KubeProxyConntrackConfiguration, out *v1alpha2.KubeProxyConntrackConfiguration, s conversion.Scope) error {
	return autoConvert_config_KubeProxyConntrackConfiguration_To_v1alpha2_KubeProxyConntrackConfiguration(in, out, s)
}

func autoConvert_v1alpha2_KubeProxyIPTablesConfiguration_To_config_KubeProxyIPTablesConfiguration(in *v1alpha2.KubeProxyIPTablesConfiguration, out *config.KubeProxyIPTablesConfiguration, s conversion.Scope) error {
	out.MasqueradeBit = (*int32)(unsafe.Pointer(in.MasqueradeBit))
	out.LocalhostNodePorts = (*bool)(unsafe.Pointer(in.LocalhostNodePorts))
	return nil
}

// Convert_v1alpha2_KubeProxyIPTablesConfiguration_To_config_KubeProxyIPTablesConfiguration is an autogenerated conversion function.
func Convert_v1alpha2_KubeProxyIPTablesConfiguration_To_config_KubeProxyIPTablesConfiguration(in *v1alpha2.KubeProxyIPTablesConfiguration, out *config.KubeProxyIPTablesConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha2_KubeProxyIPTablesConfiguration_To_config_KubeProxyIPTablesConfiguration(in, out, s)
}

func autoConvert_config_KubeProxyIPTablesConfiguration_To_v1alpha2_KubeProxyIPTablesConfiguration(in *config.KubeProxyIPTablesConfiguration, out *v1alpha2.KubeProxyIPTablesConfiguration, s conversion.Scope) error {
	out.MasqueradeBit = (*int32)(unsafe.Pointer(in.MasqueradeBit))
	out.LocalhostNodePorts = (*bool)(unsafe.Pointer(in.LocalhostNodePorts))
	return nil
}

// Convert_config_KubeProxyIPTablesConfiguration_To_v1alpha2_KubeProxyIPTablesConfiguration is an autogenerated conversion function.
func Convert_config_KubeProxyIPTablesConfiguration_To_v1alpha2_KubeProxyIPTablesConfiguration(in *config.KubeProxyIPTablesConfiguration, out *v1alpha2.KubeProxyIPTablesConfiguration, s conversion.Scope) error {
	return autoConvert_config_KubeProxyIPTablesConfiguration_To_v1alpha2_KubeProxyIPTablesConfiguration(in, out, s)
}

func autoConvert_v1alpha2_KubeProxyIPVSConfiguration_To_config_KubeProxyIPVSConfiguration(in *v1alpha2.KubeProxyIPVSConfiguration, out *config.KubeProxyIPVSConfiguration, s conversion.Scope) error {
	out.MasqueradeBit = (*int32)(unsafe.Pointer(in.MasqueradeBit))
	out.Scheduler = in.Scheduler
	out.ExcludeCIDRs = *(*[]string)(unsafe.Pointer(&in.ExcludeCIDRs))
	out.StrictARP = in.StrictARP
	out.TCPTimeout = in.TCPTimeout
	out.TCPFinTimeout = in.TCPFinTimeout
	out.UDPTimeout = in.UDPTimeout
	return nil
}

// Convert_v1alpha2_KubeProxyIPVSConfiguration_To_config_KubeProxyIPVSConfiguration is an autogenerated conversion function.
func Convert_v1alpha2_KubeProxyIPVSConfiguration_To_config_KubeProxyIPVSConfiguration(in *v1alpha2.KubeProxyIPVSConfiguration, out *config.KubeProxyIPVSConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha2_KubeProxyIPVSConfiguration_To_config_KubeProxyIPVSConfiguration(in, out, s)
}

func autoConvert_config_KubeProxyIPVSConfiguration_To_v1alpha2_KubeProxyIPVSConfiguration(in *config.KubeProxyIPVSConfiguration, out *v1alpha2.KubeProxyIPVSConfiguration, s conversion.Scope) error {
	out.MasqueradeBit = (*int32)(unsafe.Pointer(in.MasqueradeBit))
	out.Scheduler = in.Scheduler
	out.ExcludeCIDRs = *(*[]string)(unsafe.Pointer(&in.ExcludeCIDRs))
	out.StrictARP = in.StrictARP
	out.TCPTimeout = in.TCPTimeout
	out.TCPFinTimeout = in.TCPFinTimeout
	out.UDPTimeout = in.UDPTimeout
	return nil
}

// Convert_config_KubeProxyIPVSConfiguration_To_v1alpha2_KubeProxyIPVSConfiguration is an autogenerated conversion function.
func Convert_config_KubeProxyIPVSConfiguration_To_v1alpha2_KubeProxyIPVSConfiguration(in *config.KubeProxyIPVSConfiguration, out *v1alpha2.KubeProxyIPVSConfiguration, s conversion.Scope) error {
	return autoConvert_config_KubeProxyIPVSConfiguration_To_v1alpha2_KubeProxyIPVSConfiguration(in, out, s)
}

func autoConvert_v1alpha2_KubeProxyLinuxConfiguration_To_config_KubeProxyLinuxConfiguration(in *v1alpha2.KubeProxyLinuxConfiguration, out *config.KubeProxyLinuxConfiguration, s conversion.Scope) error {
	out.OOMScoreAdj = (*int32)(unsafe.Pointer(in.OOMScoreAdj))
	if err := Convert_v1alpha2_KubeProxyConntrackConfiguration_To_config_KubeProxyConntrackConfiguration(&in.Conntrack, &out.Conntrack, s); err != nil {
		return err
	}
	out.MasqueradeAll = in.MasqueradeAll
	return nil
}

// Convert_v1alpha2_KubeProxyLinuxConfiguration_To_config_KubeProxyLinuxConfiguration is an autogenerated conversion function.
func Convert_v1alpha2_KubeProxyLinuxConfiguration_To_config_KubeProxyLinuxConfiguration(in *v1alpha2.KubeProxyLinuxConfiguration, out *config.KubeProxyLinuxConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha2_KubeProxyLinuxConfiguration_To_config_KubeProxyLinuxConfiguration(in, out, s)
}

func autoConvert_config_KubeProxyLinuxConfiguration_To_v1alpha2_KubeProxyLinuxConfiguration(in *config.KubeProxyLinuxConfiguration, out *v1alpha2.KubeProxyLinuxConfiguration, s conversion.Scope) error {
	out.OOMScoreAdj = (*int32)(unsafe.Pointer(in.OOMScoreAdj))
	if err := Convert_config_KubeProxyConntrackConfiguration_To_v1alpha2_KubeProxyConntrackConfiguration(&in.Conntrack, &out.Conntrack, s); err != nil {
		return err
	}
	out.MasqueradeAll = in.MasqueradeAll
	return nil
}

// Convert_config_KubeProxyLinuxConfiguration_To_v1alpha2_KubeProxyLinuxConfiguration is an autogenerated conversion function.
func Convert_config_KubeProxyLinuxConfiguration_To_v1alpha2_KubeProxyLinuxConfiguration(in *config.KubeProxyLinuxConfiguration, out *v1alpha2.KubeProxyLinuxConfiguration, s conversion.Scope) error {
	return autoConvert_config_KubeProxyLinuxConfiguration_To_v1alpha2_KubeProxyLinuxConfiguration(in, out, s)
}

func autoConvert_v1alpha2_KubeProxyNFTablesConfiguration_To_config_KubeProxyNFTablesConfiguration(in *v1alpha2.KubeProxyNFTablesConfiguration, out *config.KubeProxyNFTablesConfiguration, s conversion.Scope) error {
	out.MasqueradeBit = (*int32)(unsafe.Pointer(in.MasqueradeBit))
	return nil
}

// Convert_v1alpha2_KubeProxyNFTablesConfiguration_To_config_KubeProxyNFTablesConfiguration is an autogenerated conversion function.
func Convert_v1alpha2_KubeProxyNFTablesConfiguration_To_config_KubeProxyNFTablesConfiguration(in *v1alpha2.KubeProxyNFTablesConfiguration, out *config.KubeProxyNFTablesConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha2_KubeProxyNFTablesConfiguration_To_config_KubeProxyNFTablesConfiguration(in, out, s)
}

func autoConvert_config_KubeProxyNFTablesConfiguration_To_v1alpha2_KubeProxyNFTablesConfiguration(in *config.KubeProxyNFTablesConfiguration, out *v1alpha2.KubeProxyNFTablesConfiguration, s conversion.Scope) error {
	out.MasqueradeBit = (*int32)(unsafe.Pointer(in.MasqueradeBit))
	return nil
}

// Convert_config_KubeProxyNFTablesConfiguration_To_v1alpha2_KubeProxyNFTablesConfiguration is an autogenerated conversion function.
func Convert_config_KubeProxyNFTablesConfiguration_To_v1alpha2_KubeProxyNFTablesConfiguration(in *config.KubeProxyNFTablesConfiguration, out *v1alpha2.KubeProxyNFTablesConfiguration, s conversion.Scope) error {
	return autoConvert_config_KubeProxyNFTablesConfiguration_To_v1alpha2_KubeProxyNFTablesConfiguration(in, out, s)
}

func autoConvert_v1alpha2_KubeProxyWindowsConfiguration_To_config_KubeProxyWindowsConfiguration(in *v1alpha2.KubeProxyWindowsConfiguration, out *config.KubeProxyWindowsConfiguration, s conversion.Scope) error {
	out.RunAsService = in.RunAsService
	return nil
}

// Convert_v1alpha2_KubeProxyWindowsConfiguration_To_config_KubeProxyWindowsConfiguration is an autogenerated conversion function.
func Convert_v1alpha2_KubeProxyWindowsConfiguration_To_config_KubeProxyWindowsConfiguration(in *v1alpha2.KubeProxyWindowsConfiguration, out *config.KubeProxyWindowsConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha2_KubeProxyWindowsConfiguration_To_config_KubeProxyWindowsConfiguration(in, out, s)
}

func autoConvert_config_KubeProxyWindowsConfiguration_To_v1alpha2_KubeProxyWindowsConfiguration(in *config.KubeProxyWindowsConfiguration, out *v1alpha2.KubeProxyWindowsConfiguration, s conversion.Scope) error {
	out.RunAsService = in.RunAsService
	return nil
}

// Convert_config_KubeProxyWindowsConfiguration_To_v1alpha2_KubeProxyWindowsConfiguration is an autogenerated conversion function.
func Convert_config_KubeProxyWindowsConfiguration_To_v1alpha2_KubeProxyWindowsConfiguration(in *config.KubeProxyWindowsConfiguration, out *v1alpha2.KubeProxyWindowsConfiguration, s conversion.Scope) error {
	return autoConvert_config_KubeProxyWindowsConfiguration_To_v1alpha2_KubeProxyWindowsConfiguration(in, out, s)
}

func autoConvert_v1alpha2_KubeProxyWinkernelConfiguration_To_config_KubeProxyWinkernelConfiguration(in *v1alpha2.KubeProxyWinkernelConfiguration, out *config.KubeProxyWinkernelConfiguration, s conversion.Scope) error {
	out.NetworkName = in.NetworkName
	out.SourceVip = in.SourceVip
	out.EnableDSR = in.EnableDSR
	out.RootHnsEndpointName = in.RootHnsEndpointName
	out.ForwardHealthCheckVip = in.ForwardHealthCheckVip
	return nil
}

// Convert_v1alpha2_KubeProxyWinkernelConfiguration_To_config_KubeProxyWinkernelConfiguration is an autogenerated conversion function.
func Convert_v1alpha2_KubeProxyWinkernelConfiguration_To_config_KubeProxyWinkernelConfiguration(in *v1alpha2.KubeProxyWinkernelConfiguration, out *config.KubeProxyWinkernelConfiguration, s conversion.Scope) error {
	return autoConvert_v1alpha2_KubeProxyWinkernelConfiguration_To_config_KubeProxyWinkernelConfiguration(in, out, s)
}

func autoConvert_config_KubeProxyWinkernelConfiguration_To_v1alpha2_KubeProxyWinkernelConfiguration(in *config.KubeProxyWinkernelConfiguration, out *v1alpha2.KubeProxyWinkernelConfiguration, s conversion.Scope) error {
	out.NetworkName = in.NetworkName
	out.SourceVip = in.SourceVip
	out.EnableDSR = in.EnableDSR
	out.RootHnsEndpointName = in.RootHnsEndpointName
	out.ForwardHealthCheckVip = in.ForwardHealthCheckVip
	return nil
}

// Convert_config_KubeProxyWinkernelConfiguration_To_v1alpha2_KubeProxyWinkernelConfiguration is an autogenerated conversion function.
func Convert_config_KubeProxyWinkernelConfiguration_To_v1alpha2_KubeProxyWinkernelConfiguration(in *config.KubeProxyWinkernelConfiguration, out *v1alpha2.KubeProxyWinkernelConfiguration, s conversion.Scope) error {
	return autoConvert_config_KubeProxyWinkernelConfiguration_To_v1alpha2_KubeProxyWinkernelConfiguration(in, out, s)
}
