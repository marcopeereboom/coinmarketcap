# coinmarketcap
Coinmarketcap CLI tool to print current market prices.

### installation
```
go get github.com/marcopeereboom/coinmarketcap
```

### flags
```
  -duration int
        time in seconds before update, <=0 is a single shot
  -filter string
        regex filter on ticker
  -h    help
  -json
        print JSON
  -max int
        display max N currencies (default 50)
```

### examples

JSON, single shot, Decred
```
$ coinmarketcap -json -filter="DCR"
{
  "id": "decred",
  "name": "Decred",
  "symbol": "DCR",
  "rank": "42",
  "price_usd": "61.5492",
  "price_btc": "0.00757582",
  "24h_volume_usd": "21375400.0",
  "market_cap_usd": "438249281.0",
  "available_supply": "7120308.0",
  "total_supply": "7540308.0",
  "percent_change_1h": "0.41",
  "percent_change_24h": "4.81",
  "percent_change_7d": "39.7",
  "last_updated": "1523636948"
}
```

Regular, single shot, Decred, Bitcoin and Litecoin
```
coinmarketcap -filter="DCR|BTC$|LTC"
+----------+-----+---------------+---------------+---------------+
|Symbol    |Rank |USD            |BTC            |Market cap USD |
+----------+-----+---------------+---------------+---------------+
|BTC       |    1|       8100.360|     1.00000000|       137.507B|
|LTC       |    5|        131.001|     0.01612430|         7.345B|
|DCR       |   42|         61.549|     0.00757582|       438.249M|
+----------+-----+---------------+---------------+---------------+
```

Regular, top 5
```
$ coinmarketcap -max=5                
+----------+-----+---------------+---------------+---------------+
|Symbol    |Rank |USD            |BTC            |Market cap USD |
+----------+-----+---------------+---------------+---------------+
|BTC       |    1|       8100.360|     1.00000000|       137.507B|
|ETH       |    2|        515.548|     0.06345650|        50.934B|
|XRP       |    3|          0.678|     0.00008349|        26.539B|
|BCH       |    4|        763.077|     0.09392380|        13.027B|
|LTC       |    5|        131.001|     0.01612430|         7.345B|
+----------+-----+---------------+---------------+---------------+
```
