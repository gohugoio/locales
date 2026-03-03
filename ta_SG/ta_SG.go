package ta_SG

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type ta_SG struct {
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

// New returns a new instance of translator for the 'ta_SG' locale
func New() locales.Translator {
	return &ta_SG{
		locale:                 "ta_SG",
		pluralsCardinal:        []locales.PluralRule{2, 6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{2, 6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "р.", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "¥", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "RM", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "$", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "฿", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "NT$", "TZS", "UAH", "UAK", "UGS", "UGX", "US$", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyPositivePrefix: " ",
		currencyNegativePrefix: " ",
		monthsAbbreviated:      []string{"", "ஜன.", "பிப்.", "மார்.", "ஏப்.", "மே", "ஜூன்", "ஜூலை", "ஆக.", "செப்.", "அக்.", "நவ.", "டிச."},
		monthsNarrow:           []string{"", "ஜ", "பி", "மா", "ஏ", "மே", "ஜூ", "ஜூ", "ஆ", "செ", "அ", "ந", "டி"},
		monthsWide:             []string{"", "ஜனவரி", "பிப்ரவரி", "மார்ச்", "ஏப்ரல்", "மே", "ஜூன்", "ஜூலை", "ஆகஸ்ட்", "செப்டம்பர்", "அக்டோபர்", "நவம்பர்", "டிசம்பர்"},
		daysAbbreviated:        []string{"ஞாயி.", "திங்.", "செவ்.", "புத.", "வியா.", "வெள்.", "சனி"},
		daysNarrow:             []string{"ஞா", "தி", "செ", "பு", "வி", "வெ", "ச"},
		daysShort:              []string{"ஞா", "தி", "செ", "பு", "வி", "வெ", "ச"},
		daysWide:               []string{"ஞாயிறு", "திங்கள்", "செவ்வாய்", "புதன்", "வியாழன்", "வெள்ளி", "சனி"},
		periodsAbbreviated:     []string{"", ""},
		timezones:              map[string]string{"ACDT": "ஆஸ்திரேலியன் மத்திய பகலொளி நேரம்", "ACST": "ஆஸ்திரேலியன் மத்திய நிலையான நேரம்", "ACT": "அக்ரே தர நேரம்", "ACWDT": "ஆஸ்திரேலியன் மத்திய மேற்கத்திய பகலொளி நேரம்", "ACWST": "ஆஸ்திரேலியன் மத்திய மேற்கத்திய நிலையான நேரம்", "ADT": "அட்லாண்டிக் பகலொளி நேரம்", "ADT Arabia": "அரேபிய பகலொளி நேரம்", "AEDT": "ஆஸ்திரேலியன் கிழக்கத்திய பகலொளி நேரம்", "AEST": "ஆஸ்திரேலியன் கிழக்கத்திய நிலையான நேரம்", "AFT": "ஆஃப்கானிஸ்தான் நேரம்", "AKDT": "அலாஸ்கா பகலொளி நேரம்", "AKST": "அலாஸ்கா நிலையான நேரம்", "AMST": "அமேசான் கோடை நேரம்", "AMST Armenia": "ஆர்மேனிய கோடை நேரம்", "AMT": "அமேசான் நிலையான நேரம்", "AMT Armenia": "ஆர்மேனிய நிலையான நேரம்", "ANAST": "அனாடையர் கோடை நேரம்", "ANAT": "அனாடையர் தர நேரம்", "ARST": "அர்ஜென்டினா கோடை நேரம்", "ART": "அர்ஜென்டினா நிலையான நேரம்", "AST": "அட்லாண்டிக் நிலையான நேரம்", "AST Arabia": "அரேபிய நிலையான நேரம்", "AWDT": "ஆஸ்திரேலியன் மேற்கத்திய பகலொளி நேரம்", "AWST": "ஆஸ்திரேலியன் மேற்கத்திய நிலையான நேரம்", "AZST": "அசர்பைஜான் கோடை நேரம்", "AZT": "அசர்பைஜான் நிலையான நேரம்", "BDT Bangladesh": "வங்கதேச கோடை நேரம்", "BNT": "புருனே டருஸ்ஸலாம் நேரம்", "BOT": "பொலிவியா நேரம்", "BRST": "பிரேசிலியா கோடை நேரம்", "BRT": "பிரேசிலியா நிலையான நேரம்", "BST Bangladesh": "வங்கதேச நிலையான நேரம்", "BT": "பூடான் நேரம்", "CAST": "CAST", "CAT": "மத்திய ஆப்பிரிக்க நேரம்", "CCT": "கோகோஸ் தீவுகள் நேரம்", "CDT": "மத்திய பகலொளி நேரம்", "CHADT": "சத்தாம் பகலொளி நேரம்", "CHAST": "சத்தாம் நிலையான நேரம்", "CHUT": "சுக் நேரம்", "CKT": "குக் தீவுகள் நிலையான நேரம்", "CKT DST": "குக் தீவுகள் அரை கோடை நேரம்", "CLST": "சிலி கோடை நேரம்", "CLT": "சிலி நிலையான நேரம்", "COST": "கொலம்பியா கோடை நேரம்", "COT": "கொலம்பியா நிலையான நேரம்", "CST": "மத்திய நிலையான நேரம்", "CST China": "சீன நிலையான நேரம்", "CST China DST": "சீன பகலொளி நேரம்", "CVST": "கேப் வெர்டே கோடை நேரம்", "CVT": "கேப் வெர்டே நிலையான நேரம்", "CXT": "கிறிஸ்துமஸ் தீவு நேரம்", "ChST": "சாமோரோ நிலையான நேரம்", "ChST NMI": "வடக்கு மரினா தீவுகள் நேரம்", "CuDT": "கியூபா பகலொளி நேரம்", "CuST": "கியூபா நிலையான நேரம்", "DAVT": "டேவிஸ் நேரம்", "DDUT": "டுமோண்ட்-டி உர்வில்லே நேரம்", "EASST": "ஈஸ்டர் தீவு கோடை நேரம்", "EAST": "ஈஸ்டர் தீவு நிலையான நேரம்", "EAT": "கிழக்கு ஆப்பிரிக்க நேரம்", "ECT": "ஈக்வடார் நேரம்", "EDT": "கிழக்கத்திய பகலொளி நேரம்", "EGDT": "கிழக்கு கிரீன்லாந்து கோடை நேரம்", "EGST": "கிழக்கு கிரீன்லாந்து நிலையான நேரம்", "EST": "கிழக்கத்திய நிலையான நேரம்", "FEET": "தூர-கிழக்கு ஐரோப்பிய நேரம்", "FJT": "ஃபிஜி நிலையான நேரம்", "FJT Summer": "ஃபிஜி கோடை நேரம்", "FKST": "ஃபாக்லாந்து தீவுகள் கோடை நேரம்", "FKT": "ஃபாக்லாந்து தீவுகள் நிலையான நேரம்", "FNST": "பெர்னான்டோ டி நோரோன்ஹா கோடை நேரம்", "FNT": "பெர்னான்டோ டி நோரோன்ஹா நிலையான நேரம்", "GALT": "கலபகோஸ் நேரம்", "GAMT": "கேம்பியர் நேரம்", "GEST": "ஜார்ஜியா கோடை நேரம்", "GET": "ஜார்ஜியா நிலையான நேரம்", "GFT": "ஃபிரஞ்சு கயானா நேரம்", "GIT": "கில்பர்ட் தீவுகள் நேரம்", "GMT": "கிரீன்விச் சராசரி நேரம்", "GNSST": "GNSST", "GNST": "GNST", "GST": "தெற்கு ஜார்ஜியா நேரம்", "GST Guam": "கம் தர நேரம்", "GYT": "கயானா நேரம்", "HADT": "ஹவாய்-அலேஷியன் நிலையான நேரம்", "HAST": "ஹவாய்-அலேஷியன் நிலையான நேரம்", "HKST": "ஹாங்காங் கோடை நேரம்", "HKT": "ஹாங்காங் நிலையான நேரம்", "HOVST": "ஹோவ்த் கோடை நேரம்", "HOVT": "ஹோவ்த் நிலையான நேரம்", "ICT": "இந்தோசீன நேரம்", "IDT": "இஸ்ரேல் பகலொளி நேரம்", "IOT": "இந்தியப் பெருங்கடல் நேரம்", "IRKST": "இர்குட்ஸ்க் கோடை நேரம்", "IRKT": "இர்குட்ஸ்க் நிலையான நேரம்", "IRST": "ஈரான் நிலையான நேரம்", "IRST DST": "ஈரான் பகலொளி நேரம்", "IST": "இந்திய நிலையான நேரம்", "IST Israel": "இஸ்ரேல் நிலையான நேரம்", "JDT": "ஜப்பான் பகலொளி நேரம்", "JST": "ஜப்பான் நிலையான நேரம்", "KOST": "கோஸ்ரே நேரம்", "KRAST": "க்ரஸ்னோயார்ஸ்க் கோடை நேரம்", "KRAT": "க்ரஸ்னோயார்ஸ்க் நிலையான நேரம்", "KST": "கொரிய நிலையான நேரம்", "KST DST": "கொரிய பகலொளி நேரம்", "LHDT": "லார்ட் ஹோவ் பகலொளி நேரம்", "LHST": "லார்ட் ஹோவ் நிலையான நேரம்", "LINT": "லைன் தீவுகள் நேரம்", "MAGST": "மகதன் கோடை நேரம்", "MAGT": "மகதன் நிலையான நேரம்", "MART": "மார்கியூசாஸ் நேரம்", "MAWT": "மாசன் நேரம்", "MDT": "மவுன்டைன் பகலொளி நேரம்", "MESZ": "மத்திய ஐரோப்பிய கோடை நேரம்", "MEZ": "மத்திய ஐரோப்பிய நிலையான நேரம்", "MHT": "மார்ஷல் தீவுகள் நேரம்", "MMT": "மியான்மர் நேரம்", "MSD": "மாஸ்கோ கோடை நேரம்", "MST": "மவுன்டைன் நிலையான நேரம்", "MUST": "மொரிஷியஸ் கோடை நேரம்", "MUT": "மொரிஷியஸ் நிலையான நேரம்", "MVT": "மாலத்தீவுகள் நேரம்", "MYT": "மலேஷிய நேரம்", "NCT": "நியூ கலிடோனியா நிலையான நேரம்", "NDT": "நியூஃபவுண்ட்லாந்து பகலொளி நேரம்", "NDT New Caledonia": "நியூ கலிடோனியா கோடை நேரம்", "NFDT": "நார்ஃபோக் தீவு பகலொளி நேரம்", "NFT": "நார்ஃபோக் தீவு நிலையான நேரம்", "NOVST": "நோவோசிபிரிஸ்க் கோடை நேரம்", "NOVT": "நோவோசிபிரிஸ்க் நிலையான நேரம்", "NPT": "நேபாள நேரம்", "NRT": "நவ்ரூ நேரம்", "NST": "நியூஃபவுண்ட்லாந்து நிலையான நேரம்", "NUT": "நியு நேரம்", "NZDT": "நியூசிலாந்து பகலொளி நேரம்", "NZST": "நியூசிலாந்து நிலையான நேரம்", "OESZ": "கிழக்கத்திய ஐரோப்பிய கோடை நேரம்", "OEZ": "கிழக்கத்திய ஐரோப்பிய நிலையான நேரம்", "OMSST": "ஓம்ஸ்க் கோடை நேரம்", "OMST": "ஓம்ஸ்க் நிலையான நேரம்", "PDT": "பசிபிக் பகலொளி நேரம்", "PDTM": "மெக்ஸிகன் பசிபிக் பகலொளி நேரம்", "PETDT": "பெட்ரோபவ்லோவ்ஸ்க் கம்சட்ஸ்கி கோடை நேரம்", "PETST": "பெட்ரோபவ்லோவ்ஸ்க் கம்சட்ஸ்கி தர நேரம்", "PGT": "பபுவா நியூ கினியா நேரம்", "PHOT": "ஃபோனிக்ஸ் தீவுகள் நேரம்", "PKT": "பாகிஸ்தான் நிலையான நேரம்", "PKT DST": "பாகிஸ்தான் கோடை நேரம்", "PMDT": "செயின்ட் பியரி & மிக்குயிலான் பகலொளி நேரம்", "PMST": "செயின்ட் பியரி & மிக்குயிலான் நிலையான நேரம்", "PONT": "போனாபே நேரம்", "PST": "பசிபிக் நிலையான நேரம்", "PST Philippine": "பிலிப்பைன் நிலையான நேரம்", "PST Philippine DST": "பிலிப்பைன் கோடை நேரம்", "PST Pitcairn": "பிட்கெய்ர்ன் நேரம்", "PSTM": "மெக்ஸிகன் பசிபிக் நிலையான நேரம்", "PWT": "பாலவ் நேரம்", "PYST": "பராகுவே கோடை நேரம்", "PYT": "பராகுவே நிலையான நேரம்", "PYT Korea": "பியாங்யாங் நேரம்", "RET": "ரீயூனியன் நேரம்", "ROTT": "ரோதேரா நேரம்", "SAKST": "சகலின் கோடை நேரம்", "SAKT": "சகலின் நிலையான நேரம்", "SAMST": "சமரா கோடை நேரம்", "SAMT": "சமரா தர நேரம்", "SAST": "தென் ஆப்பிரிக்க நிலையான நேரம்", "SBT": "சாலமன் தீவுகள் நேரம்", "SCT": "சீசெல்ஸ் நேரம்", "SGT": "சிங்கப்பூர் நிலையான நேரம்", "SLST": "லங்கா நேரம்", "SRT": "சுரினாம் நேரம்", "SST Samoa": "சமோவா நிலையான நேரம்", "SST Samoa Apia": "ஏபியா நிலையான நேரம்", "SST Samoa Apia DST": "ஏபியா பகலொளி நேரம்", "SST Samoa DST": "சமோவா பகலொளி நேரம்", "SYOT": "ஸ்யோவா நேரம்", "TAAF": "பிரெஞ்சு தெற்கத்திய & அண்டார்டிக் நேரம்", "TAHT": "தஹிதி நேரம்", "TJT": "தஜிகிஸ்தான் நேரம்", "TKT": "டோக்கெலாவ் நேரம்", "TLT": "கிழக்கு திமோர் நேரம்", "TMST": "துர்க்மெனிஸ்தான் கோடை நேரம்", "TMT": "துர்க்மெனிஸ்தான் நிலையான நேரம்", "TOST": "டோங்கா கோடை நேரம்", "TOT": "டோங்கா நிலையான நேரம்", "TVT": "துவாலு நேரம்", "TWT": "தாய்பே நிலையான நேரம்", "TWT DST": "தாய்பே பகலொளி நேரம்", "ULAST": "உலன் பாடர் கோடை நேரம்", "ULAT": "உலன் பாடர் நிலையான நேரம்", "UYST": "உருகுவே கோடை நேரம்", "UYT": "உருகுவே நிலையான நேரம்", "UZT": "உஸ்பெகிஸ்தான் நிலையான நேரம்", "UZT DST": "உஸ்பெகிஸ்தான் கோடை நேரம்", "VET": "வெனிசுலா நேரம்", "VLAST": "விளாடிவோஸ்டோக் கோடை நேரம்", "VLAT": "விளாடிவோஸ்டோக் நிலையான நேரம்", "VOLST": "வோல்கோக்ராட் கோடை நேரம்", "VOLT": "வோல்கோக்ராட் நிலையான நேரம்", "VOST": "வோஸ்டோக் நேரம்", "VUT": "வனுவாட்டு நிலையான நேரம்", "VUT DST": "வனுவாட்டு கோடை நேரம்", "WAKT": "வேக் தீவு நேரம்", "WARST": "மேற்கத்திய அர்ஜென்டினா கோடை நேரம்", "WART": "மேற்கத்திய அர்ஜென்டினா நிலையான நேரம்", "WAST": "மேற்கு ஆப்பிரிக்க நேரம்", "WAT": "மேற்கு ஆப்பிரிக்க நேரம்", "WESZ": "மேற்கத்திய ஐரோப்பிய கோடை நேரம்", "WEZ": "மேற்கத்திய ஐரோப்பிய நிலையான நேரம்", "WFT": "வாலிஸ் மற்றும் ஃப்யூடுனா நேரம்", "WGST": "மேற்கு கிரீன்லாந்து கோடை நேரம்", "WGT": "மேற்கு கிரீன்லாந்து நிலையான நேரம்", "WIB": "மேற்கத்திய இந்தோனேசிய நேரம்", "WIT": "கிழக்கத்திய இந்தோனேசிய நேரம்", "WITA": "மத்திய இந்தோனேசிய நேரம்", "YAKST": "யகுட்ஸ்க் கோடை நேரம்", "YAKT": "யகுட்ஸ்க் நிலையான நேரம்", "YEKST": "யேகாடெரின்பர்க் கோடை நேரம்", "YEKT": "யேகாடெரின்பர்க் நிலையான நேரம்", "YST": "யூகோன் நேரம்", "МСК": "மாஸ்கோ நிலையான நேரம்", "اقتاۋ": "அட்டௌ தர நேரம்", "اقتاۋ قالاسى": "அட்டௌ கோடை நேரம்", "اقتوبە": "அட்டோபே தர நேரம்", "اقتوبە قالاسى": "அட்டோபே கோடை நேரம்", "الماتى": "அல்மாடி தர நேரம்", "الماتى قالاسى": "அல்மாடி கோடை நேரம்", "باتىس قازاق ەلى": "மேற்கு கஜகஸ்தான் நேரம்", "شىعىش قازاق ەلى": "கிழக்கு கஜகஸ்தான் நேரம்", "قازاق ەلى": "கஜகஸ்தான் நேரம்", "قىرعىزستان": "கிர்கிஸ்தான் நேரம்", "قىزىلوردا": "கைஜைலோர்டா தர நேரம்", "قىزىلوردا قالاسى": "கைஜைலோர்டா கோடை நேரம்", "∅∅∅": "அசோர்ஸ் கோடை நேரம்"},
	}
}

// Locale returns the current translators string locale
func (ta *ta_SG) Locale() string {
	return ta.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'ta_SG'
func (ta *ta_SG) PluralsCardinal() []locales.PluralRule {
	return ta.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'ta_SG'
func (ta *ta_SG) PluralsOrdinal() []locales.PluralRule {
	return ta.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'ta_SG'
func (ta *ta_SG) PluralsRange() []locales.PluralRule {
	return ta.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'ta_SG'
func (ta *ta_SG) CardinalPluralRule(num float64, v uint64) locales.PluralRule {

	n := math.Abs(num)

	if n == 1 {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'ta_SG'
func (ta *ta_SG) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'ta_SG'
func (ta *ta_SG) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {

	start := ta.CardinalPluralRule(num1, v1)
	end := ta.CardinalPluralRule(num2, v2)

	if start == locales.PluralRuleOne && end == locales.PluralRuleOther {
		return locales.PluralRuleOther
	} else if start == locales.PluralRuleOther && end == locales.PluralRuleOne {
		return locales.PluralRuleOne
	}

	return locales.PluralRuleOther

}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (ta *ta_SG) MonthAbbreviated(month time.Month) string {
	if len(ta.monthsAbbreviated) == 0 {
		return ""
	}
	return ta.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (ta *ta_SG) MonthsAbbreviated() []string {
	return ta.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (ta *ta_SG) MonthNarrow(month time.Month) string {
	if len(ta.monthsNarrow) == 0 {
		return ""
	}
	return ta.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (ta *ta_SG) MonthsNarrow() []string {
	return ta.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (ta *ta_SG) MonthWide(month time.Month) string {
	if len(ta.monthsWide) == 0 {
		return ""
	}
	return ta.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (ta *ta_SG) MonthsWide() []string {
	return ta.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (ta *ta_SG) WeekdayAbbreviated(weekday time.Weekday) string {
	if len(ta.daysAbbreviated) == 0 {
		return ""
	}
	return ta.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (ta *ta_SG) WeekdaysAbbreviated() []string {
	return ta.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (ta *ta_SG) WeekdayNarrow(weekday time.Weekday) string {
	if len(ta.daysNarrow) == 0 {
		return ""
	}
	return ta.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (ta *ta_SG) WeekdaysNarrow() []string {
	return ta.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (ta *ta_SG) WeekdayShort(weekday time.Weekday) string {
	if len(ta.daysShort) == 0 {
		return ""
	}
	return ta.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (ta *ta_SG) WeekdaysShort() []string {
	return ta.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (ta *ta_SG) WeekdayWide(weekday time.Weekday) string {
	if len(ta.daysWide) == 0 {
		return ""
	}
	return ta.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (ta *ta_SG) WeekdaysWide() []string {
	return ta.daysWide
}

// Decimal returns the decimal point of number
func (ta *ta_SG) Decimal() string {
	return ta.decimal
}

// Group returns the group of number
func (ta *ta_SG) Group() string {
	return ta.group
}

// Group returns the minus sign of number
func (ta *ta_SG) Minus() string {
	return ta.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'ta_SG' and handles both Whole and Real numbers based on 'v'
func (ta *ta_SG) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ta.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ta.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ta.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'ta_SG' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (ta *ta_SG) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ta.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, ta.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, ta.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'ta_SG'
func (ta *ta_SG) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ta.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ta.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ta.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	for j := len(ta.currencyPositivePrefix) - 1; j >= 0; j-- {
		b = append(b, ta.currencyPositivePrefix[j])
	}

	for j := len(symbol) - 1; j >= 0; j-- {
		b = append(b, symbol[j])
	}

	if num < 0 {
		b = append(b, ta.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, ta.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'ta_SG'
// in accounting notation.
func (ta *ta_SG) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := ta.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, ta.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, ta.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {

		for j := len(ta.currencyNegativePrefix) - 1; j >= 0; j-- {
			b = append(b, ta.currencyNegativePrefix[j])
		}

		for j := len(symbol) - 1; j >= 0; j-- {
			b = append(b, symbol[j])
		}

		b = append(b, ta.minus[0])

	} else {

		for j := len(ta.currencyPositivePrefix) - 1; j >= 0; j-- {
			b = append(b, ta.currencyPositivePrefix[j])
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
			b = append(b, ta.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'ta_SG'
func (ta *ta_SG) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'ta_SG'
func (ta *ta_SG) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ta.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'ta_SG'
func (ta *ta_SG) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ta.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'ta_SG'
func (ta *ta_SG) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, ta.daysWide[t.Weekday()]...)
	b = append(b, []byte{0x2c, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, ta.monthsWide[t.Month()]...)
	b = append(b, []byte{0x2c, 0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'ta_SG'
func (ta *ta_SG) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ta.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ta.periodsAbbreviated[0]...)
	} else {
		b = append(b, ta.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'ta_SG'
func (ta *ta_SG) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ta.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ta.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ta.periodsAbbreviated[0]...)
	} else {
		b = append(b, ta.periodsAbbreviated[1]...)
	}

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'ta_SG'
func (ta *ta_SG) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ta.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ta.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ta.periodsAbbreviated[0]...)
	} else {
		b = append(b, ta.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'ta_SG'
func (ta *ta_SG) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	h := t.Hour()

	if h > 12 {
		h -= 12
	}

	b = strconv.AppendInt(b, int64(h), 10)
	b = append(b, ta.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, ta.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20}...)

	if t.Hour() < 12 {
		b = append(b, ta.periodsAbbreviated[0]...)
	} else {
		b = append(b, ta.periodsAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	tz, _ := t.Zone()

	if btz, ok := ta.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
