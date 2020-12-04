import os
import json
import torch
import pandas as pd
import logging as logger
from transformers import (RobertaForSequenceClassification,
                          RobertaTokenizer,
                          AdamW)
from torch.utils.data import TensorDataset, random_split
from torch.utils.data import DataLoader, RandomSampler, SequentialSampler

def model_fn(model_dir):
    #=====================Loading the Model========================
    device = torch.device('cpu')
    print('load the model')

    model_path = os.path.join(model_dir, 'model/')
    print('find cpu')

    
    
    tokenizer = RobertaTokenizer.from_pretrained(model_path)

    
    model_p = os.path.join(model_dir, "model.pth")
    print("get model pth")

    
    model = RobertaForSequenceClassification.from_pretrained(model_path,
                                                          num_labels = 2,
                                                          output_attentions = False,
                                                          output_hidden_states = False
                                                         )
    print("model dir: " + str(model_dir))
    with open(os.path.join(model_dir, 'model.pth'), 'rb') as f:
        model.load_state_dict(torch.load(f, map_location=torch.device('cpu')))

    
    print('finished reading pth files')

    model.to(device).eval()
    print('finish loading the damn model')
        
    model_dict = {'model': model, 'tokenizer':tokenizer}
    
    return model_dict

def predict_fn(input_data, model):
    #==============Setting up the Model to receive requests========================
    
    tokenizer = model['tokenizer']
    bert_model = model['model']
    
    encoded_input = tokenizer(input_data,
                              max_length = 500,
                              truncation = True,
                              return_attention_mask = True,
                              return_tensors='pt')
    
    return bert_model(**encoded_input)


def input_fn(request_body, request_content_type):
    #=====================Preparing the input for the Model========================
    
    if request_content_type == "application/json":
        request = json.loads(request_body)
    else:
        request = request_body

    return request

def output_fn(prediction, response_content_type):
    #=====================Preparing the output from the Model========================
    
    if response_content_type == "application/json":
        response = str(prediction)
    else:
        response = str(prediction)

    return response