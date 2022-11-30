import { once } from 'node:events'
import { spawn } from 'node:child_process'
import { mkdir, writeFile, chmod } from 'fs/promises'
import { join } from 'path'

const fetch = _fetch // to redefine the real fetch

const port = 5000

export const setup = async ({ randStr }) => {
  const folder = '.'
  await mkdir(`${folder}/guests`, { recursive: true })
  const randLastName = randStr()

  const createFilesIn = ({ files, dirPath: folderPath }) =>
    Promise.all(
      files.map(([fileName, content]) =>
        writeFile(`${folderPath}/${fileName}`, JSON.stringify(content)),
      ),
    )

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
    return { status, body, headers, randLastName }
  }

  return { temporaryPath: dir, createFilesIn, sendRequest }
}
const isServerThereTest = async ({ path, ctx: { server } }) => {
  server = spawn('node', [`${path}`])
  const message = await Promise.race([
    once(server.stdout, 'data'),
    Promise.race([
      once(server.stderr, 'data').then(String).then(Error),
      once(server, 'error'),
    ]).then(error => Promise.reject(error)),
  ])
  return message[0].toString().includes(port)
}
// const oneGuestTest = async ({ eq, ctx }) => {
//   const expBody = { message: 'ciao' }
//   const files = [[`mario_${ctx.randLastName}.json`, expBody]]
//   const dirName = `guests`
//   const dirPath = join(ctx.tmpPath, dirName)
//   await ctx.createFilesIn({ dirPath, files })

//   const { status, body, headers } = await ctx.sendRequest(
//     `/mario_${ctx.randLastName}`,
//     {
//       method: 'GET',
//     },
//   )
//   return eq(
//     {
//       status: status,
//       body: body,
//       contentType: headers['content-type'],
//     },
//     {
//       status: 200,
//       body: expBody,
//       contentType: 'application/json',
//     },
//   )
// }
const serverFailedTest = async ({ eq, ctx }) => {
  // change permission for existing file
  await chmod(`${ctx.tmpPath}/guests/mario_${ctx.randLastName}.json`, 0)
  const { status, body, headers } = await ctx.sendRequest(
    `/mario_${ctx.randLastName}`,
    {
      method: 'POST',
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
}
// const testGuestNotThere = async ({ eq, ctx }) => {
//   const { status, body, headers } = await ctx.sendRequest('/andrea_bianchi', {
//     method: 'GET',
//   })
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
export const tests = [
  isServerThereTest,
  //   oneGuestTest,
  serverFailedTest,
  //   testGuestNotThere,
]

Object.freeze(tests)
