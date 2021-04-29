import { join as joinPath, dirname, extname } from 'path'
import { fileURLToPath } from 'url'
import { deepStrictEqual } from 'assert'
import * as fs from 'fs'
const { readFile, writeFile } = fs.promises

global.window = global
global.fetch = url => {
  // this is a fake implementation of fetch for the tester
  // -> refer to https://devdocs.io/javascript/global_objects/fetch
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

const testNode = async ({ name }) => {
  const path = `${solutionPath}/${name}.mjs`
  return {
    path,
    url: joinPath(root, `${name}_test.mjs`),
    code: await read(path, 'student solution'),
  }
}

const runInlineTests = async ({ json, name }) => {
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
    logs.forEach((args) => console.info(...args))
    console.error(...args)
    process.exit(1)
  }

  const solution = await loadAndSanitizeSolution(name)
  for (const { description, code } of JSON.parse(json)) {
    logs.length = 0
    const [provided, tests] = code.includes('// Your code')
      ? code.split('// Your code')
      : ['', code]

    const fullCode = `
${provided ? '// Provided setup' : ''}
${provided.trim()}

// Your code
${solution.code.trim()}

// The tests
${tests.trim()}`.trim()

    try {
      eval(fullCode)
      console.info(`${description}:`, 'PASS')
    } catch (err) {
      console.info(`${description}:`, 'FAIL')
      console.info(`\n======= Code ======= \n${fullCode}`)
      console.info('\n======= Error ======')
      die(' ->', err.message)
    }
  }
}

const loadAndSanitizeSolution = async name => {
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
    fatal(`Unable to execute ${name} solution, error:\n${stackFmt(err, url)}`),
  )

  const randStr = (n = 7) => Math.random().toString(36).slice(2, n)
  const between = (min, max) => {
    max || (max = min, min = 0)
    return Math.floor(Math.random() * (max - min) + min)
  }
  const upperFirst = (str) => str[0].toUpperCase() + str.slice(1)

  const tools = { eq, fail, wait, code, path, randStr, between, upperFirst }
  tools.ctx = (await (setup && setup(tools))) || {}
  for (const [i, t] of tests.entries()) {
    try {
      if (!(await t(tools))) {
        throw Error('Test failed')
      }
    } catch (err) {
      console.log(`test #${i+1} failed:\n${t.toString()}\n\nError:`)
      fatal(stackFmt(err, url))
    }
  }
  console.log(`${name} passed (${tests.length} tests)`)
}

const main = async () => {
  const { test, mode } = await any([
    readTest(joinPath(root, `${name}.json`)),
    readTest(joinPath(root, `${name}_test.js`)),
    readTest(joinPath(root, `${name}_test.mjs`)),
  ]).catch(ifNoEnt((err) => fatal(`Missing test for ${name}`)))

  if (mode === "node") return runTests(await testNode({ test, name }))
  if (mode === "inline") return runInlineTests({ json: test, name })

  const { rawCode, code, path } = await loadAndSanitizeSolution(name)
  const parts = test.split("// /*/ // ⚡")
  const [inject, testCode] = parts.length < 2 ? ["", test] : parts
  const combined = `${inject.trim()}\n${rawCode
    .replace(inject.trim(), "")
    .trim()}\n${testCode.trim()}\n`

  const b64 = Buffer.from(combined).toString("base64")
  const url = `data:text/javascript;base64,${b64}`
  return runTests({ path, code, url })
}

main().catch(err => fatal(err.stack))
