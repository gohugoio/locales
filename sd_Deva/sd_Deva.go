package sd_Deva

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type sd_Deva struct {
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
	currencyPositivePrefix string
	currencyNegativePrefix string
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

// New returns a new instance of translator for the 'sd_Deva' locale
func New() locales.Translator {
	return &sd_Deva{
		locale:                 "sd_Deva",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "؋", "ALK", "ALL", "֏", "ANG", "Kz", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "$", "ATS", "A$", "AWG", "AZM", "₼", "BAD", "KM", "BAN", "$", "৳", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "$", "$", "Bs", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "R$", "BRN", "BRR", "BRZ", "$", "BTN", "BUK", "P", "BYB", "р.", "BYR", "$", "CA$", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "$", "CNH", "CNX", "CN¥", "$", "COU", "₡", "CSD", "CSK", "$", "$", "CVE", "CYP", "Kč", "DDM", "DEM", "DJF", "kr", "$", "DZD", "ECS", "ECV", "EEK", "E£", "ERN", "ESA", "ESB", "₧", "ETB", "€", "FIM", "$", "£", "FRF", "£", "GEK", "₾", "GHC", "GH₵", "£", "GMD", "FG", "GNS", "GQE", "GRD", "Q", "GWE", "GWP", "$", "HK$", "L", "HRD", "kn", "HTG", "Ft", "Rp", "IEP", "ILP", "ILR", "₪", "₹", "IQD", "IRR", "ISJ", "kr", "ITL", "$", "JOD", "¥", "KES", "⃀", "៛", "CF", "₩", "KRH", "KRO", "₩", "KWD", "$", "₸", "₭", "L£", "Rs", "$", "LSL", "Lt", "LTT", "LUC", "LUF", "LUL", "Ls", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "Ar", "MGF", "MKD", "MKN", "MLF", "K", "₮", "MOP", "MRO", "MRU", "MTL", "MTP", "Rs", "MVP", "MVR", "MWK", "MX$", "MXP", "MXV", "RM", "MZE", "MZM", "MZN", "$", "₦", "NIC", "C$", "NLG", "kr", "Rs", "NZ$", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "₱", "Rs", "zł", "PLZ", "PTE", "₲", "QAR", "RHD", "ROL", "lei", "RSD", "₽", "RUR", "RF", "\\u20c1", "$", "SCR", "SDD", "SDG", "SDP", "kr", "$", "£", "SIT", "SKK", "SLE", "SLL", "SOS", "$", "SRG", "£", "STD", "Db", "SUR", "SVC", "£", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "T$", "TPE", "TRL", "₺", "$", "NT$", "TZS", "₴", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "$", "UYW", "UZS", "VEB", "VED", "Bs", "VES", "₫", "VNN", "VUV", "WST", "FCFA", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "EC$", "Cg.", "XDR", "XEU", "XFO", "XFU", "F\\u202fCFA", "XPD", "CFPF", "XPT", "XRE", "XSU", "XTS", "XUA", "¤", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "R", "ZMK", "ZK", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "जन", "फर", "मार्च", "अप्रै", "मई", "जून", "जु", "अग", "सप्टे", "ऑक्टो", "नवं", "डिसं"},
		monthsNarrow:           []string{"", "ज", "फ़", "मा", "अ", "मा", "जू", "जु", "अग", "स", "ऑ", "न", "डि"},
		monthsWide:             []string{"", "जनवरी", "फेबरवरी", "मार्चु", "अप्रेल", "मई", "जून", "जुलाई", "आगस्ट", "सप्टेंबर", "आक्टोबर", "नवंबर", "डिसंबर"},
		daysAbbreviated:        []string{"आर्त", "सू", "मंग", "ॿुध", "विस", "जुम", "छंछ"},
		daysNarrow:             []string{"आ", "सू", "मं", "ॿु", "वि", "जु", "छं"},
		daysWide:               []string{"आर्तवार", "सूमर", "मंगलु", "ॿुधर", "विस्पत", "जुमो", "छंछर"},
		periodsAbbreviated:     []string{"AM", "PM"},
		timezones:              map[string]string{"ACDT": "آسٽريليا جو مرڪزي ڏينهن جو وقت", "ACST": "ACST", "ACT": "ACT", "ACWDT": "آسٽريليا جو مرڪزي مغربي ڏينهن جو وقت", "ACWST": "آسٽريليا جو مرڪزي مغربي معياري وقت", "ADT": "अटलांटिक दीं॒ह जो वक्त", "ADT Arabia": "عربين جي ڏينهن جو وقت", "AEDT": "آسٽريليا جو مشرقي ڏينهن جو وقت", "AEST": "آسٽريليا جو مشرقي معياري وقت", "AFT": "افغانستان جو وقت", "AKDT": "الاسڪا جي ڏينهن جو وقت", "AKST": "الاسڪا جو معياري وقت", "AMST": "ايميزون جي اونهاري جو وقت", "AMST Armenia": "آرمينيا جي اونهاري جو وقت", "AMT": "ايميزون جو معياري وقت", "AMT Armenia": "آرمينيا جو معياري وقت", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "ارجنٽائن جي اونهاري جو وقت", "ART": "ارجنٽائن معياري وقت", "AST": "अटलांटिक मअयारी वक्त", "AST Arabia": "عربين جو معياري وقت", "AWDT": "آسٽريليا جو مغربي ڏينهن جو وقت", "AWST": "آسٽريليا جو مغربي معياري وقت", "AZST": "آذربائيجان جي اونهاري جو وقت", "AZT": "آذربائيجان جو معياري وقت", "BDT Bangladesh": "بنگلاديش جي اونهاري جو وقت", "BNT": "برونائي دارالسلام جو وقت", "BOT": "بولويائي وقت", "BRST": "براسيليا جي اونهاري جو وقت", "BRT": "براسيليا جو معياري وقت", "BST Bangladesh": "بنگلاديش جو معياري وقت", "BT": "ڀوٽان جو وقت", "CAST": "CAST", "CAT": "مرڪزي آفريقا جو وقت", "CCT": "ڪوڪوس آئي لينڊ جو وقت", "CDT": "मरकज़ी दीं॒ह जो वक्त", "CHADT": "چئٿم جي ڏينهن جو وقت", "CHAST": "چئٿم جو معياري وقت", "CHUT": "چيوڪ جو وقت", "CKT": "ڪوڪ آئي لينڊ جو معياري وقت", "CKT DST": "ڪوڪ آئي لينڊ جي اڌ اونهاري جو وقت", "CLST": "چلي جي اونهاري جو وقت", "CLT": "چلي جو معياري وقت", "COST": "ڪولمبيا جي اونهاري جو وقت", "COT": "ڪولمبيا جو معياري وقت", "CST": "मरकज़ी मअयारी वक्त", "CST China": "چائنا جو معياري وقت", "CST China DST": "چائنا جي ڏينهن جو وقت", "CVST": "ڪيپ ورڊ جي اونهاري جو وقت", "CVT": "ڪيپ ورڊ جو معياري وقت", "CXT": "ڪرسمس آئي لينڊ جو وقت", "ChST": "چمورو جو معياري وقت", "ChST NMI": "ChST NMI", "CuDT": "ڪيوبا جي ڏينهن جو وقت", "CuST": "ڪيوبا جو معياري وقت", "DAVT": "ڊيوس جو وقت", "DDUT": "ڊومانٽ درويئل جو وقت", "EASST": "ايسٽر آئي لينڊ جي اونهاري جو وقت", "EAST": "ايسٽر آئي لينڊ جو معياري وقت", "EAT": "اوڀر آفريڪا جو وقت", "ECT": "ايڪواڊور جو وقت", "EDT": "ओभरी दीं॒ह जो वक्त", "EGDT": "مشرقي گرين لينڊ جي اونهاري جو وقت", "EGST": "مشرقي گرين لينڊ جو معياري وقت", "EST": "ओभरी मअयारी वक्त", "FEET": "وڌيڪ مشرقي يورپي وقت", "FJT": "فجي جو معياري وقت", "FJT Summer": "فجي جي اونهاري جو وقت", "FKST": "فاڪ لينڊ آئي لينڊ جي اونهاري جو وقت", "FKT": "فاڪ لينڊ آئي لينڊ جو معياري وقت", "FNST": "فرنانڊو دي نورونها جي اونهاري وقت", "FNT": "فرنانڊو دي نورونها جو معياري وقت", "GALT": "گالاپاگوز جو وقت", "GAMT": "گيمبيئر جو وقت", "GEST": "جارجيا جي اونهاري جو وقت", "GET": "جارجيا جو معياري وقت", "GFT": "فرانسيسي گيانا جو وقت", "GIT": "گلبرٽ آئي لينڊ جو وقت", "GMT": "ग्रीनविच मीन वक़्तु", "GNSST": "GNSST", "GNST": "GNST", "GST": "ڏکڻ جارجيا جو وقت", "GST Guam": "GST Guam", "GYT": "گيانائي وقت", "HADT": "هوائي اليوٽين جي ڏينهن جو وقت", "HAST": "هوائي اليوٽين جو معياري وقت", "HKST": "هانگ ڪانگ جي اونهاري جو وقت", "HKT": "هانگ ڪانگ جو معياري وقت", "HOVST": "هووڊ جي اونهاري جو وقت", "HOVT": "هووڊ جو معياري وقت", "ICT": "انڊو چائنا جو وقت", "IDT": "اسرائيل جي ڏينهن جو وقت", "IOT": "هند سمنڊ جو وقت", "IRKST": "ارڪتسڪ جي ڏينهن جو وقت", "IRKT": "ارڪتسڪ جو معياري وقت", "IRST": "ايران جو معياري وقت", "IRST DST": "ايران جي ڏينهن جو وقت", "IST": "ڀارت جو معياري وقت", "IST Israel": "اسرائيل جو معياري وقت", "JDT": "جاپان جي ڏينهن جو وقت", "JST": "جاپان جو معياري وقت", "KOST": "ڪوسرائي جو وقت", "KRAST": "ڪریسنویارسڪ جي ڏينهن جو وقت", "KRAT": "ڪریسنویارسڪ جو معياري وقت", "KST": "ڪوريا جو معياري وقت", "KST DST": "ڪوريا جي ڏينهن جو وقت", "LHDT": "لورڊ هووي جي ڏينهن جو وقت", "LHST": "لورڊ هووي جو معياري وقت", "LINT": "لائن آئي لينڊ جو وقت", "MAGST": "مگادان جي ڏينهن جي وقت", "MAGT": "مگادان جو معياري وقت", "MART": "مرڪيوسس جو وقت", "MAWT": "مائوسن جو وقت", "MDT": "पहाड़ी दीं॒ह जो वक्त", "MESZ": "मरकज़ी यूरोपी उनहारे जो वक्तु", "MEZ": "मरकज़ी यूरोपी मअयारी वक्तु", "MHT": "مارشل آئي لينڊ جو وقت", "MMT": "ميانمار جو وقت", "MSD": "ماسڪو جي ڏينهن جي وقت", "MST": "पहाड़ी मअयारी वक्त", "MUST": "موريشيس جي اونهاري جو وقت", "MUT": "موريشيس جو معياري وقت", "MVT": "مالديپ جو وقت", "MYT": "ملائيشيا جو وقت", "NCT": "نيو ڪيليڊونيا جو معياري وقت", "NDT": "نيو فائونڊ لينڊ جي ڏينهن جو وقت", "NDT New Caledonia": "نيو ڪيليڊونيا جي اونهاري جو وقت", "NFDT": "نار فوڪ آئي لينڊ جي ڏينهن جو وقت", "NFT": "نار فوڪ آئي لينڊ جو معياري وقت", "NOVST": "نوواسبئيرسڪ جي ڏينهن جو وقت", "NOVT": "نوواسبئيرسڪ جو معياري وقت", "NPT": "نيپال جو وقت", "NRT": "نائورو جو وقت", "NST": "نيو فائونڊ لينڊ جو معياري وقت", "NUT": "نيووي جو وقت", "NZDT": "نيوزيلينڊ جي ڏينهن جو وقت", "NZST": "نيوزيلينڊ جو معياري وقت", "OESZ": "ओभरी यूरोपी उनहारे जो वक्तु", "OEZ": "ओभरी यूरोपी मअयारी वक्तु", "OMSST": "اومسڪ جي ڏينهن جو وقت", "OMST": "اومسڪ جو معياري وقت", "PDT": "पेसिफिक दीं॒ह जो वक्त", "PDTM": "ميڪسيڪن پيسيفڪ جي ڏينهن جو وقت", "PETDT": "PETDT", "PETST": "PETST", "PGT": "پاپوا نيو گني جو وقت", "PHOT": "فونيڪس آئي لينڊ جو وقت", "PKT": "پاڪستان جو معياري وقت", "PKT DST": "پاڪستان جي اونهاري جو وقت", "PMDT": "سينٽ پيئر ۽ ميڪلون جي ڏينهن جو وقت", "PMST": "سينٽ پيئر ۽ ميڪلون جو معياري وقت", "PONT": "پوناپي جو وقت", "PST": "पेसिफिक मअयारी वक्त", "PST Philippine": "فلپائن جو معياري وقت", "PST Philippine DST": "فلپائن جي اونهاري جو وقت", "PST Pitcairn": "پٽڪيرن جو وقت", "PSTM": "ميڪسيڪن پيسيفڪ جو معياري وقت", "PWT": "پلائو جو وقت", "PYST": "پيراگوئي جي اونهاري جو وقت", "PYT": "پيراگوئي جو معياري وقت", "PYT Korea": "پيانگ يانگ جو وقت", "RET": "ري يونين جو وقت", "ROTT": "روٿيرا جو وقت", "SAKST": "سخالين جي ڏينهن جو وقت", "SAKT": "سخالين جو معياري وقت", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "ڏکڻ آفريڪا جو معياري وقت", "SBT": "سولومن آئي لينڊ جو وقت", "SCT": "شي شلز جو وقت", "SGT": "سنگاپور جو معياري وقت", "SLST": "SLST", "SRT": "سوري نام جو وقت", "SST Samoa": "ساموا جو معياري وقت", "SST Samoa Apia": "اپيا جو معياري وقت", "SST Samoa Apia DST": "اپيا جي ڏينهن جو وقت", "SST Samoa DST": "ساموا جي ڏينهن جو وقت", "SYOT": "سائيوا جو وقت", "TAAF": "فرانسيسي ڏاکڻيو ۽ انٽارڪٽڪ وقت", "TAHT": "تاهيٽي جو وقت", "TJT": "تاجڪستان جو وقت", "TKT": "ٽوڪيلائو جو وقت", "TLT": "اوڀر تيمور جو وقت", "TMST": "ترڪمانستان جي اونهاري جو وقت", "TMT": "ترڪمانستان جو معياري وقت", "TOST": "ٽونگا جي اونهاري جو وقت", "TOT": "ٽونگا جو معياري وقت", "TVT": "تووالو جو وقت", "TWT": "تائپي جو معياري وقت", "TWT DST": "تائپي جي ڏينهن جو وقت", "ULAST": "اولان باتر جي اونهاري جو وقت", "ULAT": "اولان باتر جو معياري وقت", "UYST": "يوروگائي جي اونهاري جو وقت", "UYT": "يوروگائي جو معياري وقت", "UZT": "ازبڪستان جو معياري وقت", "UZT DST": "ازبڪستان جي اونهاري جو وقت", "VET": "وينزويلا جو وقت", "VLAST": "اولادووستوڪ جي ڏينهن جو وقت", "VLAT": "ولادووستوڪ جو معياري وقت", "VOLST": "وولگوگراد جي ڏينهن جو وقت", "VOLT": "وولگوگراد جو معياري وقت", "VOST": "ووسٽوڪ جو وقت", "VUT": "وانواتو جو معياري وقت", "VUT DST": "وانواتو جي ڏينهن جو وقت", "WAKT": "ويڪ آئي لينڊ جو وقت", "WARST": "اولهه ارجنٽينا جي اونهاري جو وقت", "WART": "اولهه ارجنٽينا جو معياري وقت", "WAST": "اولهه آفريقا جو وقت", "WAT": "اولهه آفريقا جو وقت", "WESZ": "उलहंदो यूरोपी उनहारे जो वक्तु", "WEZ": "उलहंदो यूरोपी मअयारी वक्तु", "WFT": "ويلس ۽ فتونا جو وقت", "WGST": "مغربي گرين لينڊ جي اونهاري جو وقت", "WGT": "مغربي گرين لينڊ جو معياري وقت", "WIB": "اولهه انڊونيشيا جو وقت", "WIT": "اوڀر انڊونيشيا جو وقت", "WITA": "مرڪزي انڊونيشيا جو وقت", "YAKST": "ياڪتسڪ جي ڏينهن جو وقت", "YAKT": "ياڪتسڪ جو معياري وقت", "YEKST": "يڪاٽيرنبرگ جي ڏينهن جو وقت", "YEKT": "يڪاٽيرنبرگ جو معياري وقت", "YST": "يڪون جو وقت", "МСК": "ماسڪو جو معياري وقت", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "اولهه قازقستان جو وقت", "شىعىش قازاق ەلى": "اوڀر قزاقستان جو وقت", "قازاق ەلى": "قزاقستان وقت", "قىرعىزستان": "ڪرغزستان جو وقت", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "پيرو جي اونهاري جو وقت"},
	}
}

