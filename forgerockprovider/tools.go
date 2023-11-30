package forgerockprovider

import (
	"math/rand"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

const voc string = "abcdfghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
const numbers string = "0123456789"
const symbols string = "!@#$&*+_-="

func init() {
	rand.Seed(time.Now().UnixNano())
}

func CreatePassword(length int, hasNumbers bool, hasSymbols bool) string {
	chars := voc
	if hasNumbers {
		chars = chars + numbers
	}
	if hasSymbols {
		chars = chars + symbols
	}
	return generatePassword(length, chars)
}

func generatePassword(length int, chars string) string {
	password := ""
	for i := 0; i < length; i++ {
		password += string([]rune(chars)[rand.Intn(len(chars))])
	}
	return password
}

/*
*
Iterate on the input TF object to check if all the fields are empty or not
*/
func checkEmptyTFObject(t any) bool {

	e := reflect.ValueOf(t)

	for i := 0; i < e.NumField(); i++ {
		varType := e.Type().Field(i).Type
		switch varType.String() {
		case "basetypes.ListValue":
			if !e.Field(i).Interface().(basetypes.ListValue).IsNull() {
				return false
			}
		case "basetypes.StringValue":
			if !e.Field(i).Interface().(basetypes.StringValue).IsNull() {
				return false
			}
		case "basetypes.Int64Value":
			if !e.Field(i).Interface().(basetypes.Int64Value).IsNull() {
				return false
			}
		case "basetypes.BoolValue":
			if !e.Field(i).Interface().(basetypes.BoolValue).IsNull() {
				return false
			}
		// TODO handle types that does not have IsNull()
		case "basetypes.Float64Type":
			continue
		case "basetypes.ObjectType":
			continue
		case "basetypes.MapType":
			continue
		default:
			panic("Unknown type used in updateTerraformValue " + reflect.TypeOf(varType).String())
		}

	}
	return true
}
