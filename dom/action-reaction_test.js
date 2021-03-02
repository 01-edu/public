export const tests = []

tests.push(async ({ eq, page }) => {
  // check the initial class name of the eye left
  const eyeLeft = await page.$eval('#eye-left', (node) => node.className)
  eq(eyeLeft, 'eye')
})

tests.push(async ({ eq, page }) => {
  // click the button to close the left eye
  const button = await page.$('button')
  button.click()
  await page.waitForTimeout(150)
  const eyeLeft = await page.$eval('#eye-left', (node) => node.className)
  eq(eyeLeft, 'eye eye-closed')
})

tests.push(async ({ eq, page }) => {
  // click the button a second time to open the left eye
  const button = await page.$('button')
  button.click()
  await page.waitForTimeout(150)
  const eyeLeft = await page.$eval('#eye-left', (node) => node.className)
  eq(eyeLeft, 'eye')
})
