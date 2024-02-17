package v1alpha1

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// Config is an object that holds global clabernetes config information. Note that this CR is
// expected to effectively be a global singleton -- that is, there should be only *one* of these,
// and it *must* be named `clabernetes` -- CRD metadata spec will enforce this (via x-validation
// rules).
// +k8s:openapi-gen=true
// +kubebuilder:validation:XValidation:rule=(self.metadata.name == 'clabernetes')
type Config struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConfigSpec   `json:"spec,omitempty"`
	Status ConfigStatus `json:"status,omitempty"`
}

// ConfigSpec is the spec for a Config resource.
type ConfigSpec struct {
	// Metadata holds "global" metadata -- that is, metadata that is applied to all objects created
	// by the clabernetes controller.
	// +optional
	Metadata ConfigMetadata `json:"metadata"`
	// InClusterDNSSuffix overrides the default in cluster dns suffix used when resolving services.
	// +optional
	InClusterDNSSuffix string `json:"inClusterDNSSuffix,omitempty"`
	// ImagePull holds configurations relevant to how clabernetes launcher pods handle pulling
	// images.
	// +optional
	ImagePull ConfigImagePull `json:"imagePull"`
	// Deployment holds clabernetes deployment related configuration settings.
	// +optional
	Deployment ConfigDeployment `json:"deployment"`
	// Naming holds the global override for the "naming" setting for Topology objects -- this
	// controls whether the Topology resources have the containerlab topology name as a prefix.
	// Of course this is ignored if a Topology sets its Naming field to something not "global".
	// +kubebuilder:validation:Enum=prefixed;non-prefixed
	// +kubebuilder:default=prefixed
	// +optional
	Naming string `json:"naming"`
}

// ConfigStatus is the status for a Config resource.
type ConfigStatus struct{}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// ConfigList is a list of Config objects.
type ConfigList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`

	Items []Config `json:"items"`
}
