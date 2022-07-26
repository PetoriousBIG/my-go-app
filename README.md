# my-go-app

To build:
`docker build -t go-multi-stage .`

To run:
`docker run -p 9090:9090 go-multi-stage`

The API utilizes the following endpoint: 
`/At-A-Glance/{id:[A-Z]{3}}`
where id is any of the Alpha-3 codes in the csv file.

Testing the API endpoint:
`curl localhost/At-A-Glance/COL | jq`

Should return the following JSON:
```
{
  "country": "COL",
  "country_name": "Colombia",
  "country_id": 170,
  "average_latitude": 4,
  "average_longitude": -72
}
```