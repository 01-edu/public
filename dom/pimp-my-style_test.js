import { styles } from './subjects/pimp-my-style/pimp-my-style.data.js'

export const tests = []

const formatClass = (limit, unpimp) =>
  ['button', ...styles.slice(0, limit), unpimp && 'unpimp'].filter(Boolean)

const max = styles.length - 1

export const setup = async ({ page }) => {
  const btn = await page.$('.button')

  return {
    btn,
    getClass: async () =>
      (await (await btn.getProperty('className')).jsonValue()).split(' '),
  }
}

tests.push(async ({ page, eq, btn, getClass }) => {
  // pimp
  for (const i of styles.keys()) {
    console.log('pimp click', i + 1)
    await btn.click()
    eq(formatClass(i + 1, i === max), await getClass())
  }
})

tests.push(async ({ page, eq, btn, getClass }) => {
  // unpimp !
  for (const i of styles.keys()) {
    console.log('unpimp click', i + 1)
    await btn.click()
    eq(formatClass(max - i, i !== max), await getClass())
  }
})
