export const tests = []

tests.push(async ({ eq, page }) => {
  // check the initial class name of the eye left
  const eyeLeft = await page.$eval('#eye-left', (node) => node.className)
  eq(eyeLeft, 'eye')

  // check that the text of the button says 'close'
  const buttonText = await page.$eval('button', (node) => node.textContent)
  eq(buttonText, 'Click to close the left eye')
})

tests.push(async ({ eq, page }) => {
  // click the button to close the left eye
  const button = await page.$('button')
  button.click()

  // check that the class has been added
  await page.waitForSelector('#eye-left.eye.eye-closed', { timeout: 150 })

  // check the background color has changed
  const eyeLeft = await page.$eval(
    '#eye-left.eye.eye-closed',
    (node) => node.style.backgroundColor,
  )
  eq(eyeLeft, 'black')

  // check that the text of the button changed to 'open'
  const buttonText = await page.$eval('button', (node) => node.textContent)
  eq(buttonText, 'Click to open the left eye')
})

tests.push(async ({ eq, page }) => {
  // click the button a second time to open the left eye
  const button = await page.$('button')
  button.click()

  // check that the class has been removed
  await page.waitForSelector('#eye-left.eye:not(.eye-closed)', { timeout: 150 })

  // check the background color has changed
  const eyeLeft = await page.$eval(
    '#eye-left.eye:not(.eye-closed)',
    (node) => node.style.backgroundColor,
  )
  eq(eyeLeft, 'red')

  // check that the text of the button changed to 'close'
  const buttonText = await page.$eval('button', (node) => node.textContent)
  eq(buttonText, 'Click to close the left eye')
})
