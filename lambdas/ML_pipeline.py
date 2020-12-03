import string
import json
import boto3
import zipfile
import os
import pickle




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
    b_output = bert_output.split(" ")

    bert_result = b_output[0]
    bert_confidence = float(b_output[1])

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

def tfidf_prediction(inputText):


    pac = pickle.load(open("/tmp/pac.pkl", "rb"))
    vect = pickle.load(open("/tmp/vect.pkl", "rb"))


    input = inputText.lower()
    text= vect.transform([input])

    prediction = pac.predict(text)
    return str(prediction[1])


def lambda_handler(event, context):
    # real = "Summary of Changes April 6, 2020: Updated to align with the level of care categories defined in the Federal Healthcare Resilience Task Force Alternate Care Sites (ACS) Toolkit: First Edition, which included General (non-acute) Care and Acute Care April 23, 2020: Updated to align with the level of care categories defined in the Federal Healthcare Resilience Task Force Alternate Care Sites (ACS) Toolkit pdf icon external icon : Second Edition, which include Non-Acute Care, Hospital Care, and Acute Care Key Concepts Establishing alternate care sites will help address surge in the response to COVID-19.Since care will be provided in a non-traditional environment, it is critical to ensure these facilities can support implementation of recommended infection prevention and control practices.A local surge in the need for medical care might require jurisdictions to establish alternate care sites (ACS) where patients with COVID-19 can remain and receive medical care for the duration of their isolation period. These are typically established in non-traditional environments, such as converted hotels or mobile field medical units.1 Depending on the jurisdictional needs, ACS could provide three levels of care: Non-Acute Care: General, low-level care for mildly to moderately symptomatic COVID19 patients. These patients may require oxygen (less than or equal to 2L/min), but do not require extensive nursing care or assistance with activities of daily living (ADL). This level of care corresponds to Level 5 (ambulatory care) and Level 4 (minor acuity care) patients in medical care terminology. Hospital Care: Mid-level care for moderately symptomatic COVID-19 patients. These patients require oxygen (more than 2L/min), nursing care, and assistance with ADL. This level of care corresponds to Level 3 (medical-surgical care) patients in medical care terminology. Acute Care: Higher acuity care for COVID-19 patients. These patients require significant ventilatory support, including intensive monitoring on a ventilator. This level of care corresponds to Level 2 (step-down care) and Level 1 (intensive care unit [ICU] care) patients in medical care terminology.The expected duration of care for patients in ACS would be based on their clinical needs and potentially the timeline for discontinuation of Transmission-Based Precautions. If ACS will be used to care for both confirmed and suspected COVID-19 patients or for patients without COVID-19 who require care for other reasons, additional infection prevention and control considerations will apply. For example, planning would need to address physical separation between the cohorts and assigning different HCP with dedicated equipment to each section."

    # #test sample
    # bert_output = "1 .7".split(" ")
    # tfidf_output = "0"


    #### ****** CHANGE THE S3 BUCKET NAME/LOCATION ******
    s3_client = boto3.client('s3')
    s3_client.download_file('test-model-masterproject','pac.pkl', '/tmp/pac.pkl')
    s3_client.download_file('test-model-masterproject','vect.pkl', '/tmp/vect.pkl')


    #invoke sagemaker here for bert model
    #sample output: "1 .94234"
    bert_output = invoke_sagemaker()


    #invoke tfidf model
    #sample output: "0"
    tfidf_output = tfidf_prediction(real)

    #calculate scores and make article determination
    article_determiniation, confidence_score = calculate_result_dynamic(bert_output, tfidf_output)

    result = {
      "text_determination": article_determiniation,
      "confidence_score": confidence_score,
    }
    print(result)

    return {
        'statusCode': 200,
        'body': json.dumps(result)
    }
