package balance

type Balancer interface {
	NextPeer(nodes interface{}) (error, interface{})
}

func NewBalancer(strategy string) Balancer {
	switch strategy {
	case "round_robin":
		return &RoundRobin{}
	default:
		return &RoundRobin{}
	}
}
