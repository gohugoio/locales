package th_TH

import (
	"math"
	"strconv"
	"time"

	"github.com/gohugoio/locales"
	"github.com/gohugoio/locales/currency"
)

type th_TH struct {
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

// New returns a new instance of translator for the 'th_TH' locale
func New() locales.Translator {
	return &th_TH{
		locale:                 "th_TH",
		pluralsCardinal:        []locales.PluralRule{6},
		pluralsOrdinal:         []locales.PluralRule{6},
		pluralsRange:           []locales.PluralRule{6},
		decimal:                ".",
		group:                  ",",
		minus:                  "-",
		percent:                "%",
		timeSeparator:          ":",
		currencies:             []string{"ADP", "AED", "AFA", "AFN", "ALK", "ALL", "AMD", "ANG", "AOA", "AOK", "AON", "AOR", "ARA", "ARL", "ARM", "ARP", "ARS", "ATS", "AUD", "AWG", "AZM", "AZN", "BAD", "BAM", "BAN", "BBD", "BDT", "BEC", "BEF", "BEL", "BGL", "BGM", "BGN", "BGO", "BHD", "BIF", "BMD", "BND", "BOB", "BOL", "BOP", "BOV", "BRB", "BRC", "BRE", "BRL", "BRN", "BRR", "BRZ", "BSD", "BTN", "BUK", "BWP", "BYB", "BYN", "BYR", "BZD", "CAD", "CDF", "CHE", "CHF", "CHW", "CLE", "CLF", "CLP", "CNH", "CNX", "CNY", "COP", "COU", "CRC", "CSD", "CSK", "CUC", "CUP", "CVE", "CYP", "CZK", "DDM", "DEM", "DJF", "DKK", "DOP", "DZD", "ECS", "ECV", "EEK", "EGP", "ERN", "ESA", "ESB", "ESP", "ETB", "EUR", "FIM", "FJD", "FKP", "FRF", "GBP", "GEK", "GEL", "GHC", "GHS", "GIP", "GMD", "GNF", "GNS", "GQE", "GRD", "GTQ", "GWE", "GWP", "GYD", "HKD", "HNL", "HRD", "HRK", "HTG", "HUF", "IDR", "IEP", "ILP", "ILR", "ILS", "INR", "IQD", "IRR", "ISJ", "ISK", "ITL", "JMD", "JOD", "JPY", "KES", "KGS", "KHR", "KMF", "KPW", "KRH", "KRO", "KRW", "KWD", "KYD", "KZT", "LAK", "LBP", "LKR", "LRD", "LSL", "LTL", "LTT", "LUC", "LUF", "LUL", "LVL", "LVR", "LYD", "MAD", "MAF", "MCF", "MDC", "MDL", "MGA", "MGF", "MKD", "MKN", "MLF", "MMK", "MNT", "MOP", "MRO", "MRU", "MTL", "MTP", "MUR", "MVP", "MVR", "MWK", "MXN", "MXP", "MXV", "MYR", "MZE", "MZM", "MZN", "NAD", "NGN", "NIC", "NIO", "NLG", "NOK", "NPR", "NZD", "OMR", "PAB", "PEI", "PEN", "PES", "PGK", "PHP", "PKR", "PLN", "PLZ", "PTE", "PYG", "QAR", "RHD", "ROL", "RON", "RSD", "RUB", "RUR", "RWF", "SAR", "SBD", "SCR", "SDD", "SDG", "SDP", "SEK", "SGD", "SHP", "SIT", "SKK", "SLE", "SLL", "SOS", "SRD", "SRG", "SSP", "STD", "STN", "SUR", "SVC", "SYP", "SZL", "THB", "TJR", "TJS", "TMM", "TMT", "TND", "TOP", "TPE", "TRL", "TRY", "TTD", "TWD", "TZS", "UAH", "UAK", "UGS", "UGX", "USD", "USN", "USS", "UYI", "UYP", "UYU", "UYW", "UZS", "VEB", "VED", "VEF", "VES", "VND", "VNN", "VUV", "WST", "XAF", "XAG", "XAU", "XBA", "XBB", "XBC", "XBD", "XCD", "XCG", "XDR", "XEU", "XFO", "XFU", "XOF", "XPD", "XPF", "XPT", "XRE", "XSU", "XTS", "XUA", "XXX", "YDD", "YER", "YUD", "YUM", "YUN", "YUR", "ZAL", "ZAR", "ZMK", "ZMW", "ZRN", "ZRZ", "ZWD", "ZWG", "ZWL", "ZWR"},
		currencyNegativePrefix: "(",
		currencyNegativeSuffix: ")",
		monthsAbbreviated:      []string{"", "ม.ค.", "ก.พ.", "มี.ค.", "เม.ย.", "พ.ค.", "มิ.ย.", "ก.ค.", "ส.ค.", "ก.ย.", "ต.ค.", "พ.ย.", "ธ.ค."},
		monthsNarrow:           []string{"", "ม.ค.", "ก.พ.", "มี.ค.", "เม.ย.", "พ.ค.", "มิ.ย.", "ก.ค.", "ส.ค.", "ก.ย.", "ต.ค.", "พ.ย.", "ธ.ค."},
		monthsWide:             []string{"", "มกราคม", "กุมภาพันธ์", "มีนาคม", "เมษายน", "พฤษภาคม", "มิถุนายน", "กรกฎาคม", "สิงหาคม", "กันยายน", "ตุลาคม", "พฤศจิกายน", "ธันวาคม"},
		daysAbbreviated:        []string{"อาทิตย์", "จันทร์", "อังคาร", "พุธ", "พฤหัส", "ศุกร์", "เสาร์"},
		daysNarrow:             []string{"อา", "จ", "อ", "พ", "พฤ", "ศ", "ส"},
		daysShort:              []string{"อา.", "จ.", "อ.", "พ.", "พฤ.", "ศ.", "ส."},
		daysWide:               []string{"วันอาทิตย์", "วันจันทร์", "วันอังคาร", "วันพุธ", "วันพฤหัสบดี", "วันศุกร์", "วันเสาร์"},
		periodsAbbreviated:     []string{"", ""},
		periodsNarrow:          []string{"a", "p"},
		periodsWide:            []string{"ก่อนเที่ยง", "หลังเที่ยง"},
		erasAbbreviated:        []string{"ก่อน ค.ศ.", "ค.ศ."},
		erasNarrow:             []string{"", ""},
		erasWide:               []string{"ปีก่อนคริสตกาล", "คริสต์ศักราช"},
		timezones:              map[string]string{"ACDT": "เวลาออมแสงทางตอนกลางของออสเตรเลีย", "ACST": "เวลาฤดูร้อนอาเกร", "ACT": "เวลามาตรฐานอาเกร", "ACWDT": "เวลาออมแสงทางตะวันตกตอนกลางของออสเตรเลีย", "ACWST": "เวลามาตรฐานทางตะวันตกตอนกลางของออสเตรเลีย", "ADT": "เวลาออมแสงของแอตแลนติก", "ADT Arabia": "เวลาออมแสงอาหรับ", "AEDT": "เวลาออมแสงทางตะวันออกของออสเตรเลีย", "AEST": "เวลามาตรฐานทางตะวันออกของออสเตรเลีย", "AFT": "เวลาอัฟกานิสถาน", "AKDT": "เวลาออมแสงของอะแลสกา", "AKST": "เวลามาตรฐานอะแลสกา", "AMST": "เวลาฤดูร้อนแอมะซอน", "AMST Armenia": "เวลาฤดูร้อนอาร์เมเนีย", "AMT": "เวลามาตรฐานแอมะซอน", "AMT Armenia": "เวลามาตรฐานอาร์เมเนีย", "ANAST": "เวลาฤดูร้อนอะนาดีร์", "ANAT": "เวลามาตรฐานอะนาดีร์", "ARST": "เวลาฤดูร้อนอาร์เจนตินา", "ART": "เวลามาตรฐานอาร์เจนตินา", "AST": "เวลามาตรฐานแอตแลนติก", "AST Arabia": "เวลามาตรฐานอาหรับ", "AWDT": "เวลาออมแสงทางตะวันตกของออสเตรเลีย", "AWST": "เวลามาตรฐานทางตะวันตกของออสเตรเลีย", "AZST": "เวลาฤดูร้อนอาเซอร์ไบจาน", "AZT": "เวลามาตรฐานอาเซอร์ไบจาน", "BDT Bangladesh": "เวลาฤดูร้อนบังกลาเทศ", "BNT": "เวลาบรูไนดารุสซาลาม", "BOT": "เวลาโบลิเวีย", "BRST": "เวลาฤดูร้อนบราซิเลีย", "BRT": "เวลามาตรฐานบราซิเลีย", "BST Bangladesh": "เวลามาตรฐานบังกลาเทศ", "BT": "เวลาภูฏาน", "CAST": "เวลาเคซีย์", "CAT": "เวลาแอฟริกากลาง", "CCT": "เวลาหมู่เกาะโคโคส", "CDT": "เวลาออมแสงตอนกลางในอเมริกาเหนือ", "CHADT": "เวลาออมแสงแชทัม", "CHAST": "เวลามาตรฐานแชทัม", "CHUT": "เวลาชุก", "CKT": "เวลามาตรฐานหมู่เกาะคุก", "CKT DST": "เวลาครึ่งฤดูร้อนหมู่เกาะคุก", "CLST": "เวลาฤดูร้อนชิลี", "CLT": "เวลามาตรฐานชิลี", "COST": "เวลาฤดูร้อนโคลอมเบีย", "COT": "เวลามาตรฐานโคลอมเบีย", "CST": "เวลามาตรฐานตอนกลางในอเมริกาเหนือ", "CST China": "เวลามาตรฐานจีน", "CST China DST": "เวลาออมแสงจีน", "CVST": "เวลาฤดูร้อนเคปเวิร์ด", "CVT": "เวลามาตรฐานเคปเวิร์ด", "CXT": "เวลาเกาะคริสต์มาส", "ChST": "เวลาชามอร์โร", "ChST NMI": "เวลาหมู่เกาะมาเรียนาเหนือ", "CuDT": "เวลาออมแสงของคิวบา", "CuST": "เวลามาตรฐานคิวบา", "DAVT": "เวลาเดวิส", "DDUT": "เวลาดูมองต์ดูร์วิลล์", "EASST": "เวลาฤดูร้อนเกาะอีสเตอร์", "EAST": "เวลามาตรฐานเกาะอีสเตอร์", "EAT": "เวลาแอฟริกาตะวันออก", "ECT": "เวลาเอกวาดอร์", "EDT": "เวลาออมแสงทางตะวันออกในอเมริกาเหนือ", "EGDT": "เวลาฤดูร้อนกรีนแลนด์ตะวันออก", "EGST": "เวลามาตรฐานกรีนแลนด์ตะวันออก", "EST": "เวลามาตรฐานทางตะวันออกในอเมริกาเหนือ", "FEET": "เวลายุโรปตะวันออกไกล", "FJT": "เวลามาตรฐานฟิจิ", "FJT Summer": "เวลาฤดูร้อนฟิจิ", "FKST": "เวลาฤดูร้อนหมู่เกาะฟอล์กแลนด์", "FKT": "เวลามาตรฐานหมู่เกาะฟอล์กแลนด์", "FNST": "เวลาฤดูร้อนของหมู่เกาะเฟอร์นันโด", "FNT": "เวลามาตรฐานหมู่เกาะเฟอร์นันโด", "GALT": "เวลากาลาปาโกส", "GAMT": "เวลาแกมเบียร์", "GEST": "เวลาฤดูร้อนจอร์เจีย", "GET": "เวลามาตรฐานจอร์เจีย", "GFT": "เวลาเฟรนช์เกียนา", "GIT": "เวลาหมู่เกาะกิลเบิร์ต", "GMT": "เวลามาตรฐานกรีนิช", "GNSST": "GNSST", "GNST": "GNST", "GST": "เวลาเซาท์จอร์เจีย", "GST Guam": "เวลากวม", "GYT": "เวลากายอานา", "HADT": "เวลาออมแสงฮาวาย-อะลูเชียน", "HAST": "เวลามาตรฐานฮาวาย-อะลูเชียน", "HKST": "เวลาฤดูร้อนฮ่องกง", "HKT": "เวลามาตรฐานฮ่องกง", "HOVST": "เวลาฤดูร้อนฮอฟด์", "HOVT": "เวลามาตรฐานฮอฟด์", "ICT": "เวลาอินโดจีน", "IDT": "เวลาออมแสงอิสราเอล", "IOT": "เวลามหาสมุทรอินเดีย", "IRKST": "เวลาฤดูร้อนอีร์คุตสค์", "IRKT": "เวลามาตรฐานอีร์คุตสค์", "IRST": "เวลามาตรฐานอิหร่าน", "IRST DST": "เวลาออมแสงอิหร่าน", "IST": "เวลาอินเดีย", "IST Israel": "เวลามาตรฐานอิสราเอล", "JDT": "เวลาออมแสงญี่ปุ่น", "JST": "เวลามาตรฐานญี่ปุ่น", "KOST": "เวลาคอสไร", "KRAST": "เวลาฤดูร้อนครัสโนยาสค์", "KRAT": "เวลามาตรฐานครัสโนยาสค์", "KST": "เวลามาตรฐานเกาหลี", "KST DST": "เวลาออมแสงเกาหลี", "LHDT": "เวลาออมแสงลอร์ดโฮว์", "LHST": "เวลามาตรฐานลอร์ดโฮว์", "LINT": "เวลาหมู่เกาะไลน์", "MAGST": "เวลาฤดูร้อนมากาดาน", "MAGT": "เวลามาตรฐานมากาดาน", "MART": "เวลามาร์เคซัส", "MAWT": "เวลามอว์สัน", "MDT": "เวลาออมแสงแถบภูเขาในอเมริกาเหนือ", "MESZ": "เวลาฤดูร้อนยุโรปกลาง", "MEZ": "เวลามาตรฐานยุโรปกลาง", "MHT": "เวลาหมู่เกาะมาร์แชลล์", "MMT": "เวลาพม่า", "MSD": "เวลาฤดูร้อนมอสโก", "MST": "เวลามาตรฐานแถบภูเขาในอเมริกาเหนือ", "MUST": "เวลาฤดูร้อนของมอริเชียส", "MUT": "เวลามาตรฐานมอริเชียส", "MVT": "เวลามัลดีฟส์", "MYT": "เวลามาเลเซีย", "NCT": "เวลามาตรฐานนิวแคลิโดเนีย", "NDT": "เวลาออมแสงนิวฟันด์แลนด์", "NDT New Caledonia": "เวลาฤดูร้อนนิวแคลิโดเนีย", "NFDT": "เวลาออมแสงเกาะนอร์ฟอล์ก", "NFT": "เวลามาตรฐานเกาะนอร์ฟอล์ก", "NOVST": "เวลาฤดูร้อนโนโวซีบีสค์", "NOVT": "เวลามาตรฐานโนโวซีบีสค์", "NPT": "เวลาเนปาล", "NRT": "เวลานาอูรู", "NST": "เวลามาตรฐานนิวฟันด์แลนด์", "NUT": "เวลานีอูเอ", "NZDT": "เวลาออมแสงนิวซีแลนด์", "NZST": "เวลามาตรฐานนิวซีแลนด์", "OESZ": "เวลาฤดูร้อนยุโรปตะวันออก", "OEZ": "เวลามาตรฐานยุโรปตะวันออก", "OMSST": "เวลาฤดูร้อนออมสค์", "OMST": "เวลามาตรฐานออมสค์", "PDT": "เวลาออมแสงแปซิฟิกในอเมริกาเหนือ", "PDTM": "เวลาออมแสงแปซิฟิกเม็กซิโก", "PETDT": "เวลาฤดูร้อนเปโตรปัฟลอฟสค์-คัมชัตสกี", "PETST": "เวลาเปโตรปัฟลอฟสค์-คัมชัตสกี", "PGT": "เวลาปาปัวนิวกินี", "PHOT": "เวลาหมู่เกาะฟินิกซ์", "PKT": "เวลามาตรฐานปากีสถาน", "PKT DST": "เวลาฤดูร้อนปากีสถาน", "PMDT": "เวลาออมแสงของแซงปีแยร์และมีเกอลง", "PMST": "เวลามาตรฐานแซงปีแยร์และมีเกอลง", "PONT": "เวลาโปนาเป", "PST": "เวลามาตรฐานแปซิฟิกในอเมริกาเหนือ", "PST Philippine": "เวลามาตรฐานฟิลิปปินส์", "PST Philippine DST": "เวลาฤดูร้อนฟิลิปปินส์", "PST Pitcairn": "เวลาพิตแคร์น", "PSTM": "เวลามาตรฐานแปซิฟิกเม็กซิโก", "PWT": "เวลาปาเลา", "PYST": "เวลาฤดูร้อนปารากวัย", "PYT": "เวลามาตรฐานปารากวัย", "PYT Korea": "เวลาเปียงยาง", "RET": "เวลาเรอูนียง", "ROTT": "เวลาโรธีรา", "SAKST": "เวลาฤดูร้อนซาคาลิน", "SAKT": "เวลามาตรฐานซาคาลิน", "SAMST": "เวลาฤดูร้อนซามารา", "SAMT": "เวลามาตรฐานซามารา", "SAST": "เวลาแอฟริกาใต้", "SBT": "เวลาหมู่เกาะโซโลมอน", "SCT": "เวลาเซเชลส์", "SGT": "เวลาสิงคโปร์", "SLST": "เวลาลังกา", "SRT": "เวลาซูรินาเม", "SST Samoa": "เวลามาตรฐานซามัว", "SST Samoa Apia": "เวลามาตรฐานอาปีอา", "SST Samoa Apia DST": "เวลาออมแสงอาปีอา", "SST Samoa DST": "เวลาฤดูร้อนซามัว", "SYOT": "เวลาไซโยวา", "TAAF": "เวลาเฟรนช์เซาเทิร์นและแอนตาร์กติก", "TAHT": "เวลาตาฮีตี", "TJT": "เวลาทาจิกิสถาน", "TKT": "เวลาโตเกเลา", "TLT": "เวลาติมอร์ตะวันออก", "TMST": "เวลาฤดูร้อนเติร์กเมนิสถาน", "TMT": "เวลามาตรฐานเติร์กเมนิสถาน", "TOST": "เวลาฤดูร้อนตองกา", "TOT": "เวลามาตรฐานตองกา", "TVT": "เวลาตูวาลู", "TWT": "เวลามาตรฐานไทเป", "TWT DST": "เวลาออมแสงไทเป", "ULAST": "เวลาฤดูร้อนอูลานบาตอร์", "ULAT": "เวลามาตรฐานอูลานบาตอร์", "UYST": "เวลาฤดูร้อนอุรุกวัย", "UYT": "เวลามาตรฐานอุรุกวัย", "UZT": "เวลามาตรฐานอุซเบกิสถาน", "UZT DST": "เวลาฤดูร้อนอุซเบกิสถาน", "VET": "เวลาเวเนซุเอลา", "VLAST": "เวลาฤดูร้อนวลาดีวอสตอค", "VLAT": "เวลามาตรฐานวลาดีวอสตอค", "VOLST": "เวลาฤดูร้อนวอลโกกราด", "VOLT": "เวลามาตรฐานวอลโกกราด", "VOST": "เวลาวอสตอค", "VUT": "เวลามาตรฐานวานูอาตู", "VUT DST": "เวลาฤดูร้อนวานูอาตู", "WAKT": "เวลาเกาะเวก", "WARST": "เวลาฤดูร้อนทางตะวันตกของอาร์เจนตินา", "WART": "เวลามาตรฐานทางตะวันตกของอาร์เจนตินา", "WAST": "เวลาแอฟริกาตะวันตก", "WAT": "เวลาแอฟริกาตะวันตก", "WESZ": "เวลาฤดูร้อนยุโรปตะวันตก", "WEZ": "เวลามาตรฐานยุโรปตะวันตก", "WFT": "เวลาวาลลิสและฟุตูนา", "WGST": "เวลาฤดูร้อนกรีนแลนด์ตะวันตก", "WGT": "เวลามาตรฐานกรีนแลนด์ตะวันตก", "WIB": "เวลาอินโดนีเซียฝั่งตะวันตก", "WIT": "เวลาอินโดนีเซียฝั่งตะวันออก", "WITA": "เวลาอินโดนีเซียตอนกลาง", "YAKST": "เวลาฤดูร้อนยาคุตสค์", "YAKT": "เวลามาตรฐานยาคุตสค์", "YEKST": "เวลาฤดูร้อนเยคาเตรินบูร์ก", "YEKT": "เวลามาตรฐานเยคาเตรินบูร์ก", "YST": "เวลายูคอน", "МСК": "เวลามาตรฐานมอสโก", "اقتاۋ": "เวลามาตรฐานอัคตาอู", "اقتاۋ قالاسى": "เวลาฤดูร้อนอัคตาอู", "اقتوبە": "เวลามาตรฐานอัคโทเบ", "اقتوبە قالاسى": "เวลาฤดูร้อนอัคโทเบ", "الماتى": "เวลามาตรฐานอัลมาตี", "الماتى قالاسى": "เวลาฤดูร้อนอัลมาตี", "باتىس قازاق ەلى": "เวลาคาซัคสถานตะวันตก", "شىعىش قازاق ەلى": "เวลาคาซัคสถานตะวันออก", "قازاق ەلى": "เวลาคาซัคสถาน", "قىرعىزستان": "เวลาคีร์กีซสถาน", "قىزىلوردا": "เวลามาตรฐานคืยซิลออร์ดา", "قىزىلوردا قالاسى": "เวลาฤดูร้อนคืยซิลออร์ดา", "∅∅∅": "เวลาฤดูร้อนเปรู"},
	}
}

// Locale returns the current translators string locale
func (th *th_TH) Locale() string {
	return th.locale
}

// PluralsCardinal returns the list of cardinal plural rules associated with 'th_TH'
func (th *th_TH) PluralsCardinal() []locales.PluralRule {
	return th.pluralsCardinal
}

// PluralsOrdinal returns the list of ordinal plural rules associated with 'th_TH'
func (th *th_TH) PluralsOrdinal() []locales.PluralRule {
	return th.pluralsOrdinal
}

// PluralsRange returns the list of range plural rules associated with 'th_TH'
func (th *th_TH) PluralsRange() []locales.PluralRule {
	return th.pluralsRange
}

// CardinalPluralRule returns the cardinal PluralRule given 'num' and digits/precision of 'v' for 'th_TH'
func (th *th_TH) CardinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// OrdinalPluralRule returns the ordinal PluralRule given 'num' and digits/precision of 'v' for 'th_TH'
func (th *th_TH) OrdinalPluralRule(num float64, v uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// RangePluralRule returns the ordinal PluralRule given 'num1', 'num2' and digits/precision of 'v1' and 'v2' for 'th_TH'
func (th *th_TH) RangePluralRule(num1 float64, v1 uint64, num2 float64, v2 uint64) locales.PluralRule {
	return locales.PluralRuleOther
}

// MonthAbbreviated returns the locales abbreviated month given the 'month' provided
func (th *th_TH) MonthAbbreviated(month time.Month) string {
	return th.monthsAbbreviated[month]
}

// MonthsAbbreviated returns the locales abbreviated months
func (th *th_TH) MonthsAbbreviated() []string {
	return th.monthsAbbreviated[1:]
}

// MonthNarrow returns the locales narrow month given the 'month' provided
func (th *th_TH) MonthNarrow(month time.Month) string {
	return th.monthsNarrow[month]
}

// MonthsNarrow returns the locales narrow months
func (th *th_TH) MonthsNarrow() []string {
	return th.monthsNarrow[1:]
}

// MonthWide returns the locales wide month given the 'month' provided
func (th *th_TH) MonthWide(month time.Month) string {
	return th.monthsWide[month]
}

// MonthsWide returns the locales wide months
func (th *th_TH) MonthsWide() []string {
	return th.monthsWide[1:]
}

// WeekdayAbbreviated returns the locales abbreviated weekday given the 'weekday' provided
func (th *th_TH) WeekdayAbbreviated(weekday time.Weekday) string {
	return th.daysAbbreviated[weekday]
}

// WeekdaysAbbreviated returns the locales abbreviated weekdays
func (th *th_TH) WeekdaysAbbreviated() []string {
	return th.daysAbbreviated
}

// WeekdayNarrow returns the locales narrow weekday given the 'weekday' provided
func (th *th_TH) WeekdayNarrow(weekday time.Weekday) string {
	return th.daysNarrow[weekday]
}

// WeekdaysNarrow returns the locales narrow weekdays
func (th *th_TH) WeekdaysNarrow() []string {
	return th.daysNarrow
}

// WeekdayShort returns the locales short weekday given the 'weekday' provided
func (th *th_TH) WeekdayShort(weekday time.Weekday) string {
	return th.daysShort[weekday]
}

// WeekdaysShort returns the locales short weekdays
func (th *th_TH) WeekdaysShort() []string {
	return th.daysShort
}

// WeekdayWide returns the locales wide weekday given the 'weekday' provided
func (th *th_TH) WeekdayWide(weekday time.Weekday) string {
	return th.daysWide[weekday]
}

// WeekdaysWide returns the locales wide weekdays
func (th *th_TH) WeekdaysWide() []string {
	return th.daysWide
}

// Decimal returns the decimal point of number
func (th *th_TH) Decimal() string {
	return th.decimal
}

// Group returns the group of number
func (th *th_TH) Group() string {
	return th.group
}

// Group returns the minus sign of number
func (th *th_TH) Minus() string {
	return th.minus
}

// FmtNumber returns 'num' with digits/precision of 'v' for 'th_TH' and handles both Whole and Real numbers based on 'v'
func (th *th_TH) FmtNumber(num float64, v uint64) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, th.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, th.group[0])
				count = 1
			} else {
				count++
			}
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, th.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b)
}

