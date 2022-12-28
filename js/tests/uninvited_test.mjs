import { once } from 'node:events'
import { spawn } from 'node:child_process'
import { mkdir, writeFile, chmod } from 'fs/promises'
import { join } from 'path'
import fs from 'node:fs/promises'

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
        writeFile(`${dirPath}/${fileName}`, JSON.stringify(content), {
          flag: 'wx',
        }),
      ),
    ).catch(reason => console.log(reason))

    return true
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

const testServerRunning = async ({ path, ctx }) => {
  const { server, message } = await ctx.startServer(path)
  server.kill()
  return message[0].toString().includes(port)
}

const testRightStatusCode = async ({ path, ctx, randStr }) => {
  const { server } = await ctx.startServer(path)
  const { status } = await ctx.sendRequest(`/${ctx.randomName}`, {
    method: 'POST',
    headers: {
      'content-type': 'application/json',
    },
    body: randStr(),
  })
  server.kill()

  if (status != 201) return false
  return true
}

const testRightContentType = async ({ path, ctx, randStr }) => {
  const { server } = await ctx.startServer(path)
  const { headers } = await ctx.sendRequest(`/${ctx.randomName}`, {
    method: 'POST',
    headers: {
      'content-type': 'application/json',
    },
    body: randStr(),
  })
  server.kill()
  if (headers['content-type'] != 'application/json') return false
  return true
}

const testServerFail = async ({ path, eq, ctx, randStr }) => {
  const { server } = await ctx.startServer(path)
  await chmod(`${ctx.tmpPath}/guests`, 0).catch(reason => console.log(reason))
  const { status, body, headers } = await ctx.sendRequest(
    `/${ctx.randomName}`,
    {
      method: 'POST',
      body: randStr(),
    },
  )
  await chmod(`${ctx.tmpPath}/guests`, 0o777)
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

const testFileCreated = async ({ path, ctx, randStr }) => {
  const { server } = await ctx.startServer(path)
  const randomName = randStr()
  await ctx.sendRequest(`/${randomName}`, {
    body: randStr(),
    method: 'POST',
  })
  const dirName = 'guests'
  const dirPath = join(ctx.tmpPath, dirName)
  let accessWorked = true
  server.kill()
  await fs
    .access(`${dirPath}/${randomName}.json`, fs.constants.F_OK)
    .catch(reason => {
      accessWorked = false
      console.log(reason)
    })
  return accessWorked
}

const testFileContent = async ({ path, ctx, randStr }) => {
  const { server } = await ctx.startServer(path)
  const randomName = randStr()
  const body = randStr()
  await ctx.sendRequest(`/${randomName}`, {
    body: body,
    method: 'POST',
  })
  const dirName = 'guests'
  const dirPath = join(ctx.tmpPath, dirName)
  server.kill()
  let content = ''
  await fs
    .readFile(`./${dirPath}/${randomName}.json`, 'utf8', (err, data) => {
      if (err) {
        console.error(err)
        return 'error when reading file'
      }
      return data
    })
    .then(data => {
      if (data === 'error when reading file') return
      content = data
    })
  return body === content
}

const testBodyOnSuccess = async ({ path, ctx, eq, randStr }) => {
  const { server } = await ctx.startServer(path)
  const randomBody = { message: randStr() }
  const { body } = await ctx.sendRequest(`/${ctx.randomName}`, {
    method: 'POST',
    headers: {
      'content-type': 'application/json',
    },
    body: JSON.stringify(randomBody),
  })
  server.kill()
  return eq(
    {
      body: body,
    },
    {
      body: randomBody,
    },
  )
}
tests.push(
  testServerRunning,
  testRightStatusCode,
  testRightContentType,
  testBodyOnSuccess,
  testFileCreated,
  testFileContent,
  testServerFail,
)

Object.freeze(tests)
