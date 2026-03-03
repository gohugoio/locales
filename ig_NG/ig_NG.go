package ig_NG

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ig_NG struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	percentSuffix          string
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
	timezones              map[string]string
}

// New returns a new instance of translator for the 'ig_NG' locale
func New() locales.Translator {
	return &ig_NG{
		locale:                 "ig_NG",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		minus:                  "\u200f-",
		percent:                "٪\u200f",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "₦", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "Jen", "Feb", "Maa", "Epr", "Mee", "Juu", "Jul", "Ọgọ", "Sep", "Ọkt", "Nov", "Dis"},
		monthsNarrow:           []string{"", "J", "F", "M", "E", "M", "J", "J", "Ọ", "S", "Ọ", "N", "D"},
		monthsWide:             []string{"", "Jenụwarị", "Febrụwarị", "Maachị", "Epreel", "Mee", "Jun", "Julaị", "Ọgọọst", "Septemba", "Ọktoba", "Novemba", "Disemba"},
		daysAbbreviated:        []string{"Sọn", "Mọn", "Tiu", "Wen", "Tọọ", "Fraị", "Sat"},
		daysWide:               []string{"Sọndee", "Mọnde", "Tiuzdee", "Wenezdee", "Tọọzdee", "Fraịdee", "Satọdee"},
		timezones:              map[string]string{"ACDT": "Oge Ihe Etiti Australia", "ACST": "Oge Izugbe Etiti Australia", "ACT": "ACT", "ACWDT": "Oge Ihe Mpaghara Ọdịda Anyanwụ Etiti Australia", "ACWST": "Oge Izugbe Mpaghara Ọdịda Anyanwụ Etiti Australia", "ADT": "Oge Ihe Mpaghara Atlantic", "ADT Arabia": "Oge Ihe Arab", "AEDT": "Oge Ihe Mpaghara Ọwụwa Anyanwụ Australia", "AEST": "Oge Izugbe Mpaghara Ọwụwa Anyanwụ Australia", "AFT": "Oge Afghanistan", "AKDT": "Oge Ihe Alaska", "AKST": "Oge Izugbe Alaska", "AMST": "Oge Okpomọkụ Amazon", "AMST Armenia": "Oge Okpomọkụ Armenia", "AMT": "Oge Izugbe Amazon", "AMT Armenia": "Oge Izugbe Armenia", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Oge Okpomọkụ Argentina", "ART": "Oge Izugbe Argentina", "AST": "Oge Izugbe Mpaghara Atlantic", "AST Arabia": "Oge Izugbe Arab", "AWDT": "Oge Ihe Mpaghara Ọdịda Anyanwụ Australia", "AWST": "Oge Izugbe Mpaghara Ọdịda Anyanwụ Australia", "AZST": "Oge Okpomọkụ Azerbaijan", "AZT": "Oge Izugbe Azerbaijan", "BDT Bangladesh": "Oge Okpomọkụ Bangladesh", "BNT": "Oge Brunei Darussalam", "BOT": "Oge Bolivia", "BRST": "Oge Okpomọkụ Brasilia", "BRT": "Oge Izugbe Brasilia", "BST Bangladesh": "Oge Izugbe Bangladesh", "BT": "Oge Bhutan", "CAST": "CAST", "CAT": "Oge Etiti Afrịka", "CCT": "Oge Cocos Islands", "CDT": "Oge Ihe Mpaghara Etiti", "CHADT": "Oge Ihe Chatham", "CHAST": "Oge Izugbe Chatham", "CHUT": "Oge Chuuk", "CKT": "Oge Izugbe Cook Islands", "CKT DST": "Oge Ọkara Okpomọkụ Cook Islands", "CLST": "Oge Okpomọkụ Chile", "CLT": "Oge Izugbe Chile", "COST": "Oge Okpomọkụ Columbia", "COT": "Oge Izugbe Columbia", "CST": "Oge Izugbe Mpaghara Etiti", "CST China": "Oge Izugbe China", "CST China DST": "Oge Ihe China", "CVST": "Oge Okpomọkụ Cape Verde", "CVT": "Oge Izugbe Cape Verde", "CXT": "Oge Ekeresimesi Island", "ChST": "Oge Izugbe Chamorro", "ChST NMI": "ChST NMI", "CuDT": "Oge Ihe Mpaghara Cuba", "CuST": "Oge Izugbe Cuba", "DAVT": "Oge Davis", "DDUT": "Oge Dumont-d’Urville", "EASST": "Oge Okpomọkụ Mpaghara Ọwụwa Anyanwụ Island", "EAST": "Oge Izugbe Mpaghara Ọwụwa Anyanwụ Island", "EAT": "Oge Mpaghara Ọwụwa Anyanwụ Afrịka", "ECT": "Oge Ecuador", "EDT": "Oge Ihe Mpaghara Ọwụwa Anyanwụ", "EGDT": "Oge Okpomọkụ Mpaghara Ọwụwa Anyanwụ Greenland", "EGST": "Oge Izugbe Mpaghara Ọwụwa Anyanwụ Greenland", "EST": "Oge Izugbe Mpaghara Ọwụwa Anyanwụ", "FEET": "Further-eastern European Time", "FJT": "Oge Izugbe Fiji", "FJT Summer": "Oge Okpomọkụ Fiji", "FKST": "Oge Okpomọkụ Falkland Islands", "FKT": "Oge Izugbe Falkland Islands", "FNST": "Oge Okpomọkụ Fernando de Noronha", "FNT": "Oge Izugbe Fernando de Noronha", "GALT": "Oge Galapagos", "GAMT": "Oge Gambier", "GEST": "Oge Okpomọkụ Georgia", "GET": "Oge Izugbe Georgia", "GFT": "Oge French Guiana", "GIT": "Oge Gilbert Islands", "GMT": "Oge Mpaghara Greemwich Mean", "GNSST": "GNSST", "GNST": "GNST", "GST": "Oge South Georgia", "GST Guam": "GST Guam", "GYT": "Oge Guyana", "HADT": "Oge Ihe Hawaii-Aleutian", "HAST": "Oge Izugbe Hawaii-Aleutian", "HKST": "Oge Okpomọkụ Hong Kong", "HKT": "Oge Izugbe Hong Kong", "HOVST": "Oge Okpomọkụ Hovd", "HOVT": "Oge Izugbe Hovd", "ICT": "Oge Indochina", "IDT": "Oge Ihe Israel", "IOT": "Oge Osimiri India", "IRKST": "Oge Okpomọkụ Irkutsk", "IRKT": "Oge Izugbe Irkutsk", "IRST": "Oge Izugbe Iran", "IRST DST": "Oge Ihe Iran", "IST": "Oge Izugbe India", "IST Israel": "Oge Izugbe Israel", "JDT": "Oge Ihe Japan", "JST": "Oge Izugbe Japan", "KOST": "Oge Kosrae", "KRAST": "Oge Okpomọkụ Krasnoyarsk", "KRAT": "Oge Izugbe Krasnoyarsk", "KST": "Oge Izugbe Korea", "KST DST": "Oge Ihe Korea", "LHDT": "Oge Ihe Lord Howe", "LHST": "Oge Izugbe Lord Howe", "LINT": "Oge Line Islands", "MAGST": "Oge Okpomọkụ Magadan", "MAGT": "Oge Izugbe Magadan", "MART": "Oge Marquesas", "MAWT": "Oge Mawson", "MDT": "MDT", "MESZ": "Oge Okpomọkụ Mpaghara Etiti Europe", "MEZ": "Oge Izugbe Mpaghara Etiti Europe", "MHT": "Oge Marshall Islands", "MMT": "Oge Myanmar", "MSD": "Oge Okpomọkụ Moscow", "MST": "MST", "MUST": "Oge Okpomọkụ Mauritius", "MUT": "Oge Izugbe Mauritius", "MVT": "Oge Maldives", "MYT": "Oge Malaysia", "NCT": "Oge Izugbe New Caledonia", "NDT": "Oge Ihe Newfoundland", "NDT New Caledonia": "Oge Okpomọkụ New Caledonia", "NFDT": "Oge Okpomọkụ Norfolk Island", "NFT": "Oge Izugbe Norfolk Island", "NOVST": "Oge Okpomọkụ Novosibirsk", "NOVT": "Oge Izugbe Novosibirsk", "NPT": "Oge Nepal", "NRT": "Oge Nauru", "NST": "Oge Izugbe Newfoundland", "NUT": "Oge Niue", "NZDT": "Oge Ihe New Zealand", "NZST": "Oge Izugbe New Zealand", "OESZ": "Oge Okpomọkụ Mpaghara Ọwụwa Anyanwụ Europe", "OEZ": "Oge Izugbe Mpaghara Ọwụwa Anyanwụ Europe", "OMSST": "Oge Okpomọkụ Omsk", "OMST": "Oge Izugbe Omsk", "PDT": "Oge Ihe Mpaghara Pacific", "PDTM": "Oge Ihe Mexican Pacific", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Oge Papua New Guinea", "PHOT": "Oge Phoenix Islands", "PKT": "Oge Izugbe Pakistan", "PKT DST": "Oge Okpomọkụ Pakistan", "PMDT": "Oge Ihe St. Pierre & Miquelon", "PMST": "Oge Izugbe St. Pierre & Miquelon", "PONT": "Oge Ponape", "PST": "Oge Izugbe Mpaghara Pacific", "PST Philippine": "Oge Izugbe Philippine", "PST Philippine DST": "Oge Okpomọkụ Philippine", "PST Pitcairn": "Oge Pitcairn", "PSTM": "Oge Izugbe Mexican Pacific", "PWT": "Oge Palau", "PYST": "Oge Okpomọkụ Paraguay", "PYT": "Oge Izugbe Paraguay", "PYT Korea": "Oge Pyongyang", "RET": "Oge Réunion", "ROTT": "Oge Rothera", "SAKST": "Oge Okpomọkụ Sakhalin", "SAKT": "Oge Izugbe Sakhalin", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Oge Izugbe Mpaghara Mgbada Ugwu Afrịka", "SBT": "Oge Solomon Islands", "SCT": "Oge Seychelles", "SGT": "Oge Izugbe Singapore", "SLST": "SLST", "SRT": "Oge Suriname", "SST Samoa": "Oge Izugbe Samoa", "SST Samoa Apia": "Oge Izugbe Apia", "SST Samoa Apia DST": "Oge Ihe Apia", "SST Samoa DST": "Oge Ihe Samoa", "SYOT": "Oge Syowa", "TAAF": "Oge French Southern & Antarctic", "TAHT": "Oge Tahiti", "TJT": "Oge Tajikistan", "TKT": "Oge Tokelau", "TLT": "Oge Mpaghara Ọwụwa Anyanwụ Timor", "TMST": "Oge Okpomọkụ Turkmenist", "TMT": "Oge Izugbe Turkmenist", "TOST": "Oge Okpomọkụ Tonga", "TOT": "Oge Izugbe Tonga", "TVT": "Oge Tuvalu", "TWT": "Oge Izugbe Taipei", "TWT DST": "Oge Ihe Taipei", "ULAST": "Oge Okpomọkụ Ulaanbaatar", "ULAT": "Oge Izugbe Ulaanbaatar", "UYST": "Oge Okpomọkụ Uruguay", "UYT": "Oge Izugbe Uruguay", "UZT": "Oge Izugbe Uzbekist", "UZT DST": "Oge Okpomọkụ Uzbekist", "VET": "Oge Venezuela", "VLAST": "Oge Okpomọkụ Vladivostok", "VLAT": "Oge Izugbe Vladivostok", "VOLST": "Oge Okpomọkụ Volgograd", "VOLT": "Oge Izugbe Volgograd", "VOST": "Oge Vostok", "VUT": "Oge Izugbe Vanuatu", "VUT DST": "Oge Okpomọkụ Vanuatu", "WAKT": "Oge Wake Island", "WARST": "Oge Okpomọkụ Mpaghara Ọdịda Anyanwụ Argentina", "WART": "Oge Izugbe Mpaghara Ọdịda Anyanwụ Argentina", "WAST": "Oge Mpaghara Ọdịda Anyanwụ Afrịka", "WAT": "Oge Mpaghara Ọdịda Anyanwụ Afrịka", "WESZ": "Oge Okpomọkụ Mpaghara Ọdịda Anyanwụ Europe", "WEZ": "Oge Izugbe Mpaghara Ọdịda Anyanwụ Europe", "WFT": "Oge Wallis & Futuna", "WGST": "Oge Okpomọkụ Mpaghara Ọdịda Anyanwụ Greenland", "WGT": "Oge Izugbe Mpaghara Ọdịda Anyanwụ Greenland", "WIB": "Oge Mpaghara Ọdịda Anyanwụ Indonesia", "WIT": "Oge Mpaghara Ọwụwa Anyanwụ Indonesia", "WITA": "Oge Etiti Indonesia", "YAKST": "Oge Okpomọkụ Yakutsk", "YAKT": "Oge Izugbe Yakutsk", "YEKST": "Oge Okpomọkụ Yekaterinburg", "YEKT": "Oge Izugbe Yekaterinburg", "YST": "Oge Yukon", "МСК": "Oge Izugbe Moscow", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Oge Mpaghara Ọdịda Anyanwụ Kazakhstan", "شىعىش قازاق ەلى": "Oge Mpaghara Ọwụwa Anyanwụ Kazakhstan", "قازاق ەلى": "Oge Kazakhstan", "قىرعىزستان": "Oge Kyrgyzstan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Oge Okpomọkụ Azores"},
	}
}

