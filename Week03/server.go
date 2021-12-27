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
	// 自定义一个ServeMux，管理多个路由
	mux := http.NewServeMux()

	// homepage
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("user visiting home page")
		w.Write([]byte("Welcome to home page!"))
	})

	// shutdown page
	serverOut := make(chan struct{})
	mux.HandleFunc("/shutdown", func(w http.ResponseWriter, r *http.Request) {
		log.Println("user shutting down")
		w.Write([]byte("Bye!"))
		serverOut <- struct{}{} //shutdown时给channel赋一个空的结构体即可
	})

	server := http.Server{
		Handler: mux,
		Addr:    ":8080",
	}

	group, ctx := errgroup.WithContext(context.Background())

	// g1: 启动 server
	// g1 退出后, context 将不再阻塞，g2, g3 都会随之退出
	// 然后 main 函数中的 group.Wait() 退出，所有协程都会退出
	group.Go(func() error {
		return server.ListenAndServe()
	})

	// g2: 处理 shutdown page的情况
	// g2 退出时，调用了 shutdown，g1 会退出
	// g2 退出后, context 将不再阻塞，g3 会随之退出
	// 然后 main 函数中的 group.Wait() 退出，所有协程都会退出
	group.Go(func() error {
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

	// g3: 捕获os 退出信号
	// g3 退出后, context 将不再阻塞，g2 会随之退出
	// g2 退出时，调用了 shutdown，g1 会退出
	// 然后 main 函数中的 group.Wait() 退出，所有协程都会退出
	group.Go(func() error {
		quit := make(chan os.Signal, 0)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

		select {
		case <-ctx.Done():
			return ctx.Err()
		case sig := <-quit:
			return fmt.Errorf("get os signal: %v", sig)
		}
	})

	fmt.Printf("errgroup exiting: %+v\n", group.Wait())
}
