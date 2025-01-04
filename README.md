# Advent of Code

Mega repo to hold all of my solutions to the [Advent of Code](https://adventofcode.com/)

## Go
### Stars: 49 :star:
There is a helper script that will create the directory path for a new solution and will call the input fetch script to get the puzzle input for that day (if it doesn't already exist).

After creating the folder structure it will use the `template.go` and `template_test.go` files to create the boiler plate code to get started.

```bash
./new.sh <YEAR> <DAY>
```

Solutions are structured to be able to run in the specific day's directory and uses the relative path to the shared puzzle input. It also accepts a '-i' argument to specify different puzzle input. If 'stdin' is passed then it will read the puzzle input from there.

Example 1:
```bash
cd <YEAR>/<DAY> && go run main.go
```

Example 2:
```bash
go run <YEAR>/<DAY>/main.go -i <INPUT_PATH>
```

Each solution also comes with a test file that is mostly designed to use the sample input to validate that my solution is on the right track but also allows me to write quick sanity check tests for any helper functions created.

Example 1:
```bash
go test ./...
```

Example 2:
```bash
go test <YEAR>/<DAY>/main_test.go
```

## Python
### Stars: 49 :star:
Under construction, old solutions just moved over

## C#
### Stars: 4 :star:
Under construction, old solutions just moved over
