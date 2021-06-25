# Schema Public

## Tables

### audit

Some exercises require reviews from other students or any other member of the school. The reviewer checks that some conditions are met by the code of the student being reviewed. This process is called an `audit` and it's kept in this table.

Columns:
- id.
- groupId: The group that the audit refers to.
- auditorId: The user who's making the audit.
- attrs: The attributes of the audits keep the feedback of the auditor for each question failed.
- grade: 0 for failed and 1 for passed.
- createdAt: The timestamp of the time and date of creation.
- updatedAt: The timestamp of the last update.
- code: The code that must be used by the auditor to start the audit.
- resultId: The result of the audit.
- version: A sha of the commit that the audit refers to.
- endAt: The last date that the id refers to.
- private: Used to access the code by the auditor.

### discordToken

Columns:
- id
- accessToken
- refreshToken
- expiresAt

### event

An event anchors an objects into time.
> Note: Some [object attributes](attributes.md) are related to time, but they express time only in relative terms to the beginning of the event.
> Also, live edits (i.e., editing an event while is ongoing) it's not supported. The changes will only take effect after reloading the page.

Columns:
- id
- createdAt: The timestamp of the creation of the event.
- endAt: The timestamp for the end of the event.
- registrationId: The registration to the event.
- objectId: The object that is used to generate the event.
- parentId: The parent id of the event is based on.
- status: Confirmation that the event finished correctly (NULL or done for now).
- path: The path of the event (the relative url of the event).
- campus: The campus that contains the event.
- code: Used to unlock the exams.

### event_user

This table links users and events. A user can be in many events and an events have many users.

Columns:
- id
- createdAt
- userId
- eventId
- groupId (not used anymore, TODO: remove at some point)

### group

Some projects are required to be made by groups. Those groups are kept in this table.

Columns:
- id
- objectId: The object that is being "done" by the group. Usually projects or raids.
- eventId : The groups only exist in relation to an event
- captainId: The captain of the group.
- createdAt.
- updatedAt.
- status: See [group_status](#group_status) for possible values.
- path.
- campus.

### group_status

Columns:
- status
Possible values (for now)
  - `setup`
  - `working`
  - `audit`
  - `finished`

### group_user

This table links users to groups. A user can be in many groups and a group contains many users.

Columns:
- id
- userId
- groupId
- confirmed: Boolean. False by default. True after the user represented by the user id confirms the group.
- createdAt
- updatedAt

### match

It is used for the `bet` system in the bonus exercises. Specifically to keep a record of the matches between students who will review and make a bet about each others code.

Columns:
- id
- createdAt
- updatedAt
- objectId: The object that requires the match.
- userId: The user that is matched.
- matchId: The matched that is made to another user.
- confirmed: The same as in group_user it's a way to keep track of invitations that had not been confirmed.
- bet: boolean; For the optional exercises that require a bet. A bet can be true if the exercises is expected to pass and false otherwise.
- result: boolean; True for correct bet.
- path
- campus
- eventId: The event the object in inserted in. (Piscine, Raid, etc).

### object

Objects are generic representations of elements in the structure of the curriculum (Onboarding, campus, piscine, raids, etc). They are arranged in a hierachical structure to allow unlimited nesting (see [object_child](#object_child)). They also keep the timeline of the events relaying in durations (ex. `duration`, `eventDuration`, etc) specified by [attributes](attributes.md) (`attrs`).

Columns:
- id
- name
- type: See [object_type](#object_type) for the possible types.
- status: (Currently not in used, but it should always be `online` to avoid errors TODO: remove at some point).
- attrs: See [object attributes](attributes.md).
- childrenAttrs: Attributes that are applied to all the children objects.
- createdAt
- updatedAt
- externalRelationUrl (Maybe not needed).
- authorId (Maybe not needed).
- campus: The campus the object is related to.
- referenceId: A reference to the base object.
- referencedAt: The timestamp of the creation of the copy object.

### object_child

It keeps a child-parent relationship between two objects. It's used to encapsulate objects inside each other.

Columns:
- id
- parentId
- childId
- attrs
- key: When generating the JS object this field will be the `key` of the child.
- index: Defines the position of the child object inside the parent object.

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

It is a register of the steps in the progression of students doing the exercises and projects.

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

Keeps a register for the bans. It can also keep the record of other situations, this is done by setting the banEndAt to a date already passed and the reason for the record.

Columns:
- id
- userId
- authorId
- message
- banEndAt
- createdAt

### registration

Registrations to events (For example, an exam).

Columns:
- id
- createdAt, When the registration was created.
- startAt, When users can start registering to an event.
- endAt, When the registration ends.
- eventStartAt, When the event the registration refers to starts.
- objectId, The object the registration refers to.
- parentId, The parent object of the object the registration refers to.
- attrs
- path
- campus

### registration_user

The user registered to events. A user can register to many events and and event can have many users.

Columns:
- id
- registrationId
- userId
- createdAt

### result

The result for each exercise and project made in the platform.

Columns:
- id
- createdAt
- updatedAt
- grade
- progressId
- attrs
- type, For the possible types see the table [result_type](#result_type)
- userId
- groupId
- objectId
- path
- version, The version of the code the result refers to.
- eventId
- isLast, Used to represent transitory states. A progress is only considered finished when `isLast` is `true`.
- campus

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
- slug
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

It's a register of the rewards given to students.

Columns:
- id
- type, See [transaction_type](#transaction_type) for the possible types.
- amount:
  - type = xp; The amount of Xp that an object rewards.
  - type = up || down; a percentage of the change (10%);
- userId
- attrs
- createdAt
- path
- objectId
- eventId
- campus

### transaction_type

Columns:
- type
  - `xp`: transaction giving xp
  - `up`: transaction correspondent to reviewing someone
  - `down`: transaction correspondent to reviewing you

### user

Columns:
- id
- githubId: (Now gitea).
- githubLogin: (Now gitea).
- discordId
- discordLogin
- profile
- attrs: Extra information about the users (email, address, etc).
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