// Locale returns the current translators string locale
func (ig *ig_NG) Locale() string {
	return ig.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ig_NG'
func (ig *ig_NG) PluralsCardinal() []locales.PluralRule {
	return ig.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ig_NG'
func (ig *ig_NG) PluralsOrdinal() []locales.PluralRule {
	return ig.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ig_NG'
func (ig *ig_NG) PluralsRange() []locales.PluralRule {
	return ig.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ig_NG'
func (ig *ig_NG) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ig_NG'
func (ig *ig_NG) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ig_NG'
func (ig *ig_NG) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ig *ig_NG) MonthAbbreviated(month time.Month) string {
	if len(ig.monthsAbbreviated) == 0 {
		return ""
	}
	return ig.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ig *ig_NG) MonthsAbbreviated() []string {
	return ig.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ig *ig_NG) MonthNarrow(month time.Month) string {
	if len(ig.monthsNarrow) == 0 {
		return ""
	}
	return ig.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ig *ig_NG) MonthsNarrow() []string {
	return ig.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ig *ig_NG) MonthWide(month time.Month) string {
	if len(ig.monthsWide) == 0 {
		return ""
	}
	return ig.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ig *ig_NG) MonthsWide() []string {
	return ig.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ig *ig_NG) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(ig.daysAbbreviated) == 0 {
		return ""
	}
	return ig.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ig *ig_NG) WeekdaysAbbreviated() []string {
	return ig.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ig *ig_NG) WeekdayNarrow(weekday time.Weekday) string {
	if len(ig.daysNarrow) == 0 {
		return ""
	}
	return ig.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ig *ig_NG) WeekdaysNarrow() []string {
	return ig.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ig *ig_NG) WeekdayShort(weekday time.Weekday) string {
	if len(ig.daysShort) == 0 {
		return ""
	}
	return ig.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ig *ig_NG) WeekdaysShort() []string {
	return ig.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ig *ig_NG) WeekdayWide(weekday time.Weekday) string {
	if len(ig.daysWide) == 0 {
		return ""
	}
	return ig.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ig *ig_NG) WeekdaysWide() []string {
	return ig.daysWide
}

// Decimal returns the decimal point of number
func (ig *ig_NG) Decimal() string {
	return ig.decimal
}

// Group returns the group of number
func (ig *ig_NG) Group() string {
	return ig.group
}

// Group returns the minus sign of number
func (ig *ig_NG) Minus() string {
	return ig.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ig_NG' and handles both Whole and Real numbers based on 'v'
func (ig *ig_NG) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ig.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ig.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ig.minus) - 1; j >= 0; j-- {
			b = append(b, ig.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ig_NG' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ig *ig_NG) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 12
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ig.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(ig.minus) - 1; j >= 0; j-- {
			b = append(b, ig.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ig.percentSuffix...)

	b = append(b, ig.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ig_NG'
func (ig *ig_NG) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ig.currencies[currency]
	l := len(s) + len(symbol) + 5 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ig.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ig.group[0])
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
		for j := len(ig.minus) - 1; j >= 0; j-- {
			b = append(b, ig.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ig.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ig_NG'
// in accounting notation.
func (ig *ig_NG) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ig.currencies[currency]
	l := len(s) + len(symbol) + 7 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ig.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ig.group[0])
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

		b = append(b, ig.currencyNegativePrefix[0])

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
			b = append(b, ig.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ig.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ig.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ig.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ig.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ig.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ig.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ig.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ig.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ig.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ig.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ig_NG'
func (ig *ig_NG) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ig.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ig.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ig.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
