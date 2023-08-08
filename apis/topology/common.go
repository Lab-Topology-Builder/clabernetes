package topology

// InsecureRegistries is a slice of strings of insecure registries to configure in the launcher
// pods.
type InsecureRegistries []string

// LinkEndpointElementCount defines the expected element count for a link endpoint slice.
const LinkEndpointElementCount = 2

// LinkEndpoint is a simple struct to hold node/interface name info for a given link.
type LinkEndpoint struct {
	// NodeName is the name of the node this link resides on.
	NodeName string `json:"nodeName"`
	// InterfaceName is the name of the interface on the node this link is on.
	InterfaceName string `json:"interfaceName"`
}

// Tunnel represents a VXLAN tunnel between clabernetes nodes (as configured by containerlab).
type Tunnel struct {
	// ID is the VXLAN ID (vnid) for the tunnel.
	ID int `json:"ID"`
	// LocalNodeName is the name of the local node for this tunnel.
	LocalNodeName string `json:"localNodeName"  yaml:"localNodeName"`
	// RemoteName is the name of the service to contact the remote end of the tunnel.
	RemoteName string `json:"remoteName"     yaml:"remoteName"`
	// RemoteNodeName is the name of the remote node.
	RemoteNodeName string `json:"remoteNodeName" yaml:"remoteNodeName"`
	// LocalLinkName is the local link name for the local end of the tunnel.
	LocalLinkName string `json:"localLinkName"  yaml:"localLinkName"`
	// RemoteLinkName is the remote link name for the remote end of the tunnel.
	RemoteLinkName string `json:"remoteLinkName" yaml:"remoteLinkName"`
}

// ExposedPorts holds information about exposed ports.
type ExposedPorts struct {
	// LoadBalancerAddress holds the address assigned to the load balancer exposing ports for a
	// given node.
	LoadBalancerAddress string `json:"loadBalancerAddress"`
	// TCPPorts is a list of TCP ports exposed on the LoadBalancer service.
	// +listType=set
	TCPPorts []int `json:"TCPPorts"`
	// UDPPorts is a list of UDP ports exposed on the LoadBalancer service.
	// +listType=set
	UDPPorts []int `json:"UDPPorts"`
}

// DeepCopyInto is a copy/pasta of autogenerated DeepCopyInto funcs for our ExposedPorts object.
func (in *ExposedPorts) DeepCopyInto(out *ExposedPorts) {
	*out = *in

	if in.TCPPorts != nil {
		in, out := &in.TCPPorts, &out.TCPPorts
		*out = make([]int, len(*in))
		copy(*out, *in)
	}

	if in.UDPPorts != nil {
		in, out := &in.UDPPorts, &out.UDPPorts
		*out = make([]int, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is a copy/pasta of autogenerated DeepCopy funcs for our ExposedPorts object.
func (in *ExposedPorts) DeepCopy() *ExposedPorts {
	if in == nil {
		return nil
	}

	out := new(ExposedPorts)

	in.DeepCopyInto(out)

	return out
}
