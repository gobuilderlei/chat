package dbutil

type MyPagination struct {
	PageNumber int32
	ShowNumber int32
}

func (p *MyPagination) GetPageNumber() int32 {
	return p.PageNumber
}

func (p *MyPagination) GetShowNumber() int32 {
	return p.ShowNumber
}
