package tr

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type tr struct {
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
	timezones              map[string]string
}

// New returns a new instance of translator for the 'tr' locale
func New() locales.Translator {
	return &tr{
		locale:                 "tr",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AU$", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "L", "RSD", "RUB", "р.", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "₺", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "Oca", "Şub", "Mar", "Nis", "May", "Haz", "Tem", "Ağu", "Eyl", "Eki", "Kas", "Ara"},
		monthsNarrow:           []string{"", "O", "Ş", "M", "N", "M", "H", "T", "A", "E", "E", "K", "A"},
		monthsWide:             []string{"", "Ocak", "Şubat", "Mart", "Nisan", "Mayıs", "Haziran", "Temmuz", "Ağustos", "Eylül", "Ekim", "Kasım", "Aralık"},
		daysAbbreviated:        []string{"Paz", "Pzt", "Sal", "Çar", "Per", "Cum", "Cmt"},
		daysNarrow:             []string{"P", "P", "S", "Ç", "P", "C", "C"},
		daysShort:              []string{"Pa", "Pt", "Sa", "Ça", "Pe", "Cu", "Ct"},
		daysWide:               []string{"Pazar", "Pazartesi", "Salı", "Çarşamba", "Perşembe", "Cuma", "Cumartesi"},
		timezones:              map[string]string{"ACDT": "Orta Avustralya Yaz Saati", "ACST": "Orta Avustralya Standart Saati", "ACT": "Acre Standart Saati", "ACWDT": "İç Batı Avustralya Yaz Saati", "ACWST": "İç Batı Avustralya Standart Saati", "ADT": "Atlantik Yaz Saati", "ADT Arabia": "Arabistan Yaz Saati", "AEDT": "Doğu Avustralya Yaz Saati", "AEST": "Doğu Avustralya Standart Saati", "AFT": "Afganistan Saati", "AKDT": "Alaska Yaz Saati", "AKST": "Alaska Standart Saati", "AMST": "Amazon Yaz Saati", "AMST Armenia": "Ermenistan Yaz Saati", "AMT": "Amazon Standart Saati", "AMT Armenia": "Ermenistan Standart Saati", "ANAST": "Anadır Yaz Saati", "ANAT": "Anadır Standart Saati", "ARST": "Arjantin Yaz Saati", "ART": "Arjantin Standart Saati", "AST": "Atlantik Standart Saati", "AST Arabia": "Arabistan Standart Saati", "AWDT": "Batı Avustralya Yaz Saati", "AWST": "Batı Avustralya Standart Saati", "AZST": "Azerbaycan Yaz Saati", "AZT": "Azerbaycan Standart Saati", "BDT Bangladesh": "Bangladeş Yaz Saati", "BNT": "Brunei Darü’s-Selam Saati", "BOT": "Bolivya Saati", "BRST": "Brasilia Yaz Saati", "BRT": "Brasilia Standart Saati", "BST Bangladesh": "Bangladeş Standart Saati", "BT": "Butan Saati", "CAST": "Casey Saati", "CAT": "Orta Afrika Saati", "CCT": "Cocos Adaları Saati", "CDT": "Kuzey Amerika Merkezi Yaz Saati", "CHADT": "Chatham Yaz Saati", "CHAST": "Chatham Standart Saati", "CHUT": "Chuuk Saati", "CKT": "Cook Adaları Standart Saati", "CKT DST": "Cook Adaları Yarı Yaz Saati", "CLST": "Şili Yaz Saati", "CLT": "Şili Standart Saati", "COST": "Kolombiya Yaz Saati", "COT": "Kolombiya Standart Saati", "CST": "Kuzey Amerika Merkezi Standart Saati", "CST China": "Çin Standart Saati", "CST China DST": "Çin Yaz Saati", "CVST": "Cape Verde Yaz Saati", "CVT": "Cape Verde Standart Saati", "CXT": "Christmas Adası Saati", "ChST": "Chamorro Saati", "ChST NMI": "Kuzey Mariana Adaları Saati", "CuDT": "Küba Yaz Saati", "CuST": "Küba Standart Saati", "DAVT": "Davis Saati", "DDUT": "Dumont-d’Urville Saati", "EASST": "Paskalya Adası Yaz Saati", "EAST": "Paskalya Adası Standart Saati", "EAT": "Doğu Afrika Saati", "ECT": "Ekvador Saati", "EDT": "Kuzey Amerika Doğu Yaz Saati", "EGDT": "Doğu Grönland Yaz Saati", "EGST": "Doğu Grönland Standart Saati", "EST": "Kuzey Amerika Doğu Standart Saati", "FEET": "İleri Doğu Avrupa Saati", "FJT": "Fiji Standart Saati", "FJT Summer": "Fiji Yaz Saati", "FKST": "Falkland Adaları Yaz Saati", "FKT": "Falkland Adaları Standart Saati", "FNST": "Fernando de Noronha Yaz Saati", "FNT": "Fernando de Noronha Standart Saati", "GALT": "Galapagos Saati", "GAMT": "Gambier Saati", "GEST": "Gürcistan Yaz Saati", "GET": "Gürcistan Standart Saati", "GFT": "Fransız Guyanası Saati", "GIT": "Gilbert Adaları Saati", "GMT": "Greenwich Ortalama Saati", "GNSST": "GNSST", "GNST": "GNST", "GST": "Güney Georgia Saati", "GST Guam": "Guam Standart Saati", "GYT": "Guyana Saati", "HADT": "Hawaii-Aleut Yaz Saati", "HAST": "Hawaii-Aleut Standart Saati", "HKST": "Hong Kong Yaz Saati", "HKT": "Hong Kong Standart Saati", "HOVST": "Hovd Yaz Saati", "HOVT": "Hovd Standart Saati", "ICT": "Hindiçin Saati", "IDT": "İsrail Yaz Saati", "IOT": "Hint Okyanusu Saati", "IRKST": "İrkutsk Yaz Saati", "IRKT": "İrkutsk Standart Saati", "IRST": "İran Standart Saati", "IRST DST": "İran Yaz Saati", "IST": "Hindistan Standart Saati", "IST Israel": "İsrail Standart Saati", "JDT": "Japonya Yaz Saati", "JST": "Japonya Standart Saati", "KOST": "Kosrae Saati", "KRAST": "Krasnoyarsk Yaz Saati", "KRAT": "Krasnoyarsk Standart Saati", "KST": "Kore Standart Saati", "KST DST": "Kore Yaz Saati", "LHDT": "Lord Howe Yaz Saati", "LHST": "Lord Howe Standart Saati", "LINT": "Line Adaları Saati", "MAGST": "Magadan Yaz Saati", "MAGT": "Magadan Standart Saati", "MART": "Markiz Adaları Saati", "MAWT": "Mawson Saati", "MDT": "Kuzey Amerika Dağ Yaz Saati", "MESZ": "Orta Avrupa Yaz Saati", "MEZ": "Orta Avrupa Standart Saati", "MHT": "Marshall Adaları Saati", "MMT": "Myanmar Saati", "MSD": "Moskova Yaz Saati", "MST": "Kuzey Amerika Dağ Standart Saati", "MUST": "Mauritius Yaz Saati", "MUT": "Mauritius Standart Saati", "MVT": "Maldivler Saati", "MYT": "Malezya Saati", "NCT": "Yeni Kaledonya Standart Saati", "NDT": "Newfoundland Yaz Saati", "NDT New Caledonia": "Yeni Kaledonya Yaz Saati", "NFDT": "Norfolk Adası Yaz Saati", "NFT": "Norfolk Adası Standart Saati", "NOVST": "Novosibirsk Yaz Saati", "NOVT": "Novosibirsk Standart Saati", "NPT": "Nepal Saati", "NRT": "Nauru Saati", "NST": "Newfoundland Standart Saati", "NUT": "Niue Saati", "NZDT": "Yeni Zelanda Yaz Saati", "NZST": "Yeni Zelanda Standart Saati", "OESZ": "Doğu Avrupa Yaz Saati", "OEZ": "Doğu Avrupa Standart Saati", "OMSST": "Omsk Yaz Saati", "OMST": "Omsk Standart Saati", "PDT": "Kuzey Amerika Pasifik Yaz Saati", "PDTM": "Meksika Pasifik Kıyısı Yaz Saati", "PETDT": "Petropavlovsk-Kamçatski Yaz Saati", "PETST": "Petropavlovsk-Kamçatski Standart Saati", "PGT": "Papua Yeni Gine Saati", "PHOT": "Phoenix Adaları Saati", "PKT": "Pakistan Standart Saati", "PKT DST": "Pakistan Yaz Saati", "PMDT": "Saint Pierre ve Miquelon Yaz Saati", "PMST": "Saint Pierre ve Miquelon Standart Saati", "PONT": "Ponape Saati", "PST": "Kuzey Amerika Pasifik Standart Saati", "PST Philippine": "Filipinler Standart Saati", "PST Philippine DST": "Filipinler Yaz Saati", "PST Pitcairn": "Pitcairn Saati", "PSTM": "Meksika Pasifik Kıyısı Standart Saati", "PWT": "Palau Saati", "PYST": "Paraguay Yaz Saati", "PYT": "Paraguay Standart Saati", "PYT Korea": "Pyongyang Saati", "RET": "Reunion Saati", "ROTT": "Rothera Saati", "SAKST": "Sahalin Yaz Saati", "SAKT": "Sahalin Standart Saati", "SAMST": "Samara Yaz Saati", "SAMT": "Samara Standart Saati", "SAST": "Güney Afrika Standart Saati", "SBT": "Solomon Adaları Saati", "SCT": "Seyşeller Saati", "SGT": "Singapur Standart Saati", "SLST": "Lanka Saati", "SRT": "Surinam Saati", "SST Samoa": "Samoa Standart Saati", "SST Samoa Apia": "Apia Standart Saati", "SST Samoa Apia DST": "Apia Yaz Saati", "SST Samoa DST": "Samoa Yaz Saati", "SYOT": "Showa Saati", "TAAF": "Fransız Güney ve Antarktika Saati", "TAHT": "Tahiti Saati", "TJT": "Tacikistan Saati", "TKT": "Tokelau Saati", "TLT": "Doğu Timor Saati", "TMST": "Türkmenistan Yaz Saati", "TMT": "Türkmenistan Standart Saati", "TOST": "Tonga Yaz Saati", "TOT": "Tonga Standart Saati", "TVT": "Tuvalu Saati", "TWT": "Taipei Standart Saati", "TWT DST": "Taipei Yaz Saati", "ULAST": "Ulan Batur Yaz Saati", "ULAT": "Ulan Batur Standart Saati", "UYST": "Uruguay Yaz Saati", "UYT": "Uruguay Standart Saati", "UZT": "Özbekistan Standart Saati", "UZT DST": "Özbekistan Yaz Saati", "VET": "Venezuela Saati", "VLAST": "Vladivostok Yaz Saati", "VLAT": "Vladivostok Standart Saati", "VOLST": "Volgograd Yaz Saati", "VOLT": "Volgograd Standart Saati", "VOST": "Vostok Saati", "VUT": "Vanuatu Standart Saati", "VUT DST": "Vanuatu Yaz Saati", "WAKT": "Wake Adası Saati", "WARST": "Batı Arjantin Yaz Saati", "WART": "Batı Arjantin Standart Saati", "WAST": "Batı Afrika Saati", "WAT": "Batı Afrika Saati", "WESZ": "Batı Avrupa Yaz Saati", "WEZ": "Batı Avrupa Standart Saati", "WFT": "Wallis ve Futuna Saati", "WGST": "Batı Grönland Yaz Saati", "WGT": "Batı Grönland Standart Saati", "WIB": "Batı Endonezya Saati", "WIT": "Doğu Endonezya Saati", "WITA": "Orta Endonezya Saati", "YAKST": "Yakutsk Yaz Saati", "YAKT": "Yakutsk Standart Saati", "YEKST": "Yekaterinburg Yaz Saati", "YEKT": "Yekaterinburg Standart Saati", "YST": "Yukon Saati", "МСК": "Moskova Standart Saati", "اقتاۋ": "Aktav Standart Saati", "اقتاۋ قالاسى": "Aktav Yaz Saati", "اقتوبە": "Aktöbe Standart Saati", "اقتوبە قالاسى": "Aktöbe Yaz Saati", "الماتى": "Almatı Standart Saati", "الماتى قالاسى": "Almatı Yaz Saati", "باتىس قازاق ەلى": "Batı Kazakistan Saati", "شىعىش قازاق ەلى": "Doğu Kazakistan Saati", "قازاق ەلى": "Kazakistan Saati", "قىرعىزستان": "Kırgızistan Saati", "قىزىلوردا": "Kızılorda Standart Saati", "قىزىلوردا قالاسى": "Kızılorda Yaz Saati", "∅∅∅": "Peru Yaz Saati"},
	}
}

