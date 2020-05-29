export const tests = []
const t = (f) => tests.push(f)

t(() => typeof concatStr === 'function', 'Should be a function')
t(() => concatStr.length === 2, 'Should takes 2 arguments')
t(() => concatStr('a', 'b') === 'ab')
t(() => concatStr('yolo', 'swag') === 'yoloswag')

// handle non strings correctly
t(() => concatStr(1, 2) === '12')
t(() => concatStr(concatStr, concatStr) === String(concatStr).repeat(2))

Object.freeze(tests)
