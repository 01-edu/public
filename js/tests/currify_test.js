export const tests = []
const t = (f) => tests.push(f)

const mult2 = (el1, el2) => el1 * el2
const add3 = (el1, el2, el3) => el1 + el2 + el3
const sub4 = (el1, el2, el3, el4) => el1 - el2 - el3 - el4

t(({ eq }) => eq(currify(mult2)(2)(5), 10))
t(({ eq }) => eq(currify(mult2)(3)(6), 18))
t(({ eq }) => eq(currify(mult2)(4)(7), 28))
t(({ eq }) => eq(currify(add3)(1)(2)(3), 6))
t(({ eq }) => eq(currify(add3)(4)(5)(11), 20))
t(({ eq }) => eq(currify(add3)(4)(7)(10), 21))
t(({ eq }) => eq(currify(sub4)(4)(7)(10)(30), -43))
t(({ eq }) => eq(currify(sub4)(5)(17)(-10)(3), -5))
t(({ eq }) => eq(currify(sub4)(3)(72)(-211)(99), 43))
t(({ eq }) => eq(currify(sub4)(5)(7)(10)(26), -38))

Object.freeze(tests)
