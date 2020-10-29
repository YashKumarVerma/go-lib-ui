# go-lib-ui
A simple handy cli ui library that supports printing messages with context along with color and emoji support. 

## Why a module
this has been published as a separate module because I realized that I was using similar functionality in multiple golang projects and hence duplicating a lot of code. Also as adding and managing new features was an issue. Publishing it as a separate module then using it across multiple projects would serve the purpose.

## Usage
Include the library 
```
# fetch the library
go get github.com/YashKumarVerma/go-lib-ui
```

## Under the hood
The package uses two very popular libraries 
- [`github.com/fatih/color`](github.com/fatih/color) 
- [`github.com/kyokomi/emoji`](github.com/kyokomi/emoji)