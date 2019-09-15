import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
from scipy import stats

df = pd.read_csv("~/Documentos/L1/Analise_comparativa/TCP/dataBase1(1).csv")
df2 = pd.read_csv("~/Documentos/L1/Analise_comparativa/TCP/dataBase1(2).csv")
df3 = pd.read_csv("~/Documentos/L1/Analise_comparativa/TCP/dataBase1(3).csv")
df4 = pd.read_csv("~/Documentos/L1/Analise_comparativa/TCP/dataBase1(4).csv")
df5 = pd.read_csv("~/Documentos/L1/Analise_comparativa/TCP/dataBase1(5).csv")

dados = df["data"]
tcp1 = dados[2:]
tcp1 = tcp1.astype('float64')

dados = df2["data"]
tcp2 = dados[2:]
tcp2 = tcp2.astype('float64')

dados = df3["data"]
tcp3 = dados[2:]
tcp3 = tcp3.astype('float64')

dados = df4["data"]
tcp4 = dados[2:]
tcp4 = tcp4.astype('float64')

dados = df5["data"]
tcp5 = dados[2:]
tcp5 = tcp5.astype('float64')

df = pd.read_csv("~/Documentos/L1/Analise_comparativa/UDP/dataBase1(1).csv")
df2 = pd.read_csv("~/Documentos/L1/Analise_comparativa/UDP/dataBase1(2).csv")
df3 = pd.read_csv("~/Documentos/L1/Analise_comparativa/UDP/dataBase1(3).csv")
df4 = pd.read_csv("~/Documentos/L1/Analise_comparativa/UDP/dataBase1(4).csv")
df5 = pd.read_csv("~/Documentos/L1/Analise_comparativa/UDP/dataBase1(5).csv")

dados = df["data"]
udp1 = dados[2:]
udp1 = udp1.astype('float64')

dados = df2["data"]
udp2 = dados[2:]
udp2 = udp2.astype('float64')

dados = df3["data"]
udp3 = dados[2:]
udp3 = udp3.astype('float64')

dados = df4["data"]
udp4 = dados[2:]
udp4 = udp4.astype('float64')

dados = df5["data"]
udp5 = dados[2:]
udp5 = udp5.astype('float64')

# Média
mtcp1 = tcp1.mean()
mtcp2 = tcp2.mean()
mtcp3 = tcp3.mean()
mtcp4 = tcp4.mean()
mtcp5 = tcp5.mean()

mudp1 = udp1.mean()
mudp2 = udp2.mean()
mudp3 = udp3.mean()
mudp4 = udp4.mean()
mudp5 = udp5.mean()

# Desvio Padrao
dtcp1 = tcp1.std()
dtcp2 = tcp2.std()
dtcp3 = tcp3.std()
dtcp4 = tcp4.std()
dtcp5 = tcp5.std()

dudp1 = udp1.std()
dudp2 = udp2.std()
dudp3 = udp3.std()
dudp4 = udp4.std()
dudp5 = udp5.std()

print("1 Cliente:")
print("mediaTCP =",mtcp1)
print("devioPadraoTCP =",dtcp1)
print("mediaUDP =",mudp1)
print("desvioPadraoUDP =",dudp1)
print()

print("2 Clientes:")
print("mediaTCP =",mtcp2)
print("devioPadraoTCP =",dtcp2)
print("mediaUDP =",mudp2)
print("desvioPadraoUDP =",dudp2)
print()

print("3 Cliente2:")
print("mediaTCP =",mtcp3)
print("devioPadraoTCP =",dtcp3)
print("mediaUDP =",mudp3)
print("desvioPadraoUDP =",dudp3)
print()

print("4 Clientes:")
print("mediaTCP =",mtcp4)
print("devioPadraoTCP =",dtcp4)
print("mediaUDP =",mudp4)
print("desvioPadraoUDP =",dudp4)
print()

print("5 Clientes:")
print("mediaTCP =",mtcp5)
print("devioPadraoTCP =",dtcp5)
print("mediaUDP =",mudp5)
print("desvioPadraoUDP =",dudp5)
print()



medias = [mtcp1,mudp1]
configs = ['TCP1','UDP1']

plt.bar(configs,medias,color="blue")
plt.xticks(configs)
plt.ylabel('Tempo Médio (ms)')
plt.xlabel('Configuração')
plt.title('Configuração x Tempo médio (N=10.000)')
plt.savefig('Grafico1_1.png', format='png')
plt.show()

# medias2 = [mtcp2,mudp2]
# configs2 = ['TCP2','UDP2']

# plt.bar(configs2,medias2,color="blue")
# plt.xticks(configs2)
# plt.ylabel('Tempo Médio (ms)')
# plt.xlabel('Configuração')
# plt.title('Configuração x Tempo médio (N=10.000)')
# plt.savefig('Grafico2_1.png', format='png')
# plt.show()


# medias3 = [mtcp3,mudp3]
# configs3 = ['TCP3','UDP3']

# plt.bar(configs3,medias3,color="blue")
# plt.xticks(configs3)
# plt.ylabel('Tempo Médio (ms)')
# plt.xlabel('Configuração')
# plt.title('Configuração x Tempo médio (N=10.000)')
# plt.savefig('Grafico3_1.png', format='png')
# plt.show()


# medias4 = [mtcp4,mudp4]
# configs4 = ['TCP4','UDP4']

# plt.bar(configs4,medias4,color="blue")
# plt.xticks(configs4)
# plt.ylabel('Tempo Médio (ms)')
# plt.xlabel('Configuração')
# plt.title('Configuração x Tempo médio (N=10.000)')
# plt.savefig('Grafico4_1.png', format='png')
# plt.show()


# medias5 = [mtcp5,mudp5]
# configs5 = ['TCP5','UDP5']

# plt.bar(configs5,medias5,color="blue")
# plt.xticks(configs5)
# plt.ylabel('Tempo Médio (ms)')
# plt.xlabel('Configuração')
# plt.title('Configuração x Tempo médio (N=10.000)')
# plt.savefig('Grafico5_2.png', format='png')
# plt.show()