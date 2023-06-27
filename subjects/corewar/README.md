## Core War

### Introduction

Back in 1984 D. G. Jones had the brilliant idea to create a programming game inspired by how some viruses manage to take over a system and their internal memory.

From the very start the game was a great success and a flourishing community added new features and organized tournaments worldwide.

In this project we will recreate the necessary tools to run a more elaborate version of the game.

### General instructions

- The Virtual Machine and the Assembler can be done in any compiled language.
- You should only use the standard library of the language you choose.
- In case you use any external library it must be largely justified and explained (avoid it at any cost).
- You must not use external libraries to do the parsing part of the Assembler for any reason.
- The programs should never crash for any reason.
- The programs should never leak memory for any reason.

### Game dynamics

At its very core the game consists of having two or more programs (written by the contestants) that will be executed by the Virtual Machine in an arena (a sandbox memory space).

We will call those programs the "players" from now on since they will be the ones playing. Their creators won't have any possibility to interact with them or the arena during the match.

The player will have a limited set of instructions it could use to change the state of the arena.
It can use those instructions to try and corrupt the opponent's memory space, replicate itself or repair damages made by the enemies.

The winner will be the last player notifying the arena it was alive before the game ends.

#### Structure of the game

The Virtual Machine is a simplified reproduction of how a CPU works.
It will execute the binary code (the players) in the arena in cycles.

At the start of the game each player will have one process, and this process will be placed at the start of each player memory space (the current position of a process is called program counter, or PC).

Each process will also have a private memory space (the registers) and a carry.
A carry is a special flag that may be changed by the result of some instructions, it can be used in many clever ways.

The players will be able to create new processes using specific instructions, it is worth noting though that every instruction takes a certain amount of cycles to be executed, and creating a new process is the most expensive one.

#### End game

There is a special instruction, called `live` that notifies the VM one player is still healthy and doing well.

Every once in a while the VM will check which processes executed such instruction and kill all the processes that have not.
The processes are kept alive only if they execute at least one `live` instruction between the VM checks.

Once there is no more active processes the VM stops execution and the last player that notified it was alive is the winner of the match.

The VM will use some rules to decrease the elapsed time since the last check down to zero, this means the games can never be infinite and all processes will be killed at one point.

### The Assembler

We don't want to write code in binary, but we don't want the VM to deal with the complexity of parsing high level instructions at execution time.

To solve this issue we will create an Assembler which will take an assembly code as input and deliver a binary program as output.

You should follow those specifications:

