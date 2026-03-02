package hsb_DE

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type hsb_DE struct {
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

// New returns a new instance of translator for the 'hsb_DE' locale
func New() locales.Translator {
	return &hsb_DE{
		locale:                 "hsb_DE",
		pluralsCardinal:        []locales.PluralRule{2, 3, 4, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           nil,
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "jan.", "feb.", "měr.", "apr.", "mej.", "jun.", "jul.", "awg.", "sep.", "okt.", "now.", "dec."},
		monthsNarrow:           []string{"", "j", "f", "m", "a", "m", "j", "j", "a", "s", "o", "n", "d"},
		monthsWide:             []string{"", "januara", "februara", "měrca", "apryla", "meje", "junija", "julija", "awgusta", "septembra", "oktobra", "nowembra", "decembra"},
		daysAbbreviated:        []string{"nje", "pón", "wut", "srj", "štw", "pja", "sob"},
		daysNarrow:             []string{"n", "p", "w", "s", "š", "p", "s"},
		daysShort:              []string{"nj", "pó", "wu", "sr", "št", "pj", "so"},
		daysWide:               []string{"njedźela", "póndźela", "wutora", "srjeda", "štwórtk", "pjatk", "sobota"},
		timezones:              map[string]string{"ACDT": "srjedźoawstralski lětni čas", "ACST": "srjedźoawstralski standardny čas", "ACT": "ACT", "ACWDT": "sjedźozapadny awstralski lětni čas", "ACWST": "srjedźozapadny awstralski standardny čas", "ADT": "atlantiski lětni čas", "ADT Arabia": "arabski lětni čas", "AEDT": "wuchodoawstralski lětni čas", "AEST": "wuchodoawstralski standardny čas", "AFT": "afghanski čas", "AKDT": "alaskaski lětni čas", "AKST": "alaskaski standardny čas", "AMST": "Amaconaski lětni čas", "AMST Armenia": "armenski lětni čas", "AMT": "Amaconaski standardny čas", "AMT Armenia": "armenski standardny čas", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "argentinski lětni čas", "ART": "argentinski standardny čas", "AST": "atlantiski standardny čas", "AST Arabia": "arabski standardny čas", "AWDT": "zapadoawstralski lětni čas", "AWST": "zapadoawstralski standardny čas", "AZST": "azerbajdźanski lětni čas", "AZT": "azerbajdźanski standardny čas", "BDT Bangladesh": "bangladešski lětni čas", "BNT": "bruneiski čas", "BOT": "boliwiski čas", "BRST": "Brasiliski lětni čas", "BRT": "Brasiliski standardny čas", "BST Bangladesh": "bangladešski standardny čas", "BT": "bhutanski čas", "CAST": "CAST", "CAT": "centralnoafriski čas", "CCT": "čas Kokosowych kupow", "CDT": "sewjeroameriski centralny lětni čas", "CHADT": "chathamski lětni čas", "CHAST": "chathamski standardny čas", "CHUT": "chuukski čas", "CKT": "standardny čas Cookowych kupow", "CKT DST": "lětni čas Cookowych kupow", "CLST": "chilski lětni čas", "CLT": "chilski standardny čas", "COST": "kolumbiski lětni čas", "COT": "kolumbiski standardny čas", "CST": "sewjeroameriski centralny standardny čas", "CST China": "chinski standardny čas", "CST China DST": "chinski lětni čas", "CVST": "kapverdski lětni čas", "CVT": "kapverdski standardny čas", "CXT": "čas Hodowneje kupy", "ChST": "chamorroski čas", "ChST NMI": "ChST NMI", "CuDT": "kubaski lětni čas", "CuST": "kubaski standardny čas", "DAVT": "Daviski čas", "DDUT": "Dumont d´ Urvilleski čas", "EASST": "lětni čas Jutrowneje kupy", "EAST": "standardny čas Jutrowneje kupy", "EAT": "wuchodoafriski čas", "ECT": "ekwadorski čas", "EDT": "sewjeroameriski wuchodny lětni čas", "EGDT": "wuchodogrönlandski lětni čas", "EGST": "wuchodogrönlandski standardny čas", "EST": "sewjeroameriski wuchodny standardny čas", "FEET": "Kaliningradski čas", "FJT": "fidźiski standardny čas", "FJT Summer": "fidźiski lětni čas", "FKST": "falklandski lětni čas", "FKT": "falklandski standardny čas", "FNST": "lětni čas kupow Fernando de Noronha", "FNT": "standardny čas kupow Fernando de Noronha", "GALT": "galapagoski čas", "GAMT": "gambierski čas", "GEST": "georgiski lětni čas", "GET": "georgiski standardny čas", "GFT": "francoskoguyanski čas", "GIT": "čas Gilbertowych kupow", "GMT": "Greenwichski čas", "GNSST": "GNSST", "GNST": "GNST", "GST": "čas Persiskeho golfa", "GST Guam": "GST Guam", "GYT": "guyanski čas", "HADT": "hawaiisko-aleutski lětni čas", "HAST": "hawaiisko-aleutski standardny čas", "HKST": "Hongkongski lětni čas", "HKT": "Hongkongski standardny čas", "HOVST": "Chowdski lětni čas", "HOVT": "Chowdski standardny čas", "ICT": "indochinski čas", "IDT": "israelski lětni čas", "IOT": "indiskooceanski čas", "IRKST": "Irkutski lětni čas", "IRKT": "Irkutski standardny čas", "IRST": "iranski standardny čas", "IRST DST": "iranski lětni čas", "IST": "indiski čas", "IST Israel": "israelski standardny čas", "JDT": "japanski lětni čas", "JST": "japanski standardny čas", "KOST": "kosraeski čas", "KRAST": "Krasnojarski lětni čas", "KRAT": "Krasnojarski standardny čas", "KST": "korejski standardny čas", "KST DST": "korejski lětni čas", "LHDT": "lětni čas kupy Lord-Howe", "LHST": "standardny čas kupy Lord-Howe", "LINT": "čas Linijowych kupow", "MAGST": "Magadanski lětni čas", "MAGT": "Magadanski standardny čas", "MART": "marquesaski čas", "MAWT": "Mawsonski čas", "MDT": "MDT", "MESZ": "srjedźoeuropski lětni čas", "MEZ": "srjedźoeuropski standardny čas", "MHT": "čas Marshallowych kupow", "MMT": "myanmarski čas", "MSD": "Moskowski lětni čas", "MST": "MST", "MUST": "mauritiuski lětni čas", "MUT": "mauritiuski standardny čas", "MVT": "malediwski čas", "MYT": "malajziski čas", "NCT": "nowokaledonski standardny čas", "NDT": "nowofundlandski lětni čas", "NDT New Caledonia": "nowokaledonski lětni čas", "NFDT": "lětni čas kupy Norfolk", "NFT": "standardny čas kupy Norfolk", "NOVST": "Nowosibirski lětni čas", "NOVT": "Nowosibirski standardny čas", "NPT": "nepalski čas", "NRT": "nauruski čas", "NST": "nowofundlandski standardny čas", "NUT": "niueski čas", "NZDT": "nowoseelandski lětni čas", "NZST": "nowoseelandski standardny čas", "OESZ": "wuchodoeuropski lětni čas", "OEZ": "wuchodoeuropski standardny čas", "OMSST": "Omski lětni čas", "OMST": "Omski standardny čas", "PDT": "sewjeroameriski pacifiski lětni čas", "PDTM": "mexiski pacifiski lětni čas", "PETDT": "PETDT", "PETST": "PETST", "PGT": "papua-nowoginejski čas", "PHOT": "čas Phoenixowych kupow", "PKT": "pakistanski standardny čas", "PKT DST": "pakistanski lětni čas", "PMDT": "lětni čas kupow St. Pierre a Miquelon", "PMST": "standardny čas kupow St. Pierre a Miquelon", "PONT": "ponapeski čas", "PST": "sewjeroameriski pacifiski standardny čas", "PST Philippine": "filipinski standardny čas", "PST Philippine DST": "filipinski lětni čas", "PST Pitcairn": "čas Pitcairnowych kupow", "PSTM": "mexiski pacifiski standardny čas", "PWT": "palauski čas", "PYST": "Paraguayski lětni čas", "PYT": "Paraguayski standardny čas", "PYT Korea": "Pjöngjangski čas", "RET": "reunionski čas", "ROTT": "Rotheraski čas", "SAKST": "sachalinski lětni čas", "SAKT": "sachalinski standardny čas", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "južnoafriski čas", "SBT": "čas Salomonskich kupow", "SCT": "seychellski čas", "SGT": "Singapurski čas", "SLST": "SLST", "SRT": "surinamski čas", "SST Samoa": "samoaski standardny čas", "SST Samoa Apia": "Apiaski standardny čas", "SST Samoa Apia DST": "Apiaski lětni čas", "SST Samoa DST": "samoaski lětni čas", "SYOT": "Syowaski čas", "TAAF": "čas Francoskeho južneho a antarktiskeho teritorija", "TAHT": "tahitiski čas", "TJT": "tadźikski čas", "TKT": "tokelauski čas", "TLT": "wuchodnotimorski čas", "TMST": "turkmenski lětni čas", "TMT": "turkmenski standardny čas", "TOST": "tongaski lětni čas", "TOT": "tongaski standardny čas", "TVT": "tuvaluski čas", "TWT": "Taipehski standardny čas", "TWT DST": "Taipehski lětni čas", "ULAST": "Ulan-Batorski lětni čas", "ULAT": "Ulan-Batorski standardny čas", "UYST": "uruguayski lětni čas", "UYT": "uruguayski standardny čas", "UZT": "uzbekski standardny čas", "UZT DST": "uzbekski lětni čas", "VET": "venezuelski čas", "VLAST": "Wladiwostokski lětni čas", "VLAT": "Wladiwostokski standardny čas", "VOLST": "Wolgogradski lětni čas", "VOLT": "Wolgogradski standardny čas", "VOST": "Wostokski čas", "VUT": "vanuatuski standardny čas", "VUT DST": "vanuatuski lětni čas", "WAKT": "čas kupy Wake", "WARST": "zapadoargentinski lětni čas", "WART": "zapadoargentinski standardny čas", "WAST": "zapadoafriski čas", "WAT": "zapadoafriski čas", "WESZ": "zapadoeuropski lětni čas", "WEZ": "zapadoeuropski standardny čas", "WFT": "čas kupow Wallis a Futuna", "WGST": "zapadogrönlandski lětni čas", "WGT": "zapadogrönlandski standardny čas", "WIB": "zapadoindoneski čas", "WIT": "wuchodoindoneski", "WITA": "srjedźoindoneski čas", "YAKST": "Jakutski lětni čas", "YAKT": "Jakutski standardny čas", "YEKST": "Jekaterinburgski lětni čas", "YEKT": "Jekaterinburgski standardny čas", "YST": "Yukonowy čas", "МСК": "Moskowski standardny čas", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "zapadnokazachski čas", "شىعىش قازاق ەلى": "wuchodnokazachski čas", "قازاق ەلى": "kazachski čas", "قىرعىزستان": "kirgiski čas", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "acorski lětni čas"},
	}
}

// Locale returns the current translators string locale
func (hsb *hsb_DE) Locale() string {
	return hsb.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'hsb_DE'
func (hsb *hsb_DE) PluralsCardinal() []locales.PluralRule {
	return hsb.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'hsb_DE'
func (hsb *hsb_DE) PluralsOrdinal() []locales.PluralRule {
	return hsb.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'hsb_DE'
func (hsb *hsb_DE) PluralsRange() []locales.PluralRule {
	return hsb.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'hsb_DE'
func (hsb *hsb_DE) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)
	iMod100 := i % 100
	fMod100 := f % 100

	if (v == 0 && iMod100 == 1) || (fMod100 == 1) {
		return locales.PluralRuleOne
	} else if (v == 0 && iMod100 == 2) || (fMod100 == 2) {
		return locales.PluralRuleTwo
	} else if (v == 0 && iMod100 >= 3 && iMod100 <= 4) || (fMod100 >= 3 && fMod100 <= 4) {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'hsb_DE'
func (hsb *hsb_DE) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'hsb_DE'
func (hsb *hsb_DE) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (hsb *hsb_DE) MonthAbbreviated(month time.Month) string {
	if len(hsb.monthsAbbreviated) == 0 {
		return ""
	}
	return hsb.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (hsb *hsb_DE) MonthsAbbreviated() []string {
	return hsb.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (hsb *hsb_DE) MonthNarrow(month time.Month) string {
	if len(hsb.monthsNarrow) == 0 {
		return ""
	}
	return hsb.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (hsb *hsb_DE) MonthsNarrow() []string {
	return hsb.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (hsb *hsb_DE) MonthWide(month time.Month) string {
	if len(hsb.monthsWide) == 0 {
		return ""
	}
	return hsb.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (hsb *hsb_DE) MonthsWide() []string {
	return hsb.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (hsb *hsb_DE) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(hsb.daysAbbreviated) == 0 {
		return ""
	}
	return hsb.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (hsb *hsb_DE) WeekdaysAbbreviated() []string {
	return hsb.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (hsb *hsb_DE) WeekdayNarrow(weekday time.Weekday) string {
	if len(hsb.daysNarrow) == 0 {
		return ""
	}
	return hsb.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (hsb *hsb_DE) WeekdaysNarrow() []string {
	return hsb.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (hsb *hsb_DE) WeekdayShort(weekday time.Weekday) string {
	if len(hsb.daysShort) == 0 {
		return ""
	}
	return hsb.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (hsb *hsb_DE) WeekdaysShort() []string {
	return hsb.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (hsb *hsb_DE) WeekdayWide(weekday time.Weekday) string {
	if len(hsb.daysWide) == 0 {
		return ""
	}
	return hsb.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (hsb *hsb_DE) WeekdaysWide() []string {
	return hsb.daysWide
}

// Decimal returns the decimal point of number
func (hsb *hsb_DE) Decimal() string {
	return hsb.decimal
}

// Group returns the group of number
func (hsb *hsb_DE) Group() string {
	return hsb.group
}

// Group returns the minus sign of number
func (hsb *hsb_DE) Minus() string {
	return hsb.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'hsb_DE' and handles both Whole and Real numbers based on 'v'
func (hsb *hsb_DE) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hsb.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, hsb.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, hsb.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'hsb_DE' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (hsb *hsb_DE) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hsb.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, hsb.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, hsb.percentSuffix...)

	b = append(b, hsb.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'hsb_DE'
func (hsb *hsb_DE) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := hsb.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hsb.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, hsb.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, hsb.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, hsb.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, hsb.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'hsb_DE'
// in accounting notation.
func (hsb *hsb_DE) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := hsb.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, hsb.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, hsb.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, hsb.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, hsb.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, hsb.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, hsb.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'hsb_DE'
func (hsb *hsb_DE) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'hsb_DE'
func (hsb *hsb_DE) FmtDateMedium(t time.Time) string {
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

// FmtDateLong returns the long date representation of 't' for 'hsb_DE'
func (hsb *hsb_DE) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, hsb.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'hsb_DE'
func (hsb *hsb_DE) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, hsb.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, hsb.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'hsb_DE'
func (hsb *hsb_DE) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hsb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20, 0x68, 0x6f, 0x64, 0xc5, 0xba}...)
	b = append(b, []byte{0x2e}...)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'hsb_DE'
func (hsb *hsb_DE) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hsb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, hsb.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'hsb_DE'
func (hsb *hsb_DE) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hsb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, hsb.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'hsb_DE'
func (hsb *hsb_DE) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, hsb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, hsb.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := hsb.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