// Locale returns the current translators string locale
func (sd *sd_Deva) Locale() string {
	return sd.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'sd_Deva'
func (sd *sd_Deva) PluralsCardinal() []locales.PluralRule {
	return sd.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'sd_Deva'
func (sd *sd_Deva) PluralsOrdinal() []locales.PluralRule {
	return sd.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'sd_Deva'
func (sd *sd_Deva) PluralsRange() []locales.PluralRule {
	return sd.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'sd_Deva'
func (sd *sd_Deva) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'sd_Deva'
func (sd *sd_Deva) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'sd_Deva'
func (sd *sd_Deva) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := sd.CardinalPluralRule(num1, v1)
	end := sd.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (sd *sd_Deva) MonthAbbreviated(month time.Month) string {
	if len(sd.monthsAbbreviated) == 0 {
		return ""
	}
	return sd.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (sd *sd_Deva) MonthsAbbreviated() []string {
	return sd.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (sd *sd_Deva) MonthNarrow(month time.Month) string {
	if len(sd.monthsNarrow) == 0 {
		return ""
	}
	return sd.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (sd *sd_Deva) MonthsNarrow() []string {
	return sd.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (sd *sd_Deva) MonthWide(month time.Month) string {
	if len(sd.monthsWide) == 0 {
		return ""
	}
	return sd.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (sd *sd_Deva) MonthsWide() []string {
	return sd.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (sd *sd_Deva) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(sd.daysAbbreviated) == 0 {
		return ""
	}
	return sd.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (sd *sd_Deva) WeekdaysAbbreviated() []string {
	return sd.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (sd *sd_Deva) WeekdayNarrow(weekday time.Weekday) string {
	if len(sd.daysNarrow) == 0 {
		return ""
	}
	return sd.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (sd *sd_Deva) WeekdaysNarrow() []string {
	return sd.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (sd *sd_Deva) WeekdayShort(weekday time.Weekday) string {
	if len(sd.daysShort) == 0 {
		return ""
	}
	return sd.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (sd *sd_Deva) WeekdaysShort() []string {
	return sd.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (sd *sd_Deva) WeekdayWide(weekday time.Weekday) string {
	if len(sd.daysWide) == 0 {
		return ""
	}
	return sd.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (sd *sd_Deva) WeekdaysWide() []string {
	return sd.daysWide
}

// Decimal returns the decimal point of number
func (sd *sd_Deva) Decimal() string {
	return sd.decimal
}

// Group returns the group of number
func (sd *sd_Deva) Group() string {
	return sd.group
}

// Group returns the minus sign of number
func (sd *sd_Deva) Minus() string {
	return sd.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'sd_Deva' and handles both Whole and Real numbers based on 'v'
func (sd *sd_Deva) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sd.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sd.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sd.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'sd_Deva' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (sd *sd_Deva) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sd.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, sd.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, sd.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'sd_Deva'
func (sd *sd_Deva) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sd.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sd.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sd.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(sd.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, sd.currencyPositivePrefix[j])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		b = append(b, sd.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sd.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'sd_Deva'
// in accounting notation.
func (sd *sd_Deva) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := sd.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, sd.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, sd.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(sd.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, sd.currencyNegativePrefix[j])
		}

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		b = append(b, sd.minus[0])

	} else {

		for j := len(sd.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, sd.currencyPositivePrefix[j])
		}

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, sd.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'sd_Deva'
func (sd *sd_Deva) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Month()), 10)
	b = append(b, []byte{0x2f}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2f}...)

	if t.Year() > 9 {
		b = append(b, strconv.Itoa(t.Year())[2:]...)
	} else {
		b = append(b, strconv.Itoa(t.Year())[1:]...)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'sd_Deva'
func (sd *sd_Deva) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, sd.monthsAbbreviated[t.Month()]...)
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

// FmtDateLong returns the long date representation of 't' for 'sd_Deva'
func (sd *sd_Deva) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, sd.monthsWide[t.Month()]...)
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

// FmtDateFull returns the full date representation of 't' for 'sd_Deva'
func (sd *sd_Deva) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, sd.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, sd.monthsWide[t.Month()]...)
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

// FmtTimeShort returns the short time representation of 't' for 'sd_Deva'
func (sd *sd_Deva) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sd.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sd.periodsAbbreviated[0]...)
	} else {
		b = append(b, sd.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'sd_Deva'
func (sd *sd_Deva) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sd.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sd.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sd.periodsAbbreviated[0]...)
	} else {
		b = append(b, sd.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'sd_Deva'
func (sd *sd_Deva) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sd.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sd.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sd.periodsAbbreviated[0]...)
	} else {
		b = append(b, sd.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'sd_Deva'
func (sd *sd_Deva) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, sd.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, sd.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, sd.periodsAbbreviated[0]...)
	} else {
		b = append(b, sd.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := sd.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
