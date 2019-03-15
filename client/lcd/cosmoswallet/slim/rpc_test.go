package slim

import "testing"

func TestQOSQueryAccountGet(t *testing.T) {
	SetBlockchainEntrance("192.168.1.23:1317", "forQmoonAddr")
	addr := "address1v26ael2jh0q7aetuk45yqf3jcyyywg2g6wq2tv"
	Aout := QOSQueryAccountGet(addr)
	t.Log(Aout)
}

func TestQSCQueryAccountGet(t *testing.T) {
	SetBlockchainEntrance("192.168.1.23:1317", "forQmoonAddr")
	addr := "address13l90zvt26szkrquutwpgj7kef58mgyntfs46l2"
	Aout := QSCQueryAccountGet(addr)
	t.Log(Aout)
}

func TestQSCtransferSendStr(t *testing.T) {
	SetBlockchainEntrance("192.168.1.23:1317", "forQmoonAddr")
	addrto := "address13l90zvt26szkrquutwpgj7kef58mgyntfs46l2"
	coinstr := "10000qos"
	privkey := "xGZuHJYesaYlgNJi7yeugj9A6Sc34f6plx5on6DDTTCVRb5f7neBxIsLUHgO+13Og38maO2E4kz55kX+4obHWQ=="
	chainid := "qos-test"
	Tout := QSCtransferSendStr(addrto, coinstr, privkey, chainid)
	t.Log(Tout)

}

