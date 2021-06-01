package db

// WorkingTime 客服工作时段设定
type WorkingTime struct {
	DefaultField    `bson:",inline"`
	InboxID         string   // 用于区分不同渠道下不同的设定
	CloseByHolidays []string // 公共的假期：不区分年，应该是一个常规的日期：月/日
	CloseByDates    []string // 自定义假期：具体的日期：年/月/日
	StartWeek       int      // 周几开始
	EndWeek         int      // 周几结束
	StartTime       string   // 每天几点开始
	EndTime         string   // 每天几点结束
}
