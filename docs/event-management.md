# Events management

## Usage
> An event is associated to an object when the usage of this object implies:
>   * a limited capacity of people
>   * a limited time

> Event management require:
>   * Edition of the attributes of the object on which the event is based
>   * Creation and settings of the event associated to the initial object

## Settings for the initial object

> Important indications:
>   * Objects that doesn't have required attributes for event creation will not be open to event creation. 
> | name                 | fullfillment |
> | -------------------- | ---------    |
> | capacity             | **required** |
> | eventDuration        | **required** |
> | registrationDuration | **required** |
> | description          | optionnal    |
> | eventStartDelay      | optionnal    |
>   * All the attributes filled in the object are used as values by default for event's creation; it can be overloaded for each event related to the initial object.
>   * If the initial object has a child or children which are events itself, settings are also required for each event child.

#### Edit the object attributes:
> in *Object attributes*

<img width="1073" alt="Capture d’écran 2019-08-22 à 11 40 34" src="https://user-images.githubusercontent.com/35296671/63525316-64e0da80-c4f5-11e9-9e61-57d5a73da9b1.png">

* Add a new key **capacity** of type `Number` with the maximum number of persons you want for events related to the object by default
* Add a new key **eventDuration** of type `Number` with the duration in minutes you want for events related to the object by default
* Add a new key **registrationDuration** of type `Number` with the duration in minutes you want to allow to people to register to the event by default
* Add a new key **eventStartDelay** of type `Number`, if you want a default delay between the end of registration and the beginning of the event. This duration is expressed in minutes.
* Add a new key **description** of type `String`, if you need to associate some informations to your event (description, location, access, documents to provide, etc.)

#### Edit the children
> in *Children*

<img width="609" alt="Capture d’écran 2019-08-22 à 15 43 43" src="https://user-images.githubusercontent.com/35296671/63525543-c86b0800-c4f5-11e9-8820-60d9ff33994f.png">

* Add a new key **startAfter** of type `Number`, with the default delay you want between the beginning of the event and the beginning of the child event. This duration is expressed in minutes.

##### Example

Here is an example of the `piscine-go` settings. It presents the settings of the object attributes `piscine-go`, the settings of one of its child which is an event and the settings of the child object attributes itself. 

> NB : this object settings are provided in the admin, in the curses section: 'Piscine Go' and in the exams section 'Exam 01'.

**Piscine Go**
*Object attributes*
```json
{
  "capacity": 400,
  "eventDuration": 37440,
  "registrationDuration": 43200,
  "eventStartDelay": 240
}
```

*Children*
> In the `piscine-go`, children of type *exam* and *rush* have events itself. 

> A **startAfter** key has to be defined for each of them, in their parent object `piscine-go`. For example, the exam-01 gets this key:

```json
{
  "startAfter": 400
}
```

**Exam 01**
> The object `Exam 01`, which is a child of `Piscine Go`, has its own *Object Attributes* filled in the child object.

*Object attributes*
```json
{
  "eventDuration": 240,
  "registrationDuration": 2160,
  "eventStartDelay": 60
}
```
> NB: the **capacity** attribute is herited from the parent object `Piscine Go here`

## Create the event

### Create a new event for your object

### Settings for you event

> In the event you have created, 3 categories must be checked:
> 1. General settings
> 2. Registration's settings
> 3. Event's settings

#### General settings
#### Registration
#### Event
