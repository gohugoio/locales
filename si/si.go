package si

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type si struct {
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

// New returns a new instance of translator for the 'si' locale
func New() locales.Translator {
	return &si{
		locale:                 "si",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ".",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "රු.", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "₽", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "සිෆ්එ", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "දුරුතු", "නවම්", "මැදින්", "බක්", "වෙසක්", "පොසොන්", "ඇසළ", "නිකිණි", "බිනර", "වප්", "ඉල්", "උඳුවප්"},
		monthsNarrow:           []string{"", "දු", "න", "මැ", "බ", "වෙ", "පො", "ඇ", "නි", "බි", "ව", "ඉ", "උ"},
		monthsWide:             []string{"", "දුරුතු", "නවම්", "මැදින්", "බක්", "වෙසක්", "පොසොන්", "ඇසළ", "නිකිණි", "බිනර", "වප්", "ඉල්", "උඳුවප්"},
		daysAbbreviated:        []string{"ඉරිදා", "සඳුදා", "අඟහ", "බදාදා", "බ්\u200dරහස්", "සිකු", "සෙන"},
		daysNarrow:             []string{"ඉ", "ස", "අ", "බ", "බ්\u200dර", "සි", "සෙ"},
		daysShort:              []string{"ඉරි", "සඳු", "අඟ", "බදා", "බ්\u200dරහ", "සිකු", "සෙන"},
		daysWide:               []string{"ඉරිදා", "සඳුදා", "අඟහරුවාදා", "බදාදා", "බ්\u200dරහස්පතින්දා", "සිකුරාදා", "සෙනසුරාදා"},
		timezones:              map[string]string{"ACDT": "මධ්\u200dයම ඔස්ට්\u200dරේලියානු දහවල් වේලාව", "ACST": "ACST", "ACT": "ACT", "ACWDT": "මධ්\u200dයම බටහිර ඔස්ට්\u200dරේලියානු දහවල් වේලාව", "ACWST": "මධ්\u200dයම බටහිර ඔස්ට්\u200dරේලියානු සම්මත වේලාව", "ADT": "අත්ලාන්තික් දිවාආලෝක වේලාව", "ADT Arabia": "අරාබි දහවල් වේලාව", "AEDT": "නැඟෙනහිර ඕස්ට්\u200dරේලියානු දහවල් වේලාව", "AEST": "නැගෙනහිර ඕස්ට්\u200dරේලියානු සම්මත වේලාව", "AFT": "ඇෆ්ගනිස්ථාන වේලාව", "AKDT": "ඇලස්කා දිවාආලෝක වේලාව", "AKST": "ඇලස්කා සම්මත වේලාව", "AMST": "ඇමර්සන් ග්\u200dරීෂ්ම කාලය", "AMST Armenia": "ආමේනියානු ග්\u200dරීෂ්ම වේලාව", "AMT": "ඇමර්සන් සම්මත වේලාව", "AMT Armenia": "ආමේනියානු සම්මත වේලාව", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "ආර්ජන්ටිනා ග්\u200dරීෂ්ම කාලය", "ART": "ආර්ජන්ටිනා සම්මත වේලාව", "AST": "අත්ලාන්තික් සම්මත වේලාව", "AST Arabia": "අරාබි සම්මත වේලාව", "AWDT": "බටහිර ඔස්ට්\u200dරේලියානු දහවල් වේලාව", "AWST": "බටහිර ඕස්ට්\u200dරේලියානු සම්මත වේලාව", "AZST": "අසර්බයිජාන් ග්\u200dරීෂ්ම වේලාව", "AZT": "අසර්බයිජාන් සම්මත වේලාව", "BDT Bangladesh": "බංගලාදේශ ග්\u200dරීෂ්ම කාලය", "BNT": "බෘනායි දරුස්සලාම් වේලාව", "BOT": "බොලිවියා වේලාව", "BRST": "බ්\u200dරසීල ග්\u200dරීෂ්ම කාලය", "BRT": "බ්\u200dරසීල සම්මත වේලාව", "BST Bangladesh": "බංගලාදේශ සම්මත වේලාව", "BT": "භුතාන වේලාව", "CAST": "CAST", "CAT": "මධ්\u200dයම අප්\u200dරිකානු වේලාව", "CCT": "කොකෝස් දුපත් වේලාව", "CDT": "උතුරු ඇමරිකානු මධ්\u200dයම දිවාආලෝක වේලාව", "CHADT": "චැතම් දිවා වේලාව", "CHAST": "චැතම් සම්මත වේලාව", "CHUT": "චුක් වේලාව", "CKT": "කුක් දුපත් සම්මත වේලාව", "CKT DST": "කුක් දුපත් භාග ග්\u200dරීෂ්ම වේලාව", "CLST": "චිලී ග්\u200dරීෂ්ම කාලය", "CLT": "චිලී සම්මත වේලාව", "COST": "කොලොම්බියා ග්\u200dරීෂ්ම කාලය", "COT": "කොලොම්බියා සම්මත වේලාව", "CST": "උතුරු ඇමරිකානු මධ්\u200dයම සම්මත වේලාව", "CST China": "චීන සම්මත වේලාව", "CST China DST": "චීන දහවල් වේලාව", "CVST": "කේප්වේඩ් ග්\u200dරීෂ්ම කාලය", "CVT": "කේප්වේඩ් සම්මත වේලාව", "CXT": "ක්\u200dරිස්මස් දුපත් වේලාව", "ChST": "චමොරෝ වේලාව", "ChST NMI": "ChST NMI", "CuDT": "කියුබානු දිවාආලෝක වේලාව", "CuST": "කියුබානු සම්මත වේලාව", "DAVT": "ඩාවිස් වේලාව", "DDUT": "දුමොන්ත්-ඩ්උර්විල් වේලාව", "EASST": "ඊස්ටර් දූපත් ග්\u200dරීෂ්ම කාලය", "EAST": "ඊස්ටර් දූපත් සම්මත වේලාව", "EAT": "නැගෙනහිර අප්\u200dරිකානු වේලාව", "ECT": "ඉක්වදෝර් වේලාව", "EDT": "උතුරු ඇමරිකානු නැගෙනහිර දිවාආලෝක වේලාව", "EGDT": "නැගෙනහිර ග්\u200dරීන්ලන්ත ග්\u200dරීෂ්ම කාලය", "EGST": "නැගෙනහිර ග්\u200dරීන්ලන්ත සම්මත වේලාව", "EST": "උතුරු ඇමරිකානු නැගෙනහිර සම්මත වේලාව", "FEET": "තවත්-නැගෙනහිර යුරෝපීය වේලාව", "FJT": "ෆිජි සම්මත වේලාව", "FJT Summer": "ෆිජි ග්\u200dරීෂ්ම වේලාව", "FKST": "ෆෝක්ලන්ඩ් දූපත් ග්\u200dරීෂ්ම කාලය", "FKT": "ෆෝක්ලන්ඩ් දූපත් සම්මත වේලාව", "FNST": "ෆර්නැන්ඩෝ ඩි නොරොන්හා ග්\u200dරීෂ්ම කාලය", "FNT": "ෆර්නැන්ඩෝ ඩි නොරොන්හා සම්මත වේලාව", "GALT": "ගලපගොස් වේලාව", "GAMT": "ගැම්බියර් වේලාව", "GEST": "ජෝර්ජියානු ග්\u200dරීෂ්ම වේලාව", "GET": "ජෝර්ජියානු සම්මත වේලාව", "GFT": "ප්\u200dරංශ ගයනා වේලාව", "GIT": "ගිල්බර්ට් දුපත් වේලාව", "GMT": "ග්\u200dරිනිච් මධ්\u200dයම වේලාව", "GNSST": "GNSST", "GNST": "GNST", "GST": "ගල්ෆ් වේලාව", "GST Guam": "GST Guam", "GYT": "ගයනා වේලාව", "HADT": "හවායි-අලෙයුතියාන් සම්මත වේලාව", "HAST": "හවායි-අලෙයුතියාන් සම්මත වේලාව", "HKST": "හොංකොං ග්\u200dරීෂ්ම වේලාව", "HKT": "හොංකොං සම්මත වේලාව", "HOVST": "හොව්ඩ් ග්\u200dරීෂ්ම වේලාව", "HOVT": "හොව්ඩ් සම්මත වේලාව", "ICT": "ඉන්දුචීන වේලාව", "IDT": "ඊශ්\u200dරායල දහවල් වේලාව", "IOT": "ඉන්දියන් සාගර වේලාව", "IRKST": "ඉර්කුට්ස්ක් ග්\u200dරීෂ්ම වේලාව", "IRKT": "ඉර්කුට්ස්ක් සම්මත වේලාව", "IRST": "ඉරාන සම්මත වේලාව", "IRST DST": "ඉරාන දිවා කාලය", "IST": "ඉන්දියානු වේලාව", "IST Israel": "ඊශ්\u200dරායල සම්මත වේලාව", "JDT": "ජපාන දහවල් වේලාව", "JST": "ජපාන සම්මත වේලාව", "KOST": "කොස්රේ වේලාව", "KRAST": "ක්\u200dරස්නොයාර්ස්ක් ග්\u200dරීෂ්ම වේලාව", "KRAT": "ක්\u200dරස්නොයාර්ස්ක් සම්මත වේලාව", "KST": "කොරියානු සම්මත වේලාව", "KST DST": "කොරියානු දහවල් වේලාව", "LHDT": "ලෝර්ඩ් හෝව් දිවා වේලාව", "LHST": "ලෝර්ඩ් හෝව් සම්මත වේලාව", "LINT": "ලයින් දුපත් වේලාව", "MAGST": "මෙගඩන් ග්\u200dරීෂ්ම වේලාව", "MAGT": "මෙගඩන් සම්මත වේලාව", "MART": "මාර්කුඑසාස් වේලාව", "MAWT": "මොව්සන් වේලාව", "MDT": "උතුරු ඇමරිකානු කඳුකර දිවාආලෝක වේලාව", "MESZ": "මධ්\u200dයම යුරෝපීය ග්\u200dරීෂ්ම වේලාව", "MEZ": "මධ්\u200dයම යුරෝපීය සම්මත වේලාව", "MHT": "මාර්ෂල් දුපත් වේලාව", "MMT": "මියන්මාර් වේලාව", "MSD": "මොස්කව් ග්\u200dරීෂ්ම වේලාව", "MST": "උතුරු ඇමරිකානු කඳුකර සම්මත වේලාව", "MUST": "මුරුසි ග්\u200dරීෂ්ම කාලය", "MUT": "මුරුසි සම්මත වේලාව", "MVT": "මාලදිවයින් වේලාව", "MYT": "මැලේසියානු වේලාව", "NCT": "නව සෙලඩොනියානු සම්මත වේලාව", "NDT": "නිව්ෆවුන්ලන්ත දිවාආලෝක වේලාව", "NDT New Caledonia": "නව සෙලඩොනියානු ග්\u200dරීෂ්ම වේලාව", "NFDT": "නොෆොල්ක් දුපත් ග්\u200dරීෂ්ම වේලාව", "NFT": "නොෆොල්ක් දුපත් සම්මත වේලාව", "NOVST": "නොවසිබිර්ස්ක් ග්\u200dරීෂ්ම වේලාව", "NOVT": "නොවසිබිර්ස්ක් සම්මත වේලාව", "NPT": "නේපාල වේලාව", "NRT": "නාවුරු වේලාව", "NST": "නිව්ෆවුන්ලන්ත සම්මත වේලාව", "NUT": "නියු වේලාව", "NZDT": "නවසීලන්ත දිවා වේලාව", "NZST": "නවසීලන්ත සම්මත වේලාව", "OESZ": "නැගෙනහිර යුරෝපීය ග්\u200dරීෂ්ම වේලාව", "OEZ": "නැගෙනහිර යුරෝපීය සම්මත වේලාව", "OMSST": "ඔම්ස්ක් ග්\u200dරීෂ්ම වේලාව", "OMST": "ඔම්ස්ක් සම්මත වේලාව", "PDT": "උතුරු ඇමරිකානු පැසිෆික් දිවාආලෝක වේලාව", "PDTM": "මෙක්සිකෝ පැසිෆික් දිවාආලෝක වේලාව", "PETDT": "PETDT", "PETST": "PETST", "PGT": "පැපුවා නිව් ගිනීයා වේලාව", "PHOT": "ෆීනික්ස් දුපත් වේලාව", "PKT": "පාකිස්ථාන සම්මත වේලාව", "PKT DST": "පාකිස්ථාන ග්\u200dරීෂ්ම කාලය", "PMDT": "ශාන්ත පියරේ සහ මැකෝලන් දිවාආලෝක වේලාව", "PMST": "ශාන්ත පියරේ සහ මැකෝලන් සම්මත වේලාව", "PONT": "පොනපේ වේලාව", "PST": "උතුරු ඇමරිකානු පැසිෆික් සම්මත වේලාව", "PST Philippine": "පිලිපීන සම්මත වේලාව", "PST Philippine DST": "පිලිපීන ග්\u200dරීෂ්ම වේලාව", "PST Pitcairn": "පිට්කෙයාන් වේලාව", "PSTM": "මෙක්සිකෝ පැසිෆික් සම්මත වේලාව", "PWT": "පලාවු වේලාව", "PYST": "පැරගුවේ ග්\u200dරීෂ්ම කාලය", "PYT": "පැරගුවේ සම්මත වේලාව", "PYT Korea": "ප්යොන්ග්යන් වේලාව", "RET": "රියුනියන් වේලාව", "ROTT": "රොතෙරා වේලාව", "SAKST": "සඛලින් ග්\u200dරීෂ්ම වේලාව", "SAKT": "සඛලින් සම්මත වේලාව", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "දකුණු අප්\u200dරිකානු වේලාව", "SBT": "සොලොමන් දූපත් වේලාව", "SCT": "සීෂෙල්ස් වේලාව", "SGT": "සිංගප්පුරු වේලාව", "SLST": "ශ්\u200dරී ලංකා වේලාව", "SRT": "සුරිනාම වේලාව", "SST Samoa": "සැමෝවා සම්මත වේලාව", "SST Samoa Apia": "අපියා සම්මත වේලාව", "SST Samoa Apia DST": "අපියා දිවා වේලාව", "SST Samoa DST": "සැමෝවා ග්\u200dරීෂ්ම වේලාව", "SYOT": "ස්යෝවා වේලාව", "TAAF": "ප්\u200dරංශ දකුණුදිග සහ ඇන්ටාර්ක්ටික් වේලාව", "TAHT": "ටාහිටි වේලාව", "TJT": "ටජිකිස්තාන වේලාව", "TKT": "ටොකෙලාවු වේලාව", "TLT": "නැගෙනහිර ටිමෝර් වේලාව", "TMST": "ටර්ක්මෙනිස්තාන ග්\u200dරීෂ්ම වේලාව", "TMT": "ටර්ක්මෙනිස්තාන සම්මත වේලාව", "TOST": "ටොංගා ග්\u200dරීෂ්ම වේලාව", "TOT": "ටොංගා සම්මත වේලාව", "TVT": "ටුවාලු වේලාව", "TWT": "තායිපේ සම්මත වේලාව", "TWT DST": "තායිපේ දහවල් වේලාව", "ULAST": "උලාන් බාටර් ග්\u200dරීෂ්ම වේලාව", "ULAT": "උලාන් බාටර් සම්මත වේලාව", "UYST": "උරුගුවේ ග්\u200dරීෂ්ම කාලය", "UYT": "උරුගුවේ සම්මත වේලාව", "UZT": "උස්බෙකිස්තාන සම්මත වේලාව", "UZT DST": "උස්බෙකිස්තාන ග්\u200dරීෂ්ම වේලාව", "VET": "වෙනිසියුලා වේලාව", "VLAST": "ව්ලදිවෝස්ටෝක් ග්\u200dරීෂ්ම වේලාව", "VLAT": "ව්ලදිවෝස්ටෝක් සම්මත වේලාව", "VOLST": "වොල්ගොග්\u200dරාඩ් ග්\u200dරීෂ්ම වේලාව", "VOLT": "වොල්ගොග්\u200dරාඩ් සම්මත වේලාව", "VOST": "වොස්ටොක් වේලාව", "VUT": "වනුආටු සම්මත වේලාව", "VUT DST": "වනුආටු ගිම්හාන වේලාව", "WAKT": "වේක් දූපත් වේලාව", "WARST": "බටහිර ආර්ජන්ටිනා ග්\u200dරීෂ්ම කාලය", "WART": "බටහිර ආර්ජන්ටිනා සම්මත වේලාව", "WAST": "බටහිර අප්\u200dරිකානු වේලාව", "WAT": "බටහිර අප්\u200dරිකානු වේලාව", "WESZ": "බටහිර යුරෝපීය ග්\u200dරීෂ්ම වේලාව", "WEZ": "බටහිර යුරෝපීය සම්මත වේලාව", "WFT": "වැලිස් සහ ෆුටුනා වේලාව", "WGST": "බටහිර ග්\u200dරීන්ලන්ත ග්\u200dරීෂ්ම කාලය", "WGT": "බටහිර ග්\u200dරීන්ලන්ත සම්මත වේලාව", "WIB": "බටහිර ඉන්දුනීසියානු වේලාව", "WIT": "නැගෙනහිර ඉන්දුනීසියානු වේලාව", "WITA": "මධ්\u200dයම ඉන්දුනීසියානු වේලාව", "YAKST": "යකුට්ස්ක් ග්\u200dරීෂ්ම වේලාව", "YAKT": "යකුට්ස්ක් සම්මත වේලාව", "YEKST": "යෙකටෙරින්බර්ග් ග්\u200dරීෂ්ම වේලාව", "YEKT": "යෙකටෙරින්බර්ග් සම්මත වේලාව", "YST": "යුකොන් වේලාව", "МСК": "මොස්කව් සම්මත වේලාව", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "බටහිර කසකස්තාන වේලාව", "شىعىش قازاق ەلى": "නැගෙනහිර කසකස්තාන වේලාව", "قازاق ەلى": "කසකස්තාන වේලාව", "قىرعىزستان": "කිර්ගිස්තාන වේලාව", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "ඇසොර්ස් ග්\u200dරීෂ්ම වේලාව"},
	}
}

