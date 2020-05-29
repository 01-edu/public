Array.prototype.reduce = undefined
Array.prototype.reduceRight = undefined
// /*/ // ⚡
export const tests = []
const t = (f) => tests.push(f)

const adder = (a, b) => a + b
const ifOdd = (a, b) => (b % 2 === 0 ? a + 2 : a * 2)
const concatenate = (a = '', b) => a.concat(b)
const merger = (a, b) => ({ ...a, ...b })

t(({ eq, ctx }) => eq(fold(ctx.num1, adder, 0), 39))
t(({ eq, ctx }) =>
  eq(fold(ctx.str1, concatenate, '-> '), '-> This is a simple example'),
)
t(({ eq, ctx }) => eq(fold(ctx.num1, ifOdd, 0), 6))
t(({ eq, ctx }) => eq(fold(ctx.num1, adder, 4), 43))
t(({ eq, ctx }) =>
  eq(
    fold(ctx.str2, concatenate, ''),
    'The quick brown fox jumped over the lazy dog ',
  ),
)
t(({ eq, ctx }) => eq(fold(ctx.num1, ifOdd, 10), 26))

t(({ eq, ctx }) => eq(foldRight(ctx.num1, adder, 0), 39))
t(({ eq, ctx }) =>
  eq(foldRight(ctx.str1, concatenate, '-> '), '-> examplesimple a is This '),
)
t(({ eq, ctx }) => eq(foldRight(ctx.num1, ifOdd, 0), 12))
t(({ eq, ctx }) => eq(foldRight(ctx.num1, adder, 4), 43))
t(({ eq, ctx }) =>
  eq(
    foldRight(ctx.str2, concatenate, 'This is almost understandable. '),
    'This is almost understandable. dog lazy the over jumped fox brown quick The ',
  ),
)
t(({ eq, ctx }) => eq(foldRight(ctx.num1, ifOdd, 10), 32))

t(({ eq, ctx }) => eq(reduce(ctx.num1, adder), 39))
t(({ eq, ctx }) => eq(reduce(ctx.num2, adder), 63))
t(({ eq, ctx }) =>
  eq(reduce(ctx.str1, concatenate), 'This is a simple example'),
)

t(({ eq, ctx }) =>
  eq(
    reduce(ctx.str2, concatenate),
    'The quick brown fox jumped over the lazy dog ',
  ),
)

t(({ eq, ctx }) =>
  eq(reduce(ctx.obj, merger), {
    a: 12,
    b: 6,
    c: { d: 2, e: 3 },
    f: 'hello',
  }),
)

t(({ eq, ctx }) => eq(reduceRight(ctx.num1, adder), 39))
t(({ eq, ctx }) => eq(reduceRight(ctx.num2, adder), 63))
t(({ eq, ctx }) =>
  eq(reduceRight(ctx.str1, concatenate), 'examplesimple a is This '),
)

t(({ eq, ctx }) =>
  eq(
    reduceRight(ctx.str2, concatenate),
    'dog lazy the over jumped fox brown quick The ',
  ),
)

t(({ eq, ctx }) =>
  eq(reduceRight(ctx.obj, merger), {
    f: 'hello',
    b: 6,
    c: { d: 2, e: 3 },
    a: 12,
  }),
)

Object.freeze(tests)

export const setup = () =>
  Object.fromEntries(
    Object.entries({
      num1: [3, 10, 26, 0],
      num2: [4, 24, 10, 25],
      str1: ['This ', 'is ', 'a ', 'simple ', 'example'],
      str2: 'The quick brown fox jumped over the lazy dog'
        .split(' ')
        .map((x) => (x += ' ')),
      obj: [{ a: 12 }, { b: 6, c: { d: 2, e: 3 } }, { f: 'hello' }],
    }).map(([k, v]) => [k, Object.freeze(v)]),
  )
