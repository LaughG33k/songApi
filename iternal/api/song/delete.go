package song

import (
	"context"
	"net/http"

	"github.com/LaughG33k/songApi/iternal/model"
	"github.com/LaughG33k/songApi/pkg"
	"github.com/gin-gonic/gin"
)

func (s *Song) Delete(c *gin.Context) {

	head := model.SongHead{
		Song:  c.Param("song"),
		Group: c.Param("group"),
	}

	var timeout context.Context
	var canc context.CancelFunc

	if deadline, ok := c.Deadline(); ok {
		timeout, canc = context.WithDeadline(context.TODO(), deadline)
	} else {
		timeout, canc = context.WithTimeout(context.TODO(), s.operationTimeout)
	}

	defer canc()

	if err := s.songService.Delete(timeout, head); err != nil {

		if err.Error() == model.EmptyGroup.Error() || err.Error() == model.EmptySong.Error() {
			c.String(http.StatusBadRequest, err.Error())
			return
		}
		pkg.Log.Infof("ошибка при удаление песни: %s", err)
		c.String(http.StatusInternalServerError, model.IternalError.Error())
		return
	}

}
