from transformers import pipeline
import csv
classifier = pipeline("sentiment-analysis")

count = 0

def read_in_articles(filename, max):
    print("read in articles")
    articles = []

    count = 0
    pos = 0
    with open(filename, 'r') as f:
        reader = csv.reader(f)
        for row in reader :
            # articles.append({'title': row[0], 'content': row[1], 'writte_date': row[2], 'article_sentiment_score': row[3]})
            articles.append({'title': row[0], 'content': row[1], 'writte_date': row[2]})
            count += 1
                # if count > max:
                #     break
    return articles


def remove_new_space(articles):
    print("remove new spaces")
    pos = 0
    for article in articles:
        articles[pos]['content'] =  articles[pos]['content'].strip('\n')
        articles[pos]['content'] =  articles[pos]['content'].strip('\t')
        articles[pos]['content'] =  articles[pos]['content'].replace('\n', '')
        articles[pos]['content'] =  articles[pos]['content'].replace('\t', '')
        pos += 1

def getRatingWeightedScore(rating):
    plain = round(rating,4)*100
    if plain > 97:
        return rating*3
    elif plain > 92:
        return rating*2
    elif plain > 85:
        return rating
    else:
        return rating * .8

def get_sentiment_score(articles):
    print("get sentiment score")
    tot = 0
    for article in articles:
        pos = 0.0
        neg = 0.0
        sentences = article['content'].split('.')
        for sentence in sentences:

            if len(sentence.split()) > 200:
                continue
            result = classifier(sentence)
            label = result[0]['label']
            rating = result[0]['score']
            weightedRating = getRatingWeightedScore(rating)
            # weightedRating = rating
            if label == 'NEGATIVE':
                # neg += weightedRating
                tot += (weightedRating * -100)
            if label == 'POSITIVE':
                # pos += weightedRating
                tot += (weightedRating * 100)

        article['article_sentiment_score'] = (tot/len(sentences))
        tot = 0
        # article['positive'] = pos
        # article['negative'] = neg


def write_to_csv(article, fileName):
    print("write to csv")

    # fields = ['title', 'content', 'writte_date', 'positive', 'negative']
    fields = ['title', 'content', 'writte_date', 'article_sentiment_score']

    with open(fileName, 'w') as csvFile:
        writer = csv.DictWriter(csvFile, fieldnames=fields)
        writer.writeheader()
        writer.writerows(article)





def parse_articles():
    # article = '/Users/ysun2/masterDegree/masterProject/sentiment/sentimentAssign/v2-correctedarticles/FakeNewsScraped.csv'
    # condenced = '/Users/ysun2/masterDegree/masterProject/sentiment/sentimentAssign/v2-correctedarticles/fake-Weighted.csv'

    article = '/Users/ysun2/masterDegree/masterProject/sentiment/sentimentAssign/v2-correctedarticles/FactualArticlesScrapedV2.csv'
    condenced = '/Users/ysun2/masterDegree/masterProject/sentiment/sentimentAssign/v2-correctedarticles/real-Weighted.csv'
    articles = read_in_articles(article , 50)
    remove_new_space(articles)
    get_sentiment_score(articles)
    write_to_csv(articles, condenced)

parse_articles()

# tot = 0
# fakeCondenced = '/Users/ysun2/masterDegree/masterProject/sentiment/sentimentAssign/testFake.csv'
# realCondenced = '/Users/ysun2/masterDegree/masterProject/sentiment/sentimentAssign/v2-correctedarticles/testReal.csv'
# articles = read_in_articles(fakeCondenced, 0)
#
# for article in articles:
#     if article['article_sentiment_score'] == 'article_sentiment_score':
#         continue
#     s = float(article['article_sentiment_score'])
#     tot += s
#
# tot /= len(articles)
# print('fake avg: '+str(tot))
#
# tot = 0.0
# articles = read_in_articles(realCondenced, 0)
# for article in articles:
#     if article['article_sentiment_score'] == 'article_sentiment_score':
#         continue
#     tot += float(article['article_sentiment_score'])
#
# tot -= 99
# tot /= len(articles)
# print('real avg: '+str(tot))
