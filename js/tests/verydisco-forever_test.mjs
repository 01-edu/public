import fs from 'fs/promises'
import { join, resolve } from 'path'
import { tmpdir } from 'os'
const readFile = fs.readFile
const mkdir = fs.mkdir
import { promisify } from 'util'
import * as cp from 'child_process'

const exec = promisify(cp.exec)

const name = 'verydisco-forever'
export const tests = []
export const setup = async () => {
  const dir = tmpdir()
  // check if already exists and rm?
  await mkdir(`${dir}/${name}`)

  return { tmpPath: `${dir}/${name}` }
}
const withArgAndExpect = async ({ arg, expect, path, eq, ctx }) => {
  const scriptPath = join(resolve(), path)
  await exec(`node ${scriptPath} "${arg}"`, { cwd: ctx.tmpPath }) // ${name}.mjs
  const out = await readFile(`${ctx.tmpPath}/${name}.txt`, 'utf8').catch((err) =>
    err.code === 'ENOENT' ? 'output file not found' : err,
  )
  await exec(`rm ${name}.txt`, { cwd: ctx.tmpPath })
  return eq(out, expect)
}

tests.push(async ({ path, eq, ctx }) =>
  withArgAndExpect({ path, eq, ctx, arg: 'discovery', expect: 'verydisco' }),
)

tests.push(async ({ path, eq, ctx }) =>
  withArgAndExpect({ path, eq, ctx, arg: 'kiss cool', expect: 'sski olco' }),
)

tests.push(async ({ path, eq, ctx }) =>
  withArgAndExpect({
    path,
    eq,
    ctx,
    arg: 'Node is awesome',
    expect: 'deNo si omeawes',
  }),
)

Object.freeze(tests)
