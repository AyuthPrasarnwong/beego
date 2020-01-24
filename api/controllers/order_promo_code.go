package controllers

import (
	"api/helpers"
	"api/models"
	"api/utils/pagination"
	validation "api/validation"
	"encoding/json"
	//"errors"
	"fmt"
	//"log"
	//"strconv"

	//"reflect"
	//"regexp"
	//"strconv"
	"strings"
	"time"

	//"github.com/astaxie/beego/orm"
	"github.com/newrelic/go-agent"
	"gopkg.in/go-playground/validator.v9"

)

type OrderPromoCodeController struct {
	BaseController
}

type OrderPromoCodes []*models.OrderPromoCodeTransformer
type OrderPromoCode models.OrderPromoCode

// use a single instance , it caches struct info


// @Title GetAll
// @Description get all OrderPromoCodes
// @Success 200 {object} models.OrderPromoCode
// @router / [get]
func (this *OrderPromoCodeController) GetAll() {
	queryValues := this.Ctx.Request.URL.Query()

	var fields []string
	//var orderBy []string
	//var sortedBy []string
	var orderBy string
	var sortedBy string
	var query map[string]string = make(map[string]string)
	var limit int64 = 15
	var offset int = 0
	var includes []string
	var currentPage int = 1

	if v, err := this.GetInt64("limit"); err == nil {
		limit = v
	}

	if v, err := this.GetInt("offset"); err == nil {
		offset = v
	}

	//if v := this.GetString("orderBy"); v != "" {
	//	orderBy = strings.Split(v, ",")
	//}
	//
	//if v := this.GetString("sortedBy"); v != "" {
	//	sortedBy = strings.Split(v, ",")
	//}


	orderBy = this.GetString("orderBy")
	sortedBy = this.GetString("sortedBy")

	if v := this.GetString("include"); v != "" {
		includes = strings.Split(v, ",")
	}

	if v, err := this.GetInt("page"); err == nil {
		currentPage = v
	}

	if len(queryValues) > 0 {
		delete(queryValues, "fields")
		delete(queryValues, "orderBy")
		delete(queryValues, "sortedBy")
		delete(queryValues, "limit")
		delete(queryValues, "offset")
		delete(queryValues, "include")
		delete(queryValues, "page")

		for key, _ := range queryValues {
			query[key] = queryValues.Get(key)
		}
	}

	total, err := models.CountOrderPromoCode(query, orderBy, sortedBy)

	if err != nil {
		this.Data["json"] = err.Error()
	} else {
		paginator := pagination.SetPaginator(this.Ctx, int(limit), total)

		offset = paginator.Offset()
		totalPages := paginator.PageNums()

		s := newrelic.DatastoreSegment{
			Product: newrelic.DatastoreMySQL,
			Collection: "order_promo_codes",
			Operation: "SELECT",
		}

		txn := this.Ctx.Input.GetData("newrelic_transaction").(newrelic.Transaction)
		s.StartTime = newrelic.StartSegmentNow(txn)
		orderPromoCodes, count, err := models.GetAllOrderPromoCodes(query, fields, orderBy, sortedBy, offset, limit, includes)
		s.End()

		if err != nil {
			this.Data["json"] = err.Error()
		} else {

			var link interface{}

			data := orderPromoCodes
			perPage := int(limit)
			count := int(count)
			link = pagination.GetLink(count, currentPage, totalPages, this.Ctx.Input.Site(), paginator.PageLinkNext(), paginator.PageLinkPrev())

			if currentPage > totalPages {
				data = OrderPromoCodes{}
				count = 0
			}

			if count == 0 {
				link = []string{}
			}

			this.Data["json"] = PaginateResponse{
				data,
				pagination.Pagination{
					total,
					count,
					perPage,
					currentPage,
					totalPages,
					link,
				},
			}

		}
	}

	this.ServeJSON()
}

