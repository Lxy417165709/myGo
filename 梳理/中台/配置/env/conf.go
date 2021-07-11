package env



import (
	"encoding/json"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
)

var Conf = &conf{}

type conf struct {
	Grafana Grafana `json:"grafana"`
	Spider  Spider  `json:"spider"`
	Mysql   MySql   `json:"mysql"`
	Http    Http    `json:"http"`
	Bot     Bot     `json:"bot"`
}

func (c *conf) Load(pathOfConfFile string) error {
	confFileData, err := ioutil.ReadFile(pathOfConfFile)
	if err != nil {
		logs.Error(err)
		return err
	}
	if err := json.Unmarshal(confFileData, c); err != nil {
		logs.Error(err)
		return err
	}
	return nil
}

type Bot struct {
	WebHook                 string `json:"webhook"`
	RequestBodyTemplatePath string `json:"request_body_template_path"`
}

type Grafana struct {
	Url string `json:"url"`
}

type Http struct {
	Port int `json:"port"`
}

type MySql struct {
	Host     string `json:"host"`
	User     string `json:"user"`
	Password string `json:"password"`
	DBName   string `json:"dbname"`
}

type Spider struct {
	TraceBlockAddr         string `json:"trace_block_addr"`
	StartTakeTimeAddr      string `json:"start_take_time_addr"`
	DailyStartTakeTimeAddr string `json:"daily_start_take_time_addr"`
}
