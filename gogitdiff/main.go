package main

import (
	"fmt"
	"os"

	"github.com/hexops/gotextdiff"
	"github.com/hexops/gotextdiff/myers"
	"github.com/hexops/gotextdiff/span"
)

// error handling left out for brevity
func main() {
	file1, _ := os.ReadFile("./test/test.go")
	file2, _ := os.ReadFile("./test2/tests.go")
	edits := myers.ComputeEdits(span.URIFromPath("./test/test.go"), string(file1), string(file2))
	diff := fmt.Sprint(gotextdiff.ToUnified("./test/test.go", "./test2/tests.go", string(file1), edits))

    err := os.WriteFile("./tmp/dat1.go", []byte(diff), 0644)
    _ = (err) 

	/* input, err := ioutil.ReadFile("./tmp/dat1.go")
        if err != nil {
                log.Fatalln(err)
        }

        lines := strings.Split(string(input), "\n")
		lines[0] = ""
		lines[1] = ""
		lines[2] = ""

        for i, line := range lines {
			if(i==0){
				lines[i] = ""
			}
			if len(line)>1{
				lines[i] = line[1:]
			}
			if(i==len(lines)-2){
				lines[i] = ""
			}
        }
        output := strings.Join(lines, "\n")
        err = ioutil.WriteFile("./tmp/dat2.go", []byte(output), 0644)
        if err != nil {
                log.Fatalln(err)
        } */
}