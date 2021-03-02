import * as cp from 'child_process'
import fs from 'fs/promises'
import { join, resolve, isAbsolute } from 'path'
import { tmpdir } from 'os'
import { promisify } from 'util'
const mkdir = fs.mkdir
const rmdir = fs.rmdir
const writeFile = fs.writeFile

const exec = promisify(cp.exec)

export const tests = []
export const setup = async () => {
  const randomFilesNumber = Math.floor(Math.random() * (30 - 1) + 1)
  const dir = tmpdir()

  // check if already exists and rm?
  await mkdir(`${dir}/random`)
  for (let i = 0; i < randomFilesNumber; i++) {
    await writeFile(`${dir}/random/${i}.txt`, '', 'utf8')
  }

  return { randomFilesNumber, tmpPath: dir }
}
const countDirLength = async ({ dirLength, arg, ctx, path, eq }) => {
  const scriptPath = join(resolve(), path)
  const { stdout } = await exec(`node ${scriptPath} ${arg}`, {
    cwd: arg ? `${ctx.tmpPath}` : `${ctx.tmpPath}/random`,
  })

  // await rmdir(`${ctx.tmpPath}/random`, { recursive: true })
  return eq(Number(stdout.trim()), dirLength)
}

tests.push(async ({ path, eq, ctx }) => {
  // will execute the script with "random" as an argument
  // `random` folder has a random file number
  return countDirLength({
    path,
    eq,
    ctx,
    arg: 'random',
    dirLength: ctx.randomFilesNumber,
  })
})

tests.push(async ({ path, eq, ctx }) => {
  // will execute the script without argument
  // in the `random` folder
  return countDirLength({
    path,
    eq,
    ctx,
    arg: '',
    dirLength: ctx.randomFilesNumber,
  })
})

tests.push(async ({ path, eq, ctx }) => {
  // will execute the script with `random` folder's absolute path as argument
  return countDirLength({
    path,
    eq,
    ctx,
    arg: `${ctx.tmpPath}/random`,
    dirLength: ctx.randomFilesNumber,
  })
})

Object.freeze(tests)
