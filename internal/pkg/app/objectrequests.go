package app

import (
	"github.com/gin-gonic/gin"
	"hac/internal/app/ds"
	"net/http"
	"strconv"
)

func (a *Application) GetObjectsByFloor(gCtx *gin.Context) {
	floor := gCtx.Param("floor")
	floor_int, err := strconv.Atoi(floor)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant convert id to int"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	objects, err := a.repo.GetObjectsByFloor(floor_int)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant get all rows"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	gCtx.IndentedJSON(http.StatusOK, objects)

}

func (a *Application) GetObjectById(gCtx *gin.Context) {
	id_object := gCtx.Param("id")
	id_object_int, err := strconv.Atoi(id_object)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant convert id to int"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	product, err := a.repo.GetObjectByID(uint(id_object_int))
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant get product by id"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	gCtx.IndentedJSON(http.StatusOK, &product)
}

func (a *Application) AddFavorite(gCtx *gin.Context) {
	var params ds.Favorites
	err := gCtx.BindJSON(&params)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant parse json"}
		gCtx.IndentedJSON(http.StatusRequestedRangeNotSatisfiable, answer)
		return
	}

	err = a.repo.CreateFavorite(&params)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant create product row"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	gCtx.IndentedJSON(http.StatusOK, params)
}

func (a *Application) GetFavorite(gCtx *gin.Context) {
	id_user := gCtx.Param("id_user")
	id_user_int, err := strconv.Atoi(id_user)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant convert id to int"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	products, err := a.repo.GetFavoriteByID(uint(id_user_int))
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant get product by id"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	gCtx.IndentedJSON(http.StatusOK, products)
}

func (a *Application) DeleteFavorite(gCtx *gin.Context) {
	id_favorite := gCtx.Param("id")
	id_favorite_int, err := strconv.Atoi(id_favorite)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "id must be integer"}
		gCtx.IndentedJSON(http.StatusRequestedRangeNotSatisfiable, answer)
		return
	}
	err = a.repo.DeleteFavorite(uint(id_favorite_int))
	if err.Error() == "record not found" {
		answer := AnswerJSON{Status: "error", Description: "id not found"}
		gCtx.IndentedJSON(http.StatusRequestedRangeNotSatisfiable, answer)
		return
	}
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant delete row"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	answer := AnswerJSON{Status: "successful", Description: "row was deleted"}
	gCtx.IndentedJSON(http.StatusOK, answer)
}

func (a *Application) AddFeedback(gCtx *gin.Context) {
	var params ds.FeedBack
	err := gCtx.BindJSON(&params)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant parse json"}
		gCtx.IndentedJSON(http.StatusRequestedRangeNotSatisfiable, answer)
		return
	}

	err = a.repo.CreateFeedback(&params)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant create product row"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	gCtx.IndentedJSON(http.StatusOK, params)
}

func (a *Application) GetFeedbackByID(gCtx *gin.Context) {

}

func (a *Application) AddUser(gCtx *gin.Context) {
	var params ds.Users
	err := gCtx.BindJSON(&params)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant parse json"}
		gCtx.IndentedJSON(http.StatusRequestedRangeNotSatisfiable, answer)
		return
	}

	err = a.repo.CreateUser(&params)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant create product row"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	gCtx.IndentedJSON(http.StatusOK, params)
}

func (a *Application) AddObject(gCtx *gin.Context) {
	var params ds.Objests
	err := gCtx.BindJSON(&params)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant parse json"}
		gCtx.IndentedJSON(http.StatusRequestedRangeNotSatisfiable, answer)
		return
	}

	err = a.repo.CreateUser(&params)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant create product row"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	gCtx.IndentedJSON(http.StatusOK, params)
}

// ChangePrice godoc
// @Summary      Change price of product by id
// @Description  Change price of product by id. Price can't be 0
// @Tags         Tests
// @Produce      json
// @Success      200  {object}  ds.Goods
// @Router       /goods/{id} [put]
/*func (a *Application) ChangePrice(gCtx *gin.Context) {
	var params ds.Goods
	err := gCtx.BindJSON(&params)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant parse json, "}
		gCtx.IndentedJSON(http.StatusRequestedRangeNotSatisfiable, answer)
		return
	}

	id_product := gCtx.Param("id_good")
	id_product_int, err := strconv.Atoi(id_product)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "id must be integer"}
		gCtx.IndentedJSON(http.StatusRequestedRangeNotSatisfiable, answer)
		return
	}

	if params.Price == 0 {
		answer := AnswerJSON{Status: "error", Description: "product cant cost 0"}
		gCtx.IndentedJSON(http.StatusRequestedRangeNotSatisfiable, answer)
		return
	}

	err = a.repo.ChangeProduct(uint(id_product_int), params.Price)
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant change price"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	product, err := a.repo.GetProductByID(uint(id_product_int))
	if err != nil {
		answer := AnswerJSON{Status: "error", Description: "cant get product by id"}
		gCtx.IndentedJSON(http.StatusInternalServerError, answer)
		return
	}
	gCtx.IndentedJSON(http.StatusOK, &product)
}
*/
