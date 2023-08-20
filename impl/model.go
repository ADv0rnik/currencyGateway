package impl

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

type CurrencyRes struct {
	USDIn  string `json:"usd_in"`
	USDOut string `json:"usd_out"`
	Street string `json:"street"`
	Name   string `json:"name"`
}

type BaseGetRes struct {
	*http.Response
}

type Context struct {
	Url     string
	Headers http.Header
}

func NewCurrencyRequest(args Params) (Context, error) {
	headers := http.Header{}
	headers.Add("Accept", "*/*")
	queryParams := url.Values{}
	queryParams.Add("city", args.City)
	return Context{
		Url:     fmt.Sprintf("%s?city=%s", args.BaseUrl, queryParams.Get("city")),
		Headers: headers,
	}, nil
}

func arrayToMap(obj []byte) (CurrencyRes, error) {
	var mapData []map[string]string
	if err := json.Unmarshal(obj, &mapData); err != nil {
		errLogger().Println("Error while parsing data: ", err)
	}

	curMap := CurrencyRes{}
	for _, item := range mapData {
		if item["street"] == "Тимофеенко" {
			curMap.USDIn = item["USD_in"]
			curMap.USDOut = item["USD_out"]
			curMap.Street = item["street"]
			curMap.Name = item["name"]
		}
	}
	return curMap, nil
}
