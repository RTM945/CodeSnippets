import pandas as pd

file = 'movie_ratings_df.csv'
data = pd.read_csv(file)

users = data.groupby('userId')['userId'].nunique()
print('用户总数 {}'.format(users.count()))

movies = data.groupby('title')['title'].nunique()
print('电影总数 {}'.format(movies.count()))

# 电影被评价次数
movie_rating_count = data['title'].value_counts().reset_index(name ='count')
print(movie_rating_count)

# 被评价总分
movie_rating = data.groupby('title')['rating'].sum().reset_index(name ='total rating')
print(movie_rating.sort_values('total rating', ascending=False))

# 评分人数大于150人的电影的平均分数
import numpy as np
movie_rating_avg = data.groupby('title').agg({'rating':[np.size, np.mean]})
movie_rating_avg = movie_rating_avg[movie_rating_avg['rating']['size'] > 150]
print(movie_rating_avg.sort_values([('rating', 'mean')], ascending=False))