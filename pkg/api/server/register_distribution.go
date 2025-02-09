package server

import (
	"github.com/hanks177/podman/v4/pkg/api/handlers/compat"
	"github.com/gorilla/mux"
)

func (s *APIServer) registerDistributionHandlers(r *mux.Router) error {
	r.HandleFunc(VersionedPath("/distribution/{name}/json"), compat.UnsupportedHandler)
	// Added non version path to URI to support docker non versioned paths
	r.HandleFunc("/distribution/{name}/json", compat.UnsupportedHandler)
	return nil
}
