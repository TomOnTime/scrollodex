package reslist

type Databaser interface {
	CategoryStore(id int, data dex.Category) error
	CategoryList() ([]dex.Category, error)
}
