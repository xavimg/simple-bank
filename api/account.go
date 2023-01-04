package api

import (
	"database/sql"
	"net/http"
	"strconv"

	db "simple-bank/db/sqlc"

	"github.com/gin-gonic/gin"
)

func (s *Server) createAccount(ctx *gin.Context) {
	var req *CreateAccountRequest
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	arg := db.CreateAccountParams{
		Owner:    req.Owner,
		Currency: req.Currency,
		Balance:  0,
	}
	account, err := s.repo.CreateAccount(ctx, arg)
	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, err.Error())
		}
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (s *Server) getAccount(ctx *gin.Context) {
	idStr := ctx.Param("id")
	idInt, err := strconv.Atoi(idStr)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	account, err := s.repo.GetAccount(ctx, int64(idInt))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, account)
}

func (s *Server) listAccounts(ctx *gin.Context) {
	var req PaginationAccountsRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, err.Error())
		return
	}

	arg := db.ListAccountsParams{
		Offset: (int32(req.PageID) - 1) * req.PageSize,
		Limit:  int32(req.PageSize),
	}

	accounts, err := s.repo.ListAccounts(ctx, arg)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	ctx.JSON(http.StatusOK, accounts)
}
