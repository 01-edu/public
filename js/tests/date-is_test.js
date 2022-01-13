export const tests = []
const t = (f) => tests.push(f)

// is the date a valid date?
const invalid = (callback, ary = 1) => {
  for (const s of [
    `new Date('')`,
    `new Date(NaN)`,
    `''`,
    `'2013-01-01'`,
    `NaN`,
  ]) {
    if (callback(...Array(ary).fill(eval(s)))) {
      throw Error(`${callback.name}(${s}) should be false`)
    }
  }
}

// isValid
t(() => !invalid(isValid))
t(() => isValid(new Date()))
t(() => isValid(Date.now()))
t(() => isValid(new Date('December 17, 1995 03:24:00')))
t(() => isValid(new Date(1488370835081)))
t(() => isValid(new Date('1995-12-17T03:24:00')))
t(() => isValid(new Date('1995-12-17T03:24:00').getTime()))

// isAfter
t(() => !invalid(isAfter, 2))
t(() => !isAfter(new Date('1992-01-01'), new Date('1992-01-02')))
t(() => !isAfter(new Date('2157-11-07'), new Date('2183-04-16')))
t(() => isAfter(new Date(2321, 11, 21), new Date(Date.now())))
t(() => isAfter(123123, 526))

// isBefore
t(() => !invalid(isBefore, 2))
t(() => !isBefore(new Date(2321, 11, 21), new Date(Date.now())))
t(() => !isBefore(123123, 526))
t(() => isBefore(new Date('1992-01-01'), new Date('1992-01-02')))
t(() => isBefore(new Date('2157-11-07'), new Date('2183-04-16')))

// isFuture
t(() => !invalid(isFuture))
t(() => !isFuture(new Date('1992-01-01')))
t(() => !isFuture(new Date(Date.now() - 1)))
t(() => isFuture(new Date(2077, 11, 31)))
t(() => isFuture(new Date(Date.now() + 1)))

// isPast
t(() => !invalid(isPast))
t(() => !isPast(new Date(Date.now() + 1)))
t(() => !isPast(new Date(2077, 11, 31)))
t(() => isPast(new Date(1442, 11, 31)))
t(() => isPast(new Date(Date.now() - 1)))

Object.freeze(tests)
