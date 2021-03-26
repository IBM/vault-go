package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// VaultEndpointConfigCRL Specifies the duration for which the generated CRL should be marked valid. If the CRL is
// disabled, it will return a signed but zero-length CRL for any request. If enabled, it will re-build the CRL.
// https://www.vaultproject.io/api/secret/pki#set-crl-configuration
type VaultEndpointConfigCRL struct {
	// Expiry (string: "72h") – Specifies the time until expiration.
	Expiry string `json:"expiry" yaml:"expiry" vault:"expiry"`
	// Disable (bool: false) – Disables or enables CRL building.
	Disable string `json:"disable" yaml:"disable" vault:"disable"`
}

// VaultEndpointConfigURLs allows setting the issuing certificate endpoints, CRL distribution points, and OCSP server
// endpoints that will be encoded into issued certificates.
// https://www.vaultproject.io/api/secret/pki#set-urls
type VaultEndpointConfigURLs struct {
	//IssuingCertificates (array<string>: nil) – Specifies the URL values for the Issuing Certificate field.
	// This can be an array or a comma-separated string list.
	// https://www.vaultproject.io/api/secret/pki#set-urls
	IssuingCertificates []string `json:"issuingCertificates" yaml:"issuingCertificates" vault:"issuing_certificates"`
	// CRLDistributionPoints (array<string>: nil) – Specifies the URL values for the CRL Distribution Points
	// field. This can be an array or a comma-separated string list.
	CRLDistributionPoints []string `json:"crlDistributionPoints" yaml:"crlDistributionPoints" vault:"crl_distribution_points"`
	// OSCPServers (array<string>: nil) – Specifies the URL values for the OCSP Servers field. This can be an
	// array or a comma-separated string list.
	OSCPServers []string `json:"ocspServers" yaml:"ocspServers" vault:"ocsp_servers"`
}

