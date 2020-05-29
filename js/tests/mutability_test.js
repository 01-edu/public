const person = {
  name: 'Rick',
  age: 77,
  country: 'US',
}
// /*/ // ⚡
export const tests = []
const t = (f) => tests.push(f)

t(() => typeof samePerson === 'object')
t(() => typeof person === 'object')
t(() => typeof clone1 === 'object')
t(() => typeof clone2 === 'object')
t(({ eq }) => eq(clone1, clone2)) // equal
t(() => clone1 !== clone2) // but different !
t(() => person === samePerson) // same value
t(() => person.name === 'Rick')
t(() => person.age === 78)
t(() => person.country === 'FR')
t(() => clone1.country === 'US')
t(() => clone2.age === 77)

Object.freeze(tests)
