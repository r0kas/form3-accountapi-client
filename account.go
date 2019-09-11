package account

import "time"

// Account represents organisation account.
// Account object provides getter methods.
// For creating new account use account Builder.
type Account struct {
	id                      string
	versionIndex            int
	createdOn               time.Time
	modifiedOn              time.Time
	organizationID          string
	country                 string
	baseCurrency            string
	bankID                  string
	bankIDCode              string
	accountNumber           string
	bic                     string
	iban                    string
	customerID              string
	title                   string
	firstName               string
	bankAccountName         string
	altBankAccountNames     []string
	accountClassification   string
	jointAccount            bool
	accountMatchingOptOut   bool
	secondaryIdentification string
}

// ID unique identifier (UUID) of an account.
func (acc *Account) ID() string {
	return acc.id
}

// Version represents version number of the account in API database.
func (acc *Account) Version() int {
	return acc.versionIndex
}

// ModifiedOn returns time when account was last modified
func (acc *Account) ModifiedOn() time.Time {
	return acc.modifiedOn
}

// CreatedOn returns time when account was created
func (acc *Account) CreatedOn() time.Time {
	return acc.createdOn
}

// OrganizationID returns string of Organization ID which is unique identifier (UUID)
func (acc *Account) OrganizationID() string {
	return acc.organizationID
}

// Country returns account country code in ISO 3166 format.
func (acc *Account) Country() string {
	return acc.country
}

// BaseCurrency returns set account base currency in ISO 4217 format.
func (acc *Account) BaseCurrency() string {
	return acc.baseCurrency
}

// BankID returns Local country bank identifier. Format depends on the country.
func (acc *Account) BankID() string {
	return acc.bankID
}

// BankIDCode identifies the type of bank ID being used
func (acc *Account) BankIDCode() string {
	return acc.bankIDCode
}

// AccountNumber a unique account number.
func (acc *Account) AccountNumber() string {
	return acc.accountNumber
}

// Bic SWIFT BIC in either 8 or 11 character format e.g. 'NWBKGB22'
func (acc *Account) Bic() string {
	return acc.bic
}

// Iban of the account
func (acc *Account) Iban() string {
	return acc.iban
}

// CustomerID A free-format reference that can be used to link this account to an external system
func (acc *Account) CustomerID() string {
	return acc.customerID
}

// Title - account holder's title, e.g. Ms, Dr, Mr.
func (acc *Account) Title() string {
	return acc.title
}

// FirstName - account holder's first name.
func (acc *Account) FirstName() string {
	return acc.firstName
}

// BankAccountName - primary account name, used for Confirmation of Payee matching.
func (acc *Account) BankAccountName() string {
	return acc.bankAccountName
}

// AltBankAccountNames - Up to 3 alternative account names, used for Confirmation of Payee matching.
func (acc *Account) AltBankAccountNames() []string {
	return acc.altBankAccountNames
}

// AccountClassification - Classification of account. Can be either Personal or Business. Defaults to Personal.
func (acc *Account) AccountClassification() string {
	return acc.accountClassification
}

// IsJointAccount - True if this is a joint account.
// Defaults to false.
func (acc *Account) IsJointAccount() bool {
	return acc.jointAccount
}

// IsAccountMatchingOptOut - True if the account has opted out of account matching, e.g. Confirmation of Payee.
// Defaults to false.
func (acc *Account) IsAccountMatchingOptOut() bool {
	return acc.accountMatchingOptOut
}

// SecondaryIdentification e.g. building society roll number. Used for Confirmation of Payee.
func (acc *Account) SecondaryIdentification() string {
	return acc.secondaryIdentification
}

// used for generating rest transport structures
func (acc *Account) attributes() *accountAttributes {
	return &accountAttributes{
		Country:                 acc.Country(),
		BaseCurrency:            acc.BaseCurrency(),
		BankID:                  acc.BankID(),
		BankIDCode:              acc.BankIDCode(),
		AccountNumber:           acc.AccountNumber(),
		Bic:                     acc.Bic(),
		Iban:                    acc.Iban(),
		CustomerID:              acc.CustomerID(),
		Title:                   acc.Title(),
		FirstName:               acc.FirstName(),
		BankAccountName:         acc.BankAccountName(),
		AltBankAccountNames:     acc.AltBankAccountNames(),
		AccountClassification:   acc.AccountClassification(),
		JointAccount:            acc.IsJointAccount(),
		AccountMatchingOptOut:   acc.IsAccountMatchingOptOut(),
		SecondaryIdentification: acc.SecondaryIdentification(),
	}
}

// used for creating account object from received json transport structure
func accountFrom(response transportData) *Account {
	return &Account{
		id:                      response.ID,
		versionIndex:            response.Version,
		createdOn:               response.CreatedOn,
		modifiedOn:              response.ModifiedOn,
		organizationID:          response.OrganizationID,
		country:                 response.Attributes.Country,
		baseCurrency:            response.Attributes.BaseCurrency,
		bankID:                  response.Attributes.BankID,
		bankIDCode:              response.Attributes.BankIDCode,
		accountNumber:           response.Attributes.AccountNumber,
		bic:                     response.Attributes.Bic,
		iban:                    response.Attributes.Iban,
		customerID:              response.Attributes.CustomerID,
		title:                   response.Attributes.Title,
		firstName:               response.Attributes.FirstName,
		bankAccountName:         response.Attributes.BankAccountName,
		altBankAccountNames:     response.Attributes.AltBankAccountNames,
		accountClassification:   response.Attributes.AccountClassification,
		jointAccount:            response.Attributes.JointAccount,
		accountMatchingOptOut:   response.Attributes.AccountMatchingOptOut,
		secondaryIdentification: response.Attributes.SecondaryIdentification,
	}
}
