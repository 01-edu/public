#### Exercise 0: Environment and libraries

##### The exercise is validated is all questions of the exercise are validated.

##### Activate the virtual environment. If you used `conda` run `conda activate your_env`.

##### Run `python --version`.

###### Does it print `Python 3.x`? x >= 8

##### Do `import jupyter`, `import pandas` and `import spacy` run without any error?

---

---

#### Exercise 1: Embedding 1

##### The question 1 is validated if the embedding's shape is `(96,)`

##### The question 1 is validated if the 20 first values of the vector sum to `2.9790137708187103`

---

---

#### Exercise 2: Tokenization

##### The question 1 is validated if the tokens printed are:

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

##### The exercice is validated is all questions of the exercice are validated

##### The question 1 is validated if the embeddings of each word has a shape of `(300,)` and if the first 20 values of the embedding of laptop are:

```
array([-0.37639 , -0.075521,  0.4908  ,  0.19863 , -0.11088 , -0.076145,
    -0.30367 , -0.69663 ,  0.87048 ,  0.54388 ,  0.42523 ,  0.18045 ,
    -0.4358  , -0.32606 , -0.70702 , -0.069127, -0.42674 ,  2.4147  ,
        0.26806 ,  0.46584 ], dtype=float32)

```

##### The question 2 is validated if the output is

![alt text][logo]

[logo]: ../w3day05ex1_plot.png "Plot"

---

---

#### Exercise 4: Sentences' similarity

##### The question 1 is validated if the similarities between the sentences are:

```
sentence_1 <=> sentence 2 : 0.7073220863266589
sentence_1 <=> sentence 3: 0.42663743263528325
sentence_2 <=> sentence 3: 0.3336274235605957
```

---

---

#### Exercise 5: NER

##### The exercice is validated is all questions of the exercice are validated

##### The question 1 is validated if the ouptut of the NER is

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

##### The question 2 is validated if the output shows that the first occurence of apple is not a named entity. In my case here is what the NER returns:

    ```
    Paul 1 5 PERSON
    Apple 50 55 ORG

    ```

---

---

#### Exercise 6: Part-of-speech tags

##### The question 1 is validated if the sentences outputed are:

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
