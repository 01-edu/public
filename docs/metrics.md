## Metrics card - Export of students' results

### Where to find the data

Metrics for students from a given campus can be found in the Metrics card in the Admin dashboard.

### Important information

- Data can be exported in csv and json formats
- Data is filtered by campus, exports include all historical data for that campus.
  - We will later develop a feature that enables filtering information before downloading it
- Data is updated every hour
- The download link is valid for 24 hours
- We recommend sharing the downloaded document itself instead of the download link as the link contains the admin’s authorisation token

### Description of the data included in the exports

| Field           | Example                                 | Description                                                                                                                                                                                                                                                                                                                        |
| --------------- | --------------------------------------- | ---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| **user**        | MarieMalarme                            | Student’s github username                                                                                                                                                                                                                                                                                                          |
| **first_name**  | Marie                                   | Student’s first name                                                                                                                                                                                                                                                                                                               |
| **last_name**   | Malarme                                 | Student’s last name                                                                                                                                                                                                                                                                                                                |
| **object**      | who-are-you                             | Name of the object for which a result is given                                                                                                                                                                                                                                                                                     |
| **type**        | exercise                                | There are 3 object types: <br/> - exercise (for quests and exams) <br/> - raid <br/> - project                                                                                                                                                                                                                                     |
| **parent**      | Quest 01                                | The object’s hierarchical parent: <br/> - If the object is an exercise, the parent will be either a quest or an exam. <br/> - If the object is a raid or a project, the parent will be the module (piscine or div)                                                                                                                 |
| **module**      | piscine-go                              | The part of the curriculum in which the object is included, either a piscine (piscine Go, piscine JS, piscine Rust) or the main curriculum (Div01)                                                                                                                                                                                 |
| **path**        | /madere/piscine-go/quest-01/who-are-you | The url to that object, which outlines the path and hence the hierarchical structure to that object                                                                                                                                                                                                                                |
| **attempts**    | 2                                       | Number of submission attempts made by the student                                                                                                                                                                                                                                                                                  |
| **status**      | succeeded                               | Current status of the student's work on the object: <br/> - "available" indicates that the student has started working on the object but either hasn't submitted it yet to be corrected, or is waiting for the auditing process to be completed <br/> - "succeeded" and "failed" are attributed once the object has been corrected |
| **xp**          | 325                                     | Number of XPs gained by the student for succeeding                                                                                                                                                                                                                                                                                 |
| **base_xp**     | 325                                     | Maximum number of XPs the student can get for succeeding for that object                                                                                                                                                                                                                                                           |
| **grade**       | 1                                       | The student’s score for that object is a number between 0 (failed) and 1 (succeeded). If the object is a project, the grade can be superior to 1 when the student successfully passes bonus questions                                                                                                                              |
| **last_update** | 2020-07-27T15:39:46.368Z                | Date and time of the student’s latest activity for that object                                                                                                                                                                                                                                                                     |

### Explanation of the XP and Grade calculation

#### Exercises within quests

- _Test mode_: automatic tester
- _XPs_: fixed value. Students get all the points when succeeding at the exercise, no matter how many attempts they made.
- _Grade_: 0 (failure) or 1 (success)

#### Exercises within exams

- _Contextual information_:
  - Exams consist of many exercises which are grouped by difficulty levels. The student is randomly assigned an exercise from each level and has to succeed in order to move on to the next exercise level.
  - Each exercise level has a maximum amount of XPs which grows according to the level’s difficulty.
  - Exams are limited in time so students have to successfully complete as many exams as they can as fast as they can with the least amount of errors.
- _Test mode_: automatic tester
- _XPs_: variable value. Students get the maximum amount of points when they succeed at the first attempt. Each failed submission reduces the amount of XPs the student can get for that level.
- _Grade_: 0 (failure) or 1 (success)

#### Raids

- _Test mode_: projects are evaluated by a jury panel in 2 steps:
  - First the jury evaluates the code through an audit, a series of pass / fail questions to be answered by the auditor. Succeeding the audit gives students a certain amount of XPs.
  - The jury then interviews students and can decide to give them bonus / malus XPs based on their evaluation of the code and the students’ answers to their questions.
- _XPs_: variable value calculated as a fixed value for succeeding the project plus / minus XPs attributed by the jury at their discretion
- _Grade_: 0 (failure) or 1 (success)

#### Projects

- _Test mode_: audits, which are peer-to-peer evaluations based on a series of pass/fail questions to be answered by the auditor
- _XPs_: variable value. Students can get more than the maximum value of XPs if they complete bonus exercises. They get no XPs at all if they fail at least one of the questions in the audit
- _Grade_: number between 0 and 1, proportionate to the number of questions successfully passed in the audit. The grade can be higher than 1 if the student completed bonus exercises.
