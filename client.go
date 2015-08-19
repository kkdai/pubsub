package pubsub

type Client struct {
	channelIndex int
	channel      chan interface{}
}

func (c *Client) Get() interface{} {
	return <-c.channel
}
