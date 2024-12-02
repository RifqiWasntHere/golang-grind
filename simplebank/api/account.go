package api

import (
	"database/sql"
	"net/http"
	db "simplebank/db/sqlc"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgconn"
)

type createAccountRequest struct {
	Owner    string `json:"owner" binding:"required"`
	Currency string `json:"currency" binding:"required,oneof=USD IDR"`
	// Balance  int64  `json:"balance"` // balance is disabled because in default, the only allowed balance when creating a new account is zero
}

func (server *Server) createAccount(ctx *gin.Context) {
	var req createAccountRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err)) // TODO: errorResponse()
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}

	payload, err := server.store.CreateAccount(ctx, arg)
	if err != nil {
		if pqErr, ok := err.(*pgconn.PgError); ok {
			switch pqErr.Code {
			case "23503", "23505":
				ctx.JSON(http.StatusForbidden, errorResponse(err))
				return
			}
		}

		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, payload)
}

type getAccountRequest struct {
	ID int64 `uri:"id" binding:"required,min=1"` // this one uses `uri` instead of `json` because wellwellwell it's a uri parameter
}

func (server *Server) getAccount(ctx *gin.Context) {
	var req getAccountRequest
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload, err := server.store.GetAccount(ctx, req.ID) // again, its req.id cuz we're getting the data from uri parameter
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, payload)
}

type listAccountRequest struct {
	Page int32 `form:"page" binding:"required,min=1"` // we're using form because this time, the data is taken from query parameter.
}

func (server *Server) listAccounts(ctx *gin.Context) {
	var req listAccountRequest
	if err := ctx.ShouldBindQuery(&req); err != nil { // unidiomatically, yes. `form` requires ShouldBindQuery.
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	arg := db.ListAccountsParams{
		Limit:  5,
		Offset: (req.Page - 1) * 5,
	}

	payload, err := server.store.ListAccounts(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return

	}

	ctx.JSON(http.StatusOK, payload)
}
