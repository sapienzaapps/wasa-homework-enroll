package api

import (
	"encoding/json"
	"git.sapienzaapps.it/gamificationlab/wasa-homework-enroll/service/api/reqcontext"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"sort"
)

var ramCache = new(InMemoryCache)

// listResults returns the list of all students with their results
func (rt *_router) listResults(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var shouldUpdateCache = true

	dbresults, err := rt.db.ListResults()
	if err != nil {
		// Check if the in-memory cache has some content. If so, use it instead of returning 500
		ramCache.Lock()
		if len(ramCache.Cache) > 0 {
			dbresults = ramCache.Cache
			ramCache.Unlock()
			shouldUpdateCache = false
		} else {
			ramCache.Unlock()
			ctx.Logger.WithError(err).Error("list-results: can't get the list of results")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	if shouldUpdateCache {
		// Update the in-memory cache (only when the results is not from the cache itself), so we have something to show
		// when the upstream server is offline.
		ramCache.Lock()
		ramCache.Cache = dbresults
		ramCache.Unlock()
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
