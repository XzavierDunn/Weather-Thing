#!/usr/bin/env python3

import boto3
from botocore.exceptions import ClientError
import os
import json

region = "us-east-1"
messageType = "TRANSACTIONAL"
applicationId = os.getenv("APP_ID")
profile = 'default' # AWS Credentials profile

originationNumber = ""
registeredKeyword = ""
senderId = "ya boi"


def getMessage():
    with open("./beans.txt", 'r') as mes:
        x = json.load(mes)
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
            f"""
City: {x["name"]}\n
Current Weather: {weather[1]}
Description: {weather[2]}\n
Feels like: {main["feels_like"]}
Actual Temperature: {main["temp"]}
Max Temp: {main["temp_max"]}
Min Temp: {main["temp_min"]}
Humidity: {main["humidity"]}"""

        return final


def sendText(num):
    profile = boto3.session.Session(
        profile_name=profile)
    client = profile.client('pinpoint', region_name=region)
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
        print('Error:', e.response['Error']['Message'])
    else:
        print("Message sent")


with open("numbers.txt", 'r') as nums:
    nums = nums.readlines()
    for i in nums:
        sendText(i)
