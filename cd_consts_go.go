package cd_consts_go

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
)

const (
	FILENAME       string = "de440s.bsp"
	EXPECTEDSHA512 string = "a244335d9eddc1e4fd2f3f8ddabf360020f650bc8fca2c4e7e0f66018db7fd2691dd63f52e3652653e096d97ad74cd48c10b4587a4d5a9bb68dbae5cecf06449"
	FILELENGTH     int    = 32726016
	SIZEOFREC      int    = 1024

	//Astronomical Unit
	AU = 0.1495978707e9 // km 149597870.7

	// Здесь мы как начальное значение ставим eps = 23°26'21,448" градуса согласно CD
	// double const RAD_TO_DEG = 5.7295779513082320877e1;
	// Obliquity of the ecliptic  = 23°26'21,448"  - на 1 января 2000 года = 23.43929111111111
	// 23.43929111111111/5.7295779513082320877e1 = 0.4090928042223289
	MED_EPS = 0.4090928042223289

	SSB       = 0
	MERCURY   = 1 // 7,01° (относительно эклиптики)
	VENUS     = 2 // 3,39458° (относительно эклиптики)
	EARTH     = 3
	MARS      = 4 // 1,85061° (относительно эклиптики)
	JUPITER   = 5 // 1,304° (относительно эклиптики)
	SATURN    = 6 // 2,485 240° (относительно эклиптики)
	URANUS    = 7 // 0,772556° (относительно эклиптики)
	NEPTUNE   = 8 // 1,767975° (относительно эклиптики)
	PLUTO     = 9 // 17°,14 (относительно эклиптики)
	SUN       = 10
	MOON      = 11 // 5,14° (относительно эклиптики)
	NORTHNODE = 12
	SOUTHNODE = 13
	HIRON     = 14

	HEAD   = 0
	AJNA   = 1
	THROAT = 2
	G      = 3
	SACRAL = 4
	ROOT   = 5
	EGO    = 6
	SPLEEN = 7
	EMO    = 8

	NUMBEROFGATES    = 65 //from 1 to 64
	NUMBEROFCHANNELS = 37 //from 1 to 36

	// from 0 to 13, don't count Hiron yet
	NUMBEROFPLANETS = 14

	// from 0 to 8
	NUMBEROFCENTERS = 9
)

// main type with complete calcuated information
type CdInfo struct {
	HdInfo
	FdInfo
	AstroInfo
	NumerologyInfo
}

type HdInfo struct {
	Personality HdObjects
	Design      HdObjects
	Gates       [NUMBEROFGATES]Gate // from 1 to 64
	Channels    [NUMBEROFCHANNELS]Channel
	Centers     Centers
	Phs
	Variable string
	Psychology
	Cross      Cross
	Profile    string
	Authority  string
	Definition string
	Type       string
}

func (hd *HdInfo) Init() {

	for i := 1; i < NUMBEROFGATES; i++ {
		hd.Gates[i].Number = i

		if i < NUMBEROFCHANNELS {
			hd.Channels[i].Number = i
		}

	}

	hd.Personality.Planets.Init()
	hd.Design.Init()

	hd.Centers.Init()

}

type Gate struct {
	Number  int
	Pers    int //сколько раз активированы по личности
	Des     int // сколько раз активированы по дизайну
	Defined bool
}

// 36 каналов ДЧ
/*

  1 - 64-47
  2 - 61-24
  3 - 63-4

  4 - 17-62
  5 - 43-23
  6 - 11-56

  7 - 48-16

  8 - 57-20
  9 - 34-20
  10- 10-20
  11- 57-10
  12- 57-34
  13- 34-10

  14- 7-31
  15- 1-8
  16- 13-33
  17- 21-45
  18- 22-12
  19- 36-35

  20- 5-15
  21- 14-2
  22- 29-46
  23- 51-25

  24- 44-26
  25- 27-50
  26- 59-6
  27- 37-40

  28- 54-32
  29- 38-28
  30- 58-18

  31- 53-42
  32- 60-3
  33- 52-9

  34- 19-49
  35- 39-55
  36- 41-30

*/
type Channel struct {
	Number     int
	FirstGate  Gate
	SecondGate Gate
	Defined    bool
}

