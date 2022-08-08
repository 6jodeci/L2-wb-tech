package util

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// Flags инициализирует флаги
type Flags struct {
	FlagA     int  // -A
	FlagB     int  // -B
	FlagC     int  // -C
	FlagCount bool // -c
	Flagi     bool // -i
	Flagv     bool // -v
	FlagF     bool // -F
	Flagn     bool // -n
}

// NewFlags реализует структуру flags для работы с флагами
func NewFlags(flagA int, flagB int, flagC int, FlagCount bool, flagi bool, flagv bool, flagF bool, flagn bool) *Flags {
	return &Flags{
		FlagA:     flagA,
		FlagB:     flagB,
		FlagC:     flagC,
		FlagCount: FlagCount,
		Flagi:     flagi,
		Flagv:     flagv,
		FlagF:     flagF,
		Flagn:     flagn,
	}
}

type Checker interface {
	check(s string) bool
}

type RegexChecker struct {
	R *regexp.Regexp
}

func NewRegexChecker(r *regexp.Regexp) *RegexChecker {
	return &RegexChecker{R: r}
}

func (c *RegexChecker) check(s string) bool {
	return c.R.MatchString(s)
}

type EqualChecker struct {
	pat string
}

// NewEqualChecker проверяет точное совпадение со строкой
func NewEqualChecker(pat string) *EqualChecker {
	return &EqualChecker{pat: pat}
}

func (c *EqualChecker) check(s string) bool {
	return strings.Contains(s, c.pat)
}

// Grep собирает все аргументы
func Grep(r io.Reader, pattern string, ss *Flags) (string, error) {
	var chr Checker
	// переводит в ловеркейс и проверяет совпадение
	if ss.Flagi {
		pattern = strings.ToLower(pattern)
	}
	// проверяет точное совпадение со строкой
	if ss.FlagF {
		chr = NewEqualChecker(pattern)
	} else {
		rx, err := regexp.Compile(pattern)
		if err != nil {
			return "", err
		}
		chr = NewRegexChecker(rx)
	}

	lt := ss.FlagC
	rt := ss.FlagC
	if ss.FlagB > 0 {
		lt = ss.FlagB
	}
	if ss.FlagA > 0 {
		rt = ss.FlagA
	}

	mm := make(map[int]bool)
	var resIds []int
	var allStr []string
	counter := 0
	i := 0
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		t := sc.Text()
		allStr = append(allStr, t)
		// игнорирует регистр
		if ss.Flagi {
			t = strings.ToLower(t)
		}

		if chr.check(t) {
			counter++
			if lt > 0 {
				for j := i - lt; j < i; j++ {
					if _, ok := mm[j]; !ok {
						resIds = append(resIds, j)
						mm[j] = false
					}
				}
			}
			if rt > 0 {
				for j := i + rt; j > i; j-- {
					if _, ok := mm[j]; !ok {
						resIds = append(resIds, j)
						mm[j] = false
					}
				}
			}
			if _, ok := mm[i]; !ok {
				resIds = append(resIds, i)
			}
			mm[i] = true
		}
		i++
	}

	// выводит количество точных сопадений с указаной строкой
	if ss.FlagF {
		return fmt.Sprintf("%v совпадений", counter), nil
	}

	sort.Ints(resIds)

	var res []string
	lIndex := ""
	if !ss.Flagi {
		match := "  "
		for _, k := range resIds {
			if k > -1 && k < len(allStr) {
				if mm[k] {
					match = "> "
				} else {
					match = "  "
				}
				if ss.Flagn {
					lIndex = strconv.Itoa(k) + " "
				}
				res = append(res, fmt.Sprintf("%v%v%v", lIndex, match, allStr[k]))
			}
		}
	} else {
		for i, v := range allStr {
			if _, ok := mm[i]; !ok {
				if ss.Flagn {
					lIndex = strconv.Itoa(i) + " "
				}
				res = append(res, fmt.Sprintf("%v%v", lIndex, v))
			}
		}
	}

	return strings.Join(res, "\n"), nil
}
