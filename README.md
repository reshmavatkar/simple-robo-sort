# Package simple robo sorter

This Go module implements a package classification system for Thoughtful's robotic automation factory. The system determines how to dispatch packages based on their size (volume and dimensions) and weight.

## Classification Rules

Packages are classified into one of the following stacks:

- **STANDARD**: Package is **neither bulky nor heavy**
- **SPECIAL**: Package is **bulky or heavy** (but not both)
- **REJECTED**: Package is **both bulky and heavy**

### Definitions

- A package is **bulky** if:
  - Volume ≥ 1,000,000 cm³ (width × height × length), **or**
  - Any dimension ≥ 150 cm
- A package is **heavy** if:
  - Mass ≥ 20 kg

---
## Requirements

Go 1.18+

## Function Signature

```
func sort(width, height, length, mass int) string
```

## Interactive CLI
###  Clone the repository

```
$ git clone https://github.com/reshmavatkar/simple-robo-sort
$ cd simple-robo-sort
```
### Run the program:
```
go run main.go
```

Sample interaction:
```
Package Sorter (type 'exit' at any prompt to quit)
Width (cm): 100
Height (cm): 100
Length (cm): 100
Mass (kg): 10
Dispatch to: SPECIAL stack
```

## Unit Testing
Run tests:
```
go test
```