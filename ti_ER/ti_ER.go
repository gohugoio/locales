package ti_ER

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ti_ER struct {
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

// New returns a new instance of translator for the 'ti_ER' locale
func New() locales.Translator {
	return &ti_ER{
		locale:             "ti_ER",
		pluralsCardinal:    []locales.PluralRule{2, 6},
		pluralsOrdinal:     nil,
		pluralsRange:       nil,
		timeSeparator:      ":",
		currencies:         []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "Nfk", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		monthsAbbreviated:  []string{"", "ጥሪ", "ለካ", "መጋ", "ሚያ", "ግን", "ሰነ", "ሓም", "ነሓ", "መስ", "ጥቅ", "ሕዳ", "ታሕ"},
		monthsNarrow:       []string{"", "ጥ", "ለ", "መ", "ሚ", "ግ", "ሰ", "ሓ", "ነ", "መ", "ጥ", "ሕ", "ታ"},
		monthsWide:         []string{"", "ጥሪ", "ለካቲት", "መጋቢት", "ሚያዝያ", "ጉንበት", "ሰነ", "ሓምለ", "ነሓሰ", "መስከረም", "ጥቅምቲ", "ሕዳር", "ታሕሳስ"},
		daysAbbreviated:    []string{"ሰን", "ሰኑ", "ሰሉ", "ረቡ", "ሓሙ", "ዓር", "ቀዳ"},
		daysNarrow:         []string{"ሰ", "ሰ", "ሰ", "ረ", "ሓ", "ዓ", "ቀ"},
		daysWide:           []string{"ሰንበት", "ሰኑይ", "ሰሉስ", "ረቡዕ", "ሓሙስ", "ዓርቢ", "ቀዳም"},
		periodsAbbreviated: []string{"ቅ.ቀ.", "ድ.ቀ."},
		erasAbbreviated:    []string{"", ""},
		erasNarrow:         []string{"", ""},
		erasWide:           []string{"", ""},
		timezones:          map[string]string{"ACDT": "ናይ ማዕከላይ መዓልቲ አውስራሊያ ግዘ", "ACST": "ናይ ማዕከላይ መደበኛ አውስራሊያ ግዘ", "ACT": "ናይ መደበኛ ግዘ ኣክሪ", "ACWDT": "ናይ ምዕራባዊ መዓልቲ አውስራሊያ ግዘ", "ACWST": "ናይ ምዕራባዊ መደበኛ አውስራሊያ ግዘ", "ADT": "ናይ መዓልቲ ግዘ አትላንቲክ", "ADT Arabia": "ናይ መዓልቲ አረብ ግዘ", "AEDT": "ናይ ምብራቓዊ መዓልቲ ኣውስትራልያ ግዘ", "AEST": "ናይ ምብራቓዊ መደበኛ ኣውስትራልያ ግዘ", "AFT": "ናይ አፍጋኒስታን ግዘ", "AKDT": "ናይ መዓልቲ ግዘ አላስካ", "AKST": "ናይ መደበኛ ግዘ አላስካ", "AMST": "ግዜ ክረምቲ ኣማዞን", "AMST Armenia": "ናይ ክረምቲ አርሜኒያ ግዘ", "AMT": "ናይ መደበኛ ግዘ ኣማዞን", "AMT Armenia": "ናይ መደበኛ አርሜኒያ ግዘ", "ANAST": "ANAST", "ANAT": "ANAT", "ARST": "ግዜ ክረምቲ ኣርጀንቲና", "ART": "ናይ መደበኛ ግዘ ኣርጀንቲና", "AST": "ናይ መደበኛ ግዘ አትላንቲክ", "AST Arabia": "ናይ መደበኛ አረብ ግዘ", "AWDT": "ናይ ምዕራባዊ መዓልቲ አውስትራሊያ ግዘ", "AWST": "ናይ ምዕራባዊ መደበኛ አውስትራሊያ ግዘ", "AZST": "ናይ ክረምቲ አዘርባዣን ግዘ", "AZT": "ናይ መደበኛ አዘርባዣን ግዘ", "BDT Bangladesh": "ናይ ክረምቲ ባንግላዲሽ ግዘ", "BNT": "ናይ ብሩኔ ዳሩሳሌም ግዘ", "BOT": "ግዜ ቦሊቭያ", "BRST": "ግዜ ክረምቲ ብራዚልያ", "BRT": "ናይ መደበኛ ግዘ ብራዚልያ", "BST Bangladesh": "ናይ መደበኛ ባንግላዲሽ ግዘ", "BT": "ናይ ቡህታን ግዘ", "CAST": "CAST", "CAT": "ግዜ ማእከላይ ኣፍሪቃ", "CCT": "ናይ ኮኮስ ደሴት ግዘ", "CDT": "ናይ መዓልቲ ግዘ ማእከላይ አመሪካ", "CHADT": "ናይ መዓልቲ ቻትሃም ግዘ", "CHAST": "ናይ መደበኛ ቻትሃም ግዘ", "CHUT": "ናይ ቹክ ግዘ", "CKT": "ናይ መደበኛ ኩክ ደሴት ግዘ", "CKT DST": "ናይ ፍርቂ ክረምቲ ኩክ ደሴት ግዘ", "CLST": "ግዜ ክረምቲ ቺሌ", "CLT": "ናይ መደበኛ ግዘ ቺሌ", "COST": "ግዜ ክረምቲ ኮሎምብያ", "COT": "ናይ መደበኛ ግዘ ኮሎምብያ", "CST": "ናይ መደበኛ ግዘ ማእከላይ አመሪካ", "CST China": "ናይ መደበኛ ቻይና ግዘ", "CST China DST": "ናይ መዓልቲ ቻይና ግዘ", "CVST": "ግዜ ክረምቲ ኬፕ ቨርደ", "CVT": "ናይ መደበኛ ግዘ ኬፕ ቨርደ", "CXT": "ናይ ልደት ደሴት ግዘ", "ChST": "ናይ መደበኛ ቻሞሮ ግዘ", "ChST NMI": "ChST NMI", "CuDT": "ናይ መዓልቲ ግዘ ኩባ", "CuST": "ናይ መደበኛ ግዘ ኩባ", "DAVT": "ናይ ዴቪስ ግዘ", "DDUT": "ናይ ዱሞ-ዱርቪል ግዘ", "EASST": "ግዜ ክረምቲ ደሴት ፋሲካ", "EAST": "ናይ መደበኛ ግዘ ደሴት ፋሲካ", "EAT": "ግዜ ምብራቕ ኣፍሪቃ", "ECT": "ግዜ ኤኳዶር", "EDT": "ናይ መዓልቲ ግዘ ምብራቓዊ አመሪካ", "EGDT": "ናይ መዓልቲ ምብራቓዊ ግዘ ግሪንላንድ", "EGST": "ናይ መደበኛ ምብራቓዊ ግዘ ግሪንላንድ", "EST": "ናይ መደበኛ ግዘ ምብራቓዊ ኣመሪካ", "FEET": "ናይ ርሑቕ-ምብራቕ ኤውሮጳዊ ግዘ", "FJT": "ናይ መደበኛ ፊጂ ግዘ", "FJT Summer": "ናይ ክረምቲ ፊጂ ግዘ", "FKST": "ግዜ ከረምቲ ደሴታት ፎክላንድ", "FKT": "ናይ መደበኛ ግዘ ደሴታት ፎክላንድ", "FNST": "ግዜ ክረምቲ ፈርናንዶ ደ ኖሮንያ", "FNT": "ናይ መደበኛ ግዘ ፈርናንዶ ደ ኖሮንያ", "GALT": "ግዜ ጋላፓጎስ", "GAMT": "ናይ ጋምቢየር ግዘ", "GEST": "ናይ ክረምቲ ጆርጂያ ግዘ", "GET": "ናይ መደበኛ ጆርጂያ ግዘ", "GFT": "ግዜ ፈረንሳዊት ጊያና", "GIT": "ናይ ጊልበርት ደሴታት ግዘ", "GMT": "GMT", "GNSST": "GNSST", "GNST": "GNST", "GST": "ግዜ ደቡብ ጆርጅያ", "GST Guam": "GST Guam", "GYT": "ግዜ ጉያና", "HADT": "ናይ መዓልቲ ሃዋይ-ኣሌውቲያን ግዘ", "HAST": "ናይ መደበኛ ሃዋይ-ኣሌውቲያን ግዘ", "HKST": "ናይ ክረምቲ ሆንግ ኮንግ ግዘ", "HKT": "ናይ መደበኛ ሆንግ ኮንግ ግዘ", "HOVST": "ናይ ክረምቲ ሆቭድ ግዘ", "HOVT": "ናይ መደበኛ ሆቭድ ግዘ", "ICT": "ናይ ኢንዶቻይና ግዘ", "IDT": "ናይ መዓልቲ እስራኤል ግዘ", "IOT": "ግዜ ህንዳዊ ውቅያኖስ", "IRKST": "ናይ ክረምቲ ኢርኩትስክ ግዘ", "IRKT": "ናይ መደበኛ ኢርኩትስክ ግዘ", "IRST": "ናይ መደበኛ ኢራን ግዘ", "IRST DST": "ናይ መዓልቲ ኢራን ግዘ", "IST": "ናይ መደበኛ ህንድ ግዘ", "IST Israel": "ናይ መደበኛ እስራኤል ግዘ", "JDT": "ናይ መዓልቲ ጃፓን ግዘ", "JST": "ናይ መደበኛ ጃፓን ግዘ", "KOST": "ናይ ኮርሳይ ግዘ", "KRAST": "ናይ ክረምቲ ክራንስኖያርክ ግዘ", "KRAT": "ናይ መደበኛ ክራንስኖያርክ ግዘ", "KST": "ናይ መደበኛ ኮሪያን ግዘ", "KST DST": "ናይ መዓልቲ ኮሪያን ግዘ", "LHDT": "ናይ መዓልቲ ሎርድ ሆው ግዘ", "LHST": "ናይ መድበኛ ሎርድ ሆው ግዘ", "LINT": "ናይ ላይን ደሴታት ግዘ", "MAGST": "ናይ ክረምቲ ሜጋዳን ግዘ", "MAGT": "ናይ መደበኛ ሜጋዳን ግዘ", "MART": "ናይ ማርኩዌሳስ ግዘ", "MAWT": "ናይ ማውሶን ግዘ", "MDT": "MDT", "MESZ": "ግዜ ክረምቲ ኤውሮጳ", "MEZ": "ናይ መደበኛ ግዘ ማእከላይ ኤውሮጳ", "MHT": "ናይ ማርሻል ደሴታት ግዘ", "MMT": "ናይ ምያንማር ግዘ", "MSD": "ናይ ክረምቲ ሞስኮው ግዘ", "MST": "MST", "MUST": "ግዜ ክረምቲ ማውሪሸስ", "MUT": "ናይ መደበኛ ግዘ ማውሪሸስ", "MVT": "ናይ ሞልዲቭስ ግዘ", "MYT": "ናይ ማሌዢያ ግዘ", "NCT": "ናይ መደበኛ ኒው ካሌዶኒያ ግዘ", "NDT": "ናይ መዓልቲ ኒውፋውንድላንድ ግዘ", "NDT New Caledonia": "ናይ ክረምቲ ኒው ካሌዶኒያ ግዘ", "NFDT": "ናይ መዓልቲ ኖርፎልክ ደሴት ግዘ", "NFT": "ናይ መደበኛ ኖርፎልክ ደሴት ግዘ", "NOVST": "ናይ ክረምቲ ኖቮሲሪስክ ግዘ", "NOVT": "ናይ መደበኛ ኖቮሲሪስክ ግዘ", "NPT": "ናይ ኔፓል ግዘ", "NRT": "ናይ ናውሩ ግዘ", "NST": "ናይ መደበኛ ኒውፋውንድላንድ ግዘ", "NUT": "ናይ ኒዌ ግዘ", "NZDT": "ናይ መዓልቲ ኒው ዚላንድ ግዘ", "NZST": "ናይ መደበኛ ኒው ዚላንድ ግዘ", "OESZ": "ግዜ ክረምቲ ምብራቕ ኤውሮጳ", "OEZ": "ናይ መደበኛ ግዘ ምብራቕ ኤውሮጳ", "OMSST": "ናይ ክረምቲ ኦምስክ ግዘ", "OMST": "ናይ መደበኛ ኦምስክ ግዘ", "PDT": "ናይ መዓልቲ ግዘ ፓስፊክ", "PDTM": "ናይ መዓልቲ ሜክሲኮ ፓስፊክ ግዘ", "PETDT": "PETDT", "PETST": "PETST", "PGT": "ናይ ፓፗ ኒው ጊኒ ግዘ", "PHOT": "ናይ ፊኒክስ ደሴታት ግዘ", "PKT": "ናይ መደበኛ ፓኪስታን ግዘ", "PKT DST": "ናይ ክረምቲ ፓኪስታን ግዘ", "PMDT": "ናይ መዓልቲ ቅዱስ ፒየርን ሚከሎን ግዘ", "PMST": "ናይ መደበኛ ቅዱስ ፒየርን ሚከሎን ግዘ", "PONT": "ናይ ፖናፔ ግዘ", "PST": "ናይ መደበኛ ግዘ ፓስፊክ", "PST Philippine": "ናይ መደበኛ ፊሊፒን ግዘ", "PST Philippine DST": "ናይ ክረምቲ ፊሊፒን ግዘ", "PST Pitcairn": "ናይ ፒትቻይርን ግዘ", "PSTM": "ናይ መደበኛ ሜክሲኮ ፓስፊክ ግዘ", "PWT": "ናይ ፓላው ግዘ", "PYST": "ግዜ ክረምቲ ፓራጓይ", "PYT": "ናይ መደበኛ ግዘ ፓራጓይ", "PYT Korea": "ናይ ፕዮንግያንግ ግዘ", "RET": "ግዜ ርዩንየን", "ROTT": "ናይ ሮቴራ ግዘ", "SAKST": "ናይ ክረምቲ ሳክሃሊን ግዘ", "SAKT": "ናይ መደበኛ ሳክሃሊን ግዘ", "SAMST": "SAMST", "SAMT": "SAMT", "SAST": "ግዜ ደቡብ ኣፍሪቃ", "SBT": "ናይ ሶሎሞን ደሴታት ግዘ", "SCT": "ግዜ ሲሸልስ", "SGT": "ናይ መደበኛ ሲጋፖር ግዘ", "SLST": "SLST", "SRT": "ግዜ ሱሪናም", "SST Samoa": "ናይ መደበኛ ሳሞዋ ግዘ", "SST Samoa Apia": "ናይ መደበኛ አፒያ ግዘ", "SST Samoa Apia DST": "ናይ መዓልቲ አፒያ ግዘ", "SST Samoa DST": "ናይ መዓልቲ ሳሞዋ ግዘ", "SYOT": "ናይ ስዮዋ ግዘ", "TAAF": "ናይ ደቡባዊን ኣንታርቲክ ግዘ", "TAHT": "ናይ ቲሂቲ ግዘ", "TJT": "ናይ ታጃክስታን ግዘ", "TKT": "ናይ ቶኬላው ግዘ", "TLT": "ናይ ምብራቅ ቲሞር ግዘ", "TMST": "ናይ ክረምቲ ቱርክሜኒስታን ግዘ", "TMT": "ናይ መደበኛ ቱርክሜኒስታን ግዘ", "TOST": "ናይ ክረምቲ ቶንጋ ግዘ", "TOT": "ናይ መደበኛ ቶንጋ ግዘ", "TVT": "ናይ ቱቫሉ ግዘ", "TWT": "ናይ መደበኛ ቴፒ ግዘ", "TWT DST": "ናይ መዓልቲ ቴፒ ግዘ", "ULAST": "ናይ ክረምቲ ኡላንባታር ግዘ", "ULAT": "ናይ መደበኛ ኡላንባታር ግዘ", "UYST": "ግዜ ክረምቲ ኡራጓይ", "UYT": "ናይ መደበኛ ግዘ ኡራጓይ", "UZT": "ናይ መደበኛ ኡዝቤኪስታን ግዘ", "UZT DST": "ናይ ክረምቲ ኡዝቤኪስታን ግዘ", "VET": "ግዜ ቬኔዝዌላ", "VLAST": "ናይ ክረምቲ ቭላዲቮስቶክ ግዘ", "VLAT": "ናይ መደበኛ ቭላዲቮስቶክ ግዘ", "VOLST": "ናይ ክረምቲ ቮልጎግራድ ግዘ", "VOLT": "ናይ መደበኛ ቮልጎግራድ ግዘ", "VOST": "ናይ ቮስቶክ ግዘ", "VUT": "ናይ መደበኛ ቫኗታው ግዘ", "VUT DST": "ናይ ክረምቲ ቫኗታው ግዘ", "WAKT": "ናይ ዌክ ደሴት ግዘ", "WARST": "ግዜ ክረምቲ ምዕራባዊ ኣርጀንቲና", "WART": "ናይ መደበኛ ግዘ ምዕራባዊ ኣርጀንቲና", "WAST": "ግዜ ምዕራብ ኣፍሪቃ", "WAT": "ግዜ ምዕራብ ኣፍሪቃ", "WESZ": "ናይ ክረምቲ ምዕራባዊ ኤውሮጳዊ ግዘ", "WEZ": "ናይ መደበኛ ምዕራባዊ ኤውሮጳዊ ግዘ", "WFT": "ናይ ዌልስን ፉቷ ግዘ", "WGST": "ናይ መዓልቲ ምዕራብ ግሪንላንድ ግዘ", "WGT": "ናይ መደበኛ ምዕራብ ግሪንላንድ ግዘ", "WIB": "ናይ ምዕራባዊ ኢንዶነዥያ ግዘ", "WIT": "ናይ ምብራቓዊ ኢንዶነዥያ ግዘ", "WITA": "ናይ ማእከላይ ኢንዶነዥያ ግዘ", "YAKST": "ናይ ክረምቲ ያኩትስክ ግዘ", "YAKT": "ናይ መደበኛ ያኩትስክ ግዘ", "YEKST": "ናይ ክረምቲ ያክተርኒበርግ ግዘ", "YEKT": "ናይ መደበኛ ያክተርኒበርግ ግዘ", "YST": "ናይ ዩኮን ግዘ", "МСК": "ናይ መደበኛ ሞስኮው ግዘ", "اقتاۋ": "اقتاۋ", "اقتاۋ قالاسى": "اقتاۋ قالاسى", "اقتوبە": "اقتوبە", "اقتوبە قالاسى": "اقتوبە قالاسى", "الماتى": "الماتى", "الماتى قالاسى": "الماتى قالاسى", "باتىس قازاق ەلى": "ናይ ምዕራብ ካዛኪስታን ግዘ", "شىعىش قازاق ەلى": "ናይ ምብራቅ ካዛኪስታን ግዘ", "قازاق ەلى": "ናይ ካዛኪስታን ግዘ", "قىرعىزستان": "ናይ ክርጅስታን ግዘ", "قىزىلوردا": "قىزىلوردا", "قىزىلوردا قالاسى": "قىزىلوردا قالاسى", "∅∅∅": "ናይ ክረምቲ አዞረስ ግዘ"},
	}
}

// Locale returns the current translators string locale
func (ti *ti_ER) Locale() string {
	return ti.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ti_ER'
func (ti *ti_ER) PluralsCardinal() []locales.PluralRule {
	return ti.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ti_ER'
func (ti *ti_ER) PluralsOrdinal() []locales.PluralRule {
	return ti.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ti_ER'
func (ti *ti_ER) PluralsRange() []locales.PluralRule {
	return ti.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ti_ER'
func (ti *ti_ER) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n >= 0 && n <= 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ti_ER'
func (ti *ti_ER) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ti_ER'
func (ti *ti_ER) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleUnknown
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ti *ti_ER) MonthAbbreviated(month time.Month) string {
	return ti.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ti *ti_ER) MonthsAbbreviated() []string {
	return ti.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ti *ti_ER) MonthNarrow(month time.Month) string {
	return ti.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ti *ti_ER) MonthsNarrow() []string {
	return ti.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ti *ti_ER) MonthWide(month time.Month) string {
	return ti.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ti *ti_ER) MonthsWide() []string {
	return ti.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ti *ti_ER) WeekdayAbbreviated(weekday time.Weekday) string {
	return ti.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ti *ti_ER) WeekdaysAbbreviated() []string {
	return ti.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ti *ti_ER) WeekdayNarrow(weekday time.Weekday) string {
	return ti.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ti *ti_ER) WeekdaysNarrow() []string {
	return ti.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ti *ti_ER) WeekdayShort(weekday time.Weekday) string {
	return ti.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ti *ti_ER) WeekdaysShort() []string {
	return ti.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ti *ti_ER) WeekdayWide(weekday time.Weekday) string {
	return ti.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ti *ti_ER) WeekdaysWide() []string {
	return ti.daysWide
}

// Decimal returns the decimal point of number
func (ti *ti_ER) Decimal() string {
	return ti.decimal
}

// Group returns the group of number
func (ti *ti_ER) Group() string {
	return ti.group
}

// Group returns the minus sign of number
func (ti *ti_ER) Minus() string {
	return ti.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ti_ER' and handles both Whole and Real numbers based on 'v'
func (ti *ti_ER) FmtNumber(num float64, v uint64) string {

	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ti_ER' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ti *ti_ER) FmtPercent(num float64, v uint64) string {
	return strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ti_ER'
func (ti *ti_ER) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ti.currencies[currency]
	l := len(s) + len(symbol) + 0
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ti.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ti.group[0])
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
		b = append(b, ti.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ti.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ti_ER'
// in accounting notation.
func (ti *ti_ER) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ti.currencies[currency]
	l := len(s) + len(symbol) + 0
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ti.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ti.group[0])
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

		b = append(b, ti.minus[0])

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
			b = append(b, ti.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ti_ER'
func (ti *ti_ER) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'ti_ER'
func (ti *ti_ER) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ti.monthsAbbreviated[t.Month()]...)
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

// FmtDateLong returns the long date representation of 't' for 'ti_ER'
func (ti *ti_ER) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ti.monthsWide[t.Month()]...)
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

// FmtDateFull returns the full date representation of 't' for 'ti_ER'
func (ti *ti_ER) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ti.daysWide[t.Weekday()]...)
	b = append(b, []byte{0xe1, 0x8d, 0xa3, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ti.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ti_ER'
func (ti *ti_ER) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ti.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ti.periodsAbbreviated[0]...)
	} else {
		b = append(b, ti.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ti_ER'
func (ti *ti_ER) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ti.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ti.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ti.periodsAbbreviated[0]...)
	} else {
		b = append(b, ti.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ti_ER'
func (ti *ti_ER) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ti.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ti.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ti.periodsAbbreviated[0]...)
	} else {
		b = append(b, ti.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ti_ER'
func (ti *ti_ER) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ti.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ti.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ti.periodsAbbreviated[0]...)
	} else {
		b = append(b, ti.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ti.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
