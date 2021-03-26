package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// JWTRoleParameters parameters to feen to jwt role endpoint
type JWTRoleParameters struct {
	// RoleType	role_type (string: <optional>) - Type of role, either "oidc" (default) or "jwt".
	RoleType string `json:"roleType,omitempty" yaml:"roleType,omitempty" vault:"role_type,omitempty"`

	// BoundAudiences bound_audiences (array: <optional>) - List of aud claims to match against. Any match is sufficient. Required for "jwt" roles, optional for "oidc" roles.
	BoundAudiences []string `json:"boundAudiences,omitempty" yaml:"boundAudiences,omitempty" vault:"bound_audiences,omitempty"`

	// UserClaim user_claim (string: <required>) - The claim to use to uniquely identify the user; this will be used as the name for the Identity entity alias created due to a successful login. The claim value must be a string.
	UserClaim string `json:"userClaim,omitempty" yaml:"userClaim,omitempty" vault:"user_claim,omitempty"`

	// ClockSkewLeeway clock_skew_leeway (int: <optional>) - The amount of leeway to add to all claims to account for clock skew, in seconds. Defaults to 60 seconds if set to 0 and can be disabled if set to -1. Only applicable with "jwt" roles.
	ClockSkewLeeway int64 `json:"clockSkewLeeway,omitempty" yaml:"clockSkewLeeway,omitempty" vault:"clock_skew_leeway,omitempty"`

	// ExpirationLeeway expiration_leeway (int: <optional>) - The amount of leeway to add to expiration (exp) claims to account for clock skew, in seconds. Defaults to 150 seconds if set to 0 and can be disabled if set to -1. Only applicable with "jwt" roles.
	ExpirationLeeway int64 `json:"expirationLeeway,omitempty" yaml:"expirationLeeway,omitempty" vault:"expiration_leeway,omitempty"`

	// NotBeforeLeeway not_before_leeway (int: <optional>) - The amount of leeway to add to not before (nbf) claims to account for clock skew, in seconds. Defaults to 150 seconds if set to 0 and can be disabled if set to -1. Only applicable with "jwt" roles.
	NotBeforeLeeway int64 `json:"notBeforeLeeway,omitempty" yaml:"notBeforeLeeway,omitempty" vault:"not_before_leeway,omitempty"`

	// BoundSubject bound_subject (string: <optional>) - If set, requires that the sub claim matches this value.
	BoundSubject string `json:"boundSubject,omitempty" yaml:"boundSubject,omitempty" vault:"bound_subject,omitempty"`

	// BoundClaims bound_claims (map: <optional>) - If set, a map of claims (keys) to match against respective claim values (values). The expected value may be a single string or a list of strings. The interpretation of the bound claim values is configured with bound_claims_type. Keys support JSON pointer syntax for referencing claims.
	BoundClaims map[string]string `json:"boundClaims,omitempty" yaml:"boundClaims,omitempty" vault:"bound_claims,omitempty"`

	// BoundClaimsType bound_claims_type (string: "string") - Configures the interpretation of the bound_claims values. If "string" (the default), the values will treated as string literals and must match exactly. If set to "glob", the values will be interpreted as globs, with * matching any number of characters.
	BoundClaimsType string `json:"boundClaimsType,omitempty" yaml:"boundClaimsType,omitempty" vault:"bound_claims_type,omitempty"`

	// GroupClaim groups_claim (string: <optional>) - The claim to use to uniquely identify the set of groups to which the user belongs; this will be used as the names for the Identity group aliases created due to a successful login. The claim value must be a list of strings. Supports JSON pointer syntax for referencing claims.
	GroupClaim string `json:"groupClaim,omitempty" yaml:"groupClaim,omitempty" vault:"groups_claim,omitempty"`

	// ClaimMappings claim_mappings (map: <optional>) - If set, a map of claims (keys) to be copied to specified metadata fields (values). Keys support JSON pointer syntax for referencing claims.
	ClaimMappings map[string]string `json:"claimMappings,omitempty" yaml:"claimMappings,omitempty" vault:"claim_mappings,omitempty"`

	// OIDCScopes oidc_scopes (list: <optional>) - If set, a list of OIDC scopes to be used with an OIDC role. The standard scope "openid" is automatically included and need not be specified.
	OIDCScopes []string `json:"oidcScopes,omitempty" yaml:"oidcScopes,omitempty" vault:"oidc_scopes,omitempty"`

	// AllowedRedirectURLs allowed_redirect_uris (list: <required>) - The list of allowed values for redirect_uri during OIDC logins.
	AllowedRedirectURLs []string `json:"allowedRedirectURLs,omitempty" yaml:"allowedRedirectURLs,omitempty" vault:"allowed_redirect_uris,omitempty"`

	// tokenTTL tokenTTLng (bool: false) - Log received OIDC tokens and claims when debug-level logging is active. Not recommended in production since sensitive information may be present in OIDC responses.
	VerboseOIDCLogging bool `json:"verboseOIDCLogging,omitempty" yaml:"verboseOIDCLogging,omitempty" vault:"verbose_oidc_logging,omitempty"`

	// TokenTTL  token_ttl (integer: 0  or string: "") - The incre mental lifetime for generat ed tokens. This current value of this will be referenced at renewal time.
	TokenTTL string `json:"tokenTTL,omitempty" yaml:"tokenTTL,omitempty" vault:"token_ttl,omitempty"`

	// TokenMaxTTL token_max_ttl (integer: 0 or string: "") - The maximum lifetime for generated tokens. This current value of this will be referenced at renewal time.
	TokenMaxTTL int64 `json:"tokenMaxTTL,omitempty" yaml:"tokenMaxTTL,omitempty" vault:"token_max_ttl,omitempty"`

	// TokenPolicies token_policies (array: [] or comma-delimited string: "") - List of policies to encode onto generated tokens. Depending on the auth method, this list may be supplemented by user/group/other values.
	TokenPolicies []string `json:"tokenPolicies,omitempty" yaml:"tokenPolicies,omitempty" vault:"token_policies,omitempty"`

	// TokenBoundCIDRs token_bound_cidrs (array: [] or comma-delimited string: "") - List of CIDR blocks; if set, specifies blocks of IP addresses which can authenticate successfully, and ties the resulting token to these blocks as well.
	TokenBoundCIDRs string `json:"tokenBoundCIDRs,omitempty" yaml:"tokenBoundCIDRs,omitempty" vault:"token_bound_cidrs,omitempty"`

	// TokenExplicitMaxTTL token_explicit_max_ttl (integer: 0 or string: "") - If set, will encode an explicit max TTL onto the token. This is a hard cap even if token_ttl and token_max_ttl would otherwise allow a renewal.
	TokenExplicitMaxTTL int64 `json:"tokenExplicitMaxTTL,omitempty" yaml:"tokenExplicitMaxTTL,omitempty" vault:"token_explicit_max_ttl,omitempty"`

	// TokenNoDefaultPolicy token_no_default_policy (bool: false) - If set, the default policy will not be set on generated tokens; otherwise it will be added to the policies set in token_policies.
	TokenNoDefaultPolicy bool `json:"tokenNoDefaultPolicy,omitempty" yaml:"tokenNoDefaultPolicy,omitempty" vault:"token_no_default_policy,omitempty"`

	// TokenNumUses token_num_uses (integer: 0) - The maximum number of times a generated token may be used (within its lifetime); 0 means unlimited. If you require the token to have the ability to create child tokens, you will need to set this value to 0.
	TokenNumUses int `json:"tokenNumUses,omitempty" yaml:"tokenNumUses,omitempty" vault:"token_num_uses,omitempty"`

	// TokenPeriod token_period (integer: 0 or string: "") - The period, if any, to set on the token.
	TokenPeriod int64 `json:"tokenPeriod,omitempty" yaml:"tokenPeriod,omitempty" vault:"token_period,omitempty"`

	// TokenType token_type (string: "") - The type of token that should be generated. Can be service, batch, or default to use the mount's tuned default (which unless changed will be service tokens). For token store roles, there are two additional possibilities: default-service and default-batch which specify the type to return unless the client requests a different type at generation time.
	TokenType string `json:"tokenType,omitempty" yaml:"tokenType,omitempty" vault:"token_type,omitempty"`
}

// JWTRoleSpec defines the desired state of JWTRole
type JWTRoleSpec struct {
	AuthPath       string            `json:"authPath" yaml:"authPath"`
	RoleName       string            `json:"roleName" yaml:"roleName"`
	VaultNamespace string            `json:"vaultNamespace" yaml:"vaultNamespace"`
	Parameters     JWTRoleParameters `json:"parameters" yaml:"parameters"`
}

// JWTRoleStatus defines the observed state of JWTRole
type JWTRoleStatus struct {
}

// +kubebuilder:object:root=true

// JWTRole is the Schema for the jwtroles API
type JWTRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   JWTRoleSpec   `json:"spec,omitempty"`
	Status JWTRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// JWTRoleList contains a list of JWTRole
type JWTRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []JWTRole `json:"items"`
}

func init() {
	SchemeBuilder.Register(&JWTRole{}, &JWTRoleList{})
}
