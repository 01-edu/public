export const tests = []
const isConst = (name) => {
  try {
    eval(`${name} = 'm'`)
  return false
  } catch (err) {
    return true
  }
}
const t = (f) => tests.push(f)

// str is declared and of type string
t(() => typeof str === 'string')

// num is declared and of type number
t(() => typeof num === 'number')

// bool is declared and of type boolean
t(() => typeof bool === 'boolean')

// undef is declared and of type undefined
t(() => undef === undefined)

// check if all variables are const
t(() => ['str', 'num', 'bool', 'undef']
  .every(isConst))

Object.freeze(tests)
