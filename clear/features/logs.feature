# language: en

Feature: Get trading transaction records
  For me to access my trading records
  As an authenticated user
  I can previously login with my data

  Background: Login Form:
    Given I'm accessing the login page

  Scenario: Successful Login
    When I fill the login form
    Then I get redirected to the pit selection
    And I can access the orders page