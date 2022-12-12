## file-researcher

### Instructions

Create a file `file-researcher.sh`, which will look in the `file-researcher/facts.txt` file for the following:

- The complete sentence which begins like so `It takes 12 honey`.
- All the lines where the word `year` appears.

You can not use "echo"!

Expected Output:

```console
$ ./file-researcher.sh | cat -e
It takes 12 honeybees to make one teaspoon of honey$
6:898 tornadoes were recorded to have occurred in the United States in the year 2000.$
12:The only two days of the year in which there are no professional sportsgames (MLB, NBA, NHL, or NFL) are the day before and the day after theMajorLeague All-Star Game.$
15:All dogs are the descendant of the wolf. These wolves lived in eastern Asia about 15,000 years ago$
32:Most fleas do not live past a year old$
39:Every 238 years, the orbits of Neptune and Pluto change making Neptune at times the farthest planet from the sun$
60:A maple tree is usually tapped when the tree is at least 45 years old and has a diameter of 12 inches$
63:An acre of trees can remove about 13 tons of dust and gases every year from the surrounding environment$
67:Every year, 100 million sharks are killed by people$
70:One female mouse can produce up to 100 babies a year$
95:In Johannesburg, the average car will be involved in an accident once every four years.$
96:The youngest actress to be nominated as best actress is Keisha Castle-Hughes who was nominated at just 13 years old$
121:In an average lifetime, a person will spend 4 years traveling in an automobile and six months waiting at a red light.$
$
```

### Hints

`grep` is a command-line utility for searching plain-text data sets for lines that match a regular expression.

```console
$ head example.txt
Indonesia consists only of islands - 13,667 total
During World War II, the very first bomb dropped on Berlin by the Allies killed the only elephant in the Berlin Zoo
$ grep "Indonesia" example.txt
Indonesia consists only of islands - 13,667 total
```

> You have to use Man or Google to know more about commands flags, in order to solve this exercise!
> Google and Man will be your friends!
