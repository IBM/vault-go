package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// AuthConfig for setting TTL
type AuthConfig struct {
	DefaultLeaseTTL int64 `json:"defaultLeaseTTL,omitempty" yaml:"defaultLeaseTTL,omitempty" vault:"default_lease_ttl"`
	MaxLeaseTTL     int64 `json:"maxLeaseTTL,omitempty" yaml:"maxLeaseTTL,omitempty" vault:"max_lease_ttl"`
}

// JWTAuthConfig is config for JWT
type JWTAuthConfig struct {
	// OIDCDiscoveryURL oidc_discovery_url (string: <optional>) - The OIDC Discovery URL, without any .well-known component (base path). Cannot be used with "jwks_url" or "jwt_validation_pubkeys".
	OIDCDiscoveryURL string `json:"oidcDiscoveryURL,omitempty" yaml:"oidcDiscoveryURL,omitempty" vault:"oidc_discovery_url,omitempty"`

	// OIDCDiscoveryCAPem oidc_discovery_ca_pem (string: <optional>) - The CA certificate or chain of certificates, in PEM format, to use to validate connections to the OIDC Discovery URL. If not set, system certificates are used.
	OIDCDiscoveryCAPem string `json:"oidcDiscoveryCAPem,omitempty" yaml:"oidcDiscoveryCAPem,omitempty" vault:"oidc_discovery_ca_pem,omitempty"`

	// OIDCClientID oidc_client_id (string: <optional>) - The OAuth Client ID from the provider for OIDC roles.
	OIDCClientID string `json:"oidcClientID,omitempty" yaml:"oidcClientID,omitempty" vault:"oidc_client_id,omitempty"`

	// OIDCClientSecret oidc_client_secret (string: <optional>) - The OAuth Client Secret from the provider for OIDC roles.
	OIDCClientSecret string `json:"oidcClientSecret,omitempty" yaml:"oidcClientSecret,omitempty" vault:"oidc_client_secret,omitempty"`

	// OIDCResponseMode oidc_response_mode (string: <optional>) - The response mode to be used in the OAuth2 request. Allowed values are "query" and "form_post". Defaults to "query".
	OIDCResponseMode string `json:"oidcResponseMode,omitempty" yaml:"oidcResponseMode,omitempty" vault:"oidc_response_mode,omitempty"`

	// OIDCResponseTypes oidc_response_types (comma-separated string, or array of strings: <optional>) - The response types to request. Allowed values are "code" and "id_token". Defaults to "code". Note: "id_token" may only be used if "oidc_response_mode" is set to "form_post".
	OIDCResponseTypes string `json:"oidcResponseTypes,omitempty" yaml:"oidcResponseTypes,omitempty" vault:"oidc_response_types,omitempty"`

	// JWKSURL jwks_url (string: <optional>) - JWKS URL to use to authenticate signatures. Cannot be used with "oidc_discovery_url" or "jwt_validation_pubkeys".
	JWKSURL string `json:"JWKSURL,omitempty" yaml:"JWKSURL,omitempty" vault:"jwks_url,omitempty"`

	// JWKSCAPem jwks_ca_pem (string: <optional>) - The CA certificate or chain of certificates, in PEM format, to use to validate connections to the JWKS URL. If not set, system certificates are used.
	JWKSCAPem string `json:"JWKSCAPem,omitempty" yaml:"JWKSCAPem,omitempty" vault:"jwks_ca_pem,omitempty"`

	// JWTValidationPubkeys jwt_validation_pubkeys (comma-separated string, or array of strings: <optional>) - A list of PEM-encoded public keys to use to authenticate signatures locally. Cannot be used with "jwks_url" or "oidc_discovery_url".
	JWTValidationPubkeys string `json:"jwtValidationPubkeys,omitempty" yaml:"jwtValidationPubkeys,omitempty" vault:"jwt_validation_pubkeys,omitempty"`

	// BoundIssuer bound_issuer (string: <optional>) - The value against which to match the iss claim in a JWT.
	BoundIssuer string `json:"boundIssuer,omitempty" yaml:"boundIssuer,omitempty" vault:"bound_issuer,omitempty"`

	// JWTSupportedAlgs jwt_supported_algs (comma-separated string, or array of strings: <optional>) - A list of supported signing algorithms. Defaults to [RS256]. (Available algorithms + EdDSA)
	JWTSupportedAlgs string `json:"jwtSupportedAlgs,omitempty" yaml:"jwtSupportedAlgs,omitempty" vault:"jwt_supported_algs,omitempty"`

	// DefaultRole default_role (string: <optional>) - The default role to use if none is provided during login.
	DefaultRole string `json:"defaultRole,omitempty" yaml:"defaultRole,omitempty" vault:"default_role,omitempty"`

	// ProviderConfig provider_config (map: <optional>) - Configuration options for provider-specific handling. Providers with specific handling include Azure; the options are described in each provider's section in OIDC Provider Setup
	ProviderConfig map[string]string `json:"providerConfig,omitempty" yaml:"providerConfig,omitempty" vault:"provider_config,omitempty"`
}

// AuthData is a struct for passing data to vault rest calls
type AuthData struct {
	Type        string     `json:"type" yaml:"type" vault:"type"`
	Description string     `json:"description" yaml:"description" vault:"description"`
	Config      AuthConfig `json:"config" yaml:"config" vault:"config"`
}

// VaultAuthSpec spec for auth
// '{"type":"approle","description":"","config":{"options":null,"default_lease_ttl":"0s","max_lease_ttl":"0s","force_no_cache":false},"local":false,"seal_wrap":false,"options":null}' http://127.0.0.1:8200/v1/sys/auth/foo
type VaultAuthSpec struct {
	Data           AuthData      `json:"data,omitempty" yaml:"data,omitempty"`
	VaultNamespace string        `json:"vaultNamespace,omitempty" yaml:"vaultNamespace,omitempty"`
	Path           string        `json:"path,omitempty" yaml:"path,omitempty"`
	JWTConfig      JWTAuthConfig `json:"jwtConfig,omitempty" yaml:"jwtConfig,omitempty"`
}

// VaultAuthStatus defines the observed state of VaultAuth
type VaultAuthStatus struct {
}

// +kubebuilder:object:root=true

// VaultAuth is the Schema for the vaultauths API
type VaultAuth struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VaultAuthSpec   `json:"spec,omitempty"`
	Status VaultAuthStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VaultAuthList contains a list of VaultAuth
type VaultAuthList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VaultAuth `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VaultAuth{}, &VaultAuthList{})
}
