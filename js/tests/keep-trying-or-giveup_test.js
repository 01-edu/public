export const tests = []
const t = (f) => tests.push(f)

const fail = (q) =>
  q.then(
    (v) => Promise.reject('should fail'),
    (e) => e.message,
  )

t(async ({ eq, ctx }) => eq(await retry(0, ctx.failNTimes(0))(ctx.r), [ctx.r]))
t(async ({ eq, ctx }) => eq(await retry(3, ctx.failNTimes(3))(ctx.r), [ctx.r]))
t(async ({ eq, ctx }) =>
  eq(await retry(10, ctx.failNTimes(5))(ctx.r, ctx.r), [ctx.r, ctx.r]),
)
t(async ({ eq, ctx }) =>
  eq(await fail(retry(3, ctx.failNTimes(9))(ctx.r)), `x:${ctx.r}`),
)

t(async ({ eq, ctx }) => eq(await timeout(2, ctx.delayed(0))(ctx.r), [ctx.r]))
t(async ({ eq, ctx }) =>
  eq(await timeout(2, ctx.delayed(0))(ctx.r, ctx.r), [ctx.r, ctx.r]),
)
t(async ({ eq, ctx }) =>
  eq(await fail(timeout(2, ctx.delayed(4))(ctx.r)), 'timeout'),
)

Object.freeze(tests)

export const setup = () => ({
  r: Math.random().toString(36).slice(2),
  failNTimes: (n) => async (...v) =>
    --n < 0 ? v : Promise.reject(Error(`x:${v}`)),
  delayed: (delay) => (...v) => new Promise((s) => setTimeout(s, delay, v)),
})
