## NLP

### Overview

“NLP makes it possible for humans to talk to machines:” This branch of AI enables computers to understand, interpret, and manipulate human language. This technology is one of the most broadly applied areas of machine learning and is critical in effectively analyzing massive quantities of unstructured, text-heavy data.

### Role Play

You're a Natural Language Processing (NLP) specialist at a tech startup developing a sentiment analysis tool for social media posts. Your task is to build the preprocessing pipeline and create a bag-of-words representation for tweet analysis.

### Learning Objectives

Machine learning algorithms cannot work with raw text directly. Rather, the text must be converted into vectors of numbers. In natural language processing, a common technique for extracting features from text is to place all of the words that occur in the text in an unordered bucket. This approach is called a bag of words model or BoW for short.

It’s referred to as a “bag” of words because any information about the structure of the sentence is lost. This is useful to train usual machine learning models on text data. Other types of models as RNNs or LSTMs take as input a complete and ordered sequence.

Almost every Natural Language Processing (NLP) task requires text to be preprocessed before training a model. The article **Your Guide to Natural Language Processing (NLP)** gives a very good introduction to NLP.

Today, we we will learn to preprocess text data and to create a bag of word representation. The packages NLTK and Spacy to do the preprocessing

### Exercises of the day

- **Exercise 0:** Environment and libraries
- **Exercise 1:** Lower case
- **Exercise 2:** Punctuation
- **Exercise 3:** Tokenization
- **Exercise 4:** Stop words
- **Exercise 5:** Stemming
- **Exercise 6:** Text preprocessing
- **Exercise 7:** Bag of Word representation

### Virtual Environment

- Python 3.x
- Jupyter or JupyterLab
- Pandas
- Scikit-learn
- NLTK
- Tabulate

> We suggest to use the most recent libraries.

---

---

### Exercise 0: Environment and libraries

The goal of this exercise is to set up the Python work environment with the required libraries.

> Note: For each quest, your first exercise will be to set up the virtual environment with the required libraries.

We recommend using:

- The **last stable versions** of Python.
- The virtual environment you're the most comfortable with. `virtualenv` and `conda` are the most used in Data Science.
- one of the most recent versions of the libraries required.

1. Create a virtual environment named with a version of Python >= `3.9`, with the following libraries: `pandas`, `jupyter`, `nltk` and `scikit-learn`.

---

---

### Exercise 1: Lowercase

The goal of this exercise is to learn to lowercase text data in Python. Note that if the volume of data is low the text data can be stored in a Pandas DataFrame or Series. But, when dealing with high volumes (high but not huge), using a Pandas DataFrame or Series is not efficient. Data structures as dictionaries or list are more adapted.

```
list_ = ["This is my first NLP exercise", "wtf!!!!!"]
series_data = pd.Series(list_, name='text')

```

1. Print all texts in lowercase
2. Print all texts in upper case

> Note: Do not change the text manually!

---

---

### Exercise 2: Punctuation

The goal of this exercise is to learn to deal with punctuation. In Natural Language Processing, some basic approaches as Bag of Words model the text as an unordered combination of words. In that case the punctuation is not always useful as it doesn't add information to the model. That is why is removed.

1. Remove the punctuation from this sentence. All characters in !"#$%&'()\*+,-./:;<=>?@[\]^\_`{|}~ are considered as punctuation.

   ```

   Remove, this from .? the sentence !!!! !"#&'()*+,-./:;<=>_

   ```

---

---

### Exercise 3: Tokenization

