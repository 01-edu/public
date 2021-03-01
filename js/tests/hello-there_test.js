import { readFileSync as read } from 'fs'

// /*/ // ⚡
export const tests = [
  ({ eq }) => // code must use console.log
    read('/jail/student/hello-there.js', 'utf8').trim().includes('console.log'),

  async ({ eq, ctx }) => {
    // console.log must have been called with the right value
    const log = console.log.bind(console)
    const args = []
    console.log = (..._args) => {
      args.push(..._args)
      log(..._args)
    }
    await import('/jail/student/hello-there.js')
    console.log = log
    eq(ctx.args.join(' ').trim(), 'Hello There')
  },
]
