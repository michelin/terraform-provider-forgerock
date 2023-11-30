package test

import (
	"fmt"
	"os"
	"os/exec"
	"terraform-provider-forgerock/sdk"
	"testing"

	"github.com/hashicorp/hcl/v2/hclsimple"
)

var sdkClient *sdk.Client

// ======================= TESTS INIT =======================

func TestInit(m *testing.M) {
	authentInit(m)
	terraformInit(m)
	err := m.Run()
	

	os.Remove("./usecase1/terraform.tfvars")
	os.Remove("./usecase2/terraform.tfvars")
	os.Remove("./usecase3/terraform.tfvars")
	os.Remove("./usecase4/terraform.tfvars")

	os.Exit(err)
}

func authentInit(m *testing.M) {
	if sdkClient == nil {
		fmt.Println("Initializing client...")
		error := clientAuthentication()
		if error != nil {
			fmt.Println("Error while creating client:", error)
			os.Exit(1)
		}
	}
	fmt.Println("Starting tests...")
}

func terraformInit(m *testing.M) {

	copy("./terraform.hcl", "./usecase1/terraform.tfvars")
	copy("./terraform.hcl", "./usecase2/terraform.tfvars")
	copy("./terraform.hcl", "./usecase3/terraform.tfvars")
	copy("./terraform.hcl", "./usecase4/terraform.tfvars")

	executeCommand("terraform", "-chdir=./usecase1", "init")
	executeCommand("terraform", "-chdir=./usecase2", "init")
	executeCommand("terraform", "-chdir=./usecase3", "init")
	executeCommand("terraform", "-chdir=./usecase4", "init")
}

// ======================= TOOLS =======================

func clientAuthentication() error {

	data, error := parseTerraformVars()
	if error != nil {
		return error
	}

	c, error := sdk.NewClient(data.Url, data.Realm, data.Username, data.Password)
	if error != nil {
		return error
	}
	sdkClient = c
	return nil
}

func executeCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Dir = "./"
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		fmt.Println("Error while executing command:", err)
		return err
	}
	return nil
}

type Tfvars struct {
	Url      string `hcl:"forgerock_api"`
	Realm    string `hcl:"realm_path"`
	Username string `hcl:"username"`
	Password string `hcl:"password"`
}

func parseTerraformVars() (Tfvars, error) {
	var tfvars Tfvars
	err := hclsimple.DecodeFile("./terraform.hcl", nil, &tfvars)
	if err != nil {
		return Tfvars{}, err
	}
	return tfvars, nil
}

func copy(src string, dst string) {

    data, err := os.ReadFile(src)
    checkErr(err)
    
	err = os.WriteFile(dst, data, 0644)
    checkErr(err)
}
func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}