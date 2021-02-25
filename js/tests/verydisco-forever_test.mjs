import fs from 'fs/promises'
const readFile = fs.readFile
import { promisify } from 'util'
import * as cp from 'child_process'

const exec = promisify(cp.exec)

export const tests = []
const name = 'verydisco-forever'
const withArgAndExpect = async ({ arg, expect, path, eq }) => {
  const cwd = `${path
    .split('/')
    .slice(0, -1)
    .join('/')}/`
  await exec(`node ${name}.mjs "${arg}"`, { cwd })
  const out = await readFile(`${cwd}${name}.txt`, 'utf8').catch((err) =>
    err.code === 'ENOENT' ? 'output file not found' : err,
  )
  await exec(`rm ${name}.txt`, { cwd })
  return eq(out, expect)
}

tests.push(async ({ path, eq }) =>
  withArgAndExpect({ path, eq, arg: 'discovery', expect: 'verydisco' }),
)

tests.push(async ({ path, eq }) =>
  withArgAndExpect({ path, eq, arg: 'kiss cool', expect: 'sski olco' }),
)

tests.push(async ({ path, eq }) =>
  withArgAndExpect({
    path,
    eq,
    arg: 'Node is awesome',
    expect: 'deNo si omeawes',
  }),
)

Object.freeze(tests)
