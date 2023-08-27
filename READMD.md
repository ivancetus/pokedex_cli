This is a REPL app built using Go.   
REPL stands for "read-eval-print loop".

type a command -> evaluate -> print the result

## Test
### to test main package
    go test
### test all package
    go test ./...

### to run test in GoLand
    https://www.jetbrains.com/help/go/performing-tests.html

    selected file or folder:
    ctrl + shift + r

    selected function:
    place gutter to run test function

### Pokemon API
    https://pokeapi.co/docs/v2

### JSON to GO struct
    https://pokeapi.co/api/v2/location/

    paste the content from the above page to the following link:
    https://mholt.github.io/json-to-go/

    before using the struct, review the type definition first,
    make sure it matches your need and use case

    use this struct as json response struct