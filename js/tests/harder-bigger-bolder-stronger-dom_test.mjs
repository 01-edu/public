export const tests = []

export const setup = async ({ page }) => ({
  content: await page.$$eval('div', nodes => nodes.map(n => ({
    text: n.textContent,
    size: Number((n.style.fontSize || '').slice(0, -2)),
    weight: Number(n.style.fontWeight),
  })))
})

tests.push(({ eq, ctx }) => {
  // should contain 120 items
  eq(ctx.content.length, 120)
})

tests.push(({ eq, ctx }) => {
  // ctx.content should only be one letter long
  eq(ctx.content.reduce((total, { text }) => total + text.length, 0), 120)
})

tests.push(({ eq, ctx }) => {
  // we expect random to yield at least 10 different letters
  eq(new Set(ctx.content).size > 10, true)
})

tests.push(({ eq, ctx }) => {
  // only letters from 'A' to 'Z'
  eq(ctx.content.every(({ text }) => text >= 'A' && text <= 'Z'), true)
})

tests.push(({ eq, ctx }) => {
  // letter size should grow

  // first should be 11
  eq(ctx.content[0].size, 11)

  // last should be 120
  eq(ctx.content[119].size, 130)

  // each letter should be bigger than the previous
  let prev = 0
  for (const { size } of ctx.content) {
    if (prev >= size) {
      throw Error('Letters should grow')
    } 
  }
})

tests.push(({ eq, ctx }) => {
  // letter weight should increase in thirds
  const third = n => ({ weight }) => weight === n

  // first third should be 300
  eq(ctx.content[0].weight, 300)
  eq(ctx.content[39].weight, 300)
  eq(ctx.content.slice(0, 40).every(third(300)), true)

  // second third should be 400
  eq(ctx.content[40].weight, 400)
  eq(ctx.content[79].weight, 400)
  eq(ctx.content.slice(40, 80).every(third(400)), true)

  // last third should be 600
  eq(ctx.content[80].weight, 600)
  eq(ctx.content[119].weight, 600)
  eq(ctx.content.slice(80).every(third(600)), true)
})
