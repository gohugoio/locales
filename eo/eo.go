package eo

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type eo struct {
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

// New returns a new instance of translator for the 'eo' locale
func New() locales.Translator {
	return &eo{
		locale:                 "eo",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  "\u202f",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "₵", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "₨", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "₨", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "₨", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "₨", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: "\u202f",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: "\u202f)",
		monthsAbbreviated:      []string{"", "Jan", "Feb", "Mar", "Apr", "Maj", "Jun", "Jul", "Aŭg", "Sep", "Okt", "Nov", "Dec"},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "Januaro", "Februaro", "Marto", "Aprilo", "Majo", "Junio", "Julio", "Aŭgusto", "Septembro", "Oktobro", "Novembro", "Decembro"},
		daysAbbreviated:        []string{"di", "lu", "ma", "me", "ĵa", "ve", "sa"},
		daysNarrow:             []string{"d", "l", "m", "m", "ĵ", "v", "s"},
		daysWide:               []string{"dimanĉo", "lundo", "mardo", "merkredo", "ĵaŭdo", "vendredo", "sabato"},
		timezones:              map[string]string{"ACDT": "centra aŭstralia somera tempo", "ACST": "centra aŭstralia norma tempo", "ACT": "ACT", "ACWDT": "centrokcidenta aŭstralia somera tempo", "ACWST": "centrokcidenta aŭstralia norma tempo", "ADT": "atlantika nordamerika somera tempo", "ADT Arabia": "araba somera tempo", "AEDT": "orienta aŭstralia somera tempo", "AEST": "orienta aŭstralia norma tempo", "AFT": "afgana tempo", "AKDT": "alaska somera tempo", "AKST": "alaska norma tempo", "AMST": "amazona somera tempo", "AMST Armenia": "armena somera tempo", "AMT": "amazona norma tempo", "AMT Armenia": "armena norma tempo", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "argentina somera tempo", "ART": "argentina norma tempo", "AST": "atlantika nordamerika norma tempo", "AST Arabia": "araba norma tempo", "AWDT": "okcidenta aŭstralia somera tempo", "AWST": "okcidenta aŭstralia norma tempo", "AZST": "azerbajĝana somera tempo", "AZT": "azerbajĝana norma tempo", "BDT Bangladesh": "bangladeŝa somera tempo", "BNT": "bruneja tempo", "BOT": "bolivia tempo", "BRST": "brazilja somera tempo", "BRT": "brazilja norma tempo", "BST Bangladesh": "bangladeŝa norma tempo", "BT": "butana tempo", "CAST": "CAST", "CAT": "centrafrika tempo", "CCT": "kokosinsula tempo", "CDT": "centra nordamerika somera tempo", "CHADT": "ĉathama somera tempo", "CHAST": "ĉathama norma tempo", "CHUT": "tempo: Ĉuuk", "CKT": "kukinsula norma tempo", "CKT DST": "kukinsula somera tempo", "CLST": "ĉilia somera tempo", "CLT": "ĉilia norma tempo", "COST": "kolombia somera tempo", "COT": "kolombia norma tempo", "CST": "centra nordamerika norma tempo", "CST China": "ĉina norma tempo", "CST China DST": "ĉina somera tempo", "CVST": "kaboverda somera tempo", "CVT": "kaboverda norma tempo", "CXT": "kristnaskinsula tempo", "ChST": "ĉamora tempo", "ChST NMI": "ChST NMI", "CuDT": "Kubo (somera tempo)", "CuST": "Kubo (norma tempo)", "DAVT": "tempo: Davis", "DDUT": "tempo: Dumont d’Urville", "EASST": "paskinsula somera tempo", "EAST": "paskinsula norma tempo", "EAT": "orientafrika tempo", "ECT": "ekvadora tempo", "EDT": "orienta nordamerika somera tempo", "EGDT": "orienta gronlanda somera tempo", "EGST": "orienta gronlanda norma tempo", "EST": "orienta nordamerika norma tempo", "FEET": "ekstrem-orienteŭropa tempo", "FJT": "fiĝia norma tempo", "FJT Summer": "fiĝia somera tempo", "FKST": "falklanda somera tempo", "FKT": "falklanda norma tempo", "FNST": "Fernando de Noronha (somera tempo)", "FNT": "Fernando de Noronha (norma tempo)", "GALT": "galapaga tempo", "GAMT": "tempo: Gambier", "GEST": "kartvela somera tempo", "GET": "kartvela norma tempo", "GFT": "tempo: Franca Gujano", "GIT": "gilbertinsula tempo", "GMT": "universala tempo kunordigita", "GNSST": "GNSST", "GNST": "GNST", "GST": "arabgolfa norma tempo", "GST Guam": "GST Guam", "GYT": "gujana tempo", "HADT": "Havajo-Aleutoj (somera tempo)", "HAST": "Havajo-Aleutoj (norma tempo)", "HKST": "honkonga somera tempo", "HKT": "honkonga norma tempo", "HOVST": "ĥovda somera tempo", "HOVT": "ĥovda norma tempo", "ICT": "hindoĉina tempo", "IDT": "israela somera tempo", "IOT": "hindoceana tempo", "IRKST": "irkutska somera tempo", "IRKT": "irkutska norma tempo", "IRST": "irana norma tempo", "IRST DST": "irana somera tempo", "IST": "hinda norma tempo", "IST Israel": "israela norma tempo", "JDT": "japana somera tempo", "JST": "japana norma tempo", "KOST": "tempo: Kosrae", "KRAST": "krasnojarska somera tempo", "KRAT": "krasnojarska norma tempo", "KST": "korea norma tempo", "KST DST": "korea somera tempo", "LHDT": "Lord Howe (somera tempo)", "LHST": "Lord Howe (norma tempo)", "LINT": "tempo: Liniaj Insuloj", "MAGST": "magadana somera tempo", "MAGT": "magadana norma tempo", "MART": "markizinsula tempo", "MAWT": "tempo: Mawson", "MDT": "MDT", "MESZ": "centreŭropa somera tempo", "MEZ": "centreŭropa norma tempo", "MHT": "marŝalinsula tempo", "MMT": "birma tempo", "MSD": "moskva somera tempo", "MST": "MST", "MUST": "maŭricia somera tempo", "MUT": "maŭricia norma tempo", "MVT": "maldiva tempo", "MYT": "malajzia tempo", "NCT": "novkaledonia norma tempo", "NDT": "Novlando (somera tempo)", "NDT New Caledonia": "novkaledonia somera tempo", "NFDT": "norfolkinsula somera tempo", "NFT": "norfolkinsula norma tempo", "NOVST": "novosibirska somera tempo", "NOVT": "novosibirska norma tempo", "NPT": "nepala tempo", "NRT": "naura tempo", "NST": "Novlando (norma tempo)", "NUT": "niua tempo", "NZDT": "novzelanda somera tempo", "NZST": "novzelanda norma tempo", "OESZ": "orienteŭropa somera tempo", "OEZ": "orienteŭropa norma tempo", "OMSST": "omska somera tempo", "OMST": "omska norma tempo", "PDT": "pacifika nordamerika somera tempo", "PDTM": "pacifika meksika somera tempo", "PETDT": "PETDT", "PETST": "PETST", "PGT": "tempo: Papuo-Nov-Gvineo", "PHOT": "feniksinsula tempo", "PKT": "pakistana norma tempo", "PKT DST": "pakistana somera tempo", "PMDT": "Sankta Piero kaj Mikelono (somera tempo)", "PMST": "Sankta Piero kaj Mikelono (norma tempo)", "PONT": "tempo: Ponape", "PST": "pacifika nordamerika norma tempo", "PST Philippine": "filipina norma tempo", "PST Philippine DST": "filipina somera tempo", "PST Pitcairn": "pitkarninsula tempo", "PSTM": "pacifika meksika norma tempo", "PWT": "palaŭa tempo", "PYST": "paragvaja somera tempo", "PYT": "paragvaja norma tempo", "PYT Korea": "nord-korea tempo", "RET": "tempo: Reunio", "ROTT": "tempo: Rothera", "SAKST": "saĥalena somera tempo", "SAKT": "saĥalena norma tempo", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "sudafrika tempo", "SBT": "tempo: Salomonoj", "SCT": "sejŝela tempo", "SGT": "singapura norma tempo", "SLST": "SLST", "SRT": "surinama tempo", "SST Samoa": "uson-samoa norma tempo", "SST Samoa Apia": "samoa norma tempo", "SST Samoa Apia DST": "samoa somera tempo", "SST Samoa DST": "uson-samoa somera tempo", "SYOT": "tempo: Showa", "TAAF": "tempo: Francaj Sudaj Teritorioj", "TAHT": "tahitia tempo", "TJT": "taĝika tempo", "TKT": "tokelaa tempo", "TLT": "orient-timora tempo", "TMST": "turkmena somera tempo", "TMT": "turkmena norma tempo", "TOST": "tonga somera tempo", "TOT": "tonga norma tempo", "TVT": "tuvala tempo", "TWT": "tajvana norma tempo", "TWT DST": "tajvana somera tempo", "ULAST": "ulanbatora somera tempo", "ULAT": "ulanbatora norma tempo", "UYST": "urugvaja somera tempo", "UYT": "urugvaja norma tempo", "UZT": "uzbeka norma tempo", "UZT DST": "uzbeka somera tempo", "VET": "venezuela tempo", "VLAST": "vladivostoka somera tempo", "VLAT": "vladivostoka norma tempo", "VOLST": "volgograda somera tempo", "VOLT": "volgograda norma tempo", "VOST": "tempo: Vostok", "VUT": "vanuatua norma tempo", "VUT DST": "vanuatua somera tempo", "WAKT": "vejkinsula tempo", "WARST": "okcident-argentina somera tempo", "WART": "okcident-argentina norma tempo", "WAST": "okcidentafrika tempo", "WAT": "okcidentafrika tempo", "WESZ": "okcidenteŭropa somera tempo", "WEZ": "okcidenteŭropa norma tempo", "WFT": "tempo: Valiso kaj Futuno", "WGST": "okcidenta gronlanda somera tempo", "WGT": "okcidenta gronlanda norma tempo", "WIB": "okcident-indonezia tempo", "WIT": "orient-indonezia tempo", "WITA": "centr-indonezia tempo", "YAKST": "jakutska somera tempo", "YAKT": "jakutska norma tempo", "YEKST": "jekaterinburga somera tempo", "YEKT": "jekaterinburga norma tempo", "YST": "jukonia tempo", "МСК": "moskva norma tempo", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "okcident-kazaĥa tempo", "شىعىش قازاق ەلى": "orient-kazaĥa tempo", "قازاق ەلى": "kazaĥa tempo", "قىرعىزستان": "kirgiza tempo", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "perua somera tempo"},
	}
}

