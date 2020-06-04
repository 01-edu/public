export const tests = []
const t = (f) => tests.push(f)

// escapeStr
// is declared and of type string
t(() => typeof escapeStr === 'string')
// should include the character '
t(() => escapeStr.includes("'"))
// should include the character "
t(() => escapeStr.includes('"'))
// should include the character `
t(() => escapeStr.includes('`'))
// should include the character /
t(() => escapeStr.includes('/'))
// should include the character \
t(() => escapeStr.includes('\\'))

// arr
t(() => Array.isArray(arr)) // arr is declared and is an array
t(({ eq }) => eq(arr[0], 4)) // arr first element is 4
t(({ eq }) => eq(arr[1], '2')) // arr second element is "2"
t(({ eq }) => eq(arr.length, 2)) // arr length is 2

// obj
// obj is declared and of type object
t(() => obj.constructor === Object)
// obj is constant and can not be re-assigned
t(({ fail }) => obj && fail(() => (obj = 10)))
// obj.str is of type string
t(() => typeof obj.str === 'string')
// obj.num is of type number
t(() => typeof obj.num === 'number')
// obj.bool is of type boolean
t(() => typeof obj.bool === 'boolean')
// obj.undef is of type undefined
t(() => 'undef' in obj && typeof obj.undef === 'undefined')

// nested
const cantEdit = (fn) => {
  try {
    fn()
  } catch (err) {
    return true
  }
}

t(() => nested.constructor === Object)

// nested is constant and can not be re-assigned
t(({ fail }) => nested && fail(() => (nested = 10)))

t(() => nested.obj.constructor === Object)
t(() => typeof nested.obj.str === 'string')
t(() => typeof nested.obj.num === 'number')
t(() => typeof nested.obj.bool === 'boolean')

t(() => Array.isArray(nested.arr))
t(() => nested.arr[0] === 4)
t(() => nested.arr[1] === undefined)
t(() => nested.arr[2] === '2')
t(() => nested.arr.length === 3)

// nested is frozen and can not be changed
t(() => cantEdit(() => (nested.obj = 5)))
t(() => nested.obj !== 5)

// nested.obj is also frozen and can not be changed
t(() => cantEdit(() => (nested.obj.update = 5)))
t(() => nested.obj.update === undefined)

// nested.arr is not frozen and can be changed
t(() => cantEdit(() => nested.arr.push('hot stuff')))
t(() => nested.arr.length === 3)

Object.freeze(tests)
