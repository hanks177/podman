package libpod

import (
	"net/http"

	"github.com/hanks177/podman/v4/libpod"
	"github.com/hanks177/podman/v4/pkg/api/handlers/utils"
	api "github.com/hanks177/podman/v4/pkg/api/types"
	"github.com/hanks177/podman/v4/pkg/domain/entities"
	"github.com/hanks177/podman/v4/pkg/domain/infra/abi"
	"github.com/gorilla/schema"
	"github.com/pkg/errors"
)

func CreateSecret(w http.ResponseWriter, r *http.Request) {
	var (
		runtime = r.Context().Value(api.RuntimeKey).(*libpod.Runtime)
		decoder = r.Context().Value(api.DecoderKey).(*schema.Decoder)
	)

	query := struct {
		Name       string            `schema:"name"`
		Driver     string            `schema:"driver"`
		DriverOpts map[string]string `schema:"driveropts"`
	}{
		// override any golang type defaults
	}
	opts := entities.SecretCreateOptions{}
	if err := decoder.Decode(&query, r.URL.Query()); err != nil {
		utils.Error(w, http.StatusBadRequest, errors.Wrapf(err, "failed to parse parameters for %s", r.URL.String()))
		return
	}

	opts.Driver = query.Driver
	opts.DriverOpts = query.DriverOpts

	ic := abi.ContainerEngine{Libpod: runtime}
	report, err := ic.SecretCreate(r.Context(), query.Name, r.Body, opts)
	if err != nil {
		utils.InternalServerError(w, err)
		return
	}
	utils.WriteResponse(w, http.StatusOK, report)
}
