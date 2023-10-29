import csv
import requests

data={
    "realm":"title",
    "title_type":"feature",
    "release_date-min":"2015-02",
    "release_date-max":"2015-02",
    "view":"detailed",
    "count":"50",
    "sort":"moviemeter,asc"
}



csv_file=open('../csv/number_of_films.csv','w+')

writer=csv.writer(csv_file)

for i in range(2015,2021):
    for j in range(1,13):
        date="{}-{}".format(str(i),str(j))
        data["release_date-min"]=date
        data["release_date-max"]=date
        res=requests.post("https://www.imdb.com/search/title/",data=data)
        num=res.text.split('<div class="desc">')[1].split(" titles.</span>")[0].split(" ")[-1]
        writer.writerow([date,num.replace(",","")])