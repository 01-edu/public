import * as cp from 'child_process'
import { promisify } from 'util'

const exec = promisify(cp.exec)
export const tests = []

const testGetSecrets = async ({ randStr, path, eq }) => {
  const possibleSecrets = ['my_password', randStr()]
  const results = []

  for (let secret of possibleSecrets) {
    await exec(`echo "SECRET=${secret}" > .env`)
    const { stdout } = await exec(`node ${path}`)
    results.push(eq(stdout.trim(), secret))
  }

  try {
    await exec('rm .env')
  } finally {
    return results.every(el => el)
  }
}

const testEmpty = async ({ path, eq }) => {
  try {
    await exec('rm .env')
  } finally {
    const { stdout } = await exec(`node ${path}`)
    return eq(stdout.trim(), 'undefined')
  }
}

tests.push(testEmpty, testGetSecrets)
Object.freeze(tests)