// VaultGenerateOptions parameters to use when isuing a cert
// https://www.vaultproject.io/api/secret/pki#generate-intermediate
type VaultGenerateOptions struct {
	//	DN `json:"dn,omitempty" yaml:"dn,omitempty" vault:"dn,omitempty"`
	// start Distinguished Name ***
	// CommonName (string: <required>) – Specifies the requested CN for the certificate. If the CN is allowed by
	// role policy, it will be issued.
	CommonName string `json:"commonName,omitempty" yaml:"commonName,omitempty" vault:"common_name,omitempty"`
	// Country (string: "") – Specifies the C (Country) values in the subject field of issued certificates. This
	// is a comma-separated string or JSON array.
	Country []string `json:"country,omitempty" yaml:"country,omitempty" vault:"country,omitempty"`
	// Locality (string: "") – Specifies the L (Locality) values in the subject field of issued certificates.
	// This is a comma-separated string or JSON array.
	Locality []string `json:"locality,omitempty" yaml:"locality,omitempty" vault:"locality,omitempty"`
	// organization (string: "") – Specifies the O (Organization) values in the subject field of issued certificates.
	// This is a comma-separated string or JSON array.
	Organization []string `json:"organization,omitempty" yaml:"organization,omitempty" vault:"organization,omitempty"`
	// OU (string: "") – Specifies the OU (OrganizationalUnit) values in the subject field of issued certificates.
	// This is a comma-separated string or JSON array.
	OU []string `json:"ou,omitempty" yaml:"ou,omitempty" vault:"ou,omitempty"`
	// postalCode (string: "") – Specifies the Postal Code values in the subject field of issued certificates.
	// This is a comma-separated string or JSON array.
	PostalCode []string `json:"postalCode,omitempty" yaml:"postalCode,omitempty" vault:"postal_code,omitempty"`
	// Province (string: "") – Specifies the ST (Province) values in the subject field of issued certificates.
	// This is a comma-separated string or JSON array.
	Province []string `json:"province,omitempty" yaml:"province,omitempty" vault:"province,omitempty"`
	// SerialNumber (string: "") – Specifies the Serial Number, if any. Otherwise Vault will generate a random
	// serial for you. If you want more than one, specify alternative names in the alt_names map using OID 2.5.4.5.
	SerialNumber string `json:"serialNumber,omitempty" yaml:"serialNumber,omitempty" vault:"serial_number,omitempty"`
	// StreetAddress (string: "") – Specifies the Street Address values in the subject field of issued
	// certificates. This is a comma-separated string or JSON array.
	StreetAddress []string `json:"streetAddress,omitempty" yaml:"streetAddress,omitempty" vault:"street_address,omitempty"`
	// end Distinguished Name ***

	// AddBasicConstraints (bool) Whether to add a Basic Constraints
	// extension with CA: true. Only needed as a workaround in some compatibility
	// scenarios with Active Directory Certificate Services.
	AddBasicConstraints bool `json:"addBasicConstraints,omitempty" yaml:"addBasicConstraints,omitempty" vault:"add_basic_constraints,omitempty"`
	// AltNames (string: "") – Specifies requested Subject Alternative Names, in a comma-delimited list. These can
	// be host names or email addresses; they will be parsed into their respective fields. If any requested names
	// do not match role policy, the entire request will be denied.
	AltNames string `json:"altNames,omitempty" yaml:"altNames,omitempty" vault:"alt_names,omitempty"`
	// ExcludeCNFromSANs (bool: false) – If true, the given common_name will not be included in DNS or Email Subject
	// Alternate Names (as appropriate). Useful if the CN is not a hostname or email address, but is instead some
	// human-readable identifier.
	ExcludeCNFromSANs bool `json:"excludeCNFromSANs,omitempty" yaml:"excludeCNFromSANs,omitempty" vault:"exclude_cn_from_sans,omitempty"`
	// Format (string: "pem") – Specifies the format for returned data. This can be pem, der, or pem_bundle; defaults
	//  to pem. If der, the output is base64 encoded. If pem_bundle, the csr field will contain the private key (if exported)
	// and CSR, concatenated.
	Format string `json:"format,omitempty" yaml:"format,omitempty" vault:"format,omitempty"`
	// IPSans (string: "") – Specifies requested IP Subject Alternative Names, in a comma-delimited list. Only valid
	// if the role allows IP SANs (which is the default).
	IPSans string `json:"ipSans,omitempty" yaml:"ipSans,omitempty" vault:"ip_sans,omitempty"`
	// KeyBits (int: 2048) – Specifies the number of bits to use. This must be changed to a valid value if the key_type
	// is ec, e.g., 224 or 521.
	KeyBits string `json:"keyBits,omitempty" yaml:"keyBits,omitempty" vault:"key_bits,omitempty"`
	// KeyType (string: "rsa") – Specifies the desired key type; must be rsa or ec.
	KeyType string `json:"keyType,omitempty" yaml:"keyType,omitempty" vault:"key_type,omitempty"`
	// MaxTTL max_ttl (int64: 0) – Specifies the Time To Live value provided as seconds in int. If not set, uses the system default value or the value of max_ttl,
	// whichever is shorter.
	MaxTTL int64 `json:"maxTTL,omitempty" yaml:"maxTTL,omitempty" vault:"max_ttl,omitempty"`
	// OtherSans (string: "") – Specifies custom OID/UTF8-string SANs. These must match values specified on the role
	// in allowed_other_sans (see role creation for allowed_other_sans globbing rules). The format is the same as
	// OpenSSL: <oid>;<type>:<value> where the only current valid type is UTF8. This can be a comma-delimited list or a
	// JSON string slice.
	OtherSans string `json:"otherSans,omitempty" yaml:"otherSans,omitempty" vault:"other_sans,omitempty"`
	// PermittedDNSDomains (slice) Domains for which this certificate is allowed to sign or issue child certificates.
	// If set, all DNS names (subject and alt) on child certs must be exact matches or subsets of the given domains
	// (see https://tools.ietf.org/html/rfc5280#section-4.2.1.10).
	PermittedDNSDomains []string `json:"permittedDNSDomains,omitempty" yaml:"permittedDNSDomains,omitempty" vault:"permitted_dns_domains,omitempty"`
	// PrivateKeyFormat (string: "der") – Specifies the format for marshaling the private key. Defaults to der which will
	// return either base64-encoded DER or PEM-encoded DER, depending on the value of format. The other option is pkcs8
	// which will return the key marshalled as PEM-encoded PKCS8.
	PrivateKeyFormat string `json:"privateKeyFormat,omitempty" yaml:"privateKeyFormat,omitempty" vault:"private_key_format,omitempty"`
	// TTL ttl (int64: 0) – Specifies the Time To Live value provided as seconds in int. If not set, uses the system default value or the value of max_ttl,
	// whichever is shorter.
	TTL int64 `json:"ttl,omitempty" yaml:"ttl,omitempty" vault:"ttl,omitempty"`
	// URISans (string: "") – Specifies the requested URI Subject Alternative Names, in a comma-delimited list.
	URISans string `json:"uriSans,omitempty" yaml:"uriSans,omitempty" vault:"uri_sans,omitempty"`
}

