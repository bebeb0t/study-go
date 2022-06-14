# study-go
個人的にGo言語を学習するためのリポジトリ

---

## 学習サイト
- [A Tour Of Go](https://go-tour-jp.appspot.com/welcome/1)

---

## 目次
1. 基礎
   1. パッケージ、変数と関数
   1. for文、if文などの操作
   1. 構造体、スライス、マップなどの様々な型
1. メソッドとインターフェース
1. 並列


---
### パッケージ、変数と関数

#### パッケージ
Goプログラムはmainパッケージから開始
下記の例では、`"fmt"` と `"math/rand"` パッケージをインポートする。

```go
import (
    "fmt"
    "math/rand"
)

func main() {
    fmt.println("My favorite number is", rand.Intn(10))
    // -> 0 <= n <= 10 の数値を返す
}
```

複数のインポートの式を記述する場合
```go
import "fmt"
import "math/rand"
```

#### 関数
関数は0個以上の引数をとることができる。
変数名の __後ろ__ に型名を書くことに注意。[型がなぜこのように記述するのか](https://teratail.com/questions/188903)
```go
func add(x int, y int) int {
    return x + y
}
```

関数の２つ以上の引数が同じ型である場合、最後の型を残して省略して記述。
```go
func add(x, y int) int {
    return x + y
}
```

複数の戻り値も可能。
```go
func swap(x, y string) (string, string) {
    return y, x
}
```

返り値に名前をつけて、返却することも可能。
return式に何も書かずに戻すことを `"naked" return` と呼びます。
短い関数でのみ利用すること。可読性に悪影響あり。
```go
func split(sum int) (x, y int) {
    x = sum * 4 / 9
    y = sum - x
    return
}
```

#### 変数
varで変数を宣言。
varはパッケージ、関数で利用可能。

```go
var c, python, java bool

func main() {
    var i int
    fmt.Println(i, c, python, java)
    // 0 false false false
}
```

Go言語の組み込み関数について
```go
bool
string

int int8 int16 int32 int64
uint uint8 uint16 uint32 uint64 uintptr
// サイズ、符号なし整数の型を使うための特別な理由がない場合は基本的に`int`を使うようにする。


byte // uint8の別名
rune // int32の別名
     // Unicodeのコードポイントを表す

float32 float64
complex64 complex128 // 複素数
```

定数の宣言について
定数は、文字（character）、文字列（string）、boolean、数値（numeric）のみで使えます。なお、定数は `:=` を使って宣言できない。

```go
const Pi = 3.14

func main() {
    const World = "世界"
    fmt.Println("Hello", World)
    fmt.Println("Happy", Pi, "Day")
}
```
---
### for文、if文などの操作
#### For
forループはセミコロン ; で３つの部分に分かれている。
- 初期化ステートメント: 最初のイテレーション（繰り返し）の前に初期化が実行される(任意)
- 条件式: イテレーション毎に評価される
- 後処理ステートメント: イテレーション毎の最後に実行される(任意)
```go
func main() {
    sum := 0
    for i := 0; i < 10; i++ {
        sum += i
    }
    fmt.Println(sum)
}
```
#### If-else
```go
if x < 0 {
    fmt.Println("OK")
} else {
    fmt.Println("ELSE")
}
fmt.Println("NG")

if v := fun(x, n); x < 0 {
    fmt.Println("OK")
}
fmt.Println("NG")
```

#### Switch
各caseの最後に必要なbreakステートメントがGO言語では自動的に提供される。
switch caseは、上から下へcaseを評価する。
条件が一致すれば、そこで停止。

```go
    func main() {
        fmt.Print("Go runs on ")
        switch os := runtime.GOOS; os {
        case "darwin":
            fmt.Println("OS X.")
        case "linux":
            fmt.Println("Linux")
        default:
            fmt.Printf("%s\n", os)
        }
    }
```

#### Defer
defer ステートメントは、defer へ渡した関数の実行を、
呼び出し本の関数の終わり（returnする）まで遅延させるもの
deferへ渡した関数はLIFO(last-in-first-out)の順番で実行。
```go
func main() {
    defer fmt.Println("world")
    fmt.Println("hello")

    // -> hello
    // -> world
}
```

### 構造体、スライス、マップなどの様々な型
#### Pointer
Goはポインタを扱います。ポインタは値のメモリアドレスを指す。
変数Tのポインタは、*T型で、ゼロ値は`nil`

& オペレータは、そのオペランドへのポインタを引き出します。

```go
var p *int
i := 42
p = &i
```

#### Struct
構造体は、フィールドの集まり

```go
type Vertex struct {
    X int
    Y int
}

func main() {
    v := Vertex{1, 2}
    fmt.Println(v)
    // -> {1 2}
    v.X = 4
    fmt.Println(v)
    // -> {4 2}
}
```



引数の型を省略する記法ですが、メリットって何なんでしょうね
var constの使い分けについて
Goの定数（const）について
https://github.com/go-fsweb/practice-golang/tree/main/01_golang-ness#%E5%AE%9A%E6%95%B0%E3%81%AE%E5%AE%9A%E7%BE%A9%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6
switchのbreakを書かないと止まってくれない仕様のほうが気持ち悪いなっておもうので、Ｇｏの仕様のほうが自然なような気がする
石橋 太 から全員に 07:29 AM
golang/tree/main/01_golang-ness#%E5%AE%9A%E6%95%B0%E3%81%AE%E5%AE%9A%E7%BE%A9%E3%81%AB%E3%81%A4%E3%81%84%E3%81%A6
