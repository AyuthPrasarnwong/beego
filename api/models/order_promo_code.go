package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"regexp"
	//"sort"
	//"strconv"
	"strings"

	"time"

	"github.com/astaxie/beego/orm"
)

var (
	OrderPromoCodes map[string]*OrderPromoCode
)

type BrandName struct {
	BrandNameEn string `json:"en" validate:"required_without=BrandNameTh"`
	BrandNameTh string `json:"th" validate:"required_without=BrandNameEn"`
}

type OutletName struct {
	OutletNameEn string `json:"en" validate:"required_without=OutletNameTh"`
	OutletNameTh string `json:"th" validate:"required_without=OutletNameEn"`
}

type Locale struct {
	En string `json:"en"`
	Th string `json:"th"`
}

type OrderPromoCode struct {
	Id                  int64      `json:"id" orm:"auto;column(id);pk"`
	OrderId             int        `json:"order_id" orm:"column(order_id)"`
	OrderNo             string     `json:"order_no" orm:"column(order_no)"`
	PromoCode           string     `json:"promo_code" orm:"column(promo_code)"`
	CampaignId          string     `json:"campaign_id" orm:"column(campaign_id)"`
	CampaignName        string     `json:"campaign_name" orm:"column(campaign_name)"`
	CampaignDescription string     `json:"campaign_description" orm:"column(campaign_description)"`
	CampaignStartedAt   time.Time  `json:"campaign_started_at" orm:"column(campaign_started_at)"`
	CampaignEndedAt     time.Time  `json:"campaign_ended_at" orm:"column(campaign_ended_at)"`
	Status              string     `json:"status" orm:"column(status);type(enum);choices('success','fail','cancel','waiting')"`
	Amount              int        `json:"amount" orm:"column(amount)"`
	Discount            int        `json:"discount" orm:"column(discount)"`
	FinalAmount         int        `json:"final_amount" orm:"column(final_amount)"`
	CustomerId          int        `json:"customer_id" orm:"column(customer_id)"`
	CustomerName        string     `json:"customer_name" orm:"column(customer_name)"`
	CustomerGender      string     `json:"customer_gender" orm:"column(customer_gender);type(enum);choices('male', 'female', 'unknown')"`
	CustomerEmail       string     `json:"customer_email" orm:"null;column(customer_email)"`
	CustomerPhone       string     `json:"customer_phone" orm:"column(customer_phone)"`
	CustomerBirthday    string     `json:"customer_birthday," orm:"null;column(customer_birthday);type(date)"`
	BrandId             int        `json:"brand_id" orm:"column(brand_id)"`
	BrandName           string     `json:"brand_name" orm:"column(brand_name);type(json)"`
	OutletId            int        `json:"outlet_id" orm:"column(outlet_id)"`
	OutletName          string     `json:"outlet_name" orm:"column(outlet_name);type(json)"`
	TerminalId          int        `json:"terminal_id" orm:"column(terminal_id)"`
	TerminalNo          string     `json:"terminal_no" orm:"column(terminal_no)"`
	RedeemedAt          time.Time  `json:"redeemed_at" orm:"column(redeemed_at);"`
	Created             *time.Time `json:"created_at" orm:"column(created_at);auto_now_add;type(datetime)"`
	Updated             *time.Time `json:"updated_at" orm:"column(updated_at);auto_now;type(datetime)"`
}

