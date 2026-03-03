package el_POLYTON

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type el_POLYTON struct {
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
	periodsAbbreviated     []string
	timezones              map[string]string
}

// New returns a new instance of translator for the 'el_POLYTON' locale
func New() locales.Translator {
	return &el_POLYTON{
		locale:                 "el_POLYTON",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ",",
		group:                  ".",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "Δρχ", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositiveSuffix: " ",
		currencyNegativeSuffix: " ",
		monthsAbbreviated:      []string{"", "Ιαν", "Φεβ", "Μαρ", "Απρ", "Μαΐ", "Ιουν", "Ιουλ", "Αὐγ", "Σεπ", "Ὀκτ", "Νοε", "Δεκ"},
		monthsNarrow:           []string{"", "Ι", "Φ", "Μ", "Α", "Μ", "Ι", "Ι", "Α", "Σ", "Ο", "Ν", "Δ"},
		monthsWide:             []string{"", "Ιανουαρίου", "Φεβρουαρίου", "Μαρτίου", "Απριλίου", "Μαΐου", "Ιουνίου", "Ιουλίου", "Αὐγούστου", "Σεπτεμβρίου", "Ὀκτωβρίου", "Νοεμβρίου", "Δεκεμβρίου"},
		daysAbbreviated:        []string{"Κυρ", "Δευ", "Τρί", "Τετ", "Πέμ", "Παρ", "Σάβ"},
		daysNarrow:             []string{"Κ", "Δ", "Τ", "Τ", "Π", "Π", "Σ"},
		daysShort:              []string{"Κυ", "Δε", "Τρ", "Τε", "Πέ", "Πα", "Σά"},
		daysWide:               []string{"Κυριακή", "Δευτέρα", "Τρίτη", "Τετάρτη", "Πέμπτη", "Παρασκευή", "Σάββατο"},
		periodsAbbreviated:     []string{"π.μ.", "μ.μ."},
		timezones:              map[string]string{"ACDT": "Θερινή ώρα Κεντρικής Αυστραλίας", "ACST": "Χειμερινή ώρα Κεντρικής Αυστραλίας", "ACT": "Χειμερινή ώρα Άκρε", "ACWDT": "Θερινή ώρα Κεντροδυτικής Αυστραλίας", "ACWST": "Χειμερινή ώρα Κεντροδυτικής Αυστραλίας", "ADT": "Θερινή ώρα Ατλαντικού", "ADT Arabia": "Αραβική θερινή ώρα", "AEDT": "Θερινή ώρα Ανατολικής Αυστραλίας", "AEST": "Χειμερινή ώρα Ανατολικής Αυστραλίας", "AFT": "Ώρα Αφγανιστάν", "AKDT": "Θερινή ώρα Αλάσκας", "AKST": "Χειμερινή ώρα Αλάσκας", "AMST": "Θερινή ώρα Αμαζονίου", "AMST Armenia": "Θερινή ώρα Αρμενίας", "AMT": "Χειμερινή ώρα Αμαζονίου", "AMT Armenia": "Χειμερινή ώρα Αρμενίας", "ANAST": "Θερινή ώρα Αναντίρ", "ANAT": "Χειμερινή ώρα Αναντίρ", "ARST": "Θερινή ώρα Αργεντινής", "ART": "Χειμερινή ώρα Αργεντινής", "AST": "Χειμερινή ώρα Ατλαντικού", "AST Arabia": "Αραβική χειμερινή ώρα", "AWDT": "Θερινή ώρα Δυτικής Αυστραλίας", "AWST": "Χειμερινή ώρα Δυτικής Αυστραλίας", "AZST": "Θερινή ώρα Αζερμπαϊτζάν", "AZT": "Χειμερινή ώρα Αζερμπαϊτζάν", "BDT Bangladesh": "Θερινή ώρα Μπανγκλαντές", "BNT": "Ώρα Μπρουνέι Νταρουσαλάμ", "BOT": "Ώρα Βολιβίας", "BRST": "Θερινή ώρα Μπραζίλιας", "BRT": "Χειμερινή ώρα Μπραζίλιας", "BST Bangladesh": "Χειμερινή ώρα Μπανγκλαντές", "BT": "Ώρα Μπουτάν", "CAST": "CAST", "CAT": "Ώρα Κεντρικής Αφρικής", "CCT": "Ώρα Νήσων Κόκος", "CDT": "Κεντρική θερινή ώρα Βόρειας Αμερικής", "CHADT": "Θερινή ώρα Τσάταμ", "CHAST": "Χειμερινή ώρα Τσάταμ", "CHUT": "Ώρα Τσουκ", "CKT": "Χειμερινή ώρα Νήσων Κουκ", "CKT DST": "Θερινή ώρα Νήσων Κουκ", "CLST": "Θερινή ώρα Χιλής", "CLT": "Χειμερινή ώρα Χιλής", "COST": "Θερινή ώρα Κολομβίας", "COT": "Χειμερινή ώρα Κολομβίας", "CST": "Κεντρική χειμερινή ώρα Βόρειας Αμερικής", "CST China": "Χειμερινή ώρα Κίνας", "CST China DST": "Θερινή ώρα Κίνας", "CVST": "Θερινή ώρα Πράσινου Ακρωτηρίου", "CVT": "Χειμερινή ώρα Πράσινου Ακρωτηρίου", "CXT": "Ώρα Νήσου Χριστουγέννων", "ChST": "Ώρα Τσαμόρο", "ChST NMI": "Ώρα Νησιών Βόρειες Μαριάνες", "CuDT": "Θερινή ώρα Κούβας", "CuST": "Χειμερινή ώρα Κούβας", "DAVT": "Ώρα Ντέιβις", "DDUT": "Ώρα Ντιμόν ντ’ Ουρβίλ", "EASST": "Θερινή ώρα Νήσου Πάσχα", "EAST": "Χειμερινή ώρα Νήσου Πάσχα", "EAT": "Ώρα Ανατολικής Αφρικής", "ECT": "Ώρα Ισημερινού", "EDT": "Ανατολική θερινή ώρα Βόρειας Αμερικής", "EGDT": "Θερινή ώρα Ανατολικής Γροιλανδίας", "EGST": "Χειμερινή ώρα Ανατολικής Γροιλανδίας", "EST": "Ανατολική χειμερινή ώρα Βόρειας Αμερικής", "FEET": "Ώρα περαιτέρω Ανατολικής Ευρώπης", "FJT": "Χειμερινή ώρα Φίτζι", "FJT Summer": "Θερινή ώρα Φίτζι", "FKST": "Θερινή ώρα Νήσων Φόκλαντ", "FKT": "Χειμερινή ώρα Νήσων Φόκλαντ", "FNST": "Θερινή ώρα Φερνάρντο ντε Νορόνια", "FNT": "Χειμερινή ώρα Φερνάρντο ντε Νορόνια", "GALT": "Ώρα Γκαλάπαγκος", "GAMT": "Ώρα Γκάμπιερ", "GEST": "Θερινή ώρα Γεωργίας", "GET": "Χειμερινή ώρα Γεωργίας", "GFT": "Ώρα Γαλλικής Γουιάνας", "GIT": "Ώρα Νήσων Γκίλμπερτ", "GMT": "Μέση ώρα Γκρίνουιτς", "GNSST": "GNSST", "GNST": "GNST", "GST": "Ώρα Νότιας Γεωργίας", "GST Guam": "Ώρα Γκουάμ", "GYT": "Ώρα Γουιάνας", "HADT": "Θερινή ώρα Χαβάης-Αλεούτιων Νήσων", "HAST": "Χειμερινή ώρα Χαβάης-Αλεούτιων Νήσων", "HKST": "Θερινή ώρα Χονγκ Κονγκ", "HKT": "Χειμερινή ώρα Χονγκ Κονγκ", "HOVST": "Θερινή ώρα Χοβντ", "HOVT": "Χειμερινή ώρα Χοβντ", "ICT": "Ώρα Ινδοκίνας", "IDT": "Θερινή ώρα Ισραήλ", "IOT": "Ώρα Ινδικού Ωκεανού", "IRKST": "Θερινή ώρα Ιρκούτσκ", "IRKT": "Χειμερινή ώρα Ιρκούτσκ", "IRST": "Χειμερινή ώρα Ιράν", "IRST DST": "Θερινή ώρα Ιράν", "IST": "Ώρα Ινδίας", "IST Israel": "Χειμερινή ώρα Ισραήλ", "JDT": "Θερινή ώρα Ιαπωνίας", "JST": "Χειμερινή ώρα Ιαπωνίας", "KOST": "Ώρα Κόσραϊ", "KRAST": "Θερινή ώρα Κρασνογιάρσκ", "KRAT": "Χειμερινή ώρα Κρασνογιάρσκ", "KST": "Χειμερινή ώρα Κορέας", "KST DST": "Θερινή ώρα Κορέας", "LHDT": "Θερινή ώρα Λορντ Χάου", "LHST": "Χειμερινή ώρα Λορντ Χάου", "LINT": "Ώρα Νήσων Λάιν", "MAGST": "Θερινή ώρα Μαγκαντάν", "MAGT": "Χειμερινή ώρα Μαγκαντάν", "MART": "Ώρα Μαρκέζας", "MAWT": "Ώρα Μόσον", "MDT": "Θερινή ώρα Μακάο", "MESZ": "Θερινή ώρα Κεντρικής Ευρώπης", "MEZ": "Χειμερινή ώρα Κεντρικής Ευρώπης", "MHT": "Ώρα Νήσων Μάρσαλ", "MMT": "Ώρα Μιανμάρ", "MSD": "Θερινή ώρα Μόσχας", "MST": "Χειμερινή ώρα Μακάο", "MUST": "Θερινή ώρα Μαυρίκιου", "MUT": "Χειμερινή ώρα Μαυρίκιου", "MVT": "Ώρα Μαλδίβων", "MYT": "Ώρα Μαλαισίας", "NCT": "Χειμερινή ώρα Νέας Καληδονίας", "NDT": "Θερινή ώρα Νέας Γης", "NDT New Caledonia": "Θερινή ώρα Νέας Καληδονίας", "NFDT": "Θερινή ώρα Νήσου Νόρφολκ", "NFT": "Χειμερινή ώρα Νήσου Νόρφολκ", "NOVST": "Θερινή ώρα Νοβοσιμπίρσκ", "NOVT": "Χειμερινή ώρα Νοβοσιμπίρσκ", "NPT": "Ώρα Νεπάλ", "NRT": "Ώρα Ναούρου", "NST": "Χειμερινή ώρα Νέας Γης", "NUT": "Ώρα Νιούε", "NZDT": "Θερινή ώρα Νέας Ζηλανδίας", "NZST": "Χειμερινή ώρα Νέας Ζηλανδίας", "OESZ": "Θερινή ώρα Ανατολικής Ευρώπης", "OEZ": "Χειμερινή ώρα Ανατολικής Ευρώπης", "OMSST": "Θερινή ώρα Ομσκ", "OMST": "Χειμερινή ώρα Ομσκ", "PDT": "Θερινή ώρα Ειρηνικού", "PDTM": "Θερινή ώρα Ειρηνικού Μεξικού", "PETDT": "Θερινή ώρα Πετροπαβλόβσκ-Καμτσάτσκι", "PETST": "Χειμερινή ώρα Πετροπαβλόβσκ-Καμτσάτσκι", "PGT": "Ώρα Παπούας Νέας Γουινέας", "PHOT": "Ώρα Νήσων Φοίνιξ", "PKT": "Χειμερινή ώρα Πακιστάν", "PKT DST": "Θερινή ώρα Πακιστάν", "PMDT": "Θερινή ώρα Σεν Πιερ και Μικελόν", "PMST": "Χειμερινή ώρα Σεν Πιερ και Μικελόν", "PONT": "Ώρα Πονάπε", "PST": "Χειμερινή ώρα Ειρηνικού", "PST Philippine": "Χειμερινή ώρα Φιλιππινών", "PST Philippine DST": "Θερινή ώρα Φιλιππινών", "PST Pitcairn": "Ώρα Πίτκερν", "PSTM": "Χειμερινή ώρα Ειρηνικού Μεξικού", "PWT": "Ώρα Παλάου", "PYST": "Θερινή ώρα Παραγουάης", "PYT": "Χειμερινή ώρα Παραγουάης", "PYT Korea": "Ώρα Πιονγιάνγκ", "RET": "Ώρα Ρεϊνιόν", "ROTT": "Ώρα Ρόθερα", "SAKST": "Θερινή ώρα Σαχαλίνης", "SAKT": "Χειμερινή ώρα Σαχαλίνης", "SAMST": "Θερινή ώρα Σαμάρας", "SAMT": "Χειμερινή ώρα Σάμαρας", "SAST": "Χειμερινή ώρα Νότιας Αφρικής", "SBT": "Ώρα Νήσων Σολομώντος", "SCT": "Ώρα Σεϋχελλών", "SGT": "Ώρα Σιγκαπούρης", "SLST": "SLST", "SRT": "Ώρα Σουρινάμ", "SST Samoa": "Χειμερινή ώρα Σαμόα", "SST Samoa Apia": "Χειμερινή ώρα Απία", "SST Samoa Apia DST": "Θερινή ώρα Απία", "SST Samoa DST": "Θερινή ώρα Σαμόα", "SYOT": "Ώρα Σίοβα", "TAAF": "Ώρα Γαλλικού Νότου και Ανταρκτικής", "TAHT": "Ώρα Ταϊτής", "TJT": "Ώρα Τατζικιστάν", "TKT": "Ώρα Τοκελάου", "TLT": "Ώρα Ανατολικού Τιμόρ", "TMST": "Θερινή ώρα Τουρκμενιστάν", "TMT": "Χειμερινή ώρα Τουρκμενιστάν", "TOST": "Θερινή ώρα Τόνγκα", "TOT": "Χειμερινή ώρα Τόνγκα", "TVT": "Ώρα Τουβαλού", "TWT": "Χειμερινή ώρα Ταϊπέι", "TWT DST": "Θερινή ώρα Ταϊπέι", "ULAST": "Θερινή ώρα Ουλάν Μπατόρ", "ULAT": "Χειμερινή ώρα Ουλάν Μπατόρ", "UYST": "Θερινή ώρα Ουρουγουάης", "UYT": "Χειμερινή ώρα Ουρουγουάης", "UZT": "Χειμερινή ώρα Ουζμπεκιστάν", "UZT DST": "Θερινή ώρα Ουζμπεκιστάν", "VET": "Ώρα Βενεζουέλας", "VLAST": "Θερινή ώρα Βλαδιβοστόκ", "VLAT": "Χειμερινή ώρα Βλαδιβοστόκ", "VOLST": "Θερινή ώρα Βόλγκογκραντ", "VOLT": "Χειμερινή ώρα Βόλγκογκραντ", "VOST": "Ώρα Βόστοκ", "VUT": "Χειμερινή ώρα Βανουάτου", "VUT DST": "Θερινή ώρα Βανουάτου", "WAKT": "Ώρα Νήσου Γουέικ", "WARST": "Θερινή ώρα Δυτικής Αργεντινής", "WART": "Χειμερινή ώρα Δυτικής Αργεντινής", "WAST": "Ώρα Δυτικής Αφρικής", "WAT": "Ώρα Δυτικής Αφρικής", "WESZ": "Θερινή ώρα Δυτικής Ευρώπης", "WEZ": "Χειμερινή ώρα Δυτικής Ευρώπης", "WFT": "Ώρα Ουάλις και Φουτούνα", "WGST": "Θερινή ώρα Δυτικής Γροιλανδίας", "WGT": "Χειμερινή ώρα Δυτικής Γροιλανδίας", "WIB": "Ώρα Δυτικής Ινδονησίας", "WIT": "Ώρα Ανατολικής Ινδονησίας", "WITA": "Ώρα Κεντρικής Ινδονησίας", "YAKST": "Θερινή ώρα Γιακούτσκ", "YAKT": "Χειμερινή ώρα Γιακούτσκ", "YEKST": "Θερινή ώρα Αικατερινούπολης", "YEKT": "Χειμερινή ώρα Αικατερινούπολης", "YST": "Ώρα Γιούκον", "МСК": "Χειμερινή ώρα Μόσχας", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "Ώρα Δυτικού Καζακστάν", "شىعىش قازاق ەلى": "Ώρα Ανατολικού Καζακστάν", "قازاق ەلى": "Ώρα Καζακστάν", "قىرعىزستان": "Ώρα Κιργιστάν", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "Θερινή ώρα Περού"},
	}
}

