import os
import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
from scipy import stats


dirpath = os.curdir
outFile = open(dirpath+'/Analise/log.txt', 'w')

file1 = dirpath+'/Analise/Nosso/dataBaseDirect.csv'
file2 = dirpath+'/Analise/Nosso/dataBaseFanout.csv'
file3 = dirpath+'/Analise/Nosso/dataBaseTopic.csv'
file4 = dirpath+'/Analise/Rabbitmq/dataBaseDirect1.csv'
file5 = dirpath+'/Analise/Rabbitmq/dataBaseFanout1.csv'
file6 = dirpath+'/Analise/Rabbitmq/dataBaseTopic1.csv'

df = pd.read_csv(file1)
df2 = pd.read_csv(file2)
df3 = pd.read_csv(file3)
df4 = pd.read_csv(file4)
df5 = pd.read_csv(file5)
df6 = pd.read_csv(file6)

dados = df["data"]
dir1 = dados[500:]
dir1 = dir1.astype('float64')

dados = df2["data"]
fan1 = dados[500:]
fan1 = fan1.astype('float64')

dados = df3["data"]
top1 = dados[500:]
top1 = top1.astype('float64')

dados = df4["data"]
dir2 = dados[500:]
dir2 = dir2.astype('float64')

dados = df5["data"]
fan2 = dados[500:]
fan2 = fan2.astype('float64')

dados = df6["data"]
top2 = dados[500:]
top2 = top2.astype('float64')

# Média
mean_dir1 = dir1.mean()
mean_dir2 = dir2.mean()

mean_fan1 = fan1.mean()
mean_fan2 = fan2.mean()

mean_top1 = top1.mean()
mean_top2 = top2.mean()


# Desvio Padrao
std_dir1 = dir1.std()
std_dir2 = dir2.std()

std_fan1 = fan1.std()
std_fan2 = fan2.std()

std_top1 = top1.std()
std_top2 = top2.std()


# TESTING FOR NORMALITY
stat, p = stats.shapiro(dir1)   #shapiro-wilk test
# k2, p = stats.normaltest(mdl)  #test that combines skew and kurtosis
alpha = 0.05
print("p = {:g}".format(p))
if p <= alpha:  # null hypothesis (H0): data comes from a normal distribution
    print("reject H0, not normal.")
else:
    print("fail to reject H0, normal.")

stat, p = stats.shapiro(dir2)
# k2, p = stats.normaltest(rpc)
alpha = 0.05
print("p = {:g}".format(p))
if p <= alpha:  # null hypothesis (H0): data comes from a normal distribution
    print("reject H0, not normal.")
else:
    print("fail to reject H0, normal.")
    

# # NULL HYPOTHESIS TESTING
# # 2 sample t-test
# t, pVal1 = stats.ttest_ind(dir1,dir2)
# t, pVal2 = stats.ttest_ind(fan1,fan2)
# t, pVal3 = stats.ttest_ind(top1,top2)

# Mann whitney U test
U, pVal1 = stats.mannwhitneyu(dir1,dir2)
U, pVal2 = stats.mannwhitneyu(fan1,fan2)
U, pVal3 = stats.mannwhitneyu(top1,top2)

# print(pVal)
if pVal1 < 0.05:
    logTest = "Reject NULL hypothesis - Significant differences exist between groups."
if pVal1 > 0.05:
    logTest = "Accept NULL hypothesis - No significant difference between groups."


outFile.write('1 Cliente:\n')
outFile.write('-----------------NOSSO-----------------\n')
outFile.write('média direto = ' + str(mean_dir1) + '\n')
outFile.write('desvio Padrão direto = ' + str(std_dir1) + '\n')
outFile.write('média fanout = ' + str(mean_fan1) + '\n')
outFile.write('desvio Padrão fanout = ' + str(std_fan1) + '\n')
outFile.write('média topic = ' + str(mean_top1) + '\n')
outFile.write('desvio Padrão topic = ' + str(std_top1) + '\n')
outFile.write('---------------RABBITMQ------------\n')
outFile.write('média direto = ' + str(mean_dir2) + '\n')
outFile.write('desvio Padrão direto = ' + str(std_dir2) + '\n')
outFile.write('média fanout = ' + str(mean_fan2) + '\n')
outFile.write('desvio Padrão fanout = ' + str(std_fan2) + '\n')
outFile.write('média topic = ' + str(mean_top2) + '\n')
outFile.write('desvio Padrão topic = ' + str(std_top2) + '\n')
outFile.write('-------NULL HYPOTHESIS TESTING-------\n')
outFile.write('p1 = ' + str(pVal1) + '\n')
outFile.write('p2 = ' + str(pVal2) + '\n')
outFile.write('p3 = ' + str(pVal3) + '\n')
outFile.write('-------------------------------------\n')
outFile.close()

# ,mean_fan1,mean_fan2,mean_top1,mean_top2
medias = [mean_top1,mean_top2]
configs = ['NOSSO','RABBITMQ']

plt.bar(configs,medias,color="blue")
plt.xticks(configs)
plt.ylabel('Tempo Médio (ms)')
plt.xlabel('Middleware Utilizado')
plt.title('Middleware Utilizado x Tempo médio (N = 10.000)')
plt.savefig(dirpath+'/Analise/GraficoTopic.png', format='png')
plt.show()

