package sw_KE

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type sw_KE struct {
	locale             string
	pluralsCardinal    []locales.PluralRule
	pluralsOrdinal     []locales.PluralRule
	pluralsRange       []locales.PluralRule
	decimal            string
	group              string
	minus              string
	percent            string
	perMille           string
	timeSeparator      string
	inifinity          string
	currencies         []string // idx = enum of currency code
	monthsAbbreviated  []string
	monthsNarrow       []string
	monthsWide         []string
	daysAbbreviated    []string
	daysNarrow         []string
	daysShort          []string
	daysWide           []string
	periodsAbbreviated []string
	periodsNarrow      []string
	periodsShort       []string
	periodsWide        []string
	erasAbbreviated    []string
	erasNarrow         []string
	erasWide           []string
	timezones          map[string]string
}

// New returns a new instance of translator for the 'sw_KE' locale
func New() locales.Translator {
	return &sw_KE{
		locale:             "sw_KE",
		pluralsCardinal:    []locales.PluralRule{2, 6},
		pluralsOrdinal:     []locales.PluralRule{6},
		pluralsRange:       []locales.PluralRule{2, 6},
		decimal:            ".",
		group:              ",",
		minus:              "-",
		percent:            "%",
		timeSeparator:      ":",
		currencies:         []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		monthsAbbreviated:  []string{"", "Jan", "Feb", "Mac", "Apr", "Mei", "Jun", "Jul", "Ago", "Sep", "Okt", "Nov", "Des"},
		monthsNarrow:       []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:         []string{"", "Januari", "Februari", "Machi", "Aprili", "Mei", "Juni", "Julai", "Agosti", "Septemba", "Oktoba", "Novemba", "Desemba"},
		daysWide:           []string{"Jumapili", "Jumatatu", "Jumanne", "Jumatano", "Alhamisi", "Ijumaa", "Jumamosi"},
		periodsAbbreviated: []string{"", ""},
		periodsNarrow:      []string{"am", "pm"},
		periodsWide:        []string{"", ""},
		erasAbbreviated:    []string{"KK", "BK"},
		erasNarrow:         []string{"", ""},
		erasWide:           []string{"Kabla ya Kristo", "Baada ya Kristo"},
		timezones:          map[string]string{"ACDT": "Saa za Mchana za Australia ya Kati", "ACST": "Saa za Wastani za Australia ya Kati", "ACT": "ACT", "ACWDT": "Saa za Mchana za Magharibi mwa Australia ya Kati", "ACWST": "Saa za Wastani za Magharibi mwa Australia ya Kati", "ADT": "Saa za Mchana za Atlantiki", "ADT Arabia": "Saa za Mchana za Arabiani", "AEDT": "Saa za Mchana za Mashariki mwa Australia", "AEST": "Saa za Wastani za Mashariki mwa Australia", "AFT": "Saa za Afghanistani", "AKDT": "Saa za Mchana za Alaska", "AKST": "Saa za Wastani za Alaska", "AMST": "Saa za Majira ya Joto za Amazon", "AMST Armenia": "Saa za Majira ya Joto za Armenia", "AMT": "Saa za Wastani za Amazon", "AMT Armenia": "Saa za Wastani za Armenia", "ANAST": "Saa za Kiangazi za Anadyr", "ANAT": "Saa za Wastani za Anadyr", "ARST": "Saa za Majira Joto za Ajentina", "ART": "Saa za Wastani za Ajentina", "AST": "Saa za Wastani za Atlantiki", "AST Arabia": "Saa za Wastani za Uarabuni", "AWDT": "Saa za Mchana za Australia Magharibi", "AWST": "Saa za Wastani za Australia Magharibi", "AZST": "Saa za Majira ya Joto za Azabajani", "AZT": "Saa za Wastani za Azabajani", "BDT Bangladesh": "Saa za Majira ya Joto za Bangladeshi", "BNT": "Saa za Brunei Darussalam", "BOT": "Saa za Bolivia", "BRST": "Saa za Majira ya Joto za Brazili", "BRT": "Saa za Wastani za Brazili", "BST Bangladesh": "Saa za Wastani za Bangladeshi", "BT": "Saa za Butani", "CAST": "CAST", "CAT": "Saa za Afrika ya Kati", "CCT": "Saa za Visiwa vya Cocos", "CDT": "Saa za Mchana za Kati", "CHADT": "Saa za Mchana za Chatham", "CHAST": "Saa za Wastani za Chatham", "CHUT": "Saa za Chuuk", "CKT": "Saa za Wastani za Visiwa vya Cook", "CKT DST": "Saa za Majira Nusu ya Joto za Visiwa vya Cook", "CLST": "Saa za Majira ya joto za Chile", "CLT": "Saa za Wastani za Chile", "COST": "Saa za Majira ya Joto za Kolombia", "COT": "Saa za Wastani za Kolombia", "CST": "Saa za Wastani za Kati", "CST China": "Saa za Wastani za Uchina", "CST China DST": "Saa za Mchana za Uchina", "CVST": "Saa za Majira ya Joto za Kepuvede", "CVT": "Saa za Wastani za Kepuvede", "CXT": "Saa za Kisiwa cha Krismasi", "ChST": "Saa za Wastani za Chamorro", "ChST NMI": "ChST NMI", "CuDT": "Saa za Mchana za Kuba", "CuST": "Saa za Wastani za Kuba", "DAVT": "Saa za Davis", "DDUT": "Saa za Dumont-d’Urville", "EASST": "Saa za Majira ya Joto za Kisiwa cha Easter", "EAST": "Saa za Wastani za Kisiwa cha Easter", "EAT": "Saa za Afrika Mashariki", "ECT": "Saa za Ekwado", "EDT": "Saa za Mchana za Mashariki", "EGDT": "Saa za Majira ya Joto za Greenland Mashariki", "EGST": "Saa za Wastani za Greenland Mashariki", "EST": "Saa za Wastani za Mashariki", "FEET": "Saa za Mashariki zaidi mwa Ulaya", "FJT": "Saa za Wastani za Fiji", "FJT Summer": "Saa za Majira ya joto za Fiji", "FKST": "Saa za Majira ya joto za Visiwa vya Falkland", "FKT": "Saa za Wastani za Visiwa vya Falkland", "FNST": "Saa za Majira ya joto za Fernando de Noronha", "FNT": "Saa za Wastani za Fernando de Noronha", "GALT": "Saa za Galapagos", "GAMT": "Saa za Gambier", "GEST": "Saa za Majira ya Joto za Jiojia", "GET": "Saa za Wastani za Jiojia", "GFT": "Saa za Guiana", "GIT": "Saa za Visiwa vya Gilbert", "GMT": "Saa za Greenwich", "GNSST": "GNSST", "GNST": "GNST", "GST": "Saa za Wastani za Ghuba", "GST Guam": "GST Guam", "GYT": "Saa za Guyana", "HADT": "Saa za Wastani za Hawaii-Aleutian", "HAST": "Saa za Wastani za Hawaii-Aleutian", "HKST": "Saa za Majira ya Joto za Hong Kong", "HKT": "Saa za Wastani za Hong Kong", "HOVST": "Saa za Majira ya Joto za Hovd", "HOVT": "Saa za Wastani za Hovd", "ICT": "Saa za Indochina", "IDT": "Saa za Mchana za Israeli", "IOT": "Saa za Bahari Hindi", "IRKST": "Saa za Majira ya Joto za Irkutsk", "IRKT": "Saa za Wastani za Irkutsk", "IRST": "Saa za Wastani za Irani", "IRST DST": "Saa za Mchana za Irani", "IST": "Saa za Wastani za India", "IST Israel": "Saa za Wastani za Israeli", "JDT": "Saa za Mchana za Japani", "JST": "Saa za Wastani za Japani", "KOST": "Saa za Kosrae", "KRAST": "Saa za Majira ya Joto za Krasnoyarsk", "KRAT": "Saa za Wastani za Krasnoyask", "KST": "Saa za Wastani za Korea", "KST DST": "Saa za Mchana za Korea", "LHDT": "Saa za Mchana za Lord Howe", "LHST": "Saa za Wastani za Lord Howe", "LINT": "Saa za Visiwa vya Line", "MAGST": "Saa za Majira ya Joto za Magadan", "MAGT": "Saa za Wastani za Magadan", "MART": "Saa za Marquesas", "MAWT": "Saa za Mawson", "MDT": "MDT", "MESZ": "Saa za Majira ya Joto za Ulaya ya Kati", "MEZ": "Saa za Wastani za Ulaya ya Kati", "MHT": "Saa za Visiwa vya Marshall", "MMT": "Saa za Myanma", "MSD": "Saa za Majira ya Joto za Moscow", "MST": "MST", "MUST": "Saa za Majira ya Joto za Morisi", "MUT": "Saa za Wastani za Morisi", "MVT": "Saa za Maldivi", "MYT": "Saa za Malesia", "NCT": "Saa za Wastani za Kaledonia Mpya", "NDT": "Saa za Mchana za Newfoundland", "NDT New Caledonia": "Saa za Majira ya Joto za Kaledonia Mpya", "NFDT": "Saa za Majira ya Joto za Kisiwa cha Norfolk", "NFT": "Saa za Wastani za Kisiwa cha Norfolk", "NOVST": "Saa za Majira ya Joto za Novosibirsk", "NOVT": "Saa za Wastani za Novosibirsk", "NPT": "Saa za Nepali", "NRT": "Saa za Nauru", "NST": "Saa za Wastani za Newfoundland", "NUT": "Saa za Niue", "NZDT": "Saa za Mchana za Nyuzilandi", "NZST": "Saa za Wastani za Nyuzilandi", "OESZ": "Saa za Majira ya Joto za Mashariki mwa Ulaya", "OEZ": "Saa za Wastani za Mashariki mwa Ulaya", "OMSST": "Saa za Majira ya Joto za Omsk", "OMST": "Saa za Wastani za Omsk", "PDT": "Saa za Mchana za Pasifiki", "PDTM": "Saa za mchana za pasifiki za Meksiko", "PETDT": "Saa za Kiangazi za Petropavlovsk-Kamchatski", "PETST": "Saa za Wastani za Petropavlovsk-Kamchatski", "PGT": "Saa za Papua", "PHOT": "Saa za Visiwa vya Finiksi", "PKT": "Saa za Wastani za Pakistani", "PKT DST": "Saa za Majira ya Joto za Pakistani", "PMDT": "Saa za Mchana za Saint-Pierre na Miquelon", "PMST": "Saa za Wastani ya Saint-Pierre na Miquelon", "PONT": "Saa za Ponape", "PST": "Saa za Wastani za Pasifiki", "PST Philippine": "Saa za Wastani za Ufilipino", "PST Philippine DST": "Saa za Majira ya Joto za Ufilipino", "PST Pitcairn": "Saa za Pitcairn", "PSTM": "Saa za wastani za pasifiki za Meksiko", "PWT": "Saa za Palau", "PYST": "Saa za Majira ya Joto za Paragwai", "PYT": "Saa za Wastani za Paragwai", "PYT Korea": "Saa za Pyongyang", "RET": "Saa za Reunion", "ROTT": "Saa za Rothera", "SAKST": "Saa za Majira ya Joto za Sakhalin", "SAKT": "Saa za Wastani za Sakhalin", "SAMST": "Saa za Kiangazi za Samara", "SAMT": "Saa za Wastani za Samara", "SAST": "Saa za Wastani za Afrika Kusini", "SBT": "Saa za Visiwa vya Solomon", "SCT": "Saa za Ushelisheli", "SGT": "Saa za Wastani za Singapoo", "SLST": "SLST", "SRT": "Saa za Suriname", "SST Samoa": "Saa za Wastani za Samoa", "SST Samoa Apia": "Saa za Wastani za Apia", "SST Samoa Apia DST": "Saa za Mchana za Apia", "SST Samoa DST": "Saa za Mchana za Samoa", "SYOT": "Saa za Syowa", "TAAF": "Saa za Kusini mwa Ufaransa na Antaktiki", "TAHT": "Saa za Tahiti", "TJT": "Saaza Tajikistani", "TKT": "Saa za Tokelau", "TLT": "Saa za Timor Mashariki", "TMST": "Saa za Majira ya Joto za Turkmenistani", "TMT": "Saa za Wastani za Turkmenistani", "TOST": "Saa za Majira ya Joto za Tonga", "TOT": "Saa za Wastani za Tonga", "TVT": "Saa za Tuvalu", "TWT": "Saa za Wastani za Taipei", "TWT DST": "Saa za Mchana za Taipei", "ULAST": "Saa za Majira ya Joto za Ulaanbaatar", "ULAT": "Saa za Wastani za Ulaanbataar", "UYST": "Saa za Majira ya Joto za Urugwai", "UYT": "Saa za Wastani za Urugwai", "UZT": "Saa za wastani za Uzbekistani", "UZT DST": "Saa za Majira ya Joto za Uzbekistani", "VET": "Saa za Venezuela", "VLAST": "Saa za Majira ya Joto za Vladivostok", "VLAT": "Saa za Wastani za Vladivostok", "VOLST": "Saa za Majira ya Joto za Volgograd", "VOLT": "Saa za Wastani za Volgograd", "VOST": "Saa za Vostok", "VUT": "Saa za Wastani za Vanuatu", "VUT DST": "Saa za Majira ya Joto za Vanuatu", "WAKT": "Saa za Kisiwa cha Wake", "WARST": "Saa za Majira ya Joto za Magharibi mwa Ajentina", "WART": "Saa za Wastani za Magharibi mwa Ajentina", "WAST": "Saa za Afrika Magharibi", "WAT": "Saa za Afrika Magharibi", "WESZ": "Saa za Majira ya Joto za Magharibi mwa Ulaya", "WEZ": "Saa za Wastani za Magharibi mwa Ulaya", "WFT": "Saa za Wallis na Futuna", "WGST": "Saa za Majira ya joto za Greenland Magharibi", "WGT": "Saa za Wastani za Greenland Magharibi", "WIB": "Saa za Magharibi mwa Indonesia", "WIT": "Saa za Mashariki mwa Indonesia", "WITA": "Saa za Indonesia ya Kati", "YAKST": "Saa za Majira ya Joto za Yakutsk", "YAKT": "Saa za Wastani za Yakutsk", "YEKST": "Saa za Majira ya Joto za Yekaterinburg", "YEKT": "Saa za Wastani za Yekaterinburg", "YST": "Saa za Yukon", "МСК": "Saa za Wastani za Moscow", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Saa za Kazakistani Magharibi", "شىعىش قازاق ەلى": "Saa za Kazakistani Mashariki", "قازاق ەلى": "Saa za Kazakhstan", "قىرعىزستان": "Saa za Kyrgystan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Saa za Majira ya Joto za Peru"},
	}
}

// Locale returns the current translators string locale
func (sw *sw_KE) Locale() string {
	return sw.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sw_KE'
func (sw *sw_KE) PluralsCardinal() []locales.PluralRule {
	return sw.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sw_KE'
func (sw *sw_KE) PluralsOrdinal() []locales.PluralRule {
	return sw.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'sw_KE'
func (sw *sw_KE) PluralsRange() []locales.PluralRule {
	return sw.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sw_KE'
func (sw *sw_KE) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sw_KE'
func (sw *sw_KE) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sw_KE'
func (sw *sw_KE) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := sw.CardinalPluralRule(num1, v1)
	end := sw.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (sw *sw_KE) MonthAbbreviated(month time.Month) string {
	return sw.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (sw *sw_KE) MonthsAbbreviated() []string {
	return sw.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (sw *sw_KE) MonthNarrow(month time.Month) string {
	return sw.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (sw *sw_KE) MonthsNarrow() []string {
	return sw.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (sw *sw_KE) MonthWide(month time.Month) string {
	return sw.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (sw *sw_KE) MonthsWide() []string {
	return sw.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (sw *sw_KE) WeekdayAbbreviated(weekday time.Weekday) string {
	return sw.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (sw *sw_KE) WeekdaysAbbreviated() []string {
	return sw.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (sw *sw_KE) WeekdayNarrow(weekday time.Weekday) string {
	return sw.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (sw *sw_KE) WeekdaysNarrow() []string {
	return sw.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (sw *sw_KE) WeekdayShort(weekday time.Weekday) string {
	return sw.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (sw *sw_KE) WeekdaysShort() []string {
	return sw.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (sw *sw_KE) WeekdayWide(weekday time.Weekday) string {
	return sw.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (sw *sw_KE) WeekdaysWide() []string {
	return sw.daysWide
}

// Decimal returns the decimal point of number
func (sw *sw_KE) Decimal() string {
	return sw.decimal
}

// Group returns the group of number
func (sw *sw_KE) Group() string {
	return sw.group
}

// Group returns the minus sign of number
func (sw *sw_KE) Minus() string {
	return sw.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sw_KE' and handles both Whole and Real numbers based on 'v'
func (sw *sw_KE) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sw.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sw.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sw.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'sw_KE' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sw *sw_KE) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sw.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sw.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, sw.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sw_KE'
func (sw *sw_KE) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sw.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sw.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sw.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sw.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sw.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sw_KE'
// in accounting notation.
func (sw *sw_KE) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sw.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sw.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sw.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, sw.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sw.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, symbol...)
	} else {

		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'sw_KE'
func (sw *sw_KE) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'sw_KE'
func (sw *sw_KE) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sw.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'sw_KE'
func (sw *sw_KE) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sw.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'sw_KE'
func (sw *sw_KE) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, sw.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sw.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'sw_KE'
func (sw *sw_KE) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sw.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'sw_KE'
func (sw *sw_KE) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sw.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sw.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'sw_KE'
func (sw *sw_KE) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sw.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sw.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'sw_KE'
func (sw *sw_KE) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sw.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sw.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := sw.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
