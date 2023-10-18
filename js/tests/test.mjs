import puppeteer from 'puppeteer'
import { join as joinPath, dirname, extname } from 'path'
import { readFile, writeFile } from 'fs/promises'
import { deepStrictEqual } from 'assert'
import { fileURLToPath } from 'url'
import { tmpdir } from 'os'
import http from 'http'
import fs from 'fs'

global.window = global
global._fetch = fetch
global.fetch = url => {
  // this is a fake implementation of fetch for the tester
  const accessBody = async () => { throw Error('body unavailable') }
  return {
    ok: false,
    type: 'basic',
    status: 500,
    statusText: 'Internal Server Error',
    json: accessBody,
    text: accessBody,
  }
}

const wait = delay => new Promise(s => setTimeout(s, delay))
const fail = fn => { try { fn() } catch (err) { return true } }
const upperFirst = (str) => str[0].toUpperCase() + str.slice(1)
const randStr = (n = 7) => Math.random().toString(36).slice(2, n)
const between = (min, max) => {
  max || (max = min, min = 0)
  return Math.floor(Math.random() * (max - min) + min)
}

const props = [String,Array]
  .flatMap(({ prototype }) =>
    Object.getOwnPropertyNames(prototype)
      .map(key => ({ key, value: prototype[key], src: prototype })))
  .filter(p => typeof p.value === 'function')

const eq = (a, b) => {
  const changed = []
  for (const p of props) { !p.src[p.key] && (changed[changed.length] = p) }
  for (const p of changed) { p.src[p.key] = p.value }
  deepStrictEqual(a, b)
  for (const p of changed) { p.src[p.key] = undefined }
  return true
}

const [solutionPath, name] = process.argv.slice(2)

const tools = { eq, fail, wait, randStr, between, upperFirst }
const fatal = (...args) => {
  console.error(...args)
  process.exit(1)
}

solutionPath || fatal('missing solution-path, usage:\nnode test solution-path exercise-name')
name || fatal('missing exercise, usage:\nnode test solution-path exercise-name')

const ifNoEnt = fn => err => {
  if (err.code !== 'ENOENT') throw err
  fn(err)
}

const root = dirname(fileURLToPath(import.meta.url))
const read = (filename, description) =>
  readFile(filename, 'utf8').catch(
    ifNoEnt(() => fatal(`Missing ${description} for ${name}`)),
  )

const modes = { '.js': 'function', '.mjs': 'node', '.json': 'inline' }
const readTest = filename => readFile(filename, 'utf8')
  .then(test => ({ test, mode: modes[extname(filename)] }))

const stackFmt = (err, url) => {
  for (const p of props) { p.src[p.key] = p.value }
  if (err instanceof Error) return err.stack.split(url).join(`${name}.js`)
  throw Error(`Unexpected type thrown: ${typeof err}. usage: throw Error('my message')`)  
}

const any = arr =>
  new Promise(async (s, f) => {
    let firstError
    const setError = err => firstError || (firstError = err)
    await Promise.all(arr.map(p => p.then(s, setError)))
    f(firstError)
  })

const testNode = async () => {
  const path = `${solutionPath}/${name}.mjs`
  return {
    path,
    url: joinPath(root, `${name}_test.mjs`),
    code: await read(path, 'student solution'),
  }
}

const runInlineTests = async ({ json }) => {
  const restore = new Set()
  const equal = deepStrictEqual
  const saveArguments = (src, key) => {
    const savedArgs = []
    const fn = src[key]
    src[key] = (...args) => {
      savedArgs.push(args)
      return fn(...args)
    }

    restore.add(() => (src[key] = fn))

    return savedArgs
  }

  const logs = []
  console.log = (...args) => logs.push(args)
  const die = (...args) => {
    logs.forEach((logArgs) => console.info(...logArgs))
    fatal(...args)
  }

  const solution = await loadAndSanitizeSolution()
  for (const { description, code } of JSON.parse(json)) {
    logs.length = 0
    const [provided, tests] = code.includes('// Your code')
      ? code.split('// Your code')
      : ['', code]

    const fullCode = `
${provided ? '// Provided setup' : ''}
${provided.trim()}

// Your code
${solution.code.trim()};

// The tests
${tests.trim()};`.trim()

    try {
      eval(fullCode)
      console.info(`${description}:`, 'PASS')
    } catch (err) {
      console.info(`${description}:`, 'FAIL')
      console.info('\n======= Error ======')
      console.info(' ->', err.message, '\n')
      console.info('\n======= Code =======')
      die(fullCode)
    }
  }
}

