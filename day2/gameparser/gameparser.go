package gameparser

import (
	"log"
	"strconv"
	"strings"

	"adventofcode2023/day2/model"
)

func parseDraw(strDraw string) (draw *model.Draw) {
	draw = &model.Draw{}
	cubes := strings.Split(strDraw, ",")
	for _, cube := range cubes {
		trimmedCube := strings.Trim(cube, " ")
		cubeProperties := strings.Split(trimmedCube, " ")
		n, err := strconv.Atoi(strings.Trim(cubeProperties[0], " "))
		if err != nil {
			log.Fatalln(err)
		}
		color := cubeProperties[1]
		switch color {
		case "red":
			draw.Red = n
		case "green":
			draw.Green = n
		case "blue":
			draw.Blue = n
		}
	}
	return draw
}

func parseDraws(line string) (draws []model.Draw) {
	strDraws := strings.Split(line, ";")
	for _, draw := range strDraws {
		draws = append(draws, *parseDraw(draw))
	}
	return draws
}

func parseGame(line string) *model.Game {
	gameProperties := strings.Split(line, ":")
	gameId, err := strconv.Atoi(strings.Split(gameProperties[0], " ")[1])
	if err != nil {
		log.Fatalln(err)
	}
	draws := parseDraws(gameProperties[1])
	return &model.Game{
		Id:    gameId,
		Draws: draws,
	}
}

func ParseGames(lines []string) []model.Game {
	var games []model.Game
	for _, line := range lines {
		games = append(games, *parseGame(line))
	}
	return games
}
