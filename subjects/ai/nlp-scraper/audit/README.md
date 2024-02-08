#### NLP-enriched News Intelligence platform

##### Preliminary

###### Does the structure of the project look like the one described in the subject?

###### Does the environment contain all libraries used and their versions that are necessary to run the code?

##### Scraper

##### Run the scraper with `python scraper_news.py` and fetch 300 articles. If needed, stop the program manually when enough data has been retrieved.

###### Does it run without any error and store the articles as described in the subject?

##### Topic classifier

###### Are the learning curves provided?

###### Do the learning curves prove the topics classifier is trained correctly - without overfitting? Ask the student to explain what the term "overfitting" means and how he avoided this phenomenon.

> Additionally, you can look for external resources. For example, Wikipedia has a good page on "overfitting".

##### Ask the student to train and store the topic classifier model in a file named `topic_classifier.pkl`.

###### Can you run the topic classifier model on the test set without any error?

###### Does the topic classifier score an accuracy higher than 95% on the given datasets?

##### NLP engine output on 300 articles

###### Can you run `python nlp_enriched_news.py` without any error?

###### Does the DataFrame saved in the `csv` file contain 300 different rows?

###### Are the columns of the DataFrame as defined in the subject `Deliverable` section?

###### Does the output of the NLP engine correspond to the output defined in the subject `Deliverable` section?

##### Analyse the output: relevance of the topic(s) matched, relevance of the sentiment, relevance of the scandal detected (if detected on the three articles) and relevance of the company(ies) matched.

###### Is the information presented consistent and accurate?

##### Scandal detection

###### Does the `README.md` explain the choice of embeddings and distance?

###### Does the DataFrame flag the top 10 articles with the highest likelihood of environmental scandal?

###### Is the distance or similarity saved in the DataFrame?
