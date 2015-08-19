package pubsub

type chanMapStringList map[chan interface{}][]string
type stringMapChanList map[string][]chan interface{}

// Pubsub struct: Only content a userIndex and accessDB which content a chan map
type Pubsub struct {
	//Capacity for each chan buffer
	capacity int

	//map to store "chan -> Topic List" for find subscription
	clientMapTopics chanMapStringList
	//map to store "topic -> chan List" for publish
	topicMapClients stringMapChanList
}

//Sub: Subscribe channels, the channels could be a list of channels name
//     The channel name could be any, without define in server
func (p *Pubsub) Subscribe(topics ...string) chan interface{} {
	//init new chan using capacity as channel buffer
	workChan := make(chan interface{}, p.capacity)
	p.updateTopicMapClient(workChan, topics)
	return workChan
}

func (p *Pubsub) updateTopicMapClient(clientChan chan interface{}, topics []string) {
	var updateChanList []chan interface{}
	var ok bool
	for _, topic := range topics {
		if updateChanList, ok = p.topicMapClients[topic]; ok {
			updateChanList = append(updateChanList, clientChan)
		} else {
			updateChanList = append(updateChanList, clientChan)
		}
		p.topicMapClients[topic] = updateChanList
	}
	p.clientMapTopics[clientChan] = topics
}

//AddSubscription:  Add a new topic subscribe to specific client channel.
func (p *Pubsub) AddSubscription(clientChan chan interface{}, topics ...string) {
	p.updateTopicMapClient(clientChan, topics)
}

//Publish: Publish a content to a list of channels
//         The content could be any type.
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
