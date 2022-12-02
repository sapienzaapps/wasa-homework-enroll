package api

import (
	"errors"
	"git.sapienzaapps.it/gamificationlab/wasa-homework-enroll/service/api/reqcontext"
	"git.sapienzaapps.it/gamificationlab/wasa-homework-enroll/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func (rt *_router) getVueLog(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	studentID, err := strconv.ParseUint(ps.ByName("studentid"), 10, 64)
	if err != nil {
		ctx.Logger.WithError(err).Warn("vue-log: error parsing student ID")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	openapiLog, err := rt.db.GetVueLog(int(studentID))
	if errors.Is(err, database.ErrUserDoesNotExists) {
		ctx.Logger.WithError(err).Error("vue-log: user does not exist")
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("vue-log: error getting vue log")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("content-type", "text/plain; charset=utf-8")
	_, _ = w.Write([]byte(openapiLog))
}
