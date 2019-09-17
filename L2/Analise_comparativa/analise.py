import os
import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
from scipy import stats


 
 
dirpath = os.getcwd()
numClients = '1'
N = 10000
df = pd.read_csv(dirpath+'/MOM/dataBase'+numClients+'.csv')
df2 = pd.read_csv(dirpath+'/MOO/dataBase'+numClients+'.csv')

dados = df["data"]
mom = dados[1:]
mom = mom.astype('float64')

dados = df2["data"]
moo = dados[1:]
moo = moo.astype('float64')

# Média
mean_mom = mom.mean()
mean_moo = moo.mean()

# Variância
var_mom = mom.var(ddof=1)
var_moo = moo.var(ddof=1)

# Desvio Padrao
std_mom = mom.std()
std_moo = moo.std()
s =  np.sqrt((var_mom + var_moo)/2)

print(numClients, ' Cliente:')
print('------MOM------')
print("média =",mean_mom)
print("desvio Padrão =",std_mom)
print('------MOO------')
print("média =",mean_moo)
print("desvio Padrão =",std_moo)
print()

medias = [mean_mom,mean_moo]
configs = ['MOM','MOO']

plt.bar(configs,medias,color="blue")
plt.xticks(configs)
plt.ylabel('Tempo Médio (ms)')
plt.xlabel('Tipo do Middleware')
plt.title('Tipo do Middleware x Tempo médio (N=10.000)')
plt.savefig('Grafico'+numClients+'.png', format='png')
plt.show()

# t-statistics
t = (mean_mom - mean_moo ) / (s * np.sqrt(2/N))
## Compare with the critical t-value
#Degrees of freedom
df = (2 * N) - 2
#p-value after comparison with the t 
p = 1 - stats.t.cdf(t,df=df)
print("t = " + str(t))
print("p = " + str(2*p))
## Cross Checking with the internal scipy function
t2, p2 = stats.ttest_ind(mom,moo)
print("t = " + str(t2))
print("p = " + str(p2))