from newspaper import Article 
import pandas as pd
import json
import csv

data = pd.read_csv('FakeNews1.csv')

file = open('FakeNews1Scraped.csv', 'w', encoding='utf-8')
keywords = ["coronavirus", "pandemic", "covid19", "covid-19", "lockdown"]
csv_file = csv.DictWriter(file, fieldnames=["title", "text", "publish_date"])
csv_file.writeheader()
for x in data['links']:
	article = Article(x)
	article.download()
	article.parse()

	text={}
	text['title'] = article.title
	text['body'] = article.text
	text['date'] = article.publish_date
	for keyword in keywords:
		if keyword in article.text:
			csv_file.writerow({"title": article.title, "text": article.text, "publish_date": article.publish_date})

	print(article.title)
