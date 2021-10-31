package checks

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

type HTest struct {
	DirPattern    string
	FilePattern   []string
	IgnorePattern []string
	FileExts      []string
	BlackList     []string
	Comment       string
	URI           string
}

// DoChecks needs to read the config file (in memeory struct, don't really read
// the file each time) and then return true if the file is valid for the
// defined heuristics in the config file
//
func DoChecks(r []string) error {

	log.Println("Reading config and doing checks")

	t := "Exclude" // by default..  we don't know what the file is  (could also return an error type for this)
	uri := ""
	tests := CSDCOHTs()

	d := "./mnt/wdb"
	proj := "YUKON"

	for _, f := range r {
		dir, _ := filepath.Split(f)

		for i := range tests {
			if caselessPrefix(d, proj, dir, tests[i].DirPattern) {
				// if caselessContains(dir, tests[i].DirPattern) { // TODO should become caselessPrefix(d, proj, dir, tests[i].DirPattern)
				if fileInDir(d, proj, tests[i].DirPattern, f) {
					if caselessContainsSlice(f, tests[i].FilePattern) {
						fileext := strings.ToLower(filepath.Ext(f))
						s := tests[i].FileExts
						if contains(s, fileext) {
							// fmt.Printf("%s == %s\n", f, tests[i].Comment) //  TODO  all NewFileEntry calls should use class URI, not name like "Images"
							t = tests[i].Comment
							uri = tests[i].URI
						}
					}
				}
			}

			log.Printf("File %s is exclude value %s with uri: %s", f, t, uri)
		}

	}

	return nil
}

func fileInDir(d, proj, dp, f string) bool {
	a := fmt.Sprintf("%s/%s/%s", d, proj, dp)
	b := fmt.Sprintf("%s/", filepath.Dir(f))

	i := strings.Compare(a, b)
	e := false
	if i == 0 {
		e = true
	}

	// fmt.Printf("%t :: fileInDir: %s is in %s, %b \n", e, f, a, b)

	return e
}

func caselessContainsSlice(a string, b []string) bool {
	t := true // default to true so that 0 len string array is NOT a test.
	for i := range b {
		t = strings.Contains(strings.ToUpper(a), strings.ToUpper(b[i]))
		// 	fmt.Printf("CCS Tested %s against %s and got %t\n", a, b[i], t)
	}
	// fmt.Printf("CSS called, returning %t for %s \n", t, a)
	return t
}

func caselessContains(a, b string) bool {
	return strings.Contains(strings.ToUpper(a), strings.ToUpper(b))
}

// Test if a has prefix b    /dir1/dir2/dir3/filex  has /dir1/dir2
// To do this I need the base directory to remove from a
func caselessPrefix(base, proj, a, b string) bool {
	pref := fmt.Sprintf("%s/%s/", base, proj)
	// fmt.Printf("Directory Test: %s\n", pref)
	// fmt.Printf("Directory Test: %s\n", a)
	atl := strings.TrimPrefix(a, pref)
	// fmt.Printf("Directory Test: %s\n", atl)
	// fmt.Printf("Directory Test: In base %s in proj %s test if  %s has prefix %s result: %t \n", base, proj, strings.ToUpper(atl), strings.ToUpper(b),
	// strings.HasPrefix(strings.ToUpper(atl), strings.ToUpper(b)))
	return strings.HasPrefix(strings.ToUpper(atl), strings.ToUpper(b))
}

func contains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

// ageInYears gets the age of a file as a float64 decimal value
func ageInYears(fp string) (float64, time.Time) {
	fi, err := os.Stat(fp)
	if err != nil {
		fmt.Println(err)
	}
	stat := fi.Sys().(*syscall.Stat_t)
	// ctime := time.Unix(int64(stat.Ctim.Sec), int64(stat.Ctim.Nsec))
	ctime := time.Unix(int64(stat.Mtim.Sec), int64(stat.Mtim.Nsec))
	delta := time.Now().Sub(ctime)
	years := delta.Hours() / 24 / 365
	// fmt.Printf("Create: %v   making it %.2f  years old\n", ctime, years)
	return round2(years, 0.01), ctime
}

