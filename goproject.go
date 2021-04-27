package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func main() {
	// get the number of arguments provided
	argLength := len(os.Args[1:])

	// if more than 1, then take the first and use as module name
	if argLength >= 1 {
		name := os.Args[1]
		path := os.Args[2]

		err := os.MkdirAll(path, 0755)
		if err != nil {
			log.Fatal(err)
		}
		os.Chdir(path) // change the path to the intended destination

		out, err := exec.Command("go", "version").Output()
		if err != nil {
			log.Fatal(err)
		}

		gover := string(out[0:2]) + " " + string(out[13:17])
		moduletext := fmt.Sprintf("module %s\n\n%s\n", name, gover)

		f, err := os.Create("go.mod")
		if err != nil {
			log.Fatal(err)
		}
		defer f.Close()
		f.WriteString(moduletext)

		complete := name + ".go"

		m, err := os.Create(complete)
		if err != nil {
			log.Fatal(err)
		}
		defer m.Close()
		packagetext := fmt.Sprintf("package main\n\nimport (\n\t\"fmt\"\n)\n\nfunc main() {\n\n\tfmt.Printf(\"Hello World\")\n}")
		m.WriteString(packagetext)

		fmt.Printf("Module created in folder %s\n", path)
	} else {
		fmt.Printf("Please provide the name of your module as argument to the program, and path as the 2nd argument")
	}
}