// Locale returns the current translators string locale
func (si *si) Locale() string {
	return si.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'si'
func (si *si) PluralsCardinal() []locales.PluralRule {
	return si.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'si'
func (si *si) PluralsOrdinal() []locales.PluralRule {
	return si.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'si'
func (si *si) PluralsRange() []locales.PluralRule {
	return si.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'si'
func (si *si) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	n := math.Abs(num)
	i := int64(n)
	f := locales.F(n, v)

	if (n == 0 || n == 1) || (i == 0 && f == 1) {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'si'
func (si *si) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'si'
func (si *si) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	start := si.CardinalPluralRule(num1, v1)
	end := si.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	} else if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOther
	}

	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (si *si) MonthAbbreviated(month time.Month) string {
	return si.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (si *si) MonthsAbbreviated() []string {
	return si.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (si *si) MonthNarrow(month time.Month) string {
	return si.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (si *si) MonthsNarrow() []string {
	return si.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (si *si) MonthWide(month time.Month) string {
	return si.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (si *si) MonthsWide() []string {
	return si.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (si *si) WeekdayAbbreviated(weekday time.Weekday) string {
	return si.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (si *si) WeekdaysAbbreviated() []string {
	return si.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (si *si) WeekdayNarrow(weekday time.Weekday) string {
	return si.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (si *si) WeekdaysNarrow() []string {
	return si.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (si *si) WeekdayShort(weekday time.Weekday) string {
	return si.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (si *si) WeekdaysShort() []string {
	return si.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (si *si) WeekdayWide(weekday time.Weekday) string {
	return si.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (si *si) WeekdaysWide() []string {
	return si.daysWide
}

// Decimal returns the decimal point of number
func (si *si) Decimal() string {
	return si.decimal
}

// Group returns the group of number
func (si *si) Group() string {
	return si.group
}

// Group returns the minus sign of number
func (si *si) Minus() string {
	return si.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'si' and handles both Whole and Real numbers based on 'v'
func (si *si) FmtNumber(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, si.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, si.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, si.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'si' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (si *si) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, si.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, si.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, si.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'si'
func (si *si) FmtCurrency(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := si.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, si.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, si.group[0])
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

	if num < 0 {
		b = append(b, si.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, si.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'si'
// in accounting notation.
func (si *si) FmtAccounting(num float64, v uint64, currency currency.Type) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := si.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, si.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, si.group[0])
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

		b = append(b, si.currencyNegativePrefix[0])

	} else {
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
			b = append(b, si.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, si.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'si'
func (si *si) FmtDateShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x2d}...)

	if t.Month() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Month()), 10)

	b = append(b, []byte{0x2d}...)

	if t.Day() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateMedium returns the medium date representation of 't' for 'si'
func (si *si) FmtDateMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, si.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'si'
func (si *si) FmtDateLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, si.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'si'
func (si *si) FmtDateFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	b = append(b, []byte{0x20}...)
	b = append(b, si.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x2c, 0x20}...)
	b = append(b, si.daysWide[t.Weekday()]...)

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'si'
func (si *si) FmtTimeShort(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'si'
func (si *si) FmtTimeMedium(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'si'
func (si *si) FmtTimeLong(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'si'
func (si *si) FmtTimeFull(t time.Time) string {
	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x2e}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := si.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
