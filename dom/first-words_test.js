export const tests = []

tests.push(async ({ eq, page }) => {
  // check the class words has been added in the CSS
  await eq.css('.words', { textAlign: 'center', fontFamily: 'sans-serif' })
})

tests.push(async ({ eq, page }) => {
  // check the torso element is initially empty
  const isEmpty = await page.$eval('#torso', (node) => !node.children.length)
  eq(isEmpty, true)
})

tests.push(async ({ eq, page }) => {
  // click on the button
  const button = await page.$('button#speak-button')
  await button.click()

  // check a new text element is added in the torso
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
  // click a second time on the button
  const button = await page.$('button#speak-button')
  await button.click()

  // check a second new text element is added in the torso
  const torsoChildren = await page.$eval('#torso', (node) =>
    [...node.children].map((child) => ({
      tag: child.tagName,
      text: child.textContent,
      class: child.className,
    })),
  )
  eq(torsoChildren, [textNode, textNode])
})

const textNode = { tag: 'DIV', text: 'Hello there!', class: 'words' }
