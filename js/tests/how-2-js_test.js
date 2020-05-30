import { readFileSync as read } from 'fs'

// /*/ // ⚡

export const tests = [
  ({ eq }) =>
    eq(
      read('/jail/student/index.html', 'utf8').trim(),
      '<script type="module" src="hello-world.js"></script>',
    ),
  ({ eq }) =>
    eq(
      read('/jail/student/hello-world.js', 'utf8').trim(),
      `console.log('Hello World')`,
    ),
]
