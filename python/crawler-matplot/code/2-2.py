import csv
import matplotlib.pyplot as plt

csv_file=open('../csv/number_of_films.csv','r')

rows = csv.reader(csv_file)

x=[]
y=[]

for i in rows:
    x.append(i[0])
    y.append(int(i[1].replace(",","")))

print(x,y)

plt.bar(x, y)

plt.show()