package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
)

//start a server
func StartHttpServer(srv *http.Server) error {
	http.HandleFunc("/hello", HelloServer)
	fmt.Println("hello lisk ,you test server run")
	err := srv.ListenAndServe()
	return err
}

// add a hanlder for it
func HelloServer(w http.ResponseWriter, req *http.Request) {
	_, err := io.WriteString(w, "hello lisk!!!!!\n")
	if err != nil {
		log.Printf(err.Error())
		return
	}
}

func main() {
	ctx := context.Background()
	// 定义 withCancel -> cancel() 方法 去取消下游的 Context
	ctx, cancel := context.WithCancel(ctx)
	fmt.Println(ctx)
	// 使用 errgroup 进行 goroutine 取消
	group, errCtx := errgroup.WithContext(ctx)
	//http server
	srv := &http.Server{Addr: ":5000"}

	group.Go(func() error {
		return StartHttpServer(srv)
	})

	group.Go(func() error {
		<-errCtx.Done() //阻塞。因为 cancel、timeout、deadline 都可能导致 Done 被 close  关了会有个空消息发送也就不阻塞了
		fmt.Println("http server stop")
		return srv.Shutdown(errCtx) // 关闭 http server
	})

	chanel := make(chan os.Signal, 1) //这里要用 buffer 为1的 chan
	signal.Notify(chanel)

	group.Go(func() error {
		for {
			select {
			case <-errCtx.Done(): // 因为 cancel、timeout、deadline 都可能导致 Done 被 close
				return errCtx.Err()
			case <-chanel: // 因为 kill -9 或其他而终止
				cancel()
			}
		}
		return nil
	})

	if err := group.Wait(); err != nil {
		fmt.Println("group error: ", err)
	}
	fmt.Println("all group done!")

}
