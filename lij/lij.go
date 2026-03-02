package lij

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type lij struct {
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

// New returns a new instance of translator for the 'lij' locale
func New() locales.Translator {
	return &lij{
		locale:                 "lij",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{5, 6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "de zen.", "de fre.", "de mar.", "d’arv.", "de maz.", "de zug.", "de lug.", "d’ago.", "de set.", "d’ott.", "de nov.", "de dex."},
		monthsNarrow:           []string{"", "ZN", "FR", "MR", "AR", "MZ", "ZG", "LG", "AG", "ST", "OT", "NV", "DX"},
		monthsWide:             []string{"", "de zenâ", "de frevâ", "de marso", "d’arvî", "de mazzo", "de zugno", "de luggio", "d’agosto", "de settembre", "d’ottobre", "de novembre", "de dexembre"},
		daysAbbreviated:        []string{"dom.", "lun.", "mät.", "mäc.", "zeu.", "ven.", "sab."},
		daysNarrow:             []string{"D", "L", "M", "M", "Z", "V", "S"},
		daysWide:               []string{"domenega", "lunesdì", "mätesdì", "mäcordì", "zeuggia", "venardì", "sabbo"},
		timezones:              map[string]string{"ACDT": "oa de stæ de l’Australia de mezo", "ACST": "ACST", "ACT": "ACT", "ACWDT": "oa de stæ de l’Australia do mezo-ponente", "ACWST": "oa standard de l’Australia do mezo-ponente", "ADT": "oa de stæ de l’Atlantico nordamericaña", "ADT Arabia": "oa de stæ de l’Arabia", "AEDT": "oa de stæ de l’Australia de levante", "AEST": "oa standard de l’Australia de levante", "AFT": "oa de l’Afghanistan", "AKDT": "oa de stæ de l’Alaska", "AKST": "oa standard de l’Alaska", "AMST": "oa de stæ de l’Amassònia", "AMST Armenia": "oa de stæ de l’Ermenia", "AMT": "oa standard de l’Amassònia", "AMT Armenia": "oa standard de l’Ermenia", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "oa de stæ de l’Argentiña", "ART": "oa standard de l’Argentiña", "AST": "oa standard de l’Atlantico nordamericaña", "AST Arabia": "oa standard de l’Arabia", "AWDT": "oa de stæ de l’Australia de ponente", "AWST": "oa standard de l’Australia de ponente", "AZST": "oa de stæ de l’Azerbaigian", "AZT": "oa standard de l’Azerbaigian", "BDT Bangladesh": "oa de stæ do Bangladesh", "BNT": "oa do Brunei Darussalam", "BOT": "oa da Bolivia", "BRST": "oa de stæ de Brasilia", "BRT": "oa standard de Brasilia", "BST Bangladesh": "oa standard do Bangladesh", "BT": "oa do Bhutan", "CAST": "CAST", "CAT": "oa de l’Africa do mezo", "CCT": "oa de isoe Cocos", "CDT": "oa de stæ do mezo nordamericaña", "CHADT": "oa de stæ de Chatham", "CHAST": "oa standard de Chatham", "CHUT": "oa do Chuuk", "CKT": "oa standard de isoe Cook", "CKT DST": "oa de stæ de isoe Cook", "CLST": "oa de stæ do Cile", "CLT": "oa standard do Cile", "COST": "oa de stæ da Colombia", "COT": "oa standard da Colombia", "CST": "oa standard do mezo nordamericaña", "CST China": "oa standard da Ciña", "CST China DST": "oa de stæ da Ciña", "CVST": "oa de stæ de Cappo Verde", "CVT": "oa standard de Cappo Verde", "CXT": "oa de l’isoa Christmas", "ChST": "oa de Chamorro", "ChST NMI": "ChST NMI", "CuDT": "oa de stæ de Cubba", "CuST": "oa standard de Cubba", "DAVT": "oa de Davis", "DDUT": "oa de Dumont-d’Urville", "EASST": "oa de stæ de l’isoa de Pasqua", "EAST": "oa standard de l’isoa de Pasqua", "EAT": "oa de l’Africa de levante", "ECT": "oa de l’Ecuador", "EDT": "oa de stæ do levante nordamericaña", "EGDT": "oa de stæ da Groenlandia de levante", "EGST": "oa standard da Groenlandia de levante", "EST": "oa standard do levante nordamericaña", "FEET": "oa de Kaliningrad", "FJT": "oa standard de Figi", "FJT Summer": "oa de stæ de Figi", "FKST": "oa de stæ de isoe Malviñe", "FKT": "oa standard de isoe Malviñe", "FNST": "oa de stæ de Fernando de Noronha", "FNT": "oa standard de Fernando de Noronha", "GALT": "oa de Galapagos", "GAMT": "oa de Gambier", "GEST": "oa de stæ da Geòrgia", "GET": "oa standard da Geòrgia", "GFT": "oa da Guyana franseise", "GIT": "oa de isoe Gilbert", "GMT": "oa do meridian de Greenwich", "GNSST": "GNSST", "GNST": "GNST", "GST": "oa standard do Gorfo", "GST Guam": "GST Guam", "GYT": "oa da Guyana", "HADT": "oa standard de Hawaii-Aleutiñe", "HAST": "oa standard de Hawaii-Aleutiñe", "HKST": "oa de stæ de Hong Kong", "HKT": "oa standard de Hong Kong", "HOVST": "oa de stæ de Hovd", "HOVT": "oa standard de Hovd", "ICT": "oa de l’Indociña", "IDT": "oa de stæ d’Israele", "IOT": "oa de l’Oçeano Indian", "IRKST": "oa de stæ de Irkutsk", "IRKT": "oa standard de Irkutsk", "IRST": "oa standard de l’Iran", "IRST DST": "oa de stæ de l’Iran", "IST": "oa de l’India", "IST Israel": "oa standard d’Israele", "JDT": "oa de stæ do Giappon", "JST": "oa standard do Giappon", "KOST": "oa do Kosrae", "KRAST": "oa de stæ de Krasnoyarsk", "KRAT": "oa standard de Krasnoyarsk", "KST": "oa standard da Corea", "KST DST": "oa de stæ da Corea", "LHDT": "oa de stæ de Lord Howe", "LHST": "oa standard de Lord Howe", "LINT": "oa de isoe da Linia", "MAGST": "oa de stæ de Magadan", "MAGT": "oa standard de Magadan", "MART": "oa de Marcheixi", "MAWT": "oa de Mawson", "MDT": "MDT", "MESZ": "oa de stæ de l’Euröpa do mezo", "MEZ": "oa standard de l’Euröpa do mezo", "MHT": "oa de isoe Marshall", "MMT": "oa da Birmania", "MSD": "oa de stæ de Mosca", "MST": "MST", "MUST": "oa de stæ de Mauritius", "MUT": "oa standard de Mauritius", "MVT": "oa de Maldive", "MYT": "oa da Malesia", "NCT": "oa standard da Neuva Caledònia", "NDT": "oa de stæ de Tæraneuva", "NDT New Caledonia": "oa de stæ da Neuva Caledònia", "NFDT": "oa de stæ de l’isoa Norfolk", "NFT": "oa standard de l’isoa Norfolk", "NOVST": "oa de stæ de Novosibirsk", "NOVT": "oa standard de Novosibirsk", "NPT": "oa do Nepal", "NRT": "oa de Nauru", "NST": "oa standard de Tæraneuva", "NUT": "oa de Niue", "NZDT": "oa de stæ da Neuva Zelanda", "NZST": "oa standard da Neuva Zelanda", "OESZ": "oa de stæ de l’Euröpa de levante", "OEZ": "oa standard de l’Euröpa de levante", "OMSST": "oa de stæ de Òmsk", "OMST": "oa standard de Òmsk", "PDT": "oa de stæ do Paçifico nordamericaña", "PDTM": "oa de stæ do Paçifico mescicaña", "PETDT": "PETDT", "PETST": "PETST", "PGT": "oa da Papua Neuva Guinea", "PHOT": "oa de isoe Phoenix", "PKT": "oa standard do Pakistan", "PKT DST": "oa de stæ do Pakistan", "PMDT": "oa de stæ de San Pê e Miquelon", "PMST": "oa standard de San Pê e Miquelon", "PONT": "oa de Pohnpei", "PST": "oa standard do Paçifico nordamericaña", "PST Philippine": "oa standard de Filipiñe", "PST Philippine DST": "oa de stæ de Filipiñe", "PST Pitcairn": "oa de Pitcairn", "PSTM": "oa standard do Paçifico mescicaña", "PWT": "oa de Palau", "PYST": "oa de stæ do Paraguay", "PYT": "oa standard do Paraguay", "PYT Korea": "oa de Pyongyang", "RET": "oa da Réunion", "ROTT": "oa de Rothera", "SAKST": "oa de stæ de Sakhalin", "SAKT": "oa standard de Sakhalin", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "oa de l’Africa do meridion", "SBT": "oa de isoe Solomon", "SCT": "oa de Seychelles", "SGT": "oa de Scingapô", "SLST": "SLST", "SRT": "oa do Suriname", "SST Samoa": "oa standard de Samoa", "SST Samoa Apia": "oa standard de Apia", "SST Samoa Apia DST": "oa de stæ de Apia", "SST Samoa DST": "oa de stæ de Samoa", "SYOT": "oa de Syowa", "TAAF": "oa de Tære australe e antartiche franseixi", "TAHT": "oa de Tahiti", "TJT": "oa do Tagikistan", "TKT": "oa de Tokelau", "TLT": "oa de Timor Est", "TMST": "oa de stæ do Turkmenistan", "TMT": "oa standard do Turkmenistan", "TOST": "oa de stæ de Tonga", "TOT": "oa standard de Tonga", "TVT": "oa de Tuvalu", "TWT": "oa standard de Taipei", "TWT DST": "oa de stæ de Taipei", "ULAST": "oa de stæ d’Ulan Bator", "ULAT": "oa standard d’Ulan Bator", "UYST": "oa de stæ de l’Uruguay", "UYT": "oa standard de l’Uruguay", "UZT": "oa standard de l’Uzbekistan", "UZT DST": "oa de stæ de l’Uzbekistan", "VET": "oa do Venessuela", "VLAST": "oa de stæ de Vladivostok", "VLAT": "oa standard de Vladivostok", "VOLST": "oa de stæ de Volgograd", "VOLT": "oa standard de Volgograd", "VOST": "oa de Vostok", "VUT": "oa standard de Vanuatu", "VUT DST": "oa de stæ de Vanuatu", "WAKT": "oa de l’isoa de Wake", "WARST": "oa de stæ de l’Argentiña de ponente", "WART": "oa standard de l’Argentiña de ponente", "WAST": "oa de l’Africa de ponente", "WAT": "oa de l’Africa de ponente", "WESZ": "oa de stæ de l’Euröpa de ponente", "WEZ": "oa standard de l’Euröpa de ponente", "WFT": "oa de Wallis e Futuna", "WGST": "oa de stæ da Groenlandia de ponente", "WGT": "oa standard da Groenlandia de ponente", "WIB": "oa de l’Indonesia de ponente", "WIT": "oa de l’Indonesia de levante", "WITA": "oa de l’Indonesia de mezo", "YAKST": "oa de stæ de Yakutsk", "YAKT": "oa standard de Yakutsk", "YEKST": "oa de stæ d’Ekaterinburg", "YEKT": "oa standard d’Ekaterinburg", "YST": "oa do Yukon", "МСК": "oa standard de Mosca", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "oa do Kazakistan de ponente", "شىعىش قازاق ەلى": "oa do Kazakistan de levante", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "oa do Kirghizistan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "oa de stæ do Perù"},
	}
}

// Locale returns the current translators string locale
func (lij *lij) Locale() string {
	return lij.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'lij'
func (lij *lij) PluralsCardinal() []locales.PluralRule {
	return lij.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'lij'
func (lij *lij) PluralsOrdinal() []locales.PluralRule {
	return lij.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'lij'
func (lij *lij) PluralsRange() []locales.PluralRule {
	return lij.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'lij'
func (lij *lij) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'lij'
func (lij *lij) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n == 11 || n == 8 || (n >= 80 && n <= 89) || (n >= 800 && n <= 899) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'lij'
func (lij *lij) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	start := lij.CardinalPluralRule(num1, v1)
	end := lij.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (lij *lij) MonthAbbreviated(month time.Month) string {
	return lij.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (lij *lij) MonthsAbbreviated() []string {
	return lij.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (lij *lij) MonthNarrow(month time.Month) string {
	return lij.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (lij *lij) MonthsNarrow() []string {
	return lij.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (lij *lij) MonthWide(month time.Month) string {
	return lij.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (lij *lij) MonthsWide() []string {
	return lij.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (lij *lij) WeekdayAbbreviated(weekday time.Weekday) string {
	return lij.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (lij *lij) WeekdaysAbbreviated() []string {
	return lij.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (lij *lij) WeekdayNarrow(weekday time.Weekday) string {
	return lij.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (lij *lij) WeekdaysNarrow() []string {
	return lij.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (lij *lij) WeekdayShort(weekday time.Weekday) string {
	return lij.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (lij *lij) WeekdaysShort() []string {
	return lij.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (lij *lij) WeekdayWide(weekday time.Weekday) string {
	return lij.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (lij *lij) WeekdaysWide() []string {
	return lij.daysWide
}

// Decimal returns the decimal point of number
func (lij *lij) Decimal() string {
	return lij.decimal
}

// Group returns the group of number
func (lij *lij) Group() string {
	return lij.group
}

// Group returns the minus sign of number
func (lij *lij) Minus() string {
	return lij.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'lij' and handles both Whole and Real numbers based on 'v'
func (lij *lij) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lij.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, lij.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, lij.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'lij' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (lij *lij) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lij.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, lij.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, lij.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'lij'
func (lij *lij) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := lij.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lij.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, lij.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, lij.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, lij.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, lij.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'lij'
// in accounting notation.
func (lij *lij) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := lij.currencies[currency]
	l := len(s) + len(symbol) + 6 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, lij.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, lij.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, lij.currencyNegativePrefix[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, lij.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, lij.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, lij.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'lij'
func (lij *lij) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'lij'
func (lij *lij) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, lij.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x6f}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'lij'
func (lij *lij) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, lij.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x6f}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'lij'
func (lij *lij) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, lij.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, lij.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x6f}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'lij'
func (lij *lij) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lij.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'lij'
func (lij *lij) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lij.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lij.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'lij'
func (lij *lij) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lij.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lij.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'lij'
func (lij *lij) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, lij.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, lij.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := lij.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
