export const tests = []

const testProperties = () => {
  return (
    Array.isArray(veterinary.animalKnowledge) &&
    typeof veterinary.canTreat === 'function' &&
    typeof veterinary.respondClient === 'function'
  )
}

const testCanTreat = () => {
  veterinary.animalKnowledge.push('dog', 'cat', 'elephant')
  return (
    veterinary.canTreat('dog') &&
    veterinary.canTreat('cat') &&
    veterinary.canTreat('elephant') &&
    !veterinary.canTreat('') &&
    !veterinary.canTreat('goldfish')
  )
}

const testRespondClient = () => {
  veterinary.animalKnowledge.push('parrot')
  const vetResponse1 = veterinary.respondClient('Jack', 'parrot')
  const vetResponse2 = veterinary.respondClient('Matias', 'cobra')

  return (
    vetResponse1.includes('Jack') &&
    vetResponse1.toLowerCase().includes('yes') &&
    vetResponse2.includes('Matias') &&
    vetResponse2.toLowerCase().includes('no')
  )
}

tests.push(testProperties, testCanTreat, testRespondClient)

Object.freeze(tests)
