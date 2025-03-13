package helper

import (
	"bufio"
	"iter"
	"log"
	"os"
	"path"
	"runtime"
)

func ReadLine(filename string) iter.Seq[string] {
	_, callerPath, _, ok := runtime.Caller(1)
	if !ok {
		log.Fatal("Cannot find Caller of helper.ReadLine")
	}
	f, err := os.Open(path.Join(path.Dir(path.Dir(callerPath)), "data", filename))
	if err != nil {
		log.Fatal(err)
	}

	s := bufio.NewScanner(f)
	return func(yield func(string) bool) {
		for s.Scan() {
			if !yield(s.Text()) {
				f.Close()
				break
			}
		}
	}
}
