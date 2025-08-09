package route

import (
	"mime/multipart"
	"net/http"
	"path/filepath"
	"slices"

	"github.com/caoyong2619/elotus/internal/database"
	"github.com/caoyong2619/elotus/internal/services"
	"github.com/gabriel-vasile/mimetype"
	"github.com/gin-gonic/gin"
)

func Upload() gin.HandlerFunc {
	return func(c *gin.Context) {
		file, err := c.FormFile(`data`)
		if err != nil {
			c.JSON(http.StatusBadRequest, Error(CodeError, err.Error()))
			return
		}

		maxFileSize := int64(8 << 20) // 8MB
		if file.Size > maxFileSize {
			c.JSON(http.StatusBadRequest, Error(CodeError, `file size too large`))
			return
		}

		fp, err := file.Open()
		if err != nil {
			c.JSON(http.StatusBadRequest, Error(CodeError, err.Error()))
		}

		defer fp.Close()

		m, err := mimetype.DetectReader(fp)
		if err != nil {
			c.JSON(http.StatusBadRequest, Error(CodeError, err.Error()))
			return
		}

		allowedTypes := []string{`image/jpeg`, `image/png`}
		if !slices.Contains(allowedTypes, m.String()) {
			c.JSON(http.StatusBadRequest, Error(CodeError, `file type not allowed`))
			return
		}

		// save file
		savePath := filepath.Join(`/tmp`, file.Filename)
		if err := c.SaveUploadedFile(file, savePath); err != nil {
			c.JSON(http.StatusBadRequest, Error(CodeError, err.Error()))
			return
		}

		// record upload
		err = recordUpload(c, file, m, savePath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Error(CodeError, err.Error()))
		}

		c.JSON(http.StatusOK, Success(gin.H{
			`url`: savePath,
		}))
	}
}

func recordUpload(c *gin.Context, file *multipart.FileHeader, mime *mimetype.MIME, path string) error {
	upload := &database.Upload{
		MimeType: mime.String(),
		Size:     file.Size,
		Filepath: path,
	}
	token := c.MustGet("token").(*services.ElotusClaims)
	upload.UserId = token.ID

	_, err := database.Engine.Insert(upload)
	return err
}
