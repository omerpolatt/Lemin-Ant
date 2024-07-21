package errorslemin

import (
	"strconv"
	"strings"
)

func Mergeprocess(arr []string) string {
	if StartandEndControl(arr) != "" { // checks the existence of start and end rooms ( start ve end odalarının olup olmadığını kontrol eder )
		return StartandEndControl(arr)
	}
	if NotRepeatRoom(arr) != "" { // Controls the repetition of room names  ( Oda isimlerinin tekrar etmesini kontrol eder )
		return NotRepeatRoom(arr)
	}
	if UnknownRoom(arr) != "" { // isimsiz Oda Kontrolü ( Checks Unkown Room )
		return UnknownRoom(arr)
	}
	if CoordinateControl(arr) != "" { //  Checks if the coordinates are valid numbers ( Koordinatların geçerli sayılar olup olmaıdğını kontrol eder )
		return CoordinateControl(arr)
	}
	if RoomFirstLetterandSpaceCheck(arr) != "" { // Checks Room Name (Oda isimlerini Kontrol eder )
		return RoomFirstLetterandSpaceCheck(arr)
	}
	if HashtagControl(arr) != "" {
		return HashtagControl(arr)
	}
	if IsThereDoubleHashesMoreThanTwo(arr) != "" {
		return IsThereDoubleHashesMoreThanTwo(arr)
	}
	if RevLinkControl(arr) != "" {
		return RevLinkControl(arr)
	}
	if ConnectionControlEnd(arr) != "" {
		return ConnectionControlEnd(arr)
	}
	if AntCountControl(arr) != "" {
		return AntCountControl(arr)
	}

	return ""
}

func containstring(arr []string, val string) bool {
	for _, value := range arr {
		if value == val {
			return true
		}
	}
	return false
}

func StartandEndControl(arr []string) string {
	var new string
	if !containstring(arr, "##start") || !containstring(arr, "##end") {
		new = "Error! ##start or ##end is not in the file"
		return new
	}
	return new
}

func takefirstwords(line string) string {
	res := ""
	for _, char := range line {
		if char != ' ' {
			res += string(char)
		} else {
			break
		}
	}
	return res
}

func NotRepeatRoom(lines []string) string {
	var res string
	var roomNames []string

	for _, line := range lines {
		values := strings.Split(line, " ")
		if len(values) >= 3 {
			firstValue := takefirstwords(line)
			if containstring(roomNames, firstValue) {
				res = "Error! there is a duplicate room"
				return res
			}
			roomNames = append(roomNames, firstValue)
		}
	}
	return res
}

func UnknownRoom(lines []string) string {
	var rooms []string
	var links []string
	var unknown []string
	var res string
	for _, item := range lines {
		values := strings.Split(item, " ")
		if len((values)) >= 3 {
			rooms = append(rooms, takefirstwords(item))
		} else if len(values) == 1 && strings.Contains(item, "-") {
			links = append(links, item)
		}
	}
	for _, link := range links {
		connectedrooms := strings.Split(link, "-")
		if !containstring(rooms, connectedrooms[0]) || !containstring(rooms, connectedrooms[1]) {
			unknown = append(unknown, link)
		}
	}
	if len(unknown) > 0 {
		res = "Error! unknown and undefined room name"
		return res
	}
	return res
}

func CoordinateControl(lines []string) string {
	var control string
	for _, item := range lines {
		values := strings.Split(item, " ")
		if len(values) == 3 {
			x, err1 := strconv.Atoi(values[1])
			y, err2 := strconv.Atoi(values[2])
			if err1 != nil || err2 != nil || x < 0 || y < 0 {
				control = "Error! invalid coordinate definition"
				return control
			}
		}
	}
	return control
}

func RoomFirstLetterandSpaceCheck(lines []string) string {
	var res string

	for _, line := range lines {
		seperation := strings.Split(line, " ")
		if len(seperation) == 3 {
			for i := 0; i <= len(takefirstwords(line))-1; i++ {
				if line[0] == '#' || line[0] == 'L' {
					res = "Error! room name cannot start with # and L"
					return res
				}
			}
		}
		if len(seperation) == 4 {
			res = "Error! room name must not contain spaces"
			return res
		}
	}
	return res
}

// İkili hastaglerden sadece iki tane olup olmadığını kontrol et.
// tek hastaglelileri arrayden temizle.(yani yorumları temizle)
func HashtagControl(lines []string) string {
	var res string
	for _, line := range lines {
		hashCount := 0
		for _, char := range line {
			if char == '#' {
				hashCount++
			}
		}
		if hashCount > 2 {
			res = "There are hashes more than two!"
			return res
		}
	}
	return res
}

func IsThereDoubleHashesMoreThanTwo(lines []string) string {
	var res string
	doubleHashesCount := 0
	for _, line := range lines {
		if strings.HasPrefix(line, "##") {
			doubleHashesCount++
		}
	}

	if doubleHashesCount != 2 {
		res = "There must be only two double hashes!"
		return res
	}
	return res
}

func ClearTheComments(lines []string) []string {
	var res []string

	for _, line := range lines {
		if line[0] == '#' && line[1] != '#' {
			continue
		} else {
			res = append(res, line)
		}
	}
	return res
}

func RevLinkControl(lines []string) string {
	var res string
	for i := 0; i < len(lines); i++ {
		if IsItLink(lines[i]) {
			revLink := RevLink(lines[i])
			for j := i + 1; j < len(lines); j++ {
				if revLink == lines[j] {
					res = "Two rooms can only connect with one tunnel!" + " " + lines[i] + " " + lines[j]
					return res
				}
			}
		}
	}
	return res
}

func IsItLink(line string) bool {
	for i := 0; i < len(line); i++ {
		if line[i] == '-' && line[i-1] != ' ' && line[i+1] != ' ' {
			return true
		}
	}
	return false
}

func RevLink(str string) string {
	res := ""
	room := ""
	for _, char := range str {
		if char == '-' {
			res += room
			room = ""
		} else {
			room += string(char)
		}
	}
	return room + "-" + res
}

func ConnectionControlEnd(lines []string) string {
	var bağlanti []string
	var end []string
	var res string
	for _, item := range lines {
		if strings.Contains(item, "-") {
			bağlanti = append(bağlanti, item)
		}
	}
	for _, r := range bağlanti {
		for i := 0; i <= len(lines)-1; i++ {
			if lines[i] == "##end" {
				endcontrol := takefirstwords(lines[i+1])
				values := strings.Split(r, "-")
				if endcontrol == values[1] || endcontrol == values[0] {
					end = append(end, r)
				}
			}
		}
	}
	if len(end) == 0 {
		res = "Error! no connection to end!"
		return res
	}
	return res
}

func AntCountControl(lines []string) string {
	var res string
	value, _ := strconv.Atoi(lines[0])
	if value == 0 {
		res = "Error, invalid ant count!"
		return res
	}
	return res
}
