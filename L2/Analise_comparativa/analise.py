import os
import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
from scipy import stats


dirpath = os.getcwd() + '/Analise_comparativa'
outFile = open(dirpath+'/log.txt', 'w')

for numClients in range(1,6):
    numClients = str(numClients)
    file1 = dirpath+'/MOM/dataBase'+numClients+'.csv'
    file2 = dirpath+'/MOO/dataBase'+numClients+'.csv'
    df = pd.read_csv(file1)
    df2 = pd.read_csv(file2)

    dados = df["data"]
    mom = dados[1:-1]
    mom = mom.astype('float64')

    dados = df2["data"]
    moo = dados[1:-1]
    moo = moo.astype('float64')

    # Média
    mean_mom = mom.mean()
    mean_moo = moo.mean()

    # Desvio Padrao
    std_mom = mom.std()
    std_moo = moo.std()

    #t-value e p-value (2 sample t-test)
    t2, p2 = stats.ttest_ind(mom,moo)

    if numClients > '1':
        outFile.write('\n\n')
    outFile.write(numClients + ' Cliente:\n')
    outFile.write('-----------------MOM-----------------\n')
    outFile.write('média = ' + str(mean_mom) + '\n')
    outFile.write('desvio Padrão = ' + str(std_mom) + '\n')
    outFile.write('-----------------MOO-----------------\n')
    outFile.write('média = ' + str(mean_moo) + '\n')
    outFile.write('desvio Padrão = ' + str(std_moo) + '\n')
    outFile.write('-----------------t-Test--------------\n')
    outFile.write('t = ' + str(t2) + '\n')
    outFile.write('p = ' + str(p2) + '\n')
    outFile.write('-------------------------------------\n')

    medias = [mean_mom,mean_moo]
    configs = ['MOM','MOO']

    plt.bar(configs,medias,color="blue")
    plt.xticks(configs)
    plt.ylabel('Tempo Médio (ms)')
    plt.xlabel('Tipo do Middleware')
    plt.title('Tipo do Middleware x Tempo médio (N=10.000)')
    plt.savefig(dirpath+'/Grafico'+numClients+'.png', format='png')
    # plt.show()

outFile.close()
