/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package operation

import (
	"encoding/json"
	"time"

	"github.com/hyperledger/aries-framework-go/pkg/doc/verifiable"
)

// CreateCredentialRequest input data for edge service issuer rest api
type CreateCredentialRequest struct {
	Context []string           `json:"@context,omitempty"`
	Subject verifiable.Subject `json:"credentialSubject"`
	Issuer  verifiable.Issuer  `json:"issuer"`
	Type    []string           `json:"type,omitempty"`
	Profile string             `json:"profile,omitempty"`
}

// UpdateCredentialStatusRequest request struct for updating vc status
type UpdateCredentialStatusRequest struct {
	Credential   string `json:"credential"`
	Status       string `json:"status"`
	StatusReason string `json:"statusReason"`
}

// StoreVCRequest stores the credential with profile name
type StoreVCRequest struct {
	Profile    string `json:"profile"`
	Credential string `json:"credential"`
}

// ProfileRequest struct the input for creating profile
type ProfileRequest struct {
	Name                    string                             `json:"name"`
	URI                     string                             `json:"uri"`
	SignatureType           string                             `json:"signatureType"`
	SignatureRepresentation verifiable.SignatureRepresentation `json:"signatureRepresentation"`
	DID                     string                             `json:"did"`
	DIDPrivateKey           string                             `json:"didPrivateKey"`
	DIDKeyType              string                             `json:"didKeyType"`
	UNIRegistrar            UNIRegistrar                       `json:"uniRegistrar,omitempty"`
	DisableVCStatus         bool                               `json:"disableVCStatus"`
	OverwriteIssuer         bool                               `json:"overwriteIssuer,omitempty"`
}

// UNIRegistrar uni-registrar
type UNIRegistrar struct {
	DriverURL string            `json:"driverURL,omitempty"`
	Options   map[string]string `json:"options,omitempty"`
}

// VerifyCredentialResponse describes verify credential response
type VerifyCredentialResponse struct {
	Verified bool   `json:"verified"`
	Message  string `json:"message"`
}

// IssueCredentialRequest request for issuing credential.
type IssueCredentialRequest struct {
	Credential json.RawMessage         `json:"credential,omitempty"`
	Opts       *IssueCredentialOptions `json:"options,omitempty"`
}

// IssueCredentialOptions options for issuing credential.
type IssueCredentialOptions struct {
	// VerificationMethod is the URI of the verificationMethod used for the proof.
	// If omitted first ed25519 public key of DID (Issuer or Profile DID) will be used.
	VerificationMethod string `json:"verificationMethod,omitempty"`
	// AssertionMethod is verification method to be used for credential proof.
	// When provided along with 'VerificationMethod' property, 'VerificationMethod' takes precedence.
	// deprecated : to be removed in future, 'VerificationMethod' field will be used to pass verification method.
	AssertionMethod string `json:"assertionMethod,omitempty"`
	// ProofPurpose is purpose of the proof. If omitted "assertionMethod" will be used.
	ProofPurpose string `json:"proofPurpose,omitempty"`
	// Created date of the proof. If omitted system time will be used.
	Created *time.Time `json:"created,omitempty"`
	// Challenge is added to the proof
	Challenge string `json:"challenge,omitempty"`
	// Domain is added to the proof
	Domain string `json:"domain,omitempty"`
}

// ComposeCredentialRequest for composing and issuing credential.
type ComposeCredentialRequest struct {
	Issuer                  string          `json:"issuer,omitempty"`
	Subject                 string          `json:"subject,omitempty"`
	Types                   []string        `json:"types,omitempty"`
	IssuanceDate            *time.Time      `json:"issuanceDate,omitempty"`
	ExpirationDate          *time.Time      `json:"expirationDate,omitempty"`
	Claims                  json.RawMessage `json:"claims,omitempty"`
	Evidence                json.RawMessage `json:"evidence,omitempty"`
	TermsOfUse              json.RawMessage `json:"termsOfUse,omitempty"`
	CredentialFormat        string          `json:"credentialFormat,omitempty"`
	ProofFormat             string          `json:"proofFormat,omitempty"`
	CredentialFormatOptions json.RawMessage `json:"credentialFormatOptions,omitempty"`
	ProofFormatOptions      json.RawMessage `json:"proofFormatOptions,omitempty"`
}

