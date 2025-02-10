// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

// Code generated by "deep-copy -type DefaultActionConfigV1Alpha1 -type KubespanEndpointsConfigV1Alpha1 -type EthernetConfigV1Alpha1 -type RuleConfigV1Alpha1 -pointer-receiver -header-file ../../../../../hack/boilerplate.txt -o deep_copy.generated.go ."; DO NOT EDIT.

package network

import (
	"net/netip"
)

// DeepCopy generates a deep copy of *DefaultActionConfigV1Alpha1.
func (o *DefaultActionConfigV1Alpha1) DeepCopy() *DefaultActionConfigV1Alpha1 {
	var cp DefaultActionConfigV1Alpha1 = *o
	return &cp
}

// DeepCopy generates a deep copy of *KubespanEndpointsConfigV1Alpha1.
func (o *KubespanEndpointsConfigV1Alpha1) DeepCopy() *KubespanEndpointsConfigV1Alpha1 {
	var cp KubespanEndpointsConfigV1Alpha1 = *o
	if o.ExtraAnnouncedEndpointsConfig != nil {
		cp.ExtraAnnouncedEndpointsConfig = make([]netip.AddrPort, len(o.ExtraAnnouncedEndpointsConfig))
		copy(cp.ExtraAnnouncedEndpointsConfig, o.ExtraAnnouncedEndpointsConfig)
	}
	return &cp
}

// DeepCopy generates a deep copy of *EthernetConfigV1Alpha1.
func (o *EthernetConfigV1Alpha1) DeepCopy() *EthernetConfigV1Alpha1 {
	var cp EthernetConfigV1Alpha1 = *o
	if o.FeaturesConfig != nil {
		cp.FeaturesConfig = make(map[string]bool, len(o.FeaturesConfig))
		for k2, v2 := range o.FeaturesConfig {
			cp.FeaturesConfig[k2] = v2
		}
	}
	if o.RingsConfig != nil {
		cp.RingsConfig = new(EthernetRingsConfig)
		*cp.RingsConfig = *o.RingsConfig
		if o.RingsConfig.RX != nil {
			cp.RingsConfig.RX = new(uint32)
			*cp.RingsConfig.RX = *o.RingsConfig.RX
		}
		if o.RingsConfig.TX != nil {
			cp.RingsConfig.TX = new(uint32)
			*cp.RingsConfig.TX = *o.RingsConfig.TX
		}
		if o.RingsConfig.RXMini != nil {
			cp.RingsConfig.RXMini = new(uint32)
			*cp.RingsConfig.RXMini = *o.RingsConfig.RXMini
		}
		if o.RingsConfig.RXJumbo != nil {
			cp.RingsConfig.RXJumbo = new(uint32)
			*cp.RingsConfig.RXJumbo = *o.RingsConfig.RXJumbo
		}
		if o.RingsConfig.RXBufLen != nil {
			cp.RingsConfig.RXBufLen = new(uint32)
			*cp.RingsConfig.RXBufLen = *o.RingsConfig.RXBufLen
		}
		if o.RingsConfig.CQESize != nil {
			cp.RingsConfig.CQESize = new(uint32)
			*cp.RingsConfig.CQESize = *o.RingsConfig.CQESize
		}
		if o.RingsConfig.TXPush != nil {
			cp.RingsConfig.TXPush = new(bool)
			*cp.RingsConfig.TXPush = *o.RingsConfig.TXPush
		}
		if o.RingsConfig.RXPush != nil {
			cp.RingsConfig.RXPush = new(bool)
			*cp.RingsConfig.RXPush = *o.RingsConfig.RXPush
		}
		if o.RingsConfig.TXPushBufLen != nil {
			cp.RingsConfig.TXPushBufLen = new(uint32)
			*cp.RingsConfig.TXPushBufLen = *o.RingsConfig.TXPushBufLen
		}
		if o.RingsConfig.TCPDataSplit != nil {
			cp.RingsConfig.TCPDataSplit = new(bool)
			*cp.RingsConfig.TCPDataSplit = *o.RingsConfig.TCPDataSplit
		}
	}
	return &cp
}

// DeepCopy generates a deep copy of *RuleConfigV1Alpha1.
func (o *RuleConfigV1Alpha1) DeepCopy() *RuleConfigV1Alpha1 {
	var cp RuleConfigV1Alpha1 = *o
	if o.PortSelector.Ports != nil {
		cp.PortSelector.Ports = make([]PortRange, len(o.PortSelector.Ports))
		copy(cp.PortSelector.Ports, o.PortSelector.Ports)
	}
	if o.Ingress != nil {
		cp.Ingress = make([]IngressRule, len(o.Ingress))
		copy(cp.Ingress, o.Ingress)
	}
	return &cp
}
