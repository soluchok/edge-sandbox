/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package operation

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/trustbloc/edge-core/pkg/log"

	"github.com/trustbloc/edge-sandbox/pkg/internal/common/support"
)

const (
	// api paths
	register      = "/register"
	createAccount = "/createAccount"
)

var logger = log.New("acrp-restapi")

// Handler http handler for each controller API endpoint.
type Handler interface {
	Path() string
	Method() string
	Handle() http.HandlerFunc
}

// Operation defines handlers.
type Operation struct {
	handlers      []Handler
	dashboardHTML string
	registerHTML  string
}

// Config config.
type Config struct {
	RegisterHTML  string
	DashboardHTML string
}

// New returns acrp operation instance.
func New(config *Config) *Operation {
	op := &Operation{
		dashboardHTML: config.DashboardHTML,
		registerHTML:  config.RegisterHTML,
	}

	op.registerHandler()

	return op
}

// registerHandler register handlers to be exposed from this service as REST API endpoints
func (o *Operation) registerHandler() {
	o.handlers = []Handler{
		support.NewHTTPHandler(register, http.MethodGet, o.register),
		support.NewHTTPHandler(createAccount, http.MethodPost, o.createAccount),
	}
}

// GetRESTHandlers get all controller API handler available for this service
func (o *Operation) GetRESTHandlers() []Handler {
	return o.handlers
}

func (o *Operation) register(w http.ResponseWriter, r *http.Request) {
	o.loadHTML(w, o.registerHTML, map[string]interface{}{})
}

func (o *Operation) createAccount(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		o.writeErrorResponse(w, http.StatusInternalServerError,
			fmt.Sprintf("unable to parse form data: %s", err.Error()))

		return
	}

	// TODO remove loggger and save the user data to db
	logger.Infof("username=%s", r.FormValue("username"))
	logger.Infof("password=%s", r.FormValue("password"))
	logger.Infof("nationalID=%s", r.FormValue("nationalID"))

	o.loadHTML(w, o.dashboardHTML, map[string]interface{}{
		"UserName": r.FormValue("username"),
	})
}

func (o *Operation) loadHTML(w http.ResponseWriter, htmlFileName string, data map[string]interface{}) {
	t, err := template.ParseFiles(htmlFileName)
	if err != nil {
		o.writeErrorResponse(w, http.StatusInternalServerError,
			fmt.Sprintf("unable to load html: %s", err.Error()))

		return
	}

	if err := t.Execute(w, data); err != nil {
		logger.Errorf("failed execute %s html template: %s", htmlFileName, err.Error())
	}
}

func (o *Operation) writeErrorResponse(rw http.ResponseWriter, status int, msg string) {
	logger.Errorf(msg)

	rw.WriteHeader(status)

	write := rw.Write
	if _, err := write([]byte(msg)); err != nil {
		logger.Errorf("Unable to send error message, %s", err)
	}
}