- If no parameters are passed it will print a help message.
- The assembler takes a `file.s` as input.
- It returns a `file.cor` as output.
- `file.s` is a regular and boring text file.
- `file.cor` is a binary file, in order to inspect it we suggest you to use command-line tools like `hexdump`.
- If there is any error in the source file the Assembler should exit with an error code, print a message on `stderr` (the more specific the better) and do not create any `file.cor`.
- The binary must be written in [big-endian](https://en.wikipedia.org/wiki/Endianness).

### The Virtual Machine

The Virtual Machine role is to execute the binary programs (players) provided as `.cor` files (for a maximum of 4 players).

The VM must handle a `-d [NB_CYCLES]` flag (stating for "dump").
If the flag is specified the VM will stop execution at `NB_CYCLES` and dump the memory in the arena in hexadecimal, writing 32 bytes per row.

You should follow those specifications:

- If no parameters are passed it will print a help message.
- If one of the `.cor` file is corrupted the VM should exit with an error code, print a message on `stderr` and do not execute the programs.
- At the start of the battle the VM will print a welcome message as the one showed in the example.
- At the end of the battle the VM will write the winner (if any) as the one showed in the example.
- The players will be loaded into the arena starting by the first byte and they will be evenly spaced from each other.
- The last program passed will be the first one executed during the cycle.
- During a cycle the VM will load the instruction at the current PC and wait N cycles before to execute it (N being the cost of the instruction).
- Only when executing the instruction the VM will check for the parameters it may have and it will execute it only if the parameters are correct, otherwise it will print an error on `stderr` and continue to execute.
- If an instruction has incorrect parameters the PC will be moved forward according to the size of the parameters.
- If the instruction doesn't exist in the instruction set the PC will be moved forward of 1 byte.
- When a new process is forked it will be placed at the end of the processes and start its execution at the start of the next cycle (which means it will be the first one being executed on the next cycle).
- The entire execution is deterministic, so with the same inputs you must always have the same outputs.
- The VM assumes the binary is in big-endian and read it accordingly.

Those are the cases where a file is considered corrupted:

- Wrong signature code.
- Declared size of the program not matching the actual size.
- The size of the program is bigger that the maximum allowed size.
- The total file size is smaller than the minimum size.

### Greeting and end game messages

At the start of each game the VM should print:

```
For this match the players will be:
Player 1 ([X] bytes): [NAME] ([DESCRIPTION])
Player 2 ([X] bytes): [NAME] ([DESCRIPTION])
...
```

At the end of the game the vm should print `cycle [X]: The winner is player [X]: [NAME]!`.

> If nobody executed a valid `live` statement the end message should be `cycle [X]: Nobody wins!`.

#### Circular memory space of the arena

The memory where the players will fight is circular, this means if we want to move forward from the last address in memory (for example `4095`) we will arrive at address `0`.

It implies that moving backward of one position from `0` will bring us to the address `4095`.

> Having circular memory guarantees players they will never overflow the memory space they are playing in.

#### Stop process execution

Every `CYCLE_TO_DIE` the VM will check every process and kill all the processes that did not successfully execute any `live` instruction.

To avoid infinite games, `CYCLES_TO_DIE` will be decremented by `CYCLE_DELTA` under certain conditions:

- If during the last life loop there were at least `NBR_LIVE` successfully executed by the players.
- If it has been `MAX_CHECKS` life loops since it was decremented last time.

> Notice that some smart players may trick another player to start making `live` statements for them, this virtually means a player can still make `live` statements after all its processes were killed.

### The Assembly language

In the assembly file there must be a `.name "[Name of the player]"` field and a `.description "[Description of the player]"` before any instruction.

Every instruction has to be on a single line and may be preceded by a label declaration.
The instruction name is separated from the optional label and the parameters by one or more spaces.

The parameters are separated from each other by `,`.

The instructions will be of different sizes depending on the number and type of parameters passed.

#### Parameters types

- Register parameter

Each process has its own registers (from 1 to `REG_NUMBER`).
Each of those registers can hold 32 bits, but the register number itself is written in only 8 bits (which are enough to go from 1 to 16).

They are formatted as `r1`, `r2`, `...`, `r16`.

The VM will initialize the registers to `0` for each player except for `r1` which will have `- PLAYER_ID`, so for the first player `r1 == -1`.

- Direct parameter

A direct parameter is a value that represents itself, it is formatted by using `%` before the value. So `%10` will represent the number ten.

- Indirect parameter

An indirect parameter will give you the relative position of the address where to look for the value, think of it as an ancestor of raw pointers.

In this case `10` will say to the VM to look ten bytes forward the current instruction and take the value from there.

> Don't be scared to use the provided Asm and VM to play with those different parameters in order to be sure you understood properly their behavior.

#### Instruction Set

| Name  | Nb Params | Opcode | Nb Cycles | Has Pcode | Has Idx | Params                                                                     |
| ----- | --------- | ------ | --------- | --------- | ------- | -------------------------------------------------------------------------- |
| live  | 1         | 1      | 10        | false     | false   | Direct                                                                     |
| ld    | 2         | 2      | 5         | true      | false   | [Indirect, Direct]<br> Register                                            |
| st    | 2         | 3      | 5         | true      | false   | Register<br>[Register, Indirect]                                           |
| add   | 3         | 4      | 10        | true      | false   | Register<br> Register<br> Register                                         |
| sub   | 3         | 5      | 10        | true      | false   | Register<br> Register<br> Register                                         |
| and   | 3         | 6      | 6         | true      | false   | [Register, Indirect, Direct]<br> [Register, Indirect, Direct]<br> Register |
| or    | 3         | 7      | 6         | true      | false   | [Register, Indirect, Direct]<br> [Register, Indirect, Direct]<br> Register |
| xor   | 3         | 8      | 6         | true      | false   | [Register, Indirect, Direct]<br> [Register, Indirect, Direct]<br> Register |
| zjmp  | 1         | 9      | 20        | false     | true    | Direct                                                                     |
| ldi   | 3         | 10     | 25        | true      | true    | [Register, Indirect, Direct]<br> [Register, Direct]<br> Register           |
| sti   | 3         | 11     | 25        | true      | true    | Register<br> [Register, Indirect, Direct]<br> [Register, Direct]           |
| fork  | 1         | 12     | 800       | false     | true    | Direct                                                                     |
| lld   | 2         | 13     | 10        | true      | false   | [Indirect, Direct]<br> Register                                            |
| lldi  | 3         | 14     | 50        | true      | true    | [Register, Indirect, Direct]<br> [Register, Direct]                        |
| lfork | 1         | 15     | 1000      | false     | true    | Direct                                                                     |
| nop   | 1         | 16     | 2         | true      | false   | Register                                                                   |

- `live`: says to the VM the player with the id matching the opposite of the first parameter is alive. If the first parameter is `-2` it means player 2 is alive.
- `ld`: loads the first parameter into the registry passed as second parameter. (If the first parameter is Indirect, the VM will apply `% IDX_MOD` to it)
- `st`: writes the value in the registry passed as the first parameter into the second parameter. (If the second parameter is Indirect, the VM will apply `% IDX_MOD` to it)
- `add`: sums the first two arguments and writes the result into the third one.
- `sub`: subtract the first two arguments and writes the result into the third one.
- `and`, `or`, `xor`: apply the respective bitwise operation to the first two parameters and saves it in the third one. (If the first and/or second parameters are Indirect, the VM will apply `% IDX_MOD` to them)
- `zjmp`: moves the PC adding the value passed as the first parameter. (The VM will always apply `% IDX_MOD` to the first parameter)
- `ldi`: writes the value contained at the address obtained by summing the first two arguments into the register in the third argument. (The VM will apply `% IDX_MOD` to the obtained address)
- `sti`: writes the value of the register passed as the first argument into to address obtained by summing the last two arguments. (The VM will apply `% IDX_MOD` to the obtained address)
- `fork`: creates a new process which will be an exact (deep) copy of the current one except for the PC which will be at the address specified in the argument. (The VM will apply `% IDX_MOD` to the first argument)
- `lld`, `lldi`, `lfork`: those instructions are the long versions of `ld`, `ldi` and `fork`, which means the addresses won't be truncated by `IDX_MOD`. (Since those instructions are "long" the modulo operation won't be applied on them)
- `nop`: this operation does nothing, it can still take one argument and has a pcode so be careful on how you implement it.

