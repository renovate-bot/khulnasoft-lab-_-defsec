package network

import defsecTypes "github.com/khulnasoft-lab/defsec/pkg/types"

type Network struct {
	ElasticLoadBalancers []ElasticLoadBalancer
	LoadBalancers        []LoadBalancer
	Routers              []Router
	VpnGateways          []VpnGateway
}

type NetworkInterface struct {
	Metadata     defsecTypes.Metadata
	NetworkID    defsecTypes.StringValue
	IsVipNetwork defsecTypes.BoolValue
}
