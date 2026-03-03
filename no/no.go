package no

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type no struct {
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

// New returns a new instance of translator for the 'no' locale
func New() locales.Translator {
	return &no{
		locale:                 "no",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ",",
		group:                  " ",
		minus:                  "−",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "kr", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "L", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "( ",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "jan.", "feb.", "mars", "apr.", "mai", "juni", "juli", "aug.", "sep.", "okt.", "nov.", "des."},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "januar", "februar", "mars", "april", "mai", "juni", "juli", "august", "september", "oktober", "november", "desember"},
		daysAbbreviated:        []string{"søn.", "man.", "tir.", "ons.", "tor.", "fre.", "lør."},
		daysNarrow:             []string{"S", "M", "T", "O", "T", "F", "L"},
		daysShort:              []string{"sø.", "ma.", "ti.", "on.", "to.", "fr.", "lø."},
		daysWide:               []string{"søndag", "mandag", "tirsdag", "onsdag", "torsdag", "fredag", "lørdag"},
		timezones:              map[string]string{"ACDT": "sentralaustralsk sommertid", "ACST": "sentralaustralsk normaltid", "ACT": "Acre normaltid", "ACWDT": "vest-sentralaustralsk sommertid", "ACWST": "vest-sentralaustralsk normaltid", "ADT": "sommertid for den nordamerikanske atlanterhavskysten", "ADT Arabia": "arabisk sommertid", "AEDT": "østaustralsk sommertid", "AEST": "østaustralsk normaltid", "AFT": "afghansk tid", "AKDT": "alaskisk sommertid", "AKST": "alaskisk normaltid", "AMST": "sommertid for Amazonas", "AMST Armenia": "armensk sommertid", "AMT": "normaltid for Amazonas", "AMT Armenia": "armensk normaltid", "ANAST": "Russisk (Anadyr) sommertid", "ANAT": "Russisk (Anadyr) normaltid", "ARST": "argentinsk sommertid", "ART": "argentinsk normaltid", "AST": "normaltid for den nordamerikanske atlanterhavskysten", "AST Arabia": "arabisk standardtid", "AWDT": "vestaustralsk sommertid", "AWST": "vestaustralsk normaltid", "AZST": "aserbajdsjansk sommertid", "AZT": "aserbajdsjansk normaltid", "BDT Bangladesh": "bangladeshisk sommertid", "BNT": "tidssone for Brunei Darussalam", "BOT": "boliviansk tid", "BRST": "sommertid for Brasilia", "BRT": "normaltid for Brasilia", "BST Bangladesh": "bangladeshisk normaltid", "BT": "bhutansk tid", "CAST": "Casey-tid", "CAT": "sentralafrikansk tid", "CCT": "tidssone for Kokosøyene", "CDT": "sommertid for det sentrale Nord-Amerika", "CHADT": "sommertid for Chatham", "CHAST": "normaltid for Chatham", "CHUT": "tidssone for Chuukøyene", "CKT": "normaltid for Cookøyene", "CKT DST": "halv sommertid for Cookøyene", "CLST": "chilensk sommertid", "CLT": "chilensk normaltid", "COST": "colombiansk sommertid", "COT": "colombiansk normaltid", "CST": "normaltid for det sentrale Nord-Amerika", "CST China": "kinesisk normaltid", "CST China DST": "kinesisk sommertid", "CVST": "kappverdisk sommertid", "CVT": "kappverdisk normaltid", "CXT": "tidssone for Christmasøya", "ChST": "tidssone for Chamorro", "ChST NMI": "Nord-Marianene-tid", "CuDT": "cubansk sommertid", "CuST": "cubansk normaltid", "DAVT": "tidssone for Davis", "DDUT": "tidssone for Dumont d’Urville", "EASST": "sommertid for Påskeøya", "EAST": "normaltid for Påskeøya", "EAT": "østafrikansk tid", "ECT": "ecuadoriansk tid", "EDT": "sommertid for den nordamerikanske østkysten", "EGDT": "østgrønlandsk sommertid", "EGST": "østgrønlandsk normaltid", "EST": "normaltid for den nordamerikanske østkysten", "FEET": "fjern-østeuropeisk tid", "FJT": "fijiansk normaltid", "FJT Summer": "fijiansk sommertid", "FKST": "sommertid for Falklandsøyene", "FKT": "normaltid for Falklandsøyene", "FNST": "sommertid for Fernando de Noronha", "FNT": "normaltid for Fernando de Noronha", "GALT": "tidssone for Galápagosøyene", "GAMT": "tidssone for Gambier", "GEST": "georgisk sommertid", "GET": "georgisk normaltid", "GFT": "tidssone for Fransk Guyana", "GIT": "tidssone for Gilbertøyene", "GMT": "Greenwich middeltid", "GNSST": "GNSST", "GNST": "GNST", "GST": "tidssone for Sør-Georgia", "GST Guam": "Guam-tid", "GYT": "guyansk tid", "HADT": "normaltid for Hawaii og Aleutene", "HAST": "normaltid for Hawaii og Aleutene", "HKST": "sommertid for Hongkong", "HKT": "normaltid for Hongkong", "HOVST": "sommertid for Khovd", "HOVT": "normaltid for Khovd", "ICT": "indokinesisk tid", "IDT": "israelsk sommertid", "IOT": "tidssone for Indiahavet", "IRKST": "sommertid for Irkutsk", "IRKT": "normaltid for Irkutsk", "IRST": "iransk normaltid", "IRST DST": "iransk sommertid", "IST": "indisk tid", "IST Israel": "israelsk normaltid", "JDT": "japansk sommertid", "JST": "japansk normaltid", "KOST": "tidssone for Kosrae", "KRAST": "sommertid for Krasnojarsk", "KRAT": "normaltid for Krasnojarsk", "KST": "koreansk normaltid", "KST DST": "koreansk sommertid", "LHDT": "sommertid for Lord Howe-øya", "LHST": "normaltid for Lord Howe-øya", "LINT": "tidssone for Linjeøyene", "MAGST": "sommertid for Magadan", "MAGT": "normaltid for Magadan", "MART": "tidssone for Marquesasøyene", "MAWT": "tidssone for Mawson", "MDT": "sommertid for Rocky Mountains (USA)", "MESZ": "sentraleuropeisk sommertid", "MEZ": "sentraleuropeisk normaltid", "MHT": "marshallesisk tid", "MMT": "myanmarsk tid", "MSD": "sommertid for Moskva", "MST": "normaltid for Rocky Mountains (USA)", "MUST": "mauritisk sommertid", "MUT": "mauritisk normaltid", "MVT": "maldivisk tid", "MYT": "malaysisk tid", "NCT": "kaledonsk normaltid", "NDT": "sommertid for Newfoundland", "NDT New Caledonia": "kaledonsk sommertid", "NFDT": "sommertid for Norfolkøya", "NFT": "normaltid for Norfolkøya", "NOVST": "sommertid for Novosibirsk", "NOVT": "normaltid for Novosibirsk", "NPT": "nepalsk tid", "NRT": "naurisk tid", "NST": "normaltid for Newfoundland", "NUT": "tidssone for Niue", "NZDT": "newzealandsk sommertid", "NZST": "newzealandsk normaltid", "OESZ": "østeuropeisk sommertid", "OEZ": "østeuropeisk normaltid", "OMSST": "sommertid for Omsk", "OMST": "normaltid for Omsk", "PDT": "sommertid for den nordamerikanske Stillehavskysten", "PDTM": "sommertid for den meksikanske Stillehavskysten", "PETDT": "Russisk (Petropavlovsk-Kamtsjatskij) sommertid", "PETST": "Russisk (Petropavlovsk-Kamtsjatskij) normaltid", "PGT": "papuansk tid", "PHOT": "tidssone for Phoenixøyene", "PKT": "pakistansk normaltid", "PKT DST": "pakistansk sommertid", "PMDT": "sommertid for Saint-Pierre-et-Miquelon", "PMST": "normaltid for Saint-Pierre-et-Miquelon", "PONT": "tidssone for Pohnpei", "PST": "normaltid for den nordamerikanske Stillehavskysten", "PST Philippine": "filippinsk normaltid", "PST Philippine DST": "filippinsk sommertid", "PST Pitcairn": "tidssone for Pitcairn", "PSTM": "normaltid for den meksikanske Stillehavskysten", "PWT": "palauisk tid", "PYST": "paraguayansk sommertid", "PYT": "paraguayansk normaltid", "PYT Korea": "tidssone for Pyongyang", "RET": "tidssone for Réunion", "ROTT": "tidssone for Rothera", "SAKST": "sommertid for Sakhalin", "SAKT": "normaltid for Sakhalin", "SAMST": "Russisk (Samara) sommertid", "SAMT": "Russisk (Samara) normaltid", "SAST": "sørafrikansk tid", "SBT": "salomonsk tid", "SCT": "seychellisk tid", "SGT": "singaporsk tid", "SLST": "Lanka-tid", "SRT": "surinamsk tid", "SST Samoa": "samoansk normaltid", "SST Samoa Apia": "normaltid for Apia", "SST Samoa Apia DST": "sommertid for Apia", "SST Samoa DST": "samoansk sommertid", "SYOT": "tidssone for Syowa", "TAAF": "tidssone for De franske sørterritorier", "TAHT": "tahitisk tid", "TJT": "tadsjikisk tid", "TKT": "tidssone for Tokelau", "TLT": "østtimoresisk tid", "TMST": "turkmensk sommertid", "TMT": "turkmensk normaltid", "TOST": "tongansk sommertid", "TOT": "tongansk normaltid", "TVT": "tuvalsk tid", "TWT": "normaltid for Taipei", "TWT DST": "sommertid for Taipei", "ULAST": "sommertid for Ulan Bator", "ULAT": "normaltid for Ulan Bator", "UYST": "uruguayansk sommertid", "UYT": "uruguayansk normaltid", "UZT": "usbekisk normaltid", "UZT DST": "usbekisk sommertid", "VET": "venezuelansk tid", "VLAST": "sommertid for Vladivostok", "VLAT": "normaltid for Vladivostok", "VOLST": "sommertid for Volgograd", "VOLT": "normaltid for Volgograd", "VOST": "tidssone for Vostok", "VUT": "vanuatisk normaltid", "VUT DST": "vanuatisk sommertid", "WAKT": "tidssone for Wake Island", "WARST": "vestargentinsk sommertid", "WART": "vestargentinsk normaltid", "WAST": "vestafrikansk tid", "WAT": "vestafrikansk tid", "WESZ": "vesteuropeisk sommertid", "WEZ": "vesteuropeisk normaltid", "WFT": "tidssone for Wallis- og Futunaøyene", "WGST": "vestgrønlandsk sommertid", "WGT": "vestgrønlandsk normaltid", "WIB": "vestindonesisk tid", "WIT": "østindonesisk tid", "WITA": "sentralindonesisk tid", "YAKST": "sommertid for Jakutsk", "YAKT": "normaltid for Jakutsk", "YEKST": "sommertid for Jekaterinburg", "YEKT": "normaltid for Jekaterinburg", "YST": "tidssone for Yukon", "МСК": "normaltid for Moskva", "اقتاۋ": "Aqtau, standardtid", "اقتاۋ قالاسى": "Aqtau, sommertid", "اقتوبە": "Aqtobe, standardtid", "اقتوبە قالاسى": "Aqtobe, sommertid", "الماتى": "Almaty, standardtid", "الماتى قالاسى": "Almaty, sommertid", "باتىس قازاق ەلى": "vestkasakhstansk tid", "شىعىش قازاق ەلى": "østkasakhstansk tid", "قازاق ەلى": "kasakhstansk tid", "قىرعىزستان": "kirgisisk tid", "قىزىلوردا": "Qyzylorda, standardtid", "قىزىلوردا قالاسى": "Qyzylorda, sommertid", "∅∅∅": "peruansk sommertid"},
	}
}

