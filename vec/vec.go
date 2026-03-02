package vec

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type vec struct {
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
	currencyPositiveSuffix string
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

// New returns a new instance of translator for the 'vec' locale
func New() locales.Translator {
	return &vec{
		locale:                 "vec",
		pluralsCardinal:        []locales.PluralRule{2, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{5, 6},
		pluralsRange:           nil,
		decimal:                ",",
		group:                  " ",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "L", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "Cf", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "jen", "feb", "mar", "apr", "maj", "jug", "luj", "ago", "set", "oto", "nov", "dez"},
		monthsNarrow:           []string{"", "J", "F", "M", "A", "M", "J", "L", "A", "S", "O", "N", "D"},
		monthsWide:             []string{"", "jenaro", "febraro", "marso", "aprile", "majo", "jugno", "lujo", "agosto", "setenbre", "otobre", "novenbre", "dezenbre"},
		daysAbbreviated:        []string{"dom", "lun", "mar", "mer", "zob", "vèn", "sab"},
		daysNarrow:             []string{"D", "L", "M", "M", "Z", "V", "S"},
		daysWide:               []string{"doménega", "luni", "marti", "mèrcore", "zoba", "vènare", "sabo"},
		erasAbbreviated:        []string{"v.C.", "d.C."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"vanti Cristo", "daspò Cristo"},
		timezones:              map[string]string{"ACDT": "Ora d’istà de l’Australia sentrale", "ACST": "Ora normale de l’Australia sentrale", "ACT": "ACT", "ACWDT": "Ora d’istà de l’Australia sentro osidentale", "ACWST": "Ora normale de l’Australia sentro osidentale", "ADT": "Ora d’istà de l’Atlàntego", "ADT Arabia": "Ora d’istà de l’Arabia", "AEDT": "Ora d’istà de l’Australia orientale", "AEST": "Ora normale de l’Australia orientale", "AFT": "Ora de l’Afghanistan", "AKDT": "Ora d’istà de l’Alaska", "AKST": "Ora normale de l’Alaska", "AMST": "Ora d’istà de l’Amasonia", "AMST Armenia": "Ora d’istà de l’Armenia", "AMT": "Ora normale de l’Amasonia", "AMT Armenia": "Ora normale de l’Armenia", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "Ora d’istà de l’Arjentina", "ART": "Ora normale de l’Arjentina", "AST": "Ora normale de l’Atlàntego", "AST Arabia": "Ora normale de l’Arabia", "AWDT": "Ora d’istà de l’Australia osidentale", "AWST": "Ora normale de l’Australia osidentale", "AZST": "Ora d’istà de l’Azerbaijan", "AZT": "Ora normale de l’Azerbaijan", "BDT Bangladesh": "Ora d’istà de’l Bangladesh", "BNT": "Ora de’l Brunéi", "BOT": "Ora de la Bolivia", "BRST": "Ora d’istà de Brazilia", "BRT": "Ora normale de Brazilia", "BST Bangladesh": "Ora normale de’l Bangladesh", "BT": "Ora de’l Butan", "CAST": "CAST", "CAT": "Ora de l’Àfrega sentrale", "CCT": "Ora de le Ìzole Cocos", "CDT": "Ora d’istà de’l nord Amèrega sentrale", "CHADT": "Ora d’istà de le Ìzole Ciatem", "CHAST": "Ora normale de le Ìzole Ciatem", "CHUT": "Ora de’l Chuuk", "CKT": "Ora normale de le Ìzole Cook", "CKT DST": "Ora d’istà de le Ìzole Cook", "CLST": "Ora d’istà de’l Cile", "CLT": "Ora normale de’l Cile", "COST": "Ora d’istà de la Colonbia", "COT": "Ora normale de la Colonbia", "CST": "Ora normale de’l nord Amèrega sentrale", "CST China": "Ora normale de la Sina", "CST China DST": "Ora d’istà de la Sina", "CVST": "Ora d’istà de Cao Verdo", "CVT": "Ora normale de Cao Verdo", "CXT": "Ora de l’Ìzola de Nadale", "ChST": "Ora de Chamoro", "ChST NMI": "ChST NMI", "CuDT": "Ora d’istà de Cuba", "CuST": "Ora normale de Cuba", "DAVT": "Ora de Davis", "DDUT": "Ora de Dumont d’Urville", "EASST": "Ora d’istà de l’Ìzola de Pascua", "EAST": "Ora normale de l’Ìzola de Pascua", "EAT": "Ora de l’Àfrega orientale", "ECT": "Ora de l’Ècuador", "EDT": "Ora d’istà de’l nord Amèrega orientale", "EGDT": "Ora d’istà de la Groenlanda orientale", "EGST": "Ora normale de la Groenlanda orientale", "EST": "Ora normale de’l nord Amèrega orientale", "FEET": "Ora de l’Europa stra orientale", "FJT": "Ora normale de le Ìzole Fiji", "FJT Summer": "Ora d’istà de le Ìzole Fiji", "FKST": "Ora d’istà de le Ìzole Malvine", "FKT": "Ora normale de le Ìzole Malvine", "FNST": "Ora d’istà de Fernando de Noronha", "FNT": "Ora normale de Fernando de Noronha", "GALT": "Ora de le Galàpagos", "GAMT": "Ora de le Ìzole Gambier", "GEST": "Ora d’istà de la Jeorja", "GET": "Ora normale de la Jeorja", "GFT": "Ora de la Guyana franseze", "GIT": "Ora de le Ìzole Gilbert", "GMT": "Ora de’l meridian de Greenwich", "GNSST": "GNSST", "GNST": "GNST", "GST": "Ora de’l Golfo", "GST Guam": "GST Guam", "GYT": "Ora de la Guyana", "HADT": "Ora normale de Hawai e Aleutine", "HAST": "Ora normale de Hawai e Aleutine", "HKST": "Ora d’istà de Hong Kong", "HKT": "Ora normale de Hong Kong", "HOVST": "Ora d’istà de Hovd", "HOVT": "Ora normale de Hovd", "ICT": "Ora de l’Indosina", "IDT": "Ora d’istà de Izraele", "IOT": "Ora de l’Osèano Indian", "IRKST": "Ora d’istà de Irkutsk", "IRKT": "Ora normale de Irkutsk", "IRST": "Ora normale de l’Iran", "IRST DST": "Ora d’istà de l’Iran", "IST": "Ora de l’India", "IST Israel": "Ora normale de Izraele", "JDT": "Ora d’istà de’l Japon", "JST": "Ora normale de’l Japon", "KOST": "Ora de l’Ìzola Kosrae", "KRAST": "Ora d’istà de Krasnoyarsk", "KRAT": "Ora normale de Krasnoyarsk", "KST": "Ora normale de la Corèa", "KST DST": "Ora d’istà de la Corèa", "LHDT": "Ora d’istà de l’Ìzola Lord Howe", "LHST": "Ora normale de l’Ìzola Lord Howe", "LINT": "Ora de le Ìzole Ecuatoriali", "MAGST": "Ora d’istà de Magadan", "MAGT": "Ora normale de Magadan", "MART": "Ora de le Ìzole Marchezi", "MAWT": "Ora de Mawson", "MDT": "MDT", "MESZ": "Ora d’istà de l’Europa sentrale", "MEZ": "Ora normale de l’Europa sentrale", "MHT": "Ora de le Ìzole Marshall", "MMT": "Ora de Myanmar (Birmania)", "MSD": "Ora d’istà de Mosca", "MST": "MST", "MUST": "Ora d’istà de le Ìzole Maurisio", "MUT": "Ora normale de le Ìzole Maurisio", "MVT": "Ora de le Ìzole Maldive", "MYT": "Ora de la Malezia", "NCT": "Ora normale de la Nova Caledonia", "NDT": "Ora d’istà de Teranova", "NDT New Caledonia": "Ora d’istà de la Nova Caledonia", "NFDT": "Ora d’istà de l’Ìzola Norfolk", "NFT": "Ora normale de l’Ìzola Norfolk", "NOVST": "Ora d’istà de Novosibirsk", "NOVT": "Ora normale de Novosibirsk", "NPT": "Ora de’l Nepal", "NRT": "Ora de l’Ìzola Nauru", "NST": "Ora normale de Teranova", "NUT": "Ora de l’Ìzola Niue", "NZDT": "Ora d’istà de la Nova Zelanda", "NZST": "Ora normale de la Nova Zelanda", "OESZ": "Ora d’istà de l’Europa orientale", "OEZ": "Ora normale de l’Europa orientale", "OMSST": "Ora d’istà de Omsk", "OMST": "Ora normale de Omsk", "PDT": "Ora d’istà de’l Pasìfego", "PDTM": "Ora d’istà de’l Mèsego de’l Pasìfego", "PETDT": "PETDT", "PETST": "PETST", "PGT": "Ora de la Papua Nova Guinèa", "PHOT": "Ora de le Ìzole Fenize", "PKT": "Ora normale de’l Pakistan", "PKT DST": "Ora d’istà de’l Pakistan", "PMDT": "Ora d’istà de S. Piero e Michelon", "PMST": "Ora normale de S. Piero e Michelon", "PONT": "Ora de l’Ìzola Ponpèi", "PST": "Ora normale de’l Pasìfego", "PST Philippine": "Ora normale de le Ìzole Filipine", "PST Philippine DST": "Ora d’istà de le Ìzole Filipine", "PST Pitcairn": "Ora de le Ìzole Pitcairn", "PSTM": "Ora normale de’l Mèsego de’l Pasìfego", "PWT": "Ora de le Ìzole Palàu", "PYST": "Ora d’istà de’l Paraguài", "PYT": "Ora normale de’l Paraguài", "PYT Korea": "Ora de Pyongyang", "RET": "Ora de l’Ìzola Reunion", "ROTT": "Ora de Rothera", "SAKST": "Ora d’istà de Sakalin", "SAKT": "Ora normale de Sakalin", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "Ora de l’Àfrega meridionale", "SBT": "Ora de le Ìzole Salomon", "SCT": "Ora d’istà de le Ìzole Seisel", "SGT": "Ora de Singapore", "SLST": "SLST", "SRT": "Ora de’l Suriname", "SST Samoa": "Ora normale de le Ìzole Samòa", "SST Samoa Apia": "Ora normale de Apia", "SST Samoa Apia DST": "Ora d’istà de Apia", "SST Samoa DST": "Ora d’istà de le Ìzole Samòa", "SYOT": "Ora de Syowa", "TAAF": "Ora de le Tere fransezi de’l sud e de l’Antàrtego", "TAHT": "Ora de l’Ìzola Taiti", "TJT": "Ora de’l Tajikistan", "TKT": "Ora de le Ìzole Tokelàu", "TLT": "Ora de’l Timor Est", "TMST": "Ora d’istà de’l Turkmenistan", "TMT": "Ora normale de’l Turkmenistan", "TOST": "Ora d’istà de le Ìzole Tonga", "TOT": "Ora normale de le Ìzole Tonga", "TVT": "Ora de le Ìzole Tuvalu", "TWT": "Ora normale de Taipéi", "TWT DST": "Ora d’istà de Taipéi", "ULAST": "Ora d’istà de Ulan Bàtor", "ULAT": "Ora normale de Ulan Bàtor", "UYST": "Ora d’istà de l’Uruguài", "UYT": "Ora normale de l’Uruguài", "UZT": "Ora normale de l’Uzbekistan", "UZT DST": "Ora d’istà de l’Uzbekistan", "VET": "Ora de’l Venesuela", "VLAST": "Ora d’istà de Vladivostok", "VLAT": "Ora normale de Vladivostok", "VOLST": "Ora d’istà de Volgogrado", "VOLT": "Ora normale de Volgogrado", "VOST": "Ora de Vostok", "VUT": "Ora normale de le Ìzole Vanuatu", "VUT DST": "Ora d’istà de le Ìzole Vanuatu", "WAKT": "Ora de l’Atolo Wake", "WARST": "Ora d’istà de l’Arjentina osidentale", "WART": "Ora normale de l’Arjentina osidentale", "WAST": "Ora de l’Àfrega osidentale", "WAT": "Ora de l’Àfrega osidentale", "WESZ": "Ora d’istà de l’Europa osidentale", "WEZ": "Ora normale de l’Europa osidentale", "WFT": "Ora de le Ìzole Wallis e Futuna", "WGST": "Ora d’istà de la Groenlanda osidentale", "WGT": "Ora normale de la Groenlanda osidentale", "WIB": "Ora de l’Indonezia osidentale", "WIT": "Ora de l’Indonezia orientale", "WITA": "Ora de l’Indonezia sentrale", "YAKST": "Ora d’istà de Yakutsk", "YAKT": "Ora normale de Yakutsk", "YEKST": "Ora d’istà de Ekaterinburgo", "YEKT": "Ora normale de Ekaterinburgo", "YST": "Ora de’l Yukon", "МСК": "Ora normale de Mosca", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Ora de’l Kazakistan osidentale", "شىعىش قازاق ەلى": "Ora de’l Kazakistan orientale", "قازاق ەلى": "Ora de’l Kazakistan", "قىرعىزستان": "Ora de’l Kirghizistan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Ora d’istà de’l Perù"},
	}
}

// Locale returns the current translators string locale
func (vec *vec) Locale() string {
	return vec.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'vec'
func (vec *vec) PluralsCardinal() []locales.PluralRule {
	return vec.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'vec'
func (vec *vec) PluralsOrdinal() []locales.PluralRule {
	return vec.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'vec'
func (vec *vec) PluralsRange() []locales.PluralRule {
	return vec.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'vec'
func (vec *vec) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

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

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'vec'
func (vec *vec) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 11 || n == 8 || n == 80 || n == 800 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'vec'
func (vec *vec) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (vec *vec) MonthAbbreviated(month time.Month) string {
	return vec.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (vec *vec) MonthsAbbreviated() []string {
	return vec.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (vec *vec) MonthNarrow(month time.Month) string {
	return vec.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (vec *vec) MonthsNarrow() []string {
	return vec.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (vec *vec) MonthWide(month time.Month) string {
	return vec.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (vec *vec) MonthsWide() []string {
	return vec.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (vec *vec) WeekdayAbbreviated(weekday time.Weekday) string {
	return vec.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (vec *vec) WeekdaysAbbreviated() []string {
	return vec.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (vec *vec) WeekdayNarrow(weekday time.Weekday) string {
	return vec.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (vec *vec) WeekdaysNarrow() []string {
	return vec.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (vec *vec) WeekdayShort(weekday time.Weekday) string {
	return vec.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (vec *vec) WeekdaysShort() []string {
	return vec.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (vec *vec) WeekdayWide(weekday time.Weekday) string {
	return vec.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (vec *vec) WeekdaysWide() []string {
	return vec.daysWide
}

// Decimal returns the decimal point of number
func (vec *vec) Decimal() string {
	return vec.decimal
}

// Group returns the group of number
func (vec *vec) Group() string {
	return vec.group
}

// Group returns the minus sign of number
func (vec *vec) Minus() string {
	return vec.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'vec' and handles both Whole and Real numbers based on 'v'
func (vec *vec) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 3*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, vec.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(vec.group) - 1; j >= 0; j-- {
					b = append(b, vec.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, vec.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'vec' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (vec *vec) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, vec.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, vec.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, vec.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'vec'
func (vec *vec) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := vec.currencies[currency]
	l := len(s) + len(symbol) + 5 + 3*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, vec.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(vec.group) - 1; j >= 0; j-- {
					b = append(b, vec.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, vec.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, vec.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, vec.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'vec'
// in accounting notation.
func (vec *vec) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := vec.currencies[currency]
	l := len(s) + len(symbol) + 5 + 3*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, vec.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				for j := len(vec.group) - 1; j >= 0; j-- {
					b = append(b, vec.group[j])
				}
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, vec.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, vec.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, vec.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, vec.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'vec'
func (vec *vec) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'vec'
func (vec *vec) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, vec.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'vec'
func (vec *vec) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, vec.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'vec'
func (vec *vec) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, vec.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, vec.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'vec'
func (vec *vec) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, vec.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'vec'
func (vec *vec) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, vec.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, vec.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'vec'
func (vec *vec) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, vec.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, vec.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'vec'
func (vec *vec) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, vec.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, vec.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := vec.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
