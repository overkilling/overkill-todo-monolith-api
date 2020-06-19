package pact

import (
	"os"
	"testing"

	_ "github.com/lib/pq"
	"github.com/pact-foundation/pact-go/dsl"
	"github.com/pact-foundation/pact-go/types"
)

func TestPactProvider(t *testing.T) {
	pactURL := os.Getenv("PACT_URL")
	if pactURL == "" {
		t.Skip("set PACT_URL to run this test")
	}
	providerURL := os.Getenv("PROVIDER_URL")
	if providerURL == "" {
		t.Skip("set PROVIDER_URL to run this test")
	}

	pact := createPact()
	_, err := pact.VerifyProvider(t, types.VerifyRequest{
		ProviderBaseURL: providerURL,
		PactURLs:        []string{pactURL},
	})

	if err != nil {
		t.Log("Pact test failed")
	}
}

func createPact() dsl.Pact {
	return dsl.Pact{
		Provider: "API",
	}
}
