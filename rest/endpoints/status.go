package endpoints

import (
	"bytes"
	"encoding/json"
	"net/http"

	"git.appagile.io/services-paas/bb/vSweets/shop-backend-rest/info"
	"github.com/golang/glog"
	"github.com/labstack/echo"
)

//StatusEP endpoint implements the mandatory Grafana SimpleJson plugin endpoint
type StatusEP struct {
	Version   string
	BuildDate string
	Commit    string
}

// InitEndpoint to initialize get endpoint.
func (ps *StatusEP) InitEndpoint(e *echo.Echo, path string) error {

	var buffer bytes.Buffer

	buffer.WriteString(path)
	buffer.WriteString("v1/status")

	ps.Version = info.Version
	ps.BuildDate = info.BuildDate
	ps.Commit = info.GitCommit

	// Route => handler
	e.GET(buffer.String(), ps.GetHandlerFunc())
	glog.Info("Endpoint registered with path: ", buffer.String())

	return nil
}

// GetHandlerFunc returns the handler to deal with the rest request
func (ps *StatusEP) GetHandlerFunc() echo.HandlerFunc {
	return func(ctx echo.Context) error {

		values := ctx.QueryParams()
		glog.Info("Request contained the following values: ", values)

		resp, err := json.Marshal(ps)
		if err != nil {
			glog.Errorln("Failed to marshall info struct: ", err)
			return ctx.String(http.StatusInternalServerError, err.Error())
		}

		return ctx.JSONBlob(http.StatusOK, resp)
	}
}
