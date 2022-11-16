//go:build webui

package main

import (
	"fmt"
	"git.sapienzaapps.it/gamificationlab/wasa-homework-enroll/webui"
	"io/fs"
	"net/http"
	"strings"
)

func registerWebUI(hdl http.Handler) (http.Handler, error) {
	distDirectory, err := fs.Sub(webui.Dist, "dist")
	if err != nil {
		return nil, fmt.Errorf("error embedding WebUI dist/ directory: %w", err)
	}
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.RequestURI, "/dashboard/") {
			http.StripPrefix("/dashboard/", http.FileServer(http.FS(distDirectory))).ServeHTTP(w, r)
			return
		} else if r.RequestURI == "/" {
			// Redirect to dashboard
			http.Redirect(w, r, "/dashboard/", http.StatusTemporaryRedirect)
			return
		}
		hdl.ServeHTTP(w, r)
	}), nil
}
