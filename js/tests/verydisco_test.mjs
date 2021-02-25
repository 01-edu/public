import { promisify } from 'util'
import * as cp from 'child_process'

const exec = promisify(cp.exec)

export const tests = []

tests.push(async ({ path, eq }) => {
  const { stdout } = await exec(`node ${path} discovery`)
  return eq(stdout.trim(), 'verydisco')
})

tests.push(async ({ path, eq }) => {
  const { stdout } = await exec(`node ${path} "kiss cool"`)
  return eq(stdout.trim(), "sski olco")
})

tests.push(async ({ path, eq }) => {
  const { stdout } = await exec(`node ${path} kiss cool`)
  return eq(stdout.trim(), "sski")
})

tests.push(async ({ path, eq }) => {
  const { stdout } = await exec(`node ${path} "Node is awesome"`)
  return eq(stdout.trim(), "deNo si omeawes")
})
