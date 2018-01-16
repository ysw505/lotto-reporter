package main

import (
	"fmt"
	"net/http"
	"lotto-reporter/model"
	"io/ioutil"
	"encoding/json"
	"flag"
)

func main() {
	isRecent := flag.Bool("recent", false, "a bool")
	flag.Parse()

	var result *model.LottoResult
	var err error

	if *isRecent {
		result , err = getLottoResult(0);
	} else {
		result , err = getLottoResult(123);
	}

	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(result)
	}

}

func getLottoResult(turn int32) ( *model.LottoResult, error) {
	url := "http://www.nlotto.co.kr/common.do?method=getLottoNumber&drwNo="

	if turn != 0 {
		url = fmt.Sprintf("http://www.nlotto.co.kr/common.do?method=getLottoNumber&drwNo=%d", turn)
	}

	resp, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	respByte, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	var result model.LottoResult

	json.Unmarshal(respByte, &result)

	if err != nil {
		return nil, err
	}

	return &result, nil
}