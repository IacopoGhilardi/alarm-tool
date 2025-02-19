# Come Utilizzare DawnDash

## Guida sviluppo locale

### Docker
1. Clona il repository
2. Installa Docker
3. Naviga nella cartella `backend`
4. Copia il file `.env.example` in `.env` e configura le variabili d'ambiente
5. Avvia il progetto con `docker compose up --build`
6. Le tue API sono disponibili sulla porta `8080` e il database postgres sulla porta `5432`
7. Per stoppare il progetto `docker compose down`

### Backend (Go) (opzionale se vuoi eseguire il backend senza docker)
1. Clona il repository
2. Naviga nella cartella `backend`
3. Copia il file `.env.example` in `.env` e configura le variabili d'ambiente
4. Esegui `go mod download` per installare le dipendenze
5. Avvia il server con `go run cmd/dawndash-backend/main.go`

### App Mobile (Flutter)
1. Naviga nella cartella `mobile_app`
2. Esegui `flutter pub get` per installare le dipendenze
3. Avvia l'app con `flutter run`