// Locale returns the current translators string locale
func (el *el_POLYTON) Locale() string {
	return el.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'el_POLYTON'
func (el *el_POLYTON) PluralsCardinal() []locales.PluralRule {
	return el.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'el_POLYTON'
func (el *el_POLYTON) PluralsOrdinal() []locales.PluralRule {
	return el.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'el_POLYTON'
func (el *el_POLYTON) PluralsRange() []locales.PluralRule {
	return el.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'el_POLYTON'
func (el *el_POLYTON) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'el_POLYTON'
func (el *el_POLYTON) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'el_POLYTON'
func (el *el_POLYTON) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := el.CardinalPluralRule(num1, v1)
	end := el.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (el *el_POLYTON) MonthAbbreviated(month time.Month) string {
	if len(el.monthsAbbreviated) == 0 {
		return ""
	}
	return el.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (el *el_POLYTON) MonthsAbbreviated() []string {
	return el.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (el *el_POLYTON) MonthNarrow(month time.Month) string {
	if len(el.monthsNarrow) == 0 {
		return ""
	}
	return el.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (el *el_POLYTON) MonthsNarrow() []string {
	return el.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (el *el_POLYTON) MonthWide(month time.Month) string {
	if len(el.monthsWide) == 0 {
		return ""
	}
	return el.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (el *el_POLYTON) MonthsWide() []string {
	return el.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (el *el_POLYTON) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(el.daysAbbreviated) == 0 {
		return ""
	}
	return el.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (el *el_POLYTON) WeekdaysAbbreviated() []string {
	return el.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (el *el_POLYTON) WeekdayNarrow(weekday time.Weekday) string {
	if len(el.daysNarrow) == 0 {
		return ""
	}
	return el.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (el *el_POLYTON) WeekdaysNarrow() []string {
	return el.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (el *el_POLYTON) WeekdayShort(weekday time.Weekday) string {
	if len(el.daysShort) == 0 {
		return ""
	}
	return el.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (el *el_POLYTON) WeekdaysShort() []string {
	return el.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (el *el_POLYTON) WeekdayWide(weekday time.Weekday) string {
	if len(el.daysWide) == 0 {
		return ""
	}
	return el.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (el *el_POLYTON) WeekdaysWide() []string {
	return el.daysWide
}

// Decimal returns the decimal point of number
func (el *el_POLYTON) Decimal() string {
	return el.decimal
}

// Group returns the group of number
func (el *el_POLYTON) Group() string {
	return el.group
}

// Group returns the minus sign of number
func (el *el_POLYTON) Minus() string {
	return el.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'el_POLYTON' and handles both Whole and Real numbers based on 'v'
func (el *el_POLYTON) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, el.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, el.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, el.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'el_POLYTON' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (el *el_POLYTON) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, el.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, el.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, el.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'el_POLYTON'
func (el *el_POLYTON) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := el.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, el.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, el.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, el.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, el.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	b = append(b, el.currencyPositiveSuffix...)

	b = append(b, symbol...)

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'el_POLYTON'
// in accounting notation.
func (el *el_POLYTON) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := el.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, el.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, el.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		b = append(b, el.minus[0])

	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, el.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, el.currencyNegativeSuffix...)
		b = append(b, symbol...)
	} else {

		b = append(b, el.currencyPositiveSuffix...)
		b = append(b, symbol...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'el_POLYTON'
func (el *el_POLYTON) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'el_POLYTON'
func (el *el_POLYTON) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, el.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'el_POLYTON'
func (el *el_POLYTON) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, el.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'el_POLYTON'
func (el *el_POLYTON) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, el.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, el.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'el_POLYTON'
func (el *el_POLYTON) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, el.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, el.periodsAbbreviated[0]...)
	} else {
		b = append(b, el.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'el_POLYTON'
func (el *el_POLYTON) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, el.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, el.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, el.periodsAbbreviated[0]...)
	} else {
		b = append(b, el.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'el_POLYTON'
func (el *el_POLYTON) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, el.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, el.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, el.periodsAbbreviated[0]...)
	} else {
		b = append(b, el.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'el_POLYTON'
func (el *el_POLYTON) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, el.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, el.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0xe2, 0x80, 0xaf}...)

	if t.Hour() < 12 {
		b = append(b, el.periodsAbbreviated[0]...)
	} else {
		b = append(b, el.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := el.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
