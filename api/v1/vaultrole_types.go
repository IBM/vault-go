package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VaultRoleData will be used as data in the vault api POST request
type VaultRoleData struct {
	// k8s only audience (string: "") - Optional Audience claim to verify in the JWT.
	Audience string `json:"audience,omitempty" yaml:"audience,omitempty" vault:"audience"`
	// k8s only bound_service_account_names (array: <required>) - List of service account names able to access this role. If set to "*" all names are allowed, both this and bound_service_account_namespaces can not be "*".
	BoundServiceAccountNames []string `json:"boundServiceAccountNames,omitempty" yaml:"boundServiceAccountNames,omitempty" vault:"bound_service_account_names"`
	// k8s only bound_service_account_namespaces (array: <required>) - List of namespaces allowed to access this role. If set to "*" all namespaces are allowed, both this and bound_service_account_names can not be set to "*".
	BoundServiceAccountNamespaces []string `json:"boundServiceAccountNamespaces,omitempty" yaml:"boundServiceAccountNamespaces,omitempty" vault:"bound_service_account_namespaces"`
	// bindSecretID: true
	BindSecretID bool `json:"bindSecretID,omitempty" yaml:"bindSecretID,omitempty" vault:"bind_secret_id,omitempty"`
	// localSecretIDs: false
	LocalSecretIDs bool `json:"localSecretIDs,omitempty" yaml:"localSecretIDs,omitempty" vault:"local_secret_ids,omitempty"`
	// policies:
	// - ddt-deploy
	// - kube-deploy
	// - auth-k8s-config
	Policies []string `json:"policies,omitempty" yaml:"policies,omitempty" vault:"policies,omitempty"`
	// secretIDBoundCIDRs: null
	SecretIDBoundCIDRs []string `json:"secretIDBoundCIDRs,omitempty" yaml:"secretIDBoundCIDRs,omitempty" vault:"secret_id_bound_cidrs,omitempty"`
	// secretIDNumUses: 0
	SecretIDNumUses int `json:"secretIDNumUses,omitempty" yaml:"secretIDNumUses,omitempty" vault:"secret_id_num_uses,omitempty"`
	// secretIDTTL: 0
	SecretIDTTL int `json:"secretIDTTL,omitempty" yaml:"secretIDTTL,omitempty" vault:"secret_id_ttl,omitempty"`
	// tokenBoundCIDRs: []
	TokenBoundCIDRs []string `json:"tokenBoundCIDRs,omitempty" yaml:"tokenBoundCIDRs,omitempty" vault:"token_bound_cidrs,omitempty"`
	// tokenExplicitMaxTTL: 0
	TokenExplicitMaxTTL int `json:"tokenExplicitMaxTTL,omitempty" yaml:"tokenExplicitMaxTTL,omitempty" vault:"token_explicit_max_ttl,omitempty"`
	// tokenMaxTTL: 1800
	TokenMaxTTL int64 `json:"tokenMaxTTL,omitempty" yaml:"tokenMaxTTL,omitempty" vault:"token_max_ttl,omitempty"`
	// tokenNoDefaultPolicy: false
	TokenNoDefaultPolicy bool `json:"tokenNoDefaultPolicy,omitempty" yaml:"tokenNoDefaultPolicy,omitempty" vault:"token_no_default_policy,omitempty"`
	// tokenNumUses: 0
	TokenNumUses int `json:"tokenNumUses,omitempty" yaml:"tokenNumUses,omitempty" vault:"token_num_uses,omitempty"`
	// tokenPeriod: 0
	TokenPeriod int `json:"tokenPeriod,omitempty" yaml:"tokenPeriod,omitempty" vault:"token_period,omitempty"`
	// tokenPolicies:
	// - ddt-deploy
	// - kube-deploy
	// - auth-k8s-config
	TokenPolicies []string `json:"tokenPolicies,omitempty" yaml:"tokenPolicies,omitempty" vault:"token_policies,omitempty"`
	// tokenTTL: 1800
	TokenTTL int64 `json:"tokenTTL,omitempty" yaml:"tokenTTL,omitempty" vault:"token_ttl,omitempty"`
	// tokenType: default
	TokenType string `json:"tokenType,omitempty" yaml:"tokenType,omitempty" vault:"token_type,omitempty"`
}

// VaultRoleSpec defines the desired state of VaultRole
type VaultRoleSpec struct {
	// Foo is an example field of VaultRole. Edit VaultRole_types.go to remove/update
	AuthMethod     string        `json:"authMethod,omitempty" yaml:"authMethod,omitempty"`
	RoleName       string        `json:"roleName,omitempty" yaml:"roleName,omitempty"`
	VaultNamespace string        `json:"vaultNamespace,omitempty" yaml:"vaultNamespace,omitempty"`
	Data           VaultRoleData `json:"data,omitempty" yaml:"data,omitempty"`
}

// VaultRoleStatus defines the observed state of VaultRole
type VaultRoleStatus struct {
}

// +kubebuilder:object:root=true

// VaultRole is the Schema for the vaultroles API
type VaultRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VaultRoleSpec   `json:"spec,omitempty"`
	Status VaultRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VaultRoleList contains a list of VaultRole
type VaultRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VaultRole `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VaultRole{}, &VaultRoleList{})
}
