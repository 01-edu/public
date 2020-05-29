export const tests = []
const t = (f) => tests.push(f)
const toMatches = (results) =>
  results
    .split(' ')
    .map((k) => (/\d/.test(k) ? '_' : k))
    .reduce((a, k) => ({ ...a, [k]: (a[k] || 0) + 1 }), {})

t(({ eq }) => eq(nasa(15), '1 2 NA 4 SA NA 7 8 NA SA 11 NA 13 14 NASA'))
t(({ eq }) => eq(toMatches(nasa(60)), { NA: 16, NASA: 4, SA: 8, _: 32 }))
t(({ eq }) => eq(toMatches(nasa(100)), { NA: 27, NASA: 6, SA: 14, _: 53 }))
t(({ eq }) => eq(toMatches(nasa(300)), { NA: 80, NASA: 20, SA: 40, _: 160 }))
t(({ eq }) => eq(nasa(900).slice(-36), 'NA 892 893 NA SA 896 NA 898 899 NASA'))

Object.freeze(tests)

/*
  “Bravery comes along as a gradual accumulation of discipline”

      ― Buzz Aldrin
*/
