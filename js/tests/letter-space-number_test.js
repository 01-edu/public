export const tests = []
const t = (f) => tests.push(f)

t(({ eq }) =>
  eq(letterSpaceNumber('He is 8 or 9 years old, not 10.'), ['s 8', 'r 9']),
)
t(({ eq }) => eq(letterSpaceNumber('I like 7up.'), []))
t(({ eq }) => eq(letterSpaceNumber("It's 20 past 3"), ['t 3']))
t(({ eq }) => eq(letterSpaceNumber('example 1, example 2'), ['e 1', 'e 2']))
t(({ eq }) => eq(letterSpaceNumber(''), []))
t(({ eq }) => eq(letterSpaceNumber('Definitely 9.'), ['y 9']))

Object.freeze(tests)
