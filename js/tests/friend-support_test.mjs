import { once } from 'node:events'
import * as cp from 'node:child_process'
import fs from 'node:fs/promises'

export const tests = []
const fetch = _fetch // to redefine the real fetch

const port = 5000

// tell-me-who_test setup -- see how to have same files in testing directories
// TO BE IMPLEMENTED

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

tests.push(async ({ eq }) => {
  const { status, body, headers } = await sendRequest('/Zoe_Sierra', {
    method: 'GET',
  })
  let cont = await fs.readFile('./guests/Zoe_Sierra.json', {
    encoding: 'utf-8',
  })
  return eq(
    {
      status: status,
      body: body,
      contentType: headers['content-type'],
    },
    {
      status: 200,
      body: JSON.parse(cont),
      contentType: 'application/json',
    },
  )
})

tests.push(async ({ eq }) => {
	const { status, body, headers } = await sendRequest('/Zoe_Sierra1', {
	  method: 'GET',
	})
	console.log(body.error)
	return eq(
	  {
		status: status,
		body: body.error.trim(),
		contentType: headers['content-type'],
	  },
	  {
		status: 404,
		body: "guest not found",
		contentType: 'application/json',
	  },
	)
  })

// this test must always run to shutdown server
/* currently breaking if one test fails - it might not be a problem
since the tests are run in a separate Docker image...
*/
tests.push(({ ctx }) => {
  const { server } = ctx
  server.kill('SIGQUIT')
  return true
})

Object.freeze(tests)
