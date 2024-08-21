package groza

import "log"

func debugPrint(content string) {
	log.Printf("【groza-DEBUG】 %s", content)
}

func debugPrintWARNINGDefault(content string) {
	log.Printf("【groza-DEBUG】  %s", content)
	return
}
