package services

import (
	"context"
	"encoding/json"
)

type beefServer struct {
}

func NewBeefServer() BeefServer {
	return &beefServer{}
}

func (b *beefServer) Summary(context.Context, *EmptyRequest) (*BeefSummaryResponse, error) {
	beefs := GetBeefs()
	res := BeefSummaryResponse{
		Data: beefs,
	}
	return &res, nil
}

func (b *beefServer) mustEmbedUnimplementedBeefServer() {
	panic("unimplemented")
}

func GetBeefs() []byte {
	url := "https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text"
	respStr := getString(url)

	words := getWordsSlice(respStr)
	wordsFreq := countWords(words)
	data := map[string]interface{}{
		"beef": wordsFreq,
	}
	dataByte, err := json.Marshal(data)
	notErr(err)
	return dataByte
}
