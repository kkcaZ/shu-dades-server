package product

import (
	"encoding/json"
	"github.com/kkcaz/shu-dades-server/internal/domain"
	"github.com/kkcaz/shu-dades-server/internal/router"
	"github.com/kkcaz/shu-dades-server/pkg/models"
)

type ProductHandler struct {
	ProductUseCase domain.ProductUseCase
	Auth           domain.AuthUseCase
}

func NewProductHandler(router *router.RouterUseCase, uc domain.ProductUseCase, auth domain.AuthUseCase) {
	handler := ProductHandler{
		ProductUseCase: uc,
		Auth:           auth,
	}

	router.AddRoute("/product", models.GET, handler.Get)
	router.AddRoute("/product/all", models.GET, handler.GetAll)
	router.AddRoute("/product/search", models.GET, handler.Search)
	router.AddRoute("/product", models.POST, handler.Create)
	router.AddRoute("/product", models.PUT, handler.Update)
	router.AddRoute("/product", models.DELETE, handler.Delete)
	router.AddRoute("/product/subscribe", models.POST, handler.Subscribe)
	router.AddRoute("/product/unsubscribe", models.POST, handler.Unsubscribe)
	router.AddRoute("/product/subscriptions", models.GET, handler.GetProductSubscriptions)
}

func (p ProductHandler) Get(ctx *router.RouterContext) {
	var request models.RequestById
	err := json.Unmarshal([]byte(ctx.Body), &request)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	product, err := p.ProductUseCase.Get(request.Id)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	if product == nil {
		ctx.JSON(404, models.NewErrorResponse(404, "Product not found"))
		return
	}

	ctx.JSON(200, models.ProductResponse{
		StatusCode: 200,
		Product:    product,
	})
}

func (p ProductHandler) GetAll(ctx *router.RouterContext) {
	products, err := p.ProductUseCase.GetAll()
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	ctx.JSON(200, models.ProductListResponse{
		StatusCode: 200,
		Products:   products,
	})
}

func (p ProductHandler) Search(ctx *router.RouterContext) {
	var searchRequest models.SearchRequest
	err := json.Unmarshal([]byte(ctx.Body), &searchRequest)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	products, err := p.ProductUseCase.Search(searchRequest.PageNumber, searchRequest.PageSize, searchRequest.SortBy, searchRequest.Order)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	if products == nil {
		ctx.JSON(404, models.NewErrorResponse(404, "Products not found"))
		return
	}

	ctx.JSON(200, models.ProductListResponse{
		StatusCode: 200,
		Products:   products,
	})
}

func (p ProductHandler) Create(ctx *router.RouterContext) {
	var createProductRequest models.CreateProductRequest
	err := json.Unmarshal([]byte(ctx.Body), &createProductRequest)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	product := models.Product{
		Name:     createProductRequest.Name,
		Quantity: createProductRequest.Quantity,
	}

	err = p.ProductUseCase.Create(product)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	ctx.JSON(200, models.ProductResponse{
		StatusCode: 200,
		Product:    &product,
	})
}

func (p ProductHandler) Update(ctx *router.RouterContext) {
	var updateProductRequest models.Product
	err := json.Unmarshal([]byte(ctx.Body), &updateProductRequest)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	err = p.ProductUseCase.Update(&updateProductRequest)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	ctx.JSON(200, models.NewSuccessResponse(200, "Product updated successfully"))
}

func (p ProductHandler) Delete(ctx *router.RouterContext) {
	var request models.RequestById
	err := json.Unmarshal([]byte(ctx.Body), &request)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	err = p.ProductUseCase.Delete(request.Id)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	ctx.JSON(200, models.NewSuccessResponse(200, "Product deleted successfully"))
}

func (p ProductHandler) Subscribe(ctx *router.RouterContext) {
	var request models.ProductSubscriptionRequest
	err := json.Unmarshal([]byte(ctx.Body), &request)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	token := ctx.GetAuthToken()
	userClaim, err := p.Auth.GetUser(*token)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	err = p.ProductUseCase.Subscribe(request.ProductId, request.SubType, userClaim.UserId)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	ctx.JSON(200, models.NewSuccessResponse(200, "Subscribed to product"))
}

func (p ProductHandler) Unsubscribe(ctx *router.RouterContext) {
	var request models.ProductSubscriptionRequest
	err := json.Unmarshal([]byte(ctx.Body), &request)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	token := ctx.GetAuthToken()
	userClaim, err := p.Auth.GetUser(*token)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	err = p.ProductUseCase.Unsubscribe(request.ProductId, request.SubType, userClaim.UserId)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	ctx.JSON(200, models.NewSuccessResponse(200, "Unsubscribed from product"))
}

func (p ProductHandler) GetProductSubscriptions(ctx *router.RouterContext) {
	token := ctx.GetAuthToken()
	userClaim, err := p.Auth.GetUser(*token)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	subscriptions, err := p.ProductUseCase.GetProductSubscriptions(userClaim.UserId)
	if err != nil {
		ctx.JSON(500, models.NewInternalServerError())
		return
	}

	ctx.JSON(200, models.ProductSubscriptionListResponse{
		StatusCode:    200,
		Subscriptions: subscriptions,
	})
}
