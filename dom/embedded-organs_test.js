export const tests = []

tests.push(async ({ page, eq }) => {
  // check that the HTML structure is correct & elements are nested properly
  const elements = await page.$$eval('body', (nodes) => {
    const toNode = (el) => {
      const node = {}
      node.tag = el.tagName.toLowerCase()
      node.id = el.id
      if (el.children.length) {
        node.children = [...el.children].map(toNode)
      }
      return node
    }
    return [...nodes[0].children].map(toNode)
  })
  eq(expectedStructure, elements)
})

tests.push(async ({ page, eq }) => {
  // check the section selector style has been updated properly
  eq.css('section', {
    display: 'flex',
    justifyContent: 'center',
  })
})

tests.push(async ({ page, eq }) => {
  // check if the provided CSS has been correctly copy pasted
  eq.css('div, p', {
    border: '1px solid black',
    padding: '10px',
    margin: '0px',
    borderRadius: '30px',
  })

  eq.css('#face', { alignItems: 'center' })

  eq.css('#eyes', {
    display: 'flex',
    backgroundColor: 'yellow',
    justifyContent: 'space-between',
    alignItems: 'center',
    borderRadius: '50px',
    width: '200px',
  })

  eq.css('#torso', {
    width: '200px',
    backgroundColor: 'violet',
  })
})

const expectedStructure = [
  {
    tag: 'section',

    id: 'face',
    children: [
      {
        tag: 'div',

        id: 'eyes',
        children: [
          { tag: 'p', id: 'eye-left' },
          { tag: 'p', id: 'eye-right' },
        ],
      },
    ],
  },
  {
    tag: 'section',

    id: 'upper-body',
    children: [
      { tag: 'div', id: 'arm-left' },
      { tag: 'div', id: 'torso' },
      { tag: 'div', id: 'arm-right' },
    ],
  },
  {
    tag: 'section',

    id: 'lower-body',
    children: [
      { tag: 'div', id: 'leg-left' },
      { tag: 'div', id: 'leg-right' },
    ],
  },
]