> All addresses are relative to the current PC of the process.

> Some instructions have to truncate the addresses they are given by applying modulo `IDX_MOD`.
> This feature prevents players from reaching spaces in memory that are too far from the current PC, which means processes won't be able to attack each other straight away but will need to move by doing smaller steps, this help having more balanced games.

#### Labels

Labels are also called pseudo instructions, in the sense they won't be translated into binary code.
They are widely used to help programmers creating assembly code without needing to remember and count relative memory positions.

So when looking in the binaries the labels will be replaced in order to match their relative position in bytes from the place they were used.

As an example in the `ameba.s` player `zjmp %:hello` is exactly the same of `zjmp %-5` (`-5` being `ff fb` in hexadecimal notation).

This is because from the `zjmp` you will need to go back 5 bytes in order to match the label `hello:` declaration.

> Notice that labels can be declared at the start of an instruction or on their own line.

#### Pcode field

Some instruction may accept different kind of parameters, and those parameters may require different sizes.
For those instructions the assembler will use a byte after the opcode to inform the VM which kind of parameters to expect.

Every two bits of this byte will inform the VM about the type of a parameter.

- `01` will be used for a register parameter.
- `10` will be used for a direct parameter.
- `11` will be used for an indirect parameter.

An example of a pcode byte could be `01111000`:

Here the first parameter is expected to be a register, the second should be an indirect value and the third a direct value.

> Those numbers are in binary base, for example `10` in binary is `2` in decimal.

#### Has IDX field

Some instruction will sum up direct values and/or use them as addresses to lookup into the arena memory space. For this reason we know that a full 32 bits integer won't be necessary and to save space we do save direct values on only two bytes in those cases.

You have `Had Idx` column in the Instruction Set table, when it is true you must save and read direct parameters on two bytes, when false they will be saved on four bytes as usual.

#### The Carry flag

In all processors you will find a carry which is a flag that notify some edge case states to allow special operations, ensure security and serve many others functionalities.

The Carry of the processes in the VM is in that sense a simplified version of this concept.

The following commands will modify the carry of the process executing them: `ld`, `add`, `sub`, `and`, `or`, `xor`.

Those commands will set the carry to `true` if the value written in the register is `0`, and set it to `false` otherwise.

Only one command reads the carry and it's `zjmp`, it will do the jump if the carry is true and will do nothing otherwise.

#### The file signature

The signature, also called magic, of a file is a particular value written at the start of a binary. This signature has the role to help the OS understand which kind of file it is and how to execute it. For the `.cor` files the signature is provided in the config file.

### Your player

While the main focus for now is to create an Asm and a VM, you should provide a basic player that should be able to fight and win against the `ameba.s`, this player will be tested in the provided VM during evaluation.

### Resources

#### Config file

The specification requires a lot of constants, it would be tedious to define them in the subject every time they are mentioned.

For this reason we provide a [config file](data/config.md) which includes all the required constants. Some of them are specific to the Asm, some others to the VM, the vast majority will be helpful to both.

> While the file is language agnostic, it would be very easy to translate it to a specific language. We suggest you to do so in order to centralize the configuration, it will help to debug and extend easily your programs if needed.

#### A basic player

To help understand better how the Asm and the VM works, let's deep dive into the assembly code with a very basic player:

```
.name "ameba"
.description "not doing much"

        sti r1,%:hello,%1
        and r1,%0,r1

hello:  live %1
        zjmp %:hello
```

This player is one of the most basic living things you can imagine, it just says it is alive.

