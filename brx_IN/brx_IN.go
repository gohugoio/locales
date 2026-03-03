package brx_IN

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type brx_IN struct {
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

// New returns a new instance of translator for the 'brx_IN' locale
func New() locales.Translator {
	return &brx_IN{
		locale:                 "brx_IN",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         nil,
		pluralsRange:           nil,
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "ए.इ.दि", "AFA", "AFN", "ALK", "अल", "ए.एम.दि", "ए.एन.जि", "ए.अ.ए", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ए.आर.एस", "ATS", "AUD", "ए.दब्ल्यु.जि", "AZM", "ए.जेत.एन", "BAD", "बि.ए.एम", "BAN", "बि.बि.दि", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "बि.जि.एन", "BGO", "बि.ऐत्स.दि", "बि.आइ.एफ", "बि.एम.डि", "BND", "बि.अ.बि", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "बि.एस.दि", "BTN", "BUK", "बि.दब्ल्यु.पि", "BYB", "बि.वाई.एन", "BYR", "बि.जेद.डि", "सि.ए $", "सि.दि.एफ", "CHE", "सि.ऐत्स.एफ", "CHW", "CLE", "CLF", "सि.एल.पि", "सि.एन.ऐत्स", "CNX", "सिएन¥", "सि.अ.पि", "COU", "सि.आर.सि", "CSD", "CSK", "सि.इउ.सि", "सि.इउ.पि", "सि.भि.इ", "CYP", "सि.जेद.के", "DDM", "DEM", "दि.जे.एफ", "दि.के.के", "डि.अ.पि", "दि.जेत.दि", "ECS", "ECV", "EEK", "ई.जि.पि", "इ.आर.एन", "ESA", "ESB", "ESP", "इ.ति.बि", "EUR", "FIM", "FJD", "एफ.के.पि", "FRF", "GBP", "GEK", "जि.इ.एल", "GHC", "जि.ऐत्स.एस", "जि.आइ.पि", "जि.एम.दि", "जि.एन.एफ", "GNS", "GQE", "GRD", "जि.ति.किउ", "GWE", "GWP", "जि.वाई.दि", "ऐत्स.के$", "ऐत्स.एन.एल", "HRD", "ऐत्स.आर.के", "ऐत्स.ति.जि", "ऐत्स.इउ.एफ", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "आइ.किउ.दि", "IRR", "ISJ", "आई.एस.के", "ITL", "जे.एम.दि", "जे.अ.दि", "JPY", "के.इ.एस", "के.जि.एस", "KHR", "के.एम.एफ", "के.पि.दब्ल्यु", "KRH", "KRO", "KRW", "के.दब्ल्यु.दि", "के.वाई.दि", "के.जेत.ति", "LAK", "एल.बि.पि", "LKR", "एल.आर.दि", "एल.एस.एल", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "एल.वाई.दि", "एम.ए.दि", "MAF", "MCF", "MDC", "एम.डि.एल", "एम.जि.ए", "MGF", "एम.के.दि", "MKN", "MLF", "MMK", "एम.एन.ति", "एम.अ.पि", "MRO", "एम.आर.इउ", "MTL", "MTP", "एम.इउ.आर", "MVP", "MVR", "एम.दब्ल्यु.के", "एम.एक्स $", "MXP", "MXV", "MYR", "MZE", "MZM", "एम.जेत.एन", "एन.ए.दि", "एन.जि.एन", "NIC", "एन.आई.अ", "NLG", "एन.अ.के", "NPR", "NZD", "अ.एम.आर", "पि.ए.बि", "PEI", "पि.इ.एन", "PES", "PGK", "PHP", "PKR", "पि.एल.एन", "PLZ", "PTE", "पि.आई.जि", "किउ.ए.आर", "RHD", "ROL", "आर.अ.एन", "आर.एस.दि", "रूब", "RUR", "आर.एफ", "एस.ए.आर", "SBD", "एस.सि.आर", "SDD", "एस.दि.जि", "SDP", "एस.इ.के", "SGD", "एस.ऐत्स.पि", "SIT", "SKK", "एस.एल.इ", "एस.एल.एल", "एस.अ.एस", "एस.आर.डि", "SRG", "एस.एस.पि", "STD", "एस.ति.एन", "SUR", "SVC", "एस.वाई.पि", "एस.जेत.एल", "THB", "TJR", "ति.जे.एस", "TMM", "ति.एम.ति", "ति.एन.दि", "TOP", "TPE", "TRL", "ति.आर.वाई", "ति.ति.डि", "एन.ति$", "ति.जेत.एस", "इउ.ए.ऐत्स", "UAK", "UGS", "इउ.जि.एक्स", "$", "USN", "USS", "UYI", "UYP", "इउ.वाई.इउ", "UYW", "इउ.जेत.एस", "VEB", "VED", "VEF", "भि.इ.एस", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "इ.सि $", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "वाई.इ.आर", "YUD", "YUM", "YUN", "YUR", "ZAL", "जेत.ए.आर", "ZMK", "जेत.के", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "जान", "फेब", "मार्च", "एप्रि", "मे", "जुन", "जुल", "आग", "सेप", "अक्ट’", "नवे", "डिसे"},
		monthsNarrow:           []string{"", "ज", "फ", "म", "ए", "म", "ज", "ज", "आ", "स", "अ", "न", "ड"},
		monthsWide:             []string{"", "जानुवारी", "फेब्रूवारी", "मार्च", "एप्रिल", "मे", "जुन", "जुलाई", "आगष्ट", "सेप्थेम्बर", "अक्ट’बर", "नवेम्बर", "डिसेम्बर"},
		daysAbbreviated:        []string{"रबि", "सम", "मंगल", "बुध", "बिस्थि", "सुखुर", "सनि"},
		daysNarrow:             []string{"र", "स", "मं", "बु", "बि", "सु", "स"},
		daysWide:               []string{"रबिबार", "समबार", "मंगलबार", "बुधबार", "बिस्थिबार", "सुखुरबार", "सनिबार"},
		periodsAbbreviated:     []string{"फुं", "बेलासे"},
		timezones:              map[string]string{"ACDT": "अ’स्ट्रेलिया मिरु सानारि सम", "ACST": "अ’स्ट्रेलियानि मिरु थाखोआरि सम", "ACT": "आकर स्टैंडर्ड टाईम", "ACWDT": "अस्ट्रेलियानि मिरु सोनाबारि सानारि सम", "ACWST": "अस्ट्रेलियानि मिरु सोनाबारि थाखोआरि सम", "ADT": "आटलान्टिक सानारि सम", "ADT Arabia": "आराबीयान सानारि सम", "AEDT": "अस्ट्रेलियानि सानजायारि सानारि सम", "AEST": "अस्ट्रेलियानि सानजायारि थाखोआरि सम", "AFT": "आफगानीस्तान सम", "AKDT": "आलास्कानि सानारि सम", "AKST": "आलास्का स्टेन्डार्ड सम", "AMST": "आमाजन दैज्लां सम", "AMST Armenia": "आर्मेनिया दैज्लां सम", "AMT": "आमाजन थाखोआऱि सम", "AMT Armenia": "आर्मेनिया थाखोआरि सम", "ANAST": "अनादीर समर टाईम", "ANAT": "अनादीर स्टैंडर्ड टाईम", "ARST": "आर्जेन्टिना दैज्लां सम", "ART": "आर्जेन्टिना थाखोआरि सम", "AST": "आटलान्टिक थाखोआरि सम", "AST Arabia": "आराबीयान थाखोआरि सम", "AWDT": "अस्ट्रेलियानि सोनाबारि सानारि सम", "AWST": "अस्ट्रेलियानि सोनाबारि थाखोआरि सम", "AZST": "आजारबाईजान दैज्लां सम", "AZT": "आजारबाईजान थाखोआरि सम", "BDT Bangladesh": "बांलादेश दैज्लां सम", "BNT": "ब्रुनेई दर उस सलाम स्टैंडर्ड टाईम", "BOT": "बलिभिया सम", "BRST": "ब्राजीलिया दैज्लां सम", "BRT": "ब्राजीलिया थाखोआरि सम", "BST Bangladesh": "बांलादेश थाखोआरि सम", "BT": "भुटान सम", "CAST": "CAST", "CAT": "मिरु आफ्रिका सम", "CCT": "कक’स द्वीपफोरनि सम", "CDT": "साहा आमेरिकानि मिरु सानारि सम", "CHADT": "चेथाम सानारि सम", "CHAST": "चेथाम थाखोआरि सम", "CHUT": "चूक सम", "CKT": "कुक द्वीपफोरनि थाखोआरि सम", "CKT DST": "कुक द्वीपफोरनि खावसे दैज्लां सम", "CLST": "चीले दैज्लां सम", "CLT": "चीले थाखोआरि सम", "COST": "कलम्बिया दैज्लां सम", "COT": "कलम्बिया थाखोआरि सम", "CST": "साहा आमेरिकानि मिरु मानथाखोआरि सम", "CST China": "चाईना थाखोआरि सम", "CST China DST": "चाईना सानारि सम", "CVST": "केप भेर्दे दैज्लां सम", "CVT": "केप भेर्दे थाखोआरि सम", "CXT": "ख्रीस्टमास द्वीप सम", "ChST": "चाम’र्र’ थाखोआरि सम", "ChST NMI": "नॉर्थ मारिआना स्टैंडर्ड टाईम", "CuDT": "किउबा सानारि सम", "CuST": "किउबा थाखोआरि सम", "DAVT": "डेभिस सम", "DDUT": "डुमन्त-डीआर्भील सम", "EASST": "सानजायारि द्वीप दैज्लां सम", "EAST": "सानजायारि द्वीप थाखोआरि सम", "EAT": "सानजा आफ्रिका सम", "ECT": "एकुवाडर सम", "EDT": "साहा आमेरिकानि सानजायारि सानारि सम", "EGDT": "सानडजा ग्रीनलेण्ड दैज्लां सम", "EGST": "सानजा ग्रीनलेण्ड थाखोआरि सम", "EST": "साहा आमेरिकानि सानजायारि थाखोआरि सम", "FEET": "गोजानसिन-सानजायारि युर’पनि सम", "FJT": "फीजी थाखोआरि सम", "FJT Summer": "फीजी दैज्लां सम", "FKST": "फकलेण्ड द्वीपफोरनि दैज्लां सम", "FKT": "फकलेण्ड द्वीपफोरनि थाखोआरि सम", "FNST": "फेरनान्द’ दे नर’न्हा दैज्लां सम", "FNT": "फेरनान्द’ दे नर’न्हा थाखोआरि सम", "GALT": "गालापाग’स सम", "GAMT": "गाम्बियेर सम", "GEST": "जर्जिया दैज्लां सम", "GET": "जर्जिया थाखोआरि सम", "GFT": "फ्रेन्च गुईयाना सम", "GIT": "गीलबार्ट द्वीपफोरनि सम", "GMT": "ग्रिनवीच गेजेरारि सम", "GNSST": "GNSST", "GNST": "GNST", "GST": "खोला जर्जिया सम", "GST Guam": "गुआम स्टैंडर्ड टाईम", "GYT": "गुयाना सम", "HADT": "हावाई-एल्युतियान सानारि सम", "HAST": "हावाई-एल्युतियान थाखोआरि सम", "HKST": "हंकं दैज्लां सम", "HKT": "हंकं थाखोआरि सम", "HOVST": "ह’भ्द दैज्लां सम", "HOVT": "ह’भ्द थाखोआरि सम", "ICT": "ईंडो चइना स्टैंडर्ड टाईम", "IDT": "इज्राईल सानारि सम", "IOT": "भारत लैथोमायारि सम", "IRKST": "ईर्कुत्स्क दैज्लांं सम", "IRKT": "ईर्कुत्स्क थाखोआरि सम", "IRST": "इरान थाखोआरि सम", "IRST DST": "इरान सानारि सम", "IST": "भारतारि थाखोआरि सम", "IST Israel": "इज्राईल थाखोआरि सम", "JDT": "जापान सानारि सम", "JST": "जापान थाखोआरि सम", "KOST": "कस्रे सम", "KRAST": "क्रास्न’यार्स्क दैज्लांं सम", "KRAT": "क्रास्न’यार्स्क थाखोआरि सम", "KST": "क’रिया थाखोआरि सम", "KST DST": "क’रिया सानारि सम", "LHDT": "लर्ड हावि सानारि सम", "LHST": "लर्ड हावि थाखोआरि सम", "LINT": "लाईन द्वीपफोरनि सम", "MAGST": "मागादान दैज्लां सम", "MAGT": "मागादान थाखोआरि सम", "MART": "मार्केसास सम", "MAWT": "मौसन सम", "MDT": "माकाऊ समर टाईम", "MESZ": "मिरु युर’पनि दैज्लां सम", "MEZ": "मिरु युर’पनि थाखोआरि सम", "MHT": "मार्शेल द्वीपफोरनि सम", "MMT": "म्यानमार स्टैंडर्ड टाईम", "MSD": "मस्कौ दैज्लां सम", "MST": "माकाऊ स्टैंडर्ड टाईम", "MUST": "म’रिशियासनि दैज्लां सम", "MUT": "म’रिशियासनि थाखोआरि सम", "MVT": "मालदीभ्स सम", "MYT": "मलेशिया स्टैंडर्ड टाईम", "NCT": "गोदान केलेडनिया थाखोआरि सम", "NDT": "निउफाउन्दलेन्द सानारि सम", "NDT New Caledonia": "गोदान केलेडनिया दैज्लां सम", "NFDT": "नरफ’क द्वीप सानारि सम", "NFT": "नरफ’क द्वीप थाखोआरि सम", "NOVST": "नभ’सिबिर्स्क दैज्लां सम", "NOVT": "नभ’सिबिर्स्क थाखोआरि सम", "NPT": "नेपाल सम", "NRT": "नाऊरु सम", "NST": "निउफाउन्दलेन्द थाखोआरि सम", "NUT": "नीऊए सम", "NZDT": "निउजिलेण्ड सानारि सम", "NZST": "निउजिलेण्ड थाखोआरि सम", "OESZ": "सानजायारि युर’पनि दैज्लां सम", "OEZ": "सानजायारि युर’पनि थाखोआरि सम", "OMSST": "अम्स्क दैज्लां सम", "OMST": "अम्स्क थाखोआरि सम", "PDT": "साहा आमेरिकानि पेसिफिक सानारि सम", "PDTM": "मेक्सिक’नि पेसिफिक सानारि सम", "PETDT": "पेत्रोपावलोस्क कामचटका समर टाईम", "PETST": "पेत्रोपावलोस्क कामचटका स्टैंडर्ड टाईम", "PGT": "पापुआ निउ गिनी सम", "PHOT": "फनीक्स द्वीपफोरनि सम", "PKT": "पाकिस्तान थाखोआरि सम", "PKT DST": "पाकिस्तान दैज्लां सम", "PMDT": "सैन्ट पियेर आरो मिक्वेलन सानारि सम", "PMST": "सैन्ट पियेर आरो मिक्वेलन थाखोआरि सम", "PONT": "पनापे सम", "PST": "साहा आमेरिकानि पेसिफिक थाखोआरि सम", "PST Philippine": "फीलीपीन्स स्टैंडर्ड टाईम", "PST Philippine DST": "फीलीपीन्स समर टाईम", "PST Pitcairn": "पीटकैर्न सम", "PSTM": "मेक्सिक’नि पेसिफिक थाखोआरि सम", "PWT": "पालाउ सम", "PYST": "पारागुवे दैज्लां सम", "PYT": "पारागुवे थाखोआरि सम", "PYT Korea": "प्यंयां सम", "RET": "रियूनियन सम", "ROTT": "रथेरा सम", "SAKST": "साखालिन दैज्लां सम", "SAKT": "साखालिन थाखोआरि सम", "SAMST": "समारा समर टाईम", "SAMT": "समारा स्टैंडर्ड टाईम", "SAST": "खोला आफ्रिकानि थाखोआरि सम", "SBT": "सल’मन द्वीपफोरनि सम", "SCT": "सेशेल्स सम", "SGT": "सींगापुर स्टैंडर्ड टाईम", "SLST": "लंका स्टैंडर्ड टाईम", "SRT": "सुरीनाम सम", "SST Samoa": "साम’वा थाखोआरि सम", "SST Samoa Apia": "आपिया थाखोआरि सम", "SST Samoa Apia DST": "आपिया सान सम", "SST Samoa DST": "साम’वा सानारि सम", "SYOT": "सीअवा सम", "TAAF": "फ्रेन्च खोलायारि आरो एन्टार्कटिक सम", "TAHT": "टाहिटी सम", "TJT": "ताजिकिस्तान सम", "TKT": "टकेलौ सम", "TLT": "ईस्ट टीमोर स्टैंडर्ड टाईम", "TMST": "तुर्कमेनीस्तान दैज्लां सम", "TMT": "तुर्कमेनीस्तान थाखोआरि सम", "TOST": "टंगा दैज्लां सम", "TOT": "टंगा थाखोआरि सम", "TVT": "तुभालु सम", "TWT": "ताईपेइ थाखोआरि सम", "TWT DST": "ताईपेइ सानारि सम", "ULAST": "उलानबातार दैज्लां सम", "ULAT": "उलानबातार थाखोआरि सम", "UYST": "उरुगुवे दैज्लां सम", "UYT": "उरुगुवे थाखोआरि सम", "UZT": "उजबेकिस्तान थाखोआरि सम", "UZT DST": "उजबेकिस्तान दैज्लां सम", "VET": "भेनेजुवेला सम", "VLAST": "भ्लादिभस्त’क दैज्लां सम", "VLAT": "भ्लादिभस्त’क थाखोआरि सम", "VOLST": "भल्ग’ग्रेद दैज्लां सम", "VOLT": "भल्ग’ग्रेद थाखोआरि सम", "VOST": "भस्त’क सम", "VUT": "वानुआटु थाखोआरि सम", "VUT DST": "वानुआटु दैज्लां सम", "WAKT": "वैक द्वीप सम", "WARST": "सोनाब आर्जेन्टिना दैज्लां सम", "WART": "सोनाब आर्जेन्टिना थाखोआरि सम", "WAST": "सोनाब आफ्रिकानि सम", "WAT": "सोनाब आफ्रिकानि सम", "WESZ": "सोनाब युर’पनि दैज्लां सम", "WEZ": "सोनाब युर’पनि थाखोआरि सम", "WFT": "वालीस आरो फुतुना सम", "WGST": "सेनाब ग्रीनलेण्ड दैज्लां सम", "WGT": "सोनाब ग्रीनलेण्ड थाखोआरि सम", "WIB": "वेस्टर्न ईंडोनीशिया स्टैंडर्ड टाईम", "WIT": "ईस्टर्न ईंडोनीशिया स्टैंडर्ड टाईम", "WITA": "ईंडोनीशिया स्टैंडर्ड टाईम", "YAKST": "याकुत्स्क दैज्लां सम", "YAKT": "याकुत्स्क थाखोआरि सम", "YEKST": "येकातेरीनाबुर्ग दैज्लां सम", "YEKT": "येकातेरीनाबुर्ग थाखोआरि सम", "YST": "यूकन सम", "МСК": "मस्कौ थाखोआरि सम", "اقتاۋ": "अक़्टाऊ स्टैंडर्ड टाईम", "اقتاۋ قالاسى": "अक़्टाऊ समर टाईम", "اقتوبە": "अक़्टोबे स्टैंडर्ड टाईम", "اقتوبە قالاسى": "अक़्टोबे समर टाईम", "الماتى": "अलमाटी स्टैंडर्ड टाईम", "الماتى قالاسى": "अलमाटी समर टाईम", "باتىس قازاق ەلى": "सोनाब काजाखस्तान सम", "شىعىش قازاق ەلى": "सानजा काजाखस्तान सम", "قازاق ەلى": "قازاق ەلى", "قىرعىزستان": "कीर्गीस्तान सम", "قىزىلوردا": "क़ीज़ीलोर्डा स्टैंडर्ड टाईम", "قىزىلوردا قالاسى": "क़ीज़ीलोर्डा समर टाईम", "∅∅∅": "आज’र्सनि दैज्लां सम"},
	}
}

// Locale returns the current translators string locale
func (brx *brx_IN) Locale() string {
	return brx.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'brx_IN'
func (brx *brx_IN) PluralsCardinal() []locales.PluralRule {
	return brx.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'brx_IN'
func (brx *brx_IN) PluralsOrdinal() []locales.PluralRule {
	return brx.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'brx_IN'
func (brx *brx_IN) PluralsRange() []locales.PluralRule {
	return brx.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'brx_IN'
func (brx *brx_IN) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'brx_IN'
func (brx *brx_IN) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'brx_IN'
func (brx *brx_IN) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (brx *brx_IN) MonthAbbreviated(month time.Month) string {
	if len(brx.monthsAbbreviated) == 0 {
		return ""
	}
	return brx.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (brx *brx_IN) MonthsAbbreviated() []string {
	return brx.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (brx *brx_IN) MonthNarrow(month time.Month) string {
	if len(brx.monthsNarrow) == 0 {
		return ""
	}
	return brx.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (brx *brx_IN) MonthsNarrow() []string {
	return brx.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (brx *brx_IN) MonthWide(month time.Month) string {
	if len(brx.monthsWide) == 0 {
		return ""
	}
	return brx.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (brx *brx_IN) MonthsWide() []string {
	return brx.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (brx *brx_IN) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(brx.daysAbbreviated) == 0 {
		return ""
	}
	return brx.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (brx *brx_IN) WeekdaysAbbreviated() []string {
	return brx.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (brx *brx_IN) WeekdayNarrow(weekday time.Weekday) string {
	if len(brx.daysNarrow) == 0 {
		return ""
	}
	return brx.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (brx *brx_IN) WeekdaysNarrow() []string {
	return brx.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (brx *brx_IN) WeekdayShort(weekday time.Weekday) string {
	if len(brx.daysShort) == 0 {
		return ""
	}
	return brx.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (brx *brx_IN) WeekdaysShort() []string {
	return brx.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (brx *brx_IN) WeekdayWide(weekday time.Weekday) string {
	if len(brx.daysWide) == 0 {
		return ""
	}
	return brx.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (brx *brx_IN) WeekdaysWide() []string {
	return brx.daysWide
}

// Decimal returns the decimal point of number
func (brx *brx_IN) Decimal() string {
	return brx.decimal
}

// Group returns the group of number
func (brx *brx_IN) Group() string {
	return brx.group
}

// Group returns the minus sign of number
func (brx *brx_IN) Minus() string {
	return brx.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'brx_IN' and handles both Whole and Real numbers based on 'v'
func (brx *brx_IN) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, brx.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, brx.group[0])
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
		b = append(b, brx.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'brx_IN' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (brx *brx_IN) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, brx.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, brx.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, brx.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'brx_IN'
func (brx *brx_IN) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := brx.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, brx.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, brx.group[0])
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

	for j := len(brx.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, brx.currencyPositivePrefix[j])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		b = append(b, brx.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, brx.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'brx_IN'
// in accounting notation.
func (brx *brx_IN) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := brx.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	inSecondary := false
	groupThreshold := 3

	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, brx.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {

			if count == groupThreshold {
				b = append(b, brx.group[0])
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

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		b = append(b, brx.currencyNegativePrefix[0])

	} else {

		for j := len(brx.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, brx.currencyPositivePrefix[j])
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
			b = append(b, brx.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, brx.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'brx_IN'
func (brx *brx_IN) FmtDateShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'brx_IN'
func (brx *brx_IN) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, brx.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'brx_IN'
func (brx *brx_IN) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, brx.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'brx_IN'
func (brx *brx_IN) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, brx.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, brx.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'brx_IN'
func (brx *brx_IN) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, brx.periodsAbbreviated[0]...)
	} else {
		b = append(b, brx.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20, 0xe0, 0xa4, 0xa8, 0xe0, 0xa4, 0xbf, 0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, brx.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'brx_IN'
func (brx *brx_IN) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, brx.periodsAbbreviated[0]...)
	} else {
		b = append(b, brx.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, brx.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, brx.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'brx_IN'
func (brx *brx_IN) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, brx.periodsAbbreviated[0]...)
	} else {
		b = append(b, brx.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, brx.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, brx.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'brx_IN'
func (brx *brx_IN) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 12 {
		b = append(b, brx.periodsAbbreviated[0]...)
	} else {
		b = append(b, brx.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, brx.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, brx.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := brx.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
