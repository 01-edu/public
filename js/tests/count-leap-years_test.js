export const tests = []
const t = (f) => tests.push(f)

t(() => countLeapYears(new Date('0001-12-01')) === 0)
t(() => countLeapYears(new Date('1664-08-09')) === 403)
t(() => countLeapYears(new Date('2020-01-01')) === 489)
t(() => countLeapYears(new Date('2048-12-08')) === 496)

Object.freeze(tests)
