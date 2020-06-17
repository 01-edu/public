export const tests = []
const t = (f) => tests.push(f)

const sub32 = (el) => el - 32
const mult5 = (el) => el * 5
const div9 = (el) => el / 9
const roundDown = (el) => Math.floor(el)

const square = (nbr) => nbr * nbr
const add2 = (el) => el + 2
const mult2 = (el) => el * 2

const addAll = (...el) =>
  el.length === 1 ? el[0] : el[0] + addAll(...el.slice(1))

export const setup = () => {
  const farenheitToCelsius = flow([sub32, mult5, div9, roundDown])
  const add2Mult2Square = flow([add2, mult2, square])
  const addAllThenConvertToCelsius = flow([addAll, farenheitToCelsius])

  return { farenheitToCelsius, add2Mult2Square, addAllThenConvertToCelsius }
}

//Farenheit to Celsius

t(({ eq, ctx }) => eq(ctx.farenheitToCelsius(32), 0))
t(({ eq, ctx }) => eq(ctx.farenheitToCelsius(0), -18))
t(({ eq, ctx }) => eq(ctx.farenheitToCelsius(40), 4))
t(({ eq, ctx }) => eq(ctx.farenheitToCelsius(50), 10))
t(({ eq, ctx }) => eq(ctx.farenheitToCelsius(60), 15))
t(({ eq, ctx }) => eq(ctx.farenheitToCelsius(100), 37))

// add2Mult2Square

t(({ eq, ctx }) => eq(ctx.add2Mult2Square(32), 4624))
t(({ eq, ctx }) => eq(ctx.add2Mult2Square(0), 16))
t(({ eq, ctx }) => eq(ctx.add2Mult2Square(40), 7056))
t(({ eq, ctx }) => eq(ctx.add2Mult2Square(50), 10816))
t(({ eq, ctx }) => eq(ctx.add2Mult2Square(60), 15376))
t(({ eq, ctx }) => eq(ctx.add2Mult2Square(-100), 38416))

// addAllThenConvertToCelsius

t(({ eq, ctx }) => eq(ctx.addAllThenConvertToCelsius(20, 5, 6, 1), 0))
t(({ eq, ctx }) => eq(ctx.addAllThenConvertToCelsius(-10, -10, 20), -18))
t(({ eq, ctx }) => eq(ctx.addAllThenConvertToCelsius(10, 5, 5, 10, 5, 5), 4))
t(({ eq, ctx }) => eq(ctx.addAllThenConvertToCelsius(25, 5, 20), 10))
t(({ eq, ctx }) => eq(ctx.addAllThenConvertToCelsius(30, 30), 15))
t(({ eq, ctx }) => eq(ctx.addAllThenConvertToCelsius(99, 1), 37))

Object.freeze(tests)
