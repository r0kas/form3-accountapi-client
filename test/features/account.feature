Feature: form3 account object
  SDK must provide a way to create valid account objects

  Scenario Template: create and validate minimal account
    Given my country code is <country_code>$
    When I create an account builder
    And set random account ID
    And set random organization ID
    And set bank ID to <bank_id>$
    And set bic to <bic>$
    Then I have a valid account
    And account bank ID code is <bank_id_code>$

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

    Scenario Template: create invalid account
      Given my country code is <country_code>$
      When I create an account builder
      And set random account ID
      And set random organization ID
      And set bank ID to <bank_id>$
      And set bic to <bic>$
      Then I have an invalid account

      Examples:
        | country_code | bank_id        | bic           |
        | "GB"         | "1"            | "ABC"         |
        | "AU"         | ""             | "ABC"         |
        | "BE"         | "1"            | ""            |
        | "CA"         | ""             | "ABC"         |
        | "FR"         | "1"            | ""            |
        | "DE"         | "1"            | ""            |
        | "GR"         | "1"            | ""            |
        | "HK"         | ""             | "ABC"         |
        | "IT"         | "1"            | ""            |
        | "LU"         | "1"            | ""            |
        | "NL"         | ""             | "ABC"         |
        | "PL"         | "1"            | ""            |
        | "PT"         | "1"            | ""            |
        | "ES"         | "1"            | ""            |
        | "CH"         | "1"            | ""            |
        | "US"         | "1"            | "ABC"         |

    Scenario Template: create account with optional attributes
      Given my country code is <country_code>$
      When I create an account builder
      And set random account ID
      And set random organization ID
      And set bank ID to <bank_id>$
      And set bic to <bic>$
      And set account number to <account_number>$
      And set first name to <first_name>$
      And set business classification to <business_class>$
      And set secondary identification to <second_id>$
      Then I have a valid account
      And account bank ID code is <bank_id_code>$
      And account number is <account_number>$
      And account first name is <first_name>$
      And account business classification is <business_class>$
      And account secondary identification is <second_id>$

      Examples:
        | country_code | bank_id_code | bank_id        | bic           | account_number | first_name | business_class | second_id |
        | "GB"         | "GBDSC"      | "123456"       | "CTBAAU2S"    | "226633"       | "Alice"    | "Business"     | "SomeID"  |
        | "AU"         | "AUBSB"      | ""             | "CTBAAU2SXXX" | "123453456633" | "Marry"    | "Personal"     | "YourID"  |
        | "BE"         | "BE"         | "123"          | ""            | "2263453433"   | "Dave"     | "Business"     | "OMG"     |
        | "CA"         | "CACPA"      | ""             | "CTBAAU2S"    | "333444633"    | "John"     | "Business"     | "NoID"    |

      Scenario Template: create account with invalid optional attributes
        Given my country code is <country_code>$
        When I create an account builder
        And set random account ID
        And set random organization ID
        And set bank ID to <bank_id>$
        And set bic to <bic>$
        And set iban to <iban>$
        And set business classification to <business_class>$
        Then I have an invalid account

      Examples:
        | country_code | bank_id        | bic           | iban                       | business_class |
        | "PT"         | "12345678"     | ""            | ""                         | "Random"       |
        | "ES"         | "12345678"     | ""            | ""                         | "Public"       |
        | "CH"         | "12345"        | ""            | ""                         | "Correct"      |
        | "US"         | "123456789"    | "CTBAAU2SXXX" | "SE3550000000054910000003" | "Personal"     |
