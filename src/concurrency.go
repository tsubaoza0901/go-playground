package src

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// サンプル①
func getLuckyNum() {
	fmt.Println("...")

	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

	num := rand.Intn(10)
	fmt.Printf("Today's your lucky number is %d!\n", num)
}

func Fortune1() {
	fmt.Println("What is today's lucky number?")
	go getLuckyNum()

	// Sleepを設定し、getLuckyNumの処理が完了するまで待ってmain関数を終了。
	// ただし、Sleepの引数に設定される時間は任意の値になるため、設定時間を短くすると
	// getLuckyNumが完了する前にmain関数が終了してしまう可能性も
	time.Sleep(time.Second * 5)
}

// サンプル②：WaitGroupを使用した場合
func Fortune2() {
	fmt.Println("What is today's lucky number?")

	var wg sync.WaitGroup
	wg.Add(1) // wgの内部カウンタの値を+1

	go func() {
		defer wg.Done() // wgの内部カウンタの値を-1
		getLuckyNum()
	}()

	// Waitはsync.WaitGroupが内部に持つカウンタが0になるまで処理の終了を待つため、
	// getLuckyNumは必ず実行される
	wg.Wait()
}

// サンプル③：チャンネルを使用した場合
func getLuckyNumChan(c chan<- int) {
	fmt.Println("...")

	rand.Seed(time.Now().Unix())
	time.Sleep(time.Duration(rand.Intn(3000)) * time.Millisecond)

	num := rand.Intn(10)

	c <- num
}

func Fortune3() {
	fmt.Println("What is today's lucky number?")

	c := make(chan int)

	go getLuckyNumChan(c)

	num := <-c

	// メインゴールーチンはチャンネルcから値を受信するまでブロックされるため
	// getLuckyNumの処理が完了する前に、main関数が終了することはない
	// そのため、sync.WaitGroupを用いた待ち合わせも不要（つまり、チャンネルには同期処理の性質がある）
	fmt.Printf("Today's your lucky number is %d!\n", num)

	close(c) // 使い終わったチャンネルのClose
}

// go文の引数の有無による挙動
func ConcurrencySample1() {
	var wg sync.WaitGroup
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func() {
			defer wg.Done()
			fmt.Println(i)
		}()
	}
	wg.Wait()
}

/*（実行結果）
3
3
3
*/
// ※なんで「3」？

func ConcurrencySample2() {
	var wg sync.WaitGroup
	wg.Add(3)

	for i := 0; i < 3; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}
	wg.Wait()
}

/*（実行結果）
2
0
1
*/

// 以下、Benchmark対象↓

func NoneConcurrency1() {
	src := []int{1, 2, 3, 4, 5}
	dst := []int{}

	for _, s := range src {
		result := s * 2
		dst = append(dst, result)
	}

	// fmt.Println(dst)
}

func NoneConcurrency2() {
	src := []int{1, 2, 3, 4, 5}
	dst := make([]int, len(src))

	for i, s := range src {
		result := s * 2
		dst[i] = result
	}

	// fmt.Println(dst)
}

func Concurrency1() {
	src := []int{1, 2, 3, 4, 5}
	dst := []int{}

	c := make(chan int)
	for _, s := range src {
		go func(s int, c chan int) {
			result := s * 2
			c <- result
		}(s, c)
	}

	for _ = range src {
		num := <-c
		dst = append(dst, num)
	}

	// fmt.Println(dst)
	close(c)
}

func Concurrency2() {
	src := []int{1, 2, 3, 4, 5}
	dst := []int{}

	var mu sync.Mutex
	var wg sync.WaitGroup
	wg.Add(len(src))

	for _, s := range src {
		go func(s int) {
			defer wg.Done()
			result := s * 2
			mu.Lock()
			dst = append(dst, result)
			mu.Unlock()
		}(s)

	}
	wg.Wait()
	// fmt.Println(dst)
}

func Concurrency3() {
	src := []int{1, 2, 3, 4, 5}
	dst := []int{}

	c := make(chan int)
	for _, s := range src {
		go func(s int, c chan int) {
			result := s * 2
			c <- result
		}(s, c)

		num := <-c
		dst = append(dst, num)
	}

	// fmt.Println(dst)
	close(c)
}
