from operator import itemgetter
import sys
from collections import defaultdict
import numpy as np
import pandas as pd

file = 'movie_ratings_df.csv'
data = pd.read_csv(file)
# 定义rating > 3的为该user喜欢的电影
data["favor"] = data["rating"] > 3

# 选择前200 users
data_200 = data[data['userId'].isin(range(200))]
# 只选择喜欢的电影
favor_data = data_200[data_200["favor"] == 1]
# 将用户评论的电影放入一个集合
# frozenset 冻结集合
favor_reviews_by_users = dict((k, frozenset(v.values))
                              for k, v in favor_data.groupby("userId")["title"])

# 电影被喜欢的数量
num_favor_by_movie = data_200[["title", "favor"]].groupby("title").sum()

# 从k-1项集生成k项集
def find_frequent_itemsets(favor_reviews_by_users, k_itemsets, min_support):
    counts = defaultdict(int)
    for user, review in favor_reviews_by_users.items():
        for itemset in k_itemsets:
            if itemset.issubset(review):
                for other_reviewed_movie in review-itemset:
                    current_superset = itemset | frozenset(
                        (other_reviewed_movie,))
                    counts[current_superset] += 1
    return dict([(itemset, frequence) for itemset, frequence in counts.items() if frequence >= min_support])


frequent_itemsets = {}
min_support = 50
# 长度为1的频繁项集
frequent_itemsets[1] = dict((frozenset((title,)), row['favor'])
                            for title, row in num_favor_by_movie.iterrows() if row['favor'] > min_support)

for k in range(2, 20):
    # 通过k-1个频繁项集产生k个频繁项集
    cur_frequent_itemsets = find_frequent_itemsets(
        favor_reviews_by_users, frequent_itemsets[k-1], min_support)
    if len(cur_frequent_itemsets) == 0:
        print("找不到频繁 {} 项集".format(k))
        sys.stdout.flush()
        break
    else:
        print("找到 {} 个频繁 {} 项集".format(
            len(cur_frequent_itemsets), k))
        sys.stdout.flush()
        frequent_itemsets[k] = cur_frequent_itemsets

# 长度为1的频繁项集不需要
del frequent_itemsets[1]
print("总共 {} 个频繁项集.".format(sum(len(frequent_item)
                                for frequent_item in frequent_itemsets.values())))

# 创建关联规则
candidate_rules = []

for itemset_length, itemset_counts in frequent_itemsets.items():
    for itemset in itemset_counts.keys():
        # 将其中一个项作为结论,其他作为前提
        for conclusion in itemset:
            premise = itemset - set((conclusion,))
            candidate_rules.append((premise, conclusion))
print("总共有 {} 个关联规则.".format(len(candidate_rules)))

# 计算置信度
correct_counts = defaultdict(int)
incorrect_counts = defaultdict(int)

for user, reviews in favor_reviews_by_users.items():
    for candidate_rule in candidate_rules:
        premise, conclusion = candidate_rule
        if premise.issubset(reviews):
            if conclusion in reviews:
                correct_counts[candidate_rule] += 1
            else:
                incorrect_counts[candidate_rule] += 1

rule_confidence = {candidate_rule: correct_counts[candidate_rule] / float(correct_counts[candidate_rule] + incorrect_counts[candidate_rule])
                   for candidate_rule in candidate_rules}

# 通过最小置信度过滤
min_confidence = 0.9
rule_confidence = {rule: confidence for rule,
                   confidence in rule_confidence.items() if confidence > min_confidence}

# 排序
sort_confidence = sorted(rule_confidence.items(),
                         key=itemgetter(1), reverse=True)

# 输出关联规则
# for index in range(0, 5):
#     print("规则 #{0}:".format(index+1))
#     premise, conclusion = sort_confidence[index][0]
#     premise_name = ", ".join(premise)
#     print("如果用户喜欢 {0}, 就可能喜欢 {1}".format(
#         premise_name, conclusion))
#     print("置信度: {0:.1%}".format(sort_confidence[index][1]))

print('----------------开始测试----------------')

# 测试数据
test_data = data[~data['userId'].isin(range(200))]
test_favor = test_data[test_data["favor"]]
test_favor_by_users = dict((k, frozenset(v.values))
                           for k, v in test_favor.groupby('userId')['title'])

correct_counts = defaultdict(int)
incorrect_counts = defaultdict(int)
for user, reviews in test_favor_by_users.items():
    for candidate_rule in candidate_rules:
        premise, conclusion = candidate_rule
        if premise.issubset(reviews):
            if conclusion in reviews:
                correct_counts[candidate_rule] += 1
            else:
                incorrect_counts[candidate_rule] += 1
test_confidence = {candidate_rule: correct_counts[candidate_rule] / float(correct_counts[candidate_rule]+incorrect_counts[candidate_rule])
                   for candidate_rule in rule_confidence}

sort_test_confidence = sorted(
    test_confidence.items(), key=itemgetter(1), reverse=True)


for index in range(5):
    print("规则 #{0}:".format(index+1))
    premise, conclusion = sort_confidence[index][0]
    premise_name = ", ".join(premise)
    print("如果用户喜欢 {0}, 就可能喜欢 {1}".format(premise_name, conclusion))
    print("测试置信度: {0:.1%}".format(rule_confidence.get((premise,conclusion),-1)))
    print("实际置信度: {0:.1%}".format(test_confidence.get((premise,conclusion),-1)))
