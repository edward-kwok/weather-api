package internal

import (
	"bufio"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

//Extract a specific date weather data
func Extract(s string) string {
	c := colly.NewCollector()
	t := time.Now()
	//file, err := os.Create("result.csv")
	//defer file.Close()
	//writer := csv.NewWriter(file)
	x := ""
	//checkError("Cannot create file", err)
	//writer.Comma = ','
	c.OnHTML("body > table.skin_main_table_table01_class > tbody > tr:nth-child(1) > td.skin_main_table_td02_class > table > tbody > tr > td:nth-child(2) > pre", func(e *colly.HTMLElement) {
		scanner := bufio.NewScanner(strings.NewReader(e.Text))
		n := 0
		for scanner.Scan() {
			if n > 15 && n < 41 {
				fmt.Println(scanner.Text())
				var data []string
				var re = regexp.MustCompile(`(?m)(\w+.*[^C\s])\s+(\d+.\d) C\s+(\d+.\d)`)
				data = append(data, s)
				for i, match := range re.FindStringSubmatch(scanner.Text()) {
					if i > 0 {
						data = append(data, match)
					}

				}
				if len(data) > 1 {
					x = x + fmt.Sprintf("%s,%s,%s,%s\n", data[0], data[1], data[2], data[3])
					//checkError("Canot write to file", err)
				}

				//writer.Flush()
			}
			n++
		}

	})
	t = t.Add(time.Hour * -24)
	link := fmt.Sprintf(`https://www.hko.gov.hk/wxinfo/dailywx/yeswx/ryese%s.htm`, s)
	fmt.Println(link)
	c.Visit(link)
	c.Wait()
	return x

}

func checkError(message string, err error) {
	if err != nil {
		log.Fatal(message, err)
	}
}
