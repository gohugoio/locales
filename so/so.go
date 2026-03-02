package so

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type so struct {
	locale             string
	pluralsCardinal    []locales.PluralRule
	pluralsOrdinal     []locales.PluralRule
	pluralsRange       []locales.PluralRule
	decimal            string
	group              string
	minus              string
	percent            string
	perMille           string
	timeSeparator      string
	inifinity          string
	currencies         []string // idx = enum of currency code
	monthsAbbreviated  []string
	monthsNarrow       []string
	monthsWide         []string
	daysAbbreviated    []string
	daysNarrow         []string
	daysShort          []string
	daysWide           []string
	periodsAbbreviated []string
	periodsNarrow      []string
	periodsShort       []string
	periodsWide        []string
	erasAbbreviated    []string
	erasNarrow         []string
	erasWide           []string
	timezones          map[string]string
}

// New returns a new instance of translator for the 'so' locale
func New() locales.Translator {
	return &so{
		locale:             "so",
		pluralsCardinal:    []locales.PluralRule{2, 6},
		pluralsOrdinal:     nil,
		pluralsRange:       nil,
		timeSeparator:      ":",
		currencies:         []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "DBB", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "S", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		monthsAbbreviated:  []string{"", "Jan", "Feb", "Mar", "Abr", "May", "Jun", "Lul", "Ogs", "Seb", "Okt", "Nof", "Dis"},
		monthsNarrow:       []string{"", "J", "F", "M", "A", "M", "J", "L", "O", "S", "O", "N", "D"},
		monthsWide:         []string{"", "Janaayo", "Febraayo", "Maarso", "Abriil", "Maayo", "Juun", "Luulyo", "Agosto", "Sebtembar", "Oktoobar", "Noofeembar", "Diseembar"},
		daysAbbreviated:    []string{"Axd", "Isn", "Tldo", "Arbc", "Khms", "Jmc", "Sbti"},
		daysNarrow:         []string{"A", "I", "T", "A", "Kh", "J", "S"},
		daysWide:           []string{"Axad", "Isniin", "Talaado", "Arbaco", "Khamiis", "Jimco", "Sabti"},
		periodsAbbreviated: []string{"GH", "GD"},
		periodsNarrow:      []string{"h", "d"},
		erasAbbreviated:    []string{"BC", "AD"},
		erasNarrow:         []string{"B", "A"},
		erasWide:           []string{"Ciise Hortii", "Ciise Dabadii"},
		timezones:          map[string]string{"ACDT": "Waqtiga Dharaarta ee Bartamaha Astaraaliya", "ACST": "Wakhtiga Kulka ee Acre", "ACT": "Wakhtiga Caadiga ah ee Acre", "ACWDT": "Waqtiga Dharaarta Bartamaha Galbeedka Australiya", "ACWST": "Waqtiga Caadiga Ah ee Bartamaha Galbeedka Astaraaliya", "ADT": "Waqtiga Dharaarta ee Atlantika Waqooyiga Ameerika", "ADT Arabia": "Waqtiga Dharaarta ee Carabta", "AEDT": "Waqtiga Dharaarta ee Bariga Astaraaliya", "AEST": "Waqtiyada Caadiga ah ee Bariga Astaraaliya", "AFT": "Waqtiga Afggaanistaan", "AKDT": "Waqtiga Dharaarta ee Alaska", "AKST": "Waqtiga Caadiga Ah ee Alaska", "AMST": "Waqtiga Xagaaga ee Amason", "AMST Armenia": "Waqtiga Xagaaga ee Armeeniya", "AMT": "Waqtiga Caadiga Ah ee Amason", "AMT Armenia": "Waqtiga Caadiga Ah ee Armeeniya", "ANAST": "Wakhtiga Kulka ee Anadyr", "ANAT": "Wakhtiga Caadiga ah ee Anadyr", "ARST": "Waqtiga Xagaaga ee Arjentiina", "ART": "Waqtiga Caadiga Ah ee Arjentiina", "AST": "Waqtiga Caadiga Ah ee Atlantika Waqooyiga Ameerika", "AST Arabia": "Waqtiga Caadiga Ah ee Carabta", "AWDT": "Waqtiga Dharaarta ee Galbeedka Astaraaliya", "AWST": "Waqtiga Caadiga Ah ee Galbeedka Astaraaliya", "AZST": "Waqtiga Xagaaga ee Asarbeyjan", "AZT": "Waqtiga Caadiga Ah ee Asarbeyjan", "BDT Bangladesh": "Waqtiga Xagaaga ee Bangledeesh", "BNT": "Wakhtiga Brunei", "BOT": "Waqtiga Boliifiya", "BRST": "Waqtiga Xagaaga ee Baraasiliya", "BRT": "Waqtiga Caadiga ah ee Baraasiliya", "BST Bangladesh": "Waqtiga Caadiga Ah ee Bangledeesh", "BT": "Waqtiga Butaan", "CAST": "CAST", "CAT": "Waqtiga Bartamaha Afrika", "CCT": "Waqtiga Kokos Aylaan", "CDT": "Waqtiga Dharaarta ee Bartamaha Waqooyiga Ameerika", "CHADT": "Waqtiga Dharaarta ee Jaatam", "CHAST": "Waqtiga Caadiga Ah ee Jaatam", "CHUT": "Waqtiga Juuk", "CKT": "Waqtiga Caadiga Ah ee Jasiiradaha Cook", "CKT DST": "Waqtiga Caadiga Ah ee Jasiiradaha Cook", "CLST": "Waqtiga Xagaaga ee Jili", "CLT": "Waqtiga Caadiga Ah ee Jili", "COST": "Waqtiga Xagaaga ee Kolambiya", "COT": "Waqtiga Caadiga Ah ee Kolambiya", "CST": "Waqtiga Caadiga Ah ee Bartamaha Waqooyiga Ameerika", "CST China": "Waqtiga Caadiga Ah ee Shiinaha", "CST China DST": "Waqtiga Dharaarta ee Shiinaha", "CVST": "Waqtiga Xagaaga ee Keyb Faarde", "CVT": "Waqtiga Caadiga Ah ee Keyb Faarde", "CXT": "Waqtiga Kirismas Aylaan", "ChST": "Wakhtiga Caadiga ah ee Chamorro", "ChST NMI": "ChST NMI", "CuDT": "Waqtiga Dharaarta ee Kuuba", "CuST": "Waqtiga Caadiga Ah ee Kuuba", "DAVT": "Wakhtiga Davis", "DDUT": "Waqtiga Dumont - d’urfille", "EASST": "Waqtiga Xagaaga ee Iistar Aylaan", "EAST": "Waqtiga Caadiga Ah ee Iistar Aylaan", "EAT": "Waqtiga Bariga Afrika", "ECT": "Waqtiga Ekuwadoor", "EDT": "Waqtiga Dharaarta ee Bariga Waqooyiga Ameerika", "EGDT": "Waqtiga Xagaaga ee Bariga Giriinlaan", "EGST": "Waqtiga Caadiga ah ee Bariga Giriinlaan", "EST": "Waqtiga Caadiga Ah ee Bariga Waqooyiga Ameerika", "FEET": "Waqtiga Bariga Fog ee Yurub", "FJT": "Waqtiga Caadiga Ah ee Fiji", "FJT Summer": "Waqtiga Xagaaga ee Fiji", "FKST": "Waqtiga Xagaaga ee Faalklaan Aylaanis", "FKT": "Waqtiga Caadiga Ah ee Faalklaan Aylaanis", "FNST": "Waqtiga Xagaaga ee Farnaando de Nooronha", "FNT": "Waqtiga Caadiga Ah ee Farnaando de Nooronha", "GALT": "Waqtiga Galabagos", "GAMT": "Waqtiga Gambiyar", "GEST": "Waqtiga Xagaaga ee Joorjiya", "GET": "Waqtiga Caadiga Ah ee Joorjiya", "GFT": "Waqtiga Ferenj Guyana", "GIT": "Waqtiga Jilbeert Aylaan", "GMT": "Wakhtiga Giriinwij", "GNSST": "GNSST", "GNST": "GNST", "GST": "Waqtiga Sowt Joorjiya", "GST Guam": "GST Guam", "GYT": "Waqtiga Guyaana", "HADT": "Waqtiga Dharaarta ee Hawaay-Alutiyaan", "HAST": "Waqtiga Caadiga Ah ee Hawaay-Alutiyaan", "HKST": "Waqtiga Xagaaga ee Hoong Koong", "HKT": "Waqtiga Caadiga Ah ee Hoong Koong", "HOVST": "Waqtiga Xagaaga ee Hofud", "HOVT": "Waqtiga Caadiga Ah ee Hofud", "ICT": "Wakhtiga Indochina", "IDT": "Waqtiga Dharaarta ee Israaiil", "IOT": "Waqtiga Badweynta Hindiya", "IRKST": "Waqtiga Xagaaga ee Irkutsik", "IRKT": "Waqtiga Caadiga Ah ee Irkutsik", "IRST": "Waqtiga Caadiga Ah ee Iiraan", "IRST DST": "Waqtiga Dharaarta ee Iiraan", "IST": "Waqtiga Caadiga Ah ee Hindiya", "IST Israel": "Waqtiga Caadiga Ah ee Israaiil", "JDT": "Waqtiga Dharaarta ee Jabaan", "JST": "Waqtiga Caadiga Ah ee Jabaan", "KOST": "Waqtiga Kosriy", "KRAST": "Waqtiga Xagaaga ee Karasnoyarsik", "KRAT": "Waqtiga Caadiga Ah ee Karasnoyarsik", "KST": "Waqtiga Caadiga Ah ee Kuuriya", "KST DST": "Waqtiga Dharaarta ee Kuuriya", "LHDT": "Waqtiga Dharaarta ee Lod How", "LHST": "Waqtiga Caadiga Ah ee Lod How", "LINT": "Waqtiga Leyn Aylaan", "MAGST": "Waqtiga Xagaaga ee Magedan", "MAGT": "Waqtiga Caadiga Ah ee Magedan", "MART": "Waqtiga Marquwesas", "MAWT": "Waqtiga Mawson", "MDT": "Waqtiga Dharaarta ee Buurleyda Waqooyiga Ameerika", "MESZ": "Waqtiga Xagaaga ee Bartamaha Yurub", "MEZ": "Waqtiga Caadiga Ah ee Bartamaha Yurub", "MHT": "Waqtiga Maarshaal Aylaanis", "MMT": "Waqtiga Mayanmaar", "MSD": "Waqtiga Xagaaga ee Moskow", "MST": "Waqtiga Caadiga ah ee Buuraleyda Waqooyiga Ameerika", "MUST": "Waqtiga Xagaaga ee Morishiyaas", "MUT": "Waqtiga Caadiga Ah ee Morishiyaas", "MVT": "Waqtiga Maldifis", "MYT": "Waqtiga Maleyshiya", "NCT": "Waqtiga Caadiga Ah ee Niyuu Kaledoniya", "NDT": "Waqtiga Dharaarta ee Niyuufoonlaan", "NDT New Caledonia": "Waqtiga Xagaaga ee Niyuu Kaledoniya", "NFDT": "Waqtiga Maalinta ee Norfolk Island", "NFT": "Waqtiga Caadiga ah ee Norfolk Island", "NOVST": "Waqtiga Xagaaga ee Nofosibirsik", "NOVT": "Waqtiga Caadiga Ah ee Nofosibirsik", "NPT": "Waqtiga Neebaal", "NRT": "Waqtiga Nawroo", "NST": "Waqtiga Caadiga Ah ee Niyuufoonlaan", "NUT": "Waqtiga Niyuu", "NZDT": "Waqtiga Dharaarta ee Niyuu Si’laan", "NZST": "Waqtiga Caadiga Ah ee Niyuu Si’laan", "OESZ": "Waqtiga Xagaaga ee Bariga Yurub", "OEZ": "Waqtiga Caadiga Ah ee Bariga Yurub", "OMSST": "Waqtiga Xagaaga ee Omsk", "OMST": "Waqtiga Caadiga Ah ee Omsk", "PDT": "Waqtiga Dharaarta ee Basifika Waqooyiga Ameerika", "PDTM": "Waqtiga Dharaarta ee Baasifikada Meksiko", "PETDT": "Wakhtiga Kulka ee Petropavlovsk-Kamchatski", "PETST": "Wakhtiga Caadiga ah ee Petropavlovsk-Kamchatski", "PGT": "Waqtiga Babuw Niyuu Giniya", "PHOT": "Waqtiga Foonikis Aylaanis", "PKT": "Waqtiga Caadiga Ah ee Bakistaan", "PKT DST": "Waqtiga Xagaaga ee Bakistaan", "PMDT": "Waqtiga Dharaarta ee St. Beere & Mikiwelon", "PMST": "Waqtiga Caadiga Ah St. Beere & Mikiwelon", "PONT": "Wakhtiga Pohnpei", "PST": "Waqtiga Caadiga ah ee Basifika Waqooyiga Ameerika", "PST Philippine": "Waqtiga Caadiga Ah ee Filibiin", "PST Philippine DST": "Waqtiga Xagaaga ee Filibiin", "PST Pitcairn": "Waqtiga Bitkeen", "PSTM": "Waqtiga Caadiga Ah ee Baasifikada Meksiko", "PWT": "Waqtiga Balaw", "PYST": "Waqtiga Xagaaga ee Baragwaay", "PYT": "Waqtiga Caadiga Ah ee Baragwaay", "PYT Korea": "Waqtiga Boyongyang", "RET": "Waqtiga Riyuuniyon", "ROTT": "Waqtiga Rotera", "SAKST": "Waqtiga Xagaaga ee Sakhalin", "SAKT": "Waqtiga Caadiga Ah ee Sakhalin", "SAMST": "Wakhtiga Kulka ee Samara", "SAMT": "Wakhtiga Caadiga ah ee Samara", "SAST": "Waqtiyada Caadiga Ah ee Koonfur Afrika", "SBT": "Waqtiga Solomon Aylaanis", "SCT": "Waqtiga Siishalis", "SGT": "Waqtiga Singabuur", "SLST": "SLST", "SRT": "Waqtiga Surineym", "SST Samoa": "Waqtiga Caadiga Ah ee American Samoa", "SST Samoa Apia": "Waqtiga Caadiga Ah ee Samoa", "SST Samoa Apia DST": "Waqtiga Dharaarta ee Samoa", "SST Samoa DST": "Waqtiga Dharaarta ee American Samoa", "SYOT": "Waqtiga Siyowa", "TAAF": "Waqtiga Koonfurta Faransiiska & Antaarktik", "TAHT": "Waqtiga Tahiti", "TJT": "Waqtiga Tajikistan", "TKT": "Waqtiga Tokeluu", "TLT": "Wakhtiga Timor-Leste", "TMST": "Waqtiga Xagaaga ee Turkmenistan", "TMT": "Waqtiga Caadiga Ah ee Turkmenistan", "TOST": "Waqtiga Xagaaga ee Tonga", "TOT": "Waqtiga Caadiga Ah ee Tonga", "TVT": "Waqtiga Tufalu", "TWT": "Waqtiga Caadiga Ah ee Taiwan", "TWT DST": "Waqtiga Dharaarta ee Teybey", "ULAST": "Waqtiga Xagaaga ee Ulaanbaataar", "ULAT": "Waqtiga Caadiga Ah ee Ulaanbaataar", "UYST": "Waqtiga Xagaaga ee Urugwaay", "UYT": "Waqtiga Caadiga Ah ee Urugwaay", "UZT": "Waqtiga Caadiga Ah ee Usbekistan", "UZT DST": "Waqtiga Xagaaga ee Usbekistan", "VET": "Waqtiga Fenezuweela", "VLAST": "Waqtiga Xagaaga ee Faladifostok", "VLAT": "Waqtiga Caadiga Ah ee Faladifostok", "VOLST": "Waqtiga Xagaaga ee Folgograd", "VOLT": "Waqtiga Caadiga Ah ee Folgograd", "VOST": "Waqtiga Fostok", "VUT": "Waqtiga Caadiga Ah ee Fanuutu", "VUT DST": "Waqtiga Xagaaga ee Fanuutu", "WAKT": "Waqtiga Wayk Iylaanis", "WARST": "Waqtiga Xagaaga ee Galbeedka Arjentiina", "WART": "Waqtiga Caadiga Ah ee Galbeedka Arjentiina", "WAST": "Waqtiga Galbeedka Afrika", "WAT": "Waqtiga Galbeedka Afrika", "WESZ": "Waqtiga Xagaaga ee Galbeedka Yurub", "WEZ": "Waqtiga Caadiga Ah ee Galbeedka Yurub", "WFT": "Waqtiga Walis & Futuna", "WGST": "Waqtiga Xagaaga ee Galbeedka Giriinlaan", "WGT": "Waqtiga Caadiga Ah ee Galbeedka Giriinlaan", "WIB": "Waqtiga Galbeedka Indoneeysiya", "WIT": "Waqtiga Indoneeysiya", "WITA": "Waqtiga Bartamaha Indoneeysiya", "YAKST": "Waqtiga Xagaaga ee Yakut", "YAKT": "Waqtiga Caadiga Ah ee Yakut", "YEKST": "Waqtiga Xagaaga ee Yekaterinbaag", "YEKT": "Waqtiga Caadiga Ah ee Yekaterinbaag", "YST": "Waqtiga Yukon", "МСК": "Waqtiga Caadiga Ah ee Moskow", "اقتاۋ": "Waqtiga Caadiga ah ee Aqtau", "اقتاۋ قالاسى": "Saacada Waqtiga Kulaylaha Aqtau", "اقتوبە": "Waqtiga Caadiga ah ee Aqtobe", "اقتوبە قالاسى": "Saacada Waqtiga kulaylaha Aqtobe", "الماتى": "Waqtiga Caadiga ah ee Almaty", "الماتى قالاسى": "Saacada Waqtiga Kulaylaha ee Almaty", "باتىس قازاق ەلى": "Waqtiga Koonfurta Kasakhistan", "شىعىش قازاق ەلى": "Waqtiga Bariga Kasakhistaan", "قازاق ەلى": "Wakhtiga Kazakhistan", "قىرعىزستان": "Waqtiga Kiyrigistaan", "قىزىلوردا": "Waqtiga Caadiga ah ee Qyzylorda", "قىزىلوردا قالاسى": "Saacada Waqtiga Kulaylaha Qyzylorda", "∅∅∅": "Waqtiga Xagaaga ee Beeru"},
	}
}

