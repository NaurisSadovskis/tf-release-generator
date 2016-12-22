package main

import "os"
import "fmt"
import "bufio"
import "strings"

var filename string = "variables.tf"
// const = variable names that can't be changed later

func help() {
  fmt.Println("This is help. Run 'generator init'")
}

func prompt() {
  fmt.Println(filename + " exists! Would you like to overwrite it? (y/n/exit)")

  var input string
  fmt.Scanf("%s", &input)
  response := input

  if response == "y" || response == "Y" {
    var file, err = os.Create(filename)
    if err != nil {
      fmt.Println(err)
      os.Exit(1)
    }
    defer file.Close()
    generate()
  } else if response == "n" || response == "N" {
    fmt.Println("Aborting.")
  } else if response == "exit" || response == "EXIT" {
    fmt.Println("Exiting.")
  } else {
      prompt()
  }
}

func initialize() {
  if _, err := os.Stat(filename); os.IsNotExist(err) {
    generate()
    fmt.Println(filename + " initialized.")
  } else {
    prompt()
  }
}

func generate() {

  variables := []string{"aws_region",
                        "aws_account",
                        "environment",
                        "cost_code",
                        "project",
                        "owner",
                        "test_user"}

  template := make([]string, len(variables))

// below will be useful when generating teraform files
  // for i := 0; i < len(variables); i++ {
  //   res := "variable " + variables[i] + " {\n\tdescription = \"<modify>\"\n}\n"
  //   template[i] = res
  // }


    for i := 0; i < len(variables); i++ {
      res := variables[i] + ":\n\t\t\"<modify>\"\n"
      template[i] = res
    }

  // array has to be converted before parsing to byte[0]
  output := "\x00" + strings.Join(template, "\x00") // x20 = space and x00 = null

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

func validate() {
  inputFile, err := os.Open(filename)
  if err != nil {
    fmt.Println(err)
  }
  defer inputFile.Close()

  // m = make(map[string]string)

  scanner := bufio.NewScanner(inputFile)
  //var results []string
  for scanner.Scan() {
      // here we need to add checks
      // we basically want to read each line and add it to a variable
      fmt.Println("hello")
      }
  }



func main() {

    // version := "0.0.1alpha"
    // filename := "variables.tf"
    if len(os.Args) == 1 {
      help()
    } else if os.Args[1] == "init" {
      initialize()
    } else if os.Args[1] == "validate" {
      validate()
    } else {
        fmt.Println(len(os.Args))
    }
}
