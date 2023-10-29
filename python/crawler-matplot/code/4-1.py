import requests
import csv
import pprint

good=["January","February","March","April","May","June","July","August","September","October","November","December"]

csv_file=open('../csv/theaters.csv','w+')
writer=csv.writer(csv_file)

for i in range(6,13):
    res=requests.get("https://www.boxofficemojo.com/calendar/2019-{:02d}-01/".format(i))
    data=res.text.split("<tr class=\"mojo-group-label\">")[1:]
    for j in data:
        if good[i-1] in j.split("</th>")[0].split('">')[1]:
            pprint.pprint(j.split("</th>")[0].split('">')[1])
            for q in j.split('<h3>')[:-1]:
                url=q.split('href="')[-1][:-2]
                #rint(url)
                try:
                    res1=requests.get("https://www.boxofficemojo.com/"+url)
                    val=res1.text.split("Widest Release")[1].split("</span>")[1].split("<span>")[-1].split(" ")[0].replace(",","")
                    #print(val)
                    print(["2019-{:02d}".format(i),val])
                    writer.writerow(["2019-{:02d}".format(i),val])
                except:
                    print("no given data")
        #pprint.pprint(i.split('<h3>')[0].split('href="')[1])
for i in range(1,6):
    res=requests.get("https://www.boxofficemojo.com/calendar/2020-{:02d}-01/".format(i))
    data=res.text.split("<tr class=\"mojo-group-label\">")[1:]
    for j in data:
        if good[i-1] in j.split("</th>")[0].split('">')[1]:
            pprint.pprint(j.split("</th>")[0].split('">')[1])
            for q in j.split('<h3>')[:-1]:
                url=q.split('href="')[-1][:-2]
                #print(url)
                try:
                    res1=requests.get("https://www.boxofficemojo.com/"+url)
                    val=res1.text.split("Widest Release")[1].split("</span>")[1].split("<span>")[-1].split(" ")[0].replace(",","")
                    #print(val)
                    print(["2020-{:02d}".format(i),val])
                    writer.writerow(["2020-{:02d}".format(i),val])
                except:
                    print("no given data")