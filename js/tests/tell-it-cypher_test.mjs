import * as cp from 'child_process'
import fs from 'fs/promises'
import { join, resolve, isAbsolute } from 'path'
import { tmpdir } from 'os'
import { promisify } from 'util'
const mkdir = fs.mkdir
const rmdir = fs.rmdir
const writeFile = fs.writeFile
const readFile = fs.readFile

const exec = promisify(cp.exec)

export const tests = []
const name = 'tell-it-cypher'
// maybe get the sames from an api? like https://parser.name/
const guests = [
  'Shyam Langley',
  'Austin Harwood',
  'Reem Morgan',
  'Neal Chamberlain',
  'Ryan Walters',
  'Ocean Battle',
  'Ubaid Ballard',
  'Victoria Chan',
  'Dominika Mullen',
  'Heath Denton',
  'Lilith Hamilton',
  'Aisling Bailey',
  'Maizie Love',
  'Nathanial Franco',
  'Charmaine Bernard',
  'Sohail Downes',
  'Rabia Gomez',
  'Brendan Brennan',
  'Shannen Atherton',
  'Esa Villarreal',
  'Kayla Wynn',
  'Gladys Hardy',
  'Laaibah Rogers',
  'Zishan Randolph',
  'Connor Connolly',
  'Arabella Wooten',
  'Edna Floyd',
]
const shuffle = (arr) => {
  let i = arr.length
  let j, tmp
  while (--i > 0) {
    j = Math.floor(Math.random() * (i + 1))
    tmp = arr[j]
    arr[j] = arr[i]
    arr[i] = tmp
  }
  return arr
}
const getRandomList = (names) =>
  shuffle(names).slice(0, Math.floor(Math.random() * (names.length - 10) + 10))
const getExpected = (list) =>
  list
    .map((n) =>
      n
        .split(' ')
        .reverse()
        .join(' '),
    )
    .sort()
    .map((g, i) => `${i + 1}. ${g}`)
    .join('\n')

export const setup = async () => {
  const dir = tmpdir()

  // check if already exists and rm?
  await mkdir(`${dir}/${name}`)
  const randomList = getRandomList(guests)
  const decoded = getExpected(randomList)
  const encoded = Buffer.from(decoded).toString('base64')
  await writeFile(`${dir}/${name}/encoded.txt`, encoded)
  await writeFile(`${dir}/${name}/decoded.txt`, decoded)

  return { tmpPath: `${dir}/${name}`, decoded, encoded }
}
const cypherIt = async ({ keyword, newFile, ctx, path, eq }) => {
  const scriptPath = join(resolve(), path)
  const filename = keyword === "encode" ? "decoded.txt" : "encoded.txt"
  await exec(`node ${scriptPath} ${filename} ${keyword} ${newFile ? newFile : ""}`, {
    cwd: ctx.tmpPath,
  })
  const newFileName = newFile || (keyword === "encode" ? "cypher.txt" : "clear.txt")
  const out = await readFile(`${ctx.tmpPath}/${newFileName}`, 'utf8').catch((err) =>
    err.code === 'ENOENT' ? 'output file not found' : err,
  )
  await exec(`rm ${newFileName}`, { cwd: ctx.tmpPath })
  return eq(out, keyword === "encode" ? ctx.encoded : ctx.decoded)
}

tests.push(async ({ path, eq, ctx }) => {
  return cypherIt({
    path,
    eq,
    ctx,
    keyword: 'encode',
  })
})

tests.push(async ({ path, eq, ctx }) => {
  return cypherIt({
    path,
    eq,
    ctx,
    keyword: 'decode',
  })
})

tests.push(async ({ path, eq, ctx }) => {
  return cypherIt({
    path,
    eq,
    ctx,
    keyword: 'encode',
    newFile: 'mysecret.txt',
  })
})

tests.push(async ({ path, eq, ctx }) => {
  return cypherIt({
    path,
    eq,
    ctx,
    keyword: 'decode',
    newFile: 'pandora.txt',
  })
})

Object.freeze(tests)
