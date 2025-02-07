package compat

import (
	"net/http"

	"github.com/hanks177/podman/v4/libpod"
	"github.com/hanks177/podman/v4/pkg/api/handlers/utils"
	api "github.com/hanks177/podman/v4/pkg/api/types"
)

func PauseContainer(w http.ResponseWriter, r *http.Request) {
	runtime := r.Context().Value(api.RuntimeKey).(*libpod.Runtime)

	// /{version}/containers/(name)/pause
	name := utils.GetName(r)
	con, err := runtime.LookupContainer(name)
	if err != nil {
		utils.ContainerNotFound(w, name, err)
		return
	}

	// the api does not error if the Container is already paused, so just into it
	if err := con.Pause(); err != nil {
		utils.InternalServerError(w, err)
		return
	}
	// Success
	utils.WriteResponse(w, http.StatusNoContent, nil)
}
