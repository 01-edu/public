import * as cp from 'child_process'
import { mkdir, writeFile, readFile } from 'fs/promises'
import { join, isAbsolute } from 'path'
import { tmpdir } from 'os'
import { promisify } from 'util'

const exec = promisify(cp.exec)
export const tests = []

export const setup = async ({ path }) => {
  const dir = `${tmpdir()}/tell-it-cypher`

  await mkdir(dir)

  const run = async cmd => {
    const [filename, keyword, newFile] = cmd.split(' ')
    const filePath = isAbsolute(filename) ? filename : join(dir, filename)
    const { stdout } = await exec(
      `node ${path} ${filePath} ${keyword} ${newFile || ''}`,
    )
    const newFileName =
      newFile || (keyword === 'encode' ? 'cypher.txt' : 'clear.txt')
    const fileContent = await readFile(newFileName, 'utf8')

    return { data: fileContent }
  }

  return { tmpPath: dir, run }
}

tests.push(async ({ path, eq, ctx }) => {
  const vips = `1. Langley Shyam
  2. Harwood Austin
  3. Morgan Reem
  4. Chamberlain Neal
  5. Walters Ryan`
  const fileName = `${ctx.tmpPath}/vip.txt`
  await writeFile(fileName, vips)

  const { data } = await ctx.run(`${fileName} encode`)

  return eq(
    data,
    'MS4gTGFuZ2xleSBTaHlhbQogIDIuIEhhcndvb2QgQXVzdGluCiAgMy4gTW9yZ2FuIFJlZW0KICA0LiBDaGFtYmVybGFpbiBOZWFsCiAgNS4gV2FsdGVycyBSeWFu',
  )
})

tests.push(async ({ randStr, eq, ctx }) => {
  const vips = `1. Wynn Kayla
  2. Hardy Gladys
  3. Rogers Laaibah
  4. Randolph Zishan
  5. Connolly Connor`
  const fileName = `${ctx.tmpPath}/vip-${randStr()}.txt`
  await writeFile(fileName, vips)

  const { data } = await ctx.run(`${fileName} encode mysecret.txt`)

  return eq(
    data,
    'MS4gV3lubiBLYXlsYQogIDIuIEhhcmR5IEdsYWR5cwogIDMuIFJvZ2VycyBMYWFpYmFoCiAgNC4gUmFuZG9scGggWmlzaGFuCiAgNS4gQ29ubm9sbHkgQ29ubm9y',
  )
})

tests.push(async ({ randStr, eq, ctx }) => {
  const vipsEncoded =
    'MS4gVmlsbGFycmVhbCBFc2EKICAyLiBXeW5uIEtheWxhCiAgMy4gSGFyZHkgR2xhZHlzCiAgNC4gUm9nZXJzIExhYWliYWgKICA1LiBSYW5kb2xwaCBaaXNoYW4='
  const fileName = `${ctx.tmpPath}/vip-encoded-${randStr()}.txt`
  await writeFile(fileName, vipsEncoded)

  const { data } = await ctx.run(`${fileName} decode`)
  return eq(
    data,
    `1. Villarreal Esa
  2. Wynn Kayla
  3. Hardy Gladys
  4. Rogers Laaibah
  5. Randolph Zishan`,
  )
})

tests.push(async ({ randStr, eq, ctx }) => {
  const vipsEncoded =
    'MS4gQmVybmFyZCBDaGFybWFpbmUKICAyLiBEb3duZXMgU29oYWlsCiAgMy4gR29tZXogUmFiaWEKICA0LiBCcmVubmFuIEJyZW5kYW4KICA1LiBBdGhlcnRvbiBTaGFubmVu'
  const fileName = `${ctx.tmpPath}/vip-encoded-${randStr()}.txt`
  await writeFile(fileName, vipsEncoded)

  const { data } = await ctx.run(`${fileName} decode pandora.txt`)
  return eq(
    data,
    `1. Bernard Charmaine
  2. Downes Sohail
  3. Gomez Rabia
  4. Brennan Brendan
  5. Atherton Shannen`,
  )
})

Object.freeze(tests)