type InputCreateOrderPromoCode struct {
	OrderId             int        `json:"order_id" validate:"required,min=1,max=1000000000,uniqueDB=order_promo_codes order_id"` //min:1 max:1000000000
	OrderNo             string     `json:"order_no" validate:"required,uniqueDB=order_promo_codes order_no"`
	PromoCode           string     `json:"promo_code" validate:"required"`
	CampaignId          string     `json:"campaign_id" validate:"required"`
	CampaignName        string     `json:"campaign_name" validate:"required"`
	CampaignDescription string     `json:"campaign_description"`
	CampaignStartedAt   string     `json:"campaign_started_at" validate:"required,datetime"`
	CampaignEndedAt     string     `json:"campaign_ended_at" validate:"required,datetime"`
	Status              string     `json:"status" validate:"required,oneof=success fail cancel waiting"`
	Amount              int        `json:"amount" validate:"required,numeric"`
	Discount            int        `json:"discount" validate:"required,numeric"`
	FinalAmount         int        `json:"final_amount" validate:"required,numeric"`
	CustomerId          int        `json:"customer_id" validate:"required,min=1,max=1000000000,numeric"`
	CustomerName        string     `json:"customer_name" validate:"required"`
	CustomerGender      string     `json:"customer_gender" validate:"required,oneof=male female unknown"`
	CustomerEmail       string     `json:"customer_email" validate:"omitempty,email"`
	CustomerPhone       string     `json:"customer_phone" validate:"required,phone"`
	CustomerBirthday    string     `json:"customer_birthday" validate:"omitempty,date"`
	BrandId             int        `json:"brand_id" validate:"required,min=1,max=1000000000,numeric"`
	BrandName           BrandName  `json:"brand_name" validate:"required"`
	OutletId            int        `json:"outlet_id" validate:"required,min=1,max=1000000000,numeric"`
	OutletName          OutletName `json:"outlet_name" validate:"required"`
	TerminalId          int        `json:"terminal_id" validate:"required,min=1,max=1000000000,numeric"`
	TerminalNo          string     `json:"terminal_no" validate:"required"`
	RedeemedAt          string     `json:"redeemed_at" validate:"required,datetime,uniqueDB=order_promo_codes redeemed_at"`
}

type OrderPromoCodeTransformer struct {
	Id                  int64       `json:"id"`
	OrderId             int         `json:"order_id"`
	OrderNo             string      `json:"order_no"`
	PromoCode           string      `json:"promo_code"`
	CampaignId          string      `json:"campaign_id"`
	CampaignName        string      `json:"campaign_name"`
	CampaignDescription string      `json:"campaign_description"`
	CampaignStartedAt   time.Time   `json:"campaign_started_at"`
	CampaignEndedAt     time.Time   `json:"campaign_ended_at"`
	Status              string      `json:"status"`
	Amount              int         `json:"amount"`
	Discount            int         `json:"discount"`
	FinalAmount         int         `json:"final_amount"`
	CustomerId          int         `json:"customer_id"`
	CustomerName        string      `json:"customer_name"`
	CustomerGender      string      `json:"customer_gender"`
	CustomerEmail       interface{} `json:"customer_email"`
	CustomerPhone       string      `json:"customer_phone"`
	CustomerBirthday    interface{} `json:"customer_birthday"`
	BrandId             int         `json:"brand_id"`
	BrandName           Locale      `json:"brand_name"`
	OutletId            int         `json:"outlet_id"`
	OutletName          Locale      `json:"outlet_name"`
	TerminalId          int         `json:"terminal_id"`
	TerminalNo          string      `json:"terminal_no"`
	RedeemedAt          time.Time   `json:"redeemed_at"`
	Created             *time.Time  `json:"created_at"`
	Updated             *time.Time  `json:"updated_at"`
}

func (m *OrderPromoCode) Transformer() *OrderPromoCodeTransformer {

	brandName := Locale{}

	if err := json.Unmarshal([]byte(m.BrandName), &brandName); err != nil {
		fmt.Println("error", err)
	}

	outletName := Locale{}

	if err := json.Unmarshal([]byte(m.OutletName), &outletName); err != nil {
		fmt.Println("error", err)
	}

	model := OrderPromoCodeTransformer{
		Id:                  m.Id,
		OrderId:             m.OrderId,
		OrderNo:             m.OrderNo,
		PromoCode:           m.PromoCode,
		CampaignId:          m.CampaignId,
		CampaignName:        m.CampaignName,
		CampaignDescription: m.CampaignDescription,
		CampaignStartedAt:   m.CampaignStartedAt,
		CampaignEndedAt:     m.CampaignEndedAt,
		Status:              m.Status,
		Amount:              m.Amount,
		Discount:            m.Discount,
		FinalAmount:         m.FinalAmount,
		CustomerId:          m.CustomerId,
		CustomerName:        m.CustomerName,
		CustomerGender:      m.CustomerGender,
		CustomerPhone:       m.CustomerPhone,
		BrandId:             m.BrandId,
		BrandName:           brandName,
		OutletId:            m.OutletId,
		OutletName:          outletName,
		TerminalId:          m.TerminalId,
		TerminalNo:          m.TerminalNo,
		RedeemedAt:          m.RedeemedAt,
		Created:             m.Created,
		Updated:             m.Updated,
	}

	if m.CustomerEmail != "" {
		model.CustomerEmail = m.CustomerEmail
	}

	if m.CustomerBirthday != "" {
		model.CustomerBirthday = m.CustomerBirthday
	}

	return &model
}

