export const tests = []
const t = f => tests.push(f)

t(({ eq, ctx }) => eq(arrToSet([1, ctx, ctx, 3, 3]), new Set([1, ctx, 3])))
t(({ eq, ctx }) => eq(arrToStr([1, ctx, ctx, 3]), `${1}${ctx}${ctx}${3}`))
t(({ eq, ctx }) => eq(setToArr(new Set([1, ctx, 3])), [1, ctx, 3]))
t(({ eq, ctx }) => eq(setToStr(new Set([1, ctx, 3])), `${1}${ctx}${3}`))
t(({ eq, ctx }) => eq(strToArr(`${1}${ctx}${3}`), ['1', ...ctx, '3']))

t(({ eq, ctx }) =>
  eq(strToSet(`12${ctx[0]}23`), new Set(['1', '2', ctx[0], '3'])),
)

t(({ eq, ctx }) =>
  eq(
    mapToObj(
      new Map([
        [ctx, ctx],
        ['a', 2],
      ]),
    ),
    { [ctx]: ctx, a: 2 },
  ),
)

t(({ eq, ctx }) => eq(objToArr({ a: 1, b: ctx }), [1, ctx]))

t(({ eq, ctx }) =>
  eq(
    objToMap({ [ctx]: ctx, a: 2 }),
    new Map([
      [ctx, ctx],
      ['a', 2],
    ]),
  ),
)

t(({ eq, ctx }) =>
  eq(arrToObj([1, ctx, 'pouet']), { '0': 1, '1': ctx, '2': 'pouet' }),
)

t(({ eq, ctx }) =>
  eq(strToObj(`2${ctx[0]}2`), { '0': '2', '1': ctx[0], '2': '2' }),
)

t(({ eq }) => eq(superTypeOf(new Map()), 'Map'))
t(({ eq }) => eq(superTypeOf(new Set()), 'Set'))
t(({ eq }) => eq(superTypeOf({}), 'Object'))
t(({ eq }) => eq(superTypeOf(''), 'String'))
t(({ eq }) => eq(superTypeOf(666), 'Number'))
t(({ eq }) => eq(superTypeOf(NaN), 'Number'))
t(({ eq }) => eq(superTypeOf([]), 'Array'))
t(({ eq }) => eq(superTypeOf(null), 'null'))
t(({ eq }) => eq(superTypeOf(undefined), 'undefined'))
t(({ eq }) => eq(superTypeOf(superTypeOf), 'Function'))

export const setup = () => Math.random().toString(36).slice(2)

Object.freeze(tests)