func round2(x, unit float64) float64 {
	if x > 0 {
		return float64(int64(x/unit+0.5)) * unit
	}
	return float64(int64(x/unit-0.5)) * unit
}

// CSDCOHTs a set of tests to do on directory and file path/extensions.
func CSDCOHTs() []HTest {
	ht := []HTest{
		// HTest{DirPattern: "/Data/Sampling",
		// 	FilePattern: []string{"SRF"},
		// 	FileExts:    []string{""},
		// 	BlackList:   []string{},
		// 	Comment:     "SRF",
		// 	URI:         "http://opencoredata.org/voc/csdco/v1/SRF"},
		HTest{DirPattern: "Data/Corelyzer/",
			FilePattern: []string{},
			FileExts:    []string{".cml", ".xml"},
			BlackList:   []string{},
			Comment:     "Corelyzer files",
			URI:         "http://opencoredata.org/voc/csdco/v1/CML"},
		HTest{DirPattern: "Data/Corelyzer/",
			FilePattern: []string{},
			FileExts:    []string{".car"},
			BlackList:   []string{},
			Comment:     "Corelyzer archive files",
			URI:         "http://opencoredata.org/voc/csdco/v1/Car"},
		HTest{DirPattern: "Images/",
			FilePattern:   []string{},
			FileExts:      []string{".bmp", ".jpeg", ".jpg", "tif", "tiff"},
			IgnorePattern: []string{"tiff"},
			BlackList:     []string{},
			Comment:       "Images",
			URI:           "http://opencoredata.org/voc/csdco/v1/Image"},
		HTest{DirPattern: "Images/rgb/",
			FilePattern: []string{},
			FileExts:    []string{".csv"},
			BlackList:   []string{},
			Comment:     "RGB Image Data",
			URI:         "http://opencoredata.org/voc/csdco/v1/RGBData"},
		HTest{DirPattern: "MSCL/MSCL-S/",
			FilePattern:   []string{"_MSCL-S"},
			IgnorePattern: []string{"Other data"},
			FileExts:      []string{".xls", ".xlsx", ".csv"}, // what is the point of a black list?  I only validate on FileExts found???
			BlackList:     []string{".raw", ".dat", ".out", ".cal"},
			Comment:       "Geotek MSCL-S",
			URI:           "http://opencoredata.org/voc/csdco/v1/WholeCoreData"},
		HTest{DirPattern: "MSCL/MSCL-S_split/",
			FilePattern:   []string{"_MSCL-S_split"},
			IgnorePattern: []string{"Other data"},
			FileExts:      []string{".xls", ".xlsx", ".csv"}, // what is the point of a black list?  I only validate on FileExts found???
			BlackList:     []string{".raw", ".dat", ".out", ".cal"},
			Comment:       "Geotek MSCL-S Split-core",
			URI:           "http://opencoredata.org/voc/csdco/v1/WholeCoreData"},
		HTest{DirPattern: "MSCL/MSCL-XYZ/",
			FilePattern:   []string{"_MSCL-XYZ"},
			IgnorePattern: []string{"Other data"},
			FileExts:      []string{".xls", ".xlsx", ".csv"}, // what is the point of a black list?  I only validate on FileExts found???
			BlackList:     []string{".raw", ".dat", ".out", ".cal"},
			Comment:       "Geotek MSCL-XYZ",
			URI:           "http://opencoredata.org/voc/csdco/v1/SplitCoreData"},
		HTest{DirPattern: "ICD/",
			FilePattern: []string{},
			FileExts:    []string{".pdf"},
			BlackList:   []string{},
			Comment:     "ICD",
			URI:         "http://opencoredata.org/voc/csdco/v1/ICDFiles"},
		HTest{DirPattern: "ICD/",
			FilePattern: []string{"ICD_tabular"},
			FileExts:    []string{".xls", ".xlsx", ".csv"},
			BlackList:   []string{},
			Comment:     "ICD",
			URI:         "http://opencoredata.org/voc/csdco/v1/ICDFiles"}}

	return ht
}
