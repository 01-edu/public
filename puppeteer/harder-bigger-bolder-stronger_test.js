export const tests = []

export const setup = async ({ page }) => ({
  content: await page.$$eval('div', nodes => nodes.map(n => ({
    text: n.textContent,
    size: Number((n.style.fontSize || '').slice(0, -2)),
    weight: Number(n.style.fontWeight),
  })))
})

tests.push(({ eq, content }) => {
  // should contain 120 items
  eq(content.length, 120)
})

tests.push(({ eq, content }) => {
  // content should only be one letter long
  eq(content.reduce((total, { text }) => total + text.length, 0), 120)
})

tests.push(({ eq, content }) => {
  // we expect random to yield at least 10 different letters
  eq(new Set(content).size > 10, true)
})

tests.push(({ eq, content }) => {
  // only letters from 'A' to 'Z'
  eq(content.every(({ text }) => text >= 'A' && text <= 'Z'), true)
})

tests.push(({ eq, content }) => {
  // letter size should grow

  // first should be 11
  eq(content[0].size, 11)

  // last should be 120
  eq(content[119].size, 130)

  // each letter should be bigger than the previous
  let prev = 0
  for (const { size } of content) {
    if (prev >= size) {
      throw Error('Letters should grow')
    } 
  }
})

tests.push(({ eq, content }) => {
  // letter weight should increase in thirds
  const third = n => ({ weight }) => weight === n

  // first third should be 300
  eq(content[0].weight, 300)
  eq(content[39].weight, 300)
  eq(content.slice(0, 40).every(third(300)), true)

  // second third should be 400
  eq(content[40].weight, 400)
  eq(content[79].weight, 400)
  eq(content.slice(40, 80).every(third(400)), true)

  // last third should be 600
  eq(content[80].weight, 600)
  eq(content[119].weight, 600)
  eq(content.slice(80).every(third(600)), true)
})
