package phonescript
/*
	@ Commands: Please read "README.md"
*/

import (
	"fmt" 
	"strings"
	"strconv"
	"math/rand"
	"regexp"
	"github.com/nyaruka/phonenumbers" 
	"time"
	"encoding/csv"
	"path/filepath"
	"log"
	"os"
)


/*
	@ Generates random US phone numbers
	@ Saves to file
*/
func GenerateNumbers(limit string) string {
	var phoneNumbersOutput []string
	r1 := rand.New(rand.NewSource(time.Now().UnixNano()))
	err := os.Mkdir("Records", 0750)
	if err != nil && !os.IsExist(err){
		log.Fatal(err)
	}

	// make filename random using unixtimestamp
	randFileName := filepath.FromSlash("Records/PhoneNumbersGenerated_" + strconv.Itoa(r1.Intn(1000000000)) + "_.csv")
	csvFile, err := os.Create(randFileName)

	defer csvFile.Close()
	if err != nil{
		log.Fatal(err)
	}

	w := csv.NewWriter(csvFile)
	defer w.Flush()
	limits, _ := strconv.Atoi(limit)
	for i := 0; i <= limits; i++{
		number := rangeIn(100000000, 9999999999)
		parsedNumber, _ := phonenumbers.Parse(fmt.Sprintf("%s", number), "US")
		phoneNumbersOutput = append(phoneNumbersOutput, strings.Replace(fmt.Sprintf("%s",parsedNumber),"country_code:1 national_number:", "+1", -1))
		
	}
	w.Write(phoneNumbersOutput)
	return "Numbers saved to " + randFileName
}

/*
	@ Validate phone numbers
	@ Saves to file
*/
func ValidateNumbers(filePath string) string {
	csvData := readCsvFile(filePath)
	for _, number := range csvData[0]{
		fmt.Println(number, phoneLookup(number))
	}
	return "Done!"
}

/*
	@ Internal Method
*/
func phoneLookup(number string)string{
	re := regexp.MustCompile(`^(?:(?:\(?(?:00|\+)([1-4]\d\d|[1-9]\d?)\)?)?[\-\.\ \\\/]?)?((?:\(?\d{1,}\)?[\-\.\ \\\/]?){0,})(?:[\-\.\ \\\/]?(?:#|ext\.?|extension|x)[\-\.\ \\\/]?(\d+))?$`)
	if re.MatchString(number) == true{
		return "Valid"
	}
	return "Invalid"
}

/*
	@ Internal Method
*/
func readCsvFile(filePath string) [][]string {
    f, err := os.Open(filePath)
    if err != nil {
        log.Fatal("Unable to read input file " + filePath, err)
    }
    defer f.Close()

    csvReader := csv.NewReader(f)
    records, err := csvReader.ReadAll()
    if err != nil {
        log.Fatal("Unable to parse file as CSV for " + filePath, err)
    }

    return records
}


func rangeIn(low, hi int) int {
	return low + rand.Intn(hi - low)
}