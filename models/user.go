package models

type WorkOrderList struct {
	Id      int
	UserId  int
	Title   string
	Content string
	Type    string
	Reply   string
}

type History struct {
	Id            int
	MysqlDns      string
	MysqlUser     string
	MysqlPassword string
	MysqlPorts    int
	MysqlFrom     string
}

type UserNas struct {
	Distributionid int
	Userid         string
	Groupid        string
	Nasname        string
	Ports          int
	Starospassword string
}
