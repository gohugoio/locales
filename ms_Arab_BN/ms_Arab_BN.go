package ms_Arab_BN

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ms_Arab_BN struct {
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
	timezones              map[string]string
}

// New returns a new instance of translator for the 'ms_Arab_BN' locale
func New() locales.Translator {
	return &ms_Arab_BN{
		locale:                 "ms_Arab_BN",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         []locales.PluralRule{2, 6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "$", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "Jan", "Feb", "Mac", "Apr", "Mei", "Jun", "Jul", "Ogo", "Sep", "Okt", "Nov", "Dis"},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "O", "S", "O", "N", "D"},
		monthsWide:             []string{"", "Januari", "Februari", "Mac", "April", "Mei", "Jun", "Julai", "Ogos", "September", "Oktober", "November", "Disember"},
		daysAbbreviated:        []string{"Ahd", "Isn", "Sel", "Rab", "Kha", "Jum", "Sab"},
		daysNarrow:             []string{"A", "I", "S", "R", "K", "J", "S"},
		daysShort:              []string{"Ah", "Is", "Se", "Ra", "Kh", "Ju", "Sa"},
		daysWide:               []string{"Ahad", "Isnin", "Selasa", "Rabu", "Khamis", "Jumaat", "Sabtu"},
		periodsAbbreviated:     []string{"PG", "PTG"},
		timezones:              map[string]string{"ACDT": "Waktu Siang Australia Tengah", "ACST": "Waktu Piawai Australia Tengah", "ACT": "Waktu Piawai Acre", "ACWDT": "Waktu Siang Barat Tengah Australia", "ACWST": "Waktu Piawai Barat Tengah Australia", "ADT": "Waktu Siang Atlantik", "ADT Arabia": "Waktu Siang Arab", "AEDT": "Waktu Siang Australia Timur", "AEST": "Waktu Piawai Timur Australia", "AFT": "Waktu Afghanistan", "AKDT": "Waktu Siang Alaska", "AKST": "Waktu Piawai Alaska", "AMST": "Waktu Musim Panas Amazon", "AMST Armenia": "Waktu Musim Panas Armenia", "AMT": "Waktu Piawai Amazon", "AMT Armenia": "Waktu Piawai Armenia", "ANAST": "Waktu Musim Panas Anadyr", "ANAT": "Waktu Piawai Anadyr", "ARST": "Waktu Musim Panas Argentina", "ART": "Waktu Piawai Argentina", "AST": "Waktu Piawai Atlantik", "AST Arabia": "Waktu Piawai Arab", "AWDT": "Waktu Siang Australia Barat", "AWST": "Waktu Piawai Australia Barat", "AZST": "Waktu Musim Panas Azerbaijan", "AZT": "Waktu Piawai Azerbaijan", "BDT Bangladesh": "Waktu Musim Panas Bangladesh", "BNT": "Waktu Brunei Darussalam", "BOT": "Waktu Bolivia", "BRST": "Waktu Musim Panas Brasilia", "BRT": "Waktu Piawai Brasilia", "BST Bangladesh": "Waktu Piawai Bangladesh", "BT": "Waktu Bhutan", "CAST": "CAST", "CAT": "Waktu Afrika Tengah", "CCT": "Waktu Kepulauan Cocos", "CDT": "Waktu Siang Tengah", "CHADT": "Waktu Siang Chatham", "CHAST": "Waktu Piawai Chatham", "CHUT": "Waktu Chuuk", "CKT": "Waktu Piawai Kepulauan Cook", "CKT DST": "Waktu Musim Panas Separuh Kepulauan Cook", "CLST": "Waktu Musim Panas Chile", "CLT": "Waktu Piawai Chile", "COST": "Waktu Musim Panas Colombia", "COT": "Waktu Piawai Colombia", "CST": "Waktu Piawai Pusat", "CST China": "Waktu Piawai China", "CST China DST": "Waktu Siang China", "CVST": "Waktu Musim Panas Tanjung Verde", "CVT": "Waktu Piawai Tanjung Verde", "CXT": "Waktu Pulau Christmas", "ChST": "Waktu Piawai Chamorro", "ChST NMI": "Waktu Kepulauan Mariana Utara", "CuDT": "Waktu Siang Cuba", "CuST": "Waktu Piawai Cuba", "DAVT": "Waktu Davis", "DDUT": "Waktu Dumont-d’Urville", "EASST": "Waktu Musim Panas Pulau Easter", "EAST": "Waktu Piawai Pulau Easter", "EAT": "Waktu Afrika Timur", "ECT": "Waktu Ecuador", "EDT": "Waktu Siang Timur", "EGDT": "Waktu Musim Panas Greenland Timur", "EGST": "Waktu Piawai Greenland Timur", "EST": "Waktu Piawai Timur", "FEET": "Waktu Eropah ceruk timur", "FJT": "Waktu Piawai Fiji", "FJT Summer": "Waktu Musim Panas Fiji", "FKST": "Waktu Musim Panas Kepulauan Falkland", "FKT": "Waktu Piawai Kepulauan Falkland", "FNST": "Waktu Musim Panas Fernando de Noronha", "FNT": "Waktu Piawai Fernando de Noronha", "GALT": "Waktu Galapagos", "GAMT": "Waktu Gambier", "GEST": "Waktu Musim Panas Georgia", "GET": "Waktu Piawai Georgia", "GFT": "Waktu Guyana Perancis", "GIT": "Waktu Kepulauan Gilbert", "GMT": "Waktu Min Greenwich", "GNSST": "GNSST", "GNST": "GNST", "GST": "Waktu Piawai Teluk", "GST Guam": "GST Guam", "GYT": "Waktu Guyana", "HADT": "Waktu Siang Hawaii-Aleutian", "HAST": "Waktu Piawai Hawaii-Aleutian", "HKST": "Waktu Musim Panas Hong Kong", "HKT": "Waktu Piawai Hong Kong", "HOVST": "Waktu Musim Panas Hovd", "HOVT": "Waktu Piawai Hovd", "ICT": "Waktu Indochina", "IDT": "Waktu Siang Israel", "IOT": "Waktu Lautan Hindi", "IRKST": "Waktu Musim Panas Irkutsk", "IRKT": "Waktu Piawai Irkutsk", "IRST": "Waktu Piawai Iran", "IRST DST": "Waktu Siang Iran", "IST": "Waktu Piawai India", "IST Israel": "Waktu Piawai Israel", "JDT": "Waktu Siang Jepun", "JST": "Waktu Piawai Jepun", "KOST": "Waktu Kosrae", "KRAST": "Waktu Musim Panas Krasnoyarsk", "KRAT": "Waktu Piawai Krasnoyarsk", "KST": "Waktu Piawai Korea", "KST DST": "Waktu Siang Korea", "LHDT": "Waktu Siang Lord Howe", "LHST": "Waktu Piawai Lord Howe", "LINT": "Waktu Kepulauan Line", "MAGST": "Waktu Musim Panas Magadan", "MAGT": "Waktu Piawai Magadan", "MART": "Waktu Marquesas", "MAWT": "Waktu Mawson", "MDT": "Waktu Musim Panas Macao", "MESZ": "Waktu Musim Panas Eropah Tengah", "MEZ": "Waktu Piawai Eropah Tengah", "MHT": "Waktu Kepulauan Marshall", "MMT": "Waktu Myanmar", "MSD": "Waktu Musim Panas Moscow", "MST": "Waktu Piawai Macao", "MUST": "Waktu Musim Panas Mauritius", "MUT": "Waktu Piawai Mauritius", "MVT": "Waktu Maldives", "MYT": "Waktu Malaysia", "NCT": "Waktu Piawai New Caledonia", "NDT": "Waktu Siang Newfoundland", "NDT New Caledonia": "Waktu Musim Panas New Caledonia", "NFDT": "Waktu Siang Kepulauan Norfolk", "NFT": "Waktu Piawai Kepulauan Norfolk", "NOVST": "Waktu Musim Panas Novosibirsk", "NOVT": "Waktu Piawai Novosibirsk", "NPT": "Waktu Nepal", "NRT": "Waktu Nauru", "NST": "Waktu Piawai Newfoundland", "NUT": "Waktu Niue", "NZDT": "Waktu Siang New Zealand", "NZST": "Waktu Piawai New Zealand", "OESZ": "Waktu Musim Panas Eropah Timur", "OEZ": "Waktu Piawai Eropah Timur", "OMSST": "Waktu Musim Panas Omsk", "OMST": "Waktu Piawai Omsk", "PDT": "Waktu Siang Pasifik", "PDTM": "Waktu Siang Pasifik Mexico", "PETDT": "Waktu Musim Panas Petropavlovsk-Kamchatski", "PETST": "Waktu Piawai Petropavlovsk-Kamchatski", "PGT": "Waktu Papua New Guinea", "PHOT": "Waktu Kepulauan Phoenix", "PKT": "Waktu Piawai Pakistan", "PKT DST": "Waktu Musim Panas Pakistan", "PMDT": "Waktu Siang Saint Pierre dan Miquelon", "PMST": "Waktu Piawai Saint Pierre dan Miquelon", "PONT": "Waktu Ponape", "PST": "Waktu Piawai Pasifik", "PST Philippine": "Waktu Piawai Filipina", "PST Philippine DST": "Waktu Musim Panas Filipina", "PST Pitcairn": "Waktu Pitcairn", "PSTM": "Waktu Piawai Pasifik Mexico", "PWT": "Waktu Palau", "PYST": "Waktu Musim Panas Paraguay", "PYT": "Waktu Piawai Paraguay", "PYT Korea": "Waktu Pyongyang", "RET": "Waktu Reunion", "ROTT": "Waktu Rothera", "SAKST": "Waktu Musim Panas Sakhalin", "SAKT": "Waktu Piawai Sakhalin", "SAMST": "Waktu Musim Panas Samara", "SAMT": "Waktu Piawai Samara", "SAST": "Waktu Piawai Afrika Selatan", "SBT": "Waktu Kepulauan Solomon", "SCT": "Waktu Seychelles", "SGT": "Waktu Piawai Singapura", "SLST": "SLST", "SRT": "Waktu Suriname", "SST Samoa": "Waktu Piawai Samoa", "SST Samoa Apia": "Waktu Piawai Apia", "SST Samoa Apia DST": "Waktu Siang Apia", "SST Samoa DST": "Waktu Musim Panas Samoa", "SYOT": "Waktu Syowa", "TAAF": "Waktu Perancis Selatan dan Antartika", "TAHT": "Waktu Tahiti", "TJT": "Waktu Tajikistan", "TKT": "Waktu Tokelau", "TLT": "Waktu Timor Timur", "TMST": "Waktu Musim Panas Turkmenistan", "TMT": "Waktu Piawai Turkmenistan", "TOST": "Waktu Musim Panas Tonga", "TOT": "Waktu Piawai Tonga", "TVT": "Waktu Tuvalu", "TWT": "Waktu Piawai Taipei", "TWT DST": "Waktu Siang Taipei", "ULAST": "Waktu Musim Panas Ulan Bator", "ULAT": "Waktu Piawai Ulan Bator", "UYST": "Waktu Musim Panas Uruguay", "UYT": "Waktu Piawai Uruguay", "UZT": "Waktu Piawai Uzbekistan", "UZT DST": "Waktu Musim Panas Uzbekistan", "VET": "Waktu Venezuela", "VLAST": "Waktu Musim Panas Vladivostok", "VLAT": "Waktu Piawai Vladivostok", "VOLST": "Waktu Musim Panas Volgograd", "VOLT": "Waktu Piawai Volgograd", "VOST": "Waktu Vostok", "VUT": "Waktu Piawai Vanuatu", "VUT DST": "Waktu Musim Panas Vanuatu", "WAKT": "Waktu Pulau Wake", "WARST": "Waktu Musim Panas Argentina Barat", "WART": "Waktu Piawai Argentina Barat", "WAST": "Waktu Afrika Barat", "WAT": "Waktu Afrika Barat", "WESZ": "Waktu Musim Panas Eropah Barat", "WEZ": "Waktu Piawai Eropah Barat", "WFT": "Waktu Wallis dan Futuna", "WGST": "Waktu Musim Panas Greenland Barat", "WGT": "Waktu Piawai Greenland Barat", "WIB": "Waktu Indonesia Barat", "WIT": "Waktu Indonesia Timur", "WITA": "Waktu Indonesia Tengah", "YAKST": "Waktu Musim Panas Yakutsk", "YAKT": "Waktu Piawai Yakutsk", "YEKST": "Waktu Musim Panas Yekaterinburg", "YEKT": "Waktu Piawai Yekaterinburg", "YST": "Masa Yukon", "МСК": "Waktu Piawai Moscow", "اقتاۋ": "Waktu Standard Aqtau", "اقتاۋ قالاسى": "Waktu Musim Panas Aqtau", "اقتوبە": "Waktu Piawai Aqtobe", "اقتوبە قالاسى": "Waktu Musim Panas Aqtobe", "الماتى": "Waktu Piawai Almaty", "الماتى قالاسى": "Waktu Musim Panas Almaty", "باتىس قازاق ەلى": "Waktu Kazakhstan Barat", "شىعىش قازاق ەلى": "Waktu Kazakhstan Timur", "قازاق ەلى": "Waktu Kazakhstan", "قىرعىزستان": "Waktu Kyrgystan", "قىزىلوردا": "Waktu Piawai Qyzylorda", "قىزىلوردا قالاسى": "Waktu Musim Panas Qyzylorda", "∅∅∅": "Waktu Musim Panas Peru"},
	}
}

