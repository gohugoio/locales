package vi_VN

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type vi_VN struct {
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
	currencyPositiveSuffix string
	currencyNegativeSuffix string
	monthsAbbreviated      []string
	monthsNarrow           []string
	monthsWide             []string
	daysAbbreviated        []string
	daysNarrow             []string
	daysShort              []string
	daysWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'vi_VN' locale
func New() locales.Translator {
	return &vi_VN{
		locale:                 "vi_VN",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         []locales.PluralRule{2, 6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "thg 1", "thg 2", "thg 3", "thg 4", "thg 5", "thg 6", "thg 7", "thg 8", "thg 9", "thg 10", "thg 11", "thg 12"},
		monthsWide:             []string{"", "tháng 1", "tháng 2", "tháng 3", "tháng 4", "tháng 5", "tháng 6", "tháng 7", "tháng 8", "tháng 9", "tháng 10", "tháng 11", "tháng 12"},
		daysAbbreviated:        []string{"CN", "Thứ 2", "Thứ 3", "Thứ 4", "Thứ 5", "Thứ 6", "Thứ 7"},
		daysNarrow:             []string{"CN", "T2", "T3", "T4", "T5", "T6", "T7"},
		daysShort:              []string{"CN", "T2", "T3", "T4", "T5", "T6", "T7"},
		daysWide:               []string{"Chủ Nhật", "Thứ Hai", "Thứ Ba", "Thứ Tư", "Thứ Năm", "Thứ Sáu", "Thứ Bảy"},
		timezones:              map[string]string{"ACDT": "Giờ Mùa Hè Miền Trung Australia", "ACST": "Giờ Chuẩn Miền Trung Australia", "ACT": "Giờ Chuẩn Acre", "ACWDT": "Giờ Mùa Hè Miền Trung Tây Australia", "ACWST": "Giờ Chuẩn Miền Trung Tây Australia", "ADT": "Giờ mùa hè Đại Tây Dương", "ADT Arabia": "Giờ Mùa Hè Ả Rập", "AEDT": "Giờ Mùa Hè Miền Đông Australia", "AEST": "Giờ Chuẩn Miền Đông Australia", "AFT": "Giờ Afghanistan", "AKDT": "Giờ Mùa Hè Alaska", "AKST": "Giờ Chuẩn Alaska", "AMST": "Giờ Mùa Hè Amazon", "AMST Armenia": "Giờ Mùa Hè Armenia", "AMT": "Giờ Chuẩn Amazon", "AMT Armenia": "Giờ Chuẩn Armenia", "ANAST": "Giờ mùa hè Anadyr", "ANAT": "Giờ Chuẩn Anadyr", "ARST": "Giờ Mùa Hè Argentina", "ART": "Giờ Chuẩn Argentina", "AST": "Giờ Chuẩn Đại Tây Dương", "AST Arabia": "Giờ chuẩn Ả Rập", "AWDT": "Giờ Mùa Hè Miền Tây Australia", "AWST": "Giờ Chuẩn Miền Tây Australia", "AZST": "Giờ Mùa Hè Azerbaijan", "AZT": "Giờ Chuẩn Azerbaijan", "BDT Bangladesh": "Giờ Mùa Hè Bangladesh", "BNT": "Giờ Brunei Darussalam", "BOT": "Giờ Bolivia", "BRST": "Giờ Mùa Hè Brasilia", "BRT": "Giờ Chuẩn Brasilia", "BST Bangladesh": "Giờ Chuẩn Bangladesh", "BT": "Giờ Bhutan", "CAST": "Giờ Casey", "CAT": "Giờ Trung Phi", "CCT": "Giờ Quần Đảo Cocos", "CDT": "Giờ mùa hè miền Trung", "CHADT": "Giờ Mùa Hè Chatham", "CHAST": "Giờ Chuẩn Chatham", "CHUT": "Giờ Chuuk", "CKT": "Giờ Chuẩn Quần Đảo Cook", "CKT DST": "Giờ Nửa Mùa Hè Quần Đảo Cook", "CLST": "Giờ Mùa Hè Chile", "CLT": "Giờ Chuẩn Chile", "COST": "Giờ Mùa Hè Colombia", "COT": "Giờ Chuẩn Colombia", "CST": "Giờ chuẩn miền Trung", "CST China": "Giờ Chuẩn Trung Quốc", "CST China DST": "Giờ Mùa Hè Trung Quốc", "CVST": "Giờ Mùa Hè Cape Verde", "CVT": "Giờ Chuẩn Cape Verde", "CXT": "Giờ Đảo Christmas", "ChST": "Giờ Chamorro", "ChST NMI": "Giờ Quần Đảo Bắc Mariana", "CuDT": "Giờ Mùa Hè Cuba", "CuST": "Giờ Chuẩn Cuba", "DAVT": "Giờ Davis", "DDUT": "Giờ Dumont-d’Urville", "EASST": "Giờ Mùa Hè Đảo Phục Sinh", "EAST": "Giờ Chuẩn Đảo Phục Sinh", "EAT": "Giờ Đông Phi", "ECT": "Giờ Ecuador", "EDT": "Giờ mùa hè miền Đông", "EGDT": "Giờ Mùa Hè Miền Đông Greenland", "EGST": "Giờ Chuẩn Miền Đông Greenland", "EST": "Giờ chuẩn miền Đông", "FEET": "Giờ Viễn đông Châu Âu", "FJT": "Giờ Chuẩn Fiji", "FJT Summer": "Giờ Mùa Hè Fiji", "FKST": "Giờ Mùa Hè Quần Đảo Falkland", "FKT": "Giờ Chuẩn Quần Đảo Falkland", "FNST": "Giờ Mùa Hè Fernando de Noronha", "FNT": "Giờ Chuẩn Fernando de Noronha", "GALT": "Giờ Galapagos", "GAMT": "Giờ Gambier", "GEST": "Giờ Mùa Hè Georgia", "GET": "Giờ Chuẩn Georgia", "GFT": "Giờ Guiana thuộc Pháp", "GIT": "Giờ Quần Đảo Gilbert", "GMT": "Giờ Trung bình Greenwich", "GNSST": "GNSST", "GNST": "GNST", "GST": "Giờ Nam Georgia", "GST Guam": "Giờ Chuẩn Guam", "GYT": "Giờ Guyana", "HADT": "Giờ Mùa Hè Hawaii-Aleut", "HAST": "Giờ Chuẩn Hawaii-Aleut", "HKST": "Giờ Mùa Hè Hồng Kông", "HKT": "Giờ Chuẩn Hồng Kông", "HOVST": "Giờ Mùa Hè Hovd", "HOVT": "Giờ Chuẩn Hovd", "ICT": "Giờ Đông Dương", "IDT": "Giờ Mùa Hè Israel", "IOT": "Giờ Ấn Độ Dương", "IRKST": "Giờ Mùa Hè Irkutsk", "IRKT": "Giờ Chuẩn Irkutsk", "IRST": "Giờ Chuẩn Iran", "IRST DST": "Giờ Mùa Hè Iran", "IST": "Giờ Chuẩn Ấn Độ", "IST Israel": "Giờ Chuẩn Israel", "JDT": "Giờ Mùa Hè Nhật Bản", "JST": "Giờ Chuẩn Nhật Bản", "KOST": "Giờ Kosrae", "KRAST": "Giờ Mùa Hè Krasnoyarsk", "KRAT": "Giờ Chuẩn Krasnoyarsk", "KST": "Giờ Chuẩn Hàn Quốc", "KST DST": "Giờ Mùa Hè Hàn Quốc", "LHDT": "Giờ Mùa Hè Lord Howe", "LHST": "Giờ Chuẩn Lord Howe", "LINT": "Giờ Quần Đảo Line", "MAGST": "Giờ mùa hè Magadan", "MAGT": "Giờ Chuẩn Magadan", "MART": "Giờ Marquesas", "MAWT": "Giờ Mawson", "MDT": "Giờ Mùa Hè Ma Cao", "MESZ": "Giờ mùa hè Trung Âu", "MEZ": "Giờ chuẩn Trung Âu", "MHT": "Giờ Quần Đảo Marshall", "MMT": "Giờ Myanmar", "MSD": "Giờ Mùa Hè Matxcơva", "MST": "Giờ Chuẩn Ma Cao", "MUST": "Giờ Mùa Hè Mauritius", "MUT": "Giờ Chuẩn Mauritius", "MVT": "Giờ Maldives", "MYT": "Giờ Malaysia", "NCT": "Giờ Chuẩn New Caledonia", "NDT": "Giờ Mùa Hè Newfoundland", "NDT New Caledonia": "Giờ Mùa Hè New Caledonia", "NFDT": "Giờ Mùa Hè Đảo Norfolk", "NFT": "Giờ Chuẩn Đảo Norfolk", "NOVST": "Giờ mùa hè Novosibirsk", "NOVT": "Giờ chuẩn Novosibirsk", "NPT": "Giờ Nepal", "NRT": "Giờ Nauru", "NST": "Giờ Chuẩn Newfoundland", "NUT": "Giờ Niue", "NZDT": "Giờ Mùa Hè New Zealand", "NZST": "Giờ Chuẩn New Zealand", "OESZ": "Giờ mùa hè Đông Âu", "OEZ": "Giờ chuẩn Đông Âu", "OMSST": "Giờ mùa hè Omsk", "OMST": "Giờ chuẩn Omsk", "PDT": "Giờ mùa hè Thái Bình Dương", "PDTM": "Giờ Mùa Hè Thái Bình Dương Mexico", "PETDT": "Giờ mùa hè Petropavlovsk-Kamchatski", "PETST": "Giờ chuẩn Petropavlovsk-Kamchatski", "PGT": "Giờ Papua New Guinea", "PHOT": "Giờ Quần Đảo Phoenix", "PKT": "Giờ Chuẩn Pakistan", "PKT DST": "Giờ Mùa Hè Pakistan", "PMDT": "Giờ Mùa Hè Saint Pierre và Miquelon", "PMST": "Giờ Chuẩn St. Pierre và Miquelon", "PONT": "Giờ Ponape", "PST": "Giờ chuẩn Thái Bình Dương", "PST Philippine": "Giờ Chuẩn Philippin", "PST Philippine DST": "Giờ Mùa Hè Philippin", "PST Pitcairn": "Giờ Pitcairn", "PSTM": "Giờ Chuẩn Thái Bình Dương Mexico", "PWT": "Giờ Palau", "PYST": "Giờ Mùa Hè Paraguay", "PYT": "Giờ Chuẩn Paraguay", "PYT Korea": "Giờ Bình Nhưỡng", "RET": "Giờ Reunion", "ROTT": "Giờ Rothera", "SAKST": "Giờ mùa hè Sakhalin", "SAKT": "Giờ Chuẩn Sakhalin", "SAMST": "Giờ mùa hè Samara", "SAMT": "Giờ Chuẩn Samara", "SAST": "Giờ Chuẩn Nam Phi", "SBT": "Giờ Quần Đảo Solomon", "SCT": "Giờ Seychelles", "SGT": "Giờ Singapore", "SLST": "Giờ Lanka", "SRT": "Giờ Suriname", "SST Samoa": "Giờ Chuẩn Samoa", "SST Samoa Apia": "Giờ Chuẩn Apia", "SST Samoa Apia DST": "Giờ Mùa Hè Apia", "SST Samoa DST": "Giờ ban ngày Samoa", "SYOT": "Giờ Syowa", "TAAF": "Giờ Nam Cực và Nam Nước Pháp", "TAHT": "Giờ Tahiti", "TJT": "Giờ Tajikistan", "TKT": "Giờ Tokelau", "TLT": "Giờ Đông Timor", "TMST": "Giờ Mùa Hè Turkmenistan", "TMT": "Giờ Chuẩn Turkmenistan", "TOST": "Giờ Mùa Hè Tonga", "TOT": "Giờ Chuẩn Tonga", "TVT": "Giờ Tuvalu", "TWT": "Giờ Chuẩn Đài Bắc", "TWT DST": "Giờ Mùa Hè Đài Bắc", "ULAST": "Giờ mùa hè Ulan Bator", "ULAT": "Giờ chuẩn Ulan Bator", "UYST": "Giờ Mùa Hè Uruguay", "UYT": "Giờ Chuẩn Uruguay", "UZT": "Giờ Chuẩn Uzbekistan", "UZT DST": "Giờ Mùa Hè Uzbekistan", "VET": "Giờ Venezuela", "VLAST": "Giờ mùa hè Vladivostok", "VLAT": "Giờ Chuẩn Vladivostok", "VOLST": "Giờ Mùa Hè Volgograd", "VOLT": "Giờ Chuẩn Volgograd", "VOST": "Giờ Vostok", "VUT": "Giờ Chuẩn Vanuatu", "VUT DST": "Giờ Mùa Hè Vanuatu", "WAKT": "Giờ Đảo Wake", "WARST": "Giờ mùa hè miền tây Argentina", "WART": "Giờ chuẩn miền tây Argentina", "WAST": "Giờ Tây Phi", "WAT": "Giờ Tây Phi", "WESZ": "Giờ mùa hè Tây Âu", "WEZ": "Giờ Chuẩn Tây Âu", "WFT": "Giờ Wallis và Futuna", "WGST": "Giờ Mùa Hè Miền Tây Greenland", "WGT": "Giờ Chuẩn Miền Tây Greenland", "WIB": "Giờ Miền Tây Indonesia", "WIT": "Giờ Miền Đông Indonesia", "WITA": "Giờ Miền Trung Indonesia", "YAKST": "Giờ mùa hè Yakutsk", "YAKT": "Giờ Chuẩn Yakutsk", "YEKST": "Giờ mùa hè Yekaterinburg", "YEKT": "Giờ Chuẩn Yekaterinburg", "YST": "Giờ Yukon", "МСК": "Giờ Chuẩn Matxcơva", "اقتاۋ": "Giờ Chuẩn Aqtau", "اقتاۋ قالاسى": "Giờ Mùa Hè Aqtau", "اقتوبە": "Giờ Chuẩn Aqtobe", "اقتوبە قالاسى": "Giờ Mùa Hè Aqtobe", "الماتى": "Giờ Chuẩn Almaty", "الماتى قالاسى": "Giờ Mùa Hè Almaty", "باتىس قازاق ەلى": "Giờ Miền Tây Kazakhstan", "شىعىش قازاق ەلى": "Giờ Miền Đông Kazakhstan", "قازاق ەلى": "Giờ Kazakhstan", "قىرعىزستان": "Giờ Kyrgystan", "قىزىلوردا": "Giờ Chuẩn Qyzylorda", "قىزىلوردا قالاسى": "Giờ Mùa Hè Qyzylorda", "∅∅∅": "Giờ mùa hè Azores"},
	}
}

// Locale returns the current translators string locale
func (vi *vi_VN) Locale() string {
	return vi.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'vi_VN'
func (vi *vi_VN) PluralsCardinal() []locales.PluralRule {
	return vi.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'vi_VN'
func (vi *vi_VN) PluralsOrdinal() []locales.PluralRule {
	return vi.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'vi_VN'
func (vi *vi_VN) PluralsRange() []locales.PluralRule {
	return vi.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'vi_VN'
func (vi *vi_VN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'vi_VN'
func (vi *vi_VN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'vi_VN'
func (vi *vi_VN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (vi *vi_VN) MonthAbbreviated(month time.Month) string {
	return vi.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (vi *vi_VN) MonthsAbbreviated() []string {
	return vi.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (vi *vi_VN) MonthNarrow(month time.Month) string {
	return vi.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (vi *vi_VN) MonthsNarrow() []string {
	return nil
}

// MonthWide returns the locales wide month given the 'month' provided
func (vi *vi_VN) MonthWide(month time.Month) string {
	return vi.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (vi *vi_VN) MonthsWide() []string {
	return vi.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (vi *vi_VN) WeekdayAbbreviated(weekday time.Weekday) string {
	return vi.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (vi *vi_VN) WeekdaysAbbreviated() []string {
	return vi.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (vi *vi_VN) WeekdayNarrow(weekday time.Weekday) string {
	return vi.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (vi *vi_VN) WeekdaysNarrow() []string {
	return vi.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (vi *vi_VN) WeekdayShort(weekday time.Weekday) string {
	return vi.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (vi *vi_VN) WeekdaysShort() []string {
	return vi.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (vi *vi_VN) WeekdayWide(weekday time.Weekday) string {
	return vi.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (vi *vi_VN) WeekdaysWide() []string {
	return vi.daysWide
}

// Decimal returns the decimal point of number
func (vi *vi_VN) Decimal() string {
	return vi.decimal
}

// Group returns the group of number
func (vi *vi_VN) Group() string {
	return vi.group
}

// Group returns the minus sign of number
func (vi *vi_VN) Minus() string {
	return vi.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'vi_VN' and handles both Whole and Real numbers based on 'v'
func (vi *vi_VN) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, vi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, vi.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, vi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'vi_VN' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (vi *vi_VN) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, vi.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, vi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, vi.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'vi_VN'
func (vi *vi_VN) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := vi.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, vi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, vi.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, vi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, vi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, vi.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'vi_VN'
// in accounting notation.
func (vi *vi_VN) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := vi.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, vi.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, vi.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, vi.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, vi.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, vi.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, vi.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'vi_VN'
func (vi *vi_VN) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'vi_VN'
func (vi *vi_VN) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, vi.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'vi_VN'
func (vi *vi_VN) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, vi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'vi_VN'
func (vi *vi_VN) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, vi.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, vi.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'vi_VN'
func (vi *vi_VN) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, vi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'vi_VN'
func (vi *vi_VN) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, vi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, vi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'vi_VN'
func (vi *vi_VN) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, vi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, vi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'vi_VN'
func (vi *vi_VN) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, vi.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, vi.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := vi.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
