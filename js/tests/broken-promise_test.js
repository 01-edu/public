export const tests = []

const testOne = async () => {
  return (await doubleIt(() => new Promise(resolve => resolve(1)))) === 2
}

const testFifteen = async () => {
  return (await doubleIt(() => new Promise(resolve => resolve(15)))) === 30
}

const testRejectNoNumbers = async () => {
  return (
    (await doubleIt(() => new Promise((resolve, reject) => reject("No numbers")))) === "Error: No numbers"
  )
}

const testRejectNumberless = async () => {
  return (
    (await doubleIt(() => new Promise((resolve, reject) => reject("Numberless")))) === "Error: Numberless"
  )
}

tests.push(testOne, testFifteen, testRejectNoNumbers, testRejectNumberless)

Object.freeze(tests)
