const hashCode = str =>
  (
    [...str].reduce((h, c) => (h = (h << 5) - h + c.charCodeAt(0)) & h, 0) >>> 0
  ).toString(36)
// /*/ // ⚡
export const tests = []
const t = (f) => tests.push(f)

t(({ eq }) => {
  const { chain, ...block } = blockChain({ a: 1 })
  return eq(block, {
    index: 1,
    data: { a: 1 },
    prev: { index: 0, hash: '0' },
    hash: '1103f27',
  })
})

t(({ eq }) => {
  const block = blockChain({ a: 1 }).chain({ hello: 'world' })

  const chain = block
    .chain({ value: 4455 })
    .chain({ some: 'data' })
    .chain({ cool: 'stuff' })

  const fork = block
    .chain({ value: 335 })
    .chain({ some: 'data' })
    .chain({ cool: 'stuff' })

  return eq(
    {
      chainRoot: chain.prev.prev.prev,
      forkRoot: chain.prev.prev.prev,
      chainHash: chain.hash,
      forkHash: fork.hash,
      index: chain.index,
      data: chain.data,
    },
    {
      chainRoot: block,
      forkRoot: block,
      chainHash: '1qr3qfs',
      forkHash: '1x9gsc1',
      index: fork.index,
      data: fork.data,
    },
  )
})

t(({ eq }) => {
  const random = Math.random()
  const hash = hashCode(`21103f27{"random":${random}}`)
  const { prev, chain, ...block } = blockChain({ a: 1 }).chain({ random })
  return eq(block, { index: 2, data: { random }, hash })
})

Object.freeze(tests)
