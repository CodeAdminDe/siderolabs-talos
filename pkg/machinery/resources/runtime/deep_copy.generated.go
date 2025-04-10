// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by "deep-copy -type DevicesStatusSpec -type DiagnosticSpec -type EventSinkConfigSpec -type ExtensionServiceConfigSpec -type ExtensionServiceConfigStatusSpec -type KernelModuleSpecSpec -type KernelParamSpecSpec -type KernelParamStatusSpec -type KmsgLogConfigSpec -type MaintenanceServiceConfigSpec -type MaintenanceServiceRequestSpec -type MachineResetSignalSpec -type MachineStatusSpec -type MetaKeySpec -type MountStatusSpec -type PlatformMetadataSpec -type SecurityStateSpec -type MetaLoadedSpec -type UniqueMachineTokenSpec -type VersionSpec -type WatchdogTimerConfigSpec -type WatchdogTimerStatusSpec -header-file ../../../../hack/boilerplate.txt -o deep_copy.generated.go ."; DO NOT EDIT.

package runtime

import (
	"net/netip"
	"net/url"
)

// DeepCopy generates a deep copy of DevicesStatusSpec.
func (o DevicesStatusSpec) DeepCopy() DevicesStatusSpec {
	var cp DevicesStatusSpec = o
	return cp
}

// DeepCopy generates a deep copy of DiagnosticSpec.
func (o DiagnosticSpec) DeepCopy() DiagnosticSpec {
	var cp DiagnosticSpec = o
	if o.Details != nil {
		cp.Details = make([]string, len(o.Details))
		copy(cp.Details, o.Details)
	}
	return cp
}

// DeepCopy generates a deep copy of EventSinkConfigSpec.
func (o EventSinkConfigSpec) DeepCopy() EventSinkConfigSpec {
	var cp EventSinkConfigSpec = o
	return cp
}

// DeepCopy generates a deep copy of ExtensionServiceConfigSpec.
func (o ExtensionServiceConfigSpec) DeepCopy() ExtensionServiceConfigSpec {
	var cp ExtensionServiceConfigSpec = o
	if o.Files != nil {
		cp.Files = make([]ExtensionServiceConfigFile, len(o.Files))
		copy(cp.Files, o.Files)
	}
	if o.Environment != nil {
		cp.Environment = make([]string, len(o.Environment))
		copy(cp.Environment, o.Environment)
	}
	return cp
}

// DeepCopy generates a deep copy of ExtensionServiceConfigStatusSpec.
func (o ExtensionServiceConfigStatusSpec) DeepCopy() ExtensionServiceConfigStatusSpec {
	var cp ExtensionServiceConfigStatusSpec = o
	return cp
}

// DeepCopy generates a deep copy of KernelModuleSpecSpec.
func (o KernelModuleSpecSpec) DeepCopy() KernelModuleSpecSpec {
	var cp KernelModuleSpecSpec = o
	if o.Parameters != nil {
		cp.Parameters = make([]string, len(o.Parameters))
		copy(cp.Parameters, o.Parameters)
	}
	return cp
}

// DeepCopy generates a deep copy of KernelParamSpecSpec.
func (o KernelParamSpecSpec) DeepCopy() KernelParamSpecSpec {
	var cp KernelParamSpecSpec = o
	return cp
}

// DeepCopy generates a deep copy of KernelParamStatusSpec.
func (o KernelParamStatusSpec) DeepCopy() KernelParamStatusSpec {
	var cp KernelParamStatusSpec = o
	return cp
}

// DeepCopy generates a deep copy of KmsgLogConfigSpec.
func (o KmsgLogConfigSpec) DeepCopy() KmsgLogConfigSpec {
	var cp KmsgLogConfigSpec = o
	if o.Destinations != nil {
		cp.Destinations = make([]*url.URL, len(o.Destinations))
		copy(cp.Destinations, o.Destinations)
		for i2 := range o.Destinations {
			if o.Destinations[i2] != nil {
				cp.Destinations[i2] = new(url.URL)
				*cp.Destinations[i2] = *o.Destinations[i2]
				if o.Destinations[i2].User != nil {
					cp.Destinations[i2].User = new(url.Userinfo)
					*cp.Destinations[i2].User = *o.Destinations[i2].User
				}
			}
		}
	}
	return cp
}

