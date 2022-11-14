export const tests = []

export const setup = async () => {
    const getNotes = async (page) =>
        await page.$$eval('.note', (nodes) => {
          return nodes.map((note) => note.textContent)
        })
    return { getNotes }
}

const characters = `didyouhandlethekeydowneventcorrectly`

tests.push(async ({ page, eq, ctx }) => {
  // check that a note is created and matches the right letter when a key is pressed
  for (const [i, character] of characters.split('').entries()) {
    await page.keyboard.down(character)
    const typed = characters.slice(0, i + 1).split('')
    eq(await ctx.getNotes(page), typed)
  }
})

tests.push(async ({ page, eq, ctx }) => {
  // check that the last note is removed when Backspace key is pressed
  let step = 1
  while (step < 10) {
    await page.keyboard.down('Backspace')
    const typed = characters.slice(0, characters.length - step).split('')
    eq(await ctx.getNotes(page), typed)
    step++
  }
})

tests.push(async ({ page, eq, ctx }) => {
  // check that all the notes are cleared when Escape key is pressed
  await page.keyboard.down('Escape')
  const cleared = (await ctx.getNotes(page)).length === 0
  eq(cleared, true)
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
    return nodes.map((note) => note.style.backgroundColor)
  })

  const colors = [...new Set(getNotesBg)]
  const allDifferent = colors.length === test.length
  eq(allDifferent, true)
})
