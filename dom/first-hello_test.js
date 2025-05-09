export const tests = []

tests.push(async ({ eq, page }) => {
  // Check if the class 'words' has been added in the CSS
  await eq.css('.words', { textAlign: 'center', fontFamily: 'sans-serif' })
})

tests.push(async ({ eq, page }) => {
  // Check if the torso element is initially empty
  const isEmpty = await page.$eval('#torso', (node) => !node.children.length)
  eq(isEmpty, true)
})

tests.push(async ({ eq, page }) => {
  // Click on the button
  const button = await page.$('button#speak-button')
  await button.click()

  // Check if a new text element is added in the torso
  const torsoChildren = await page.$eval('#torso', (node) =>
    [...node.children].map((child) => ({
      tag: child.tagName,
      text: child.textContent,
      class: child.className,
    })),
  )
  eq(torsoChildren, [textNode])
})

tests.push(async ({ eq, page }) => {
  // Click a second time on the button
  const button = await page.$('button#speak-button')
  await button.click()

  // Check if the text element is removed from the torso
  const isEmpty = await page.$eval('#torso', (node) => !node.children.length)
  eq(isEmpty, true)
})

tests.push(async ({ eq, page }) => {
  // Click the button once more to ensure the text is added again
  const button = await page.$('button#speak-button')
  await button.click()

  // Check if a new text element is added in the torso
  const torsoChildren = await page.$eval('#torso', (node) =>
    [...node.children].map((child) => ({
      tag: child.tagName,
      text: child.textContent,
      class: child.className,
    })),
  )
  eq(torsoChildren, [textNode])
})

const textNode = { tag: 'DIV', text: 'Hello World', class: 'words' }