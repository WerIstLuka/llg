`CheckDebug()` check if the environmental variable 'DEBUG' is set to `true`

`GetDebug() *bool` returns a pointer to the debug variable to allow setting it manually

`PrintDebug(v ...any)` acts similiar to `fmt.Println` but only prints when debug is true \
to set debug either use `CheckDebug` or set it manually with `GetDebug` \
messages are prepended with `[DEBUG]`

`PrintError(v ...any)` acts similiar to `fmt.Println` but exits with code 1 \
messages are prepended with `[ERROR]`

`MaxStringLength(source []string) int` returns the length of the longest string in a slice of strings

`NextMultipleInt(source int, multiple int) int` returns the next multiple of source
