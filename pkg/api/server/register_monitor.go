package server

import (
	"github.com/hanks177/podman/v4/pkg/api/handlers/compat"
	"github.com/gorilla/mux"
)

func (s *APIServer) registerMonitorHandlers(r *mux.Router) error {
	r.Handle(VersionedPath("/monitor"), s.APIHandler(compat.UnsupportedHandler))
	// Added non version path to URI to support docker non versioned paths
	r.Handle("/monitor", s.APIHandler(compat.UnsupportedHandler))
	return nil
}
