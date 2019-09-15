package account

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"

	"github.com/google/uuid"
	"github.com/pkg/errors"
)

type (
	// HTTPClient handles execution and transport of API commands.
	HTTPClient struct {
		httpClient  *http.Client
		apiHost     *url.URL
		apiEndpoint *url.URL
	}

	// PaginationSettings represents settings for pagination feature on List command.
	PaginationSettings struct {
		// Enabled turns on pagination functionality in API Client
		Enabled bool
		// PageNumber sets which page to return. Supports int numbers or 'first', 'last' tags.
		PageNumber string
		// PageSize sets maximum size of the returned data set
		PageSize int
	}
)

// NewHTTPClient creates HTTP Client instance.
///////
// Enables SDK user to pre configure http.Client e.g. Timeout settings
// Validates provided host and endpoint parameters
///////
func NewHTTPClient(httpClient *http.Client, apiHost, apiEndpoint string) (*HTTPClient, error) {
	host, err := url.Parse(apiHost)
	if err != nil {
		return nil, err
	}
	endpoint, err := url.Parse(apiEndpoint)
	if err != nil {
		return nil, err
	}
	if httpClient == nil {
		httpClient = &http.Client{
			Timeout: 30,
		}
	}
	return &HTTPClient{
		httpClient:  httpClient,
		apiHost:     host,
		apiEndpoint: endpoint,
	}, nil
}

// Create handles execution of create command against accounts API.
// Accepts context and account objects.
// Returns account object which will be received from API server after successful operation.
///////
// Context adds additional request configuration flexibility for SDK user.
//////
func (c *HTTPClient) Create(ctx context.Context, account *Account) (*Account, error) {
	if err := c.validateClient(); err != nil {
		return nil, err
	}
	if account == nil {
		return nil, errors.New("cannot create account without initialized account object. Use account builder")
	}

	request, err := c.newRequest(ctx, http.MethodPost, c.clientAPIRequestURL("", nil), c.createTransportData(account))
	if err != nil {
		return nil, err
	}

	responseJSON := new(restTransport)
	err = c.doRequest(request, http.StatusCreated, responseJSON)
	if err != nil {
		return nil, err
	}

	return accountFrom(responseJSON.Data), nil
}

// Fetch handles execution of fetch command against accounts API.
// Accepts context and account ID - UUID format 4.
// Returns account object which will be received if account is found on API server.
///////
// Context adds additional request configuration flexibility for SDK user.
//////
func (c *HTTPClient) Fetch(ctx context.Context, accountID string) (*Account, error) {
	if err := c.validateClient(); err != nil {
		return nil, err
	}
	if _, err := uuid.Parse(accountID); err != nil {
		return nil, errors.Wrap(err, "provided account ID must be a valid UUID")
	}

	request, err := c.newRequest(ctx, http.MethodGet, c.clientAPIRequestURL(accountID, nil), nil)
	if err != nil {
		return nil, err
	}

	responseJSON := new(restTransport)
	err = c.doRequest(request, http.StatusOK, responseJSON)
	if err != nil {
		return nil, err
	}

	return accountFrom(responseJSON.Data), nil
}

// List handles execution of list command against accounts API.
// Accepts context and pagination settings. If pagination settings are not required - parameter can be set to nil.
// Returns slice of requested account objects.
///////
// Context adds additional request configuration flexibility for SDK user.
//////
func (c *HTTPClient) List(ctx context.Context, paging *PaginationSettings) ([]Account, error) {
	if err := c.validateClient(); err != nil {
		return nil, err
	}
	if paging == nil {
		paging = &PaginationSettings{}
	}

	request, err := c.newRequest(ctx, http.MethodGet, c.clientAPIRequestURL("", c.pagingParameters(paging)), nil)
	if err != nil {
		return nil, err
	}

	responseJSON := new(restTransportList)
	err = c.doRequest(request, http.StatusOK, responseJSON)
	if err != nil {
		return nil, err
	}
	accounts := make([]Account, 0)
	for _, accountJSON := range responseJSON.Data {
		accounts = append(accounts, *accountFrom(accountJSON))
	}
	return accounts, nil
}

