package intel

// Catalog is the mindmap order shown in the UI.
func Catalog() []CatalogStage {
	return []CatalogStage{
		{
			ID:    "subdomains_passive",
			Title: "Поддомены",
			Hint:  "crt.sh: имена из публичных сертификатов. Таймаут и лимит задаются в модуле.",
			Mode:  ModePassive,
		},
		{
			ID:    "dns_brute",
			Title: "DNS-имена",
			Hint:  "Активный перебор DNS-префиксов (www, api, staging…) или свой словарь.",
			Mode:  ModeActive,
		},
		{
			ID:    "live_hosts",
			Title: "Живые хосты",
			Hint:  "Проверка найденных имён по HTTP/HTTPS. Схемы и таймаут — в настройках.",
			Mode:  ModeActive,
		},
		{
			ID:    "web_ports",
			Title: "Веб-порты",
			Hint:  "Скан веб-портов на живых хостах. Список портов свой; «найдено» только если ответил HTTP.",
			Mode:  ModeActive,
		},
		{
			ID:    "tech",
			Title: "Технологии",
			Hint:  "Отпечатки стека: Server, cookies, HTML, известные CDN/auth. Таймаут на запрос.",
			Mode:  ModePassive,
		},
		{
			ID:    "urls_passive",
			Title: "Старые URL",
			Hint:  "Wayback CDX: исторические URL. По умолчанию только этот домен, без мусорных путей.",
			Mode:  ModePassive,
		},
		{
			ID:    "scrape",
			Title: "HTML и JavaScript",
			Hint:  "Ссылки и API из HTML/JS стартовой страницы. Чужие хосты и битые пути отбрасываются.",
			Mode:  ModeActive,
		},
		{
			ID:    "dirs",
			Title: "Скрытые пути",
			Hint:  "Перебор путей (как gobuster dir): словарь, workers, RPS, скрытие статус-кодов.",
			Mode:  ModeActive,
		},
		{
			ID:    "params",
			Title: "Параметры",
			Hint:  "Перебор query-параметров на базовом URL. Трекинг (utm, fbclid) не считается находкой.",
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
