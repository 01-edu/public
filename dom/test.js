import http from 'http'
import fs from 'fs'
import path from 'path'
import { deepStrictEqual } from 'assert'
import puppeteer from 'puppeteer'

const exercise = process.argv[2]
if (!exercise) throw Error(`usage: node test EXERCISE_NAME`)
const PORT = 9898
const config = {
  args: [
    '--no-sandbox',
    '--disable-setuid-sandbox',

    // This will write shared memory files into /tmp instead of /dev/shm,
    // because Docker’s default for /dev/shm is 64MB
    '--disable-dev-shm-usage',
  ],
  headless: !process.env.SOLUTION_PATH,
}

const solutionPath = process.env.SOLUTION_PATH || '/jail/student'
const mediaTypes = {
  jpg: 'image/jpeg',
  png: 'image/png',
  html: 'text/html',
  css: 'text/css',
  js: 'application/javascript',
  json: 'application/json',
}

const random = (min, max = min) => {
  max === min && (min = 0)
  min = Math.ceil(min)
  return Math.floor(Math.random() * (Math.floor(max) - min + 1)) + min
}

const rgbToHsl = rgbStr => {
  const [r, g, b] = rgbStr.slice(4, -1).split(',').map(Number)
  const max = Math.max(r, g, b)
  const min = Math.min(r, g, b)
  const l = (max + min) / ((0xff * 2) / 100)

  if (max === min) return [0, 0, l]

  const d = max - min
  const s = (d / (l > 50 ? 0xff * 2 - max - min : max + min)) * 100
  if (max === r) return [((g - b) / d + (g < b && 6)) * 60, s, l]
  return max === g
    ? [((b - r) / d + 2) * 60, s, l]
    : [((r - g) / d + 4) * 60, s, l]
}

const pathMap = {
  [`/${exercise}/${exercise}.js`]: path.join(solutionPath, `${exercise}.js`),
}

const ifNotExists = (p, fn) => {
  try {
    fs.statSync(p)
  } catch (err) {
    if (err.code !== 'ENOENT') throw err
    fn()
  }
}

ifNotExists(`./subjects/${exercise}/index.html`, () => {
  const indexPath = path.join(solutionPath, `${exercise}.html`)
  pathMap[`/${exercise}/index.html`] = indexPath
  ifNotExists(indexPath, () => {
    console.error(`missing student ${exercise}.html file`)
    process.exit(1)
  })
})

const server = http
  .createServer(({ url, method }, response) => {
    console.log(method + ' ' + url)
    if (url.endsWith('/favicon.ico')) return response.end()
    const filepath = pathMap[url] || path.join('./subjects', url)
    const ext = path.extname(filepath)
    response.setHeader('Content-Type', mediaTypes[ext.slice(1)] || 'text/plain')

    const stream = fs
      .createReadStream(filepath)
      .pipe(response)
      .once('error', err => {
        console.log(err)
        response.statusCode = 500 // handle 404 ?
        response.end('oopsie')
      })
  })
  .listen(PORT, async err => {
    let browser,
      code = 0
    try {
      err && (console.error(err.stack) || process.exit(1))
      const { setup = () => {}, tests } = await import(`./${exercise}_test.js`)
      browser = await puppeteer.launch(config)

      const [page] = await browser.pages()
      await page.goto(`http://localhost:${PORT}/${exercise}/index.html`)
      const baseContext = {
        page,
        browser,
        eq: deepStrictEqual,
        random,
        rgbToHsl,
      }
      const context = await setup(baseContext)

      browser
        .defaultBrowserContext()
        .overridePermissions(`http://localhost:${PORT}`, ['clipboard-read'])

      for (const [n, test] of tests.entries()) {
        try {
          await test({ ...baseContext, ...context })
        } catch (err) {
          console.log(`test #${n} failed:`)
          console.log(test.toString())
          throw err
        }
      }
    } catch (err) {
      code = 1
      console.log(err.stack)
    } finally {
      await (browser && browser.close())
      server.close()
      process.exit(code)
    }
  })
