# FORM3 Account API SDK Client

## Prerequisites
Using SKD in your project requires `go 1.12`
[Get GoLang](https://golang.org/dl/)

To run tests `docker` is required.
[Get started with docker](https://www.docker.com/get-started)

## Instructions
### Create API client instance
To create an instance of Account API client use: `NewHTTPClient(httpClient *http.Client, apiHost, apiEndpoint string)`

* httpClient - http client can be provided with custom configuration. 
If `nil` is provided - Default Http client with 30 seconds timeout will be used.
* apiHost - string value of valid url which points to server hosting account API.
* apiEndpoint - string value of valid url path which represents accounts API endpoint on the server.

Returns pointer to `account.HTTPClient` and `error` if provided URLs fail to parse.

### Create account resource builder
Builder is created based on what country accounts it will create.
Supported countries are listed under `Country` enum. 
One of available entries of the enum must be provided as parameter in Builder constructor: `NewBuilder(Country)`

Returns pointer to `account.Builder` with pre-set Country code and BankID code.

### How to use account builder
Builder provides functions to set account attributes and then validate them.

For making account creation seamless essential attributes are separated from optional.
Every setter step returns same Builder which can be used for setting next attribute or validating.
If all required attributes are set and in right format - `account.Account` will be returned by `Validate()` method.

Example:
```
gbAccount, err := accountBuilder.
    SetID("0911be7a-f7da-4f7e-b692-d1fbdf1aa7cf").
    SetOrganizationID("cac625ac-9aa6-4557-a495-2d8ea7882c4f").
    SetBankID("601613").
    SetBic("ABBYGB2LGTB").
    SetOptionalAttribute().SetBaseCurrency(currency.EUR).
    Validate()
```
Attributes and their format differ by country.
For more detailed information visit [API docs.](https://api-docs.form3.tech/api.html#organisation-accounts-create)

### Available Account API client methods

#### Create
Register an existing bank account with Form3 or create a new one.

Method contract -`Create(ctx context.Context, account *Account) (*Account, error)`

* ctx - provide a context for request customization
* account - request payload created with account builder

Returns created account or error if request was unsuccessful.

#### Fetch
Get a single account using the account ID.

Method contact - `Fetch(ctx context.Context, accountID string) (*Account, error)`

* ctx - provide a context for request customization
* accountID - valid uuid linked to desired account

Returns requested account or error if request was unsuccessful.

#### List
List accounts with the ability to page.

Method contact - `List(ctx context.Context, paging *PaginationSettings) ([]Account, error)`

* ctx - provide a context for request customization
* paging - PaginationSettings object which describes size of the page and which page is required.
If set to nil will request for all available accounts.

Returns slice of requested accounts or error if request was unsuccessful.

#### Delete
Delete a single account using the account ID.

Method contact - `Delete(ctx context.Context, accountID string, version int) error`

* ctx - provide a context for request customization
* accountID - valid uuid linked to desired account
* version - version number of record

Returns error if request was unsuccessful.

### Modify fetched account
Account object provides just getter methods. 
If there is a need to modify fetched account - use account builder constructor `CastBuilderFrom(*Account)`

Account builder with values from provided account will be returned.
That can be used to set desired attributes and validate them in order to receive a transformed account object.
 