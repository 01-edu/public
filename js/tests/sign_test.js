Math.sign = undefined
// /*/ // ⚡
export const tests = []
const t = (f) => tests.push(f)

t(() => typeof sign === 'function')
t(() => sign.length === 1)
t(() => sign !== Math.sign)
t(() => sign(-2) === -1)
t(() => sign(10) === 1)
t(() => sign(0) === 0)
t(() => sign(132) === 1)

t(() => typeof sameSign === 'function')
t(() => sameSign.length === 2)
t(() => sameSign(-2, -1))
t(() => sameSign(0, 0))
t(() => sameSign(12, 3232))
t(() => !sameSign(1, -1))
t(() => !sameSign(-231, 1))
t(() => !sameSign(-231, 0))
t(() => !sameSign(0, 231))
t(() => !sameSign(231, -233))

Object.freeze(tests)