// VaultMountTuneOptions Tune the default lease for the PKI secrets engine:
// https://www.vaultproject.io/api-docs/system/mounts#tune-mount-configuration
type VaultMountTuneOptions struct {
	// AllowedResponseHeaders (slice)
	// A list of headers to whitelist and allow a plugin to set on responses.
	AllowedResponseHeaders []string `json:"AllowedResponseHeaders" yaml:"AllowedResponseHeaders" vault:"allowed_response_headers"`
	// AuditNonHMACRequestKeyss (string: "") - Comma-separated string or list of keys that will not be HMAC'd
	// by audit devices in the request data object.
	AuditNonHMACRequestKeys string `json:"auditNonHMACRequestKeys" yaml:"auditNonHMACRequestKeys" vault:"audit-non-hmac-request-keys"`
	// AuditNonHMACResponseKeys (string: "") - Comma-separated string or list of keys that will not be HMAC'd
	// by audit devices in the response data object.
	AuditNonHMACResponseKeys string `json:"auditNonHMACResponseKeys" yaml:"auditNonHMACResponseKeys" vault:"audit-non-hmac-response-keys"`
	// DefaultLeaseTTL (duration: "") - The default lease TTL for this secrets engine. If unspecified, this
	// defaults to the Vault server's globally configured default lease TTL, or a previously configured value for
	// the secrets engine.
	DefaultLeaseTTL int64 `json:"defaultLeaseTTL,omitempty" yaml:"defaultLeaseTTL,omitempty" vault:"default_lease_ttl,omitempty"`
	// Description (string)
	// User-friendly description for this credential backend.
	Description string `json:"description,omitempty" yaml:"description,omitempty" vault:"description,omitempty"`
	// ForceNoCache
	ForceNoCache bool `json:"forceNoCache,omitempty" yaml:"forceNoCache,omitempty" vault:"force_no_cache,omitempty"`
	// ListingVisibility (string)
	// Determines the visibility of the mount in the UI-specific listing endpoint. Accepted value are 'unauth' and ''.
	ListingVisibility string `json:"listingVisibility,omitempty" yaml:"listingVisibility,omitempty" vault:"listingVisibility,omitempty"`
	// MaxLeaseTTL (duration: "") - The maximum lease TTL for this secrets engine. If unspecified, this defaults
	// to the Vault server's globally configured maximum lease TTL, or a previously configured value for the secrets engine.
	MaxLeaseTTL int64 `json:"maxLeaseTTL,omitempty" yaml:"maxLeaseTTL,omitempty" vault:"max_lease_ttl,omitempty"`
}