type HdObjects struct {
	Planets
	Centers Centers
	TimeData
	Authority string
}

type Planet struct {
	Longitude float64 //in Radians
	Name      string
	Number    int
	HdStructure
	FdStructure
	ZodiacStructure
}

type Planets struct {
	Planet [NUMBEROFPLANETS]Planet
}

func (pl *Planets) Init() {
	pl.Planet[0] = Planet{Name: "SSB", Number: 0}
	pl.Planet[1] = Planet{Name: "Mercury", Number: 1}
	pl.Planet[2] = Planet{Name: "Venus", Number: 2}
	pl.Planet[3] = Planet{Name: "Earth", Number: 3}
	pl.Planet[4] = Planet{Name: "Mars", Number: 4}
	pl.Planet[5] = Planet{Name: "Jupiter", Number: 5}
	pl.Planet[6] = Planet{Name: "Saturn", Number: 6}
	pl.Planet[7] = Planet{Name: "Uranus", Number: 7}
	pl.Planet[8] = Planet{Name: "Neptune", Number: 8}
	pl.Planet[9] = Planet{Name: "Pluto", Number: 9}
	pl.Planet[10] = Planet{Name: "Sun", Number: 10}
	pl.Planet[11] = Planet{Name: "Moon", Number: 11}
	pl.Planet[12] = Planet{Name: "NorthNode", Number: 12}
	pl.Planet[13] = Planet{Name: "SouthNde", Number: 13}
	//{Name: "Hiron", Number: 14},

}

/*
func InitPlanets() Planets {

	return Planets{[NUMBEROFPLANETS]Planet{
		{Name: "SSB", Number: 0},
		{Name: "Mercury", Number: 1},
		{Name: "Venus", Number: 2},
		{Name: "Earth", Number: 3},
		{Name: "Mars", Number: 4},
		{Name: "Jupiter", Number: 5},
		{Name: "Saturn", Number: 6},
		{Name: "Uranus", Number: 7},
		{Name: "Neptune", Number: 8},
		{Name: "Pluto", Number: 9},
		{Name: "Sun", Number: 10},
		{Name: "Moon", Number: 11},
		{Name: "NorthNode", Number: 12},
		{Name: "SouthNde", Number: 13},
		//{Name: "Hiron", Number: 14},
	}}

}
*/

// simple Date structure
// ok
type GregDate struct {
	Year    int `json:"year"`
	Month   int `json:"month"`
	Day     int `json:"day"`
	Hour    int `json:"hour"`
	Minutes int `json:"minutes"`
	Seconds int `json:"seconds"`
}

func (gd GregDate) String() string {

	var monthsFormatted string
	if gd.Month > 9 {
		monthsFormatted = "" + fmt.Sprint(gd.Month)

	} else {
		monthsFormatted = "0" + fmt.Sprint(gd.Month)
	}

	var daysFormatted string
	if gd.Day > 9 {
		daysFormatted = "" + fmt.Sprint(gd.Day)

	} else {
		daysFormatted = "0" + fmt.Sprint(gd.Day)
	}

	var hoursFormatted string
	if gd.Hour > 9 {
		hoursFormatted = "" + fmt.Sprint(gd.Hour)

	} else {
		hoursFormatted = "0" + fmt.Sprint(gd.Hour)
	}

	var minutesFormatted string
	if gd.Minutes > 9 {
		minutesFormatted = "" + fmt.Sprint(gd.Minutes)

	} else {
		minutesFormatted = "0" + fmt.Sprint(gd.Minutes)
	}

	var secondsFormatted string
	if gd.Seconds > 9 {
		secondsFormatted = "" + fmt.Sprint(gd.Seconds)

	} else {
		secondsFormatted = "0" + fmt.Sprint(gd.Seconds)
	}

	return "Date:  " + daysFormatted + "." + monthsFormatted + "." + fmt.Sprint(gd.Year) +
		"  Time: " + hoursFormatted + ":" + minutesFormatted + ":" + secondsFormatted
}

