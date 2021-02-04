### Weather Thing GO

## Getting Setup

1. git clone 
2. Create the numbers.txt file with one number per line 
    Ex:
        +1-123-456-7891 
        \n
        +1-322-435-2342
3. Create beans.txt (holds weather data)
4. Setup environment vars 
    - APP_ID (AWS PinPoint Project ID)
    - API_URL (OpenWeatherAPI full URL EX: https://api.openweathermap.org/data/2.5/weather?q=yourCity,yourContryCode&units=imperial&APPID=yourappid
    - PYFILE (./pyText.py) 
5. python3 -m venv venv => . ./venv/bin/activate
6. pip install -r requirements.txt or pip install boto3 


### Next
1. go run main.go (Run Go server)
    Endpoints:
        - /weather
            1. Grabs data from OpenWeatherApi
            2. Saves to txt file
            3. Runs pyScript to send texts
            4. Returns weather data

