# Schema Public

## Tables

### audit
Columns:
- id
- groupId, the group that the audit refers to
- auditorId, the user who's making the audit
- attrs, what are used for. What are the possible ones? TODO: look through the code and take them from functions.js
- grade, either 1 or 0
- createdAt, the timestamp of the time and date of creation
- updatedAt, the timestamp of the last update
- code, the code that should be used by the auditor to start the audit
- resultId, the result of the audit
- version, a sha of the commit that the audit refers to
- endAt, the last date that the id refers to
- private, used to access the code by the auditor (needs to be refactored because the requirements changes)

### discordToken
Columns:
- id
- accessToken
- refreshToken
- expiresAt

### event
Columns:
- id
- createdAt, the timestamp of the creation of the event
- endAt, the timestamp with the time the event finishes at
- registrationId, the registration to the event
- objectId, the object that is used to generate the event
- parentId, the parent id of the event is based on
- status, confirmation that the event finished correctly (NULL or done for now)
- path, the path of the event (the relative url of the event)
- campus, the campus the event is inserted in
- code, only used to unlock the exams

### event_user
Columns:
- id
- createdAt
- userId
- eventId
- groupId (not used anymore, TODO: remove at some point)

### group
Columns:
- id
- objectId, The object that is being done by the group (Usually projects or raids)
- eventId , The groups only exist in relation to an event
- captainId, The captain of the group
- createdAt
- updatedAt
- status, See group_status for possible values
- path
- campus (Auto generated for by the path)

### group_status
Columns:
- status
Possible values (for now)
  - `setup`
  - `working`
  - `audit`
  - `finished`

### group_user
Columns:
- id
- userId
- groupId
- confirmed, Transitory state while it's not confirmed means
- createdAt
- updatedAt

### match
Columns:
- id
- createdAt
- updatedAt
- objectId, The object that requires the match
- userId, The user that is matched
- matchId, The matched that is made to another user
- confirmed (The same as in group_user it's way to keep track of invitations that had not been confirmed)
- bet (For the optional exercises that require a bet. A bet can be true if the exercises is expected to pass)
- result, says if the exercise was validated as succeeded by the tester (TODO: reward xp for getting a bet right)
- path
- campus
- eventId, the event the object in inserted in. (Piscine, Raid, etc)

### object
Anything that defines the structure of the curriculum? (Onboarding, campus, piscine, raids, etc)
Columns:
- id
- name
- type, See object type for the possible values
- status (Not used: always online TODO: remove at some point)
- attrs, Same as attrs in the object look for them in `functions.js`
- childrenAttrs Attributes that are applied to all the children objects
- createdAt
- updatedAt
- externalRelationUrl (Not used?)
- authorId (Not used?)
- campus, The campus the object is related to
- referenceId, A reference to the base object
- referencedAt, The timestamp of the creation of the copy object

#### Attributes (attrs)
##### Objects
- inScope: This attribute is used to determine if the exercise gives XP or not. If an exercise is not in scope then it doesn't give XP.

- status: It can have the values `blocked`, `available`, `succeeded` or `failed`
TODO: Finish documenting the json attributes that can be used
##### Quests
- `duration`, `startDay`, `start`, `scopeExtraDuration`, `hasStarted`, `special`, `scopeEnd`, `rootName`, `repositoryPath`, `status`, `scopeStart`, `difficulty`, `difficultyMod`, `baseXp`, `language`, `group`, `exerciseType`, `preRequisiteId`, `delay` `parentType`, `inScope`, `xpIndex`, `level`, `difficulty`

### object_child
Created a relationship parent child between two objects, it's used to encapsulate objects inside each other.
Columns:
- id
- parentId
- childId
- attr
- key, The 'key' in the object child relationship
- index, Defines the position of the child object inside the parent object

### object_status
Columns:
- status
Possible values:
  - `online`
  - `offline`
  - `draft`

### object_type
Columns:
- type
Possible values:
  - `onboarding`
  - `campus`
  - `exercise`
  - `quest`
  - `signup`
  - `exam`
  - `raid`
  - `project`
  - `piscine`

### progress
What is a progress?
Columns:
- id
- createdAt
- updateAt
- userId
- groupId
- eventId
- version
- grade
- isDone
- path
- campus
- objectId

### record
Record of what?
Is it just for banned people?
Columns:
- id
- userId
- authorId
- message
- banEndAt
- createdAt

### registration
Registrations to events
Columns:
- id
- createdAt, When the registration was created
- startAt, When user can start registering to an event
- endAt, When the registration ends
- eventStartAt, When the event the registration refers to starts
- objectId, The object the registration refers to
- parentId, The parent object of the object the registration refers to
- attrs, What do the attributes of the registration are used for?
- path
- campus

### registration_user
The user registered to events
Columns:
- id
- registrationId
- userId
- createdAt

### result
The result
Columns:
- id
- createdAt
- updatedAt
- grade
- progressId, The result is liked to a progress?
- attrs, What are they used for?
- type, For the possible types see the table result_type
- userId
- groupId
- objectId
- path
- version, The version of the code the result refers to.
- eventId
- isLast, What is it used for?
- campus, Again isn't it available in the object already?

### result_type
Columns:
- type
  - `tester`
  - `user_audit`
  - `admit_audit`
  - `admin_selection`
  - `status`

### role
Columns:
- id
- slug (slug vs. name?)
- name
- description
- createdAt
- updatedAt

### token
Columns:
- id
- status
- createdAt
- updatedAt

### token_status
Columns:
- status
  - `active`
  - `expired`

### transaction
Where is the link between the table xp and the transaction?
Columns:
- id
- type, See transaction_type for the possible types
- amount (does it work for up and down also?)
- userId
- attrs (TODO: specify and document)
- createdAt
- path
- objectId
- eventId
- campus

### transaction_type
Columns:
- type
  - `xp`, transaction giving xp
  - `up`, transaction correspondent to reviewing someone
  - `down`, transaction correspondent to reviewing you

### user
Columns:
- id
- githubId
- githubLogin
- discordId
- discordLogin
- profile
- attrs
- createdAt
- updateAt
- discordDMChannelld
- campus
- audits

### user_role
Columns:
- id
- userId
- roleId

### xp
Columns:
- userId
- eventId
- amount
