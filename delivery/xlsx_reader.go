package delivery

import (
	"archive/zip"
	"encoding/xml"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// readSheetCells membaca file .xlsx dan mengembalikan map koordinat sel (mis. "A3") -> nilai string,
// beserta nomor baris terbesar. Reader ini membaca setiap sel berdasarkan koordinatnya,
// sehingga tetap benar walau file punya struktur <row> yang tidak standar
// (mis. export TikTok/Tokopedia yang membungkus tiap sel dalam <row> terpisah).
func readSheetCells(path, sheetName string) (map[string]string, int, error) {
	zr, err := zip.OpenReader(path)
	if err != nil {
		return nil, 0, err
	}
	defer zr.Close()

	files := map[string]*zip.File{}
	for _, f := range zr.File {
		files[f.Name] = f
	}

	// shared strings (opsional, dipakai jika sel bertipe t="s")
	shared := []string{}
	if f, ok := files["xl/sharedStrings.xml"]; ok {
		shared = readSharedStrings(f)
	}

	target, err := resolveSheetTarget(files, sheetName)
	if err != nil {
		return nil, 0, err
	}

	f, ok := files[target]
	if !ok {
		return nil, 0, fmt.Errorf("worksheet %s tidak ditemukan", target)
	}

	rc, err := f.Open()
	if err != nil {
		return nil, 0, err
	}
	defer rc.Close()

	var sheet struct {
		Cells []struct {
			R string `xml:"r,attr"`
			T string `xml:"t,attr"`
			V string `xml:"v"`
			Is struct {
				T string `xml:"t"`
			} `xml:"is"`
		} `xml:"sheetData>row>c"`
	}
	if err := xml.NewDecoder(rc).Decode(&sheet); err != nil {
		return nil, 0, err
	}

	cells := map[string]string{}
	maxRow := 0
	for _, cell := range sheet.Cells {
		if cell.R == "" {
			continue
		}
		var val string
		switch cell.T {
		case "s":
			if idx, e := strconv.Atoi(strings.TrimSpace(cell.V)); e == nil && idx >= 0 && idx < len(shared) {
				val = shared[idx]
			}
		case "inlineStr":
			val = cell.Is.T
		default:
			val = cell.V
		}
		cells[cell.R] = val
		if r := rowOf(cell.R); r > maxRow {
			maxRow = r
		}
	}

	return cells, maxRow, nil
}

func readSharedStrings(f *zip.File) []string {
	rc, err := f.Open()
	if err != nil {
		return nil
	}
	defer rc.Close()

	var sst struct {
		SI []struct {
			T string `xml:"t"`
			R []struct {
				T string `xml:"t"`
			} `xml:"r"`
		} `xml:"si"`
	}
	if err := xml.NewDecoder(rc).Decode(&sst); err != nil {
		return nil
	}

	out := make([]string, 0, len(sst.SI))
	for _, si := range sst.SI {
		if len(si.R) > 0 {
			var b strings.Builder
			for _, run := range si.R {
				b.WriteString(run.T)
			}
			out = append(out, b.String())
		} else {
			out = append(out, si.T)
		}
	}
	return out
}

// resolveSheetTarget mencari path worksheet XML untuk sheetName tertentu
// via xl/workbook.xml (name -> r:id) dan xl/_rels/workbook.xml.rels (r:id -> target).
func resolveSheetTarget(files map[string]*zip.File, sheetName string) (string, error) {
	wbFile, ok := files["xl/workbook.xml"]
	if !ok {
		return "", fmt.Errorf("xl/workbook.xml tidak ditemukan")
	}

	var wb struct {
		Sheets []struct {
			Name string `xml:"name,attr"`
			RID  string `xml:"http://schemas.openxmlformats.org/officeDocument/2006/relationships id,attr"`
		} `xml:"sheets>sheet"`
	}
	if err := decodeZip(wbFile, &wb); err != nil {
		return "", err
	}

	rid := ""
	for _, s := range wb.Sheets {
		if s.Name == sheetName {
			rid = s.RID
			break
		}
	}
	if rid == "" {
		return "", fmt.Errorf("sheet %q tidak ada di workbook", sheetName)
	}

	relFile, ok := files["xl/_rels/workbook.xml.rels"]
	if !ok {
		return "", fmt.Errorf("workbook.xml.rels tidak ditemukan")
	}
	var rels struct {
		Rel []struct {
			ID     string `xml:"Id,attr"`
			Target string `xml:"Target,attr"`
		} `xml:"Relationship"`
	}
	if err := decodeZip(relFile, &rels); err != nil {
		return "", err
	}

	for _, r := range rels.Rel {
		if r.ID == rid {
			return normalizeTarget(r.Target), nil
		}
	}
	return "", fmt.Errorf("relationship %s tidak ditemukan", rid)
}

func normalizeTarget(t string) string {
	t = strings.TrimPrefix(t, "/")
	if strings.HasPrefix(t, "xl/") {
		return t
	}
	return "xl/" + t
}

func decodeZip(f *zip.File, v interface{}) error {
	rc, err := f.Open()
	if err != nil {
		return err
	}
	defer rc.Close()
	data, err := io.ReadAll(rc)
	if err != nil {
		return err
	}
	return xml.Unmarshal(data, v)
}

// rowOf mengambil nomor baris dari koordinat sel, mis. "AD12" -> 12.
func rowOf(ref string) int {
	i := 0
	for i < len(ref) && (ref[i] < '0' || ref[i] > '9') {
		i++
	}
	n, _ := strconv.Atoi(ref[i:])
	return n
}
