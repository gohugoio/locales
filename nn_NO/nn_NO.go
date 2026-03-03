package nn_NO

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type nn_NO struct {
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

// New returns a new instance of translator for the 'nn_NO' locale
func New() locales.Translator {
	return &nn_NO{
		locale:                 "nn_NO",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  " ",
		minus:                  "−",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "kr", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "lei", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "$", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "( ",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "jan.", "feb.", "mars", "apr.", "mai", "juni", "juli", "aug.", "sep.", "okt.", "nov.", "des."},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "januar", "februar", "mars", "april", "mai", "juni", "juli", "august", "september", "oktober", "november", "desember"},
		daysAbbreviated:        []string{"sø.", "må.", "ty.", "on.", "to.", "fr.", "la."},
		daysNarrow:             []string{"S", "M", "T", "O", "T", "F", "L"},
		daysShort:              []string{"sø.", "må.", "ty.", "on.", "to.", "fr.", "la."},
		daysWide:               []string{"søndag", "måndag", "tysdag", "onsdag", "torsdag", "fredag", "laurdag"},
		timezones:              map[string]string{"ACDT": "sentralaustralsk sommartid", "ACST": "sentralaustralsk standardtid", "ACT": "Acre normaltid", "ACWDT": "vest-sentralaustralsk sommartid", "ACWST": "vest-sentralaustralsk standardtid", "ADT": "sommartid for den nordamerikanske atlanterhavskysten", "ADT Arabia": "arabisk sommartid", "AEDT": "austaustralsk sommartid", "AEST": "austaustralsk standardtid", "AFT": "afghansk tid", "AKDT": "alaskisk sommartid", "AKST": "alaskisk normaltid", "AMST": "sommartid for Amazonas", "AMST Armenia": "armensk sommartid", "AMT": "normaltid for Amazonas", "AMT Armenia": "armensk normaltid", "ANAST": "Russisk (Anadyr) sommertid", "ANAT": "Russisk (Anadyr) normaltid", "ARST": "argentinsk sommartid", "ART": "argentinsk normaltid", "AST": "normaltid for den nordamerikanske atlanterhavskysten", "AST Arabia": "arabisk normaltid", "AWDT": "vestaustralsk sommartid", "AWST": "vestaustralsk standardtid", "AZST": "aserbajdsjansk sommartid", "AZT": "aserbajdsjansk normaltid", "BDT Bangladesh": "bangladeshisk sommartid", "BNT": "tidssone for Brunei Darussalam", "BOT": "boliviansk tid", "BRST": "sommartid for Brasilia", "BRT": "normaltid for Brasilia", "BST Bangladesh": "bangladeshisk normaltid", "BT": "bhutansk tid", "CAST": "Casey-tid", "CAT": "sentralafrikansk tid", "CCT": "tidssone for Kokosøyane", "CDT": "sommartid for sentrale Nord-Amerika", "CHADT": "sommartid for Chatham", "CHAST": "normaltid for Chatham", "CHUT": "tidssone for Chuukøyane", "CKT": "normaltid for Cookøyane", "CKT DST": "sommartid for Cookøyane", "CLST": "chilensk sommartid", "CLT": "chilensk normaltid", "COST": "kolombiansk sommartid", "COT": "kolombiansk normaltid", "CST": "normaltid for sentrale Nord-Amerika", "CST China": "kinesisk normaltid", "CST China DST": "kinesisk sommartid", "CVST": "kappverdisk sommartid", "CVT": "kappverdisk normaltid", "CXT": "tidssone for Christmasøya", "ChST": "tidssone for Chamorro", "ChST NMI": "Nord-Marianene-tid", "CuDT": "kubansk sommartid", "CuST": "kubansk normaltid", "DAVT": "tidssone for Davis", "DDUT": "tidssone for Dumont-d’Urville", "EASST": "sommartid for Påskeøya", "EAST": "normaltid for Påskeøya", "EAT": "austafrikansk tid", "ECT": "ecuadoriansk tid", "EDT": "sommartid for den nordamerikanske austkysten", "EGDT": "austgrønlandsk sommartid", "EGST": "austgrønlandsk normaltid", "EST": "normaltid for den nordamerikanske austkysten", "FEET": "fjern-austeuropeisk tid", "FJT": "fijiansk normaltid", "FJT Summer": "fijiansk sommartid", "FKST": "sommartid for Falklandsøyane", "FKT": "normaltid for Falklandsøyane", "FNST": "sommartid for Fernando de Noronha", "FNT": "normaltid for Fernando de Noronha", "GALT": "tidssone for Galápagosøyane", "GAMT": "tidssone for Gambier", "GEST": "georgisk sommartid", "GET": "georgisk normaltid", "GFT": "tidssone for Fransk Guyana", "GIT": "tidssone for Gilbertøyane", "GMT": "Greenwich middeltid", "GNSST": "GNSST", "GNST": "GNST", "GST": "tidssone for Sør-Georgia", "GST Guam": "Guam-tid", "GYT": "guyansk tid", "HADT": "sommartid for Hawaii og Aleutene", "HAST": "normaltid for Hawaii og Aleutene", "HKST": "hongkongkinesisk sommartid", "HKT": "hongkongkinesisk normaltid", "HOVST": "sommartid for Khovd", "HOVT": "normaltid for Khovd", "ICT": "indokinesisk tid", "IDT": "israelsk sommartid", "IOT": "tidssone for Indiahavet", "IRKST": "sommartid for Irkutsk", "IRKT": "normaltid for Irkutsk", "IRST": "iransk normaltid", "IRST DST": "iransk sommartid", "IST": "indisk tid", "IST Israel": "israelsk normaltid", "JDT": "japansk sommartid", "JST": "japansk normaltid", "KOST": "tidssone for Kosrae", "KRAST": "sommartid for Krasnojarsk", "KRAT": "normaltid for Krasnojarsk", "KST": "koreansk normaltid", "KST DST": "koreansk sommartid", "LHDT": "sommartid for Lord Howe-øya", "LHST": "normaltid for Lord Howe-øya", "LINT": "tidssone for Lineøyane", "MAGST": "sommartid for Magadan", "MAGT": "normaltid for Magadan", "MART": "tidssone for Marquesasøyane", "MAWT": "tidssone for Mawson", "MDT": "Macau, sommertid", "MESZ": "sentraleuropeisk sommartid", "MEZ": "sentraleuropeisk standardtid", "MHT": "marshallesisk tid", "MMT": "myanmarsk tid", "MSD": "sommartid for Moskva", "MST": "Macau, standardtid", "MUST": "mauritisk sommartid", "MUT": "mauritisk normaltid", "MVT": "maldivisk tid", "MYT": "malaysisk tid", "NCT": "kaledonsk normaltid", "NDT": "sommartid for Newfoundland", "NDT New Caledonia": "kaledonsk sommartid", "NFDT": "sommartid for Norfolkøya", "NFT": "normaltid for Norfolkøya", "NOVST": "sommartid for Novosibirsk", "NOVT": "normaltid for Novosibirsk", "NPT": "nepalsk tid", "NRT": "naurisk tid", "NST": "normaltid for Newfoundland", "NUT": "tidssone for Niue", "NZDT": "nyzealandsk sommartid", "NZST": "nyzealandsk normaltid", "OESZ": "austeuropeisk sommartid", "OEZ": "austeuropeisk standardtid", "OMSST": "sommartid for Omsk", "OMST": "normaltid for Omsk", "PDT": "sommartid for den nordamerikanske stillehavskysten", "PDTM": "sommartid for den meksikanske stillehavskysten", "PETDT": "Russisk (Petropavlovsk-Kamtsjatskij) sommertid", "PETST": "Russisk (Petropavlovsk-Kamtsjatskij) normaltid", "PGT": "papuansk tid", "PHOT": "tidssone for Phoenixøyane", "PKT": "pakistansk normaltid", "PKT DST": "pakistansk sommartid", "PMDT": "sommartid for Saint-Pierre-et-Miquelon", "PMST": "normaltid for Saint-Pierre-et-Miquelon", "PONT": "tidssone for Pohnpei", "PST": "normaltid for den nordamerikanske stillehavskysten", "PST Philippine": "filippinsk normaltid", "PST Philippine DST": "filippinsk sommartid", "PST Pitcairn": "tidssone for Pitcairn", "PSTM": "normaltid for den meksikanske stillehavskysten", "PWT": "palauisk tid", "PYST": "paraguayansk sommartid", "PYT": "paraguayansk normaltid", "PYT Korea": "tidssone for Pyongyang", "RET": "tidssone for Réunion", "ROTT": "tidssone for Rothera", "SAKST": "sommartid for Sakhalin", "SAKT": "normaltid for Sakhalin", "SAMST": "Russisk (Samara) sommertid", "SAMT": "Russisk (Samara) normaltid", "SAST": "sørafrikansk tid", "SBT": "salomonsk tid", "SCT": "seychellisk tid", "SGT": "singaporsk tid", "SLST": "Lanka-tid", "SRT": "surinamsk tid", "SST Samoa": "samoansk normaltid", "SST Samoa Apia": "normaltid for Apia", "SST Samoa Apia DST": "sommartid for Apia", "SST Samoa DST": "samoansk sommartid", "SYOT": "tidssone for Syowa", "TAAF": "tidssone for Dei franske sørterritoria", "TAHT": "tahitisk tid", "TJT": "tadsjikisk tid", "TKT": "tidssone for Tokelau", "TLT": "austtimoresisk tid", "TMST": "turkmensk sommartid", "TMT": "turkmensk normaltid", "TOST": "tongansk sommartid", "TOT": "tongansk normaltid", "TVT": "tuvalsk tid", "TWT": "normaltid for Taipei", "TWT DST": "sommartid for Taipei", "ULAST": "sommartid for Ulan Bator", "ULAT": "normaltid for Ulan Bator", "UYST": "uruguayansk sommartid", "UYT": "uruguayansk normaltid", "UZT": "usbekisk normaltid", "UZT DST": "usbekisk sommartid", "VET": "venezuelansk tid", "VLAST": "sommartid for Vladivostok", "VLAT": "normaltid for Vladivostok", "VOLST": "sommartid for Volgograd", "VOLT": "normaltid for Volgograd", "VOST": "tidssone for Vostok", "VUT": "vanuatisk normaltid", "VUT DST": "vanuatisk sommartid", "WAKT": "tidssone for Wake Island", "WARST": "vestargentinsk sommartid", "WART": "vestargentinsk normaltid", "WAST": "vestafrikansk tid", "WAT": "vestafrikansk tid", "WESZ": "vesteuropeisk sommartid", "WEZ": "vesteuropeisk standardtid", "WFT": "tidssone for Wallis- og Futunaøyane", "WGST": "vestgrønlandsk sommartid", "WGT": "vestgrønlandsk normaltid", "WIB": "vestindonesisk tid", "WIT": "austindonesisk tid", "WITA": "sentralindonesisk tid", "YAKST": "sommartid for Jakutsk", "YAKT": "normaltid for Jakutsk", "YEKST": "sommartid for Jekaterinburg", "YEKT": "normaltid for Jekaterinburg", "YST": "tidssone for Yukon", "МСК": "normaltid for Moskva", "اقتاۋ": "Aqtau, standardtid", "اقتاۋ قالاسى": "Aqtau, sommertid", "اقتوبە": "Aqtobe, standardtid", "اقتوبە قالاسى": "Aqtobe, sommertid", "الماتى": "Almaty, standardtid", "الماتى قالاسى": "Almaty, sommertid", "باتىس قازاق ەلى": "vestkasakhstansk tid", "شىعىش قازاق ەلى": "austkasakhstansk tid", "قازاق ەلى": "kasakhstansk tid", "قىرعىزستان": "kirgisisk tid", "قىزىلوردا": "Qyzylorda, standardtid", "قىزىلوردا قالاسى": "Qyzylorda, sommertid", "∅∅∅": "asorisk sommartid"},
	}
}

