# import json
# import boto3
# import zipfile
# import os
# import pickle



# def lambda_handler(event, context):
#     s3_client = boto3.client('s3')
#     s3_client.download_file('sagemaker-tensormodelmasters','pac.pkl', '/tmp/pac.pkl')
#     s3_client.download_file('sagemaker-tensormodelmasters','vect.pkl', '/tmp/vect.pkl')

#     fake_news_2 = 'THIS IS FAKE FAKE FAKE FAKE NEWS'

#     pacz = pickle.load(open("/tmp/pac.pkl", "rb"))
#     vect = pickle.load(open("/tmp/vect.pkl", "rb"))


#     input = fake_news_2.lower()
#     text= vect.transform([input])

#     print(pacz.predict(text))

#     # return{
#     #     'statusCode': 200,
#     #     'body': json.loads(text)
#     # }


import boto3
import json
import re
import string
import zipfile
import os
import pickle



# ======================== Calculation ========================
def str2bool(v):
  return v.lower() in ("1")

def calculate_result_static(bert_output, tfidf_output):
    bert_result = bert_output[0]
    bert_confidence = float(bert_output[1])

    if tfidf_output == bert_result:
        print("here")
        #when the two classifiers' result match

        final_result = (bert_confidence * 70) + (bert_confidence *30)
        return final_result
    else:
        return 0

    return 0

def weight_multiplier_diff(weight, multiplier):
    b_weight = weight * multiplier
    t_weight = abs(100-(b_weight*multiplier))
    if multiplier < 0.8:
        t_weight = abs(100-b_weight) * multiplier
    return b_weight, t_weight


def calculate_bert_conf_diff(b_weight, t_weight, conf):
    return b_weight * conf + (t_weight / 2)

def confidence_weight_diff(input_confidence):
    bert_weight = 80
    tfidf_weight = 20

    if input_confidence > .98:
        return weight_multiplier_diff(bert_weight, 1.15)
    elif input_confidence > .93:
        return weight_multiplier_diff(bert_weight, 1.07)
    elif input_confidence > .85:
        return weight_multiplier_diff(bert_weight, 1)
    elif input_confidence > .75:
        return weight_multiplier_diff(bert_weight, .94)
    elif input_confidence > .65:
        return weight_multiplier_diff(bert_weight, .8)
    elif input_confidence > .40:
        return weight_multiplier_diff(bert_weight, .65)
    else:
        return weight_multiplier_diff(bert_weight, .5)


def weight_multiplier_same(weight, multiplier):
    b_weight = weight * multiplier
    t_weight = abs(100-(b_weight))
    if multiplier < 0.8:
        t_weight = abs(100-b_weight) * multiplier
    return b_weight, t_weight

def calculate_bert_conf_same(b_weight, t_weight, conf):
    return b_weight * conf + t_weight

def confidence_weight_same(input_confidence):
    # when bert and tfidf results the same conclusion
    bert_weight = 80
    tfidf_weight = 20

    if input_confidence > .95:
        return weight_multiplier_same(bert_weight, 1)
    elif input_confidence > .9:
        return weight_multiplier_same(bert_weight, .95)
    elif input_confidence > .85:
        return weight_multiplier_same(bert_weight, .87)
    elif input_confidence > .75:
        return weight_multiplier_same(bert_weight, .8)
    elif input_confidence > .65:
        return weight_multiplier_same(bert_weight, .77)
    elif input_confidence > .40:
        return weight_multiplier_same(bert_weight, .7)
    else:
        return weight_multiplier_same(bert_weight, .6)


def calculate_result_dynamic(bert_output, tfidf_output):
    # b_output = bert_output.split(" ")

    bert_result = bert_output[0]
    bert_confidence = float(bert_output[1])

    determiniation = 1
    conf_result = 0.0

    if tfidf_output == bert_result:
        #when the two classifiers' result match
        bert_weight, tfidf_weight = confidence_weight_same(bert_confidence)
        conf_result = calculate_bert_conf_same(bert_weight, tfidf_weight, bert_confidence)
    else:
        bert_weight, tfidf_weight = confidence_weight_diff(bert_confidence)
        conf_result = calculate_bert_conf_diff(bert_weight, tfidf_weight, bert_confidence)

    if bert_confidence < .3:
        determiniation = tfidf_output
    else:
        determiniation = bert_result
    return determiniation, conf_result
# ======================== Calculation ========================



