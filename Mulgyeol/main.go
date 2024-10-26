package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

// 요청 바디 구조체
type RequestBody struct {
    GrantType string `json:"grant_type"`
    AppKey    string `json:"appkey"`
    SecretKey string `json:"secretkey"`
}

// 응답 바디 구조체
type ResponseBody struct {
    ApprovalKey string `json:"approval_key"`
}

func getApprovalKey() (string, error) {
    // API URL 설정 (모의투자 환경)
    url := "https://openapivts.koreainvestment.com:29443/oauth2/Approval"

    // 요청 바디 생성
    body := RequestBody{
        GrantType: "client_credentials",
        AppKey:    "PSg5dctL9dKPo727J13Ur405OSXXXXXXXXXX", // 발급받은 AppKey 사용
        SecretKey: "yo2t8zS68zpdjGuWvFyM9VikjXE0i0CbgPEamnqPA00G0bIfrd...", // 발급받은 SecretKey 사용
    }

    // JSON 직렬화
    jsonData, _ := json.Marshal(body)

    // HTTP 요청 생성
    req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
    if err != nil {
        return "", err
    }

    req.Header.Set("Content-Type", "application/json; charset=utf-8")

    // 클라이언트 생성 및 요청 실행
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    // 응답 바디 읽기
    responseData, _ := io.ReadAll(resp.Body)

    // 응답 데이터 파싱
    var responseBody ResponseBody
    if err := json.Unmarshal(responseData, &responseBody); err != nil {
        return "", err
    }

    return responseBody.ApprovalKey, nil
}

func main() {
    approvalKey, err := getApprovalKey()
    if err != nil {
        fmt.Println("Error getting approval key:", err)
        return
    }
    fmt.Println("Approval Key:", approvalKey)
}
