```shell
# 运行例子
$ go run example.go
# 压测
$ wrk -c 500 -t 30 -d 1m  http://127.0.0.1:9998/buzoumei@gmail.coom
# 进入 tool 查看 pprof
$ go tool pprof http://127.0.0.1:9998/debug/pprof/profile

$ 进入 pprof 
# top 查看资源占用top
(pprof) top
Showing nodes accounting for 115.70s, 91.63% of 126.27s total
Dropped 415 nodes (cum <= 0.63s)
Showing top 10 nodes out of 91
      flat  flat%   sum%        cum   cum%
   106.41s 84.27% 84.27%    106.48s 84.33%  syscall.syscall
     2.07s  1.64% 85.91%      2.10s  1.66%  runtime.usleep
     2.05s  1.62% 87.53%      2.05s  1.62%  runtime.pthread_kill
     1.45s  1.15% 88.68%      1.45s  1.15%  runtime.kevent
     0.84s  0.67% 89.35%      0.88s   0.7%  runtime.nanotime1
     0.80s  0.63% 89.98%      1.47s  1.16%  runtime.scanobject
     0.64s  0.51% 90.49%      0.64s  0.51%  runtime.pthread_cond_wait
     0.60s  0.48% 90.96%      0.72s  0.57%  runtime.step
     0.45s  0.36% 91.32%      1.19s  0.94%  runtime.pcvalue
     0.39s  0.31% 91.63%      2.92s  2.31%  runtime.gentraceback

# list handler 查看代码运行耗时
(pprof) list handler
Total: 2.10mins
ROUTINE ======================== main.handler in /Users/fotoable/GolandProjects/github.com/chen-shiwei/learn-notes/go-examples/pprof/example.go
      30ms      3.43s (flat, cum)  2.72% of Total
         .          .      5:   "net/http"
         .          .      6:   _ "net/http/pprof"
         .          .      7:   "regexp"
         .          .      8:)
         .          .      9:
      10ms       10ms     10:func handler(wr http.ResponseWriter, r *http.Request) {
         .      3.11s     11:   var pattern = regexp.MustCompile(`^(\w+)@gmail.com$`)
      10ms       10ms     12:   account := r.URL.Path[1:]
      10ms      230ms     13:   res := pattern.FindSubmatch([]byte(account))
         .          .     14:   if len(res) > 1 {
         .       70ms     15:           wr.Write(res[1])
         .          .     16:   } else {
         .          .     17:           wr.Write([]byte("None"))
         .          .     18:   }
         .          .     19:}
         .          .     20:
ROUTINE ======================== net/http.(*ServeMux).handler in /Users/fotoable/go/go1.15.6/src/net/http/server.go
         0       10ms (flat, cum) 0.0079% of Total
         .          .   2393:   // Host-specific pattern takes precedence over generic ones
         .          .   2394:   if mux.hosts {
         .          .   2395:           h, pattern = mux.match(host + path)
         .          .   2396:   }
         .          .   2397:   if h == nil {
         .       10ms   2398:           h, pattern = mux.match(path)
         .          .   2399:   }
         .          .   2400:   if h == nil {
         .          .   2401:           h, pattern = NotFoundHandler(), ""
         .          .   2402:   }
         .          .   2403:   return
# 更为直观的火焰图
$ go-torch http://127.0.0.1:9998/debug/pprof/profile
# 浏览器打开 
open 打开torch.svg
```
