package id_ID

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type id_ID struct {
	locale            string
	pluralsCardinal   []locales.PluralRule
	pluralsOrdinal    []locales.PluralRule
	pluralsRange      []locales.PluralRule
	decimal           string
	group             string
	minus             string
	percent           string
	timeSeparator     string
	currencies        []string // idx = enum of currency code
	monthsAbbreviated []string
	monthsNarrow      []string
	monthsWide        []string
	daysAbbreviated   []string
	daysNarrow        []string
	daysShort         []string
	daysWide          []string
	timezones         map[string]string
}

// New returns a new instance of translator for the 'id_ID' locale
func New() locales.Translator {
	return &id_ID{
		locale:            "id_ID",
		pluralsCardinal:   []locales.PluralRule{6},
		pluralsOrdinal:    []locales.PluralRule{6},
		pluralsRange:      []locales.PluralRule{6},
		decimal:           ",",
		group:             ".",
		minus:             "-",
		percent:           "%",
		timeSeparator:     ".",
		currencies:        []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		monthsAbbreviated: []string{"", "Jan", "Feb", "Mar", "Apr", "Mei", "Jun", "Jul", "Agu", "Sep", "Okt", "Nov", "Des"},
		monthsNarrow:      []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:        []string{"", "Januari", "Februari", "Maret", "April", "Mei", "Juni", "Juli", "Agustus", "September", "Oktober", "November", "Desember"},
		daysAbbreviated:   []string{"Min", "Sen", "Sel", "Rab", "Kam", "Jum", "Sab"},
		daysNarrow:        []string{"M", "S", "S", "R", "K", "J", "S"},
		daysWide:          []string{"Minggu", "Senin", "Selasa", "Rabu", "Kamis", "Jumat", "Sabtu"},
		timezones:         map[string]string{"ACDT": "Waktu Musim Panas Tengah Australia", "ACST": "Waktu Musim Panas Acre", "ACT": "Waktu Standar Acre", "ACWDT": "Waktu Musim Panas Barat Tengah Australia", "ACWST": "Waktu Standar Barat Tengah Australia", "ADT": "Waktu Musim Panas Atlantik", "ADT Arabia": "Waktu Musim Panas Arab", "AEDT": "Waktu Musim Panas Timur Australia", "AEST": "Waktu Standar Timur Australia", "AFT": "Waktu Afganistan", "AKDT": "Waktu Musim Panas Alaska", "AKST": "Waktu Standar Alaska", "AMST": "Waktu Musim Panas Amazon", "AMST Armenia": "Waktu Musim Panas Armenia", "AMT": "Waktu Standar Amazon", "AMT Armenia": "Waktu Standar Armenia", "ANAST": "Waktu Musim Panas Anadyr", "ANAT": "Waktu Standar Anadyr", "ARST": "Waktu Musim Panas Argentina", "ART": "Waktu Standar Argentina", "AST": "Waktu Standar Atlantik", "AST Arabia": "Waktu Standar Arab", "AWDT": "Waktu Musim Panas Barat Australia", "AWST": "Waktu Standar Barat Australia", "AZST": "Waktu Musim Panas Azerbaijan", "AZT": "Waktu Standar Azerbaijan", "BDT Bangladesh": "Waktu Musim Panas Bangladesh", "BNT": "Waktu Brunei Darussalam", "BOT": "Waktu Bolivia", "BRST": "Waktu Musim Panas Brasil", "BRT": "Waktu Standar Brasil", "BST Bangladesh": "Waktu Standar Bangladesh", "BT": "Waktu Bhutan", "CAST": "Waktu Casey", "CAT": "Waktu Afrika Tengah", "CCT": "Waktu Kepulauan Cocos", "CDT": "Waktu Musim Panas Tengah", "CHADT": "Waktu Musim Panas Chatham", "CHAST": "Waktu Standar Chatham", "CHUT": "Waktu Chuuk", "CKT": "Waktu Standar Kep. Cook", "CKT DST": "Waktu Tengah Musim Panas Kep. Cook", "CLST": "Waktu Musim Panas Cile", "CLT": "Waktu Standar Cile", "COST": "Waktu Musim Panas Kolombia", "COT": "Waktu Standar Kolombia", "CST": "Waktu Standar Tengah", "CST China": "Waktu Standar Tiongkok", "CST China DST": "Waktu Musim Panas Tiongkok", "CVST": "Waktu Musim Panas Tanjung Verde", "CVT": "Waktu Standar Tanjung Verde", "CXT": "Waktu Pulau Natal", "ChST": "Waktu Standar Chamorro", "ChST NMI": "Waktu Kep. Mariana Utara", "CuDT": "Waktu Musim Panas Kuba", "CuST": "Waktu Standar Kuba", "DAVT": "Waktu Davis", "DDUT": "Waktu Dumont-d’Urville", "EASST": "Waktu Musim Panas Pulau Paskah", "EAST": "Waktu Standar Pulau Paskah", "EAT": "Waktu Afrika Timur", "ECT": "Waktu Ekuador", "EDT": "Waktu Musim Panas Timur", "EGDT": "Waktu Musim Panas Greenland Timur", "EGST": "Waktu Standar Greenland Timur", "EST": "Waktu Standar Timur", "FEET": "Waktu Eropa Timur Jauh", "FJT": "Waktu Standar Fiji", "FJT Summer": "Waktu Musim Panas Fiji", "FKST": "Waktu Musim Panas Kepulauan Falkland", "FKT": "Waktu Standar Kepulauan Falkland", "FNST": "Waktu Musim Panas Fernando de Noronha", "FNT": "Waktu Standar Fernando de Noronha", "GALT": "Waktu Galapagos", "GAMT": "Waktu Gambier", "GEST": "Waktu Musim Panas Georgia", "GET": "Waktu Standar Georgia", "GFT": "Waktu Guyana Prancis", "GIT": "Waktu Kep. Gilbert", "GMT": "Greenwich Mean Time", "GNSST": "GNSST", "GNST": "GNST", "GST": "Waktu Georgia Selatan", "GST Guam": "Waktu Guam", "GYT": "Waktu Guyana", "HADT": "Waktu Musim Panas Hawaii-Aleutian", "HAST": "Waktu Standar Hawaii-Aleutian", "HKST": "Waktu Musim Panas Hong Kong", "HKT": "Waktu Standar Hong Kong", "HOVST": "Waktu Musim Panas Hovd", "HOVT": "Waktu Standar Hovd", "ICT": "Waktu Indochina", "IDT": "Waktu Musim Panas Israel", "IOT": "Waktu Samudera Hindia", "IRKST": "Waktu Musim Panas Irkutsk", "IRKT": "Waktu Standar Irkutsk", "IRST": "Waktu Standar Iran", "IRST DST": "Waktu Musim Panas Iran", "IST": "Waktu India", "IST Israel": "Waktu Standar Israel", "JDT": "Waktu Musim Panas Jepang", "JST": "Waktu Standar Jepang", "KOST": "Waktu Kosrae", "KRAST": "Waktu Musim Panas Krasnoyarsk", "KRAT": "Waktu Standar Krasnoyarsk", "KST": "Waktu Standar Korea", "KST DST": "Waktu Musim Panas Korea", "LHDT": "Waktu Musim Panas Lord Howe", "LHST": "Waktu Standar Lord Howe", "LINT": "Waktu Kep. Line", "MAGST": "Waktu Musim Panas Magadan", "MAGT": "Waktu Standar Magadan", "MART": "Waktu Marquesas", "MAWT": "Waktu Mawson", "MDT": "Waktu Musim Panas Pegunungan", "MESZ": "Waktu Musim Panas Eropa Tengah", "MEZ": "Waktu Standar Eropa Tengah", "MHT": "Waktu Kep. Marshall", "MMT": "Waktu Myanmar", "MSD": "Waktu Musim Panas Moskow", "MST": "Waktu Standar Pegunungan", "MUST": "Waktu Musim Panas Mauritius", "MUT": "Waktu Standar Mauritius", "MVT": "Waktu Maladewa", "MYT": "Waktu Malaysia", "NCT": "Waktu Standar Kaledonia Baru", "NDT": "Waktu Musim Panas Newfoundland", "NDT New Caledonia": "Waktu Musim Panas Kaledonia Baru", "NFDT": "Waktu Musim Panas Pulau Norfolk", "NFT": "Waktu Standar Pulau Norfolk", "NOVST": "Waktu Musim Panas Novosibirsk", "NOVT": "Waktu Standar Novosibirsk", "NPT": "Waktu Nepal", "NRT": "Waktu Nauru", "NST": "Waktu Standar Newfoundland", "NUT": "Waktu Niue", "NZDT": "Waktu Musim Panas Selandia Baru", "NZST": "Waktu Standar Selandia Baru", "OESZ": "Waktu Musim Panas Eropa Timur", "OEZ": "Waktu Standar Eropa Timur", "OMSST": "Waktu Musim Panas Omsk", "OMST": "Waktu Standar Omsk", "PDT": "Waktu Musim Panas Pasifik", "PDTM": "Waktu Musim Panas Pasifik Meksiko", "PETDT": "Waktu Musim Panas Petropavlovsk-Kamchatski", "PETST": "Waktu Standar Petropavlovsk-Kamchatsky", "PGT": "Waktu Papua Nugini", "PHOT": "Waktu Kepulauan Phoenix", "PKT": "Waktu Standar Pakistan", "PKT DST": "Waktu Musim Panas Pakistan", "PMDT": "Waktu Musim Panas Saint Pierre dan Miquelon", "PMST": "Waktu Standar Saint Pierre dan Miquelon", "PONT": "Waktu Ponape", "PST": "Waktu Standar Pasifik", "PST Philippine": "Waktu Standar Filipina", "PST Philippine DST": "Waktu Musim Panas Filipina", "PST Pitcairn": "Waktu Pitcairn", "PSTM": "Waktu Standar Pasifik Meksiko", "PWT": "Waktu Palau", "PYST": "Waktu Musim Panas Paraguay", "PYT": "Waktu Standar Paraguay", "PYT Korea": "Waktu Pyongyang", "RET": "Waktu Reunion", "ROTT": "Waktu Rothera", "SAKST": "Waktu Musim Panas Sakhalin", "SAKT": "Waktu Standar Sakhalin", "SAMST": "Waktu Musim Panas Samara", "SAMT": "Waktu Standar Samara", "SAST": "Waktu Standar Afrika Selatan", "SBT": "Waktu Kepulauan Solomon", "SCT": "Waktu Seychelles", "SGT": "Waktu Standar Singapura", "SLST": "Waktu Lanka", "SRT": "Waktu Suriname", "SST Samoa": "Waktu Standar Samoa", "SST Samoa Apia": "Waktu Standar Apia", "SST Samoa Apia DST": "Waktu Musim Panas Apia", "SST Samoa DST": "Waktu Musim Panas Samoa", "SYOT": "Waktu Syowa", "TAAF": "Waktu Wilayah Selatan dan Antarktika Prancis", "TAHT": "Waktu Tahiti", "TJT": "Waktu Tajikistan", "TKT": "Waktu Tokelau", "TLT": "Waktu Timor Leste", "TMST": "Waktu Musim Panas Turkmenistan", "TMT": "Waktu Standar Turkmenistan", "TOST": "Waktu Musim Panas Tonga", "TOT": "Waktu Standar Tonga", "TVT": "Waktu Tuvalu", "TWT": "Waktu Standar Taipei", "TWT DST": "Waktu Musim Panas Taipei", "ULAST": "Waktu Musim Panas Ulan Bator", "ULAT": "Waktu Standar Ulan Bator", "UYST": "Waktu Musim Panas Uruguay", "UYT": "Waktu Standar Uruguay", "UZT": "Waktu Standar Uzbekistan", "UZT DST": "Waktu Musim Panas Uzbekistan", "VET": "Waktu Venezuela", "VLAST": "Waktu Musim Panas Vladivostok", "VLAT": "Waktu Standar Vladivostok", "VOLST": "Waktu Musim Panas Volgograd", "VOLT": "Waktu Standar Volgograd", "VOST": "Waktu Vostok", "VUT": "Waktu Standar Vanuatu", "VUT DST": "Waktu Musim Panas Vanuatu", "WAKT": "Waktu Kepulauan Wake", "WARST": "Waktu Musim Panas Argentina Bagian Barat", "WART": "Waktu Standar Argentina Bagian Barat", "WAST": "Waktu Afrika Barat", "WAT": "Waktu Afrika Barat", "WESZ": "Waktu Musim Panas Eropa Barat", "WEZ": "Waktu Standar Eropa Barat", "WFT": "Waktu Wallis dan Futuna", "WGST": "Waktu Musim Panas Greenland Barat", "WGT": "Waktu Standar Greenland Barat", "WIB": "Waktu Indonesia Barat", "WIT": "Waktu Indonesia Timur", "WITA": "Waktu Indonesia Tengah", "YAKST": "Waktu Musim Panas Yakutsk", "YAKT": "Waktu Standar Yakutsk", "YEKST": "Waktu Musim Panas Yekaterinburg", "YEKT": "Waktu Standar Yekaterinburg", "YST": "Waktu Yukon", "МСК": "Waktu Standar Moskow", "اقتاۋ": "Waktu Standar Aqtau", "اقتاۋ قالاسى": "Waktu Musim Panas Aqtau", "اقتوبە": "Waktu Standar Aqtobe", "اقتوبە قالاسى": "Waktu Musim Panas Aqtobe", "الماتى": "Waktu Standar Almaty", "الماتى قالاسى": "Waktu Musim Panas Almaty", "باتىس قازاق ەلى": "Waktu Kazakhstan Barat", "شىعىش قازاق ەلى": "Waktu Kazakhstan Timur", "قازاق ەلى": "Waktu Kazakhstan", "قىرعىزستان": "Waktu Kirgizstan", "قىزىلوردا": "Waktu Standar Qyzylorda", "قىزىلوردا قالاسى": "Waktu Musim Panas Qyzylorda", "∅∅∅": "Waktu Musim Panas Azores"},
	}
}

