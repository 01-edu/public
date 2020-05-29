export const tests = []
const t = (f) => tests.push(f)

// it should work with an empty array
t(async ({ eq }) => eq(await series([]), []))

// it should work with synchronous values
t(async ({ eq }) => eq(await series([() => 1, () => true]), [1, true]))

// it should work with pending promises
t(async ({ eq }) =>
  eq(await series([() => Promise.resolve(1), () => Promise.resolve(true)]), [
    1,
    true,
  ]),
)

// it should work with mixed values
t(async ({ eq }) =>
  eq(await series([() => Promise.resolve(1), () => true]), [1, true]),
)

// ensure callbacks are run one at a time
t(async ({ eq, ctx, wait }) => {
  const delay = Math.round(ctx * 4 + 1)
  const start = Date.now()
  let pending = false
  const run = async () => {
    if (pending)
      throw Error(
        'attempts to run function concurently instead of one at a time',
      )
    pending = true
    await wait(delay)
    pending = false
    return ctx
  }
  const callbacks = Object.freeze([run, run, run, run])
  const result = eq(await series(callbacks), [ctx, ctx, ctx, ctx])
  if (!result) throw Error('invalid return value')
  const delta = Date.now() - start
  const min = delay * callbacks.length
  if (delta >= min) return true
  throw Error(`promises should take ~${min}ms but was ${delta}ms`)
})

// it should fail if one of the promises reject
t(async ({ eq }) =>
  eq(
    await series([
      async () => 1,
      async () => {
        throw Error('oops')
      },
    ]).catch((err) => err.message),
    'oops',
  ),
)

Object.freeze(tests)

export const setup = Math.random
