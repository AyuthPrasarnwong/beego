package validation

import (
	"api/helpers"
	"fmt"
	"reflect"
	"regexp"
	"strings"

	enTranslations "gopkg.in/go-playground/validator.v9/translations/en"
	//jaTranslations "gopkg.in/go-playground/validator.v9/translations/ja"
	"github.com/astaxie/beego/orm"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"

	"gopkg.in/go-playground/validator.v9"
)


var (
	Uni *ut.UniversalTranslator
	Validate *validator.Validate

)

func Run() {
	Validate = validator.New()
	Validate.RegisterValidation("datetime", dateTime)
	Validate.RegisterValidation("date", date)
	Validate.RegisterValidation("phone", phone)
	Validate.RegisterValidation("uniqueDB", uniqueDB)

	en := en.New()
	Uni = ut.New(en, en)

	trans, _ := Uni.GetTranslator("en")

	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	_ = Validate.RegisterTranslation("required", trans, func(ut ut.Translator) error {
		return ut.Add("required", "The {0} field is required", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {

		fmt.Println("validator.FieldError", fe)
		t, _ := ut.T("required", fe.Field())
		return t
	})

	_ = Validate.RegisterTranslation("required_without", trans, func(ut ut.Translator) error {
		return ut.Add("required_without", "The {0} field is required when {1} is not present", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {

		param :=  helpers.ToSnakeCase(fe.Param())

		field :=  helpers.ToSnakeCase(fe.Field())

		enum := []string{"en", "th"}

		if helpers.InSlice(field, enum) == true {

			namespace := fe.Namespace()

			keys := strings.Split(namespace, ".")

			len := len(keys)

			field = keys[len - 2] + "." + keys[len - 1]

			field = helpers.ToSnakeCase(field)

			param = keys[len - 2] + "." + param
		}

		t, _ := ut.T("required_without", field, param)
		return t
	})

	_ = Validate.RegisterTranslation("oneof", trans, func(ut ut.Translator) error {
		return ut.Add("oneof", "The {0} must be one of the following types: {1}", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		params := strings.Replace(fe.Param(), " ", ", ", -1)
		t, _ := ut.T("oneof", fe.Field(), params)
		return t
	})

	_ = Validate.RegisterTranslation("email", trans, func(ut ut.Translator) error {
		return ut.Add("email", "The {0} must be a valid email address", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("email", fe.Field())
		return t
	})

	_ = Validate.RegisterTranslation("datetime", trans, func(ut ut.Translator) error {
		return ut.Add("datetime", "The {0} does not match the format {1}", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		format := "Y-m-d H:i:s"
		t, _ := ut.T("datetime", fe.Field(), format)
		return t
	})

	_ = Validate.RegisterTranslation("date", trans, func(ut ut.Translator) error {
		return ut.Add("date", "The {0} does not match the format {1}", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		format := "Y-m-d"
		t, _ := ut.T("date", fe.Field(), format)
		return t
	})

	_ = Validate.RegisterTranslation("phone", trans, func(ut ut.Translator) error {
		return ut.Add("phone", "The {0} must be a valid telephone or mobile phone number", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("phone", fe.Field())
		return t
	})

	_ = Validate.RegisterTranslation("uniqueDB", trans, func(ut ut.Translator) error {
		return ut.Add("uniqueDB", "The {0} has already been taken", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("uniqueDB", fe.Field())
		return t
	})

	_ = Validate.RegisterTranslation("max", trans, func(ut ut.Translator) error {
		return ut.Add("max", "The {0} may not be greater than {1}", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("max", fe.Field(), fe.Param())
		return t
	})

	_ = Validate.RegisterTranslation("min", trans, func(ut ut.Translator) error {
		return ut.Add("min", "The {0} must be at least {1}", true) // see universal-translator for details
	}, func(ut ut.Translator, fe validator.FieldError) string {
		t, _ := ut.T("min", fe.Field(), fe.Param())
		return t
	})

	// this is usually know or extracted from http 'Accept-Language' header
	// also see uni.FindTranslator(...)

	enTranslations.RegisterDefaultTranslations(Validate, trans)
}


func dateTime(fl validator.FieldLevel) bool {

	var dateTimePattern = regexp.MustCompile(`^(19|2[0-9])[0-9]{2}-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01]) (0[0-9]|1[0-9]|2[0-3]):([0-5][0-9]):([0-5][0-9])((\\+|-)[0-1][0-9]{3})?$`)

	input := fl.Field().String()

	match := dateTimePattern.FindString(input)

	if match == "" {

		return false
	}

	return true
}

func date(fl validator.FieldLevel) bool {

	var datePattern = regexp.MustCompile(`^([12]\d{3}-(0[1-9]|1[0-2])-(0[1-9]|[12]\d|3[01]))?$`)

	input := fl.Field().String()

	match := datePattern.FindString(input)

	if match == "" {

		return false
	}

	return true
}

func mobile(fl validator.FieldLevel) bool {

	var mobilePattern = regexp.MustCompile(`^(\+\d{1,3}[- ]?)?\d{10}$`)

	input := fl.Field().String()

	match := mobilePattern.FindString(input)

	if match == "" {
		return false
	}

	return true
}

func tel(fl validator.FieldLevel) bool {

	var telPattern = regexp.MustCompile(`^(0\d{2,3}(\-)?)?\d{7,8}$`)

	input := fl.Field().String()

	match := telPattern.FindString(input)

	if match == "" {
		return false
	}

	return true
}

func phone(fl validator.FieldLevel) bool {

	var mobilePattern = regexp.MustCompile(`^(\+\d{1,3}[- ]?)?\d{10}$`)

	var telPattern = regexp.MustCompile(`^(0\d{2,3}(\-)?)?\d{9}$`)

	input := fl.Field().String()

	matchMobile := mobilePattern.FindString(input)

	matchTel := telPattern.FindString(input)

	if matchTel == "" && matchMobile == "" {
		return false
	}

	return true
}

func uniqueDB(fl validator.FieldLevel) bool {

	v, kind, _ := fl.ExtractType(fl.Field())

	var input interface{}
	switch kind {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		input = v.Int()
	case reflect.String:
		input = v.String()
	}

	param := strings.Split(fl.Param(), " ")
	table := param[0]
	field := param[1]

	o := orm.NewOrm()
	exist := o.QueryTable(table).Filter(field, input).Exist()

	if exist {
		return false
	}

	return true
}