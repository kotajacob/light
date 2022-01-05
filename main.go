// backlight - Print or change backlight brightness percentage
// Part of the useless software collection. License is found in root directory.
package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strconv"
	"strings"
)

const nowPath string = "/sys/class/backlight/intel_backlight/brightness"
const maxPath string = "/sys/class/backlight/intel_backlight/max_brightness"

// read brightness from file and return as int
func get(path string) int {
	b, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "backlight: reading %s\n", path)
		os.Exit(1)
	}
	s := strings.TrimSuffix(string(b), "\n")
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Fprintf(os.Stderr, "backlight: corrupt file %s\n", path)
		os.Exit(1)
	}
	return i
}

// read brightness value and mode from argument
func read(arg string) (int, int) {
	switch a := string(arg[0]); a {
	case "+":
		m := 1
		v := strip(arg)
		return m, v
	case "-":
		m := 2
		v := strip(arg)
		return m, v
	default:
		m := 0
		v := strip(arg)
		return m, v
	}
}

// strip the argument of its mode operator and convert to int
func strip(s string) int {
	i, err := strconv.Atoi(strings.Trim(s, "+-"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "backlight: bad argument\n")
		os.Exit(1)
	}
	return i
}

// check if new brightness value and mode are valid and return a proper int
func check(m, v, n int) int {
	switch m {
	case 1:
		// add
		i := n + v
		if i > 100 {
			i = 100
		}
		return i
	case 2:
		// subtract
		i := n - v
		if i < 0 {
			i = 0
		}
		return i
	default:
		// set
		if v > 100 || v < 0 {
			fmt.Println("Error: backlight: bad argument")
			return n
		} else {
			return v
		}
	}
}

// set the new brightness value
func set(v, m int) {
	v = int(float32(m) * (float32(v) / 100))
	s := strconv.Itoa(v) + "\n"
	d := []byte(s)
	err := ioutil.WriteFile(nowPath, d, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "backlight: writing %s\n", nowPath)
		os.Exit(1)
	}
}

// return an int percent from two int values
func percent(n, m int) int {
	return int(math.Ceil(float64(n) / (float64(m) / 100)))
}

func main() {
	switch a := len(os.Args); a {
	case 1:
		// print the brightness
		fmt.Println(percent(get(nowPath), get(maxPath)))
	case 2:
		// set the brightness
		rNow := get(nowPath)
		rMax := get(maxPath)
		now := percent(rNow, rMax)
		mode, value := read(os.Args[1])
		value = check(mode, value, now)
		set(value, rMax)
	default:
		// error for too many args
		fmt.Println("Error: backlight: too many arguments")
		os.Exit(1)
	}
}
