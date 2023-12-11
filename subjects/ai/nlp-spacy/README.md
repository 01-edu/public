# Natural Language processing with Spacy

`spaCy` is a natural language processing (NLP) library for Python designed to have fast performance, and with word embedding models built in, it’s perfect for a quick and easy start. I don't need to detail what spaCy does, it is perfectly summarized by spaCy in this article: **spaCy 101: Everything you need to know**.

Today, we will learn to use a pre-trained embedding to convert a text into a vector to compute similarity between words or sentences. Remember, embeddings translate large sparse vectors into a lower-dimensional space that preserves semantic relationships.
Word embeddings is a technique where individual words of a domain or language are represented as real-valued vectors in a lower dimensional space. The BoW representation's dimension depends on the size of the vocabulary. But it can easily reach 10k words. We will also learn to use NER and Part-of-speech. NER allows to identify and segment the named entities and classify or categorize them under various predefined classes. Part-of-speech is a special label assigned to each token (word) in a text corpus to indicate the part of speech and often also other grammatical categories such as tense, number (plural/singular), case etc.

### Exercises of the day

- Exercise 0: Environment and libraries
- Exercise 1: Embedding 1
- Exercise 2: Tokenization
- Exercise 3: Embeddings 2
- Exercise 4: Sentences' similarity
- Exercise 5: NER
- Exercise 6: Part-of-speech tags

### Virtual Environment

- Python 3.x
- Jupyter or JupyterLab
- Pandas
- spaCy
- Scikit-learn
- Matplotlib

I suggest using the most recent libraries.

### **Resources**

- https://spacy.io/usage/spacy-101
- https://spacy.io/api/doc
- https://www.analyticsvidhya.com/blog/2021/06/nlp-application-named-entity-recognition-ner-in-python-with-spacy/
- https://medium.com/mlearning-ai/nlp-04-part-of-speech-tagging-in-spacy-dc3e239c2726

---

---

# Exercise 0: Environment and libraries

The goal of this exercise is to set up the Python work environment with the required libraries.

**Note:** For each quest, your first exercise will be to set up the virtual environment with the required libraries.

I recommend to use:

- the **last stable versions** of Python.
- the virtual environment you're the most comfortable with. `virtualenv` and `conda` are the most used in Data Science.
- one of the most recent versions of the libraries required

1. Create a virtual environment named with a version of Python >= `3.8`, with the following libraries: `pandas`, `jupyter`, `spaCy 3.4.0`, `sklearn`, `matplotlib`.

---

---

# Exercise 1: Embedding 1

The goal of this exercise is to learn to load an embedding on `spaCy`.

1. Install and load `en_core_web_sm` version `3.4.1` [embedding](https://github.com/explosion/spacy-models/releases/tag/en_core_web_sm-3.4.1). Compute the embedding of `car`.

---

---

# Exercise 2: Tokenization

The goal of this exercise is to learn to tokenize a document using `spaCy`. We did this using NLTK yesterday.

1. Tokenize the text below and print the tokens

   ```
       text = "Tokenize this sentence. And this one too."

   ```

---

---

# Exercise 3: Embeddings 2

The goal of this exercise is to learn to use `spaCy` embedding on a document.

1. Compute the embedding of all the words in this sentence. The language model considered is `en_core_web_md` version 3.4.1.

```
"laptop computer coffee tea water liquid dog cat kitty"
```

2. Plot the pairwise cosine distances between all the words in a HeatMap.

![alt text][logo]

[logo]: ./w3day05ex1_plot.png "Plot"

https://medium.com/datadriveninvestor/cosine-similarity-cosine-distance-6571387f9bf8

---

---

# Exercise 4: Sentences' similarity

The goal of this exerice is to learn to compute the similarity between two sentences. As explained in the documentation: **The word embedding of a full sentence is simply the average over all different words**. This is how `similarity` works in SpaCy. This small use case is very interesting because if we build a corpus of sentences that express an intention as **buy shoes**, then we can detect this intention and use it to propose shoes advertisement for customers. The language model used in this exercise is `en_core_web_sm`.

1. Compute the similarities (3 in total) between these sentences:

   ```
   sentence_1 = "I want to buy shoes"
   sentence_2 = "I would love to purchase running shoes"
   sentence_3 = "I am in my room"

   ```

---

---

# Exercise 5: NER

The goal of this exercise is to learn to use a Named entity recognition algorithm to detect entities.

```
Apple Inc. is an American multinational technology company headquartered in Cupertino, California, that designs, develops, and sells consumer electronics, computer software, and online services. It is considered one of the Big Five companies in the U.S. information technology industry, along with Amazon, Google, Microsoft, and Facebook.
Apple was founded by Steve Jobs, Steve Wozniak, and Ronald Wayne in April 1976 to develop and sell Wozniak's Apple I personal computer, though Wayne sold his share back within 12 days. It was incorporated as Apple Computer, Inc., in January 1977, and sales of its computers, including the Apple I and Apple II, grew quickly.
```

1. Extract all named entities in the text as well as the label of the named entity.

2. The NER is also useful to remove ambiguous entities. From a conceptual standpoint, disambiguation is the process of determining the most probable meaning of a specific phrase. For example in the sentence below, the word `apple` is present twice in the sentence. The first time to mention the fruit and the second to mention a company. Run the NER on this sentence and print the named entity, the `start_char`, the `end_char` and the label of the named entity.

```
Paul eats an apple while watching a movie on his Apple device.
```

https://en.wikipedia.org/wiki/Named-entity_recognition

---

---

# Exercise 6: Part-of-speech tags

The goal of this exercise is to learn to use the Part-of-speech tags (**POS TAG**) using `spaCy`. As explained on Wikipedia, the POS TAG is the process of marking up a word in a text (corpus) as corresponding to a particular part of speech, based on both its definition and its context.

Example

The sentence: **"Heat water in a large vessel"** is tagged this way after the POS TAG:

- heat verb (noun)
- water noun (verb)
- in prep (noun, adv)
- a det (noun)
- large adj (noun)
- vessel noun

The data `news_amazon.txt` used is a newspaper about Amazon.

1. Return all sentences mentioning **Bezos** as a NNP (tag).