// DeepCopy generates a deep copy of MaintenanceServiceConfigSpec.
func (o MaintenanceServiceConfigSpec) DeepCopy() MaintenanceServiceConfigSpec {
	var cp MaintenanceServiceConfigSpec = o
	if o.ReachableAddresses != nil {
		cp.ReachableAddresses = make([]netip.Addr, len(o.ReachableAddresses))
		copy(cp.ReachableAddresses, o.ReachableAddresses)
	}
	return cp
}

// DeepCopy generates a deep copy of MaintenanceServiceRequestSpec.
func (o MaintenanceServiceRequestSpec) DeepCopy() MaintenanceServiceRequestSpec {
	var cp MaintenanceServiceRequestSpec = o
	return cp
}

// DeepCopy generates a deep copy of MachineResetSignalSpec.
func (o MachineResetSignalSpec) DeepCopy() MachineResetSignalSpec {
	var cp MachineResetSignalSpec = o
	return cp
}

// DeepCopy generates a deep copy of MachineStatusSpec.
func (o MachineStatusSpec) DeepCopy() MachineStatusSpec {
	var cp MachineStatusSpec = o
	if o.Status.UnmetConditions != nil {
		cp.Status.UnmetConditions = make([]UnmetCondition, len(o.Status.UnmetConditions))
		copy(cp.Status.UnmetConditions, o.Status.UnmetConditions)
	}
	return cp
}

// DeepCopy generates a deep copy of MetaKeySpec.
func (o MetaKeySpec) DeepCopy() MetaKeySpec {
	var cp MetaKeySpec = o
	return cp
}

// DeepCopy generates a deep copy of MountStatusSpec.
func (o MountStatusSpec) DeepCopy() MountStatusSpec {
	var cp MountStatusSpec = o
	if o.Options != nil {
		cp.Options = make([]string, len(o.Options))
		copy(cp.Options, o.Options)
	}
	if o.EncryptionProviders != nil {
		cp.EncryptionProviders = make([]string, len(o.EncryptionProviders))
		copy(cp.EncryptionProviders, o.EncryptionProviders)
	}
	return cp
}

// DeepCopy generates a deep copy of PlatformMetadataSpec.
func (o PlatformMetadataSpec) DeepCopy() PlatformMetadataSpec {
	var cp PlatformMetadataSpec = o
	if o.Tags != nil {
		cp.Tags = make(map[string]string, len(o.Tags))
		for k2, v2 := range o.Tags {
			cp.Tags[k2] = v2
		}
	}
	return cp
}

// DeepCopy generates a deep copy of SecurityStateSpec.
func (o SecurityStateSpec) DeepCopy() SecurityStateSpec {
	var cp SecurityStateSpec = o
	return cp
}

// DeepCopy generates a deep copy of MetaLoadedSpec.
func (o MetaLoadedSpec) DeepCopy() MetaLoadedSpec {
	var cp MetaLoadedSpec = o
	return cp
}

// DeepCopy generates a deep copy of UniqueMachineTokenSpec.
func (o UniqueMachineTokenSpec) DeepCopy() UniqueMachineTokenSpec {
	var cp UniqueMachineTokenSpec = o
	return cp
}

// DeepCopy generates a deep copy of VersionSpec.
func (o VersionSpec) DeepCopy() VersionSpec {
	var cp VersionSpec = o
	return cp
}

// DeepCopy generates a deep copy of WatchdogTimerConfigSpec.
func (o WatchdogTimerConfigSpec) DeepCopy() WatchdogTimerConfigSpec {
	var cp WatchdogTimerConfigSpec = o
	return cp
}

// DeepCopy generates a deep copy of WatchdogTimerStatusSpec.
func (o WatchdogTimerStatusSpec) DeepCopy() WatchdogTimerStatusSpec {
	var cp WatchdogTimerStatusSpec = o
	return cp
}
