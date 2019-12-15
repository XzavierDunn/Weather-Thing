import boto3
from botocore.exceptions import ClientError
import os
import json

region = "us-east-1"
messageType = "TRANSACTIONAL"
applicationId = os.getenv("APP_ID")

originationNumber = ""
registeredKeyword = ""
senderId = ""

def getMessage():
    with open("./beans.txt", 'r') as mes:
        weather = []
        final = []

        # Grabs current weather and description
        y = dict(x["weather"][0])
        for k, v in enumerate(y):
            weather.append(y[v])

        # Grabs main points
        main = {}
        y = dict(x["main"])
        for k, v in enumerate(y):
            main[v] = y[v]
        
        final = \
f"""Xzavier's boujee weather app\n
City: {x["name"]}\n
Current Weather: {weather[1]} - {weather[2]}\n
Feels like: {main["feels_like"]}
Actual Temperature: {main["temp"]}
Max Temp: {main["temp_max"]}
Min Temp: {main["temp_min"]}
Humidity: {main["humidity"]}"""

        return final


def sendText(num):
    client = boto3.client('pinpoint',region_name=region)
    try:
        response = client.send_messages(
            ApplicationId=applicationId,
            MessageRequest={
                'Addresses': {
                    num: {
                        'ChannelType': 'SMS'
                    }
                },
                'MessageConfiguration': {
                    'SMSMessage': {
                        'Body': getMessage(),
                        'Keyword': registeredKeyword,
                        'MessageType': messageType,
                        'OriginationNumber': originationNumber,
                        'SenderId': senderId
                    }
                }
            }
        )
    except ClientError as e:
        print(e.response['Error']['Message'])
    else:
        print("Message sent")
