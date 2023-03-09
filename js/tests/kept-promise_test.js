export const tests = []

const testZero = async () => {
  return (await processInfo(() => new Promise(resolve => resolve(0)))) === 'Ok!'
}

const testOddNumber = async () => {
  return (
    (await processInfo(() => new Promise(resolve => resolve(1)))) === 'Error!'
  )
}

const testEvenNumber = async () => {
  return (await processInfo(() => new Promise(resolve => resolve(4)))) === 'Ok!'
}

tests.push(testZero, testOddNumber, testEvenNumber)

Object.freeze(tests)
