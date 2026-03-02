package dsb_DE

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type dsb_DE struct {
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

// New returns a new instance of translator for the 'dsb_DE' locale
func New() locales.Translator {
	return &dsb_DE{
		locale:                 "dsb_DE",
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
		monthsAbbreviated:      []string{"", "jan.", "feb.", "měr.", "apr.", "maj.", "jun.", "jul.", "awg.", "sep.", "okt.", "now.", "dec."},
		monthsNarrow:           []string{"", "j", "f", "m", "a", "m", "j", "j", "a", "s", "o", "n", "d"},
		monthsWide:             []string{"", "januara", "februara", "měrca", "apryla", "maja", "junija", "julija", "awgusta", "septembra", "oktobra", "nowembra", "decembra"},
		daysAbbreviated:        []string{"nje", "pón", "wał", "srj", "stw", "pět", "sob"},
		daysNarrow:             []string{"n", "p", "w", "s", "s", "p", "s"},
		daysShort:              []string{"nj", "pó", "wa", "sr", "st", "pě", "so"},
		daysWide:               []string{"njeźela", "pónjeźele", "wałtora", "srjoda", "stwórtk", "pětk", "sobota"},
		timezones:              map[string]string{"ACDT": "Srjejźoawstralski lěśojski cas", "ACST": "Srjejźoawstralski standardny cas", "ACT": "ACT", "ACWDT": "Srjejźopódwjacorny awstralski lěśojski cas", "ACWST": "Srjejźopódwjacorny awstralski standardny cas", "ADT": "Atlantiski lěśojski cas", "ADT Arabia": "Arabiski lěśojski cas", "AEDT": "Pódzajtšnoawstralski lěśojski cas", "AEST": "Pódzajtšnoawstralski standardny cas", "AFT": "Afghaniski cas", "AKDT": "Alaskojski lěśojski cas", "AKST": "Alaskojski standardny cas", "AMST": "Amaconaski lěśojski cas", "AMST Armenia": "Armeński lěśojski cas", "AMT": "Amaconaski standardny cas", "AMT Armenia": "Armeński standardny cas", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Argentinski lěśojski cas", "ART": "Argentinski standardny cas", "AST": "Atlantiski standardny cas", "AST Arabia": "Arabiski standardny cas", "AWDT": "Pódwjacornoawstralski lěśojski cas", "AWST": "Pódwjacornoawstralski standardny cas", "AZST": "Azerbajdžaniski lěśojski cas", "AZT": "Azerbajdžaniski standardny cas", "BDT Bangladesh": "Bangladešski lěśojski cas", "BNT": "Bruneiski cas", "BOT": "Boliwiski cas", "BRST": "Brasília lěśojski cas", "BRT": "Brasília standardny cas", "BST Bangladesh": "Bangladešski standardny cas", "BT": "Bhutański cas", "CAST": "CAST", "CAT": "Srjejźoafriski cas", "CCT": "cas Kokosowych kupow", "CDT": "Pódpołnocnoameriski centralny lěśojski cas", "CHADT": "Chathamski lěśojski cas", "CHAST": "Chathamski standardny cas", "CHUT": "Chuukski cas", "CKT": "Standardny cas Cookowych kupow", "CKT DST": "lěśojski cas Cookowych kupow", "CLST": "Chilski lěśojski cas", "CLT": "Chilski standardny cas", "COST": "Kolumbiski lěśojski cas", "COT": "Kolumbiski standardny cas", "CST": "Pódpołnocnoameriski centralny standardny cas", "CST China": "Chinski standardny cas", "CST China DST": "Chinski lěśojski cas", "CVST": "Kapverdski lěśojski cas", "CVT": "Kapverdski standardny cas", "CXT": "cas Gódownych kupow", "ChST": "Chamorrski cas", "ChST NMI": "ChST NMI", "CuDT": "Kubański lěśojski cas", "CuST": "Kubański standardny cas", "DAVT": "Davis cas", "DDUT": "DumontDUrville cas", "EASST": "lěśojski cas Jatšowneje kupy", "EAST": "standardny cas Jatšowneje kupy", "EAT": "Pódzajtšnoafriski cas", "ECT": "Ekuadorski cas", "EDT": "Pódpołnocnoameriski pódzajtšny lěśojski cas", "EGDT": "Pódzajtšnogrönlandski lěśojski cas", "EGST": "Pódzajtšnogrönlandski standardny cas", "EST": "Pódpołnocnoameriski pódzajtšny standardny cas", "FEET": "Kaliningradski cas", "FJT": "Fidźiski standardny cas", "FJT Summer": "Fidźiski lěśojski cas", "FKST": "Falklandski lěśojski cas", "FKT": "Falklandski standardny cas", "FNST": "lěśojski cas Fernando de Noronha", "FNT": "standardny cas Fernando de Noronha", "GALT": "Galapagoski cas", "GAMT": "Gambierski cas", "GEST": "Georgiski lěśojski cas", "GET": "Georgiski standardny cas", "GFT": "Francojskoguyański cas", "GIT": "cas Gilbertowych kupow", "GMT": "Greenwichski cas", "GNSST": "GNSST", "GNST": "GNST", "GST": "cas Persiskego golfa", "GST Guam": "GST Guam", "GYT": "Guyański cas", "HADT": "Hawaiisko-aleutski standardny cas", "HAST": "Hawaiisko-aleutski standardny cas", "HKST": "Hongkongski lěśojski cas", "HKT": "Hongkongski standardny cas", "HOVST": "Chowdski lěśojski cas", "HOVT": "Chowdski standardny cas", "ICT": "Indochinski cas", "IDT": "Israelski lěśojski cas", "IOT": "Indiskooceaniski cas", "IRKST": "Irkutski lěśojski cas", "IRKT": "Irkutski standardny cas", "IRST": "Irański standardny cas", "IRST DST": "Irański lěśojski cas", "IST": "Indiski cas", "IST Israel": "Israelski standardny cas", "JDT": "Japański lěśojski cas", "JST": "Japański standardny cas", "KOST": "Kosraeski cas", "KRAST": "Krasnojarski lěśojski cas", "KRAT": "Krasnojarski standardny cas", "KST": "Korejski standardny cas", "KST DST": "Korejski lěśojski cas", "LHDT": "lěśojski cas kupy Lord-Howe", "LHST": "Standardny cas kupy Lord-Howe", "LINT": "cas Linijowych kupow", "MAGST": "Magadański lěśojski cas", "MAGT": "Magadański standardny cas", "MART": "Marqueski cas", "MAWT": "Mawson cas", "MDT": "MDT", "MESZ": "Srjejźnoeuropejski lěśojski cas", "MEZ": "Srjejźnoeuropejski standardny cas", "MHT": "cas Marshallowych kupow", "MMT": "Myanmarski cas", "MSD": "Moskowski lěśojski cas", "MST": "MST", "MUST": "Mauriciski lěśojski cas", "MUT": "Mauriciski standardny cas", "MVT": "Malediwski cas", "MYT": "Malajziski cas", "NCT": "Nowokaledoniski standardny cas", "NDT": "Nowofundlandski lěśojski cas", "NDT New Caledonia": "Nowokaledoniski lěśojski cas", "NFDT": "lěśojski cas kupy Norfolk", "NFT": "standardny cas kupy Norfolk", "NOVST": "Nowosibirski lěśojski cas", "NOVT": "Nowosibirski standardny cas", "NPT": "Nepalski cas", "NRT": "Nauruski cas", "NST": "Nowofundlandski standardny cas", "NUT": "Niueski cas", "NZDT": "Nowoseelandski lěśojski cas", "NZST": "Nowoseelandski standardny cas", "OESZ": "Pódzajtšnoeuropejski lěśojski cas", "OEZ": "Pódzajtšnoeuropejski standardny cas", "OMSST": "Omski lěśojski cas", "OMST": "Omski standardny cas", "PDT": "Pódpołnocnoameriski pacifiski lěśojski cas", "PDTM": "Mexiski pacifiski lěśojski cas", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Papua-Nowoginejski cas", "PHOT": "cas Phoenixowych kupow", "PKT": "Pakistański standardny cas", "PKT DST": "Pakistański lěśojski cas", "PMDT": "St.-Pierre-a-Miqueloński lěśojski cas", "PMST": "St.-Pierre-a-Miqueloński standardny cas", "PONT": "Ponapski cas", "PST": "Pódpołnocnoameriski pacifiski standardny cas", "PST Philippine": "Filipinski standardny cas", "PST Philippine DST": "Filipinski lěśojski cas", "PST Pitcairn": "cas Pitcairnowych kupow", "PSTM": "Mexiski pacifiski standardny cas", "PWT": "Palauski cas", "PYST": "Paraguayski lěśojski cas", "PYT": "Paraguayski standardny cas", "PYT Korea": "Pjöngjangski cas", "RET": "Reunionski cas", "ROTT": "cas Rothera", "SAKST": "Sachalinski lěśojski cas", "SAKT": "Sachalinski standardny cas", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Pódpołdnjowoafriski cas", "SBT": "Salomoński cas", "SCT": "Seychelski cas", "SGT": "Singapurski cas", "SLST": "SLST", "SRT": "Surinamski cas", "SST Samoa": "Samoaski standardny cas", "SST Samoa Apia": "Apiaski standardny cas", "SST Samoa Apia DST": "Apiaski lěśojski cas", "SST Samoa DST": "Samoaski lěśojski cas", "SYOT": "Syowa cas", "TAAF": "cas francojskego pódpołdnjowego a antarktiskeho teritoriuma", "TAHT": "Tahitiski cas", "TJT": "Tadźikiski cas", "TKT": "Tokelauski cas", "TLT": "Pódzajtšnotimorski cas", "TMST": "Turkmeniski lěśojski cas", "TMT": "Turkmeniski standardny cas", "TOST": "Tongaski lěśojski cas", "TOT": "Tongaski standardny cas", "TVT": "Tuvalski cas", "TWT": "Tchajpejski standardny cas", "TWT DST": "Tchajpejski lěśojski cas", "ULAST": "Ulan-Batorski lěśojski cas", "ULAT": "Ulan-Batorski standardny cas", "UYST": "Uruguayski lěśojski cas", "UYT": "Uruguayski standardny cas", "UZT": "Uzbekiski standardny cas", "UZT DST": "Uzbekiski lěśojski cas", "VET": "Venezuelski cas", "VLAST": "Wladiwostokski lěśojski cas", "VLAT": "Wladiwostokski standardny cas", "VOLST": "Wolgogradski lěśojski cas", "VOLT": "Wolgogradski standardny cas", "VOST": "cas Wostok", "VUT": "Vanuatski standardny cas", "VUT DST": "Vanuatski lěśojski cas", "WAKT": "cas kupy Wake", "WARST": "Pódwjacornoargentinski lěśojski cas", "WART": "Pódwjacornoargentinski standardny cas", "WAST": "Pódwjacornoafriski cas", "WAT": "Pódwjacornoafriski cas", "WESZ": "Pódwjacornoeuropejski lěśojski cas", "WEZ": "Pódwjacornoeuropejski standardny cas", "WFT": "cas kupow Wallis a Futuna", "WGST": "Pódwjacornogrönlandski lěśojski cas", "WGT": "Pódwjacornogrönlandski standardny cas", "WIB": "Pódwjacornoindoneski cas", "WIT": "Pódzajtšnoindoneski", "WITA": "Srjejźoindoneski cas", "YAKST": "Jakutski lěśojski cas", "YAKT": "Jakutski standardny cas", "YEKST": "Jekaterinburgski lěśojski cas", "YEKT": "Jekaterinburgski standardny cas", "YST": "Yukonowy cas", "МСК": "Moskowski standardny cas", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Pódwjacornokazachski cas", "شىعىش قازاق ەلى": "Pódzajtšnokazachski cas", "قازاق ەلى": "kazachiski cas", "قىرعىزستان": "Kirgiski cas", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Peruski lěśojski cas"},
	}
}

// Locale returns the current translators string locale
func (dsb *dsb_DE) Locale() string {
	return dsb.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'dsb_DE'
func (dsb *dsb_DE) PluralsCardinal() []locales.PluralRule {
	return dsb.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'dsb_DE'
func (dsb *dsb_DE) PluralsOrdinal() []locales.PluralRule {
	return dsb.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'dsb_DE'
func (dsb *dsb_DE) PluralsRange() []locales.PluralRule {
	return dsb.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'dsb_DE'
func (dsb *dsb_DE) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
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

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'dsb_DE'
func (dsb *dsb_DE) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'dsb_DE'
func (dsb *dsb_DE) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (dsb *dsb_DE) MonthAbbreviated(month time.Month) string {
	return dsb.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (dsb *dsb_DE) MonthsAbbreviated() []string {
	return dsb.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (dsb *dsb_DE) MonthNarrow(month time.Month) string {
	return dsb.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (dsb *dsb_DE) MonthsNarrow() []string {
	return dsb.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (dsb *dsb_DE) MonthWide(month time.Month) string {
	return dsb.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (dsb *dsb_DE) MonthsWide() []string {
	return dsb.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (dsb *dsb_DE) WeekdayAbbreviated(weekday time.Weekday) string {
	return dsb.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (dsb *dsb_DE) WeekdaysAbbreviated() []string {
	return dsb.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (dsb *dsb_DE) WeekdayNarrow(weekday time.Weekday) string {
	return dsb.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (dsb *dsb_DE) WeekdaysNarrow() []string {
	return dsb.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (dsb *dsb_DE) WeekdayShort(weekday time.Weekday) string {
	return dsb.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (dsb *dsb_DE) WeekdaysShort() []string {
	return dsb.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (dsb *dsb_DE) WeekdayWide(weekday time.Weekday) string {
	return dsb.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (dsb *dsb_DE) WeekdaysWide() []string {
	return dsb.daysWide
}

// Decimal returns the decimal point of number
func (dsb *dsb_DE) Decimal() string {
	return dsb.decimal
}

// Group returns the group of number
func (dsb *dsb_DE) Group() string {
	return dsb.group
}

// Group returns the minus sign of number
func (dsb *dsb_DE) Minus() string {
	return dsb.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'dsb_DE' and handles both Whole and Real numbers based on 'v'
func (dsb *dsb_DE) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, dsb.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, dsb.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, dsb.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'dsb_DE' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (dsb *dsb_DE) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, dsb.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, dsb.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, dsb.percentSuffix...)

	b = append(b, dsb.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'dsb_DE'
func (dsb *dsb_DE) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := dsb.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, dsb.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, dsb.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, dsb.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, dsb.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, dsb.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'dsb_DE'
// in accounting notation.
func (dsb *dsb_DE) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := dsb.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, dsb.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, dsb.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, dsb.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, dsb.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, dsb.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, dsb.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'dsb_DE'
func (dsb *dsb_DE) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'dsb_DE'
func (dsb *dsb_DE) FmtDateMedium(t time.Time) string {
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

// FmtDateLong returns the long date representation of 't' for 'dsb_DE'
func (dsb *dsb_DE) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, dsb.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'dsb_DE'
func (dsb *dsb_DE) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, dsb.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2e, 0x20}...)
	b = append(b, dsb.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'dsb_DE'
func (dsb *dsb_DE) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, dsb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'dsb_DE'
func (dsb *dsb_DE) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, dsb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, dsb.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'dsb_DE'
func (dsb *dsb_DE) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, dsb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, dsb.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'dsb_DE'
func (dsb *dsb_DE) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, dsb.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, dsb.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := dsb.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
