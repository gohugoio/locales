package scn

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type scn struct {
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

// New returns a new instance of translator for the 'scn' locale
func New() locales.Translator {
	return &scn{
		locale:            "scn",
		pluralsCardinal:   []locales.PluralRule{2, 5, 6},
		pluralsOrdinal:    []locales.PluralRule{5, 6},
		pluralsRange:      []locales.PluralRule{2, 6},
		decimal:           ",",
		group:             ".",
		minus:             "-",
		percent:           "%",
		timeSeparator:     ":",
		currencies:        []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		monthsAbbreviated: []string{"", "jin", "fri", "mar", "apr", "maj", "giu", "gnt", "agu", "sit", "utt", "nuv", "dic"},
		monthsNarrow:      []string{"", "J", "F", "M", "A", "M", "G", "G", "A", "S", "U", "N", "D"},
		monthsWide:        []string{"", "jinnaru", "frivaru", "marzu", "aprili", "maju", "giugnu", "giugnettu", "agustu", "sittèmmiru", "uttùviru", "nuvèmmiru", "dicèmmiru"},
		daysAbbreviated:   []string{"dum", "lun", "mar", "mer", "jov", "ven", "sab"},
		daysNarrow:        []string{"d", "l", "m", "m", "j", "v", "s"},
		daysShort:         []string{"du", "lu", "ma", "me", "jo", "ve", "sa"},
		daysWide:          []string{"dumìnica", "lunnidìa", "martidìa", "mercuridìa", "jovidìa", "venniridìa", "sàbbatu"},
		timezones:         map[string]string{"ACDT": "Ura ligali di l’Australia cintrali", "ACST": "Ura sulari di l’Australia cintrali", "ACT": "ACT", "ACWDT": "Ura ligali di l’Australia cintrali di punenti", "ACWST": "Ura sulari di l’Australia cintrali di punenti", "ADT": "Ura ligali di l’Atlànticu", "ADT Arabia": "Ura ligali Àrabba", "AEDT": "Ura ligali di l’Australia di livanti", "AEST": "Ura sulari di l’Australia di livanti", "AFT": "Ura di l’Afghànistan", "AKDT": "Ura ligali di l’Alaska", "AKST": "Ura sulari di l’Alaska", "AMST": "Ura ligali di l’Amazzonia", "AMST Armenia": "Ura ligali di l’Armenia", "AMT": "Ura sulari di l’Amazzonia", "AMT Armenia": "Ura sulari di l’Armenia", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Ura ligali di l’Argintina", "ART": "Ura sulari di l’Argintina", "AST": "Ura sulari di l’Atlànticu", "AST Arabia": "Ura sulari Àrabba", "AWDT": "Ura ligali di l’Australia di punenti", "AWST": "Ura sulari di l’Australia di punenti", "AZST": "Ura ligali di l’Azerbaijan", "AZT": "Ura sulari di l’Azerbaijan", "BDT Bangladesh": "Ura ligali dû Bàngladesh", "BNT": "Ura dû Brunei", "BOT": "Ura dâ Bulivia", "BRST": "Ura ligali di Brasilia", "BRT": "Ura sulari di Brasilia", "BST Bangladesh": "Ura sulari dû Bàngladesh", "BT": "Ura dû Bhutan", "CAST": "CAST", "CAT": "Ura di l’Àfrica Cintrali", "CCT": "Ura di l’Ìsuli Cocos", "CDT": "Ura ligali cintrali", "CHADT": "Ura ligali di Chatham", "CHAST": "Ura sulari di Chatham", "CHUT": "Ura di Chuuk", "CKT": "Ura sulari di l’Ìsuli Cook", "CKT DST": "Ura ligali di l’Ìsuli Cook", "CLST": "Ura ligali dû Cili", "CLT": "Ura sulari dû Cili", "COST": "Ura ligali dâ Culommia", "COT": "Ura sulari dâ Culommia", "CST": "Ura sulari cintrali", "CST China": "Ura sulari dâ Cina", "CST China DST": "Ura ligali dâ Cina", "CVST": "Ura ligali di Capu Virdi", "CVT": "Ura sulari di Capu Virdi", "CXT": "Ura di l’Ìsula di Natali", "ChST": "Ura di Chamorro", "ChST NMI": "ChST NMI", "CuDT": "Ura ligali di Cubba", "CuST": "Ura sulari di Cubba", "DAVT": "Ura di Davis", "DDUT": "Ura di Dumont d’Urville", "EASST": "Ura ligali di l’Ìsula di Pasca", "EAST": "Ura sulari di l’Ìsula di Pasca", "EAT": "Ura di l’Àfrica di Livanti", "ECT": "Ura di l’Ècuador", "EDT": "Ura ligali livantina", "EGDT": "Ura livantina dâ Gruillannia livantina", "EGST": "Ura sulari dâ Gruillannia livantina", "EST": "Ura sulari livantina", "FEET": "Ura di l’Europa cchiù a Livanti", "FJT": "Ura sulari dî Figi", "FJT Summer": "Ura ligali dî Figi", "FKST": "Ura ligali di l’Ìsuli Falkland", "FKT": "Ura sulari di l’Ìsuli Falkland", "FNST": "Ura ligali di Fernando di Noronha", "FNT": "Ura sulari di Fernando di Noronha", "GALT": "Ura dî Galàpagos", "GAMT": "Ura di Gambier", "GEST": "Ura ligali dâ Giorgia", "GET": "Ura sulari dâ Giorgia", "GFT": "Ura dâ Guiana Francisi", "GIT": "Ura di l’Ìsuli Gilbert", "GMT": "Ura Minzana di Greenwich", "GNSST": "Ura ligali dâ Gruillannia", "GNST": "Ura sulari dâ Gruillannia", "GST": "Ura dâ Giorgia di sciroccu", "GST Guam": "GST Guam", "GYT": "Ura dâ Guiana", "HADT": "Ura ligali di l’Hawaai", "HAST": "Ura sulari di l’Hawaai", "HKST": "Ura ligali di Hong Kong", "HKT": "Ura sulari di Hong Kong", "HOVST": "Ura ligali di Khovd", "HOVT": "Ura sulari di Khovd", "ICT": "Ura di l’Innucina", "IDT": "Ura ligali di Isdraeli", "IOT": "Ura di l’Ucianu Innianu", "IRKST": "Ura ligali di Irtkutsk", "IRKT": "Ura sulari di Irtkutsk", "IRST": "Ura sulari di l’Iran", "IRST DST": "Ura ligali di l’Iran", "IST": "Ura sulari di l’Ìnnia", "IST Israel": "Ura sulari di Isdraeli", "JDT": "Ura ligali dû Giappuni", "JST": "Ura sulari dû Giappuni", "KOST": "Ura di Kosrae", "KRAST": "Ura ligali di Krasnoyarsk", "KRAT": "Ura sulari di Krasnoyarsk", "KST": "Ura sulari dâ Curìa", "KST DST": "Ura ligali dâ Curìa", "LHDT": "Ura ligali di Lord Howe", "LHST": "Ura sulari di Lord Howe", "LINT": "Ura di l’Ìsuli Line", "MAGST": "Ura ligali di Magdan", "MAGT": "Ura sulari di Magdan", "MART": "Ura dî Marchesi", "MAWT": "Ura di Mawson", "MDT": "MDT", "MESZ": "Ura ligali Cintrali Eurupea", "MEZ": "Ura sulari Cintrali Eurupea", "MHT": "Ura di l’Ìsuli Marshall", "MMT": "Ura di Myanmar", "MSD": "Ura ligali di Mosca", "MST": "MST", "MUST": "Ura ligali di Mauritius", "MUT": "Ura sulari di Mauritius", "MVT": "Ura dî Mardivi", "MYT": "Ura dâ Malisia", "NCT": "Ura sulari dâ Nova Calidonia", "NDT": "Ura ligali di Tirranova", "NDT New Caledonia": "Ura ligali dâ Nova Calidonia", "NFDT": "Ura ligali di l’Ìsula Norfolk", "NFT": "Ura sulari di l’Ìsula Norfolk", "NOVST": "Ura ligali di Novosibirsk", "NOVT": "Ura sulari di Novosibirsk", "NPT": "Ura dû Nepal", "NRT": "Ura di Nauru", "NST": "Ura sulari di Tirranova", "NUT": "Ura di Niue", "NZDT": "Ura ligali dâ Nova Zilannia", "NZST": "Ura sulari dâ Nova Zilannia", "OESZ": "Ura ligali di l’Europa di Livanti", "OEZ": "Ura sulari di l’Europa di Livanti", "OMSST": "Ura ligali di Omsk", "OMST": "Ura sulari di Omsk", "PDT": "Ura ligali dû Pacìficu", "PDTM": "Ura ligali dû Mèssicu Pacìficu", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Ura dâ Papua Nova Guinìa", "PHOT": "Ura di l’Ìsuli Phoenix", "PKT": "Ura sulari dû Pàkistan", "PKT DST": "Ura ligali dû Pàkistan", "PMDT": "Ura ligali di S. Pierre e Miquelon", "PMST": "Ura sulari di S. Pierre e Miquelon", "PONT": "Ura di Pohnpei", "PST": "Ura sulari dû Pacìficu", "PST Philippine": "Ura sulari dî Filippini", "PST Philippine DST": "Ura ligali dî Filippini", "PST Pitcairn": "Ura di Pitcairn", "PSTM": "Ura sulari dû Mèssicu Pacìficu", "PWT": "Ura di Palau", "PYST": "Ura ligali dû Paraguay", "PYT": "Ura sulari dû Paraguay", "PYT Korea": "Ura dâ Curia di Tramuntana", "RET": "Ura di Réunion", "ROTT": "Ura di Rothera", "SAKST": "Ura ligali di Sakhalin", "SAKT": "Ura sulari di Sakhalin", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Ura Nurmali di l’Àfrica di Sciroccu", "SBT": "Ura di l’Ìsuli Salumuni", "SCT": "Ura dî Seychelles", "SGT": "Ura di Singapuri", "SLST": "SLST", "SRT": "Ura dû Surinami", "SST Samoa": "Ura sulari dî Samoa miricani", "SST Samoa Apia": "Ura sulari di Samoa", "SST Samoa Apia DST": "Ura ligali di Samoa", "SST Samoa DST": "Ura ligali dî Samoa miricani", "SYOT": "Ura di Syowa", "TAAF": "Ura Francisi di Sciroccu e di l’Antàrtidi", "TAHT": "Ura di Tahiti", "TJT": "Ura dû Taggìkistan", "TKT": "Ura di Tukilau", "TLT": "Ura di Timor di Livanti", "TMST": "Ura ligali dû Turkmènistan", "TMT": "Ura sulari dû Turkmènistan", "TOST": "Ura ligali di Tonga", "TOT": "Ura sulari di Tonga", "TVT": "Ura di Tuvalu", "TWT": "Ura sulari di Taiwan", "TWT DST": "Ura ligali di Taiwan", "ULAST": "Ura ligali di Ulaanbaatar", "ULAT": "Ura sulari di Ulaanbaatar", "UYST": "Ura ligali di l’Uruguay", "UYT": "Ura sulari di l’Uruguay", "UZT": "Ura sulari di l’Uzbèkistan", "UZT DST": "Ura ligali di l’Uzbèkistan", "VET": "Ura dû Vinizzuela", "VLAST": "Ura ligali di Vladìvustok", "VLAT": "Ura sulari di Vladìvustok", "VOLST": "Ura ligali di Vòlgugrad", "VOLT": "Ura sulari di Vòlgugrad", "VOST": "Ura di Vostok", "VUT": "Ura sulari di Vanuatu", "VUT DST": "Ura ligali di Vanuatu", "WAKT": "Ura di l’Ìsula Wake", "WARST": "Ura ligali di l’Argintina di punenti", "WART": "Ura sulari di l’Argintina di punenti", "WAST": "Ura di l’Àfrica di Punenti", "WAT": "Ura di l’Àfrica di Punenti", "WESZ": "Ura ligali di l’Europa di Punenti", "WEZ": "Ura sulari di l’Europa di Punenti", "WFT": "Ura di Wallis e Futuna", "WGST": "Ura ligali dâ Gruillannia punintina", "WGT": "Ura sulari dâ Gruillannia punintina", "WIB": "Ura di l’Innunesia di punenti", "WIT": "Ura di l’Innunesia di livanti", "WITA": "Ura di l’Innunesia cintrali", "YAKST": "Ura ligali di Yakutsk", "YAKT": "Ura sulari di Yakutsk", "YEKST": "Ura ligali di Yekatirimmurgu", "YEKT": "Ura sulari di Yekatirimmurgu", "YST": "Ura dû Yukon", "МСК": "Ura sulari di Mosca", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Ura dû Kazzàkistan di Punenti", "شىعىش قازاق ەلى": "Ura dû Kazzàkistan di Livanti", "قازاق ەلى": "Ura dû Kazzàkistan", "قىرعىزستان": "Ura dû Kirghìzzistan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Ura ligali di l’Azzorri"},
	}
}

// Locale returns the current translators string locale
func (scn *scn) Locale() string {
	return scn.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'scn'
func (scn *scn) PluralsCardinal() []locales.PluralRule {
	return scn.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'scn'
func (scn *scn) PluralsOrdinal() []locales.PluralRule {
	return scn.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'scn'
func (scn *scn) PluralsRange() []locales.PluralRule {
	return scn.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'scn'
func (scn *scn) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)
	i := int64(n)
	e := int64(0)
	iMod1000000 := i % 1000000

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	} else if (e == 0 && i != 0 && iMod1000000 == 0 && v == 0) || (e < 0 || e > 5) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'scn'
func (scn *scn) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 11 || n == 8 || (n >= 80 && n <= 89) || (n >= 800 && n <= 899) {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'scn'
func (scn *scn) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := scn.CardinalPluralRule(num1, v1)
	end := scn.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (scn *scn) MonthAbbreviated(month time.Month) string {
	if len(scn.monthsAbbreviated) == 0 {
		return ""
	}
	return scn.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (scn *scn) MonthsAbbreviated() []string {
	return scn.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (scn *scn) MonthNarrow(month time.Month) string {
	if len(scn.monthsNarrow) == 0 {
		return ""
	}
	return scn.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (scn *scn) MonthsNarrow() []string {
	return scn.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (scn *scn) MonthWide(month time.Month) string {
	if len(scn.monthsWide) == 0 {
		return ""
	}
	return scn.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (scn *scn) MonthsWide() []string {
	return scn.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (scn *scn) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(scn.daysAbbreviated) == 0 {
		return ""
	}
	return scn.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (scn *scn) WeekdaysAbbreviated() []string {
	return scn.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (scn *scn) WeekdayNarrow(weekday time.Weekday) string {
	if len(scn.daysNarrow) == 0 {
		return ""
	}
	return scn.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (scn *scn) WeekdaysNarrow() []string {
	return scn.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (scn *scn) WeekdayShort(weekday time.Weekday) string {
	if len(scn.daysShort) == 0 {
		return ""
	}
	return scn.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (scn *scn) WeekdaysShort() []string {
	return scn.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (scn *scn) WeekdayWide(weekday time.Weekday) string {
	if len(scn.daysWide) == 0 {
		return ""
	}
	return scn.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (scn *scn) WeekdaysWide() []string {
	return scn.daysWide
}

// Decimal returns the decimal point of number
func (scn *scn) Decimal() string {
	return scn.decimal
}

// Group returns the group of number
func (scn *scn) Group() string {
	return scn.group
}

// Group returns the minus sign of number
func (scn *scn) Minus() string {
	return scn.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'scn' and handles both Whole and Real numbers based on 'v'
func (scn *scn) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, scn.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, scn.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, scn.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'scn' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (scn *scn) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, scn.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, scn.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, scn.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'scn'
func (scn *scn) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := scn.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'scn'
// in accounting notation.
func (scn *scn) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := scn.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtDateShort returns the short date representation of 't' for 'scn'
func (scn *scn) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'scn'
func (scn *scn) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, scn.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'scn'
func (scn *scn) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, scn.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'scn'
func (scn *scn) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, scn.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, scn.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'scn'
func (scn *scn) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, scn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'scn'
func (scn *scn) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, scn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, scn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'scn'
func (scn *scn) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, scn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, scn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'scn'
func (scn *scn) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, scn.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, scn.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := scn.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
