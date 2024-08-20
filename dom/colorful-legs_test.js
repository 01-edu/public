export const tests = []

tests.push(async ({ eq, page }) => {
  // Click on the button to change the robot's leg colors
  const button = await page.$('button#leg-color')
  await button.click()

  // Get the new colors of both legs
  const legLeftColor = await page.$eval('#leg-left', (node) => getComputedStyle(node).backgroundColor)
  const legRightColor = await page.$eval('#leg-right', (node) => getComputedStyle(node).backgroundColor)

  // Check if both legs have been assigned the same new color
  eq(legLeftColor, legRightColor)

  // Ensure the new color is not black
  eq(legLeftColor !== 'rgb(0, 0, 0)', true, 'The color of the legs should not be black')
})

tests.push(async ({ eq, page }) => {
  // Get the initial colors of the legs before clicking the button
  const initialLegLeftColor = await page.$eval('#leg-left', (node) => getComputedStyle(node).backgroundColor)
  const initialLegRightColor = await page.$eval('#leg-right', (node) => getComputedStyle(node).backgroundColor)

  // Click on the button to change the robot's leg colors
  const button = await page.$('button#leg-color')
  await button.click()

  // Get the new colors of both legs
  const newLegLeftColor = await page.$eval('#leg-left', (node) => getComputedStyle(node).backgroundColor)
  const newLegRightColor = await page.$eval('#leg-right', (node) => getComputedStyle(node).backgroundColor)

  // Check if both legs have been assigned the same new color
  eq(newLegLeftColor, newLegRightColor)

  // Ensure the new color is different from the initial color
  eq(newLegLeftColor !== initialLegLeftColor, true, 'The color of the legs should be different from the initial color')
})
