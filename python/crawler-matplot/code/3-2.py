import csv
import matplotlib.pyplot as plt

AU=open('../csv/Australia.csv','r')
JP=open('../csv/Japan.csv','r')
DE=open('../csv/Germany.csv','r')
NA=open('../csv/NorthAmerica.csv','r')

AU_rows = csv.reader(AU)
JP_rows = csv.reader(JP)
DE_rows = csv.reader(DE)
NA_rows = csv.reader(NA)

month=[]
money=[]

for i in NA_rows:
    month.append(i[0])
    money.append(int(i[1]))

plt.plot(month,money,color = 'g', label="NA")

month=[]
money=[]

for i in AU_rows:
    month.append(i[0])
    money.append(int(i[1]))

plt.plot(month,money,color = 'b', label="AU")

month=[]
money=[]

for i in JP_rows:
    month.append(i[0])
    money.append(int(i[1]))

plt.plot(month,money,color = 'r', label="JP")

month=[]
money=[]

for i in DE_rows:
    month.append(i[0])
    money.append(int(i[1]))

plt.plot(month,money,color = 'y', label="DE")



plt.show()