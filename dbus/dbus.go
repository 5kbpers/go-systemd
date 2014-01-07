// Integration with the systemd D-Bus API.  See http://www.freedesktop.org/wiki/Software/systemd/dbus/
package dbus

import (
	"strings"
	"sync"

	"github.com/guelfey/go.dbus"
)

const signalBuffer = 100

// ObjectPath creates a dbus.ObjectPath using the rules that systemd uses for
// serializing special characters.
func ObjectPath(path string) dbus.ObjectPath {
	path = strings.Replace(path, ".", "_2e", -1)
	path = strings.Replace(path, "-", "_2d", -1)
	path = strings.Replace(path, "@", "_40", -1)

	return dbus.ObjectPath(path)
}

type Conn struct {
	sysconn     *dbus.Conn
	sysobj      *dbus.Object
	jobListener struct {
		jobs map[dbus.ObjectPath]chan string
		sync.Mutex
	}
	subscriber struct {
		updateCh chan<- *SubStateUpdate
		errCh    chan<- error
		sync.Mutex
		ignore      map[dbus.ObjectPath]int64
		cleanIgnore int64
	}
	dispatch map[string]func(dbus.Signal)
}

func New() (*Conn, error) {
	c := new(Conn)

	if err := c.initConnection(); err != nil {
		return nil, err
	}

	c.initJobs()
	return c, nil
}

func (c *Conn) initConnection() error {
	var err error
	c.sysconn, err = dbus.SystemBusPrivate()
	if err != nil {
		return err
	}

	err = c.sysconn.Auth(nil)
	if err != nil {
		c.sysconn.Close()
		return err
	}

	err = c.sysconn.Hello()
	if err != nil {
		c.sysconn.Close()
		return err
	}

	c.sysobj = c.sysconn.Object("org.freedesktop.systemd1", dbus.ObjectPath("/org/freedesktop/systemd1"))

	return nil
}


