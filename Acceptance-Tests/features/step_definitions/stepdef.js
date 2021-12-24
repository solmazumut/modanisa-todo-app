const assert = require('assert');
const { Given, When, Then, AfterAll } = require('@cucumber/cucumber');
const { Builder, By, Capabilities, Key } = require('selenium-webdriver');
const { expect } = require('chai');

// driver setup
const driver = new Builder().forBrowser("firefox").build();



Given('Open ToDo app {string}', async function (url) {
  // Write code here that turns the phrase above into concrete actions
  await driver.get(url);
});

When('Reset todos and add {string}', async function (task) {
  const resetButton = await driver.findElement(By.id('reset-button'));
  await resetButton.click();

  const inputBox = await driver.findElement(By.id('input-box'));
  await inputBox.sendKeys(task, Key.RETURN);

  const addButton = await driver.findElement(By.id('add-button'));
  await addButton.click();
});

Then('I should see the {string}', {timeout: 60 * 1000}, async function (task) {
  const taskName = await driver.findElement(By.tagName("td")).getText();
  await driver.quit();
  expect(task).to.equal(taskName);
});
