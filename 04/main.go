package main

import "fmt"

//var Wait sync.WaitGroup
//var Counter int = 0

func main() {
	//for routine := 2; routine <= 10; routine++ {
	//	Wait.Add(1)
	//	go Routine(routine)
	//}
	//Wait.Wait()
	//fmt.Println("1111", Counter)
	var i int
	i++
	fmt.Println(i)
}

//func Routine(id int) {
//	for count := 0; count < 2; count++ {
//		value := Counter
//		value++
//		Counter = value
//	}
//	Wait.Done()
//}

//type Tracker struct {
//	ch   chan string
//	stop chan struct{}
//}
//
//func NewTracker() *Tracker {
//	return &Tracker{
//		ch: make(chan string, 10),
//	}
//}
//
//func (t *Tracker) Event(ctx context.Context, data string) error {
//	select {
//	case t.ch <- data:
//		return nil
//	case <-ctx.Done():
//		return ctx.Err()
//	}
//}
//
//func (t *Tracker) Run() {
//	for data := range t.ch {
//		time.Sleep(1 * time.Second)
//		fmt.Println(data)
//	}
//	t.stop <- struct{}{}
//}
//
//func (t *Tracker) Shutdown(ctx context.Context) {
//	close(t.ch)
//	select {
//	case <-t.stop:
//		fmt.Println("结束")
//	case <-ctx.Done():
//		fmt.Println("终止")
//	}
//}
//
//func main() {
//	tr := NewTracker()
//	go tr.Run()
//
//	_ = tr.Event(context.Background(), "test1")
//	_ = tr.Event(context.Background(), "test2")
//	_ = tr.Event(context.Background(), "test3")
//	time.Sleep(3 * time.Second)
//	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(5*time.Second))
//	defer cancel()
//	tr.Shutdown(ctx)
//}
