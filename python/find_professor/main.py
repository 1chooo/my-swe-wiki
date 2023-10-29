import requests
from bs4 import BeautifulSoup
import pandas as pd
import re

prof_url = []
name = re.compile("\w* \w* *\w*")
good = re.compile("專長 *\/ *(\w*\/*\w*、*)*(\w*\/*\w*,* *)*(\w*\/*\w*\(*\w*\)*，* *)*\w*")
mail = re.compile("E(-)*mail: *\w*.*\w*@\w*.\w*.*.\w*.*.\w{2}")

edict = {'name':"", 'email':"",'speciality':""}

def makelist(whre, what):
    count = 0
    back = []
    for x in what:
        if count >= 49:
            break
        c = re.search(whre, str(x))
        if c:
            back.append(c.group())
            count += 1
        else:
            continue
    return back, count

url = "https://www.csie.ncu.edu.tw/department/member"
response = requests.get(url)
soup = BeautifulSoup(response.text, "html.parser")

a = soup.find_all("h4")
prof_name, counta = makelist(name, a)

b = soup.find_all("p")
prof_good, countb = makelist(good, b)
prof_email, countc = makelist(mail, b)

edict["name"] = prof_name
edict["email"] = prof_email
edict["speciality"] = prof_good

df = pd.DataFrame(edict)
df.to_csv("./prof_name.csv", encoding="utf-8-sig")
