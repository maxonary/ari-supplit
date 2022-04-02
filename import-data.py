import csv
import pandas as pd



from flask import Flask

app = Flask(__name__)

@app.route("/")
def hello_world():

    df = pd.read_csv (r"./material/sneaker-numbers.csv")
    df.to_json (r"./material/sneaker.json")
    return "<p>Hello, World!</p>"

#with open("./material/sneaker-numbers.csv") as custfile:
   # rows = csv.reader(custfile,delimiter=',')
  #  for r in rows:
 #       r = r.replace(u'\xa0', u' ')
        #print(r)

