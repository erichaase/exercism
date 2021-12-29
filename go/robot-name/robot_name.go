package robotname

import (
	"errors"
	"fmt"
	"math/rand"
)

var names []string

func init() {
	// build list of available names
	for c1 := 'A'; c1 <= 'Z'; c1++ {
		for c2 := 'A'; c2 <= 'Z'; c2++ {
			for i := 0; i <= 999; i++ {
				k := fmt.Sprintf("%c%c%03d", c1, c2, i)
				names = append(names, k)
			}
		}
	}

	// shuffle list of available names
	for i := range names {
		j := rand.Intn(i + 1)
		names[i], names[j] = names[j], names[i]
	}
}

type Robot struct {
	name string
}

func (r *Robot) Name() (string, error) {
	if r.name != "" {
		return r.name, nil
	}
	return r.assignNewName()
}

func (r *Robot) Reset() {
	r.assignNewName()
}

func (r *Robot) assignNewName() (string, error) {
	n := len(names)
	if n <= 0 {
		return "", errors.New("No more available names")
	}

	i := n - 1
	r.name = names[i]
	names = names[:i]
	return r.name, nil
}
