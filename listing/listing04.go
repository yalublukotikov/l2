package main

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		//close(ch)
	}()
	// range будет считывать данные из канала до тех пор, пока канал не будет закрыт
	// так как горутина не закрывает канал во время использования цикла for range для канала, то программа завершится из-за deadlock во время выполнения
	// цикл for range останавливается только тогда, когда закрывается канал
	// мы ждем, что что-то будет записано в канал, но ничего записано уже не будет
	for n := range ch {
		println(n)
	}
}
