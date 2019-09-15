Feature: form3 api client
  API Client must send valid requests and handle responses

  Background:
    Given api host is "http://localhost:8080"
    And api endpoint is "/v1/organisation/accounts"
    Then api client is created
    And api client is healthy

  Scenario Template: Create new account
    Given my country code is <country_code>$
    And I create an account builder
    And set random account ID
    And set random organization ID
    And set bank ID to <bank_id>$
    And set bic to <bic>$
    And I have a valid account
    When I run api client Create command
    Then account country code is <country_code>$
    And account bank ID code is <bank_id_code>$
    And account bank id is <bank_id>$
    And account bic is <bic>$

    Examples:
      | country_code | bank_id_code | bank_id        | bic           |
      | "GB"         | "GBDSC"      | "123456"       | "CTBAAU2S"    |
      | "AU"         | "AUBSB"      | ""             | "CTBAAU2SXXX" |
      | "BE"         | "BE"         | "123"          | ""            |
      | "CA"         | "CACPA"      | ""             | "CTBAAU2S"    |
      | "FR"         | "FR"         | "0123456789"   | ""            |
      | "DE"         | "DEBLZ"      | "12345678"     | ""            |
      | "GR"         | "GRBIC"      | "1234567"      | ""            |
      | "HK"         | "HKNCC"      | ""             | "CTBAAU2S"    |
      | "IT"         | "ITNCC"      | "01234567890"  | ""            |
      | "LU"         | "LULUX"      | "123"          | ""            |
      | "NL"         | ""           | ""             | "CTBAAU2S"    |
      | "PL"         | "PLKNR"      | "12345678"     | ""            |
      | "PT"         | "PTNCC"      | "12345678"     | ""            |
      | "ES"         | "ESNCC"      | "12345678"     | ""            |
      | "CH"         | "CHBCC"      | "12345"        | ""            |
      | "US"         | "USABA"      | "123456789"    | "CTBAAU2SXXX" |

  Scenario Template: Fetch created account
    Given my country code is <country_code>$
    And I create an account builder
    And set random account ID
    And set random organization ID
    And set bank ID to <bank_id>$
    And set bic to <bic>$
    And I have a valid account
    And I run api client Create command
    When I run api client Fetch command for same ID
    Then account country code is <country_code>$
    And account bank ID code is <bank_id_code>$
    And account bank id is <bank_id>$
    And account bic is <bic>$

    Examples:
      | country_code | bank_id_code | bank_id        | bic           |
      | "GB"         | "GBDSC"      | "123456"       | "CTBAAU2S"    |
      | "AU"         | "AUBSB"      | ""             | "CTBAAU2SXXX" |
      | "BE"         | "BE"         | "123"          | ""            |
      | "CA"         | "CACPA"      | ""             | "CTBAAU2S"    |
      | "FR"         | "FR"         | "0123456789"   | ""            |
      | "DE"         | "DEBLZ"      | "12345678"     | ""            |
      | "GR"         | "GRBIC"      | "1234567"      | ""            |
      | "HK"         | "HKNCC"      | ""             | "CTBAAU2S"    |
      | "IT"         | "ITNCC"      | "01234567890"  | ""            |
      | "LU"         | "LULUX"      | "123"          | ""            |
      | "NL"         | ""           | ""             | "CTBAAU2S"    |
      | "PL"         | "PLKNR"      | "12345678"     | ""            |
      | "PT"         | "PTNCC"      | "12345678"     | ""            |
      | "ES"         | "ESNCC"      | "12345678"     | ""            |
      | "CH"         | "CHBCC"      | "12345"        | ""            |
      | "US"         | "USABA"      | "123456789"    | "CTBAAU2SXXX" |

  Scenario Template: Delete created account
    Given my country code is <country_code>$
    And I create an account builder
    And set random account ID
    And set random organization ID
    And set bank ID to <bank_id>$
    And set bic to <bic>$
    And I have a valid account
    And I run api client Create command
    When I run api client Delete command for same ID
    Then api client command Fetch fails for same ID

    Examples:
      | country_code | bank_id        | bic           |
      | "GB"         | "123456"       | "CTBAAU2S"    |
      | "AU"         | ""             | "CTBAAU2SXXX" |
      | "BE"         | "123"          | ""            |
      | "CA"         | ""             | "CTBAAU2S"    |
      | "FR"         | "0123456789"   | ""            |
      | "DE"         | "12345678"     | ""            |
      | "GR"         | "1234567"      | ""            |
      | "HK"         | ""             | "CTBAAU2S"    |
      | "IT"         | "01234567890"  | ""            |
      | "LU"         | "123"          | ""            |
      | "NL"         | ""             | "CTBAAU2S"    |
      | "PL"         | "12345678"     | ""            |
      | "PT"         | "12345678"     | ""            |
      | "ES"         | "12345678"     | ""            |
      | "CH"         | "12345"        | ""            |
      | "US"         | "123456789"    | "CTBAAU2SXXX" |

  Scenario: List, Delete, Create and List accounts again
    Given I List available accounts
    And Delete all Listed accounts
    When I List available accounts
    And I have 0 account/s in my list
    Then I Create 10 random accounts
    When I List available accounts
    Then I have 10 account/s in my list
    When I List "first" page with Page Size 3
    Then I have 3 account/s in my list
    When I List "last" page with Page Size 3
    Then I have 1 account/s in my list
    When I List "2" page with Page Size 3
    Then I have 3 account/s in my list

  Scenario Template: API returns an error with wrong input
    Given my country code is <country_code>$
    And I create an account builder
    And set random account ID
    And set random organization ID
    And set bank ID to <bank_id>$
    And set bic to <bic>$
    And set iban to <iban>$
    And I have a valid account
    When API returns an error on Create command

    Examples:
      | country_code | bank_id        | bic           | iban   |
      | "GB"         | "123456"       | "CTBAAU2S"    | "xxxx" |
      | "BE"         | "123"          | ""            | "xxxx" |
      | "FR"         | "0123456789"   | ""            | "xxxx" |
      | "DE"         | "12345678"     | ""            | "xxxx" |
      | "GR"         | "1234567"      | ""            | "xxxx" |
      | "IT"         | "01234567890"  | ""            | "xxxx" |
      | "LU"         | "123"          | ""            | "xxxx" |
      | "NL"         | ""             | "CTBAAU2S"    | "xxxx" |
      | "PL"         | "12345678"     | ""            | "xxxx" |
      | "PT"         | "12345678"     | ""            | "xxxx" |
      | "ES"         | "12345678"     | ""            | "xxxx" |
      | "CH"         | "12345"        | ""            | "xxxx" |
