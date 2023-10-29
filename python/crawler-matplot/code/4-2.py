import csv
import matplotlib.pyplot as plt

csv_file=open('../csv/theaters.csv','r')

rows = csv.reader(csv_file)

xLabel = []
yLabel = [str(i) for i in range(5000,0,-1000)]

data=[ [ 0 for j in range(12)] for i in range(len(yLabel))]

for i in rows:
    if i[0] not in xLabel:
        xLabel.append(i[0])
    inp=-1
    for j in range(len(xLabel)):
        if xLabel[j]==i[0]:
            inp=j
    data[4-int(i[1])//1000][inp]+=1



fig = plt.figure()
ax = fig.add_subplot(111)
ax.set_yticks(range(len(yLabel)))
ax.set_yticklabels(yLabel)
ax.set_xticks(range(len(xLabel)))
ax.set_xticklabels(xLabel)

for i in range(len(yLabel)):
    for j in range(len(xLabel)):
        text = ax.text(j, i, data[i][j],ha="center", va="center", color="b")

im = ax.imshow(data, cmap=plt.cm.hot_r)
plt.colorbar(im)

plt.show()