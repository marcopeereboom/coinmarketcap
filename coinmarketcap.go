package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type Ticker struct {
	Id                   string `json:"id"`
	Name                 string `json:"name"`
	Symbol               string `json:"symbol"`
	Rank                 string `json:"rank"`
	PriceUsd             string `json:"price_usd"`
	PriceBtc             string `json:"price_btc"`
	TwentyfourHourVolume string `json:"24h_volume_usd"`
	MarketCapUsd         string `json:"market_cap_usd"`
	AvailableSupply      string `json:"available_supply"`
	TotalSupply          string `json:"total_supply"`
	PercentChange1Hour   string `json:"percent_change_1h"`
	PercentChange24Hour  string `json:"percent_change_24h"`
	PercentChange7Day    string `json:"percent_change_7d"`
	LastUpdated          string `json:"last_updated"`
}

func (t Ticker) String() string {
	usd, _ := strconv.ParseFloat(t.PriceUsd, 64)
	btc, _ := strconv.ParseFloat(t.PriceBtc, 64)
	cap, _ := strconv.ParseFloat(t.MarketCapUsd, 64)
	p := " "
	if cap > 1000 {
		cap /= 1000
		p = "K"
	}
	if cap > 1000 {
		cap /= 1000
		p = "M"
	}
	if cap > 1000 {
		cap /= 1000
		p = "B"
	}
	return fmt.Sprintf("|%-10v|%5v|%15.3f|%15.08f|%14.03f%v|",
		t.Symbol, t.Rank, usd, btc, cap, p)
}

func printSeparator() {
	fmt.Printf("+%-10v+%-5v+%-15v+%-15v+%-15v+\n",
		dash10, dash5, dash15, dash15, dash15)
}

var (
	tickerRoute = "https://api.coinmarketcap.com/v1/ticker/"

	dash5  = strings.Repeat("-", 5)
	dash10 = strings.Repeat("-", 10)
	dash15 = strings.Repeat("-", 15)
)

func ticker(j bool, max int, re *regexp.Regexp) error {
	r, err := http.Get(tickerRoute)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	d := json.NewDecoder(r.Body)

	var t []Ticker
	err = d.Decode(&t)
	if err != nil {
		return fmt.Errorf("Could node unmarshal Ticker: %v", err)
	}
	if !j {
		printSeparator()
		fmt.Printf("|%-10v|%-5v|%-15v|%-15v|%-15v|\n",
			"Symbol", "Rank", "USD", "BTC", "Market cap USD")
		printSeparator()
	}
	count := 0
	for _, v := range t {
		if re != nil {
			if !re.MatchString(v.Symbol) {
				continue
			}
		}
		if j {
			// Just reencode it
			jo, err := json.MarshalIndent(v, "", "  ")
			if err != nil {
				return err
			}
			fmt.Printf("%v\n", string(jo))
		} else {
			fmt.Printf("%v\n", v)
		}
		count++
		if count >= max {
			break
		}
	}
	if !j {
		printSeparator()
	}

	return nil
}

func _main() error {
	duration := flag.Int("duration", 0, "time in seconds before update, <=0 is a single shot")
	filter := flag.String("filter", "", "regex filter on ticker")
	max := flag.Int("max", 50, "display max N currencies")
	useJson := flag.Bool("json", false, "print JSON")
	help := flag.Bool("h", false, "help")
	flag.Parse()

	if *help {
		flag.PrintDefaults()
		return nil
	}

	if *duration <= 0 {
		*duration = 0
	}

	var (
		re  *regexp.Regexp
		err error
	)
	if *filter != "" {
		re, err = regexp.Compile(*filter)
		if err != nil {
			return err
		}
	}

	for {
		err = ticker(*useJson, *max, re)
		if err != nil {
			return err
		}
		if *duration == 0 {
			break
		} else {
			time.Sleep(time.Duration(*duration) * time.Second)
		}
	}

	return nil
}

func main() {
	err := _main()
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
		os.Exit(1)
	}
}