// GenerateKeyPairResponse contains response from KMS generate keypair API.
type GenerateKeyPairResponse struct {
	PublicKey string `json:"publicKey,omitempty"`
	KeyID     string `json:"keyID,omitempty"`
}

// CredentialsVerificationRequest request for verifying credential.
type CredentialsVerificationRequest struct {
	Credential json.RawMessage                 `json:"verifiableCredential,omitempty"`
	Opts       *CredentialsVerificationOptions `json:"options,omitempty"`
}

// CredentialsVerificationOptions options for credential verifications.
type CredentialsVerificationOptions struct {
	Domain    string   `json:"domain,omitempty"`
	Challenge string   `json:"challenge,omitempty"`
	Checks    []string `json:"checks,omitempty"`
}

// CredentialsVerificationSuccessResponse resp when credential verification is success.
type CredentialsVerificationSuccessResponse struct {
	Checks []string `json:"checks,omitempty"`
}

// CredentialsVerificationFailResponse resp when credential verification is failed.
type CredentialsVerificationFailResponse struct {
	Checks []CredentialsVerificationCheckResult `json:"checks,omitempty"`
}

// CredentialsVerificationCheckResult resp containing failure check details.
type CredentialsVerificationCheckResult struct {
	Check              string `json:"check,omitempty"`
	Error              string `json:"error,omitempty"`
	VerificationMethod string `json:"verificationMethod,omitempty"`
}

// VerifyPresentationRequest request for verifying presentation.
type VerifyPresentationRequest struct {
	Presentation json.RawMessage            `json:"verifiablePresentation,omitempty"`
	Opts         *VerifyPresentationOptions `json:"options,omitempty"`
}

// VerifyPresentationOptions options for presentation verifications.
type VerifyPresentationOptions struct {
	Domain    string   `json:"domain,omitempty"`
	Challenge string   `json:"challenge,omitempty"`
	Checks    []string `json:"checks,omitempty"`
}

// VerifyPresentationSuccessResponse resp when presentation verification is success.
type VerifyPresentationSuccessResponse struct {
	Checks []string `json:"checks,omitempty"`
}

// VerifyPresentationFailureResponse resp when presentation verification is failed.
type VerifyPresentationFailureResponse struct {
	Checks []VerifyPresentationCheckResult `json:"checks,omitempty"`
}

// VerifyPresentationCheckResult resp containing failure check details.
type VerifyPresentationCheckResult struct {
	Check              string `json:"check,omitempty"`
	Error              string `json:"error,omitempty"`
	VerificationMethod string `json:"verificationMethod,omitempty"`
}

// ErrorResponse to send error message in the response
type ErrorResponse struct {
	Message string `json:"errMessage,omitempty"`
}

// HolderProfileRequest holder mode profile request
type HolderProfileRequest struct {
	Name                    string                             `json:"name"`
	SignatureType           string                             `json:"signatureType"`
	SignatureRepresentation verifiable.SignatureRepresentation `json:"signatureRepresentation"`
	DID                     string                             `json:"did"`
	DIDPrivateKey           string                             `json:"didPrivateKey"`
	DIDKeyType              string                             `json:"didKeyType"`
	UNIRegistrar            UNIRegistrar                       `json:"uniRegistrar,omitempty"`
}

// SignPresentationRequest request for signing a presentation.
type SignPresentationRequest struct {
	Presentation json.RawMessage          `json:"presentation,omitempty"`
	Opts         *SignPresentationOptions `json:"options,omitempty"`
}

// SignPresentationOptions options for signing a presentation.
type SignPresentationOptions struct {
	VerificationMethod string     `json:"verificationMethod,omitempty"`
	AssertionMethod    string     `json:"assertionMethod,omitempty"`
	ProofPurpose       string     `json:"proofPurpose,omitempty"`
	Created            *time.Time `json:"created,omitempty"`
	Challenge          string     `json:"challenge,omitempty"`
	Domain             string     `json:"domain,omitempty"`
}
