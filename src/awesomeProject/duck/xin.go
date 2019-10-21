package duck

type Xin struct {
	Contents string
}

func (x Xin) Get(url string) string {
	return x.Contents
}

func (x Xin) Post(url string, form map[string]string) string {
	x.Contents = form["name"]
	return "ok"
}
