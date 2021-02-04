### Weather Thing GO

## Getting Setup

1. git clone 
2. Create the numbers.txt file with one number per line 
    Ex:
        - +1-123-456-7891
        - +1-322-435-2342
3. Create beans.txt (holds weather data)
4. Setup environment vars 
    - APP_ID (AWS PinPoint Project ID)
    - API_URL (OpenWeatherAPI full URL EX: https://api.openweathermap.org/data/2.5/weather?q=<yourCity>,<yourContryCode>&units=imperial&APPID=<yourappid>
    - PYFILE (./pyText.py) 
5. go run main.go
    - /weather (grabs weather data, saves it to beans.txt, runs py script)
6. pyScript 
    - loops over numbers.txt file and sends a text to each with the weather data

