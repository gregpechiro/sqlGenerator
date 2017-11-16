package main

import (
	"bytes"
	"flag"
	"io/ioutil"
	"log"
	"strings"

	"github.com/gregpechiro/sqlGenerator"
)

var (
	out = flag.String("o", "", "what file to write")
	pkg = flag.String("pkg", ".", "what package to get the interface from")
)

func main() {
	flag.Parse()
	log.SetFlags(0)

	st := flag.Arg(0)

	if st == "" {
		log.Fatal("need to specify a struct name")
	}

	gen, err := sqlGenerator.NewSqlGenerator(*pkg, st)
	if err != nil {
		log.Fatalf("main.go >> unmarchalmap.NewGenerator() >> %v\n", err)
	}

	var buf bytes.Buffer

	log.Printf("Generating SQL Service Layer for %s\n", st)
	err = gen.Write(&buf)
	if err != nil {
		log.Fatalf("main.g0 >> gen.Write() >> %v\n", err)
	}

	if *out == "" {
		*out = ToLowerFirst(st) + "SqlService.go"

	}
	if err := ioutil.WriteFile(*out, buf.Bytes(), 0666); err != nil {
		log.Fatalf("main.go >> ioutil.WriteFile() >> %v\n", err)
	}
}

func ToLowerFirst(s string) string {
	if len(s) > 1 {
		return strings.ToLower(string(s[0])) + s[1:]
	}
	return strings.ToLower(s)
}
