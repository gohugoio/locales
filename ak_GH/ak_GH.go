package ak_GH

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ak_GH struct {
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
	periodsAbbreviated     []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'ak_GH' locale
func New() locales.Translator {
	return &ak_GH{
		locale:                 "ak_GH",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsNarrow:           []string{"", "Ɔp", "Ɔg", "Ɔb", "O", "K", "A", "Ku", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "Ɔpɛpɔn", "Ɔgyefoɔ", "Ɔbɛnem", "Oforisuo", "Kɔtɔnimma", "Ayɛwohomumu", "Kutawonsa", "Ɔsanaa", "Ɛbɔ", "Ahinime", "Obubuo", "Ɔpɛnimma"},
		daysAbbreviated:        []string{"Kwa", "Dwo", "Ben", "Wuk", "Yaw", "Fia", "Mem"},
		daysNarrow:             []string{"K", "D", "B", "W", "Y", "F", "M"},
		daysShort:              []string{"Kwa", "Dwo", "Ben", "Wuk", "Yaw", "Fia", "Mem"},
		daysWide:               []string{"Sun", "Dwoada", "Benada", "Wukuada", "Yawoada", "Fiada", "Memeneda"},
		periodsAbbreviated:     []string{"AN", "ANW"},
		timezones:              map[string]string{"ACDT": "Ɔstrelia Mfinimfini Awia Berɛ", "ACST": "Ɔstrelia Mfinimfini Susudua Berɛ", "ACT": "ACT", "ACWDT": "Ɔstrelia Mfinimfini Atɔeeɛ Awia Berɛ", "ACWST": "Ɔstrelia Mfinimfini Atɔeeɛ Susudua Berɛ", "ADT": "Atlantik Awia Berɛ", "ADT Arabia": "Arabia Awia Berɛ", "AEDT": "Ɔstrelia Apueeɛ Awia Berɛ", "AEST": "Ɔstrelia Apueeɛ Susudua Berɛ", "AFT": "Afganistan Berɛ", "AKDT": "Alaska Awia Berɛ", "AKST": "Alaska Susudua Berɛ", "AMST": "Amazon Awia Berɛ", "AMST Armenia": "Aamenia Awia Berɛ", "AMT": "Amazon Susudua Berɛ", "AMT Armenia": "Aamenia Susudua Berɛ", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Agyɛntina Awia Berɛ", "ART": "Agyɛntina Susudua Berɛ", "AST": "Atlantik Susudua Berɛ", "AST Arabia": "Arabia Susudua Berɛ", "AWDT": "Ɔstrelia Atɔeeɛ Awia Berɛ", "AWST": "Ɔstrelia Atɔeeɛ Susudua Berɛ", "AZST": "Asabegyan Awia Berɛ", "AZT": "Asabegyan Susudua Berɛ", "BDT Bangladesh": "Bangladɛhye Awia Berɛ", "BNT": "Brunei Berɛ", "BOT": "Bolivia Berɛ", "BRST": "Brasilia Awia Berɛ", "BRT": "Brasilia Susudua Berɛ", "BST Bangladesh": "Bangladɛhye Susudua Berɛ", "BT": "Butan Berɛ", "CAST": "CAST", "CAT": "Afrika Finimfin Berɛ", "CCT": "Kokoso Aeland Berɛ", "CDT": "Mfinimfini Awia Berɛ", "CHADT": "Kyatam Awia Berɛ", "CHAST": "Kyatam Susudua Berɛ", "CHUT": "Kyuuk Berɛ", "CKT": "Kuk Aeland Susudua Berɛ", "CKT DST": "Kuk Aeland Awia Fa Berɛ", "CLST": "Kyili Awia Berɛ", "CLT": "Kyili Susudua Berɛ", "COST": "Kolombia Awia Berɛ", "COT": "Kolombia Susudua Berɛ", "CST": "Mfinimfini Susudua Berɛ", "CST China": "Kyaena Susudua Berɛ", "CST China DST": "Kyaena Awia Berɛ", "CVST": "Kepvɛde Awia Berɛ", "CVT": "Kepvɛde Susudua Berɛ", "CXT": "Buronya Aeland Berɛ", "ChST": "Kyamoro Susudua Berɛ", "ChST NMI": "ChST NMI", "CuDT": "Kuba Awia Berɛ", "CuST": "Kuba Susudua Berɛ", "DAVT": "Davis Berɛ", "DDUT": "Dumont-d’Urville Berɛ", "EASST": "Easta Aeland Awia Berɛ", "EAST": "Easta Aeland Susudua Berɛ", "EAT": "Afrika Apueeɛ Berɛ", "ECT": "Yikuwedɔ Berɛ", "EDT": "Apueeɛ Awia Berɛ", "EGDT": "Greenland Apueeɛ Awia Berɛ", "EGST": "Greenland Apueeɛ Susudua Berɛ", "EST": "Apueeɛ Susudua Berɛ", "FEET": "Yuropu Apueeɛ Nohoa Berɛ", "FJT": "Figyi Susudua Berɛ", "FJT Summer": "Figyi Awia Berɛ", "FKST": "Fɔkman Aeland Awia Berɛ", "FKT": "Fɔkman Aeland Susudua Berɛ", "FNST": "Fernando de Noronha Awia Berɛ", "FNT": "Fernando de Noronha Susudua Berɛ", "GALT": "Galapagɔs Berɛ", "GAMT": "Gambier Berɛ", "GEST": "Gyɔgyea Awia Berɛ", "GET": "Gyɔgyea Susudua Berɛ", "GFT": "Frɛnkye Gayana Berɛ", "GIT": "Geebɛt Aeland Berɛ", "GMT": "Greenwich Mean Berɛ", "GNSST": "GNSST", "GNST": "GNST", "GST": "Gyɔɔgyia Anaafoɔ Berɛ", "GST Guam": "GST Guam", "GYT": "Gayana Berɛ", "HADT": "Hawaii-Aleutian Awia Berɛ", "HAST": "Hawaii-Aleutian Susudua Berɛ", "HKST": "Hɔnkɔn Awia Berɛ", "HKT": "Hɔnkɔn Susudua Berɛ", "HOVST": "Hovd Awia Berɛ", "HOVT": "Hovd Susudua Berɛ", "ICT": "Indɔkyina Berɛ", "IDT": "Israel Awia Berɛ", "IOT": "India Po Berɛ", "IRKST": "Yiikusk Awia Berɛ", "IRKT": "Yiikusk Susudua Berɛ", "IRST": "Iran Susudua Berɛ", "IRST DST": "Iran Awia Berɛ", "IST": "India Susudua Berɛ", "IST Israel": "Israel Susudua Berɛ", "JDT": "Gyapan Awia Berɛ", "JST": "Gyapan Susudua Berɛ", "KOST": "Kosrae Berɛ", "KRAST": "Krasnoyarsk Awia Berɛ", "KRAT": "Krasnoyarsk Susudua Berɛ", "KST": "Korean Susudua Berɛ", "KST DST": "Korean Awia Berɛ", "LHDT": "Lɔd Howe Awia Berɛ", "LHST": "Lɔd Howe Susudua Berɛ", "LINT": "Lai Aeland Berɛ", "MAGST": "Magadan Awia Berɛ", "MAGT": "Magadan Susudua Berɛ", "MART": "Makesase Berɛ", "MAWT": "Mɔɔson Berɛ", "MDT": "MDT", "MESZ": "Yuropu Mfinimfini Awia Berɛ", "MEZ": "Yuropu Mfinimfini Susudua Berɛ", "MHT": "Mahyaa Aeland Berɛ", "MMT": "Mayaama Berɛ", "MSD": "Mɔsko Awia Berɛ", "MST": "MST", "MUST": "Mɔrihyiɔso Awia Berɛ", "MUT": "Mɔrihyiɔso Susudua Berɛ", "MVT": "Maldives Berɛ", "MYT": "Malehyia Berɛ", "NCT": "Kaledonia Foforɔ Susudua Berɛ", "NDT": "Newfoundland Awia Berɛ", "NDT New Caledonia": "Kaledonia Foforɔ Awia Berɛ", "NFDT": "Nɔɔfɔk Aeland Awia Berɛ", "NFT": "Nɔɔfɔk Aeland Susudua Berɛ", "NOVST": "Novosibirsk Awia Berɛ", "NOVT": "Novosibirsk Susudua Berɛ", "NPT": "Nɛpal Berɛ", "NRT": "Nauru Berɛ", "NST": "Newfoundland Susudua Berɛ", "NUT": "Niue Berɛ", "NZDT": "Ziland Foforɔ Awia Berɛ", "NZST": "Ziland Foforɔ Susudua Berɛ", "OESZ": "Yuropu Apueeɛ Awia Berɛ", "OEZ": "Yuropu Apueeɛ Susudua Berɛ", "OMSST": "Omsk Awia Berɛ", "OMST": "Omsk Susudua Berɛ", "PDT": "Pasifik Awia Berɛ", "PDTM": "Mɛksiko Pasifik Awia Berɛ", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Papua Gini Foforɔ Berɛ", "PHOT": "Finise Aeland Berɛ", "PKT": "Pakistan Susudua Berɛ", "PKT DST": "Pakistan Awia Berɛ", "PMDT": "St. Pierre & Miquelon Awia Berɛ", "PMST": "St. Pierre & Miquelon Susudua Berɛ", "PONT": "Ponape Berɛ", "PST": "Pasifik Susudua Berɛ", "PST Philippine": "Filipin Susudua Berɛ", "PST Philippine DST": "Filipin Awia Berɛ", "PST Pitcairn": "Pitkairn Berɛ", "PSTM": "Mɛksiko Pasifik Susudua Berɛ", "PWT": "Palau Berɛ", "PYST": "Paraguae Awia Berɛ", "PYT": "Paraguae Susudua Berɛ", "PYT Korea": "Pyongyang Berɛ", "RET": "Réunion Berɛ", "ROTT": "Rotera Berɛ", "SAKST": "Sakhalin Awia Berɛ", "SAKT": "Sakhalin Susudua Berɛ", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Afrika Anaafoɔ Susudua Berɛ", "SBT": "Solomon Aeland Berɛ", "SCT": "Seyhyɛl Berɛ", "SGT": "Singapɔ Susudua Berɛ", "SLST": "SLST", "SRT": "Suriname Berɛ", "SST Samoa": "Samoa Susudua Berɛ", "SST Samoa Apia": "Apia Susudua Berɛ", "SST Samoa Apia DST": "Apia Awia Berɛ", "SST Samoa DST": "Samoa Awia Berɛ", "SYOT": "Syowa Berɛ", "TAAF": "Frɛnkye Anaafoɔ ne Antaatik Berɛ", "TAHT": "Tahiti Berɛ", "TJT": "Tagyikistan Berɛ", "TKT": "Tokelau Berɛ", "TLT": "Timɔɔ Apueeɛ Berɛ", "TMST": "Tɛkmɛnistan Awia Berɛ", "TMT": "Tɛkmɛnistan Susudua Berɛ", "TOST": "Tonga Awia Berɛ", "TOT": "Tonga Susudua Berɛ", "TVT": "Tuvalu Berɛ", "TWT": "Taipei Susudua Berɛ", "TWT DST": "Taipei Awia Berɛ", "ULAST": "Yulanbata Awia Berɛ", "ULAT": "Yulanbata Susudua Berɛ", "UYST": "Yurugwae Awia Berɛ", "UYT": "Yurugwae Susudua Berɛ", "UZT": "Usbɛkistan Susudua Berɛ", "UZT DST": "Usbɛkistan Awia Berɛ", "VET": "Venezuela Berɛ", "VLAST": "Vladivostok Awia Berɛ", "VLAT": "Vladivostok Susudua Berɛ", "VOLST": "Volgograd Awia Berɛ", "VOLT": "Volgograd Susudua Berɛ", "VOST": "Vostok Berɛ", "VUT": "Vanuatu Susudua Berɛ", "VUT DST": "Vanuatu Awia Berɛ", "WAKT": "Wake Aeland Berɛ", "WARST": "Agyɛntina Atɔeeɛ Awia Berɛ", "WART": "Agyɛntina Atɔeeɛ Susudua Berɛ", "WAST": "Afrika Atɔeɛ Berɛ", "WAT": "Afrika Atɔeɛ Berɛ", "WESZ": "Yuropu Atɔeeɛ Awia Berɛ", "WEZ": "Yuropu Atɔeeɛ Susudua Berɛ", "WFT": "Wallis ne Futuna Berɛ", "WGST": "Greenland Atɔeɛ Awia Berɛ", "WGT": "Greenland Atɔeɛ Susudua Berɛ", "WIB": "Indɔnehyia Atɔeeɛ Berɛ", "WIT": "Indɔnehyia Apueeɛ Berɛ", "WITA": "Indɔnehyia Mfinimfini Berɛ", "YAKST": "Yakutsk Awia Berɛ", "YAKT": "Yakutsk Susudua Berɛ", "YEKST": "Yɛkatɛrinbɛg Awia Berɛ", "YEKT": "Yɛkatɛrinbɛg Susudua Berɛ", "YST": "Yukɔn Berɛ", "МСК": "Mɔsko Susudua Berɛ", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Kazakstan Atɔeɛ Berɛ", "شىعىش قازاق ەلى": "Kazakstan Apueeɛ Berɛ", "قازاق ەلى": "Kazakstan Berɛ", "قىرعىزستان": "Kɛɛgestan Berɛ", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Azores Awia Berɛ"},
	}
}

// Locale returns the current translators string locale
func (ak *ak_GH) Locale() string {
	return ak.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ak_GH'
func (ak *ak_GH) PluralsCardinal() []locales.PluralRule {
	return ak.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ak_GH'
func (ak *ak_GH) PluralsOrdinal() []locales.PluralRule {
	return ak.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ak_GH'
func (ak *ak_GH) PluralsRange() []locales.PluralRule {
	return ak.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ak_GH'
func (ak *ak_GH) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ak_GH'
func (ak *ak_GH) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ak_GH'
func (ak *ak_GH) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	start := ak.CardinalPluralRule(num1, v1)
	end := ak.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ak *ak_GH) MonthAbbreviated(month time.Month) string {
	if len(ak.monthsAbbreviated) == 0 {
		return ""
	}
	return ak.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ak *ak_GH) MonthsAbbreviated() []string {
	return nil
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ak *ak_GH) MonthNarrow(month time.Month) string {
	if len(ak.monthsNarrow) == 0 {
		return ""
	}
	return ak.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ak *ak_GH) MonthsNarrow() []string {
	return ak.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ak *ak_GH) MonthWide(month time.Month) string {
	if len(ak.monthsWide) == 0 {
		return ""
	}
	return ak.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ak *ak_GH) MonthsWide() []string {
	return ak.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ak *ak_GH) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(ak.daysAbbreviated) == 0 {
		return ""
	}
	return ak.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ak *ak_GH) WeekdaysAbbreviated() []string {
	return ak.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ak *ak_GH) WeekdayNarrow(weekday time.Weekday) string {
	if len(ak.daysNarrow) == 0 {
		return ""
	}
	return ak.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ak *ak_GH) WeekdaysNarrow() []string {
	return ak.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ak *ak_GH) WeekdayShort(weekday time.Weekday) string {
	if len(ak.daysShort) == 0 {
		return ""
	}
	return ak.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ak *ak_GH) WeekdaysShort() []string {
	return ak.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ak *ak_GH) WeekdayWide(weekday time.Weekday) string {
	if len(ak.daysWide) == 0 {
		return ""
	}
	return ak.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ak *ak_GH) WeekdaysWide() []string {
	return ak.daysWide
}

// Decimal returns the decimal point of number
func (ak *ak_GH) Decimal() string {
	return ak.decimal
}

// Group returns the group of number
func (ak *ak_GH) Group() string {
	return ak.group
}

// Group returns the minus sign of number
func (ak *ak_GH) Minus() string {
	return ak.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ak_GH' and handles both Whole and Real numbers based on 'v'
func (ak *ak_GH) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ak.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ak.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ak.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ak_GH' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ak *ak_GH) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ak.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ak.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ak.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ak_GH'
func (ak *ak_GH) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ak.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ak.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ak.group[0])
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
		b = append(b, ak.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ak.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ak_GH'
// in accounting notation.
func (ak *ak_GH) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ak.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ak.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ak.group[0])
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

		b = append(b, ak.currencyNegativePrefix[0])

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
			b = append(b, ak.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ak.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ak_GH'
func (ak *ak_GH) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'ak_GH'
func (ak *ak_GH) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, ak.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ak_GH'
func (ak *ak_GH) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, ak.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ak_GH'
func (ak *ak_GH) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, ak.daysAbbreviated[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, ak.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ak_GH'
func (ak *ak_GH) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ak.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, ak.periodsAbbreviated[0]...)
	} else {
		b = append(b, ak.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ak_GH'
func (ak *ak_GH) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ak.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ak.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, ak.periodsAbbreviated[0]...)
	} else {
		b = append(b, ak.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ak_GH'
func (ak *ak_GH) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ak.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ak.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, ak.periodsAbbreviated[0]...)
	} else {
		b = append(b, ak.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ak_GH'
func (ak *ak_GH) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ak.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ak.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, ak.periodsAbbreviated[0]...)
	} else {
		b = append(b, ak.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ak.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
