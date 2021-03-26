package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VaultNamespaceSpec defines the desired state of VaultNamespace
type VaultNamespaceSpec struct {
	// NamespaceBase specifies the base path of the namespace. Use "root" for root or no namespace.
	NamespaceBase string `json:"namespaceBase,omitempty" yaml:"namespaceBase,omitempty"`
	NamespaceName string `json:"namespaceName,omitempty" yaml:"namespaceName,omitempty"`
}

// VaultNamespaceStatus defines the observed state of VaultNamespace
type VaultNamespaceStatus struct {
}

// +kubebuilder:object:root=true

// VaultNamespace is the Schema for the vaultnamespaces API
type VaultNamespace struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VaultNamespaceSpec   `json:"spec,omitempty"`
	Status VaultNamespaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VaultNamespaceList contains a list of VaultNamespace
type VaultNamespaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VaultNamespace `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VaultNamespace{}, &VaultNamespaceList{})
}