type TimeData struct {
	LocalTime GregDate //для design всегда 0
	UtcTime   GregDate

	TypeOfTyme    int    //Изначальный источник данных 2 - local time, 1- UTC Time,  0 - Ephemeries time
	Offset        int    //смещение локального времени от UTC в секундах
	SecFromJd2000 int64  // Ephemeries time
	Place         string // не пустой, только если время изначально Local, для design всегда пустой
}

type HdStructure struct {

	//номера округляются вверх

	Hex int // соответствует воротам

	Line float64

	Color float64

	Tone float64

	Base float64

	NumberOfPassedDegrees float64 //сколько пройдено в градусах от начала гексаграммы
}

func (hd HdStructure) String() string {

	return strconv.Itoa(hd.Hex) + "." + strconv.Itoa(int(math.Ceil(hd.Line))) + "." + strconv.Itoa(int(math.Ceil(hd.Color))) + "." + strconv.Itoa(int(math.Ceil(hd.Tone))) + "." + strconv.Itoa(int(math.Ceil(hd.Base)))
}

type FdInfo struct {
}

type FdStructure struct {
	Power     int
	Direction string // D - директное  R - ретроградное  S - стационарное
}

type AstroInfo struct {
}

type ZodiacStructure struct {
	Degrees int
	Minutes int
	Seconds int
	Zodiac  string
}

type Phs struct {
	Theme     string
	NutrType  string
	Cognition string
}

type Psychology struct {
	Motivation string
	Mind       string
}

type Cross struct {
	First  int
	Second int
	Third  int
	Forth  int
}

type NumerologyInfo struct {
}

type Centers struct {
	Center map[string]bool
}

func (cent *Centers) Init() {
	/*
		cent.Centers[0] = Center{Name: "Head"}
		cent.Centers[1] = Center{Name: "Ajna"}
		cent.Centers[2] = Center{Name: "Throat"}
		cent.Centers[3] = Center{Name: "G"}
		cent.Centers[4] = Center{Name: "Sacral"}
		cent.Centers[5] = Center{Name: "Root"}
		cent.Centers[6] = Center{Name: "Ego"}
		cent.Centers[7] = Center{Name: "Spleen"}
		cent.Centers[8] = Center{Name: "Emo"}

	*/
	cent.Center = make(map[string]bool, 9)
	cent.Center["Head"] = false
	cent.Center["Ajna"] = false
	cent.Center["Throat"] = false
	cent.Center["G"] = false
	cent.Center["Sacral"] = false
	cent.Center["Root"] = false
	cent.Center["Ego"] = false
	cent.Center["Spleen"] = false
	cent.Center["Emo"] = false

}

/*
http://astro.ukho.gov.uk/nao/miscellanea/DeltaT/
https://ru.wikipedia.org/wiki/%D0%94%D0%B5%D0%BB%D1%8C%D1%82%D0%B0_T
https://eclipse.gsfc.nasa.gov/SEhelp/deltatpoly2004.html
https://en.wikipedia.org/wiki/%CE%94T
*/
type DeltaTTableStructure struct {
	Year    int
	Seconds float64
}

// первый и последнй года таблицы значений Дельта Т для быстрого доступа и сама таблица
type DeltaTTable struct {
	FirstYear int
	LastYear  int
	Table     []DeltaTTableStructure
}

type BspFile struct {
	FilePtr     *bytes.Reader
	FileInfo    *FileInfo
	NodesCoords *[]NodesJsonStruct
	DeltaTTable *DeltaTTable
}

// [-4733494022,"north"],[-4732252235,"south"]
type NodesJsonStruct struct {
	Time  float64
	Which string
}

type FileInfo struct {
	PathToDir           string
	FileName            string
	Length              int64
	Sha512              string
	FirstSummaryRec     int
	FileRecordStruct    FileRecordStruct
	SummaryRecordStruct SummaryRecordStruct
	SummariesLineStruct []SummariesLines
}

