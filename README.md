# Password Generator

A simple and flexible password generator written in Go.

## Features

- Generate passwords of any length
- Customizable character sets:
  - Digits (0-9)
  - Special characters (!@#$%^&*()_+{}|:<>?-=[]\;',./)
  - Lowercase letters (a-z)
  - Uppercase letters (A-Z)
- Ability to include or exclude any character set

## Usage

To use the password generator in your Go code:

```go
import "your-module-name/passwordgenerator"

// Generate a password of length 20 with all character sets
password := passwordgenerator.Generate(20, true, true, true, true)
```

The `Generate` function takes the following parameters:

1. `length`: Length of the password
2. `isDigit`: Include digits
3. `isSpecial`: Include special characters
4. `isLowercase`: Include lowercase letters
5. `isUppercase`: Include uppercase letters

## Example

```go
package main

import (
    "fmt"
    "your-module-name/passwordgenerator"
)

func main() {
    // Generate a 16-character password with digits and letters (no special characters)
    password := passwordgenerator.Generate(16, true, false, true, true)
    fmt.Println("Generated Password:", password)
}
```

## Note

This generator uses Go's `math/rand` package with its default source. For cryptographically secure random numbers, consider using `crypto/rand` instead.

## License

This project is licensed under the GNU General Public License v3.0 (GNU GPLv3). For more details, see the [LICENSE](https://www.gnu.org/licenses/gpl-3.0.en.html) file.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
