from pymongo import *
import os
import csv

hostname = os.environ['HOST']
port =  os.environ['PORT']
username = os.environ['USER']
password = os.environ['PASS']
database_name = os.environ['DB']

def csv_reader(file_obj):
    token_list = []
    reader = csv.reader(file_obj)
    for row in reader:
        obj = {
            "symbol": row[0],
            "address": row[1],
            "decimal": row[2]
        }
        token_list.append(obj)
    return token_list



# read *csv with token list
with open("tokenList.csv", "r") as f_obj:
    values = csv_reader(f_obj)

client = MongoClient(hostname, int(port))
db = client[database_name]
db.authenticate(username, password)

# add entries to collection
for i in values:
    db.token_list.insert_one(i)