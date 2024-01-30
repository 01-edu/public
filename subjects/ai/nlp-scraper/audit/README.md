#### NLP-enriched News Intelligence platform

##### Preliminary

###### Does the structure of the project look like the one described in the subject?

###### Does the environment contain all libraries used and their versions that are necessary to run the code?

##### Scraper

##### There are at least 300 news articles stored in the file system or the database.

##### Run the scraper with `python scraper_news.py` and fetch 3 documents. The scraper is not expected to fetch 3 documents and stop by itself, you can stop it manually. 

###### Does it run without any error and store the 3 files as expected?

##### Topic classifier

###### Are the learning curves provided?

###### Do the learning curves prove the topics classifier is trained correctly - without overfitting?

###### Can you run the topic classifier model on the test set without any error?

###### Does the topic classifier score an accuracy higher than 95%?

##### Scandal detection

###### Does the `README.md` explain the choice of embeddings and distance?

###### Does the DataFrame flag the top 10 articles with the highest likelihood of environmental scandal?

###### Is the distance or similarity saved in the DataFrame?

##### NLP engine output on 300 articles

###### Does the DataFrame contain 300 different rows?

###### Are the columns of the DataFrame as defined in the subject `Deliverable` section?

##### Analyse the DataFrame with 300 articles: relevance of the topics matched, relevance of the sentiment, relevance of the scandal detected and relevance of the companies matched. The algorithms are not 100% accurate, so you should expect a few issues in the results.

##### NLP engine on 3 articles

###### Can you run `python nlp_enriched_news.py` without any error?

###### Does the output of the NLP engine correspond to the output defined in the subject `Deliverable` section?

##### Analyse the output: relevance of the topic(s) matched, relevance of the sentiment, relevance of the scandal detected (if detected on the three articles) and relevance of the company(ies) matched.