// Locale returns the current translators string locale
func (nn *nn_NO) Locale() string {
	return nn.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'nn_NO'
func (nn *nn_NO) PluralsCardinal() []locales.PluralRule {
	return nn.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'nn_NO'
func (nn *nn_NO) PluralsOrdinal() []locales.PluralRule {
	return nn.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'nn_NO'
func (nn *nn_NO) PluralsRange() []locales.PluralRule {
	return nn.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'nn_NO'
func (nn *nn_NO) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'nn_NO'
func (nn *nn_NO) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'nn_NO'
func (nn *nn_NO) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (nn *nn_NO) MonthAbbreviated(month time.Month) string {
	if len(nn.monthsAbbreviated) == 0 {
		return ""
	}
	return nn.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (nn *nn_NO) MonthsAbbreviated() []string {
	return nn.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (nn *nn_NO) MonthNarrow(month time.Month) string {
	if len(nn.monthsNarrow) == 0 {
		return ""
	}
	return nn.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (nn *nn_NO) MonthsNarrow() []string {
	return nn.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (nn *nn_NO) MonthWide(month time.Month) string {
	if len(nn.monthsWide) == 0 {
		return ""
	}
	return nn.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (nn *nn_NO) MonthsWide() []string {
	return nn.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (nn *nn_NO) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(nn.daysAbbreviated) == 0 {
		return ""
	}
	return nn.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (nn *nn_NO) WeekdaysAbbreviated() []string {
	return nn.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (nn *nn_NO) WeekdayNarrow(weekday time.Weekday) string {
	if len(nn.daysNarrow) == 0 {
		return ""
	}
	return nn.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (nn *nn_NO) WeekdaysNarrow() []string {
	return nn.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (nn *nn_NO) WeekdayShort(weekday time.Weekday) string {
	if len(nn.daysShort) == 0 {
		return ""
	}
	return nn.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (nn *nn_NO) WeekdaysShort() []string {
	return nn.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (nn *nn_NO) WeekdayWide(weekday time.Weekday) string {
	if len(nn.daysWide) == 0 {
		return ""
	}
	return nn.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (nn *nn_NO) WeekdaysWide() []string {
	return nn.daysWide
}

// Decimal returns the decimal point of number
func (nn *nn_NO) Decimal() string {
	return nn.decimal
}

// Group returns the group of number
func (nn *nn_NO) Group() string {
	return nn.group
}

// Group returns the minus sign of number
func (nn *nn_NO) Minus() string {
	return nn.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'nn_NO' and handles both Whole and Real numbers based on 'v'
func (nn *nn_NO) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nn.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(nn.group) - 1; j >= 0; j-- {
					b = append(b, nn.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(nn.minus) - 1; j >= 0; j-- {
			b = append(b, nn.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'nn_NO' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (nn *nn_NO) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 7
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nn.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(nn.minus) - 1; j >= 0; j-- {
			b = append(b, nn.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, nn.percentSuffix...)

	b = append(b, nn.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'nn_NO'
func (nn *nn_NO) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := nn.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nn.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(nn.group) - 1; j >= 0; j-- {
					b = append(b, nn.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(nn.minus) - 1; j >= 0; j-- {
			b = append(b, nn.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, nn.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, nn.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'nn_NO'
// in accounting notation.
func (nn *nn_NO) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := nn.currencies[currency]
	l := len(s) + len(symbol) + 8 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, nn.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(nn.group) - 1; j >= 0; j-- {
					b = append(b, nn.group[j])
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

		for j := len(nn.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, nn.currencyNegativePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, nn.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, nn.currencyNegativeSuffix...)
	} else {

		b = append(b, nn.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'nn_NO'
func (nn *nn_NO) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'nn_NO'
func (nn *nn_NO) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, nn.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'nn_NO'
func (nn *nn_NO) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, nn.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'nn_NO'
func (nn *nn_NO) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, nn.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, nn.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'nn_NO'
func (nn *nn_NO) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'nn_NO'
func (nn *nn_NO) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, nn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'nn_NO'
func (nn *nn_NO) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, nn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'nn_NO'
func (nn *nn_NO) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, nn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, nn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := nn.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
