# errlist

Wrapping errors in Go can be tricky. To elaborate on the package-level `ErrNotFound`,

```golang
var (
	ErrNotFound = errors.New("not found")
)
```

one possible way is to call `fmt.Errorf`.

```golang
func ElaborateNotFound(username string) error {
	return fmt.Errorf("%w: user %s", ErrNotFound, username)
}

func main() {
	err := ElaborateNotFound("foobar") 
	fmt.Println(errors.Is(err, ErrNotFound)) // true
}
```

However, with this approach it's not possible to question whether
the _wrapper_ of `ErrNotFound` is of certain type (or matches certain package-level variable).

A bit of boilterplate is needed to allow nested wrapping:

```golang
type ErrOutside struct {
	Err error
}

func (e ErrOutside) Error() string {
	return fmt.Sprintf("outside: %s", e.Err)
}

func (e ErrOutside) Is(target error) bool {
	if _, ok := target.(ErrOutside); ok {
		return true
	}
	return errors.Is(e.Err, target)
}

var ErrInside = errors.New("inside")

func main() {
	insideWithReason := fmt.Errorf("%w: %s", ErrInside, "whatever reason")
	outside := ErrOutside{Err: insideWithReason}
	fmt.Println(errors.Is(outside, ErrInside)) // true
	fmt.Println(errors.Is(outside, ErrOutside{})) // true
}
```

## And here comes errlist

No need for fully fledged type definition any more. Errors defined with `errors.New` are just as fine for nested wrapping!

```golang
var (
	ErrOutside = errors.New("outside")
	ErrInside  = errors.New("inside")
)

func main() {
	err := errlist.New(ErrOutside, ErrInside, errors.New("whatever reason"))
	fmt.Println(errors.Is(err, ErrInside)) // true
	fmt.Println(errors.Is(err, ErrOutside)) // true
}
```
