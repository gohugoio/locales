package lld_IT

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type lld_IT struct {
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

// New returns a new instance of translator for the 'lld_IT' locale
func New() locales.Translator {
	return &lld_IT{
		locale:                 "lld_IT",
		pluralsCardinal:        []locales.PluralRule{2, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{5, 6},
		pluralsRange:           nil,
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "jená", "forá", "merz", "aurí", "ma", "jügn", "messé", "aost", "set", "oto", "nov", "dez"},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "M", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "de jená", "de forá", "de merz", "d’aurí", "de ma", "de jügn", "de messé", "d’aost", "de setëmber", "d’otober", "de novëmber", "de dezëmber"},
		daysAbbreviated:        []string{"dom", "lön", "mert", "merc", "jöb", "vën", "sab"},
		daysNarrow:             []string{"D", "L", "M", "M", "J", "V", "S"},
		daysWide:               []string{"domënia", "lönesc", "mertesc", "mercui", "jöbia", "vëndres", "sabeda"},
		timezones:              map[string]string{"ACDT": "Ora da d’isté dl’Australia zentrala", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Ora da d’isté dl’Australia zënter-ozidentala", "ACWST": "Ora standard dl’Australia zënter-ozidentala", "ADT": "Ora da d’isté dl Atlantich", "ADT Arabia": "Ora da d’isté dl’Arabia", "AEDT": "Ora da d’isté dl’Australia orientala", "AEST": "Ora standard dl’Australia orientala", "AFT": "Ora dl Afghanistan", "AKDT": "Ora da d’isté dl’Alaska", "AKST": "Ora standard dl’Alaska", "AMST": "Ora da d’isté dl’Amazonia", "AMST Armenia": "Ora da d’isté dl’Armenia", "AMT": "Ora standard dl’Amazonia", "AMT Armenia": "Ora standard dl’Armenia", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Ora da d’isté dl’Argentina", "ART": "Ora standard dl’Argentina", "AST": "Ora standard dl Atlantich", "AST Arabia": "Ora standard dl’Arabia", "AWDT": "Ora da d’isté dl’Australia ozidentala", "AWST": "Ora standard dl’Australia ozidentala", "AZST": "Ora da d’isté dl Azerbaijan", "AZT": "Ora standard dl Azerbaijan", "BDT Bangladesh": "Ora da d’isté dl Bangladesc", "BNT": "Ora dl Brunei Darussalam", "BOT": "Ora dla Bolivia", "BRST": "Ora da d’isté de Brasilia", "BRT": "Ora standard de Brasilia", "BST Bangladesh": "Ora standard dl Bangladesc", "BT": "Ora dl Bhutan", "CAST": "CAST", "CAT": "Ora dl’Africa zentrala", "CCT": "Ora dles Isoles Cocos", "CDT": "Ora da d’isté dl’America dl Nord zentrala", "CHADT": "Ora da d’isté dles Isoles Chatham", "CHAST": "Ora standard dles Isoles Chatham", "CHUT": "Ora dl Chuuk", "CKT": "Ora standard dles Isoles Cook", "CKT DST": "Ora da d’isté mesana dles Isoles Cook", "CLST": "Ora da d’isté dl Cile", "CLT": "Ora standard dl Cile", "COST": "Ora da d’isté dla Colombia", "COT": "Ora standard dla Colombia", "CST": "Ora standard dl’America dl Nord zentrala", "CST China": "Ora standard dla Cina", "CST China DST": "Ora da d’isté dla Cina", "CVST": "Ora da d’isté de Capo Verde", "CVT": "Ora standard de Capo Verde", "CXT": "Ora dla Isola dl Nadé", "ChST": "Ora standard de Chamorro", "ChST NMI": "ChST NMI", "CuDT": "Ora da d’isté de Cuba", "CuST": "Ora standard de Cuba", "DAVT": "Ora de Davis", "DDUT": "Ora de Dumont-d’Urville", "EASST": "Ora da d’isté dl’Isola de Pasca", "EAST": "Ora standard dl’Isola de Pasca", "EAT": "Ora dl’Africa orientala", "ECT": "Ora dl Ecuador", "EDT": "Ora da d’isté dl’America dl Nord orientala", "EGDT": "Ora da d’isté dla Groenlandia orientala", "EGST": "Ora standard dla Groenlandia orientala", "EST": "Ora standard dl’America dl Nord orientala", "FEET": "Ora de Kaliningrad", "FJT": "Ora standard dles Fiji", "FJT Summer": "Ora da d’isté dles Fiji", "FKST": "Ora da d’isté dles Isoles Falkland", "FKT": "Ora standard dles Isoles Falkland", "FNST": "Ora da d’isté de Fernando de Noronha", "FNT": "Ora standard de Fernando de Noronha", "GALT": "Ora dles Galapagos", "GAMT": "Ora dles Isoles Gambier", "GEST": "Ora da d’isté dla Georgia", "GET": "Ora standard dla Georgia", "GFT": "Ora dla Guyana franzeja", "GIT": "Ora dles Isoles Gilbert", "GMT": "Ora mesana de Greenwich", "GNSST": "GNSST", "GNST": "GNST", "GST": "Ora dla Georgia dl Süd", "GST Guam": "GST Guam", "GYT": "Ora dla Guyana", "HADT": "Ora standard dles Isoles Hawaii-Aleutines", "HAST": "Ora standard dles Isoles Hawaii-Aleutines", "HKST": "Ora da d’isté de Hong Kong", "HKT": "Ora standard de Hong Kong", "HOVST": "Ora da d’isté de Hovd", "HOVT": "Ora standard de Hovd", "ICT": "Ora dl’Indocina", "IDT": "Ora da d’isté d’Israel", "IOT": "Ora dl Ozean indian", "IRKST": "Ora da d’isté de Irkutsk", "IRKT": "Ora standard de Irkutsk", "IRST": "Ora standard dl Iran", "IRST DST": "Ora da d’isté dl Iran", "IST": "Ora standard dl’India", "IST Israel": "Ora standard d’Israel", "JDT": "Ora da d’isté dl Iapan", "JST": "Ora standard dl Iapan", "KOST": "Ora dl’Isola Kosrae", "KRAST": "Ora da d’isté de Krasnoyarsk", "KRAT": "Ora standard de Krasnoyarsk", "KST": "Ora standard dla Corea", "KST DST": "Ora da d’isté dla Corea", "LHDT": "Ora da d’isté de Lord Howe", "LHST": "Ora standard de Lord Howe", "LINT": "Ora dles Isoles Line", "MAGST": "Ora da d’isté de Magadan", "MAGT": "Ora standard de Magadan", "MART": "Ora dles Isoles Marchejes", "MAWT": "Ora de Mawson", "MDT": "MDT", "MESZ": "Ora da d’isté dl’Europa zentrala", "MEZ": "Ora standard dl’Europa zentrala", "MHT": "Ore dles Isoles Marshall", "MMT": "Ora dl Myanmar", "MSD": "Ora da d’isté de Mosca", "MST": "MST", "MUST": "Ora da d’isté de Mauritius", "MUT": "Ora standard de Mauritius", "MVT": "Ora dles Maldives", "MYT": "Ora dla Malesia", "NCT": "Ora standard dla Nöia Caledonia", "NDT": "Ora da d’isté de Newfoundland", "NDT New Caledonia": "Ora da d’isté dla Nöia Caledonia", "NFDT": "Ora da d’isté dl’Isola Norfolk", "NFT": "Ora standard dl’Isola Norfolk", "NOVST": "Ora da d’isté de Novosibirsk", "NOVT": "Ora standard de Novosibirsk", "NPT": "Ora dl Nepal", "NRT": "Ora de Nauru", "NST": "Ora standard de Newfoundland", "NUT": "Ora de Niue", "NZDT": "Ora da d’isté dla Nöia Zelanda", "NZST": "Ora standard dla Nöia Zelanda", "OESZ": "Ora da d’isté dl’Europa orientala", "OEZ": "Ora standard dl’Europa orientala", "OMSST": "Ora da d’isté de Omsk", "OMST": "Ora standard de Omsk", "PDT": "Ora da d’isté dl’America dl Nord ozidentala", "PDTM": "Ora da d’isté dla spona pazifica dl Messich", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Ora dla Papua Nöia Guinea", "PHOT": "Ora dles Isoles Phoenix", "PKT": "Ora standard dl Pakistan", "PKT DST": "Ora da d’isté dl Pakistan", "PMDT": "Ora da d’isté de St. Pierre y Miquelon", "PMST": "Ora standard de St. Pierre y Miquelon", "PONT": "Ora de Ponape", "PST": "Ora standard dl’America dl Nord ozidentala", "PST Philippine": "Ora standard dles Filipines", "PST Philippine DST": "Ora da d’isté dles Filipines", "PST Pitcairn": "Ora dles Isoles Pitcairn", "PSTM": "Ora standard dla spona pazifica dl Messich", "PWT": "Ora de Palau", "PYST": "Ora da d’isté dl Paraguay", "PYT": "Ora standard dl Paraguay", "PYT Korea": "Ora de Pjöngjang", "RET": "Ora de Réunion", "ROTT": "Ora de Rothera", "SAKST": "Ora da d’isté de Sachalin", "SAKT": "Ora standard de Sachalin", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Ora dl’Africa dl Süd", "SBT": "Ora dles Isoles Salomon", "SCT": "Ora dles Seychelles", "SGT": "Ora standard de Singapur", "SLST": "SLST", "SRT": "Ora dl Surinam", "SST Samoa": "Ora standard de Samoa", "SST Samoa Apia": "Ora standard de Apia", "SST Samoa Apia DST": "Ora da d’isté de Apia", "SST Samoa DST": "Ora da d’isté de Samoa", "SYOT": "Ora de Syowa", "TAAF": "Ora di Teritori franzesc dl Süd y dl’Antartica", "TAHT": "Ora de Tahiti", "TJT": "Ora dl Tajikistan", "TKT": "Ora de Tokelau", "TLT": "Ora de Timor Ost", "TMST": "Ora da d’isté dl Turkmenistan", "TMT": "Ora standard dl Turkmenistan", "TOST": "Ora da d’isté de Tonga", "TOT": "Ora standard de Tonga", "TVT": "Ora de Tuvalu", "TWT": "Ora standard de Taipei", "TWT DST": "Ora da d’isté de Taipei", "ULAST": "Ora da d’isté de Ulan Bator", "ULAT": "Ora standard de Ulan Bator", "UYST": "Ora da d’isté dl Uruguay", "UYT": "Ora standard dl Uruguay", "UZT": "Ora standard dl Uzbekistan", "UZT DST": "Ora da d’isté dl Uzbekistan", "VET": "Ora dl Venezuela", "VLAST": "Ora da d’isté de Vladivostok", "VLAT": "Ora standard de Vladivostok", "VOLST": "Ora da d’isté de Volgograd", "VOLT": "Ora standard de Volgograd", "VOST": "Ora de Vostok", "VUT": "Ora standard dl Vanuatu", "VUT DST": "Ora da d’isté dl Vanuatu", "WAKT": "Ora dl’Isola de Wake", "WARST": "Ora da d’isté dl’Argentina ozidentala", "WART": "Ora standard dl’Argentina ozidentala", "WAST": "Ora dl’Africa ozidentala", "WAT": "Ora dl’Africa ozidentala", "WESZ": "Ora da d’isté dl’Europa ozidentala", "WEZ": "Ora standard dl’Europa ozidentala", "WFT": "Ora de Wallis y Futuna", "WGST": "Ora da d’isté dla Groenlandia ozidentala", "WGT": "Ora standard dla Groenlandia ozidentala", "WIB": "Ora dl’Indonesia ozidentala", "WIT": "Ora dl’Indonesia orientala", "WITA": "Ora dl’Indonesia zentrala", "YAKST": "Ora da d’isté de Iakutsk", "YAKT": "Ora standard de Iakutsk", "YEKST": "Ora da d’isté de Yekaterinburg", "YEKT": "Ora standard de Yekaterinburg", "YST": "Ora dl Yukon", "МСК": "Ora standard de Mosca", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Ora dl Kazakhstan ozidental", "شىعىش قازاق ەلى": "Ora dl Kazakhstan oriental", "قازاق ەلى": "Ora dl Kazakhstan", "قىرعىزستان": "Ora dl Kyrgystan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Ora da d’isté dles Azores"},
	}
}

// Locale returns the current translators string locale
func (lld *lld_IT) Locale() string {
	return lld.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'lld_IT'
func (lld *lld_IT) PluralsCardinal() []locales.PluralRule {
	return lld.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'lld_IT'
func (lld *lld_IT) PluralsOrdinal() []locales.PluralRule {
	return lld.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'lld_IT'
func (lld *lld_IT) PluralsRange() []locales.PluralRule {
	return lld.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'lld_IT'
func (lld *lld_IT) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	e := int64(0)
	iMod1000000 := i % 1000000

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	} else if (e == 0 && i != 0 && iMod1000000 == 0 && v == 0) || (e < 0 || e > 5) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'lld_IT'
func (lld *lld_IT) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 11 || n == 8 || n == 80 || n == 800 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'lld_IT'
func (lld *lld_IT) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (lld *lld_IT) MonthAbbreviated(month time.Month) string {
	if len(lld.monthsAbbreviated) == 0 {
		return ""
	}
	return lld.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (lld *lld_IT) MonthsAbbreviated() []string {
	return lld.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (lld *lld_IT) MonthNarrow(month time.Month) string {
	if len(lld.monthsNarrow) == 0 {
		return ""
	}
	return lld.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (lld *lld_IT) MonthsNarrow() []string {
	return lld.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (lld *lld_IT) MonthWide(month time.Month) string {
	if len(lld.monthsWide) == 0 {
		return ""
	}
	return lld.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (lld *lld_IT) MonthsWide() []string {
	return lld.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (lld *lld_IT) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(lld.daysAbbreviated) == 0 {
		return ""
	}
	return lld.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (lld *lld_IT) WeekdaysAbbreviated() []string {
	return lld.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (lld *lld_IT) WeekdayNarrow(weekday time.Weekday) string {
	if len(lld.daysNarrow) == 0 {
		return ""
	}
	return lld.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (lld *lld_IT) WeekdaysNarrow() []string {
	return lld.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (lld *lld_IT) WeekdayShort(weekday time.Weekday) string {
	if len(lld.daysShort) == 0 {
		return ""
	}
	return lld.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (lld *lld_IT) WeekdaysShort() []string {
	return lld.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (lld *lld_IT) WeekdayWide(weekday time.Weekday) string {
	if len(lld.daysWide) == 0 {
		return ""
	}
	return lld.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (lld *lld_IT) WeekdaysWide() []string {
	return lld.daysWide
}

// Decimal returns the decimal point of number
func (lld *lld_IT) Decimal() string {
	return lld.decimal
}

// Group returns the group of number
func (lld *lld_IT) Group() string {
	return lld.group
}

// Group returns the minus sign of number
func (lld *lld_IT) Minus() string {
	return lld.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'lld_IT' and handles both Whole and Real numbers based on 'v'
func (lld *lld_IT) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lld.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, lld.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, lld.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'lld_IT' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (lld *lld_IT) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lld.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, lld.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, lld.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'lld_IT'
func (lld *lld_IT) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := lld.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lld.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, lld.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, lld.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, lld.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, lld.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'lld_IT'
// in accounting notation.
func (lld *lld_IT) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := lld.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lld.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, lld.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, lld.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, lld.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, lld.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, lld.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'lld_IT'
func (lld *lld_IT) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'lld_IT'
func (lld *lld_IT) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, lld.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'lld_IT'
func (lld *lld_IT) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, lld.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'lld_IT'
func (lld *lld_IT) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, lld.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, lld.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x6c}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'lld_IT'
func (lld *lld_IT) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lld.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'lld_IT'
func (lld *lld_IT) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lld.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lld.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'lld_IT'
func (lld *lld_IT) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lld.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lld.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'lld_IT'
func (lld *lld_IT) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lld.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lld.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := lld.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