// Locale returns the current translators string locale
func (so *so) Locale() string {
	return so.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'so'
func (so *so) PluralsCardinal() []locales.PluralRule {
	return so.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'so'
func (so *so) PluralsOrdinal() []locales.PluralRule {
	return so.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'so'
func (so *so) PluralsRange() []locales.PluralRule {
	return so.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'so'
func (so *so) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'so'
func (so *so) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'so'
func (so *so) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (so *so) MonthAbbreviated(month time.Month) string {
	return so.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (so *so) MonthsAbbreviated() []string {
	return so.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (so *so) MonthNarrow(month time.Month) string {
	return so.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (so *so) MonthsNarrow() []string {
	return so.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (so *so) MonthWide(month time.Month) string {
	return so.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (so *so) MonthsWide() []string {
	return so.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (so *so) WeekdayAbbreviated(weekday time.Weekday) string {
	return so.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (so *so) WeekdaysAbbreviated() []string {
	return so.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (so *so) WeekdayNarrow(weekday time.Weekday) string {
	return so.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (so *so) WeekdaysNarrow() []string {
	return so.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (so *so) WeekdayShort(weekday time.Weekday) string {
	return so.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (so *so) WeekdaysShort() []string {
	return so.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (so *so) WeekdayWide(weekday time.Weekday) string {
	return so.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (so *so) WeekdaysWide() []string {
	return so.daysWide
}

// Decimal returns the decimal point of number
func (so *so) Decimal() string {
	return so.decimal
}

// Group returns the group of number
func (so *so) Group() string {
	return so.group
}

// Group returns the minus sign of number
func (so *so) Minus() string {
	return so.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'so' and handles both Whole and Real numbers based on 'v'
func (so *so) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'so' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (so *so) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'so'
func (so *so) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := so.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'so'
// in accounting notation.
func (so *so) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := so.currencies[currency]
	return string(append(append([]byte{}, symbol...), s...))
}

// FmtDateShort returns the short date representation of 't' for 'so'
func (so *so) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'so'
func (so *so) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)
	b = append(b, so.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2d}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'so'
func (so *so) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, so.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'so'
func (so *so) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, so.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, so.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'so'
func (so *so) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, so.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, so.periodsAbbreviated[0]...)
	} else {
		b = append(b, so.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'so'
func (so *so) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, so.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, so.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, so.periodsAbbreviated[0]...)
	} else {
		b = append(b, so.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'so'
func (so *so) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, so.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, so.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, so.periodsAbbreviated[0]...)
	} else {
		b = append(b, so.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'so'
func (so *so) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, so.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, so.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, so.periodsAbbreviated[0]...)
	} else {
		b = append(b, so.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := so.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
