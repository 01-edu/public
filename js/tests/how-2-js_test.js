import { readFileSync as read } from 'fs'

// /*/ // ⚡

export const tests = [
  ({ eq }) =>
    eq(
      read('/jail/student/index.html', 'utf8').trim(),
      '<script type="module" src="how-2-js.js"></script>',
    ),
  ({ eq }) =>
    eq(
      read('/jail/student/how-2-js.js', 'utf8').trim(),
      `console.log('Hello World')`,
    ),
]
