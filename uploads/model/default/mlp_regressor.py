import os
import pickle

import numpy as np
import pandas as pd
import sklearn
from sklearn.model_selection import train_test_split
from sklearn.neural_network import MLPRegressor


def read_dataset(target):
    dataset = pd.read_csv(os.path.join(os.path.dirname(__file__), "..", "data/data.csv"))

    y = dataset[target]
    x = dataset.drop(target, axis=1)

    return x, y


def split_dataset(x, y):
    return train_test_split(x, y, test_size=0.3, random_state=0)


def train_dataset(x_train, y_train):
    clf = MLPRegressor()
    clf.fit(x_train, y_train)

    return clf


def evaluate_classifier_model(classifier, x_test, y_test):
    y_pred = classifier.predict(x_test)
    y_score = classifier.predict_proba(x_test)

    accuracy = sklearn.metrics.accuracy_score(y_test, y_pred)
    recall = sklearn.metrics.recall_score(y_test, y_pred, average="micro")
    f1_score = sklearn.metrics.f1_score(y_test, y_pred, average="micro")
    confusion_matrix = sklearn.metrics.confusion_matrix(y_test, y_pred)

    return {
        "accuracy": accuracy,
        "recall": recall,
        "f1_score": f1_score,
        "confusion_matrix": confusion_matrix.tolist()
    }


def evaluate_regressor_model(classifier, x_test, y_test):
    y_pred = classifier.predict(x_test)

    mean_squared_error = sklearn.metrics.mean_squared_error(y_test, y_pred)
    root_mean_squared_error = np.sqrt(sklearn.metrics.mean_squared_error(y_test, y_pred))
    mean_absolute_error = sklearn.metrics.mean_absolute_error(y_test, y_pred)
    r2_score = sklearn.metrics.r2_score(y_test, y_pred)
    explained_variance_score = sklearn.metrics.explained_variance_score(y_test, y_pred)
    max_error = sklearn.metrics.max_error(y_test, y_pred)

    return {
        "mean_squared_error": mean_squared_error,
        "root_mean_squared_error": root_mean_squared_error,
        "mean_absolute_error": mean_absolute_error,
        "r2_score": r2_score,
        "explained_variance_score": explained_variance_score,
        "max_error": max_error,
    }


model_path = os.path.join(os.path.dirname(__file__), "..", "model/model.pickle")


def save_model_file(model):
    pickle.dump(model, open(model_path, "wb"))


def run(config):
    x, y = read_dataset(config["target"])
    x_train, x_test, y_train, y_test = split_dataset(x, y)

    if not os.path.exists(model_path):
        model = train_dataset(x_train, y_train)
        save_model_file(model)
    else:
        with open(model_path, 'rb') as model_file:
            model = pickle.load(model_file)

    evaluation = evaluate_regressor_model(model, x_test, y_test)
    return evaluation
