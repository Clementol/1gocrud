package validators

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"

	"github.com/Clementol/1gocrud/models"
	// "github.com/thedevsaddam/govalidator"
)

type RequiredValidator struct {
	Value string
}

type LengthValidator struct {
	Min int
	Max int
}
type Validator interface {
	Validate(interface{}) (bool, error)
}

type DefaultValidator struct {
}

type EmailValidator struct {
	Email string
}

func (v DefaultValidator) Validate(val interface{}) (bool, error) {
	return true, nil
}

func (v RequiredValidator) Validate(val interface{}) (bool, error) {
	valLength := len(val.(string))
	if valLength == 0 {
		return false, fmt.Errorf("is required")
	}
	return true, nil
}

func (v LengthValidator) Validate(val interface{}) (bool, error) {
	// fmt.Println("value", v.Value)
	fmt.Println("num val", reflect.TypeOf(val), v.Max)
	valLength := len(val.(string))
	if len(val.(string)) == 0 {
		return false, fmt.Errorf("is required")
	}

	numCheck := regexp.MustCompile(`[0-9]+`)

	fmt.Println(val)
	if numCheck.MatchString(val.(string)) {
		return false, fmt.Errorf("must be text")

	}

	if valLength < v.Min || valLength > v.Max {
		return false, fmt.Errorf("must be between %v to %v characters long", v.Min, v.Max)
	}
	return true, nil
}

func (v EmailValidator) Validate(val interface{}) (bool, error) {
	emailToString := val.(string)
	if len(emailToString) == 0 {
		return false, fmt.Errorf("is required")
	}
	var mailRe = regexp.MustCompile(`\A[\w+\-.]+@[a-z\d\-]+(\.[a-z]+)*\.[a-z]+\z`)
	if !mailRe.MatchString(emailToString) {
		return false, fmt.Errorf("is not a valid email address")
	}
	return true, nil
}

func getValidatorFromTag(tag string) Validator {

	args := strings.Split(tag, ",")
	// required, _ := regexp.MatchString("required", tag)
	minMax, _ := regexp.MatchString("min=([0-9]+),max=([0-9]+)", tag)
	email, _ := regexp.MatchString("email", tag)

	fmt.Println(minMax)
	switch {
	case minMax:
		validator := LengthValidator{}
		fmt.Sscanf(strings.Join(args[1:], ","), "min=%d,max=%d",
			&validator.Min, &validator.Max)
		return validator
	case email:
		validator := EmailValidator{}
		return validator
	}

	return DefaultValidator{}

}

func EmployeeValidation(employee models.Employee) error {
	errs := []error{}
	var errMsg error
	v := reflect.ValueOf(employee)
	// fmt.Println(v.Field(1).Interface())
	for i := 0; i < v.NumField(); i++ {
		tag := v.Type().Field(i).Tag.Get("validate")
		// fmt.Println(tag)
		if tag == "" || tag == "-" {
			continue
		}

		validator := getValidatorFromTag(tag)
		valid, err := validator.Validate(v.Field(i).Interface())
		if !valid && err != nil {
			errs = append(errs, fmt.Errorf("%s %s", v.Type().Field(i).Name, err.Error()))

		}
	}

	fmt.Println(errs)

	if len(errs) > 0 {
		errMsg = errs[0]
	}
	return errMsg

}
