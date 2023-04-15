package downloader

type Output struct {
	channel chan string
}

func (o *Output) Println(msg string) {
	o.channel <- msg
}
