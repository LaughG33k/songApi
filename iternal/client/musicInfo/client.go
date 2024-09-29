package musicinfo

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/LaughG33k/songApi/iternal/model"
	"github.com/goccy/go-json"
)

var musicInfoUrl string = os.Getenv("musicInfoUrl")

func GetDetailInfo(ctx context.Context, song, group string) (model.DetailInfo, error) {

	detailInfo := model.DetailInfo{}

	r, err := http.NewRequestWithContext(ctx, http.MethodGet, fmt.Sprintf("%s/?song=%s&group=%s", musicInfoUrl, song, group), nil)

	if err != nil {
		return detailInfo, nil
	}

	client := &http.Client{}

	resp, err := client.Do(r)

	if err != nil {
		return detailInfo, err
	}

	defer resp.Body.Close()

	if resp.StatusCode == 400 {
		return detailInfo, model.BadRequest
	} else if resp.StatusCode == 500 {
		return detailInfo, model.IternalError
	}

	bytes, err := io.ReadAll(resp.Body)

	if err != nil {
		return detailInfo, err
	}

	if err := json.Unmarshal(bytes, &detailInfo); err != nil {
		return detailInfo, err
	}

	return detailInfo, nil

}
