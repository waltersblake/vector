# vector

Using Go generics to provide a slice wrapper with some helper methods similar to a C++ Vector

<br>

## Usage

**Creating a 0 value variable**

```go
var numbers vector.Vector[int]
```

**Using the New function, passing in length and capacity**

```go
words := vector.New[string](0,10)
```

**Add one or many values. Capacity will auto-grow when needed just as a normal slice does, and no need to re-assign it to itself**

```go
words.Push("Hello", "there")
```

**Access the underlying slice to do For Range and other standard Go slice operations**

```go
for _, v := range words.Slice {
    fmt.Println(v) // "Hello" on first loop, "there" on second
}
```

**Removing and returning the most recently added value**

```go
word := words.Pop() // word == "there"
```

**Getting length and capacity**

```go
capacity := words.Capacity() // 10
length := words.Size() // 1
```

**Lowering capacity to match current length**

```go
words.Shrink()
capacity := words.Capacity() // 1
```
