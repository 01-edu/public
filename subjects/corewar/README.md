## Core War

### Introduction

Back in 1984 D. G. Jones had the brilliant idea to create a programming game
inspired by how some viruses manage to take over a system and their internal
memory.

From the very start the game was a great success and a flourishing community
added new features and organized tournaments worldwide.

In this project we will recreate the necessary tools to run a modified version
of the game.

### Game dynamics

At its very base the game consists on having two or more programs (written
by the contestants) that will be executed by the Virtual Machine in an arena (a
sandbox memory space).

We will call those programs the "players" from now on since they will be the ones
playing, while their creators won't have any possibility to interact with them
or the arena during the match.

The player will have a limited set of instructions it could use to change the
state of the arena.
It can use those instructions to try and corrupt the opponent's memory space,
replicate itself or repair damages made by the enemies.

The winner will be the last player notifying the arena it was alive before the
game ends.

#### Structure of the game

<!-- About CPU abstraction and toolchain -->

#### End game

### The Assembler

### The Virtual Machine

### Resources

#### Instruction Set

#### Config file

#### A basic player

#### Testing environment