type ArrayInfo struct {
	Init   float64 //:= 0.0;      // start time of the first record in array
	Intlen float64 //:= 0.0;    // the length of one record (seconds)
	Rsize  float64 //:= 0.0;     // number of elements in one record
	N      float64 //:= 0.0;         // number of records in segment
}

type SummariesLines struct {
	Name string //:= ""

	//просто порядковый номер в файле
	Number int //:= -1

	SEGMENT_START_TIME int64 // := 0; // always the same if only 1 Record in file
	SEGMENT_LAST_TIME  int64 //:= 0; // always the same if only 1 Record in file

	TargetCode int //:= -1;
	CenterCode int //:= -1;

	RefFrame   int //:= 0;     // always 1 in planet SPK ???
	TypeOfData int //: = 0;  // always 2 in planet SPK

	RecordStartAddress int //:= -1; // counted in elements, one need to multiply this by 8 to obtain adress in file
	RecordLastAddress  int //:= -1;
}

type SummaryRecordStruct struct {
	TotalSummariesNumber int
	NextRecordNumber     int
	PreviousRecordNumber int
}

type FileRecordStruct struct {

	//start of the file
	// данные в комментариях написаны для файла de430.bsp

	// 1. LOCIDW (8 characters, 8 bytes) An identification word (`DAF/SPK') 7+'\0'[Address 0]
	Locidw string //= "";

	// 2. ND(1 integer, 4 bytes) : The number of double prec. components in each array summary.[Address 8] nd = 2;
	Nd int

	// 3. NI ( 1 integer, 4 bytes): The number of integer components in each array summary. [Address 12] ni = 6;
	Ni int

	// 4. LOCIFN (60 characters, 60 bytes):
	// The internal name or description of the array file. 7+'\0' NIO2SPK
	Locifn string //= "";

	// 5. FWARD ( 1 integer, 4 bytes): The record number of the initial summary record in the file. [Address 76] fward = 4
	Fward int //= 0;

	// 6. BWARD ( 1 integer, 4 bytes): The record number of the final summary record in the file. [Address 80] bward = 4;
	Bward int //= 0;

	// 7. FREE(1 integer, 4 bytes) :
	// The first free address in the file.This is the address at which
	// the first element of the next array to be added to the file will be stored. free = 14967465;
	Free int //= 0;

	// 8. LOCFMT(8 characters, 8 bytes) :
	// The character string that indicates the numeric binary format of the DAF.
	// The string has value "LTL-IEEE" 8+'\0' одна буква не влезает
	// переделывать структуру не хочется. и так сойдет :) LTL-IEEE
	Locfmt string //= "";

	// 10. FTPSTR(28 characters, 28 bytes) : The FTP validation string.
	// ftpstr : "FTPSTR:\r:\n:\r\n:\r\x00:\x81:\x10\xce:ENDFTP",
}

type Position struct {
	X float64
	Y float64
	Z float64

	VelocityX float64
	VelocityY float64
	VelocityZ float64
}

type PolarPosition struct {
	Longitude float64
	Latitude  float64
	Radius    float64

	VelocityX float64
	VelocityY float64
	VelocityZ float64
}

// months from 1 to 12
var MonthsArr = [13]string{"", "January", "February", "March", "April", "May", "June", "July", "August", "September", "October", "November", "December"}

// zodiac names from 0 to 11
var ZodiacNames = [13]string{
	"",
	"Aries",
	"Taurus",
	"Gemini",
	"Cancer",
	"Leo",
	"Virgo",
	"Libra",
	"Scorpio",
	"Sagittarius",
	"Capricorn",
	"Aquarius",
	"Pisces"}

