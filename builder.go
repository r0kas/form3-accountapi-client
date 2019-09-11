package account

import (
	"github.com/pkg/errors"
	"golang.org/x/text/currency"
	"gopkg.in/go-playground/validator.v9"
)

type (
	// Builder provides methods to create and validate account objects.
	///////
	// I was after creating a Builder which would have compile time validation,
	// but as there was vast amount of attributes and format differences for given countries,
	// compile time validation would leave with ugly struct design with numerous interfaces
	// which would not be maintainable in the long run.
	// Hence chose runtime attributes validation written in clean idiomatic manner using field tags in struct.
	// Also builder is the only way of creating and setting account objects, which mitigates the misuse of SDK.
	///////
	Builder struct {
		essential *essentialAttributes
		optional  *optionalAttributes
		validate  *validator.Validate
	}

	///////
	// Validation tags are based on country codes, hence it's easy to differentiate and
	// also trivial to set the validator tag as country code is the core attribute of account builder.
	//////
	essentialAttributes struct {
		ID             string `validate:"uuid,required"`
		OrganizationID string `validate:"uuid,required"`
		Country        string `validate:"eq=GB|eq=AU|eq=BE|eq=CA|eq=FR|eq=DE|eq=GR|eq=HK|eq=IT|eq=LU|eq=NL|eq=PL|eq=PT|eq=ES|eq=ES|eq=CH|eq=US"`
		BankIDCode     string
		BankID         string `GB:"len=6" BE:"len=3" FR:"len=10" DE:"len=8" GR:"len=7" IT:"len=10|len=11" LU:"len=3" NL:"len=0" PL:"len=8" PT:"len=8" ES:"len=8" CH:"len=5" US:"len=9"`
		Bic            string `GB:"len=8|len=11" AU:"len=8|len=11" CA:"len=8|len=11" HK:"len=8|len=11" NL:"len=8|len=11" US:"len=8|len=11"`
		Iban           string `AU:"len=0" CA:"len=0" HK:"len=0" US:"len=0"`
	}

	///////
	// Chose not to add all attributes on the same builder level.
	// Having a structure and separating optional attributes from essential ones lowers cognitive load for client user.
	///////
	optionalAttributes struct {
		Builder                 *Builder
		VersionIndex            int
		AccountNumber           string
		BaseCurrency            string
		CustomerID              string
		Title                   string   `validate:"max=40"`
		FirstName               string   `validate:"max=40"`
		BankAccountName         string   `validate:"max=140"`
		AltBankAccountNames     []string `validate:"max=3,dive,max=140"`
		AccountClassification   string   `validate:"eq=Personal|eq=Business"`
		SecondaryIdentification string   `validate:"max=140"`
		JointAccount            bool
		AccountMatchingOptOut   bool
	}
	// OptionalAttributes has a collection of methods to set optional account attributes
	///////
	// This interface enabled to hide public optional attributes struct fields which were required for validation.
	///////
	OptionalAttributes interface {
		SetVersion(int) *Builder
		SetAccountNumber(string) *Builder
		SetBaseCurrency(currency.Unit) *Builder
		SetCustomerID(string) *Builder
		SetTitle(string) *Builder
		SetFirstName(string) *Builder
		SetBankAccountName(string) *Builder
		SetAltBankAccountNames(...string) *Builder
		SetAccountClassification(string) *Builder
		SetJointAccount(bool) *Builder
		SetAccountMatchingOptOut(bool) *Builder
		SetSecondaryIdentification(string) *Builder
	}
)

// NewBuilder creates account builder from provided Country.
// Country sets validation rules for created accounts.
func NewBuilder(country Country) *Builder {
	return &Builder{
		essential: &essentialAttributes{
			Country:    country.Code(),
			BankIDCode: country.BankIDCode(),
		},
		optional: &optionalAttributes{
			AccountClassification: "Personal",
		},
		validate: validator.New(),
	}
}

// CastBuilderFrom creates an account builder from existing account object.
// This enables to modify, validate and create new account object.
///////
// This function would be essential for PATCH operation.
///////
func CastBuilderFrom(account *Account) *Builder {
	return &Builder{
		essential: &essentialAttributes{
			ID:             account.ID(),
			OrganizationID: account.OrganizationID(),
			Country:        account.Country(),
			BankIDCode:     account.BankIDCode(),
			BankID:         account.BankID(),
			Bic:            account.Bic(),
			Iban:           account.Iban(),
		},
		optional: &optionalAttributes{
			VersionIndex:            account.Version(),
			AccountNumber:           account.AccountNumber(),
			BaseCurrency:            account.BaseCurrency(),
			CustomerID:              account.CustomerID(),
			Title:                   account.Title(),
			FirstName:               account.FirstName(),
			BankAccountName:         account.BankAccountName(),
			AltBankAccountNames:     account.AltBankAccountNames(),
			AccountClassification:   account.AccountClassification(),
			SecondaryIdentification: account.SecondaryIdentification(),
			JointAccount:            account.IsJointAccount(),
			AccountMatchingOptOut:   account.IsAccountMatchingOptOut(),
		},
		validate: validator.New(),
	}
}

// SetID of an account. Unique identifier (UUID) string - required for all accounts.
func (b *Builder) SetID(id string) *Builder {
	b.essential.ID = id
	return b
}

// SetOrganizationID of an account. Unique identifier (UUID) string - required for all accounts.
func (b *Builder) SetOrganizationID(id string) *Builder {
	b.essential.OrganizationID = id
	return b
}

// SetBankID - Local country bank identifier. Format depends on the country. Required for most countries.
func (b *Builder) SetBankID(bankID string) *Builder {
	b.essential.BankID = bankID
	return b
}

