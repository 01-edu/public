export const tests = []

export const setup = async ({ page }) => ({
  getNotes: async () =>
    await page.$$eval('.note', (nodes) => {
      return nodes.map((note) => note.textContent)
    }),
})

const characters = `didyouhandlethekeydowneventcorrectly`

tests.push(async ({ page, eq, getNotes }) => {
  // check that a note is created and matches the right letter when a key is pressed
  for (const [i, character] of characters.split('').entries()) {
    await page.keyboard.down(character)
    const typed = characters.slice(0, i + 1).split('')
    eq(await getNotes(), typed)
  }
})

tests.push(async ({ page, eq, getNotes }) => {
  // check that the last note is removed when Backspace key is pressed
  let step = 1
  while (step < 10) {
    await page.keyboard.down('Backspace')
    const typed = characters.slice(0, characters.length - step).split('')
    eq(await getNotes(), typed)
    step++
  }
})

tests.push(async ({ page, eq, getNotes }) => {
  // check that all the notes are cleared when Escape key is pressed
  await page.keyboard.down('Escape')
  const cleared = (await getNotes()).length === 0
  eq(await cleared, true)
})

tests.push(async ({ page, eq }) => {
  // check that notes have different background colors
  const test = 'abcdefghijklmnopqrstuvwxyz'
  let step = 0
  while (step < test.length) {
    await page.keyboard.down(test[step])
    step++
  }

  const getNotesBg = await page.$$eval('.note', (nodes) => {
    return nodes.map((note) => note.style.background)
  })

  const colors = [...new Set(getNotesBg)]
  const allDifferent = colors.length === test.length
  eq(allDifferent, true)
})