Here is how it proceeds in order to do so:

- The `sti` instruction will copy the value in the first register (which is the opposite of its id, so `-1` if this player is the first one).
- It is interesting that this number will be copied at the address in memory corresponding to the label `hello` + 1 byte.
- Said in other words this instruction is changing the argument of the `live` statement.
- The `and` instruction will copy the result of the AND logic operator into the first register.
- This result will be zero since any number AND 0 is zero. This has the side effect of setting the carry of this process to true.
- The `live` statement notify the VM the player is alive. Notice the specific value of the argument is irrelevant in this context since the `sti` instruction override it before the `live` is executed.
- The `zjmp` brings the PC back to the start of the live instruction (creating a loop).
- Notice `zjmp` only operates the jump if the carry is not equal to zero. This is why we executed the `and` instruction before.

Now let's see what this player binary looks like with `hexdump -C ameba.cor`:

```
00000000  00 ea 83 f3 61 6d 65 62  61 00 00 00 00 00 00 00  |....ameba.......|
00000010  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
*
00000080  00 00 00 00 00 00 00 00  00 00 00 17 6e 6f 74 20  |............not |
00000090  64 6f 69 6e 67 20 6d 75  63 68 00 00 00 00 00 00  |doing much......|
000000a0  00 00 00 00 00 00 00 00  00 00 00 00 00 00 00 00  |................|
*
00000890  0b 68 01 00 0f 00 01 06  64 01 00 00 00 00 01 01  |.h......d.......|
000008a0  00 00 00 01 09 ff fb                              |.......|
000008a7
```

- All numbers are represented in hexadecimal.
- At the start of each line there is the position of the first byte of the line in hexadecimal.
- This means that for the second line `10` hex is actually 16 in decimal notation.
- The right panel shows a `.` if the byte is not a valid ASCII character.
- Repeated lines are aggregated under `*`, this is why we pass from `10` to `80`.
- Each `xx` is a byte (a byte is 8 bits, `ff` in hexadecimal is `1111 1111` in binary, `255` in decimal).

Let's read it through:

- The first 4 bytes `00 ea 83 f3` are the program signature (also called the magic), this is a 32 bits integer, represented as four 8 bits slices.
- Then the next 128 bits are for the name of the program, the first five `61 6d 65 62 61` are the actual name, the others are zeros because we don't actually need it.
- Then we have 4 bytes `00 00 00 17` (so 23) which will be a 32 bits integer with the size in bytes of the program to be executed in the arena of the VM (so only the size of the instructions, without the name, description, signature and so on).
- Then we have the description which works exactly the same as the name but it will add four bytes padding at the end of it.

The actual instructions start from there:

- The instruction `sti` is made by those bytes: `0b 68 01 00 0f 00 01`.
- `0b` is the opcode, `68` is the pcode, `01` is the first parameter, `00 0f` is the second parameter and `00 01` is the third parameter.
- Notice the last two parameters are written on two bytes each because `sti` has `Has Idx` set to true.
- Next instruction is `and`, and it is `06 64 01 00 00 00 00 01`.
- `06` is the opcode, `64` is the pcode, the first parameter is `01`, the second one is `00 00 00 00` and the third one is `01`.
- Here comes the `live` instruction: `01 00 00 00 01`.
- `01` is the opcode, `00 00 00 01` is the first parameter, notice the parameter memory space will be overridden by the `sti` instruction.
- The last instruction is `zjmp`: `09 ff fb`.
- `01` is the opcode while `ff fb` is the first parameter (on two bytes because of `Has Idx` once again).

#### Testing environment

You will be provided with a [playground](data/playground.zip) containing:

- A VM and an Assembler binaries to use as references.
- A Dockerfile to run the binaries in a container.
- Instructions on how to execute the references.
- Some players to use as tests.

> The provided VM has an extra flag `-v` which you can use to print the state of the VM at every cycle, this should greatly help you during development and debugging.

### Some advices for the road

While the functioning of the VM and the Assembler may seem quite strange and obscure, it is a perfect scenario to reproduce a CPU and understand what a computer really does at its core.

In this project you are indeed creating a Turing Complete machine largely inspired by Von Neumann architecture (which is still how today's processors works).

### Bonus

Those bonuses will be evaluated only if the Assembler and the VM are perfect.
You can use external libraries (for example graphic libraries) in order to achieve them.

- Create a disassembler that will take binary as input and return a `.s` file as output.
- Create a visualizer to show what is happening in the VM in real time.
- Introduce arithmetic operations support in the Assembly language.
- Introduce a simple macro system in the Assembly language.

### Resources

- [Turing Machine](https://en.wikipedia.org/wiki/Turing_machine)
- [Von Neumann Architecture](https://en.wikipedia.org/wiki/Von_Neumann_architecture)
