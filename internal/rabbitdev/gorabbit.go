package rabbitdev

import "github.com/streadway/amqp"

type RabbitMqDev struct {
	user       string
	password   string
	host       string
	port       string
	Connection *amqp.Connection
}

func NewRabbitMqDevConn(options ...func(*RabbitMqDev)) *RabbitMqDev {
	rbmd := &RabbitMqDev{}
	for _, o := range options {
		o(rbmd)
	}
	return rbmd
}

func WithUser(user string) func(*RabbitMqDev) {
	return func(rbmd *RabbitMqDev) {
		rbmd.user = user
	}
}

func WithPassword(password string) func(*RabbitMqDev) {
	return func(rbmd *RabbitMqDev) {
		rbmd.password = password
	}
}

func WithHost(host string) func(*RabbitMqDev) {
	return func(rbmd *RabbitMqDev) {
		rbmd.host = host
	}
}

func WithPort(port string) func(*RabbitMqDev) {
	return func(rbmd *RabbitMqDev) {
		rbmd.port = port
	}
}

func (rbmd *RabbitMqDev) Connect() {
	connection, err := amqp.Dial("amqp://" + rbmd.user + ":" + rbmd.password + "@" + rbmd.host + ":" + rbmd.port + "/")
	if err != nil {
		panic(err)
	}
	rbmd.Connection = connection
}
