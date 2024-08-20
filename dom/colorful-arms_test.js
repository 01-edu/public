export const tests = []

tests.push(async ({ eq, page }) => {
  // Check if the button with id 'arm-color' exists
  const buttonExists = await page.$('button#arm-color')
  eq(!!buttonExists, true)
})

tests.push(async ({ eq, page }) => {
  // Check if the left and right arms exist
  const leftArmExists = await page.$('#arm-left')
  const rightArmExists = await page.$('#arm-right')
  eq(!!leftArmExists && !!rightArmExists, true)
})

tests.push(async ({ eq, page }) => {
  // Get the initial background colors of the arms
  const initialLeftArmColor = await page.$eval('#arm-left', node => getComputedStyle(node).backgroundColor)
  const initialRightArmColor = await page.$eval('#arm-right', node => getComputedStyle(node).backgroundColor)

  // Click the 'arm-color' button
  const button = await page.$('button#arm-color')
  await button.click()

  // Get the new background colors of the arms after clicking the button
  const newLeftArmColor = await page.$eval('#arm-left', node => getComputedStyle(node).backgroundColor)
  const newRightArmColor = await page.$eval('#arm-right', node => getComputedStyle(node).backgroundColor)

  // Check if the colors have changed and are now different from the initial colors
  eq(initialLeftArmColor !== newLeftArmColor, true)
  eq(initialRightArmColor !== newRightArmColor, true)
  eq(newLeftArmColor, newRightArmColor) // Check if both arms have the same color
})

tests.push(async ({ eq, page }) => {
  // Click the 'arm-color' button multiple times to ensure the colors keep changing
  const button = await page.$('button#arm-color')

  const armColors = []
  for (let i = 0; i < 3; i++) {
    await button.click()
    const leftArmColor = await page.$eval('#arm-left', node => getComputedStyle(node).backgroundColor)
    const rightArmColor = await page.$eval('#arm-right', node => getComputedStyle(node).backgroundColor)
    armColors.push({ leftArmColor, rightArmColor })
  }

  // Check if the colors are different in each click
  eq(new Set(armColors.map(c => c.leftArmColor)).size, armColors.length)
  eq(new Set(armColors.map(c => c.rightArmColor)).size, armColors.length)
  // Check if the arms always have the same color after each click
  armColors.forEach(colorPair => eq(colorPair.leftArmColor, colorPair.rightArmColor))
})
