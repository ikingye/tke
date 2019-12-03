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
	config "tkestack.io/tke/pkg/registry/apis/config"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*FileSystemStorage)(nil), (*config.FileSystemStorage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_FileSystemStorage_To_config_FileSystemStorage(a.(*FileSystemStorage), b.(*config.FileSystemStorage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.FileSystemStorage)(nil), (*FileSystemStorage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_FileSystemStorage_To_v1_FileSystemStorage(a.(*config.FileSystemStorage), b.(*FileSystemStorage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*InMemoryStorage)(nil), (*config.InMemoryStorage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_InMemoryStorage_To_config_InMemoryStorage(a.(*InMemoryStorage), b.(*config.InMemoryStorage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.InMemoryStorage)(nil), (*InMemoryStorage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_InMemoryStorage_To_v1_InMemoryStorage(a.(*config.InMemoryStorage), b.(*InMemoryStorage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Redis)(nil), (*config.Redis)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Redis_To_config_Redis(a.(*Redis), b.(*config.Redis), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Redis)(nil), (*Redis)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Redis_To_v1_Redis(a.(*config.Redis), b.(*Redis), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RegistryConfiguration)(nil), (*config.RegistryConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_RegistryConfiguration_To_config_RegistryConfiguration(a.(*RegistryConfiguration), b.(*config.RegistryConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.RegistryConfiguration)(nil), (*RegistryConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_RegistryConfiguration_To_v1_RegistryConfiguration(a.(*config.RegistryConfiguration), b.(*RegistryConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*S3Storage)(nil), (*config.S3Storage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_S3Storage_To_config_S3Storage(a.(*S3Storage), b.(*config.S3Storage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.S3Storage)(nil), (*S3Storage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_S3Storage_To_v1_S3Storage(a.(*config.S3Storage), b.(*S3Storage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Security)(nil), (*config.Security)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Security_To_config_Security(a.(*Security), b.(*config.Security), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Security)(nil), (*Security)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Security_To_v1_Security(a.(*config.Security), b.(*Security), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Storage)(nil), (*config.Storage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1_Storage_To_config_Storage(a.(*Storage), b.(*config.Storage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.Storage)(nil), (*Storage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_Storage_To_v1_Storage(a.(*config.Storage), b.(*Storage), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1_FileSystemStorage_To_config_FileSystemStorage(in *FileSystemStorage, out *config.FileSystemStorage, s conversion.Scope) error {
	out.RootDirectory = in.RootDirectory
	out.MaxThreads = (*int64)(unsafe.Pointer(in.MaxThreads))
	return nil
}

// Convert_v1_FileSystemStorage_To_config_FileSystemStorage is an autogenerated conversion function.
func Convert_v1_FileSystemStorage_To_config_FileSystemStorage(in *FileSystemStorage, out *config.FileSystemStorage, s conversion.Scope) error {
	return autoConvert_v1_FileSystemStorage_To_config_FileSystemStorage(in, out, s)
}

func autoConvert_config_FileSystemStorage_To_v1_FileSystemStorage(in *config.FileSystemStorage, out *FileSystemStorage, s conversion.Scope) error {
	out.RootDirectory = in.RootDirectory
	out.MaxThreads = (*int64)(unsafe.Pointer(in.MaxThreads))
	return nil
}

// Convert_config_FileSystemStorage_To_v1_FileSystemStorage is an autogenerated conversion function.
func Convert_config_FileSystemStorage_To_v1_FileSystemStorage(in *config.FileSystemStorage, out *FileSystemStorage, s conversion.Scope) error {
	return autoConvert_config_FileSystemStorage_To_v1_FileSystemStorage(in, out, s)
}

func autoConvert_v1_InMemoryStorage_To_config_InMemoryStorage(in *InMemoryStorage, out *config.InMemoryStorage, s conversion.Scope) error {
	return nil
}

// Convert_v1_InMemoryStorage_To_config_InMemoryStorage is an autogenerated conversion function.
func Convert_v1_InMemoryStorage_To_config_InMemoryStorage(in *InMemoryStorage, out *config.InMemoryStorage, s conversion.Scope) error {
	return autoConvert_v1_InMemoryStorage_To_config_InMemoryStorage(in, out, s)
}

func autoConvert_config_InMemoryStorage_To_v1_InMemoryStorage(in *config.InMemoryStorage, out *InMemoryStorage, s conversion.Scope) error {
	return nil
}

// Convert_config_InMemoryStorage_To_v1_InMemoryStorage is an autogenerated conversion function.
func Convert_config_InMemoryStorage_To_v1_InMemoryStorage(in *config.InMemoryStorage, out *InMemoryStorage, s conversion.Scope) error {
	return autoConvert_config_InMemoryStorage_To_v1_InMemoryStorage(in, out, s)
}

func autoConvert_v1_Redis_To_config_Redis(in *Redis, out *config.Redis, s conversion.Scope) error {
	out.Addr = in.Addr
	out.Password = in.Password
	out.DB = in.DB
	out.ReadTimeoutMillisecond = (*int64)(unsafe.Pointer(in.ReadTimeoutMillisecond))
	out.DialTimeoutMillisecond = (*int64)(unsafe.Pointer(in.DialTimeoutMillisecond))
	out.WriteTimeoutMillisecond = (*int64)(unsafe.Pointer(in.WriteTimeoutMillisecond))
	out.PoolMaxIdle = (*int32)(unsafe.Pointer(in.PoolMaxIdle))
	out.PoolMaxActive = (*int32)(unsafe.Pointer(in.PoolMaxActive))
	out.PoolIdleTimeoutSeconds = (*int64)(unsafe.Pointer(in.PoolIdleTimeoutSeconds))
	return nil
}

// Convert_v1_Redis_To_config_Redis is an autogenerated conversion function.
func Convert_v1_Redis_To_config_Redis(in *Redis, out *config.Redis, s conversion.Scope) error {
	return autoConvert_v1_Redis_To_config_Redis(in, out, s)
}

func autoConvert_config_Redis_To_v1_Redis(in *config.Redis, out *Redis, s conversion.Scope) error {
	out.Addr = in.Addr
	out.Password = in.Password
	out.DB = in.DB
	out.ReadTimeoutMillisecond = (*int64)(unsafe.Pointer(in.ReadTimeoutMillisecond))
	out.DialTimeoutMillisecond = (*int64)(unsafe.Pointer(in.DialTimeoutMillisecond))
	out.WriteTimeoutMillisecond = (*int64)(unsafe.Pointer(in.WriteTimeoutMillisecond))
	out.PoolMaxIdle = (*int32)(unsafe.Pointer(in.PoolMaxIdle))
	out.PoolMaxActive = (*int32)(unsafe.Pointer(in.PoolMaxActive))
	out.PoolIdleTimeoutSeconds = (*int64)(unsafe.Pointer(in.PoolIdleTimeoutSeconds))
	return nil
}

// Convert_config_Redis_To_v1_Redis is an autogenerated conversion function.
func Convert_config_Redis_To_v1_Redis(in *config.Redis, out *Redis, s conversion.Scope) error {
	return autoConvert_config_Redis_To_v1_Redis(in, out, s)
}

func autoConvert_v1_RegistryConfiguration_To_config_RegistryConfiguration(in *RegistryConfiguration, out *config.RegistryConfiguration, s conversion.Scope) error {
	if err := Convert_v1_Storage_To_config_Storage(&in.Storage, &out.Storage, s); err != nil {
		return err
	}
	if err := Convert_v1_Security_To_config_Security(&in.Security, &out.Security, s); err != nil {
		return err
	}
	out.Redis = (*config.Redis)(unsafe.Pointer(in.Redis))
	out.DefaultTenant = in.DefaultTenant
	out.DomainSuffix = in.DomainSuffix
	return nil
}

// Convert_v1_RegistryConfiguration_To_config_RegistryConfiguration is an autogenerated conversion function.
func Convert_v1_RegistryConfiguration_To_config_RegistryConfiguration(in *RegistryConfiguration, out *config.RegistryConfiguration, s conversion.Scope) error {
	return autoConvert_v1_RegistryConfiguration_To_config_RegistryConfiguration(in, out, s)
}

func autoConvert_config_RegistryConfiguration_To_v1_RegistryConfiguration(in *config.RegistryConfiguration, out *RegistryConfiguration, s conversion.Scope) error {
	if err := Convert_config_Storage_To_v1_Storage(&in.Storage, &out.Storage, s); err != nil {
		return err
	}
	if err := Convert_config_Security_To_v1_Security(&in.Security, &out.Security, s); err != nil {
		return err
	}
	out.Redis = (*Redis)(unsafe.Pointer(in.Redis))
	out.DefaultTenant = in.DefaultTenant
	out.DomainSuffix = in.DomainSuffix
	return nil
}

// Convert_config_RegistryConfiguration_To_v1_RegistryConfiguration is an autogenerated conversion function.
func Convert_config_RegistryConfiguration_To_v1_RegistryConfiguration(in *config.RegistryConfiguration, out *RegistryConfiguration, s conversion.Scope) error {
	return autoConvert_config_RegistryConfiguration_To_v1_RegistryConfiguration(in, out, s)
}

func autoConvert_v1_S3Storage_To_config_S3Storage(in *S3Storage, out *config.S3Storage, s conversion.Scope) error {
	out.Bucket = in.Bucket
	out.Region = in.Region
	out.AccessKey = (*string)(unsafe.Pointer(in.AccessKey))
	out.SecretKey = (*string)(unsafe.Pointer(in.SecretKey))
	out.RegionEndpoint = (*string)(unsafe.Pointer(in.RegionEndpoint))
	out.Encrypt = (*bool)(unsafe.Pointer(in.Encrypt))
	out.KeyID = (*string)(unsafe.Pointer(in.KeyID))
	out.Secure = (*bool)(unsafe.Pointer(in.Secure))
	out.SkipVerify = (*bool)(unsafe.Pointer(in.SkipVerify))
	out.V4Auth = (*bool)(unsafe.Pointer(in.V4Auth))
	out.ChunkSize = (*int64)(unsafe.Pointer(in.ChunkSize))
	out.MultipartCopyChunkSize = (*int64)(unsafe.Pointer(in.MultipartCopyChunkSize))
	out.MultipartCopyMaxConcurrency = (*int64)(unsafe.Pointer(in.MultipartCopyMaxConcurrency))
	out.MultipartCopyThresholdSize = (*int64)(unsafe.Pointer(in.MultipartCopyThresholdSize))
	out.RootDirectory = (*string)(unsafe.Pointer(in.RootDirectory))
	out.StorageClass = (*config.S3StorageClass)(unsafe.Pointer(in.StorageClass))
	out.UserAgent = (*string)(unsafe.Pointer(in.UserAgent))
	out.ObjectACL = (*string)(unsafe.Pointer(in.ObjectACL))
	return nil
}

// Convert_v1_S3Storage_To_config_S3Storage is an autogenerated conversion function.
func Convert_v1_S3Storage_To_config_S3Storage(in *S3Storage, out *config.S3Storage, s conversion.Scope) error {
	return autoConvert_v1_S3Storage_To_config_S3Storage(in, out, s)
}

func autoConvert_config_S3Storage_To_v1_S3Storage(in *config.S3Storage, out *S3Storage, s conversion.Scope) error {
	out.Bucket = in.Bucket
	out.Region = in.Region
	out.AccessKey = (*string)(unsafe.Pointer(in.AccessKey))
	out.SecretKey = (*string)(unsafe.Pointer(in.SecretKey))
	out.RegionEndpoint = (*string)(unsafe.Pointer(in.RegionEndpoint))
	out.Encrypt = (*bool)(unsafe.Pointer(in.Encrypt))
	out.KeyID = (*string)(unsafe.Pointer(in.KeyID))
	out.Secure = (*bool)(unsafe.Pointer(in.Secure))
	out.SkipVerify = (*bool)(unsafe.Pointer(in.SkipVerify))
	out.V4Auth = (*bool)(unsafe.Pointer(in.V4Auth))
	out.ChunkSize = (*int64)(unsafe.Pointer(in.ChunkSize))
	out.MultipartCopyChunkSize = (*int64)(unsafe.Pointer(in.MultipartCopyChunkSize))
	out.MultipartCopyMaxConcurrency = (*int64)(unsafe.Pointer(in.MultipartCopyMaxConcurrency))
	out.MultipartCopyThresholdSize = (*int64)(unsafe.Pointer(in.MultipartCopyThresholdSize))
	out.RootDirectory = (*string)(unsafe.Pointer(in.RootDirectory))
	out.StorageClass = (*S3StorageClass)(unsafe.Pointer(in.StorageClass))
	out.UserAgent = (*string)(unsafe.Pointer(in.UserAgent))
	out.ObjectACL = (*string)(unsafe.Pointer(in.ObjectACL))
	return nil
}

// Convert_config_S3Storage_To_v1_S3Storage is an autogenerated conversion function.
func Convert_config_S3Storage_To_v1_S3Storage(in *config.S3Storage, out *S3Storage, s conversion.Scope) error {
	return autoConvert_config_S3Storage_To_v1_S3Storage(in, out, s)
}

func autoConvert_v1_Security_To_config_Security(in *Security, out *config.Security, s conversion.Scope) error {
	out.TokenPrivateKeyFile = in.TokenPrivateKeyFile
	out.TokenPublicKeyFile = in.TokenPublicKeyFile
	out.TokenExpiredHours = (*int64)(unsafe.Pointer(in.TokenExpiredHours))
	out.HTTPSecret = in.HTTPSecret
	out.AdminUsername = in.AdminUsername
	out.AdminPassword = in.AdminPassword
	return nil
}

// Convert_v1_Security_To_config_Security is an autogenerated conversion function.
func Convert_v1_Security_To_config_Security(in *Security, out *config.Security, s conversion.Scope) error {
	return autoConvert_v1_Security_To_config_Security(in, out, s)
}

func autoConvert_config_Security_To_v1_Security(in *config.Security, out *Security, s conversion.Scope) error {
	out.TokenPrivateKeyFile = in.TokenPrivateKeyFile
	out.TokenPublicKeyFile = in.TokenPublicKeyFile
	out.TokenExpiredHours = (*int64)(unsafe.Pointer(in.TokenExpiredHours))
	out.HTTPSecret = in.HTTPSecret
	out.AdminUsername = in.AdminUsername
	out.AdminPassword = in.AdminPassword
	return nil
}

// Convert_config_Security_To_v1_Security is an autogenerated conversion function.
func Convert_config_Security_To_v1_Security(in *config.Security, out *Security, s conversion.Scope) error {
	return autoConvert_config_Security_To_v1_Security(in, out, s)
}

func autoConvert_v1_Storage_To_config_Storage(in *Storage, out *config.Storage, s conversion.Scope) error {
	out.FileSystem = (*config.FileSystemStorage)(unsafe.Pointer(in.FileSystem))
	out.InMemory = (*config.InMemoryStorage)(unsafe.Pointer(in.InMemory))
	out.S3 = (*config.S3Storage)(unsafe.Pointer(in.S3))
	return nil
}

// Convert_v1_Storage_To_config_Storage is an autogenerated conversion function.
func Convert_v1_Storage_To_config_Storage(in *Storage, out *config.Storage, s conversion.Scope) error {
	return autoConvert_v1_Storage_To_config_Storage(in, out, s)
}

func autoConvert_config_Storage_To_v1_Storage(in *config.Storage, out *Storage, s conversion.Scope) error {
	out.FileSystem = (*FileSystemStorage)(unsafe.Pointer(in.FileSystem))
	out.InMemory = (*InMemoryStorage)(unsafe.Pointer(in.InMemory))
	out.S3 = (*S3Storage)(unsafe.Pointer(in.S3))
	return nil
}

// Convert_config_Storage_To_v1_Storage is an autogenerated conversion function.
func Convert_config_Storage_To_v1_Storage(in *config.Storage, out *Storage, s conversion.Scope) error {
	return autoConvert_config_Storage_To_v1_Storage(in, out, s)
}
