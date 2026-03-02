package sc_IT

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type sc_IT struct {
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

// New returns a new instance of translator for the 'sc_IT' locale
func New() locales.Translator {
	return &sc_IT{
		locale:                 "sc_IT",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{5, 6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "ghe", "fre", "mar", "abr", "maj", "làm", "trì", "aus", "cab", "stG", "stA", "nad"},
		monthsNarrow:           []string{"", "G", "F", "M", "A", "M", "L", "T", "A", "C", "S", "S", "N"},
		monthsWide:             []string{"", "ghennàrgiu", "freàrgiu", "martzu", "abrile", "maju", "làmpadas", "trìulas", "austu", "cabudanni", "santugaine", "santandria", "nadale"},
		daysAbbreviated:        []string{"dom", "lun", "mar", "mèr", "giò", "che", "sàb"},
		daysNarrow:             []string{"D", "L", "M", "M", "G", "C", "S"},
		daysWide:               []string{"domìniga", "lunis", "martis", "mèrcuris", "giòbia", "chenàbura", "sàbadu"},
		timezones:              map[string]string{"ACDT": "Ora legale de s’Austràlia tzentrale", "ACST": "Ora legale de Acre", "ACT": "Ora istandard de Acre", "ACWDT": "Ora legale de s’Austràlia tzentru-otzidentale", "ACWST": "Ora istandard de s’Austràlia tzentru-otzidentale", "ADT": "Ora legale de s’Atlànticu", "ADT Arabia": "Ora legale àraba", "AEDT": "Ora legale de s’Austràlia orientale", "AEST": "Ora istandard de s’Austràlia orientale", "AFT": "Ora de s’Afghànistan", "AKDT": "Ora legale de s’Alaska", "AKST": "Ora istandard de s’Alaska", "AMST": "Ora legale de s’Amatzònia", "AMST Armenia": "Ora legale de s’Armènia", "AMT": "Ora istandard de s’Amatzònia", "AMT Armenia": "Ora istandard de s’Armènia", "ANAST": "Ora legale de Anadyr", "ANAT": "Ora istandard de Anadyr", "ARST": "Ora legale de s’Argentina", "ART": "Ora istandard de s’Argentina", "AST": "Ora istandard de s’Atlànticu", "AST Arabia": "Ora istandard àraba", "AWDT": "Ora legale de s’Austràlia otzidentale", "AWST": "Ora istandard de s’Austràlia otzidentale", "AZST": "Ora legale de s’Azerbaigiàn", "AZT": "Ora istandard de s’Azerbaigiàn", "BDT Bangladesh": "Ora legale de su Bangladesh", "BNT": "Ora de su Brunei", "BOT": "Ora de sa Bolìvia", "BRST": "Ora legale de Brasìlia", "BRT": "Ora istandard de Brasìlia", "BST Bangladesh": "Ora istandard de su Bangladesh", "BT": "Ora de su Bhutàn", "CAST": "Ora de Casey", "CAT": "Ora de s’Àfrica tzentrale", "CCT": "Ora de sas Ìsulas Cocos", "CDT": "Ora legale tzentrale USA", "CHADT": "Ora legale de sas Chatham", "CHAST": "Ora istandard de sas Chatham", "CHUT": "Ora de su Chuuk", "CKT": "Ora istandard de sas Ìsulas Cook", "CKT DST": "Ora legale de sas Ìsulas Cook", "CLST": "Ora legale de su Tzile", "CLT": "Ora istandard de su Tzile", "COST": "Ora legale de sa Colòmbia", "COT": "Ora istandard de sa Colòmbia", "CST": "Ora istandard tzentrale USA", "CST China": "Ora istandard de sa Tzina", "CST China DST": "Ora legale de sa Tzina", "CVST": "Ora legale de su Cabu Birde", "CVT": "Ora istandard de su Cabu Birde", "CXT": "Ora de s’Ìsula de sa Natividade", "ChST": "Ora istandard de Chamorro", "ChST NMI": "Ora de sas Ìsulas Mariannas Setentrionales", "CuDT": "Ora legale de Cuba", "CuST": "Ora istandard de Cuba", "DAVT": "Ora de Davis", "DDUT": "Ora de Dumont-d’Urville", "EASST": "Ora legale de s’Ìsula de Pasca", "EAST": "Ora istandard de s’Ìsula de Pasca", "EAT": "Ora de s’Àfrica orientale", "ECT": "Ora de s’Ecuador", "EDT": "Ora legale orientale USA", "EGDT": "Ora legale de sa Groenlàndia orientale", "EGST": "Ora istandard de sa Groenlàndia orientale", "EST": "Ora istandard orientale USA", "FEET": "Ora de s’estremu oriente europeu (Kaliningrad)", "FJT": "Ora istandard de sas Fiji", "FJT Summer": "Ora legale de sas Fiji", "FKST": "Ora legale de sas Ìsulas Falkland", "FKT": "Ora istandard de sas Ìsulas Falkland", "FNST": "Ora legale de su Fernando de Noronha", "FNT": "Ora istandard de su Fernando de Noronha", "GALT": "Ora de sas Galàpagos", "GAMT": "Ora de Gambier", "GEST": "Ora legale de sa Geòrgia", "GET": "Ora istandard de sa Geòrgia", "GFT": "Ora de sa Guiana Frantzesa", "GIT": "Ora de sas Ìsulas Gilbert", "GMT": "Ora de su meridianu de Greenwich", "GNSST": "Ora legale de sa Groenlàndia", "GNST": "Ora istandard de sa Groenlàndia", "GST": "Ora istandard de su Gulfu", "GST Guam": "Ora istandard de Guàm", "GYT": "Ora de sa Guyana", "HADT": "Ora istandard de sas ìsulas Hawaii-Aleutinas", "HAST": "Ora istandard de sas ìsulas Hawaii-Aleutinas", "HKST": "Ora legale de Hong Kong", "HKT": "Ora istandard de Hong Kong", "HOVST": "Ora legale de Hovd", "HOVT": "Ora istandard de Hovd", "ICT": "Ora de s’Indotzina", "IDT": "Ora legale de Israele", "IOT": "Ora de s’Otzèanu Indianu", "IRKST": "Ora legale de Irkutsk", "IRKT": "Ora istandard de Irkutsk", "IRST": "Ora istandard de s’Iràn", "IRST DST": "Ora legale de s’Iràn", "IST": "Ora istandard de s’Ìndia", "IST Israel": "Ora istandard de Israele", "JDT": "Ora legale de su Giapone", "JST": "Ora istandard de su Giapone", "KOST": "Ora de Kosrae", "KRAST": "Ora legale de Krasnoyarsk", "KRAT": "Ora istandard de Krasnoyarsk", "KST": "Ora istandard coreana", "KST DST": "Ora legale coreana", "LHDT": "Ora legale de Lord Howe", "LHST": "Ora istandard de Lord Howe", "LINT": "Ora de sas Ìsulas de sa Lìnia", "MAGST": "Ora legale de Magadan", "MAGT": "Ora istandard de Magadan", "MART": "Ora de sas Marchesas", "MAWT": "Ora de Mawson", "MDT": "Ora legale de Macao", "MESZ": "Ora legale de s’Europa tzentrale", "MEZ": "Ora istandard de s’Europa tzentrale", "MHT": "Ora de sas Ìsulas Marshall", "MMT": "Ora de su Myanmàr", "MSD": "Ora legale de Mosca", "MST": "Ora istandard de Macao", "MUST": "Ora legale de sas Maurìtzius", "MUT": "Ora istandard de sas Maurìtzius", "MVT": "Ora de sas Maldivas", "MYT": "Ora de sa Malèsia", "NCT": "Ora istandard de sa Caledònia Noa", "NDT": "Ora legale de Terranova", "NDT New Caledonia": "Ora legale de sa Caledònia Noa", "NFDT": "Ora legale de s’Ìsula Norfolk", "NFT": "Ora istandard de s’Ìsula Norfolk", "NOVST": "Ora legale de Novosibirsk", "NOVT": "Ora istandard de Novosibirsk", "NPT": "Ora de su Nepal", "NRT": "Ora de Nauru", "NST": "Ora istandard de Terranova", "NUT": "Ora de Niue", "NZDT": "Ora legale de sa Zelanda Noa", "NZST": "Ora istandard de sa Zelanda Noa", "OESZ": "Ora legale de s’Europa orientale", "OEZ": "Ora istandard de s’Europa orientale", "OMSST": "Ora legale de Omsk", "OMST": "Ora istandard de Omsk", "PDT": "Ora legale de su Patzìficu USA", "PDTM": "Ora legale de su Patzìficu (Mèssicu)", "PETDT": "Ora legale de Petropavlovsk-Kamchatski", "PETST": "Ora istandard de Petropavlovsk-Kamchatski", "PGT": "Ora de sa Pàpua Guinea Noa", "PHOT": "Ora de sas Ìsulas de sa Fenìtzie", "PKT": "Ora istandard de su Pàkistan", "PKT DST": "Ora legale de su Pàkistan", "PMDT": "Ora legale de Saint-Pierre e Miquelon", "PMST": "Ora istandard de Saint-Pierre e Miquelon", "PONT": "Ora de Pohnpei", "PST": "Ora istandard de su Patzìficu USA", "PST Philippine": "Ora istandard de sas Filipinas", "PST Philippine DST": "Ora legale de sas Filipinas", "PST Pitcairn": "Ora de sas Pitcairn", "PSTM": "Ora istandard de su Patzìficu (Mèssicu)", "PWT": "Ora de Palau", "PYST": "Ora legale de su Paraguay", "PYT": "Ora istandard de su Paraguay", "PYT Korea": "Ora de Pyongyang", "RET": "Ora de sa Reunione", "ROTT": "Ora de Rothera", "SAKST": "Ora legale de Sakhalin", "SAKT": "Ora istandard de Sakhalin", "SAMST": "Ora legale de Samara", "SAMT": "Ora istandard de Samara", "SAST": "Ora istandard de s’Àfrica meridionale", "SBT": "Ora de sas Ìsulas Salomone", "SCT": "Ora de sas Seychelles", "SGT": "Ora de Singapore", "SLST": "Ora de Lanka", "SRT": "Ora de su Suriname", "SST Samoa": "Ora istandard de sas Samoa", "SST Samoa Apia": "Ora istandard de Apia", "SST Samoa Apia DST": "Ora legale de Apia", "SST Samoa DST": "Ora legale de sas Samoa", "SYOT": "Ora de Syowa", "TAAF": "Ora de sa Terras australes e antàrticas frantzesas", "TAHT": "Ora de Tahiti", "TJT": "Ora de su Tagìkistan", "TKT": "Ora de su Tokelau", "TLT": "Ora de su Timor Est", "TMST": "Ora legale de su Turkmènistan", "TMT": "Ora istandard de su Turkmènistan", "TOST": "Ora legale de su Tonga", "TOT": "Ora istandard de su Tonga", "TVT": "Ora de su Tuvalu", "TWT": "Ora istandard de Taipei", "TWT DST": "Ora legale de Taipei", "ULAST": "Ora legale de Ulàn Bator", "ULAT": "Ora istandard de Ulàn Bator", "UYST": "Ora legale de s’Uruguay", "UYT": "Ora istandard de s’Uruguay", "UZT": "Ora istandard de s’Uzbèkistan", "UZT DST": "Ora legale de s’Uzbèkistan", "VET": "Ora de su Venetzuela", "VLAST": "Ora legale de Vladivostok", "VLAT": "Ora istandard de Vladivostok", "VOLST": "Ora legale de Volgograd", "VOLT": "Ora istandard de Volgograd", "VOST": "Ora de Vostok", "VUT": "Ora istandard de su Vanuatu", "VUT DST": "Ora legale de su Vanuatu", "WAKT": "Ora de sas Ìsulas Wake", "WARST": "Ora legale de s’Argentina otzidentale", "WART": "Ora istandard de s’Argentina otzidentale", "WAST": "Ora de s’Àfrica otzidentale", "WAT": "Ora de s’Àfrica otzidentale", "WESZ": "Ora legale de s’Europa otzidentale", "WEZ": "Ora istandard de s’Europa otzidentale", "WFT": "Ora de Wallis e Futuna", "WGST": "Ora legale de sa Groenlàndia otzidentale", "WGT": "Ora istandard de sa Groenlàndia otzidentale", "WIB": "Ora de s’Indonèsia otzidentale", "WIT": "Ora de s’Indonèsia orientale", "WITA": "Ora de s’Indonèsia tzentrale", "YAKST": "Ora legale de Yakutsk", "YAKT": "Ora istandard de Yakutsk", "YEKST": "Ora legale de Yekaterinburg", "YEKT": "Ora istandard de Yekaterinburg", "YST": "Ora de su Yukon", "МСК": "Ora istandard de Mosca", "اقتاۋ": "Ora istandard de Aktau", "اقتاۋ قالاسى": "Ora legale de Aktau", "اقتوبە": "Ora istandard de Aktobe", "اقتوبە قالاسى": "Ora legale de Aktobe", "الماتى": "Ora istandard de Almaty", "الماتى قالاسى": "Ora legale de Almaty", "باتىس قازاق ەلى": "Ora de su Kazàkistan otzidentale", "شىعىش قازاق ەلى": "Ora de su Kazàkistan orientale", "قازاق ەلى": "Ora de su Kazàkistan", "قىرعىزستان": "Ora de su Kirghìzistan", "قىزىلوردا": "Ora istandard de Qyzylorda", "قىزىلوردا قالاسى": "Ora legale de Qyzylorda", "∅∅∅": "Ora legale de su Perù"},
	}
}

// Locale returns the current translators string locale
func (sc *sc_IT) Locale() string {
	return sc.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sc_IT'
func (sc *sc_IT) PluralsCardinal() []locales.PluralRule {
	return sc.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sc_IT'
func (sc *sc_IT) PluralsOrdinal() []locales.PluralRule {
	return sc.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'sc_IT'
func (sc *sc_IT) PluralsRange() []locales.PluralRule {
	return sc.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sc_IT'
func (sc *sc_IT) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	i := int64(n)

	if i == 1 && v == 0 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sc_IT'
func (sc *sc_IT) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n == 11 || n == 8 || n == 80 || n == 800 {
		return locales.PluralRuleMany
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sc_IT'
func (sc *sc_IT) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	start := sc.CardinalPluralRule(num1, v1)
	end := sc.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (sc *sc_IT) MonthAbbreviated(month time.Month) string {
	return sc.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (sc *sc_IT) MonthsAbbreviated() []string {
	return sc.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (sc *sc_IT) MonthNarrow(month time.Month) string {
	return sc.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (sc *sc_IT) MonthsNarrow() []string {
	return sc.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (sc *sc_IT) MonthWide(month time.Month) string {
	return sc.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (sc *sc_IT) MonthsWide() []string {
	return sc.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (sc *sc_IT) WeekdayAbbreviated(weekday time.Weekday) string {
	return sc.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (sc *sc_IT) WeekdaysAbbreviated() []string {
	return sc.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (sc *sc_IT) WeekdayNarrow(weekday time.Weekday) string {
	return sc.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (sc *sc_IT) WeekdaysNarrow() []string {
	return sc.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (sc *sc_IT) WeekdayShort(weekday time.Weekday) string {
	return sc.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (sc *sc_IT) WeekdaysShort() []string {
	return sc.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (sc *sc_IT) WeekdayWide(weekday time.Weekday) string {
	return sc.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (sc *sc_IT) WeekdaysWide() []string {
	return sc.daysWide
}

// Decimal returns the decimal point of number
func (sc *sc_IT) Decimal() string {
	return sc.decimal
}

// Group returns the group of number
func (sc *sc_IT) Group() string {
	return sc.group
}

// Group returns the minus sign of number
func (sc *sc_IT) Minus() string {
	return sc.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sc_IT' and handles both Whole and Real numbers based on 'v'
func (sc *sc_IT) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sc.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sc.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sc.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'sc_IT' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sc *sc_IT) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sc.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sc.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, sc.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sc_IT'
func (sc *sc_IT) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sc.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sc.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sc.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sc.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sc.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, sc.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sc_IT'
// in accounting notation.
func (sc *sc_IT) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sc.currencies[currency]
	l := len(s) + len(symbol) + 6 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sc.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sc.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sc.currencyNegativePrefix[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sc.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, sc.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, sc.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'sc_IT'
func (sc *sc_IT) FmtDateShort(t time.Time) string {
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

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'sc_IT'
func (sc *sc_IT) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, sc.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'sc_IT'
func (sc *sc_IT) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, sc.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20, 0x73, 0x75}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'sc_IT'
func (sc *sc_IT) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, sc.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20}...)
	b = append(b, sc.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65}...)
	b = append(b, []byte{0x20, 0x73, 0x75}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'sc_IT'
func (sc *sc_IT) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sc.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'sc_IT'
func (sc *sc_IT) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sc.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sc.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'sc_IT'
func (sc *sc_IT) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sc.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sc.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'sc_IT'
func (sc *sc_IT) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, sc.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sc.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := sc.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
