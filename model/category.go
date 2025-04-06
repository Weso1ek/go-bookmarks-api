package model

type Category struct {
	Id        int32
	Title     string
	Body      *string
	MetaTitle string
	MetaDesc  string
	Logo      *string
	Path      *string
	SLug      *string
}
