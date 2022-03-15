// Package api provides primitives to interact with the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen version v1.9.1 DO NOT EDIT.
package api

import (
	"bytes"
	"compress/gzip"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"path"
	"strings"

	"github.com/deepmap/oapi-codegen/pkg/runtime"
	"github.com/getkin/kin-openapi/openapi3"
	"github.com/labstack/echo/v4"
)

// ServerInterface represents all server handlers.
type ServerInterface interface {
	// Delete
	// (DELETE /userAccount/{uuid})
	DeleteAccount(ctx echo.Context, uuid string) error
	// Get an existing User Account by UUID
	// (GET /userAccount/{uuid})
	GetUserAccount(ctx echo.Context, uuid string) error
	// Modify an existing User Account
	// (PATCH /userAccount/{uuid})
	ModifyUserAccount(ctx echo.Context, uuid string) error
	// Verify / Unverify
	// (PUT /userAccount/{uuid}/verification)
	VerifyAccount(ctx echo.Context, uuid string) error
	// Gets all Users
	// (GET /userAccounts)
	GetUserAccounts(ctx echo.Context, params GetUserAccountsParams) error
	// Create a User Account
	// (POST /userAccounts)
	CreateAccount(ctx echo.Context) error
}

// ServerInterfaceWrapper converts echo contexts to parameters.
type ServerInterfaceWrapper struct {
	Handler ServerInterface
}

// DeleteAccount converts echo context to params.
func (w *ServerInterfaceWrapper) DeleteAccount(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid string

	err = runtime.BindStyledParameterWithLocation("simple", false, "uuid", runtime.ParamLocationPath, ctx.Param("uuid"), &uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.DeleteAccount(ctx, uuid)
	return err
}

// GetUserAccount converts echo context to params.
func (w *ServerInterfaceWrapper) GetUserAccount(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid string

	err = runtime.BindStyledParameterWithLocation("simple", false, "uuid", runtime.ParamLocationPath, ctx.Param("uuid"), &uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetUserAccount(ctx, uuid)
	return err
}

// ModifyUserAccount converts echo context to params.
func (w *ServerInterfaceWrapper) ModifyUserAccount(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid string

	err = runtime.BindStyledParameterWithLocation("simple", false, "uuid", runtime.ParamLocationPath, ctx.Param("uuid"), &uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.ModifyUserAccount(ctx, uuid)
	return err
}

// VerifyAccount converts echo context to params.
func (w *ServerInterfaceWrapper) VerifyAccount(ctx echo.Context) error {
	var err error
	// ------------- Path parameter "uuid" -------------
	var uuid string

	err = runtime.BindStyledParameterWithLocation("simple", false, "uuid", runtime.ParamLocationPath, ctx.Param("uuid"), &uuid)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter uuid: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.VerifyAccount(ctx, uuid)
	return err
}

// GetUserAccounts converts echo context to params.
func (w *ServerInterfaceWrapper) GetUserAccounts(ctx echo.Context) error {
	var err error

	// Parameter object where we will unmarshal all parameters from the context
	var params GetUserAccountsParams
	// ------------- Optional query parameter "limit" -------------

	err = runtime.BindQueryParameter("form", true, false, "limit", ctx.QueryParams(), &params.Limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, fmt.Sprintf("Invalid format for parameter limit: %s", err))
	}

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.GetUserAccounts(ctx, params)
	return err
}

// CreateAccount converts echo context to params.
func (w *ServerInterfaceWrapper) CreateAccount(ctx echo.Context) error {
	var err error

	// Invoke the callback with all the unmarshalled arguments
	err = w.Handler.CreateAccount(ctx)
	return err
}

// This is a simple interface which specifies echo.Route addition functions which
// are present on both echo.Echo and echo.Group, since we want to allow using
// either of them for path registration
type EchoRouter interface {
	CONNECT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	DELETE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	GET(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	HEAD(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	OPTIONS(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PATCH(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	POST(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	PUT(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
	TRACE(path string, h echo.HandlerFunc, m ...echo.MiddlewareFunc) *echo.Route
}

// RegisterHandlers adds each server route to the EchoRouter.
func RegisterHandlers(router EchoRouter, si ServerInterface) {
	RegisterHandlersWithBaseURL(router, si, "")
}

// Registers handlers, and prepends BaseURL to the paths, so that the paths
// can be served under a prefix.
func RegisterHandlersWithBaseURL(router EchoRouter, si ServerInterface, baseURL string) {

	wrapper := ServerInterfaceWrapper{
		Handler: si,
	}

	router.DELETE(baseURL+"/userAccount/:uuid", wrapper.DeleteAccount)
	router.GET(baseURL+"/userAccount/:uuid", wrapper.GetUserAccount)
	router.PATCH(baseURL+"/userAccount/:uuid", wrapper.ModifyUserAccount)
	router.PUT(baseURL+"/userAccount/:uuid/verification", wrapper.VerifyAccount)
	router.GET(baseURL+"/userAccounts", wrapper.GetUserAccounts)
	router.POST(baseURL+"/userAccounts", wrapper.CreateAccount)

}

// Base64 encoded, gzipped, json marshaled Swagger object
var swaggerSpec = []string{

	"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",
	"P+1hQTggAAA=",
}

// GetSwagger returns the content of the embedded swagger specification file
// or error if failed to decode
func decodeSpec() ([]byte, error) {
	zipped, err := base64.StdEncoding.DecodeString(strings.Join(swaggerSpec, ""))
	if err != nil {
		return nil, fmt.Errorf("error base64 decoding spec: %s", err)
	}
	zr, err := gzip.NewReader(bytes.NewReader(zipped))
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, fmt.Errorf("error decompressing spec: %s", err)
	}

	return buf.Bytes(), nil
}

var rawSpec = decodeSpecCached()

// a naive cached of a decoded swagger spec
func decodeSpecCached() func() ([]byte, error) {
	data, err := decodeSpec()
	return func() ([]byte, error) {
		return data, err
	}
}

// Constructs a synthetic filesystem for resolving external references when loading openapi specifications.
func PathToRawSpec(pathToFile string) map[string]func() ([]byte, error) {
	var res = make(map[string]func() ([]byte, error))
	if len(pathToFile) > 0 {
		res[pathToFile] = rawSpec
	}

	return res
}

// GetSwagger returns the Swagger specification corresponding to the generated code
// in this file. The external references of Swagger specification are resolved.
// The logic of resolving external references is tightly connected to "import-mapping" feature.
// Externally referenced files must be embedded in the corresponding golang packages.
// Urls can be supported but this task was out of the scope.
func GetSwagger() (swagger *openapi3.T, err error) {
	var resolvePath = PathToRawSpec("")

	loader := openapi3.NewLoader()
	loader.IsExternalRefsAllowed = true
	loader.ReadFromURIFunc = func(loader *openapi3.Loader, url *url.URL) ([]byte, error) {
		var pathToFile = url.String()
		pathToFile = path.Clean(pathToFile)
		getSpec, ok := resolvePath[pathToFile]
		if !ok {
			err1 := fmt.Errorf("path not found: %s", pathToFile)
			return nil, err1
		}
		return getSpec()
	}
	var specData []byte
	specData, err = rawSpec()
	if err != nil {
		return
	}
	swagger, err = loader.LoadFromData(specData)
	if err != nil {
		return
	}
	return
}
