export const getArchitects = () => {
  const architects = [...document.getElementsByTagName('a')]
  const others = [...document.getElementsByTagName('span')]
  return [architects, others]
}

export const getClassical = () => {
  const classicals = [...document.getElementsByClassName('classical')]
  const others = [...document.querySelectorAll('a:not(.classical)')]
  return [classicals, others]
}

export const getActive = () => {
  const active = [...document.querySelectorAll('.classical.active')]
  const others = [...document.querySelectorAll('.classical:not(.active)')]
  return [active, others]
}

export const getBonannoPisano = () => {
  const bonanno = document.getElementById('BonannoPisano')
  const others = [
    ...document.querySelectorAll('a.classical.active:not(#BonannoPisano)'),
  ]
  return [bonanno, others]
}