const loadAndSanitizeSolution = async () => {
  const path = `${solutionPath}/${name}.js`
  const rawCode = await read(path, "student solution")

  // this is a very crude and basic removal of comments
  // since checking code is only use to prevent cheating
  // it's not that important if it doesn't work 100% of the time.
  const code = rawCode.replace(/\/\*[\s\S]*?\*\/|\/\/.*/g, "").trim()
  if (code.includes("import")) fatal("import keyword not allowed")
  return { code, rawCode, path }
}

const runTests = async ({ url, path, code }) => {
  const { setup, tests } = await import(url).catch(err =>
    fatal(`Unable to execute ${name}, error:\n${stackFmt(err, url)}`),
  )

  Object.assign(tools, { code, path })
  tools.ctx = (await (setup && setup(tools))) || {}
  const isDOM = name.endsWith('-dom')
  if (isDOM) {
    Object.assign(tools, await prepareForDOM({ code }))
  }
  let timeout
  for (const [i, t] of tests.entries()) {
    try {
      const waitWithTimeout = Promise.race([
        t(tools),
        new Promise((s, f) => {
          timeout = setTimeout(f, 60000, Error('Time limit reached (1min)'))
        }),
      ])
      if (!(await waitWithTimeout) && !isDOM) {
        throw Error('Test failed')
      }
    } catch (err) {
      console.info(`test #${i+1} failed:\n${t.toString()}\n`)
      fatal(stackFmt(err, url))
    } finally {
      clearTimeout(timeout)
    }
  }
  console.info(`${name} passed (${tests.length} tests)`)
}

// add puppeteer tests as JS language:
const PORT = 9898
const config = {
  args: [
    '--no-sandbox',
    '--disable-setuid-sandbox',

    // This will write shared memory files into /tmp instead of /dev/shm,
    // because Docker’s default for /dev/shm is 64MB
    '--disable-dev-shm-usage',
  ],
  headless: !process.env.DEBUG_PUPPETTEER,
}

// LEGACY random, use between instead (only used by dom exercise, to be replaced)
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

const prepareForDOM = ({ code }, server) => new Promise((s, f) => (server = http
  .createServer(({ url, method }, response) => {
    console.info(method + ' ' + url)
    // Loading either the `index.html` or the js code (student solution)
    response.setHeader('Content-Type', 'text/html')
    return response.end(`<script type="module">${code}</script>`)
  }))
  .listen(PORT, async listenErr => {
    if (listenErr) return f(listenErr)
    try {
      const browser = await puppeteer.launch(config)
      const [page] = await browser.pages()
      await page.goto(`http://localhost:${PORT}/index.html`)
      deepStrictEqual.$ = async (selector, props) => {
        const keys = Object.keys(props)
        const extractProps = (node, props) => {
          const fromProps = (a, b) => Object.fromEntries(Object.keys(b).map(k => [
            k,
            typeof b[k] === 'object' ? fromProps(a[k], b[k]) : a[k],
          ]))
          return fromProps(node, props)
        }
        const domProps = await page.$eval(selector, extractProps, props)
        return deepStrictEqual(props, domProps)
      }

      deepStrictEqual.css = async (selector, props) => {
        const cssProps = await page.evaluate((selector, props) => {
          const styles = Object.fromEntries([...document.styleSheets]
            .flatMap(({ cssRules }) => [...cssRules].map(r => [r.selectorText, r.style])))

          if (!styles[selector]) {
            throw Error(`css ${selector} did not match any declarations`)
          }

          return Object.fromEntries(Object.keys(props).map(k => [k, styles[selector][k]]))
        }, selector, props)

        return deepStrictEqual(props, cssProps)
      }

      browser
        .defaultBrowserContext()
        .overridePermissions(`http://localhost:${PORT}`, ['clipboard-read'])

      s({ page, browser, random, rgbToHsl, eq: deepStrictEqual, server })
    } catch (err) {
      f(err)
    }
  }))

const main = async () => {
  const { test, mode } = await any([
    readTest(joinPath(root, `${name}.json`)),
    readTest(joinPath(root, `${name}_test.js`)),
    readTest(joinPath(root, `${name}_test.mjs`)),
  ]).catch(ifNoEnt((err) => fatal(`Missing test for ${name}`)))

  if (mode === "node") return runTests(await testNode())
  if (mode === "inline") return runInlineTests({ json: test })

  const { rawCode, code, path } = await loadAndSanitizeSolution()
  const parts = test.split("// /*/ // ⚡")
  const [inject, testCode] = parts.length < 2 ? ["", test] : parts
  const combined = `${inject.trim()}\n${rawCode
    .replace(inject.trim(), "")
    .trim()}\n;${testCode.trim()}\n`

  const url = `${tmpdir()}/${name}.mjs`
  await writeFile(url, combined)
  return runTests({ path, code, url })
}

main().then(
  () => process.exit(0),
  err => fatal(err?.stack || Error('').stack),
)
