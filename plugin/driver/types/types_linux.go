package types

import (
	"github.com/containernetworking/plugins/pkg/ns"

	terwayTypes "github.com/AliyunContainerService/terway/types"
)

type CheckConfig struct {
	DP DataPath

	RecordPodEvent

	NetNS ns.NetNS

	HostVETHName    string
	ContainerIfName string

	ContainerIPNet *terwayTypes.IPNetSet
	HostIPSet      *terwayTypes.IPNetSet
	GatewayIP      *terwayTypes.IPSet

	ENIIndex int32 // phy device
	TrunkENI bool
	MTU      int

	DefaultRoute bool
	MultiNetwork bool
}
