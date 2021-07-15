package plugin_test

import (
	"fmt"
	"log"

	"github.com/Reisender/go-plugin"
)

//go:generate go build --buildmode=plugin -o ./example/some-plugin.so ./example/hello.go

func ExampleLoad() {
	type pluginInterface struct {
		Hello func(name string) string
		Add   func(int, int) int
		Mult  func(int, int) int `lookup:"Multiply"`
	}

	loadedPlugin := pluginInterface{}

	err := plugin.Load(&loadedPlugin, "./example/some-plugin.so")
	if err != nil {
		log.Fatal(fmt.Errorf("try running `go generate` first : %v", err))
	}

	fmt.Println(loadedPlugin.Hello("world"))
	// Output: Hello world!
}
