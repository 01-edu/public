import { once } from 'node:events'
import { spawn } from 'node:child_process'
import { mkdir, writeFile, chmod } from 'fs/promises'
import fs from 'fs'
import { join } from 'path'

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

  return { tmpPath: dir, createFilesIn, sendRequest, startServer }
}

const isServerRunningWell = async ({ path, ctx }) => {
  const { server, message } = await ctx.startServer(path)
  server.kill()
  return message[0].toString().includes(port)
}

const testGoodRequests = async ({ path, eq, ctx }) => {
  let isTestOk = true
  const expectedBody = {
    answer: 'yes',
    drink: 'alcohol',
    food: 'bats',
  }
  const dirName = 'guests'
  const dirPath = join(ctx.tmpPath, dirName)
  const { server } = await ctx.startServer(path)

  {
    const { status, body, headers } = await ctx.sendRequest(`/Ana_Riber`, {
      method: 'POST',
      headers: {
        authorization:
          'Basic ' +
          Buffer.from('Caleb_Squires:abracadabra').toString('base64'),
        body: JSON.stringify(expectedBody),
      },
    })

    fs.access(`${dirPath}/Ana_Riber.json`, fs.F_OK, err => {
      if (err) {
        console.error(err)
        isTestOk = false
      }
    })
    if (
      status != 200 ||
      headers['content-type'] != 'application/json' ||
      !eq({ body: body }, { body: expectedBody })
    ) {
      isTestOk = false
    }
  }
  {
    const { status, body, headers } = await ctx.sendRequest(`/Rob_Frie`, {
      method: 'POST',
      headers: {
        authorization:
          'Basic ' +
          Buffer.from('Tyrique_Dalton:abracadabra').toString('base64'),
        body: JSON.stringify(expectedBody),
      },
    })

    fs.access(`${dirPath}/Rob_Frie.json`, fs.F_OK, err => {
      if (err) {
        console.error(err)
        isTestOk = false
      }
    })
    if (
      status != 200 ||
      headers['content-type'] != 'application/json' ||
      !eq({ body: body }, { body: expectedBody })
    ) {
      isTestOk = false
    }
  }
  {
    const { status, body, headers } = await ctx.sendRequest(`/George_Harl`, {
      method: 'POST',
      headers: {
        authorization:
          'Basic ' + Buffer.from('Rahima_Young:abracadabra').toString('base64'),
        body: JSON.stringify(expectedBody),
      },
    })

    fs.access(`${dirPath}/George_Harl.json`, fs.F_OK, err => {
      if (err) {
        console.error(err)
        isTestOk = false
      }
    })
    if (
      status != 200 ||
      headers['content-type'] != 'application/json' ||
      !eq({ body: body }, { body: expectedBody })
    ) {
      isTestOk = false
    }
  }
  server.kill()
  return isTestOk
}

const testUnauthorizedRequests = async ({ path, ctx }) => {
  let isTestOk = true
  const { server } = await ctx.startServer(path)
  {
    const { status } = await ctx.sendRequest(`/Rahima_Young`, {
      method: 'POST',
    })
    if (status != 401) {
      isTestOk = false
    }
  }
  {
    const { status } = await ctx.sendRequest(``, {
      method: 'POST',
    })
    if (status != 401) {
      isTestOk = false
    }
  }
  {
    const { status } = await ctx.sendRequest(
      `/Rahima_Young:wrongpass`,
      {
        method: 'POST',
      },
    )
    if (status != 401) {
      isTestOk = false
    }
  }
  {
    const { status } = await ctx.sendRequest(
      `/Anonymus:abracadabra`,
      {
        method: 'POST',
      },
    )
    if (status != 401) {
      isTestOk = false
    }
  }
  server.kill()
  return isTestOk
}

tests.push(isServerRunningWell, testGoodRequests, testUnauthorizedRequests)

Object.freeze(tests)
