package main

import (
	"encoding/csv"
	"encoding/xml"
	"io"
	"os"
	"strconv"
	"strings"
)

// начало решения

type Organization struct {
	XMLName     xml.Name     `xml:"organization"`
	Departments []Department `xml:"department"`
}

// Department — элемент департамента
type Department struct {
	Code      string     `xml:"code"`
	Employees []Employee `xml:"employees>employee"` // Путь через вложенный тег <employees>
}

// Employee — элемент сотрудника
type Employee struct {
	ID     string `xml:"id,attr"` // Читаем атрибут id="..."
	Name   string `xml:"name"`
	City   string `xml:"city"`
	Salary int    `xml:"salary"`
}

// ConvertEmployees преобразует XML-документ с информацией об организации
// в плоский CSV-документ с информацией о сотрудниках
func ConvertEmployees(outCSV io.Writer, inXML io.Reader) error {

	var org Organization

	decoder := xml.NewDecoder(inXML)

	if err := decoder.Decode(&org); err != nil {
		return err
	}

	encoder := csv.NewWriter(outCSV)
	if err := encoder.Write([]string{"id", "name", "city", "department", "salary"}); err != nil {
		return err
	}

	for _, dept := range org.Departments {
		for _, emp := range dept.Employees {
			if err := encoder.Write([]string{emp.ID, emp.Name, emp.City, dept.Code, strconv.Itoa(emp.Salary)}); err != nil {
				return err
			}
		}
	}
	encoder.Flush()

	if err := encoder.Error(); err != nil {
		return err
	}

	return nil
}

// конец решения

func main() {
	src := `<organization>
    <department>
        <code>hr</code>
        <employees>
            <employee id="11">
                <name>Дарья</name>
                <city>Самара</city>
                <salary>70</salary>
            </employee>
            <employee id="12">
                <name>Борис</name>
                <city>Самара</city>
                <salary>78</salary>
            </employee>
        </employees>
    </department>
    <department>
        <code>it</code>
        <employees>
            <employee id="21">
                <name>Елена</name>
                <city>Самара</city>
                <salary>84</salary>
            </employee>
            <employee id="22">
                <name>Елена2</name>
            </employee>
        </employees>
    </department>
</organization>`

	in := strings.NewReader(src)
	out := os.Stdout
	ConvertEmployees(out, in)
	/*
		id,name,city,department,salary
		11,Дарья,Самара,hr,70
		12,Борис,Самара,hr,78
		21,Елена,Самара,it,84
	*/
}

// --------------------------
// (from solution):

// // Organization описывает организацию
// type Organization struct {
//     Departments []Department `xml:"department"`
// }

// // Department описывает департамент организации
// type Department struct {
//     Code      string     `xml:"code"`
//     Employees []Employee `xml:"employees>employee"`
// }

// // Employee описывает сотрудника департамента
// type Employee struct {
//     Id     int     `xml:"id,attr"`
//     Name   string  `xml:"name"`
//     City   string  `xml:"city"`
//     Salary float64 `xml:"salary"`
// }

// // decodeOrganization декодирует организацию из XML-документа
// func decodeOrganization(in io.Reader) (Organization, error) {
//     var org Organization
//     decoder := xml.NewDecoder(in)
//     err := decoder.Decode(&org)
//     return org, err
// }

// // employeeWriter записывает сотрудников в CSV
// type employeeWriter struct {
//     w   *csv.Writer
//     err error
// }

// // writeHeader записывает заголовок
// func (ew *employeeWriter) writeHeader() {
//     if ew.err != nil {
//         return
//     }
//     header := []string{"id", "name", "city", "department", "salary"}
//     ew.err = ew.w.Write(header)
// }

// // writeEmployee записывает сотрудника в строку
// func (ew *employeeWriter) writeEmployee(depCode string, emp Employee) {
//     if ew.err != nil {
//         return
//     }
//     fields := []string{
//         strconv.Itoa(emp.Id),
//         emp.Name,
//         emp.City,
//         depCode,
//         strconv.FormatFloat(emp.Salary, 'f', -1, 64),
//     }
//     ew.err = ew.w.Write(fields)
// }

// // flush финализирует данные
// func (ew *employeeWriter) flush() error {
//     ew.w.Flush()
//     if ew.err == nil {
//         ew.err = ew.w.Error()
//     }
//     return ew.err
// }

// // newEmployeeWriter создает нового писателя сотрудников в CSV
// func newEmployeeWriter(w io.Writer) *employeeWriter {
//     return &employeeWriter{w: csv.NewWriter(w)}
// }

// // ConvertEmployees преобразует XML-документ с информацией об организации
// // в плоский CSV-документ с информацией о сотрудниках
// func ConvertEmployees(outCSV io.Writer, inXML io.Reader) error {
//     org, err := decodeOrganization(inXML)
//     if err != nil {
//         return fmt.Errorf("failed to parse xml: %w", err)
//     }

//     w := newEmployeeWriter(outCSV)
//     w.writeHeader()

//     for _, dep := range org.Departments {
//         for _, emp := range dep.Employees {
//             w.writeEmployee(dep.Code, emp)
//         }
//     }

//     if err := w.flush(); err != nil {
//         return fmt.Errorf("failed writing csv: %w", err)
//     }

//     return nil
// }