// Locale returns the current translators string locale
func (eo *eo) Locale() string {
	return eo.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'eo'
func (eo *eo) PluralsCardinal() []locales.PluralRule {
	return eo.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'eo'
func (eo *eo) PluralsOrdinal() []locales.PluralRule {
	return eo.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'eo'
func (eo *eo) PluralsRange() []locales.PluralRule {
	return eo.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'eo'
func (eo *eo) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'eo'
func (eo *eo) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'eo'
func (eo *eo) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (eo *eo) MonthAbbreviated(month time.Month) string {
	if len(eo.monthsAbbreviated) == 0 {
		return ""
	}
	return eo.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (eo *eo) MonthsAbbreviated() []string {
	return eo.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (eo *eo) MonthNarrow(month time.Month) string {
	if len(eo.monthsNarrow) == 0 {
		return ""
	}
	return eo.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (eo *eo) MonthsNarrow() []string {
	return eo.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (eo *eo) MonthWide(month time.Month) string {
	if len(eo.monthsWide) == 0 {
		return ""
	}
	return eo.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (eo *eo) MonthsWide() []string {
	return eo.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (eo *eo) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(eo.daysAbbreviated) == 0 {
		return ""
	}
	return eo.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (eo *eo) WeekdaysAbbreviated() []string {
	return eo.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (eo *eo) WeekdayNarrow(weekday time.Weekday) string {
	if len(eo.daysNarrow) == 0 {
		return ""
	}
	return eo.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (eo *eo) WeekdaysNarrow() []string {
	return eo.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (eo *eo) WeekdayShort(weekday time.Weekday) string {
	if len(eo.daysShort) == 0 {
		return ""
	}
	return eo.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (eo *eo) WeekdaysShort() []string {
	return eo.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (eo *eo) WeekdayWide(weekday time.Weekday) string {
	if len(eo.daysWide) == 0 {
		return ""
	}
	return eo.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (eo *eo) WeekdaysWide() []string {
	return eo.daysWide
}

// Decimal returns the decimal point of number
func (eo *eo) Decimal() string {
	return eo.decimal
}

// Group returns the group of number
func (eo *eo) Group() string {
	return eo.group
}

// Group returns the minus sign of number
func (eo *eo) Minus() string {
	return eo.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'eo' and handles both Whole and Real numbers based on 'v'
func (eo *eo) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 3*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, eo.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(eo.group) - 1; j >= 0; j-- {
					b = append(b, eo.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, eo.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'eo' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (eo *eo) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, eo.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, eo.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, eo.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'eo'
func (eo *eo) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := eo.currencies[currency]
	l := len(s) + len(symbol) + 5 + 3*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, eo.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(eo.group) - 1; j >= 0; j-- {
					b = append(b, eo.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, eo.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, eo.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, eo.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'eo'
// in accounting notation.
func (eo *eo) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := eo.currencies[currency]
	l := len(s) + len(symbol) + 7 + 3*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, eo.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(eo.group) - 1; j >= 0; j-- {
					b = append(b, eo.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, eo.currencyNegativePrefix[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, eo.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, eo.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, eo.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'eo'
func (eo *eo) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'eo'
func (eo *eo) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2d}...)
	b = append(b, eo.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2d}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'eo'
func (eo *eo) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2d}...)
	b = append(b, eo.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2d}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'eo'
func (eo *eo) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, eo.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20, 0x6c, 0x61}...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d, 0x61}...)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, eo.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'eo'
func (eo *eo) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, eo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'eo'
func (eo *eo) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, eo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, eo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'eo'
func (eo *eo) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, eo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, eo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'eo'
func (eo *eo) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, eo.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, eo.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := eo.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
