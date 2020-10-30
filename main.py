import os
import csv
import pandas as pd
from tqdm import tqdm
# from simpletransformers.classification import LanguageModelingModel
from simpletransformers.classification import ClassificationModel
import logging


# from sklearn.model_selection import train_test_split
#
# def fnc(path_headlines, path_bodies):
#
#     map = {'agree': 0, 'disagree':1, 'discuss':2, 'unrelated':3}
#
#     with open(path_bodies, encoding='utf_8') as fb:  # Body ID,articleBody
#         body_dict = {}
#         lines_b = csv.reader(fb)
#         for i, line in enumerate(tqdm(list(lines_b), ncols=80, leave=False)):
#             if i > 0:
#                 body_id = int(line[0].strip())
#                 body_dict[body_id] = line[1]
#
#     with open(path_headlines, encoding='utf_8') as fh: # Headline,Body ID,Stance
#         lines_h = csv.reader(fh)
#         h = []
#         b = []
#         l = []
#         for i, line in enumerate(tqdm(list(lines_h), ncols=80, leave=False)):
#             if i > 0:
#                 body_id = int(line[1].strip())
#                 label = line[2].strip()
#                 if label in map and body_id in body_dict:
#                     h.append(line[0])
#                     l.append(map[line[2]])
#                     b.append(body_dict[body_id])
#     return h, b, l
#
# curDir = os.getcwd()
# data_dir = curDir + "/dataset"
# headlines, bodies, labels = fnc(
#     os.path.join(data_dir, 'combined_stances_train.csv'),
#     os.path.join(data_dir, 'combined_bodies_train.csv')
# )
#
# list_of_tuples = list(zip(headlines, bodies, labels))
# df = pd.DataFrame(list_of_tuples, columns=['text_a', 'text_b', 'label'])
# train_df, val_df = train_test_split(df)
# labels_val = pd.Series(val_df['label']).to_numpy()
#
# headlines, bodies, labels = fnc(
#     os.path.join(data_dir, 'combined_stances_test.csv'),
#     os.path.join(data_dir, 'combined_bodies_test.csv')
# )
#
# list_of_tuples = list(zip(headlines, bodies, labels))
# test_df = pd.DataFrame(list_of_tuples, columns=['text_a', 'text_b', 'label'])
# labels_test = pd.Series(test_df['label']).to_numpy()
#
#
#
#
#
#
#
#
#
# model.train_model(train_df)





#============ sample ==================
logging.basicConfig(level=logging.INFO)
transformers_logger = logging.getLogger("transformers")
transformers_logger.setLevel(logging.WARNING)

# Train and Evaluation data needs to be in a Pandas Dataframe containing at least two columns. If the Dataframe has a header, it should contain a 'text' and a 'labels' column. If no header is present, the Dataframe should contain at least two columns, with the first column is the text with type str, and the second column in the label with type int.
train_data = [['Example sentence belonging to class 1', 1], ['Example sentence belonging to class 0', 0], ['Example eval senntence belonging to class 2', 2]]
train_df = pd.DataFrame(train_data)

eval_data = [['Example eval sentence belonging to class 1', 1], ['Example eval sentence belonging to class 0', 0], ['Example eval senntence belonging to class 2', 2]]
eval_df = pd.DataFrame(eval_data)

# Create a ClassificationModel
model = ClassificationModel('bert', 'bert-base-cased', num_labels=2, args={'reprocess_input_data': True, 'overwrite_output_dir': True, 'fp16':False}, use_cuda=False)
# You can set class weights by using the optional weight argument

# Train the model
model.train_model(train_df)

# Evaluate the model
result, model_outputs, wrong_predictions = model.eval_model(eval_df)

predictions, raw_outputs = model.predict(["Some arbitary sentence"])