// SetBic - SWIFT BIC in either 8 or 11 character format e.g. 'NWBKGB22'
func (b *Builder) SetBic(bic string) *Builder {
	b.essential.Bic = bic
	return b
}

// SetIban - IBAN of the account. Will be calculated from other fields if not supplied.
func (b *Builder) SetIban(iban string) *Builder {
	b.essential.Iban = iban
	return b
}

// Validate checks set fields based on country code.
// Returns account object if no errors are generated during validation.
func (b *Builder) Validate() (*Account, error) {
	b.validate.SetTagName(b.essential.Country)
	if err := validateStruct(b.validate, b.essential); err != nil {
		return nil, err
	}
	b.validate.SetTagName("validate")
	if err := validateStruct(b.validate, b.essential); err != nil {
		return nil, err
	}
	if err := validateStruct(b.validate, b.optional); err != nil {
		return nil, err
	}
	return &Account{
		id:                      b.essential.ID,
		organizationID:          b.essential.OrganizationID,
		versionIndex:            b.optional.VersionIndex,
		country:                 b.essential.Country,
		bankIDCode:              b.essential.BankIDCode,
		bankID:                  b.essential.BankID,
		bic:                     b.essential.Bic,
		iban:                    b.essential.Iban,
		baseCurrency:            b.optional.BaseCurrency,
		accountNumber:           b.optional.AccountNumber,
		customerID:              b.optional.CustomerID,
		title:                   b.optional.Title,
		firstName:               b.optional.FirstName,
		bankAccountName:         b.optional.BankAccountName,
		altBankAccountNames:     b.optional.AltBankAccountNames,
		accountClassification:   b.optional.AccountClassification,
		jointAccount:            b.optional.JointAccount,
		accountMatchingOptOut:   b.optional.AccountMatchingOptOut,
		secondaryIdentification: b.optional.SecondaryIdentification,
	}, nil
}

// SetOptionalAttribute returns a list of methods for setting optional account attributes.
func (b *Builder) SetOptionalAttribute() OptionalAttributes {
	return &optionalAttributes{
		Builder: b,
	}
}

// SetAccountNumber - A unique account number will automatically be generated if not provided.
func (opt *optionalAttributes) SetAccountNumber(accountNumber string) *Builder {
	opt.Builder.optional.AccountNumber = accountNumber
	return opt.Builder
}

// SetVersion - version number of account object. Needs to be incremented when Patching an existing account.
func (opt *optionalAttributes) SetVersion(version int) *Builder {
	opt.Builder.optional.VersionIndex = version
	return opt.Builder
}

// SetBaseCurrency - ISO 4217 code used to identify the base currency of the account, e.g. 'GBP', 'EUR'
// Provide currency unit object from golang text library.
func (opt *optionalAttributes) SetBaseCurrency(unit currency.Unit) *Builder {
	opt.Builder.optional.BaseCurrency = unit.String()
	return opt.Builder
}

// SetCustomerID - A free-format reference that can be used to link this account to an external system
func (opt *optionalAttributes) SetCustomerID(customerID string) *Builder {
	opt.Builder.optional.CustomerID = customerID
	return opt.Builder
}

// SetTitle - The account holder's title, e.g. Ms, Dr, Mr.
// Valid up to string[40]
func (opt *optionalAttributes) SetTitle(title string) *Builder {
	opt.Builder.optional.Title = title
	return opt.Builder
}

// SetFirstName - The account holder's first name.
// Valid up to string[40]
func (opt *optionalAttributes) SetFirstName(firstName string) *Builder {
	opt.Builder.optional.FirstName = firstName
	return opt.Builder
}

// SetBankAccountName - Primary account name, used for Confirmation of Payee matching.
// Required if Confirmation of Payee is enabled for the organisation.
// Valid up to string[140]
func (opt *optionalAttributes) SetBankAccountName(accountName string) *Builder {
	opt.Builder.optional.BankAccountName = accountName
	return opt.Builder
}

// SetAltBankAccountNames - Up to 3 alternative account names, used for Confirmation of Payee matching.
// Each element valid up to string[140]
func (opt *optionalAttributes) SetAltBankAccountNames(names ...string) *Builder {
	opt.Builder.optional.AltBankAccountNames = names
	return opt.Builder
}

// SetAccountClassification - Classification of account. Can be either Personal or Business.
// Defaults to Personal.
func (opt *optionalAttributes) SetAccountClassification(accountClassification string) *Builder {
	opt.Builder.optional.AccountClassification = accountClassification
	return opt.Builder
}

// SetJointAccount - set to True if this is a joint account.
// Defaults to false.
func (opt *optionalAttributes) SetJointAccount(isJointAccount bool) *Builder {
	opt.Builder.optional.JointAccount = isJointAccount
	return opt.Builder
}

// SetAccountMatchingOptOut - set to True if the account has opted out of account matching, e.g. Confirmation of Payee.
// Defaults to false.
func (opt *optionalAttributes) SetAccountMatchingOptOut(isAccountMatching bool) *Builder {
	opt.Builder.optional.AccountMatchingOptOut = isAccountMatching
	return opt.Builder
}

// SetSecondaryIdentification - Secondary identification, e.g. building society roll number.
// Used for Confirmation of Payee.
// Valid up to string[140]
func (opt *optionalAttributes) SetSecondaryIdentification(secondaryID string) *Builder {
	opt.Builder.optional.SecondaryIdentification = secondaryID
	return opt.Builder
}

func validateStruct(validate *validator.Validate, s interface{}) (err error) {
	err = validate.Struct(s)
	if err != nil {
		if len(err.(validator.ValidationErrors)) > 0 {
			return errors.New(err.(validator.ValidationErrors).Error())
		}
	}
	return
}
