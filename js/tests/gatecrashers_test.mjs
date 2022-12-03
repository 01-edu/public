import { once } from 'node:events'
import * as cp from 'node:child_process'
import { mkdir, writeFile, chmod } from 'fs/promises'

export const tests = []
const fetch = _fetch // to redefine the real fetch

const port = 5000

export const setup = async ({}) => {
  const dir = '.'

  await mkdir(`${dir}/guests`, { recursive: true })

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
    return { status, body, headers }
  }

  return { tmpPath: dir, createFilesIn, sendRequest }
}

// Test the server is running and writes the port in stdout
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

tests.push(async ({ eq, ctx }) => {
  {
    const { status, body, headers } = await ctx.sendRequest(`/Rahima_Young`, {
      method: 'POST',
      headers: {
        authorization:
          'Basic ' +
          Buffer.from('Caleb_Squires:abracadabra').toString('base64'),
      },
    })
    if (status != 200 || headers['content-type'] != 'application/json') {
      return false
    }
  }
  {
    const { status, body, headers } = await ctx.sendRequest(`/Rahima_Young`, {
      method: 'POST',
      headers: {
        authorization:
          'Basic ' +
          Buffer.from('Tyrique_Dalton:abracadabra').toString('base64'),
      },
    })
    if (status != 200 || headers['content-type'] != 'application/json') {
      return false
    }
  }
  {
    const { status, body, headers } = await ctx.sendRequest(`/Rahima_Young`, {
      method: 'POST',
      headers: {
        authorization:
          'Basic ' + Buffer.from('Rahima_Young:abracadabra').toString('base64'),
      },
    })
    if (status != 200 || headers['content-type'] != 'application/json') {
      return false
    }
  }
  return true
})

// Unauthorized requests
tests.push(async ({ eq, ctx }) => {
  {
    const { status, body, headers } = await ctx.sendRequest(`/Rahima_Young`, {
      method: 'POST',
    })
    if (status != 401) {
      return false
    }
  }

  {
    const { status, body, headers } = await ctx.sendRequest(``, {
      method: 'POST',
    })
    if (status != 401) {
      return false
    }
  }

  {
    const { status, body, headers } = await ctx.sendRequest(
      `/Rahima_Young:wrongpass`,
      {
        method: 'POST',
      },
    )
    if (status != 401) {
      return false
    }
  }

  {
    const { status, body, headers } = await ctx.sendRequest(
      `/Anonymus:abracadabra`,
      {
        method: 'POST',
      },
    )
    if (status != 401) {
      return false
    }
  }
  return true
})

Object.freeze(tests)
