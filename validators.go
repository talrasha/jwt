package jwt

import (
	"errors"
	"time"
)

var (
	// ErrAudValidation is the error for an invalid "aud" claim.
	ErrAudValidation = errors.New("jwt: aud claim is invalid")
	// ErrExpValidation is the error for an invalid "exp" claim.
	ErrExpValidation = errors.New("jwt: exp claim is invalid")
	// ErrIatValidation is the error for an invalid "iat" claim.
	ErrIatValidation = errors.New("jwt: iat claim is invalid")
	// ErrIssValidation is the error for an invalid "iss" claim.
	ErrIssValidation = errors.New("jwt: iss claim is invalid")
	// ErrJtiValidation is the error for an invalid "jti" claim.
	ErrJtiValidation = errors.New("jwt: jti claim is invalid")
	// ErrNbfValidation is the error for an invalid "nbf" claim.
	ErrNbfValidation = errors.New("jwt: nbf claim is invalid")
	// ErrSubValidation is the error for an invalid "sub" claim.
	ErrSubValidation = errors.New("jwt: sub claim is invalid")
)

// ValidatorFunc is a function for running extra
// validators when parsing a Payload string.
type ValidatorFunc func(*Payload) error

// AudienceValidator validates the "aud" claim.
// It checks if at least one of the audiences within the
// JWT's payload is listed in the server's audience whitelist.
func AudienceValidator(aud Audience) ValidatorFunc {
	return func(p *Payload) error {
		for _, serverAud := range aud {
			for _, clientAud := range p.Audience {
				if clientAud == serverAud {
					return nil
				}
			}
		}
		return ErrAudValidation
	}
}

// ExpirationTimeValidator validates the "exp" claim.
func ExpirationTimeValidator(now time.Time, validateZero bool) ValidatorFunc {
	return func(p *Payload) error {
		expint := p.ExpirationTime
		if !validateZero && expint == 0 {
			return nil
		}
		if exp := time.Unix(expint, 0); now.After(exp) {
			return ErrExpValidation
		}
		return nil
	}
}

// IssuedAtValidator validates the "iat" claim.
func IssuedAtValidator(now time.Time) ValidatorFunc {
	return func(p *Payload) error {
		if iat := time.Unix(p.IssuedAt, 0); now.Before(iat) {
			return ErrIatValidation
		}
		return nil
	}
}

// IssuerValidator validates the "iss" claim.
func IssuerValidator(iss string) ValidatorFunc {
	return func(p *Payload) error {
		if p.Issuer != iss {
			return ErrIssValidation
		}
		return nil
	}
}

// JWTIDValidator validates the "jti" claim.
func JWTIDValidator(jti string) ValidatorFunc {
	return func(p *Payload) error {
		if p.JWTID != jti {
			return ErrJtiValidation
		}
		return nil
	}
}

// NotBeforeValidator validates the "nbf" claim.
func NotBeforeValidator(now time.Time) ValidatorFunc {
	return func(p *Payload) error {

		if nbf := time.Unix(p.NotBefore, 0); now.Before(nbf) {
			return ErrNbfValidation
		}
		return nil
	}
}

// SubjectValidator validates the "sub" claim.
func SubjectValidator(sub string) ValidatorFunc {
	return func(p *Payload) error {
		if p.Subject != sub {
			return ErrSubValidation
		}
		return nil
	}
}
