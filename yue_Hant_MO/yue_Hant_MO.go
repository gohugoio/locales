package yue_Hant_MO

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type yue_Hant_MO struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	timeSeparator          string
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
	timezones              map[string]string
}

// New returns a new instance of translator for the 'yue_Hant_MO' locale
func New() locales.Translator {
	return &yue_Hant_MO{
		locale:                 "yue_Hant_MO",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AU$", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "￦", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "р.", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsWide:             []string{"", "1月", "2月", "3月", "4月", "5月", "6月", "7月", "8月", "9月", "10月", "11月", "12月"},
		daysNarrow:             []string{"日", "一", "二", "三", "四", "五", "六"},
		daysShort:              []string{"日", "一", "二", "三", "四", "五", "六"},
		daysWide:               []string{"星期日", "星期一", "星期二", "星期三", "星期四", "星期五", "星期六"},
		periodsAbbreviated:     []string{"上晝", "下晝"},
		timezones:              map[string]string{"ACDT": "澳洲中部夏令時間", "ACST": "澳洲中部標準時間", "ACT": "艾克標準時間", "ACWDT": "澳洲中西部夏令時間", "ACWST": "澳洲中西部標準時間", "ADT": "大西洋夏令時間", "ADT Arabia": "阿拉伯夏令時間", "AEDT": "澳洲東部夏令時間", "AEST": "澳洲東部標準時間", "AFT": "阿富汗時間", "AKDT": "阿拉斯加夏令時間", "AKST": "阿拉斯加標準時間", "AMST": "亞馬遜夏令時間", "AMST Armenia": "亞美尼亞夏令時間", "AMT": "亞馬遜標準時間", "AMT Armenia": "亞美尼亞標準時間", "ANAST": "阿那底河夏令時間", "ANAT": "阿那底河標準時間", "ARST": "阿根廷夏令時間", "ART": "阿根廷標準時間", "AST": "大西洋標準時間", "AST Arabia": "阿拉伯標準時間", "AWDT": "澳洲西部夏令時間", "AWST": "澳洲西部標準時間", "AZST": "亞塞拜然夏令時間", "AZT": "亞塞拜然標準時間", "BDT Bangladesh": "孟加拉夏令時間", "BNT": "汶萊時間", "BOT": "玻利維亞時間", "BRST": "巴西利亞夏令時間", "BRT": "巴西利亞標準時間", "BST Bangladesh": "孟加拉標準時間", "BT": "不丹時間", "CAST": "凱西站時間", "CAT": "中非時間", "CCT": "科科斯群島時間", "CDT": "中部夏令時間", "CHADT": "查坦群島夏令時間", "CHAST": "查坦群島標準時間", "CHUT": "楚克島時間", "CKT": "庫克群島標準時間", "CKT DST": "庫克群島半夏令時間", "CLST": "智利夏令時間", "CLT": "智利標準時間", "COST": "哥倫比亞夏令時間", "COT": "哥倫比亞標準時間", "CST": "中部標準時間", "CST China": "中國標準時間", "CST China DST": "中國夏令時間", "CVST": "維德角夏令時間", "CVT": "維德角標準時間", "CXT": "聖誕島時間", "ChST": "查莫洛時間", "ChST NMI": "北馬里亞納群島時間", "CuDT": "古巴夏令時間", "CuST": "古巴標準時間", "DAVT": "戴維斯時間", "DDUT": "杜蒙杜比爾時間", "EASST": "復活節島夏令時間", "EAST": "復活節島標準時間", "EAT": "東非時間", "ECT": "厄瓜多時間", "EDT": "東部夏令時間", "EGDT": "格陵蘭東部夏令時間", "EGST": "格陵蘭東部標準時間", "EST": "東部標準時間", "FEET": "歐洲遠東時間", "FJT": "斐濟標準時間", "FJT Summer": "斐濟夏令時間", "FKST": "福克蘭群島夏令時間", "FKT": "福克蘭群島標準時間", "FNST": "費爾南多 - 迪諾羅尼亞夏令時間", "FNT": "費爾南多 - 迪諾羅尼亞標準時間", "GALT": "加拉巴哥群島時間", "GAMT": "甘比爾群島時間", "GEST": "喬治亞夏令時間", "GET": "喬治亞標準時間", "GFT": "法屬圭亞那時間", "GIT": "吉爾伯特群島時間", "GMT": "格林威治標準時間", "GNSST": "GNSST", "GNST": "GNST", "GST": "波斯灣海域標準時間", "GST Guam": "關島標準時間", "GYT": "蓋亞那時間", "HADT": "夏威夷-阿留申夏令時間", "HAST": "夏威夷-阿留申標準時間", "HKST": "香港夏令時間", "HKT": "香港標準時間", "HOVST": "科布多夏令時間", "HOVT": "科布多標準時間", "ICT": "印度支那時間", "IDT": "以色列夏令時間", "IOT": "印度洋時間", "IRKST": "伊爾庫次克夏令時間", "IRKT": "伊爾庫次克標準時間", "IRST": "伊朗標準時間", "IRST DST": "伊朗夏令時間", "IST": "印度標準時間", "IST Israel": "以色列標準時間", "JDT": "日本夏令時間", "JST": "日本標準時間", "KOST": "科斯瑞時間", "KRAST": "克拉斯諾亞爾斯克夏令時間", "KRAT": "克拉斯諾亞爾斯克標準時間", "KST": "韓國標準時間", "KST DST": "韓國夏令時間", "LHDT": "豪勳爵島夏令時間", "LHST": "豪勳爵島標準時間", "LINT": "萊恩群島時間", "MAGST": "馬加丹夏令時間", "MAGT": "馬加丹標準時間", "MART": "馬可薩斯時間", "MAWT": "莫森時間", "MDT": "澳門夏令時間", "MESZ": "中歐夏令時間", "MEZ": "中歐標準時間", "MHT": "馬紹爾群島時間", "MMT": "緬甸時間", "MSD": "莫斯科夏令時間", "MST": "澳門標準時間", "MUST": "模里西斯夏令時間", "MUT": "模里西斯標準時間", "MVT": "馬爾地夫時間", "MYT": "馬來西亞時間", "NCT": "新喀里多尼亞標準時間", "NDT": "紐芬蘭夏令時間", "NDT New Caledonia": "新喀里多尼亞群島夏令時間", "NFDT": "諾福克島夏令時間", "NFT": "諾福克島標準時間", "NOVST": "新西伯利亞夏令時間", "NOVT": "新西伯利亞標準時間", "NPT": "尼泊爾時間", "NRT": "諾魯時間", "NST": "紐芬蘭標準時間", "NUT": "紐埃島時間", "NZDT": "紐西蘭夏令時間", "NZST": "紐西蘭標準時間", "OESZ": "東歐夏令時間", "OEZ": "東歐標準時間", "OMSST": "鄂木斯克夏令時間", "OMST": "鄂木斯克標準時間", "PDT": "太平洋夏令時間", "PDTM": "墨西哥太平洋夏令時間", "PETDT": "彼得羅巴甫洛夫斯克日光節約時間", "PETST": "彼得羅巴甫洛夫斯克標準時間", "PGT": "巴布亞紐幾內亞時間", "PHOT": "鳳凰群島時間", "PKT": "巴基斯坦標準時間", "PKT DST": "巴基斯坦夏令時間", "PMDT": "聖皮埃爾和密克隆群島夏令時間", "PMST": "聖皮埃爾和密克隆群島標準時間", "PONT": "波納佩時間", "PST": "太平洋標準時間", "PST Philippine": "菲律賓標準時間", "PST Philippine DST": "菲律賓夏令時間", "PST Pitcairn": "皮特肯時間", "PSTM": "墨西哥太平洋標準時間", "PWT": "帛琉時間", "PYST": "巴拉圭夏令時間", "PYT": "巴拉圭標準時間", "PYT Korea": "平壤時間", "RET": "留尼旺時間", "ROTT": "羅瑟拉時間", "SAKST": "庫頁島夏令時間", "SAKT": "庫頁島標準時間", "SAMST": "薩馬拉夏令時間", "SAMT": "薩馬拉標準時間", "SAST": "南非標準時間", "SBT": "索羅門群島時間", "SCT": "塞席爾時間", "SGT": "新加坡標準時間", "SLST": "蘭卡時間", "SRT": "蘇利南時間", "SST Samoa": "薩摩亞標準時間", "SST Samoa Apia": "阿皮亞標準時間", "SST Samoa Apia DST": "阿皮亞夏令時間", "SST Samoa DST": "薩摩亞夏令時間", "SYOT": "昭和基地時間", "TAAF": "法國南方及南極時間", "TAHT": "大溪地時間", "TJT": "塔吉克時間", "TKT": "托克勞群島時間", "TLT": "東帝汶時間", "TMST": "土庫曼夏令時間", "TMT": "土庫曼標準時間", "TOST": "東加夏令時間", "TOT": "東加標準時間", "TVT": "吐瓦魯時間", "TWT": "台北標準時間", "TWT DST": "台北夏令時間", "ULAST": "烏蘭巴托夏令時間", "ULAT": "烏蘭巴托標準時間", "UYST": "烏拉圭夏令時間", "UYT": "烏拉圭標準時間", "UZT": "烏茲別克標準時間", "UZT DST": "烏茲別克夏令時間", "VET": "委內瑞拉時間", "VLAST": "海參崴夏令時間", "VLAT": "海參崴標準時間", "VOLST": "伏爾加格勒夏令時間", "VOLT": "伏爾加格勒標準時間", "VOST": "沃斯托克時間", "VUT": "萬那杜標準時間", "VUT DST": "萬那杜夏令時間", "WAKT": "威克島時間", "WARST": "阿根廷西部夏令時間", "WART": "阿根廷西部標準時間", "WAST": "西非時間", "WAT": "西非時間", "WESZ": "西歐夏令時間", "WEZ": "西歐標準時間", "WFT": "瓦利斯和富圖納群島時間", "WGST": "格陵蘭西部夏令時間", "WGT": "格陵蘭西部標準時間", "WIB": "印尼西部時間", "WIT": "印尼東部時間", "WITA": "印尼中部時間", "YAKST": "雅庫次克夏令時間", "YAKT": "雅庫次克標準時間", "YEKST": "葉卡捷琳堡夏令時間", "YEKT": "葉卡捷琳堡標準時間", "YST": "育空時間", "МСК": "莫斯科標準時間", "اقتاۋ": "阿克陶標準時間", "اقتاۋ قالاسى": "阿克陶夏令時間", "اقتوبە": "阿克托比標準時間", "اقتوبە قالاسى": "阿克托比夏令時間", "الماتى": "阿拉木圖標準時間", "الماتى قالاسى": "阿拉木圖夏令時間", "باتىس قازاق ەلى": "西哈薩克時間", "شىعىش قازاق ەلى": "東哈薩克時間", "قازاق ەلى": "哈薩克時間", "قىرعىزستان": "吉爾吉斯時間", "قىزىلوردا": "克孜勒奧爾達標準時間", "قىزىلوردا قالاسى": "克孜勒奧爾達夏令時間", "∅∅∅": "亞速爾群島夏令時間"},
	}
}

