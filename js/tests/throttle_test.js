export const tests = []
const t = (f) => tests.push(f)

const add = (arr, el) => arr?.push(el)
const run = (callback, callLimit, nbr) =>
  new Promise((r) => {
    let arr = []
    let inter = setInterval(() => callback(arr, 1), callLimit)
    setTimeout(() => {
      clearInterval(inter)
      r(arr.length)
    }, callLimit * nbr)
  })

t(async ({ eq }) =>
  // wait 26ms and execute 4 times every 16ms, executes with a wait time of 26
  eq(await run(throttle(add, 26), 16, 4), 2)
)

t(async ({ eq }) =>
  // wait 20ms and execute 2 times every 10ms, executes with a wait time of 26
  eq(await run(throttle(add, 20), 10, 2), 1)
)

t(async ({ eq }) =>
  // wait 16ms and execute 5 times every 26ms, will execute with out waiting
  eq(await run(throttle(add, 16), 26, 5), 4)
)

t(async ({ eq }) =>
  // it works concurrently
  eq(
    await Promise.all([
      run(throttle(add, 16), 26, 6),
      run(throttle(add, 16), 26, 6),
    ]),
    [5, 5]
  )
)

t(async ({ eq }) =>
  // tests the trailing option
  eq(await run(opThrottle(add, 26, { trailing: true }), 16, 4), 1)
)

t(async ({ eq }) =>
  // tests the leading option with wait time in the leading edge of the timeout
  eq(await run(opThrottle(add, 15, { leading: true }), 10, 10), 5)
)

t(async ({ eq }) =>
  // tests the leading option with wait time not in the leading edge of the timeout
  eq(await run(opThrottle(add, 26, { leading: true }), 16, 4), 2)
)

t(async ({ eq }) =>
  // tests without options
  eq(await run(opThrottle(add, 10), 5, 2), 0)
)

t(async ({ eq }) =>
  // tests with both options true
  eq(
    await run(opThrottle(add, 26, { trailing: true, leading: true }), 16, 4),
    2
  )
)

Object.freeze(tests)
