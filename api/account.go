package api

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	db "github.com/nadamalash/bank-backend/db/sqlc"
)

// createAccountRequest is a struct to hold the request body for creating an account
type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD EUR"`
}

// createAccount is a handler to create a new account
// extends the server struct
func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest

	// bind the request body to the req struct
	if err := ctx.ShouldBindJSON(&req); err != nil {
		// return a bad request error response, which is a helper function
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// create a new account in the database
	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}
	account, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		// return a server error response, which is a helper function
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// return a success response
	// the response body will be like this:
	//	{
	//		"id": 1,
	//		"owner": "nadamalash",
	//		"balance": 0,
	//		"currency": "USD"
	//	}
	ctx.JSON(http.StatusOK, account)

}

// getAccountRequest is a struct to hold the request body for getting an account
type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"`
}

func (server *Server) getAccount(ctx *gin.Context) {

	var req getAccountRequest

	// bind the request body to the req struct
	if err := ctx.ShouldBindUri(&req); err != nil {
		// return a bad request error response, which is a helper function
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// get the account from the database
	account, err := server.store.GetAccount(ctx, req.ID)
	if err != nil {
		if err == sql.ErrNoRows {
			// return a not found error response, which is a helper function
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		// return a server error response, which is a helper function
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// return a success response
	// the response body will be like this:
	//	{
	//		"id": 1,
	//		"owner": "nadamalash",
	//		"balance": 0,
	//		"currency": "USD"
	//	}
	ctx.JSON(http.StatusOK, account)
}

type listAccountRequest struct {
	PageID   int32 `form:"page_id" binding:"required,min=1"`
	PageSize int32 `form:"page_size" binding:"required,min=5,max=10"`
}

func (server *Server) listAccounts(ctx *gin.Context) {
	var req listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	// Log the request parameters for debugging
	log.Printf("Received request with parameters: %+v", req)

	arg := db.ListAccountsParams{
		Limit:  req.PageSize,
		Offset: (req.PageID - 1) * req.PageSize,
	}
	// Log the database query parameters for debugging
	log.Printf("Querying database with parameters: %+v", arg)

	accounts, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// Log the retrieved accounts for debugging
	log.Printf("Retrieved accounts from the database: %+v", accounts)

	ctx.JSON(http.StatusOK, accounts)
}
