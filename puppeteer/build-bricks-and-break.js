const body = document.querySelector('body')

const create = (tag) => {
  const element = document.createElement(tag)
  return element
}

export const build = (amount = 54) => {
  let count = 1
  const intervalID = setInterval(() => {
    const brick = create('div')
    brick.title = 'brick'
    brick.id = `brick-${count}`
    if (count % 3 === 2) {
      brick.dataset.foundation = true
    }
    brick.append(count)
    body.append(brick)

    if (count === amount) {
      window.clearInterval(intervalID)
      return
    }

    count++
  }, 100)
}

export const repair = (...ids) => {
  ids.forEach((id) => {
    const toRepair = document.getElementById(id)
    if (toRepair) {
      toRepair.dataset.repaired = toRepair.hasAttribute('data-foundation')
        ? 'in progress'
        : true
    }
  })
}

export const destroy = () => {
  const bricks = [...document.querySelectorAll('[title="brick"]')]
  const toRemove = bricks[bricks.length - 1]

  if (toRemove) {
    toRemove.remove()
  }
}
