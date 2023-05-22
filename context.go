//
//
// @filename: context.go
// COPYRIGHT 2023 Krisna Pranav, NanoBlocksDevelopers
//
//

package nanoweb

import "net/http"

type Context struct {
	Request        *http.Request
	Response       http.ResponseWriter
	Params         Parameter
	statusCode     int
	App            *Application
	Session        Session
	IsSent         bool
	templateLoader string
}
