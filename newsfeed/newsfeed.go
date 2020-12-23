package newsfeed

type Feed struct {
	Title string `json:"title"`
	Post  string `json:"post"`
}

type Repo struct {
	feeds []Feed
}

func New() Repo {
	return Repo{}
}

func (r *Repo) Add(f Feed) {
	r.feeds = append(r.feeds, f)
}

func (r *Repo) GetAll() []Feed {
	return r.feeds
}
