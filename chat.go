
func ChatLog(currentConn *WebSocketConnection, room string, messsage string) error {
	f, err := os.OpenFile(room, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)
	log.Println(fmt.Sprintf("%s : %s", currentConn.Username, messsage))
	return nil
}

func Readlog(room string) {
	file, err := os.OpenFile(room, os.O_RDWR, 0644)
	if err != nil {
		return
	}
	defer file.Close()

	text := make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			if err != nil {
				break
			}
		}
		if n == 0 {
			break
		}
	}
	if err != nil {
		return
	}

	fmt.Println("==> file berhasil dibaca")
	temp := string(text)
	var chatTemp interface{}
	chatTemp = temp

}
