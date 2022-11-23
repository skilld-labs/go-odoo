package odoo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/cookiejar"
	"strings"

	"github.com/imdario/mergo"
)

// List is a shortcut to []interface{}
type List []interface{}

// Client holds client connection and the server metadata
type Client struct {
	baseURL    string
	httpClient *http.Client
	Auth       *AuthResponse
	Context    map[string]interface{}
}

// IsValid checks if a Client is in a valid state
func (c *Client) IsValid() bool {
	return c.httpClient != nil
}

// GetBaseURL gets client's connection base URL
func (c *Client) GetBaseURL() string {
	if c.Auth == nil || c.Auth.Result == nil {
		return ""
	}
	return string(c.Auth.Result.WebBaseURL)
}

// GetServerVersion gets Odoo server's version
func (c *Client) GetServerVersion() string {
	if c.Auth == nil || c.Auth.Result == nil {
		return ""
	}
	return string(c.Auth.Result.ServerVersion)
}

// Call calls remote Odoo method
func (c *Client) Call(model, method string, args []interface{}, kwargs map[string]interface{}) (interface{}, error) {
	url := fmt.Sprintf("%s/web/dataset/call_kw/%s/%s", c.baseURL, model, method)

	context := map[string]interface{}{}
	err := mergo.Merge(&context, c.Auth.Result.UserContext)
	if err != nil {
		return nil, err
	}

	if ctxParam, ok := kwargs["context"]; ok {
		if err = mergo.Merge(&context, ctxParam); err != nil {
			return nil, err
		}
	}
	kwargs["context"] = context

	requestPayload := map[string]interface{}{
		"id":      1,
		"jsonrpc": "2.0",
		"method":  "call",
		"params": map[string]interface{}{
			"model":  model,
			"method": method,
			"args":   args,
			"kwargs": kwargs,
		},
	}

	payloadBuf, err := json.Marshal(requestPayload)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", url, bytes.NewReader(payloadBuf))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")

	httpResponse, err := c.httpClient.Do(request)
	if err != nil {
		return nil, err
	}
	defer httpResponse.Body.Close()

	response := &RPCResponse{}
	err = json.NewDecoder(httpResponse.Body).Decode(response)
	if err != nil {
		return nil, err
	}

	if response.Error != nil {
		return nil, fmt.Errorf("RPC Error: %s", response.Error.Data.Message)
	}

	if response.Result == nil {
		return nil, fmt.Errorf("RPC returned no result")
	}

	return response.Result, nil
}

// Read reads records for the specified model with specified IDs
// The read will include only specified fields, or if empty will return
// all fields. Beware that retrieving all fields will normally
// take longer time than selective fields, depending on the model's
// schema complexity
func (c *Client) Read(model string, ids []int, fields []string) ([]interface{}, error) {
	args := []interface{}{ids}
	kwargs := map[string]interface{}{}
	if len(fields) > 0 {
		kwargs["fields"] = fields
	}
	respIface, err := c.Call(model, "read", args, kwargs)
	if err != nil {
		return nil, err
	}

	return respIface.([]interface{}), nil
}

// Search searches records in the specified model, matching the
// criteria specified in domain, which is similar to domain in Odoo
// python syntax
func (c *Client) Search(model string, domain []interface{}) ([]int, error) {
	args := []interface{}{domain}
	kwargs := map[string]interface{}{}
	respIface, err := c.Call(model, "search", args, kwargs)
	if err != nil {
		return nil, err
	}
	recs := respIface.([]interface{})
	var ids []int
	for _, rec := range recs {
		idf := rec.(float64)
		ids = append(ids, int(idf))
	}
	return ids, nil
}

// SearchRead is a convenience method that search and read in one go
func (c *Client) SearchRead(model string, domain []interface{}, fields []string) ([]interface{}, error) {
	args := []interface{}{domain}
	kwargs := map[string]interface{}{}
	if len(fields) > 0 {
		kwargs["fields"] = fields
	}
	respIface, err := c.Call(model, "search_read", args, kwargs)
	if err != nil {
		return nil, err
	}

	return respIface.([]interface{}), nil
}

// Write updates records
func (c *Client) Write(model string, ids []int, values map[string]interface{}) (bool, error) {
	args := []interface{}{ids, values}
	kwargs := map[string]interface{}{}
	respIface, err := c.Call(model, "write", args, kwargs)
	if err != nil {
		return false, err
	}

	return respIface.(bool), nil
}

// Create creates a new record and returns the new record ID or error
func (c *Client) Create(model string, values map[string]interface{}) (int, error) {
	args := []interface{}{values}
	kwargs := map[string]interface{}{}
	respIface, err := c.Call(model, "create", args, kwargs)
	if err != nil {
		return -1, err
	}
	return int(respIface.(float64)), nil
}

