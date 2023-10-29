import csv
import requests
import pprint


exchange={
    'Jan':'01','Feb':'02','Mar':'03','Apr':'04','May':'05','Jun':'06','Jul':'07','Aug':'08','Sep':'09','Oct':'10','Nov':'11','Dec':'12'
}



csv_file=open('../csv/Australia.csv','w+')
writer=csv.writer(csv_file)

res=requests.get("https://www.boxofficemojo.com/weekend/by-year/2019/?area=AU")
good=['Jun','Jul','Aug','Sep','Oct','Nov','Dec']
for i in res.text.split('<tr>')[2:][::-1]:
    date=i.split("</a>")[0].split("\">")[-1]
    val=i.split("</td>")[3].split(">$")[1].replace(",","")
    for j in good:
        if j in date.split('-')[0]:
            date="2019-"+exchange[date[:3]]+"-"+"{:02d}".format(int(date.split('-')[0].split(' ')[1]))
            print(date,val)
            writer.writerow([date,val])

good=['Jan','Feb','Mar','Apr','May']

res=requests.get("https://www.boxofficemojo.com/weekend/by-year/2020/?area=AU")
for i in res.text.split('<tr>')[2:][::-1]:
    date=i.split("</a>")[0].split("\">")[-1]
    val=i.split("</td>")[3].split(">$")[1].replace(",","")
    for j in good:
        if j in date.split('-')[0]:
            date="2020-"+exchange[date[:3]]+"-"+"{:02d}".format(int(date.split('-')[0].split(' ')[1]))
            print(date,val)
            writer.writerow([date,val])






csv_file=open('../csv/Japan.csv','w+')
writer=csv.writer(csv_file)

res=requests.get("https://www.boxofficemojo.com/weekend/by-year/2019/?area=JP")
good=['Jun','Jul','Aug','Sep','Oct','Nov','Dec']
for i in res.text.split('<tr>')[2:][::-1]:
    date=i.split("</a>")[0].split("\">")[-1]
    val=i.split("</td>")[3].split(">$")[1].replace(",","")
    for j in good:
        if j in date.split('-')[0]:
            date="2019-"+exchange[date[:3]]+"-"+"{:02d}".format(int(date.split('-')[0].split(' ')[1]))
            print(date,val)
            writer.writerow([date,val])

good=['Jan','Feb','Mar','Apr','May']

res=requests.get("https://www.boxofficemojo.com/weekend/by-year/2020/?area=JP")
for i in res.text.split('<tr>')[2:][::-1]:
    date=i.split("</a>")[0].split("\">")[-1]
    val=i.split("</td>")[3].split(">$")[1].replace(",","")
    for j in good:
        if j in date.split('-')[0]:
            date="2020-"+exchange[date[:3]]+"-"+"{:02d}".format(int(date.split('-')[0].split(' ')[1]))
            print(date,val)
            writer.writerow([date,val])







csv_file=open('../csv/Germany.csv','w+')
writer=csv.writer(csv_file)

res=requests.get("https://www.boxofficemojo.com/weekend/by-year/2019/?area=DE")
good=['Jun','Jul','Aug','Sep','Oct','Nov','Dec']
for i in res.text.split('<tr>')[2:][::-1]:
    date=i.split("</a>")[0].split("\">")[-1]
    val=i.split("</td>")[3].split(">$")[1].replace(",","")
    for j in good:
        if j in date.split('-')[0]:
            date="2019-"+exchange[date[:3]]+"-"+"{:02d}".format(int(date.split('-')[0].split(' ')[1]))
            print(date,val)
            writer.writerow([date,val])

good=['Jan','Feb','Mar','Apr','May']

res=requests.get("https://www.boxofficemojo.com/weekend/by-year/2020/?area=DE")
for i in res.text.split('<tr>')[2:][::-1]:
    date=i.split("</a>")[0].split("\">")[-1]
    val=i.split("</td>")[3].split(">$")[1].replace(",","")
    for j in good:
        if j in date.split('-')[0]:
            date="2020-"+exchange[date[:3]]+"-"+"{:02d}".format(int(date.split('-')[0].split(' ')[1]))
            print(date,val)
            writer.writerow([date,val])






csv_file=open('../csv/NorthAmerica.csv','w+')
writer=csv.writer(csv_file)

res=requests.get("https://www.boxofficemojo.com/daily/2019/?view=year")
good=['Jun','Jul','Aug','Sep','Oct','Nov','Dec']
for i in res.text.split('<tr>')[2:][::-1]:
    date=i.split("</a>")[0].split("\">")[-1]
    val=i.split("</td>")[-2].split(">$")[1].replace(",","")
    for j in good:
        if j in date.split(' ')[0]:
            date="2019-"+exchange[date[:3]]+"-"+"{:02d}".format(int(date.split(',')[0].split(' ')[1]))
            print(date,val)
            writer.writerow([date,val])

res=requests.get("https://www.boxofficemojo.com/date/")
good=['Jan','Feb','Mar','Apr','May']
for i in res.text.split('<tr>')[2:][::-1]:
    date=i.split("</a>")[0].split("\">")[-1]
    val=i.split("</td>")[-2].split(">$")[1].replace(",","")
    for j in good:
        if j in date.split(' ')[0]:
            date="2020-"+exchange[date[:3]]+"-"+"{:02d}".format(int(date.split(',')[0].split(' ')[1]))
            print(date,val)
            writer.writerow([date,val])
