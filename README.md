# lighter
a lightweight cli and launcher

# usage
The pattern is to create your flags, run Parse, then use your flags. 

The supported flag types are string, bool, and int64. To create a flag use NewStringFlag, NewBoolFlag, and NewInt64Flag.
These functions require the name and description of the flag as well as declaring whether the flag is required or not. 
Flags that are set as required and are not entered as a command line argument will return an error when parsed. 

# example
```go
// create a required string flag 
name, err := NewStringFlag("name", "set the user's name", true)
if err != nil {
    // display help message
    lighter.HelpWithError(err)
    os.Exit(1)
}

// create a required int flag 
age, err := NewInt64Flag("age", "set user's age", true)
if err != nil {
    // display help message
    lighter.HelpWithError(err)
    os.Exit(1)
}

// create an optional bool flag 
admin, err := NewBoolFlag("admin", "set admin privileges", false )
if err != nil {
    // display help message
    lighter.HelpWithError(err)
    os.Exit(1)
}

if err := lighter.Parse; err != nil {
    // display help message
    lighter.HelpWithError(err)
    os.Exit(1)
}

fmt.Println("name:", name.Value())
fmt.Println("age:", age.Value())
fmt.Println("is admin:", admin.Value())
```

The following:
```
./MyApp --name john --age 32
```

Will result in:
```
name: john
age: 32
is admin: false
```
