# コード修正対応の回答

## swift

誤 : `for i in 1..100`  
正 : `for i in 1...100`

## Rust

誤 : `main() {`  
正 : `fn main() {`

## Ruby

ここは同じ内容で2箇所

誤 : `elif i % 3 == 0`  
正 : `elsif i % 3 == 0`

誤 : `elif i % 5 == 0`  
正 : `elsif i % 5 == 0`

## Python

最後の`fizzbuzz()`で不要なインデントが入り正しくコードが理解されていない

## PHP

if文を見ると`FizzBuzz`と表示されるべき

`15で割った際に余りが0となる`

記載が抜けている。追加が必要

## kotlin

誤 : `func main() {`  
正 : `fun main() {`

## JavaScript

function,for,ifの3つのコードブロックがありますが、閉じるための中括弧が一つ足りない

## java

誤 : `System.out.println(i)`  
正 : `System.out.println(i);`

## go

誤 : `for (i := 1; i <= 100; i++) {`  
正 : `for i := 1; i <= 100; i++ {`

## C#

FizzBuzzと表示されるべき15で割った際の余りが0であることの判定が最後になってしまっている。
