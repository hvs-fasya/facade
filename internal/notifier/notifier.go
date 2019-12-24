package notifier

type NotifierIface interface {
	NotifyCreateUser(string)
}

type notifier struct{}

func (n *notifier) NotifyCreateUser(userName string) {
	println("user '" + userName + "' created")
}

func NewNotifier() NotifierIface {
	return &notifier{}
}
