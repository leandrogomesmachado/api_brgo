// This is a wrapper file for the brgo implementation
// It allows the Go module system to recognize the package

package api

import (
	"net/http"
)

// ConfigurarRotas returns an HTTP handler that configures all API routes
// The actual implementation is in the .brgo files
func ConfigurarRotas() http.Handler {
	// This is just a stub - the real implementation is in rotas.brgo
	// This definition is here only to satisfy the Go compiler
	return nil
}