func GetName(targetCode int) string {

	NaifCodes := map[int]string{
		0:       "SSB",
		1:       "MERCURY_BARYCENTER", //'MERCURY_BARYCENTER'  'MERCURY BARYCENTER'
		2:       "VENUS_BARYCENTER",   //        'VENUS_BARYCENTER' 'VENUS BARYCENTER'
		3:       "EARTH_BARYCENTER",   // 'EARTH_BARYCENTER' 'EMB' 'EARTH MOON BARYCENTER' 'EARTH-MOON BARYCENTER' 'EARTH BARYCENTER'
		4:       "MARS_BARYCENTER",    //          'MARS_BARYCENTER' 'MARS BARYCENTER'
		5:       "JUPITER_BARYCENTER", //          'JUPITER_BARYCENTER' 'JUPITER BARYCENTER'
		6:       "SATURN_BARYCENTER",  //          'SATURN_BARYCENTER' 'SATURN BARYCENTER'
		7:       "URANUS_BARYCENTER",  //          'URANUS_BARYCENTER' 'URANUS BARYCENTER'
		8:       "NEPTUNE_BARYCENTER", //         'NEPTUNE_BARYCENTER' 'NEPTUNE BARYCENTER'
		9:       "PLUTO_BARYCENTER",   //          'PLUTO_BARYCENTER' 'PLUTO BARYCENTER'
		10:      "SUN",                //         'SUN'
		199:     "MERCURY",            // 'MERCURY'
		299:     "VENUS",              // 'VENUS'
		399:     "EARTH",              // 'EARTH'
		301:     "MOON",               // 'MOON'
		499:     "MARS",               // 'MARS'
		599:     "JUPITER",            // 'JUPITER'
		699:     "SATURN",             // 'SATURN'
		799:     "URANUS",             // 'URANUS'
		899:     "NEPTUNE",            // 'NEPTUNE'
		999:     "PLUTO",              // 'PLUTO'
		2002060: "CHIRON",
	}

	if val, ok := NaifCodes[targetCode]; ok {
		return val
	} else {
		return "Unknown Object"
	}

}

