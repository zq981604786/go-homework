package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//gruop, _ := errgroup.WithContext(context.Background())
	//serverOut := make(chan struct{})
	//
	//// group1
	//gruop.Go(func() error {
	//	Server(9001, "1111", true, serverOut)
	//	return nil
	//})
	//
	//// group2
	//gruop.Go(func() error {
	//
	//	Server(9002, "2222", false, nil)
	//	<-serverOut
	//	return nil
	//})
	//
	//// group3
	//gruop.Go(func() error {
	//
	//	Server(9003, "3333", false, nil)
	//	return nil
	//})
	//
	//if err := gruop.Wait(); err != nil {
	//	fmt.Printf("main group error:%v", err)
	//}
	//
	//fmt.Printf("main exit...")

	g, ctx := errgroup.WithContext(context.Background())

	mux := http.NewServeMux()
	mux.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("pong"))
	})

	// 模拟单个服务错误退出
	serverOut := make(chan struct{})
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		serverOut <- struct{}{}
	})

	server := http.Server{
		Handler: mux,
		Addr:    ":9000",
	}

	g.Go(func() error {
		return server.ListenAndServe()
	})

	g.Go(func() error {
		select {
		case <-ctx.Done():
			log.Println("errgroup exit...")
		case <-serverOut:
			log.Println("server will out...")
		}

		timeoutCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		log.Println("shutting down server...")
		return server.Shutdown(timeoutCtx)
	})

	g.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			return fmt.Errorf("get os signal: %v", sig)
		}
	})

	fmt.Printf("errgroup exiting: %+v\n", g.Wait())
}

//func Server(addr int, content string, shutdown bool, ch chan<- struct{}) {
//	mux := http.NewServeMux()
//	mux.HandleFunc("/ping", func(writer http.ResponseWriter, request *http.Request) {
//		writer.Write([]byte(content))
//	})
//	if shutdown {
//		mux.HandleFunc("/shutdown", func(writer http.ResponseWriter, request *http.Request) {
//			ch <- struct{}{}
//			writer.Write([]byte("shutdown..."))
//		})
//	}
//	server := http.Server{
//		Addr:    fmt.Sprintf(":%v", addr),
//		Handler: mux,
//	}
//	server.ListenAndServe()
//}
