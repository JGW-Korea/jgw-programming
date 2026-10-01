package main

func main() {
	Loop1:
	for i := 0; i <= 10; i++ {
	
		Loop2:
		for j := 0; j <= 10; j++ {
			
			// 1. i가 2이면서 j가 4인 경우, Loop1의 다음 반복으로 넘어간다.
			if i == 2 && j == 4 {
				continue Loop1
			}
			
			// 2. j가 8 이상인 경우, Loop2 반복문을 중단한다.
			if j > 8 {
				break Loop2
			}
			
			// 3. i - 7이 7 이상인 경우, Loop1 반복문을 중단한다.
			if (i - 7) > 7 {
				break Loop1
			}
		}
	}
}