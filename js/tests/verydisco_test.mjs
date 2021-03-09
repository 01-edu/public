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

tests.push(async ({ path, eq, randStr }) => {
  const tic = randStr(7)
  const tac = randStr(7)
  const { stdout } = await exec(`node ${path} "${tic}${tac}"`)
  return eq(stdout.trim(), `${tac}${tic}`)
})

tests.push(async ({ path, eq, randStr }) => {
  const ying = randStr(8)
  const yang = randStr(7)
  const { stdout } = await exec(`node ${path} "${ying}${yang}"`)
  return eq(stdout.trim(), `${yang}${ying}`)
})

tests.push(async ({ path, eq, randStr }) => {
  const tic = randStr(5)
  const tac = randStr(5)
  const ying = randStr(3)
  const yang = randStr(2)
  const { stdout } = await exec(`node ${path} "${tic}${tac} ${ying}${yang}"`)
  return eq(stdout.trim(), `${tac}${tic} ${yang}${ying}`)
})
