package ga_GB

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ga_GB struct {
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

// New returns a new instance of translator for the 'ga_GB' locale
func New() locales.Translator {
	return &ga_GB{
		locale:                 "ga_GB",
		pluralsCardinal:        []locales.PluralRule{2, 3, 4, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 6},
		pluralsRange:           []locales.PluralRule{2, 3, 4, 5, 6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "Ean", "Feabh", "Márta", "Aib", "Beal", "Meith", "Iúil", "Lún", "MFómh", "DFómh", "Samh", "Noll"},
		monthsNarrow:           []string{"", "E", "F", "M", "A", "B", "M", "I", "L", "M", "D", "S", "N"},
		monthsWide:             []string{"", "Eanáir", "Feabhra", "Márta", "Aibreán", "Bealtaine", "Meitheamh", "Iúil", "Lúnasa", "Meán Fómhair", "Deireadh Fómhair", "Samhain", "Nollaig"},
		daysAbbreviated:        []string{"Domh", "Luan", "Máirt", "Céad", "Déar", "Aoine", "Sath"},
		daysNarrow:             []string{"D", "L", "M", "C", "D", "A", "S"},
		daysShort:              []string{"Do", "Lu", "Má", "Cé", "Dé", "Ao", "Sa"},
		daysWide:               []string{"Dé Domhnaigh", "Dé Luain", "Dé Máirt", "Dé Céadaoin", "Déardaoin", "Dé hAoine", "Dé Sathairn"},
		periodsAbbreviated:     []string{"r.n.", "i.n."},
		erasAbbreviated:        []string{"RC", "AD"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"Roimh Chríost", "Anno Domini"},
		timezones:              map[string]string{"ACDT": "Am Samhraidh Lár na hAstráile", "ACST": "Am Samhraidh Acre", "ACT": "Am Caighdeánach Acre", "ACWDT": "Am Samhraidh Mheániarthar na hAstráile", "ACWST": "Am Caighdeánach Mheániarthar na hAstráile", "ADT": "Am Samhraidh an Atlantaigh", "ADT Arabia": "Am Samhraidh na hAraibe", "AEDT": "Am Samhraidh Oirthear na hAstráile", "AEST": "Am Caighdeánach Oirthear na hAstráile", "AFT": "Am na hAfganastáine", "AKDT": "Am Samhraidh Alasca", "AKST": "Am Caighdeánach Alasca", "AMST": "Am Samhraidh na hAmasóine", "AMST Armenia": "Am Samhraidh na hAirméine", "AMT": "Am Caighdeánach na hAmasóine", "AMT Armenia": "Am Caighdeánach na hAirméine", "ANAST": "Am Samhraidh Anadyr", "ANAT": "Am Caighdeánach Anadyr", "ARST": "Am Samhraidh na hAirgintíne", "ART": "Am Caighdeánach na hAirgintíne", "AST": "Am Caighdeánach an Atlantaigh", "AST Arabia": "Am Caighdeánach na hAraibe", "AWDT": "Am Samhraidh Iarthar na hAstráile", "AWST": "Am Caighdeánach Iarthar na hAstráile", "AZST": "Am Samhraidh na hAsarbaiseáine", "AZT": "Am Caighdeánach na hAsarbaiseáine", "BDT Bangladesh": "Am Samhraidh na Banglaidéise", "BNT": "Am Bhrúiné Darasalám", "BOT": "Am na Bolaive", "BRST": "Am Samhraidh Bhrasília", "BRT": "Am Caighdeánach Bhrasília", "BST Bangladesh": "Am Caighdeánach na Banglaidéise", "BT": "Am na Bútáine", "CAST": "Am Stáisiún Casey", "CAT": "Am na hAfraice Láir", "CCT": "Am Oileáin Cocos", "CDT": "Am Samhraidh Lárnach Mheiriceá Thuaidh", "CHADT": "Am Samhraidh Chatham", "CHAST": "Am Caighdeánach Chatham", "CHUT": "Am Chuuk", "CKT": "Am Caighdeánach Oileáin Cook", "CKT DST": "Am Samhraidh Oileáin Cook", "CLST": "Am Samhraidh na Sile", "CLT": "Am Caighdeánach na Sile", "COST": "Am Samhraidh na Colóime", "COT": "Am Caighdeánach na Colóime", "CST": "Am Caighdeánach Lárnach Mheiriceá Thuaidh", "CST China": "Am Caighdeánach na Síne", "CST China DST": "Am Samhraidh na Síne", "CVST": "Am Samhraidh Rinn Verde", "CVT": "Am Caighdeánach Rinn Verde", "CXT": "Am Oileán na Nollag", "ChST": "Am Caighdeánach Seamórach", "ChST NMI": "Am na nOileán Máirianach Thuaidh", "CuDT": "Am Samhraidh Chúba", "CuST": "Am Caighdeánach Chúba", "DAVT": "Am Davis", "DDUT": "Am Dumont-d’Urville", "EASST": "Am Samhraidh Oileán na Cásca", "EAST": "Am Caighdeánach Oileán na Cásca", "EAT": "Am Oirthear na hAfraice", "ECT": "Am Eacuadór", "EDT": "Am Samhraidh Oirthearach Mheiriceá Thuaidh", "EGDT": "Am Samhraidh Oirthear na Graonlainne", "EGST": "Am Caighdeánach Oirthear na Graonlainne", "EST": "Am Caighdeánach Oirthearach Mheiriceá Thuaidh", "FEET": "Am Chianoirthear na hEorpa", "FJT": "Am Caighdeánach Fhidsí", "FJT Summer": "Am Samhraidh Fhidsí", "FKST": "Am Samhraidh Oileáin Fháclainne", "FKT": "Am Caighdeánach Oileáin Fháclainne", "FNST": "Am Samhraidh Fernando de Noronha", "FNT": "Am Caighdeánach Fernando de Noronha", "GALT": "Am Oileáin Galápagos", "GAMT": "Am Gambier", "GEST": "Am Samhraidh na Seoirsia", "GET": "Am Caighdeánach na Seoirsia", "GFT": "Am Ghuáin na Fraince", "GIT": "Am Chiribeas", "GMT": "Meán-Am Greenwich", "GNSST": "GNSST", "GNST": "GNST", "GST": "Am Caighdeánach na Murascaille", "GST Guam": "Am Caighdeánach Ghuam", "GYT": "Am na Guáine", "HADT": "Am Caighdeánach Haváí-Ailiúit", "HAST": "Am Caighdeánach Haváí-Ailiúit", "HKST": "Am Samhraidh Hong Cong", "HKT": "Am Caighdeánach Hong Cong", "HOVST": "Am Samhraidh Khovd", "HOVT": "Am Caighdeánach Khovd", "ICT": "Am na hInd-Síne", "IDT": "Am Samhraidh Iosrael", "IOT": "Am an Aigéin Indiaigh", "IRKST": "Am Samhraidh Irkutsk", "IRKT": "Am Caighdeánach Irkutsk", "IRST": "Am Caighdeánach na hIaráine", "IRST DST": "Am Samhraidh na hIaráine", "IST": "Am Caighdeánach na hIndia", "IST Israel": "Am Caighdeánach Iosrael", "JDT": "Am Samhraidh na Seapáine", "JST": "Am Caighdeánach na Seapáine", "KOST": "Am Kosrae", "KRAST": "Am Samhraidh Krasnoyarsk", "KRAT": "Am Caighdeánach Krasnoyarsk", "KST": "Am Caighdeánach na Cóiré", "KST DST": "Am Samhraidh na Cóiré", "LHDT": "Am Samhraidh Lord Howe", "LHST": "Am Caighdeánach Lord Howe", "LINT": "Am Oileáin na Líne", "MAGST": "Am Samhraidh Mhagadan", "MAGT": "Am Caighdeánach Mhagadan", "MART": "Am na nOileán Marcasach", "MAWT": "Am Mawson", "MDT": "Am Samhraidh Mhacao", "MESZ": "Am Samhraidh Lár na hEorpa", "MEZ": "Am Caighdeánach Lár na hEorpa", "MHT": "Am Oileáin Marshall", "MMT": "Am Mhaenmar", "MSD": "Am Samhraidh Mhoscó", "MST": "Am Caighdeánach Mhacao", "MUST": "Am Samhraidh Oileán Mhuirís", "MUT": "Am Caighdeánach Oileán Mhuirís", "MVT": "Am Oileáin Mhaildíve", "MYT": "Am na Malaeisia", "NCT": "Am Caighdeánach na Nua-Chaladóine", "NDT": "Am Samhraidh Thalamh an Éisc", "NDT New Caledonia": "Am Samhraidh na Nua-Chaladóine", "NFDT": "Am Samhraidh Oileán Norfolk", "NFT": "Am Caighdeánach Oileán Norfolk", "NOVST": "Am Samhraidh Novosibirsk", "NOVT": "Am Caighdeánach Novosibirsk", "NPT": "Am Neipeal", "NRT": "Am Nárú", "NST": "Am Caighdeánach Thalamh an Éisc", "NUT": "Am Niue", "NZDT": "Am Samhraidh na Nua-Shéalainne", "NZST": "Am Caighdeánach na Nua-Shéalainne", "OESZ": "Am Samhraidh Oirthear na hEorpa", "OEZ": "Am Caighdeánach Oirthear na hEorpa", "OMSST": "Am Samhraidh Omsk", "OMST": "Am Caighdeánach Omsk", "PDT": "Am Samhraidh an Aigéin Chiúin", "PDTM": "Am Samhraidh Meicsiceach an Aigéin Chiúin", "PETDT": "Am Samhraidh Phetropavlovsk-Kamchatski", "PETST": "Am Caighdeánach Phetropavlovsk-Kamchatski", "PGT": "Am Nua-Ghuine Phapua", "PHOT": "Am Oileáin an Fhéinics", "PKT": "Am Caighdeánach na Pacastáine", "PKT DST": "Am Samhraidh na Pacastáine", "PMDT": "Am Samhraidh Saint-Pierre-et-Miquelon", "PMST": "Am Caighdeánach Saint-Pierre-et-Miquelon", "PONT": "Am Pohnpei", "PST": "Am Caighdeánach an Aigéin Chiúin", "PST Philippine": "Am Caighdeánach na nOileán Filipíneach", "PST Philippine DST": "Am Samhraidh na nOileán Filipíneach", "PST Pitcairn": "Am Pitcairn", "PSTM": "Am Caighdeánach Meicsiceach an Aigéin Chiúin", "PWT": "Am Oileáin Palau", "PYST": "Am Samhraidh Pharagua", "PYT": "Am Caighdeánach Pharagua", "PYT Korea": "Am Pyongyang", "RET": "Am Réunion", "ROTT": "Am Rothera", "SAKST": "Am Samhraidh Shacailín", "SAKT": "Am Caighdeánach Shacailín", "SAMST": "Am Samhraidh Shamara", "SAMT": "Am Caighdeánach Shamara", "SAST": "Am na hAfraice Theas", "SBT": "Am Oileáin Sholaimh", "SCT": "Am na Séiséal", "SGT": "Am Caighdeánach Shingeapór", "SLST": "Am Shrí Lanca", "SRT": "Am Shuranam", "SST Samoa": "Am Caighdeánach Shamó Mheiriceá", "SST Samoa Apia": "Am Caighdeánach Shamó", "SST Samoa Apia DST": "Am Samhraidh Shamó", "SST Samoa DST": "Am Samhraidh Shamó Mheiriceá", "SYOT": "Am Syowa", "TAAF": "Am Francach Dheisceart an Domhain agus an Antartaigh", "TAHT": "Am Thaihítí", "TJT": "Am na Táidsíceastáine", "TKT": "Am Oileáin Thócaláú", "TLT": "Am Thíomór Thoir", "TMST": "Am Samhraidh na Tuircméanastáine", "TMT": "Am Caighdeánach na Tuircméanastáine", "TOST": "Am Samhraidh Thonga", "TOT": "Am Caighdeánach Thonga", "TVT": "Am Thúvalú", "TWT": "Am Caighdeánach Taipei", "TWT DST": "Am Samhraidh Taipei", "ULAST": "Am Samhraidh Ulan Bator", "ULAT": "Am Caighdeánach Ulan Bator", "UYST": "Am Samhraidh Uragua", "UYT": "Am Caighdeánach Uragua", "UZT": "Am Caighdeánach na hÚisbéiceastáine", "UZT DST": "Am Samhraidh na hÚisbéiceastáine", "VET": "Am Veiniséala", "VLAST": "Am Samhraidh Vladivostok", "VLAT": "Am Caighdeánach Vladivostok", "VOLST": "Am Samhraidh Volgograd", "VOLT": "Am Caighdeánach Volgograd", "VOST": "Am Vostok", "VUT": "Am Caighdeánach Vanuatú", "VUT DST": "Am Samhraidh Vanuatú", "WAKT": "Am Oileán Wake", "WARST": "Am Samhraidh Iartharach na hAirgintíne", "WART": "Am Caighdeánach Iartharach na hAirgintíne", "WAST": "Am Iarthar na hAfraice", "WAT": "Am Iarthar na hAfraice", "WESZ": "Am Samhraidh Iarthar na hEorpa", "WEZ": "Am Caighdeánach Iarthar na hEorpa", "WFT": "Am Wallis agus Futuna", "WGST": "Am Samhraidh Iarthar na Graonlainne", "WGT": "Am Caighdeánach Iarthar na Graonlainne", "WIB": "Am Iarthar na hIndinéise", "WIT": "Am Oirthear na hIndinéise", "WITA": "Am Lár na hIndinéise", "YAKST": "Am Samhraidh Iacútsc", "YAKT": "Am Caighdeánach Iacútsc", "YEKST": "Am Samhraidh Yekaterinburg", "YEKT": "Am Caighdeánach Yekaterinburg", "YST": "Am Yukon", "МСК": "Am Caighdeánach Mhoscó", "اقتاۋ": "Am Caighdeánach Aqtau", "اقتاۋ قالاسى": "Am Samhraidh Aqtau", "اقتوبە": "Am Caighdeánach Aqtobe", "اقتوبە قالاسى": "Am Samhraidh Aqtobe", "الماتى": "Am Caighdeánach Almaty", "الماتى قالاسى": "Am Samhraidh Almaty", "باتىس قازاق ەلى": "Am Iarthar na Casacstáine", "شىعىش قازاق ەلى": "Am Oirthear na Casacstáine", "قازاق ەلى": "Am na Casacstáine", "قىرعىزستان": "Am na Cirgeastáine", "قىزىلوردا": "Am Caighdeánach Qyzylorda", "قىزىلوردا قالاسى": "Am Samhraidh Qyzylorda", "∅∅∅": "Am Samhraidh Pheiriú"},
	}
}

// Locale returns the current translators string locale
func (ga *ga_GB) Locale() string {
	return ga.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ga_GB'
func (ga *ga_GB) PluralsCardinal() []locales.PluralRule {
	return ga.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ga_GB'
func (ga *ga_GB) PluralsOrdinal() []locales.PluralRule {
	return ga.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ga_GB'
func (ga *ga_GB) PluralsRange() []locales.PluralRule {
	return ga.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ga_GB'
func (ga *ga_GB) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	} else if n >= 3 && n <= 6 {
		return locales.PluralRuleFew
	} else if n >= 7 && n <= 10 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ga_GB'
func (ga *ga_GB) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ga_GB'
func (ga *ga_GB) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	start := ga.CardinalPluralRule(num1, v1)
	end := ga.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleTwo {
		return locales.PluralRuleTwo
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleTwo && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleTwo && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleTwo && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleMany && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleTwo {
		return locales.PluralRuleTwo
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ga *ga_GB) MonthAbbreviated(month time.Month) string {
	return ga.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ga *ga_GB) MonthsAbbreviated() []string {
	return ga.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ga *ga_GB) MonthNarrow(month time.Month) string {
	return ga.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ga *ga_GB) MonthsNarrow() []string {
	return ga.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ga *ga_GB) MonthWide(month time.Month) string {
	return ga.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ga *ga_GB) MonthsWide() []string {
	return ga.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ga *ga_GB) WeekdayAbbreviated(weekday time.Weekday) string {
	return ga.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ga *ga_GB) WeekdaysAbbreviated() []string {
	return ga.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ga *ga_GB) WeekdayNarrow(weekday time.Weekday) string {
	return ga.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ga *ga_GB) WeekdaysNarrow() []string {
	return ga.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ga *ga_GB) WeekdayShort(weekday time.Weekday) string {
	return ga.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ga *ga_GB) WeekdaysShort() []string {
	return ga.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ga *ga_GB) WeekdayWide(weekday time.Weekday) string {
	return ga.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ga *ga_GB) WeekdaysWide() []string {
	return ga.daysWide
}

// Decimal returns the decimal point of number
func (ga *ga_GB) Decimal() string {
	return ga.decimal
}

// Group returns the group of number
func (ga *ga_GB) Group() string {
	return ga.group
}

// Group returns the minus sign of number
func (ga *ga_GB) Minus() string {
	return ga.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ga_GB' and handles both Whole and Real numbers based on 'v'
func (ga *ga_GB) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ga.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ga.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ga.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ga_GB' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ga *ga_GB) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ga.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ga.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ga.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ga_GB'
func (ga *ga_GB) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ga.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ga.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ga.group[0])
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
		b = append(b, ga.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ga.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ga_GB'
// in accounting notation.
func (ga *ga_GB) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ga.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ga.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ga.group[0])
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

		b = append(b, ga.currencyNegativePrefix[0])

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
			b = append(b, ga.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ga.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ga_GB'
func (ga *ga_GB) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'ga_GB'
func (ga *ga_GB) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ga.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ga_GB'
func (ga *ga_GB) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ga.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ga_GB'
func (ga *ga_GB) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, ga.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ga.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ga_GB'
func (ga *ga_GB) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ga.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ga_GB'
func (ga *ga_GB) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ga.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ga.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ga_GB'
func (ga *ga_GB) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ga.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ga.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ga_GB'
func (ga *ga_GB) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ga.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ga.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ga.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
