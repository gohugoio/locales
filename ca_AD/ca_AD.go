package ca_AD

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ca_AD struct {
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

// New returns a new instance of translator for the 'ca_AD' locale
func New() locales.Translator {
	return &ca_AD{
		locale:                 "ca_AD",
		pluralsCardinal:        []locales.PluralRule{2, 5, 6},
		pluralsOrdinal:         []locales.PluralRule{2, 3, 4, 6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositiveSuffix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: " )",
		monthsAbbreviated:      []string{"", "de gen.", "de febr.", "de març", "d’abr.", "de maig", "de juny", "de jul.", "d’ag.", "de set.", "d’oct.", "de nov.", "de des."},
		monthsNarrow:           []string{"", "GN", "FB", "MÇ", "AB", "MG", "JN", "JL", "AG", "ST", "OC", "NV", "DS"},
		monthsWide:             []string{"", "de gener", "de febrer", "de març", "d’abril", "de maig", "de juny", "de juliol", "d’agost", "de setembre", "d’octubre", "de novembre", "de desembre"},
		daysAbbreviated:        []string{"dg.", "dl.", "dt.", "dc.", "dj.", "dv.", "ds."},
		daysNarrow:             []string{"dg.", "dl.", "dt.", "dc.", "dj.", "dv.", "ds."},
		daysWide:               []string{"diumenge", "dilluns", "dimarts", "dimecres", "dijous", "divendres", "dissabte"},
		timezones:              map[string]string{"ACDT": "Hora d’estiu d’Austràlia central", "ACST": "ACST", "ACT": "ACT", "ACWDT": "Hora d’estiu d’Austràlia centre-occidental", "ACWST": "Hora estàndard d’Austràlia centre-occidental", "ADT": "Hora d’estiu de l’Atlàntic", "ADT Arabia": "Hora d’estiu àrab", "AEDT": "Hora d’estiu d’Austràlia oriental", "AEST": "Hora estàndard d’Austràlia oriental", "AFT": "Hora de l’Afganistan", "AKDT": "Hora d’estiu d’Alaska", "AKST": "Hora estàndard d’Alaska", "AMST": "Hora d’estiu de l’Amazones", "AMST Armenia": "Hora d’estiu d’Armènia", "AMT": "Hora estàndard de l’Amazones", "AMT Armenia": "Hora estàndard d’Armènia", "ANAST": "Horari d’estiu d’Anadyr", "ANAT": "Hora estàndard d’Anadyr", "ARST": "Hora d’estiu de l’Argentina", "ART": "Hora estàndard de l’Argentina", "AST": "Hora estàndard de l’Atlàntic", "AST Arabia": "Hora estàndard àrab", "AWDT": "Hora d’estiu d’Austràlia occidental", "AWST": "Hora estàndard d’Austràlia occidental", "AZST": "Hora d’estiu de l’Azerbaidjan", "AZT": "Hora estàndard de l’Azerbaidjan", "BDT Bangladesh": "Hora d’estiu de Bangladesh", "BNT": "Hora de Brunei Darussalam", "BOT": "Hora de Bolívia", "BRST": "Hora d’estiu de Brasília", "BRT": "Hora estàndard de Brasília", "BST Bangladesh": "Hora estàndard de Bangladesh", "BT": "Hora de Bhutan", "CAST": "CAST", "CAT": "Hora de l’Àfrica central", "CCT": "Hora de les illes Cocos", "CDT": "Hora d’estiu central d’Amèrica del Nord", "CHADT": "Hora d’estiu de Chatham", "CHAST": "Hora estàndard de Chatham", "CHUT": "Hora de Chuuk", "CKT": "Hora estàndard de les illes Cook", "CKT DST": "Hora de mig estiu de les illes Cook", "CLST": "Hora d’estiu de Xile", "CLT": "Hora estàndard de Xile", "COST": "Hora d’estiu de Colòmbia", "COT": "Hora estàndard de Colòmbia", "CST": "Hora estàndard central d’Amèrica del Nord", "CST China": "Hora estàndard de la Xina", "CST China DST": "Hora d’estiu de la Xina", "CVST": "Hora d’estiu de Cap Verd", "CVT": "Hora estàndard de Cap Verd", "CXT": "Hora de Kiritimati", "ChST": "Hora estàndard de Chamorro", "ChST NMI": "ChST NMI", "CuDT": "Hora d’estiu de Cuba", "CuST": "Hora estàndard de Cuba", "DAVT": "Hora de Davis", "DDUT": "Hora de Dumont d’Urville", "EASST": "Hora d’estiu de l’illa de Pasqua", "EAST": "Hora estàndard de l’illa de Pasqua", "EAT": "Hora de l’Àfrica oriental", "ECT": "Hora de l’Equador", "EDT": "Hora d’estiu oriental d’Amèrica del Nord", "EGDT": "Hora d’estiu de l’Est de Groenlàndia", "EGST": "Hora estàndard de l’Est de Groenlàndia", "EST": "Hora estàndard oriental d’Amèrica del Nord", "FEET": "Hora de l’extrem oriental d’Europa", "FJT": "Hora estàndard de Fiji", "FJT Summer": "Hora d’estiu de Fiji", "FKST": "Hora d’estiu de les illes Falkland", "FKT": "Hora estàndard de les illes Falkland", "FNST": "Hora d’estiu de Fernando de Noronha", "FNT": "Hora estàndard de Fernando de Noronha", "GALT": "Hora de Galápagos", "GAMT": "Hora de Gambier", "GEST": "Hora d’estiu de Geòrgia", "GET": "Hora estàndard de Geòrgia", "GFT": "Hora de la Guaiana Francesa", "GIT": "Hora de les illes Gilbert", "GMT": "Hora del meridià de Greenwich", "GNSST": "Hora d’estiu de Groenlàndia", "GNST": "Hora estàndard de Groenlàndia", "GST": "Hora de Geòrgia del Sud", "GST Guam": "GST Guam", "GYT": "Hora de Guyana", "HADT": "Hora d’estiu de Hawaii-Aleutianes", "HAST": "Hora estàndard de Hawaii-Aleutianes", "HKST": "Hora d’estiu de Hong Kong", "HKT": "Hora estàndard de Hong Kong", "HOVST": "Hora d’estiu de Khovd", "HOVT": "Hora estàndard de Khovd", "ICT": "Hora de l’Indoxina", "IDT": "Hora d’estiu d’Israel", "IOT": "Hora de l’oceà Índic", "IRKST": "Hora d’estiu d’Irkutsk", "IRKT": "Hora estàndard d’Irkutsk", "IRST": "Hora estàndard de l’Iran", "IRST DST": "Hora d’estiu de l’Iran", "IST": "Hora de l’Índia", "IST Israel": "Hora estàndard d’Israel", "JDT": "Hora d’estiu del Japó", "JST": "Hora estàndard del Japó", "KOST": "Hora de Kosrae", "KRAST": "Hora d’estiu de Krasnoiarsk", "KRAT": "Hora estàndard de Krasnoiarsk", "KST": "Hora estàndard de Corea", "KST DST": "Hora d’estiu de Corea", "LHDT": "Horari d’estiu de Lord Howe", "LHST": "Hora estàndard de Lord Howe", "LINT": "Hora de les illes Line", "MAGST": "Hora d’estiu de Magadan", "MAGT": "Hora estàndard de Magadan", "MART": "Hora de les Marqueses", "MAWT": "Hora de Mawson", "MDT": "Hora d’estiu de muntanya d’Amèrica del Nord", "MESZ": "Hora d’estiu d’Europa central", "MEZ": "Hora estàndard d’Europa central", "MHT": "Hora de les illes Marshall", "MMT": "Hora de Myanmar", "MSD": "Hora d’estiu de Moscou", "MST": "Hora estàndard de muntanya d’Amèrica del Nord", "MUST": "Hora d’estiu de Maurici", "MUT": "Hora estàndard de Maurici", "MVT": "Hora de les Maldives", "MYT": "Hora de Malàisia", "NCT": "Hora estàndard de Nova Caledònia", "NDT": "Hora d’estiu de Terranova", "NDT New Caledonia": "Hora d’estiu de Nova Caledònia", "NFDT": "Hora d’estiu de l’illa Norfolk", "NFT": "Hora estàndard de l’illa Norfolk", "NOVST": "Hora d’estiu de Novossibirsk", "NOVT": "Hora estàndard de Novossibirsk", "NPT": "Hora del Nepal", "NRT": "Hora de Nauru", "NST": "Hora estàndard de Terranova", "NUT": "Hora de Niue", "NZDT": "Hora d’estiu de Nova Zelanda", "NZST": "Hora estàndard de Nova Zelanda", "OESZ": "Hora d’estiu d’Europa oriental", "OEZ": "Hora estàndard d’Europa oriental", "OMSST": "Hora d’estiu d’Omsk", "OMST": "Hora estàndard d’Omsk", "PDT": "Hora d’estiu del Pacífic d’Amèrica del Nord", "PDTM": "Hora d’estiu del Pacífic de Mèxic", "PETDT": "Horari d’estiu de Petropavlovsk de Kamtxatka", "PETST": "Hora estàndard de Petropavlovsk de Kamtxatka", "PGT": "Hora de Papua Nova Guinea", "PHOT": "Hora de les illes Phoenix", "PKT": "Hora estàndard del Pakistan", "PKT DST": "Hora d’estiu del Pakistan", "PMDT": "Hora d’estiu de Saint-Pierre-et-Miquelon", "PMST": "Hora estàndard de Saint-Pierre-et-Miquelon", "PONT": "Hora de Ponape", "PST": "Hora estàndard del Pacífic d’Amèrica del Nord", "PST Philippine": "Hora estàndard de les Filipines", "PST Philippine DST": "Hora d’estiu de les Filipines", "PST Pitcairn": "Hora de Pitcairn", "PSTM": "Hora estàndard del Pacífic de Mèxic", "PWT": "Hora de Palau", "PYST": "Hora d’estiu del Paraguai", "PYT": "Hora estàndard del Paraguai", "PYT Korea": "Hora de Pyongyang", "RET": "Hora de Reunió", "ROTT": "Hora de Rothera", "SAKST": "Hora d’estiu de Sakhalín", "SAKT": "Hora estàndard de Sakhalín", "SAMST": "Hora d’estiu de Samara", "SAMT": "Hora estàndard de Samara", "SAST": "Hora estàndard de l’Àfrica meridional", "SBT": "Hora de les illes Salomó", "SCT": "Hora de les Seychelles", "SGT": "Hora de Singapur", "SLST": "SLST", "SRT": "Hora de Surinam", "SST Samoa": "Hora estàndard de Samoa", "SST Samoa Apia": "Hora estàndard d’Apia", "SST Samoa Apia DST": "Hora d’estiu d’Apia", "SST Samoa DST": "Hora d’estiu de Samoa", "SYOT": "Hora de Syowa", "TAAF": "Hora d’Antàrtida i de les Terres Australs Antàrtiques Franceses", "TAHT": "Hora de Tahití", "TJT": "Hora del Tadjikistan", "TKT": "Hora de Tokelau", "TLT": "Hora de Timor Oriental", "TMST": "Hora d’estiu del Turkmenistan", "TMT": "Hora estàndard del Turkmenistan", "TOST": "Hora d’estiu de Tonga", "TOT": "Hora estàndard de Tonga", "TVT": "Hora de Tuvalu", "TWT": "Hora estàndard de Taipei", "TWT DST": "Hora d’estiu de Taipei", "ULAST": "Hora d’estiu d’Ulaanbaatar", "ULAT": "Hora estàndard d’Ulaanbaatar", "UYST": "Hora d’estiu de l’Uruguai", "UYT": "Hora estàndard de l’Uruguai", "UZT": "Hora estàndard de l’Uzbekistan", "UZT DST": "Hora d’estiu de l’Uzbekistan", "VET": "Hora de Veneçuela", "VLAST": "Hora d’estiu de Vladivostok", "VLAT": "Hora estàndard de Vladivostok", "VOLST": "Hora d’estiu de Volgograd", "VOLT": "Hora estàndard de Volgograd", "VOST": "Hora de Vostok", "VUT": "Hora estàndard de Vanatu", "VUT DST": "Hora d’estiu de Vanatu", "WAKT": "Hora de les illes Wake", "WARST": "Hora d’estiu de l’oest de l’Argentina", "WART": "Hora estàndard de l’oest de l’Argentina", "WAST": "Hora de l’Àfrica occidental", "WAT": "Hora de l’Àfrica occidental", "WESZ": "Hora d’estiu d’Europa occidental", "WEZ": "Hora estàndard d’Europa occidental", "WFT": "Hora de Wallis i Futuna", "WGST": "Hora d’estiu de l’Oest de Groenlàndia", "WGT": "Hora estàndard de l’Oest de Groenlàndia", "WIB": "Hora de l’oest d’Indonèsia", "WIT": "Hora de l’est d’Indonèsia", "WITA": "Hora central d’Indonèsia", "YAKST": "Hora d’estiu de Iakutsk", "YAKT": "Hora estàndard de Iakutsk", "YEKST": "Hora d’estiu de Iekaterinburg", "YEKT": "Hora estàndard de Iekaterinburg", "YST": "Hora de Yukon", "МСК": "Hora estàndard de Moscou", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Hora de l’oest del Kazakhstan", "شىعىش قازاق ەلى": "Hora de l’est del Kazakhstan", "قازاق ەلى": "Hora del Kazakhstan", "قىرعىزستان": "Hora del Kirguizstan", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Hora d’estiu de les Açores"},
	}
}

// Locale returns the current translators string locale
func (ca *ca_AD) Locale() string {
	return ca.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ca_AD'
func (ca *ca_AD) PluralsCardinal() []locales.PluralRule {
	return ca.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ca_AD'
func (ca *ca_AD) PluralsOrdinal() []locales.PluralRule {
	return ca.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ca_AD'
func (ca *ca_AD) PluralsRange() []locales.PluralRule {
	return ca.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ca_AD'
func (ca *ca_AD) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
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

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ca_AD'
func (ca *ca_AD) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)

	if n == 1 || n == 3 {
		return locales.PluralRuleOne
	} else if n == 2 {
		return locales.PluralRuleTwo
	} else if n == 4 {
		return locales.PluralRuleFew
	}

	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ca_AD'
func (ca *ca_AD) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ca *ca_AD) MonthAbbreviated(month time.Month) string {
	return ca.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ca *ca_AD) MonthsAbbreviated() []string {
	return ca.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ca *ca_AD) MonthNarrow(month time.Month) string {
	return ca.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ca *ca_AD) MonthsNarrow() []string {
	return ca.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ca *ca_AD) MonthWide(month time.Month) string {
	return ca.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ca *ca_AD) MonthsWide() []string {
	return ca.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ca *ca_AD) WeekdayAbbreviated(weekday time.Weekday) string {
	return ca.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ca *ca_AD) WeekdaysAbbreviated() []string {
	return ca.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ca *ca_AD) WeekdayNarrow(weekday time.Weekday) string {
	return ca.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ca *ca_AD) WeekdaysNarrow() []string {
	return ca.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ca *ca_AD) WeekdayShort(weekday time.Weekday) string {
	return ca.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ca *ca_AD) WeekdaysShort() []string {
	return ca.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ca *ca_AD) WeekdayWide(weekday time.Weekday) string {
	return ca.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ca *ca_AD) WeekdaysWide() []string {
	return ca.daysWide
}

// Decimal returns the decimal point of number
func (ca *ca_AD) Decimal() string {
	return ca.decimal
}

// Group returns the group of number
func (ca *ca_AD) Group() string {
	return ca.group
}

// Group returns the minus sign of number
func (ca *ca_AD) Minus() string {
	return ca.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ca_AD' and handles both Whole and Real numbers based on 'v'
func (ca *ca_AD) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ca.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ca.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ca.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ca_AD' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ca *ca_AD) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ca.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ca.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ca.percentSuffix...)

	b = append(b, ca.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ca_AD'
func (ca *ca_AD) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ca.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ca.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ca.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ca.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ca.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, ca.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ca_AD'
// in accounting notation.
func (ca *ca_AD) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ca.currencies[currency]
	l := len(s) + len(symbol) + 6 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ca.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ca.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ca.currencyNegativePrefix[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ca.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, ca.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, ca.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ca_AD'
func (ca *ca_AD) FmtDateShort(t time.Time) string {
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

// FmtDateMedium returns the medium date representation of 't' for 'ca_AD'
func (ca *ca_AD) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ca.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ca_AD'
func (ca *ca_AD) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ca.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65, 0x6c}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ca_AD'
func (ca *ca_AD) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = append(b, ca.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ca.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20, 0x64, 0x65, 0x6c}...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ca_AD'
func (ca *ca_AD) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ca.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ca_AD'
func (ca *ca_AD) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ca.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ca.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ca_AD'
func (ca *ca_AD) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ca.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ca.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ca_AD'
func (ca *ca_AD) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, ca.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ca.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0x28}...)

	tz, _ := t.Zone()

	if btz, ok := ca.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	b = append(b, []byte{0x29}...)

	return string(b)
}
