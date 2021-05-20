# goEMO
change function name and variables into emoji letters.

`Go` actually supports using some of emojis as ident. (ref: [Does Go (golang) actually permit unicode characters in variable names?](https://stackoverflow.com/questions/33947871/does-go-golang-actually-permit-unicode-characters-in-variable-names))

# Example
One possible scenario, when choosing `Math Script Bold` font:
```go
// Before goEMO
var SayHi = "hi"

func LetsGoCrazy() {
    println(SayHi)
}

// After goEMO
var 𝓢𝓪𝔂𝓗𝓲 = "hi"

func 𝓛𝓮𝓽𝓼𝓖𝓸𝓒𝓻𝓪𝔃𝔂() {
    println(𝓢𝓪𝔂𝓗𝓲)
}
```

# Feature
- Can replace all variable names and function names (Except for `main` function in `main` package)
- Support various fonts
- Support change files (ended with `.go`) or a single file (also, ended with `.go`)
- Most files should still work after been transformed. 
- Let's think...

# Why doing this
'Cause I'm an EMO boy!

# License
GPL v3.0