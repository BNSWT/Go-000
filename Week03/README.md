> 基于errgroup实现一个http server的启动和关闭，以及linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出

## 代码逻辑说明

通过mux管理多个路由

```go
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
```

创建带context的errgroup，以实现一个协程推出，所有协程都会推出

```go
group, ctx := errgroup.WithContext(context.Background())
```

首先是启动server的协程

```go
// g1: 启动 server
// g1 退出后, context 将不再阻塞，g2, g3 都会随之退出
// 然后 main 函数中的 group.Wait() 退出，所有协程都会退出
group.Go(func() error {
    return server.ListenAndServe()
})
```

然后是关闭server的协程

```go
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
```

然后是linux signal信号的注册和处理，这里实现的是siginit信号与sigterm信号的注册和处理

```go
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
```

最后打印使得errgroup中所有协程退出的原因

```go
fmt.Printf("errgroup exiting: %+v\n", group.Wait())
```

## 测试

访问homepage

![image-20211227224125775](/home/yuyangz/.config/Typora/typora-user-images/image-20211227224125775.png)

访问shutdown page

![image-20211227224151956](/home/yuyangz/.config/Typora/typora-user-images/image-20211227224151956.png)

![image-20211227224216897](/home/yuyangz/.config/Typora/typora-user-images/image-20211227224216897.png)

interrupt的情况

![image-20211227224046168](/home/yuyangz/.config/Typora/typora-user-images/image-20211227224046168.png)

terminate的情况

![image-20211227225353328](/home/yuyangz/.config/Typora/typora-user-images/image-20211227225353328.png)

![image-20211227225425849](/home/yuyangz/.config/Typora/typora-user-images/image-20211227225425849.png)
