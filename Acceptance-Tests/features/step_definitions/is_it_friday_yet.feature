Feature: Adding ToDo
Test Adding ToDo

  Scenario: Add Task
    Given Open ToDo app "http://localhost:8080/"
    When Reset todos and add "test-task"
    Then I should see the "test-task"

