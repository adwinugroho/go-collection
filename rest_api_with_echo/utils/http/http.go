package http

import (
	"bytes"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type (
	// Request request parameter
	Request struct {
		Protocol string
		Host     string
		Port     int
		Path     string
		Body     interface{}
		Param    string
	}
	// Response response body
	Response struct {
		Status        bool        `json:"status"`
		Code          int         `json:"code"`
		ErrorMessage  string      `json:"errMessage,omitempty"`
		Data          interface{} `json:"data,omitempty"`
		Message       interface{} `json:"message,omitempty"`
		Total         int64       `json:"total,omitempty"`
		TotalFiltered int64       `json:"totalFiltered,omitempty"`
	}
	ErrorResponse map[string]interface{}
)

func (r *ErrorResponse) Error() string {
	return ""
}

func Post(req *Request) (Response, error) {
	body, err := json.Marshal(req.Body)
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}

	resp, err := http.Post(fmt.Sprintf("%s://%s:%d%s", req.Protocol, req.Host, req.Port, req.Path), "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}

	defer resp.Body.Close()
	var response Response
	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return response, err
	}

	if resp.StatusCode == 200 {
		err = json.Unmarshal(responseBody, &response)
		if err != nil {
			fmt.Println(err)
			return response, err
		}
	} else {
		response = Response{Code: resp.StatusCode, Status: false, ErrorMessage: fmt.Sprintf("%s", responseBody)}
	}

	return response, nil
}

func PostDynamic(req *Request) (interface{}, error) {
	body, err := json.Marshal(req.Body)
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}

	resp, err := http.Post(fmt.Sprintf("%s://%s:%d%s", req.Protocol, req.Host, req.Port, req.Path), "application/json", bytes.NewBuffer(body))
	if err != nil {
		fmt.Println(err)
		return Response{}, err
	}

	defer resp.Body.Close()
	var response interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return response, err
	}
	return response, nil
}

//PostWithHeader call to specific url using POST methods with Headers
func PostWithHeader(req *Request, headers map[string]string) (map[string]interface{}, error) {
	body, err := json.Marshal(req.Body)
	if err != nil {
		log.Printf("[http]: error message: %v", err)
		return nil, err
	}
	log.Printf("[http]: req: %+v headers: %+v", req, headers)
	client := &http.Client{}
	if req.Protocol == "https" {
		tr := &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: false},
		}
		client = &http.Client{Transport: tr}
	}
	request, err := http.NewRequest("POST", fmt.Sprintf("%s://%s:%d%s", req.Protocol, req.Host, req.Port, req.Path), bytes.NewBuffer(body))
	if err != nil {
		log.Printf("[http]: error message: %v", err)
		return nil, err
	}
	for key, value := range headers {
		request.Header.Add(key, value)
	}
	resp, err := client.Do(request)
	if err != nil {
		log.Printf("[http]: error message: %v", err)
		return nil, err
	}
	if resp.StatusCode >= 500 {
		log.Printf("[http]: error message: %+v", resp)
		var errorResp ErrorResponse
		err = json.NewDecoder(resp.Body).Decode(&errorResp)
		if err != nil {
			log.Printf("[http]: error message: %v", err)
			return nil, err
		}
		return nil, &errorResp
	}
	defer resp.Body.Close()
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Printf("[http]: error message: %v", err)
		return nil, err
	}
	return response, nil
}

//GetWithHeaders call to specific url using GET methods with Headers
func GetWithHeaders(req *Request, headers map[string]string) (*map[string]interface{}, error) {
	var client http.Client
	var url string
	if req.Port == 0 {
		url = fmt.Sprintf("%s://%s%s%s", req.Protocol, req.Host, req.Path, req.Param)
	} else {
		url = fmt.Sprintf("%s://%s:%d%s%s", req.Protocol, req.Host, req.Port, req.Path, req.Param)
	}
	if req.Protocol == "" {
		url = fmt.Sprintf("%s:%d%s%s", req.Host, req.Port, req.Path, req.Param)
	}
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Printf("[http]: error message: %v", err)
		return nil, err
	}
	for key, value := range headers {
		request.Header.Add(key, value)
	}
	log.Printf("[http]: GET URL: %v headers: %+v", url, headers)
	resp, err := client.Do(request)
	if err != nil {
		log.Printf("[http]: error message: %v", err)
		return nil, err
	}
	defer resp.Body.Close()
	var response map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		log.Printf("[http]: error message: %v", err)
		return nil, err
	}
	return &response, nil
}
