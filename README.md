# globant

##Setup Instruction:
1. Look for appsetting.json file :
2. Change port number by editing the port section of appsetting.json in case you want a different port number
3. Change the API key incase it has change by editing the key section on the appsetting.json


##To run the test cases:
1. Open the terminal or the command prompt
2. Navigate to the tests  folder in the terminal
3. Run this command   go test configuration_test.go midleware_test.go

##To run the program: 
1. Open the Terminal or command prompt
2. Navigate to the globant folder
3. Run this command go run main.go
4. Either open Postman or browser
Enter this url localhost:8080/weatherforecast/city_name,county_code  to get the forecast for the city of your choice
Example :
 localhost:8080/weatherforecast/Toronto,ca
 
6. You can also get forecast for the next 5 days and 3 horse forecast for each day using the url below:
localhost:808/weatherforecast/city_name,country_code/cont
 localhost:8080/weatherforecast/Toronto,ca/2
 
 
                 
 
 
 

