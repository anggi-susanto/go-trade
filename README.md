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
### 1 calculate max profit
Create http request to POST `http://localhost:8080/max-trade`
using this body:
```
{
    "file_url":"https://gist.githubusercontent.com/Jekiwijaya/c72c2de532203965bf818e5a4e5e43e3/raw/2631344d08b044a4b833caeab8a42486b87cc19a/gistfile1.txt"
}
```

and the response will be like this:
```
{
  "max_profit":299992438
}
```
### 2 unique string

reate http request to POST `http://localhost:8080/unique-string`
using this body:
```
{
    "file_url":"https://gist.githubusercontent.com/Jekiwijaya/0b85de3b9ff551a879896dd78256e9b8/raw/e9d58da5d4df913ad62e6e8dd83c936090ee6ef4/gistfile1.txt"
}
```

and the response will be like this:
```
{
    "first_occurence": "sfgxclryzidpuvejaqbtwmhkno",
    "smallest_lexicographical_order": "abcdefghijklmnopqrstuvwxyz"
}
```
