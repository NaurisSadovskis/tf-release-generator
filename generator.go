package main

import "os"
import "fmt"
import "bufio"

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

func initialize() {
  filename := "file.txt"
  if _, err := os.Stat(filename); os.IsNotExist(err) {
    os.Create(filename)
    fmt.Println(filename + " initialized.")
  } else {
    prompt(filename)
  }
}

func generate(filename string) {
  // make template from a list
  // save it as template
  variables := [6]string{"aws_region", "aws_account", "aws_test", "aws_what", "environment", "dev"}
  template := make([]string, 6)
  for i := 0; i < len(variables); i++ {
    res := "variable " + variables[i] + " {\n\tdescription = \"fffffffff\"\n}"
    template[i] = res
  }

  slice := template[0:len(template)]

  fmt.Println(template)
  fmt.Println(slice)
  raw := slice[0:len(slice)]
  fmt.Println(raw)
  file, err := os.Create(filename)
  if err != nil {
      fmt.Println("Cannot create file", err)
  }
  defer file.Close()
  // fmt.Fprintf(file, slice)
  d2 := []byte("zen")
  n2, err := file.Write(d2)
  fmt.Printf("wrote %d bytes\n", n2)
  if err != nil {
  	fmt.Println(err)
    os.Exit(1)
  }

}

func main() {

    // version := "0.0.1alpha"

    if len(os.Args) == 1 {
      help()
    } else if os.Args[1] == "init" {
      initialize()
    } else {
        fmt.Println(len(os.Args))
    }
}
