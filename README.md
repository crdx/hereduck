# hereduck

**hereduck** is a heredoc library for Go.

## Usage

```go
import "github.com/crdx/hereduck"

func main() {
    doc = hereduck.D(`
        Hello
            world
    `)
}
```

This would result in the following string in `doc`, including the trailing new line.

```
Hello
    world
```

The minimum indentation level shared by all lines is stripped out, but subsequent indentation is preserved.

There is also `hereduck.Df` which follows the standard formatting string pattern.

## Contributions

Open an [issue](https://github.com/crdx/hereduck/issues) or send a [pull request](https://github.com/crdx/hereduck/pulls).

## Licence

[MIT](LICENCE.md).
