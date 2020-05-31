export const tests = []
const t = (f) => tests.push(f)
const fakeFetch = async ({ data, error, ...opts } = {}) => ({
  ok: !opts.status,
  type: 'basic',
  status: 200,
  statusText: 'OK',
  json: async () => ({ data, error }),
  text: async () => JSON.stringify({ data, error }),
  ...opts,
})

// check url parsing
t(async ({ eq }) => {
  let url
  fetch = async (arg) => fakeFetch({ url: (url = arg) })
  const pending = getJSON('/test', { query: 'hello world', b: 5 })
  return eq(url, '/test?query=hello%20world&b=5')
})

// check that it return the given value
t(async ({ eq }) => {
  const data = Math.random()
  fetch = (url) => fakeFetch({ url, data })
  return eq(await getJSON('/', { q: 1 }), data)
})

// check that it throw an error with the correct message
t(async ({ eq }) => {
  const error = `oops: ${Math.random()}`
  fetch = (url) => fakeFetch({ url, error })

  return eq(
    await getJSON('/', { q: 1 }).then(
      () => Promise.reject(Error('Should fail')),
      (err) => err.message,
    ),
    error,
  )
})

// check that it throw if the request is not ok
t(async ({ eq }) => {
  fetch = (url) =>
    fakeFetch({ url, status: 500, statusText: 'Internal Server Error' })

  return eq(
    await getJSON('/', { q: 1 }).then(
      () => Promise.reject(Error('Should fail')),
      (err) => err.message,
    ),
    'Internal Server Error',
  )
})

// if fetch fail, the error should not be handled
t(async ({ eq }) => {
  const error = `oops: ${Math.random()}`
  fetch = (url) => Promise.reject(Error(error))

  return eq(
    await getJSON('/', { q: 1 }).then(
      () => Promise.reject(Error('Should fail')),
      (err) => err.message,
    ),
    error,
  )
})

export const setup = () => globalThis.fetch
