#### Exercise 0: Environment and libraries

##### The exercise is validated if all questions of the exercise are validated

##### Activate the virtual environment. If you used `conda` run `conda activate your_env`.

##### Run `python --version`.

###### Does it print `Python 3.x`? x >= 9

###### Do `import jupyter`, `import pandas`, `import nltk` and `import sklearn` run without any error?

---

---

#### Exercise 1: Lower case

##### The exercise is validated if all questions of the exercise are validated

###### For question 1, is the output the following?

```
0    this is my first nlp exercise
1                         wtf!!!!!
Name: text, dtype: object
```

###### For question 2, is the output the following?

```
0    THIS IS MY FIRST NLP EXERCISE
1                         WTF!!!!!
Name: text, dtype: object
```

---

---

#### Exercise 2: Punctuation

###### For question 1, is validated if the ouptut doesn't contain punctuation `` !"#$%&'()*+,-./:;<=>?@[]^_`{|}~ ``. Is the previous statement true? Do not take into account the spaces in the output. The output should be as:

```
Remove this from  the sentence
```

---

---

#### Exercise 3: Tokenization

##### The exercise is validated if all questions of the exercise are validated

###### For question 1, is output the following?

```
['Bitcoin is a cryptocurrency invented in 2008 by an unknown person or group of people using the name Satoshi Nakamoto.',
'The currency began use in 2009 when its implementation was released as open-source software.']

```

###### For question 2, is the output the following?

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

###### For question 1, is the output the following? (using NLTK)

```
['The', 'goal', 'exercise', 'learn', 'remove', 'stop', 'words', 'NLTK', '.', 'Stop', 'words', 'usually', 'refers', 'common', 'words', 'language', '.']
```

---

---

#### Exercise 5: Stemming

###### For question 1, is the output the following? (using NLTK)

```
['the', 'interview', 'interview', 'the', 'presid', 'in', 'an', 'interview']
```

---

---

#### Exercise 6: Text preprocessing

###### For question 1, is the output the following?

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

##### The exercise is validated if all questions of the exercise are validated

###### For question 1, is the output of the `CountVectorizer` the following?

```
<6588x500 sparse matrix of type '<class 'numpy.int64'>'
	with 37334 stored elements in Compressed Sparse Row format>
```

###### For question 2, is the output of `print(count_vecotrized_df.iloc[:3,400:403].to_markdown())` the following?

```python
    |    |   someth |   son |   song |
    |---:|---------:|------:|-------:|
    |  0 |        0 |     0 |      0 |
    |  1 |        0 |     0 |      0 |
    |  2 |        0 |     0 |      0 |
```

###### For question 3, is the output matching with the following one?

```python
cant    1
deal    1
end     1
find    1
keep    1
like    1
may     1
say     1
talk    1
Name: 3, dtype: Sparse[int64, 0]
```

###### For question 4, is the output matching with the following one?

```python
tomorrow    1126
go           733
day          667
night        641
may          533
tonight      501
see          439
time         429
im           422
get          398
today        389
game         382
saturday     379
friday       375
sunday       368
dtype: int64
```

###### For question 5, is the output of `print(count_vectorized_df.iloc[350:354,499:501].to_markdown())` the following?

```python
|     |   your |   label |
|----:|-------:|--------:|
| 350 |      0 |       1 |
| 351 |      1 |      -1 |
| 352 |      0 |       1 |
| 353 |      0 |       0 |
```
