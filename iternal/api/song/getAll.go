package song

import (
	"context"
	"net/http"
	"strconv"

	"github.com/LaughG33k/songApi/iternal/model"
	"github.com/LaughG33k/songApi/pkg"
	"github.com/gin-gonic/gin"
)

func (s *Song) GetAll(c *gin.Context) {

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

	filter := model.Song{
		SongHead: model.SongHead{
			Song:  c.Query("song"),
			Group: c.Query("group"),
		},
		DetailInfo: model.DetailInfo{
			Text:        c.Query("text"),
			Link:        c.Query("link"),
			RealeseDate: c.Query("realese"),
		},
	}

	var timeout context.Context
	var canc context.CancelFunc

	if deadline, ok := c.Deadline(); ok {
		timeout, canc = context.WithDeadline(context.TODO(), deadline)
	} else {
		timeout, canc = context.WithTimeout(context.TODO(), s.operationTimeout)
	}

	defer canc()

	songs, err := s.songService.GetAll(timeout, filter, limit, offset)

	if err != nil {
		pkg.Log.Infof("произошла ошибка при получение списка песен: %s", err)
		c.String(http.StatusInternalServerError, model.IternalError.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]any{
		"songs": songs,
	})

}
