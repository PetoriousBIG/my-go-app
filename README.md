# my-go-app

To build:
`docker build -t go-multi-stage .`

To run:
`docker run -p 9090:9090 go-multi-stage`

Then open a browser to:
`localhost:9090/ping`

And you should see:
`Pong`