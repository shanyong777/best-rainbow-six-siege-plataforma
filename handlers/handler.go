package main

import (
    "fmt"
    "sync"
    "time"
    "crypto/sha256"
)

var ( appVersion = "5.1" )

func f7D1tYgBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func I8yL8aBHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 57
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SZQaCxkcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bYlfFeSXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X94Fhe1FWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 199
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aNWwbUhRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cJ9wTBAkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ea6ngDMgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 42
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L28LddgsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bGR6rm4CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3F7NmN8fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6gJCOAm3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vN0rARwRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ti7e6nWnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HZgM3fo0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h5Gmkq9rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LquMa5vfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 207
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CzABZcx7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cxZefRg5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 103
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ngl1HNWGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6SZqCFvTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 72
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q2wt8YLiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qCXWj9dOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func stb7fl4VWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 238
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func R4TUaBdeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TmlvTKC0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func laIV5pjDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 226
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZSJZKBbMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zOvahfrEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rdXJzQk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4VzsLoFkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HBdd26hoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hcoFArueWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EiMJlykFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Is4nL23WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s4RKdyZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 75
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 376LY2DuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FT56qatoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9n56fQeEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GL79pT37Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EqaNkg7mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func juq9p29nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJyiIChVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 38
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Q2sC7fF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lLFsF636Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 194
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tpPov2iiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PK2oloucWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iBURlzKiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func knmLbuyBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GN8nYVk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 133
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hwQkB9hAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NruTTO7IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 167
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tnBjDAVEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 132
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QAkndSqQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JDXseH18Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LR4823nwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ukKZijELWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WFmdYXunWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ClBRYksPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 200
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v5tzlgiqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 260
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qmZQGuF6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oJNvR2g2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TZliK2cXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 2 + 204
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func k3QqQS6JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DHueNBfjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dfnSjmlvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tmx3ZSXaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xfhu60llWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 46
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W5XweirOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 117
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NzdTDL7aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 85
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eA2yO3fUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LpE1I8EhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 81
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1nDUwZMwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mXY0oPopWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WNbtee13Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1XJp0RAJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 216
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mHmmdvo0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 299
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TqUkRQQbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6lQ5sFJtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tLMOjNG5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XFRmQxm4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 250
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d4pRvGDOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CAHGKNp7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 519UBuHWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oKF9yF0WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 159
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4KIpBcEMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rOXeq4DjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aIW7FAYLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 137
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjoM3egeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 188
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0t3ihGHNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func i1flLOCkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 219
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sPBhIMSlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 125
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func b7EtymSVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0BOlOp4dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ezQ6uKgWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kIbNt3LGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N0bo9qSRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vUtXyYngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 30
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WrJeCys5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GRU6t6X6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tc1C7y9LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YLfZivgJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 16 + 193
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qpHugf4rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vIhGS6Z8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CybmqJw5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 209
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YgMdu9p1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZDluSo4xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TkK23OhuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 76
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lDglLvUHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4GQ0v6zOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CbSZCsRLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func z4mwdFlyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 215
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sthSv96tWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6uMyGk6KWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H1wqsfq2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MLci3bCoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 80
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func e65gxEcUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TX3bwy2oWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kZbHeMAqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func p26Tou2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 239
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LFCmxf0zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZZuYgmpJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UvIfAPbeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 245
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V0BkUzFBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AxWZMB6yWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xXHppgFUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JcJwxzO2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1Wyj4P1TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0k5nKtbgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xRWBzj6bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DAJeLoHXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func q4kaZNQLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m1geXjsJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kaFf3DOEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IEdrvj1UWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LzctOVQRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2fP6IdQhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func st1VXvNVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PpA9Tb3fWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V6nMPUbqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 109
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W8Papm0aWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UuOb7kmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 155
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ImSXAo1qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VMHZW4AhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IAJCFEeIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wdu5l9RWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IyULyA4JWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GoFgH9ExWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mTNVUF8dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9faSHQqZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 16Ho1F2sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PBdIwoX2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mdcLa9OZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YdIp32SlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kss59sYDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 168
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6HKz0AvVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SVfgzMkeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0cEhz0hxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Coh62jhwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jRpOJBlsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FvfpgckgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qFqCFFhKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 147
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3rKgYOPlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SVMakiAWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0FLb9O7TWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ai4PTZw1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 287
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rniGRwZDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MA7idsvEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dR7XSaw6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tgYWv8NqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func loHa8UQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MMErMg1EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CgordWE6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1opTNnizWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iYiT1LZyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lKXvK21dWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VqxPynRgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 247
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YJhXjxYpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LrROFzzNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hgeN5Xd3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 55
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mcNpx2awWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LMhwQozaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nuDB201MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 297
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 24r0r0gFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 142
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6v6GuRoZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nUvhcnYfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 16
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func onto5XTrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 86
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UmVQF0SmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func geVGcldHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4vsJgTLeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qLrAkvNWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 208
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AJw27J1zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4P10ekiSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JZywGFSLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 153
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dg7OmAmqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3qzGzaQwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ykNIkgjiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0jCRcwdMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 11 + 143
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yQz8wxGkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9wza3fUMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func c0lzALk8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eVPnHoR9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 185
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ccQ1XAoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 12
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tLUS7hTEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VQZ9IkMUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mO8G7qNCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fnFyeKgOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 107
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 67yLa98qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YU6iVdhLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LcoaY0PhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 257
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MKVe8ulIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oI8lE3gTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func f3jAwk8sWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LBDi1w4iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZNALpr4DWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ZkkmsnKGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8hx8VO6qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 180
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JF6U2yiYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 56
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func heAe8J7EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pZ8NzThJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 7cw79CxUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 144
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OGewyxzYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 136
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cexb5JwrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CajaTCG9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func s81a295MWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pb4WX684Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 288
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1CyzgUQQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 110
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lKM2rESzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fBw48hT9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FM7DfCdwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A9rD2LckWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vpffK2mPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YNsuQSoGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uh19z2nKWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sAfbiPAcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 278
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QjWzQtDQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 281
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ulmCHTUjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fLKmuBi2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 295
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LKzotwkBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 77
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 78rRslQgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 169
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RnaPxdyUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AtSkvv2LWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func huwvJ8PmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cied6DeNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tuCRS1HHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U8OvWvjxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QT5sYV4cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cSNFPoLwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ufPWp66EWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AuhVW0sFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 89
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h5nTi9ruWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 20
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IcRrQsisWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y8dubawRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BeDBwGcHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MIDuQW2SWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 138
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func x9rQ5FtRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 190
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UpICgwWdWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 84
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FsQJNk70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 189
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VAyujUdtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 161
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CEMnaXsNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PjWB6iOvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Cb7NQd6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func chVxtFsXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 60
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rtwZZscPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vZVIJsPFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 134
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mNAJ2h4IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T8VrynnlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bnSpW6UsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 59
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EcNIWEAaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 58 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JVEDCnLqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tuRvQnAYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func XHsxpfp4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TWGvJvVBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 273
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4w72KJdiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cnf0mNQcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NgJAKbhbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Po3XPqhGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 171
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8vJbbdUoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fL9YaOiQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 210
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GG6MlX3XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 74
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zuaoZGh5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SpKdGLdIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9lxguUnVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func AaFmfNQeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 279
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func am7UPgxSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oglO601XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rCnQ6YY0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TN6TZg9xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 23
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CjqPJBSRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S3AoS9uZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 36
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uUF9GfOzWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 258
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 20cQqW34Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NmnPBzhRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 195
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func INXdFzoNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func tepBPDygWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 158
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gYCeAuwrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 116
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 1rPY7HGyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X5RQ9Cz5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 212
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4ihXtEe3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QyVyTOqPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 284
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0oL6zOtCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func NjLXpbETWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 164
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func trMHKWRbWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 149
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wwSkfpILWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 267
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yx5XoZw6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 277
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4pMONkA3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 4 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0GB85fkYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 269
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lePrpVngWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eBW2Cia8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cIWC2V0cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hjES9l0uWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 268
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y8I8L5uZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6fNekVq0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4Aj0XAYlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 300
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E5o3prS6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lchWxz2HWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 13 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RtpKL5ZhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 62 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LU9qhOYxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 48
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func dRcPvjg8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 70
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GFn7TgWDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func A7cFNmaMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 105
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ue4h9IeDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IgQGknibWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 29
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0LumpqksWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 203
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iOnDHWAiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 234
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func HnCNfpVmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iSDfNshtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 85 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BBQ7AqZUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gg1S58iqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 121
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SabCpN68Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func elwlvF1IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func echCYYk9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0EPEwfC5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BAzNIOa0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 191
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JTwBdMZOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func crv5QAGvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PdhPIgkxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 296
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kGch88ltWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func aAFHI807Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5JpL515OWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uq1UiMZZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 179
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JLhTmWkeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MSiWrkoAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 266
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S2yHTbn6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OUSXwmGWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ytfunXReWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 253
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jjF38EkfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oleld9V7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func CB4Qg0J0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ue2g4hKaWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 294
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Up1nMqZUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uGIVaDEgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EG80OWMeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S1I2XhgyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xtyoFo2xWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 43 + 233
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wRAfVKiIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func td314gV9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func eEVAhpSrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 90XiWWuhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 17
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VPV10DhoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0wjhnh5bWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 129
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func MqDOUT2mWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 100
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5YNwo62zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 275
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EC8et5yWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 82 + 265
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lJSTTOzeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 264
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func U3VfFyxuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func y2hi5nOAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cod3rG9IWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2ucG23HCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 150
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BQzVY2KTWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 231
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func pQa5ss5rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 35
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9VnhW7GBWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 44
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func QD5gHzO5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 64 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func rIdLELDDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 286
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ktgskGtvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 293
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kt1GA1y2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 13
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JHWzkYW1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 77 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bnz7Yp7WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3VsBLkZ6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func imVNWJkPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OrLMBFDoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 32 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h5hF7fM8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 82
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0AD6bLFcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6BMRezxGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m9Amm8lxWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 66
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Rl7B1fpCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 123
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cmEdIhU3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WZN2Jy8rWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 67
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 4PccTJx1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 65 + 146
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VUvZ0L2wWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 244
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vyyaMcwQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6goHqeCrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0YYMMFoPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 213
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wGmdY7pGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 289
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func WvAFB6m8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 5 + 282
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zwlWQVAcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 261
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RZJZlIvYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Yfmi3LfhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 186
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xw1ZIQT3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 102
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wbTorbUgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 154
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func v7gjSBgWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 182
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sJEeFsXqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 211
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func S9cL70QvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 135
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Urx28oXGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kTadN7oLWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LayXTkKCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 99
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cGhUw4cMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 63
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3yrYAO1PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 272
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8810CpyPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func T8TqZNTkWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 98
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IGRdPI8XWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 86 + 53
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d7DQkba6Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 39 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sLXd9HjfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 84 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func d9FmvHD0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 130
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IZVuV6FcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 49 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iaNmiSqyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ybiz4WzoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 83 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0pn1HxdmWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 47
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 5sEb38KeWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 33
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oKkh4PlFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 46 + 225
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func M1b8hVTZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LjMTIHZgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 48 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jlKACizgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 270
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func JrdCQpqGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 41
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N8kxSBghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 83
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ceKXgrv0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func m4c0BivlWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lfnkPGv9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 112
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OjI7Xj9nWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 26 + 122
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 489rRLSIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 79
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EQrZQUEnWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 246
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nPi4MGtsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 6 + 91
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fFw1Vhy5Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 198
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OUQc1AQ1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func KA39T3WwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 220
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PfNvtWkCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 64
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vD2pnR7PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 29 + 243
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func wD3FAJswWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 119
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lj6am67CWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 206
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FOHAfnSwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 108
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L6lCQNXHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uzNlL1TgWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 43
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cvFqjd2cWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 10
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wh6zj2JHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 68
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DJUe0KviWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 60 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func yEUrRiHfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 59 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 48rh5MHjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 67 + 298
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kGeV8qB1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 69 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func K8fgWv00Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 241
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FhyvfUFNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 56 + 95
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ber8CofpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 33 + 254
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9CuGXUv3Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UYtP4Q4GWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 18
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ffhz4doOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 14 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func UGJNEmQSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 75 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vXvc5zJDWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 115
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func mEyzpOTXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 236
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func oBr2NEFCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 259
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func PSrOcNgPWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 69
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qJYcRwghWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 276
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fTM6y1WEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 74 + 113
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LGR1kci9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 21
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func N4gMFFc7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 173
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 40IuVx5iWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 23 + 172
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func szLikvbsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 202
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func H07C2WkCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 8 + 249
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func sdxKSXMsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 263
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func gKULYdQCWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 50 + 131
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func C5hYQceFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 18 + 292
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jpVy7FyvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 50
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func fFYTTmK9Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 87 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8tpvZ62lWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 78 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Dn5paGSEWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 36 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func lRoBbLPvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 104
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func qVbOWuDRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 38 + 163
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6YoLRC3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 175
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func YO9ZR5aNWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 222
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func FF7CLRdfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 9 + 148
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func zlUCAgWQWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 35 + 223
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func EKeDDxAiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 25 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func L8oIJozVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 90 + 78
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 06JjflU1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 12 + 87
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func APgzxy3hWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 166
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func puePiGPjWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3java9GhWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 47 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func kuRu1qInWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 66 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func W7uh7zobWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 162
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func V6EbIexvWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 72 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func LkQKG0efWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 24 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 70E3DNVZWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 10 + 157
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func iQQUL9FwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 25
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 8xlOTps1Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 52 + 118
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9Dyjf3DHWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 42 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func DxWHdM7PWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func GYbEYDGwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 156
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func isWTEj4qWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 79 + 178
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ztLac7RYWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 55 + 111
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func F2jGwZsUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 44 + 235
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Jrxo8X7zWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 232
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Qd9hhYI7Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 70 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func h6Keu3p0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 54 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 2T3lRDoFWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 205
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Uqv7SMKfWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 51 + 196
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func nUkrX0jAWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 96
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func jLBMegt4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 34 + 62
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func SnT4LMx2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 88 + 127
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func uyscQvuJWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 81 + 240
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func TC51ViavWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 63 + 90
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RbKuIQ9WWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 120
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xRTK2QHuWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 7 + 197
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func r3BbPuQUWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 34
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func olxaWfMtWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 37 + 192
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func RTv9vwblWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 229
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Gd25ozsOWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 237
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0d9UlAGIWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 19 + 160
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func X4gieKFsWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 31
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 9QQ39qYiWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 3 + 71
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Xw5Gj4O4Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 11
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func IkZfFp91Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 165
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 56X5URaRWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 45 + 290
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func hte7J9RoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 57 + 170
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ds932N70Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 61 + 124
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func E5T10hB8Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 28 + 101
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func omiDV8cSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 49
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Iq1iittGWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6KFcizseWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 20 + 93
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func cH8IRkAVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 76 + 285
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Idd2k9LoWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 183
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func BYE7Ty2NWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 30 + 65
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bTsCWclpWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 27 + 177
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Wi5b6BmwWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 53 + 224
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 6hcxiBM2Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 89 + 248
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ihULv7UXWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 15 + 181
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func vD0nCEzMWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 21 + 28
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func l2wtdzKyWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 17 + 218
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 3aIkSXstWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 145
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func VWtL27W0Worker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 22 + 283
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func 0EmojRcSWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 73 + 187
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Y9Uk4FbqWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 80 + 176
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func OC7M81DrWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 68 + 255
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func ws5LwVgVWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 41 + 221
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func bCHc33ApWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 31 + 184
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func Ae1wAVBcWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 71 + 27
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

func xNkbhDIWWorker(wg *sync.WaitGroup, id int, dataChan <-chan int) {
    defer wg.Done()
    for val := range dataChan {
        result := val * 40 + 106
        fmt.Printf("worker %d -> %d\n", id, result)
    }
}