The goal of this exercise is to learn [to tokenize](https://en.wikipedia.org/wiki/Lexical_analysis#Tokenization) as text. This step is important because it splits the text into token. A token could be a sentence or a word.

```
text = """Bitcoin is a cryptocurrency invented in 2008 by an unknown person or group of people using the name Satoshi Nakamoto. The currency began use in 2009 when its implementation was released as open-source software."""

```

1. Tokenize this text using `sent_tokenize` from NLTK.

2. Tokenize this text using `word_tokenize` from NLTK.

---

---

### Exercise 4: Stop words

The goal of this exercise is to learn to remove stop words with NLTK. Stop words usually refers to the most common words in a language. For example: "and", "is", "a" are stop words and do not add information to a sentence.

```
text = """
The goal of this exercise is to learn to remove stop words with NLTK.  Stop words usually refers to the most common words in a language.
"""
```

1. Remove stop words from this sentence and return the list of work tokens without stop words.

---

---

### Exercise 5: Stemming

The goal of this exercise is to learn to use stemming using NLTK. As explained in details in the article, stemming is the process of reducing inflection in words to their root forms such as mapping a group of words to the same stem even if the stem itself is not a valid word in the Language.

> Note: The output of a stemmer is a word that may not exist in the dictionary.

```
text = """
The interviewer interviews the president in an interview
"""
```

1. Return the list of the stemmed tokens.

---

---

### Exercise 6: Text preprocessing

The goal of this exercise is to learn to create a function to prepocess and clean a text using NLTK.

Put this text in a variable:

```
01 Edu System presents an innovative curriculum in software engineering and programming. With a renowned industry-leading reputation, the curriculum has been rigorously designed for learning skills of the digital world and technology industry. Taking a different approach than the classic teaching methods today, learning is facilitated through a collective and co-creative process in a professional environment.

```

1. Write a function that takes as input the text and returns it preprocessed.

The preprocessing is composed of:

    1. Lowercase
    2. Removing Punctuation
    3. Tokenization
    4. Stopword Filtering
    5. Stemming

_ [Resources](https://towardsdatascience.com/nlp-preprocessing-with-nltk-3c04ee00edc0_)

---

---

### Exercise 7: Bag of Word representation

The goal of this exercise is to understand the creation of a Bag of Word (BoW) model for a corpus of texts and create a labeled dataset from textual data using a word count matrix.

_Resources: [Gentle Introduction to Bag of Words Model](https://machinelearningmastery.com/gentle-introduction-bag-words-model/)_

The Bag of Word representation assumes that word order in a text is irrelevant. There are various types of Bag of Words representations:

- Boolean: Each document represented by a boolean vector
- Word count: Each document represented by a word count vector
- TFIDF: Each document represented by a score vector (more detailed in the next exercise)

The data file [tweets_train.txt](resources/tweets_train.txt) contains labeled tweets indicating their positivity.

Steps:

1. Preprocess the data using the function implemented earlier. Then, use `CountVectorizer` from scikit-learn with `max_features=500` to compute the word count of the tweets. The output is a sparse matrix.

   - Check the shape of the word count matrix.
   - Set **max_features** to 500, considering the initial size of the dictionary.

> Note: Given that a data set is often described as an m x n matrix in which m is the number of rows and n is the number of columns: features. It is strongly recommended to work with m >> n. The value of the ratio depends on the signal existing in the data set and on the model complexity.

2. Using `from_spmatrix` from Pandas, create a DataFrame `count_vecotrized_df` using the output features names as column names. The final results should be similar to the below one.

|     | and | boat | compute |
| --: | --: | ---: | ------: |
|   0 |   0 |    2 |       0 |
|   1 |   0 |    0 |       1 |
|   2 |   1 |    0 |       0 |

> Note: The sample 3x3 table mentioned is a small representation of the expected output for demonstration purposes. It's not necessary to drop columns in this context.

3. Show the token counts (obtained with the above-mentioned steps) of the fourth tweet.

4. Using the word counter, show the 15 most used tokenized words in the datasets' tweets

5. Add to your `count_vecotrized_df` a `label` column considering the following:

   - 1: Positive
   - 0: Neutral
   - -1: Negative

   The final DataFrame should be similar to the below:

|     | ... | label |
| --: | --: | ----: |
|   0 | ... |     1 |
|   1 | ... |    -1 |
|   2 | ... |    -1 |
|   3 | ... |    -1 |

### Resources

- [Nlp 1](https://towardsdatascience.com/your-guide-to-natural-language-processing-nlp-48ea2511f6e1)

- [Nlp 2](https://towardsdatascience.com/word-embeddings-for-nlp-5b72991e01d4)

- [Sklearn.feature_extraction.text.CountVectorizer](https://scikit-learn.org/stable/modules/generated/sklearn.feature_extraction.text.CountVectorizer.html)
