package doc_doc_webapi

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/nemo260/doc-doc-webapi/internal/db_service"
)

func (this *implDocsAPI) AddDocument(ctx *gin.Context) {
	value, exists := ctx.Get("db_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}
	db, ok := value.(db_service.DbService[Document])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	document := Document{}
	err := ctx.BindJSON(&document)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		return
	}

	if document.Id == "" {
		document.Id = uuid.New().String()
	}

	err = db.CreateDocument(ctx, document.Id, &document)

	switch err {
	case nil:
		ctx.JSON(
			http.StatusCreated,
			document,
		)
	case db_service.ErrConflict:
		ctx.JSON(
			http.StatusConflict,
			gin.H{
				"status":  "Conflict",
				"message": "Document already exists",
				"error":   err.Error(),
			},
		)
	default:
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to create document in database",
				"error":   err.Error(),
			},
		)
	}

}

func (this *implDocsAPI) DeleteDocument(ctx *gin.Context) {
	value, exists := ctx.Get("db_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}
	db, ok := value.(db_service.DbService[Document])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	documentId := ctx.Param("id")
	err := db.DeleteDocument(ctx, documentId)
	if err != nil {
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to delete document from database",
				"error":   err.Error(),
			})
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{
			"status":  "OK",
			"message": "Document deleted",
		},
	)
}

func (this *implDocsAPI) GetDocumentById(ctx *gin.Context) {
	value, exists := ctx.Get("db_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}
	db, ok := value.(db_service.DbService[Document])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	documentId := ctx.Param("id")
	document, err := db.FindDocument(ctx, documentId)
	if err != nil {
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to retrieve document from database",
				"error":   err.Error(),
			})
		return
	}

	ctx.JSON(
		http.StatusOK,
		document,
	)
}

func (this *implDocsAPI) GetDocuments(ctx *gin.Context) {
	value, exists := ctx.Get("db_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}
	db, ok := value.(db_service.DbService[Document])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	documents, err := db.FindAllDocuments(ctx)
	if err != nil {
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to retrieve documents from database",
				"error":   err.Error(),
			})
		return
	}

	ctx.JSON(
		http.StatusOK,
		documents,
	)
}

func (this *implDocsAPI) UpdateDocument(ctx *gin.Context) {
	value, exists := ctx.Get("db_service")
	if !exists {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db not found",
				"error":   "db not found",
			})
		return
	}
	db, ok := value.(db_service.DbService[Document])
	if !ok {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{
				"status":  "Internal Server Error",
				"message": "db context is not of required type",
				"error":   "cannot cast db context to db_service.DbService",
			})
		return
	}

	document := Document{}
	err := ctx.BindJSON(&document)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			gin.H{
				"status":  "Bad Request",
				"message": "Invalid request body",
				"error":   err.Error(),
			})
		return
	}

	err = db.UpdateDocument(ctx, document.Id, &document)
	if err != nil {
		ctx.JSON(
			http.StatusBadGateway,
			gin.H{
				"status":  "Bad Gateway",
				"message": "Failed to update document in database",
				"error":   err.Error(),
			})
		return
	}

	ctx.JSON(
		http.StatusOK,
		document,
	)
}
