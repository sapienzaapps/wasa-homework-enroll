package api

import "regexp"

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
