Promise.race = undefined
// /*/ // ⚡
export const tests = []
const t = (f) => tests.push(f)

//// RACE
t(async ({ eq, wait }) => {
  // empty array should never resolve, just hang, forever...
  let notCalled = true
  race([]).then(() => (notCalled = false))
  await wait(5)
  return eq(notCalled, true)
})

t(async ({ eq, wait }) =>
  // it should return the first value to resolve
  eq(await race([Promise.resolve(2), wait(5).then(() => 5)]), 2)
)

t(async ({ eq, wait }) => {
  // it should not wait for every results to wait
  const start = Date.now()
  const result = await race([Promise.resolve(2), wait(50)])
  if (Date.now() - start > 5) throw Error("that's too long !")
  return eq(result, 2)
})

t(async ({ eq, wait }) =>
  // it should fail if the first promise reject
  eq(
    await race([wait(5), Promise.reject(Error('oops'))]).catch(
      (err) => err.message
    ),
    'oops'
  )
)

t(async ({ eq, wait, ctx }) =>
  // it should not fail if the first promise did not reject
  eq(
    await race([
      wait(5).then(() => Promise.reject(Error('oops'))),
      Promise.resolve(ctx),
    ]).catch((err) => err.message),
    ctx
  )
)

//// SOME
t(async ({ eq }) =>
  // empty array returns an empty array
  eq(await some([], 10), [])
)

t(async ({ eq }) =>
  // a count of less than 1 also returns an empty array
  eq(await some([1, 2, 3], 0), [])
)

t(async ({ eq, wait }) =>
  // it should return the first value to resolve
  eq(await some([Promise.resolve(2), wait(5).then(() => 5)], 1), [2])
)

t(async ({ eq, wait }) => {
  // it should not wait for every results to wait, the order should be preserved.
  const start = Date.now()
  const result = await some([wait(1), wait(50), Promise.resolve(5)], 2)
  if (Date.now() - start > 5) throw Error("that's too long !")
  return eq(result, [undefined, 5])
})

Object.freeze(tests)

export const setup = Math.random
