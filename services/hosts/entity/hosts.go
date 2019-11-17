/*
--------------------------------------------------
 File Name: hostsEntity.go
 Author: hanxu
 AuthorSite: http://www.googx.top/
 GitSource: https://github.com/googx/linuxShell
 Created Time: 2019-11-17-上午9:25
---------------------说明--------------------------

---------------------------------------------------
*/

package entity

import (
	"net"
)

type Hosts struct {
	index      int16
	hostip     net.IP
	hostdomain string
	comment    string
	enable     bool
}
