package test

import (
	"errors"

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

func countryCodeIsEqual(c string) error {
	accountCountry = account.Country(c)
	return nil
}
func createAccountBuilder() error {
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

func setBankID(bankID string) error {
	accountBuilder.SetBankID(bankID)
	return nil
}

func setBic(bic string) error {
	accountBuilder.SetBic(bic)
	return nil
}

func isValidAccount() (err error) {
	theAccount, err = accountBuilder.Validate()
	return err
}

func isInvalidAccount() error {
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

func setAccountNumber(number string) error {
	accountBuilder.SetOptionalAttribute().SetAccountNumber(number)
	return nil
}

func setFirstName(name string) error {
	accountBuilder.SetOptionalAttribute().SetFirstName(name)
	return nil
}

func setAlternativeBankAccountName(accName string) error {
	accountBuilder.SetOptionalAttribute().SetAltBankAccountNames(accName)
	return nil
}

func setBusinessClassification(businessType string) error {
	accountBuilder.SetOptionalAttribute().SetAccountClassification(businessType)
	return nil
}

func setSecondaryIdentification(secondaryID string) error {
	accountBuilder.SetOptionalAttribute().SetSecondaryIdentification(secondaryID)
	return nil
}

func accountNumberEquals(accNumber string) error {
	if theAccount.AccountNumber() != accNumber {
		return errors.New("wrong account number")
	}
	return nil
}

func accountFirstNameEquals(name string) error {
	if theAccount.FirstName() != name {
		return errors.New("wrong first name")
	}
	return nil
}

func setIban(iban string) error {
	accountBuilder.SetIban(iban)
	return nil
}

func accountBusinessClassificationEquals(businessClass string) error {
	if theAccount.AccountClassification() != businessClass {
		return errors.New("wrong classification")
	}
	return nil
}

func accountSecondaryIdentificationEquals(secondaryID string) error {
	if theAccount.SecondaryIdentification() != secondaryID {
		return errors.New("wrong secondary identification")
	}
	return nil
}

func accountCountryCodeEquals(countryCode string) error {
	if theAccount.Country() != countryCode {
		return errors.New("wrong country code")
	}
	return nil
}

func accountBankIdEquals(bankID string) error {
	if theAccount.BankID() != bankID {
		return errors.New("wrong bank ID")
	}
	return nil
}

func accountBicEquals(bic string) error {
	if theAccount.Bic() != bic {
		return errors.New("wrong bic")
	}
	return nil
}

func createAPIClient() (err error) {
	apiClient, err = account.NewHTTPClient(nil, apiHost, apiEndpoint)
	return
}

func createAccount() (err error) {
	theAccount, err = apiClient.Create(nil, theAccount)
	return
}

func apiHostEquals(hostname string) error {
	apiHost = hostname
	return nil
}

func apiEndpointEquals(endpoint string) error {
	apiEndpoint = endpoint
	return nil
}

func fetchAccount() (err error) {
	theAccount, err = apiClient.Fetch(nil, theAccount.ID())
	return
}

func deleteAccount() (err error) {
	err = apiClient.Delete(nil, theAccount.ID(), theAccount.Version())
	return
}

func fetchAccountFails() error {
	_, err := apiClient.Fetch(nil, theAccount.ID())
	if err != nil {
		return nil
	}
	return errors.New("account was found. Which is not expected")
}

func listAccounts() (err error) {
	accountsList, err = apiClient.List(nil, nil)
	return
}

func deleteListedAccounts() error {
	for _, acc := range accountsList {
		err := apiClient.Delete(nil, acc.ID(), acc.Version())
		if err != nil {
			return err
		}
	}
	return nil
}

func createRandomAccounts(accountsCount int) error {
	countryCodeIsEqual("BE")
	createAccountBuilder()
	setBankID("123")
	for i := 0; i < accountsCount; i++ {
		setRandomAccountID()
		setRandomOrganizationID()
		err := isValidAccount()
		if err != nil {
			return err
		}
		err = createAccount()
		if err != nil {
			return err
		}
	}
	return nil
}

func accountsInList(accountsCount int) error {
	if len(accountsList) != accountsCount {
		return errors.New("expected accounts count does not match")
	}
	return nil
}

func listAccountsWithPageSize(pageNumber string, pageSize int) (err error) {
	pagination := &account.PaginationSettings{
		Enabled:    true,
		PageNumber: pageNumber,
		PageSize:   pageSize,
	}
	accountsList, err = apiClient.List(nil, pagination)
	return
}

func createAccountFails() error {
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
	s.Step(`^my country code is "([^"]*)"\$$`, countryCodeIsEqual)
	s.Step(`^I create an account builder$`, createAccountBuilder)
	s.Step(`^set random account ID$`, setRandomAccountID)
	s.Step(`^set random organization ID$`, setRandomOrganizationID)
	s.Step(`^set bank ID to "([^"]*)"\$$`, setBankID)
	s.Step(`^set bic to "([^"]*)"\$$`, setBic)
	s.Step(`^I have a valid account$`, isValidAccount)
	s.Step(`^I have an invalid account$`, isInvalidAccount)
	s.Step(`^account bank ID code is "([^"]*)"\$$`, accountBankIDCodeIs)
	s.Step(`^set account number to "([^"]*)"\$$`, setAccountNumber)
	s.Step(`^set first name to "([^"]*)"\$$`, setFirstName)
	s.Step(`^set alternative bank account name "([^"]*)"\$$`, setAlternativeBankAccountName)
	s.Step(`^set business classification to "([^"]*)"\$$`, setBusinessClassification)
	s.Step(`^set secondary identification to "([^"]*)"\$$`, setSecondaryIdentification)
	s.Step(`^account number is "([^"]*)"\$$`, accountNumberEquals)
	s.Step(`^account first name is "([^"]*)"\$$`, accountFirstNameEquals)
	s.Step(`^account business classification is "([^"]*)"\$$`, accountBusinessClassificationEquals)
	s.Step(`^account secondary identification is "([^"]*)"\$$`, accountSecondaryIdentificationEquals)
	s.Step(`^set iban to "([^"]*)"\$$`, setIban)
	s.Step(`^account country code is "([^"]*)"\$$`, accountCountryCodeEquals)
	s.Step(`^account bank id is "([^"]*)"\$$`, accountBankIdEquals)
	s.Step(`^account bic is "([^"]*)"\$$`, accountBicEquals)
	s.Step(`^I run api client Create command$`, createAccount)
	s.Step(`^api host is "([^"]*)"$`, apiHostEquals)
	s.Step(`^api endpoint is "([^"]*)"$`, apiEndpointEquals)
	s.Step(`^api client is created$`, createAPIClient)
	s.Step(`^I run api client Fetch command for same ID$`, fetchAccount)
	s.Step(`^I run api client Delete command for same ID$`, deleteAccount)
	s.Step(`^api client command Fetch fails for same ID$`, fetchAccountFails)
	s.Step(`^I List available accounts$`, listAccounts)
	s.Step(`^Delete all Listed accounts$`, deleteListedAccounts)
	s.Step(`^I Create (\d+) random accounts$`, createRandomAccounts)
	s.Step(`^I have (\d+) account\/s in my list$`, accountsInList)
	s.Step(`^I List "([^"]*)" page with Page Size (\d+)$`, listAccountsWithPageSize)
	s.Step(`^API returns an error on Create command$`, createAccountFails)
	s.Step(`^api client is healthy$`, apiClientIsHealthy)
}
