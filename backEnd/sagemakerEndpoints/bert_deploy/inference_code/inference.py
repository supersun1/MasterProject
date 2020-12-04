import os
import json
import torch
import pandas as pd
from transformers import (BertForSequenceClassification,
                          BertTokenizer,
                          AdamW)

def model_fn(model_dir):
    """
    Load the model for inference
    """
    device = torch.device('cpu')
    print('load the model')

    model_path = os.path.join(model_dir, 'model/')
    print('find cpu')

    # Load BERT tokenizer from disk.
    tokenizer = BertTokenizer.from_pretrained(model_path)
    
    # Load BERT model from disk.
    model = BertForSequenceClassification.from_pretrained(model_path,
                                                          num_labels = 2,
                                                          output_attentions = False,
                                                          output_hidden_states = False
                                                         )
    
    print('finished loading the model files')
    model.to(device).eval()
        
    model_dict = {'model': model, 'tokenizer':tokenizer}
    
    return model_dict

def predict_fn(input_data, model):
    """
    Apply model to the incoming request
    """
    
    tokenizer = BertTokenizer.from_pretrained("bert-base-uncased")
    model = model['model']
    
    encoded_input = tokenizer(input_data,  
                              max_length = 500,
                              truncation = True,
                              return_attention_mask = True,
                              return_tensors='pt')
    
    bert_result = model(**encoded_input)
    bert_final = bert_result[0]
    bert_final = bert_final.detach().cpu().numpy()
    
    return bert_final[0]

def input_fn(request_body, request_content_type):
    """
    Deserialize and prepare the prediction input
    """
    
    if request_content_type == "application/json":
        request = json.loads(request_body)
    else:
        request = request_body

    return request

def output_fn(prediction, response_content_type):
    """
    Serialize and prepare the prediction output
    """
    
    if response_content_type == "application/json":
        response = str(prediction)
    else:
        response = str(prediction)

    return response