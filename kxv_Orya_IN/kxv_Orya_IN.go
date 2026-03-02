package kxv_Orya_IN

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type kxv_Orya_IN struct {
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
	currencyPositivePrefix string
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
	periodsNarrow          []string
	periodsShort           []string
	periodsWide            []string
	erasAbbreviated        []string
	erasNarrow             []string
	erasWide               []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'kxv_Orya_IN' locale
func New() locales.Translator {
	return &kxv_Orya_IN{
		locale:                 "kxv_Orya_IN",
		pluralsCardinal:        nil,
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: "( ",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "pusu", "maha", "pagu", "hire", "bese", "jaṭṭa", "aasaḍi", "srabĩ", "bado", "dasara", "divi", "pande"},
		monthsNarrow:           []string{"", "pu", "ma", "pa", "hi", "be", "ja", "aa", "sra", "b", "da", "di", "pa"},
		monthsWide:             []string{"", "pusu lenju", "maha lenju", "pagu lenju", "hire lenju", "bese lenju", "jaṭṭa lenju", "aasaḍi lenju", "srabĩ lenju", "bado lenju", "dasara lenju", "divi lenju", "pande lenju"},
		daysAbbreviated:        []string{"aadi", "smba", "manga", "puda", "laki", "sukru", "sani"},
		daysNarrow:             []string{"aa", "s", "ma", "pu", "la", "su", "sa"},
		daysShort:              []string{"aa", "s", "ma", "pu", "laki", "su", "sa"},
		daysWide:               []string{"aadi vara", "smbara", "mangaḍa", "pudara", "laki vara", "sukru vara", "sani vara"},
		periodsAbbreviated:     []string{"am", "pm"},
		periodsNarrow:          []string{"a", "p"},
		erasAbbreviated:        []string{"bc", "ad"},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"krisṭ purb nki", "krisṭabd"},
		timezones:              map[string]string{"ACDT": "āstreliāti mādinā delāiṭ belā", "ACST": "āstreliāti mādinā mānānka belā", "ACT": "ACT", "ACWDT": "āstreliāti mādinā weḍākuṇupu delāiṭ belā", "ACWST": "āstreliāti mādinā weḍākuṇupu mānānka belā", "ADT": "āṭlānṭik delāiṭ belā", "ADT Arabia": "ārbiāti delāit belā", "AEDT": "āstreliāti weḍāhapū delāiti belā", "AEST": "āstreliāti weḍāhapū mānānka belā", "AFT": "aapganistan belā", "AKDT": "alaska ḍelaaiṭ belā", "AKST": "alaska mananka belā", "AMST": "āmajon karã masa belā", "AMST Armenia": "ārmeniā kār~ā belā", "AMT": "āmajon mananka belā", "AMT Armenia": "ārmeniā mānānka belā", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "ārjenṭinā kār~ā belā belā", "ART": "ārjenṭinā mānānka belā", "AST": "āṭlānṭik mānānka belā", "AST Arabia": "ārbiāti mānānka belā", "AWDT": "āstreliāti weḍākuṇupū delāiṭ belā", "AWST": "āstreliāti weḍākuṇupū mānānka belā", "AZST": "ājerbāijān kār~ā belā", "AZT": "ājerbāijān mānānka belā", "BDT Bangladesh": "bānglādes kār~ā belā", "BNT": "bruneti dārusālām belā", "BOT": "bolwiā belā", "BRST": "brājiliā kār~ā belā", "BRT": "brājiliā mānānka belā", "BST Bangladesh": "bānglādes mānānka belā", "BT": "butān belā", "CAST": "CAST", "CAT": "madini aaprika belā", "CCT": "kokos dīp belā", "CDT": "madinī ḍelaaiṭ belā", "CHADT": "cyātām delāit belā", "CHAST": "cyātām mānānka belā", "CHUT": "cuk belā", "CKT": "kuk dīp mānānka belā", "CKT DST": "kuk dīp ādā kār~ā belā", "CLST": "cili kār~ā belā", "CLT": "cili mānānka belā", "COST": "kolombiā kār~ā belā", "COT": "kolombiā mānānka belā", "CST": "madinī mananka belā", "CST China": "cin mānānka belā", "CST China DST": "cin delāiṭ belā", "CVST": "kep bḍ kār~ā belā", "CVT": "kep bḍ mānānka belā", "CXT": "krismās dīp belā", "ChST": "cāmor mānānka belā", "ChST NMI": "ChST NMI", "CuDT": "kubā delāiṭ belā", "CuST": "kubā mānānka belā", "DAVT": "debis belā", "DDUT": "ḍumonṭ ḍi arwilē belā", "EASST": "isṭr dīp kār~ā belā", "EAST": "isṭr dīp mānānka belā", "EAT": "veḍa hpu aaprika belā", "ECT": "iquāḍor belā", "EDT": "veḍahapu ḍelaaiṭ belā", "EGDT": "weḍāhpu grinlānd kār~ā belā", "EGST": "weḍāhpu grinlānd mānānka belā", "EST": "veḍahapu mananka belā", "FEET": "ar weḍāhpu yuropati belā", "FJT": "piji mānānka belā", "FJT Summer": "piji kār~ā belā", "FKST": "paklyānḍ dīpati kār~ā belā", "FKT": "paklyānḍ dīpati mānānka belā", "FNST": "parn~ḍo ḍe norohā~ kār~ā belā", "FNT": "parn~ḍo ḍe norohā~ mānānka belā", "GALT": "gālāpogs belā", "GAMT": "gambiyr belā", "GEST": "jarjiā kār~ā belā", "GET": "jarjiā mānānka belā", "GFT": "prench guyān belā", "GIT": "gīlbrṭ dīp belā", "GMT": "grinwic mīn belā", "GNSST": "GNSST", "GNST": "GNST", "GST": "dkīṇa jarjīā belā", "GST Guam": "GST Guam", "GYT": "guyān belā", "HADT": "hāwāe- aalūsān mānānka belā", "HAST": "hāwāe- aalūsān mānānka belā", "HKST": "hang kang kār~ā belā", "HKT": "hang kang mānānka belā", "HOVST": "hwoḍ kār~ā belā", "HOVT": "hwoḍ mānānka belā", "ICT": "inḍocinā belā", "IDT": "ijarāīl ḍelāiṭ belā", "IOT": "bārat kājā belā", "IRKST": "īrkustak kār~ā belā", "IRKT": "īrkustak mānānka belā", "IRST": "irān mānānka belā", "IRST DST": "irān ḍelāiṭ belā", "IST": "bārat mānānka belā", "IST Israel": "ijarāīl mānānka belā", "JDT": "jāpān ḍelāiṭ belā", "JST": "jāpān mānānka belā", "KOST": "kasrāē belā", "KRAST": "krāsnōrska kār~ā belā", "KRAT": "krāsnōrska mānānka belā", "KST": "koriān mānānka belā", "KST DST": "koriān ḍelāiṭ belā", "LHDT": "laṛ hawe ḍelāiṭ belā", "LHST": "laṛ hawe mānānka belā", "LINT": "lāin dīp belā", "MAGST": "māgādan ḍelāit belā", "MAGT": "māgādan mānānka belā", "MART": "mārksas belā", "MAWT": "māwosn belā", "MDT": "harka ḍelaaiṭ belā", "MESZ": "mādinā yuropiāti kār~ā belā", "MEZ": "mādinā yuropiāti mānānka belā", "MHT": "mārsāl dīpa belā", "MMT": "miñyāmār belā", "MSD": "mosko kār~ā belā", "MST": "harka mananka belā", "MUST": "marīsas kār~ā belā", "MUT": "marīsasmānānka belā", "MVT": "māldīp belā", "MYT": "malesiā belā", "NCT": "niū keleḍoniā mānānka belā", "NDT": "niūpaunḍlyānḍ ḍelāiṭ belā", "NDT New Caledonia": "niū keleḍoniā kār~ā belā", "NFDT": "norpok dīp ḍelāiṭ belā", "NFT": "norpok dīp mānānka belā", "NOVST": "nawosibisrk kār~ā belā", "NOVT": "nawosibisrk mānānka belā", "NPT": "nepāl belā", "NRT": "nāurū belā", "NST": "niūpaunḍlyānḍ mānānka belā", "NUT": "niū belā", "NZDT": "niūjilānḍ ḍelāiṭ belā", "NZST": "niūjilānḍ mānānka belā", "OESZ": "weḍāhpu yuropiāti kār~ā belā", "OEZ": "weḍāhpu yuropiāti mānānka belā", "OMSST": "amska kār~ā belā", "OMST": "amska mānānka belā", "PDT": "pesipik ḍelaaiṭ belā", "PDTM": "meksikān pesipic ḍelāiṭ belā", "PETDT": "PETDT", "PETST": "PETST", "PGT": "pāpuā niu gunīā belā", "PHOT": "piniksa dīpati belā", "PKT": "pākistān mānānka belā", "PKT DST": "pākistān kār~ā belā", "PMDT": "sēnṭ pierē aḍē mekwēlān ḍelāiṭ belā", "PMST": "sēnṭ pierē aḍē mekwēlān mānānka belā", "PONT": "ponāpe belā", "PST": "pesipik mananka belā", "PST Philippine": "pilipin mānānka belā", "PST Philippine DST": "pilipin kār~ā belā", "PST Pitcairn": "piṭkēran belā", "PSTM": "meksikān pesipic mānānka belā", "PWT": "pālāu belā", "PYST": "pārāguyē kār~ā belā", "PYT": "pārāguyē mānānka belā", "PYT Korea": "pyongayāng belā", "RET": "rīūnīan belā", "ROTT": "rotērā belā", "SAKST": "sakālin kār~ā belā", "SAKT": "sakālin mānānka belā", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "dkiṇ aaprika belā", "SBT": "soloman dīpati belā", "SCT": "sēsels belā", "SGT": "singāpur mānānka belā", "SLST": "SLST", "SRT": "surīnām belā", "SST Samoa": "saāmoa mānānka belā", "SST Samoa Apia": "epiā mānānka belā", "SST Samoa Apia DST": "epiā delāit belā", "SST Samoa DST": "saāmoa ḍelāiṭ", "SYOT": "sawā belā", "TAAF": "prenc dakiṇ aḍe aanṭārkṭik belā", "TAHT": "tāhiti belā", "TJT": "tājikistān belā", "TKT": "ṭokelāu belā", "TLT": "weḍāhapu timor belā", "TMST": "turkmenistān kār~ā belā", "TMT": "turkmenistān manaka belā", "TOST": "ṭangā kār~ā belā", "TOT": "ṭangā mānānka belā", "TVT": "tuwalū belā", "TWT": "tāipē mānānka belā", "TWT DST": "tāipē ḍelāiṭ", "ULAST": "ūlānbator kār~ā belā", "ULAT": "ūlānbator mānānka belā", "UYST": "urguwē kār~ā belā", "UYT": "urguwē manaka belā", "UZT": "uzwēkistān mānānka belā", "UZT DST": "uzwēkistān kārã belā", "VET": "wenēzuelā belā", "VLAST": "wlādiwostak kār~ā belā", "VLAT": "wlādiwostak mānānka belā", "VOLST": "wālgogrāḍ kār~ā belā", "VOLT": "wālgogrāḍ mānānka belā", "VOST": "wostak belā", "VUT": "wanuātū mānānka belā", "VUT DST": "wanuātū kār~ā belā", "WAKT": "waka dīpa belā", "WARST": "weḍākuṇupu ārjenṭinā kār~ā belā", "WART": "weḍākuṇupu ārjenṭinā mānānka belā", "WAST": "veḍakuṇpu aaprika belā", "WAT": "veḍakuṇpu aaprika belā", "WESZ": "weḍākūṇūpu yuropiāti kār~ā belā", "WEZ": "weḍākūṇūpu yuropiāti mānānka belā", "WFT": "walis aḍē puṭunā", "WGST": "weḍakūṇūp grinlānḍ kār~ā belā", "WGT": "weḍakūṇūp grinlānḍ mānānka belā", "WIB": "weḍākūṇpū inḍnesiā belā", "WIT": "weḍāhpu inḍnesiā belā", "WITA": "mādini inḍnesiā belā", "YAKST": "yakustuk kār~ā belā", "YAKT": "yakustuk mānānka belā", "YEKST": "yakaterinbarg kār~ā belā", "YEKT": "yakaterinbarg mānānka belā", "YST": "YST", "МСК": "mosko mānānka belā", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "weḍākūṇpū kājākstān belā", "شىعىش قازاق ەلى": "weḍāhapu kājākstān belā", "قازاق ەلى": "kājākstān belā", "قىرعىزستان": "kirgijstān belā", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "ajores kār~ā belā"},
	}
}

