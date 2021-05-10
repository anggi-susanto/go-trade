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
    "max_profit": 299994700,
    "buy_hour": 11789,
    "buy_price": 100004284,
    "sell_hour": 19090,
    "sell_price": 399998984
}
```
### 2 unique string

create http request to POST `http://localhost:8080/unique-string`
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
### 3 find H(n) of six

create http request to POST `http://localhost:8080/factor-six`
using this body:
```
{
    "number":262144
}
```

and the response will be like this:
```
{
    "mathched_factor_count": 13208
}
```

*note : on step 3 needed to refactor, it's the simplest way to find the matching factor. Need further refactor because submission time limit.