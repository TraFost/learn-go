# Go Learning Journal

## 2025-09-27

### Why Move to Go

- Writing JS for both frontend and backend started to feel repetitive
- Mainly because i wanted to learn go from such a long time ago, but i haven't got the time yet until now xdd

### Concept Learned

#### Go Packages and Binaries

- When building a Go package, the output is an executable binary (`.exe` on Windows)
- This feels different from Node.js, which typically produces a `/dist` folder with JS bundles
- The compiled binary can run directly without needing Node.js installed (can you believe it?)

#### Variable Types: Explicit and Inferred

- The holy grail Walrus Operator
- We only can declare an inferred variable(Walrus Operator) inside a function
- If we want to declare variable outside of a function, we use either var or const

## 2025-09-28

#### Declaring Variable

- we can declare multiple variable by using parantheses, example below:

  ```
  const (
      a = 1
      b = 2
      c = 3
  )
  ```

- There's no string template literals in go, but there's a built-in function inside "fmt" package that we can use to do the job.

  ```
    name := "Big"
    age := 23

    msg := fmt.Sprintf("Hello %s, you are %d old", name, age)

    <!-- each types holds their own placeholder values, keep it in mind -->
    %s → string
    %d → integer
    %f → float
    %v → generic (Go will figure it out)
  ```

- Formal Way: iota is Go’s auto-incrementer for constants, often used to implement enums and bitmasks.
- Go doesn't have enums, but they got iota. iota is an auto incrementing number that will increments by 1 on each new line.
  - Started at 0 in every each constants variable block.
  - Build Enums or bit flags without hardcoding numbers.
  - Empty lines repeat the previous expression with the new iota
  - Each left shift doubles the number (2ⁿ), so the sequence is 1, 2, 4, 8, 16… until you run out of bit width (e.g., 8 bits maxes at 128 = 1000 0000)

```
    << shifts bits left → multiplies by 2 each step.
    >> shifts bits right → divides by 2 each step.
    Formula: X << n = X * (2^n) and X >> n = X / (2^n) (integer division).

	const (
		test = 1 << iota // 0001
		execute // 0010
		final // 0100
		anotherVal // 1000
	)
```

---
