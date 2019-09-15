package test

import (
	"errors"
	"net/http"

	"github.com/DATA-DOG/godog"
	"github.com/google/uuid"

	account "github.com/r0kas/form3-accountapi-client"
)

var apiHost string
var apiEndpoint string
var apiClient *account.HTTPClient
var accountBuilder *account.Builder
var accountCountry account.Country
var theAccount *account.Account
var accountsList []account.Account

func myCountryCodeIs(c string) error {
	accountCountry = account.Country(c)
	return nil
}

func iCreateAnAccountBuilder() error {
	accountBuilder = account.NewBuilder(accountCountry)
	return nil
}

func setRandomAccountID() error {
	accountBuilder.SetID(uuid.New().String())
	return nil
}

func setRandomOrganizationID() error {
	accountBuilder.SetOrganizationID(uuid.New().String())
	return nil
}

func setBankIDTo(bankID string) error {
	accountBuilder.SetBankID(bankID)
	return nil
}

func setBicTo(bic string) error {
	accountBuilder.SetBic(bic)
	return nil
}

func iHaveAValidAccount() (err error) {
	theAccount, err = accountBuilder.Validate()
	return err
}

func iHaveAnInvalidAccount() error {
	_, expectingAnError := accountBuilder.Validate()
	if expectingAnError != nil {
		return nil
	}
	return errors.New("account is valid")
}

func accountBankIDCodeIs(bankIDCode string) error {
	if bankIDCode != theAccount.BankIDCode() {
		return errors.New("not valid bank ID code")
	}
	return nil
}

func setAccountNumberTo(number string) error {
	accountBuilder.SetOptionalAttribute().SetAccountNumber(number)
	return nil
}

func setFirstNameTo(name string) error {
	accountBuilder.SetOptionalAttribute().SetFirstName(name)
	return nil
}

func setAlternativeBankAccountName(accName string) error {
	accountBuilder.SetOptionalAttribute().SetAltBankAccountNames(accName)
	return nil
}

func setBusinessClassificationTo(businessType string) error {
	accountBuilder.SetOptionalAttribute().SetAccountClassification(businessType)
	return nil
}

func setSecondaryIdentificationTo(secondaryID string) error {
	accountBuilder.SetOptionalAttribute().SetSecondaryIdentification(secondaryID)
	return nil
}

func accountNumberIs(accNumber string) error {
	if theAccount.AccountNumber() != accNumber {
		return errors.New("wrong account number")
	}
	return nil
}

func accountFirstNameIs(name string) error {
	if theAccount.FirstName() != name {
		return errors.New("wrong first name")
	}
	return nil
}

func setIbanTo(iban string) error {
	accountBuilder.SetIban(iban)
	return nil
}

func accountBusinessClassificationIs(businessClass string) error {
	if theAccount.AccountClassification() != businessClass {
		return errors.New("wrong classification")
	}
	return nil
}

func accountSecondaryIdentificationIs(secondaryID string) error {
	if theAccount.SecondaryIdentification() != secondaryID {
		return errors.New("wrong secondary identification")
	}
	return nil
}

func accountCountryCodeIs(countryCode string) error {
	if theAccount.Country() != countryCode {
		return errors.New("wrong country code")
	}
	return nil
}

func accountBankIdIs(bankID string) error {
	if theAccount.BankID() != bankID {
		return errors.New("wrong bank ID")
	}
	return nil
}

func accountBicIs(bic string) error {
	if theAccount.Bic() != bic {
		return errors.New("wrong bic")
	}
	return nil
}

func apiClientIsCreatedWithApiHostAndApiEndpoint(hostname, endpoint string) (err error) {
	apiClient, err = account.NewHTTPClient(nil, hostname, endpoint)
	return
}

func iRunApiClientCreateCommand() (err error) {
	theAccount, err = apiClient.Create(nil, theAccount)
	return
}

func apiHostIs(hostname string) error {
	apiHost = hostname
	return nil
}

func apiEndpointIs(endpoint string) error {
	apiEndpoint = endpoint
	return nil
}

func apiClientIsCreated() (err error) {
	apiClient, err = account.NewHTTPClient(http.DefaultClient, apiHost, apiEndpoint)
	return nil
}

func iRunApiClientFetchCommandForSameID() (err error) {
	theAccount, err = apiClient.Fetch(nil, theAccount.ID())
	return
}

func iRunApiClientDeleteCommandForSameID() (err error) {
	err = apiClient.Delete(nil, theAccount.ID(), theAccount.Version())
	return
}

func apiClientCommandFetchFailsForSameID() error {
	_, err := apiClient.Fetch(nil, theAccount.ID())
	if err != nil {
		return nil
	}
	return errors.New("account was found. Which is not expected")
}

func iListAvailableAccounts() (err error) {
	accountsList, err = apiClient.List(nil, nil)
	return
}

