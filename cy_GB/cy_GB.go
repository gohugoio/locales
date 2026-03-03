package cy_GB

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type cy_GB struct {
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

// New returns a new instance of translator for the 'cy_GB' locale
func New() locales.Translator {
	return &cy_GB{
		locale:                 "cy_GB",
		pluralsCardinal:        []locales.PluralRule{1, 2, 3, 4, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{1, 2, 3, 4, 5, 6},
		pluralsRange:           []locales.PluralRule{2, 3, 4, 5, 6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "TK", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HK$", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "Ion", "Chwef", "Maw", "Ebr", "Mai", "Meh", "Gorff", "Awst", "Medi", "Hyd", "Tach", "Rhag"},
		monthsNarrow:           []string{"", "I", "Ch", "M", "E", "M", "M", "G", "A", "M", "H", "T", "Rh"},
		monthsWide:             []string{"", "Ionawr", "Chwefror", "Mawrth", "Ebrill", "Mai", "Mehefin", "Gorffennaf", "Awst", "Medi", "Hydref", "Tachwedd", "Rhagfyr"},
		daysAbbreviated:        []string{"Sul", "Llun", "Maw", "Mer", "Iau", "Gwen", "Sad"},
		daysNarrow:             []string{"S", "Ll", "M", "M", "I", "G", "S"},
		daysShort:              []string{"Su", "Ll", "Ma", "Me", "Ia", "Gw", "Sa"},
		daysWide:               []string{"Dydd Sul", "Dydd Llun", "Dydd Mawrth", "Dydd Mercher", "Dydd Iau", "Dydd Gwener", "Dydd Sadwrn"},
		timezones:              map[string]string{"ACDT": "Amser Haf Canolbarth Awstralia", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Amser Haf Canolbarth Gorllewin Awstralia", "ACWST": "Amser Safonol Canolbarth Gorllewin Awstralia", "ADT": "Amser Haf Cefnfor yr Iwerydd", "ADT Arabia": "Amser Haf Arabaidd", "AEDT": "Amser Haf Dwyrain Awstralia", "AEST": "Amser Safonol Dwyrain Awstralia", "AFT": "Amser Afghanistan", "AKDT": "Amser Haf Alaska", "AKST": "Amser Safonol Alaska", "AMST": "Amser Haf Amazonas", "AMST Armenia": "Amser Haf Armenia", "AMT": "Amser Safonol Amazonas", "AMT Armenia": "Amser Safonol Armenia", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Amser Haf Ariannin", "ART": "Amser Safonol Ariannin", "AST": "Amser Safonol Cefnfor yr Iwerydd", "AST Arabia": "Amser Safonol Arabaidd", "AWDT": "Amser Haf Gorllewin Awstralia", "AWST": "Amser Safonol Gorllewin Awstralia", "AZST": "Amser Haf Aserbaijan", "AZT": "Amser Safonol Aserbaijan", "BDT Bangladesh": "Amser Haf Bangladesh", "BNT": "Amser Brunei Darussalam", "BOT": "Amser Bolifia", "BRST": "Amser Haf Brasília", "BRT": "Amser Safonol Brasília", "BST Bangladesh": "Amser Safonol Bangladesh", "BT": "Amser Bhutan", "CAST": "CAST", "CAT": "Amser Canolbarth Affrica", "CCT": "Amser Ynysoedd Cocos", "CDT": "Amser Haf Canolbarth Gogledd America", "CHADT": "Amser Haf Chatham", "CHAST": "Amser Safonol Chatham", "CHUT": "Amser Chuuk", "CKT": "Amser Safonol Ynysoedd Cook", "CKT DST": "Amser Hanner Haf Ynysoedd Cook", "CLST": "Amser Haf Chile", "CLT": "Amser Safonol Chile", "COST": "Amser Haf Colombia", "COT": "Amser Safonol Colombia", "CST": "Amser Safonol Canolbarth Gogledd America", "CST China": "Amser Safonol Tsieina", "CST China DST": "Amser Haf Tsieina", "CVST": "Amser Haf Cabo Verde", "CVT": "Amser Safonol Cabo Verde", "CXT": "Amser Ynys Y Nadolig", "ChST": "Amser Chamorro", "ChST NMI": "ChST NMI", "CuDT": "Amser Haf Ciwa", "CuST": "Amser Safonol Ciwba", "DAVT": "Amser Davis", "DDUT": "Amser Dumont-d’Urville", "EASST": "Amser Haf Ynys y Pasg", "EAST": "Amser Safonol Ynys y Pasg", "EAT": "Amser Dwyrain Affrica", "ECT": "Amser Ecuador", "EDT": "Amser Haf Dwyrain Gogledd America", "EGDT": "Amser Haf Dwyrain yr Ynys Las", "EGST": "Amser Safonol Dwyrain yr Ynys Las", "EST": "Amser Safonol Dwyrain Gogledd America", "FEET": "Amser Dwyrain Pell Ewrop", "FJT": "Amser Safonol Fiji", "FJT Summer": "Amser Haf Fiji", "FKST": "Amser Haf Ynysoedd Falklands/Malvinas", "FKT": "Amser Safonol Ynysoedd Falklands/Malvinas", "FNST": "Amser Haf Fernando de Noronha", "FNT": "Amser Safonol Fernando de Noronha", "GALT": "Amser Galapagos", "GAMT": "Amser Gambier", "GEST": "Amser Haf Georgia", "GET": "Amser Safonol Georgia", "GFT": "Amser Guyane Ffrengig", "GIT": "Amser Ynysoedd Gilbert", "GMT": "Amser Safonol Greenwich", "GNSST": "GNSST", "GNST": "GNST", "GST": "Amser De Georgia", "GST Guam": "GST Guam", "GYT": "Amser Guyana", "HADT": "Amser Haf Hawaii-Aleutian", "HAST": "Amser Safonol Hawaii-Aleutian", "HKST": "Amser Haf Hong Kong", "HKT": "Amser Safonol Hong Kong", "HOVST": "Amser Haf Hovd", "HOVT": "Amser Safonol Hovd", "ICT": "Amser Indo-Tsieina", "IDT": "Amser Haf Israel", "IOT": "Amser Cefnfor India", "IRKST": "Amser Haf Irkutsk", "IRKT": "Amser Safonol Irkutsk", "IRST": "Amser Safonol Iran", "IRST DST": "Amser Haf Iran", "IST": "Amser India", "IST Israel": "Amser Safonol Israel", "JDT": "Amser Haf Japan", "JST": "Amser Safonol Japan", "KOST": "Amser Kosrae", "KRAST": "Amser Haf Krasnoyarsk", "KRAT": "Amser Safonol Krasnoyarsk", "KST": "Amser Safonol Corea", "KST DST": "Amser Haf Corea", "LHDT": "Amser Haf yr Arglwydd Howe", "LHST": "Amser Safonol yr Arglwydd Howe", "LINT": "Amser Ynysoedd Line", "MAGST": "Amser Haf Magadan", "MAGT": "Amser Safonol Magadan", "MART": "Amser Marquises", "MAWT": "Amser Mawson", "MDT": "MDT", "MESZ": "Amser Haf Canolbarth Ewrop", "MEZ": "Amser Safonol Canolbarth Ewrop", "MHT": "Amser Ynysoedd Marshall", "MMT": "Amser Myanmar", "MSD": "Amser Haf Moscfa", "MST": "MST", "MUST": "Amser Haf Mauritius", "MUT": "Amser Safonol Mauritius", "MVT": "Amser Y Maldives", "MYT": "Amser Malaysia", "NCT": "Amser Safonol Caledonia Newydd", "NDT": "Amser Haf Newfoundland", "NDT New Caledonia": "Amser Haf Caledonia Newydd", "NFDT": "Amser Haf Ynys Norfolk", "NFT": "Amser Safonol Ynys Norfolk", "NOVST": "Amser Haf Novosibirsk", "NOVT": "Amser Safonol Novosibirsk", "NPT": "Amser Nepal", "NRT": "Amser Nauru", "NST": "Amser Safonol Newfoundland", "NUT": "Amser Niue", "NZDT": "Amser Haf Seland Newydd", "NZST": "Amser Safonol Seland Newydd", "OESZ": "Amser Haf Dwyrain Ewrop", "OEZ": "Amser Safonol Dwyrain Ewrop", "OMSST": "Amser Haf Omsk", "OMST": "Amser Safonol Omsk", "PDT": "Amser Haf Cefnfor Tawel Gogledd America", "PDTM": "Amser Haf Pasiffig Mecsico", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Amser Papua Guinea Newydd", "PHOT": "Amser Ynysoedd Phoenix", "PKT": "Amser Safonol Pakistan", "PKT DST": "Amser Haf Pakistan", "PMDT": "Amser Haf Saint-Pierre-et-Miquelon", "PMST": "Amser Safonol Saint-Pierre-et-Miquelon", "PONT": "Amser Pohnpei", "PST": "Amser Safonol Cefnfor Tawel Gogledd America", "PST Philippine": "Amser Safonol Pilipinas", "PST Philippine DST": "Amser Haf Pilipinas", "PST Pitcairn": "Amser Pitcairn", "PSTM": "Amser Safonol Pasiffig Mecsico", "PWT": "Amser Palau", "PYST": "Amser Haf Paraguay", "PYT": "Amser Safonol Paraguay", "PYT Korea": "Amser Pyongyang", "RET": "Amser Réunion", "ROTT": "Amser Rothera", "SAKST": "Amser Haf Sakhalin", "SAKT": "Amser Safonol Sakhalin", "SAMST": "Amser Haf Samara", "SAMT": "Amser Safonol Samara", "SAST": "Amser Safonol De Affrica", "SBT": "Amser Ynysoedd Solomon", "SCT": "Amser Seychelles", "SGT": "Amser Singapore", "SLST": "SLST", "SRT": "Amser Suriname", "SST Samoa": "Amser Safonol Samoa", "SST Samoa Apia": "Amser Safonol Apia", "SST Samoa Apia DST": "Amser Haf Apia", "SST Samoa DST": "Amser Haf Samoa", "SYOT": "Amser Syowa", "TAAF": "Amser Tiroedd Ffrainc yn y De a’r Antarctig", "TAHT": "Amser Tahiti", "TJT": "Amser Tajicistan", "TKT": "Amser Tokelau", "TLT": "Amser Dwyrain Timor", "TMST": "Amser Haf Tyrcmenistan", "TMT": "Amser Safonol Tyrcmenistan", "TOST": "Amser Haf Tonga", "TOT": "Amser Safonol Tonga", "TVT": "Amser Tuvalu", "TWT": "Amser Safonol Taipei", "TWT DST": "Amser Haf Taipei", "ULAST": "Amser Haf Ulan Bator", "ULAT": "Amser Safonol Ulan Bator", "UYST": "Amser Haf Uruguay", "UYT": "Amser Safonol Uruguay", "UZT": "Amser Safonol Uzbekistan", "UZT DST": "Amser Haf Uzbekistan", "VET": "Amser Venezuela", "VLAST": "Amser Haf Vladivostok", "VLAT": "Amser Safonol Vladivostok", "VOLST": "Amser Haf Volgograd", "VOLT": "Amser Safonol Volgograd", "VOST": "Amser Vostok", "VUT": "Amser Safonol Vanuatu", "VUT DST": "Amser Haf Vanuatu", "WAKT": "Amser Ynys Wake", "WARST": "Amser Haf Gorllewin Ariannin", "WART": "Amser Safonol Gorllewin Ariannin", "WAST": "Amser Gorllewin Affrica", "WAT": "Amser Gorllewin Affrica", "WESZ": "Amser Haf Gorllewin Ewrop", "WEZ": "Amser Safonol Gorllewin Ewrop", "WFT": "Amser Wallis a Futuna", "WGST": "Amser Haf Gorllewin yr Ynys Las", "WGT": "Amser Safonol Gorllewin yr Ynys Las", "WIB": "Amser Gorllewin Indonesia", "WIT": "Amser Dwyrain Indonesia", "WITA": "Amser Canolbarth Indonesia", "YAKST": "Amser Haf Yakutsk", "YAKT": "Amser Safonol Yakutsk", "YEKST": "Amser Haf Yekaterinburg", "YEKT": "Amser Safonol Yekaterinburg", "YST": "Amser Yukon", "МСК": "Amser Safonol Moscfa", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Amser Gorllewin Kazakhstan", "شىعىش قازاق ەلى": "Amser Dwyrain Kazakhstan", "قازاق ەلى": "Amser Kazakhstan", "قىرعىزستان": "Amser Kyrgyzstan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Amser Haf Periw"},
	}
}

// Locale returns the current translators string locale
func (cy *cy_GB) Locale() string {
	return cy.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'cy_GB'
func (cy *cy_GB) PluralsCardinal() []locales.PluralRule {
	return cy.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'cy_GB'
func (cy *cy_GB) PluralsOrdinal() []locales.PluralRule {
	return cy.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'cy_GB'
func (cy *cy_GB) PluralsRange() []locales.PluralRule {
	return cy.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'cy_GB'
func (cy *cy_GB) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 0 {
		return locales.PluralRuleZero
	} else if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	} else if n == 3 {
		return locales.PluralRuleFew
	} else if n == 6 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'cy_GB'
func (cy *cy_GB) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 0 || n == 7 || n == 8 || n == 9 {
		return locales.PluralRuleZero
	} else if n == 1 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	} else if n == 3 || n == 4 {
		return locales.PluralRuleFew
	} else if n == 5 || n == 6 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'cy_GB'
func (cy *cy_GB) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := cy.CardinalPluralRule(num1, v1)
	end := cy.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleZero && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleTwo {
		return locales.PluralRuleTwo
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleZero && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleTwo {
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
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleMany {
		return locales.PluralRuleMany
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
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
func (cy *cy_GB) MonthAbbreviated(month time.Month) string {
	if len(cy.monthsAbbreviated) == 0 {
		return ""
	}
	return cy.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (cy *cy_GB) MonthsAbbreviated() []string {
	return cy.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (cy *cy_GB) MonthNarrow(month time.Month) string {
	if len(cy.monthsNarrow) == 0 {
		return ""
	}
	return cy.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (cy *cy_GB) MonthsNarrow() []string {
	return cy.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (cy *cy_GB) MonthWide(month time.Month) string {
	if len(cy.monthsWide) == 0 {
		return ""
	}
	return cy.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (cy *cy_GB) MonthsWide() []string {
	return cy.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (cy *cy_GB) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(cy.daysAbbreviated) == 0 {
		return ""
	}
	return cy.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (cy *cy_GB) WeekdaysAbbreviated() []string {
	return cy.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (cy *cy_GB) WeekdayNarrow(weekday time.Weekday) string {
	if len(cy.daysNarrow) == 0 {
		return ""
	}
	return cy.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (cy *cy_GB) WeekdaysNarrow() []string {
	return cy.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (cy *cy_GB) WeekdayShort(weekday time.Weekday) string {
	if len(cy.daysShort) == 0 {
		return ""
	}
	return cy.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (cy *cy_GB) WeekdaysShort() []string {
	return cy.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (cy *cy_GB) WeekdayWide(weekday time.Weekday) string {
	if len(cy.daysWide) == 0 {
		return ""
	}
	return cy.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (cy *cy_GB) WeekdaysWide() []string {
	return cy.daysWide
}

// Decimal returns the decimal point of number
func (cy *cy_GB) Decimal() string {
	return cy.decimal
}

// Group returns the group of number
func (cy *cy_GB) Group() string {
	return cy.group
}

// Group returns the minus sign of number
func (cy *cy_GB) Minus() string {
	return cy.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'cy_GB' and handles both Whole and Real numbers based on 'v'
func (cy *cy_GB) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, cy.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, cy.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, cy.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'cy_GB' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (cy *cy_GB) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, cy.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, cy.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, cy.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'cy_GB'
func (cy *cy_GB) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := cy.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, cy.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, cy.group[0])
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
		b = append(b, cy.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, cy.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'cy_GB'
// in accounting notation.
func (cy *cy_GB) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := cy.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, cy.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, cy.group[0])
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

		b = append(b, cy.currencyNegativePrefix[0])

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
			b = append(b, cy.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, cy.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'cy_GB'
func (cy *cy_GB) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'cy_GB'
func (cy *cy_GB) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, cy.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'cy_GB'
func (cy *cy_GB) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, cy.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'cy_GB'
func (cy *cy_GB) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, cy.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, cy.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'cy_GB'
func (cy *cy_GB) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, cy.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'cy_GB'
func (cy *cy_GB) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, cy.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, cy.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'cy_GB'
func (cy *cy_GB) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, cy.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, cy.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'cy_GB'
func (cy *cy_GB) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, cy.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, cy.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := cy.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
