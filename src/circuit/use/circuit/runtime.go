package circuit

type runtime interface {
	WorkerAddr() Addr
	SetBoot(interface{})
	Spawn(Host, []string, Func, ...interface{}) ([]interface{}, Addr, error)
	Daemonize(func())
	Kill(Addr) error
	Dial(Addr, string) X
	DialSelf(string) interface{}
	TryDial(Addr, string) (X, error)
	Listen(string, interface{})
	Export(...interface{}) interface{}
	Import(interface{}) ([]interface{}, string, error)
	Ref(interface{}) X
	PermRef(interface{}) XPerm
}