// Delete deletes records
func (c *Client) Delete(model string, ids []int) (bool, error) {
	args := []interface{}{ids}
	kwargs := map[string]interface{}{}
	respIface, err := c.Call(model, "unlink", args, kwargs)
	if err != nil {
		return false, err
	}
	return respIface.(bool), nil
}

// WithContext returns a new Client with added context
func (c *Client) WithContext(ctx map[string]interface{}) *Client {
	newContext := map[string]interface{}{}
	for k, v := range c.Context {
		newContext[k] = v
	}
	for k, v := range ctx {
		newContext[k] = v
	}
	newClient := &Client{
		baseURL:    c.baseURL,
		Auth:       c.Auth,
		httpClient: c.httpClient,
		Context:    newContext,
	}
	return newClient
}

// Connect attempts to connect to Odoo server and returns the
// client or error
func Connect(baseURL, db, login, password string) (*Client, error) {
	url := baseURL + "/web/session/authenticate"

	jar, err := cookiejar.New(nil)
	if err != nil {
		return nil, err
	}
	client := &http.Client{
		Jar: jar,
	}

	authParams := AuthParams{
		DB:       db,
		Login:    login,
		Password: password,
	}

	authRequest := AuthRequest{
		ID:      "1",
		JSONRPC: "2.0",
		Method:  "call",
		Params:  authParams,
	}

	payload := &strings.Builder{}
	err = json.NewEncoder(payload).Encode(authRequest)
	if err != nil {
		return nil, err
	}

	request, err := http.NewRequest("POST", url, strings.NewReader(payload.String()))
	if err != nil {
		return nil, err
	}
	request.Header.Add("Content-Type", "application/json")
	request.Header.Add("Accept", "application/json")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	data := AuthResponse{}
	err = json.NewDecoder(response.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	if data.Error != nil {
		return nil, fmt.Errorf("Auth error: %s", data.Error.Data.Message)
	}

	if data.Result == nil {
		return nil, fmt.Errorf("Auth returned no result")
	}

	odooClient := &Client{
		baseURL:    baseURL,
		Auth:       &data,
		httpClient: client,
		Context:    make(map[string]interface{}),
	}

	return odooClient, nil
}

// AuthParams struct holds authentication parameters required for an auth request
type AuthParams struct {
	DB       string `json:"db"`
	Login    string `json:"login"`
	Password string `json:"password"`
}

// AuthRequest struct used to hold data for an authentication request
type AuthRequest struct {
	ID      string     `json:"id"`
	JSONRPC string     `json:"jsonrpc"`
	Method  string     `json:"method"`
	Params  AuthParams `json:"params"`
}

// AuthResponse struct holds response data returned from Odoo server
type AuthResponse struct {
	ID      string      `json:"id"`
	JSONRPC string      `json:"jsonrpc"`
	Result  *AuthResult `json:"result"`
	Error   *AuthError  `json:"error"`
}

// AuthError struct holds data for an authentication error response
type AuthError struct {
	Code    int            `json:"code,omitempty"`
	Message string         `json:"message,omitempty"`
	Data    *AuthErrorData `json:"data,omitempty"`
}

// AuthErrorData struct holds data from an authentication error response
type AuthErrorData struct {
	Name          string        `json:"name,omitempty"`
	Debug         string        `json:"debug,omitempty"`
	Message       string        `json:"message,omitempty"`
	Arguments     []interface{} `json:"arguments,omitempty"`
	ExceptionType string        `json:"exception_type,omitempty"`
}

// AuthResult struct hold authentication data returned after successfull login
type AuthResult struct {
	UID                int                    `json:"uid"`
	IsSystem           bool                   `json:"is_system"`
	IsAdmin            bool                   `json:"is_admin"`
	UserContext        map[string]interface{} `json:"user_context"`
	DB                 string                 `json:"db"`
	ServerVersion      string                 `json:"server_version"`
	ServerVersionInfo  []interface{}          `json:"server_version_info"`
	Name               string                 `json:"name"`
	UserName           string                 `json:"username"`
	PartnerDisplayName string                 `json:"partner_display_name"`
	CompanyID          int                    `json:"company_id"`
	PartnerID          int                    `json:"partner_id"`
	WebBaseURL         string                 `json:"web.base.url"`
}

// RPCErrorData struct holds error data from a request error
type RPCErrorData struct {
	Name          string        `json:"name"`
	Debug         string        `json:"debug"`
	Message       string        `json:"message"`
	Arguments     []interface{} `json:"arguments"`
	ExceptionType string        `json:"exception_type"`
}

// RPCError struct holds data related to RPC error
type RPCError struct {
	Code    int          `json:"code"`
	Message string       `json:"message"`
	Data    RPCErrorData `json:"data"`
}

// RPCResponse struct holds data returned from server on a successfull request
type RPCResponse struct {
	ID      int         `json:"id"`
	JSONRPC string      `json:"jsonrpc"`
	Result  interface{} `json:"result"`
	Error   *RPCError   `json:"error"`
}
