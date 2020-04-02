import pandas as pd
import numpy as np
from datetime import datetime, timedelta
from sklearn.metrics import mean_squared_error
from scipy.optimize import curve_fit
from scipy.optimize import fsolve
import matplotlib.pyplot as plt
# %matplotlib inline
import io
import requests

# https://raw.githubusercontent.com/pcm-dpc/COVID-19/master/dati-andamento-nazionale/dpc-covid19-ita-andamento-nazionale.csv
url = "https://cdn.jsdelivr.net/gh/pcm-dpc/COVID-19/dati-andamento-nazionale/dpc-covid19-ita-andamento-nazionale.csv"
df = pd.read_csv(url)
# df = df[df['data'] <= '2020-03-07T18:00:00']
df = df.loc[:, ['data', 'totale_casi']]
df = df
FMT = '%Y-%m-%dT%H:%M:%S'
date = df['data']
df['data'] = date.map(lambda x: (datetime.strptime(
    x, FMT) - datetime.strptime("2020-01-01T00:00:00", FMT)).days)


def logistic_model(x, a, b, c):
    return c/(1+np.exp(-(x-b)/a))


print(df)
x = list(df.iloc[:, 0])
y = list(df.iloc[:, 1])
fit = curve_fit(logistic_model, x, y, p0=[2, 100, 200000])
print(fit)
errors = [np.sqrt(fit[1][i][i]) for i in [0, 1, 2]]
print(errors)
f = fit[0]
a = f[0]
b = f[1]
c = f[2]
sol = int(fsolve(lambda x: logistic_model(x, a, b, c) - int(c), b))
print(sol)