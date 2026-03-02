package yue_Hans

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type yue_Hans struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	perMille               string
	timeSeparator          string
	inifinity              string
	currencies             []string // idx = enum of currency code
	currencyNegativePrefix string
	currencyNegativeSuffix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	periodsAbbreviated     []string
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'yue_Hans' locale
func New() locales.Translator {
	return &yue_Hans{
		locale:                 "yue_Hans",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		perMille:               "‰",
		timeSeparator:          ":",
		inifinity:              "∞",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AU$", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "￦", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "р.", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsNarrow:           []string{"", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
		monthsWide:             []string{"", "1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
		daysNarrow:             []string{"日", "一", "二", "三", "四", "五", "六"},
		daysShort:              []string{"日", "一", "二", "三", "四", "五", "六"},
		daysWide:               []string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"},
		periodsAbbreviated:     []string{"上昼", "下昼"},
		periodsNarrow:          []string{"上昼", "下昼"},
		periodsWide:            []string{"上昼", "下昼"},
		erasAbbreviated:        []string{"西元前", "西元"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"", ""},
		timezones:              map[string]string{"ACDT": "澳洲中部夏令时间", "ACST": "艾克夏令时间", "ACT": "艾克标准时间", "ACWDT": "澳洲中西部夏令时间", "ACWST": "澳洲中西部标准时间", "ADT": "大西洋夏令时间", "ADT Arabia": "阿拉伯夏令时间", "AEDT": "澳洲东部夏令时间", "AEST": "澳洲东部标准时间", "AFT": "阿富汗时间", "AKDT": "阿拉斯加夏令时间", "AKST": "阿拉斯加标准时间", "AMST": "亚马逊夏令时间", "AMST Armenia": "亚美尼亚夏令时间", "AMT": "亚马逊标准时间", "AMT Armenia": "亚美尼亚标准时间", "ANAST": "阿那底河夏令时间", "ANAT": "阿那底河标准时间", "ARST": "阿根廷夏令时间", "ART": "阿根廷标准时间", "AST": "大西洋标准时间", "AST Arabia": "阿拉伯标准时间", "AWDT": "澳洲西部夏令时间", "AWST": "澳洲西部标准时间", "AZST": "亚塞拜然夏令时间", "AZT": "亚塞拜然标准时间", "BDT Bangladesh": "孟加拉夏令时间", "BNT": "汶莱时间", "BOT": "玻利维亚时间", "BRST": "巴西利亚夏令时间", "BRT": "巴西利亚标准时间", "BST Bangladesh": "孟加拉标准时间", "BT": "不丹时间", "CAST": "凯西站时间", "CAT": "中非时间", "CCT": "科科斯群岛时间", "CDT": "中部夏令时间", "CHADT": "查坦群岛夏令时间", "CHAST": "查坦群岛标准时间", "CHUT": "楚克岛时间", "CKT": "库克群岛标准时间", "CKT DST": "库克群岛半夏令时间", "CLST": "智利夏令时间", "CLT": "智利标准时间", "COST": "哥伦比亚夏令时间", "COT": "哥伦比亚标准时间", "CST": "中部标准时间", "CST China": "中国标准时间", "CST China DST": "中国夏令时间", "CVST": "维德角夏令时间", "CVT": "维德角标准时间", "CXT": "圣诞岛时间", "ChST": "查莫洛时间", "ChST NMI": "北马里亚纳群岛时间", "CuDT": "古巴夏令时间", "CuST": "古巴标准时间", "DAVT": "戴维斯时间", "DDUT": "杜蒙杜比尔时间", "EASST": "复活节岛夏令时间", "EAST": "复活节岛标准时间", "EAT": "东非时间", "ECT": "厄瓜多时间", "EDT": "东部夏令时间", "EGDT": "格陵兰东部夏令时间", "EGST": "格陵兰东部标准时间", "EST": "东部标准时间", "FEET": "欧洲远东时间", "FJT": "斐济标准时间", "FJT Summer": "斐济夏令时间", "FKST": "福克兰群岛夏令时间", "FKT": "福克兰群岛标准时间", "FNST": "费尔南多 - 迪诺罗尼亚夏令时间", "FNT": "费尔南多 - 迪诺罗尼亚标准时间", "GALT": "加拉巴哥群岛时间", "GAMT": "甘比尔群岛时间", "GEST": "乔治亚夏令时间", "GET": "乔治亚标准时间", "GFT": "法属圭亚那时间", "GIT": "吉尔伯特群岛时间", "GMT": "格林威治标准时间", "GNSST": "GNSST", "GNST": "GNST", "GST": "波斯湾海域标准时间", "GST Guam": "关岛标准时间", "GYT": "盖亚那时间", "HADT": "夏威夷-阿留申夏令时间", "HAST": "夏威夷-阿留申标准时间", "HKST": "香港夏令时间", "HKT": "香港标准时间", "HOVST": "科布多夏令时间", "HOVT": "科布多标准时间", "ICT": "印度支那时间", "IDT": "以色列夏令时间", "IOT": "印度洋时间", "IRKST": "伊尔库次克夏令时间", "IRKT": "伊尔库次克标准时间", "IRST": "伊朗标准时间", "IRST DST": "伊朗夏令时间", "IST": "印度标准时间", "IST Israel": "以色列标准时间", "JDT": "日本夏令时间", "JST": "日本标准时间", "KOST": "科斯瑞时间", "KRAST": "克拉斯诺亚尔斯克夏令时间", "KRAT": "克拉斯诺亚尔斯克标准时间", "KST": "韩国标准时间", "KST DST": "韩国夏令时间", "LHDT": "豪勋爵岛夏令时间", "LHST": "豪勋爵岛标准时间", "LINT": "莱恩群岛时间", "MAGST": "马加丹夏令时间", "MAGT": "马加丹标准时间", "MART": "马可萨斯时间", "MAWT": "莫森时间", "MDT": "山区夏令时间", "MESZ": "中欧夏令时间", "MEZ": "中欧标准时间", "MHT": "马绍尔群岛时间", "MMT": "缅甸时间", "MSD": "莫斯科夏令时间", "MST": "山区标准时间", "MUST": "模里西斯夏令时间", "MUT": "模里西斯标准时间", "MVT": "马尔地夫时间", "MYT": "马来西亚时间", "NCT": "新喀里多尼亚标准时间", "NDT": "纽芬兰夏令时间", "NDT New Caledonia": "新喀里多尼亚群岛夏令时间", "NFDT": "诺福克岛夏令时间", "NFT": "诺福克岛标准时间", "NOVST": "新西伯利亚夏令时间", "NOVT": "新西伯利亚标准时间", "NPT": "尼泊尔时间", "NRT": "诺鲁时间", "NST": "纽芬兰标准时间", "NUT": "纽埃岛时间", "NZDT": "纽西兰夏令时间", "NZST": "纽西兰标准时间", "OESZ": "东欧夏令时间", "OEZ": "东欧标准时间", "OMSST": "鄂木斯克夏令时间", "OMST": "鄂木斯克标准时间", "PDT": "太平洋夏令时间", "PDTM": "墨西哥太平洋夏令时间", "PETDT": "彼得罗巴甫洛夫斯克日光节约时间", "PETST": "彼得罗巴甫洛夫斯克标准时间", "PGT": "巴布亚纽几内亚时间", "PHOT": "凤凰群岛时间", "PKT": "巴基斯坦标准时间", "PKT DST": "巴基斯坦夏令时间", "PMDT": "圣皮埃尔和密克隆群岛夏令时间", "PMST": "圣皮埃尔和密克隆群岛标准时间", "PONT": "波纳佩时间", "PST": "太平洋标准时间", "PST Philippine": "菲律宾标准时间", "PST Philippine DST": "菲律宾夏令时间", "PST Pitcairn": "皮特肯时间", "PSTM": "墨西哥太平洋标准时间", "PWT": "帛琉时间", "PYST": "巴拉圭夏令时间", "PYT": "巴拉圭标准时间", "PYT Korea": "平壤时间", "RET": "留尼旺时间", "ROTT": "罗瑟拉时间", "SAKST": "库页岛夏令时间", "SAKT": "库页岛标准时间", "SAMST": "萨马拉夏令时间", "SAMT": "萨马拉标准时间", "SAST": "南非标准时间", "SBT": "索罗门群岛时间", "SCT": "塞席尔时间", "SGT": "新加坡标准时间", "SLST": "兰卡时间", "SRT": "苏利南时间", "SST Samoa": "萨摩亚标准时间", "SST Samoa Apia": "阿皮亚标准时间", "SST Samoa Apia DST": "阿皮亚夏令时间", "SST Samoa DST": "萨摩亚夏令时间", "SYOT": "昭和基地时间", "TAAF": "法国南方及南极时间", "TAHT": "大溪地时间", "TJT": "塔吉克时间", "TKT": "托克劳群岛时间", "TLT": "东帝汶时间", "TMST": "土库曼夏令时间", "TMT": "土库曼标准时间", "TOST": "东加夏令时间", "TOT": "东加标准时间", "TVT": "吐瓦鲁时间", "TWT": "台北标准时间", "TWT DST": "台北夏令时间", "ULAST": "乌兰巴托夏令时间", "ULAT": "乌兰巴托标准时间", "UYST": "乌拉圭夏令时间", "UYT": "乌拉圭标准时间", "UZT": "乌兹别克标准时间", "UZT DST": "乌兹别克夏令时间", "VET": "委内瑞拉时间", "VLAST": "海参崴夏令时间", "VLAT": "海参崴标准时间", "VOLST": "伏尔加格勒夏令时间", "VOLT": "伏尔加格勒标准时间", "VOST": "沃斯托克时间", "VUT": "万那杜标准时间", "VUT DST": "万那杜夏令时间", "WAKT": "威克岛时间", "WARST": "阿根廷西部夏令时间", "WART": "阿根廷西部标准时间", "WAST": "西非时间", "WAT": "西非时间", "WESZ": "西欧夏令时间", "WEZ": "西欧标准时间", "WFT": "瓦利斯和富图纳群岛时间", "WGST": "格陵兰西部夏令时间", "WGT": "格陵兰西部标准时间", "WIB": "印尼西部时间", "WIT": "印尼东部时间", "WITA": "印尼中部时间", "YAKST": "雅库次克夏令时间", "YAKT": "雅库次克标准时间", "YEKST": "叶卡捷琳堡夏令时间", "YEKT": "叶卡捷琳堡标准时间", "YST": "育空时间", "МСК": "莫斯科标准时间", "اقتاۋ": "阿克陶标准时间", "اقتاۋ قالاسى": "阿克陶夏令时间", "اقتوبە": "阿克托比标准时间", "اقتوبە قالاسى": "阿克托比夏令时间", "الماتى": "阿拉木图标准时间", "الماتى قالاسى": "阿拉木图夏令时间", "باتىس قازاق ەلى": "西哈萨克时间", "شىعىش قازاق ەلى": "东哈萨克时间", "قازاق ەلى": "哈萨克时间", "قىرعىزستان": "吉尔吉斯时间", "قىزىلوردا": "克孜勒奥尔达标准时间", "قىزىلوردا قالاسى": "克孜勒奥尔达夏令时间", "∅∅∅": "秘鲁夏令时间"},
	}
}

// Locale returns the current translators string locale
func (yue *yue_Hans) Locale() string {
	return yue.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'yue_Hans'
func (yue *yue_Hans) PluralsCardinal() []locales.PluralRule {
	return yue.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'yue_Hans'
func (yue *yue_Hans) PluralsOrdinal() []locales.PluralRule {
	return yue.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'yue_Hans'
func (yue *yue_Hans) PluralsRange() []locales.PluralRule {
	return yue.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'yue_Hans'
func (yue *yue_Hans) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'yue_Hans'
func (yue *yue_Hans) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'yue_Hans'
func (yue *yue_Hans) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (yue *yue_Hans) MonthAbbreviated(month time.Month) string {
	return yue.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (yue *yue_Hans) MonthsAbbreviated() []string {
	return nil
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (yue *yue_Hans) MonthNarrow(month time.Month) string {
	return yue.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (yue *yue_Hans) MonthsNarrow() []string {
	return yue.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (yue *yue_Hans) MonthWide(month time.Month) string {
	return yue.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (yue *yue_Hans) MonthsWide() []string {
	return yue.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (yue *yue_Hans) WeekdayAbbreviated(weekday time.Weekday) string {
	return yue.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (yue *yue_Hans) WeekdaysAbbreviated() []string {
	return yue.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (yue *yue_Hans) WeekdayNarrow(weekday time.Weekday) string {
	return yue.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (yue *yue_Hans) WeekdaysNarrow() []string {
	return yue.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (yue *yue_Hans) WeekdayShort(weekday time.Weekday) string {
	return yue.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (yue *yue_Hans) WeekdaysShort() []string {
	return yue.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (yue *yue_Hans) WeekdayWide(weekday time.Weekday) string {
	return yue.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (yue *yue_Hans) WeekdaysWide() []string {
	return yue.daysWide
}

// Decimal returns the decimal point of number
func (yue *yue_Hans) Decimal() string {
	return yue.decimal
}

// Group returns the group of number
func (yue *yue_Hans) Group() string {
	return yue.group
}

// Group returns the minus sign of number
func (yue *yue_Hans) Minus() string {
	return yue.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'yue_Hans' and handles both Whole and Real numbers based on 'v'
func (yue *yue_Hans) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, yue.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, yue.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, yue.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'yue_Hans' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (yue *yue_Hans) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, yue.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, yue.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, yue.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'yue_Hans'
func (yue *yue_Hans) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := yue.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, yue.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, yue.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		b = append(b, yue.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, yue.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'yue_Hans'
// in accounting notation.
func (yue *yue_Hans) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := yue.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, yue.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, yue.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		b = append(b, yue.currencyNegativePrefix[0])

	} else {
		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, yue.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, yue.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'yue_Hans'
func (yue *yue_Hans) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'yue_Hans'
func (yue *yue_Hans) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xe5, 0xb9, 0xb4}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0xe6, 0x9c, 0x88}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xe6, 0x97, 0xa5}...)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'yue_Hans'
func (yue *yue_Hans) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xe5, 0xb9, 0xb4}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0xe6, 0x9c, 0x88}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xe6, 0x97, 0xa5}...)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'yue_Hans'
func (yue *yue_Hans) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0xe5, 0xb9, 0xb4}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0xe6, 0x9c, 0x88}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0xe6, 0x97, 0xa5, 0x20}...)
	b = append(b, yue.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'yue_Hans'
func (yue *yue_Hans) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yue.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'yue_Hans'
func (yue *yue_Hans) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yue.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, yue.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'yue_Hans'
func (yue *yue_Hans) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yue.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, yue.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x5b}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	b = append(b, []byte{0x5d}...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'yue_Hans'
func (yue *yue_Hans) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, yue.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, yue.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x5b}...)

	tz, _ := t.Zone()

	if btz, ok := yue.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	b = append(b, []byte{0x5d}...)

	return string(b)
}
