import { once } from 'node:events'
import { spawn } from 'node:child_process'
import { mkdir, writeFile, chmod } from 'fs/promises'
import { join } from 'path'

export const tests = []
const fetch = _fetch // to redefine the real fetch

const port = 5000

export const setup = async ({ randStr }) => {
  const dir = '.'

  await mkdir(`${dir}/guests`, { recursive: true })

  const randomName = randStr()

  const createFilesIn = ({ files, dirPath }) => {
    Promise.all(
      files.map(([fileName, content]) =>
        writeFile(`${dirPath}/${fileName}`, JSON.stringify(content)),
      ),
    )
  }

  const sendRequest = async (path, options) => {
    const response = await fetch(`http://localhost:${port}${path}`, options)
    const { status } = response
    const headers = Object.fromEntries(response.headers)
    let body = ''
    try {
      body = await response.json()
    } catch (err) {
      body = err
    }
    return { status, body, headers }
  }

  const startServer = async path => {
    const server = spawn('node', [`${path}`])
    const message = await Promise.race([
      once(server.stdout, 'data'),
      Promise.race([
        once(server.stderr, 'data').then(String).then(Error),
        once(server, 'error'),
      ]).then(result => Promise.reject(result)),
    ])
    return { server, message }
  }

  return { tmpPath: dir, createFilesIn, randomName, sendRequest, startServer }
}

const isServerRunningWell = async ({ path, ctx }) => {
  const { server, message } = await ctx.startServer(path)
  server.kill()
  return message[0].toString().includes(port)
}
const isRightStatusCode = async ({ eq, ctx, randStr }) => {
  const { status } = await ctx.sendRequest(`/${ctx.randomName}`, {
    method: 'POST',
    headers: {
      'content-type': 'application/json',
    },
    body: {
      message: randStr(),
    },
  })
  server.kill()
  if (status != 201) {
    return false
  }
  return true
}
const isRightContentType = async ({ eq, ctx, randStr }) => {
  const { headers } = await ctx.sendRequest(`/Rahima_Young`, {
    method: 'POST',
    headers: {
      'content-type': 'application/json',
    },
    body: randStr(),
  })
  if (headers['content-type'] != 'application/json') {
    return false
  }
  server.kill()
  return true
}

const testServerFail = async ({ path, eq, ctx }) => {
  const { server } = await ctx.startServer(path)
  await chmod(`${ctx.tmpPath}/guests/${ctx.randomName}.json`, 0)
  const { status, body, headers } = await ctx.sendRequest(
    `/mario_${ctx.randomName}`,
    {
      method: 'GET',
    },
  )
  server.kill()
  return eq(
    {
      status: status,
      body: body,
      'content-type': headers['content-type'],
    },
    {
      status: 500,
      body: { error: 'server failed' },
      'content-type': 'application/json',
    },
  )
}

tests.push(
  isServerRunningWell,
  isRightStatusCode,
  isRightContentType,
  testServerFail,
)

Object.freeze(tests)

// const testGuestNotThere = async ({ path, eq, ctx }) => {
//   const { server } = await ctx.startServer(path)
//   const { status, body, headers } = await ctx.sendRequest('/andrea_bianchi', {
//     method: 'GET',
//   })
//   server.kill()
//   return eq(
//     {
//       status: status,
//       body: body,
//       contentType: headers['content-type'],
//     },
//     {
//       status: 404,
//       body: { error: 'guest not found' },
//       contentType: 'application/json',
//     },
//   )
// }

