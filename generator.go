package main

import "os"
import "fmt"
import "bufio"
import "strings"
import "io/ioutil"

func help() {
  fmt.Println("This is help. Run 'generator init'")
}

func prompt(filename string) {
  fmt.Println(filename + " exists! Would you like to overwrite it? (y/n/exit)")
  reader := bufio.NewReader(os.Stdin)
  response, _ := reader.ReadString('\n')
  if response == "y\n" || response == "Y\n" {
    var file, err = os.Create(filename)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
    defer file.Close()
  } else if response == "n\n" || response == "N\n" {
    fmt.Println("Aborting.")
  } else if response == "exit\n" || response == "EXIT\n" {
    fmt.Println("Exiting.")
  } else {
      prompt(filename)
  }
  generate(filename)
}

func initialize(filename string) {
  if _, err := os.Stat(filename); os.IsNotExist(err) {
    generate(filename)
    fmt.Println(filename + " initialized.")
  } else {
    prompt(filename)
  }
}

func generate(filename string) {

  variables := []string{"aws_region",
                        "aws_account",
                        "environment",
                        "cost_code",
                        "project",
                        "owner",
                        "test_user"}

  template := make([]string, len(variables))

  for i := 0; i < len(variables); i++ {
    res := "variable " + variables[i] + " {\n\tdescription = \"<modify>\"\n}\n"
    template[i] = res
  }

  // array has to be converted before parsing to byte[0]
  output := "\x00" + strings.Join(template, "\x20\x00") // x20 = space and x00 = null

  file, err := os.Create(filename)
  if err != nil {
      fmt.Println("Cannot create file", err)
      os.Exit(1)
  }

  defer file.Close()

  d2 := []byte(output)
  n2, err := file.Write(d2)
  fmt.Printf("wrote %d bytes\n", n2)
  if err != nil {
  	fmt.Println(err)

  }
}

func validate(filename string) {
  file, err := os.Open(filename) // For read access.
  if err != nil {
  	fmt.Println(err)
  }
  b, err := ioutil.ReadFile(file)
  fmt.Print(b)
  // str := byte(file) // convert content to a 'string'
  //
  // fmt.Println(file) // print the content as a 'string'
}

func main() {

    // version := "0.0.1alpha"
    filename := "variables.tf"
    if len(os.Args) == 1 {
      help()
    } else if os.Args[1] == "init" {
      initialize(filename)
    } else if os.Args[1] == "validate" {
      validate(filename)
    } else {
        fmt.Println(len(os.Args))
    }
}
