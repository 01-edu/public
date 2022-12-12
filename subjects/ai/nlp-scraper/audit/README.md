#### NLP-enriched News Intelligence platform

##### Preliminary

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

###### Does the structure of the project is as below ?

###### Does the readme file give an introduction of the project, show the username, describe the feature engineering and show the best score the on the leaderboard ?

###### Does the environment contain all libraries used and their versions that are necessary to run the code ?

##### Scrapper

##### There are at least 300 news articles stored in the file system or the database.

##### Run the scrapper with `python scrapper_news.py` and fetch 3 documents. The scrapper is not expected to fetch 3 documents and stop by itself, you can stop it manually. It runs without any error and stores the 3 files as expected.

##### Topic classifier

###### Are the learning curves provided ?

###### Do the learning curves prove the topics classifier is trained without correctly - without overfitting ?

###### Can you run the topic classfier model on the test set without any error ?

###### Does the topic classifier score an accuracy higher than 95% ?

##### Scandal detection

###### Does the `README.md` explain the choice of embeddings and distance ?

###### Does the DataFrame flag the top 10 articles with the highest likelihood of environmental scandal ?

###### Is the distance or similarity saved in the DataFrame ?

#####

##### NLP engine output on 300 articles

###### Does the DataFrame contain 300 different rows ?

###### Does the columns of the DataFrame are as expected ?

```
Date scrapped (date)
Title (str)
URL (str)
Body (str)
Org (str)
Topics (list str)
Sentiment (list float or float)
Scandal_distance (float)
Top_10 (bool)

```

##### Analyse the DataFrame with 300 articles: relevance of the topics matched, relevance of the sentiment, relevance of the scandal detected and relevance of the companies matched. The algorithms are not 100% accurate so you should expect a few issues in the results.

##### NLP engine on 3 articles

###### Can you run `python nlp_enriched_news.py` without any error ?

###### Does the output of the nlp engine correspond to the output below?

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

##### Analyse the output: relevance of the topic(s) matched, relevance of the sentiment, relevance of the scandal detected (if detected on the three articles) and relevance of the companie(s) matched.
