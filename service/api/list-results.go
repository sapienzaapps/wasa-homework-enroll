package api

import (
	"encoding/json"
	"git.sapienzaapps.it/gamificationlab/wasa-homework-enroll/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sort"
)

// listResults returns the list of all students with their results
func (rt *_router) listResults(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	dbresults, err := rt.db.ListResults()
	if err != nil {
		ctx.Logger.WithError(err).Error("list-results: can't get the list of results")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var output = make([]HomeworkResult, 0, len(dbresults))
	for _, dbresult := range dbresults {
		var h HomeworkResult
		h.FromDB(dbresult)
		output = append(output, h)
	}

	sort.Slice(output, func(i, j int) bool {
		return output[i].StudentID < output[j].StudentID
	})

	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(output)
}
