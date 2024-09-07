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

To use the password generator, run the program with the following command-line arguments:

```
go run main.go <length> <useDigits> <useSpecial> <useLower> <useUpper>
```

### Arguments:

1. `length`: An integer specifying the length of the password
2. `useDigits`: `true` or `false` to include digits
3. `useSpecial`: `true` or `false` to include special characters
4. `useLower`: `true` or `false` to include lowercase letters
5. `useUpper`: `true` or `false` to include uppercase letters

### Example:

```
go run main.go 20 true false false true
```

This command will generate a 20-character password using only digits and uppercase letters.

### Note:

- All arguments are required.
- Boolean arguments must be specified as `true` or `false`.
- At least one character set (digits, special, lowercase, or uppercase) must be set to `true`.

## Example in Go Code

If you want to use the password generator in your Go code:

```go
import "your-module-name/passwordgenerator"

// Generate a password of length 20 with all character sets
password := passwordgenerator.Generate(20, true, true, true, true)
```

## Note

This generator uses Go's `math/rand` package with its default source. For cryptographically secure random numbers, consider using `crypto/rand` instead.

## License

This project is licensed under the GNU General Public License v3.0 (GNU GPLv3). For more details, see the [LICENSE](https://www.gnu.org/licenses/gpl-3.0.en.html) file.

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
