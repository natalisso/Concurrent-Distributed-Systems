import os
import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
from scipy import stats
 
 
dirpath = os.getcwd()
numClients = '1'
df = pd.read_csv(dirpath+'/MOM/dataBase'+numClients+'.csv')
df2 = pd.read_csv(dirpath+'/RPC/dataBase'+numClients+'.csv')

dados = df["data"]
mom = dados[2:]
mom = mom.astype('float64')

dados = df2["data"]
rpc = dados[2:]
rpc = rpc.astype('float64')

# Média
mean_mom = mom.mean()
mean_rpc = rpc.mean()

# Desvio Padrao
std_mom = mom.std()
std_rpc = rpc.std()

print(numClients, ' Cliente:')
print('------MOM------')
print("média =",mean_mom)
print("desvio Padrão =",std_mom)
print('------RPC------')
print("média =",mean_rpc)
print("desvio Padrão =",std_rpc)
print()

medias = [mean_mom,mean_rpc]
configs = ['MOM','RPC']

plt.bar(configs,medias,color="blue")
plt.xticks(configs)
plt.ylabel('Tempo Médio (ms)')
plt.xlabel('Tipo do Middleware')
plt.title('Tipo do Middleware x Tempo médio (N=10.000)')
plt.savefig('Grafico'+numClients+'.png', format='png')
plt.show()

