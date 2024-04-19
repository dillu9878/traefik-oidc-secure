package traefik_oidc_secure

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestServeHTTP(t *testing.T) {
	// Setup
	config := CreateConfig()
	config.ProviderURL = "http://auth.fracty.in/application/o/fracty-admin/"
	config.ClientID = "CYrfMfD4lggR1sWenXGPm5pWxbQij1dxOiLds0GJ"
	config.ClientSecret = "8Pe1zhlhLWmPB2aLcB6jJ5rrkUqhzcMKmEyE7IYz5AM51swuDbkkKwz5xa7UDPp0vNhJUtWyVCa25OSvqMpwfpgnXU4o5tcB6ePD0NIBTVYsh2QSYWRpKCNhr7zrVG4M"

	// Create a new instance of our middleware
	ProviderMiddleware, err := New(context.TODO(), http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusOK)
	}), config, "provider-openid")
	if err != nil {
		t.Fatal("Expected no error while creating OpenID Provider middleware, got:", err)
	}

	fmt.Printf("%+v\n", ProviderMiddleware)
	req, err := http.NewRequest("GET", "https://admin.fracty.in/", nil)
	if err != nil {
		t.Fatal("Expected no error while creating http request, got:", err)
	}

	rw := httptest.NewRecorder()

	// Test
	ProviderMiddleware.ServeHTTP(rw, req)

	fmt.Printf("%+v\n", rw)
	fmt.Printf("==>>>%+v\n", req)
}