func De440sFile() *FileInfo {
	//заполняем согласно спецификации файла de430.bsp_auto_description.txt

	var fi FileInfo

	fi.SummaryRecordStruct.TotalSummariesNumber = 14
	fi.SummaryRecordStruct.NextRecordNumber = 0
	fi.SummaryRecordStruct.PreviousRecordNumber = 0

	fi.SummariesLineStruct = make([]SummariesLines, fi.SummaryRecordStruct.TotalSummariesNumber+1)
	for i := 0; i <= fi.SummaryRecordStruct.TotalSummariesNumber; i++ {

		var summariesLineStructTmp SummariesLines

		//0-ую позицию заполняем SOLAR SYSTEM BARYCENTER, везде 0
		if i == 0 {

			summariesLineStructTmp.SEGMENT_START_TIME = 0.0 // always
			summariesLineStructTmp.SEGMENT_LAST_TIME = 0.0  // always
			summariesLineStructTmp.TargetCode = 0
			summariesLineStructTmp.CenterCode = 0
			summariesLineStructTmp.RefFrame = 0   // always ?
			summariesLineStructTmp.TypeOfData = 0 // always
			summariesLineStructTmp.RecordStartAddress = 0
			summariesLineStructTmp.RecordLastAddress = 0

			summariesLineStructTmp.Name = GetName(summariesLineStructTmp.TargetCode)
			summariesLineStructTmp.Number = i

			fi.SummariesLineStruct[i] = summariesLineStructTmp

			continue
		}

		// предполагаем что SEGMENT_START_TIME, SEGMENT_LAST_TIME всегда одни и те же для всего файла

		summariesLineStructTmp.SEGMENT_START_TIME = -4734072000.0 // start time of segment in our case only 1 segment os of the whole file
		summariesLineStructTmp.SEGMENT_LAST_TIME = 4735368000.0   // end time of segment in our case only 1 segment os of the whole file

		//всегда refFrame = 1 ,  typeOfData = 2

		summariesLineStructTmp.RefFrame = 1
		summariesLineStructTmp.TypeOfData = 2

		summariesLineStructTmp.Number = i

		switch i {

		case 1:

			// MERCURY BARYCENTER
			summariesLineStructTmp.TargetCode = 1
			summariesLineStructTmp.CenterCode = 0
			summariesLineStructTmp.RecordStartAddress = 8065
			summariesLineStructTmp.RecordLastAddress = 610868

		case 2:
			// VENUS BARYCENTER

			summariesLineStructTmp.TargetCode = 2
			summariesLineStructTmp.CenterCode = 0
			summariesLineStructTmp.RecordStartAddress = 610869
			summariesLineStructTmp.RecordLastAddress = 830072

		case 3:
			// EARTH-MOON BARY
			summariesLineStructTmp.TargetCode = 3
			summariesLineStructTmp.CenterCode = 0
			summariesLineStructTmp.RecordStartAddress = 830073
			summariesLineStructTmp.RecordLastAddress = 1110926

		case 4:
			// MARS BARYCENTER

			summariesLineStructTmp.TargetCode = 4
			summariesLineStructTmp.CenterCode = 0
			summariesLineStructTmp.RecordStartAddress = 1110927
			summariesLineStructTmp.RecordLastAddress = 1230805

		case 5:
			// JUPITER BARYCENTER
			summariesLineStructTmp.TargetCode = 5
			summariesLineStructTmp.CenterCode = 0
			summariesLineStructTmp.RecordStartAddress = 1230806
			summariesLineStructTmp.RecordLastAddress = 1319859

		case 6:
			// SATURN BARYCENTER
			summariesLineStructTmp.TargetCode = 6
			summariesLineStructTmp.CenterCode = 0
			summariesLineStructTmp.RecordStartAddress = 1319860
			summariesLineStructTmp.RecordLastAddress = 1398638

		case 7:
			// URANUS BARYCENTER
			summariesLineStructTmp.TargetCode = 7
			summariesLineStructTmp.CenterCode = 0
			summariesLineStructTmp.RecordStartAddress = 1398639
			summariesLineStructTmp.RecordLastAddress = 1467142

		case 8:
			// NEPTUNE BARYCENTER
			summariesLineStructTmp.TargetCode = 8
			summariesLineStructTmp.CenterCode = 0
			summariesLineStructTmp.RecordStartAddress = 1467143
			summariesLineStructTmp.RecordLastAddress = 1535646

		case 9:
			// PLUTO BARYCENTER
			summariesLineStructTmp.TargetCode = 9
			summariesLineStructTmp.CenterCode = 0
			summariesLineStructTmp.RecordStartAddress = 1535647
			summariesLineStructTmp.RecordLastAddress = 1604150

		case 10:
			// SUN
			summariesLineStructTmp.TargetCode = 10
			summariesLineStructTmp.CenterCode = 0
			summariesLineStructTmp.RecordStartAddress = 1604151
			summariesLineStructTmp.RecordLastAddress = 1843904

		case 11:
			// MOON
			summariesLineStructTmp.TargetCode = 301
			summariesLineStructTmp.CenterCode = 3
			summariesLineStructTmp.RecordStartAddress = 1843905
			summariesLineStructTmp.RecordLastAddress = 2967308

		case 12:
			// EARTH
			summariesLineStructTmp.TargetCode = 399
			summariesLineStructTmp.CenterCode = 3
			summariesLineStructTmp.RecordStartAddress = 2967309
			summariesLineStructTmp.RecordLastAddress = 4090712

		case 13:
			// MERCURY

			summariesLineStructTmp.TargetCode = 199
			summariesLineStructTmp.CenterCode = 1
			summariesLineStructTmp.RecordStartAddress = 4090713
			summariesLineStructTmp.RecordLastAddress = 4090724

		case 14:
			// VENUS

			summariesLineStructTmp.TargetCode = 299
			summariesLineStructTmp.CenterCode = 2
			summariesLineStructTmp.RecordStartAddress = 4090725
			summariesLineStructTmp.RecordLastAddress = 4090736

		default:

			fmt.Printf("De430File defualt switch happend")
		}

		summariesLineStructTmp.Name = GetName(summariesLineStructTmp.TargetCode)

		fi.SummariesLineStruct[i] = summariesLineStructTmp

	}

	return &fi
}
