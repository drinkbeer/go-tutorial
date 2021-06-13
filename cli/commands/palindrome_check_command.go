package cli

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type PalindromeCheck struct {
	inPath string
	strict bool
	flags  *flag.FlagSet
}

func NewPalindromeCheck() *PalindromeCheck {
	cmd := &PalindromeCheck{}
	flags := flag.NewFlagSet("Setup command", flag.ExitOnError)
	flags.StringVar(&cmd.inPath, "i", "", "The path to the test file with strings.")
	flags.BoolVar(&cmd.strict, "s", false, "If strict check or loose check the strings.")
	cmd.flags = flags

	return cmd
}

func (ps *PalindromeCheck) Name() string {
	return "palindrome-check"
}

func (ps *PalindromeCheck) Usage() {
	fmt.Printf("Usage: %s -i <path_to_fixtures> -s <strict_check_or_not>\n", ps.Name())
	ps.flags.PrintDefaults()
}

func (ps *PalindromeCheck) Description() string {
	return "The command checks if the string is palindrome."
}

func (ps *PalindromeCheck) Run(args ...string) error {
	err := ps.flags.Parse(args)
	if err != nil {
		return fmt.Errorf("Failed to parse arguments %s, error %s", args, err.Error())
	}

	fmt.Println("Start checking palindrome in file", ps.inPath, ", strict check", ps.strict)
	files := mustReadDir(ps.inPath)
	for _, f := range files {
		if strings.Index(f, "palindrome") == -1 {
			continue
		}
		file, _ := os.Open(f)
		result := checkFile(file, ps.strict)
		fmt.Println("Number of palindrome in path", f, "is", result)
	}
	return nil
}

func mustReadDir(path string) []string {
	dir, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer func(dir *os.File) {
		err := dir.Close()
		if err != nil {
			panic(err)
		}
	}(dir)

	names, err := dir.Readdirnames(0)
	if err != nil {
		panic(err)
	}

	// Construct full path.
	for i := range names {
		names[i] = filepath.Join(path, names[i])
	}

	return names
}

func checkFile(path *os.File, strict bool) int {
	count := 0
	scanner := bufio.NewScanner(path)
	for scanner.Scan() {
		word := scanner.Text()
		if len(word) == 0 {
			continue
		}
		result := checkString(word, strict)
		if result {
			fmt.Println(word)
			count++
		}
	}
	return count
}

func checkString(str string, strict bool) bool {
	if !strict {
		str = strings.Trim(str, " ")
		str = strings.Trim(str, ".")
		str = strings.Trim(str, ",")
		str = strings.ToLower(str)
	}

	for i := 0; i < len(str)/2; i++ {
		if str[i] != str[len(str)-1-i] {
			return false
		}
	}

	return true
}
