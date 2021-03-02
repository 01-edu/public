import fs from 'fs/promises'
import { tmpdir } from 'os'
import { join, resolve } from 'path'
const readFile = fs.readFile
const writeFile = fs.writeFile
const mkdir = fs.mkdir
import { promisify } from 'util'
import * as cp from 'child_process'

const exec = promisify(cp.exec)

export const tests = []
const name = 'verydisco-reverso'
export const setup = async () => {
  const dir = tmpdir()
  // check if already exists and rm?
  await mkdir(`${dir}/${name}`)

  return { tmpPath: `${dir}/${name}` }
}
const hasContentAndExpect = async ({ content, expect, path, eq, ctx }) => {
  const scriptPath = join(resolve(), path)

  await writeFile(`${ctx.tmpPath}/${name}.txt`, content, 'utf8')
  const { stdout } = await exec(`node ${scriptPath} "${name}.txt"`, { cwd: ctx.tmpPath })
  await exec(`rm ${name}.txt`, { cwd: ctx.tmpPath })

  return eq(stdout.trim(), expect)
}

tests.push(async ({ path, eq, ctx }) =>
  hasContentAndExpect({
    path,
    eq,
    ctx,
    content: `deNo si omeawes`,
    expect: 'Node is awesome',
  }),
)

// tests.push(async ({ path, eq, ctx }) => {
//   await hasContentAndExpect({
//     path,
//     eq,
//     ctx,
//     content: // newfilename,
//     expect: // reverso
//   })
// })

Object.freeze(tests)
