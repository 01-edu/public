import { styles } from './data.js'

let count = 0
let increment = 1

export const pimp = () => {
  const button = document.querySelector('.button')

  const ceiling = count === styles.length - 1
  const floor = !count

  const increasing = increment > 0

  if (increasing || (floor && !increment)) {
    button.classList.add(styles[count])
  } else {
    button.classList.remove(styles[count])
  }

  if (ceiling) {
    increment = increment ? 0 : -1
  }
  if (floor) {
    increment = increment < 0 ? 0 : 1
  }

  button.classList.toggle('unpimp', increment < 0 || (!increment && ceiling))

  count += increment
}
