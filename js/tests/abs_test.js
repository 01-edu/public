Math.abs = undefined
// /*/ // ⚡
export const tests = []
const t = (f) => tests.push(f)

t(() => isPositive(3))
t(() => isPositive(1998790))
t(() => !isPositive(-1))
t(() => !isPositive(-0.7))
t(() => !isPositive(-787823))
t(() => !isPositive(0))

t(({ eq }) => eq(abs(0), 0))
t(({ eq }) => eq(abs(-1), 1))
t(({ eq }) => eq(abs(-13.2), 13.2))
t(({ eq }) => eq(abs(132), 132))

Object.freeze(tests)