// Locale returns the current translators string locale
func (tr *tr) Locale() string {
	return tr.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'tr'
func (tr *tr) PluralsCardinal() []locales.PluralRule {
	return tr.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'tr'
func (tr *tr) PluralsOrdinal() []locales.PluralRule {
	return tr.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'tr'
func (tr *tr) PluralsRange() []locales.PluralRule {
	return tr.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'tr'
func (tr *tr) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'tr'
func (tr *tr) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'tr'
func (tr *tr) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := tr.CardinalPluralRule(num1, v1)
	end := tr.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (tr *tr) MonthAbbreviated(month time.Month) string {
	if len(tr.monthsAbbreviated) == 0 {
		return ""
	}
	return tr.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (tr *tr) MonthsAbbreviated() []string {
	return tr.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (tr *tr) MonthNarrow(month time.Month) string {
	if len(tr.monthsNarrow) == 0 {
		return ""
	}
	return tr.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (tr *tr) MonthsNarrow() []string {
	return tr.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (tr *tr) MonthWide(month time.Month) string {
	if len(tr.monthsWide) == 0 {
		return ""
	}
	return tr.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (tr *tr) MonthsWide() []string {
	return tr.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (tr *tr) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(tr.daysAbbreviated) == 0 {
		return ""
	}
	return tr.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (tr *tr) WeekdaysAbbreviated() []string {
	return tr.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (tr *tr) WeekdayNarrow(weekday time.Weekday) string {
	if len(tr.daysNarrow) == 0 {
		return ""
	}
	return tr.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (tr *tr) WeekdaysNarrow() []string {
	return tr.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (tr *tr) WeekdayShort(weekday time.Weekday) string {
	if len(tr.daysShort) == 0 {
		return ""
	}
	return tr.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (tr *tr) WeekdaysShort() []string {
	return tr.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (tr *tr) WeekdayWide(weekday time.Weekday) string {
	if len(tr.daysWide) == 0 {
		return ""
	}
	return tr.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (tr *tr) WeekdaysWide() []string {
	return tr.daysWide
}

// Decimal returns the decimal point of number
func (tr *tr) Decimal() string {
	return tr.decimal
}

// Group returns the group of number
func (tr *tr) Group() string {
	return tr.group
}

// Group returns the minus sign of number
func (tr *tr) Minus() string {
	return tr.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'tr' and handles both Whole and Real numbers based on 'v'
func (tr *tr) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, tr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, tr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, tr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'tr' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (tr *tr) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, tr.decimal[0])
			inWhole = true

			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, tr.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, tr.minus[0])
	}

	b = append(b, tr.percent[0])

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'tr'
func (tr *tr) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := tr.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, tr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, tr.group[0])
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
		b = append(b, tr.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, tr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'tr'
// in accounting notation.
func (tr *tr) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := tr.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, tr.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, tr.group[0])
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

		b = append(b, tr.currencyNegativePrefix[0])

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
			b = append(b, tr.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, tr.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'tr'
func (tr *tr) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2e}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'tr'
func (tr *tr) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, tr.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'tr'
func (tr *tr) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, tr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'tr'
func (tr *tr) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, tr.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, tr.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'tr'
func (tr *tr) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, tr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'tr'
func (tr *tr) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, tr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, tr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'tr'
func (tr *tr) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, tr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, tr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'tr'
func (tr *tr) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, tr.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, tr.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := tr.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
