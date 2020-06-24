Promise.all = undefined
// /*/ // ⚡
export const tests = []
const t = (f) => tests.push(f)

t(async ({ eq }) =>
  // it should work with an empty object
  eq(await all({}), {})
)

t(async ({ eq }) =>
  // it should work with synchronous values
  eq(await all({ a: 1, b: true }), { a: 1, b: true })
)

t(async ({ eq }) =>
  // it should work with pending promises
  eq(
    await all({
      a: Promise.resolve(1),
      b: Promise.resolve(true),
    }),
    { a: 1, b: true }
  )
)

t(async ({ eq, ctx }) =>
  // it should work with pending promises and synchronous values
  eq(
    await all(
      Object.freeze({
        a: Promise.resolve(ctx).then((v) => v + 1),
        b: ctx,
      })
    ),
    { a: ctx + 1, b: ctx }
  )
)

t(async ({ eq }) =>
  // it should fail if one of the promises reject
  eq(
    await all({
      a: Promise.resolve(1),
      b: Promise.reject(Error('oops')),
    }).catch((err) => err.message),
    'oops'
  )
)

Object.freeze(tests)

export const setup = Math.random
