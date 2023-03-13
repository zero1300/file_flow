package calc

import "file_flow/models"

func Offset(p *models.Paginate) {
	p.Page = (p.Page - 1) * p.PageSize
}
