# Wiz Repository

Collection of boilerplate Go code

### Design Mandates

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

AES.go
```
AESEncrypt(data []byte, key []byte) ([]byte, error)
AESDecrypt(stream []byte, key []byte) ([]byte, error)
```
Args.go
```
Args() []string //Return command line arguments passed to program
```
ASCII.go
```
ASCII(b []byte) (string, bool)
Printable(b []byte) (string, bool)
StripNonASCII(in string) string
StripNonPrintableASCII(in string) string
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
Print(items ...interface{})
```
Ed25519.go
```
NewEdKeyPair(seed []byte) ([]byte, []byte, error)
EdSign(data, publicKey, privateKey []byte) ([]byte, error)
EdVerify(data, signature, publicKey []byte) error
```
Files.go
```
Executable() string
ProgramName() string
Dir() string
FileExists(file string) bool
FolderExists(file string) bool
DeleteFile(file string) error
MkDir(dir string) error
WriteFile(file string, data []byte) error
ReadFile(file string) ([]byte, error)
```
Hash.go
```
Hash(data []byte) []byte
HashMatch(data []byte, hash []byte) bool
```
Hex.go
```
BytesToHex(data []byte) string
HexToBytes(data string) ([]byte, error)
```
HTTP.go
```
SplitURL(url string) []string
ServeSimple(ln net.Listener, getter func([]string) (int, []byte), poster func([]string, []byte) (int, []byte)) error
NewClient(c *http.Client, timeout int) Client

type Client
Client.Get(url string) ([]byte, error)
Client.GetStruct(url string, responseVessel interface{}) error
Client.Post(url string, requestBody []byte) ([]byte, error)
Client.PostStruct(url string, requestPayload interface{}, responseVessel interface{}) error
```
JSON.go
```
CompactJSON(data []byte) ([]byte, error)
NeatJSON(data []byte) ([]byte, error)
Marshal(payload interface{}) ([]byte, error)
MarshalNeat(payload interface{}) ([]byte, error)
```
Random.go
```
RandomBytes(len int) ([]byte, error)
```
SQLite.go
```
SQLiteOpen(err *error, dbName string) SimpleDatabase

type SimpleDatabase
SimpleDatabase.Close()
SimpleDatabase.MakeTable(table string) error
SimpleDatabase.ClearTable(table string) error
SimpleDatabase.AddItemAt(table string, key uint64, data string) error
SimpleDatabase.GetItemAt(table string, key uint64) (string, error)
SimpleDatabase.DeleteItemAt(table string, key uint64) error
SimpleDatabase.GetKeys(table string) ([]uint64, error)
SimpleDatabase.CheckOrder(table string) (uint64, error)
SimpleDatabase.Latest(table string) (uint64, error)

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
Sleep(seconds int)
```
Uint64.go
```
Uint64(interface{}) (uint64, error)
```
