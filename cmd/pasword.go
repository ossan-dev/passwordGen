/*
Copyright © 2024 Richard Fearnhead solutions@richard-fearnhead.com
*/
package cmd

import (
	"fmt"
	"math/rand"

	"github.com/spf13/cobra"
)

// TODO: write some tests
// TODO: extrapolate logic into "internal" pkg
// TODO: give to the user the option to specify the character set to use
// TODO: exclude the compiled binary in the .gitignore file

var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate random passwords",
	Long:  "Generate random passwords with cusomizable options. For example: passwordGen generate -l 12 -d -s",

	Run: generatePassword,
}

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("length", "l", 8, "Length of the generated password")
	generateCmd.Flags().BoolP("digits", "d", false, "include digits in the generated password")
	generateCmd.Flags().BoolP("special-chars", "s", false, "Include special chars in the generated password")
}

// FIXME: stick to primitive types so your tests will be less flaky.
func generatePassword(cmd *cobra.Command, args []string) {
	// FIXME: extrapolate these logic into a 'setupCommand' func
	length, _ := cmd.Flags().GetInt("length")
	isDigits, _ := cmd.Flags().GetBool("digits")
	isSpecialChars, _ := cmd.Flags().GetBool("special-chars")

	// FIXME: this can be either a pkg constant or provided as a param to the func
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

	if isDigits {
		charset += "0123456789"
	}

	if isSpecialChars {
		charset += "!@#$%^&*()_+{}[]|;:,.<>?-="
	}

	password := make([]byte, length)

	for i := range password {
		password[i] = charset[rand.Intn(len(charset))]
	}

	// FIXME: if you'd like to print something, stick to fmt.Fprint... functions. What you've used is not unit-testable
	fmt.Println("Generating Password")
	fmt.Println(string(password))
}
