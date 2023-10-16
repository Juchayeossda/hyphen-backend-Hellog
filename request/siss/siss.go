package siss

import (
	"bytes"
	"encoding/json"
	"hyphen-backend-hellog/cerrors"
	"hyphen-backend-hellog/model"
	"hyphen-backend-hellog/verifier"
	"io"
	"mime/multipart"
	"net/http"
)

func RegisterImage(image []byte) (string, error) {

	// net/http package를 이용해서 siss서버에 요청

	// request body 설정하는 방법
	var requestBody bytes.Buffer
	multipartWriter := multipart.NewWriter(&requestBody)

	// file filed 설정
	part, err := multipartWriter.CreateFormFile("image", "image.txt")
	if err != nil {
		return "", err
	}

	// 파일 데이터를 MultiPart Form 데이터에 쓰기
	_, err = part.Write(image)
	if err != nil {
		return "", err
	}

	// MultiPart Form 마무리
	err = multipartWriter.Close()
	if err != nil {
		return "", err
	}

	// HTTP POST 요청 만들기
	targetURL := "http://101.101.217.155:8083/api/siss/images/image"
	req, err := http.NewRequest("POST", targetURL, &requestBody)

	// Content-Type 설정
	req.Header.Set("Content-Type", multipartWriter.FormDataContentType())

	// HTTP 클라이언트 생성
	client := &http.Client{}

	// 요청 보내기
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	// body parsing
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// json Unmarshal
	respJSON := new(model.SISSImageModel)
	err = json.Unmarshal(respBody, respJSON)
	if err != nil {
		return "", err
	}

	// 유효성 검사
	verifier.Validate(respJSON)

	// 응답 처리

	// 응답에 실패했으면
	if respJSON.Code != 201 {
		return "", &cerrors.ReqeustFailedError{ErrorMessage: respJSON.Message}
	}

	return respJSON.Data.ID, nil
}
