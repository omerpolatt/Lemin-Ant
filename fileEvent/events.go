package events

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Room struct {
	Name        string
	Coordinates []int
}

type Connection struct {
	From string
	To   string
}

type ParseResult struct {
	AntCount    int
	StartRoom   Room
	EndRoom     Room
	Rooms       []Room
	Connections []Connection
}

func ParseInputFile(filename string) (ParseResult, error) {
	file, err := os.Open(filename)
	if err != nil {
		return ParseResult{}, fmt.Errorf("file not read: %w", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var antCount int
	var start Room
	var end Room
	var rooms []Room
	var connections []Connection

	firstLine := true
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "##comment") || strings.HasPrefix(line, "#comment") || strings.HasPrefix(line, "#another comment") {
			continue
		}

		if firstLine {
			antCount, err = strconv.Atoi(line)
			if err != nil {
				return ParseResult{}, fmt.Errorf("ant count not read: %w", err)
			}
			firstLine = false
			continue
		}

		if strings.Contains(line, "-") {
			parts := strings.Split(line, "-")
			connections = append(connections, Connection{From: parts[0], To: parts[1]})
		} else {
			if strings.HasPrefix(line, "##start") {
				scanner.Scan()
				start = ParseRoom(scanner.Text())
				rooms = append(rooms, start)
			} else if strings.HasPrefix(line, "##end") {
				scanner.Scan()
				end = ParseRoom(scanner.Text())
				rooms = append(rooms, end)
			} else {
				rooms = append(rooms, ParseRoom(line))
			}
		}
	}

	return ParseResult{
		AntCount:    antCount,
		StartRoom:   start,
		EndRoom:     end,
		Rooms:       rooms,
		Connections: connections,
	}, nil
}

func ParseRoom(line string) Room {
	parts := strings.Fields(line)
	name := parts[0]
	var coordinates []int
	for _, part := range parts[1:] {
		coordinate, _ := strconv.Atoi(part)
		coordinates = append(coordinates, coordinate)
	}
	return Room{Name: name, Coordinates: coordinates}
}

func BuildGraph(rooms []Room, connections []Connection) map[string]map[string]bool {
	graph := make(map[string]map[string]bool)
	for _, room := range rooms {
		graph[room.Name] = make(map[string]bool)
	}
	for _, conn := range connections {
		graph[conn.From][conn.To] = true
		graph[conn.To][conn.From] = true
	}
	return graph
}
