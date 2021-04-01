package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// KVKey is a key in the kv secret path
type KVKey struct {
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
}

// KVPath a path in the kv secret engine
type KVPath struct {
	Path string  `json:"path,omitempty" yaml:"path,omitempty"`
	Keys []KVKey `json:"keys,omitempty" yaml:"keys,omitempty"`
}

// PKIConfig a path in the pki secret engine
type PKIACL struct {
	Path     string `json:"path,omitempty" yaml:"path,omitempty"`
	RoleName string `json:"roleName,omitempty" yaml:"roleName,omitempty"`
}

type SecretTypeEnum string

const (
	SecretTypeNone SecretTypeEnum = ""
	SecretTypeKVV2 SecretTypeEnum = "kv-v2"
	SecretTypePKI  SecretTypeEnum = "pki"
)

// SecretMetaSpec defines the desired state of SecretMeta
type SecretMetaSpec struct {
	// Deletable means that it is ok to remove any secrets and afterwards to remove the documentation. Inactive should be true for clarity.
	Deletable bool `json:"deletable" yaml:"deletable"`
	// DocURL location of external documentation
	DocURL string `json:"docURL" yaml:"docURL"`
	// Inactive means that secrets can be deleted from vault without breaking anything.
	Inactive bool `json:"inactive" yaml:"inactive"`
	// KVPath describes the kv path in vault minus the taxonomy
	KVPath KVPath `json:"kvPath,omitempty" yaml:"kvPath,omitempty"`
	// OwnerName is the Team owner name of this secret
	OwnerName string `json:"ownerName" yaml:"ownerName"`
	// Notes is a notes field
	Notes string `json:"notes" yaml:"notes"`
	// PKIACL is holds the role and path of the endpoint
	PKIACL PKIACL `json:"pkiACL,omitempty" yaml:"pkiACL,omitempty"`
	// Purpose is the reason the secret exists
	// Purpose a text field for reports
	Purpose string `json:"purpose" yaml:"purpose"`
	// Type is a SecretTypeEnum
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

// SecretMetaStatus defines the observed state of SecretMeta
type SecretMetaStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// SecretMeta is the Schema for the secretmeta API
type SecretMeta struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   SecretMetaSpec   `json:"spec,omitempty"`
	Status SecretMetaStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// SecretMetaList contains a list of SecretMeta
type SecretMetaList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []SecretMeta `json:"items"`
}

func init() {
	SchemeBuilder.Register(&SecretMeta{}, &SecretMetaList{})
}
