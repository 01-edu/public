// fake `getJSON` function
let getJSON = async (url) => url
// /*/ // ⚡
export const tests = []
const t = (f) => tests.push(f)

// queryServers with main server be fastest
t(async ({ eq, ctx }) => {
  ctx.setTimings({ pouet_backup: 2 })
  return eq(await queryServers('pouet', ctx.r), `/pouet?q=${ctx.r}`)
})

// queryServers with backup server be fastest
t(async ({ eq, ctx }) => {
  ctx.setTimings({ pouet: 2 })
  return eq(await queryServers('pouet', ctx.r), `/pouet_backup?q=${ctx.r}`)
})

// gougleSearch fast enough
t(async ({ eq, ctx }) => {
  ctx.setTimings({ web_backup: 3, image: 2, video_backup: 4 })
  return eq(await gougleSearch(ctx.r), {
    web: `/web?q=${ctx.r}`,
    image: `/image_backup?q=${ctx.r}`,
    video: `/video?q=${ctx.r}`,
  })
})

// gougleSearch fast enough, alternate timings
t(async ({ eq, ctx }) => {
  ctx.setTimings({ web: 3, image_backup: 1, video: 4 })
  return eq(await gougleSearch(ctx.r), {
    web: `/web_backup?q=${ctx.r}`,
    image: `/image?q=${ctx.r}`,
    video: `/video_backup?q=${ctx.r}`,
  })
})

// gougleSearch too slow !
t(async ({ eq, ctx }) => {
  ctx.setTimings({ web: 85, web_backup: 99 })
  return eq(
    await gougleSearch(ctx.r).then(
      () => Promise.reject(Error('Should fail')),
      (err) => err.message,
    ),
    'timeout',
  )
})

Object.freeze(tests)

export const setup = () => ({
  r: Math.random().toString(36).slice(2),
  setTimings: (timings) =>
    (getJSON = (url) =>
      new Promise((s) =>
        setTimeout(s, timings[url.split(/\/([^?]+)?/)[1]] || 0, url),
      )),
})
