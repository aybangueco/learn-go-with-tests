# learn-go-with-tests

This repository serves as a way to document my journey on diving deep in Go. I will be learning Go along with Test Driven Development approach using this wonderful [course](https://quii.gitbook.io/learn-go-with-tests) by quii. This takes a lot of reading but that's actually good for me, so I could get comfortable reading long ass documentations.

In this repository, I will be doing guided tasks based on the course, while wrapping my head around it and building mental models needed. I will also be documenting my learnings and realizations. I'm excited on this journey as this will be my secondary language along with TypeScript.

## Motivation

I love Golang ever since I wrote and refactored an existing codebase during my internship. It's so simple and ✨elegant✨ to write. I just want to be good at things I do and I want to be better than I was yesterday because I feel like I'm already hitting my comfort zone again, and I need to step away from it.

## Hello, World: Learnings

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world

Here I learned the existence of go doc and [pkgsite](https://github.com/golang/pkgsite). Such a great way to see code documentation even from your own projects. I learned the concepts of test driven development alone in this chapter. Well, its simple, just write a failing test, write the minimum code to make the test pass and clean up the code while keeping the tests running error free.

## Integers: Learnings

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/integers

Here I learned and got familiar with TDD workflow and also exposed me to write better code documentation via writing comments and testable examples.

## Iteration: Learnings

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/iteration

Here I learned about benchmarking. Surprise! Yeah its not just about implementing test driven development principles, it has benchmark lessons too.

To run benchmark

```bash
go test -bench=. ./iteration

go test -bench=. -benchmem ./iteration # additional flags for reporting memory allocations
```

Example of typical setup for benchmarking

```go
func Benchmark(b *testing.B) {
	//... setup ...
	for b.Loop() {
		//... code to measure ...
	}
	//... cleanup ...
}
```

## Arrays and Slices: Learnings

https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices

Here in this chapter, just another section being exposed on the workflow of TDD. I'm getting confident in writing tests, and realized how important writing tests are. I also learned the existence of test coverage which is an additional options when running tests in Go.
