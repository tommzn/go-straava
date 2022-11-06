package strava

// NewPagination returns a pagination struct.
// If page and per_page are zero it will return nil.
func NewPagination(page, per_page int) *Pagination {

	if page == 0 && per_page == 0 {
		return nil
	}

	pagination := &Pagination{}
	if page > 0 {
		pagination.Page = &page
	}
	if per_page > 0 {
		pagination.PerPage = &per_page
	}
	return pagination
}

// NextPage, helper to move to next page.
func (pagination *Pagination) NextPage() {
	if pagination.Page != nil {
		*pagination.Page++
	}
}
