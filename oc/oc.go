package oc

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type oc struct {
	locale                 string
	pluralsCardinal        []locales.PluralRule
	pluralsOrdinal         []locales.PluralRule
	pluralsRange           []locales.PluralRule
	decimal                string
	group                  string
	minus                  string
	percent                string
	percentSuffix          string
	perMille               string
	timeSeparator          string
	inifinity              string
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
	periodsAbbreviated     []string
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'oc' locale
func New() locales.Translator {
	return &oc{
		locale:                 "oc",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          "'h'",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "$AR", "ATS", "$AU", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "$BM", "$BN", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "$BZ", "$CA", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "$CL", "CNH", "CNX", "CNY", "$CO", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "£E", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "$FJ", "£FK", "FRF", "£GB", "GEK", "GEL", "GHC", "GHS", "£GI", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "FC", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "£LB", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "$MX", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "$NA", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "$NZ", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "L", "RSD", "RUB", "RUR", "FR", "SAR", "$SB", "SCR", "SDD", "SDG", "SDP", "SEK", "$SG", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "$SR", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "$T", "TPE", "TRL", "LT", "$TT", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "$US", "USN", "USS", "UYI", "UYP", "$UY", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "$WS", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "FCFP", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "Kw", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "gen.", "feb.", "març", "abr.", "mai", "junh", "jul.", "ago.", "set.", "oct.", "nov.", "dec."},
		monthsNarrow:           []string{"", "G", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "de genièr", "de febrièr", "de març", "d’abril", "de mai", "de junh", "de julhet", "d’agost", "de setembre", "d’octòbre", "de novembre", "de decembre"},
		daysNarrow:             []string{"Dg", "Dl", "Dm", "Dc", "Dj", "Dv", "Ds"},
		daysShort:              []string{"dg", "dl", "dm", "dc", "dj", "dv", "ds"},
		daysWide:               []string{"dimenge", "diluns", "dimars", "dimècres", "dijòus", "divendres", "dissabte"},
		erasAbbreviated:        []string{"Ab. J.C.", "de. J.-C."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"Abans Jèsus-Crist", "despús Jèsus-Crist"},
		timezones:              map[string]string{"ACDT": "ora d’estiu d’Austràlia centrala", "ACST": "ACST", "ACT": "ACT", "ACWDT": "ora d’estiu estandarda d’Austràlia centreoccidentala", "ACWST": "ora estandard d’Austràlia centreoccidentala", "ADT": "ora d’estiu de l’Atlantic", "ADT Arabia": "ora d’estiu d’Arabia", "AEDT": "ora d’estiu d’Austràlia oriental", "AEST": "ora estandard d’Austràlia oriental", "AFT": "ora d’Afganistan", "AKDT": "ora d’estiu d’Alaska", "AKST": "ora estandarda d’Alaska", "AMST": "ora d’estiu l’Amazònes", "AMST Armenia": "ora d’estiu d’Armenia", "AMT": "ora estandarda de l’Amazònes", "AMT Armenia": "ora estandarda d’Armenia", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "ora d’estiu d’Argentina", "ART": "ora estandarda d’Argentina", "AST": "ora estandarda de l’Atlantic", "AST Arabia": "ora estandarda d’Arabia", "AWDT": "ora d’estiu d’Austràlia occidental", "AWST": "ora estandard d’Austràlia occidental", "AZST": "ora d’estiu d’Azerbaijan", "AZT": "ora estandarda d’Azerbaijan", "BDT Bangladesh": "ora d’estiu de Bangaldesh", "BNT": "ora de Brunei", "BOT": "ora de Bolivia", "BRST": "ora d’estiu de Brasilia", "BRT": "ora estandarda de Brasilia", "BST Bangladesh": "ora estandarda de Bangaldesh", "BT": "ora de Butan", "CAST": "CAST", "CAT": "ora d’Africa centrala", "CCT": "ora de l’illa de Cocos", "CDT": "ora d’estiu centrala", "CHADT": "ora d’estiu de Chatham", "CHAST": "ora estandard de Chatham", "CHUT": "ora de Chuuk", "CKT": "ora estandard de les Illes Cook", "CKT DST": "ora d’estiu mieja de les Illes Cook", "CLST": "ora d’estiu de Chile", "CLT": "ora estandarda de Chile", "COST": "ora d’estiu de Colòmbia", "COT": "ora estandarda de Colòmbia", "CST": "ora estandarda centrala", "CST China": "ora estandarda de China", "CST China DST": "ora d’estiu de China", "CVST": "ora d’estiu de Cap-Verd", "CVT": "ora estandarda de Cap-Verd", "CXT": "ora de l’illa de Nadau", "ChST": "ora estandard de Chamorro", "ChST NMI": "ChST NMI", "CuDT": "ora d’estiu de Cuba", "CuST": "ora estandarda de Cuba", "DAVT": "ora de Davis", "DDUT": "ora de Dumont-d’Urville", "EASST": "ora d’estiu de la illa de Pascua", "EAST": "ora estandarda de la illa de Pascua", "EAT": "ora d’Africa de l’èst", "ECT": "ora d’Equator", "EDT": "ora d’estiu de l’Èst", "EGDT": "ora d’estiu de l’èst de Groenlandia", "EGST": "ora estandarda de l’èst de Groenlandia", "EST": "ora estandarda de l’Èst", "FEET": "ora de l’extrem d’Euròpa orientala", "FJT": "ora estandard de Fiji", "FJT Summer": "ora d’estiu de Fiji", "FKST": "ora d’estiu de les illes Malvinas", "FKT": "ora estandarda de les illes Malvinas", "FNST": "ora d’estiu de Fernando Noronha", "FNT": "ora estandarda de Fernando Noronha", "GALT": "ora de Galapagos", "GAMT": "ora de Gambier", "GEST": "ora d’estiu de Geòrgia", "GET": "ora estandarda de Geòrgia", "GFT": "ora de la Guayana Francesa", "GIT": "ora de les illes Gilbert", "GMT": "ora al meridian de Greenwich", "GNSST": "GNSST", "GNST": "GNST", "GST": "ora estandard deth Gòlf", "GST Guam": "GST Guam", "GYT": "ora de la Guayana", "HADT": "ora d’estiu de Hawai-Aleutianes", "HAST": "ora estandarda de Hawai-Aleutianes", "HKST": "ora d’estiu de Hong Kong", "HKT": "ora estandarda de Hong Kong", "HOVST": "ora d’estiu de Hovd", "HOVT": "ora estandarda de Hovd", "ICT": "ora d’Indochina", "IDT": "ora d’estiu d’Israèl", "IOT": "ora de l’Ocean Indic", "IRKST": "ora d’estiu d’Irkutsk", "IRKT": "ora estandarda d’Irkutsk", "IRST": "ora estandarda d’Iran", "IRST DST": "ora d’estiu d’Iran", "IST": "ora estandarda de l’India", "IST Israel": "ora estandarda d’Israèl", "JDT": "ora d’estiu de Japon", "JST": "ora estandarda de Japon", "KOST": "ora de Kosrae", "KRAST": "ora d’estiu de Krasnoyarsk", "KRAT": "ora estandarda de Krasnoyarsk", "KST": "ora estandarda de Corèa", "KST DST": "ora d’estiu de Corèa", "LHDT": "ora d’estiu de Lord Howe", "LHST": "ora estandard de Lord Howe", "LINT": "ora de les Espòrades Equatorials", "MAGST": "ora d’estiu de Magadan", "MAGT": "ora estandarda de Magadan", "MART": "ora de les Illes Marqueses", "MAWT": "ora de Mawson", "MDT": "MDT", "MESZ": "ora d’estiu d’Euròpa centrala", "MEZ": "ora estandarda d’Euròpa centrala", "MHT": "ora de les Illes Marshall", "MMT": "ora de Myanmar", "MSD": "ora d’estiu de Moscòu", "MST": "MST", "MUST": "ora d’estiu de Maurici", "MUT": "ora estandarda de Maurici", "MVT": "ora des Malvines", "MYT": "ora de Malàisia", "NCT": "ora estandard de Nòva Caledònia", "NDT": "ora d’estiu de Terra-Nòva", "NDT New Caledonia": "ora d’estiu de Nòva Caledònia", "NFDT": "ora d’estiu de l’illa Norfolk", "NFT": "ora estandard de l’illa Norfolk", "NOVST": "ora d’estiu de Novosibirsk", "NOVT": "ora estandarda de Novosibirsk", "NPT": "ora de Nepal", "NRT": "ora de Nauru", "NST": "ora estandarda de Terra-Nòva", "NUT": "ora de Niue", "NZDT": "ora d’estiu de Nòva Zelanda", "NZST": "ora estandard de Nòva Zelanda", "OESZ": "ora d’estiu d’Euròpa de l’èst", "OEZ": "ora estandarda d’Euròpa de l’èst", "OMSST": "ora d’estiu d’Omsk", "OMST": "ora estandarda d’Omsk", "PDT": "ora d’estiu del Pacific", "PDTM": "ora d’estiu del Pacific Mexican", "PETDT": "PETDT", "PETST": "PETST", "PGT": "ora de Papua Nòva Guinèa", "PHOT": "ora de les Illes Fénix", "PKT": "ora estandarda de Pakistan", "PKT DST": "ora d’estiu de Pakistan", "PMDT": "ora d’estiu de St. Pierre e Miquelon", "PMST": "ora estandarda de St. Pierre e Miquelon", "PONT": "ora de Pohnpei", "PST": "ora estandarda del Pacific", "PST Philippine": "ora estandarda de Filipines", "PST Philippine DST": "ora d’estiu de Filipines", "PST Pitcairn": "ora de Pitcairn", "PSTM": "ora estandarda del Pacific Mexican", "PWT": "ora de Palaos", "PYST": "ora d’estiu de Paraguay", "PYT": "ora estandarda de Paraguay", "PYT Korea": "ora de Pyongyang", "RET": "ora de la Reünion", "ROTT": "ora de Rothera", "SAKST": "ora d’estiu de Sajalin", "SAKT": "ora estandarda de Sajalin", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "ora de Sudafrica", "SBT": "ora de les Illes Salomon", "SCT": "ora de Seychelles", "SGT": "ora de Singapor", "SLST": "SLST", "SRT": "ora de Surinam", "SST Samoa": "ora estandard de Samoa", "SST Samoa Apia": "ora estandard d’Apia", "SST Samoa Apia DST": "ora d’esiu d’Apia", "SST Samoa DST": "ora d’estiu de Samoa", "SYOT": "ora de Syowa", "TAAF": "ora de les Terres australes e antartiques franceses", "TAHT": "ora de Tahiti", "TJT": "ora de Tajikistan", "TKT": "ora de Tokelau", "TLT": "ora de Timòr oriental", "TMST": "ora d’estiu de Turkmenistan", "TMT": "ora estandarda de Turkmenistan", "TOST": "ora d’estiu de Tonga", "TOT": "ora estandard de Tonga", "TVT": "ora de Tuvalu", "TWT": "ora estandarda de Taipei", "TWT DST": "ora d’estiu de Taipei", "ULAST": "ora d’estiu de Ulan Bator", "ULAT": "ora estandarda de Ulan Bator", "UYST": "ora d’estiu de l’Uruguay", "UYT": "ora estandarda de l’Uruguay", "UZT": "ora estandarda de Uzbekistan", "UZT DST": "ora d’estiu de Uzbekistan", "VET": "ora de Veneçuela", "VLAST": "ora d’estiu de Vladivostok", "VLAT": "ora estandarda de Vladivostok", "VOLST": "ora d’estiu de Volgograd", "VOLT": "ora estandarda de Volgograd", "VOST": "ora de Vostok", "VUT": "ora estandard de Vanuatu", "VUT DST": "ora d’estiu de Vanuatu", "WAKT": "ora de l’Illa Wake", "WARST": "ora d’estiu de l’oèst d’Argentina", "WART": "ora estandarda de l’oèst d’Argentina", "WAST": "ora d’Africa occidentala", "WAT": "ora d’Africa occidentala", "WESZ": "ora d’estiu d’Euròpa de l’oèst", "WEZ": "ora estandarda d’Euròpa de l’oèst", "WFT": "ora de Wallis e Futuna", "WGST": "ora d’estiu de l’oèst de Groenlandia", "WGT": "ora estandarda de l’oèst de Groenlandia", "WIB": "hora d’Indonesia occidentala", "WIT": "hora d’Indonesia orientala", "WITA": "hora d’Indonesia centrala", "YAKST": "ora d’estiu de Yakutsk", "YAKT": "ora estandarda de Yakutsk", "YEKST": "ora d’estiu d’Ekaterimburg", "YEKT": "ora estandarda d’Ekaterimburg", "YST": "ora de Yukon", "МСК": "ora estandarda de Moscòu", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "ora de Kazajistan occidenatal", "شىعىش قازاق ەلى": "ora de Kazajistan orientala", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "ora de Kirguistan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "ora d’estiu de les Açòres"},
	}
}

// Locale returns the current translators string locale
func (oc *oc) Locale() string {
	return oc.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'oc'
func (oc *oc) PluralsCardinal() []locales.PluralRule {
	return oc.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'oc'
func (oc *oc) PluralsOrdinal() []locales.PluralRule {
	return oc.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'oc'
func (oc *oc) PluralsRange() []locales.PluralRule {
	return oc.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'oc'
func (oc *oc) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'oc'
func (oc *oc) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'oc'
func (oc *oc) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (oc *oc) MonthAbbreviated(month time.Month) string {
	return oc.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (oc *oc) MonthsAbbreviated() []string {
	return oc.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (oc *oc) MonthNarrow(month time.Month) string {
	return oc.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (oc *oc) MonthsNarrow() []string {
	return oc.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (oc *oc) MonthWide(month time.Month) string {
	return oc.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (oc *oc) MonthsWide() []string {
	return oc.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (oc *oc) WeekdayAbbreviated(weekday time.Weekday) string {
	return oc.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (oc *oc) WeekdaysAbbreviated() []string {
	return oc.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (oc *oc) WeekdayNarrow(weekday time.Weekday) string {
	return oc.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (oc *oc) WeekdaysNarrow() []string {
	return oc.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (oc *oc) WeekdayShort(weekday time.Weekday) string {
	return oc.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (oc *oc) WeekdaysShort() []string {
	return oc.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (oc *oc) WeekdayWide(weekday time.Weekday) string {
	return oc.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (oc *oc) WeekdaysWide() []string {
	return oc.daysWide
}

// Decimal returns the decimal point of number
func (oc *oc) Decimal() string {
	return oc.decimal
}

// Group returns the group of number
func (oc *oc) Group() string {
	return oc.group
}

// Group returns the minus sign of number
func (oc *oc) Minus() string {
	return oc.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'oc' and handles both Whole and Real numbers based on 'v'
func (oc *oc) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, oc.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(oc.group) - 1; j >= 0; j-- {
					b = append(b, oc.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, oc.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'oc' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (oc *oc) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 6
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, oc.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, oc.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, oc.percentSuffix...)

	b = append(b, oc.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'oc'
func (oc *oc) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := oc.currencies[currency]
	l := len(s) + len(symbol) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, oc.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(oc.group) - 1; j >= 0; j-- {
					b = append(b, oc.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, oc.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, oc.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, oc.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'oc'
// in accounting notation.
func (oc *oc) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := oc.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, oc.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(oc.group) - 1; j >= 0; j-- {
					b = append(b, oc.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, oc.currencyNegativePrefix[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, oc.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, oc.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, oc.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'oc'
func (oc *oc) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'oc'
func (oc *oc) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, oc.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'oc'
func (oc *oc) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, oc.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'oc'
func (oc *oc) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, oc.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, oc.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'oc'
func (oc *oc) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x68}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'oc'
func (oc *oc) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, oc.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, oc.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'oc'
func (oc *oc) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, oc.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, oc.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'oc'
func (oc *oc) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, oc.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, oc.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := oc.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