// Locale returns the current translators string locale
func (yue *yue_Hant_MO) Locale() string {
	return yue.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'yue_Hant_MO'
func (yue *yue_Hant_MO) PluralsCardinal() []locales.PluralRule {
	return yue.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'yue_Hant_MO'
func (yue *yue_Hant_MO) PluralsOrdinal() []locales.PluralRule {
	return yue.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'yue_Hant_MO'
func (yue *yue_Hant_MO) PluralsRange() []locales.PluralRule {
	return yue.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'yue_Hant_MO'
func (yue *yue_Hant_MO) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'yue_Hant_MO'
func (yue *yue_Hant_MO) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'yue_Hant_MO'
func (yue *yue_Hant_MO) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (yue *yue_Hant_MO) MonthAbbreviated(month time.Month) string {
	if len(yue.monthsAbbreviated) == 0 {
		return ""
	}
	return yue.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (yue *yue_Hant_MO) MonthsAbbreviated() []string {
	return nil
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (yue *yue_Hant_MO) MonthNarrow(month time.Month) string {
	if len(yue.monthsNarrow) == 0 {
		return ""
	}
	return yue.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (yue *yue_Hant_MO) MonthsNarrow() []string {
	return nil
}

// MonthWide returns the locales wide month given the 'month' provided
func (yue *yue_Hant_MO) MonthWide(month time.Month) string {
	if len(yue.monthsWide) == 0 {
		return ""
	}
	return yue.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (yue *yue_Hant_MO) MonthsWide() []string {
	return yue.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (yue *yue_Hant_MO) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(yue.daysAbbreviated) == 0 {
		return ""
	}
	return yue.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (yue *yue_Hant_MO) WeekdaysAbbreviated() []string {
	return yue.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (yue *yue_Hant_MO) WeekdayNarrow(weekday time.Weekday) string {
	if len(yue.daysNarrow) == 0 {
		return ""
	}
	return yue.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (yue *yue_Hant_MO) WeekdaysNarrow() []string {
	return yue.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (yue *yue_Hant_MO) WeekdayShort(weekday time.Weekday) string {
	if len(yue.daysShort) == 0 {
		return ""
	}
	return yue.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (yue *yue_Hant_MO) WeekdaysShort() []string {
	return yue.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (yue *yue_Hant_MO) WeekdayWide(weekday time.Weekday) string {
	if len(yue.daysWide) == 0 {
		return ""
	}
	return yue.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (yue *yue_Hant_MO) WeekdaysWide() []string {
	return yue.daysWide
}

// Decimal returns the decimal point of number
func (yue *yue_Hant_MO) Decimal() string {
	return yue.decimal
}

// Group returns the group of number
func (yue *yue_Hant_MO) Group() string {
	return yue.group
}

// Group returns the minus sign of number
func (yue *yue_Hant_MO) Minus() string {
	return yue.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'yue_Hant_MO' and handles both Whole and Real numbers based on 'v'
func (yue *yue_Hant_MO) FmtNumber(num float64, v uint64) string {

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

// FmtPercent returns 'num' with digits/precision of 'v' for 'yue_Hant_MO' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (yue *yue_Hant_MO) FmtPercent(num float64, v uint64) string {
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

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'yue_Hant_MO'
func (yue *yue_Hant_MO) FmtCurrency(num float64, v uint64, currency currency.Type) string {

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

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'yue_Hant_MO'
// in accounting notation.
func (yue *yue_Hant_MO) FmtAccounting(num float64, v uint64, currency currency.Type) string {

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

// FmtDateShort returns the short date representation of 't' for 'yue_Hant_MO'
func (yue *yue_Hant_MO) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'yue_Hant_MO'
func (yue *yue_Hant_MO) FmtDateMedium(t time.Time) string {

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

// FmtDateLong returns the long date representation of 't' for 'yue_Hant_MO'
func (yue *yue_Hant_MO) FmtDateLong(t time.Time) string {

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

// FmtDateFull returns the full date representation of 't' for 'yue_Hant_MO'
func (yue *yue_Hant_MO) FmtDateFull(t time.Time) string {

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

// FmtTimeShort returns the short time representation of 't' for 'yue_Hant_MO'
func (yue *yue_Hant_MO) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, yue.periodsAbbreviated[0]...)
	} else {
		b = append(b, yue.periodsAbbreviated[1]...)
	}

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, yue.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'yue_Hant_MO'
func (yue *yue_Hant_MO) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, yue.periodsAbbreviated[0]...)
	} else {
		b = append(b, yue.periodsAbbreviated[1]...)
	}

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
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

// FmtTimeLong returns the long time representation of 't' for 'yue_Hant_MO'
func (yue *yue_Hant_MO) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, yue.periodsAbbreviated[0]...)
	} else {
		b = append(b, yue.periodsAbbreviated[1]...)
	}

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
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

// FmtTimeFull returns the full time representation of 't' for 'yue_Hant_MO'
func (yue *yue_Hant_MO) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, yue.periodsAbbreviated[0]...)
	} else {
		b = append(b, yue.periodsAbbreviated[1]...)
	}

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
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
