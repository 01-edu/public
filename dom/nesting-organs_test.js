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

  // check the section selector style has been updated properly
  const sectionSelectorStyle = await page.evaluate(() => {
    const target = [...window.document.styleSheets[0].cssRules].find(
      (rule) => rule.selectorText === 'section',
    )
    const { display, justifyContent } = target.style
    return { display, justifyContent }
  })
  eq({ display: 'flex', justifyContent: 'center' }, sectionSelectorStyle)

  // check if the given CSS has been correctly copy pasted

  // div,
  // p {
  //   border: solid 1px black;
  //   padding: 10px;
  //   margin: 0;
  //   border-radius: 30px;
  // }

  // #face {
  //   align-items: center;
  // }

  // #eyes {
  //   display: flex;
  //   background-color: yellow;
  //   justify-content: space-between;
  //   align-items: center;
  //   border-radius: 50px;
  //   width: 200px;
  // }

  // #torso {
  //   width: 200px;
  //   background-color: violet;
  // }
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