// Locale returns the current translators string locale
func (no *no) Locale() string {
	return no.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'no'
func (no *no) PluralsCardinal() []locales.PluralRule {
	return no.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'no'
func (no *no) PluralsOrdinal() []locales.PluralRule {
	return no.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'no'
func (no *no) PluralsRange() []locales.PluralRule {
	return no.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'no'
func (no *no) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'no'
func (no *no) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'no'
func (no *no) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (no *no) MonthAbbreviated(month time.Month) string {
	if len(no.monthsAbbreviated) == 0 {
		return ""
	}
	return no.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (no *no) MonthsAbbreviated() []string {
	return no.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (no *no) MonthNarrow(month time.Month) string {
	if len(no.monthsNarrow) == 0 {
		return ""
	}
	return no.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (no *no) MonthsNarrow() []string {
	return no.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (no *no) MonthWide(month time.Month) string {
	if len(no.monthsWide) == 0 {
		return ""
	}
	return no.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (no *no) MonthsWide() []string {
	return no.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (no *no) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(no.daysAbbreviated) == 0 {
		return ""
	}
	return no.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (no *no) WeekdaysAbbreviated() []string {
	return no.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (no *no) WeekdayNarrow(weekday time.Weekday) string {
	if len(no.daysNarrow) == 0 {
		return ""
	}
	return no.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (no *no) WeekdaysNarrow() []string {
	return no.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (no *no) WeekdayShort(weekday time.Weekday) string {
	if len(no.daysShort) == 0 {
		return ""
	}
	return no.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (no *no) WeekdaysShort() []string {
	return no.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (no *no) WeekdayWide(weekday time.Weekday) string {
	if len(no.daysWide) == 0 {
		return ""
	}
	return no.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (no *no) WeekdaysWide() []string {
	return no.daysWide
}

// Decimal returns the decimal point of number
func (no *no) Decimal() string {
	return no.decimal
}

// Group returns the group of number
func (no *no) Group() string {
	return no.group
}

// Group returns the minus sign of number
func (no *no) Minus() string {
	return no.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'no' and handles both Whole and Real numbers based on 'v'
func (no *no) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, no.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(no.group) - 1; j >= 0; j-- {
					b = append(b, no.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(no.minus) - 1; j >= 0; j-- {
			b = append(b, no.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'no' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (no *no) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 7
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, no.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(no.minus) - 1; j >= 0; j-- {
			b = append(b, no.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, no.percentSuffix...)

	b = append(b, no.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'no'
func (no *no) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := no.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, no.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(no.group) - 1; j >= 0; j-- {
					b = append(b, no.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(no.minus) - 1; j >= 0; j-- {
			b = append(b, no.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, no.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, no.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'no'
// in accounting notation.
func (no *no) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := no.currencies[currency]
	l := len(s) + len(symbol) + 8 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, no.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(no.group) - 1; j >= 0; j-- {
					b = append(b, no.group[j])
				}
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

		for j := len(no.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, no.currencyNegativePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, no.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, no.currencyNegativeSuffix...)
	} else {

		b = append(b, no.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'no'
func (no *no) FmtDateShort(t time.Time) string {

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

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'no'
func (no *no) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, no.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'no'
func (no *no) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, no.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'no'
func (no *no) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, no.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, no.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'no'
func (no *no) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, no.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'no'
func (no *no) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, no.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, no.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'no'
func (no *no) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, no.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, no.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'no'
func (no *no) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, no.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, no.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := no.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
