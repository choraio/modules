Feature: Content

  Scenario: valid agent
    Given agent
    """
    {
      "address": "chora1r05jctvye5rweu4f6n3gle0ycaww6xpueccqpy63erz75gfy3p7snu2hw4",
      "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
    }
    """
    When validate agent
    Then expect no error

  Scenario: an error is returned if address is empty
    Given agent
    """
    {}
    """
    When validate agent
    Then expect the error
    """
    address: empty address string is not allowed: parse error
    """

  Scenario: an error is returned if agent is empty
    Given agent
    """
    {
      "address": "chora1r05jctvye5rweu4f6n3gle0ycaww6xpueccqpy63erz75gfy3p7snu2hw4"
    }
    """
    When validate agent
    Then expect the error
    """
    admin: empty address string is not allowed: parse error
    """

  Scenario: an error is returned if metadata is empty
    Given agent
    """
    {
      "address": "chora1r05jctvye5rweu4f6n3gle0ycaww6xpueccqpy63erz75gfy3p7snu2hw4",
      "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    When validate agent
    Then expect the error
    """
    metadata: empty string is not allowed: parse error
    """

  Scenario: an error is returned if metadata exceeds 128 characters
    Given agent
    """
    {
      "address": "chora1r05jctvye5rweu4f6n3gle0ycaww6xpueccqpy63erz75gfy3p7snu2hw4",
      "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    And metadata with length "129"
    When validate agent
    Then expect the error
    """
    metadata: exceeds max length 128: parse error
    """
