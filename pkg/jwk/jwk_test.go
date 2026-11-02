package jwk

import (
	"errors"
	"testing"

	"github.com/pluralsh/soft-serve/pkg/config"
)

func TestBadNewPair(t *testing.T) {
	_, err := NewPair(nil)
	if !errors.Is(err, config.ErrNilConfig) {
		t.Errorf("NewPair(nil) => %v, want %v", err, config.ErrNilConfig)
	}
}

func TestGoodNewPair(t *testing.T) {
	cfg := config.DefaultConfig()
	pair, err := NewPair(cfg)
	if err != nil {
		t.Fatalf("NewPair(cfg) => _, %v, want nil error", err)
	}

	jwk := pair.JWK()
	if !jwk.Valid() {
		t.Error("JWK() returned an invalid JSON Web Key")
	}
	if got, want := jwk.Algorithm, SigningMethod.Alg(); got != want {
		t.Errorf("JWK().Algorithm = %q, want %q", got, want)
	}
	if jwk.KeyID == "" {
		t.Error("JWK().KeyID is empty")
	}
}
