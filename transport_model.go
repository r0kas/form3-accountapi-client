package account

import "time"

type (
	restTransport struct {
		Data  transportData `json:"data"`
		Links links         `json:"links,omitempty"`
	}

	restTransportList struct {
		Data  []transportData `json:"data"`
		Links links           `json:"links,omitempty"`
	}

	transportData struct {
		Type           string            `json:"type"`
		ID             string            `json:"id"`
		OrganizationID string            `json:"organisation_id"`
		Version        int               `json:"version,omitempty"`
		CreatedOn      time.Time         `json:"created_on,omitempty"`
		ModifiedOn     time.Time         `json:"modified_on,omitempty"`
		Attributes     accountAttributes `json:"attributes"`
	}

	accountAttributes struct {
		Country                 string   `json:"country"`
		BaseCurrency            string   `json:"base_currency,omitempty"`
		BankID                  string   `json:"bank_id,omitempty"`
		BankIDCode              string   `json:"bank_id_code,omitempty"`
		AccountNumber           string   `json:"account_number,omitempty"`
		Bic                     string   `json:"bic,omitempty"`
		Iban                    string   `json:"iban,omitempty"`
		CustomerID              string   `json:"customer_id,omitempty"`
		Title                   string   `json:"title,omitempty"`
		FirstName               string   `json:"first_name,omitempty"`
		BankAccountName         string   `json:"bank_account_name,omitempty"`
		AltBankAccountNames     []string `json:"alternative_bank_account_names,omitempty"`
		AccountClassification   string   `json:"account_classification"`
		JointAccount            bool     `json:"joint_account"`
		AccountMatchingOptOut   bool     `json:"account_matching_opt_out"`
		SecondaryIdentification string   `json:"secondary_identification,omitempty"`
	}

	links struct {
		First string `json:"first,omitempty"`
		Last  string `json:"last,omitempty"`
		Next  string `json:"next,omitempty"`
		Self  string `json:"self,omitempty"`
	}
)
