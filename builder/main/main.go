package main

import (
	"bytes"
	"flag"
	"github.com/orbs-network/contract-external-libraries-go/builder/templates/project/javascript"
	"io/ioutil"
	"os"
	"path"
	"strings"
	"text/template"
)

func main() {
	name := flag.String("name", "App", "app name")
	projectPath := flag.String("path", ".", "path to new project")

	flag.Parse()

	nameLowercase := strings.ToLower(*name)
	argsMap := getTemplateArgumentsMap(*name, nameLowercase)

	if err := writeFile(path.Join(*projectPath, "src", "contract"), "contract.go",
		renderTemplate(javascript.CONTRACT_SOURCE, argsMap)); err != nil {
			panic(err)
	}

	if err := writeFile(path.Join(*projectPath, "src"), "deploy.js",
		renderTemplate(javascript.JAVASCRIPT_DEPLOYMENT_SOURCE, argsMap)); err != nil {
		panic(err)
	}

	if err := writeFile(path.Join(*projectPath, "test"), nameLowercase + "_test.js",
		renderTemplate(javascript.JAVASCRIPT_TEST_SOURCE, argsMap)); err != nil {
		panic(err)
	}

	if err := writeFile(path.Join(*projectPath, "src"), nameLowercase + ".js",
		renderTemplate(javascript.JAVASCRIPT_INTERFACE_SOURCE, argsMap)); err != nil {
		panic(err)
	}

	if err := writeFile(path.Join(*projectPath), "package.json",
		renderTemplate(javascript.PACKAGE_JSON_SOURCE, argsMap)); err != nil {
			panic(err)
	}

}

func getTemplateArgumentsMap(name string, nameLowercase string) map[string]interface{} {
	return map[string]interface{}{
		"AppName": name,
		"AppNameLowercase": nameLowercase,
	}
}

func renderTemplate(source string, argsMap map[string]interface{}) string {
	t, err := template.New("").Parse(source)
	if err != nil {
		panic(err)
	}

	writer := bytes.NewBufferString("")
	t.Execute(writer, argsMap)
	return writer.String()
}

func writeFile(dir string, filename, contents string) error {
	os.MkdirAll(dir, 0744)
	return ioutil.WriteFile(path.Join(dir, filename), []byte(contents), 0644)
}