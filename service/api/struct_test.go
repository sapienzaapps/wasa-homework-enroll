package api

import "testing"

func TestStudentEnrollRequest_Valid(t *testing.T) {
	t.Run("empty", func(t *testing.T) {
		var s StudentEnrollRequest
		if s.Valid() {
			t.Fatal("StudentEnrollRequest empty validated")
		}
	})

	t.Run("invalid ID - zero", func(t *testing.T) {
		var s = StudentEnrollRequest{
			StudentID: 0,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "doe.1234567@studenti.uniroma1.it",
			RepoURL:   "git@github.com:jdoe/homework.git",
		}
		if s.Valid() {
			t.Fatal("Student ID zero accepted")
		}
	})

	t.Run("invalid ID - outside range (too small)", func(t *testing.T) {
		var s = StudentEnrollRequest{
			StudentID: -33,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "doe.1234567@studenti.uniroma1.it",
			RepoURL:   "git@github.com:jdoe/homework.git",
		}
		if s.Valid() {
			t.Fatal("Negative Student ID accepted")
		}
	})

	t.Run("invalid ID - outside range (too big)", func(t *testing.T) {
		var s = StudentEnrollRequest{
			StudentID: 99999999,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "doe.1234567@studenti.uniroma1.it",
			RepoURL:   "git@github.com:jdoe/homework.git",
		}
		if s.Valid() {
			t.Fatal("Big Student ID accepted")
		}
	})

	t.Run("missing first name", func(t *testing.T) {
		var s = StudentEnrollRequest{
			StudentID: 5,
			LastName:  "Doe",
			Email:     "doe.1234567@studenti.uniroma1.it",
			RepoURL:   "git@github.com:jdoe/homework.git",
		}
		if s.Valid() {
			t.Fatal("First Name empty accepted")
		}
	})

	t.Run("first name short", func(t *testing.T) {
		var s = StudentEnrollRequest{
			StudentID: 5,
			FirstName: "A",
			LastName:  "Doe",
			Email:     "doe.1234567@studenti.uniroma1.it",
			RepoURL:   "git@github.com:jdoe/homework.git",
		}
		if s.Valid() {
			t.Fatal("Short First Name accepted")
		}
	})

	t.Run("first name long", func(t *testing.T) {
		var s = StudentEnrollRequest{
			StudentID: 5,
			FirstName: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			LastName:  "Doe",
			Email:     "doe.1234567@studenti.uniroma1.it",
			RepoURL:   "git@github.com:jdoe/homework.git",
		}
		if s.Valid() {
			t.Fatal("Long First Name accepted")
		}
	})

	t.Run("missing last name", func(t *testing.T) {
		var s = StudentEnrollRequest{
			StudentID: 5,
			FirstName: "John",
			Email:     "doe.1234567@studenti.uniroma1.it",
			RepoURL:   "git@github.com:jdoe/homework.git",
		}
		if s.Valid() {
			t.Fatal("Missing Last Name accepted")
		}
	})

	t.Run("last name short", func(t *testing.T) {
		var s = StudentEnrollRequest{
			StudentID: 5,
			FirstName: "John",
			LastName:  "A",
			Email:     "doe.1234567@studenti.uniroma1.it",
			RepoURL:   "git@github.com:jdoe/homework.git",
		}
		if s.Valid() {
			t.Fatal("Short Last Name accepted")
		}
	})

	t.Run("last name long", func(t *testing.T) {
		var s = StudentEnrollRequest{
			StudentID: 5,
			FirstName: "John",
			LastName:  "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA",
			Email:     "doe.1234567@studenti.uniroma1.it",
			RepoURL:   "git@github.com:jdoe/homework.git",
		}
		if s.Valid() {
			t.Fatal("Long Last Name accepted")
		}
	})

	t.Run("empty email", func(t *testing.T) {
		var s = StudentEnrollRequest{
			StudentID: 5,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "",
			RepoURL:   "git@github.com:jdoe/homework.git",
		}
		if s.Valid() {
			t.Fatal("empty email accepted")
		}
	})

	t.Run("empty email", func(t *testing.T) {
		var s = StudentEnrollRequest{
			StudentID: 5,
			FirstName: "John",
			LastName:  "Doe",
			RepoURL:   "git@github.com:jdoe/homework.git",
		}
		if s.Valid() {
			t.Fatal("empty email accepted")
		}
	})

	t.Run("wrong email", func(t *testing.T) {
		var s = StudentEnrollRequest{
			StudentID: 5,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "sfsfdfdsfsdfsdf",
			RepoURL:   "git@github.com:jdoe/homework.git",
		}
		if s.Valid() {
			t.Fatal("wrong email accepted")
		}
	})

	t.Run("email on gmail", func(t *testing.T) {
		var s = StudentEnrollRequest{
			StudentID: 5,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "jdoe@gmail.com",
			RepoURL:   "git@github.com:jdoe/homework.git",
		}
		if s.Valid() {
			t.Fatal("email on gmail accepted")
		}
	})

	t.Run("repo url empty", func(t *testing.T) {
		var s = StudentEnrollRequest{
			StudentID: 5,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "doe.1234567@studenti.uniroma1.it",
		}
		if s.Valid() {
			t.Fatal("repo url empty accepted")
		}
	})

	t.Run("repo url invalid", func(t *testing.T) {
		var s = StudentEnrollRequest{
			StudentID: 5,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "doe.1234567@studenti.uniroma1.it",
			RepoURL:   "https://github.com/jdoe/homework",
		}
		if s.Valid() {
			t.Fatal("repo url invalid accepted")
		}
	})

	t.Run("good github", func(t *testing.T) {
		var s = StudentEnrollRequest{
			StudentID: 5,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "doe.1234567@studenti.uniroma1.it",
			RepoURL:   "git@github.com:jdoe/homework.git",
		}
		if !s.Valid() {
			t.Fatal("good request rejected")
		}
	})

	t.Run("good gitlab", func(t *testing.T) {
		var s = StudentEnrollRequest{
			StudentID: 5,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "doe.1234567@studenti.uniroma1.it",
			RepoURL:   "git@gitlab.com:jdoe/homework.git",
		}
		if !s.Valid() {
			t.Fatal("good request rejected")
		}
	})

	t.Run("good bitbucket", func(t *testing.T) {
		var s = StudentEnrollRequest{
			StudentID: 5,
			FirstName: "John",
			LastName:  "Doe",
			Email:     "doe.1234567@studenti.uniroma1.it",
			RepoURL:   "ssh://git@bitbucket.com/jdoe/homework.git",
		}
		if !s.Valid() {
			t.Fatal("good request rejected")
		}
	})
}
