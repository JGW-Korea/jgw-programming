package main

import (
	"fmt"
	"os"
)

// 오류 구분 #1. 대응 가능한 일반적인 오류
func loadUserConfig(path string) ([]byte, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		// 파일을 읽지 못했지만 호출한 쪽에서
		// 기본 설정 파일 사용 등의 대응이 가능하므로 error 반환
		return nil, err
	}

	return data, nil
}

func main() {
	data, err := loadUserConfig("./config.json")
	if err != nil {
		fmt.Println("설정 파일을 읽지 못해 기본 설정을 사용합니다.")

		data = []byte(`{"theme":"default"}`)
	}

	fmt.Println(string(data))

	// 오류에 대응한 뒤 이후 작업을 계속 수행
	fmt.Println("프로그램 실행을 계속합니다.")
}