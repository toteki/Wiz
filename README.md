# Wiz Repository

Collection of boilerplate Go code

### Best practices note:

It is best practice for go packages to have a single purpose which is evident from their name *(e.g. net/http)*

This is the exact opposite of ambiguously named utilities packages, like this one.

Nevertheless, I have yet to encounter a single project which could not be vastly accelerated by a utilities package. Sacrifices must be made.

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