type InputUpdateOrderPromoCode struct {
	Status string `json:"status" validate:"required,oneof=success fail cancel waiting"`
}

func (m *OrderPromoCode) TableName() string {
	return "order_promo_codes"
}

func (m OrderPromoCode) DefaultIncludes() []string {
	return []string{}
}

func (m OrderPromoCode) Attributes() []string {

	e := reflect.ValueOf(&m).Elem()

	var attributes []string

	for i := 0; i < e.NumField(); i++ {

		field := e.Type().Field(i).Name

		isRelation := false

		if value, ok := e.Type().Field(i).Tag.Lookup("orm"); ok {

			ormValues := strings.Split(value, ";")

			for _, v := range ormValues {

				if matched, err := regexp.MatchString(`^rel\((fk|m2m)\)$`, v); err == nil {
					if matched {
						isRelation = true
					}
				}

			}

		}

		isDefaultInclude := Contains(m.DefaultIncludes(), field)

		if !isRelation || isDefaultInclude == true {
			attributes = append(attributes, field)
		}
	}

	return attributes
}

func FindOrderPromoCodeById(id int64) (m *OrderPromoCodeTransformer, error error, statusCode int) {

	o := orm.NewOrm()
	model := OrderPromoCode{Id: id}
	err := o.Read(&model)
	//num, err := o.LoadRelated(&outlet, "Brand")

	if err == orm.ErrNoRows {
		message := fmt.Sprintf("The order promo code with id %d doesn't exist", id)
		return nil, errors.New(message), 404
	} else if err == orm.ErrMissPK {
		return nil, errors.New("No primary key found."), 404
	} else {
		//fmt.Println("Brand nums: ", num)
		return model.Transformer(), nil, 200
	}
}

func CreateOrderPromoCode(input OrderPromoCode) (m *OrderPromoCodeTransformer, e error, statusCode int) {
	o := orm.NewOrm()

	if id, err := o.Insert(&input); err == nil {
		return FindOrderPromoCodeById(id)
	} else {
		fmt.Println(err)
		return nil, err, 417
	}
}

func GetAllOrderPromoCodes(query map[string]string, fields []string, orderBy string, sortedBy string,
	offset int, limit int64, includes []string) (models []*OrderPromoCodeTransformer, count int64, err error) {

	o := orm.NewOrm()
	//
	qs := o.QueryTable(new(OrderPromoCode))

	qs = where(qs, query)

	if orderBy != "" {
		sort := orderBy

		if sortedBy == "desc" {
			sort = "-" + sort
		}

		qs = qs.OrderBy(sort)
	}

	model := OrderPromoCode{}

	attributes := model.Attributes()

	if len(includes) > 0 {
		for _, field := range includes {
			attributes = append(attributes, field)
		}
	}

	var all []OrderPromoCode

	if len(model.DefaultIncludes()) > 0 {

		relations := strings.Join(model.DefaultIncludes(), ", ")

		if count, err := qs.Limit(limit, offset).RelatedSel(relations).All(&all, attributes...); err == nil {

			for _, model := range all {
				models = append(models, model.Transformer())
			}

			return models, count, nil
		}
	} else {
		if count, err := qs.Limit(limit, offset).All(&all, attributes...); err == nil {

			for _, model := range all {
				models = append(models, model.Transformer())
			}

			return models, count, nil
		}
	}

	return models, count, nil
}

