package pagination

import (
	"github.com/astaxie/beego/context"
)

// SetPaginator Instantiates a Paginator and assigns it to context.Input.Data("paginator").
func SetPaginator(context *context.Context, per int, nums int64) (paginator *Paginator) {
	paginator = NewPaginator(context.Request, per, nums)
	context.Input.SetData("paginator", &paginator)
	return
}
