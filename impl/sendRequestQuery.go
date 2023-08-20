package impl

import (
	"io"
	"net/http"
)

func SendRequest(args map[string]string) (CurrencyRes, error) {
	params, err := RetrieveParams(args)
	if err != nil {
		errLogger().Println("No input params")
	}

	curReq, err := NewCurrencyRequest(params)
	resp, err := http.Get(curReq.Url)
	if err != nil {
		errLogger().Println("Bad request: ", err)
	}
	body, err := io.ReadAll(resp.Body)
	bodyString := string(body)

	data, err := arrayToMap([]byte(bodyString))
	if err != nil {
		errLogger().Println("Error while parsing data: ", err)
	}
	//xType := fmt.Sprintf("%T", data)
	//fmt.Println(xType)
	//jsonData, err := json.Marshal(data)
	//fmt.Printf("%s", jsonData)
	return data, nil
}
