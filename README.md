# go-trade
this is simple golang implementation using simple REST API to calculate max difference from remote file

## 1 installing dependencies
```
make install
```

## 2 testing the code
```
make test
```

## 3 run te server
```
go run main.go
```
or 
```
make build
./go-trade
```

## How to use the endpoint
Create http request to POST `http://localhost:8080/calculate`
using this body:
```
{
    "file_url":"https://gist.githubusercontent.com/Jekiwijaya/c72c2de532203965bf818e5a4e5e43e3/raw/2631344d08b044a4b833caeab8a42486b87cc19a/gistfile1.txt"
}
```

and the response will be like this:
```
{"max_profit":299992438}
```