// Delete handles execution of delete command against accounts API.
// Accepts context and account ID - UUID format 4.
// Returns no error if execution is successful.
///////
// Context adds additional request configuration flexibility for SDK user.
//////
func (c *HTTPClient) Delete(ctx context.Context, accountID string, version int) error {
	if err := c.validateClient(); err != nil {
		return err
	}
	if _, err := uuid.Parse(accountID); err != nil {
		return errors.Wrap(err, "provided account ID must be a valid UUID")
	}

	request, err := c.newRequest(ctx, http.MethodDelete, c.clientAPIRequestURL(accountID, c.versionParameters(version)), nil)
	if err != nil {
		return err
	}

	return c.doRequest(request, http.StatusNoContent, nil)
}

// Health queries API for it's status
// Returns a boolean based on API health status
func (c *HTTPClient) Health(ctx context.Context) bool {
	if err := c.validateClient(); err != nil {
		return false
	}
	request, err := c.newRequest(ctx, http.MethodGet, c.healthEndpoint(), nil)
	if err != nil {
		return false
	}

	err = c.doRequest(request, http.StatusOK, nil)
	if err != nil {
		return false
	}
	return true
}

// checks if API client was properly initialized, so that API commands would fail fast if something is missing.
func (c *HTTPClient) validateClient() error {
	if c.httpClient == nil || c.apiHost == nil || c.apiEndpoint == nil {
		return errors.New("accounts api client was not initialised. Use constructor method")
	}
	return nil
}

// used for constructing any new request
func (c *HTTPClient) newRequest(ctx context.Context, method string, url *url.URL, body interface{}) (*http.Request, error) {
	buf, err := c.bodyBuffer(body)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(method, url.String(), buf)
	if err != nil {
		return nil, err
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}
	if body != nil {
		req.Header.Set("Content-Type", "application/vnd.api+json")
	}
	req.Header.Set("Accept", "application/vnd.api+json")
	return req, nil
}

func (c *HTTPClient) bodyBuffer(body interface{}) (*bytes.Buffer, error) {
	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}
	return buf, nil
}

func (c *HTTPClient) clientAPIRequestURL(addToPath string, parameters map[string]string) *url.URL {
	endpoint := *c.apiEndpoint
	if addToPath != "" {
		endpoint.Path = path.Join(endpoint.Path, addToPath)
	}
	c.addParameters(&endpoint, parameters)
	return c.apiHost.ResolveReference(&endpoint)
}

func (c *HTTPClient) healthEndpoint() *url.URL {
	healthEndpoint, _ := url.Parse("/v1/health")
	return c.apiHost.ResolveReference(healthEndpoint)
}

func (c *HTTPClient) addParameters(url *url.URL, parameters map[string]string) {
	if parameters != nil && len(parameters) > 0 {
		queryString := url.Query()
		for key, value := range parameters {
			queryString.Set(key, value)
		}
		url.RawQuery = queryString.Encode()
	}
}

// creates body structure for API requests
func (c *HTTPClient) createTransportData(account *Account) *restTransport {
	return &restTransport{
		Data: transportData{
			Type:           "accounts",
			ID:             account.id,
			OrganizationID: account.organizationID,
			Version:        account.versionIndex,
			Attributes:     *account.attributes(),
		},
	}
}

// handles SDK's http transport.
// accepts expected http response code, so it can be used through different commands to choose if
// received response is a happy day scenario - if not returns error together with status code and API response message.
func (c *HTTPClient) doRequest(request *http.Request, expectedResponseCode int, responseData interface{}) error {
	response, err := c.httpClient.Do(request)
	if err != nil {
		return err
	}
	defer response.Body.Close()

	if response.StatusCode != expectedResponseCode {
		return errors.Wrap(c.errorFromStream(response.Body), response.Status)
	}
	if responseData != nil {
		return json.NewDecoder(response.Body).Decode(responseData)
	}
	return nil
}

func (c *HTTPClient) errorFromStream(stream io.ReadCloser) error {
	bodyBytes, err := ioutil.ReadAll(stream)
	if err != nil {
		return errors.Wrap(err, "failed to read response stream.")
	}
	return errors.New(string(bodyBytes))
}

func (c *HTTPClient) pagingParameters(paging *PaginationSettings) map[string]string {
	parameters := make(map[string]string)
	if paging.Enabled {
		parameters["page[number]"] = paging.PageNumber
		parameters["page[size]"] = strconv.Itoa(paging.PageSize)
	}
	return parameters
}

func (c *HTTPClient) versionParameters(version int) map[string]string {
	parameters := make(map[string]string)
	parameters["version"] = strconv.Itoa(version)
	return parameters
}
