defer file.Close()

mu.Lock()
defer.mu.Lock()

printHeader()
defer printHeader()

defer disconnectFromDB()