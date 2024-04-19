package traefik_oidc_secure

import (
	"fmt"
	"testing"
)

func TestGetOIDCDiscovery(t *testing.T) {
	// Replace "https://your-openid-provider" with the actual URL of your OpenID provider
	providerURL := "http://auth.fracty.in/application/o/fracty-admin/"

	// Call the GetOIDCDiscovery function to retrieve OIDC discovery endpoints
	discovery, err := GetOIDCDiscovery(providerURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Print the discovered OIDC endpoints
	fmt.Printf("Issuer: %s\n", discovery.Issuer)
	fmt.Printf("Authorization Endpoint: %s\n", discovery.AuthorizationEndpoint)
	fmt.Printf("Token Endpoint: %s\n", discovery.TokenEndpoint)
	fmt.Printf("JWKS URI: %s\n", discovery.JWKSURI)
}
