package ro_MD

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ro_MD struct {
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

// New returns a new instance of translator for the 'ro_MD' locale
func New() locales.Translator {
	return &ro_MD{
		locale:                 "ro_MD",
		pluralsCardinal:        []locales.PluralRule{2, 4, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 6},
		pluralsRange:           []locales.PluralRule{4, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "L", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "ian.", "feb.", "mar.", "apr.", "mai", "iun.", "iul.", "aug.", "sept.", "oct.", "nov.", "dec."},
		monthsNarrow:           []string{"", "I", "F", "M", "A", "M", "I", "I", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "ianuarie", "februarie", "martie", "aprilie", "mai", "iunie", "iulie", "august", "septembrie", "octombrie", "noiembrie", "decembrie"},
		daysAbbreviated:        []string{"Dum", "Lun", "Mar", "Mie", "Joi", "Vin", "Sâm"},
		daysNarrow:             []string{"D", "L", "Ma", "Mi", "J", "V", "S"},
		daysShort:              []string{"Du", "Lu", "Ma", "Mi", "Jo", "Vi", "Sâ"},
		daysWide:               []string{"duminică", "luni", "marți", "miercuri", "joi", "vineri", "sâmbătă"},
		timezones:              map[string]string{"ACDT": "Ora de vară a Australiei Centrale", "ACST": "Ora de vară Acre", "ACT": "Ora standard Acre", "ACWDT": "Ora de vară a Australiei Central Occidentale", "ACWST": "Ora standard a Australiei Central Occidentale", "ADT": "Ora de vară în zona Atlantic nord-americană", "ADT Arabia": "Ora de vară arabă", "AEDT": "Ora de vară a Australiei Orientale", "AEST": "Ora standard a Australiei Orientale", "AFT": "Ora Afganistanului", "AKDT": "Ora de vară din Alaska", "AKST": "Ora standard din Alaska", "AMST": "Ora de vară a Amazonului", "AMST Armenia": "Ora de vară a Armeniei", "AMT": "Ora standard a Amazonului", "AMT Armenia": "Ora standard a Armeniei", "ANAST": "Ora de vară din Anadyr", "ANAT": "Ora standard din Anadyr", "ARST": "Ora de vară a Argentinei", "ART": "Ora standard a Argentinei", "AST": "Ora standard în zona Atlantic nord-americană", "AST Arabia": "Ora standard arabă", "AWDT": "Ora de vară a Australiei Occidentale", "AWST": "Ora standard a Australiei Occidentale", "AZST": "Ora de vară a Azerbaidjanului", "AZT": "Ora standard a Azerbaidjanului", "BDT Bangladesh": "Ora de vară din Bangladesh", "BNT": "Ora din Brunei Darussalam", "BOT": "Ora Boliviei", "BRST": "Ora de vară a Brasiliei", "BRT": "Ora standard a Brasiliei", "BST Bangladesh": "Ora standard din Bangladesh", "BT": "Ora Bhutanului", "CAST": "CAST", "CAT": "Ora Africii Centrale", "CCT": "Ora Insulelor Cocos", "CDT": "Ora de vară centrală nord-americană", "CHADT": "Ora de vară din Chatham", "CHAST": "Ora standard din Chatham", "CHUT": "Ora din Chuuk", "CKT": "Ora standard a Insulelor Cook", "CKT DST": "Ora de vară a Insulelor Cook", "CLST": "Ora de vară din Chile", "CLT": "Ora standard din Chile", "COST": "Ora de vară a Columbiei", "COT": "Ora standard a Columbiei", "CST": "Ora standard centrală nord-americană", "CST China": "Ora standard a Chinei", "CST China DST": "Ora de vară a Chinei", "CVST": "Ora de vară din Capul Verde", "CVT": "Ora standard din Capul Verde", "CXT": "Ora din Insula Christmas", "ChST": "Ora din Chamorro", "ChST NMI": "ChST NMI", "CuDT": "Ora de vară a Cubei", "CuST": "Ora standard a Cubei", "DAVT": "Ora din Davis", "DDUT": "Ora din Dumont-d’Urville", "EASST": "Ora de vară din Insula Paștelui", "EAST": "Ora standard din Insula Paștelui", "EAT": "Ora Africii Orientale", "ECT": "Ora Ecuadorului", "EDT": "Ora de vară orientală nord-americană", "EGDT": "Ora de vară a Groenlandei orientale", "EGST": "Ora standard a Groenlandei orientale", "EST": "Ora standard orientală nord-americană", "FEET": "Ora Europei de Est îndepărtate", "FJT": "Ora standard din Fiji", "FJT Summer": "Ora de vară din Fiji", "FKST": "Ora de vară din Insulele Falkland", "FKT": "Ora standard din Insulele Falkland", "FNST": "Ora de vară din Fernando de Noronha", "FNT": "Ora standard din Fernando de Noronha", "GALT": "Ora din Galapagos", "GAMT": "Ora din Gambier", "GEST": "Ora de vară a Georgiei", "GET": "Ora standard a Georgiei", "GFT": "Ora din Guyana Franceză", "GIT": "Ora Insulelor Gilbert", "GMT": "Ora de Greenwich", "GNSST": "GNSST", "GNST": "GNST", "GST": "Ora standard a Golfului", "GST Guam": "GST Guam", "GYT": "Ora din Guyana", "HADT": "Ora standard din Hawaii-Aleutine", "HAST": "Ora standard din Hawaii-Aleutine", "HKST": "Ora de vară din Hong Kong", "HKT": "Ora standard din Hong Kong", "HOVST": "Ora de vară din Hovd", "HOVT": "Ora standard din Hovd", "ICT": "Ora Indochinei", "IDT": "Ora de vară a Israelului", "IOT": "Ora Oceanului Indian", "IRKST": "Ora de vară din Irkuțk", "IRKT": "Ora standard din Irkuțk", "IRST": "Ora standard a Iranului", "IRST DST": "Ora de vară a Iranului", "IST": "Ora Indiei", "IST Israel": "Ora standard a Israelului", "JDT": "Ora de vară a Japoniei", "JST": "Ora standard a Japoniei", "KOST": "Ora din Kosrae", "KRAST": "Ora de vară din Krasnoiarsk", "KRAT": "Ora standard din Krasnoiarsk", "KST": "Ora standard a Coreei", "KST DST": "Ora de vară a Coreei", "LHDT": "Ora de vară din Lord Howe", "LHST": "Ora standard din Lord Howe", "LINT": "Ora din Insulele Line", "MAGST": "Ora de vară din Magadan", "MAGT": "Ora standard din Magadan", "MART": "Ora Insulelor Marchize", "MAWT": "Ora din Mawson", "MDT": "Ora de vară în zona montană nord-americană", "MESZ": "Ora de vară a Europei Centrale", "MEZ": "Ora standard a Europei Centrale", "MHT": "Ora Insulelor Marshall", "MMT": "Ora Myanmarului", "MSD": "Ora de vară a Moscovei", "MST": "Ora standard în zona montană nord-americană", "MUST": "Ora de vară din Mauritius", "MUT": "Ora standard din Mauritius", "MVT": "Ora din Maldive", "MYT": "Ora din Malaysia", "NCT": "Ora standard a Noii Caledonii", "NDT": "Ora de vară din Newfoundland", "NDT New Caledonia": "Ora de vară a Noii Caledonii", "NFDT": "Ora de vară a Insulei Norfolk", "NFT": "Ora standard a Insulei Norfolk", "NOVST": "Ora de vară din Novosibirsk", "NOVT": "Ora standard din Novosibirsk", "NPT": "Ora Nepalului", "NRT": "Ora din Nauru", "NST": "Ora standard din Newfoundland", "NUT": "Ora din Niue", "NZDT": "Ora de vară a Noii Zeelande", "NZST": "Ora standard a Noii Zeelande", "OESZ": "Ora de vară a Europei de Est", "OEZ": "Ora standard a Europei de Est", "OMSST": "Ora de vară din Omsk", "OMST": "Ora standard din Omsk", "PDT": "Ora de vară în zona Pacific nord-americană", "PDTM": "Ora de vară a zonei Pacific mexicane", "PETDT": "Ora de vară din Petropavlovsk-Kamciațki", "PETST": "Ora standard din Petropavlovsk-Kamciațki", "PGT": "Ora din Papua Noua Guinee", "PHOT": "Ora Insulelor Phoenix", "PKT": "Ora standard a Pakistanului", "PKT DST": "Ora de vară a Pakistanului", "PMDT": "Ora de vară din Saint-Pierre și Miquelon", "PMST": "Ora standard din Saint-Pierre și Miquelon", "PONT": "Ora din Ponape", "PST": "Ora standard în zona Pacific nord-americană", "PST Philippine": "Ora standard din Filipine", "PST Philippine DST": "Ora de vară din Filipine", "PST Pitcairn": "Ora din Pitcairn", "PSTM": "Ora standard a zonei Pacific mexicane", "PWT": "Ora din Palau", "PYST": "Ora de vară din Paraguay", "PYT": "Ora standard din Paraguay", "PYT Korea": "Ora din Pyongyang", "RET": "Ora din Reunion", "ROTT": "Ora din Rothera", "SAKST": "Ora de vară din Sahalin", "SAKT": "Ora standard din Sahalin", "SAMST": "Ora de vară din Samara", "SAMT": "Ora standard din Samara", "SAST": "Ora Africii Meridionale", "SBT": "Ora Insulelor Solomon", "SCT": "Ora din Seychelles", "SGT": "Ora din Singapore", "SLST": "SLST", "SRT": "Ora Surinamului", "SST Samoa": "Ora standard din Samoa", "SST Samoa Apia": "Ora standard din Samoa", "SST Samoa Apia DST": "Ora de vară din Samoa", "SST Samoa DST": "Ora de vară din Samoa", "SYOT": "Ora din Syowa", "TAAF": "Ora din Teritoriile Australe și Antarctice Franceze", "TAHT": "Ora din Tahiti", "TJT": "Ora din Tadjikistan", "TKT": "Ora din Tokelau", "TLT": "Ora Timorului de Est", "TMST": "Ora de vară din Turkmenistan", "TMT": "Ora standard din Turkmenistan", "TOST": "Ora de vară din Tonga", "TOT": "Ora standard din Tonga", "TVT": "Ora din Tuvalu", "TWT": "Ora standard din Taipei", "TWT DST": "Ora de vară din Taipei", "ULAST": "Ora de vară din Ulan Bator", "ULAT": "Ora standard din Ulan Bator", "UYST": "Ora de vară a Uruguayului", "UYT": "Ora standard a Uruguayului", "UZT": "Ora standard din Uzbekistan", "UZT DST": "Ora de vară din Uzbekistan", "VET": "Ora Venezuelei", "VLAST": "Ora de vară din Vladivostok", "VLAT": "Ora standard din Vladivostok", "VOLST": "Ora de vară din Volgograd", "VOLT": "Ora standard din Volgograd", "VOST": "Ora din Vostok", "VUT": "Ora standard din Vanuatu", "VUT DST": "Ora de vară din Vanuatu", "WAKT": "Ora Insulei Wake", "WARST": "Ora de vară a Argentinei Occidentale", "WART": "Ora standard a Argentinei Occidentale", "WAST": "Ora Africii Occidentale", "WAT": "Ora Africii Occidentale", "WESZ": "Ora de vară a Europei de Vest", "WEZ": "Ora standard a Europei de Vest", "WFT": "Ora din Wallis și Futuna", "WGST": "Ora de vară a Groenlandei occidentale", "WGT": "Ora standard a Groenlandei occidentale", "WIB": "Ora Indoneziei de Vest", "WIT": "Ora Indoneziei de Est", "WITA": "Ora Indoneziei Centrale", "YAKST": "Ora de vară din Iakuțk", "YAKT": "Ora standard din Iakuțk", "YEKST": "Ora de vară din Ekaterinburg", "YEKT": "Ora standard din Ekaterinburg", "YST": "Ora din Yukon", "МСК": "Ora standard a Moscovei", "اقتاۋ": "Ora standard Aqtau", "اقتاۋ قالاسى": "Ora de vară a zonei Aqtau", "اقتوبە": "Ora standard Aqtobe", "اقتوبە قالاسى": "Ora de vară a zonei Aqtobe", "الماتى": "Ora standard Almaty", "الماتى قالاسى": "Ora de vară Almaty", "باتىس قازاق ەلى": "Ora din Kazahstanul de Vest", "شىعىش قازاق ەلى": "Ora din Kazahstanul de Est", "قازاق ەلى": "Ora din Kazahstan", "قىرعىزستان": "Ora din Kârgâzstan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Ora de vară din Peru"},
	}
}

// Locale returns the current translators string locale
func (ro *ro_MD) Locale() string {
	return ro.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ro_MD'
func (ro *ro_MD) PluralsCardinal() []locales.PluralRule {
	return ro.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ro_MD'
func (ro *ro_MD) PluralsOrdinal() []locales.PluralRule {
	return ro.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ro_MD'
func (ro *ro_MD) PluralsRange() []locales.PluralRule {
	return ro.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ro_MD'
func (ro *ro_MD) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	i := int64(n)
	nMod100 := math.Mod(n, 100)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	} else if (v != 0) || (n == 0) || (n != 1 && nMod100 >= 1 && nMod100 <= 19) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ro_MD'
func (ro *ro_MD) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ro_MD'
func (ro *ro_MD) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	start := ro.CardinalPluralRule(num1, v1)
	end := ro.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOne {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	} else if start == locales.PluralRuleFew && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleFew {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ro *ro_MD) MonthAbbreviated(month time.Month) string {
	return ro.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ro *ro_MD) MonthsAbbreviated() []string {
	return ro.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ro *ro_MD) MonthNarrow(month time.Month) string {
	return ro.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ro *ro_MD) MonthsNarrow() []string {
	return ro.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ro *ro_MD) MonthWide(month time.Month) string {
	return ro.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ro *ro_MD) MonthsWide() []string {
	return ro.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ro *ro_MD) WeekdayAbbreviated(weekday time.Weekday) string {
	return ro.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ro *ro_MD) WeekdaysAbbreviated() []string {
	return ro.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ro *ro_MD) WeekdayNarrow(weekday time.Weekday) string {
	return ro.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ro *ro_MD) WeekdaysNarrow() []string {
	return ro.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ro *ro_MD) WeekdayShort(weekday time.Weekday) string {
	return ro.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ro *ro_MD) WeekdaysShort() []string {
	return ro.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ro *ro_MD) WeekdayWide(weekday time.Weekday) string {
	return ro.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ro *ro_MD) WeekdaysWide() []string {
	return ro.daysWide
}

// Decimal returns the decimal point of number
func (ro *ro_MD) Decimal() string {
	return ro.decimal
}

// Group returns the group of number
func (ro *ro_MD) Group() string {
	return ro.group
}

// Group returns the minus sign of number
func (ro *ro_MD) Minus() string {
	return ro.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ro_MD' and handles both Whole and Real numbers based on 'v'
func (ro *ro_MD) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ro.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ro.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ro.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ro_MD' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ro *ro_MD) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ro.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ro.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ro.percentSuffix...)

	b = append(b, ro.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ro_MD'
func (ro *ro_MD) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ro.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ro.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ro.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ro.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ro.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ro.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ro_MD'
// in accounting notation.
func (ro *ro_MD) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ro.currencies[currency]
	l := len(s) + len(symbol) + 6 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ro.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ro.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ro.currencyNegativePrefix[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ro.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ro.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ro.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ro_MD'
func (ro *ro_MD) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'ro_MD'
func (ro *ro_MD) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ro.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ro_MD'
func (ro *ro_MD) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ro.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ro_MD'
func (ro *ro_MD) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, ro.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ro.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ro_MD'
func (ro *ro_MD) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ro.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ro_MD'
func (ro *ro_MD) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ro.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ro.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ro_MD'
func (ro *ro_MD) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ro.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ro.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ro_MD'
func (ro *ro_MD) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ro.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ro.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ro.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
