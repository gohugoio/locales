package to_TO

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type to_TO struct {
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
	currencyPositivePrefix string
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
	periodsAbbreviated     []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'to_TO' locale
func New() locales.Translator {
	return &to_TO{
		locale:                 "to_TO",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD$", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "F$", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD$", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "S$", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "T$", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		percentSuffix:          " ",
		currencyPositivePrefix: " ",
		currencyPositiveSuffix: "a",
		currencyNegativePrefix: " ",
		currencyNegativeSuffix: "a",
		monthsAbbreviated:      []string{"", "Sān", "Fēp", "Maʻa", "ʻEpe", "Mē", "Sun", "Siu", "ʻAok", "Sēp", "ʻOka", "Nōv", "Tīs"},
		monthsNarrow:           []string{"", "S", "F", "M", "ʻE", "M", "S", "S", "ʻA", "S", "ʻO", "N", "T"},
		monthsWide:             []string{"", "Sānuali", "Fēpueli", "Maʻasi", "ʻEpeleli", "Mē", "Sune", "Siulai", "ʻAokosi", "Sēpitema", "ʻOkatopa", "Nōvema", "Tīsema"},
		daysAbbreviated:        []string{"Sāp", "Mōn", "Tūs", "Pul", "Tuʻa", "Fal", "Tok"},
		daysNarrow:             []string{"S", "M", "T", "P", "T", "F", "T"},
		daysWide:               []string{"Sāpate", "Mōnite", "Tūsite", "Pulelulu", "Tuʻapulelulu", "Falaite", "Tokonaki"},
		periodsAbbreviated:     []string{"HH", "EA"},
		timezones:              map[string]string{"ACDT": "houa fakaʻaositelēlia-loto taimi liliu", "ACST": "houa fakaʻakelī taimi liliu", "ACT": "houa fakaʻakelī taimi totonu", "ACWDT": "houa fakaʻaositelēlia-loto-hihifo taimi liliu", "ACWST": "houa fakaʻaositelēlia-loto-hihifo taimi totonu", "ADT": "houa fakaʻamelika-tokelau ʻatalanitiki taimi liliu", "ADT Arabia": "houa fakaʻalepea taimi liliu", "AEDT": "houa fakaʻaositelēlia-hahake taimi liliu", "AEST": "houa fakaʻaositelēlia-hahake taimi totonu", "AFT": "houa fakaʻafikānisitani", "AKDT": "houa fakaʻalasika taimi liliu", "AKST": "houa fakaʻalasika taimi totonu", "AMST": "houa fakaʻamasōne taimi liliu", "AMST Armenia": "houa fakaʻāmenia taimi liliu", "AMT": "houa fakaʻamasōne taimi totonu", "AMT Armenia": "houa fakaʻāmenia taimi totonu", "ANAST": "houa fakalūsia-ʻanatili taimi liliu", "ANAT": "houa fakalūsia-ʻanatili taimi totonu", "ARST": "houa fakaʻasenitina taimi liliu", "ART": "houa fakaʻasenitina taimi totonu", "AST": "houa fakaʻamelika-tokelau ʻatalanitiki taimi totonu", "AST Arabia": "houa fakaʻalepea taimi totonu", "AWDT": "houa fakaʻaositelēlia-hihifo taimi liliu", "AWST": "houa fakaʻaositelēlia-hihifo taimi totonu", "AZST": "houa fakaʻasapaisani taimi liliu", "AZT": "houa fakaʻasapaisani taimi totonu", "BDT Bangladesh": "houa fakapāngilātesi taimi liliu", "BNT": "houa fakapulunei", "BOT": "houa fakapolīvia", "BRST": "houa fakapalāsila taimi liliu", "BRT": "houa fakapalāsila taimi totonu", "BST Bangladesh": "houa fakapāngilātesi taimi totonu", "BT": "houa fakapūtani", "CAST": "houa fakakeesi", "CAT": "houa fakaʻafelika-loto", "CCT": "houa fakamotukokosi", "CDT": "houa fakaʻamelika-tokelau loto taimi liliu", "CHADT": "houa fakasatihami taimi liliu", "CHAST": "houa fakasatihami taimi totonu", "CHUT": "houa fakatūke", "CKT": "houa fakaʻotumotukuki taimi totonu", "CKT DST": "houa fakaʻotumotukuki taimi liliu", "CLST": "houa fakasili taimi liliu", "CLT": "houa fakasili taimi totonu", "COST": "houa fakakolomipia taimi liliu", "COT": "houa fakakolomipia taimi totonu", "CST": "houa fakaʻamelika-tokelau loto taimi totonu", "CST China": "houa fakasiaina taimi totonu", "CST China DST": "houa fakasiaina taimi liliu", "CVST": "houa fakamuiʻi-vēte taimi liliu", "CVT": "houa fakamuiʻi-vēte taimi totonu", "CXT": "houa fakamotukilisimasi", "ChST": "houa fakakamolo", "ChST NMI": "houa fakamalianatokelau", "CuDT": "houa fakakiupa taimi liliu", "CuST": "houa fakakiupa taimi totonu", "DAVT": "houa fakatavisi", "DDUT": "houa fakatūmoni-tūvile", "EASST": "houa fakalapanui taimi liliu", "EAST": "houa fakalapanui taimi totonu", "EAT": "houa fakaʻafelika-hahake", "ECT": "houa fakaʻekuetoa", "EDT": "houa fakaʻamelika-tokelau hahake taimi liliu", "EGDT": "houa fakafonuamata-hahake taimi liliu", "EGST": "houa fakafonuamata-hahake taimi totonu", "EST": "houa fakaʻamelika-tokelau hahake taimi totonu", "FEET": "houa fakaʻeulope-hahake-ange", "FJT": "houa fakafisi taimi totonu", "FJT Summer": "houa fakafisi taimi liliu", "FKST": "houa fakaʻotumotu-fokulani taimi liliu", "FKT": "houa fakaʻotumotu-fokulani taimi totonu", "FNST": "houa fakafēnanito-te-nolōnia taimi liliu", "FNT": "houa fakafēnanito-te-nolōnia taimi totonu", "GALT": "houa fakakalapakosi", "GAMT": "houa fakakamipiē", "GEST": "houa fakaseōsia taimi liliu", "GET": "houa fakaseōsia taimi totonu", "GFT": "houa fakakuiana-fakafalanisē", "GIT": "houa fakakilipasi", "GMT": "houa fakakiliniuisi mālie", "GNSST": "GNSST", "GNST": "GNST", "GST": "houa fakasiosiatonga", "GST Guam": "houa fakakuami", "GYT": "houa fakakuiana", "HADT": "houa fakahauaiʻi-aleuti taimi liliu", "HAST": "houa fakahauaiʻi-aleuti taimi totonu", "HKST": "houa fakahongi-kongi taimi liliu", "HKT": "houa fakahongi-kongi taimi totonu", "HOVST": "houa fakahovite taimi liliu", "HOVT": "houa fakahovite taimi totonu", "ICT": "houa fakaʻinitosiaina", "IDT": "houa fakaʻisileli taimi liliu", "IOT": "houa fakamoanaʻinitia", "IRKST": "houa fakalūsia-ʻīkutisiki taimi liliu", "IRKT": "houa fakalūsia-ʻīkutisiki taimi totonu", "IRST": "houa fakaʻilaani taimi totonu", "IRST DST": "houa fakaʻilaani taimi liliu", "IST": "houa fakaʻinitia", "IST Israel": "houa fakaʻisileli taimi totonu", "JDT": "houa fakasiapani taimi liliu", "JST": "houa fakasiapani taimi totonu", "KOST": "houa fakakosilae", "KRAST": "houa fakalūsia-kalasinoiāsiki taimi liliu", "KRAT": "houa fakalūsia-kalasinoiāsiki taimi totonu", "KST": "houa fakakōlea taimi totonu", "KST DST": "houa fakakōlea taimi liliu", "LHDT": "houa fakamotuʻeikihoue taimi liliu", "LHST": "houa fakamotuʻeikihoue taimi totonu", "LINT": "houa fakaʻotumotulaine", "MAGST": "houa fakalūsia-makatani taimi liliu", "MAGT": "houa fakalūsia-makatani taimi totonu", "MART": "houa fakamākesasi", "MAWT": "houa fakamausoni", "MDT": "houa fakaʻamelika-tokelau moʻunga taimi liliu", "MESZ": "houa fakaʻeulope-loto taimi liliu", "MEZ": "houa fakaʻeulope-loto taimi totonu", "MHT": "houa fakaʻotumotumasolo", "MMT": "houa fakapema", "MSD": "houa fakalūsia-mosikou taimi liliu", "MST": "houa fakaʻamelika-tokelau moʻunga taimi totonu", "MUST": "houa fakamaulitiusi taimi liliu", "MUT": "houa fakamaulitiusi taimi totonu", "MVT": "houa fakamalativisi", "MYT": "houa fakamaleisia", "NCT": "houa fakakaletōniafoʻou taimi totonu", "NDT": "houa fakafonuaʻilofoʻou taimi liliu", "NDT New Caledonia": "houa fakakaletōniafoʻou taimi liliu", "NFDT": "houa fakanoafōki taimi liliu", "NFT": "houa fakanoafōki taimi totonu", "NOVST": "houa fakalūsia-novosipīsiki taimi liliu", "NOVT": "houa fakalūsia-novosipīsiki taimi totonu", "NPT": "houa fakanepali", "NRT": "houa fakanaulu", "NST": "houa fakafonuaʻilofoʻou taimi totonu", "NUT": "houa fakaniuē", "NZDT": "houa fakanuʻusila taimi liliu", "NZST": "houa fakanuʻusila taimi totonu", "OESZ": "houa fakaʻeulope-hahake taimi liliu", "OEZ": "houa fakaʻeulope-hahake taimi totonu", "OMSST": "houa fakalūsia-ʻomisiki taimi liliu", "OMST": "houa fakalūsia-ʻomisiki taimi totonu", "PDT": "houa fakaʻamelika-tokelau pasifika taimi liliu", "PDTM": "houa fakamekisikou-pasifika taimi liliu", "PETDT": "houa fakalūsia-petelopavilovisiki taimi liliu", "PETST": "houa fakalūsia-petelopavilovisiki taimi totonu", "PGT": "houa fakapapuaniukini", "PHOT": "houa fakaʻotumotufoinikisi", "PKT": "houa fakapākisitani taimi totonu", "PKT DST": "houa fakapākisitani taimi liliu", "PMDT": "houa fakasā-piea-mo-mikeloni taimi liliu", "PMST": "houa fakasā-piea-mo-mikeloni taimi totonu", "PONT": "houa fakapōnapē", "PST": "houa fakaʻamelika-tokelau pasifika taimi totonu", "PST Philippine": "houa fakafilipaine taimi totonu", "PST Philippine DST": "houa fakafilipaine taimi liliu", "PST Pitcairn": "houa fakapitikani", "PSTM": "houa fakamekisikou-pasifika taimi totonu", "PWT": "houa fakapalau", "PYST": "houa fakapalakuai taimi liliu", "PYT": "houa fakapalakuai taimi totonu", "PYT Korea": "houa fakapiongiangi", "RET": "houa fakalēunioni", "ROTT": "houa fakalotela", "SAKST": "houa fakalūsia-sakāline taimi liliu", "SAKT": "houa fakalūsia-sakāline taimi totonu", "SAMST": "houa fakalūsia-samala taimi liliu", "SAMT": "houa fakalūsia-samala taimi totonu", "SAST": "houa fakaʻafelika-tonga", "SBT": "houa fakaʻotumotusolomone", "SCT": "houa fakaʻotumotu-seiseli", "SGT": "houa fakasingapoa", "SLST": "houa fakalangikā", "SRT": "houa fakasuliname", "SST Samoa": "houa fakahaʻamoa taimi totonu", "SST Samoa Apia": "houa fakaapia taimi totonu", "SST Samoa Apia DST": "houa fakaapia taimi liliu", "SST Samoa DST": "houa fakahaʻamoa taimi liliu", "SYOT": "houa fakasioua", "TAAF": "houa fakaʻanetātikafalanisē", "TAHT": "houa fakatahisi", "TJT": "houa fakatasikitani", "TKT": "houa fakatokelau", "TLT": "houa fakatimoa-hahake", "TMST": "houa fakatūkimenisitani taimi liliu", "TMT": "houa fakatūkimenisitani taimi totonu", "TOST": "houa fakatonga taimi liliu", "TOT": "houa fakatonga taimi totonu", "TVT": "houa fakatūvalu", "TWT": "houa fakataipei taimi totonu", "TWT DST": "houa fakataipei taimi liliu", "ULAST": "houa fakaʻulānipātā taimi liliu", "ULAT": "houa fakaʻulānipātā taimi totonu", "UYST": "houa fakaʻulukuai taimi liliu", "UYT": "houa fakaʻulukuai taimi totonu", "UZT": "houa fakaʻusipekitani taimi totonu", "UZT DST": "houa fakaʻusipekitani taimi liliu", "VET": "houa fakavenesuela", "VLAST": "houa fakalūsia-valativositoki taimi liliu", "VLAT": "houa fakalūsia-valativositoki taimi totonu", "VOLST": "houa fakalūsia-volikokalati taimi liliu", "VOLT": "houa fakalūsia-volikokalati taimi totonu", "VOST": "houa fakavositoki", "VUT": "houa fakavanuatu taimi totonu", "VUT DST": "houa fakavanuatu taimi liliu", "WAKT": "houa fakamotuueke", "WARST": "houa fakaʻasenitina-hihifo taimi liliu", "WART": "houa fakaʻasenitina-hihifo taimi totonu", "WAST": "houa fakaʻafelika-hihifo", "WAT": "houa fakaʻafelika-hihifo", "WESZ": "houa fakaʻeulope-hihifo taimi liliu", "WEZ": "houa fakaʻeulope-hihifo taimi totonu", "WFT": "houa fakaʻuvea mo futuna", "WGST": "houa fakafonuamata-hihifo taimi liliu", "WGT": "houa fakafonuamata-hihifo taimi totonu", "WIB": "houa fakaʻinitonisia-hihifo", "WIT": "houa fakaʻinitonisia-hahake", "WITA": "houa fakaʻinitonisia-loto", "YAKST": "houa fakalūsia-ʻiākutisiki taimi liliu", "YAKT": "houa fakalūsia-ʻiākutisiki taimi totonu", "YEKST": "houa fakalūsia-ʻiekatelinepūki taimi liliu", "YEKT": "houa fakalūsia-ʻiekatelinepūki taimi totonu", "YST": "houa fakaiukoni", "МСК": "houa fakalūsia-mosikou taimi totonu", "اقتاۋ": "houa fakaʻakitau taimi totonu", "اقتاۋ قالاسى": "houa fakaʻakitau taimi liliu", "اقتوبە": "houa fakaʻakitōpe taimi totonu", "اقتوبە قالاسى": "houa fakaʻakitōpe taimi liliu", "الماتى": "houa fakaʻalamati taimi totonu", "الماتى قالاسى": "houa fakaʻalamati taimi liliu", "باتىس قازاق ەلى": "houa fakakasakitani-hihifo", "شىعىش قازاق ەلى": "houa fakakasakitani-hahake", "قازاق ەلى": "houa fakakasakitani", "قىرعىزستان": "houa fakakīkisitani", "قىزىلوردا": "houa fakakisilōta taimi totonu", "قىزىلوردا قالاسى": "houa fakakisilōta taimi liliu", "∅∅∅": "houa fakapelū taimi liliu"},
	}
}

// Locale returns the current translators string locale
func (to *to_TO) Locale() string {
	return to.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'to_TO'
func (to *to_TO) PluralsCardinal() []locales.PluralRule {
	return to.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'to_TO'
func (to *to_TO) PluralsOrdinal() []locales.PluralRule {
	return to.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'to_TO'
func (to *to_TO) PluralsRange() []locales.PluralRule {
	return to.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'to_TO'
func (to *to_TO) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'to_TO'
func (to *to_TO) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'to_TO'
func (to *to_TO) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (to *to_TO) MonthAbbreviated(month time.Month) string {
	if len(to.monthsAbbreviated) == 0 {
		return ""
	}
	return to.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (to *to_TO) MonthsAbbreviated() []string {
	return to.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (to *to_TO) MonthNarrow(month time.Month) string {
	if len(to.monthsNarrow) == 0 {
		return ""
	}
	return to.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (to *to_TO) MonthsNarrow() []string {
	return to.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (to *to_TO) MonthWide(month time.Month) string {
	if len(to.monthsWide) == 0 {
		return ""
	}
	return to.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (to *to_TO) MonthsWide() []string {
	return to.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (to *to_TO) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(to.daysAbbreviated) == 0 {
		return ""
	}
	return to.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (to *to_TO) WeekdaysAbbreviated() []string {
	return to.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (to *to_TO) WeekdayNarrow(weekday time.Weekday) string {
	if len(to.daysNarrow) == 0 {
		return ""
	}
	return to.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (to *to_TO) WeekdaysNarrow() []string {
	return to.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (to *to_TO) WeekdayShort(weekday time.Weekday) string {
	if len(to.daysShort) == 0 {
		return ""
	}
	return to.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (to *to_TO) WeekdaysShort() []string {
	return to.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (to *to_TO) WeekdayWide(weekday time.Weekday) string {
	if len(to.daysWide) == 0 {
		return ""
	}
	return to.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (to *to_TO) WeekdaysWide() []string {
	return to.daysWide
}

// Decimal returns the decimal point of number
func (to *to_TO) Decimal() string {
	return to.decimal
}

// Group returns the group of number
func (to *to_TO) Group() string {
	return to.group
}

// Group returns the minus sign of number
func (to *to_TO) Minus() string {
	return to.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'to_TO' and handles both Whole and Real numbers based on 'v'
func (to *to_TO) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, to.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, to.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, to.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'to_TO' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (to *to_TO) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 5
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, to.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, to.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, to.percentSuffix...)

	b = append(b, to.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'to_TO'
func (to *to_TO) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := to.currencies[currency]
	l := len(s) + len(symbol) + 5

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, to.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	for j := len(to.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, to.currencyPositivePrefix[j])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		b = append(b, to.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, to.currencyPositiveSuffix...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'to_TO'
// in accounting notation.
func (to *to_TO) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := to.currencies[currency]
	l := len(s) + len(symbol) + 5

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, to.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(to.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, to.currencyNegativePrefix[j])
		}

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		b = append(b, to.minus[0])

	} else {

		for j := len(to.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, to.currencyPositivePrefix[j])
		}

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if num < 0 {
		b = append(b, to.currencyNegativeSuffix...)
	} else {

		b = append(b, to.currencyPositiveSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'to_TO'
func (to *to_TO) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'to_TO'
func (to *to_TO) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, to.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'to_TO'
func (to *to_TO) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, to.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'to_TO'
func (to *to_TO) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, to.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, to.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'to_TO'
func (to *to_TO) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, to.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, to.periodsAbbreviated[0]...)
	} else {
		b = append(b, to.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'to_TO'
func (to *to_TO) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, to.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, to.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, to.periodsAbbreviated[0]...)
	} else {
		b = append(b, to.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'to_TO'
func (to *to_TO) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, to.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, to.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, to.periodsAbbreviated[0]...)
	} else {
		b = append(b, to.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'to_TO'
func (to *to_TO) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, to.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, to.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, to.periodsAbbreviated[0]...)
	} else {
		b = append(b, to.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := to.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
