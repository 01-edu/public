export const tests = []

tests.push(async ({ eq, page }) => {
  // check the initial class name of the eye left
  const eyeLeft = await page.$eval('#eye-left', (node) => node.className)
  eq(eyeLeft, 'eye')

  const buttonText = await page.$eval('button', (node) => node.textContent)
  // check that the text of the button says 'close'
  eq(buttonText, 'Click to close the left eye')
})

tests.push(async ({ eq, page }) => {
  // click the button to close the left eye
  const button = await page.$('button')
  button.click()
  await page.waitForTimeout(150)

  // check that the class has been added
  const eyeLeft = await page.$eval('#eye-left', (node) => node.className)
  eq(eyeLeft, 'eye eye-closed')

  // check that the text of the button changed to 'open'
  const buttonText = await page.$eval('button', (node) => node.textContent)
  eq(buttonText, 'Click to open the left eye')
})

tests.push(async ({ eq, page }) => {
  // click the button a second time to open the left eye
  const button = await page.$('button')
  button.click()
  await page.waitForTimeout(150)

  // check that the class has been removed
  const eyeLeft = await page.$eval('#eye-left', (node) => node.className)
  eq(eyeLeft, 'eye')

  const buttonText = await page.$eval('button', (node) => node.textContent)
  // check that the text of the button changed to 'close'
  eq(buttonText, 'Click to close the left eye')
})
