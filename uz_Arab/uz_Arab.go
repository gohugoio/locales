package uz_Arab

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type uz_Arab struct {
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

// New returns a new instance of translator for the 'uz_Arab' locale
func New() locales.Translator {
	return &uz_Arab{
		locale:                 "uz_Arab",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "؋", "ALK", "ALL", "֏", "ANG", "Kz", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "$", "ATS", "A$", "AWG", "AZM", "₼", "BAD", "KM", "BAN", "$", "৳", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "$", "$", "Bs", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "R$", "BRN", "BRR", "BRZ", "$", "BTN", "BUK", "P", "BYB", "р.", "BYR", "$", "CA$", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "$", "CNH", "CNX", "CN¥", "$", "COU", "₡", "CSD", "CSK", "$", "$", "CVE", "CYP", "Kč", "DDM", "DEM", "DJF", "kr", "$", "DZD", "ECS", "ECV", "EEK", "E£", "ERN", "ESA", "ESB", "₧", "ETB", "€", "FIM", "$", "£", "FRF", "£", "GEK", "₾", "GHC", "GH₵", "£", "GMD", "FG", "GNS", "GQE", "GRD", "Q", "GWE", "GWP", "$", "HK$", "L", "HRD", "kn", "HTG", "Ft", "Rp", "IEP", "ILP", "ILR", "₪", "₹", "IQD", "IRR", "ISJ", "kr", "ITL", "$", "JOD", "JP¥", "KES", "⃀", "៛", "CF", "₩", "KRH", "KRO", "₩", "KWD", "$", "₸", "₭", "L£", "Rs", "$", "LSL", "Lt", "LTT", "LUC", "LUF", "LUL", "Ls", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "Ar", "MGF", "MKD", "MKN", "MLF", "K", "₮", "MOP", "MRO", "MRU", "MTL", "MTP", "Rs", "MVP", "MVR", "MWK", "MX$", "MXP", "MXV", "RM", "MZE", "MZM", "MZN", "$", "₦", "NIC", "C$", "NLG", "kr", "Rs", "NZ$", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "₱", "Rs", "zł", "PLZ", "PTE", "₲", "QAR", "RHD", "ROL", "lei", "RSD", "₽", "RUR", "RF", "\u20c1", "$", "SCR", "SDD", "SDG", "SDP", "kr", "$", "£", "SIT", "SKK", "SLE", "SLL", "SOS", "$", "SRG", "£", "STD", "Db", "SUR", "SVC", "£", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "T$", "TPE", "TRL", "₺", "$", "NT$", "TZS", "₴", "UAK", "UGS", "UGX", "US$", "USN", "USS", "UYI", "UYP", "$", "UYW", "soʻm", "VEB", "VED", "Bs", "VES", "₫", "VNN", "VUV", "WST", "FCFA", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "EC$", "Cg.", "XDR", "XEU", "XFO", "XFU", "F\u202fCFA", "XPD", "CFPF", "XPT", "XRE", "XSU", "XTS", "XUA", "¤", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "R", "ZMK", "ZK", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "جنو", "فبر", "مار", "اپر", "می", "جون", "جول", "اگس", "سپت", "اکت", "نوم", "دسم"},
		monthsNarrow:           []string{"", "1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12"},
		monthsWide:             []string{"", "جنوری", "فبروری", "مارچ", "اپریل", "می", "جون", "جولای", "اگست", "سپتمبر", "اکتوبر", "نومبر", "دسمبر"},
		daysAbbreviated:        []string{"ی.", "د.", "س.", "چ.", "پ.", "ج.", "ش."},
		daysNarrow:             []string{"S", "M", "T", "W", "T", "F", "S"},
		daysShort:              []string{"Ya", "Du", "Se", "Ch", "Pa", "Ju", "Sh"},
		daysWide:               []string{"یکشنبه", "دوشنبه", "سه\u200cشنبه", "چهارشنبه", "پنجشنبه", "جمعه", "شنبه"},
		timezones:              map[string]string{"ACDT": "Markaziy Avstraliya yozgi vaqti", "ACST": "Markaziy Avstraliya standart vaqti", "ACT": "ACT", "ACWDT": "Markaziy Avstraliya g‘arbiy yozgi vaqti", "ACWST": "Markaziy Avstraliya g‘arbiy standart vaqti", "ADT": "Atlantika yozgi vaqti", "ADT Arabia": "Saudiya Arabistoni yozgi vaqti", "AEDT": "Sharqiy Avstraliya yozgi vaqti", "AEST": "Sharqiy Avstraliya standart vaqti", "AFT": "افغانستان وقتی", "AKDT": "Alyaska yozgi vaqti", "AKST": "Alyaska standart vaqti", "AMST": "Amazonka yozgi vaqti", "AMST Armenia": "Armaniston yozgi vaqti", "AMT": "Amazonka standart vaqti", "AMT Armenia": "Armaniston standart vaqti", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Argentina yozgi vaqti", "ART": "Argentina standart vaqti", "AST": "Atlantika standart vaqti", "AST Arabia": "Saudiya Arabistoni standart vaqti", "AWDT": "G‘arbiy Avstraliya yozgi vaqti", "AWST": "G‘arbiy Avstraliya standart vaqti", "AZST": "Ozarbayjon yozgi vaqti", "AZT": "Ozarbayjon standart vaqti", "BDT Bangladesh": "Bangladesh yozgi vaqti", "BNT": "Bruney-Dorussalom vaqti", "BOT": "Boliviya vaqti", "BRST": "Braziliya yozgi vaqti", "BRT": "Braziliya standart vaqti", "BST Bangladesh": "Bangladesh standart vaqti", "BT": "Butan vaqti", "CAST": "CAST", "CAT": "Markaziy Afrika vaqti", "CCT": "Kokos orollari vaqti", "CDT": "Markaziy Amerika yozgi vaqti", "CHADT": "Chatem yozgi vaqti", "CHAST": "Chatem standart vaqti", "CHUT": "Chuuk vaqti", "CKT": "Kuk orollari standart vaqti", "CKT DST": "Kuk orollari yarim yozgi vaqti", "CLST": "Chili yozgi vaqti", "CLT": "Chili standart vaqti", "COST": "Kolumbiya yozgi vaqti", "COT": "Kolumbiya standart vaqti", "CST": "Markaziy Amerika standart vaqti", "CST China": "Xitoy standart vaqti", "CST China DST": "Xitoy yozgi vaqti", "CVST": "Kabo-Verde yozgi vaqti", "CVT": "Kabo-Verde standart vaqti", "CXT": "Rojdestvo oroli vaqti", "ChST": "Chamorro standart vaqti", "ChST NMI": "ChST NMI", "CuDT": "Kuba yozgi vaqti", "CuST": "Kuba standart vaqti", "DAVT": "Deyvis vaqti", "DDUT": "Dyumon-d’Yurvil vaqti", "EASST": "Pasxa oroli yozgi vaqti", "EAST": "Pasxa oroli standart vaqti", "EAT": "Sharqiy Afrika vaqti", "ECT": "Ekvador vaqti", "EDT": "Sharqiy Amerika yozgi vaqti", "EGDT": "Sharqiy Grenlandiya yozgi vaqti", "EGST": "Sharqiy Grenlandiya standart vaqti", "EST": "Sharqiy Amerika standart vaqti", "FEET": "Kaliningrad va Minsk vaqti", "FJT": "Fiji standart vaqti", "FJT Summer": "Fiji yozgi vaqti", "FKST": "Folklend orollari yozgi vaqti", "FKT": "Folklend orollari standart vaqti", "FNST": "Fernandu-di-Noronya yozgi vaqti", "FNT": "Fernandu-di-Noronya standart vaqti", "GALT": "Galapagos vaqti", "GAMT": "Gambye vaqti", "GEST": "Gruziya yozgi vaqti", "GET": "Gruziya standart vaqti", "GFT": "Fransuz Gvianasi vaqti", "GIT": "Gilbert orollari vaqti", "GMT": "Grinvich o‘rtacha vaqti", "GNSST": "GNSST", "GNST": "GNST", "GST": "Janubiy Georgiya vaqti", "GST Guam": "GST Guam", "GYT": "Gayana vaqti", "HADT": "Gavayi-aleut standart vaqti", "HAST": "Gavayi-aleut standart vaqti", "HKST": "Gonkong yozgi vaqti", "HKT": "Gonkong standart vaqti", "HOVST": "Xovd yozgi vaqti", "HOVT": "Xovd standart vaqti", "ICT": "Hindixitoy vaqti", "IDT": "Isroil yozgi vaqti", "IOT": "Hind okeani vaqti", "IRKST": "Irkutsk yozgi vaqti", "IRKT": "Irkutsk standart vaqti", "IRST": "Eron standart vaqti", "IRST DST": "Eron yozgi vaqti", "IST": "Hindiston standart vaqti", "IST Israel": "Isroil standart vaqti", "JDT": "Yaponiya yozgi vaqti", "JST": "Yaponiya standart vaqti", "KOST": "Kosrae vaqti", "KRAST": "Krasnoyarsk yozgi vaqti", "KRAT": "Krasnoyarsk standart vaqti", "KST": "Koreya standart vaqti", "KST DST": "Koreya yozgi vaqti", "LHDT": "Lord-Xau yozgi vaqti", "LHST": "Lord-Xau standart vaqti", "LINT": "Layn orollari vaqti", "MAGST": "Magadan yozgi vaqti", "MAGT": "Magadan standart vaqti", "MART": "Markiz orollari vaqti", "MAWT": "Mouson vaqti", "MDT": "Tog‘ yozgi vaqti (AQSH)", "MESZ": "Markaziy Yevropa yozgi vaqti", "MEZ": "Markaziy Yevropa standart vaqti", "MHT": "Marshall orollari vaqti", "MMT": "Myanma vaqti", "MSD": "Moskva yozgi vaqti", "MST": "Tog‘ standart vaqti (AQSH)", "MUST": "Mavrikiy yozgi vaqti", "MUT": "Mavrikiy standart vaqti", "MVT": "Maldiv orollari vaqti", "MYT": "Malayziya vaqti", "NCT": "Yangi Kaledoniya standart vaqti", "NDT": "Nyufaundlend yozgi vaqti", "NDT New Caledonia": "Yangi Kaledoniya yozgi vaqti", "NFDT": "Norfolk oroli yozgi vaqti", "NFT": "Norfolk oroli standart vaqti", "NOVST": "Novosibirsk yozgi vaqti", "NOVT": "Novosibirsk standart vaqti", "NPT": "Nepal vaqti", "NRT": "Nauru vaqti", "NST": "Nyufaundlend standart vaqti", "NUT": "Niuye vaqti", "NZDT": "Yangi Zelandiya yozgi vaqti", "NZST": "Yangi Zelandiya standart vaqti", "OESZ": "Sharqiy Yevropa yozgi vaqti", "OEZ": "Sharqiy Yevropa standart vaqti", "OMSST": "Omsk yozgi vaqti", "OMST": "Omsk standart vaqti", "PDT": "Tinch okeani yozgi vaqti", "PDTM": "Meksika Tinch okeani yozgi vaqti", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Papua-Yangi Gvineya vaqti", "PHOT": "Feniks orollari vaqti", "PKT": "Pokiston standart vaqti", "PKT DST": "Pokiston yozgi vaqti", "PMDT": "Sen-Pyer va Mikelon yozgi vaqti", "PMST": "Sen-Pyer va Mikelon standart vaqti", "PONT": "Ponape vaqti", "PST": "Tinch okeani standart vaqti", "PST Philippine": "Filippin standart vaqti", "PST Philippine DST": "Filippin yozgi vaqti", "PST Pitcairn": "Pitkern vaqti", "PSTM": "Meksika Tinch okeani standart vaqti", "PWT": "Palau vaqti", "PYST": "Paragvay yozgi vaqti", "PYT": "Paragvay standart vaqti", "PYT Korea": "Pxenyan vaqti", "RET": "Reyunion vaqti", "ROTT": "Rotera vaqti", "SAKST": "Saxalin yozgi vaqti", "SAKT": "Saxalin standart vaqti", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Janubiy Afrika standart vaqti", "SBT": "Solomon orollari vaqti", "SCT": "Seyshel orollari vaqti", "SGT": "Singapur vaqti", "SLST": "SLST", "SRT": "Surinam vaqti", "SST Samoa": "Samoa standart vaqti", "SST Samoa Apia": "Apia standart vaqti", "SST Samoa Apia DST": "Apia yozgi vaqti", "SST Samoa DST": "Samoa yozgi vaqti", "SYOT": "Syova vaqti", "TAAF": "Fransuz Janubiy hududlari va Antarktika vaqti", "TAHT": "Taiti vaqti", "TJT": "Tojikiston vaqti", "TKT": "Tokelau vaqti", "TLT": "Sharqiy Timor vaqti", "TMST": "Turkmaniston yozgi vaqti", "TMT": "Turkmaniston standart vaqti", "TOST": "Tonga yozgi vaqti", "TOT": "Tonga standart vaqti", "TVT": "Tuvalu vaqti", "TWT": "Tayvan standart vaqti", "TWT DST": "Tayvan yozgi vaqti", "ULAST": "Ulan-Bator yozgi vaqti", "ULAT": "Ulan-Bator standart vaqti", "UYST": "Urugvay yozgi vaqti", "UYT": "Urugvay standart vaqti", "UZT": "O‘zbekiston standart vaqti", "UZT DST": "O‘zbekiston yozgi vaqti", "VET": "Venesuela vaqti", "VLAST": "Vladivostok yozgi vaqti", "VLAT": "Vladivostok standart vaqti", "VOLST": "Volgograd yozgi vaqti", "VOLT": "Volgograd standart vaqti", "VOST": "Vostok vaqti", "VUT": "Vanuatu standart vaqti", "VUT DST": "Vanuatu yozgi vaqti", "WAKT": "Ueyk oroli vaqti", "WARST": "Gʻarbiy Argentina yozgi vaqti", "WART": "Gʻarbiy Argentina standart vaqti", "WAST": "Gʻarbiy Afrika vaqti", "WAT": "Gʻarbiy Afrika vaqti", "WESZ": "G‘arbiy Yevropa yozgi vaqti", "WEZ": "G‘arbiy Yevropa standart vaqti", "WFT": "Uollis va Futuna vaqti", "WGST": "G‘arbiy Grenlandiya yozgi vaqti", "WGT": "G‘arbiy Grenlandiya standart vaqti", "WIB": "Gʻarbiy Indoneziya vaqti", "WIT": "Sharqiy Indoneziya vaqti", "WITA": "Markaziy Indoneziya vaqti", "YAKST": "Yakutsk yozgi vaqti", "YAKT": "Yakutsk standart vaqti", "YEKST": "Yekaterinburg yozgi vaqti", "YEKT": "Yekaterinburg standart vaqti", "YST": "Yukon vaqti", "МСК": "Moskva standart vaqti", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Gʻarbiy Qozogʻiston vaqti", "شىعىش قازاق ەلى": "Sharqiy Qozogʻiston vaqti", "قازاق ەلى": "Qozogʻiston vaqti", "قىرعىزستان": "Qirgʻiziston vaqti", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Peru yozgi vaqti"},
	}
}

// Locale returns the current translators string locale
func (uz *uz_Arab) Locale() string {
	return uz.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'uz_Arab'
func (uz *uz_Arab) PluralsCardinal() []locales.PluralRule {
	return uz.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'uz_Arab'
func (uz *uz_Arab) PluralsOrdinal() []locales.PluralRule {
	return uz.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'uz_Arab'
func (uz *uz_Arab) PluralsRange() []locales.PluralRule {
	return uz.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'uz_Arab'
func (uz *uz_Arab) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'uz_Arab'
func (uz *uz_Arab) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'uz_Arab'
func (uz *uz_Arab) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := uz.CardinalPluralRule(num1, v1)
	end := uz.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (uz *uz_Arab) MonthAbbreviated(month time.Month) string {
	if len(uz.monthsAbbreviated) == 0 {
		return ""
	}
	return uz.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (uz *uz_Arab) MonthsAbbreviated() []string {
	return uz.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (uz *uz_Arab) MonthNarrow(month time.Month) string {
	if len(uz.monthsNarrow) == 0 {
		return ""
	}
	return uz.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (uz *uz_Arab) MonthsNarrow() []string {
	return uz.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (uz *uz_Arab) MonthWide(month time.Month) string {
	if len(uz.monthsWide) == 0 {
		return ""
	}
	return uz.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (uz *uz_Arab) MonthsWide() []string {
	return uz.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (uz *uz_Arab) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(uz.daysAbbreviated) == 0 {
		return ""
	}
	return uz.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (uz *uz_Arab) WeekdaysAbbreviated() []string {
	return uz.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (uz *uz_Arab) WeekdayNarrow(weekday time.Weekday) string {
	if len(uz.daysNarrow) == 0 {
		return ""
	}
	return uz.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (uz *uz_Arab) WeekdaysNarrow() []string {
	return uz.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (uz *uz_Arab) WeekdayShort(weekday time.Weekday) string {
	if len(uz.daysShort) == 0 {
		return ""
	}
	return uz.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (uz *uz_Arab) WeekdaysShort() []string {
	return uz.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (uz *uz_Arab) WeekdayWide(weekday time.Weekday) string {
	if len(uz.daysWide) == 0 {
		return ""
	}
	return uz.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (uz *uz_Arab) WeekdaysWide() []string {
	return uz.daysWide
}

// Decimal returns the decimal point of number
func (uz *uz_Arab) Decimal() string {
	return uz.decimal
}

// Group returns the group of number
func (uz *uz_Arab) Group() string {
	return uz.group
}

// Group returns the minus sign of number
func (uz *uz_Arab) Minus() string {
	return uz.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'uz_Arab' and handles both Whole and Real numbers based on 'v'
func (uz *uz_Arab) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uz.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, uz.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, uz.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'uz_Arab' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (uz *uz_Arab) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uz.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, uz.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, uz.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'uz_Arab'
func (uz *uz_Arab) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := uz.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uz.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, uz.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, uz.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, uz.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, uz.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'uz_Arab'
// in accounting notation.
func (uz *uz_Arab) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := uz.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, uz.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, uz.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, uz.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, uz.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, uz.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, uz.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'uz_Arab'
func (uz *uz_Arab) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'uz_Arab'
func (uz *uz_Arab) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, uz.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'uz_Arab'
func (uz *uz_Arab) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0xd9, 0x86, 0xda, 0x86, 0xdb, 0x8c, 0x20}...)
	b = append(b, uz.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'uz_Arab'
func (uz *uz_Arab) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20, 0xd9, 0x86, 0xda, 0x86, 0xdb, 0x8c, 0x20, 0xdb, 0x8c, 0xdb, 0x8c, 0xd9, 0x84, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0xd9, 0x86, 0xda, 0x86, 0xdb, 0x8c, 0x20}...)
	b = append(b, uz.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = append(b, uz.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20, 0xda, 0xa9, 0xd9, 0x88, 0xd9, 0x86, 0xdb, 0x8c}...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'uz_Arab'
func (uz *uz_Arab) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'uz_Arab'
func (uz *uz_Arab) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'uz_Arab'
func (uz *uz_Arab) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	b = append(b, []byte{0x29}...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'uz_Arab'
func (uz *uz_Arab) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, uz.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()

	if btz, ok := uz.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	b = append(b, []byte{0x29}...)

	return string(b)
}
