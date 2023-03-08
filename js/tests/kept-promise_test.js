export const tests = []

const testZero = async _ => {
  return (await processInfo(_ => new Promise(resolve => resolve(0)))) === 'Ok!'
}

const testOddNumber = async _ => {
  return (
    (await processInfo(_ => new Promise(resolve => resolve(1)))) === 'Error!'
  )
}

const testEvenNumber = async _ => {
  return (await processInfo(_ => new Promise(resolve => resolve(4)))) === 'Ok!'
}

tests.push(testZero, testOddNumber, testEvenNumber)

Object.freeze(tests)
