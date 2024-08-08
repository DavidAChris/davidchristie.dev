templ:
	@templ generate

tailwind:
	pnpm tailwindcss -i .\tailwind.css -o static/style.css

run: templ
	@go run .

dev:
	@templ generate --watch --proxy=http://localhost:3000 --cmd='go run .'

clean:
	@del /q /s *_templ.*

tidy:
	go mod tidy
	