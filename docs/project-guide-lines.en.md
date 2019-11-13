# Project Guide Lines

## project shape

- create a new table for the questions of review

```
const project = {
  // take the exercise structure
  id: 3001,
  name: 'ascii-art',
  type: 'project',
  key: 'ascii-art',
  index: 1,
  path: 'piscine-go/quest-01/ascii-art',
  attrs: {
    subject: {
      name: 'translate',
      type: 'function',
    },
    'subject-en': 'https://public.01-edu.org/subjects/ascii-art.en',
    'subject-fr': 'https://public.01-edu.org/subjects/ascii-art.fr',
    expectedFiles: ['hello.sh'],
    allowedFunctions: [],
    language: 'go',
    difficulty: 0.5,
    xpIndex: 0,
    level: 0,
    baseXp: 175,
    parentCursus: '01-alem',
    parentType: 'cursus',
    exerciseType: 'required', // equivalent on the graph = 'mandatory', 'optional', 'asynchronous'
  },
  // pass review questions as children of the project
  children: [
    // generate form input for each question according to its type (boolean / numeric / string...)
    { questionId: 1, weight: 1, type: 'mandatory' },
    { questionId: 2, weight: 3, type: 'mandatory' },
    { questionId: 3, weight: 2, type: 'mandatory' },
    { questionId: 4, weight: 1, type: 'optional' },
    { questionId: 5, weight: 5, type: 'optional' },
  ],
}

const inputObject = 
  {
    id: 1,
    label: 'Is the website coded in Go ?',
    type: 'checkbox',
    tag: 'functional',
    required: true,
  }
```


## peer review system

- algorithm(not friends/enemies, location, no correction before)

## calendar with slots

- 3 people/slot (1 one at the time to begin)

- 10 people available are picked by the algorithm

- 3 to 10 reviewers/project(acording to difficulty)
