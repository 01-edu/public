export const tests = []
const t = (f) => tests.push(f)

t(() => circular.constructor === Object)
t(() => circular.circular === circular)
t(() => circular.circular.circular === circular)
t(() => circular.circular.circular.circular === circular)
t(() => circular.circular.circular.circular.circular === circular)

/*
  To understand recursion,
  one must first understand recursion.
*/

Object.freeze(tests)
