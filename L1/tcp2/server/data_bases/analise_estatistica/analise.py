import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
from scipy import stats

df = pd.read_csv("~/Documentos/ProgConcDist/L1/tcp2/server/data_bases/dataBase1.csv")
df2 = pd.read_csv("~/Documentos/ProgConcDist/L1/tcp2/server/data_bases/dataBase2.csv")
df3 = pd.read_csv("~/Documentos/ProgConcDist/L1/tcp2/server/data_bases/dataBase3.csv")
df4 = pd.read_csv("~/Documentos/ProgConcDist/L1/tcp2/server/data_bases/dataBase4.csv")
df5 = pd.read_csv("~/Documentos/ProgConcDist/L1/tcp2/server/data_bases/dataBase5.csv")

dados = df["data"]
tempos1 = dados[2:]
tempos1 = tempos1.astype('float64')

dados = df2["data"]
tempos2 = dados[2:]
tempos2 = tempos2.astype('float64')

dados = df3["data"]
tempos3 = dados[2:]
tempos3 = tempos3.astype('float64')

dados = df4["data"]
tempos4 = dados[2:]
tempos4 = tempos4.astype('float64')

dados = df5["data"]
tempos5 = dados[2:]
tempos5 = tempos5.astype('float64')

t1 =tempos1.mean()
t2 =tempos2.mean()
t3 =tempos3.mean()
t4 =tempos4.mean()
t5 =tempos5.mean()

medias = [t1,t2,t3,t4,t5]
configs = ['TCP1','TCP2','TCP3','TCP4','TCP5']

plt.bar(configs,medias,color="blue")
plt.xticks(configs)
plt.ylabel('Tempo Médio (us)')
plt.xlabel('Configuração')
plt.title('Configuração x Tempo médio (N=10.000)')
plt.savefig('teste.png', format='png')
plt.show()