// @Title Get
// @Description find outlet by id
// @Param	id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.OrderPromoCode
// @Failure 403 :id is empty
// @router /:id [get]
func (this *OrderPromoCodeController) Get() {

	id, _ := this.GetInt64(":id")

	s := newrelic.DatastoreSegment{
		Product: newrelic.DatastoreMySQL,
		Collection: "order_promo_codes",
		Operation: "SELECT",
	}

	txn := this.Ctx.Input.GetData("newrelic_transaction").(newrelic.Transaction)
	s.StartTime = newrelic.StartSegmentNow(txn)
	orderPromoCode, err, statusCode := models.FindOrderPromoCodeById(id)
	s.End()

	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(statusCode)
		this.Data["json"] = ErrorResponse{
			Errors{
					Code: statusCode,
					Message: err.Error(),
				},
		}

	} else {
		this.Data["json"] = Response{
			Data: orderPromoCode,
		}
	}

	this.ServeJSON()
}

// @Title Create
// @Description create outlet
// @Param	body		body 	models.OrderPromoCode	true		"body for user content"
// @Success 200 {int} models.OrderPromoCode.Id
// @Failure 403 body is empty
// @router / [post]
func (this *OrderPromoCodeController) Post() {

	var input models.InputCreateOrderPromoCode
	var model models.OrderPromoCode

	json.Unmarshal(this.Ctx.Input.RequestBody, &input)

	err := validation.Validate.Struct(&input)
	if err != nil {

		// this check is only needed when your code could produce
		// an invalid value for validation such as interface with nil
		// value most including myself do not usually have code like this.
		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		errs := err.(validator.ValidationErrors)

		for _, e := range errs {

			this.Ctx.ResponseWriter.WriteHeader(422)

			field :=  helpers.ToSnakeCase(e.Field())

			enum := []string{"en", "th"}

			if helpers.InSlice(field, enum) == true {

				namespace := e.Namespace()

				keys := strings.Split(namespace, ".")

				len := len(keys)

				field = keys[len - 2] + "." + keys[len - 1]

				field = helpers.ToSnakeCase(field)
			}

			trans, _ := validation.Uni.GetTranslator("en")

			msg := e.Translate(trans)

			this.Data["json"] = ErrorResponse{
				Errors{422, msg, map[string][]string{field: []string{msg}}},
			}

			this.ServeJSON()
			return

			//fmt.Println(e.Translate(trans))


		}

		// from here you can create your own error messages in whatever language you wish
		return
	}

	//dateLayout := "2006-01-02"
	dateTimeLayout := "2006-01-02 15:04:05"

	campaignStartedAt, _ := time.Parse(dateTimeLayout, input.CampaignStartedAt)
	campaignEndedAt, _ := time.Parse(dateTimeLayout, input.CampaignEndedAt)
	redeemedAt, _ := time.Parse(dateTimeLayout, input.RedeemedAt)

	//customerBirthday, _ := time.Parse(dateLayout, input.CustomerBirthday)
	//
	brandName := &input.BrandName
	brandNameJson, _ := json.Marshal(brandName)
	outletName := &input.OutletName
	outletNameJson, _ := json.Marshal(outletName)

	customerGender := "unknown"

	if input.CustomerGender != "" {
		customerGender = input.CustomerGender
	}

	//fmt.Println("input.BrandName", input.BrandName.En)

	model.OrderId = input.OrderId
	model.OrderNo = input.OrderNo
	model.PromoCode = input.PromoCode
	model.CampaignId = input.CampaignId
	model.CampaignName = input.CampaignName
	model.CampaignDescription = input.CampaignDescription
	model.CampaignStartedAt = campaignStartedAt.Add(time.Hour * -7)
	model.CampaignEndedAt = campaignEndedAt.Add(time.Hour * -7)
	model.Status = input.Status
	model.Amount = input.Amount
	model.Discount = input.Discount
	model.FinalAmount = input.FinalAmount
	model.CustomerId = input.CustomerId
	model.CustomerName = input.CustomerName
	model.CustomerGender = customerGender
	model.CustomerEmail = input.CustomerEmail
	model.CustomerPhone = input.CustomerPhone
	model.CustomerBirthday = input.CustomerBirthday

	model.BrandId = input.BrandId
	model.BrandName = string(brandNameJson)
	model.OutletId = input.OutletId
	model.OutletName = string(outletNameJson)
	model.TerminalId = input.TerminalId
	model.TerminalNo = input.TerminalNo
	model.RedeemedAt = redeemedAt.Add(time.Hour * -7)

	s := newrelic.DatastoreSegment{
		Product: newrelic.DatastoreMySQL,
		Collection: "order_promo_codes",
		Operation: "INSERT",
	}

	txn := this.Ctx.Input.GetData("newrelic_transaction").(newrelic.Transaction)
	s.StartTime = newrelic.StartSegmentNow(txn)
	orderPromoCode, err, statusCode := models.CreateOrderPromoCode(model)
	s.End()

	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(statusCode)
		this.Data["json"] = ErrorResponse{
			Errors{
				Code: statusCode,
				Message: err.Error(),
			},
		}
	} else {
		this.Ctx.ResponseWriter.WriteHeader(201)
		this.Data["json"] = Response{
			Data: orderPromoCode,
		}
	}

	this.ServeJSON()
}

