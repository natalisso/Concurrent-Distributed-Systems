import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
from scipy import stats

df = pd.read_csv("~/Documentos/ProgConcDist/L1/tcp2/server/data_bases/analise_estatistica/dataBase1.csv")
df2 = pd.read_csv("~/Documentos/ProgConcDist/L1/tcp2/server/data_bases/analise_estatistica/dataBase2.csv")
df3 = pd.read_csv("~/Documentos/ProgConcDist/L1/tcp2/server/data_bases/analise_estatistica/dataBase3.csv")
df4 = pd.read_csv("~/Documentos/ProgConcDist/L1/tcp2/server/data_bases/analise_estatistica/dataBase4.csv")
df5 = pd.read_csv("~/Documentos/ProgConcDist/L1/tcp2/server/data_bases/analise_estatistica/dataBase5.csv")

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

# Média
mt1 = tempos1.mean()
mt2 = tempos2.mean()
mt3 = tempos3.mean()
mt4 = tempos4.mean()
mt5 = tempos5.mean()

# Desvio Padrao
dt1 = tempos1.std()
dt2 = tempos2.std()
dt3 = tempos3.std()
dt4 = tempos4.std()
dt5 = tempos5.std()

print("mt1 =",mt1)
print("mt2 =",mt2)
print("mt3 =",mt3)
print("mt4 =",mt4)
print("mt5 =",mt5)
print()
print("dt1 =",dt1)
print("dt2 =",dt2)
print("dt3 =",dt3)
print("dt4 =",dt4)
print("dt5 =",dt5)

medias = [mt1,mt2,mt3,mt4,mt5]
configs = ['TCP1','TCP2','TCP3','TCP4','TCP5']

plt.bar(configs,medias,color="blue")
plt.xticks(configs)
plt.ylabel('Tempo Médio (us)')
plt.xlabel('Configuração')
plt.title('Configuração x Tempo médio (N=10.000)')
plt.savefig('teste.png', format='png')
plt.show()