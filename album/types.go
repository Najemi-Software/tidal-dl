package album

import (
	"github.com/najemi-software/tidal-dl/v4/common"
)

type Album struct {
	common.BaseAlbum
	Artist common.ItemArtist `json:"artist"`
}
