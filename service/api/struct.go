package api

import (
	"git.sapienzaapps.it/gamificationlab/wasa-homework-enroll/service/database"
	"regexp"
	"time"
)

var (
	emailRx = regexp.MustCompile(`@studenti\.uniroma1\.it$`)
	repoRx  = regexp.MustCompile(`^(ssh://)?git@`)
)

type StudentEnrollRequest struct {
	StudentID int    `json:"studentID"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	RepoURL   string `json:"repoURL"`
}

func (s *StudentEnrollRequest) Valid() bool {
	return s.StudentID >= 1 && s.StudentID <= 9999999 &&
		len(s.FirstName) >= 2 && len(s.FirstName) <= 32 &&
		len(s.LastName) >= 2 && len(s.LastName) <= 32 &&
		len(s.Email) >= 25 && len(s.Email) <= 70 && emailRx.MatchString(s.Email) &&
		len(s.RepoURL) >= 5 && len(s.RepoURL) <= 200 && repoRx.MatchString(s.RepoURL)
}

type StudentEnrollResult struct {
	PublicKey string `json:"publicKey"`
}

type HomeworkResult struct {
	StudentID uint64    `json:"studentID"`
	Hash      string    `json:"hash"`
	OpenAPI   int       `json:"openAPI"`
	Go        int       `json:"go"`
	Vue       int       `json:"vue"`
	Docker    int       `json:"docker"`
	LastCheck time.Time `json:"lastCheck"`
}

func (h *HomeworkResult) FromDB(r database.HomeworkResult) {
	h.StudentID = r.StudentID
	h.Hash = r.Hash
	h.OpenAPI = r.OpenAPI
	h.Go = r.Go
	h.Vue = r.Vue
	h.Docker = r.Docker
	h.LastCheck = r.LastCheck
}