# ======================== BERT AND ROBERTA ========================
def get_result(bert_result, roberta_result):
    f_value = bert_result[0] + roberta_result[0]
    t_value = bert_result[1] + roberta_result[1]

    if f_value > t_value:
        print("The statement is False!")
        result = float(f_value)/11
        return (0, result)

    else:
        print("The statement is True")
        result = float(t_value)/11
        return (1, result)

def parse_output(input):
    input_values = re.findall(r"[-+]?\d*\.\d+|\d+", input)
    result_tuple = (float(input_values[0]), float(input_values[1]))
    return result_tuple


def get_bert_roberta_result(input_string):
    sagemaker_client = boto3.client('sagemaker-runtime')

    bert_response = sagemaker_client.invoke_endpoint(EndpointName="bert-project9",
                              Body=input_string.encode(encoding='UTF-8'),
                              ContentType='text/csv')

    bert_r = bert_response['Body'].read().decode("utf-8")

    bert_result = parse_output(bert_r)

    roberta_response = sagemaker_client.invoke_endpoint(EndpointName="roberta-project1",
                              Body=input_string.encode(encoding='UTF-8'),
                              ContentType='text/csv')

    roberta_r = roberta_response['Body'].read().decode("utf-8")

    roberta_result = parse_output(roberta_r) #(.001, .001)

    return  get_result(bert_result, roberta_result)
# ======================== BERT AND ROBERTA ========================



# ======================== TF-IDF ========================
def tfidf_prediction(input_text):


    pac = pickle.load(open("/tmp/pac.pkl", "rb"))
    vect = pickle.load(open("/tmp/vect.pkl", "rb"))


    input = input_text.lower()
    text= vect.transform([input])

    prediction = pac.predict(text)
    return str(prediction[0])

def get_tf_idf_prediction(input_text):
    s3_client = boto3.client('s3')
    s3_client.download_file('sagemaker-tensormodelmasters','pac.pkl', '/tmp/pac.pkl')
    s3_client.download_file('sagemaker-tensormodelmasters','vect.pkl', '/tmp/vect.pkl')
    return tfidf_prediction(input_text)

# ======================== TF-IDF ========================






## primary function
def lambda_handler(event, context):
    # # lambda receives the input from the web app as an event in json format
    # sentences = event['body']
    # tokenized_sentences = review_to_words(sentences)
    # payload = {"instances" : [tokenized_sentences]}

    # # The SageMaker runtime is what allows us to invoke the endpoint that we've created.
    # runtime = boto3.Session().client('sagemaker-runtime')

    # # Now we use the SageMaker runtime to invoke our endpoint, sending the text we were given
    # response = runtime.invoke_endpoint(EndpointName = 'blazingtext-2020-08-08-09-19-31-404',# The name of the endpoint we created
    #                                   ContentType = 'application/json',                 # The data format that is expected by BlazingText
    #                                   Body = json.dumps(payload))

    # # The response is an HTTP response whose body contains the result of our inference
    # output = json.loads(response['Body'].read().decode('utf-8'))
    # prob = output[0]['prob'][0]*100
    # label = output[0]['label'][0].split('__label__')[1]
    # output = 'The predicted label is {} with a probability of {:.1f}%'.format(label, prob)


    input_text = "The best part of Amazon SageMaker is that it makes machine learning easy"
    b_r_result = get_bert_roberta_result(input_text)
    tf_result = get_tf_idf_prediction(input_text)


    article_determiniation, confidence_score = calculate_result_dynamic(b_r_result, tf_result)

    result = {
      "text_determination": article_determiniation,
      "confidence_score": confidence_score,
    }
    # s3_client = boto3.client('s3')
    # s3_client.download_file('sagemaker-tensormodelmasters','pac.pkl', '/tmp/pac.pkl')
    # s3_client.download_file('sagemaker-tensormodelmasters','vect.pkl', '/tmp/vect.pkl')

    # fake_news_2 = 'THIS IS FAKE FAKE FAKE FAKE NEWS'

    # pacz = pickle.load(open("/tmp/pac.pkl", "rb"))
    # vect = pickle.load(open("/tmp/vect.pkl", "rb"))


    # input = fake_news_2.lower()
    # text= vect.transform([input])

    # print(pacz.predict(text))

    # we return the output in a format expected by API Gateway
    return {
        'statusCode' : 200,
        'headers' : { 'Content-Type' : 'text/plain', 'Access-Control-Allow-Origin' : '*' },
        'body' : result
    }
