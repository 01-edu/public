> Due to file size reason, the solution might be uploaded on GitHub instead of Gitea!

#### General

##### If, at any point during the audit, coupled code is detected that could have been resolved using the Observer pattern (`event dispatchers`), the audit will stop, and the student will fail the project.

###### Is the game following a coherent theme ?

#### Main Menu

###### Is the main menu widget on a separate level/map from the main game map?

###### Does the main menu contain a Start game button?

###### Does the main menu contain an Exit game button?

#### Player Character

###### Does the player character have a skeletal mesh?

###### Can the player character move only along two axes: left-right and up-down?

###### Does the player character have a basic locomotion system with all the required animations?

###### Does the player character transition smoothly between walking and running animations based on their speed?

#### Collectible

###### Does the collectible Actor have a static mesh?

###### Does the collectible rotate around an axis?

###### Does the collectible have a box collider that acts as a trigger?

###### Is the collectible collected when the player enters its box trigger?

###### Does the collectible play a sound when collected?

#### HUD

###### Does the HUD widget display a value related to the `collectible`?

###### Is the HUD updated each time a `collectible` is picked up?

###### Is the HUD also updated when picking a `collectible` that spawned from killing an enemy?

#### Enemy

###### Does the enemy character have a skeletal mesh?

###### Does the enemy character have at least a walking cycle?

###### Does the enemy character have a simple AI that patrols between two set points?

###### Does the enemy character kill the player on collision?

###### Is the enemy character killed if the player lands on top of it?

###### Does the enemy character play a sound when killed?

###### Does the enemy character create an instance of a `Collectible` at the place of death?

#### Game loop

###### Is there a spawn point where the player starts?

###### Does the player respawn at the spawn point when dead?

###### Is there an end point that defines the player’s goal to finish the level?

###### When the player reaches the end point, is a menu displayed with options to restart or quit the game?

#### Level design

###### Does the level design avoid being just a long, empty run to the finish point?

###### Does the level include mechanics chosen by the student? `(e.g. moving platforms skill-based platforming)`

###### Does the level provide a fun and challenging experience for the player?

#### Bonus

###### +Has a health system been added to both the player character and the enemy character using Components?

###### +Has a simple crouch been added to the player character's locomotion system?

###### +Have additional mechanics been implemented to fit the overall game design?

###### +Has a simple Game Design Document been created that describes the mechanics and outlines the plan for implementing them?
