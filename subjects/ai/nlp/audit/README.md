#### Exercise 0: Environment and libraries

##### The exercise is validated is all questions of the exercise are validated

##### Activate the virtual environment. If you used `conda` run `conda activate your_env`.

##### Run `python --version`.

###### Does it print `Python 3.x`? x >= 8

##### Does `import jupyter`, `import pandas`, `import nltk` and `import sklearn` run without any error?

---

---

#### Exercise 1: Lower case

##### The exercise is validated is all questions of the exercise are validated

##### The question 1 is validated if the output is:

```
0    this is my first nlp exercise
1                         wtf!!!!!
Name: text, dtype: object
```

##### The question 2 is validated if the output is:

```
0    THIS IS MY FIRST NLP EXERCISE
1                         WTF!!!!!
Name: text, dtype: object
```

---

---

#### Exercise 2: Punctuation

##### The question 1 is validated if the ouptut doesn't contain punctuation `` !"#$%&'()*+,-./:;<=>?@[]^_`{|}~ ``. Do not take into account the spaces in the output. The output should be as:

```
Remove this from  the sentence
```

---

---

#### Exercise 3: Tokenization

##### The exercise is validated is all questions of the exercise are validated

##### The question 1 is validated if the ouptut is:

```
['Bitcoin is a cryptocurrency invented in 2008 by an unknown person or group of people using the name Satoshi Nakamoto.',
'The currency began use in 2009 when its implementation was released as open-source software.']

```

##### The question 2 is validated if the ouptut is:

```
['Bitcoin',
'is',
'a',
'cryptocurrency',
'invented',
'in',
'2008',
'by',
'an',
'unknown',
'person',
'or',
'group',
'of',
'people',
'using',
'the',
'name',
'Satoshi',
'Nakamoto',
'.',
'The',
'currency',
'began',
'use',
'in',
'2009',
'when',
'its',
'implementation',
'was',
'released',
'as',
'open-source',
'software',
'.']

```

---

---

#### Exercise 4: Stop words

##### The question 1 is validated if, using NLTK, the ouptut is:

```
['The', 'goal', 'exercise', 'learn', 'remove', 'stop', 'words', 'NLTK', '.', 'Stop', 'words', 'usually', 'refers', 'common', 'words', 'language', '.']
```

---

---

#### Exercise 5: Stemming

##### The question 1 is validated if, using NLTK, the output is:

```
['the', 'interview', 'interview', 'the', 'presid', 'in', 'an', 'interview']
```

---

---

#### Exercise 6: Text preprocessing

##### The question 1 is validated if the output is:

```
['01',
 'edu',
 'system',
 'present',
 'innov',
 'curriculum',
 'softwar',
 'engin',
 'program',
 'renown',
 'industrylead',
 'reput',
 'curriculum',
 'rigor',
 'design',
 'learn',
 'skill',
 'digit',
 'world',
 'technolog',
 'industri',
 'take',
 'differ',
 'approach',
 'classic',
 'teach',
 'method',
 'today',
 'learn',
 'facilit',
 'collect',
 'cocré',
 'process',
 'profession',
 'environ']

```

---

---

#### Exercise 7: Bag of Word representation

##### The exercice is validated is all questions of the exercice are validated

##### The question 1 is validated if the output of the CountVectorizer is

```
<6588x500 sparse matrix of type '<class 'numpy.int64'>'
	with 79709 stored elements in Compressed Sparse Row format>
```

##### The question 2 is validated if the output of `print(df.iloc[:3,400:403].to_markdown())` is:

    |    |   talk |   team |   tell |
    |---:|-------:|-------:|-------:|
    |  0 |      0 |      0 |      0 |
    |  1 |      0 |      0 |      0 |
    |  2 |      0 |      0 |      0 |

##### The question 3 is validated if the shape of the wordcount DataFrame `(6588, 501)` is and if the output of `print(df.iloc[300:304,499:501].to_markdown())` is:

    |     |   youtube |   label |
    |----:|----------:|--------:|
    | 300 |         0 |       0 |
    | 301 |         0 |      -1 |
    | 302 |         1 |       0 |
    | 303 |         0 |       1 |
