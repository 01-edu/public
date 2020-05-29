export const tests = []
const t = (f) => tests.push(f)

t(({ eq }) => vowels.test('a') && !vowels.test('c'))
t(({ eq }) => eq(vowelDots('something'), 'so.me.thi.ng'))
t(({ eq }) => eq(vowelDots(''), ''))
t(({ eq }) => eq(vowelDots('rhythm'), 'rhythm'))
t(({ eq }) => eq(vowelDots('Algorithm'), 'A.lgo.ri.thm'))

Object.freeze(tests)