// FmtPercent returns 'num' with digits/precision of 'v' for 'th_TH' and handles both Whole and Real numbers based on 'v'
// NOTE: 'num' passed into FmtPercent is assumed to be in percent already
func (th *th_TH) FmtPercent(num float64, v uint64) string {
	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	l := len(s) + 3
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, th.decimal[0])
			continue
		}

		b = append(b, s[i])
	}

	if num < 0 {
		b = append(b, th.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	b = append(b, th.percent...)

	return string(b)
}

// FmtCurrency returns the currency representation of 'num' with digits/precision of 'v' for 'th_TH'
func (th *th_TH) FmtCurrency(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := th.currencies[currency]
	l := len(s) + len(symbol) + 2 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, th.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, th.group[0])
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
		b = append(b, th.minus[0])
	}

	// reverse
	for i, j := 0, len(b)-1; i < j; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	if int(v) < 2 {

		if v == 0 {
			b = append(b, th.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	return string(b)
}

// FmtAccounting returns the currency representation of 'num' with digits/precision of 'v' for 'th_TH'
// in accounting notation.
func (th *th_TH) FmtAccounting(num float64, v uint64, currency currency.Type) string {

	s := strconv.FormatFloat(math.Abs(num), 'f', int(v), 64)
	symbol := th.currencies[currency]
	l := len(s) + len(symbol) + 4 + 1*len(s[:len(s)-int(v)-1])/3
	count := 0
	inWhole := v == 0
	b := make([]byte, 0, l)

	for i := len(s) - 1; i >= 0; i-- {

		if s[i] == '.' {
			b = append(b, th.decimal[0])
			inWhole = true
			continue
		}

		if inWhole {
			if count == 3 {
				b = append(b, th.group[0])
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

		b = append(b, th.currencyNegativePrefix[0])

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
			b = append(b, th.decimal...)
		}

		for i := 0; i < 2-int(v); i++ {
			b = append(b, '0')
		}
	}

	if num < 0 {
		b = append(b, th.currencyNegativeSuffix...)
	}

	return string(b)
}

// FmtDateShort returns the short date representation of 't' for 'th_TH'
func (th *th_TH) FmtDateShort(t time.Time) string {

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

// FmtDateMedium returns the medium date representation of 't' for 'th_TH'
func (th *th_TH) FmtDateMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, th.monthsAbbreviated[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateLong returns the long date representation of 't' for 'th_TH'
func (th *th_TH) FmtDateLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, th.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() < 0 {
		b = append(b, th.erasAbbreviated[0]...)
	} else {
		b = append(b, th.erasAbbreviated[1]...)
	}

	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtDateFull returns the full date representation of 't' for 'th_TH'
func (th *th_TH) FmtDateFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = append(b, th.daysWide[t.Weekday()]...)
	b = append(b, []byte{0xe0, 0xb8, 0x97, 0xe0, 0xb8, 0xb5, 0xe0, 0xb9, 0x88, 0x20}...)
	b = strconv.AppendInt(b, int64(t.Day()), 10)
	b = append(b, []byte{0x20}...)
	b = append(b, th.monthsWide[t.Month()]...)
	b = append(b, []byte{0x20}...)

	if t.Year() < 0 {
		b = append(b, th.erasWide[0]...)
	} else {
		b = append(b, th.erasWide[1]...)
	}

	b = append(b, []byte{0x20}...)

	if t.Year() > 0 {
		b = strconv.AppendInt(b, int64(t.Year()), 10)
	} else {
		b = strconv.AppendInt(b, int64(-t.Year()), 10)
	}

	return string(b)
}

// FmtTimeShort returns the short time representation of 't' for 'th_TH'
func (th *th_TH) FmtTimeShort(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, th.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)

	return string(b)
}

// FmtTimeMedium returns the medium time representation of 't' for 'th_TH'
func (th *th_TH) FmtTimeMedium(t time.Time) string {

	b := make([]byte, 0, 32)

	if t.Hour() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, th.timeSeparator...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, th.timeSeparator...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)

	return string(b)
}

// FmtTimeLong returns the long time representation of 't' for 'th_TH'
func (th *th_TH) FmtTimeLong(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x20, 0xe0, 0xb8, 0x99, 0xe0, 0xb8, 0xb2, 0xe0, 0xb8, 0xac, 0xe0, 0xb8, 0xb4, 0xe0, 0xb8, 0x81, 0xe0, 0xb8, 0xb2, 0x20}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20, 0xe0, 0xb8, 0x99, 0xe0, 0xb8, 0xb2, 0xe0, 0xb8, 0x97, 0xe0, 0xb8, 0xb5, 0x20}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0xe0, 0xb8, 0xa7, 0xe0, 0xb8, 0xb4, 0xe0, 0xb8, 0x99, 0xe0, 0xb8, 0xb2, 0xe0, 0xb8, 0x97, 0xe0, 0xb8, 0xb5, 0x20}...)

	tz, _ := t.Zone()
	b = append(b, tz...)

	return string(b)
}

