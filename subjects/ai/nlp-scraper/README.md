# NLP-enriched News Intelligence platform

The goal of this project is to build an NLP-enriched News Intelligence platform. News analysis is a trending and important topic. The analysts get their information from the news and the amount of available information is limitless. Having a platform that helps to detect the relevant information is definitely valuable.

The platform connects to a news data source, detects the entities, detects the topic of the article, analyse the sentiment and ...

### Scrapper

News data source:

- Find a news website that is easy to scrap. I could have chosen the website but the news' websites change their scraping policy frequently.
- Store it:

  - File system per day:
    - URL, date unique id
    - headline
    - body of the article
  - SQL database (optional)

  from the last week otherwise the volume may be to high

### NLP engine

In production architectures, the NLP engine delivers a live output based on the news that are delivered in a live stream data by the scrapper. However, it required advanced Python skills that are not a pre-requisite for the AI branch.
To simplify this step the scrapper and the NLP engine are used independently in the project. The scrapper fetches the news and store them in the data structure (either the file systeme or the SQL database) and then, the NLP engine runs on the stored data.

Here how the NLP engine should process the news:

### **1. Entities detection:**

The goal is to detect all the entities in the document (headline and body). The type of entity we focus on is `ORG`. This corresponds to companies and organisations. This information should be stored.

    - Detect all companies using SpaCy NER on the body of the text.

https://towardsdatascience.com/named-entity-recognition-with-nltk-and-spacy-8c4a7d88e7da

### **2. Topic detection:**

The goal is to detect what the article is dealing with: Tech, Sport, Business, Entertainment or Politics. To do so, a labelled dataset is provided. From this dataset, build a classifier that learns to detect the right topic in the article. The trained model should be stored as `topic_classifier.pkl`. Make sure the model can be used easily (with the preprocessing pipeline built for instance) because the audit requires the auditor to test the model.

Save the plot of learning curves (`learning_curves.png`) in `results` to prove that the model is trained correctly and not overfitted.

- Learning constraints: **Score on test: > 95%**

- **Optional**: If you want to train a news' topic classifier based on a more challenging dataset, you can use the following which based on 200k news headlines. https://www.kaggle.com/rmisra/news-category-dataset.

### **3. Sentiment analysis:**

The goal is to detect the sentiment of the news articles. To do so, use a pre-trained sentiment model. I suggest to use: NLTK.
There are 3 reasons for which we use a pre-trained model:

1. As a Data Scientist, you should learn to use a pre-trained model. There are so many models available and trained that sometimes you don't need to train one from scratch.
2. Labelled news data for sentiment analysis are very expensive. Companies as SESAMm provide this kind of services.
3. You already know how to train a sentiment analysis classifier ;-)

### **4. Scandal detection **

The goal is to detect environmental disaster for the detected companies. Here is the methodology that should be used:

- Define keywords that correspond to environmental disaster that may be caused by companies: pollution, deforestation etc ... Here is an example of disaster we want to detect: https://en.wikipedia.org/wiki/MV_Erika. Pay attention to not use ambigous words that make sense in the context of an environmental disaster but also in another context. This would lead to detect a false positive natural disaster.

- Compute the embeddings of the keywords.

- Compute the distance between the embeddings of the keywords and all sentences that contain an entity. Explain in the `README.md` the embeddings chosen and why. Similarly explain the distance or similarity chosen and why.

- Save the distance

- Flag the top 10 articles.

### 5. **Source analysis (optional)**

The goal is to show insights about the news' source you scrapped.
This requires to scrap data on at least 5 days (a week ideally). Save the plots in the `results` folder.

Here are examples of insights:

- Per day:

  - Proportion of topics per day
  - Number of articles
  - Number of companies mentioned
  - Sentiment per day

- Per companies:

  - Companies mentioned the most
  - Sentiment per companies

### Deliverables

The structure of the project is:

```
project
│   README.md
│   environment.yml
│
└───data
│   │   topic_classification_data.csv
│
└───results
│   │   topic_classifier.pkl
│   │   learning_curves.png
│   │   enhanced_news.csv
|
|───nlp_engine
│

```

1.  Run the scrapper until it fetches at least 300 articles

```
python scrapper_news.py

1. scrapping <URL>
        requesting ...
        parsing ...
        saved in <path>

2. scrapping <URL>
        requesting ...
        parsing ...
        saved in <path>

```

2. Run on this 300 articles the NLP engine.

Save a DataFrame:

Date scrapped (date)
Title (str)
URL (str)
Body (str)
Org (str)
Topics (list str)
Sentiment (list float or float)
Scandal_distance (float)
Top_10 (bool)

```prompt
python nlp_enriched_news.py

Enriching <URL>:

Cleaning document ... (optional)

---------- Detect entities ----------

Detected <X> companies which are <company_1> and <company_2>

---------- Topic detection ----------

Text preprocessing ...

The topic of the article is: <topic>

---------- Sentiment analysis ----------

Text preprocessing ... (optional)
The title which is <title> is <sentiment>
The body of the article is <sentiment>

---------- Scandal detection ----------

Computing embeddings and distance ...

Environmental scandal detected for <entity>
```

I strongly suggest to create a data structure (dictionary for example) to save all the intermediate result. Then, a boolean argument `cache` fetched the intermediate results when they are already computed.

Ressources:

- https://www.youtube.com/watch?v=XVv6mJpFOb0
