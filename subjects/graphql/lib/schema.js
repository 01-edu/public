const backImg =
  'https://cdn1.iconfinder.com/data/icons/thin-ui-1/100/Noun_Project_100Icon_1px_grid_thin_ic_arrow_left_simple-512.png'
const body = document.body
const docDisplay = document.querySelector('.display-doc')
const exist = () => document.querySelector('.back')
const removeComments = (obj) =>
  Object.fromEntries(obj.filter((v) => v[0] !== 'comment'))

const getData = async (pathToFile) => {
  const data = await fetch(pathToFile)
    .then((data) => data.json())
    .catch(console.log)
  return data
}

const generateLinks = (k, v) => {
  const h4 = document.createElement('h4')
  const a = document.createElement('a')
  a.innerText = v
  h4.innerHTML = k
  h4.appendChild(a)
  docDisplay.appendChild(h4)
  return a
}

const generateAllQueries = (data) => {
  const docTitle = document.querySelector('.doc-title')
  docDisplay.innerHTML = ''
  const back = backArrow()
  docTitle.appendChild(back)
  for (let obj of data.types) {
    for (const [k, _] of Object.entries(removeComments(Object.entries(obj)))) {
      generateContent(k, obj)
    }
  }
}

const loadRootTypes = async () => {
  const data = await getData('./lib/schema.json')
  for (let v of data.root_types) {
    const h4 = document.createElement('h4')
    const a = document.createElement('a')
    a.addEventListener('click', (event) => changeContent(event, data))

    for (const [k, value] of Object.entries(v)) {
      a.innerText = value
      h4.innerHTML = `${k}: `
      h4.appendChild(a)
      docDisplay.appendChild(h4)
    }
  }
  const a = generateLinks('all: ', 'all_available_queries')
  a.addEventListener('click', () => generateAllQueries(data))
}

const generateContent = (t, data) => {
  docDisplay.innerHTML += `type ${t.toUpperCase()} {`
  for (const [k, v] of Object.entries(data[t])) {
    generateLinks(`     ${k}: `, v)
  }
  docDisplay.innerHTML += `}<br><br>`
  if (t === 'query') {
    for (let v of document.getElementsByTagName('a')) {
      v.addEventListener('click', (e) => changeContent(e, data))
    }
  }
}

const backArrow = () => {
  const title = document.querySelector('.title')
  const typeDescription = document.querySelector('.comment')
  const contType = document.querySelector('.doc-type')
  const origType = contType.innerText
  const origTitle = title.innerText
  const origTypeDescription = typeDescription.innerText
  const back = document.createElement('img')
  back.className = 'back'
  back.src = backImg

  back.addEventListener('click', () => {
    back.remove()
    docDisplay.innerHTML = ''
    title.innerText = origTitle
    contType.innerText = origType
    typeDescription.innerText = origTypeDescription
    loadRootTypes()
  })
  return back
}

const allArgs = (data) => {
  for (const [k, v] of Object.entries(data.arguments)) {
    docDisplay.innerHTML += `<h4>${k}: <a>(${v.description})</a> // ${v.comment}</h4>`
  }
}

const changeContent = ({ target }, data) => {
  const title = document.querySelector('.title')
  const typeDescription = document.querySelector('.comment')
  const contType = document.querySelector('.doc-type')
  const docTitle = document.querySelector('.doc-title')

  const t = target.innerText
  const back = backArrow()

  title.innerText = t.toUpperCase()
  contType.innerText = `${t.toUpperCase()} fields`
  docDisplay.innerHTML = ''
  if (/\[\w*\!\]\!/g.test(t)) {
    const b = t.replace(/\[|\]|\!/g, '')
    const obj = data.types.filter((v) => v[b])[0]
    typeDescription.innerText = obj.comment
    generateContent(b, obj)
  } else if (t === 'arguments') {
    allArgs(data)
    !exist() && docTitle.appendChild(back)
  } else {
    generateContent(t, data)
    !exist() && docTitle.appendChild(back)
  }
}

loadRootTypes()

/* search */
const updateResult = (dataToSearch, { target }) => {
  docDisplay.innerHTML = ''

  dataToSearch.search.map((type) =>
    target.value.split(' ').map((w) => {
      if (type.toLowerCase().indexOf(w.toLowerCase()) !== -1) {
        docDisplay.innerHTML += `<li><a class="query-result">${type}</a></li>`
        document
          .querySelectorAll('.query-result')
          .forEach((ele) =>
            ele.addEventListener('click', (e) => changeContent(e, dataToSearch))
          )
      }
    })
  )
}

const initSearch = async () => {
  const data = await getData('./lib/schema.json')
  const input = document.querySelector('.search')

  input.addEventListener('focus', () => {
    const bg = document.querySelector('.back')
    if (!bg) {
      const back = backArrow()
      document.querySelector('.doc-title').appendChild(back)
    }
    docDisplay.innerHTML = ''
  })
  input.oninput = (e) => updateResult(data, e)
}

initSearch()
/*--------------*/
