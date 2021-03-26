package v1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// PKIRoleBackendConfig configures the role definition. Note that the allowed_domains, allow_subdomains,
// allow_glob_domains, and allow_any_name attributes are additive; between them nearly and across multiple roles
// nearly any issuing policy can be accommodated. server_flag, client_flag, and code_signing_flag are additive
// as well. If a client requests a certificate that is not allowed by the CN policy in the role, the request is
// denied.
type PKIRoleBackendConfig struct {
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

	// AllowAnyName (bool: false) – Specifies if clients can request any CN. Useful in some circumstances,
	// but make sure you understand whether it is appropriate for your installation before enabling it.
	AllowAnyName bool `json:"allowAnyName,omitempty" yaml:"allowAnyName,omitempty" vault:"allow_any_name,omitempty"`
	// AllowBareDomains (bool: false) – Specifies if clients can request certificates matching the value of the
	// actual domains themselves; e.g. if a configured domain set with allowed_domains is example.com, this
	// allows clients to actually request a certificate containing the name example.com as one of the DNS values
	// on the final certificate. In some scenarios, this can be considered a security risk.
	AllowBareDomains bool `json:"allowBareDomains,omitempty" yaml:"allowBareDomains,omitempty" vault:"allow_bare_domains,omitempty"`
	// AllowGlobDomains (bool: false) - Allows names specified in allowed_domains to contain glob patterns
	// (e.g. ftp*.example.com). Clients will be allowed to request certificates with names matching the glob
	// patterns.
	AllowGlobDomains bool `json:"allowGlobDomains,omitempty" yaml:"allowGlobDomains,omitempty" vault:"allow_glob_domains,omitempty"`
	// AllowIPSANs (bool: true) – Specifies if clients can request IP Subject Alternative Names. No authorization
	// checking is performed except to verify that the given values are valid IP addresses.
	AllowIPSANs bool `json:"allowIPSANs" yaml:"allowIPSANs" vault:"allow_ip_sans"`
	// AllowLocalhost (bool: true) – Specifies if clients can request certificates for localhost as one of the
	// requested common names. This is useful for testing and to allow clients on a single host to talk securely.
	AllowLocalhost bool `json:"allowLocalhost" yaml:"allowLocalhost" vault:"allow_localhost"`
	// AllowSubdomains (bool: false) – Specifies if clients can request certificates with CNs that are subdomains
	// of the CNs allowed by the other role options. This includes wildcard subdomains. For example, an
	// allowed_domains value of example.com with this option set to true will allow foo.example.com and
	// bar.example.com as well as *.example.com. This is redundant when using the allow_any_name option.
	AllowSubdomains bool `json:"allowSubdomains,omitempty" yaml:"allowSubdomains,omitempty" vault:"allow_subdomains"`
	// AllowedDomains (list: []) – Specifies the domains of the role. This is used with the allow_bare_domains
	// and allow_subdomains options.
	AllowedDomains []string `json:"allowedDomains,omitempty" yaml:"allowedDomains,omitempty" vault:"allowed_domains,omitempty"`
	// AllowedOtherSANs (string: "") – Defines allowed custom OID/UTF8-string SANs. This can be a comma-delimited
	// list or a JSON string slice, where each element has the same format as OpenSSL: <oid>;<type>:<value>, but
	// the only valid type is UTF8 or UTF-8. The value part of an element may be a * to allow any value with that
	// OID. Alternatively, specifying a single * will allow any other_sans input.
	AllowedOtherSANs string `json:"allowedOtherSANs,omitempty" yaml:"allowedOtherSANs,omitempty" vault:"allowed_other_sans,omitempty"`
	// AllowedURISANs (string: "") - Defines allowed URI Subject Alternative Names. No authorization checking is
	// performed except to verify that the given values are valid URIs. This can be a comma-delimited list or a
	// JSON string slice. Values can contain glob patterns (e.g. spiffe://hostname/*).
	AllowedURISANs []string `json:"allowedURISANs,omitempty" yaml:"allowedURISANs,omitempty" vault:"allowed_uri_sans,omitempty"`
	// BasicConstraintesValidForNonCA (bool: false) - Mark Basic Constraints valid when issuing non-CA certificates.
	BasicConstraintesValidForNonCA bool `json:"basicConstraintesValidForNonCA,omitempty" yaml:"basicConstraintesValidForNonCA,omitempty" vault:"basic_constraints_valid_for_non_ca,omitempty"`
	// ClientFlag (bool: true) – Specifies if certificates are flagged for client use.
	ClientFlag bool `json:"clientFlag" yaml:"clientFlag" vault:"client_flag"`
	// CodeSigningFlag (bool: false) – Specifies if certificates are flagged for code signing use.
	CodeSigningFlag bool `json:"codeSigningFlag,omitempty" yaml:"codeSigningFlag,omitempty" vault:"code_signing_flag,omitempty"`
	// yaml doenst understand annon structs
	// DN (DN) is a term that describes the identifying information in a certificate and is part of the certificate
	// itself. A certificate contains DN information for both the owner or requestor of the certificate (called
	// the Subject DN) and the CA that issues the certificate (called the Issuer DN).
	// DN `json:"dn,omitempty,omitempty" yaml:"dn,omitempty,omitempty" vault:"dn,omitempty"`
	// email_protection_flag (bool: false) – Specifies if certificates are flagged for email protection use.
	EmailProtectionFlag bool `json:"emailProtectionFlag,omitempty" yaml:"emailProtectionFlag,omitempty" vault:"email_protection_flag,omitempty"`
	// EnforceHostnames (bool: true) – Specifies if only valid host names are allowed for CNs, DNS SANs, and the
	// host part of email addresses.
	EnforceHostnames bool `json:"enforceHostnames" yaml:"enforceHostnames" vault:"enforce_hostnames"`
	// ExtKeyUsage (list: []) – Specifies the allowed extended key usage constraint on issued certificates. Valid
	// values can be found at https://golang.org/pkg/crypto/x509/#ExtKeyUsage - simply drop the ExtKeyUsage part of
	// the value. Values are not case-sensitive. To specify no key usage constraints, set this to an empty list.
	ExtKeyUsage []string `json:"extKeyUsage,omitempty" yaml:"extKeyUsage,omitempty" vault:"ext_key_usage,omitempty"`
	// ExtKeyUsageOIDs (string: "") - A comma-separated string or list of extended key usage oids.
	ExtKeyUsageOIDs []string `json:"extKeyUsageOIDs,omitempty" yaml:"extKeyUsageOIDs,omitempty" vault:"ext_key_usage_oids,omitempty"`
	// GenerateLease (bool: false) – Specifies if certificates issued/signed against this role will have Vault
	// leases attached to them. Certificates can be added to the CRL by vault revoke <lease_id> when certificates
	// are associated with leases. It can also be done using the pki/revoke endpoint. However, when lease
	// generation is disabled, invoking pki/revoke would be the only way to add the certificates to the CRL.
	GenerateLease bool `json:"generateLease,omitempty" yaml:"generateLease,omitempty" vault:"generate_lease,omitempty"`
	// KeyBits (int: 2048) – Specifies the number of bits to use for the generated keys. This will need to be
	// changed for ec keys, e.g., 224 or 521.
	KeyBits int `json:"keyBits,omitempty" yaml:"keyBits,omitempty" vault:"key_bits,omitempty"`
	// KeyType (string: "rsa") – Specifies the type of key to generate for generated private keys and the type
	// of key expected for submitted CSRs. Currently, rsa and ec are supported, or when signing CSRs any can be
	// specified to allow keys of either type and with any bit size (subject to > 1024 bits for RSA keys).
	KeyType string `json:"keyType,omitempty" yaml:"keyType,omitempty" vault:"key_type,omitempty"`
	// KeyUsage (list: ["DigitalSignature", "KeyAgreement", "KeyEncipherment"]) – Specifies the allowed key
	// usage constraint on issued certificates. Valid values can be found at
	// https://golang.org/pkg/crypto/x509/#KeyUsage - simply drop the KeyUsage part of the value. Values are not
	// case-sensitive. To specify no key usage constraints, set this to an empty list.
	KeyUsage []string `json:"keyUsage" yaml:"keyUsage" vault:"key_usage"`
	// MaxTTL (int: 0) – Specifies the maximum Time To Live provided as int number of seconds.
	// Hour is the largest suffix. If not set, defaults to the system maximum lease TTL.
	MaxTTL int64 `json:"maxTTL,omitempty" yaml:"maxTTL,omitempty" vault:"max_ttl,omitempty"`
	// NoStore (bool: false) – If set, certificates issued/signed against this role will not be stored in the
	// storage backend. This can improve performance when issuing large numbers of certificates. However,
	// certificates issued in this way cannot be enumerated or revoked, so this option is recommended only for
	// certificates that are non-sensitive, or extremely short-lived. This option implies a value of false for
	// generate_lease.
	NoStore bool `json:"noStore,omitempty" yaml:"noStore,omitempty" vault:"no_store,omitempty"`
	// NotBeforeDurationn (int: 30) – Specifies the seconds by which to backdate the NotBefore property.
	NotBeforeDuration int `json:"notBeforeDuration,omitempty" yaml:"notBeforeDuration,omitempty" vault:"not_before_duration,omitempty"`
	// PolicyIdentifiers (list: []) – A comma-separated string or list of policy OIDs.
	PolicyIdentifiers []string `json:"policyIdentifiers,omitempty" yaml:"policyIdentifiers,omitempty" vault:"policy_identifiers,omitempty"`
	// RequireCN (bool: true) - If set to false, makes the common_name field optional while generating a certificate.
	RequireCN bool `json:"requireCN" yaml:"requireCN" vault:"require_cn"`
	// TTL ttl (int64: 0) – Specifies the Time To Live value provided as seconds in int. If not set, uses the system default value or the value of max_ttl,
	// whichever is shorter.
	TTL int64 `json:"ttl,omitempty" yaml:"ttl,omitempty" vault:"ttl,omitempty"`
	// ServerFlag (bool: true) – Specifies if certificates are flagged for server use.
	ServerFlag bool `json:"serverFlag" yaml:"serverFlag" vault:"server_flag"`
	// UseCSRCommonName (bool: true) – When used with the CSR signing endpoint, the common name in the CSR will
	// be used instead of taken from the JSON data. This does not include any requested SANs in the CSR; use
	// use_csr_sans for that.
	UseCSRCommonName bool `json:"UseCSRCommonName" yaml:"UseCSRCommonName" vault:"use_csr_common_name"`
	// UseCSRSANs (bool: true) – When used with the CSR signing endpoint, the subject alternate names in the CSR
	// will be used instead of taken from the JSON data. This does not include the common name in the CSR; use
	// use_csr_common_name for that.
	UseCSRSANs bool `json:"UseCSRSANs" yaml:"UseCSRSANs" vault:"use_csr_sans"`
}

// PKIRoleSpec defines the desired state of PKIRole
type PKIRoleSpec struct {
	// Name (string: <required>) – Specifies the name of the role to create. This is part of the request URL.
	IssuerPath     string               `json:"issuerPath" yaml:"issuerPath"`
	RoleName       string               `json:"roleName" yaml:"roleName"`
	VaultNamespace string               `json:"vaultNamespace" yaml:"vaultNamespace"`
	Config         PKIRoleBackendConfig `json:"config" yaml:"config"`
}

// PKIRoleStatus defines the observed state of PKIRole
type PKIRoleStatus struct {
}

// +kubebuilder:object:root=true

// PKIRole is the Schema for the pkiroles API
type PKIRole struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   PKIRoleSpec   `json:"spec,omitempty"`
	Status PKIRoleStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// PKIRoleList contains a list of PKIRole
type PKIRoleList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []PKIRole `json:"items"`
}

func init() {
	SchemeBuilder.Register(&PKIRole{}, &PKIRoleList{})
}
