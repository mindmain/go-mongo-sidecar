package types

import (
	"os"
	"strconv"
	"strings"
)

type EvirontmentVariable string

func (e EvirontmentVariable) get() string {
	return os.Getenv(string(e))
}

func (e EvirontmentVariable) Get() string {

	if e.IsEmpty() {

		if defaultValue, ok := defaultValues[e]; ok {
			return defaultValue
		}
	}

	return e.get()
}

func (e EvirontmentVariable) IsEmpty() bool {
	return e.get() == ""
}

func (e EvirontmentVariable) IsNotEmpty() bool {
	return !e.IsEmpty()
}

func (e EvirontmentVariable) IsSet() bool {
	return e.IsNotEmpty()
}

func (e EvirontmentVariable) List() []string {
	return strings.Split(e.get(), ",")
}

func (e EvirontmentVariable) Int64() int64 {

	s := e.Get()

	if s == "" {
		return 0
	}

	i, err := strconv.Atoi(s)

	if err != nil {
		return 0
	}

	return int64(i)

}