func UpdateOrderPromoCode(id int64, input InputUpdateOrderPromoCode) (m *OrderPromoCodeTransformer, err error, statusCode int) {

	o := orm.NewOrm()

	model := OrderPromoCode{Id: id}

	if err = o.Read(&model); err == nil {
		var num int64
		model.Status = input.Status

		if num, err = o.Update(&model); err == nil {
			fmt.Println("Number of records updated in database:", num)
			//if num, err = o.LoadRelated(&orderPromoCode, "Brand"); err == nil {
			//	return &promotion, nil
			//}
			return model.Transformer(), nil, 200
		} else {
			fmt.Println("err:", err)
			return nil, err, 417
		}
	}

	if err == orm.ErrNoRows {
		message := fmt.Sprintf("The order promo code with id %d doesn't exist", id)
		return nil, errors.New(message), 404
	} else if err == orm.ErrMissPK {
		return nil, errors.New("No primary key found."), 404
	} else {
		return model.Transformer(), nil, 200
	}
}

//
//func DeleteOrderPromoCode(id int64) (err error) {
//	o := orm.NewOrm()
//
//	promotion := OrderPromoCode{Id: id}
//
//	if num, err := o.Delete(&promotion); err == nil {
//		fmt.Println(num)
//		if num > 0 {
//			return nil
//		} else {
//			return errors.New("No result found.")
//		}
//	}
//
//	return err
//}

func CountOrderPromoCode(query map[string]string, orderBy string, sortedBy string) (count int64, err error) {

	o := orm.NewOrm()

	qs := o.QueryTable(new(OrderPromoCode))

	qs = where(qs, query)

	if orderBy != "" {
		sort := orderBy

		if sortedBy == "desc" {
			sort = "-" + sort
		}

		qs = qs.OrderBy(sort)
	}

	if count, err := qs.Count(); err == nil {
		return count, nil
	}

	return 0, err
}

func where(qs orm.QuerySeter, query map[string]string) orm.QuerySeter {

	fmt.Println("len(query)", len(query))

	if len(query) > 0 {
		for k, v := range query {
			// rewrite dot-notation to Object__Attribute
			k = strings.Replace(k, ".", "__", -1)

			switch {
			case k == "redeemed_start_at":
				qs = qs.Filter("redeemed_at__gte", v)
			case k == "redeemed_end_at":
				qs = qs.Filter("redeemed_at__lte", v)
			case k == "order_no":
				qs = qs.Filter("order_no__icontains", v)
			case k == "customer_name":
				qs = qs.Filter("customer_name__icontains", v)
			case k == "customer_email":
				qs = qs.Filter("customer_email__icontains", v)

			default:
				qs = qs.Filter(k, v)
			}
		}
	}

	return qs
}

func sort(qs orm.QuerySeter, orderBy []string, sortedBy []string) (orm.QuerySeter, error) {

	var sortFields []string
	if len(orderBy) != 0 {
		if len(orderBy) == len(sortedBy) {
			// 1) for each sort field, there is an associated order
			for i, v := range orderBy {
				orderby := ""
				if sortedBy[i] == "desc" {
					orderby = "-" + v
				} else if sortedBy[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(orderBy) != len(sortedBy) && len(sortedBy) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range orderBy {
				orderby := ""
				if sortedBy[0] == "desc" {
					orderby = "-" + v
				} else if sortedBy[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(orderBy) != len(sortedBy) && len(sortedBy) != 1 {
			return nil, errors.New("Error: 'OrderBy', 'SortedBy' sizes mismatch or 'OrderBy' size is not 1")
		}
	} else {
		if len(sortedBy) != 0 {
			return nil, errors.New("Error: unused 'SortedBy' fields")
		}
	}

	qs = qs.OrderBy(sortFields...)

	return qs, nil
}

func init() {
	orm.RegisterModel(new(OrderPromoCode))
}
