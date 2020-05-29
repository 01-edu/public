Math.imul = undefined
// /*/ // ⚡
export const tests = []
const t = (f) => tests.push(f)

// * is not allowed for this exercise
t(({ code }) => !code.includes('*'))

t(() => multiply(34, 78) === 2652)
t(() => multiply(123, 0) === 0)
t(() => multiply(0, -230) === 0)
t(() => multiply(0, 0) === 0)
t(() => multiply(123, -22) === -2706)
t(() => multiply(-22, 123) === -2706)
t(() => multiply(-22, -123) === 2706)

// / is not allowed for this exercise
t(({ code }) => !code.includes('/'))

t(() => divide(34, 78) === 0)
t(() => divide(78, 34) === 2)
t(() => divide(123, 22) === 5)
t(() => divide(123, -22) === -5)
t(() => divide(-123, 22) === -5)
t(() => divide(-123, -22) === 5)

// % is not allowed for this exercise
t(({ code }) => !code.includes('%'))

t(() => modulo(34, 78) === 34)
t(() => modulo(78, 34) === 10)
t(() => modulo(123, 22) === 13)
t(() => modulo(123, -22) === 13)
t(() => modulo(-123, 22) === -13)
t(() => modulo(-123, -22) === -13)

Object.freeze(tests)