// VaultMountOptions the options for mounting secrets
// https://www.vaultproject.io/api-docs/system/mounts#enable-secrets-engine
type VaultMountOptions struct {
	// Type (string: <required>) – Specifies the type of the backend, such as "aws".
	Type string `json:"type,omitempty" yaml:"type,omitempty" vault:"type,omitempty"`
	// Description (string: "") – Specifies the human-friendly description of the mount.
	Description string `json:"description,omitempty" yaml:"description,omitempty" vault:"description,omitempty"`
	// Config (map<string|string>: nil) – Specifies configuration options for this mount; if set on a
	// specific mount, values will override any global defaults (e.g. the system TTL/Max TTL)
	// Local (bool: false) – Specifies if the secrets engine is a local mount only. Local mounts are not
	// replicated nor (if a secondary) removed by replication.
	Local bool `json:"local,omitempty" yaml:"local,omitempty" vault:"local,omitempty"`
	// SealWrap (bool: false) - Enable seal wrapping for the mount, causing values stored by the mount to
	// be wrapped by the seal's encryption capability.
	SealWrap bool `json:"sealWrap,omitempty" yaml:"sealWrap,omitempty" vault:"seal_wrap,omitempty"`
	// ExternalEntropyAccess (bool: false) - Enable the secrets engine to access Vault's external entropy source.
	ExternalEntropyAccess bool `json:"externalEntropyAccess,omitempty" yaml:"externalEntropyAccess,omitempty" vault:"external_entropy_access,omitempty"`
}

type IntermediateOptions struct {
	RootCANamespace string                `json:"rootCANamespace,omitempty" yaml:"rootCANamespace,omitempty"`
	RootCAPath      string                `json:"rootCAPath,omitempty" yaml:"rootCAPath,omitempty"`
	GenerateOptions *VaultGenerateOptions `json:"generateOptions,omitempty" yaml:"generateOptions,omitempty"`
}
type RootOptions struct {
	GenerateOptions *VaultGenerateOptions `json:"generateOptions,omitempty" yaml:"generateOptions,omitempty"`
}

// PKIConfig is the pki mount configuration
type PKIConfig struct {
	// ExportPrivateKey (bool: false) –
	// If true, the private key will be returned in the response;
	// if false the private key will not be returned and cannot be retrieved later.
	ExportPrivateKey    bool                     `json:"exportPrivateKey,omitempty" yaml:"exportPrivateKey,omitempty"`
	IntermediateOptions IntermediateOptions      `json:"intermediateOptions,omitempty" yaml:"intermediateOptions,omitempty"`
	RootOptions         RootOptions              `json:"rootOptions,omitempty" yaml:"rootOptions,omitempty"`
	URLs                *VaultEndpointConfigURLs `json:"urls,omitempty" yaml:"urls,omitempty"`
	CRL                 *VaultEndpointConfigCRL  `json:"crl,omitempty" yaml:"crl,omitempty"`
}

// VaultEndpointSpec defines the desired state of VaultEndpoint
type VaultEndpointSpec struct {
	VaultNamespace string `json:"vaultNamespace,omitempty" yaml:"vaultNamespace,omitempty"`
	Path           string `json:"path" yaml:"path"`
	//AllowedRoles      []string              `json:"allowedRoles" yaml:"allowedRoles"`
	MountOptions VaultMountOptions     `json:"mountOptions,omitempty" yaml:"mountOptions,omitempty"`
	TuneOptions  VaultMountTuneOptions `json:"tuneOptions,omitempty" yaml:"tuneOptions,omitempty"`
	PKIConfig    PKIConfig             `json:"pkiConfig,omitempty" yaml:"pkiConfig,omitempty"`
}

// VaultEndpointStatus defines the observed state of VaultEndpoint
type VaultEndpointStatus struct {
}

// +kubebuilder:object:root=true

// VaultEndpoint is the Schema for the vaultendpoints API
type VaultEndpoint struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VaultEndpointSpec   `json:"spec,omitempty"`
	Status VaultEndpointStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VaultEndpointList contains a list of VaultEndpoint
type VaultEndpointList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VaultEndpoint `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VaultEndpoint{}, &VaultEndpointList{})
}
