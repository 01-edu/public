export const tests = []
const t = (f) => tests.push(f)

const add = (arr, el) => arr.push(el)

// async setInterval is not precise enough so I made this synchronous one
const setIntervalSync = (fn, delay, limit = 0) => new Promise(s => {
  let run = true
  let count = 1
  const start = Date.now()
  const loop = () => {
    const tick = Date.now()
    const elapsed = tick - start
    if (elapsed > count * delay) {
      fn()
      count++
    }
    elapsed < limit
      ? setTimeout(loop)
      : s()
  }
  setTimeout(loop)
})

// it uses the array to better test the leading and trailing edge of the time limit
// so if the leading edge is true it will execute the callback
// if the trailing edge is true it will execute the callback before returning the array
const run = async (callback, { delay, count }) => {
  const arr = []
  await setIntervalSync(() => callback(arr, 1), delay, count * delay + 5)
  return arr.length
}

t(async ({ eq }) =>
  // test with debounce wait limit inferior to wait time call (how much time we wait to the function be called again)
  // it works concurrently
  eq(
    await Promise.all([
      run(debounce(add, 50), { delay: 100, count: 5 }),
      run(debounce(add, 20), { delay: 50, count: 10 }),
    ]),
    [4, 9]
  )
)
t(async ({ eq }) =>
  // testing with wait limit superior to wait time call
  // execution on the trailing edge, after wait limit has elapsed
  eq(await run(debounce(add, 100), { delay: 50, count: 5 }), 0)
)

t(async ({ eq }) =>
  // it works concurrently
  eq(
    await Promise.all([
      run(opDebounce(add, 40), { delay: 20, count: 5 }),
      run(opDebounce(add, 40), { delay: 20, count: 2 }),
    ]),
    [0, 0]
  )
)

t(async ({ eq }) =>
  // leading edge as true
  // it works concurrently
  eq(
    await Promise.all([
      run(opDebounce(add, 200, { leading: true }), { delay: 70, count: 3 }),
      run(opDebounce(add, 100, { leading: true }), { delay: 140, count: 3 }),
    ]),
    [1, 3]
  )
)

Object.freeze(tests)
