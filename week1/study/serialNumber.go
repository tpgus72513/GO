package main // 이 파일이 독립적으로 실행 가능한 '메인' 프로그램임을 선언합니다.

import (
	"fmt"     // 입출력(Scan, Println)을 담당하는 패키지입니다.
	"sort"    // 슬라이스 정렬(Slice) 기능을 제공하는 패키지입니다.
	"unicode" // 문자가 숫자인지 판별(IsDigit)하는 기능을 위해 가져옵니다.
)

func getDigitSum(s string) int { // string을 입력받아 int를 반환하는 함수입니다.
	sum := 0 // 숫자의 합을 저장할 변수를 0으로 초기화합니다.
	for _, r := range s { // 문자열 s를 순회합니다. _는 인덱스 번호를 안 쓰겠다는 뜻이고, r은 각 문자입니다.
		if unicode.IsDigit(r) { // 만약 문자 r이 '0'~'9' 사이의 숫자라면 실행합니다.
			// r - '0'은 아스키 코드 값을 이용해 문자 '5'를 실제 정수 5로 바꾸는 테크닉입니다.
			sum += int(r - '0') 
		}
	}
	return sum // 최종 계산된 합계를 반환합니다.
}

func main() {
	var n int
	fmt.Scan(&n) // 첫 번째 줄에서 시리얼 번호의 개수 N을 입력받습니다.

	// n개의 문자열을 담을 수 있는 '슬라이스(동적 배열)'를 만듭니다.
	serials := make([]string, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&serials[i]) // N번 반복하며 각 시리얼 번호를 슬라이스에 저장합니다.
	}
	// sort.Slice(대상, 비교함수): 비교함수가 true를 반환하면 i번째가 j번째보다 앞에 옵니다.
	sort.Slice(serials, func(i, j int) bool {
		// [조건 1] 문자열의 길이가 다르면, 더 짧은 것이 무조건 앞으로 옵니다.
		if len(serials[i]) != len(serials[j]) {
			return len(serials[i]) < len(serials[j])
		}

		// [조건 2] 길이가 같다면, 위에서 만든 함수를 써서 숫자의 합을 비교합니다.
		sumI := getDigitSum(serials[i])
		sumJ := getDigitSum(serials[j])
		if sumI != sumJ {
			return sumI < sumJ // 합이 더 작은 것이 앞으로 옵니다.
		}

		// [조건 3] 1, 2번으로도 순위가 안 정해지면, 기본 사전순(A < B)으로 비교합니다.
		return serials[i] < serials[j]
	})
	// 정렬이 완료된 슬라이스를 순회하며 하나씩 출력합니다.
	for _, s := range serials {
		fmt.Println(s)
	}
}