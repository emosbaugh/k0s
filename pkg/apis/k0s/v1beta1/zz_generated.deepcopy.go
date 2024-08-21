//go:build !ignore_autogenerated

/*
Copyright k0s authors

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
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *APISpec) DeepCopyInto(out *APISpec) {
	*out = *in
	if in.ExtraArgs != nil {
		in, out := &in.ExtraArgs, &out.ExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
	if in.SANs != nil {
		in, out := &in.SANs, &out.SANs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new APISpec.
func (in *APISpec) DeepCopy() *APISpec {
	if in == nil {
		return nil
	}
	out := new(APISpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CaResponse) DeepCopyInto(out *CaResponse) {
	*out = *in
	if in.Key != nil {
		in, out := &in.Key, &out.Key
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.Cert != nil {
		in, out := &in.Cert, &out.Cert
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.SAKey != nil {
		in, out := &in.SAKey, &out.SAKey
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
	if in.SAPub != nil {
		in, out := &in.SAPub, &out.SAPub
		*out = make([]byte, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CaResponse.
func (in *CaResponse) DeepCopy() *CaResponse {
	if in == nil {
		return nil
	}
	out := new(CaResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Calico) DeepCopyInto(out *Calico) {
	*out = *in
	if in.EnvVars != nil {
		in, out := &in.EnvVars, &out.EnvVars
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Calico.
func (in *Calico) DeepCopy() *Calico {
	if in == nil {
		return nil
	}
	out := new(Calico)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CalicoImageSpec) DeepCopyInto(out *CalicoImageSpec) {
	*out = *in
	out.CNI = in.CNI
	out.Node = in.Node
	out.KubeControllers = in.KubeControllers
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CalicoImageSpec.
func (in *CalicoImageSpec) DeepCopy() *CalicoImageSpec {
	if in == nil {
		return nil
	}
	out := new(CalicoImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Chart) DeepCopyInto(out *Chart) {
	*out = *in
	if in.ForceUpgrade != nil {
		in, out := &in.ForceUpgrade, &out.ForceUpgrade
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Chart.
func (in *Chart) DeepCopy() *Chart {
	if in == nil {
		return nil
	}
	out := new(Chart)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ChartsSettings) DeepCopyInto(out *ChartsSettings) {
	{
		in := &in
		*out = make(ChartsSettings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ChartsSettings.
func (in ChartsSettings) DeepCopy() ChartsSettings {
	if in == nil {
		return nil
	}
	out := new(ChartsSettings)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfig) DeepCopyInto(out *ClusterConfig) {
	*out = *in
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.TypeMeta = in.TypeMeta
	if in.Spec != nil {
		in, out := &in.Spec, &out.Spec
		*out = new(ClusterSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Status != nil {
		in, out := &in.Status, &out.Status
		*out = new(ClusterConfigStatus)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfig.
func (in *ClusterConfig) DeepCopy() *ClusterConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterConfig) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigList) DeepCopyInto(out *ClusterConfigList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ClusterConfig, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigList.
func (in *ClusterConfigList) DeepCopy() *ClusterConfigList {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ClusterConfigList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfigStatus) DeepCopyInto(out *ClusterConfigStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfigStatus.
func (in *ClusterConfigStatus) DeepCopy() *ClusterConfigStatus {
	if in == nil {
		return nil
	}
	out := new(ClusterConfigStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterExtensions) DeepCopyInto(out *ClusterExtensions) {
	*out = *in
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(StorageExtension)
		**out = **in
	}
	if in.Helm != nil {
		in, out := &in.Helm, &out.Helm
		*out = new(HelmExtensions)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterExtensions.
func (in *ClusterExtensions) DeepCopy() *ClusterExtensions {
	if in == nil {
		return nil
	}
	out := new(ClusterExtensions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterImages) DeepCopyInto(out *ClusterImages) {
	*out = *in
	out.Konnectivity = in.Konnectivity
	out.PushGateway = in.PushGateway
	out.MetricsServer = in.MetricsServer
	out.KubeProxy = in.KubeProxy
	out.CoreDNS = in.CoreDNS
	out.Pause = in.Pause
	out.Calico = in.Calico
	out.KubeRouter = in.KubeRouter
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterImages.
func (in *ClusterImages) DeepCopy() *ClusterImages {
	if in == nil {
		return nil
	}
	out := new(ClusterImages)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterSpec) DeepCopyInto(out *ClusterSpec) {
	*out = *in
	if in.API != nil {
		in, out := &in.API, &out.API
		*out = new(APISpec)
		(*in).DeepCopyInto(*out)
	}
	if in.ControllerManager != nil {
		in, out := &in.ControllerManager, &out.ControllerManager
		*out = new(ControllerManagerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Scheduler != nil {
		in, out := &in.Scheduler, &out.Scheduler
		*out = new(SchedulerSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Storage != nil {
		in, out := &in.Storage, &out.Storage
		*out = new(StorageSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(Network)
		(*in).DeepCopyInto(*out)
	}
	if in.WorkerProfiles != nil {
		in, out := &in.WorkerProfiles, &out.WorkerProfiles
		*out = make(WorkerProfiles, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Telemetry != nil {
		in, out := &in.Telemetry, &out.Telemetry
		*out = new(ClusterTelemetry)
		**out = **in
	}
	if in.Install != nil {
		in, out := &in.Install, &out.Install
		*out = new(InstallSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Images != nil {
		in, out := &in.Images, &out.Images
		*out = new(ClusterImages)
		**out = **in
	}
	if in.Extensions != nil {
		in, out := &in.Extensions, &out.Extensions
		*out = new(ClusterExtensions)
		(*in).DeepCopyInto(*out)
	}
	if in.Konnectivity != nil {
		in, out := &in.Konnectivity, &out.Konnectivity
		*out = new(KonnectivitySpec)
		**out = **in
	}
	if in.FeatureGates != nil {
		in, out := &in.FeatureGates, &out.FeatureGates
		*out = make(FeatureGates, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterSpec.
func (in *ClusterSpec) DeepCopy() *ClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterTelemetry) DeepCopyInto(out *ClusterTelemetry) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterTelemetry.
func (in *ClusterTelemetry) DeepCopy() *ClusterTelemetry {
	if in == nil {
		return nil
	}
	out := new(ClusterTelemetry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ControllerManagerSpec) DeepCopyInto(out *ControllerManagerSpec) {
	*out = *in
	if in.ExtraArgs != nil {
		in, out := &in.ExtraArgs, &out.ExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ControllerManagerSpec.
func (in *ControllerManagerSpec) DeepCopy() *ControllerManagerSpec {
	if in == nil {
		return nil
	}
	out := new(ControllerManagerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DualStack) DeepCopyInto(out *DualStack) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DualStack.
func (in *DualStack) DeepCopy() *DualStack {
	if in == nil {
		return nil
	}
	out := new(DualStack)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EnvoyProxy) DeepCopyInto(out *EnvoyProxy) {
	*out = *in
	if in.Image != nil {
		in, out := &in.Image, &out.Image
		*out = new(ImageSpec)
		**out = **in
	}
	if in.KonnectivityServerBindPort != nil {
		in, out := &in.KonnectivityServerBindPort, &out.KonnectivityServerBindPort
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EnvoyProxy.
func (in *EnvoyProxy) DeepCopy() *EnvoyProxy {
	if in == nil {
		return nil
	}
	out := new(EnvoyProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdConfig) DeepCopyInto(out *EtcdConfig) {
	*out = *in
	if in.ExternalCluster != nil {
		in, out := &in.ExternalCluster, &out.ExternalCluster
		*out = new(ExternalCluster)
		(*in).DeepCopyInto(*out)
	}
	if in.ExtraArgs != nil {
		in, out := &in.ExtraArgs, &out.ExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdConfig.
func (in *EtcdConfig) DeepCopy() *EtcdConfig {
	if in == nil {
		return nil
	}
	out := new(EtcdConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdRequest) DeepCopyInto(out *EtcdRequest) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdRequest.
func (in *EtcdRequest) DeepCopy() *EtcdRequest {
	if in == nil {
		return nil
	}
	out := new(EtcdRequest)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EtcdResponse) DeepCopyInto(out *EtcdResponse) {
	*out = *in
	in.CA.DeepCopyInto(&out.CA)
	if in.InitialCluster != nil {
		in, out := &in.InitialCluster, &out.InitialCluster
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EtcdResponse.
func (in *EtcdResponse) DeepCopy() *EtcdResponse {
	if in == nil {
		return nil
	}
	out := new(EtcdResponse)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalCluster) DeepCopyInto(out *ExternalCluster) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalCluster.
func (in *ExternalCluster) DeepCopy() *ExternalCluster {
	if in == nil {
		return nil
	}
	out := new(ExternalCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FeatureGate) DeepCopyInto(out *FeatureGate) {
	*out = *in
	if in.Components != nil {
		in, out := &in.Components, &out.Components
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureGate.
func (in *FeatureGate) DeepCopy() *FeatureGate {
	if in == nil {
		return nil
	}
	out := new(FeatureGate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in FeatureGates) DeepCopyInto(out *FeatureGates) {
	{
		in := &in
		*out = make(FeatureGates, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FeatureGates.
func (in FeatureGates) DeepCopy() FeatureGates {
	if in == nil {
		return nil
	}
	out := new(FeatureGates)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *HelmExtensions) DeepCopyInto(out *HelmExtensions) {
	*out = *in
	if in.Repositories != nil {
		in, out := &in.Repositories, &out.Repositories
		*out = make(RepositoriesSettings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Charts != nil {
		in, out := &in.Charts, &out.Charts
		*out = make(ChartsSettings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new HelmExtensions.
func (in *HelmExtensions) DeepCopy() *HelmExtensions {
	if in == nil {
		return nil
	}
	out := new(HelmExtensions)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImageSpec) DeepCopyInto(out *ImageSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImageSpec.
func (in *ImageSpec) DeepCopy() *ImageSpec {
	if in == nil {
		return nil
	}
	out := new(ImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *InstallSpec) DeepCopyInto(out *InstallSpec) {
	*out = *in
	if in.SystemUsers != nil {
		in, out := &in.SystemUsers, &out.SystemUsers
		*out = new(SystemUser)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new InstallSpec.
func (in *InstallSpec) DeepCopy() *InstallSpec {
	if in == nil {
		return nil
	}
	out := new(InstallSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KineConfig) DeepCopyInto(out *KineConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KineConfig.
func (in *KineConfig) DeepCopy() *KineConfig {
	if in == nil {
		return nil
	}
	out := new(KineConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KonnectivitySpec) DeepCopyInto(out *KonnectivitySpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KonnectivitySpec.
func (in *KonnectivitySpec) DeepCopy() *KonnectivitySpec {
	if in == nil {
		return nil
	}
	out := new(KonnectivitySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxy) DeepCopyInto(out *KubeProxy) {
	*out = *in
	if in.IPTables != nil {
		in, out := &in.IPTables, &out.IPTables
		*out = new(KubeProxyIPTablesConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.IPVS != nil {
		in, out := &in.IPVS, &out.IPVS
		*out = new(KubeProxyIPVSConfiguration)
		(*in).DeepCopyInto(*out)
	}
	if in.NodePortAddresses != nil {
		in, out := &in.NodePortAddresses, &out.NodePortAddresses
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxy.
func (in *KubeProxy) DeepCopy() *KubeProxy {
	if in == nil {
		return nil
	}
	out := new(KubeProxy)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyIPTablesConfiguration) DeepCopyInto(out *KubeProxyIPTablesConfiguration) {
	*out = *in
	if in.MasqueradeBit != nil {
		in, out := &in.MasqueradeBit, &out.MasqueradeBit
		*out = new(int32)
		**out = **in
	}
	if in.LocalhostNodePorts != nil {
		in, out := &in.LocalhostNodePorts, &out.LocalhostNodePorts
		*out = new(bool)
		**out = **in
	}
	out.SyncPeriod = in.SyncPeriod
	out.MinSyncPeriod = in.MinSyncPeriod
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyIPTablesConfiguration.
func (in *KubeProxyIPTablesConfiguration) DeepCopy() *KubeProxyIPTablesConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeProxyIPTablesConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeProxyIPVSConfiguration) DeepCopyInto(out *KubeProxyIPVSConfiguration) {
	*out = *in
	out.SyncPeriod = in.SyncPeriod
	out.MinSyncPeriod = in.MinSyncPeriod
	if in.ExcludeCIDRs != nil {
		in, out := &in.ExcludeCIDRs, &out.ExcludeCIDRs
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	out.TCPTimeout = in.TCPTimeout
	out.TCPFinTimeout = in.TCPFinTimeout
	out.UDPTimeout = in.UDPTimeout
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeProxyIPVSConfiguration.
func (in *KubeProxyIPVSConfiguration) DeepCopy() *KubeProxyIPVSConfiguration {
	if in == nil {
		return nil
	}
	out := new(KubeProxyIPVSConfiguration)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeRouter) DeepCopyInto(out *KubeRouter) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeRouter.
func (in *KubeRouter) DeepCopy() *KubeRouter {
	if in == nil {
		return nil
	}
	out := new(KubeRouter)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KubeRouterImageSpec) DeepCopyInto(out *KubeRouterImageSpec) {
	*out = *in
	out.CNI = in.CNI
	out.CNIInstaller = in.CNIInstaller
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KubeRouterImageSpec.
func (in *KubeRouterImageSpec) DeepCopy() *KubeRouterImageSpec {
	if in == nil {
		return nil
	}
	out := new(KubeRouterImageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Network) DeepCopyInto(out *Network) {
	*out = *in
	if in.Calico != nil {
		in, out := &in.Calico, &out.Calico
		*out = new(Calico)
		(*in).DeepCopyInto(*out)
	}
	out.DualStack = in.DualStack
	if in.KubeProxy != nil {
		in, out := &in.KubeProxy, &out.KubeProxy
		*out = new(KubeProxy)
		(*in).DeepCopyInto(*out)
	}
	if in.KubeRouter != nil {
		in, out := &in.KubeRouter, &out.KubeRouter
		*out = new(KubeRouter)
		**out = **in
	}
	if in.NodeLocalLoadBalancing != nil {
		in, out := &in.NodeLocalLoadBalancing, &out.NodeLocalLoadBalancing
		*out = new(NodeLocalLoadBalancing)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Network.
func (in *Network) DeepCopy() *Network {
	if in == nil {
		return nil
	}
	out := new(Network)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeLocalLoadBalancing) DeepCopyInto(out *NodeLocalLoadBalancing) {
	*out = *in
	if in.EnvoyProxy != nil {
		in, out := &in.EnvoyProxy, &out.EnvoyProxy
		*out = new(EnvoyProxy)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeLocalLoadBalancing.
func (in *NodeLocalLoadBalancing) DeepCopy() *NodeLocalLoadBalancing {
	if in == nil {
		return nil
	}
	out := new(NodeLocalLoadBalancing)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in RepositoriesSettings) DeepCopyInto(out *RepositoriesSettings) {
	{
		in := &in
		*out = make(RepositoriesSettings, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new RepositoriesSettings.
func (in RepositoriesSettings) DeepCopy() RepositoriesSettings {
	if in == nil {
		return nil
	}
	out := new(RepositoriesSettings)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Repository) DeepCopyInto(out *Repository) {
	*out = *in
	if in.Insecure != nil {
		in, out := &in.Insecure, &out.Insecure
		*out = new(bool)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Repository.
func (in *Repository) DeepCopy() *Repository {
	if in == nil {
		return nil
	}
	out := new(Repository)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SchedulerSpec) DeepCopyInto(out *SchedulerSpec) {
	*out = *in
	if in.ExtraArgs != nil {
		in, out := &in.ExtraArgs, &out.ExtraArgs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SchedulerSpec.
func (in *SchedulerSpec) DeepCopy() *SchedulerSpec {
	if in == nil {
		return nil
	}
	out := new(SchedulerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageExtension) DeepCopyInto(out *StorageExtension) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageExtension.
func (in *StorageExtension) DeepCopy() *StorageExtension {
	if in == nil {
		return nil
	}
	out := new(StorageExtension)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageSpec) DeepCopyInto(out *StorageSpec) {
	*out = *in
	if in.Etcd != nil {
		in, out := &in.Etcd, &out.Etcd
		*out = new(EtcdConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.Kine != nil {
		in, out := &in.Kine, &out.Kine
		*out = new(KineConfig)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageSpec.
func (in *StorageSpec) DeepCopy() *StorageSpec {
	if in == nil {
		return nil
	}
	out := new(StorageSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SystemUser) DeepCopyInto(out *SystemUser) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SystemUser.
func (in *SystemUser) DeepCopy() *SystemUser {
	if in == nil {
		return nil
	}
	out := new(SystemUser)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WorkerProfile) DeepCopyInto(out *WorkerProfile) {
	*out = *in
	if in.Config != nil {
		in, out := &in.Config, &out.Config
		*out = new(runtime.RawExtension)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerProfile.
func (in *WorkerProfile) DeepCopy() *WorkerProfile {
	if in == nil {
		return nil
	}
	out := new(WorkerProfile)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in WorkerProfiles) DeepCopyInto(out *WorkerProfiles) {
	{
		in := &in
		*out = make(WorkerProfiles, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WorkerProfiles.
func (in WorkerProfiles) DeepCopy() WorkerProfiles {
	if in == nil {
		return nil
	}
	out := new(WorkerProfiles)
	in.DeepCopyInto(out)
	return *out
}
