const add4 = '+4'
const mul2 = '*2'
// /*/ // ⚡
export const tests = []
const t = (f) => tests.push(f)

const result = (expression) =>
  expression
    .slice(2)
    .split(' ')
    .reduce((total, op) => {
      if (op === '+4') return total + 4
      if (op === '*2') return total * 2
      throw Error(`unknown op ${op}`)
    }, 1)

t(
  ({ code }) =>
    !/[5-9]/g.test(code) && code.includes('add4') && code.includes('mul2'),
)

t(({ eq }) => eq(result(findExpression(8)), 8))
t(({ eq }) => eq(result(findExpression(14)), 14))
t(({ eq }) => eq(result(findExpression(60)), 60))
t(({ eq }) => eq(result(findExpression(100)), 100))
t(({ eq }) => eq(result(findExpression(100)), 100))
t(({ eq }) => eq(result(findExpression(280)), 280))
t(({ eq }) => eq(result(findExpression(110)), 110))
t(({ eq }) => eq(result(findExpression(144)), 144))
t(({ eq }) => eq(result(findExpression(200)), 200))
t(({ eq }) => eq(result(findExpression(104)), 104))

t(({ eq }) => eq(findExpression(7), undefined))
t(({ eq }) => eq(findExpression(63), undefined))
t(({ eq }) => eq(findExpression(23), undefined))
t(({ eq }) => eq(findExpression(103), undefined))

Object.freeze(tests)
