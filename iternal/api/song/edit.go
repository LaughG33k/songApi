package song

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/LaughG33k/songApi/iternal/model"
	"github.com/LaughG33k/songApi/pkg"
	"github.com/gin-gonic/gin"
)

func (s *Song) Edit(c *gin.Context) {

	filter := model.SongHead{
		Song:  c.Param("song"),
		Group: c.Param("group"),
	}
	editFields := model.Song{}

	if err := json.NewDecoder(c.Request.Body).Decode(&editFields); err != nil {
		c.String(http.StatusBadRequest, model.BadRequest.Error())
		return
	}

	var timeout context.Context
	var canc context.CancelFunc

	if deadline, ok := c.Deadline(); ok {
		timeout, canc = context.WithDeadline(context.TODO(), deadline)
	} else {
		timeout, canc = context.WithTimeout(context.TODO(), s.operationTimeout)
	}

	defer canc()

	if err := s.songService.Edit(timeout, filter, editFields); err != nil {
		if err.Error() == model.EmptyGroup.Error() || err.Error() == model.EmptySong.Error() {
			c.String(http.StatusBadRequest, err.Error())
			return
		}

		pkg.Log.Infof("ошибка при редактирование песни: %s", err)
		if err.Error() == model.SongExists.Error() {
			c.String(http.StatusInternalServerError, model.SongExists.Error())
			return
		}
		c.String(http.StatusInternalServerError, model.IternalError.Error())
		return
	}

}
