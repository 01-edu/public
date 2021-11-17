#### Functional

###### Has the requirement for the allowed packages been respected? (Reminder for this project: only [standard packages](https://golang.org/pkg/))

##### Try passing as arguments `"banana standard --output test00.txt"`

```
Usage: go run . [STRING] [BANNER] [OPTION]

EX: go run . something standard --output=<fileName.txt>
```

###### Does it display the same result as above?

##### Try passing as arguments `"First\nTest" shadow --output=test00.txt`

```
student$ cat test00.txt
                                       $
_|_|_|_| _|                     _|     $
_|          _|  _|_|   _|_|_| _|_|_|_| $
_|_|_|   _| _|_|     _|_|       _|     $
_|       _| _|           _|_|   _|     $
_|       _| _|       _|_|_|       _|_| $
                                       $
                                       $
                                      $
_|_|_|_|_|                     _|     $
    _|       _|_|     _|_|_| _|_|_|_| $
    _|     _|_|_|_| _|_|       _|     $
    _|     _|           _|_|   _|     $
    _|       _|_|_| _|_|_|       _|_| $
                                      $
                                      $
$
```

###### Does it save the right output in the right file?

##### Try passing as arguments `"hello" standard --output=test01.txt`

```
student$ cat test01.txt
 _              _   _          $
| |            | | | |         $
| |__     ___  | | | |   ___   $
|  _ \   / _ \ | | | |  / _ \  $
| | | | |  __/ | | | | | (_) | $
|_| |_|  \___| |_| |_|  \___/  $
                               $
                               $
$
```

###### Does it save the right output in the right file?

##### Try passing as arguments `"123 -> #$%" standard --output=test02.txt`

```
student$ cat test02.txt
                                   __             _  _      _     _   __ $
 _   ____    _____                 \ \          _| || |_   | |   (_) / / $
/ | |___ \  |___ /         ______   \ \        |_  __  _| / __)     / /  $
| |   __) |   |_ \        |______|   > >        _| || |_  \__ \    / /   $
| |  / __/   ___) |                 / /        |_  __  _| (   /   / / _  $
|_| |_____| |____/                 /_/           |_||_|    |_|   /_/ (_) $
                                                                         $
                                                                         $
$
```

###### Does it save the right output in the right file?

##### Try passing as arguments `"432 -> #$%&@" shadow --output=test03.txt`

```
student$ cat test03.txt
                                                                                                                  $
_|  _|   _|_|_|     _|_|                    _|             _|  _|     _|   _|_|    _|   _|           _|_|_|_|_|   $
_|  _|         _| _|    _|                    _|         _|_|_|_|_| _|_|_| _|_|  _|   _|  _|       _|          _| $
_|_|_|_|   _|_|       _|         _|_|_|_|_|     _|         _|  _|   _|_|       _|       _|_|  _| _|    _|_|_|  _| $
    _|         _|   _|                        _|         _|_|_|_|_|   _|_|   _|  _|_| _|    _|   _|  _|    _|  _| $
    _|   _|_|_|   _|_|_|_|                  _|             _|  _|   _|_|_| _|    _|_|   _|_|  _| _|    _|_|_|_|   $
                                                                      _|                           _|             $
                                                                                                     _|_|_|_|_|_| $
$
```

###### Does it save the right output in the right file?

##### Try passing as arguments `"There" shadow --output=test04.txt`

```
student$ cat test04.txt
                                               $
_|_|_|_|_| _|                                  $
    _|     _|_|_|     _|_|   _|  _|_|   _|_|   $
    _|     _|    _| _|_|_|_| _|_|     _|_|_|_| $
    _|     _|    _| _|       _|       _|       $
    _|     _|    _|   _|_|_| _|         _|_|_| $
                                               $
                                               $
$
```

###### Does it save the right output in the right file?

##### Try passing as arguments `"123 -> \"#$%@" thinkertoy --output=test05.txt`

```
student$ cat test05.txt
                                    o o         | |               $
  0    --  o-o            o         | |  | |   -O-O-      O   o   $
 /|   o  o    |            \            -O-O- o | |   o  /   / \  $
o |     /   oo              O            | |   -O-O-    /   o O-o $
  |    /      |       o-o  /            -O-O-   | | o  /  o  \    $
o-o-o o--o o-o            o              | |   -O-O-  O       o-  $
                                                | |               $
                                                                  $
$
```

###### Does it save the right output in the right file?

##### Try passing as arguments `"2 you" thinkertoy --output=test06.txt`

```
student$ cat test06.txt
                         $
 --                      $
o  o                     $
  /        o  o o-o o  o $
 /         |  | | | |  | $
o--o       o--O o-o o--o $
              |          $
           o--o          $
$
```

###### Does it save the right output in the right file?

##### Try passing as arguments `"Testing long output!" standard --output=test07.txt`

```
student$ cat test07.txt
 _______                _     _                         _                                                 _                     _     _  $
|__   __|              | |   (_)                       | |                                               | |                   | |   | | $
   | |      ___   ___  | |_   _   _ __     __ _        | |   ___    _ __     __ _          ___    _   _  | |_   _ __    _   _  | |_  | | $
   | |     / _ \ / __| | __| | | | '_ \   / _` |       | |  / _ \  | '_ \   / _` |        / _ \  | | | | | __| | '_ \  | | | | | __| | | $
   | |    |  __/ \__ \ \ |_  | | | | | | | (_| |       | | | (_) | | | | | | (_| |       | (_) | | |_| | \ |_  | |_) | | |_| | \ |_  |_| $
   |_|     \___| |___/  \__| |_| |_| |_|  \__, |       |_|  \___/  |_| |_|  \__, |        \___/   \__,_|  \__| | .__/   \__,_|  \__| (_) $
                                           __/ |                             __/ |                             | |                       $
                                          |___/                             |___/                              |_|                       $
$
```

###### Does it save the right output in the right file?

##### Try passing as arguments a random string with lower and upper case letters, and the output flag ("--output=") followed by a random file name.

###### Does it save the right output in the right file?

##### Try passing as arguments a random string with lower case letters, numbers and spaces, and the output flag ("--output=") followed by a random file name.

###### Does it save the right output in the right file?

##### Try passing as arguments a random string with special characters, and the output flag ("--output=") followed by a random file name.

###### Does it save the right output in the right file?

##### Try passing as arguments a random string with lower, upper case, spaces and numbers letters, and the output flag ("--output=") followed by a random file name.

###### Does it save the right output in the right file?

###### Are all the test files tested saved?

###### As an auditor, is this project up to every standard? If not, why are you failing the project?(Empty Work, Incomplete Work, Invalid compilation, Cheating, Crashing, Leaks)

#### Basic

###### +Does the project runs quickly and effectively (Favoring of recursive, no unnecessary data requests, etc.)?

###### +Is the output of the program well structured? The characters are correctly in line?

###### +Is there a test file for this code?

###### +Are the tests checking each possible case?

###### +Does the code obey the [good practices](../../good-practices/README.md)?

#### Social

###### +Did you learn anything from this project?

###### +Would you recommend/nominate this program as an example for the rest of the school?
