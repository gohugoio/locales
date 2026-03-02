package sv_FI

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type sv_FI struct {
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

// New returns a new instance of translator for the 'sv_FI' locale
func New() locales.Translator {
	return &sv_FI{
		locale:                 "sv_FI",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ",",
		group:                  " ",
		minus:                  "−",
		percent:                "%",
		timeSeparator:          ".",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "jan.", "feb.", "mars", "apr.", "maj", "juni", "juli", "aug.", "sep.", "okt.", "nov.", "dec."},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "J", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "januari", "februari", "mars", "april", "maj", "juni", "juli", "augusti", "september", "oktober", "november", "december"},
		daysAbbreviated:        []string{"sön", "mån", "tis", "ons", "tors", "fre", "lör"},
		daysNarrow:             []string{"S", "M", "T", "O", "T", "F", "L"},
		daysShort:              []string{"sö", "må", "ti", "on", "to", "fr", "lö"},
		daysWide:               []string{"söndag", "måndag", "tisdag", "onsdag", "torsdag", "fredag", "lördag"},
		timezones:              map[string]string{"ACDT": "centralaustralisk sommartid", "ACST": "västbrasiliansk sommartid", "ACT": "västbrasiliansk normaltid", "ACWDT": "västcentralaustralisk sommartid", "ACWST": "västcentralaustralisk normaltid", "ADT": "nordamerikansk atlantsommartid", "ADT Arabia": "saudiarabisk sommartid", "AEDT": "östaustralisk sommartid", "AEST": "östaustralisk normaltid", "AFT": "afghansk tid", "AKDT": "Alaska, sommartid", "AKST": "Alaska, normaltid", "AMST": "Amazonas, sommartid", "AMST Armenia": "armenisk sommartid", "AMT": "Amazonas, normaltid", "AMT Armenia": "armenisk normaltid", "ANAST": "Anadyrsommartid", "ANAT": "Anadyrnormaltid", "ARST": "östargentinsk sommartid", "ART": "östargentinsk normaltid", "AST": "nordamerikansk atlantnormaltid", "AST Arabia": "saudiarabisk normaltid", "AWDT": "västaustralisk sommartid", "AWST": "västaustralisk normaltid", "AZST": "azerbajdzjansk sommartid", "AZT": "azerbajdzjansk normaltid", "BDT Bangladesh": "bangladeshisk sommartid", "BNT": "Bruneitid", "BOT": "boliviansk tid", "BRST": "Brasilia, sommartid", "BRT": "Brasilia, normaltid", "BST Bangladesh": "bangladeshisk normaltid", "BT": "bhutansk tid", "CAST": "Caseytid", "CAT": "centralafrikansk tid", "CCT": "Keelingöarnas tid", "CDT": "centralnordamerikansk sommartid", "CHADT": "Chatham, sommartid", "CHAST": "Chatham, normaltid", "CHUT": "Chuuktid", "CKT": "Cooköarnas normaltid", "CKT DST": "Cooköarnas sommartid", "CLST": "chilensk sommartid", "CLT": "chilensk normaltid", "COST": "colombiansk sommartid", "COT": "colombiansk normaltid", "CST": "centralnordamerikansk normaltid", "CST China": "kinesisk normaltid", "CST China DST": "kinesisk sommartid", "CVST": "Kap Verde, sommartid", "CVT": "Kap Verde, normaltid", "CXT": "Julöns tid", "ChST": "Chamorrotid", "ChST NMI": "Nordmarianernas tid", "CuDT": "kubansk sommartid", "CuST": "kubansk normaltid", "DAVT": "Davistid", "DDUT": "Dumont d’Urville-tid", "EASST": "Påskön, sommartid", "EAST": "Påskön, normaltid", "EAT": "östafrikansk tid", "ECT": "ecuadoriansk tid", "EDT": "östnordamerikansk sommartid", "EGDT": "östgrönländsk sommartid", "EGST": "östgrönländsk normaltid", "EST": "östnordamerikansk normaltid", "FEET": "Kaliningradtid", "FJT": "Fiji, normaltid", "FJT Summer": "Fiji, sommartid", "FKST": "Falklandsöarna, sommartid", "FKT": "Falklandsöarna, normaltid", "FNST": "Fernando de Noronha, sommartid", "FNT": "Fernando de Noronha, normaltid", "GALT": "Galápagostid", "GAMT": "Gambiertid", "GEST": "georgisk sommartid", "GET": "georgisk normaltid", "GFT": "Franska Guyanatid", "GIT": "Kiribatitid", "GMT": "Greenwichtid", "GNSST": "GNSST", "GNST": "GNST", "GST": "Persiska vikentid", "GST Guam": "Guamtid", "GYT": "Guyanatid", "HADT": "aleutisk sommartid", "HAST": "aleutisk normaltid", "HKST": "Hongkong, sommartid", "HKT": "Hongkong, normaltid", "HOVST": "Chovd, sommartid", "HOVT": "Chovd, normaltid", "ICT": "indokinesisk tid", "IDT": "israelisk sommartid", "IOT": "Brittiska Indiska oceanöarnas tid", "IRKST": "Irkutsk, sommartid", "IRKT": "Irkutsk, normaltid", "IRST": "iransk normaltid", "IRST DST": "iransk sommartid", "IST": "indisk tid", "IST Israel": "israelisk normaltid", "JDT": "japansk sommartid", "JST": "japansk normaltid", "KOST": "Kosraetid", "KRAST": "Krasnojarsk, sommartid", "KRAT": "Krasnojarsk, normaltid", "KST": "koreansk normaltid", "KST DST": "koreansk sommartid", "LHDT": "Lord Howe, sommartid", "LHST": "Lord Howe, normaltid", "LINT": "Lineöarnas tid", "MAGST": "Magadan, sommartid", "MAGT": "Magadan, normaltid", "MART": "Marquesastid", "MAWT": "Mawsontid", "MDT": "Macaosommartid", "MESZ": "centraleuropeisk sommartid", "MEZ": "centraleuropeisk normaltid", "MHT": "Marshallöarnas tid", "MMT": "burmesisk tid", "MSD": "Moskva, sommartid", "MST": "Macaonormaltid", "MUST": "Mauritius, sommartid", "MUT": "Mauritius, normaltid", "MVT": "Maldivernatid", "MYT": "malaysisk tid", "NCT": "Nya Kaledonien, normaltid", "NDT": "Newfoundland, sommartid", "NDT New Caledonia": "Nya Kaledonien, sommartid", "NFDT": "Norfolköns sommartid", "NFT": "Norfolköns normaltid", "NOVST": "Novosibirsk, sommartid", "NOVT": "Novosibirsk, normaltid", "NPT": "nepalesisk tid", "NRT": "Naurutid", "NST": "Newfoundland, normaltid", "NUT": "Niuetid", "NZDT": "nyzeeländsk sommartid", "NZST": "nyzeeländsk normaltid", "OESZ": "östeuropeisk sommartid", "OEZ": "östeuropeisk normaltid", "OMSST": "Omsk, sommartid", "OMST": "Omsk, normaltid", "PDT": "västnordamerikansk sommartid", "PDTM": "mexikansk stillahavstid, sommartid", "PETDT": "Kamtjatkasommartid", "PETST": "Kamtjatkanormaltid", "PGT": "Papua Nya Guineas tid", "PHOT": "Enderburytid", "PKT": "pakistansk normaltid", "PKT DST": "pakistansk sommartid", "PMDT": "Saint-Pierre-et-Miquelon, sommartid", "PMST": "Saint-Pierre-et-Miquelon, normaltid", "PONT": "Ponapetid", "PST": "västnordamerikansk normaltid", "PST Philippine": "filippinsk normaltid", "PST Philippine DST": "filippinsk sommartid", "PST Pitcairn": "Pitcairntid", "PSTM": "mexikansk stillahavstid, normaltid", "PWT": "Palautid", "PYST": "paraguayansk sommartid", "PYT": "paraguayansk normaltid", "PYT Korea": "Pyongyangtid", "RET": "Réuniontid", "ROTT": "Rotheratid", "SAKST": "Sachalin, sommartid", "SAKT": "Sachalin, normaltid", "SAMST": "Samarasommartid", "SAMT": "Samaranormaltid", "SAST": "sydafrikansk tid", "SBT": "Salomonöarnas tid", "SCT": "Seychellernatid", "SGT": "Singaporetid", "SLST": "Sri Lankatid", "SRT": "Surinamtid", "SST Samoa": "samoansk normaltid", "SST Samoa Apia": "Apia, normaltid", "SST Samoa Apia DST": "Apia, sommartid", "SST Samoa DST": "samoansk sommartid", "SYOT": "Syowatid", "TAAF": "Franska Sydterritoriernas tid", "TAHT": "Tahititid", "TJT": "Tadzjikistantid", "TKT": "Tokelautid", "TLT": "östtimorisk tid", "TMST": "turkmensk sommartid", "TMT": "turkmensk normaltid", "TOST": "Tonga, sommartid", "TOT": "Tonga, normaltid", "TVT": "Tuvalutid", "TWT": "Taipei, normaltid", "TWT DST": "Taipei, sommartid", "ULAST": "Ulaanbaatar, sommartid", "ULAT": "Ulaanbaatar, normaltid", "UYST": "uruguayansk sommartid", "UYT": "uruguayansk normaltid", "UZT": "uzbekisk normaltid", "UZT DST": "uzbekisk sommartid", "VET": "venezuelansk tid", "VLAST": "Vladivostok, sommartid", "VLAT": "Vladivostok, normaltid", "VOLST": "Volgograd, sommartid", "VOLT": "Volgograd, normaltid", "VOST": "Vostoktid", "VUT": "Vanuatu, normaltid", "VUT DST": "Vanuatu, sommartid", "WAKT": "Wakeöarnas tid", "WARST": "västargentinsk sommartid", "WART": "västargentinsk normaltid", "WAST": "västafrikansk tid", "WAT": "västafrikansk tid", "WESZ": "västeuropeisk sommartid", "WEZ": "västeuropeisk normaltid", "WFT": "Wallis- och Futunaöarnas tid", "WGST": "västgrönländsk sommartid", "WGT": "västgrönländsk normaltid", "WIB": "västindonesisk tid", "WIT": "östindonesisk tid", "WITA": "centralindonesisk tid", "YAKST": "Jakutsk, sommartid", "YAKT": "Jakutsk, normaltid", "YEKST": "Jekaterinburg, sommartid", "YEKT": "Jekaterinburg, normaltid", "YST": "Yukontid", "МСК": "Moskva, normaltid", "اقتاۋ": "Aqtaunormaltid", "اقتاۋ قالاسى": "Aqtausommartid", "اقتوبە": "Aqtöbenormaltid", "اقتوبە قالاسى": "Aqtöbesommartid", "الماتى": "Almatynormaltid", "الماتى قالاسى": "Almatysommartid", "باتىس قازاق ەلى": "västkazakstansk tid", "شىعىش قازاق ەلى": "östkazakstansk tid", "قازاق ەلى": "kazakstansk tid", "قىرعىزستان": "kirgizisk tid", "قىزىلوردا": "Qyzylordanormaltid", "قىزىلوردا قالاسى": "Qyzylordasommartid", "∅∅∅": "azorisk sommartid"},
	}
}

