import os
import pandas as pd
import numpy as np
import matplotlib.pyplot as plt
from scipy import stats


dirpath = os.curdir
outFile = open(dirpath+'/log.txt', 'w')

for numClients in range(1,6):
    numClients = str(numClients)
    file1 = dirpath+'/MOM/dataBase'+numClients+'.csv'
    file2 = dirpath+'/RPC/dataBase'+numClients+'.csv'
    df = pd.read_csv(file1)
    df2 = pd.read_csv(file2)

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


    # TESTING FOR NORMALITY
    stat, p = stats.shapiro(mom)   #shapiro-wilk test
    # k2, p = stats.normaltest(mom)  #test that combines skew and kurtosis
    alpha = 0.05
    print("p = {:g}".format(p))
    if p <= alpha:  # null hypothesis (H0): data comes from a normal distribution
        print("reject H0, not normal.")
    else:
        print("fail to reject H0, normal.")
    
    stat, p = stats.shapiro(rpc)
    # k2, p = stats.normaltest(rpc)
    alpha = 0.05
    print("p = {:g}".format(p))
    if p <= alpha:  # null hypothesis (H0): data comes from a normal distribution
        print("reject H0, not normal.")
    else:
        print("fail to reject H0, normal.")
       

    # # NULL HYPOTHESIS TESTING
    # # 2 sample t-test
    # t, pVal = stats.ttest_ind(mom,rpc)

    # Mann whitney U test
    U, pVal = stats.mannwhitneyu(mom,rpc)

    # print(pVal)
    if pVal < 0.05:
        logTest = "Reject NULL hypothesis - Significant differences exist between groups."
    if pVal > 0.05:
        logTest = "Accept NULL hypothesis - No significant difference between groups."


    if numClients > '1':
        outFile.write('\n\n')
    outFile.write(numClients + ' Cliente:\n')
    outFile.write('-----------------RPC-----------------\n')
    outFile.write('média = ' + str(mean_rpc) + '\n')
    outFile.write('desvio Padrão = ' + str(std_rpc) + '\n')
    outFile.write('-----------------MOM-----------------\n')
    outFile.write('média = ' + str(mean_mom) + '\n')
    outFile.write('desvio Padrão = ' + str(std_mom) + '\n')
    outFile.write('-------NULL HYPOTHESIS TESTING-------\n')
    outFile.write('p = ' + str(pVal) + '\n')
    outFile.write(logTest + '\n')
    outFile.write('-------------------------------------\n')
    outFile.close()

    medias = [mean_rpc,mean_mom]
    configs = ['RPC','MOM']

    plt.bar(configs,medias,color="blue")
    plt.xticks(configs)
    plt.ylabel('Tempo Médio (ms)')
    plt.xlabel('Tipo do Middleware')
    plt.title('Tipo do Middleware x Tempo médio (N=10.000)')
    plt.savefig(dirpath+'/Grafico'+numClients+'.png', format='png')
    # plt.show()

