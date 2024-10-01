#### Exercise 0: Environment and libraries

##### The exercise is validated if all questions of the exercise are validated.

##### Activate the virtual environment. If you used `conda` run `conda activate your_env`.

##### Run `python --version`.

###### Does it print `Python 3.x`? x >= 9

##### Do `import jupyter`, `import pandas` and `import spacy` run without any error?

---

---

#### Exercise 1: Embedding 1

###### For question 1, is the embedding's shape `(96,)`?

###### For question 2, do the 20 first values of the vector sum to `(99959115, 33554432)`?

---

---

#### Exercise 2: Tokenization

###### For question 1, are the tokens printed like the following?

    ```
    Tokenize
    this
    sentence
    .
    And
    this
    one
    too
    .
    ```

---

---

#### Exercise 3: Embeddings 2

##### The exercise is validated if all questions of the exercise are validated

###### For question 1, have the embeddings of each word a shape of `(300,)` and is the sum of the first 20 values of the embedding of laptop the following?

```
5.710388
```

###### For question 2, is the output the following?

![alt text][logo]

[logo]: ../w3day05ex1_plot.png "Plot"

---

---

#### Exercise 4: Sentences' similarity

###### For question 1, are the similarities between the sentences the following?

```
sentence_1 <=> sentence 2 : 0.7073220863266589
sentence_1 <=> sentence 3: 0.42663743263528325
sentence_2 <=> sentence 3: 0.3336274235605957
```

---

---

#### Exercise 5: NER

##### The exercise is validated if all questions of the exercise are validated

###### For question 1, is the ouptut of the NER the following?

    ```
    Apple Inc. ORG
    American NORP
    Cupertino GPE
    California GPE
    Five CARDINAL
    U.S. GPE
    Amazon ORG
    Google ORG
    Microsoft ORG
    Facebook ORG
    Apple ORG
    Steve Jobs PERSON
    Steve Wozniak PERSON
    Ronald Wayne PERSON
    April 1976 DATE
    Wozniak PERSON
    Apple ORG
    Wayne PERSON
    12 days DATE
    Apple Computer, Inc. ORG
    January 1977 DATE
    Apple ORG
    Apple II ORG
    ```

###### For question 2, does the output show that the first occurrence of apple is not a named entity? In my case here is what the NER returns:

    ```
    Paul 1 5 PERSON
    Apple 50 55 ORG

    ```

---

---

#### Exercise 6: Part-of-speech tags

###### For question 1, are the sentences outputed the following?

```
INFO:  Bezos PROPN NNP
Sentence:  Amazon (AMZN) enters 2021 with plenty of big opportunities, but is losing its lauded Chief Executive Jeff Bezos, who announced his plan to step aside in the third quarter.


INFO:  Bezos PROPN NNP
Sentence:  Bezos will hand off his role as chief executive to Andy Jassy, the CEO of its cloud computing unit.


INFO:  Bezos PROPN NNP
Sentence:  He's not leaving, as Bezos will transition to the role of Executive Chairman and remain active.


INFO:  Bezos PROPN NNP
Sentence:  "When you look at our financial results, what you're actually seeing are the long-run cumulative results of invention," Bezos said in written remarks with the Amazon earnings release.
```
