package song

import (
	"context"
	"net/http"
	"strconv"

	"github.com/LaughG33k/songApi/iternal/model"
	"github.com/LaughG33k/songApi/pkg"
	"github.com/gin-gonic/gin"
)

func (s *Song) GetText(c *gin.Context) {

	limit, err := strconv.Atoi(c.Query("limit"))

	if err != nil {
		c.String(http.StatusBadRequest, "limit must be integer")
		return
	}

	offset, err := strconv.Atoi(c.Query("offset"))

	if err != nil {
		c.String(http.StatusBadRequest, "offset must be integer")
		return
	}

	filter := model.SongHead{
		Song:  c.Query("song"),
		Group: c.Query("group"),
	}

	var timeout context.Context
	var canc context.CancelFunc

	if deadline, ok := c.Deadline(); ok {
		timeout, canc = context.WithDeadline(context.TODO(), deadline)
	} else {
		timeout, canc = context.WithTimeout(context.TODO(), s.operationTimeout)
	}

	defer canc()

	couplets, err := s.songService.GetText(timeout, filter, limit, offset)

	if err != nil {
		if err.Error() == model.EmptyGroup.Error() || err.Error() == model.EmptySong.Error() {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		pkg.Log.Infof("произошла ошибка при получение текста: %s", err)
		c.String(http.StatusInternalServerError, model.IternalError.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"couplets": couplets,
	})

}
