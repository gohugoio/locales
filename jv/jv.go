package jv

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type jv struct {
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
	currencyPositivePrefix string
	currencyNegativePrefix string
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

// New returns a new instance of translator for the 'jv' locale
func New() locales.Translator {
	return &jv{
		locale:                 "jv",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  ".",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "Rp", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Agt", "Sep", "Okt", "Nov", "Des"},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"},
		daysAbbreviated:        []string{"Ahad", "Sen", "Sel", "Rab", "Kam", "Jum", "Sab"},
		daysNarrow:             []string{"A", "S", "S", "R", "K", "J", "S"},
		daysWide:               []string{"Ahad", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"},
		periodsAbbreviated:     []string{"Isuk", "Wengi"},
		erasAbbreviated:        []string{"SM", "M"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"Sakdurunge Masehi", "Masehi"},
		timezones:              map[string]string{"ACDT": "Wektu Ketigo Australia Tengah", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Wektu Ketigo Australia Tengah sisih Kulon", "ACWST": "Wektu Standar Australia Tengah sisih Kulon", "ADT": "Wektu Ketigo Atlantik", "ADT Arabia": "Wektu Ketigo Arab", "AEDT": "Wektu Ketigo Australia sisih Wetan", "AEST": "Wektu Standar Australia sisih Wetan", "AFT": "Wektu Afghanistan", "AKDT": "Wektu Ketigo Alaska", "AKST": "Wektu Standar Alaska", "AMST": "Wektu Ketigo Amazon", "AMST Armenia": "Wektu Ketigo Armenia", "AMT": "Wektu Standar Amazon", "AMT Armenia": "Wektu Standar Armenia", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Wektu Ketigo Argentina", "ART": "Wektu Standar Argentina", "AST": "Wektu Standar Atlantik", "AST Arabia": "Wektu Standar Arab", "AWDT": "Wektu Ketigo Australia sisih Kulon", "AWST": "Wektu Standar Australia sisih Kulon", "AZST": "Wektu Ketigo Azerbaijan", "AZT": "Wektu Standar Azerbaijan", "BDT Bangladesh": "Wektu Ketigo Bangladesh", "BNT": "Wektu Brunei", "BOT": "Wektu Bolivia", "BRST": "Wektu Ketigo Brasilia", "BRT": "Wektu Standar Brasilia", "BST Bangladesh": "Wektu Standar Bangladesh", "BT": "Wektu Bhutan", "CAST": "CAST", "CAT": "Wektu Afrika Tengah", "CCT": "Wektu Kepuloan Cocos", "CDT": "Wektu Ketigo Tengah", "CHADT": "Wektu Ketigo Chatham", "CHAST": "Wektu Standar Chatham", "CHUT": "Wektu Chuuk", "CKT": "Wektu Standar Kepuloan Cook", "CKT DST": "Wektu Ketiga Kepuloan Cook", "CLST": "Wektu Ketigo Chili", "CLT": "Wektu Standar Chili", "COST": "Wektu Ketigo Kolombia", "COT": "Wektu Standar Kolombia", "CST": "Wektu Standar Tengah", "CST China": "Wektu Standar Cina", "CST China DST": "Wektu Ketiga Cina", "CVST": "Wektu Ketigo Tanjung Verde", "CVT": "Wektu Standar Tanjung Verde", "CXT": "Wektu Pulo Natal", "ChST": "Wektu Chamorro", "ChST NMI": "ChST NMI", "CuDT": "Wektu Ketigo Kuba", "CuST": "Wektu Standar Kuba", "DAVT": "Wektu Davis", "DDUT": "Wektu Dumont d’Urville", "EASST": "Wektu Ketigo Pulo Paskah", "EAST": "Wektu Standar Pulo Paskah", "EAT": "Wektu Afrika Wetan", "ECT": "Wektu Ekuador", "EDT": "Wektu Ketigo sisih Wetah", "EGDT": "Wektu Ketigo Grinland Wetan", "EGST": "Wektu Standar Grinland Wetan", "EST": "Wektu Standar sisih Wetan", "FEET": "Wektu Eropa sisih Wetan seng Luwih Adoh", "FJT": "Wektu Standar Fiji", "FJT Summer": "Wektu Ketigo Fiji", "FKST": "Wektu Ketigo Kepuloan Falkland", "FKT": "Wektu Standar Kepuloan Falkland", "FNST": "Wektu Ketigo Fernando de Noronha", "FNT": "Wektu Standar Fernando de Noronha", "GALT": "Wektu Galapagos", "GAMT": "Wektu Gambier", "GEST": "Wektu Ketigo Georgia", "GET": "Wektu Standar Georgia", "GFT": "Wektu Guiana Prancis", "GIT": "Wektu Kepuloan Gilbert", "GMT": "Wektu Rerata Greenwich", "GNSST": "GNSST", "GNST": "GNST", "GST": "Wektu Georgia Kidul", "GST Guam": "GST Guam", "GYT": "Wektu Guyana", "HADT": "Wektu Ketigo Hawaii-Aleutian", "HAST": "Wektu Standar Hawaii-Aleutian", "HKST": "Wektu Ketiga Hong Kong", "HKT": "Wektu Standar Hong Kong", "HOVST": "Wektu Ketiga Hovd", "HOVT": "Wektu Standar Hovd", "ICT": "Wektu Indocina", "IDT": "Wektu Ketigo Israel", "IOT": "Wektu Segoro Hindia", "IRKST": "Wektu Ketigo Irkutsk", "IRKT": "Wektu Standar Irkutsk", "IRST": "Wektu Standar Iran", "IRST DST": "Wektu Ketigo Iran", "IST": "Wektu Standar India", "IST Israel": "Wektu Standar Israel", "JDT": "Wektu Ketiga Jepang", "JST": "Wektu Standar Jepang", "KOST": "Wektu Kosrae", "KRAST": "Wektu Ketigo Krasnoyarsk", "KRAT": "Wektu Standar Krasnoyarsk", "KST": "Wektu Standar Korea", "KST DST": "Wektu Ketiga Korea", "LHDT": "Wektu Ketigo Lord Howe", "LHST": "Wektu Standar Lord Howe", "LINT": "Wektu Kepuloan Line", "MAGST": "Wektu Ketigo Magadan", "MAGT": "Wektu Standar Magadan", "MART": "Wektu Marquesas", "MAWT": "Wektu Mawson", "MDT": "MDT", "MESZ": "Wektu Ketigo Eropa Tengah", "MEZ": "Wektu Standar Eropa Tengah", "MHT": "Wektu Kepuloan Marshall", "MMT": "Wektu Myanmar", "MSD": "Wektu Ketigo Moscow", "MST": "MST", "MUST": "Wektu Ketigo Mauritius", "MUT": "Wektu Standar Mauritius", "MVT": "Wektu Maladewa", "MYT": "Wektu Malaysia", "NCT": "Wektu Standar Kaledonia Anyar", "NDT": "Wektu Ketigo Newfoundland", "NDT New Caledonia": "Wektu Ketigo Kaledonia Anyar", "NFDT": "Wektu Ketigo Pulo Norfolk", "NFT": "Wektu Standar Pulo Norfolk", "NOVST": "Wektu Ketigo Novosibirsk", "NOVT": "Wektu Standar Novosibirsk", "NPT": "Wektu Nepal", "NRT": "Wektu Nauru", "NST": "Wektu Standar Newfoundland", "NUT": "Wektu Niue", "NZDT": "Wektu Ketigo Selandia Anyar", "NZST": "Wektu Standar Selandia Anyar", "OESZ": "Wektu Ketigo Eropa sisih Wetan", "OEZ": "Wektu Standar Eropa sisih Wetan", "OMSST": "Wektu Ketigo Omsk", "OMST": "Wektu Standar Omsk", "PDT": "Wektu Ketigo Pasifik", "PDTM": "Wektu Ketigo Pasifik Meksiko", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Wektu Papua Nugini", "PHOT": "Wektu Kepuloan Phoenix", "PKT": "Wektu Standar Pakistan", "PKT DST": "Wektu Ketigo Pakistan", "PMDT": "Wektu Ketigo Santa Pierre lan Miquelon", "PMST": "Wektu Standar Santa Pierre lan Miquelon", "PONT": "Wektu Pohnpei", "PST": "Wektu Standar Pasifik", "PST Philippine": "Wektu Standar Filipina", "PST Philippine DST": "Wektu Ketigo Filipina", "PST Pitcairn": "Wektu Pitcairn", "PSTM": "Wektu Standar Pasifik Meksiko", "PWT": "Wektu Palau", "PYST": "Wektu Ketigo Paraguay", "PYT": "Wektu Standar Paraguay", "PYT Korea": "Wektu Korea Lor", "RET": "Wektu Reunion", "ROTT": "Wektu Rothera", "SAKST": "Wektu Ketigo Sakhalin", "SAKT": "Wektu Standar Sakhalin", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Wektu Standar Afrika Kidul", "SBT": "Wektu Kepuloan Solomon", "SCT": "Wektu Seichelles", "SGT": "Wektu Singapura", "SLST": "SLST", "SRT": "Wektu Suriname", "SST Samoa": "Wektu Standar Samoa Amerika", "SST Samoa Apia": "Wektu Standar Samoa", "SST Samoa Apia DST": "Wektu Ketiga Samoa", "SST Samoa DST": "Wektu Ketiga Samoa Amerika", "SYOT": "Wektu Syowa", "TAAF": "Wektu Antartika lan Prancis sisih Kidul", "TAHT": "Wektu Tahiti", "TJT": "Wektu Tajikistan", "TKT": "Wektu Tokelau", "TLT": "Wektu Timor-Leste", "TMST": "Wektu Ketigo Turkmenistan", "TMT": "Wektu Standar Turkmenistan", "TOST": "Wektu Ketigo Tonga", "TOT": "Wektu Standar Tonga", "TVT": "Wektu Tuvalu", "TWT": "Wektu Standar Taiwan", "TWT DST": "Wektu Awan Taiwan", "ULAST": "Wektu Ketiga Ulaanbaatar", "ULAT": "Wektu Standar Ulaanbaatar", "UYST": "Wektu Ketigo Uruguay", "UYT": "Wektu Standar Uruguay", "UZT": "Wektu Standar Usbekistan", "UZT DST": "Wektu Ketigo Usbekistan", "VET": "Wektu Venezuela", "VLAST": "Wektu Ketigo Vladivostok", "VLAT": "Wektu Standar Vladivostok", "VOLST": "Wektu Ketigo Volgograd", "VOLT": "Wektu Standar Volgograd", "VOST": "Wektu Vostok", "VUT": "Wektu Standar Vanuatu", "VUT DST": "Wektu Ketigo Vanuatu", "WAKT": "Wektu Pulo Wake", "WARST": "Wektu Ketigo Argentina sisih Kulon", "WART": "Wektu Standar Argentina sisih Kulon", "WAST": "Wektu Afrika Kulon", "WAT": "Wektu Afrika Kulon", "WESZ": "Wektu Ketigo Eropa sisih Kulon", "WEZ": "Wektu Standar Eropa sisih Kulon", "WFT": "Wektu Wallis lan Futuna", "WGST": "Wektu Ketigo Grinland Kulon", "WGT": "Wektu Standar Grinland Kulon", "WIB": "Wektu Indonesia sisih Kulon", "WIT": "Wektu Indonesia sisih Wetan", "WITA": "Wektu Indonesia Tengah", "YAKST": "Wektu Ketigo Yakutsk", "YAKT": "Wektu Standar Yakutsk", "YEKST": "Wektu Ketigo Yekaterinburg", "YEKT": "Wektu Standar Yekaterinburg", "YST": "Wektu Yukon", "МСК": "Wektu Standar Moscow", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Wektu Kazakhstan Kulon", "شىعىش قازاق ەلى": "Wektu Kazakhstan Wetan", "قازاق ەلى": "Wektu Kazakhstan", "قىرعىزستان": "Wektu Kirgizstan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Wektu Ketigo Azores"},
	}
}

// Locale returns the current translators string locale
func (jv *jv) Locale() string {
	return jv.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'jv'
func (jv *jv) PluralsCardinal() []locales.PluralRule {
	return jv.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'jv'
func (jv *jv) PluralsOrdinal() []locales.PluralRule {
	return jv.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'jv'
func (jv *jv) PluralsRange() []locales.PluralRule {
	return jv.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'jv'
func (jv *jv) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'jv'
func (jv *jv) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'jv'
func (jv *jv) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (jv *jv) MonthAbbreviated(month time.Month) string {
	return jv.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (jv *jv) MonthsAbbreviated() []string {
	return jv.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (jv *jv) MonthNarrow(month time.Month) string {
	return jv.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (jv *jv) MonthsNarrow() []string {
	return jv.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (jv *jv) MonthWide(month time.Month) string {
	return jv.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (jv *jv) MonthsWide() []string {
	return jv.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (jv *jv) WeekdayAbbreviated(weekday time.Weekday) string {
	return jv.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (jv *jv) WeekdaysAbbreviated() []string {
	return jv.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (jv *jv) WeekdayNarrow(weekday time.Weekday) string {
	return jv.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (jv *jv) WeekdaysNarrow() []string {
	return jv.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (jv *jv) WeekdayShort(weekday time.Weekday) string {
	return jv.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (jv *jv) WeekdaysShort() []string {
	return jv.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (jv *jv) WeekdayWide(weekday time.Weekday) string {
	return jv.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (jv *jv) WeekdaysWide() []string {
	return jv.daysWide
}

// Decimal returns the decimal point of number
func (jv *jv) Decimal() string {
	return jv.decimal
}

// Group returns the group of number
func (jv *jv) Group() string {
	return jv.group
}

// Group returns the minus sign of number
func (jv *jv) Minus() string {
	return jv.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'jv' and handles both Whole and Real numbers based on 'v'
func (jv *jv) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'jv' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (jv *jv) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'jv'
func (jv *jv) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := jv.currencies[currency]
	l := len(s) + len(symbol) + 3 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, jv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, jv.group[0])
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

	for j := len(jv.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, jv.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, jv.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, jv.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'jv'
// in accounting notation.
func (jv *jv) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := jv.currencies[currency]
	l := len(s) + len(symbol) + 3 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, jv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, jv.group[0])
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

		for j := len(jv.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, jv.currencyNegativePrefix[j])
		}

		b = append(b, jv.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(jv.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, jv.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, jv.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'jv'
func (jv *jv) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'jv'
func (jv *jv) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, jv.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'jv'
func (jv *jv) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, jv.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'jv'
func (jv *jv) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, jv.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, jv.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'jv'
func (jv *jv) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, jv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'jv'
func (jv *jv) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, jv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, jv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'jv'
func (jv *jv) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, jv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, jv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'jv'
func (jv *jv) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, jv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, jv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := jv.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