//@Title Update
//@Description update the outlet
//@Param	uid		path 	string	true		"The id you want to update"
//@Param	body		body 	models.OrderPromoCode	true		"body for user content"
//@Success 200 {object} models.OrderPromoCode
//@Failure 403 :id is not int
//@router /:id [put]
func (this *OrderPromoCodeController) Put() {
	id, _ := this.GetInt64(":id")

	var input models.InputUpdateOrderPromoCode

	json.Unmarshal(this.Ctx.Input.RequestBody, &input)

	err := validation.Validate.Struct(&input)
	if err != nil {

		if _, ok := err.(*validator.InvalidValidationError); ok {
			fmt.Println(err)
			return
		}

		errs := err.(validator.ValidationErrors)

		for _, e := range errs {

			this.Ctx.ResponseWriter.WriteHeader(422)

			field :=  helpers.ToSnakeCase(e.Field())

			enum := []string{"en", "th"}

			if helpers.InSlice(field, enum) == true {

				namespace := e.Namespace()

				keys := strings.Split(namespace, ".")

				len := len(keys)

				field = keys[len - 2] + "." + keys[len - 1]

				field = helpers.ToSnakeCase(field)
			}

			trans, _ := validation.Uni.GetTranslator("en")

			msg := e.Translate(trans)

			this.Data["json"] = ErrorResponse{
				Errors{422, msg, map[string][]string{field: []string{msg}}},
			}
			this.ServeJSON()
			return

		}

		return
	}


	s := newrelic.DatastoreSegment{
		Product: newrelic.DatastoreMySQL,
		Collection: "order_promo_codes",
		Operation: "UPDATE",
	}

	txn := this.Ctx.Input.GetData("newrelic_transaction").(newrelic.Transaction)
	s.StartTime = newrelic.StartSegmentNow(txn)
	orderPromoCode, err, statusCode := models.UpdateOrderPromoCode(id, input)
	s.End()

	if err != nil {
		this.Ctx.ResponseWriter.WriteHeader(statusCode)
		this.Data["json"] = ErrorResponse{
			Errors{
				Code: statusCode,
				Message: err.Error(),
			},
		}
	} else {
		this.Data["json"] = Response{
			Data: orderPromoCode,
		}
	}
	this.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 204 body is empty
// @Failure 403 outletId is empty
// @router /:id [delete]
//func (this *OrderPromoCodeController) Delete() {
//
//	id, _ := this.GetInt64(":id")
//
//	s := newrelic.DatastoreSegment{
//		Product: newrelic.DatastoreMySQL,
//		Collection: "outlets",
//		Operation: "DELETE",
//	}
//
//	txn := this.Ctx.Input.GetData("newrelic_transaction").(newrelic.Transaction)
//	s.StartTime = newrelic.StartSegmentNow(txn)
//	err := models.DeletePromotion(id)
//	s.End()
//
//	if err != nil {
//		this.Data["json"] = err.Error()
//	} else {
//		this.Ctx.ResponseWriter.WriteHeader(204)
//	}
//
//	this.ServeJSON()
//}
