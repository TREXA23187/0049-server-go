import json
import os
import pickle
import warnings

import pandas as pd
import sklearn
from sklearn.ensemble import RandomForestClassifier
from sklearn.metrics import classification_report
from sklearn.model_selection import GridSearchCV
from sklearn.model_selection import train_test_split

warnings.filterwarnings("ignore")


def read_dataset(data, target):
    dataset = pd.read_csv(os.path.join(os.path.dirname(__file__), "..", "data/data.csv"))

    target_labels = dataset[target]

    label_to_int = {label: i for i, label in enumerate(target_labels)}
    int_to_label = {i: label for label, i in label_to_int.items()}

    label_int_tag = {
        "label_to_int": label_to_int,
        "int_to_label": int_to_label
    }
    with open('model/label_int_tag.json', 'w', encoding='utf-8') as label_int_tag_file:
        json.dump(label_int_tag, label_int_tag_file)

    encoded_labels = [label_to_int[label] for label in target_labels]
    dataset[target] = encoded_labels

    y = dataset[target]
    x = dataset[data]

    return x, y


def split_dataset(x, y):
    return train_test_split(x, y, test_size=0.3, random_state=0)


def train_dataset(x_train, y_train):
    clf = RandomForestClassifier()
    clf.fit(x_train, y_train)

    return clf


def evaluate_classifier_model(classifier, x_test, y_test):
    y_pred = classifier.predict(x_test)

    confusion_matrix = sklearn.metrics.confusion_matrix(y_test, y_pred)

    return {
        "classification_report": classification_report(y_test, y_pred, output_dict=True),
        "confusion_matrix": confusion_matrix.tolist()
    }


model_path = os.path.join(os.path.dirname(__file__), "..", "model/model.pickle")


def save_model_file(model):
    pickle.dump(model, open(model_path, "wb"))


def advance_operate(x, y, hyper_parameters):
    grid_search = GridSearchCV(RandomForestClassifier(), hyper_parameters, cv=10, scoring='accuracy')

    grid_search.fit(x, y)

    return grid_search.best_params_, grid_search.best_score_


def run(config):
    x, y = read_dataset(config["dataLabel"], config["targetLabel"])

    x_train, x_test, y_train, y_test = split_dataset(x, y)

    if not os.path.exists(model_path):
        model = train_dataset(x_train, y_train)
        save_model_file(model)
    else:
        with open(model_path, 'rb') as model_file:
            model = pickle.load(model_file)

    evaluation = evaluate_classifier_model(model, x_test, y_test)

    if bool(config["hyperParameters"]):
        optimal_hyper_parameters, optimal_accuracy = advance_operate(x, y, config["hyperParameters"])
        advance_evaluation = {"optimal_hyper_parameters": optimal_hyper_parameters,
                              "optimal_accuracy": optimal_accuracy}
        return {**evaluation, **advance_evaluation}
    else:
        return evaluation