// FmtTimeFull returns the full time representation of 't' for 'th_TH'
func (th *th_TH) FmtTimeFull(t time.Time) string {

	b := make([]byte, 0, 32)

	b = strconv.AppendInt(b, int64(t.Hour()), 10)
	b = append(b, []byte{0x20, 0xe0, 0xb8, 0x99, 0xe0, 0xb8, 0xb2, 0xe0, 0xb8, 0xac, 0xe0, 0xb8, 0xb4, 0xe0, 0xb8, 0x81, 0xe0, 0xb8, 0xb2, 0x20}...)

	if t.Minute() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Minute()), 10)
	b = append(b, []byte{0x20, 0xe0, 0xb8, 0x99, 0xe0, 0xb8, 0xb2, 0xe0, 0xb8, 0x97, 0xe0, 0xb8, 0xb5, 0x20}...)

	if t.Second() < 10 {
		b = append(b, '0')
	}

	b = strconv.AppendInt(b, int64(t.Second()), 10)
	b = append(b, []byte{0x20, 0xe0, 0xb8, 0xa7, 0xe0, 0xb8, 0xb4, 0xe0, 0xb8, 0x99, 0xe0, 0xb8, 0xb2, 0xe0, 0xb8, 0x97, 0xe0, 0xb8, 0xb5, 0x20}...)

	tz, _ := t.Zone()

	if btz, ok := th.timezones[tz]; ok {
		b = append(b, btz...)
	} else {
		b = append(b, tz...)
	}

	return string(b)
}
