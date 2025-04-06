package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/meshenka/nimble/internal/log"
)

func writeJSON(ctx context.Context, w http.ResponseWriter, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Ctx(ctx).Error("could not write response body", log.Err(err))
		panic(http.ErrAbortHandler)
	}
}