// Locale returns the current translators string locale
func (ms *ms_Arab_BN) Locale() string {
	return ms.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ms_Arab_BN'
func (ms *ms_Arab_BN) PluralsCardinal() []locales.PluralRule {
	return ms.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ms_Arab_BN'
func (ms *ms_Arab_BN) PluralsOrdinal() []locales.PluralRule {
	return ms.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ms_Arab_BN'
func (ms *ms_Arab_BN) PluralsRange() []locales.PluralRule {
	return ms.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ms_Arab_BN'
func (ms *ms_Arab_BN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ms_Arab_BN'
func (ms *ms_Arab_BN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ms_Arab_BN'
func (ms *ms_Arab_BN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ms *ms_Arab_BN) MonthAbbreviated(month time.Month) string {
	if len(ms.monthsAbbreviated) == 0 {
		return ""
	}
	return ms.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ms *ms_Arab_BN) MonthsAbbreviated() []string {
	return ms.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ms *ms_Arab_BN) MonthNarrow(month time.Month) string {
	if len(ms.monthsNarrow) == 0 {
		return ""
	}
	return ms.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ms *ms_Arab_BN) MonthsNarrow() []string {
	return ms.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ms *ms_Arab_BN) MonthWide(month time.Month) string {
	if len(ms.monthsWide) == 0 {
		return ""
	}
	return ms.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ms *ms_Arab_BN) MonthsWide() []string {
	return ms.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ms *ms_Arab_BN) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(ms.daysAbbreviated) == 0 {
		return ""
	}
	return ms.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ms *ms_Arab_BN) WeekdaysAbbreviated() []string {
	return ms.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ms *ms_Arab_BN) WeekdayNarrow(weekday time.Weekday) string {
	if len(ms.daysNarrow) == 0 {
		return ""
	}
	return ms.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ms *ms_Arab_BN) WeekdaysNarrow() []string {
	return ms.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ms *ms_Arab_BN) WeekdayShort(weekday time.Weekday) string {
	if len(ms.daysShort) == 0 {
		return ""
	}
	return ms.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ms *ms_Arab_BN) WeekdaysShort() []string {
	return ms.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ms *ms_Arab_BN) WeekdayWide(weekday time.Weekday) string {
	if len(ms.daysWide) == 0 {
		return ""
	}
	return ms.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ms *ms_Arab_BN) WeekdaysWide() []string {
	return ms.daysWide
}

// Decimal returns the decimal point of number
func (ms *ms_Arab_BN) Decimal() string {
	return ms.decimal
}

// Group returns the group of number
func (ms *ms_Arab_BN) Group() string {
	return ms.group
}

// Group returns the minus sign of number
func (ms *ms_Arab_BN) Minus() string {
	return ms.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ms_Arab_BN' and handles both Whole and Real numbers based on 'v'
func (ms *ms_Arab_BN) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ms.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ms.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ms.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ms_Arab_BN' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ms *ms_Arab_BN) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ms.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ms.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ms.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ms_Arab_BN'
func (ms *ms_Arab_BN) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ms.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ms.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ms.group[0])
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

	for j := len(ms.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, ms.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, ms.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ms.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ms_Arab_BN'
// in accounting notation.
func (ms *ms_Arab_BN) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ms.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ms.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ms.group[0])
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

		for j := len(ms.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ms.currencyNegativePrefix[j])
		}

		b = append(b, ms.minus[0])

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(ms.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, ms.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ms.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ms_Arab_BN'
func (ms *ms_Arab_BN) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'ms_Arab_BN'
func (ms *ms_Arab_BN) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ms.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ms_Arab_BN'
func (ms *ms_Arab_BN) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ms.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ms_Arab_BN'
func (ms *ms_Arab_BN) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ms.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ms_Arab_BN'
func (ms *ms_Arab_BN) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ms.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, ms.periodsAbbreviated[0]...)
	} else {
		b = append(b, ms.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ms_Arab_BN'
func (ms *ms_Arab_BN) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ms.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ms.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, ms.periodsAbbreviated[0]...)
	} else {
		b = append(b, ms.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ms_Arab_BN'
func (ms *ms_Arab_BN) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ms.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ms.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, ms.periodsAbbreviated[0]...)
	} else {
		b = append(b, ms.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ms_Arab_BN'
func (ms *ms_Arab_BN) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ms.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ms.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, ms.periodsAbbreviated[0]...)
	} else {
		b = append(b, ms.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ms.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
