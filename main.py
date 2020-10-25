import os
import csv
import pandas as pd
from tqdm import tqdm
from simpletransformers.language_modeling import LanguageModelingModel

from sklearn.model_selection import train_test_split

def fnc(path_headlines, path_bodies):

    map = {'agree': 0, 'disagree':1, 'discuss':2, 'unrelated':3}

    with open(path_bodies, encoding='utf_8') as fb:  # Body ID,articleBody
        body_dict = {}
        lines_b = csv.reader(fb)
        for i, line in enumerate(tqdm(list(lines_b), ncols=80, leave=False)):
            if i > 0:
                body_id = int(line[0].strip())
                body_dict[body_id] = line[1]

    with open(path_headlines, encoding='utf_8') as fh: # Headline,Body ID,Stance
        lines_h = csv.reader(fh)
        h = []
        b = []
        l = []
        for i, line in enumerate(tqdm(list(lines_h), ncols=80, leave=False)):
            if i > 0:
                body_id = int(line[1].strip())
                label = line[2].strip()
                if label in map and body_id in body_dict:
                    h.append(line[0])
                    l.append(map[line[2]])
                    b.append(body_dict[body_id])
    return h, b, l

curDir = os.getcwd()
data_dir = curDir + "/dataset"
headlines, bodies, labels = fnc(
    os.path.join(data_dir, 'combined_stances_train.csv'),
    os.path.join(data_dir, 'combined_bodies_train.csv')
)

list_of_tuples = list(zip(headlines, bodies, labels))
df = pd.DataFrame(list_of_tuples, columns=['text_a', 'text_b', 'label'])
train_df, val_df = train_test_split(df)
labels_val = pd.Series(val_df['label']).to_numpy()

headlines, bodies, labels = fnc(
    os.path.join(data_dir, 'combined_stances_test.csv'),
    os.path.join(data_dir, 'combined_bodies_test.csv')
)

list_of_tuples = list(zip(headlines, bodies, labels))
test_df = pd.DataFrame(list_of_tuples, columns=['text_a', 'text_b', 'label'])
labels_test = pd.Series(test_df['label']).to_numpy()


model = LanguageModelingModel('roberta', 'roberta-base', num_labels=4, args={
    'learning_rate':1e-5,
    'num_train_epochs': 10,
    'reprocess_input_data': True,
    'overwrite_output_dir': True,
    'process_count': 10,
    'train_batch_size': 4,
    'eval_batch_size': 4,
    'max_seq_length': 512,
    'fp16': True
})
model.train_model(train_df)
