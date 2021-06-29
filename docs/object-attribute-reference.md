# Object Attributes (attrs)

# This is still a work in progress (mistakes are common).

## For objects that have associated events

> Required in **bold**

- **capacity**: number; The number of users that can be admitted in the event associated to the event.
- **registrationDuration**: number; The time of minutes a registrations lasts.
- **eventDuration**: number; The duration in minutes of the event associated to the object.
- eventStartDelay: number: The time in minutes between the end of the registration and the beginning of the event.
- expectedXp: number; The amount of `Xp` an object gives.
- baseXp: number; The starting point of the `Xp`. It's used to calculate the `Xp` per child object.

## Some specific attributes for object types:

- Exercise
  - **language**: string; The programming language of the exercise, ex: 'go'
  - start: type; Description
  - delay: number; The amount of minutes that has to pass after the previous object before the current one is unlocked.
  - special: boolean; True for a hackaton.
  - startDay: number; The day relative to the beginning of the parent event.
  - scopeStart
  - scopeEnd
  - scopeExtraDuration
  - auditsRequired: number; The number of approved audits that need to be done to consider an exercise as approved
  - exerciseType: string
    - `required`
    - `enrichment`
    - `bonus`
    - `asynchronous`
    - `synchronous`
  - inScope, This attribute is used to determine if the exercise gives `Xp` or not. If an exercise is not in scope then it doesn't give `Xp`.
  - status: It can have the values `blocked`, `available`, `succeeded` or `failed`.
  - hasStarted
  - rootName
  - repositoryPath
  - difficulty: A number that represents how hard it should be to complete an exercise.
  - difficultyMod: A multiplier of the difficulty

- Exam exercises:
On top of the exercise of the attributes of a normal exercise exam exercises need the following attributes.
  - **group**: number; Identifies the position that an exercise can take in an exam. The exam takes exercises randomly for each group. For instance: the first exercise of the exam will be an exercise with `group` = 1, the second one will be an exercise with `group` = 2, and so on.

- Raid
  - groupSize: number; The number of elements wanted in a raid group. The actual number of elements in a group will be depenedent in the availability of candidates. For example if the groupSize is 3 and there are 8 users in the event, you'll have groups with sizes: 3-3-2.

- Project
  - groupMax: number; The max amount of members a group should have to be created.
  - groupMin: number; The minimum amount of members a group should have to be created.

- Campus
  - TODO

- Quest:
  - TODO

- Piscine:
  - TODO

- Onboarding:
  - TODO

- Signup:
  - TODO
