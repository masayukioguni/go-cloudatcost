package cloudatcost

type StandardResponse struct {
	Status   string `json:"status"`
	Time     int    `json:"time"`
	API      string `json:"api"`
	Serverid string `json:"serverid"`
	Result   string `json:"result"`
}

type ConsoleResponse struct {
	Status   string `json:"status"`
	Time     int    `json:"time"`
	API      string `json:"api"`
	Serverid string `json:"serverid"`
	Console  string `json:"console"`
}

type PowerOperationResponse struct {
	Status   string `json:"status"`
	Time     int    `json:"time"`
	API      string `json:"api"`
	Action   string `json:"action"`
	Serverid string `json:"serverid"`
	Taskid   int64  `json:"taskid"`
	Result   string `json:"result"`
}

type ListTask struct {
	Cid        string `json:"cid"`
	Idf        string `json:"idf"`
	Serverid   string `json:"serverid"`
	Action     string `json:"action"`
	Status     string `json:"status"`
	Starttime  string `json:"starttime"`
	Finishtime string `json:"finishtime"`
}

type ListTasksResponse struct {
	Status string     `json:"status"`
	Time   int        `json:"time"`
	API    string     `json:"api"`
	Cid    string     `json:"cid"`
	Action string     `json:"action"`
	Data   []ListTask `json:"data"`
}

type ListTemplate struct {
	ID     string `json:"id"`
	Detail string `json:"detail"`
}

type ListTemplatesResponse struct {
	Status string         `json:"status"`
	Time   int            `json:"time"`
	API    string         `json:"api"`
	Action string         `json:"action"`
	Data   []ListTemplate `json:"data"`
}

type ListServer struct {
	Sid        string      `json:"sid"`
	ID         string      `json:"id"`
	Packageid  string      `json:"packageid"`
	Servername string      `json:"servername"`
	Lable      interface{} `json:"lable"`
	Vmname     string      `json:"vmname"`
	IP         string      `json:"ip"`
	Netmask    string      `json:"netmask"`
	Gateway    string      `json:"gateway"`
	Portgroup  string      `json:"portgroup"`
	Hostname   string      `json:"hostname"`
	Rootpass   string      `json:"rootpass"`
	Vncport    string      `json:"vncport"`
	Vncpass    string      `json:"vncpass"`
	Servertype string      `json:"servertype"`
	Template   string      `json:"template"`
	CPU        string      `json:"cpu"`
	Cpuusage   string      `json:"cpuusage"`
	RAM        string      `json:"ram"`
	Ramusage   string      `json:"ramusage"`
	Storage    string      `json:"storage"`
	Hdusage    string      `json:"hdusage"`
	Sdate      string      `json:"sdate"`
	Status     string      `json:"status"`
	PanelNote  string      `json:"panel_note"`
	Mode       string      `json:"mode"`
	UID        string      `json:"uid"`
}

type ListServersResponse struct {
	Status string       `json:"status"`
	Time   int          `json:"time"`
	API    string       `json:"api"`
	Action string       `json:"action"`
	Data   []ListServer `json:"data"`
}

type CloudProServerResponse struct {
	Status string `json:"status"`
	Time   int    `json:"time"`
	API    string `json:"api"`
	Action string `json:"action"`
	Taskid int    `json:"taskid"`
	Result string `json:"result"`
}

type CloudProResourcesTotalResponse struct {
	CPU     string `json:"cpu_total"`
	Ram     string `json:"ram_total"`
	Stroage string `json:"storage_total"`
}

type CloudProResourcesUsedResponse struct {
	CPU     string `json:"cpu_used"`
	Ram     string `json:"ram_used"`
	Stroage string `json:"storage_used"`
}

type CloudProResourcesData struct {
	Total []CloudProResourcesTotalResponse `json:"total"`
	Used  CloudProResourcesUsedResponse    `json:"used"`
}

type CloudProResourcesResponse struct {
	Status string                `json:"status"`
	Time   int                   `json:"time"`
	API    string                `json:"api"`
	Action string                `json:"action"`
	Data   CloudProResourcesData `json:"data"`
}

type CreateServerOptions struct {
	Cpu     string `json:"cpu"`
	Ram     string `json:"ram"`
	Storage string `json:"storage"`
	OS      string `json:"os"`
}