// Locale returns the current translators string locale
func (kxv *kxv_Orya_IN) Locale() string {
	return kxv.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'kxv_Orya_IN'
func (kxv *kxv_Orya_IN) PluralsCardinal() []locales.PluralRule {
	return kxv.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'kxv_Orya_IN'
func (kxv *kxv_Orya_IN) PluralsOrdinal() []locales.PluralRule {
	return kxv.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'kxv_Orya_IN'
func (kxv *kxv_Orya_IN) PluralsRange() []locales.PluralRule {
	return kxv.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'kxv_Orya_IN'
func (kxv *kxv_Orya_IN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'kxv_Orya_IN'
func (kxv *kxv_Orya_IN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'kxv_Orya_IN'
func (kxv *kxv_Orya_IN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (kxv *kxv_Orya_IN) MonthAbbreviated(month time.Month) string {
	return kxv.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (kxv *kxv_Orya_IN) MonthsAbbreviated() []string {
	return kxv.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (kxv *kxv_Orya_IN) MonthNarrow(month time.Month) string {
	return kxv.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (kxv *kxv_Orya_IN) MonthsNarrow() []string {
	return kxv.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (kxv *kxv_Orya_IN) MonthWide(month time.Month) string {
	return kxv.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (kxv *kxv_Orya_IN) MonthsWide() []string {
	return kxv.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (kxv *kxv_Orya_IN) WeekdayAbbreviated(weekday time.Weekday) string {
	return kxv.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (kxv *kxv_Orya_IN) WeekdaysAbbreviated() []string {
	return kxv.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (kxv *kxv_Orya_IN) WeekdayNarrow(weekday time.Weekday) string {
	return kxv.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (kxv *kxv_Orya_IN) WeekdaysNarrow() []string {
	return kxv.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (kxv *kxv_Orya_IN) WeekdayShort(weekday time.Weekday) string {
	return kxv.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (kxv *kxv_Orya_IN) WeekdaysShort() []string {
	return kxv.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (kxv *kxv_Orya_IN) WeekdayWide(weekday time.Weekday) string {
	return kxv.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (kxv *kxv_Orya_IN) WeekdaysWide() []string {
	return kxv.daysWide
}

// Decimal returns the decimal point of number
func (kxv *kxv_Orya_IN) Decimal() string {
	return kxv.decimal
}

// Group returns the group of number
func (kxv *kxv_Orya_IN) Group() string {
	return kxv.group
}

// Group returns the minus sign of number
func (kxv *kxv_Orya_IN) Minus() string {
	return kxv.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'kxv_Orya_IN' and handles both Whole and Real numbers based on 'v'
func (kxv *kxv_Orya_IN) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 0
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kxv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, kxv.group[0])
				count = 1

				if !inSecondary {
					inSecondary = true
					groupThreshold = 2
				}
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, kxv.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'kxv_Orya_IN' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (kxv *kxv_Orya_IN) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'kxv_Orya_IN'
func (kxv *kxv_Orya_IN) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kxv.currencies[currency]
	l := len(s) + len(symbol) + 2
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kxv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, kxv.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	for j := len(kxv.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, kxv.currencyPositivePrefix[j])
	}

	if num < 0 {
		b = append(b, kxv.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, kxv.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'kxv_Orya_IN'
// in accounting notation.
func (kxv *kxv_Orya_IN) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := kxv.currencies[currency]
	l := len(s) + len(symbol) + 4
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, kxv.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, kxv.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(kxv.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, kxv.currencyNegativePrefix[j])
		}

	} else {

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		for j := len(kxv.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, kxv.currencyPositivePrefix[j])
		}

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, kxv.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, kxv.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'kxv_Orya_IN'
func (kxv *kxv_Orya_IN) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'kxv_Orya_IN'
func (kxv *kxv_Orya_IN) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kxv.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'kxv_Orya_IN'
func (kxv *kxv_Orya_IN) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kxv.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'kxv_Orya_IN'
func (kxv *kxv_Orya_IN) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, kxv.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, kxv.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'kxv_Orya_IN'
func (kxv *kxv_Orya_IN) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, kxv.periodsAbbreviated[0]...)
	} else {
		b = append(b, kxv.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'kxv_Orya_IN'
func (kxv *kxv_Orya_IN) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, kxv.periodsAbbreviated[0]...)
	} else {
		b = append(b, kxv.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'kxv_Orya_IN'
func (kxv *kxv_Orya_IN) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, kxv.periodsAbbreviated[0]...)
	} else {
		b = append(b, kxv.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'kxv_Orya_IN'
func (kxv *kxv_Orya_IN) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, kxv.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, kxv.periodsAbbreviated[0]...)
	} else {
		b = append(b, kxv.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := kxv.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
