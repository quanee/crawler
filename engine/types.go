package engine

type Callback func(contents []byte, url string) Result

type Request struct {
	URL      string
	Callback Callback
}

type Result struct {
	Requests []Request
	Items    []Item
}

type Item struct {
	URL     string
	Type    string
	Id      string
	Payload interface{}
}
