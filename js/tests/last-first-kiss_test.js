export const tests = []
const t = (f) => tests.push(f)

t(({ eq }) => eq(last([2, 42]), 42)) // Oh
t(({ eq }) => eq(last(['pouet', 4, true]), true))
t(({ eq }) => eq(last([last]), last)) // I wanna be last, yeah
t(({ eq }) => eq(last('salut'), 't')) // Baby let me be your last
t(({ eq }) => eq(last([]), undefined)) // Your last first kiss
t(({ eq }) => eq(first([2, 42]), 2))
t(({ eq }) => eq(first(['pouet', 4, true]), 'pouet'))
t(({ eq }) => eq(first([first]), first))
t(({ eq }) => eq(first('salut'), 's'))
t(({ eq }) => eq(first([]), undefined))
t(({ eq }) => eq(kiss([1, 2, 3, 4, 5, 6]), [6, 1]))
t(({ eq }) => eq(kiss([eq, kiss, first]), [first, eq]))
t(({ eq }) => eq(kiss([]), [undefined, undefined]))

Object.freeze(tests)
