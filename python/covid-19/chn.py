import pandas as pd
import numpy as np
from datetime import datetime, timedelta
from sklearn.metrics import mean_squared_error
from scipy.optimize import curve_fit
from scipy.optimize import fsolve
import matplotlib.pyplot as plt
# %matplotlib inline

# https://raw.githubusercontent.com/CSSEGISandData/COVID-19/master/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_confirmed_global.csv
url = 'https://cdn.jsdelivr.net/gh/CSSEGISandData/COVID-19/csse_covid_19_data/csse_covid_19_time_series/time_series_covid19_confirmed_global.csv'
df = pd.read_csv(url)
df = df[df['Country/Region'] == 'China']
df.drop(['Country/Region', 'Lat', 'Long',
         'Province/State'], axis=1, inplace=True)

FMT = '%m/%d/%y'
FMT_NZP = '%#m/%#d/%y'  # no zero-padding
start = datetime.strptime('1/22/20', FMT)
end = datetime.now() - timedelta(days=1)
d = {}
while start < end:
    day = start.strftime(FMT_NZP)
    d[day] = df[day].sum()
    start += timedelta(days=1)

ndf = pd.DataFrame(d.items(), columns=['date', 'confirmed'])
date = ndf['date']
ndf['date'] = date.map(lambda x: (datetime.strptime(
    x, FMT) - datetime.strptime("1/21/20", FMT)).days)
print(ndf)
# ndf.plot(x='date', y='confirmed')
# plt.show()

def logistic_model(x, a, b, c):
    return c/(1+np.exp(-(x-b)/a))

x = list(ndf.iloc[:, 0])
y = list(ndf.iloc[:, 1])
fit = curve_fit(logistic_model, x, y, p0=[3, 100, 100000])
print(fit)
errors = [np.sqrt(fit[1][i][i]) for i in [0, 1, 2]]
print(errors)