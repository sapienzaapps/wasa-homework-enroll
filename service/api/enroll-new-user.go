package api

import (
	"encoding/json"
	"errors"
	"git.sapienzaapps.it/gamificationlab/wasa-homework-enroll/service/api/reqcontext"
	"git.sapienzaapps.it/gamificationlab/wasa-homework-enroll/service/database"
	"git.sapienzaapps.it/gamificationlab/wasa-homework-enroll/service/sshkey"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

// enrollNewUser enrolls a new student in the system, and provides the public key.
func (rt *_router) enrollNewUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext) {
	var request StudentEnrollRequest
	err := json.NewDecoder(r.Body).Decode(&request)
	_ = r.Body.Close()
	if err != nil {
		ctx.Logger.WithError(err).Error("enroll: error decoding JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	} else if !request.Valid() {
		ctx.Logger.Error("enroll: error validating JSON")
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Generate a new key pair
	publicKey, privateKey, err := sshkey.GenerateSSHKeyPair()
	if err != nil {
		ctx.Logger.WithError(err).Error("enroll: error generating SSH key")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Create student in the DB
	err = rt.db.CreateStudent(request.StudentID, request.FirstName, request.LastName, request.Email, request.RepoURL, publicKey, privateKey)
	if errors.Is(err, database.ErrUserExists) {
		w.WriteHeader(http.StatusConflict)
		return
	} else if err != nil {
		ctx.Logger.WithError(err).Error("enroll: error creating user in DB")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Send the response
	var response = StudentEnrollResult{PublicKey: publicKey}
	w.Header().Set("content-type", "application/json")
	_ = json.NewEncoder(w).Encode(response)
}
