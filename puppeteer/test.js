import http from 'http'
import fs from 'fs'
import path from 'path'
import { deepStrictEqual } from 'assert'
import puppeteer from 'puppeteer-core'

const exercise = process.argv[2]
if (!exercise) throw Error(`usage: node test EXERCISE_NAME`)
const PORT = 9898
const config = {
  headless: false,
  executablePath: process.env.CHROME_PATH || '/usr/bin/google-chrome',
}

const mediaTypes = {
  jpg: 'image/jpeg',
  png: 'image/png',
  html: 'text/html',
  css: 'text/css',
  js: 'application/javascript',
  json: 'application/json',
}

const server = http.createServer(({ url, method }, response) => {
  console.log(method + ' ' + url)
  const filepath = url.endsWith(`${exercise}/${exercise}.js`)
    ? path.join('.', url.slice(exercise.length + 1))
    : path.join('./subjects', url)
  const ext = path.extname(filepath)
  response.setHeader('Content-Type', mediaTypes[ext.slice(1)] || 'text/plain')

  const stream = fs.createReadStream(filepath)
    .pipe(response)
    .once('error', err => {
      console.log(err)
      response.statusCode = 500 // handle 404 ?
      response.end('oopsie')
    })
}).listen(PORT, async (err) => {
  let browser, code = 0
  try {
    err && (console.error(err.stack) || process.exit(1))
    const { setup = () => {}, tests } = await import(`./${exercise}_test.js`)
    browser = await puppeteer.launch(config)
    const [page] = await browser.pages()
    await page.goto(`http://localhost:${PORT}/${exercise}/index.html`)
    const context = await setup({ page })
    for (const [n, test] of tests.entries()) {
      try {
        await test({ page, eq: deepStrictEqual, ...context })
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
