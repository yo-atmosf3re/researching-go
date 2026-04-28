package main

import (
	"researching-go/lessons"
	"researching-go/pkg/logger"
)

func main() {
	worker := lessons.NewWorker("Martin", "IKEA", 25, true, 5.5)
	logger.Pt(worker)
	logger.Ptc(worker)
}
