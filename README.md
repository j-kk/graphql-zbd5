### Projekt 5

Rozwiązanie napisane w golangu, zawiera 3 serwisy udostępniające api graphql
pod ścieżką /gql i prosty front pod ścieżką /

Obecnie do odpalenia należy dla każdego programu:
- ustawić pod zmienną środowiskową DATABASE_URL url do bazy danych postgresql (wraz z loginem/hasłem) (np: DATABASE_URL=postgresql://user:pass@localhost:5432/user)
- użyć skryptu create.sql do stworzenia tabel w bazie
- ustawić pod zmienną środowiskową PORT na jakim porcie mają być wystawione usługi (domyślnie 8080)

Działanie programów wg. suffixów: 
- user - dodawanie użytkowników
- ad - dodawanie reklam
- views - dodawanie informacji o wyswietleniu (i ew. odpytanie o użytkownika/reklamę)

Każdy z programów można zbudować za pomocą `go build server.go` i odpalić za pomocą 
`go run server.go`. Załączone dockerfile generują poprawne obrazy, a za pomocą docker-compose.yml
można uruchomić całą apkę (TODO).

Specyfikacje graphql znajdują się dla każdej aplikacji w pliku schema.graphqls

#### TODO
- program 4. do query
- docker-compose
- ujednolicenie bibliotek dostępowych do bazy