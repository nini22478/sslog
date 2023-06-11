package Console

import (
	"jdudp/logics"

	"github.com/robfig/cron/v3"
)

/**
 * Created by Goland
 * User: wkk alisleepy@hotmail.com
 * Time: 2023/1/8 - 02:02
 * Desc: <统一定时任务管理>
 */

// Conrs 定时器
var Conrs *cron.Cron

// HandleCorn 定时任务入口
func init() {
	// dump.P("开始处理定时任务")
	Conrs = cron.New() // 定时任务
	Conrs.Start()
	// 删除离职员工
	_, err := Conrs.AddFunc("@every 5m", logics.Get) // 每隔1分钟执行一次DeleteStaffs方法
	if err != nil {
		// dump.P("删除员工定时任务失败。。。")
		return
	}
}