func deleteAllListedAccounts() error {
	for _, acc := range accountsList {
		err := apiClient.Delete(nil, acc.ID(), acc.Version())
		if err != nil {
			return err
		}
	}
	return nil
}

func iCreateRandomAccounts(accountsCount int) error {
	myCountryCodeIs("BE")
	iCreateAnAccountBuilder()
	setBankIDTo("123")
	for i := 0; i < accountsCount; i++ {
		setRandomAccountID()
		setRandomOrganizationID()
		err := iHaveAValidAccount()
		if err != nil {
			return err
		}
		err = iRunApiClientCreateCommand()
		if err != nil {
			return err
		}
	}
	return nil
}

func iHaveAccountsInMyList(accountsCount int) error {
	if len(accountsList) != accountsCount {
		return errors.New("expected accounts count does not match")
	}
	return nil
}

func iListPageWithPageSize(pageNumber string, pageSize int) (err error) {
	pagination := &account.PaginationSettings{
		Enabled:    true,
		PageNumber: pageNumber,
		PageSize:   pageSize,
	}
	accountsList, err = apiClient.List(nil, pagination)
	return
}

func aPIReturnsAnErrorOnCreateCommand() error {
	_, err := apiClient.Create(nil, theAccount)
	if err == nil {
		return errors.New("expected and error. Got nil")
	}
	return nil
}

func apiClientIsHealthy() error {
	if apiClient.IsHealthy(nil) {
		return nil
	}
	return errors.New("API is not healthy")
}

func FeatureContext(s *godog.Suite) {
	s.Step(`^my country code is "([^"]*)"\$$`, myCountryCodeIs)
	s.Step(`^I create an account builder$`, iCreateAnAccountBuilder)
	s.Step(`^set random account ID$`, setRandomAccountID)
	s.Step(`^set random organization ID$`, setRandomOrganizationID)
	s.Step(`^set bank ID to "([^"]*)"\$$`, setBankIDTo)
	s.Step(`^set bic to "([^"]*)"\$$`, setBicTo)
	s.Step(`^I have a valid account$`, iHaveAValidAccount)
	s.Step(`^I have an invalid account$`, iHaveAnInvalidAccount)
	s.Step(`^account bank ID code is "([^"]*)"\$$`, accountBankIDCodeIs)
	s.Step(`^set account number to "([^"]*)"\$$`, setAccountNumberTo)
	s.Step(`^set first name to "([^"]*)"\$$`, setFirstNameTo)
	s.Step(`^set alternative bank account name "([^"]*)"\$$`, setAlternativeBankAccountName)
	s.Step(`^set business classification to "([^"]*)"\$$`, setBusinessClassificationTo)
	s.Step(`^set secondary identification to "([^"]*)"\$$`, setSecondaryIdentificationTo)
	s.Step(`^account number is "([^"]*)"\$$`, accountNumberIs)
	s.Step(`^account first name is "([^"]*)"\$$`, accountFirstNameIs)
	s.Step(`^account business classification is "([^"]*)"\$$`, accountBusinessClassificationIs)
	s.Step(`^account secondary identification is "([^"]*)"\$$`, accountSecondaryIdentificationIs)
	s.Step(`^set iban to "([^"]*)"\$$`, setIbanTo)
	s.Step(`^account country code is "([^"]*)"\$$`, accountCountryCodeIs)
	s.Step(`^account bank id is "([^"]*)"\$$`, accountBankIdIs)
	s.Step(`^account bic is "([^"]*)"\$$`, accountBicIs)
	s.Step(`^I run api client Create command$`, iRunApiClientCreateCommand)
	s.Step(`^api client is created with api host "([^"]*)" and api endpoint "([^"]*)"$`, apiClientIsCreatedWithApiHostAndApiEndpoint)
	s.Step(`^api host is "([^"]*)"$`, apiHostIs)
	s.Step(`^api endpoint is "([^"]*)"$`, apiEndpointIs)
	s.Step(`^api client is created$`, apiClientIsCreated)
	s.Step(`^I run api client Fetch command for same ID$`, iRunApiClientFetchCommandForSameID)
	s.Step(`^I run api client Delete command for same ID$`, iRunApiClientDeleteCommandForSameID)
	s.Step(`^api client command Fetch fails for same ID$`, apiClientCommandFetchFailsForSameID)
	s.Step(`^I List available accounts$`, iListAvailableAccounts)
	s.Step(`^Delete all Listed accounts$`, deleteAllListedAccounts)
	s.Step(`^I Create (\d+) random accounts$`, iCreateRandomAccounts)
	s.Step(`^I have (\d+) account\/s in my list$`, iHaveAccountsInMyList)
	s.Step(`^I List "([^"]*)" page with Page Size (\d+)$`, iListPageWithPageSize)
	s.Step(`^API returns an error on Create command$`, aPIReturnsAnErrorOnCreateCommand)
	s.Step(`^api client is healthy$`, apiClientIsHealthy)
}
