# Wiz Repository

Collection of boilerplate Go code

### Development Mandates

The wiz repository must first and foremost remain compatible with the following two use cases:

1. **import** github.com/toteki/wiz

    Exported functions can be called like so:
    ```
    s = wiz.Args()
    ```
2. **copy** individual files and paste into a project anywhere

    The first line of each file:
    ```
    package wiz
    ```
    must be changed to match the destination. Examples:
    ```
    package main
    ```
    ```
    package foo
    ```
    Beware name collisions with the target code in that case.

To provide stable support for both those use cases, additional patterns will be followed. Here is a working list:

- wiz functions in different files cannot call each other *(to support individual file copying)*
- error messages from wiz functions will use only the function name, rather than prefixing with *wiz.*
- unexported functions should be avoided, lest they become available in the target code in the file copying use case

### Best practices note:

It is best practice for packages to have a single purpose which is evident from their name *(e.g. github.com/pkg/errors)*

This is the exact opposite of ambiguously named utilities packages *e.g. github.com/toteki/wiz*

Nevertheless, I have yet to encounter a single project which could not be vastly accelerated by accessible boilerplate. Sacrifices must be made.

(Said sacrifices must be *unmade* efficiently when the dependency needs to be eliminated. This is the reasoning behind the second development mandate.)

### Wiz/Examples

The examples folder is **not** meant to be imported, nor is any subdirectory. It is for cannibalism only.

See examples/README.md for specifics.

### Exposed Functions By File

Args.go
```
Args() []string
Executable() string
ProgramName() string
```
Client.go
```
Get(err *error, url string) []byte
StructGet(err *error, url string, responsePayload interface{})
Post(err *error, url string, requestBody []byte) []byte
StructPost(err *error, url string, requestPayload interface{}, responsePayload interface{})
```
Console.go
```
SilentPrompt(prompt string) string
Prompt(prompt string) string
Red(items ...interface{})
Yellow(items ...interface{})
Green(items ...interface{})
Blue(items ...interface{})
Purple(items ...interface{})
White(items ...interface{})
Grey(items ...interface{})
```
Ed25519.go
```
NewEdKeyPair(err *error, seed []byte) ([]byte, []byte)
EdSign(err *error, data []byte, privateKey []byte, publicKey []byte) ([]byte)
EdVerify(err *error, data []byte, signature []byte, publicKey []byte)
```
Files.go
```
DeleteFile(err *error, file string)
MkDir(err *error, dir string)
WriteFile(err *error, file string, data []byte)
ReadFile(err *error, file string) []byte
```
Hash.go
```
Hash(data []byte) []byte
HashMatch(data []byte, hash []byte) bool
```
Hex.go
```
BytesToHex(data []byte) string
HexToBytes(err *error, data string) []byte
```
JSON.go
```
Marshal(err *error, payload interface{}) []byte
MarshalNeat(err *error, payload interface{}) []byte
StringMarshalNeat(err *error, payload interface{}) string
```
Random.go
```
RandomBytes(err *error, len int) []byte
```
Server.go
```
ServeHTTP(err *error, addr string, getter func([]string) (int, []byte), poster func([]string, []byte) (int, []byte)) (served string)
SplitURL(url string) []string
```
SQLite.go
```
type Database struct
SQLiteOpen(err *error, dbName string) Database
(*Database).Close()
(*Database).MakeTable(err *error, tableName string)
(*Database).ClearTable(err *error, tableName string)
(*Database).AddItemAt(err *error, tableName string, primaryKey uint64, data string)
(*Database).GetItemAt(err *error, tableName string, primaryKey uint64) string
(*Database).DeleteItemAt(err *error, tableName string, primaryKey uint64)
(*Database).GetKeys(err *error, tableName string) uint64
(*Database).CheckOrder(err *error, tableName string) uint64
(*Database).Latest(err *error, tableName string) uint64
```
Strings.go
```
Lowercase(string) string
Uppercase(string) string
String(interface{}) string
```
Time.go
```
Now() uint64
Sleep(milliseconds int)
```
Uint64.go
```
Uint64(err *error, interface{}) uint64
```
