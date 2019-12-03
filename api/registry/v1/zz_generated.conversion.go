// +build !ignore_autogenerated

/*
 * Tencent is pleased to support the open source community by making TKEStack
 * available.
 *
 * Copyright (C) 2012-2019 Tencent. All Rights Reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use
 * this file except in compliance with the License. You may obtain a copy of the
 * License at
 *
 * https://opensource.org/licenses/Apache-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS, WITHOUT
 * WARRANTIES OF ANY KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations under the License.
 */

// Code generated by conversion-gen. DO NOT EDIT.

package v1

import (
	unsafe "unsafe"

	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
	registry "tkestack.io/tke/api/registry"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*Chart)(nil), (*registry.Chart)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Chart_To_registry_Chart(a.(*Chart), b.(*registry.Chart), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*registry.Chart)(nil), (*Chart)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_registry_Chart_To_v1_Chart(a.(*registry.Chart), b.(*Chart), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ChartGroup)(nil), (*registry.ChartGroup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ChartGroup_To_registry_ChartGroup(a.(*ChartGroup), b.(*registry.ChartGroup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*registry.ChartGroup)(nil), (*ChartGroup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_registry_ChartGroup_To_v1_ChartGroup(a.(*registry.ChartGroup), b.(*ChartGroup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ChartGroupList)(nil), (*registry.ChartGroupList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ChartGroupList_To_registry_ChartGroupList(a.(*ChartGroupList), b.(*registry.ChartGroupList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*registry.ChartGroupList)(nil), (*ChartGroupList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_registry_ChartGroupList_To_v1_ChartGroupList(a.(*registry.ChartGroupList), b.(*ChartGroupList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ChartGroupSpec)(nil), (*registry.ChartGroupSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ChartGroupSpec_To_registry_ChartGroupSpec(a.(*ChartGroupSpec), b.(*registry.ChartGroupSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*registry.ChartGroupSpec)(nil), (*ChartGroupSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_registry_ChartGroupSpec_To_v1_ChartGroupSpec(a.(*registry.ChartGroupSpec), b.(*ChartGroupSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ChartGroupStatus)(nil), (*registry.ChartGroupStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ChartGroupStatus_To_registry_ChartGroupStatus(a.(*ChartGroupStatus), b.(*registry.ChartGroupStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*registry.ChartGroupStatus)(nil), (*ChartGroupStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_registry_ChartGroupStatus_To_v1_ChartGroupStatus(a.(*registry.ChartGroupStatus), b.(*ChartGroupStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ChartList)(nil), (*registry.ChartList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ChartList_To_registry_ChartList(a.(*ChartList), b.(*registry.ChartList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*registry.ChartList)(nil), (*ChartList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_registry_ChartList_To_v1_ChartList(a.(*registry.ChartList), b.(*ChartList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ChartSpec)(nil), (*registry.ChartSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ChartSpec_To_registry_ChartSpec(a.(*ChartSpec), b.(*registry.ChartSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*registry.ChartSpec)(nil), (*ChartSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_registry_ChartSpec_To_v1_ChartSpec(a.(*registry.ChartSpec), b.(*ChartSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ChartStatus)(nil), (*registry.ChartStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ChartStatus_To_registry_ChartStatus(a.(*ChartStatus), b.(*registry.ChartStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*registry.ChartStatus)(nil), (*ChartStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_registry_ChartStatus_To_v1_ChartStatus(a.(*registry.ChartStatus), b.(*ChartStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ChartVersion)(nil), (*registry.ChartVersion)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ChartVersion_To_registry_ChartVersion(a.(*ChartVersion), b.(*registry.ChartVersion), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*registry.ChartVersion)(nil), (*ChartVersion)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_registry_ChartVersion_To_v1_ChartVersion(a.(*registry.ChartVersion), b.(*ChartVersion), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConfigMap)(nil), (*registry.ConfigMap)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ConfigMap_To_registry_ConfigMap(a.(*ConfigMap), b.(*registry.ConfigMap), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*registry.ConfigMap)(nil), (*ConfigMap)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_registry_ConfigMap_To_v1_ConfigMap(a.(*registry.ConfigMap), b.(*ConfigMap), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ConfigMapList)(nil), (*registry.ConfigMapList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_ConfigMapList_To_registry_ConfigMapList(a.(*ConfigMapList), b.(*registry.ConfigMapList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*registry.ConfigMapList)(nil), (*ConfigMapList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_registry_ConfigMapList_To_v1_ConfigMapList(a.(*registry.ConfigMapList), b.(*ConfigMapList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Namespace)(nil), (*registry.Namespace)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Namespace_To_registry_Namespace(a.(*Namespace), b.(*registry.Namespace), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*registry.Namespace)(nil), (*Namespace)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_registry_Namespace_To_v1_Namespace(a.(*registry.Namespace), b.(*Namespace), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NamespaceList)(nil), (*registry.NamespaceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_NamespaceList_To_registry_NamespaceList(a.(*NamespaceList), b.(*registry.NamespaceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*registry.NamespaceList)(nil), (*NamespaceList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_registry_NamespaceList_To_v1_NamespaceList(a.(*registry.NamespaceList), b.(*NamespaceList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NamespaceSpec)(nil), (*registry.NamespaceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_NamespaceSpec_To_registry_NamespaceSpec(a.(*NamespaceSpec), b.(*registry.NamespaceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*registry.NamespaceSpec)(nil), (*NamespaceSpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_registry_NamespaceSpec_To_v1_NamespaceSpec(a.(*registry.NamespaceSpec), b.(*NamespaceSpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NamespaceStatus)(nil), (*registry.NamespaceStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_NamespaceStatus_To_registry_NamespaceStatus(a.(*NamespaceStatus), b.(*registry.NamespaceStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*registry.NamespaceStatus)(nil), (*NamespaceStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_registry_NamespaceStatus_To_v1_NamespaceStatus(a.(*registry.NamespaceStatus), b.(*NamespaceStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Repository)(nil), (*registry.Repository)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Repository_To_registry_Repository(a.(*Repository), b.(*registry.Repository), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*registry.Repository)(nil), (*Repository)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_registry_Repository_To_v1_Repository(a.(*registry.Repository), b.(*Repository), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RepositoryList)(nil), (*registry.RepositoryList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RepositoryList_To_registry_RepositoryList(a.(*RepositoryList), b.(*registry.RepositoryList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*registry.RepositoryList)(nil), (*RepositoryList)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_registry_RepositoryList_To_v1_RepositoryList(a.(*registry.RepositoryList), b.(*RepositoryList), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RepositorySpec)(nil), (*registry.RepositorySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RepositorySpec_To_registry_RepositorySpec(a.(*RepositorySpec), b.(*registry.RepositorySpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*registry.RepositorySpec)(nil), (*RepositorySpec)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_registry_RepositorySpec_To_v1_RepositorySpec(a.(*registry.RepositorySpec), b.(*RepositorySpec), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RepositoryStatus)(nil), (*registry.RepositoryStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RepositoryStatus_To_registry_RepositoryStatus(a.(*RepositoryStatus), b.(*registry.RepositoryStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*registry.RepositoryStatus)(nil), (*RepositoryStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_registry_RepositoryStatus_To_v1_RepositoryStatus(a.(*registry.RepositoryStatus), b.(*RepositoryStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RepositoryTag)(nil), (*registry.RepositoryTag)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RepositoryTag_To_registry_RepositoryTag(a.(*RepositoryTag), b.(*registry.RepositoryTag), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*registry.RepositoryTag)(nil), (*RepositoryTag)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_registry_RepositoryTag_To_v1_RepositoryTag(a.(*registry.RepositoryTag), b.(*RepositoryTag), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_Chart_To_registry_Chart(in *Chart, out *registry.Chart, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_ChartSpec_To_registry_ChartSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ChartStatus_To_registry_ChartStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Chart_To_registry_Chart is an autogenerated conversion function.
func Convert_v1_Chart_To_registry_Chart(in *Chart, out *registry.Chart, s conversion.Scope) error {
	return autoConvert_v1_Chart_To_registry_Chart(in, out, s)
}

func autoConvert_registry_Chart_To_v1_Chart(in *registry.Chart, out *Chart, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_registry_ChartSpec_To_v1_ChartSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_registry_ChartStatus_To_v1_ChartStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_registry_Chart_To_v1_Chart is an autogenerated conversion function.
func Convert_registry_Chart_To_v1_Chart(in *registry.Chart, out *Chart, s conversion.Scope) error {
	return autoConvert_registry_Chart_To_v1_Chart(in, out, s)
}

func autoConvert_v1_ChartGroup_To_registry_ChartGroup(in *ChartGroup, out *registry.ChartGroup, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_ChartGroupSpec_To_registry_ChartGroupSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_ChartGroupStatus_To_registry_ChartGroupStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_ChartGroup_To_registry_ChartGroup is an autogenerated conversion function.
func Convert_v1_ChartGroup_To_registry_ChartGroup(in *ChartGroup, out *registry.ChartGroup, s conversion.Scope) error {
	return autoConvert_v1_ChartGroup_To_registry_ChartGroup(in, out, s)
}

func autoConvert_registry_ChartGroup_To_v1_ChartGroup(in *registry.ChartGroup, out *ChartGroup, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_registry_ChartGroupSpec_To_v1_ChartGroupSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_registry_ChartGroupStatus_To_v1_ChartGroupStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_registry_ChartGroup_To_v1_ChartGroup is an autogenerated conversion function.
func Convert_registry_ChartGroup_To_v1_ChartGroup(in *registry.ChartGroup, out *ChartGroup, s conversion.Scope) error {
	return autoConvert_registry_ChartGroup_To_v1_ChartGroup(in, out, s)
}

func autoConvert_v1_ChartGroupList_To_registry_ChartGroupList(in *ChartGroupList, out *registry.ChartGroupList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]registry.ChartGroup)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_ChartGroupList_To_registry_ChartGroupList is an autogenerated conversion function.
func Convert_v1_ChartGroupList_To_registry_ChartGroupList(in *ChartGroupList, out *registry.ChartGroupList, s conversion.Scope) error {
	return autoConvert_v1_ChartGroupList_To_registry_ChartGroupList(in, out, s)
}

func autoConvert_registry_ChartGroupList_To_v1_ChartGroupList(in *registry.ChartGroupList, out *ChartGroupList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ChartGroup)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_registry_ChartGroupList_To_v1_ChartGroupList is an autogenerated conversion function.
func Convert_registry_ChartGroupList_To_v1_ChartGroupList(in *registry.ChartGroupList, out *ChartGroupList, s conversion.Scope) error {
	return autoConvert_registry_ChartGroupList_To_v1_ChartGroupList(in, out, s)
}

func autoConvert_v1_ChartGroupSpec_To_registry_ChartGroupSpec(in *ChartGroupSpec, out *registry.ChartGroupSpec, s conversion.Scope) error {
	out.Name = in.Name
	out.TenantID = in.TenantID
	out.DisplayName = in.DisplayName
	out.Visibility = registry.Visibility(in.Visibility)
	return nil
}

// Convert_v1_ChartGroupSpec_To_registry_ChartGroupSpec is an autogenerated conversion function.
func Convert_v1_ChartGroupSpec_To_registry_ChartGroupSpec(in *ChartGroupSpec, out *registry.ChartGroupSpec, s conversion.Scope) error {
	return autoConvert_v1_ChartGroupSpec_To_registry_ChartGroupSpec(in, out, s)
}

func autoConvert_registry_ChartGroupSpec_To_v1_ChartGroupSpec(in *registry.ChartGroupSpec, out *ChartGroupSpec, s conversion.Scope) error {
	out.Name = in.Name
	out.TenantID = in.TenantID
	out.DisplayName = in.DisplayName
	out.Visibility = Visibility(in.Visibility)
	return nil
}

// Convert_registry_ChartGroupSpec_To_v1_ChartGroupSpec is an autogenerated conversion function.
func Convert_registry_ChartGroupSpec_To_v1_ChartGroupSpec(in *registry.ChartGroupSpec, out *ChartGroupSpec, s conversion.Scope) error {
	return autoConvert_registry_ChartGroupSpec_To_v1_ChartGroupSpec(in, out, s)
}

func autoConvert_v1_ChartGroupStatus_To_registry_ChartGroupStatus(in *ChartGroupStatus, out *registry.ChartGroupStatus, s conversion.Scope) error {
	out.Locked = (*bool)(unsafe.Pointer(in.Locked))
	out.ChartCount = in.ChartCount
	return nil
}

// Convert_v1_ChartGroupStatus_To_registry_ChartGroupStatus is an autogenerated conversion function.
func Convert_v1_ChartGroupStatus_To_registry_ChartGroupStatus(in *ChartGroupStatus, out *registry.ChartGroupStatus, s conversion.Scope) error {
	return autoConvert_v1_ChartGroupStatus_To_registry_ChartGroupStatus(in, out, s)
}

func autoConvert_registry_ChartGroupStatus_To_v1_ChartGroupStatus(in *registry.ChartGroupStatus, out *ChartGroupStatus, s conversion.Scope) error {
	out.Locked = (*bool)(unsafe.Pointer(in.Locked))
	out.ChartCount = in.ChartCount
	return nil
}

// Convert_registry_ChartGroupStatus_To_v1_ChartGroupStatus is an autogenerated conversion function.
func Convert_registry_ChartGroupStatus_To_v1_ChartGroupStatus(in *registry.ChartGroupStatus, out *ChartGroupStatus, s conversion.Scope) error {
	return autoConvert_registry_ChartGroupStatus_To_v1_ChartGroupStatus(in, out, s)
}

func autoConvert_v1_ChartList_To_registry_ChartList(in *ChartList, out *registry.ChartList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]registry.Chart)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_ChartList_To_registry_ChartList is an autogenerated conversion function.
func Convert_v1_ChartList_To_registry_ChartList(in *ChartList, out *registry.ChartList, s conversion.Scope) error {
	return autoConvert_v1_ChartList_To_registry_ChartList(in, out, s)
}

func autoConvert_registry_ChartList_To_v1_ChartList(in *registry.ChartList, out *ChartList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Chart)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_registry_ChartList_To_v1_ChartList is an autogenerated conversion function.
func Convert_registry_ChartList_To_v1_ChartList(in *registry.ChartList, out *ChartList, s conversion.Scope) error {
	return autoConvert_registry_ChartList_To_v1_ChartList(in, out, s)
}

func autoConvert_v1_ChartSpec_To_registry_ChartSpec(in *ChartSpec, out *registry.ChartSpec, s conversion.Scope) error {
	out.Name = in.Name
	out.TenantID = in.TenantID
	out.ChartGroupName = in.ChartGroupName
	out.DisplayName = in.DisplayName
	out.Visibility = registry.Visibility(in.Visibility)
	return nil
}

// Convert_v1_ChartSpec_To_registry_ChartSpec is an autogenerated conversion function.
func Convert_v1_ChartSpec_To_registry_ChartSpec(in *ChartSpec, out *registry.ChartSpec, s conversion.Scope) error {
	return autoConvert_v1_ChartSpec_To_registry_ChartSpec(in, out, s)
}

func autoConvert_registry_ChartSpec_To_v1_ChartSpec(in *registry.ChartSpec, out *ChartSpec, s conversion.Scope) error {
	out.Name = in.Name
	out.TenantID = in.TenantID
	out.ChartGroupName = in.ChartGroupName
	out.DisplayName = in.DisplayName
	out.Visibility = Visibility(in.Visibility)
	return nil
}

// Convert_registry_ChartSpec_To_v1_ChartSpec is an autogenerated conversion function.
func Convert_registry_ChartSpec_To_v1_ChartSpec(in *registry.ChartSpec, out *ChartSpec, s conversion.Scope) error {
	return autoConvert_registry_ChartSpec_To_v1_ChartSpec(in, out, s)
}

func autoConvert_v1_ChartStatus_To_registry_ChartStatus(in *ChartStatus, out *registry.ChartStatus, s conversion.Scope) error {
	out.Locked = (*bool)(unsafe.Pointer(in.Locked))
	out.PullCount = in.PullCount
	out.Versions = *(*[]registry.ChartVersion)(unsafe.Pointer(&in.Versions))
	return nil
}

// Convert_v1_ChartStatus_To_registry_ChartStatus is an autogenerated conversion function.
func Convert_v1_ChartStatus_To_registry_ChartStatus(in *ChartStatus, out *registry.ChartStatus, s conversion.Scope) error {
	return autoConvert_v1_ChartStatus_To_registry_ChartStatus(in, out, s)
}

func autoConvert_registry_ChartStatus_To_v1_ChartStatus(in *registry.ChartStatus, out *ChartStatus, s conversion.Scope) error {
	out.Locked = (*bool)(unsafe.Pointer(in.Locked))
	out.PullCount = in.PullCount
	out.Versions = *(*[]ChartVersion)(unsafe.Pointer(&in.Versions))
	return nil
}

// Convert_registry_ChartStatus_To_v1_ChartStatus is an autogenerated conversion function.
func Convert_registry_ChartStatus_To_v1_ChartStatus(in *registry.ChartStatus, out *ChartStatus, s conversion.Scope) error {
	return autoConvert_registry_ChartStatus_To_v1_ChartStatus(in, out, s)
}

func autoConvert_v1_ChartVersion_To_registry_ChartVersion(in *ChartVersion, out *registry.ChartVersion, s conversion.Scope) error {
	out.Version = in.Version
	out.ChartSize = in.ChartSize
	out.TimeCreated = in.TimeCreated
	return nil
}

// Convert_v1_ChartVersion_To_registry_ChartVersion is an autogenerated conversion function.
func Convert_v1_ChartVersion_To_registry_ChartVersion(in *ChartVersion, out *registry.ChartVersion, s conversion.Scope) error {
	return autoConvert_v1_ChartVersion_To_registry_ChartVersion(in, out, s)
}

func autoConvert_registry_ChartVersion_To_v1_ChartVersion(in *registry.ChartVersion, out *ChartVersion, s conversion.Scope) error {
	out.Version = in.Version
	out.ChartSize = in.ChartSize
	out.TimeCreated = in.TimeCreated
	return nil
}

// Convert_registry_ChartVersion_To_v1_ChartVersion is an autogenerated conversion function.
func Convert_registry_ChartVersion_To_v1_ChartVersion(in *registry.ChartVersion, out *ChartVersion, s conversion.Scope) error {
	return autoConvert_registry_ChartVersion_To_v1_ChartVersion(in, out, s)
}

func autoConvert_v1_ConfigMap_To_registry_ConfigMap(in *ConfigMap, out *registry.ConfigMap, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Data = *(*map[string]string)(unsafe.Pointer(&in.Data))
	out.BinaryData = *(*map[string][]byte)(unsafe.Pointer(&in.BinaryData))
	return nil
}

// Convert_v1_ConfigMap_To_registry_ConfigMap is an autogenerated conversion function.
func Convert_v1_ConfigMap_To_registry_ConfigMap(in *ConfigMap, out *registry.ConfigMap, s conversion.Scope) error {
	return autoConvert_v1_ConfigMap_To_registry_ConfigMap(in, out, s)
}

func autoConvert_registry_ConfigMap_To_v1_ConfigMap(in *registry.ConfigMap, out *ConfigMap, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	out.Data = *(*map[string]string)(unsafe.Pointer(&in.Data))
	out.BinaryData = *(*map[string][]byte)(unsafe.Pointer(&in.BinaryData))
	return nil
}

// Convert_registry_ConfigMap_To_v1_ConfigMap is an autogenerated conversion function.
func Convert_registry_ConfigMap_To_v1_ConfigMap(in *registry.ConfigMap, out *ConfigMap, s conversion.Scope) error {
	return autoConvert_registry_ConfigMap_To_v1_ConfigMap(in, out, s)
}

func autoConvert_v1_ConfigMapList_To_registry_ConfigMapList(in *ConfigMapList, out *registry.ConfigMapList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]registry.ConfigMap)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_ConfigMapList_To_registry_ConfigMapList is an autogenerated conversion function.
func Convert_v1_ConfigMapList_To_registry_ConfigMapList(in *ConfigMapList, out *registry.ConfigMapList, s conversion.Scope) error {
	return autoConvert_v1_ConfigMapList_To_registry_ConfigMapList(in, out, s)
}

func autoConvert_registry_ConfigMapList_To_v1_ConfigMapList(in *registry.ConfigMapList, out *ConfigMapList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]ConfigMap)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_registry_ConfigMapList_To_v1_ConfigMapList is an autogenerated conversion function.
func Convert_registry_ConfigMapList_To_v1_ConfigMapList(in *registry.ConfigMapList, out *ConfigMapList, s conversion.Scope) error {
	return autoConvert_registry_ConfigMapList_To_v1_ConfigMapList(in, out, s)
}

func autoConvert_v1_Namespace_To_registry_Namespace(in *Namespace, out *registry.Namespace, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_NamespaceSpec_To_registry_NamespaceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_NamespaceStatus_To_registry_NamespaceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Namespace_To_registry_Namespace is an autogenerated conversion function.
func Convert_v1_Namespace_To_registry_Namespace(in *Namespace, out *registry.Namespace, s conversion.Scope) error {
	return autoConvert_v1_Namespace_To_registry_Namespace(in, out, s)
}

func autoConvert_registry_Namespace_To_v1_Namespace(in *registry.Namespace, out *Namespace, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_registry_NamespaceSpec_To_v1_NamespaceSpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_registry_NamespaceStatus_To_v1_NamespaceStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_registry_Namespace_To_v1_Namespace is an autogenerated conversion function.
func Convert_registry_Namespace_To_v1_Namespace(in *registry.Namespace, out *Namespace, s conversion.Scope) error {
	return autoConvert_registry_Namespace_To_v1_Namespace(in, out, s)
}

func autoConvert_v1_NamespaceList_To_registry_NamespaceList(in *NamespaceList, out *registry.NamespaceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]registry.Namespace)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_NamespaceList_To_registry_NamespaceList is an autogenerated conversion function.
func Convert_v1_NamespaceList_To_registry_NamespaceList(in *NamespaceList, out *registry.NamespaceList, s conversion.Scope) error {
	return autoConvert_v1_NamespaceList_To_registry_NamespaceList(in, out, s)
}

func autoConvert_registry_NamespaceList_To_v1_NamespaceList(in *registry.NamespaceList, out *NamespaceList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Namespace)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_registry_NamespaceList_To_v1_NamespaceList is an autogenerated conversion function.
func Convert_registry_NamespaceList_To_v1_NamespaceList(in *registry.NamespaceList, out *NamespaceList, s conversion.Scope) error {
	return autoConvert_registry_NamespaceList_To_v1_NamespaceList(in, out, s)
}

func autoConvert_v1_NamespaceSpec_To_registry_NamespaceSpec(in *NamespaceSpec, out *registry.NamespaceSpec, s conversion.Scope) error {
	out.Name = in.Name
	out.TenantID = in.TenantID
	out.DisplayName = in.DisplayName
	out.Visibility = registry.Visibility(in.Visibility)
	return nil
}

// Convert_v1_NamespaceSpec_To_registry_NamespaceSpec is an autogenerated conversion function.
func Convert_v1_NamespaceSpec_To_registry_NamespaceSpec(in *NamespaceSpec, out *registry.NamespaceSpec, s conversion.Scope) error {
	return autoConvert_v1_NamespaceSpec_To_registry_NamespaceSpec(in, out, s)
}

func autoConvert_registry_NamespaceSpec_To_v1_NamespaceSpec(in *registry.NamespaceSpec, out *NamespaceSpec, s conversion.Scope) error {
	out.Name = in.Name
	out.TenantID = in.TenantID
	out.DisplayName = in.DisplayName
	out.Visibility = Visibility(in.Visibility)
	return nil
}

// Convert_registry_NamespaceSpec_To_v1_NamespaceSpec is an autogenerated conversion function.
func Convert_registry_NamespaceSpec_To_v1_NamespaceSpec(in *registry.NamespaceSpec, out *NamespaceSpec, s conversion.Scope) error {
	return autoConvert_registry_NamespaceSpec_To_v1_NamespaceSpec(in, out, s)
}

func autoConvert_v1_NamespaceStatus_To_registry_NamespaceStatus(in *NamespaceStatus, out *registry.NamespaceStatus, s conversion.Scope) error {
	out.Locked = (*bool)(unsafe.Pointer(in.Locked))
	out.RepoCount = in.RepoCount
	return nil
}

// Convert_v1_NamespaceStatus_To_registry_NamespaceStatus is an autogenerated conversion function.
func Convert_v1_NamespaceStatus_To_registry_NamespaceStatus(in *NamespaceStatus, out *registry.NamespaceStatus, s conversion.Scope) error {
	return autoConvert_v1_NamespaceStatus_To_registry_NamespaceStatus(in, out, s)
}

func autoConvert_registry_NamespaceStatus_To_v1_NamespaceStatus(in *registry.NamespaceStatus, out *NamespaceStatus, s conversion.Scope) error {
	out.Locked = (*bool)(unsafe.Pointer(in.Locked))
	out.RepoCount = in.RepoCount
	return nil
}

// Convert_registry_NamespaceStatus_To_v1_NamespaceStatus is an autogenerated conversion function.
func Convert_registry_NamespaceStatus_To_v1_NamespaceStatus(in *registry.NamespaceStatus, out *NamespaceStatus, s conversion.Scope) error {
	return autoConvert_registry_NamespaceStatus_To_v1_NamespaceStatus(in, out, s)
}

func autoConvert_v1_Repository_To_registry_Repository(in *Repository, out *registry.Repository, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_v1_RepositorySpec_To_registry_RepositorySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_v1_RepositoryStatus_To_registry_RepositoryStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1_Repository_To_registry_Repository is an autogenerated conversion function.
func Convert_v1_Repository_To_registry_Repository(in *Repository, out *registry.Repository, s conversion.Scope) error {
	return autoConvert_v1_Repository_To_registry_Repository(in, out, s)
}

func autoConvert_registry_Repository_To_v1_Repository(in *registry.Repository, out *Repository, s conversion.Scope) error {
	out.ObjectMeta = in.ObjectMeta
	if err := Convert_registry_RepositorySpec_To_v1_RepositorySpec(&in.Spec, &out.Spec, s); err != nil {
		return err
	}
	if err := Convert_registry_RepositoryStatus_To_v1_RepositoryStatus(&in.Status, &out.Status, s); err != nil {
		return err
	}
	return nil
}

// Convert_registry_Repository_To_v1_Repository is an autogenerated conversion function.
func Convert_registry_Repository_To_v1_Repository(in *registry.Repository, out *Repository, s conversion.Scope) error {
	return autoConvert_registry_Repository_To_v1_Repository(in, out, s)
}

func autoConvert_v1_RepositoryList_To_registry_RepositoryList(in *RepositoryList, out *registry.RepositoryList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]registry.Repository)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_v1_RepositoryList_To_registry_RepositoryList is an autogenerated conversion function.
func Convert_v1_RepositoryList_To_registry_RepositoryList(in *RepositoryList, out *registry.RepositoryList, s conversion.Scope) error {
	return autoConvert_v1_RepositoryList_To_registry_RepositoryList(in, out, s)
}

func autoConvert_registry_RepositoryList_To_v1_RepositoryList(in *registry.RepositoryList, out *RepositoryList, s conversion.Scope) error {
	out.ListMeta = in.ListMeta
	out.Items = *(*[]Repository)(unsafe.Pointer(&in.Items))
	return nil
}

// Convert_registry_RepositoryList_To_v1_RepositoryList is an autogenerated conversion function.
func Convert_registry_RepositoryList_To_v1_RepositoryList(in *registry.RepositoryList, out *RepositoryList, s conversion.Scope) error {
	return autoConvert_registry_RepositoryList_To_v1_RepositoryList(in, out, s)
}

func autoConvert_v1_RepositorySpec_To_registry_RepositorySpec(in *RepositorySpec, out *registry.RepositorySpec, s conversion.Scope) error {
	out.Name = in.Name
	out.TenantID = in.TenantID
	out.NamespaceName = in.NamespaceName
	out.DisplayName = in.DisplayName
	out.Visibility = registry.Visibility(in.Visibility)
	return nil
}

// Convert_v1_RepositorySpec_To_registry_RepositorySpec is an autogenerated conversion function.
func Convert_v1_RepositorySpec_To_registry_RepositorySpec(in *RepositorySpec, out *registry.RepositorySpec, s conversion.Scope) error {
	return autoConvert_v1_RepositorySpec_To_registry_RepositorySpec(in, out, s)
}

func autoConvert_registry_RepositorySpec_To_v1_RepositorySpec(in *registry.RepositorySpec, out *RepositorySpec, s conversion.Scope) error {
	out.Name = in.Name
	out.TenantID = in.TenantID
	out.NamespaceName = in.NamespaceName
	out.DisplayName = in.DisplayName
	out.Visibility = Visibility(in.Visibility)
	return nil
}

// Convert_registry_RepositorySpec_To_v1_RepositorySpec is an autogenerated conversion function.
func Convert_registry_RepositorySpec_To_v1_RepositorySpec(in *registry.RepositorySpec, out *RepositorySpec, s conversion.Scope) error {
	return autoConvert_registry_RepositorySpec_To_v1_RepositorySpec(in, out, s)
}

func autoConvert_v1_RepositoryStatus_To_registry_RepositoryStatus(in *RepositoryStatus, out *registry.RepositoryStatus, s conversion.Scope) error {
	out.Locked = (*bool)(unsafe.Pointer(in.Locked))
	out.PullCount = in.PullCount
	out.Tags = *(*[]registry.RepositoryTag)(unsafe.Pointer(&in.Tags))
	return nil
}

// Convert_v1_RepositoryStatus_To_registry_RepositoryStatus is an autogenerated conversion function.
func Convert_v1_RepositoryStatus_To_registry_RepositoryStatus(in *RepositoryStatus, out *registry.RepositoryStatus, s conversion.Scope) error {
	return autoConvert_v1_RepositoryStatus_To_registry_RepositoryStatus(in, out, s)
}

func autoConvert_registry_RepositoryStatus_To_v1_RepositoryStatus(in *registry.RepositoryStatus, out *RepositoryStatus, s conversion.Scope) error {
	out.Locked = (*bool)(unsafe.Pointer(in.Locked))
	out.PullCount = in.PullCount
	out.Tags = *(*[]RepositoryTag)(unsafe.Pointer(&in.Tags))
	return nil
}

// Convert_registry_RepositoryStatus_To_v1_RepositoryStatus is an autogenerated conversion function.
func Convert_registry_RepositoryStatus_To_v1_RepositoryStatus(in *registry.RepositoryStatus, out *RepositoryStatus, s conversion.Scope) error {
	return autoConvert_registry_RepositoryStatus_To_v1_RepositoryStatus(in, out, s)
}

func autoConvert_v1_RepositoryTag_To_registry_RepositoryTag(in *RepositoryTag, out *registry.RepositoryTag, s conversion.Scope) error {
	out.Name = in.Name
	out.Digest = in.Digest
	out.TimeCreated = in.TimeCreated
	return nil
}

// Convert_v1_RepositoryTag_To_registry_RepositoryTag is an autogenerated conversion function.
func Convert_v1_RepositoryTag_To_registry_RepositoryTag(in *RepositoryTag, out *registry.RepositoryTag, s conversion.Scope) error {
	return autoConvert_v1_RepositoryTag_To_registry_RepositoryTag(in, out, s)
}

func autoConvert_registry_RepositoryTag_To_v1_RepositoryTag(in *registry.RepositoryTag, out *RepositoryTag, s conversion.Scope) error {
	out.Name = in.Name
	out.Digest = in.Digest
	out.TimeCreated = in.TimeCreated
	return nil
}

// Convert_registry_RepositoryTag_To_v1_RepositoryTag is an autogenerated conversion function.
func Convert_registry_RepositoryTag_To_v1_RepositoryTag(in *registry.RepositoryTag, out *RepositoryTag, s conversion.Scope) error {
	return autoConvert_registry_RepositoryTag_To_v1_RepositoryTag(in, out, s)
}
