# GoLang Assessment

## Requirements
  - Go 1.19
  - Gin 1.8.2
 
 ## Description
  Based on the assessment, an api is created to calculate the start and end. You can see the server configuration in `main.go` file. The login is written 
 in flights package. 
  In `flights.go` you see a function which is named `GetStartAndEnd`. We have two struct, one of them is for the json body and the other one is for the 
  algorithm. Airports struct map to flight struct. We have two map named `starts` and `end`. I created two map, starts is for all starting points of the
  flight and ends is for ending points for the flights.
  In the next step these two maps compare to each other and duplicated values in maps tagged with `true` value. Non duplicated values would be the answer.
  
  The application runs on localhost port 8080.
  
  You can send your request like below:
  
  <img width="850" alt="image" src="https://user-images.githubusercontent.com/41889003/211646576-36ce0dbe-b8fc-4121-8d93-ffb637c3c7f5.png">

cURL request would be like below:
```bash
curl --location --request POST '127.0.0.1:8080/calculate/' \
--header 'Content-Type: application/json' \
--data-raw '{
    "paths": "[[\"IND\", \"EWR\"],[\"SFO\", \"ATL\"],[\"GSO\", \"IND\"],[\"ATL\", \"GSO\"]]"
}'
```
 
