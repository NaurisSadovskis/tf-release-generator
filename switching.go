package main

import "fmt"
import "os"

func main() {

  // this asks for input and validates it

  var inp_aws_region, inp_aws_account, inp_environment, inp_cost_code, inp_owner, inp_project_name string

  fmt.Print("region: ")
  fmt.Scanf("%s", &inp_aws_region)

  fmt.Print("aws account: ")
  fmt.Scanf("%s", &inp_aws_account)

  fmt.Print("aws project name: ")
  fmt.Scanf("%s", &inp_project_name)

  fmt.Print("environemnt: ")
  fmt.Println("environemnt: ")
  fmt.Scanf("%s", &inp_environment)

  fmt.Print("cost code: ")
  fmt.Scanf("%s", &inp_cost_code)


  // this is a map
  aws := map[string]map[string]string{
      "main": map[string]string {
          "aws_region"  : inp_aws_region,
          "aws_account" : inp_aws_account,
          "environment" : inp_environment,
          "cost_code"   : inp_cost_code,
          "project_name": inp_project_name,
          "owner"       : inp_owner,
      },
    "s3": map[string]string {
          "bucket_name"  : inp_project_name + "-" + inp_environment + "-" + "backup",
          "backup_user"  :"Hydrogen",
        },
    "ec2": map[string]string {
          "ec2_name"      : inp_project_name + "-" + inp_environment,
          "ec2_hostname"  :"Hydrogen",
        },
    "rds": map[string]string {
          "rds_name"      : inp_project_name + "-" + inp_environment + "-" + "db",
          "ec2_hostname"  :"Hydrogen",
        },
      }

    os.Mkdir("aws",0777)
    os.Mkdir("aws" + "/" + "s3" + "-" + aws["s3"]["bucket_name"], 0777)
    os.Mkdir("aws" + "/" + "ec2" + "-" + aws["ec2"]["ec2_name"], 0777)
    os.Mkdir("aws" + "/" + "rds" + "-" + aws["rds"]["rds_name"], 0777)

}
