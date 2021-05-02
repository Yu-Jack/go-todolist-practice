package network

import (
	"encoding/json"
	"jack-test/internal/usecase"
)

type Response struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

var messageMap = map[int]string{
	0:   "success",
	100: "Failed",
	200: "Invalid format.",
}

func (response *Response) convertToBytes() []byte {
	responseJson, _ := json.MarshalIndent(response, "", "    ")
	return responseJson
}

func ToJson(response interface{}) []byte {
	responseJson, _ := json.MarshalIndent(response, "", "    ")
	return responseJson

}

func buildResponse(status int) (response Response) {
	return Response{
		Status:  status,
		Message: messageMap[status],
	}
}

func NewSuccessResponse() (response Response) {
	return buildResponse(0)
}

func NewFailedResponse(status int) (response Response) {
	return buildResponse(status)
}

type Network struct {
	usecaseapi usecase.Usecaseapi
}

func NewNetwork(usecaseapi usecase.Usecaseapi) (network *Network) {
	return &Network{
		usecaseapi: usecaseapi,
	}
}
