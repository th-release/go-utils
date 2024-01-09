package utils

import (
	"fmt"
	"net/url"
)

func ParseQueryParameters(link string) (map[string]string, error) {
	// URL 파싱
	parsedURL, err := url.Parse(link)
	if err != nil {
		return nil, fmt.Errorf("URL parsing error: %v", err)
	}

	// 쿼리 매개변수 가져오기
	queryParams := parsedURL.Query()

	// 쿼리 매개변수를 map으로 변환
	paramsMap := make(map[string]string)
	for key, values := range queryParams {
		if len(values) > 0 {
			paramsMap[key] = values[0]
		}
	}

	return paramsMap, nil
}
