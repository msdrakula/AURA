package intel

// Catalog is the mindmap order shown in the UI.
func Catalog() []CatalogStage {
	return []CatalogStage{
		{
			ID:    "subdomains_passive",
			Title: "Поддомены",
			Hint:  "Пассивный поиск имён в открытых сертификатах (crt.sh).",
			Mode:  ModePassive,
		},
		{
			ID:    "dns_brute",
			Title: "DNS-имена",
			Hint:  "Короткий список типичных префиксов (www, api, staging…). Нужно согласие.",
			Mode:  ModeActive,
		},
		{
			ID:    "live_hosts",
			Title: "Живые хосты",
			Hint:  "HTTP/HTTPS-проверка найденных имён.",
			Mode:  ModeActive,
		},
		{
			ID:    "web_ports",
			Title: "Веб-порты",
			Hint:  "Проверка популярных веб-портов на живых хостах, не полное сканирование сети.",
			Mode:  ModeActive,
		},
		{
			ID:    "tech",
			Title: "Технологии",
			Hint:  "Стек по заголовкам и HTML живых ответов.",
			Mode:  ModePassive,
		},
		{
			ID:    "urls_passive",
			Title: "Старые URL",
			Hint:  "Публичный архив Wayback Machine — уже известные пути.",
			Mode:  ModePassive,
		},
		{
			ID:    "scrape",
			Title: "HTML и JavaScript",
			Hint:  "Ссылки, API-пути и скрипты со стартовой страницы.",
			Mode:  ModeActive,
		},
		{
			ID:    "dirs",
			Title: "Скрытые пути",
			Hint:  "Перебор короткого встроенного словаря директорий (как gobuster dir).",
			Mode:  ModeActive,
		},
		{
			ID:    "params",
			Title: "Параметры",
			Hint:  "Поиск типичных query-параметров на базовом URL (как ffuf ?FUZZ=).",
			Mode:  ModeActive,
		},
	}
}

func stageByID(id string) (CatalogStage, bool) {
	for _, s := range Catalog() {
		if s.ID == id {
			return s, true
		}
	}
	return CatalogStage{}, false
}
