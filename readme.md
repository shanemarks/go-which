# Go Which

## About

I built this project as a way to familiarise myself with the Go programming language. My approach to learning new languages involves rebuilding a simple, real application. In this case I chose the "which" command.

Which is a tool that lets you find where an executable may exist in your path directory. It is a nice test project because it requires working with the filesystem, which means you need to rely on composition to make the logic testable, instead of coupling it with the logic.

## My Learnings

Through the hour that I spent on this project, I found that Go and its tooling just simply worked with VSCode (autocompletion, linting, debugging etc) - no "farming" required. I also think it's great that the testing libraries and command-line argument libraries come standard with a vanilla setup.

I was surprised there was no map function like you find in JavaScript, or support for enums, but I quickly found solutions for that with the help of google.

## My Approach

Most of the logic exists int the main.go file, which 2 small helpers in the the helper.go file. 

### Helpers
- **Map**:  This implements a basic generic mapping function which we use to append the filename to a list of paths. Perhaps this was overkill, because we only really used this once

- **FileExists (and its delegate)**: I wanted to be able to mock out the file existance check, so I could check against a mock file system for stable testing, and so I decomposed this into a small function and a type definition that would allow me to swap out this piece of logic in the main code base.

### Flow Control

#### Command Line Argument Parsing
In our main function I relied on the built in `Flag` functionality to generate and parse commandline arguments.

```go
allPointer := flag.Bool("a", false, "Returns all matches")
flag.Parse()

//any arguments that aren't flags come here
arguments := flag.Args()
target := ""
if len(arguments) > 0 {
    target = arguments[0]
}
```

#### Running the Business Logic
Perhaps also overkill, I  built a map of "Actions" and calulated which action to execute, instead of relying on a chain of if/else in the main body

```go
runMode := calculateRunMode(len(os.Args) >= 2, *allPointer)
actions := buildActions()

action, exists := actions[runMode]

if exists {
    action(target)
}

```


### Testing
For testing, I used the `Testify` package so that I could use asserts, which I am used to.  I also used a delegate that let me switch out the check for file existance in my unit tests.
```go
```