// Locale returns the current translators string locale
func (sv *sv_FI) Locale() string {
	return sv.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sv_FI'
func (sv *sv_FI) PluralsCardinal() []locales.PluralRule {
	return sv.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sv_FI'
func (sv *sv_FI) PluralsOrdinal() []locales.PluralRule {
	return sv.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'sv_FI'
func (sv *sv_FI) PluralsRange() []locales.PluralRule {
	return sv.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sv_FI'
func (sv *sv_FI) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sv_FI'
func (sv *sv_FI) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	nMod10 := math.Mod(n, 10)
	nMod100 := math.Mod(n, 100)

	if (nMod10 == 1 || nMod10 == 2) && (nMod100 != 11 && nMod100 != 12) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sv_FI'
func (sv *sv_FI) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (sv *sv_FI) MonthAbbreviated(month time.Month) string {
	if len(sv.monthsAbbreviated) == 0 {
		return ""
	}
	return sv.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (sv *sv_FI) MonthsAbbreviated() []string {
	return sv.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (sv *sv_FI) MonthNarrow(month time.Month) string {
	if len(sv.monthsNarrow) == 0 {
		return ""
	}
	return sv.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (sv *sv_FI) MonthsNarrow() []string {
	return sv.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (sv *sv_FI) MonthWide(month time.Month) string {
	if len(sv.monthsWide) == 0 {
		return ""
	}
	return sv.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (sv *sv_FI) MonthsWide() []string {
	return sv.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (sv *sv_FI) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(sv.daysAbbreviated) == 0 {
		return ""
	}
	return sv.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (sv *sv_FI) WeekdaysAbbreviated() []string {
	return sv.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (sv *sv_FI) WeekdayNarrow(weekday time.Weekday) string {
	if len(sv.daysNarrow) == 0 {
		return ""
	}
	return sv.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (sv *sv_FI) WeekdaysNarrow() []string {
	return sv.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (sv *sv_FI) WeekdayShort(weekday time.Weekday) string {
	if len(sv.daysShort) == 0 {
		return ""
	}
	return sv.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (sv *sv_FI) WeekdaysShort() []string {
	return sv.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (sv *sv_FI) WeekdayWide(weekday time.Weekday) string {
	if len(sv.daysWide) == 0 {
		return ""
	}
	return sv.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (sv *sv_FI) WeekdaysWide() []string {
	return sv.daysWide
}

// Decimal returns the decimal point of number
func (sv *sv_FI) Decimal() string {
	return sv.decimal
}

// Group returns the group of number
func (sv *sv_FI) Group() string {
	return sv.group
}

// Group returns the minus sign of number
func (sv *sv_FI) Minus() string {
	return sv.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sv_FI' and handles both Whole and Real numbers based on 'v'
func (sv *sv_FI) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 4 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(sv.group) - 1; j >= 0; j-- {
					b = append(b, sv.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(sv.minus) - 1; j >= 0; j-- {
			b = append(b, sv.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'sv_FI' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sv *sv_FI) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 7
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sv.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(sv.minus) - 1; j >= 0; j-- {
			b = append(b, sv.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, sv.percentSuffix...)

	b = append(b, sv.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sv_FI'
func (sv *sv_FI) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sv.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(sv.group) - 1; j >= 0; j-- {
					b = append(b, sv.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(sv.minus) - 1; j >= 0; j-- {
			b = append(b, sv.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sv.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, sv.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sv_FI'
// in accounting notation.
func (sv *sv_FI) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sv.currencies[currency]
	l := len(s) + len(symbol) + 6 + 2*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(sv.group) - 1; j >= 0; j-- {
					b = append(b, sv.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		for j := len(sv.minus) - 1; j >= 0; j-- {
			b = append(b, sv.minus[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sv.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, sv.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, sv.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'sv_FI'
func (sv *sv_FI) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'sv_FI'
func (sv *sv_FI) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sv.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'sv_FI'
func (sv *sv_FI) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sv.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'sv_FI'
func (sv *sv_FI) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, sv.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, sv.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'sv_FI'
func (sv *sv_FI) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'sv_FI'
func (sv *sv_FI) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'sv_FI'
func (sv *sv_FI) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'sv_FI'
func (sv *sv_FI) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := sv.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
