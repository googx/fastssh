/*
--------------------------------------------------
 File Name: groups.go
 Author: hanxu
 AuthorSite: http://www.googx.top/
 GitSource: https://github.com/googx/linuxShell
 Created Time: 2019-11-17-上午9:29
---------------------说明--------------------------
 主机组是一个抽象的概念,实际直接和生成的hosts文件无关,只是在系统中用来方便管理不同组别的hosts而已,不同组之间的host可以重复.
---------------------------------------------------
*/

package entity

type hostgroup struct {
	index uint
}
