# language: en

Feature: Get trading transaction records
  For me to access my trading records
  As an authenticated user
  I can previously login with my data

  Background: Login Process:
    Given I'm accessing the login page
    And I fill the login form
    And I get redirected to the pit selection

  Scenario: Extract Day Trade Orders
    Given I can access orders page
    Then I can filter orders from <start> to <end>
    And I can extract day trade orders

    Examples:
      | start      | end        |
      | 24/05/2020 | 24/05/2020 |