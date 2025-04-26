package zipcode

//testing
import (
	"bytes"
	"fmt"
	"os/exec"
	"strconv"
	"strings"
	"testing"

	"github.com/monopolly/file"
	//testing
	//go test -bench=.
	//go test --timeout 9999999999999s
)

func TestConvert(u *testing.T) {
	__(u)

	return

	// var list []string

	var res bytes.Buffer

	var c int
	res.WriteString("package zipcode\n")
	res.WriteString("var list = []Code{\n")
	// "zip","lat","lng","city","state_id","state_name","zcta","parent_zcta","population","density","county_fips","county_name","county_weights","county_names_all","county_fips_all","imprecise","military","timezone"
	file.CSV("uszips.csv", ',', func(line []string) (stop bool) {
		if len(line) == 0 {
			return
		}
		id, _ := strconv.Atoi(line[0])
		if id == 0 {
			return
		}

		// {"00501", 40.922326, -72.637078, "Holtsville", "NY", "Suffolk"},
		res.WriteString("{")
		p := strings.Join(
			[]string{
				fmt.Sprintf(`"%s"`, line[0]),
				line[1],
				line[2],
				fmt.Sprintf(`"%s"`, line[4]),
				fmt.Sprintf(`"%s"`, line[5]),
				fmt.Sprintf(`"%s"`, line[11]),
			},
			",",
		)
		res.WriteString(p)
		res.WriteString("},\n")
		c++
		// if c == 10 {
		// 	return true
		// }
		return
	})
	res.WriteString("}")

	// var res bytes.Buffer
	// pp := csv.NewWriter(&res)

	// pp.WriteAll(list)

	// file.Save("codes2025.csv", res.Bytes())
	file.Save("codes2025.go", res.Bytes())

}
func TestMain(u *testing.T) {
	__(u)

	a := New()

	from := 90210
	to := 90302

	fmt.Println(a.Get(from))
	fmt.Println(a.Get(to))

	fmt.Println(a.DirectMiles(from, to))
	fmt.Println(a.DriveMiles(from, to))

}

func Benchmark1(u *testing.B) {
	u.ReportAllocs()
	u.ResetTimer()
	for n := 0; n < u.N; n++ {

	}
}

func Benchmark2(u *testing.B) {
	u.RunParallel(func(pb *testing.PB) {
		for pb.Next() {

		}
	})
}

func __(u *testing.T) {
	fmt.Printf("\033[1;32m%s\033[0m\n", strings.ReplaceAll(u.Name(), "Test", ""))
}

func cmd(name string, v ...string) {
	c := exec.Command(name, v...)
	r, err := c.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println(string(r))
}
