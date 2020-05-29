Array.prototype.reverse = undefined
// /*/ // ⚡
export const tests = []
const t = (f) => tests.push(f)

t(({ eq }) => eq(reverse([1, 2, 3]), [3, 2, 1]))
t(({ eq, ctx }) => eq(reverse([1, eq, 3, ctx]), [ctx, 3, eq, 1]))
t(({ eq }) => eq(reverse('pouet'), 'teuop'))
t(({ eq }) => eq(reverse("salut c'est cool"), "looc tse'c tulas"))

Object.freeze(tests)

export const setup = Math.random
