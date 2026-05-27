## Профилирование по CPU
```GOGC=off go test -bench="." -cpuprofile="cpu.prof"```

Windows (separate command)
```
set GOGC=off
go test -bench="." -cpuprofile="cpu.prof"
```
Открыть собранный профиль в интерактивной утилите:
```
go tool pprof cpu.prof
```

(pprof) Дальше команды утилиты ppof
```
(pprof) top10
(pprof) top10 -cum
(pprof) peek words.UniqWords  // range function children
(pprof) web // failed on Win
```

Открыть в браузере:
`go tool pprof -http=localhost:8080 cpu.prof`


(pprof) peek JoinWords results:
```
...
                                             4.18s   100% |   github.com/.../profiler/joiner.BenchmarkJoinWords.func1
         0     0%  0.18%      4.18s 73.59%                | github.com/.../profiler/joiner.JoinWords
                                             2.45s 58.61% |   github.com/.../profiler/joiner.sorted (inline)
                                             0.76s 18.18% |   github.com/.../profiler/joiner.uniq (inline)
                                             0.55s 13.16% |   github.com/.../profiler/joiner.split (inline)
                                             0.23s  5.50% |   github.com/.../profiler/joiner.join (inline)
                                             0.19s  4.55% |   github.com/.../profiler/joiner.lower (inline)
```

## Профилирование по потреблению памяти:
```
GOGC=off go test -bench="." -memprofile="mem.prof"
```
Дальше те же комманды. Для запуска pprof:

```go tool pprof mem.prof```

И те же без доп флагов для функций 

result:
```
(pprof) peek joiner.JoinWords
Showing nodes accounting for 1220.17MB, 100% of 1220.17MB total
----------------------------------------------------------+-------------
      flat  flat%   sum%        cum   cum%   calls calls% + context              
----------------------------------------------------------+-------------
                                         1215.05MB   100% |   github.com/.../profiler/joiner.BenchmarkJoinWords.func1
         0     0%     0%  1215.05MB 99.58%                | github.com/.../profiler/joiner.JoinWords
                                          708.83MB 58.34% |   github.com/.../profiler/joiner.uniq (inline)
                                          257.15MB 21.16% |   github.com/.../profiler/joiner.split (inline)
                                          249.07MB 20.50% |   github.com/.../profiler/joiner.join (inline)
----------------------------------------------------------+-------------
```
