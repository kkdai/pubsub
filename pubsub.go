package pubsub

type chanMapStringList map[chan interface{}][]string
type stringMapChanList map[string][]chan interface{}

// Pubsub struct: Only content a userIndex and accessDB which content a chan map
type Pubsub struct {
	//Capacity for each chan
	capacity int

	//DB Level
	clientMapTopics chanMapStringList
	topicMapClients stringMapChanList
}

//Sub: Subscribe channels, the channels could be a list of channels name
//     The channel name could be any, without define in server
func (p *Pubsub) Subscribe(topics ...string) chan interface{} {
	//init new chan using capacity as channel buffer
	workChan := make(chan interface{}, p.capacity)

	var topicList []string
	for _, topic := range topics {
		if chanList, ok := p.topicMapClients[topic]; ok {
			chanList = append(chanList, workChan)
		} else {
			var newChanList []chan interface{}
			newChanList = append(newChanList, workChan)
			p.topicMapClients[topic] = newChanList
		}
		topicList = append(topicList, topic)
	}

	//Add in  server DB
	p.clientMapTopics[workChan] = topicList
	return workChan
}

//Pub: Publish a content to a list of channels
//     The content could be any type.
func (p *Pubsub) Publish(content interface{}, topics ...string) {
	for _, topic := range topics {
		if chanList, ok := p.topicMapClients[topic]; ok {
			//Someone has subscribed this topic
			for _, channel := range chanList {
				channel <- content
			}
		}
	}
}

// Create a pubsub with expect init size, but the size could be extend.
func NewPubsub(initChanCapacity int) *Pubsub {
	initClientMapTopics := make(chanMapStringList)
	initTopicMapClients := make(stringMapChanList)

	server := Pubsub{clientMapTopics: initClientMapTopics, topicMapClients: initTopicMapClients}
	server.capacity = initChanCapacity
	return &server
}
