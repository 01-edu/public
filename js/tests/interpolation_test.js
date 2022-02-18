export const tests = []
const t = (f) => tests.push(f)
const run = async ({ step, start, end, duration, waitTime = 15 }) => {
  let arr = []
  const round = (nbr) => Math.round(nbr * 100) / 100
  const callback = (el) =>
    arr.push(Array.isArray(el) ? el.map(round) : round(el))
  interpolation({ step, start, end, callback, duration })
  await new Promise((s) => setTimeout(s, waitTime))
  return arr
}

t(async ({ eq }) => {
  // testing duration time, forbid loops
  const { length } = await run({ step: 5, start: 0, end: 4, duration: 50 })
  return eq(length, 1)
})

t(async ({ eq }) => {
  // testing duration time stamp
  const { length } = await run({
    step: 2,
    start: 0,
    end: 4,
    duration: 10,
    waitTime: 0,
  })
  return eq(length, 0)
})

t(async ({ eq }) => {
  // testing the amount of times the callback was called
  const { length } = await run({ step: 5, start: 0, end: 1, duration: 10 })
  return eq(length, 5)
})

t(async ({ eq }) =>
  // testing the interpolation points
  eq(await run({ step: 5, start: 0, end: 1, duration: 10 }), [
    [0, 2],
    [0.2, 4],
    [0.4, 6],
    [0.6, 8],
    [0.8, 10],
  ])
)

t(async ({ eq }) =>
  // testing with different values
  eq(await run({ step: 3, start: 1, end: 2, duration: 10 }), [
    [1, 3.33],
    [1.33, 6.67],
    [1.67, 10],
  ])
)

t(async ({ eq }) =>
  // testing with `duration` inferior to `step`
  eq(await run({ step: 10, start: 2, end: 6, duration: 4 }), [
    [2, 0.4],
    [2.4, 0.8],
    [2.8, 1.2],
    [3.2, 1.6],
    [3.6, 2],
    [4, 2.4],
    [4.4, 2.8],
    [4.8, 3.2],
    [5.2, 3.6],
    [5.6, 4],
  ])
)

t(async ({ eq }) =>
  // testing with `start` superior to `end`
  // inverted straight line
  eq(await run({ step: 5, start: 6, end: 2, duration: 6, waitTime: 10 }), [
    [6, 1.2],
    [5.2, 2.4],
    [4.4, 3.6],
    [3.6, 4.8],
    [2.8, 6],
  ])
)

Object.freeze(tests)
