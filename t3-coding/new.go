package main

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

type LogRecord struct {
	Timestamp time.Time
	ThreadID  int
	Elapsed   time.Duration
	IsFinal   bool
}

func main() {
	// 1. 读取X值
	var x int
	fmt.Print("请输入线程数量 X: ")
	_, err := fmt.Scanln(&x)
	if err != nil {
		fmt.Println("输入错误:", err)
		return
	}

	// 2. 读取日志文件名
	var logfile string
	fmt.Print("请输入日志文件名: ")
	_, err = fmt.Scanln(&logfile)
	if err != nil {
		fmt.Println("输入错误:", err)
		return
	}

	// 创建日志文件
	file, err := os.Create("F:\\tmp\\" + logfile + ".txt")
	if err != nil {
		fmt.Println("创建文件失败:", err)
		return
	}
	defer file.Close()

	// 共享资源保护
	var (
		mu          sync.Mutex
		logCh       = make(chan LogRecord, x*10)
		ctx, cancel = context.WithCancel(context.Background())
		wg          sync.WaitGroup
	)

	// 信号处理
	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT)

	// 日志写入协程
	go func() {
		for log := range logCh {
			mu.Lock()
			var status string
			if log.IsFinal {
				status = "[FINAL]"
			}
			fmt.Fprintf(file, "%s %s Thread-%d %.3fs\n",
				log.Timestamp.Format("2006-01-02 15:04:05.000"),
				status,
				log.ThreadID,
				log.Elapsed.Seconds())
			mu.Unlock()
		}
	}()

	// 3. 创建X个线程
	for i := 0; i < x; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()

			// 初始化随机数生成器
			r := rand.New(rand.NewSource(time.Now().UnixNano()))

			// 运行时间控制
			maxDuration := time.Duration(2*x+r.Intn(x)) * time.Second
			timeout := time.After(maxDuration)

			var (
				lastLog = time.Now()
				timer   = time.NewTimer(0)
			)
			defer timer.Stop()

			for {
				select {
				case <-timer.C:
					now := time.Now()
					elapsed := now.Sub(lastLog)

					// 发送日志记录
					logCh <- LogRecord{
						Timestamp: now,
						ThreadID:  id,
						Elapsed:   elapsed,
					}

					// 更新最后记录时间
					lastLog = now

					// 设置下一个随机间隔（平均2秒）
					interval := time.Duration(r.ExpFloat64() * 2 * float64(time.Second))
					timer.Reset(interval)

				case <-timeout:
					// 正常超时退出
					logCh <- LogRecord{
						Timestamp: time.Now(),
						ThreadID:  id,
						Elapsed:   time.Since(lastLog),
						IsFinal:   true,
					}
					return

				case <-ctx.Done():
					// 被中断退出
					logCh <- LogRecord{
						Timestamp: time.Now(),
						ThreadID:  id,
						Elapsed:   time.Since(lastLog),
						IsFinal:   true,
					}
					return
				}
			}
		}(i)
	}

	// 等待信号或完成
	select {
	case <-sigCh:
		fmt.Println("\n收到中断信号，停止所有线程...")
		cancel()
	case <-time.After(time.Duration(3*x) * time.Second):
		fmt.Println("所有线程正常完成")
	}

	// 等待所有线程退出
	wg.Wait()
	close(logCh)
}
