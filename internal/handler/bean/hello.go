package bean

import gginutil "gitlab.ftsview.com/fotoable-go/ggin-util"

type HelloReq struct {
	Header gginutil.HeaderToC `json:"-"`

	ID int64 `json:"id"`
}
