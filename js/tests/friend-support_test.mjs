import { once } from 'node:events'
import * as cp from 'node:child_process'
import { mkdir, writeFile, chmod } from 'fs/promises'
import { tmpdir } from 'node:os'
import { join } from 'path'

export const tests = []
const fetch = _fetch // to redefine the real fetch

const port = 5000

export const setup = async ({ randStr }) => {
  const dir = '.'

  await mkdir(`${dir}/guests`, { recursive: true })

  const randLastName = randStr()

  const createFilesIn = ({ files, dirPath }) =>
    Promise.all(
      files.map(([fileName, content]) =>
        writeFile(`${dirPath}/${fileName}`, JSON.stringify(content)),
      ),
    )

  const sendRequest = async (path, options) => {
    const response = await fetch(`http://localhost:${port}${path}`, options)
    const { status, statusText, ok } = response
    const headers = Object.fromEntries(response.headers)
    let body = ''
    try {
      body = await response.json()
    } catch (err) {
      body = err
    }
    return { status, body, headers, randLastName }
  }

  return { tmpPath: dir, createFilesIn, sendRequest }
}

tests.push(async ({ path, ctx }) => {
  ctx.server = cp.spawn('node', [`${path}`])
  const message = await Promise.race([
    once(ctx.server.stdout, 'data'),
    Promise.race([
      once(ctx.server.stderr, 'data').then(String).then(Error),
      once(ctx.server, 'error'),
    ]).then(x => Promise.reject(x)),
  ])
  return message[0].toString().includes(port)
})

tests.push(async ({ eq, ctx, randStr }) => {
  // test for one guest
  const expBody = { message: 'ciao' }
  const files = [[`mario_${ctx.randLastName}.json`, expBody]]
  const dirName = `guests`
  const dirPath = join(ctx.tmpPath, dirName)
  await ctx.createFilesIn({ dirPath, files })

  const { status, body, headers } = await ctx.sendRequest(
    `/mario_${ctx.randLastName}`,
    {
      method: 'GET',
    },
  )
  return eq(
    {
      status: status,
      body: body,
      contentType: headers['content-type'],
    },
    {
      status: 200,
      body: expBody,
      contentType: 'application/json',
    },
  )
})

tests.push(async ({ eq, ctx, randStr }) => {
  // test server failed
  // change permission for existing file
  await chmod(`${ctx.tmpPath}/guests/mario_${ctx.randLastName}.json`, 0)
  const { status, body, headers } = await ctx.sendRequest(
    `/mario_${ctx.randLastName}`,
    {
      method: 'GET',
    },
  )
  return eq(
    {
      status: status,
      body: body,
      contentType: headers['content-type'],
    },
    {
      status: 500,
      body: { error: 'server failed' },
      contentType: 'application/json',
    },
  )
})

tests.push(async ({ eq, ctx }) => {
  // test guest not there
  const { status, body, headers } = await ctx.sendRequest('/andrea_bianchi', {
    method: 'GET',
  })
  return eq(
    {
      status: status,
      body: body,
      contentType: headers['content-type'],
    },
    {
      status: 404,
      body: { error: 'guest not found' },
      contentType: 'application/json',
    },
  )
})

Object.freeze(tests)
