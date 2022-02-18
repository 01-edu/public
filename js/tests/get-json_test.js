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

t(async ({ eq }) => {
  // check url parsing
  let url
  fetch = async (arg) => fakeFetch({ url: (url = arg) })
  const pending = await getJSON('/test', { query: 'hello world', b: 5 })
  return eq(url, '/test?query=hello+world&b=5')
})

t(async ({ eq }) => {
  // check that it return the given value
  const data = Math.random()
  fetch = (url) => fakeFetch({ url, data })
  return eq(await getJSON('/', { q: 1 }), data)
})

t(async ({ eq }) => {
  // check that it throw an error with the correct message
  const error = `oops: ${Math.random()}`
  fetch = (url) => fakeFetch({ url, error })

  return eq(
    await getJSON('/', { q: 1 }).then(
      () => Promise.reject(Error('Should fail')),
      (err) => err.message
    ),
    error
  )
})

t(async ({ eq }) => {
  // check that it throw if the request is not ok
  fetch = (url) =>
    fakeFetch({ url, status: 500, statusText: 'Internal Server Error' })

  return eq(
    await getJSON('/', { q: 1 }).then(
      () => Promise.reject(Error('Should fail')),
      (err) => err.message
    ),
    'Internal Server Error'
  )
})

t(async ({ eq }) => {
  // if fetch fail, the error should not be handled
  const error = `oops: ${Math.random()}`
  fetch = (url) => Promise.reject(Error(error))

  return eq(
    await getJSON('/', { q: 1 }).then(
      () => Promise.reject(Error('Should fail')),
      (err) => err.message
    ),
    error
  )
})

export const setup = () => globalThis.fetch
