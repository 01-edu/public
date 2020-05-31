Promise.race = undefined
// /*/ // ⚡
export const tests = []
const t = (f) => tests.push(f)

//// RACE
// empty array should never resolve, just hang, forever...
t(async ({ eq, wait }) => {
  let notCalled = true
  race([]).then(() => (notCalled = false))
  await wait(5)
  return eq(notCalled, true)
})

// it should return the first value to resolve
t(async ({ eq, wait }) =>
  eq(await race([Promise.resolve(2), wait(5).then(() => 5)]), 2),
)

// it should not wait for every results to wait
t(async ({ eq, wait }) => {
  const start = Date.now()
  const result = await race([Promise.resolve(2), wait(50)])
  if (Date.now() - start > 5) throw Error("that's too long !")
  return eq(result, 2)
})

// it should fail if the first promise reject
t(async ({ eq, wait }) =>
  eq(
    await race([wait(5), Promise.reject(Error('oops'))]).catch(
      (err) => err.message,
    ),
    'oops',
  ),
)

// it should not fail if the first promise did not reject
t(async ({ eq, wait, ctx }) =>
  eq(
    await race([
      wait(5).then(() => Promise.reject(Error('oops'))),
      Promise.resolve(ctx),
    ]).catch((err) => err.message),
    ctx,
  ),
)

//// SOME
// empty array returns an empty array
t(async ({ eq }) => eq(await some([], 10), []))

// a count of less than 1 also returns an empty array
t(async ({ eq }) => eq(await some([1, 2, 3], 0), []))

// it should return the first value to resolve
t(async ({ eq, wait }) =>
  eq(await some([Promise.resolve(2), wait(5).then(() => 5)], 1), [2]),
)

// it should not wait for every results to wait, the order should be preserved.
t(async ({ eq, wait }) => {
  const start = Date.now()
  const result = await some([wait(1), wait(50), Promise.resolve(5)], 2)
  if (Date.now() - start > 5) throw Error("that's too long !")
  return eq(result, [undefined, 5])
})

Object.freeze(tests)

export const setup = Math.random
