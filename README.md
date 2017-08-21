# go-uava

Basic pre-conditions checking utilities, similar to a subset of Google Guava (minus the whole generics thing). 


Usage Example:
	
~~~~
func fnDef(arg1 string, arg2 bool) {
    err := check.Preconditions(check.NotEmptyString(arg1), check.Argument(arg2))
    if err != nil {
        ...
    }
    ...
}

~~~~