// Locale returns the current translators string locale
func (id *id_ID) Locale() string {
	return id.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'id_ID'
func (id *id_ID) PluralsCardinal() []locales.PluralRule {
	return id.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'id_ID'
func (id *id_ID) PluralsOrdinal() []locales.PluralRule {
	return id.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'id_ID'
func (id *id_ID) PluralsRange() []locales.PluralRule {
	return id.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'id_ID'
func (id *id_ID) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'id_ID'
func (id *id_ID) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'id_ID'
func (id *id_ID) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (id *id_ID) MonthAbbreviated(month time.Month) string {
	return id.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (id *id_ID) MonthsAbbreviated() []string {
	return id.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (id *id_ID) MonthNarrow(month time.Month) string {
	return id.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (id *id_ID) MonthsNarrow() []string {
	return id.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (id *id_ID) MonthWide(month time.Month) string {
	return id.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (id *id_ID) MonthsWide() []string {
	return id.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (id *id_ID) WeekdayAbbreviated(weekday time.Weekday) string {
	return id.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (id *id_ID) WeekdaysAbbreviated() []string {
	return id.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (id *id_ID) WeekdayNarrow(weekday time.Weekday) string {
	return id.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (id *id_ID) WeekdaysNarrow() []string {
	return id.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (id *id_ID) WeekdayShort(weekday time.Weekday) string {
	return id.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (id *id_ID) WeekdaysShort() []string {
	return id.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (id *id_ID) WeekdayWide(weekday time.Weekday) string {
	return id.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (id *id_ID) WeekdaysWide() []string {
	return id.daysWide
}

// Decimal returns the decimal point of number
func (id *id_ID) Decimal() string {
	return id.decimal
}

// Group returns the group of number
func (id *id_ID) Group() string {
	return id.group
}

// Group returns the minus sign of number
func (id *id_ID) Minus() string {
	return id.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'id_ID' and handles both Whole and Real numbers based on 'v'
func (id *id_ID) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, id.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, id.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, id.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'id_ID' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (id *id_ID) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, id.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, id.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, id.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'id_ID'
func (id *id_ID) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := id.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, id.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, id.group[0])
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
		b = append(b, id.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, id.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'id_ID'
// in accounting notation.
func (id *id_ID) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := id.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, id.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, id.group[0])
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

		b = append(b, id.minus[0])

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
			b = append(b, id.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'id_ID'
func (id *id_ID) FmtDateShort(t time.Time) string {
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

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'id_ID'
func (id *id_ID) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, id.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'id_ID'
func (id *id_ID) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, id.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'id_ID'
func (id *id_ID) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, id.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, id.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'id_ID'
func (id *id_ID) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'id_ID'
func (id *id_ID) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'id_ID'
func (id *id_ID) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'id_ID'
func (id *id_ID) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := id.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
