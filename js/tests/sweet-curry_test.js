export const tests = []
const t = (f) => tests.push(f)

t(({ eq }) => eq(mult2(2)(5), 10))
t(({ eq }) => eq(mult2(3)(6), 18))
t(({ eq }) => eq(mult2(4)(7), 28))
t(({ eq }) => eq(add3(1)(2)(3), 6))
t(({ eq }) => eq(add3(4)(5)(11), 20))
t(({ eq }) => eq(add3(4)(7)(10), 21))
t(({ eq }) => eq(sub4(4)(7)(10)(30), -43))
t(({ eq }) => eq(sub4(5)(17)(-10)(3), -5))
t(({ eq }) => eq(sub4(3)(72)(-211)(99), 43))
t(({ eq }) => eq(sub4(5)(7)(10)(26), -38))

Object.freeze(tests)
